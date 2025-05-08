package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/mavrw/farm-rest-rpg/backend/config"
	"github.com/mavrw/farm-rest-rpg/backend/internal/auth"
	"github.com/mavrw/farm-rest-rpg/backend/internal/db"
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

	// create router and register middleware
	router := gin.Default()
	router.Use(middleware.RequestLogger())
	// router.Use(middleware.CORSMiddleWare())
	// router.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
	router.Use(middleware.RLS()) // injects current_user_id into context

	// register routes and pass deps
	auth.RegisterRoutes(router, dbPool, cfg.Auth)

	// start the server
	addr := ":" + cfg.Server.Port
	router.Run(addr)
}
