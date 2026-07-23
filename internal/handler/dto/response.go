package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse adalah wrapper standard untuk semua API responses.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *MetaData   `json:"meta,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

// MetaData berisi metadata pagination.
type MetaData struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// SuccessResponse mengembalikan response sukses.
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
	})
}

// SuccessResponseWithMeta mengembalikan response sukses dengan pagination.
func SuccessResponseWithMeta(c *gin.Context, data interface{}, meta *MetaData) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

// CreatedResponse mengembalikan response created (201).
func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Data:    data,
	})
}

// ErrorResponse mengembalikan response error.
func ErrorResponse(c *gin.Context, statusCode int, message string, errors ...string) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// BadRequestResponse mengembalikan response 400.
func BadRequestResponse(c *gin.Context, message string, errors ...string) {
	ErrorResponse(c, http.StatusBadRequest, message, errors...)
}

// NotFoundResponse mengembalikan response 404.
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message)
}

// UnauthorizedResponse mengembalikan response 401.
func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message)
}

// ForbiddenResponse mengembalikan response 403.
func ForbiddenResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message)
}

// InternalErrorResponse mengembalikan response 500.
func InternalErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message)
}

// ReleaseResponse adalah response DTO untuk data release.
type ReleaseResponse struct {
	ID              string                   `json:"id"`
	Title           string                   `json:"title"`
	Slug            string                   `json:"slug"`
	Description     string                   `json:"description,omitempty"`
	ReleaseType     string                   `json:"release_type"`
	Status          string                   `json:"status"`
	Year            int                      `json:"year"`
	CoverImageURL   string                   `json:"cover_image_url,omitempty"`
	Tags            []string                 `json:"tags,omitempty"`
	ViewCount       int                      `json:"view_count"`
	PublishedAt     *string                  `json:"published_at,omitempty"`
	CreatedBy       *string                  `json:"created_by,omitempty"`
	DatasetMetadata *DatasetMetadataResponse `json:"dataset_metadata,omitempty"`
	ArticleMetadata *ArticleMetadataResponse `json:"article_metadata,omitempty"`
	CreatedAt       string                   `json:"created_at"`
	UpdatedAt       string                   `json:"updated_at"`
}

// DatasetMetadataResponse adalah response DTO untuk dataset metadata.
type DatasetMetadataResponse struct {
	FileURL         string  `json:"file_url"`
	FileFormat      string  `json:"file_format"`
	FileSize        int64   `json:"file_size"`
	RowCount        *int    `json:"row_count,omitempty"`
	ColumnCount     *int    `json:"column_count,omitempty"`
	DataSource      string  `json:"data_source,omitempty"`
	DataPeriodStart *string `json:"data_period_start,omitempty"`
	DataPeriodEnd   *string `json:"data_period_end,omitempty"`
	UpdateFrequency string  `json:"update_frequency,omitempty"`
	IsGeospatial    bool    `json:"is_geospatial"`
}

// ArticleMetadataResponse adalah response DTO untuk article metadata.
type ArticleMetadataResponse struct {
	Content     string `json:"content"`
	Excerpt     string `json:"excerpt,omitempty"`
	AuthorName  string `json:"author_name,omitempty"`
	ReadingTime *int   `json:"reading_time_minutes,omitempty"`
	IsFeatured  bool   `json:"is_featured"`
	Category    string `json:"category,omitempty"`
}

// StandardResponse adalah response DTO untuk standard.
type StandardResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Year        int    `json:"year"`
	FileURL     string `json:"file_url,omitempty"`
	FileSize    int64  `json:"file_size,omitempty"`
	Status      string `json:"status"`
	Version     string `json:"version"`
	IsCurrent   bool   `json:"is_current"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// TokenResponse adalah response DTO untuk autentikasi.
type TokenResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	TokenType    string       `json:"token_type"`
	ExpiresIn    int          `json:"expires_in"`
	User         UserResponse `json:"user"`
}

// OrganizationResponse adalah response DTO untuk organisasi.
type OrganizationResponse struct {
	ID           string `json:"id"`
	Slug         string `json:"slug"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
	Website      string `json:"website,omitempty"`
	DatasetCount int    `json:"dataset_count,omitempty"`
}

// GroupResponse adalah response DTO untuk grup/kategori.
type GroupResponse struct {
	ID           string `json:"id"`
	Slug         string `json:"slug"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
	DatasetCount int    `json:"dataset_count,omitempty"`
}

// SearchSuggestResponse adalah response DTO untuk search suggest.
type SearchSuggestResponse struct {
	Suggestions []string          `json:"suggestions"`
	Results     []ReleaseResponse `json:"results,omitempty"`
}

// UploadResponse adalah response DTO untuk upload file.
type UploadResponse struct {
	URL        string `json:"url"`
	FileName   string `json:"file_name"`
	FileSize   int64  `json:"file_size"`
	FileFormat string `json:"file_format"`
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
