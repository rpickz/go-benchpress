package go_benchpress

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

func RasterRenderTypeFromString(str string) (RasterRenderType, error) {
	switch str {
	case "PNG":
		return PNG, nil
	case "SVG":
		return SVG, nil
	default:
		return -1, fmt.Errorf("raster render type %q not supported: %w", str, ErrUnknownRasterRenderType)
	}
}

// RasterRenderer outputs a raster graphic based representation of the benchmarks, compared against one another.
type RasterRenderer struct {
	Title string
	Height int
	BarWidth int
	RenderType RasterRenderType

	// barChartRenderFunc is used to isolate unit testing - in non-testing usage, points to `renderGraphicalBarChart`.
	barChartRenderFunc barChartBenchmarkRenderer
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
		return fmt.Errorf("render type %q not supported: %w", r.RenderType, ErrUnknownRasterRenderType)
	}

	return graph.Render(renderer, writer)
}
