package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/pkg/logger"
	"go.uber.org/zap"
)

// LoggerMiddleware mencatat semua HTTP request dengan Zap logger.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		fields := []zap.Field{
			zap.Int("status", statusCode),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", raw),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		}

		if len(c.Errors) > 0 {
			fields = append(fields, zap.Strings("errors", c.Errors.Errors()))
		}

		switch {
		case statusCode >= 500:
			logger.GetLogger().Error("Server error", fields...)
		case statusCode >= 400:
			logger.GetLogger().Warn("Client error", fields...)
		default:
			logger.GetLogger().Info("Request completed", fields...)
		}
	}
}
