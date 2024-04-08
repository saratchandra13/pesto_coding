package services

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"time"
)

func RateLimitMiddleware(max int64) gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(float64(max), &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.AbortWithStatusJSON(httpError.StatusCode, gin.H{"error": httpError.Message})
			return
		}
		c.Next()
	}
}
