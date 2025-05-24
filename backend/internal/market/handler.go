package market

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/middleware"
)

type MarketHandler struct {
	svc *MarketService
}

func NewMarketHandler(svc *MarketService) *MarketHandler {
	return &MarketHandler{svc: svc}
}

func RegisterRoutes(r *gin.RouterGroup, pool *pgxpool.Pool) {
	q := repository.New(pool)

	svc := NewMarketService(q, pool)
	h := NewMarketHandler(svc)

	grp := r.Group("/market")
	grp.GET("/listing/:item_id", h.GetMarketListing)
	grp.GET("/listing/all", h.ListMarketListings)
	grp.POST("/listing/:item_id/buy/:buy_quantity", h.BuyMarketListing)
	grp.POST("/listing/:item_id/sell/:sell_quantity", h.SellMarketListing)
}

func (h *MarketHandler) GetMarketListing(c *gin.Context) {
	itemIDstr := c.Param("item_id")

	itemID, err := strconv.Atoi(itemIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item_id"})
		return
	}

	listing, err := h.svc.GetMarketListing(c.Request.Context(), int32(itemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, listing)
}

func (h *MarketHandler) ListMarketListings(c *gin.Context) {
	listings, err := h.svc.ListMarketListings(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, listings)
}

func (h *MarketHandler) BuyMarketListing(c *gin.Context) {
	itemIDstr := c.Param("item_id")
	buyQtystr := c.Param("buy_quantity")

	userID := c.GetInt(middleware.ContextUserIDKey)

	itemID, err := strconv.Atoi(itemIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item_id"})
		return
	}

	buyQty, err := strconv.Atoi(buyQtystr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid buy quantity"})
		return
	}

	marketTXResult, err := h.svc.BuyMarketListing(c.Request.Context(), int32(userID), int32(itemID), int32(buyQty))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, marketTXResult)
}

func (h *MarketHandler) SellMarketListing(c *gin.Context) {
	itemIDstr := c.Param("item_id")
	sellQtystr := c.Param("sell_quantity")

	userID := c.GetInt(middleware.ContextUserIDKey)

	itemID, err := strconv.Atoi(itemIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item_id"})
		return
	}

	sellQty, err := strconv.Atoi(sellQtystr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid buy quantity"})
		return
	}

	marketTXResult, err := h.svc.SellMarketListing(c.Request.Context(), int32(userID), int32(itemID), int32(sellQty))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, marketTXResult)
}
