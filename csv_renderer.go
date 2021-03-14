package go_benchpress

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/tools/benchmark/parse"
	"io"
	"strconv"
)

type CSVRenderer struct {

}

func (c *CSVRenderer) Render(writer io.Writer, parentBenchmark string, dimension RenderDimension, benchmarks []parse.Benchmark) error {
	csvWriter := csv.NewWriter(writer)

	// Write header
	header := []string{"Name", "N", "NsPerOp", "AllocedBytesPerOp", "AllocsPerOp", "MBPerS", "Measured", "Ord"}
	err := csvWriter.Write(header)
	if err != nil {
		return err
	}

	// Write records
	for _, benchmark := range benchmarks {
		n := strconv.Itoa(benchmark.N)
		nsPerOp := fmt.Sprintf("%.12f", benchmark.NsPerOp)
		allocedBytesPerOp := strconv.FormatUint(benchmark.AllocedBytesPerOp, 10)
		allocsPerOp := strconv.FormatUint(benchmark.AllocsPerOp, 10)
		mbPerS := fmt.Sprintf("%.12f", benchmark.MBPerS)
		measured := strconv.Itoa(benchmark.Measured)
		ord := strconv.Itoa(benchmark.Ord)

		record := []string{benchmark.Name, n, nsPerOp, allocedBytesPerOp, allocsPerOp, mbPerS, measured, ord}
		err := csvWriter.Write(record)
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()

	return csvWriter.Error()
}



