package plot

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/middleware"
)

type PlotHandler struct {
	svc *PlotService
}

func NewPlotHandler(svc *PlotService) *PlotHandler {
	return &PlotHandler{svc: svc}
}

func RegisterRoutes(r *gin.RouterGroup, pool *pgxpool.Pool) {
	q := repository.New(pool)

	svc := NewPlotService(q, pool)
	h := NewPlotHandler(svc)

	r.GET("/plots/:plot_id", h.GetPlot)
	r.GET("/farm/:farm_id/plots", h.GetAllPlots)
	r.POST("/farm/:farm_id/plots", h.BuyPlot)
	r.POST("/plots/:plot_id/plant/:crop_id", h.PlantPlot)
	r.POST("/plots/:plot_id/harvest", h.HarvestPlot)
}

func (h *PlotHandler) GetPlot(c *gin.Context) {
	plotIDstr := c.Param("plot_id")

	userId := c.GetInt(middleware.ContextUserIDKey)

	plotID, err := strconv.Atoi(plotIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plot_id"})
		return
	}

	plot, err := h.svc.GetPlotStateByID(c.Request.Context(), int32(userId), int32(plotID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plot)
}

func (h *PlotHandler) GetAllPlots(c *gin.Context) {
	farmIDstr := c.Param("farm_id")

	userId := c.GetInt(middleware.ContextUserIDKey)

	farmID, err := strconv.Atoi(farmIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid farm_id"})
		return
	}

	plots, err := h.svc.GetAllPlotStates(c.Request.Context(), int32(userId), int32(farmID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plots)
}

func (h *PlotHandler) BuyPlot(c *gin.Context) {
	farmIDstr := c.Param("farm_id")

	userId := c.GetInt(middleware.ContextUserIDKey)

	farmID, err := strconv.Atoi(farmIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid farm_id"})
		return
	}

	plot, err := h.svc.BuyPlot(c.Request.Context(), int32(userId), int32(farmID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plot)
}

func (h *PlotHandler) PlantPlot(c *gin.Context) {
	plotIDstr := c.Param("plot_id")
	cropIDstr := c.Param("crop_id")

	userId := c.GetInt(middleware.ContextUserIDKey)

	plotID, err := strconv.Atoi(plotIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plot_id"})
		return
	}
	cropID, err := strconv.Atoi(cropIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid crop_id"})
		return
	}

	plot, err := h.svc.PlantPlot(c.Request.Context(), int32(userId), int32(plotID), int32(cropID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plot)
}

func (h *PlotHandler) HarvestPlot(c *gin.Context) {
	plotIDstr := c.Param("plot_id")

	userId := c.GetInt(middleware.ContextUserIDKey)

	plotID, err := strconv.Atoi(plotIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plot_id"})
		return
	}

	plot, err := h.svc.HarvestPlot(c.Request.Context(), int32(userId), int32(plotID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plot)
}
