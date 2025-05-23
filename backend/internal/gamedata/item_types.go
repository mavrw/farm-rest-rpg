package gamedata

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

var ItemTypeDefinitions = []repository.CreateItemTypeParams{
	{
		Name: "seed",
	},
	{
		Name: "crop",
	},
}
