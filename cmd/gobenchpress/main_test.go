package main

import (
	"github.com/rpickz/go-benchpress"
	"io/ioutil"
	"os"
	"testing"
)

func TestSVGOutput(t *testing.T) {
	benchmarkFile := setupBenchmarkInput(t)
	defer benchmarkFile.Close()

	file := setupOutputFile(t)
	defer file.Close()

	// Call program entry point.
	main()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Could not read output file - error: %v", err)
	}

	got := len(content)
	wantMoreThan := 3000
	if got < wantMoreThan {
		t.Errorf("Want more than %d content length, got %d content length", wantMoreThan, len(content))
	}
}

func TestPNGOutput(t *testing.T) {
	benchmarkFile := setupBenchmarkInput(t)
	defer benchmarkFile.Close()

	file := setupOutputFile(t)
	defer file.Close()

	setupRasterRenderType(go_benchpress.PNG)

	// Call program entry point.
	main()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Could not read output file - error: %v", err)
	}

	got := len(content)
	wantMoreThan := 3000
	if got < wantMoreThan {
		t.Errorf("Want more than %d content length, got %d content length", wantMoreThan, len(content))
	}
}

func setupRasterRenderType(t go_benchpress.RasterRenderType) {
	render := new(string)
	*render = t.String()
	renderType = render
}

func setupOutputFile(t *testing.T) *os.File {
	file, err := os.CreateTemp("", "benchmark-output-*.svg")
	if err != nil {
		t.Fatalf("Could not create temporary file for output data")
	}
	outputFile := new(string)
	*outputFile = file.Name()
	outputFilename = outputFile
	return file
}

func setupBenchmarkInput(t *testing.T) *os.File {
	fileContent := `
goos: darwin
goarch: amd64
pkg: go-benchpress/m/v2/cmd/examples/csvparser
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkParseCSVLineFields/10_Fields-12                 5764971               193.9 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFields/20_Fields-12                 5929747               204.8 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFields/40_Fields-12                 5663358               210.2 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFields/80_Fields-12                 5247201               243.9 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFields/160_Fields-12                4742551               251.8 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFields/320_Fields-12                4086375               319.8 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFields/640_Fields-12                2326442               462.0 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFields/1280_Fields-12               1947211               609.3 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFieldLength/Length_10-12            5744911               207.4 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFieldLength/Length_20-12            3311017               370.3 ns/op           387 B/op          2 allocs/op
BenchmarkParseCSVLineFieldLength/Length_40-12            5159778               229.6 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFieldLength/Length_80-12            2851275               425.0 ns/op           387 B/op          2 allocs/op
BenchmarkParseCSVLineFieldLength/Length_160-12           1929837               622.6 ns/op           711 B/op          2 allocs/op
BenchmarkParseCSVLineFieldLength/Length_320-12           2768451               432.9 ns/op            64 B/op          2 allocs/op
BenchmarkParseCSVLineFieldLength/Length_640-12             12278             95948 ns/op            8025 B/op          5 allocs/op
BenchmarkParseCSVLineFieldLength/Length_1280-12            10000            109451 ns/op            5638 B/op         23 allocs/op
PASS
ok      go-benchpress/m/v2/cmd/examples/csvparser       25.236s
`

	file, err := os.CreateTemp("", "benchmark-input-*.txt")
	if err != nil {
		t.Fatalf("Could not create temporary file for benchmark data - error: %v", err)
	}

	err = ioutil.WriteFile(file.Name(), []byte(fileContent), 0775)
	if err != nil {
		t.Fatalf("Could not benchmark data into temporary file - error: %v", err)
	}

	name := new(string)
	*name = file.Name()
	input = name

	return file
}

