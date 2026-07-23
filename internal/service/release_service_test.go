package service

import (
	"context"
	"testing"

	"github.com/satudata/backend/internal/domain"
	"github.com/satudata/backend/internal/repository"
	"github.com/satudata/backend/pkg/cache"
)

// mockReleaseRepo adalah mock untuk ReleaseRepositoryInterface.
type mockReleaseRepo struct {
	releases map[string]*domain.Release
	stats    map[string]int64
}

func newMockReleaseRepo() *mockReleaseRepo {
	return &mockReleaseRepo{
		releases: make(map[string]*domain.Release),
		stats:    make(map[string]int64),
	}
}

func (m *mockReleaseRepo) GetReleases(ctx context.Context, filter repository.ReleaseFilter) ([]domain.Release, int64, error) {
	var result []domain.Release
	for _, r := range m.releases {
		result = append(result, *r)
	}
	return result, int64(len(result)), nil
}

func (m *mockReleaseRepo) GetReleaseByID(ctx context.Context, id string) (*domain.Release, error) {
	if r, ok := m.releases[id]; ok {
		return r, nil
	}
	return nil, nil
}

func (m *mockReleaseRepo) GetReleaseBySlug(ctx context.Context, slug string) (*domain.Release, error) {
	for _, r := range m.releases {
		if r.Slug == slug {
			return r, nil
		}
	}
	return nil, nil
}

func (m *mockReleaseRepo) CreateRelease(ctx context.Context, release *domain.Release) error {
	m.releases[release.ID.String()] = release
	return nil
}

func (m *mockReleaseRepo) UpdateRelease(ctx context.Context, release *domain.Release) error {
	m.releases[release.ID.String()] = release
	return nil
}

func (m *mockReleaseRepo) DeleteRelease(ctx context.Context, id string) error {
	delete(m.releases, id)
	return nil
}

func (m *mockReleaseRepo) GetReleaseStats(ctx context.Context) (map[string]int64, error) {
	return m.stats, nil
}

func TestCreateRelease(t *testing.T) {
	// Pastikan cache tidak nil (beri dummy)
	_ = &cache.RedisCache{}

	repo := newMockReleaseRepo()
	svc := NewReleaseService(repo)

	req := CreateReleaseRequest{
		Title:       "Test Release",
		ReleaseType: "dataset",
		Year:        2025,
		Description: "A test release",
	}

	release, err := svc.CreateRelease(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if release.Title != "Test Release" {
		t.Errorf("expected title 'Test Release', got: %s", release.Title)
	}

	if release.ReleaseType != domain.ReleaseTypeDataset {
		t.Errorf("expected type 'dataset', got: %s", release.ReleaseType)
	}

	if release.Slug == "" {
		t.Error("expected slug to be generated, got empty")
	}

	if release.Status != domain.ReleaseStatusDraft {
		t.Errorf("expected status 'draft', got: %s", release.Status)
	}
}

func TestCreateRelease_InvalidType(t *testing.T) {
	repo := newMockReleaseRepo()
	svc := NewReleaseService(repo)

	req := CreateReleaseRequest{
		Title:       "Invalid Release",
		ReleaseType: "invalid-type",
		Year:        2025,
	}

	_, err := svc.CreateRelease(context.Background(), req)
	if err == nil {
		t.Fatal("expected error for invalid release type, got nil")
	}
}

func TestCreateRelease_MissingTitle(t *testing.T) {
	repo := newMockReleaseRepo()
	svc := NewReleaseService(repo)

	req := CreateReleaseRequest{
		ReleaseType: "dataset",
		Year:        2025,
	}

	_, err := svc.CreateRelease(context.Background(), req)
	if err == nil {
		t.Fatal("expected error for missing title, got nil")
	}
}

func TestGetReleases(t *testing.T) {
	repo := newMockReleaseRepo()
	svc := NewReleaseService(repo)

	// Add a release first
	req := CreateReleaseRequest{
		Title:       "Test Release",
		ReleaseType: "dataset",
		Year:        2025,
	}
	_, err := svc.CreateRelease(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to create release: %v", err)
	}

	filter := ReleaseFilterRequest{
		Page:  1,
		Limit: 10,
	}

	releases, total, err := svc.GetReleases(context.Background(), filter)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if total != 1 {
		t.Errorf("expected total 1, got: %d", total)
	}

	if len(releases) != 1 {
		t.Errorf("expected 1 release, got: %d", len(releases))
	}
}

func TestDeleteRelease(t *testing.T) {
	repo := newMockReleaseRepo()
	svc := NewReleaseService(repo)

	// Add a release first
	req := CreateReleaseRequest{
		Title:       "To Delete",
		ReleaseType: "article",
		Year:        2025,
	}
	release, err := svc.CreateRelease(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to create release: %v", err)
	}

	// Delete it
	err = svc.DeleteRelease(context.Background(), release.ID.String())
	if err != nil {
		t.Fatalf("expected no error on delete, got: %v", err)
	}

	// Verify it's gone
	_, err = svc.GetReleaseByID(context.Background(), release.ID.String())
	if err == nil {
		t.Fatal("expected error when getting deleted release, got nil")
	}
}
