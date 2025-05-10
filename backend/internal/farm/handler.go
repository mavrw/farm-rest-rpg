package farm

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/middleware"
)

type FarmHandler struct {
	svc *FarmService
}

func NewFarmHandler(svc *FarmService) *FarmHandler {
	return &FarmHandler{svc: svc}
}

func RegisterRoutes(r *gin.RouterGroup, pool *pgxpool.Pool) {
	q := repository.New(pool)

	svc := NewFarmService(q)
	h := NewFarmHandler(svc)

	r.GET("/farm/get", h.GetFarm)
	r.POST("/farm/create", h.CreateFarm)
}

func (h *FarmHandler) GetFarm(c *gin.Context) {
	userId := c.GetInt(middleware.ContextUserIDKey)

	farm, err := h.svc.Get(c.Request.Context(), int32(userId))
	if err != nil {
		if err == ErrFarmNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, farm)
}

func (h *FarmHandler) CreateFarm(c *gin.Context) {
	userId := c.GetInt(middleware.ContextUserIDKey)
	var in CreateFarmInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	farm, err := h.svc.Create(c.Request.Context(), int32(userId), in)
	if err != nil {
		if err == ErrFarmAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, farm)
}
