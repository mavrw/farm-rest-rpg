package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/jwtutil"
)

const ContextUserIDKey = "current_user_id"

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "missing or malformed auth header"},
			)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		_, claims, err := jwtutil.Parse(tokenStr, secret)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "invalid or expired token"},
			)
			return
		}

		userID, exists := claims["sub"].(float64)
		if !exists {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "invalid token claims"},
			)
			return
		}

		// ! This broke when casting userID to int32
		// ! I believe it's because under the hood, Gin makes a type
		// ! assertion to int when we later call `gin.Context.GetInt`
		// ! which would fail when this value was set as int32.
		c.Set(ContextUserIDKey, int(userID))
		c.Next()
	}
}
