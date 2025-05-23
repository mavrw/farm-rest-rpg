package gamedata

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

const (
	Minute = 60
	Hour   = Minute * 60
	Day    = Hour * 24
)

// TODO: Figure out how to introduce iota to make these definitions more clear
//       - Perhaps define iota with definition name, then use that iota
//		 - value as a map key for the definition's data

var CropDefinitions = []repository.CreateCropDefinitionParams{
	{
		Name:              "Carrot",
		GrowthTimeSeconds: 1 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       1,
	},
	{
		Name:              "Potato",
		GrowthTimeSeconds: 3 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       4,
	},
	{
		Name:              "Tomato",
		GrowthTimeSeconds: 4 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       6,
	},
	{
		Name:              "Corn",
		GrowthTimeSeconds: 6 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       8,
	},
	{
		Name:              "Soy Bean",
		GrowthTimeSeconds: 6 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       12,
	},
}
