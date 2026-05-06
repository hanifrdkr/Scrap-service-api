package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func RequestIDMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		requestID := uuid.New().String()
		c.Set("request_id", requestID)
		c.Header("Request-ID", requestID)

		c.Next()
	}
}
