package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mavrw/farm-rest-rpg/backend/config"
	"github.com/mavrw/farm-rest-rpg/backend/internal/db"
	"github.com/mavrw/farm-rest-rpg/backend/internal/gamedata"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
)

type SeederFunc func(ctx context.Context, q *repository.Queries)

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

	// TODO: Add consistency check to determine if any section can be skipped

	seedingOrder := []SeederFunc{
		seedItemTypeDefinitions,
		seedItemDefinitions,
		seedCurrencyTypeDefinitions,
		seedCropDefinitions,
		seedMarketCatalogDefinitions,
	}

	for _, step := range seedingOrder {
		step(ctx, q)
	}
}

// TODO: Figure out a way to de-duplicate some of the logic in the seeder funcs
//		- generics over-complicate things, and the reflection removes
//		- type safety. A happy medium can probably be reached with
//		- some helper functions for error checking and logging.

func seedCropDefinitions(ctx context.Context, q *repository.Queries) {
	for idx, crop := range gamedata.CropDefinitions {
		crop.ID = idx
		_, err := q.CreateCropDefinition(ctx, crop)
		if err == pgx.ErrNoRows {
			log.Printf("crop data already exists for crop: %s (id: %d)\n", crop.Name, crop.ID)
		} else if err != nil {
			log.Printf("failed to seed crop %+v: %v\n", crop, err)
		}
	}
	log.Println("Crop Definitions seeded successfully")
}

func seedCurrencyTypeDefinitions(ctx context.Context, q *repository.Queries) {
	for idx, currency := range gamedata.CurrencyTypeDefinitions {
		currency.ID = idx
		_, err := q.CreateCurrencyType(ctx, currency)
		if err == pgx.ErrNoRows {
			log.Printf("currency type already exists for currency: %s (id: %d)\n", currency.Name, currency.ID)
		} else if err != nil {
			log.Printf("failed to seed currency type %+v: %v\n", currency, err)
		}
	}
	log.Println("Currency Types seeded successfully")
}

func seedItemTypeDefinitions(ctx context.Context, q *repository.Queries) {
	for idx, itemType := range gamedata.ItemTypeDefinitions {
		itemType.ID = idx
		_, err := q.CreateItemType(ctx, itemType)
		if err == pgx.ErrNoRows {
			log.Printf("item type data already exists for item type: %s (id: %d)\n", itemType.Name, itemType.ID)
		} else if err != nil {
			log.Printf("failed to seed item type %+v: %v\n", itemType, err)
		}
	}
	log.Println("Item Types seeded successfully")
}

func seedItemDefinitions(ctx context.Context, q *repository.Queries) {
	for idx, item := range gamedata.ItemDefinitions {
		item.ID = idx
		_, err := q.CreateItemDefinition(ctx, item)
		if err == pgx.ErrNoRows {
			log.Printf("item data already exists for currency: %s (id: %d)\n", item.Name, item.ID)
		} else if err != nil {
			log.Printf("failed to seed item definition %+v: %v\n", item, err)
		}
	}
	log.Println("Item Definitions seeded successfully")
}

func seedMarketCatalogDefinitions(ctx context.Context, q *repository.Queries) {
	for itemID, catalogItem := range gamedata.MarketCatalogDefinitions {
		catalogItem.ItemID = itemID
		_, err := q.CreateMarketItem(ctx, catalogItem)
		if err == pgx.ErrNoRows {
			log.Printf("market catalog item data already exists for currency: (item_id: %d)\n", catalogItem.ItemID)
		} else if err != nil {
			log.Printf("failed to seed market catalog item definition %+v: %v\n", catalogItem, err)
		}
	}
	log.Println("Market Catalog Item Definitions seeded successfully")
}
