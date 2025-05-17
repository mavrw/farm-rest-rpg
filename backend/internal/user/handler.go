package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/middleware"
)

type UserHandler struct {
	svc *UserService
}

func NewUserHandler(svc *UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func RegisterRoutes(r *gin.RouterGroup, pool *pgxpool.Pool) {
	q := repository.New(pool)

	svc := NewUserService(q)
	h := NewUserHandler(svc)

	// Trying a group instead of prefixing path with '/user' each time
	grp := r.Group("/users")
	grp.GET("/me", h.GetMe)
	grp.PATCH("/me", h.UpdateMe)
	grp.DELETE("/me", h.DeleteMe)
}

func (h *UserHandler) GetMe(c *gin.Context) {
	userId := c.GetInt(middleware.ContextUserIDKey)

	user, err := h.svc.GetMe(c.Request.Context(), int32(userId))
	if err == errs.ErrUserNotFound {
		// ! This should 401 from the auth middleware before this
		// ! happens, but we'll return a 404 in case, i guess???
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	resp := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) UpdateMe(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": errs.ErrNotImplemented.Error()})
}

func (h *UserHandler) DeleteMe(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": errs.ErrNotImplemented.Error()})
}
