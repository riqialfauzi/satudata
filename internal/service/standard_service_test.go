package service

import (
	"context"
	"testing"

	"github.com/satudata/backend/internal/domain"
)

// mockStandardRepo adalah mock untuk StandardRepositoryInterface.
type mockStandardRepo struct {
	standards []domain.Standard
}

func newMockStandardRepo() *mockStandardRepo {
	return &mockStandardRepo{}
}

func (m *mockStandardRepo) GetStandards(ctx context.Context) ([]domain.Standard, error) {
	return m.standards, nil
}

func (m *mockStandardRepo) GetStandardByYear(ctx context.Context, year int) (*domain.Standard, error) {
	for _, s := range m.standards {
		if s.Year == year {
			return &s, nil
		}
	}
	return nil, nil
}

func (m *mockStandardRepo) CreateStandard(ctx context.Context, standard *domain.Standard) error {
	m.standards = append(m.standards, *standard)
	return nil
}

func (m *mockStandardRepo) UpdateStandard(ctx context.Context, standard *domain.Standard) error {
	for i, s := range m.standards {
		if s.ID == standard.ID {
			m.standards[i] = *standard
			return nil
		}
	}
	return nil
}

func TestCreateStandard(t *testing.T) {
	repo := newMockStandardRepo()
	svc := NewStandardService(repo)

	req := CreateStandardRequest{
		Title:       "Data Standard 2025",
		Description: "Standard data untuk tahun 2025",
		Year:        2025,
		Version:     "1.0",
		IsCurrent:   true,
	}

	std, err := svc.CreateStandard(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if std.Title != "Data Standard 2025" {
		t.Errorf("expected title 'Data Standard 2025', got: %s", std.Title)
	}

	if std.Status != domain.StandardStatusActive {
		t.Errorf("expected status 'active', got: %s", std.Status)
	}

	if std.Year != 2025 {
		t.Errorf("expected year 2025, got: %d", std.Year)
	}
}

func TestCreateStandard_InvalidYear(t *testing.T) {
	repo := newMockStandardRepo()
	svc := NewStandardService(repo)

	req := CreateStandardRequest{
		Title: "Invalid Standard",
		Year:  1999,
	}

	_, err := svc.CreateStandard(context.Background(), req)
	if err == nil {
		t.Fatal("expected error for invalid year, got nil")
	}
}

func TestGetActiveStandards(t *testing.T) {
	repo := newMockStandardRepo()
	svc := NewStandardService(repo)

	// Add two standards - one active, one archived
	_, _ = svc.CreateStandard(context.Background(), CreateStandardRequest{
		Title: "Active Standard",
		Year:  2025,
	})

	archivedStd := &domain.Standard{
		Title:  "Archived Standard",
		Year:   2024,
		Status: domain.StandardStatusArchived,
	}
	// Set archived manually - the service defaults to active
	archivedStd.Status = domain.StandardStatusArchived
	_ = repo.CreateStandard(context.Background(), archivedStd)

	active, err := svc.GetActiveStandards(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Should only return active ones
	for _, s := range active {
		if s.Status != domain.StandardStatusActive {
			t.Errorf("expected all standards to be active, got: %s", s.Status)
		}
	}
}
