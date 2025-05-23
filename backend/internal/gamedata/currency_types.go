package gamedata

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

const (
	GoldCurrency int32 = iota + 1
)

var CurrencyTypeDefinitions = map[int32]repository.CreateCurrencyTypeParams{
	GoldCurrency: {
		Name: "Gold",
	},
}
