package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/satudata/backend/internal/domain"
	"gorm.io/gorm"
)

// OrganizationRepository adalah implementasi dari OrganizationRepositoryInterface.
type OrganizationRepository struct {
	db    *gorm.DB
	cache *CacheRepository
}

// NewOrganizationRepository membuat instance baru OrganizationRepository.
func NewOrganizationRepository(db *gorm.DB, cache *CacheRepository) *OrganizationRepository {
	return &OrganizationRepository{
		db:    db,
		cache: cache,
	}
}

// GetOrganizations mengambil daftar semua organisasi.
func (r *OrganizationRepository) GetOrganizations(ctx context.Context) ([]domain.Organization, error) {
	cacheKey := "organizations:list"

	if r.cache != nil {
		var cached []domain.Organization
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil && cached != nil {
			return cached, nil
		}
	}

	var organizations []domain.Organization
	if err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Order("name ASC").
		Find(&organizations).Error; err != nil {
		return nil, fmt.Errorf("failed to get organizations: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, organizations, 30*time.Minute)
	}

	return organizations, nil
}

// GetOrganizationBySlug mengambil organisasi berdasarkan slug.
func (r *OrganizationRepository) GetOrganizationBySlug(ctx context.Context, slug string) (*domain.Organization, error) {
	cacheKey := fmt.Sprintf("organization:slug:%s", slug)

	if r.cache != nil {
		var cached domain.Organization
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil {
			return &cached, nil
		}
	}

	var org domain.Organization
	if err := r.db.WithContext(ctx).
		Where("slug = ? AND deleted_at IS NULL", slug).
		First(&org).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get organization: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, org, 30*time.Minute)
	}

	return &org, nil
}

// CreateOrganization membuat organisasi baru.
func (r *OrganizationRepository) CreateOrganization(ctx context.Context, org *domain.Organization) error {
	if err := r.db.WithContext(ctx).Create(org).Error; err != nil {
		return fmt.Errorf("failed to create organization: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Delete(ctx, "organizations:list")
	}

	return nil
}

// UpdateOrganization memperbarui organisasi.
func (r *OrganizationRepository) UpdateOrganization(ctx context.Context, org *domain.Organization) error {
	if err := r.db.WithContext(ctx).Save(org).Error; err != nil {
		return fmt.Errorf("failed to update organization: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Delete(ctx, "organizations:list")
		_ = r.cache.Delete(ctx, fmt.Sprintf("organization:slug:%s", org.Slug))
	}

	return nil
}

// DeleteOrganization menghapus organisasi (soft delete).
func (r *OrganizationRepository) DeleteOrganization(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.Organization{}).Error; err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Delete(ctx, "organizations:list")
	}

	return nil
}

// GetOrganizationDatasetCount mengembalikan jumlah dataset per organisasi.
func (r *OrganizationRepository) GetOrganizationDatasetCount(ctx context.Context, orgID string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.Release{}).
		Where("created_by IN (SELECT id FROM users WHERE ...) OR ...", orgID).
		Count(&count).Error; err != nil {
		return 0, nil
	}
	return count, nil
}
