package main

import (
	"flag"
	"github.com/rpickz/go-benchpress"
	"golang.org/x/tools/benchmark/parse"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var input = flag.String("input", "STDIN", "The input filename")
var outputFilename = flag.String("output", "output_{}", "The output filename")
var renderType = flag.String("renderType", "SVG", "The render type - can be 'SVG', 'PNG', 'JSON', or 'CSV'")
var dimension = flag.String("dimension", "NS_PER_OP", "The dimension to compare - can be 'NS_PER_OP', 'BYTES_PER_OP', 'ALLOCS_PER_OP'")

var _logError = logError

func main() {
	flag.Parse()

	var reader io.Reader
	if *input == "STDIN" {
		reader = os.Stdin
	} else {
		file, err := os.Open(*input)
		if err != nil {
			_logError("Could not open output.txt for reading - error: %v", err)
		}
		defer file.Close()
		reader = file
	}

	separatedBenchmarks, err := go_benchpress.ReadAndSeparateBenchmarks(reader)
	if err != nil {
		_logError("Could not read benchmarks from input - error: %v", err)
	}

	dim, err := go_benchpress.RenderDimensionFromString(*dimension)
	if err != nil {
		_logError("Render dimension %q invalid", *dimension)
	}

	for name, benchmarks := range separatedBenchmarks {
		writeBenchmarks(name, benchmarks, dim, *outputFilename)
	}
}

func writeBenchmarks(name string, benchmarks []parse.Benchmark, dimension go_benchpress.RenderDimension, outputFilename string) {

	outputName := strings.ReplaceAll(outputFilename, "{}", name)

	renderType, err := go_benchpress.RenderTypeFromString(*renderType)
	if err != nil {
		_logError("Could not determine valid render type - error: %v", err)
	}

	outputName = determineOutputFilename(outputName, renderType)

	renderer, err := renderType.Renderer(name)
	if err != nil {
		_logError("Could not find renderer for type %q - error: %v", renderType, err)
	}

	file, err := os.Create(outputName)
	if err != nil {
		_logError("Could not open file for writing - error: %v", err)
	}
	defer file.Close()

	err = renderer.Render(file, name, dimension, benchmarks)
	if err != nil {
		// TODO: Update error detection method once merge request has been merged and released.
		// TODO: Currently, when there is no range between the data points, `go-chart` errors using `fmt.Errorf`.
		// TODO: This is the merge request: https://github.com/wcharczuk/go-chart/pull/169
		_logError("Could not output chart - error: %v", err)
	}
}

// determineOutputFilename corrects the filename to output to if the wrong file extension is provided.
func determineOutputFilename(outputName string, renderType go_benchpress.RenderType) string {
	result := outputName

	// If the file extension provided is different from that of the format, change the filename to use the correct
	// file extension.
	ext := filepath.Ext(result)
	formatExt := renderType.FileExtension()
	if ext != formatExt {
		replaceFrom := strings.LastIndex(result, ext)
		if replaceFrom != -1 {
			result = result[:replaceFrom]
		}
		result += formatExt
	}

	return result
}

func logError(format string, vars ...interface{}) {
	log.Fatalf(format, vars...)
}
