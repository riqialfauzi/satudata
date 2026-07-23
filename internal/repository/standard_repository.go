package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/satudata/backend/internal/domain"
	"gorm.io/gorm"
)

// StandardRepository adalah implementasi dari StandardRepositoryInterface.
type StandardRepository struct {
	db    *gorm.DB
	cache *CacheRepository
}

// NewStandardRepository membuat instance baru StandardRepository.
func NewStandardRepository(db *gorm.DB, cache *CacheRepository) *StandardRepository {
	return &StandardRepository{
		db:    db,
		cache: cache,
	}
}

// GetStandards mengambil semua standard (kecuali yang soft-deleted).
func (r *StandardRepository) GetStandards(ctx context.Context) ([]domain.Standard, error) {
	cacheKey := "standards:all"

	if r.cache != nil {
		var cached []domain.Standard
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil && cached != nil {
			return cached, nil
		}
	}

	var standards []domain.Standard
	err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Order("year DESC, created_at DESC").
		Find(&standards).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get standards: %w", err)
	}

	// Cache TTL: 1 jam
	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, standards, 1*time.Hour)
	}

	return standards, nil
}

// GetStandardByYear mengambil standard berdasarkan tahun.
func (r *StandardRepository) GetStandardByYear(ctx context.Context, year int) (*domain.Standard, error) {
	cacheKey := fmt.Sprintf("standards:year:%d", year)

	if r.cache != nil {
		var cached domain.Standard
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil && cached.ID.String() != "" {
			return &cached, nil
		}
	}

	var standard domain.Standard
	err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Where("year = ?", year).
		First(&standard).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get standard by year: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, standard, 1*time.Hour)
	}

	return &standard, nil
}

// CreateStandard membuat standard baru.
func (r *StandardRepository) CreateStandard(ctx context.Context, standard *domain.Standard) error {
	if err := r.db.WithContext(ctx).Create(standard).Error; err != nil {
		return fmt.Errorf("failed to create standard: %w", err)
	}

	// Invalidate cache
	if r.cache != nil {
		_ = r.cache.Delete(ctx, "standards:all")
	}

	return nil
}

// UpdateStandard memperbarui standard yang ada.
func (r *StandardRepository) UpdateStandard(ctx context.Context, standard *domain.Standard) error {
	if err := r.db.WithContext(ctx).Save(standard).Error; err != nil {
		return fmt.Errorf("failed to update standard: %w", err)
	}

	// Invalidate cache
	if r.cache != nil {
		_ = r.cache.Delete(ctx,
			"standards:all",
			fmt.Sprintf("standards:year:%d", standard.Year),
		)
	}

	return nil
}
