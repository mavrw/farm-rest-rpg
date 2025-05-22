package main

import (
	"context"
	"log"
	"time"

	"github.com/mavrw/farm-rest-rpg/backend/config"
	"github.com/mavrw/farm-rest-rpg/backend/internal/db"
	"github.com/mavrw/farm-rest-rpg/backend/internal/db_seed"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	pool, err := db.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("database load error: %v", err)
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q := repository.New(pool)

	// seed crop data
	for idx, crop := range db_seed.InitialCrops {
		crop.ID = int32(idx + 1)
		if err := q.CreateCrop(ctx, crop); err != nil {
			log.Printf("failed to seed crop %+v: %v\n", crop, err)
		}
	}
	log.Println("Crops seeded successfully")

	// seed currency types
	for idx, currency := range db_seed.InitialCurrencyTypes {
		currency.ID = int32(idx + 1)
		if err := q.CreateCurrencyType(ctx, currency); err != nil {
			log.Printf("failed to seed currency type %+v: %v\n", currency, err)
		}
	}
	log.Println("Currency Types seeded successfully")
}
