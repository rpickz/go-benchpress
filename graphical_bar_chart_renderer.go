package benchpress

import (
	"errors"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/tools/benchmark/parse"
	"strings"
)

type barChartRenderer func(title string, height, barWidth int, benchmarks []parse.Benchmark) (*chart.BarChart, error)

func renderGraphicalBarChart(title string, height int, barWidth int, benchmarks []parse.Benchmark) (*chart.BarChart, error) {

	if len(benchmarks) == 0 {
		return nil, errors.New("could not render benchmarks - no benchmarks provided")
	}

	values := make([]chart.Value, 0)

	for _, benchmark := range benchmarks {
		name := benchmark.Name

		parts := strings.Split(name, "/")
		if len(parts) < 2 {
			continue
		}

		name = parts[1]

		values = append(values, chart.Value{
			Style: chart.Style {
				Show: true,
			},
			Label: name,
			Value: benchmark.NsPerOp,
		})
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
		Height:   height,
		BarWidth: barWidth,
		Bars: values,
	}

	return &graph, nil
}
