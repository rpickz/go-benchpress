package main

import (
	"bytes"
	"github.com/wcharczuk/go-chart"
	"go-benchpress/m/v2"
	"golang.org/x/tools/benchmark/parse"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("./output.txt")
	if err != nil {
		log.Fatalf("Could not open output.txt for reading - error: %v", err)
	}
	defer file.Close()

	separatedBenchmarks, err := benchpress.ReadAndSeparateBenchmarks(file)
	if err != nil {
		log.Fatalf("Could not read benchmarks from input - error: %v", err)
	}

	for i, benchmarks := range separatedBenchmarks {
		err = chartBenchmarks(benchmarks, i+".png")
		if err != nil {
			log.Fatalf("Could not output chart - error: %v", err)
		}
	}
}

func chartBenchmarks(benchmarks []parse.Benchmark, outputFile string) error {

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
		Title: "Test Bar Chart",
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Name: "SomethingY",
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
		Bars: values,
	}

	var buf bytes.Buffer
	err := graph.Render(chart.PNG, &buf)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputFile, buf.Bytes(), 0777)
}
