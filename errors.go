package go_benchpress

import "errors"

var (
	ErrNoBenchmarksProvided    = errors.New("could not render benchmarks - no benchmarks provided")
	ErrUnknownRasterRenderType = errors.New("unknown raster render type")
	ErrCouldNotParseLine       = errors.New("could not parse benchmark line")
	ErrUnknownDimensionType    = errors.New("unknown render dimension type")
)
