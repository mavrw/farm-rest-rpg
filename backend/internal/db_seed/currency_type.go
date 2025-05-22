package db_seed

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

var InitialCurrencyTypes = []repository.CreateCurrencyTypeParams{
	{
		Name: "Gold",
	},
}
