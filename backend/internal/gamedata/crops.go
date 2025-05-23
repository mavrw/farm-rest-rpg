package gamedata

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

const (
	Minute = 60
	Hour   = Minute * 60
	Day    = Hour * 24
)

const (
	CarrotCrop int32 = iota + 1
	PotatoCrop
	TomatoCrop
	CornCrop
	SoyBeanCrop
)

var CropDefinitions = map[int32]repository.CreateCropDefinitionParams{
	CarrotCrop: {
		Name:              "Carrot",
		GrowthTimeSeconds: 1 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       1,
	},
	PotatoCrop: {
		Name:              "Potato",
		GrowthTimeSeconds: 3 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       4,
	},
	TomatoCrop: {
		Name:              "Tomato",
		GrowthTimeSeconds: 4 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       6,
	},
	CornCrop: {
		Name:              "Corn",
		GrowthTimeSeconds: 6 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       8,
	},
	SoyBeanCrop: {
		Name:              "Soy Bean",
		GrowthTimeSeconds: 6 * Minute,
		SeedItemID:        1,
		YieldItemID:       2,
		YieldAmount:       12,
	},
}
