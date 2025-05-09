package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/config"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/cookieutil"
)

type AuthHandler struct {
	svc *AuthService
}

func NewAuthHandler(svc *AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func RegisterRoutes(r *gin.Engine, pool *pgxpool.Pool, cfg config.AuthConfig) {
	q := repository.New(pool)

	svc := NewAuthService(q, cfg.JWTSecret)
	h := NewAuthHandler(svc)

	grp := r.Group("/api/v1/auth")
	grp.POST("/register", h.Register)
	grp.POST("/login", h.Login)
	grp.POST("/refresh", h.Refresh)
	grp.POST("/logout", h.Logout)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var in RegisterInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Register(c.Request.Context(), in); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var in LoginInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.svc.Login(c.Request.Context(), in)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	cookie, err := c.Request.Cookie(cookieutil.RefreshCookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
		return
	}

	accessToken, newRefresh, err := h.svc.Refresh(c.Request.Context(), cookie.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cookieutil.SetRefreshCookie(c, newRefresh, RefreshTokenExpiryHours)
	c.JSON(http.StatusOK, gin.H{"token": accessToken})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie(cookieutil.RefreshCookieName)
	if err == nil {
		_ = h.svc.RevokeRefreshToken(c.Request.Context(), cookie.Value)
	}

	cookieutil.ClearRefreshTokenCookie(c)
	c.Status(http.StatusNoContent)
}
