package currency

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
)

type CurrencyService struct {
	q *repository.Queries
}

func NewCurrencyService(q *repository.Queries) *CurrencyService {
	return &CurrencyService{q: q}
}

func (s *CurrencyService) GetBalance(ctx context.Context, userID, currencyTypeID int32) (*repository.CurrencyBalance, error) {
	params := repository.GetUserCurrencyBalanceByTypeParams{
		UserID:         userID,
		CurrencyTypeID: currencyTypeID,
	}
	balance, err := s.q.GetUserCurrencyBalanceByType(ctx, params)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrNoBalanceFound
	} else if err != nil {
		return nil, err
	}

	// double-check balance is owned by user (should always be the case)
	if balance.UserID != userID {
		return nil, errs.ErrBalanceNotOwned
	}

	return &balance, nil
}

func (s *CurrencyService) ListBalances(ctx context.Context, userID int32) (*[]repository.CurrencyBalance, error) {
	balances, err := s.q.ListUserCurrencyBalances(ctx, userID)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrNoBalanceFound
	} else if err != nil {
		return nil, err
	}

	return &balances, nil
}
