package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/internal/domain"
	"github.com/satudata/backend/internal/handler/dto"
	"github.com/satudata/backend/internal/service"
)

// ReleaseHandler handles HTTP requests for releases.
type ReleaseHandler struct {
	releaseService service.ReleaseServiceInterface
}

// NewReleaseHandler membuat instance baru ReleaseHandler.
func NewReleaseHandler(releaseService service.ReleaseServiceInterface) *ReleaseHandler {
	return &ReleaseHandler{
		releaseService: releaseService,
	}
}

// GetReleases godoc
// @Summary Daftar releases
// @Description Mengembalikan daftar releases dengan filter, pagination, dan sorting
// @Tags Public
// @Accept json
// @Produce json
// @Param type query string false "Filter by type: dataset, article, infographic"
// @Param status query string false "Filter by status: draft, published, archived"
// @Param year query int false "Filter by year"
// @Param search query string false "Search by title or description"
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10, max: 100)"
// @Param sort_by query string false "Sort field (default: created_at)"
// @Param sort_dir query string false "Sort direction: asc, desc (default: desc)"
// @Success 200 {object} dto.APIResponse{data=[]dto.ReleaseResponse}
// @Router /public/releases [get]
func (h *ReleaseHandler) GetReleases(c *gin.Context) {
	var req dto.GetReleasesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid query parameters")
		return
	}

	svcReq := service.ReleaseFilterRequest{
		Type:    req.Type,
		Status:  req.Status,
		Year:    req.Year,
		Search:  req.Search,
		Page:    req.Page,
		Limit:   req.Limit,
		SortBy:  req.SortBy,
		SortDir: req.SortDir,
	}

	releases, total, err := h.releaseService.GetReleases(c.Request.Context(), svcReq)
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	// Map to response DTOs
	var response []dto.ReleaseResponse
	for _, release := range releases {
		response = append(response, mapReleaseToResponse(release))
	}

	page := req.Page
	if page < 1 {
		page = 1
	}
	limit := req.Limit
	if limit < 1 {
		limit = 10
	}
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	dto.SuccessResponseWithMeta(c, response, &dto.MetaData{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	})
}

// GetReleaseByID godoc
// @Summary Detail release by ID
// @Description Mengembalikan detail release berdasarkan ID
// @Tags Public
// @Produce json
// @Param id path string true "Release ID"
// @Success 200 {object} dto.APIResponse{data=dto.ReleaseResponse}
// @Failure 404 {object} dto.APIResponse
// @Router /public/releases/{id} [get]
func (h *ReleaseHandler) GetReleaseByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.BadRequestResponse(c, "Release ID is required")
		return
	}

	release, err := h.releaseService.GetReleaseByID(c.Request.Context(), id)
	if err != nil {
		dto.NotFoundResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, mapReleaseToResponse(*release))
}

// GetReleaseBySlug godoc
// @Summary Detail release by Slug
// @Description Mengembalikan detail release berdasarkan slug (untuk URL SEO-friendly)
// @Tags Public
// @Produce json
// @Param slug path string true "Release slug"
// @Success 200 {object} dto.APIResponse{data=dto.ReleaseResponse}
// @Failure 404 {object} dto.APIResponse
// @Router /public/releases/slug/{slug} [get]
func (h *ReleaseHandler) GetReleaseBySlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		dto.BadRequestResponse(c, "Release slug is required")
		return
	}

	release, err := h.releaseService.GetReleaseBySlug(c.Request.Context(), slug)
	if err != nil {
		dto.NotFoundResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, mapReleaseToResponse(*release))
}

// CreateRelease godoc
// @Summary Buat release baru
// @Description Membuat release baru (dataset, article, atau infographic)
// @Tags Protected
// @Accept json
// @Produce json
// @Param request body dto.CreateReleaseRequest true "Data release"
// @Success 201 {object} dto.APIResponse{data=dto.ReleaseResponse}
// @Failure 400 {object} dto.APIResponse
// @Security BearerAuth
// @Router /protected/releases [post]
func (h *ReleaseHandler) CreateRelease(c *gin.Context) {
	var req dto.CreateReleaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.CreateReleaseRequest{
		Title:           req.Title,
		Description:     req.Description,
		ReleaseType:     req.ReleaseType,
		Year:            req.Year,
		CoverImageURL:   req.CoverImageURL,
		Tags:            req.Tags,
		FileURL:         req.FileURL,
		FileFormat:      req.FileFormat,
		FileSize:        req.FileSize,
		DataSource:      req.DataSource,
		DataPeriodStart: req.DataPeriodStart,
		DataPeriodEnd:   req.DataPeriodEnd,
		UpdateFrequency: req.UpdateFrequency,
		IsGeospatial:    req.IsGeospatial,
		Content:         req.Content,
		Excerpt:         req.Excerpt,
		AuthorName:      req.AuthorName,
		Category:        req.Category,
		IsFeatured:      req.IsFeatured,
	}

	release, err := h.releaseService.CreateRelease(c.Request.Context(), svcReq)
	if err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.CreatedResponse(c, mapReleaseToResponse(*release))
}

