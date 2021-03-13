package go_benchpress

import (
	"errors"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/tools/benchmark/parse"
	"reflect"
	"testing"
)

// ===== renderGraphicalBarChart tests =====

func TestRenderGraphicalBarChart(t *testing.T) {
	nsPerOpVal := chart.Value{
		Style: chart.StyleShow(),
		Label: "SubBenchmark",
		Value: 1000,
	}

	bytesPerOpVal := chart.Value{
		Style: chart.StyleShow(),
		Label: "SubBenchmark",
		Value: 10000,
	}

	allocsPerOpVal := chart.Value{
		Style: chart.StyleShow(),
		Label: "SubBenchmark",
		Value: 100000,
	}

	baseNsPerOpVal := chart.Value{
		Style: chart.StyleShow(),
		Label: "BenchmarkOne",
		Value: 1000,
	}

	nestedNsPerOpVal := chart.Value{
		Style: chart.StyleShow(),
		Label: "SubBenchmark/SubSubBenchmark",
		Value: 1000,
	}

	tests := []struct {
		name             string
		title            string
		height, barWidth int
		dimension        RenderDimension
		benchmarks       []parse.Benchmark
		want             *chart.BarChart
		wantErr          error

		wantCalled    bool
		wantTitle     string
		wantHeight    int
		wantBarWidth  int
		wantDimension RenderDimension
		wantValues    []chart.Value
	}{
		{
			name:      "ns per op",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderNsPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
			want: &chart.BarChart{
				Title:      "ExampleTitle",
				TitleStyle: chart.Style{Show: true},
				XAxis: chart.Style{
					Show: true,
				},
				YAxis: chart.YAxis{
					Name: RenderNsPerOp.String(),
					Style: chart.Style{
						Show: true,
					},
				},
				Background: chart.Style{
					Padding: chart.Box{
						Top: 40,
					},
				},
				Height:   512,
				BarWidth: 60,
				Bars:     []chart.Value{nsPerOpVal},
			},
			wantCalled:   true,
			wantTitle:    "ExampleTitle",
			wantHeight:   512,
			wantBarWidth: 60,
			wantValues:   []chart.Value{nsPerOpVal},
		},
		{
			name:      "bytes per op",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderBytesPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
			want: &chart.BarChart{
				Title:      "ExampleTitle",
				TitleStyle: chart.Style{Show: true},
				XAxis: chart.Style{
					Show: true,
				},
				YAxis: chart.YAxis{
					Name: RenderNsPerOp.String(),
					Style: chart.Style{
						Show: true,
					},
				},
				Background: chart.Style{
					Padding: chart.Box{
						Top: 40,
					},
				},
				Height:   512,
				BarWidth: 60,
				Bars:     []chart.Value{nsPerOpVal},
			},
			wantCalled:    true,
			wantTitle:     "ExampleTitle",
			wantHeight:    512,
			wantBarWidth:  60,
			wantDimension: RenderBytesPerOp,
			wantValues:    []chart.Value{bytesPerOpVal},
		},
		{
			name:      "allocs per op",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderAllocsPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
			want: &chart.BarChart{
				Title:      "ExampleTitle",
				TitleStyle: chart.Style{Show: true},
				XAxis: chart.Style{
					Show: true,
				},
				YAxis: chart.YAxis{
					Name: RenderNsPerOp.String(),
					Style: chart.Style{
						Show: true,
					},
				},
				Background: chart.Style{
					Padding: chart.Box{
						Top: 40,
					},
				},
				Height:   512,
				BarWidth: 60,
				Bars:     []chart.Value{nsPerOpVal},
			},
			wantCalled:    true,
			wantTitle:     "ExampleTitle",
			wantHeight:    512,
			wantBarWidth:  60,
			wantDimension: RenderAllocsPerOp,
			wantValues:    []chart.Value{allocsPerOpVal},
		},
		{
			name:       "no benchmarks",
			benchmarks: []parse.Benchmark{},
			wantErr:    ErrNoBenchmarksProvided,
		},
		{
			name:      "unknown dimension",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderDimension(1000),
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
			wantErr:       ErrUnknownDimensionType,
		},
		{
			name:      "no nesting",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderNsPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
			want: &chart.BarChart{
				Title:      "ExampleTitle",
				TitleStyle: chart.Style{Show: true},
				XAxis: chart.Style{
					Show: true,
				},
				YAxis: chart.YAxis{
					Name: RenderNsPerOp.String(),
					Style: chart.Style{
						Show: true,
					},
				},
				Background: chart.Style{
					Padding: chart.Box{
						Top: 40,
					},
				},
				Height:   512,
				BarWidth: 60,
				Bars:     []chart.Value{nsPerOpVal},
			},
			wantCalled:   true,
			wantTitle:    "ExampleTitle",
			wantHeight:   512,
			wantBarWidth: 60,
			wantValues:   []chart.Value{baseNsPerOpVal},
		},
		{
			name:      "multiply nested",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderNsPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark/SubSubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
			want: &chart.BarChart{
				Title:      "ExampleTitle",
				TitleStyle: chart.Style{Show: true},
				XAxis: chart.Style{
					Show: true,
				},
				YAxis: chart.YAxis{
					Name: RenderNsPerOp.String(),
					Style: chart.Style{
						Show: true,
					},
				},
				Background: chart.Style{
					Padding: chart.Box{
						Top: 40,
					},
				},
				Height:   512,
				BarWidth: 60,
				Bars:     []chart.Value{nsPerOpVal},
			},
			wantCalled:   true,
			wantTitle:    "ExampleTitle",
			wantHeight:   512,
			wantBarWidth: 60,
			wantValues:   []chart.Value{nestedNsPerOpVal},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Setup fake renderer
			fakeRenderer := fakeBarChartRenderer{}
			fakeRenderer.replyWith = test.want
			_renderBarChart = fakeRenderer.renderBarChart
			// Ensure _renderBarChart is pointing at original implementation at end of test.
			defer func() {
				_renderBarChart = renderBarChart
			}()

			got, err := renderGraphicalBarChart(test.title, test.height, test.barWidth, test.dimension, test.benchmarks)
			if err != nil {
				if !errors.Is(err, test.wantErr) {
					t.Errorf("Want error '%v', got error '%v'", test.wantErr, err)
				}
			}

			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("want %v, got %v", test.want, got)
			}

			fakeRenderer.assertWantsMet(t, test.wantCalled, test.wantTitle, test.wantHeight, test.wantBarWidth, test.wantDimension, test.wantValues)
		})
	}
}

