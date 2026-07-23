package service

import (
	"context"
	"fmt"

	"github.com/satudata/backend/internal/domain"
	"github.com/satudata/backend/internal/repository"
)

// StandardService adalah implementasi dari StandardServiceInterface.
type StandardService struct {
	standardRepo repository.StandardRepositoryInterface
}

// NewStandardService membuat instance baru StandardService.
func NewStandardService(standardRepo repository.StandardRepositoryInterface) *StandardService {
	return &StandardService{
		standardRepo: standardRepo,
	}
}

// GetStandards mengambil semua standards.
func (s *StandardService) GetStandards(ctx context.Context) ([]domain.Standard, error) {
	return s.standardRepo.GetStandards(ctx)
}

// GetActiveStandards mengambil standards yang statusnya active.
func (s *StandardService) GetActiveStandards(ctx context.Context) ([]domain.Standard, error) {
	standards, err := s.standardRepo.GetStandards(ctx)
	if err != nil {
		return nil, err
	}

	var active []domain.Standard
	for _, std := range standards {
		if std.Status == domain.StandardStatusActive {
			active = append(active, std)
		}
	}

	return active, nil
}

// CreateStandard membuat standard baru.
func (s *StandardService) CreateStandard(ctx context.Context, req CreateStandardRequest) (*domain.Standard, error) {
	if req.Title == "" {
		return nil, fmt.Errorf("title is required")
	}
	if req.Year < 2000 || req.Year > 2100 {
		return nil, fmt.Errorf("invalid year: %d", req.Year)
	}

	standard := &domain.Standard{
		Title:       req.Title,
		Description: req.Description,
		Year:        req.Year,
		FileURL:     req.FileURL,
		FileSize:    req.FileSize,
		Status:      domain.StandardStatusActive,
		Version:     "1.0",
		IsCurrent:   req.IsCurrent,
	}

	if req.Version != "" {
		standard.Version = req.Version
	}

	if err := s.standardRepo.CreateStandard(ctx, standard); err != nil {
		return nil, fmt.Errorf("failed to create standard: %w", err)
	}

	return standard, nil
}

// UpdateStandard memperbarui standard yang ada.
func (s *StandardService) UpdateStandard(ctx context.Context, id string, req UpdateStandardRequest) (*domain.Standard, error) {
	// Get standard by year (as ID for simplicity)
	standards, err := s.standardRepo.GetStandards(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get standards: %w", err)
	}

	var existing *domain.Standard
	for i, std := range standards {
		if std.ID.String() == id {
			existing = &standards[i]
			break
		}
	}

	if existing == nil {
		return nil, fmt.Errorf("standard not found")
	}

	if req.Title != "" {
		existing.Title = req.Title
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	if req.Status != "" {
		ss := domain.StandardStatus(req.Status)
		if !ss.IsValid() {
			return nil, fmt.Errorf("invalid status: %s", req.Status)
		}
		existing.Status = ss
	}
	if req.FileURL != "" {
		existing.FileURL = req.FileURL
	}
	if req.FileSize > 0 {
		existing.FileSize = req.FileSize
	}
	if req.Version != "" {
		existing.Version = req.Version
	}
	existing.IsCurrent = req.IsCurrent

	if err := s.standardRepo.UpdateStandard(ctx, existing); err != nil {
		return nil, fmt.Errorf("failed to update standard: %w", err)
	}

	return existing, nil
}
