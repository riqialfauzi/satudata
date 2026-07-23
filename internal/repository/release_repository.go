package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/satudata/backend/internal/domain"
	"gorm.io/gorm"
)

// ReleaseRepository adalah implementasi dari ReleaseRepositoryInterface.
type ReleaseRepository struct {
	db    *gorm.DB
	cache *CacheRepository
}

// NewReleaseRepository membuat instance baru ReleaseRepository.
func NewReleaseRepository(db *gorm.DB, cache *CacheRepository) *ReleaseRepository {
	return &ReleaseRepository{
		db:    db,
		cache: cache,
	}
}

// GetReleases mengambil daftar releases dengan filter dan pagination.
func (r *ReleaseRepository) GetReleases(ctx context.Context, filter ReleaseFilter) ([]domain.Release, int64, error) {
	cacheKey := fmt.Sprintf("releases:list:%+v", filter)

	// Coba ambil dari cache
	if r.cache != nil {
		var cached []domain.Release
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil && cached != nil {
			return cached, 0, nil
		}
	}

	query := r.db.WithContext(ctx).Model(&domain.Release{})

	// Apply filters
	if filter.Type != "" {
		query = query.Where("release_type = ?", filter.Type)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	} else {
		query = query.Where("status != ?", "archived")
	}
	if filter.Year > 0 {
		query = query.Where("year = ?", filter.Year)
	}
	if filter.Search != "" {
		// Try full-text search first, fallback to ILIKE
		search := "%" + filter.Search + "%"
		query = query.Where(
			"search_vector @@ plainto_tsquery('indonesian', ?) OR title ILIKE ? OR description ILIKE ?",
			filter.Search, search, search,
		)
	}

	// Count total
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count releases: %w", err)
	}

	// Pagination defaults
	page := filter.Page
	if page < 1 {
		page = 1
	}
	limit := filter.Limit
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit

	// Sorting
	sortBy := "created_at"
	if filter.SortBy != "" {
		sortBy = filter.SortBy
	}
	sortDir := "DESC"
	if filter.SortDir == "asc" {
		sortDir = "ASC"
	}

	// Exclude soft-deleted
	query = query.Where("deleted_at IS NULL")

	var releases []domain.Release
	if err := query.Order(fmt.Sprintf("%s %s", sortBy, sortDir)).
		Preload("DatasetMetadata").
		Preload("ArticleMetadata").
		Offset(offset).
		Limit(limit).
		Find(&releases).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get releases: %w", err)
	}

	// Simpan ke cache (TTL: 5 menit)
	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, releases, 5*time.Minute)
	}

	return releases, total, nil
}

// GetReleaseByID mengambil release berdasarkan ID.
func (r *ReleaseRepository) GetReleaseByID(ctx context.Context, id string) (*domain.Release, error) {
	cacheKey := fmt.Sprintf("releases:id:%s", id)

	// Coba ambil dari cache
	if r.cache != nil {
		var cached domain.Release
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil && cached.ID.String() != "" {
			return &cached, nil
		}
	}

	var release domain.Release
	err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Preload("DatasetMetadata").
		Preload("ArticleMetadata").
		First(&release, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get release by ID: %w", err)
	}

	// Simpan ke cache (TTL: 10 menit)
	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, release, 10*time.Minute)
	}

	return &release, nil
}

// GetReleaseBySlug mengambil release berdasarkan slug.
func (r *ReleaseRepository) GetReleaseBySlug(ctx context.Context, slug string) (*domain.Release, error) {
	cacheKey := fmt.Sprintf("releases:slug:%s", slug)

	if r.cache != nil {
		var cached domain.Release
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil && cached.ID.String() != "" {
			return &cached, nil
		}
	}

	var release domain.Release
	err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Preload("DatasetMetadata").
		Preload("ArticleMetadata").
		First(&release, "slug = ?", slug).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get release by slug: %w", err)
	}

	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, release, 10*time.Minute)
	}

	return &release, nil
}

// CreateRelease membuat release baru.
func (r *ReleaseRepository) CreateRelease(ctx context.Context, release *domain.Release) error {
	if err := r.db.WithContext(ctx).Create(release).Error; err != nil {
		return fmt.Errorf("failed to create release: %w", err)
	}

	// Invalidate cache
	r.invalidateListCache(ctx)
	return nil
}

