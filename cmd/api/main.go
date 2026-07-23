// @title Satudata API
// @version 1.0
// @description Portal data terbuka API — dataset, artikel, standar data, dan infografis.
// @termsOfService https://satudata.go.id/terms

// @contact.name Tim Satudata
// @contact.email info@satudata.go.id

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token dengan format: `Bearer <token>`
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/internal/config"
	"github.com/satudata/backend/internal/router"
	"github.com/satudata/backend/pkg/cache"
	"github.com/satudata/backend/pkg/database"
	"github.com/satudata/backend/pkg/logger"
	"github.com/satudata/backend/pkg/storage"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	if err := logger.Init(cfg.App.LogLevel, cfg.App.Environment); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	l := logger.GetLogger()
	l.Info("Starting Satudata API server...",
		zap.String("environment", cfg.App.Environment),
		zap.Int("port", cfg.Server.Port),
	)

	// Initialize database
	db, err := database.Init(cfg.Database)
	if err != nil {
		l.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer database.Close()

	// Initialize Redis cache
	redisCache, err := cache.Init(cfg.Redis)
	if err != nil {
		l.Warn("Failed to initialize Redis cache, continuing without cache", zap.Error(err))
	} else {
		defer cache.Close()
	}

	// Initialize MinIO storage
	minioClient, err := storage.Init(cfg.MinIO)
	if err != nil {
		l.Warn("Failed to initialize MinIO storage, continuing without storage", zap.Error(err))
		minioClient = nil
	}

	// Setup Gin with middleware and routes
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		healthStatus := gin.H{
			"status": "ok",
			"time":   time.Now().UTC(),
		}

		if err := database.HealthCheck(); err != nil {
			healthStatus["database"] = "unhealthy"
		} else {
			healthStatus["database"] = "healthy"
		}

		if redisCache != nil {
			if err := redisCache.Ping(context.Background()); err != nil {
				healthStatus["redis"] = "unhealthy"
			} else {
				healthStatus["redis"] = "healthy"
			}
		} else {
			healthStatus["redis"] = "not configured"
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    healthStatus,
		})
	})

	// Setup all API routes
	router.SetupRoutes(r, cfg, db, redisCache, minioClient)

	// Setup HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Graceful shutdown
	go func() {
		l.Info("Server listening", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	l.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		l.Fatal("Server forced to shutdown", zap.Error(err))
	}

	l.Info("Server exited")
}
