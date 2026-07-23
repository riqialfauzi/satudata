package dto

// GetReleasesRequest adalah query params untuk daftar releases.
type GetReleasesRequest struct {
	Type    string `form:"type"`
	Status  string `form:"status"`
	Year    int    `form:"year"`
	Search  string `form:"search"`
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
	SortBy  string `form:"sort_by"`
	SortDir string `form:"sort_dir"`
}

// CreateReleaseRequest adalah request body untuk membuat release baru.
type CreateReleaseRequest struct {
	Title         string   `json:"title" binding:"required"`
	Description   string   `json:"description"`
	ReleaseType   string   `json:"release_type" binding:"required"`
	Year          int      `json:"year" binding:"required"`
	CoverImageURL string   `json:"cover_image_url"`
	Tags          []string `json:"tags"`

	// Dataset-specific
	FileURL         string `json:"file_url"`
	FileFormat      string `json:"file_format"`
	FileSize        int64  `json:"file_size"`
	DataSource      string `json:"data_source"`
	DataPeriodStart string `json:"data_period_start"`
	DataPeriodEnd   string `json:"data_period_end"`
	UpdateFrequency string `json:"update_frequency"`
	IsGeospatial    bool   `json:"is_geospatial"`

	// Article-specific
	Content    string `json:"content"`
	Excerpt    string `json:"excerpt"`
	AuthorName string `json:"author_name"`
	Category   string `json:"category"`
	IsFeatured bool   `json:"is_featured"`
}

// UpdateReleaseRequest adalah request body untuk memperbarui release.
type UpdateReleaseRequest struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Status        string   `json:"status"`
	Year          int      `json:"year"`
	CoverImageURL string   `json:"cover_image_url"`
	Tags          []string `json:"tags"`
}

// CreateStandardRequest adalah request body untuk membuat standard baru.
type CreateStandardRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Year        int    `json:"year" binding:"required"`
	FileURL     string `json:"file_url"`
	FileSize    int64  `json:"file_size"`
	Version     string `json:"version"`
	IsCurrent   bool   `json:"is_current"`
}

// UpdateStandardRequest adalah request body untuk memperbarui standard.
type UpdateStandardRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	FileURL     string `json:"file_url"`
	FileSize    int64  `json:"file_size"`
	Version     string `json:"version"`
	IsCurrent   bool   `json:"is_current"`
}

// LoginRequest adalah request body untuk login.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest adalah request body untuk registrasi.
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
}

// RefreshTokenRequest adalah request body untuk refresh token.
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// CreateOrganizationRequest adalah request body untuk membuat organisasi baru.
type CreateOrganizationRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Website     string `json:"website"`
}

// UpdateOrganizationRequest adalah request body untuk memperbarui organisasi.
type UpdateOrganizationRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url"`
	Website     *string `json:"website"`
}

// CreateGroupRequest adalah request body untuk membuat grup baru.
type CreateGroupRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// UpdateGroupRequest adalah request body untuk memperbarui grup.
type UpdateGroupRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url"`
}
