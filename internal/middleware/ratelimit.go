package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satudata/backend/pkg/cache"
)

// RateLimiterConfig holds configuration for rate limiting.
type RateLimiterConfig struct {
	MaxRequests int
	WindowTime  time.Duration
}

// RateLimiterMiddleware adalah middleware untuk rate limiting berbasis IP.
// Menggunakan Redis untuk sliding window counter.
func RateLimiterMiddleware(redisCache *cache.RedisCache, config RateLimiterConfig) gin.HandlerFunc {
	if redisCache == nil {
		// If Redis is not available, skip rate limiting
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "ratelimit:" + ip
		ctx := context.Background()

		count, err := redisCache.Incr(ctx, key)
		if err != nil {
			// If Redis fails, allow the request
			c.Next()
			return
		}

		// Set expiry on first request
		if count == 1 {
			_ = redisCache.Expire(ctx, key, config.WindowTime)
		}

		if count > int64(config.MaxRequests) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Too many requests. Please try again later.",
			})
			return
		}

		c.Next()
	}
}
