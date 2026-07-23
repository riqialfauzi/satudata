package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/internal/config"
	"github.com/satudata/backend/internal/handler"
	"github.com/satudata/backend/internal/middleware"
	"github.com/satudata/backend/internal/repository"
	"github.com/satudata/backend/internal/service"
	"github.com/satudata/backend/pkg/cache"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	_ "github.com/satudata/backend/internal/docs" // Swagger docs
)

// SetupRoutes mengkonfigurasi semua routes API.
func SetupRoutes(
	router *gin.Engine,
	cfg *config.Config,
	db *gorm.DB,
	redisCache *cache.RedisCache,
) {
	// Initialize repositories
	cacheRepo := repository.NewCacheRepository(redisCache)

	releaseRepo := repository.NewReleaseRepository(db, cacheRepo)
	standardRepo := repository.NewStandardRepository(db, cacheRepo)
	userRepo := repository.NewUserRepository(db)
	_ = repository.NewAuditLogRepository(db) // Will be used later

	// Initialize services
	releaseService := service.NewReleaseService(releaseRepo)
	standardService := service.NewStandardService(standardRepo)
	authService := service.NewAuthService(userRepo, cfg.JWT)

	// Initialize handlers
	releaseHandler := handler.NewReleaseHandler(releaseService)
	standardHandler := handler.NewStandardHandler(standardService)
	authHandler := handler.NewAuthHandler(authService)

	// Apply global middleware
	router.Use(middleware.CORSMiddleware(cfg.App.CORSOrigins))
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.MetricsMiddleware())

	// Apply rate limiter if Redis is available
	if redisCache != nil {
		router.Use(middleware.RateLimiterMiddleware(redisCache, middleware.RateLimiterConfig{
			MaxRequests: 100,
			WindowTime:  1 * time.Minute,
		}))
	}

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes
	apiV1 := router.Group("/api/v1")
	{
		// ========== Public Endpoints (no auth) ==========
		public := apiV1.Group("/public")
		{
			// Health check
			public.GET("/health", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"success": true,
					"message": "Satudata API v1",
				})
			})

			// Releases
			public.GET("/releases", releaseHandler.GetReleases)
			public.GET("/releases/:id", releaseHandler.GetReleaseByID)
			public.GET("/releases/stats", releaseHandler.GetReleaseStats)

			// Standards
			public.GET("/standards", standardHandler.GetStandards)
		}

		// ========== Auth Endpoints ==========
		auth := apiV1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.POST("/logout", authHandler.Logout)
		}

		// ========== Protected Endpoints (auth required) ==========
		protected := apiV1.Group("/protected")
		protected.Use(middleware.AuthMiddleware(authService))
		{
			// Profile
			protected.GET("/profile", authHandler.GetProfile)

			// Releases CRUD
			protected.POST("/releases", releaseHandler.CreateRelease)
			protected.PUT("/releases/:id", releaseHandler.UpdateRelease)

			// Standards CRUD
			protected.POST("/standards", standardHandler.CreateStandard)
			protected.PUT("/standards/:id", standardHandler.UpdateStandard)
		}

		// ========== Admin Endpoints (admin only) ==========
		admin := apiV1.Group("/admin")
		admin.Use(middleware.AuthMiddleware(authService))
		admin.Use(middleware.AdminMiddleware())
		{
			// Releases - Delete
			admin.DELETE("/releases/:id", releaseHandler.DeleteRelease)

			// Users management
			admin.GET("/users", func(c *gin.Context) {
				c.JSON(200, gin.H{"success": true, "message": "Users list - to be implemented"})
			})
			admin.PUT("/users/:id/role", func(c *gin.Context) {
				c.JSON(200, gin.H{"success": true, "message": "Update user role - to be implemented"})
			})

			// Audit logs
			admin.GET("/audit-logs", func(c *gin.Context) {
				c.JSON(200, gin.H{"success": true, "message": "Audit logs - to be implemented"})
			})
		}
	}

	// Slug-based route (outside /public prefix for cleaner URLs)
	router.GET("/api/v1/public/releases/slug/:slug", releaseHandler.GetReleaseBySlug)
}
