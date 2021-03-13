package benchpress

import (
	"fmt"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/tools/benchmark/parse"
	"io"
)

type RasterRenderType int

const (
	PNG RasterRenderType = iota
	SVG
)

func (r RasterRenderType) String() string {
	switch r {
	case PNG:
		return "PNG"
	case SVG:
		return "SVG"
	default:
		return fmt.Sprintf("Unknown (%d)", r)
	}
}

// RasterRenderer outputs a raster graphic based representation of the benchmarks, compared against one another.
type RasterRenderer struct {
	Title string
	Height int
	BarWidth int
	RenderType RasterRenderType

	// barChartRenderFunc is used to isolate unit testing - in non-testing usage, points to `renderGraphicalBarChart`.
	barChartRenderFunc barChartRenderer
}

func NewRasterRenderer(title string, renderType RasterRenderType) *RasterRenderer {
	return &RasterRenderer{
		Title:    title,
		Height:   512,
		BarWidth: 60,
		RenderType: renderType,
		barChartRenderFunc: renderGraphicalBarChart,
	}
}

func (r *RasterRenderer) Render(writer io.Writer, parentBenchmark string, benchmarks []parse.Benchmark) error {

	if len(benchmarks) == 0 {
		return ErrNoBenchmarksProvided
	}

	title := r.Title
	if title == "" {
		title = parentBenchmark
	}

	graph, err := r.barChartRenderFunc(title, r.Height, r.BarWidth, benchmarks)
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
		return fmt.Errorf("render type %q not supported: %w", r.RenderType, ErrUnknownRasterRenderType)
	}

	return graph.Render(renderer, writer)
}
