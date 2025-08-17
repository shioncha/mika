package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/repository"
)

type RateLimitMiddleware struct {
	rateLimitRepo repository.RateLimitRepository
	limit         int64
	ttl           time.Duration
}

func NewRateLimitMiddleware(rateLimitRepo repository.RateLimitRepository, limit int64, ttl time.Duration) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		rateLimitRepo: rateLimitRepo,
		limit:         limit,
		ttl:           ttl,
	}
}

func (m *RateLimitMiddleware) IPBasedRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "ratelimit:" + ip

		count, err := m.rateLimitRepo.Increment(c.Request.Context(), key, m.ttl)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		if count > m.limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}
