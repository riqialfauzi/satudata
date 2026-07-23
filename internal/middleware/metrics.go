package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/pkg/logger"
	"go.uber.org/zap"
)

// MetricsMiddleware mengumpulkan metrik request.
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		// Log metrics untuk monitoring
		logger.GetLogger().Info("metrics",
			zap.String("metric_type", "request"),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", statusCode),
			zap.Float64("latency_ms", float64(latency.Microseconds())/1000.0),
		)

		// Set response headers untuk debugging
		c.Header("X-Response-Time", strconv.FormatFloat(float64(latency.Microseconds())/1000.0, 'f', 2, 64)+"ms")
	}
}
