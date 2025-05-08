package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ctxKey string

const userIDKey ctxKey = "current_user_id"

// inject current_user_id from Gin into standard ctx
func RLS() gin.HandlerFunc {
	return func(c *gin.Context) {
		if uid, ok := c.Get("userID"); ok {
			ctx := context.WithValue(c.Request.Context(), userIDKey, uid.(int32))
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}

func PGXBeforeAcquire(ctx context.Context, conn *pgx.Conn) bool {
	if v := ctx.Value(userIDKey); v != nil {
		conn.Exec(ctx, "SELECT util.set_current_user_id($1)", v.(int32))
	}
	return true
}

func PGXAfterRelease(conn *pgx.Conn) bool {
	conn.Exec(context.Background(), "SELECT util.clear_user_id()")
	return true
}
