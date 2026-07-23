package service

// ReleaseFilterRequest adalah request DTO untuk filter daftar releases.
type ReleaseFilterRequest struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Year    int    `json:"year"`
	Search  string `json:"search"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	SortBy  string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
}

// CreateReleaseRequest adalah request DTO untuk membuat release baru.
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

// UpdateReleaseRequest adalah request DTO untuk memperbarui release.
type UpdateReleaseRequest struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Status        string   `json:"status"`
	Year          int      `json:"year"`
	CoverImageURL string   `json:"cover_image_url"`
	Tags          []string `json:"tags"`
}

// CreateStandardRequest adalah request DTO untuk membuat standard baru.
type CreateStandardRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Year        int    `json:"year" binding:"required"`
	FileURL     string `json:"file_url"`
	FileSize    int64  `json:"file_size"`
	Version     string `json:"version"`
	IsCurrent   bool   `json:"is_current"`
}

// UpdateStandardRequest adalah request DTO untuk memperbarui standard.
type UpdateStandardRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	FileURL     string `json:"file_url"`
	FileSize    int64  `json:"file_size"`
	Version     string `json:"version"`
	IsCurrent   bool   `json:"is_current"`
}

// LoginRequest adalah request DTO untuk login.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest adalah request DTO untuk registrasi.
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
}

// TokenResponse adalah response DTO untuk autentikasi.
type TokenResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	TokenType    string       `json:"token_type"`
	ExpiresIn    int          `json:"expires_in"`
	User         UserResponse `json:"user"`
}

// UserResponse adalah response DTO untuk data user.
type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

// ReleaseStatsResponse adalah response DTO untuk statistik releases.
type ReleaseStatsResponse struct {
	Total  int64            `json:"total"`
	ByType map[string]int64 `json:"by_type"`
	ByYear map[string]int64 `json:"by_year"`
}

// JWTClaims adalah claims yang ada di dalam JWT token.
type JWTClaims struct {
	UserID  string `json:"user_id"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	TokenID string `json:"token_id"`
}

// UploadFileRequest adalah request DTO untuk upload file.
type UploadFileRequest struct {
	FileName    string
	FileSize    int64
	ContentType string
	Data        []byte
}

// CreateOrganizationRequest adalah request DTO untuk membuat organisasi baru.
type CreateOrganizationRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Website     string `json:"website"`
}

// UpdateOrganizationRequest adalah request DTO untuk memperbarui organisasi.
type UpdateOrganizationRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url"`
	Website     *string `json:"website"`
}

// CreateGroupRequest adalah request DTO untuk membuat grup baru.
type CreateGroupRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// UpdateGroupRequest adalah request DTO untuk memperbarui grup.
type UpdateGroupRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	ImageURL    *string `json:"image_url"`
}
