package repository

import (
	"context"

	"github.com/satudata/backend/internal/domain"
)

// ReleaseRepositoryInterface mendefinisikan operasi untuk repository release.
type ReleaseRepositoryInterface interface {
	GetReleases(ctx context.Context, filter ReleaseFilter) ([]domain.Release, int64, error)
	GetReleaseByID(ctx context.Context, id string) (*domain.Release, error)
	GetReleaseBySlug(ctx context.Context, slug string) (*domain.Release, error)
	CreateRelease(ctx context.Context, release *domain.Release) error
	UpdateRelease(ctx context.Context, release *domain.Release) error
	DeleteRelease(ctx context.Context, id string) error
	GetReleaseStats(ctx context.Context) (map[string]int64, error)
	GetRelatedReleases(ctx context.Context, releaseID string, limit int) ([]domain.Release, error)
}

// StandardRepositoryInterface mendefinisikan operasi untuk repository standard.
type StandardRepositoryInterface interface {
	GetStandards(ctx context.Context) ([]domain.Standard, error)
	GetStandardByYear(ctx context.Context, year int) (*domain.Standard, error)
	CreateStandard(ctx context.Context, standard *domain.Standard) error
	UpdateStandard(ctx context.Context, standard *domain.Standard) error
}

// UserRepositoryInterface mendefinisikan operasi untuk repository user.
type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
}

// AuditLogRepositoryInterface mendefinisikan operasi untuk repository audit log.
type AuditLogRepositoryInterface interface {
	Create(ctx context.Context, log *domain.AuditLog) error
	GetAuditLogs(ctx context.Context, filter AuditLogFilter) ([]domain.AuditLog, int64, error)
}

// ReleaseFilter adalah filter untuk query releases.
type ReleaseFilter struct {
	Type    string `form:"type"`
	Status  string `form:"status"`
	Year    int    `form:"year"`
	Search  string `form:"search"`
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
	SortBy  string `form:"sort_by"`
	SortDir string `form:"sort_dir"` // "asc" or "desc"
}

// AuditLogFilter adalah filter untuk query audit logs.
type AuditLogFilter struct {
	UserID   string `form:"user_id"`
	Action   string `form:"action"`
	Resource string `form:"resource"`
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
}

// OrganizationRepositoryInterface mendefinisikan operasi untuk repository organization.
type OrganizationRepositoryInterface interface {
	GetOrganizations(ctx context.Context) ([]domain.Organization, error)
	GetOrganizationBySlug(ctx context.Context, slug string) (*domain.Organization, error)
	CreateOrganization(ctx context.Context, org *domain.Organization) error
	UpdateOrganization(ctx context.Context, org *domain.Organization) error
	DeleteOrganization(ctx context.Context, id string) error
}

// GroupRepositoryInterface mendefinisikan operasi untuk repository group.
type GroupRepositoryInterface interface {
	GetGroups(ctx context.Context) ([]domain.Group, error)
	GetGroupBySlug(ctx context.Context, slug string) (*domain.Group, error)
	CreateGroup(ctx context.Context, group *domain.Group) error
	UpdateGroup(ctx context.Context, group *domain.Group) error
	DeleteGroup(ctx context.Context, id string) error
}
