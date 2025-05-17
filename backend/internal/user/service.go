package user

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
)

type UserService struct {
	q *repository.Queries
}

func NewUserService(q *repository.Queries) *UserService {
	return &UserService{q: q}
}

func (s *UserService) GetMe(ctx context.Context, userID int32) (*repository.User, error) {
	user, err := s.q.GetUserByID(ctx, userID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) UpdateMe(ctx context.Context, userID int32) (*repository.User, error) {
	return nil, errs.ErrNotImplemented
}

func (s *UserService) DeleteMe(ctx context.Context, userID int32) error {
	return errs.ErrNotImplemented
}
