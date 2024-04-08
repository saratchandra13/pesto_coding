package tests

import (
	"github.com/didip/tollbooth/v7"
	limiter2 "github.com/didip/tollbooth/v7/limiter"
	"github.com/gin-gonic/gin"
	"github.com/pesto_coding/product_service/app/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimitMiddleware(t *testing.T) {
	// Set up a Gin router with the rate limit middleware
	router := gin.Default()
	lmt := tollbooth.NewLimiter(float64(5), &limiter2.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	router.Use(services.RateLimitMiddleware(lmt)) // Set the rate limit to 5 requests per hour
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Create a test server
	server := httptest.NewServer(router)
	defer server.Close()

	// Create a client
	client := server.Client()

	// Send 5 requests within a short period of time
	for i := 0; i < 5; i++ {
		resp, err := client.Get(server.URL + "/test")
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Check if the response status code is 200
		assert.Equal(t, http.StatusOK, resp.StatusCode)

	}

	// Send another request
	resp, err := client.Get(server.URL + "/test")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is 429, indicating that the rate limit has been exceeded
	assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
}
