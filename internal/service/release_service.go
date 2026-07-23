package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/satudata/backend/internal/domain"
	"github.com/satudata/backend/internal/repository"
)

// ReleaseService adalah implementasi dari ReleaseServiceInterface.
type ReleaseService struct {
	releaseRepo repository.ReleaseRepositoryInterface
}

// NewReleaseService membuat instance baru ReleaseService.
func NewReleaseService(releaseRepo repository.ReleaseRepositoryInterface) *ReleaseService {
	return &ReleaseService{
		releaseRepo: releaseRepo,
	}
}

// GetReleases mengambil daftar releases dengan filter.
func (s *ReleaseService) GetReleases(ctx context.Context, filter ReleaseFilterRequest) ([]domain.Release, int64, error) {
	// Validasi filter defaults
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 || filter.Limit > 100 {
		filter.Limit = 10
	}
	if filter.SortBy == "" {
		filter.SortBy = "created_at"
	}
	if filter.SortDir == "" {
		filter.SortDir = "desc"
	}

	// Validasi release type
	if filter.Type != "" {
		rt := domain.ReleaseType(filter.Type)
		if !rt.IsValid() {
			return nil, 0, fmt.Errorf("invalid release type: %s", filter.Type)
		}
	}

	repoFilter := repository.ReleaseFilter{
		Type:    filter.Type,
		Status:  filter.Status,
		Year:    filter.Year,
		Search:  filter.Search,
		Page:    filter.Page,
		Limit:   filter.Limit,
		SortBy:  filter.SortBy,
		SortDir: filter.SortDir,
	}

	return s.releaseRepo.GetReleases(ctx, repoFilter)
}

// GetReleaseByID mengambil release berdasarkan ID.
func (s *ReleaseService) GetReleaseByID(ctx context.Context, id string) (*domain.Release, error) {
	if id == "" {
		return nil, fmt.Errorf("release ID is required")
	}

	release, err := s.releaseRepo.GetReleaseByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get release: %w", err)
	}
	if release == nil {
		return nil, fmt.Errorf("release not found")
	}

	return release, nil
}

// GetReleaseBySlug mengambil release berdasarkan slug.
func (s *ReleaseService) GetReleaseBySlug(ctx context.Context, slug string) (*domain.Release, error) {
	if slug == "" {
		return nil, fmt.Errorf("release slug is required")
	}

	release, err := s.releaseRepo.GetReleaseBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("failed to get release: %w", err)
	}
	if release == nil {
		return nil, fmt.Errorf("release not found")
	}

	return release, nil
}

// CreateRelease membuat release baru.
func (s *ReleaseService) CreateRelease(ctx context.Context, req CreateReleaseRequest) (*domain.Release, error) {
	// Validasi
	if req.Title == "" {
		return nil, fmt.Errorf("title is required")
	}
	if req.ReleaseType == "" {
		return nil, fmt.Errorf("release type is required")
	}

	rt := domain.ReleaseType(req.ReleaseType)
	if !rt.IsValid() {
		return nil, fmt.Errorf("invalid release type: %s", req.ReleaseType)
	}

	if req.Year < 2000 || req.Year > 2100 {
		return nil, fmt.Errorf("invalid year: %d", req.Year)
	}

	slug := generateSlug(req.Title)

	release := &domain.Release{
		Title:         req.Title,
		Slug:          slug,
		Description:   req.Description,
		ReleaseType:   rt,
		Status:        domain.ReleaseStatusDraft,
		Year:          req.Year,
		CoverImageURL: req.CoverImageURL,
		Tags:          req.Tags,
	}

	// Set metadata berdasarkan tipe
	if rt == domain.ReleaseTypeDataset {
		release.DatasetMetadata = &domain.DatasetMetadata{
			FileURL:         req.FileURL,
			FileFormat:      req.FileFormat,
			FileSize:        req.FileSize,
			DataSource:      req.DataSource,
			UpdateFrequency: req.UpdateFrequency,
			IsGeospatial:    req.IsGeospatial,
		}
		if req.DataPeriodStart != "" {
			if t, err := time.Parse("2006-01-02", req.DataPeriodStart); err == nil {
				release.DatasetMetadata.DataPeriodStart = &t
			}
		}
		if req.DataPeriodEnd != "" {
			if t, err := time.Parse("2006-01-02", req.DataPeriodEnd); err == nil {
				release.DatasetMetadata.DataPeriodEnd = &t
			}
		}
	} else if rt == domain.ReleaseTypeArticle {
		release.ArticleMetadata = &domain.ArticleMetadata{
			Content:    req.Content,
			Excerpt:    req.Excerpt,
			AuthorName: req.AuthorName,
			Category:   req.Category,
			IsFeatured: req.IsFeatured,
		}
	}

	if err := s.releaseRepo.CreateRelease(ctx, release); err != nil {
		return nil, fmt.Errorf("failed to create release: %w", err)
	}

	return release, nil
}

