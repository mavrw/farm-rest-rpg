package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/mavrw/farm-rest-rpg/backend/config"
	"github.com/mavrw/farm-rest-rpg/backend/internal/auth"
	"github.com/mavrw/farm-rest-rpg/backend/internal/db"
	"github.com/mavrw/farm-rest-rpg/backend/internal/farm"
	"github.com/mavrw/farm-rest-rpg/backend/internal/inventory"
	"github.com/mavrw/farm-rest-rpg/backend/internal/plot"
	"github.com/mavrw/farm-rest-rpg/backend/internal/user"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/middleware"
)

func main() {
	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	// connect to databsae
	dbPool, err := db.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("database load error: %v", err)
	}
	defer dbPool.Close()

	// create router and register general middleware
	router := gin.Default()
	router.Use(middleware.RequestLogger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// --- PUBLIC ROUTES ---
	public := router.Group("/api/v1")
	auth.RegisterRoutes(public, dbPool, cfg.Auth)

	// --- PROTECTED ROUTES ---
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret, dbPool))
	protected.Use(middleware.RLS()) // injects current_user_id into context

	// TODO: Make routes for REST conformant
	user.RegisterRoutes(protected, dbPool)
	farm.RegisterRoutes(protected, dbPool)
	plot.RegisterRoutes(protected, dbPool)
	inventory.RegisterRoutes(protected, dbPool)

	// start the server
	addr := ":" + cfg.Server.Port
	router.Run(addr)
}
