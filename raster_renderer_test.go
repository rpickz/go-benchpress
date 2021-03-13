package go_benchpress

import (
	"bytes"
	"errors"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/tools/benchmark/parse"
	"reflect"
	"testing"
)

func TestRasterRenderType_String(t *testing.T) {
	tests := []struct {
		name  string
		input RasterRenderType
		want  string
	}{
		{
			name:  "png",
			input: PNG,
			want:  "PNG",
		},
		{
			name:  "svg",
			input: SVG,
			want:  "SVG",
		},
		{
			name:  "svg",
			input: RasterRenderType(1000),
			want:  "Unknown (1000)",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.String()
			if test.want != got {
				t.Errorf("want %q, got %q", test.want, got)
			}
		})
	}
}

func TestRasterRenderTypeFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    RasterRenderType
		wantErr error
	}{
		{
			name:  "png",
			input: "PNG",
			want:  PNG,
		},
		{
			name:  "svg",
			input: "SVG",
			want:  SVG,
		},
		{
			name:    "unknown",
			input:   "abc123",
			want:    RasterRenderType(-1),
			wantErr: ErrUnknownRasterRenderType,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := RasterRenderTypeFromString(test.input)
			if err != nil {
				if !errors.Is(err, test.wantErr) {
					t.Errorf("Want error '%v', got error '%v'", test.wantErr, err)
				}
			}
			if test.want != got {
				t.Errorf("want %q, got %q", test.want, got)
			}
		})
	}
}

// ===== RasterRenderer tests =====

