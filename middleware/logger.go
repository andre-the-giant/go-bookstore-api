package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path

		c.Next()

		duration := time.Since(startTime)
		status := c.Writer.Status()
		println("➡️", method, path, "⬅️", status, duration)
	}
}