// UpdateRelease memperbarui release yang ada.
func (r *ReleaseRepository) UpdateRelease(ctx context.Context, release *domain.Release) error {
	if err := r.db.WithContext(ctx).Save(release).Error; err != nil {
		return fmt.Errorf("failed to update release: %w", err)
	}

	// Invalidate cache
	r.invalidateCache(ctx, release.ID.String(), release.Slug)
	return nil
}

// DeleteRelease melakukan soft delete release.
func (r *ReleaseRepository) DeleteRelease(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).
		Model(&domain.Release{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now())

	if result.Error != nil {
		return fmt.Errorf("failed to delete release: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("release not found")
	}

	// Invalidate cache
	r.invalidateListCache(ctx)
	r.invalidateCacheByID(ctx, id)
	return nil
}

// GetReleaseStats mengembalikan statistik releases (total by type, year).
func (r *ReleaseRepository) GetReleaseStats(ctx context.Context) (map[string]int64, error) {
	cacheKey := "releases:stats"

	if r.cache != nil {
		var cached map[string]int64
		if err := r.cache.Get(ctx, cacheKey, &cached); err == nil && cached != nil {
			return cached, nil
		}
	}

	stats := make(map[string]int64)

	// Stats by type
	type typeStat struct {
		Type  string
		Count int64
	}
	var byType []typeStat
	r.db.WithContext(ctx).Model(&domain.Release{}).
		Select("release_type as type, COUNT(*) as count").
		Where("deleted_at IS NULL").
		Group("release_type").
		Scan(&byType)

	for _, ts := range byType {
		stats[fmt.Sprintf("type_%s", ts.Type)] = ts.Count
	}

	// Stats by year
	var byYear []typeStat
	r.db.WithContext(ctx).Model(&domain.Release{}).
		Select("CAST(year AS VARCHAR) as type, COUNT(*) as count").
		Where("deleted_at IS NULL").
		Group("year").
		Order("year DESC").
		Scan(&byYear)

	for _, ys := range byYear {
		stats[fmt.Sprintf("year_%s", ys.Type)] = ys.Count
	}

	// Total count
	var total int64
	r.db.WithContext(ctx).Model(&domain.Release{}).
		Where("deleted_at IS NULL").
		Count(&total)
	stats["total"] = total

	if r.cache != nil {
		_ = r.cache.Set(ctx, cacheKey, stats, 5*time.Minute)
	}

	return stats, nil
}

// GetRelatedReleases mengambil releases terkait (berdasarkan tags yang sama, exclude ID tertentu).
func (r *ReleaseRepository) GetRelatedReleases(ctx context.Context, releaseID string, limit int) ([]domain.Release, error) {
	// Cari tags dari release ini
	var currentRelease domain.Release
	if err := r.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", releaseID).
		First(&currentRelease).Error; err != nil {
		return nil, nil
	}

	// Cari release lain dengan tags yang sama
	var releases []domain.Release
	query := r.db.WithContext(ctx).
		Where("deleted_at IS NULL AND status = 'published' AND id != ?", releaseID)

	// If release has tags, find by similar tags
	if len(currentRelease.Tags) > 0 {
		// Build a JSON containment query for tags
		for _, tag := range currentRelease.Tags {
			query = query.Or("tags @> ?", fmt.Sprintf(`["%s"]`, tag))
		}
	} else {
		// Fallback: same type
		query = query.Where("release_type = ?", currentRelease.ReleaseType)
	}

	if err := query.
		Preload("DatasetMetadata").
		Preload("ArticleMetadata").
		Limit(limit).
		Find(&releases).Error; err != nil {
		return nil, fmt.Errorf("failed to get related releases: %w", err)
	}

	return releases, nil
}

// invalidateListCache menghapus cache daftar releases.
func (r *ReleaseRepository) invalidateListCache(ctx context.Context) {
	if r.cache == nil {
		return
	}
	// Pattern-based cache invalidation would need Redis SCAN
	// For simplicity, we use a version key approach
	r.cache.Incr(ctx, "releases:cache_version")
}

// invalidateCache menghapus cache untuk release tertentu.
func (r *ReleaseRepository) invalidateCache(ctx context.Context, id, slug string) {
	if r.cache == nil {
		return
	}
	_ = r.cache.Delete(ctx,
		fmt.Sprintf("releases:id:%s", id),
		fmt.Sprintf("releases:slug:%s", slug),
	)
}

// invalidateCacheByID menghapus cache untuk release berdasarkan ID.
func (r *ReleaseRepository) invalidateCacheByID(ctx context.Context, id string) {
	if r.cache == nil {
		return
	}
	_ = r.cache.Delete(ctx, fmt.Sprintf("releases:id:%s", id))
}
