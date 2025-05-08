package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		status := c.Writer.Status()

		// write request info to console
		log.Printf("%s %s %d %s",
			c.Request.Method, c.Request.URL.Path, status, latency,
		)
	}
}
