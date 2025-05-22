package inventory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/middleware"
)

type InventoryHandler struct {
	svc *InventoryService
}

func NewInventoryHandler(svc *InventoryService) *InventoryHandler {
	return &InventoryHandler{svc: svc}
}

func RegisterRoutes(r *gin.RouterGroup, pool *pgxpool.Pool) {
	q := repository.New(pool)

	svc := NewInventoryService(q)
	h := NewInventoryHandler(svc)

	grp := r.Group("/inventory")
	grp.GET("/items/:item_id", h.GetItem)
	grp.GET("/items/all", h.ListItems)
}

func (h *InventoryHandler) GetItem(c *gin.Context) {
	itemIDstr := c.Param("item_id")

	userID := c.GetInt(middleware.ContextUserIDKey)

	itemID, err := strconv.Atoi(itemIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item_id"})
		return
	}

	item, err := h.svc.GetItem(c.Request.Context(), int32(userID), int32(itemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *InventoryHandler) ListItems(c *gin.Context) {
	userID := c.GetInt(middleware.ContextUserIDKey)

	items, err := h.svc.ListItems(c.Request.Context(), int32(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}
