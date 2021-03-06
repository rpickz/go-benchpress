package go_benchpress

import (
	"bufio"
	"fmt"
	"golang.org/x/tools/benchmark/parse"
	"io"
	"strings"
)

// ReadBenchmarks uses the provided reader, and reads the benchmarks from the read lines.
// If lines cannot be read, or parsed, an error is returned.
func ReadBenchmarks(reader io.Reader) ([]parse.Benchmark, error) {
	var results []parse.Benchmark
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Benchmark") {
			continue
		}

		benchmark, err := parse.ParseLine(line)
		if err != nil {
			return nil, fmt.Errorf("parsing error: %w", ErrCouldNotParseLine)
		}
		results = append(results, *benchmark)
	}
	return results, scanner.Err()
}

// BenchmarkSets represents a number of benchmarks, grouped by the parent benchmark.
type BenchmarkSets map[string][]parse.Benchmark

// ReadAndSeparateBenchmarks reads benchmarks from the provided reader, and groups them by benchmark name
// (which is split from sub-benchmark names).  If the operation did not succeed, an error is returned.
func ReadAndSeparateBenchmarks(reader io.Reader) (BenchmarkSets, error) {
	benchmarks, err := ReadBenchmarks(reader)
	if err != nil {
		return nil, err
	}

	results := make(map[string][]parse.Benchmark)
	for _, val := range benchmarks {
		parts := strings.Split(val.Name, "/")
		benchName := parts[0]

		s, ok := results[benchName]
		if !ok {
			s = make([]parse.Benchmark, 0)
		}
		s = append(s, val)
		results[benchName] = s
	}
	return results, nil
}