func TestRasterRenderer_Render(t *testing.T) {
	benchmark := parse.Benchmark{
		Name:     "Benchmark1",
		N:        100,
		NsPerOp:  100,
		Measured: 1,
	}

	barChartRenderErr := errors.New("something went wrong rendering bar chart")

	tests := []struct {
		name       string
		benchmarks []parse.Benchmark
		renderType RasterRenderType
		dimension  RenderDimension

		rendererTitle string

		barChartRenderErr error

		wantError error

		wantRenderCalled     bool
		wantRenderTitle      string
		wantRenderHeight     int
		wantRenderBarWidth   int
		wantDimension        RenderDimension
		wantRenderBenchmarks []parse.Benchmark
	}{
		{
			name:                 "single benchmark",
			rendererTitle:        "RasterRendererTitle",
			benchmarks:           []parse.Benchmark{benchmark},
			wantRenderCalled:     true,
			wantRenderTitle:      "RasterRendererTitle",
			wantRenderHeight:     512,
			wantRenderBarWidth:   60,
			wantRenderBenchmarks: []parse.Benchmark{benchmark},
		},
		{
			name:      "no benchmarks",
			wantError: ErrNoBenchmarksProvided,
		},
		{
			name:                 "no renderer title chooses benchmark title",
			benchmarks:           []parse.Benchmark{benchmark},
			wantRenderCalled:     true,
			wantRenderTitle:      "ParentBenchmark",
			wantRenderHeight:     512,
			wantRenderBarWidth:   60,
			wantRenderBenchmarks: []parse.Benchmark{benchmark},
		},
		{
			name:                 "bar chart rendering error",
			benchmarks:           []parse.Benchmark{benchmark},
			barChartRenderErr:    barChartRenderErr,
			wantError:            barChartRenderErr,
			wantRenderCalled:     true,
			wantRenderTitle:      "ParentBenchmark",
			wantRenderHeight:     512,
			wantRenderBarWidth:   60,
			wantRenderBenchmarks: []parse.Benchmark{benchmark},
		},
		{
			name:                 "render svg",
			rendererTitle:        "RasterRendererTitle",
			benchmarks:           []parse.Benchmark{benchmark},
			renderType:           SVG,
			wantRenderCalled:     true,
			wantRenderTitle:      "RasterRendererTitle",
			wantRenderHeight:     512,
			wantRenderBarWidth:   60,
			wantRenderBenchmarks: []parse.Benchmark{benchmark},
		},
		{
			name:                 "unknown render type",
			rendererTitle:        "RasterRendererTitle",
			benchmarks:           []parse.Benchmark{benchmark},
			renderType:           RasterRenderType(100),
			wantError:            ErrUnknownRasterRenderType,
			wantRenderCalled:     true,
			wantRenderTitle:      "RasterRendererTitle",
			wantRenderHeight:     512,
			wantRenderBarWidth:   60,
			wantRenderBenchmarks: []parse.Benchmark{benchmark},
		},
		{
			name:                 "bytes per op dimension",
			rendererTitle:        "RasterRendererTitle",
			benchmarks:           []parse.Benchmark{benchmark},
			dimension:            RenderBytesPerOp,
			renderType:           RasterRenderType(100),
			wantError:            ErrUnknownRasterRenderType,
			wantRenderCalled:     true,
			wantRenderTitle:      "RasterRendererTitle",
			wantRenderHeight:     512,
			wantRenderBarWidth:   60,
			wantDimension:        RenderBytesPerOp,
			wantRenderBenchmarks: []parse.Benchmark{benchmark},
		},
		{
			name:                 "allocs per op dimension",
			rendererTitle:        "RasterRendererTitle",
			benchmarks:           []parse.Benchmark{benchmark},
			dimension:            RenderAllocsPerOp,
			renderType:           RasterRenderType(100),
			wantError:            ErrUnknownRasterRenderType,
			wantRenderCalled:     true,
			wantRenderTitle:      "RasterRendererTitle",
			wantRenderHeight:     512,
			wantRenderBarWidth:   60,
			wantDimension:        RenderAllocsPerOp,
			wantRenderBenchmarks: []parse.Benchmark{benchmark},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBufferString("")

			rasterRenderer := NewRasterRenderer(test.rendererTitle, test.renderType)
			fakeBarRenderer := newDefaultFakeBarChartRenderer()
			fakeBarRenderer.replyWithError = test.barChartRenderErr
			rasterRenderer.barChartRenderFunc = fakeBarRenderer.fakeRenderGraphicalBarChart

			err := rasterRenderer.Render(buf, "ParentBenchmark", test.dimension, test.benchmarks)
			if err != nil {
				if !errors.Is(err, test.wantError) {
					t.Errorf("Want error '%v', got error '%v'", test.wantError, err)
				}
			}

			if test.wantRenderCalled != fakeBarRenderer.called {
				t.Errorf("Want render called %v, got render called %v", test.wantRenderCalled, fakeBarRenderer.called)
			}

			if test.wantRenderTitle != fakeBarRenderer.title {
				t.Errorf("Want render title %q, got render title %q", test.wantRenderTitle, fakeBarRenderer.title)
			}

			if test.wantRenderHeight != fakeBarRenderer.height {
				t.Errorf("Want render height %d, got render height %d", test.wantRenderHeight, fakeBarRenderer.height)
			}

			if test.wantRenderBarWidth != fakeBarRenderer.barWidth {
				t.Errorf("Want render bar width %d, got render bar width %d", test.wantRenderBarWidth, fakeBarRenderer.barWidth)
			}

			if test.wantDimension != fakeBarRenderer.dimension {
				t.Errorf("Want render dimension %q, got render dimension %q", test.wantDimension, fakeBarRenderer.dimension)
			}

			if !reflect.DeepEqual(test.wantRenderBenchmarks, fakeBarRenderer.benchmarks) {
				t.Errorf("Want render benchmarks %v, got render benchmarks %v", test.wantRenderBenchmarks, fakeBarRenderer.benchmarks)
			}
		})
	}
}

// ===== fakeBarChartBenchmarkRenderer =====

type fakeBarChartBenchmarkRenderer struct {
	replyWithChart *chart.BarChart
	replyWithError error

	called     bool
	title      string
	height     int
	barWidth   int
	dimension  RenderDimension
	benchmarks []parse.Benchmark
}

func newDefaultFakeBarChartRenderer() fakeBarChartBenchmarkRenderer {
	return fakeBarChartBenchmarkRenderer{
		replyWithChart: &chart.BarChart{
			Bars: []chart.Value{
				{Value: 1.0, Label: "A"},
				{Value: 2.0, Label: "B"},
			},
		},
	}
}

func (f *fakeBarChartBenchmarkRenderer) fakeRenderGraphicalBarChart(title string, height int, barWidth int, dimension RenderDimension, benchmarks []parse.Benchmark) (*chart.BarChart, error) {
	f.called = true
	f.title = title
	f.height = height
	f.barWidth = barWidth
	f.dimension = dimension
	f.benchmarks = benchmarks
	return f.replyWithChart, f.replyWithError
}