// UpdateRelease godoc
// @Summary Update release
// @Description Memperbarui data release yang sudah ada
// @Tags Protected
// @Accept json
// @Produce json
// @Param id path string true "Release ID"
// @Param request body dto.UpdateReleaseRequest true "Data update"
// @Success 200 {object} dto.APIResponse{data=dto.ReleaseResponse}
// @Failure 400 {object} dto.APIResponse
// @Security BearerAuth
// @Router /protected/releases/{id} [put]
func (h *ReleaseHandler) UpdateRelease(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.BadRequestResponse(c, "Release ID is required")
		return
	}

	var req dto.UpdateReleaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.UpdateReleaseRequest{
		Title:         req.Title,
		Description:   req.Description,
		Status:        req.Status,
		Year:          req.Year,
		CoverImageURL: req.CoverImageURL,
		Tags:          req.Tags,
	}

	release, err := h.releaseService.UpdateRelease(c.Request.Context(), id, svcReq)
	if err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, mapReleaseToResponse(*release))
}

// GetRelatedReleases godoc
// @Summary Releases terkait
// @Description Mengembalikan daftar releases terkait berdasarkan tags yang sama
// @Tags Public
// @Produce json
// @Param id path string true "Release ID"
// @Param limit query int false "Number of related items (default: 5, max: 20)"
// @Success 200 {object} dto.APIResponse{data=[]dto.ReleaseResponse}
// @Router /public/releases/{id}/related [get]
func (h *ReleaseHandler) GetRelatedReleases(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.BadRequestResponse(c, "Release ID is required")
		return
	}

	limit := 5
	if l, err := strconv.Atoi(c.Query("limit")); err == nil && l > 0 {
		limit = l
	}

	releases, err := h.releaseService.GetRelatedReleases(c.Request.Context(), id, limit)
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	var response []dto.ReleaseResponse
	for _, release := range releases {
		response = append(response, mapReleaseToResponse(release))
	}

	dto.SuccessResponse(c, response)
}

// DeleteRelease godoc
// @Summary Hapus release (soft-delete)
// @Description Menghapus release secara soft-delete (admin only)
// @Tags Admin
// @Produce json
// @Param id path string true "Release ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.APIResponse
// @Security BearerAuth
// @Router /admin/releases/{id} [delete]
func (h *ReleaseHandler) DeleteRelease(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.BadRequestResponse(c, "Release ID is required")
		return
	}

	if err := h.releaseService.DeleteRelease(c.Request.Context(), id); err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, gin.H{"message": "Release deleted successfully"})
}

// GetReleaseStats godoc
// @Summary Statistik releases
// @Description Mengembalikan statistik releases (total by type, by year)
// @Tags Public
// @Produce json
// @Success 200 {object} dto.APIResponse{data=dto.ReleaseStatsResponse}
// @Router /public/releases/stats [get]
func (h *ReleaseHandler) GetReleaseStats(c *gin.Context) {
	stats, err := h.releaseService.GetReleaseStats(c.Request.Context())
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.ReleaseStatsResponse{
		Total:  stats.Total,
		ByType: stats.ByType,
		ByYear: stats.ByYear,
	})
}

// mapReleaseToResponse memetakan domain Release ke ReleaseResponse DTO.
func mapReleaseToResponse(release domain.Release) dto.ReleaseResponse {
	resp := dto.ReleaseResponse{
		ID:            release.ID.String(),
		Title:         release.Title,
		Slug:          release.Slug,
		Description:   release.Description,
		ReleaseType:   string(release.ReleaseType),
		Status:        string(release.Status),
		Year:          release.Year,
		CoverImageURL: release.CoverImageURL,
		Tags:          release.Tags,
		ViewCount:     release.ViewCount,
		CreatedAt:     release.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:     release.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	if release.PublishedAt != nil {
		pub := release.PublishedAt.Format("2006-01-02T15:04:05Z")
		resp.PublishedAt = &pub
	}

	if release.CreatedBy != nil {
		cb := release.CreatedBy.String()
		resp.CreatedBy = &cb
	}

	if release.DatasetMetadata != nil {
		dm := release.DatasetMetadata
		resp.DatasetMetadata = &dto.DatasetMetadataResponse{
			FileURL:         dm.FileURL,
			FileFormat:      dm.FileFormat,
			FileSize:        dm.FileSize,
			RowCount:        dm.RowCount,
			ColumnCount:     dm.ColumnCount,
			DataSource:      dm.DataSource,
			UpdateFrequency: dm.UpdateFrequency,
			IsGeospatial:    dm.IsGeospatial,
		}
		if dm.DataPeriodStart != nil {
			s := dm.DataPeriodStart.Format("2006-01-02")
			resp.DatasetMetadata.DataPeriodStart = &s
		}
		if dm.DataPeriodEnd != nil {
			s := dm.DataPeriodEnd.Format("2006-01-02")
			resp.DatasetMetadata.DataPeriodEnd = &s
		}
	}

	if release.ArticleMetadata != nil {
		am := release.ArticleMetadata
		resp.ArticleMetadata = &dto.ArticleMetadataResponse{
			Content:     am.Content,
			Excerpt:     am.Excerpt,
			AuthorName:  am.AuthorName,
			ReadingTime: am.ReadingTimeMin,
			IsFeatured:  am.IsFeatured,
			Category:    am.Category,
		}
	}

	return resp
}
