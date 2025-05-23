package gamedata

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

var ItemDefinitions = []repository.CreateItemDefinitionParams{
	{
		Name:        "seed_01",
		Description: "this seed is a test seed to test with while seeding test seeds",
		Rarity:      repository.ItemRarityCommon,
		Type:        1,
	},
	{
		Name:        "crop_01",
		Description: "this crop is a test crop to test with while seeding test crops",
		Rarity:      repository.ItemRarityCommon,
		Type:        2,
	},
}
