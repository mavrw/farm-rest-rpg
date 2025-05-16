package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/jwtutil"
)

const ContextUserIDKey = "current_user_id"

func AuthMiddleware(secret string, db *pgxpool.Pool) gin.HandlerFunc {
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

		// verify that the userID from the claim is exists in the db.
		// ? Can this be improved to restore as much stateless-ness as
		// ? possible? Caching could be done, but that needs special
		// ? attention to ensure it is being invalidated correctly. This
		// ? hybrid approach might be the best option for now. Maybe a
		// ? better solution would be a specialized query on only the
		// ? indexed `id` column so that the query is as fast as posisble?
		q := repository.New(db)
		if _, err := q.GetUserByID(c.Request.Context(), int32(userID)); err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "user no longer exists"},
			)
			return
		}

		// ! This broke when casting userID to int32
		// ! I believe it's because under the hood, Gin makes a type
		// ! assertion to int when we later call `gin.Context.GetInt`
		// ! which would fail when this value was set as int32.
		// ! The real solution is to stop using int32, but try and
		// ! stop me though...
		c.Set(ContextUserIDKey, int(userID))
		c.Next()
	}
}
