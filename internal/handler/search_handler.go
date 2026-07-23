package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/internal/handler/dto"
	"github.com/satudata/backend/internal/service"
)

// SearchHandler handles HTTP requests for search.
type SearchHandler struct {
	searchService  *service.SearchService
	releaseService service.ReleaseServiceInterface
}

// NewSearchHandler membuat instance baru SearchHandler.
func NewSearchHandler(searchService *service.SearchService, releaseService service.ReleaseServiceInterface) *SearchHandler {
	return &SearchHandler{
		searchService:  searchService,
		releaseService: releaseService,
	}
}

// SearchSuggest godoc
// @Summary Saran pencarian (autocomplete)
// @Description Mengembalikan saran judul dataset berdasarkan query pencarian
// @Tags Public
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} dto.APIResponse{data=dto.SearchSuggestResponse}
// @Router /public/search/suggest [get]
func (h *SearchHandler) SearchSuggest(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		dto.SuccessResponse(c, dto.SearchSuggestResponse{
			Suggestions: []string{},
			Results:     []dto.ReleaseResponse{},
		})
		return
	}

	result, err := h.searchService.SearchSuggest(c.Request.Context(), q)
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	var releases []dto.ReleaseResponse
	for _, release := range result.Datasets {
		releases = append(releases, mapReleaseToResponse(release))
	}

	dto.SuccessResponse(c, dto.SearchSuggestResponse{
		Suggestions: result.Suggestions,
		Results:     releases,
	})
}

// Search godoc
// @Summary Pencarian global
// @Description Mencari dataset, artikel, dan infografis
// @Tags Public
// @Produce json
// @Param q query string true "Search query"
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10)"
// @Success 200 {object} dto.APIResponse{data=[]dto.ReleaseResponse}
// @Router /public/search [get]
func (h *SearchHandler) Search(c *gin.Context) {
	q := c.Query("q")
	page := 1
	limit := 10

	if q == "" {
		dto.SuccessResponseWithMeta(c, []dto.ReleaseResponse{}, &dto.MetaData{
			Page:       1,
			Limit:      10,
			Total:      0,
			TotalPages: 0,
		})
		return
	}

	releases, total, err := h.releaseService.GetReleases(c.Request.Context(), service.ReleaseFilterRequest{
		Search:  q,
		Page:    page,
		Limit:   limit,
		SortBy:  "created_at",
		SortDir: "desc",
	})
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	var response []dto.ReleaseResponse
	for _, release := range releases {
		response = append(response, mapReleaseToResponse(release))
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