// UpdateRelease memperbarui release yang ada.
func (s *ReleaseService) UpdateRelease(ctx context.Context, id string, req UpdateReleaseRequest) (*domain.Release, error) {
	existing, err := s.releaseRepo.GetReleaseByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get release: %w", err)
	}
	if existing == nil {
		return nil, fmt.Errorf("release not found")
	}

	if req.Title != "" {
		existing.Title = req.Title
		existing.Slug = generateSlug(req.Title)
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	if req.Status != "" {
		rs := domain.ReleaseStatus(req.Status)
		if !rs.IsValid() {
			return nil, fmt.Errorf("invalid status: %s", req.Status)
		}
		existing.Status = rs
	}
	if req.Year > 0 {
		existing.Year = req.Year
	}
	if req.CoverImageURL != "" {
		existing.CoverImageURL = req.CoverImageURL
	}
	if req.Tags != nil {
		existing.Tags = req.Tags
	}

	if err := s.releaseRepo.UpdateRelease(ctx, existing); err != nil {
		return nil, fmt.Errorf("failed to update release: %w", err)
	}

	return existing, nil
}

// DeleteRelease melakukan soft delete release.
func (s *ReleaseService) DeleteRelease(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("release ID is required")
	}

	return s.releaseRepo.DeleteRelease(ctx, id)
}

// GetRelatedReleases mengembalikan releases terkait.
func (s *ReleaseService) GetRelatedReleases(ctx context.Context, releaseID string, limit int) ([]domain.Release, error) {
	if limit <= 0 || limit > 20 {
		limit = 5
	}

	releases, err := s.releaseRepo.GetRelatedReleases(ctx, releaseID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get related releases: %w", err)
	}

	return releases, nil
}

// GetReleaseStats mengembalikan statistik releases.
func (s *ReleaseService) GetReleaseStats(ctx context.Context) (*ReleaseStatsResponse, error) {
	stats, err := s.releaseRepo.GetReleaseStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get release stats: %w", err)
	}

	response := &ReleaseStatsResponse{
		Total:  stats["total"],
		ByType: make(map[string]int64),
		ByYear: make(map[string]int64),
	}

	for k, v := range stats {
		if strings.HasPrefix(k, "type_") {
			response.ByType[strings.TrimPrefix(k, "type_")] = v
		} else if strings.HasPrefix(k, "year_") {
			response.ByYear[strings.TrimPrefix(k, "year_")] = v
		}
	}

	return response, nil
}

// generateSlug membuat slug dari title.
func generateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, ".", "-")

	// Remove special characters
	var result strings.Builder
	for _, ch := range slug {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '-' {
			result.WriteRune(ch)
		}
	}

	slug = result.String()
	slug = strings.Trim(slug, "-")

	// Remove consecutive dashes
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	// Add unique suffix
	if len(slug) > 100 {
		slug = slug[:100]
	}
	slug = fmt.Sprintf("%s-%s", slug, uuid.New().String()[:8])

	return slug
}
