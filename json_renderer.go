package go_benchpress

import (
	"encoding/json"
	"golang.org/x/tools/benchmark/parse"
	"io"
)

type JSONRenderer struct {

}

type benchmarksJSON struct {
	ParentBenchmark string
	Benchmarks []parse.Benchmark
}

func (j *JSONRenderer) Render(writer io.Writer, parentBenchmark string, dimension RenderDimension, benchmarks []parse.Benchmark) error {
	b := benchmarksJSON{
		ParentBenchmark: parentBenchmark,
		Benchmarks:      benchmarks,
	}
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}
