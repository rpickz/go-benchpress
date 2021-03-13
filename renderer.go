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

func RenderDimensionFromString(str string) (RenderDimension, error) {
	switch str {
	case "NS_PER_OP":
		return RenderNsPerOp, nil
	case "BYTES_PER_OP":
		return RenderBytesPerOp, nil
	case "ALLOCS_PER_OP":
		return RenderAllocsPerOp, nil
	default:
		return -1, fmt.Errorf("render dimension %q not supported: %w", str, ErrUnknownDimensionType)
	}
}

type Renderer interface {
	Render(writer io.Writer, parentBenchmark string, dimension RenderDimension, benchmarks []parse.Benchmark) error
}
