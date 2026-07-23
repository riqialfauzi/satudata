package service

import (
	"context"

	"github.com/satudata/backend/internal/domain"
)

// ReleaseServiceInterface mendefinisikan operasi untuk service release.
type ReleaseServiceInterface interface {
	GetReleases(ctx context.Context, filter ReleaseFilterRequest) ([]domain.Release, int64, error)
	GetReleaseByID(ctx context.Context, id string) (*domain.Release, error)
	GetReleaseBySlug(ctx context.Context, slug string) (*domain.Release, error)
	CreateRelease(ctx context.Context, req CreateReleaseRequest) (*domain.Release, error)
	UpdateRelease(ctx context.Context, id string, req UpdateReleaseRequest) (*domain.Release, error)
	DeleteRelease(ctx context.Context, id string) error
	GetReleaseStats(ctx context.Context) (*ReleaseStatsResponse, error)
	GetRelatedReleases(ctx context.Context, releaseID string, limit int) ([]domain.Release, error)
}

// StandardServiceInterface mendefinisikan operasi untuk service standard.
type StandardServiceInterface interface {
	GetStandards(ctx context.Context) ([]domain.Standard, error)
	GetActiveStandards(ctx context.Context) ([]domain.Standard, error)
	CreateStandard(ctx context.Context, req CreateStandardRequest) (*domain.Standard, error)
	UpdateStandard(ctx context.Context, id string, req UpdateStandardRequest) (*domain.Standard, error)
}

// AuthServiceInterface mendefinisikan operasi untuk service autentikasi.
type AuthServiceInterface interface {
	Login(ctx context.Context, req LoginRequest) (*TokenResponse, error)
	Register(ctx context.Context, req RegisterRequest) (*domain.User, error)
	ValidateToken(ctx context.Context, tokenString string) (*JWTClaims, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error)
	Logout(ctx context.Context, token string) error
}

// StorageServiceInterface mendefinisikan operasi untuk service storage/file.
type StorageServiceInterface interface {
	UploadDataset(ctx context.Context, file UploadFileRequest) (string, error)
	UploadArticleImage(ctx context.Context, file UploadFileRequest) (string, error)
	UploadStandardDoc(ctx context.Context, file UploadFileRequest) (string, error)
	DeleteFile(ctx context.Context, url string) error
	GeneratePresignedURL(ctx context.Context, key string, expiry int32) (string, error)
}

// OrganizationServiceInterface mendefinisikan operasi untuk service organization.
type OrganizationServiceInterface interface {
	GetOrganizations(ctx context.Context) ([]domain.Organization, error)
	GetOrganizationBySlug(ctx context.Context, slug string) (*domain.Organization, error)
	CreateOrganization(ctx context.Context, req CreateOrganizationRequest) (*domain.Organization, error)
	UpdateOrganization(ctx context.Context, id string, req UpdateOrganizationRequest) (*domain.Organization, error)
	DeleteOrganization(ctx context.Context, id string) error
}

// GroupServiceInterface mendefinisikan operasi untuk service group.
type GroupServiceInterface interface {
	GetGroups(ctx context.Context) ([]domain.Group, error)
	GetGroupBySlug(ctx context.Context, slug string) (*domain.Group, error)
	CreateGroup(ctx context.Context, req CreateGroupRequest) (*domain.Group, error)
	UpdateGroup(ctx context.Context, id string, req UpdateGroupRequest) (*domain.Group, error)
	DeleteGroup(ctx context.Context, id string) error
}
