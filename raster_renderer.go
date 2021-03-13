package benchpress

import (
	"bytes"
	"errors"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/tools/benchmark/parse"
	"io"
	"strings"
)

// RasterRenderer outputs a raster graphic based representation of the benchmarks, compared against one another.
type RasterRenderer struct {
	Title string
	Height int
	BarWidth int
}

func NewRasterRenderer(title string) *RasterRenderer {
	return &RasterRenderer{
		Title:    title,
		Height:   512,
		BarWidth: 60,
	}
}

func (r *RasterRenderer) Render(writer io.Writer, benchmarks []parse.Benchmark) error {

	if len(benchmarks) == 0 {
		return errors.New("could not render benchmarks - no benchmarks provided")
	}

	values := make([]chart.Value, 0)
	var parentBenchmarkTitle string

	for _, benchmark := range benchmarks {
		name := benchmark.Name

		parts := strings.Split(name, "/")
		if len(parts) < 2 {
			continue
		}

		name = parts[1]
		parentBenchmarkTitle = parts[0]

		values = append(values, chart.Value{
			Style: chart.Style {
				Show: true,
			},
			Label: name,
			Value: benchmark.NsPerOp,
		})
	}

	title := r.Title
	if title == "" {
		title = parentBenchmarkTitle
	}

	graph := chart.BarChart{
		Title: title,
		TitleStyle: chart.Style{Show: true},

		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Name: "Ns Per Op",
			Style: chart.Style{
				Show: true,
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   r.Height,
		BarWidth: r.BarWidth,
		Bars: values,
	}

	var buf bytes.Buffer
	err := graph.Render(chart.PNG, &buf)
	if err != nil {
		return err
	}

	_, err = writer.Write(buf.Bytes())
	return err
}
