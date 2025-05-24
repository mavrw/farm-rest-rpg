package farm

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mavrw/farm-rest-rpg/backend/internal/db"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
)

const (
	numStarterPlots = 6
)

type FarmService struct {
	q      *repository.Queries
	dbPool *pgxpool.Pool
}

func NewFarmService(q *repository.Queries, pool *pgxpool.Pool) *FarmService {
	return &FarmService{
		q:      q,
		dbPool: pool,
	}
}

func (s *FarmService) Get(ctx context.Context, userId int32) (*repository.Farm, error) {
	farm, err := s.q.GetFarmByUserID(ctx, userId)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrFarmNotFound
	}

	return &farm, err
}

func (s *FarmService) Create(ctx context.Context, userID int32, in CreateFarmInput) (*repository.Farm, error) {
	_, err := s.q.GetFarmByUserID(ctx, userID)
	if err == nil {
		return nil, errs.ErrFarmAlreadyExists
	}
	if err != pgx.ErrNoRows {
		return nil, err
	}

	// TODONE: Move these operations into a transaction
	// !	 - this way if something fails when adding starter plots
	// !	 - to the user's farm, they don't accidentally end up with
	// !	 - a plot with less starter plots than intended, or with
	// !	 - no plots at all.

	// ! start transaciton
	tx, err := s.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer db.AutoRollbackTx(ctx, tx, &err)

	qtx := s.q.WithTx(tx)

	// Create farm
	params := repository.CreateFarmParams{
		UserID: userID,
		Name:   in.Name,
	}
	farm, err := qtx.CreateFarm(ctx, params)
	if err != nil {
		return nil, err
	}

	// ! Quick sanity check to make sure the user doesn't somehow
	// ! - already have plots
	plots, err := s.q.GetPlotsByFarmID(ctx, farm.ID)
	if len(plots) > 0 {
		return nil, errs.ErrFarmAlreadyHasPlots
	}

	// Give the user free starter plots when their farm is created
	for range numStarterPlots {
		_, err := qtx.CreatePlot(ctx, farm.ID)
		if err != nil {
			// failed to create a starter plot, roll everything back
			// so farm creation can be tried again
			return nil, errs.ErrPlotCreationFailed
		}
	}

	// ? Something is stinking here, but I can't quite pin down
	// ? quite what it is yet...

	// ! commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &farm, nil

}
