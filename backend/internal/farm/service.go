package farm

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
)

var (
	ErrFarmNotFound      = errors.New("farm not found")
	ErrFarmAlreadyExists = errors.New("farm already exists")
)

type FarmService struct {
	q *repository.Queries
}

func NewFarmService(q *repository.Queries) *FarmService {
	return &FarmService{q: q}
}

func (s *FarmService) Get(ctx context.Context, userId int32) (*repository.Farm, error) {
	farm, err := s.q.GetFarmByUserID(ctx, userId)
	if err == sql.ErrNoRows {
		return nil, ErrFarmNotFound
	}

	return &farm, err
}

func (s *FarmService) Create(ctx context.Context, userID int32, in CreateFarmInput) (*repository.Farm, error) {
	_, err := s.q.GetFarmByUserID(ctx, userID)
	if err == nil {
		return nil, ErrFarmAlreadyExists
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	params := repository.CreateFarmParams{
		UserID: userID,
		Name:   in.Name,
	}

	farm, err := s.q.CreateFarm(ctx, params)

	// Does this really need to return a pointer? WHich is better practice?
	return &farm, err

}
