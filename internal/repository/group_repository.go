package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/satudata/backend/internal/domain"
	"gorm.io/gorm"
)

// GroupRepository adalah implementasi dari GroupRepositoryInterface.
type GroupRepository struct {
	db    *gorm.DB
	cache *CacheRepository
}

// NewGroupRepository membuat instance baru GroupRepository.
func NewGroupRepository(db *gorm.DB, cache *CacheRepository) *GroupRepository {
	return &GroupRepository{
		db:    db,
		cache: cache,
	}
}

// GetGroups mengambil daftar semua grup/kategori.
func (r *GroupRepository) GetGroups(ctx context.Context) ([]domain.Group, error) {
	cacheKey := "groups:list"

	if r.cache != nil {
		var cached []domain.Group
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil && cached != nil {
			return cached, nil
		}
	}

	var groups []domain.Group
	if err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Order("name ASC").
		Find(&groups).Error; err != nil {
		return nil, fmt.Errorf("failed to get groups: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, groups, 30*time.Minute)
	}

	return groups, nil
}

// GetGroupBySlug mengambil grup berdasarkan slug.
func (r *GroupRepository) GetGroupBySlug(ctx context.Context, slug string) (*domain.Group, error) {
	cacheKey := fmt.Sprintf("group:slug:%s", slug)

	if r.cache != nil {
		var cached domain.Group
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil {
			return &cached, nil
		}
	}

	var group domain.Group
	if err := r.db.WithContext(ctx).
		Where("slug = ? AND deleted_at IS NULL", slug).
		First(&group).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get group: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, group, 30*time.Minute)
	}

	return &group, nil
}

// CreateGroup membuat grup baru.
func (r *GroupRepository) CreateGroup(ctx context.Context, group *domain.Group) error {
	if err := r.db.WithContext(ctx).Create(group).Error; err != nil {
		return fmt.Errorf("failed to create group: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Delete(ctx, "groups:list")
	}

	return nil
}

// UpdateGroup memperbarui grup.
func (r *GroupRepository) UpdateGroup(ctx context.Context, group *domain.Group) error {
	if err := r.db.WithContext(ctx).Save(group).Error; err != nil {
		return fmt.Errorf("failed to update group: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Delete(ctx, "groups:list")
		_ = r.cache.Delete(ctx, fmt.Sprintf("group:slug:%s", group.Slug))
	}

	return nil
}

// DeleteGroup menghapus grup (soft delete).
func (r *GroupRepository) DeleteGroup(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.Group{}).Error; err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Delete(ctx, "groups:list")
	}

	return nil
}
