package errs

import "errors"

var (
	ErrNotImplemented = errors.New("action not implemented")

	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrEmailAlreadyExists  = errors.New("email already registered")
	ErrUsernameTaken       = errors.New("username is taken")
	ErrTokenAlreadyRevoked = errors.New("token already revoked")

	ErrFarmNotFound        = errors.New("farm not found")
	ErrFarmAlreadyExists   = errors.New("farm already exists")
	ErrFarmNotOwnedByUser  = errors.New("farm not owned by user")
	ErrFarmAlreadyHasPlots = errors.New("farm already has plots")
	ErrFarmHasNoPlots      = errors.New("farm has no plots")

	ErrPlotNotFound       = errors.New("plot(s) not found")
	ErrPlotCreationFailed = errors.New("error creating plot")
	ErrPlotNotOwnedByUser = errors.New("plot not owned by user")
	ErrPlotAlreadyPlanted = errors.New("plot already planted")
	ErrPlotNotFullyGrown  = errors.New("plot not ready for harvest")
	ErrCannotAffordPlot   = errors.New("user cannot afford new plot")

	ErrCropNotFound = errors.New("crop data not found")

	ErrInventoryItemNotFound    = errors.New("inventory item not found")
	ErrInventoryItemNotOwned    = errors.New("inventory item not owned")
	ErrInventoryEmpty           = errors.New("inventory empty")
	ErrInsufficientItemQuantity = errors.New("insufficient item quantity")
	ErrInvalidItemQuantity      = errors.New("invalid item quantity")

	ErrNoBalanceFound      = errors.New("balance(s) not found")
	ErrBalanceNotOwned     = errors.New("balance not owned")
	ErrInsufficientBalance = errors.New("insufficient balance")

	ErrListingNotFound    = errors.New("no listing(s) found")
	ErrItemNotPurchasable = errors.New("item cannot be purchased")
	ErrItemNotSellable    = errors.New("item cannot be sold")
)
