package benchpress

import (
	"fmt"
	"golang.org/x/tools/benchmark/parse"
	"io"
)

type RenderDimension int

const (
	RenderNsPerOp RenderDimension = iota
	RenderBytesPerOp
	RenderAllocsPerOp
)

func (r RenderDimension) String() string {
	switch r {
	case RenderNsPerOp:
		return "Ns Per Op"
	case RenderBytesPerOp:
		return "Bytes Per Op"
	case RenderAllocsPerOp:
		return "Allocs Per Op"
	default:
		return fmt.Sprintf("Unknown (%d)", r)
	}
}

type Renderer interface {
	Render(writer io.Writer, parentBenchmark string, dimension RenderDimension, benchmarks []parse.Benchmark) error
}
