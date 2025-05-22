package inventory

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mavrw/farm-rest-rpg/backend/internal/repository"
	"github.com/mavrw/farm-rest-rpg/backend/pkg/errs"
)

type InventoryService struct {
	q *repository.Queries
}

func NewInventoryService(q *repository.Queries) *InventoryService {
	return &InventoryService{q: q}
}

func (s *InventoryService) GetItem(ctx context.Context, userId, itemId int32) (*repository.Inventory, error) {
	item, err := s.q.GetItem(ctx, repository.GetItemParams{UserID: userId, ItemID: itemId})
	if err == pgx.ErrNoRows {
		return nil, errs.ErrInventoryItemNotFound
	} else if err != nil {
		return nil, err
	}

	// this shouldn't happen, but I'm going to explicitly check anyways
	if item.UserID != userId {
		return nil, errs.ErrInventoryItemNotOwned
	}

	return &item, nil
}

func (s *InventoryService) ListItems(ctx context.Context, userId int32) (*[]repository.Inventory, error) {
	items, err := s.q.ListItems(ctx, userId)
	if err == pgx.ErrNoRows {
		return nil, errs.ErrInventoryEmpty
	} else if err != nil {
		return nil, err
	}

	return &items, nil
}
