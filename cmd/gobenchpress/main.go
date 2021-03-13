package main

import (
	"flag"
	"go-benchpress/m/v2"
	"golang.org/x/tools/benchmark/parse"
	"io"
	"log"
	"os"
)

var input = flag.String("input", "STDIN", "The input filename")
var outputFilename = flag.String("output", "output.svg", "The output filename")
var dimension = flag.String("dimension", "NS_PER_OP", "The dimension to compare - can be 'NS_PER_OP', 'BYTES_PER_OP', 'ALLOCS_PER_OP'")

func main() {
	flag.Parse()

	var reader io.Reader
	if *input == "STDIN" {
		reader = os.Stdin
	} else {
		file, err := os.Open("./output.txt")
		if err != nil {
			log.Fatalf("Could not open output.txt for reading - error: %v", err)
		}
		defer file.Close()
		reader = file
	}

	separatedBenchmarks, err := benchpress.ReadAndSeparateBenchmarks(reader)
	if err != nil {
		log.Fatalf("Could not read benchmarks from input - error: %v", err)
	}

	dim, err := benchpress.RenderDimensionFromString(*dimension)
	if err != nil {
		log.Fatalf("Render dimension %q invalid", *dimension)
	}

	for name, benchmarks := range separatedBenchmarks {
		writeBenchmarks(name, benchmarks, dim, *outputFilename)
	}
}

func writeBenchmarks(name string, benchmarks []parse.Benchmark, dimension benchpress.RenderDimension, outputFilename string) {

	renderer := benchpress.NewRasterRenderer(name, benchpress.SVG)

	file, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("Could not open file for writing - error: %v", err)
	}
	defer file.Close()

	err = renderer.Render(file, name, benchpress.RenderNsPerOp, benchmarks)
	if err != nil {
		// TODO: Update error detection method once merge request has been merged and released.
		// TODO: Currently, when there is no range between the data points, `go-chart` errors using `fmt.Errorf`.
		// TODO: This is the merge request: https://github.com/wcharczuk/go-chart/pull/169
		log.Fatalf("Could not output chart - error: %v", err)
	}
}
