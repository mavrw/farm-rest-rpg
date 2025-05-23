package gamedata

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

const (
	SeedItem int32 = iota + 1
	CropItem
)

var ItemTypeDefinitions = map[int32]repository.CreateItemTypeParams{
	SeedItem: {
		Name: "seed",
	},
	CropItem: {
		Name: "crop",
	},
}
