package go_benchpress

import (
	"encoding/xml"
	"golang.org/x/tools/benchmark/parse"
	"io"
)

type XMLRenderer struct {

}

type xmlBenchmarkRecord struct {
	ParentBenchmark string
	Benchmarks []parse.Benchmark
}

func (x *XMLRenderer) Render(writer io.Writer, parentBenchmark string, dimension RenderDimension, benchmarks []parse.Benchmark) error {

	record := xmlBenchmarkRecord{
		ParentBenchmark: parentBenchmark,
		Benchmarks:      benchmarks,
	}

	data, err := xml.Marshal(record)
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	return err
}



