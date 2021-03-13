package benchpress

import (
	"golang.org/x/tools/benchmark/parse"
	"io"
)

type Renderer interface {
	Render(io.Writer, []parse.Benchmark) error
}
