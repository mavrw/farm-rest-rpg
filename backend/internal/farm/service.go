package farm

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
)

type FarmService struct {
	q *repository.Queries
}

func NewFarmService(q *repository.Queries) *FarmService {
	return &FarmService{q: q}
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

	// TODO: Move these operations into a transaction
	// !	 - this way if something fails when adding starter plots
	// !	 - to the user's farm, they don't accidentally end up with
	// !	 - a plot with less starter plots than intended, or with
	// !	 - no plots at all.
	params := repository.CreateFarmParams{
		UserID: userID,
		Name:   in.Name,
	}
	farm, err := s.q.CreateFarm(ctx, params)
	if err != nil {
		return nil, err
	}

	// TODO: Give the user 4 to 6 free starter plots when their farm is created

	return &farm, nil

}
