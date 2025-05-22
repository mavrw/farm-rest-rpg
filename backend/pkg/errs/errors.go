package errs

import "errors"

var (
	ErrNotImplemented = errors.New("action not implemented")

	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrEmailAlreadyExists  = errors.New("email already registered")
	ErrUsernameTaken       = errors.New("username is taken")
	ErrTokenAlreadyRevoked = errors.New("token already revoked")

	ErrFarmNotFound       = errors.New("farm not found")
	ErrFarmAlreadyExists  = errors.New("farm already exists")
	ErrFarmNotOwnedByUser = errors.New("farm not owned by user")

	ErrPlotNotFound       = errors.New("plot(s) not found")
	ErrPlotNotOwnedByUser = errors.New("plot not owned by user")
	ErrPlotAlreadyPlanted = errors.New("plot already planted")
	ErrPlotNotFullyGrown  = errors.New("plot not ready for harvest")

	ErrCropNotFound = errors.New("crop data not found")

	ErrInventoryItemNotFound = errors.New("inventory item not found")
	ErrInventoryItemNotOwned = errors.New("inventory item not owned")
	ErrInventoryEmpty        = errors.New("inventory empty")
)
