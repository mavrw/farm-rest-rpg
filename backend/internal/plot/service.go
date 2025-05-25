package plot

import (
	"context"
	"math"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/currency"
	"github.com/mavrw/farm-rest-rpg/backend/internal/db"
	"github.com/mavrw/farm-rest-rpg/backend/internal/gamedata"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
)

const (
	NumStarterPlots = 6
	BasePlotCost    = 100.0
	CostGrowthRate  = 1.15
)

type PlotService struct {
	q      *repository.Queries
	dbPool *pgxpool.Pool
}

func NewPlotService(q *repository.Queries, pool *pgxpool.Pool) *PlotService {
	return &PlotService{
		q:      q,
		dbPool: pool,
	}
}

func (s *PlotService) BuyPlot(ctx context.Context, userID, farmID int32) (*repository.Plot, error) {
	farm, err := s.q.GetFarmByID(ctx, userID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrFarmNotFound
	} else if err != nil {
		return nil, err
	}

	if farm.UserID != userID {
		return nil, errs.ErrFarmNotOwnedByUser
	}

	// get the user's gold balance
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

	// get the number of plots owned by the user
	farmPlots, err := s.q.GetPlotsByFarmID(ctx, farm.ID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrFarmHasNoPlots
	} else if err != nil {
		return nil, err
	}
	plotsOwned := int32(len(farmPlots))

	// calculate the cost of the new plot
	newPlotCost := CalculatePlotCost(plotsOwned)

	// ? Does the player have a sufficient balance ?
	if playerBalance.Balance < newPlotCost {
		return nil, errs.ErrCannotAffordPlot
	}

	// ! start transaciton
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer db.AutoRollbackTx(ctx, tx, &err)

	qtx := s.q.WithTx(tx)

	plot, err := qtx.CreatePlot(ctx, farm.ID)
	if err != nil {
		return nil, err
	}

	adjustBalParams := repository.AdjustUserCurrencyBalanceByTypeParams{
		UserID:         userID,
		Amount:         currency.Debit(newPlotCost),
		CurrencyTypeID: gamedata.GoldCurrency,
	}
	_, err = qtx.AdjustUserCurrencyBalanceByType(ctx, adjustBalParams)
	if err != nil {
		return nil, err
	}

	// ! commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Create `PlotTransactionResult` struct and return that instead
	// !	- this would enable the logging and auditing of plot transactions

	return &plot, nil
}

func (s *PlotService) PlantPlot(ctx context.Context, userID, plotID, cropID int32) (*repository.Plot, error) {
	// Get plot by id
	plot, err := s.q.GetPlotByID(ctx, plotID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrPlotNotFound
	} else if err != nil {
		return nil, err
	}

	cropInfo, err := s.q.GetCropDefinition(ctx, cropID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrCropNotFound
	} else if err != nil {
		return nil, err
	}

	// get player's inventory for the crop's seed item
	userItemQtyParams := repository.GetInventoryItemParams{
		UserID: userID,
		ItemID: cropInfo.SeedItemID,
	}
	playerSeedItemInventory, err := s.q.GetInventoryItem(ctx, userItemQtyParams)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrInventoryItemNotFound
	} else if err != nil {
		return nil, err
	}

	// ? Does the player have enough seeds
	// TODO: Consider different seed quantities to plant various crops
	if playerSeedItemInventory.Quantity <= 0 {
		return nil, errs.ErrInsufficientSeedQuantity
	}

	// ! start transaciton
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer db.AutoRollbackTx(ctx, tx, &err)

	qtx := s.q.WithTx(tx)

	// TODO: Consider caching farmID at the handler layer
	// TODO: Consider 'Ownership' middleware for farm and plot ownership checks
	// TODO: Consider `assertPlotOwnership` helper func instead

	// check that plot belongs to user via farmID
	farm, err := qtx.GetFarmByUserID(ctx, userID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrFarmNotFound
	} else if err != nil {
		return nil, err
	}

	if plot.FarmID != farm.ID {
		return nil, errs.ErrPlotNotOwnedByUser
	}

	// verify that plot is currently empty
	if plot.CropID != nil {
		return &plot, errs.ErrPlotAlreadyPlanted
	}

	// fetch crop information by cropID
	crop, err := qtx.GetCropDefinition(ctx, cropID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrCropNotFound
	}

	// sow plot with crop information
	sownPlot, err := qtx.SowPlotByID(ctx, repository.SowPlotByIDParams{
		ID:        plot.ID,
		CropID:    &crop.ID,
		PlantedAt: time.Now(),
		HarvestAt: time.Now().Add(time.Duration(crop.GrowthTimeSeconds) * time.Second),
	})
	if err == pgx.ErrNoRows {
		// ! If this happens, something got really fucked up
		return nil, errs.ErrPlotNotFound
	} else if err != nil {
		return nil, err
	}

	setInvItemParams := repository.SetInventoryItemQuantityParams{
		UserID:   userID,
		ItemID:   cropInfo.SeedItemID,
		Quantity: playerSeedItemInventory.Quantity - 1,
	}
	_, err = qtx.SetInventoryItemQuantity(ctx, setInvItemParams)
	if err != nil {
		return nil, errs.ErrUpdatingInventoryItem
	}

	// ! commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &sownPlot, nil
}

