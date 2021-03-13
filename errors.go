package benchpress

import "errors"

var (
	ErrNoBenchmarksProvided    = errors.New("could not render benchmarks - no benchmarks provided")
	ErrUnknownRasterRenderType = errors.New("unknown raster render type")
)
