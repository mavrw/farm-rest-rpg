package currency

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/middleware"
)

type CurrencyHandler struct {
	svc *CurrencyService
}

func NewCurrencyHandler(svc *CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{svc: svc}
}

func RegisterRoutes(r *gin.RouterGroup, pool *pgxpool.Pool) {
	q := repository.New(pool)

	svc := NewCurrencyService(q)
	h := NewCurrencyHandler(svc)

	grp := r.Group("/currency")
	grp.GET("/balance/:currency_type_id", h.GetBalance)
	grp.GET("/balance/all", h.ListBalances)
}

func (h *CurrencyHandler) GetBalance(c *gin.Context) {
	currencyTypeIDstr := c.Param("currency_type_id")

	userID := c.GetInt(middleware.ContextUserIDKey)

	currencyTypeID, err := strconv.Atoi(currencyTypeIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid currency_type_id"})
		return
	}

	balance, err := h.svc.GetBalance(c.Request.Context(), int32(userID), int32(currencyTypeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, balance)
}

func (h *CurrencyHandler) ListBalances(c *gin.Context) {
	userID := c.GetInt(middleware.ContextUserIDKey)

	balances, err := h.svc.ListBalances(c.Request.Context(), int32(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, balances)
}