func BenchmarkRenderGraphicalBarChart(b *testing.B) {
	benchmarks := []struct {
		name             string
		title            string
		height, barWidth int
		dimension        RenderDimension
		benchmarks       []parse.Benchmark
	}{

		{
			name:      "ns per op",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderNsPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
		},
		{
			name:      "bytes per op",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderBytesPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
		},
		{
			name:      "allocs per op",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderAllocsPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
		},
		{
			name:       "no benchmarks",
			benchmarks: []parse.Benchmark{},
		},
		{
			name:      "unknown dimension",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderDimension(1000),
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
		},
		{
			name:      "no nesting",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderNsPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
		},
		{
			name:      "multiply nested",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderNsPerOp,
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark/SubSubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultBarChart, _ = renderGraphicalBarChart(bm.title, bm.height, bm.barWidth, bm.dimension, bm.benchmarks)
			}
		})
	}
}

// ===== fakeBarChartRenderer =====

type fakeBarChartRenderer struct {
	replyWith *chart.BarChart

	called           bool
	title            string
	height, barWidth int
	dimension        RenderDimension
	values           []chart.Value
}

func (f *fakeBarChartRenderer) renderBarChart(title string, height, barWidth int, dimension RenderDimension, values []chart.Value) *chart.BarChart {
	f.called = true
	f.title = title
	f.height = height
	f.barWidth = barWidth
	f.dimension = dimension
	f.values = values
	return f.replyWith
}

