package benchpress

import (
	"fmt"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/tools/benchmark/parse"
	"strings"
)

type barChartBenchmarkRenderer func(title string, height, barWidth int, dimension RenderDimension, benchmarks []parse.Benchmark) (*chart.BarChart, error)

type barChartRenderer func(title string, height, barWidth int, dimension RenderDimension, values []chart.Value) *chart.BarChart

// Defined for testing purposes - to isolate testing of the renderer and the construction of go-chart Bar charts.
var _renderBarChart barChartRenderer = renderBarChart

func renderGraphicalBarChart(title string, height int, barWidth int, dimension RenderDimension, benchmarks []parse.Benchmark) (*chart.BarChart, error) {

	if len(benchmarks) == 0 {
		return nil, ErrNoBenchmarksProvided
	}

	values := make([]chart.Value, 0)

	for _, benchmark := range benchmarks {
		name := benchmark.Name

		parts := strings.Split(name, "/")

		name = parts[0]
		if len(parts) > 1 {
			name = strings.Join(parts[1:], "/")
		}

		var value float64
		switch dimension {
		case RenderNsPerOp:
			value = benchmark.NsPerOp
		case RenderBytesPerOp:
			value = float64(benchmark.AllocedBytesPerOp)
		case RenderAllocsPerOp:
			value = float64(benchmark.AllocsPerOp)
		default:
			return nil, fmt.Errorf("render dimension type %q not supported: %w", dimension, ErrUnknownDimensionType)
		}

		values = append(values, chart.Value{
			Style: chart.StyleShow(),
			Label: name,
			Value: value,
		})
	}

	graph := _renderBarChart(title, height, barWidth, dimension, values)

	return graph, nil
}

func renderBarChart(title string, height, barWidth int, dimension RenderDimension, values []chart.Value) *chart.BarChart {
	return &chart.BarChart{
		Title:      title,
		TitleStyle: chart.Style{Show: true},
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Name: dimension.String(),
			Style: chart.Style{
				Show: true,
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   height,
		BarWidth: barWidth,
		Bars:     values,
	}
}
