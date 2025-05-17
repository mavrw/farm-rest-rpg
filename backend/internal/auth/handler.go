package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/config"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/cookieutil"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
)

type AuthHandler struct {
	svc *AuthService
}

func NewAuthHandler(svc *AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func RegisterRoutes(r *gin.RouterGroup, pool *pgxpool.Pool, cfg config.AuthConfig) {
	q := repository.New(pool)

	svc := NewAuthService(q, cfg.JWTSecret)
	h := NewAuthHandler(svc)

	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)
	r.POST("/auth/refresh", h.Refresh)
	r.POST("/auth/logout", h.Logout)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var in RegisterInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Register(c.Request.Context(), in); err != nil {
		switch err {
		case errs.ErrEmailAlreadyExists, errs.ErrUsernameTaken:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "account created successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var in LoginInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.svc.Login(c.Request.Context(), in)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	cookieutil.SetRefreshCookie(c, refreshToken, RefreshTokenExpiryHours)
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
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
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie(cookieutil.RefreshCookieName)
	if err == nil {
		_ = h.svc.RevokeRefreshToken(c.Request.Context(), cookie.Value)
	}

	cookieutil.ClearRefreshTokenCookie(c)
	c.Status(http.StatusNoContent)
}
