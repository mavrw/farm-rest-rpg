package db_seed

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

const (
	Minute = 60
	Hour   = Minute * 60
	Day    = Hour * 24
)

var InitialCrops = []repository.CreateCropParams{
	{
		Name:              "Carrot",
		GrowthTimeSeconds: 1 * Minute,
		YieldAmount:       1,
	},
	{
		Name:              "Potato",
		GrowthTimeSeconds: 3 * Minute,
		YieldAmount:       4,
	},
	{
		Name:              "Tomato",
		GrowthTimeSeconds: 4 * Minute,
		YieldAmount:       6,
	},
	{
		Name:              "Corn",
		GrowthTimeSeconds: 6 * Minute,
		YieldAmount:       8,
	},
	{
		Name:              "Soy Bean",
		GrowthTimeSeconds: 6 * Minute,
		YieldAmount:       12,
	},
}
