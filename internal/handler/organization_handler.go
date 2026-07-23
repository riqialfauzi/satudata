package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/internal/handler/dto"
	"github.com/satudata/backend/internal/service"
)

// OrganizationHandler handles HTTP requests for organizations.
type OrganizationHandler struct {
	orgService service.OrganizationServiceInterface
}

// NewOrganizationHandler membuat instance baru OrganizationHandler.
func NewOrganizationHandler(orgService service.OrganizationServiceInterface) *OrganizationHandler {
	return &OrganizationHandler{
		orgService: orgService,
	}
}

// GetOrganizations godoc
// @Summary Daftar organisasi
// @Description Mengembalikan daftar semua organisasi/unit kerja
// @Tags Public
// @Produce json
// @Success 200 {object} dto.APIResponse{data=[]dto.OrganizationResponse}
// @Router /public/organizations [get]
func (h *OrganizationHandler) GetOrganizations(c *gin.Context) {
	organizations, err := h.orgService.GetOrganizations(c.Request.Context())
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	var response []dto.OrganizationResponse
	for _, org := range organizations {
		response = append(response, dto.OrganizationResponse{
			ID:          org.ID.String(),
			Slug:        org.Slug,
			Name:        org.Name,
			Description: org.Description,
			ImageURL:    org.ImageURL,
			Website:     org.Website,
		})
	}

	dto.SuccessResponse(c, response)
}

// GetOrganizationBySlug godoc
// @Summary Detail organisasi
// @Description Mengembalikan detail organisasi berdasarkan slug
// @Tags Public
// @Produce json
// @Param slug path string true "Organization slug"
// @Success 200 {object} dto.APIResponse{data=dto.OrganizationResponse}
// @Failure 404 {object} dto.APIResponse
// @Router /public/organizations/{slug} [get]
func (h *OrganizationHandler) GetOrganizationBySlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		dto.BadRequestResponse(c, "Organization slug is required")
		return
	}

	org, err := h.orgService.GetOrganizationBySlug(c.Request.Context(), slug)
	if err != nil {
		dto.NotFoundResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.OrganizationResponse{
		ID:          org.ID.String(),
		Slug:        org.Slug,
		Name:        org.Name,
		Description: org.Description,
		ImageURL:    org.ImageURL,
		Website:     org.Website,
	})
}

// CreateOrganization godoc
// @Summary Buat organisasi baru
// @Description Membuat organisasi baru (admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Param request body dto.CreateOrganizationRequest true "Data organisasi"
// @Success 201 {object} dto.APIResponse{data=dto.OrganizationResponse}
// @Security BearerAuth
// @Router /admin/organizations [post]
func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var req dto.CreateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.CreateOrganizationRequest{
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		Website:     req.Website,
	}

	org, err := h.orgService.CreateOrganization(c.Request.Context(), svcReq)
	if err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.CreatedResponse(c, dto.OrganizationResponse{
		ID:          org.ID.String(),
		Slug:        org.Slug,
		Name:        org.Name,
		Description: org.Description,
		ImageURL:    org.ImageURL,
		Website:     org.Website,
	})
}

// UpdateOrganization godoc
// @Summary Update organisasi
// @Description Memperbarui data organisasi (admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Param request body dto.UpdateOrganizationRequest true "Data update"
// @Success 200 {object} dto.APIResponse{data=dto.OrganizationResponse}
// @Security BearerAuth
// @Router /admin/organizations/{id} [put]
func (h *OrganizationHandler) UpdateOrganization(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.BadRequestResponse(c, "Organization ID is required")
		return
	}

	var req dto.UpdateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.UpdateOrganizationRequest{
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		Website:     req.Website,
	}

	org, err := h.orgService.UpdateOrganization(c.Request.Context(), id, svcReq)
	if err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.OrganizationResponse{
		ID:          org.ID.String(),
		Slug:        org.Slug,
		Name:        org.Name,
		Description: org.Description,
		ImageURL:    org.ImageURL,
		Website:     org.Website,
	})
}

// DeleteOrganization godoc
// @Summary Hapus organisasi
// @Description Menghapus organisasi (admin only)
// @Tags Admin
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} dto.APIResponse
// @Security BearerAuth
// @Router /admin/organizations/{id} [delete]
func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.BadRequestResponse(c, "Organization ID is required")
		return
	}

	if err := h.orgService.DeleteOrganization(c.Request.Context(), id); err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, gin.H{"message": "Organization deleted successfully"})
}
