// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package repository

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ItemRarity string

const (
	ItemRarityCommon    ItemRarity = "common"
	ItemRarityUncommon  ItemRarity = "uncommon"
	ItemRarityRare      ItemRarity = "rare"
	ItemRarityEpic      ItemRarity = "epic"
	ItemRarityLegendary ItemRarity = "legendary"
)

func (e *ItemRarity) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ItemRarity(s)
	case string:
		*e = ItemRarity(s)
	default:
		return fmt.Errorf("unsupported scan type for ItemRarity: %T", src)
	}
	return nil
}

type NullItemRarity struct {
	ItemRarity ItemRarity
	Valid      bool // Valid is true if ItemRarity is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullItemRarity) Scan(value interface{}) error {
	if value == nil {
		ns.ItemRarity, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ItemRarity.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullItemRarity) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ItemRarity), nil
}

type Crop struct {
	ID                int32
	Name              string
	GrowthTimeSeconds int32
	SeedItemID        int32
	YieldItemID       int32
	YieldAmount       int32
}

type CurrencyBalance struct {
	ID             int32
	UserID         int32
	CurrencyTypeID int32
	Balance        int32
}

type CurrencyType struct {
	ID   int32
	Name string
}

type Farm struct {
	ID        int32
	UserID    int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Inventory struct {
	ID       int32
	UserID   int32
	ItemID   int32
	Quantity int32
}

type Item struct {
	ID          int32
	Name        string
	Description string
	Rarity      ItemRarity
	Type        int32
	EffectJson  []byte
}

type ItemType struct {
	ID   int32
	Name string
}

type MarketListing struct {
	ItemID    int32
	BuyPrice  *int32
	SellPrice *int32
}

type Plot struct {
	ID        int32
	FarmID    int32
	CropID    *int32
	PlantedAt *time.Time
	HarvestAt *time.Time
}

type RefreshToken struct {
	ID        uuid.UUID
	UserID    int32
	Token     string
	ExpiresAt time.Time
	Revoked   *bool
	CreatedAt *time.Time
}

type User struct {
	ID           int32
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
