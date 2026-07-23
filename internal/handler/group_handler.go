package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/internal/handler/dto"
	"github.com/satudata/backend/internal/service"
)

// GroupHandler handles HTTP requests for groups/categories.
type GroupHandler struct {
	groupService service.GroupServiceInterface
}

// NewGroupHandler membuat instance baru GroupHandler.
func NewGroupHandler(groupService service.GroupServiceInterface) *GroupHandler {
	return &GroupHandler{
		groupService: groupService,
	}
}

// GetGroups godoc
// @Summary Daftar grup/kategori
// @Description Mengembalikan daftar semua grup/kategori dataset
// @Tags Public
// @Produce json
// @Success 200 {object} dto.APIResponse{data=[]dto.GroupResponse}
// @Router /public/groups [get]
func (h *GroupHandler) GetGroups(c *gin.Context) {
	groups, err := h.groupService.GetGroups(c.Request.Context())
	if err != nil {
		dto.InternalErrorResponse(c, err.Error())
		return
	}

	var response []dto.GroupResponse
	for _, g := range groups {
		response = append(response, dto.GroupResponse{
			ID:          g.ID.String(),
			Slug:        g.Slug,
			Name:        g.Name,
			Description: g.Description,
			ImageURL:    g.ImageURL,
		})
	}

	dto.SuccessResponse(c, response)
}

// GetGroupBySlug godoc
// @Summary Detail grup/kategori
// @Description Mengembalikan detail grup/kategori berdasarkan slug
// @Tags Public
// @Produce json
// @Param slug path string true "Group slug"
// @Success 200 {object} dto.APIResponse{data=dto.GroupResponse}
// @Failure 404 {object} dto.APIResponse
// @Router /public/groups/{slug} [get]
func (h *GroupHandler) GetGroupBySlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		dto.BadRequestResponse(c, "Group slug is required")
		return
	}

	group, err := h.groupService.GetGroupBySlug(c.Request.Context(), slug)
	if err != nil {
		dto.NotFoundResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.GroupResponse{
		ID:          group.ID.String(),
		Slug:        group.Slug,
		Name:        group.Name,
		Description: group.Description,
		ImageURL:    group.ImageURL,
	})
}

// CreateGroup godoc
// @Summary Buat grup baru
// @Description Membuat grup/kategori baru (admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Param request body dto.CreateGroupRequest true "Data grup"
// @Success 201 {object} dto.APIResponse{data=dto.GroupResponse}
// @Security BearerAuth
// @Router /admin/groups [post]
func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var req dto.CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.CreateGroupRequest{
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
	}

	group, err := h.groupService.CreateGroup(c.Request.Context(), svcReq)
	if err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.CreatedResponse(c, dto.GroupResponse{
		ID:          group.ID.String(),
		Slug:        group.Slug,
		Name:        group.Name,
		Description: group.Description,
		ImageURL:    group.ImageURL,
	})
}

// UpdateGroup godoc
// @Summary Update grup
// @Description Memperbarui data grup/kategori (admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param request body dto.UpdateGroupRequest true "Data update"
// @Success 200 {object} dto.APIResponse{data=dto.GroupResponse}
// @Security BearerAuth
// @Router /admin/groups/{id} [put]
func (h *GroupHandler) UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.BadRequestResponse(c, "Group ID is required")
		return
	}

	var req dto.UpdateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.UpdateGroupRequest{
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
	}

	group, err := h.groupService.UpdateGroup(c.Request.Context(), id, svcReq)
	if err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.GroupResponse{
		ID:          group.ID.String(),
		Slug:        group.Slug,
		Name:        group.Name,
		Description: group.Description,
		ImageURL:    group.ImageURL,
	})
}

// DeleteGroup godoc
// @Summary Hapus grup
// @Description Menghapus grup/kategori (admin only)
// @Tags Admin
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} dto.APIResponse
// @Security BearerAuth
// @Router /admin/groups/{id} [delete]
func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		dto.BadRequestResponse(c, "Group ID is required")
		return
	}

	if err := h.groupService.DeleteGroup(c.Request.Context(), id); err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, gin.H{"message": "Group deleted successfully"})
}
