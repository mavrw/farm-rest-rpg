package market

import "github.com/mavrw/farm-rest-rpg/backend/internal/repository"

// MarketTransactionResult captures the result of a buy/sell transaction in the market.
// It includes the item transacted, quantity, total cost, updated currency balance, and inventory state.
type MarketTransactionResult struct {
	ItemID        int32                       `json:"item_id"`
	Quantity      int32                       `json:"quantity"`
	TotalCost     int32                       `json:"total_cost"`
	NewBalance    *repository.CurrencyBalance `json:"new_balance"`
	InventoryItem *repository.Inventory       `json:"inventory_item"`
}
