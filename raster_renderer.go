package go_benchpress

import (
	"fmt"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/tools/benchmark/parse"
	"io"
)

// RasterRenderer outputs a raster graphic based representation of the benchmarks, compared against one another.
type RasterRenderer struct {
	Title      string
	Height     int
	BarWidth   int
	RenderType RenderType

	// barChartRenderFunc is used to isolate unit testing - in non-testing usage, points to `renderGraphicalBarChart`.
	barChartRenderFunc barChartBenchmarkRenderer
}

func NewRasterRenderer(title string, renderType RenderType) *RasterRenderer {
	return &RasterRenderer{
		Title:    title,
		Height:   512,
		BarWidth: 60,
		RenderType: renderType,
		barChartRenderFunc: renderGraphicalBarChart,
	}
}

func (r *RasterRenderer) Render(writer io.Writer, parentBenchmark string, renderDimension RenderDimension, benchmarks []parse.Benchmark) error {

	if len(benchmarks) == 0 {
		return ErrNoBenchmarksProvided
	}

	title := r.Title
	if title == "" {
		title = parentBenchmark
	}

	graph, err := r.barChartRenderFunc(title, r.Height, r.BarWidth, renderDimension, benchmarks)
	if err != nil {
		return err
	}

	var renderer chart.RendererProvider
	switch r.RenderType {
	case PNG:
		renderer = chart.PNG
	case SVG:
		renderer = chart.SVG
	default:
		return fmt.Errorf("render type %q not supported: %w", r.RenderType, ErrUnknownRenderType)
	}

	return graph.Render(renderer, writer)
}