func (s *PlotService) HarvestPlot(ctx context.Context, userID, plotID int32) (*repository.Plot, error) {
	// get plot by ID
	plot, err := s.q.GetPlotByID(ctx, plotID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrPlotNotFound
	} else if err != nil {
		return nil, err
	}

	// check that plot belongs to user via farmID
	farm, err := s.q.GetFarmByUserID(ctx, userID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrFarmNotFound
	} else if err != nil {
		return nil, err
	}

	if plot.FarmID != farm.ID {
		return nil, errs.ErrPlotNotOwnedByUser
	}

	// check that plot is ready for harvest
	if time.Until(plot.HarvestAt) > 0 {
		return &plot, errs.ErrPlotNotFullyGrown
	}

	// get sown crop info
	cropInfo, err := s.q.GetCropDefinition(ctx, *plot.CropID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrCropNotFound
	} else if err != nil {
		return nil, err
	}

	// ! start transaciton
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer db.AutoRollbackTx(ctx, tx, &err)

	qtx := s.q.WithTx(tx)

	// update player inventory for crop yield item
	upsertInvItemParams := repository.UpsertInventoryItemParams{
		UserID:   userID,
		ItemID:   cropInfo.YieldItemID,
		Quantity: cropInfo.YieldAmount,
	}
	_, err = qtx.UpsertInventoryItem(ctx, upsertInvItemParams)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrUpdatingInventoryItem
	} else if err != nil {
		return nil, err
	}

	// clear crop info
	plot, err = qtx.HarvestPlotByID(ctx, plot.ID)
	if err == pgx.ErrNoRows {
		// ! If this happens, something got really fucked up
		return nil, errs.ErrPlotNotFound
	} else if err != nil {
		return nil, err
	}

	// ! commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	// return the plot's new state
	return &plot, nil
}

func (s *PlotService) GetAllPlotStates(ctx context.Context, userID, farmID int32) (*[]repository.Plot, error) {
	// get the user's farm
	farm, err := s.q.GetFarmByID(ctx, farmID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrFarmNotFound
	} else if err != nil {
		return nil, err
	}

	if farm.UserID != userID {
		return nil, errs.ErrFarmNotOwnedByUser
	}

	// get all plots for farm_id
	plots, err := s.q.GetPlotsByFarmID(ctx, farm.ID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrPlotNotFound
	} else if err != nil {
		return nil, err
	}

	// return plots
	return &plots, nil
}

func (s *PlotService) GetPlotStateByID(ctx context.Context, userID, plotID int32) (*repository.Plot, error) {
	// get plot by ID
	plot, err := s.q.GetPlotByID(ctx, plotID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrPlotNotFound
	} else if err != nil {
		return nil, err
	}

	// check that plot belongs to user via farmID
	farm, err := s.q.GetFarmByUserID(ctx, userID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrFarmNotFound
	} else if err != nil {
		return nil, err
	}

	if plot.FarmID != farm.ID {
		return nil, errs.ErrPlotNotOwnedByUser
	}

	// return plot
	return &plot, nil
}

func CalculatePlotCost(plotsOwned int32) int32 {
	plotsPurchased := float64(plotsOwned - NumStarterPlots)
	if plotsPurchased < 0 {
		plotsPurchased = 0
	}

	cost := BasePlotCost * math.Pow(CostGrowthRate, plotsPurchased)
	return int32(cost)
}
