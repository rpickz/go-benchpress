package main

import (
	"go-benchpress/m/v2"
	"golang.org/x/tools/benchmark/parse"
	"log"
	"os"
)

func main() {

	file, err := os.Open("./output.txt")
	if err != nil {
		log.Fatalf("Could not open output.txt for reading - error: %v", err)
	}
	defer file.Close()

	separatedBenchmarks, err := benchpress.ReadAndSeparateBenchmarks(file)
	if err != nil {
		log.Fatalf("Could not read benchmarks from input - error: %v", err)
	}

	for name, benchmarks := range separatedBenchmarks {
		writeBenchmarks(name, benchmarks)
	}
}

func writeBenchmarks(name string, benchmarks []parse.Benchmark) {

	renderer := benchpress.NewRasterRenderer(name, benchpress.SVG)

	file, err := os.Create(name + ".svg")
	if err != nil {
		log.Fatalf("Could not open file for writing - error: %v", err)
	}
	defer file.Close()

	err = renderer.Render(file, name, benchpress.RenderBytesPerOp, benchmarks)
	if err != nil {
		// TODO: Update error detection method once merge request has been merged and released.
		// TODO: Currently, when there is no range between the data points, `go-chart` errors using `fmt.Errorf`.
		// TODO: This is the merge request: https://github.com/wcharczuk/go-chart/pull/169
		log.Fatalf("Could not output chart - error: %v", err)
	}
}
