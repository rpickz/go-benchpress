package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/rpickz/go-benchpress"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestSVGOutput(t *testing.T) {
	benchmarkFile := setupBenchmarkInput(t)
	defer benchmarkFile.Close()

	file := setupOutputFile(t, go_benchpress.SVG)
	defer file.Close()

	// Call program entry point.
	main()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Could not read output file - error: %v", err)
	}

	// Unmarshal SVG as XML - test file not corrupt, or wrong format.
	xmlData := make([]interface{}, 0)
	err = xml.Unmarshal(content, &xmlData)
	if err != nil {
		t.Errorf("Error unmarshalling SVG as XML - error: %v", err)
	}
}

func TestPNGOutput(t *testing.T) {
	benchmarkFile := setupBenchmarkInput(t)
	defer benchmarkFile.Close()

	file := setupOutputFile(t, go_benchpress.PNG)
	defer file.Close()

	setupRenderType(go_benchpress.PNG)

	// Call program entry point.
	main()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Could not read output file - error: %v", err)
	}

	_, err = png.Decode(bytes.NewBuffer(content))
	if err != nil {
		t.Errorf("Could not decode PNG file - error: %v", err)
	}
}

func TestJSONOutput(t *testing.T) {
	benchmarkFile := setupBenchmarkInput(t)
	defer benchmarkFile.Close()

	file := setupOutputFile(t, go_benchpress.JSON)
	defer file.Close()

	setupRenderType(go_benchpress.JSON)

	// Call program entry point.
	main()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Could not read output file - error: %v", err)
	}

	data := new(interface{})
	err = json.Unmarshal(content, &data)
	if err != nil {
		t.Errorf("Could not decode JSON file - error: %v", err)
	}
}

func TestCSVOutput(t *testing.T) {
	benchmarkFile := setupBenchmarkInput(t)
	defer benchmarkFile.Close()

	file := setupOutputFile(t, go_benchpress.CSV)
	defer file.Close()

	setupRenderType(go_benchpress.CSV)

	// Call program entry point.
	main()

	csvReader := csv.NewReader(file)
	_, err := csvReader.ReadAll()
	if err != nil {
		t.Errorf("Could not decode CSV file - error: %v", err)
	}
}

func TestInvalidRenderType(t *testing.T) {
	wantErr := `Could not determine valid render type - error: render type "Unknown (1000)" not supported: unknown render type`
	errorLogger := fakeErrorLogger{}

	defer func() {
		p := recover()
		if p != nil {
			if !errorLogger.called {
				t.Fatalf("Unexpected panic - not triggered by error logger - panic: %v", p)
			}

			if wantErr != errorLogger.msg {
				t.Errorf("Wanted error msg %q, got error msg %q", wantErr, errorLogger.msg)
			}
		}

		// Ensure error logger is called, regardless of if panic occurs.
		if !errorLogger.called {
			t.Error("Error logger not called - expected error")
		}
	}()

	_logError = errorLogger.logError

	benchmarkFile := setupBenchmarkInput(t)
	defer benchmarkFile.Close()

	renderType := go_benchpress.RenderType(1000)

	file := setupOutputFile(t, renderType)
	defer file.Close()

	setupRenderType(renderType)

	// Call program entry point.
	main()
}

func TestInvalidRenderDimension(t *testing.T) {
	wantErr := `Render dimension "Unknown (1000)" invalid`
	errorLogger := fakeErrorLogger{}

	defer func() {
		p := recover()
		if p != nil {
			if !errorLogger.called {
				t.Fatalf("Unexpected panic - not triggered by error logger - panic: %v", p)
			}

			if wantErr != errorLogger.msg {
				t.Errorf("Wanted error msg %q, got error msg %q", wantErr, errorLogger.msg)
			}
		}

		// Ensure error logger is called, regardless of if panic occurs.
		if !errorLogger.called {
			t.Error("Error logger not called - expected error")
		}
	}()

	_logError = errorLogger.logError

	benchmarkFile := setupBenchmarkInput(t)
	defer benchmarkFile.Close()

	setupRenderDimension(go_benchpress.RenderDimension(1000))

	file := setupOutputFile(t, go_benchpress.SVG)
	defer file.Close()

	setupRenderType(go_benchpress.SVG)

	// Call program entry point.
	main()
}

// ===== determineOutputFilename tests =====

func TestDetermineOutputFilename(t *testing.T) {
	tests := []struct {
		name           string
		outputFilename string
		renderType     go_benchpress.RenderType
		want           string
	}{
		// === Tests not correcting file extension ===
		{
			name:           "png",
			outputFilename: "hello.png",
			renderType:     go_benchpress.PNG,
			want:           "hello.png",
		},
		{
			name:           "svg",
			outputFilename: "hello.svg",
			renderType:     go_benchpress.SVG,
			want:           "hello.svg",
		},
		{
			name:           "json",
			outputFilename: "hello.json",
			renderType:     go_benchpress.JSON,
			want:           "hello.json",
		},
		{
			name:           "csv",
			outputFilename: "hello.csv",
			renderType:     go_benchpress.CSV,
			want:           "hello.csv",
		},
		// === Incorrect file extension correction tests ===
		{
			name:           "incorrect png",
			outputFilename: "hello.something",
			renderType:     go_benchpress.PNG,
			want:           "hello.png",
		},
		{
			name:           "incorrect svg",
			outputFilename: "hello.something",
			renderType:     go_benchpress.SVG,
			want:           "hello.svg",
		},
		{
			name:           "incorrect json",
			outputFilename: "hello.something",
			renderType:     go_benchpress.JSON,
			want:           "hello.json",
		},
		{
			name:           "incorrect csv",
			outputFilename: "hello.something",
			renderType:     go_benchpress.CSV,
			want:           "hello.csv",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := determineOutputFilename(test.outputFilename, test.renderType)
			if test.want != got {
				t.Errorf("Wanted %q, got %q", test.want, got)
			}
		})
	}
}

// ===== Test utilities =====

func setupRenderDimension(t go_benchpress.RenderDimension) {
	*dimension = t.String()
}

func setupRenderType(t go_benchpress.RenderType) {
	*renderType = t.String()
}

func setupOutputFile(t *testing.T, renderType go_benchpress.RenderType) *os.File {
	file, err := os.CreateTemp("", "benchmark-output-*"+renderType.FileExtension())
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

// ===== fakeErrorLogger =====

type fakeErrorLogger struct {
	called bool
	msg    string
}

func (f *fakeErrorLogger) logError(format string, vars ...interface{}) {
	f.called = true
	f.msg = fmt.Sprintf(format, vars...)
	panic("error logged")
}
