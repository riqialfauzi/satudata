package service

import (
	"context"
	"fmt"

	"github.com/satudata/backend/internal/domain"
	"github.com/satudata/backend/internal/repository"
)

// OrganizationService adalah implementasi dari OrganizationServiceInterface.
type OrganizationService struct {
	orgRepo repository.OrganizationRepositoryInterface
}

// NewOrganizationService membuat instance baru OrganizationService.
func NewOrganizationService(orgRepo repository.OrganizationRepositoryInterface) *OrganizationService {
	return &OrganizationService{
		orgRepo: orgRepo,
	}
}

// GetOrganizations mengambil daftar semua organisasi.
func (s *OrganizationService) GetOrganizations(ctx context.Context) ([]domain.Organization, error) {
	return s.orgRepo.GetOrganizations(ctx)
}

// GetOrganizationBySlug mengambil organisasi berdasarkan slug.
func (s *OrganizationService) GetOrganizationBySlug(ctx context.Context, slug string) (*domain.Organization, error) {
	if slug == "" {
		return nil, fmt.Errorf("organization slug is required")
	}

	org, err := s.orgRepo.GetOrganizationBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization: %w", err)
	}
	if org == nil {
		return nil, fmt.Errorf("organization not found")
	}

	return org, nil
}

// CreateOrganization membuat organisasi baru.
func (s *OrganizationService) CreateOrganization(ctx context.Context, req CreateOrganizationRequest) (*domain.Organization, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("organization name is required")
	}

	slug := generateSlug(req.Name)

	org := &domain.Organization{
		Name:        req.Name,
		Slug:        slug,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		Website:     req.Website,
	}

	if err := s.orgRepo.CreateOrganization(ctx, org); err != nil {
		return nil, fmt.Errorf("failed to create organization: %w", err)
	}

	return org, nil
}

// UpdateOrganization memperbarui organisasi.
func (s *OrganizationService) UpdateOrganization(ctx context.Context, id string, req UpdateOrganizationRequest) (*domain.Organization, error) {
	// Since we don't have GetByID on the interface, we'll get by listing first
	orgs, err := s.orgRepo.GetOrganizations(ctx)
	if err != nil {
		return nil, err
	}

	var existing *domain.Organization
	for i := range orgs {
		if orgs[i].ID.String() == id {
			existing = &orgs[i]
			break
		}
	}
	if existing == nil {
		return nil, fmt.Errorf("organization not found")
	}

	if req.Name != "" {
		existing.Name = req.Name
		existing.Slug = generateSlug(req.Name)
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}
	if req.ImageURL != nil {
		existing.ImageURL = *req.ImageURL
	}
	if req.Website != nil {
		existing.Website = *req.Website
	}

	if err := s.orgRepo.UpdateOrganization(ctx, existing); err != nil {
		return nil, fmt.Errorf("failed to update organization: %w", err)
	}

	return existing, nil
}

// DeleteOrganization menghapus organisasi.
func (s *OrganizationService) DeleteOrganization(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("organization ID is required")
	}
	return s.orgRepo.DeleteOrganization(ctx, id)
}
