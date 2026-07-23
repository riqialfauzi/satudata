package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/internal/handler/dto"
	"github.com/satudata/backend/internal/middleware"
	"github.com/satudata/backend/internal/service"
)

// AuthHandler handles HTTP requests for authentication.
type AuthHandler struct {
	authService service.AuthServiceInterface
}

// NewAuthHandler membuat instance baru AuthHandler.
func NewAuthHandler(authService service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login godoc
// @Summary Login user
// @Description Autentikasi user dengan email dan password, mengembalikan JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Credentials"
// @Success 200 {object} dto.APIResponse{data=dto.TokenResponse}
// @Failure 401 {object} dto.APIResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	tokenResp, err := h.authService.Login(c.Request.Context(), svcReq)
	if err != nil {
		dto.UnauthorizedResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.TokenResponse{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		TokenType:    tokenResp.TokenType,
		ExpiresIn:    tokenResp.ExpiresIn,
		User: dto.UserResponse{
			ID:       tokenResp.User.ID,
			Email:    tokenResp.User.Email,
			FullName: tokenResp.User.FullName,
			Role:     tokenResp.User.Role,
		},
	})
}

// Register godoc
// @Summary Register user baru
// @Description Mendaftarkan user baru
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Data registrasi"
// @Success 201 {object} dto.APIResponse{data=dto.UserResponse}
// @Failure 400 {object} dto.APIResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	svcReq := service.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
	}

	user, err := h.authService.Register(c.Request.Context(), svcReq)
	if err != nil {
		dto.BadRequestResponse(c, err.Error())
		return
	}

	dto.CreatedResponse(c, dto.UserResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		FullName: user.FullName,
		Role:     string(user.Role),
	})
}

// RefreshToken godoc
// @Summary Refresh token
// @Description Memperbarui access token menggunakan refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} dto.APIResponse{data=dto.TokenResponse}
// @Failure 401 {object} dto.APIResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dto.BadRequestResponse(c, "Invalid request body", err.Error())
		return
	}

	tokenResp, err := h.authService.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		dto.UnauthorizedResponse(c, err.Error())
		return
	}

	dto.SuccessResponse(c, dto.TokenResponse{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		TokenType:    tokenResp.TokenType,
		ExpiresIn:    tokenResp.ExpiresIn,
		User: dto.UserResponse{
			ID:       tokenResp.User.ID,
			Email:    tokenResp.User.Email,
			FullName: tokenResp.User.FullName,
			Role:     tokenResp.User.Role,
		},
	})
}

// Logout godoc
// @Summary Logout user
// @Description Menghapus session token
// @Tags Auth
// @Produce json
// @Success 200 {object} dto.APIResponse
// @Security BearerAuth
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" && len(authHeader) > 7 {
		token := authHeader[7:]
		_ = h.authService.Logout(c.Request.Context(), token)
	}

	dto.SuccessResponse(c, gin.H{"message": "Logged out successfully"})
}

// GetProfile godoc
// @Summary Profile user
// @Description Mengembalikan informasi profile user yang sedang login
// @Tags Protected
// @Produce json
// @Success 200 {object} dto.APIResponse
// @Security BearerAuth
// @Router /protected/profile [get]
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	dto.SuccessResponse(c, gin.H{
		"user_id": userID,
		"role":    userRole,
	})
}
