package errs

import "errors"

var (
	ErrFarmNotFound       = errors.New("farm not found")
	ErrFarmAlreadyExists  = errors.New("farm already exists")
	ErrFarmNotOwnedByUser = errors.New("farm not owned by user")

	ErrPlotNotFound       = errors.New("plot(s) not found")
	ErrPlotNotOwnedByUser = errors.New("plot not owned by user")
	ErrPlotAlreadyPlanted = errors.New("plot already planted")
	ErrPlotNotFullyGrown  = errors.New("plot not ready for harvest")

	ErrCropNotFound = errors.New("crop data not found")
)
