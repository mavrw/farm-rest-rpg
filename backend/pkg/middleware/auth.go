package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c.Request)
		userID, err := parseJWT(token, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}

// TODO: MUST BE IMPLEMENTED
func extractToken(req *http.Request) string {
	return ""
}

// TODO: MUST BE IMPLEMENTED
func parseJWT(token, secret string) (int32, error) {
	return -1, nil
}
