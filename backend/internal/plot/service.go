package plot

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/db"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
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

	// ! start transaciton
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer db.AutoRollbackTx(ctx, tx, &err)

	qtx := s.q.WithTx(tx)

	// TODO: Check that the user has enough money to buy the new plot

	plot, err := qtx.CreatePlot(ctx, farm.ID)
	if err != nil {
		return nil, err
	}

	// TODO: Deduct cost of plot from the user's money

	// ! commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

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

	// TODO: verify that user has seeds in inventory

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
	crop, err := qtx.GetCropByID(ctx, cropID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrCropNotFound
	}

	// sow plot with crop information
	plot, err = qtx.SowPlotByID(ctx, repository.SowPlotByIDParams{
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

	// TODO: deduct seeds from user inventory

	// ! commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &plot, nil
}

func (s *PlotService) HarvestPlot(ctx context.Context, userID, plotID int32) (*repository.Plot, error) {
	// get plot by ID
	plot, err := s.q.GetPlotByID(ctx, plotID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrPlotNotFound
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

	// check that plot is ready for harvest
	if time.Until(plot.HarvestAt) > 0 {
		return &plot, errs.ErrPlotNotFullyGrown
	}

	// TODO: Add harvested crops to player inventory

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
