package service

import (
	"context"
	"fmt"

	"github.com/satudata/backend/internal/domain"
	"github.com/satudata/backend/internal/repository"
)

// GroupService adalah implementasi dari GroupServiceInterface.
type GroupService struct {
	groupRepo repository.GroupRepositoryInterface
}

// NewGroupService membuat instance baru GroupService.
func NewGroupService(groupRepo repository.GroupRepositoryInterface) *GroupService {
	return &GroupService{
		groupRepo: groupRepo,
	}
}

// GetGroups mengambil daftar semua grup/kategori.
func (s *GroupService) GetGroups(ctx context.Context) ([]domain.Group, error) {
	return s.groupRepo.GetGroups(ctx)
}

// GetGroupBySlug mengambil grup berdasarkan slug.
func (s *GroupService) GetGroupBySlug(ctx context.Context, slug string) (*domain.Group, error) {
	if slug == "" {
		return nil, fmt.Errorf("group slug is required")
	}

	group, err := s.groupRepo.GetGroupBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("failed to get group: %w", err)
	}
	if group == nil {
		return nil, fmt.Errorf("group not found")
	}

	return group, nil
}

// CreateGroup membuat grup baru.
func (s *GroupService) CreateGroup(ctx context.Context, req CreateGroupRequest) (*domain.Group, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("group name is required")
	}

	group := &domain.Group{
		Name:        req.Name,
		Slug:        generateSlug(req.Name),
		Description: req.Description,
		ImageURL:    req.ImageURL,
	}

	if err := s.groupRepo.CreateGroup(ctx, group); err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	return group, nil
}

// UpdateGroup memperbarui grup.
func (s *GroupService) UpdateGroup(ctx context.Context, id string, req UpdateGroupRequest) (*domain.Group, error) {
	groups, err := s.groupRepo.GetGroups(ctx)
	if err != nil {
		return nil, err
	}

	var existing *domain.Group
	for i := range groups {
		if groups[i].ID.String() == id {
			existing = &groups[i]
			break
		}
	}
	if existing == nil {
		return nil, fmt.Errorf("group not found")
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

	if err := s.groupRepo.UpdateGroup(ctx, existing); err != nil {
		return nil, fmt.Errorf("failed to update group: %w", err)
	}

	return existing, nil
}

// DeleteGroup menghapus grup.
func (s *GroupService) DeleteGroup(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("group ID is required")
	}
	return s.groupRepo.DeleteGroup(ctx, id)
}
