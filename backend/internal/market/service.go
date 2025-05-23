package market

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/currency"
	"github.com/mavrw/farm-rest-rpg/backend/internal/db"
	"github.com/mavrw/farm-rest-rpg/backend/internal/gamedata"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
)

type MarketService struct {
	q      *repository.Queries
	dbPool *pgxpool.Pool
}

func NewMarketService(q *repository.Queries, pool *pgxpool.Pool) *MarketService {
	return &MarketService{
		q:      q,
		dbPool: pool,
	}
}

func (s *MarketService) GetMarketListing(ctx context.Context, itemID int32) (*repository.MarketListing, error) {
	// get market listing from database
	listing, err := s.q.GetMarketListing(ctx, itemID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrListingNotFound
	} else if err != nil {
		return nil, err
	}

	return &listing, nil
}

func (s *MarketService) ListMarketListings(ctx context.Context) (*[]repository.MarketListing, error) {
	// get all market listings from
	listings, err := s.q.ListMarketListings(ctx)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrListingNotFound
	} else if err != nil {
		return nil, err
	}

	return &listings, nil
}

func (s *MarketService) BuyMarketListing(ctx context.Context, userID, itemID, quantity int32) (*MarketTransactionResult, error) {
	// perform checks to confirm that the transaction is possible (item purchasable, enough gold, etc)
	// ? Is the quantity valid ?
	if quantity <= 0 {
		return nil, errs.ErrInvalidItemQuantity
	}

	// ? Does listing exist ?
	buyPrice, err := s.q.GetMarketListingBuyPrice(ctx, itemID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrListingNotFound
	} else if err != nil {
		return nil, err
	}

	// ? Can listing be purchased / does it have a buy price ?
	if buyPrice == nil {
		return nil, errs.ErrItemNotPurchasable
	}

	// TODO: Support multiple currencies, or defining required currency in listing
	userBalParams := repository.GetUserCurrencyBalanceByTypeParams{
		UserID:         userID,
		CurrencyTypeID: gamedata.GoldCurrency,
	}
	playerBalance, err := s.q.GetUserCurrencyBalanceByType(ctx, userBalParams)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrNoBalanceFound
	} else if err != nil {
		return nil, err
	}

	// ? Does the player have a sufficient balance ?
	totalCost := *buyPrice * quantity
	if playerBalance.Balance < totalCost {
		return nil, errs.ErrInsufficientBalance
	}

	// ! start transaction
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer db.AutoRollbackTx(ctx, tx, &err)

	qtx := s.q.WithTx(tx)

	// ! Deduct currency
	adjustBalParams := repository.AdjustUserCurrencyBalanceByTypeParams{
		UserID:         userID,
		Amount:         currency.Debit(totalCost),
		CurrencyTypeID: gamedata.GoldCurrency,
	}
	newBalance, err := qtx.AdjustUserCurrencyBalanceByType(ctx, adjustBalParams)
	if err != nil {
		return nil, err
	}

	// ! Add inventory items
	upsertInvParams := repository.UpsertInventoryItemParams{
		UserID:   userID,
		ItemID:   itemID,
		Quantity: quantity,
	}
	newInventoryItem, err := qtx.UpsertInventoryItem(ctx, upsertInvParams)
	if err != nil {
		return nil, err
	}

	// ! commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	txResult := MarketTransactionResult{
		ItemID:        itemID,
		Quantity:      quantity,
		TotalCost:     totalCost,
		NewBalance:    &newBalance,
		InventoryItem: &newInventoryItem,
	}
	return &txResult, nil
}

func (s *MarketService) SellMarketListing(ctx context.Context, userID, itemID, quantity int32) (*MarketTransactionResult, error) {
	// perform checks to confirm that the transaction is possible (item sellable, enough of item, etc)
	// ? Is the quantity valid ?
	if quantity <= 0 {
		return nil, errs.ErrInvalidItemQuantity
	}

	// ? Does listing exist ?
	sellPrice, err := s.q.GetMarketListingSellPrice(ctx, itemID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrListingNotFound
	} else if err != nil {
		return nil, err
	}

	// ? Can listing be sold / does it have a sell price ?
	if sellPrice == nil {
		return nil, errs.ErrItemNotSellable
	}

	userItemQtyParams := repository.GetInventoryItemParams{
		UserID: userID,
		ItemID: itemID,
	}
	playerInvQuantity, err := s.q.GetInventoryItem(ctx, userItemQtyParams)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrInventoryItemNotFound
	} else if err != nil {
		return nil, err
	}

	// ? Does the player have a sufficient quantity of the item ?
	if playerInvQuantity.Quantity < quantity {
		return nil, errs.ErrInsufficientItemQuantity
	}

	// ! start transaction
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer db.AutoRollbackTx(ctx, tx, &err)

	qtx := s.q.WithTx(tx)

	// ! Deduct items from inventory
	newInvQty := playerInvQuantity.Quantity - quantity
	setInvItemParams := repository.SetInventoryItemQuantityParams{
		UserID:   userID,
		ItemID:   itemID,
		Quantity: newInvQty,
	}
	newInventoryItem, err := qtx.SetInventoryItemQuantity(ctx, setInvItemParams)
	if err != nil {
		return nil, err
	}

	// ! Add total sale price to player balance
	creditAmt := *sellPrice * quantity
	adjBalanceParams := repository.AdjustUserCurrencyBalanceByTypeParams{
		UserID:         userID,
		Amount:         currency.Credit(creditAmt),
		CurrencyTypeID: gamedata.GoldCurrency,
	}
	newBalance, err := qtx.AdjustUserCurrencyBalanceByType(ctx, adjBalanceParams)
	if err != nil {
		return nil, err
	}

	// ! commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	txResult := MarketTransactionResult{
		ItemID:        itemID,
		Quantity:      quantity,
		TotalCost:     creditAmt,
		NewBalance:    &newBalance,
		InventoryItem: &newInventoryItem,
	}

	return &txResult, nil
}