func (f *fakeBarChartRenderer) assertWantsMet(t *testing.T, wantCalled bool, wantTitle string, wantHeight, wantBarWidth int, wantDimension RenderDimension, wantValues []chart.Value) {

	if wantCalled != f.called {
		t.Errorf("Want render called %v, got render called %v", wantCalled, f.called)
	}

	if wantTitle != f.title {
		t.Errorf("Want render title %q, got render title %q", wantTitle, f.title)
	}

	if wantHeight != f.height {
		t.Errorf("Want render height %d, got render height %d", wantHeight, f.height)
	}

	if wantBarWidth != f.barWidth {
		t.Errorf("Want render bar width %d, got render bar width %d", wantBarWidth, f.barWidth)
	}

	if wantDimension != f.dimension {
		t.Errorf("Want render dimension %q, got render dimension %q", wantDimension, f.dimension)
	}

	if !reflect.DeepEqual(wantValues, f.values) {
		t.Errorf("Want render values %v, got render values %v", wantValues, f.values)
	}
}

// ===== renderBarChart tests =====

func TestRenderBarChart(t *testing.T) {
	value := chart.Value{
		Style: chart.StyleShow(),
		Label: "Data Point",
		Value: 10,
	}

	tests := []struct {
		name             string
		title            string
		height, barWidth int
		dimension        RenderDimension
		values           []chart.Value
		want             *chart.BarChart
	}{
		{
			name:      "ns per op",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderNsPerOp,
			values:    []chart.Value{value},
			want: &chart.BarChart{
				Title:      "ExampleTitle",
				TitleStyle: chart.Style{Show: true},
				XAxis: chart.Style{
					Show: true,
				},
				YAxis: chart.YAxis{
					Name: RenderNsPerOp.String(),
					Style: chart.Style{
						Show: true,
					},
				},
				Background: chart.Style{
					Padding: chart.Box{
						Top: 40,
					},
				},
				Height:   512,
				BarWidth: 60,
				Bars:     []chart.Value{value},
			},
		},
		{
			name:      "bytes per op",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderBytesPerOp,
			values:    []chart.Value{value},
			want: &chart.BarChart{
				Title:      "ExampleTitle",
				TitleStyle: chart.Style{Show: true},
				XAxis: chart.Style{
					Show: true,
				},
				YAxis: chart.YAxis{
					Name: RenderBytesPerOp.String(),
					Style: chart.Style{
						Show: true,
					},
				},
				Background: chart.Style{
					Padding: chart.Box{
						Top: 40,
					},
				},
				Height:   512,
				BarWidth: 60,
				Bars:     []chart.Value{value},
			},
		},
		{
			name:      "allocs per op",
			title:     "ExampleTitle",
			height:    512,
			barWidth:  60,
			dimension: RenderAllocsPerOp,
			values:    []chart.Value{value},
			want: &chart.BarChart{
				Title:      "ExampleTitle",
				TitleStyle: chart.Style{Show: true},
				XAxis: chart.Style{
					Show: true,
				},
				YAxis: chart.YAxis{
					Name: RenderAllocsPerOp.String(),
					Style: chart.Style{
						Show: true,
					},
				},
				Background: chart.Style{
					Padding: chart.Box{
						Top: 40,
					},
				},
				Height:   512,
				BarWidth: 60,
				Bars:     []chart.Value{value},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := renderBarChart(test.title, test.height, test.barWidth, test.dimension, test.values)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("got %v, want %v", test.want, got)
			}
		})
	}
}

var resultBarChart *chart.BarChart

func BenchmarkRenderBarChart(b *testing.B) {
	values := []chart.Value{
		{
			Style: chart.StyleShow(),
			Label: "Data Point",
			Value: 10,
		},
	}

	for i := 0; i < b.N; i++ {
		resultBarChart = renderBarChart("Example Title", 512, 60, RenderNsPerOp, values)
	}
}
