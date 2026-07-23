package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/internal/handler/dto"
	"github.com/satudata/backend/internal/service"
)

// StandardHandler handles HTTP requests for standards.
type StandardHandler struct {
	standardService service.StandardServiceInterface
}

// NewStandardHandler membuat instance baru StandardHandler.
func NewStandardHandler(standardService service.StandardServiceInterface) *StandardHandler {
	return &StandardHandler{
		standardService: standardService,
	}
}

// GetStandards menangani GET /api/v1/public/standards
func (h *StandardHandler) GetStandards(c *gin.Context) {
	standards, err := h.standardService.GetStandards(c.Request.Context())
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	var response []dto.StandardResponse
	for _, std := range standards {
		response = append(response, dto.StandardResponse{
			ID:          std.ID.String(),
			Title:       std.Title,
			Description: std.Description,
			Year:        std.Year,
			FileURL:     std.FileURL,
			FileSize:    std.FileSize,
			Status:      string(std.Status),
			Version:     std.Version,
			IsCurrent:   std.IsCurrent,
			CreatedAt:   std.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   std.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	dto.SuccessResponse(c, response)
}

// CreateStandard menangani POST /api/v1/protected/standards
func (h *StandardHandler) CreateStandard(c *gin.Context) {
	var req dto.CreateStandardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.CreateStandardRequest{
		Title:       req.Title,
		Description: req.Description,
		Year:        req.Year,
		FileURL:     req.FileURL,
		FileSize:    req.FileSize,
		Version:     req.Version,
		IsCurrent:   req.IsCurrent,
	}

	std, err := h.standardService.CreateStandard(c.Request.Context(), svcReq)
	if err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.CreatedResponse(c, dto.StandardResponse{
		ID:          std.ID.String(),
		Title:       std.Title,
		Description: std.Description,
		Year:        std.Year,
		FileURL:     std.FileURL,
		FileSize:    std.FileSize,
		Status:      string(std.Status),
		Version:     std.Version,
		IsCurrent:   std.IsCurrent,
		CreatedAt:   std.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   std.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	})
}

// UpdateStandard menangani PUT /api/v1/protected/standards/:id
func (h *StandardHandler) UpdateStandard(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.BadRequestResponse(c, "Standard ID is required")
		return
	}

	var req dto.UpdateStandardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.UpdateStandardRequest{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		FileURL:     req.FileURL,
		FileSize:    req.FileSize,
		Version:     req.Version,
		IsCurrent:   req.IsCurrent,
	}

	std, err := h.standardService.UpdateStandard(c.Request.Context(), id, svcReq)
	if err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.StandardResponse{
		ID:          std.ID.String(),
		Title:       std.Title,
		Description: std.Description,
		Year:        std.Year,
		FileURL:     std.FileURL,
		FileSize:    std.FileSize,
		Status:      string(std.Status),
		Version:     std.Version,
		IsCurrent:   std.IsCurrent,
		CreatedAt:   std.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   std.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	})
}
