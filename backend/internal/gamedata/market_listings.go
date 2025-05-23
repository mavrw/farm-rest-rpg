package gamedata

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

var MarketCatalogDefinitions = map[int32]repository.CreateMarketListingParams{
	Seed_01: {
		BuyPrice:  Int32Ptr(2),
		SellPrice: Int32Ptr(1),
	},
	Crop_01: {
		SellPrice: Int32Ptr(6),
	},
}
