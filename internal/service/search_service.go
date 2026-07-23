package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/satudata/backend/internal/domain"
	"gorm.io/gorm"
)

// SearchService adalah service untuk pencarian dan autocomplete.
type SearchService struct {
	db *gorm.DB
}

// NewSearchService membuat instance baru SearchService.
func NewSearchService(db *gorm.DB) *SearchService {
	return &SearchService{
		db: db,
	}
}

// SearchSuggest mengembalikan saran pencarian berdasarkan query.
func (s *SearchService) SearchSuggest(ctx context.Context, query string) (*SearchSuggestResponse, error) {
	if query == "" || len(strings.TrimSpace(query)) < 2 {
		return &SearchSuggestResponse{
			Suggestions: []string{},
			Datasets:    []domain.Release{},
		}, nil
	}

	q := "%" + query + "%"

	// Cari title suggestions dari releases (max 5)
	var titles []string
	if err := s.db.WithContext(ctx).
		Model(&domain.Release{}).
		Where("deleted_at IS NULL AND status = 'published'").
		Where("title ILIKE ?", q).
		Limit(5).
		Pluck("title", &titles).Error; err != nil {
		return nil, fmt.Errorf("failed to search suggest: %w", err)
	}

	// Cari releases yang cocok (max 5)
	var releases []domain.Release
	if err := s.db.WithContext(ctx).
		Where("deleted_at IS NULL AND status = 'published'").
		Where("search_vector @@ plainto_tsquery('indonesian', ?) OR title ILIKE ? OR description ILIKE ?", query, q, q).
		Preload("DatasetMetadata").
		Preload("ArticleMetadata").
		Limit(5).
		Find(&releases).Error; err != nil {
		return nil, fmt.Errorf("failed to search releases: %w", err)
	}

	return &SearchSuggestResponse{
		Suggestions: titles,
		Datasets:    releases,
	}, nil
}

// SearchSuggestResponse adalah response DTO untuk search suggest.
type SearchSuggestResponse struct {
	Suggestions []string         `json:"suggestions"`
	Datasets    []domain.Release `json:"datasets"`
}
