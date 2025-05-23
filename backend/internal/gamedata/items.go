package gamedata

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

const (
	Seed_01 int32 = iota + 1
	Crop_01
)

var ItemDefinitions = map[int32]repository.CreateItemDefinitionParams{
	Seed_01: {
		Name:        "seed_01",
		Description: "this seed is a test seed to test with while seeding test seeds",
		Rarity:      repository.ItemRarityCommon,
		Type:        SeedItem,
	},
	Crop_01: {
		Name:        "crop_01",
		Description: "this crop is a test crop to test with while seeding test crops",
		Rarity:      repository.ItemRarityCommon,
		Type:        CropItem,
	},
}
