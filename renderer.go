package benchpress

import (
	"golang.org/x/tools/benchmark/parse"
	"io"
)

type Renderer interface {
	Render(writer io.Writer, parentBenchmark string, benchmarks []parse.Benchmark) error
}
