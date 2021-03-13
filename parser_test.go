package benchpress

import (
	"golang.org/x/tools/benchmark/parse"
	"io"
	"reflect"
	"strings"
	"testing"
)

// ===== ReadBenchmarks Tests =====

func TestReadBenchmarks(t *testing.T) {
	tests := []struct {
		name  string
		input io.Reader
		want  []parse.Benchmark
	}{
		{
			name:  "SingleLine",
			input: strings.NewReader("BenchmarkSomething/SubBenchmark-12   	   10000	     10000 ns/op"),
			want: []parse.Benchmark{
				{
					Name:     "BenchmarkSomething/SubBenchmark-12",
					N:        10000,
					NsPerOp:  10000,
					Measured: 1,
				},
			},
		},
		{
			name: "MultipleLines",
			input: strings.NewReader(`BenchmarkSomething/SubBenchmark1-12   	   10000	     10000 ns/op
BenchmarkSomething/SubBenchmark2-12   	   10000	     10000 ns/op
BenchmarkSomething/SubBenchmark3-12   	   10000	     10000 ns/op`),
			want: []parse.Benchmark{
				{
					Name:     "BenchmarkSomething/SubBenchmark1-12",
					N:        10000,
					NsPerOp:  10000,
					Measured: 1,
				},
				{
					Name:     "BenchmarkSomething/SubBenchmark2-12",
					N:        10000,
					NsPerOp:  10000,
					Measured: 1,
				},
				{
					Name:     "BenchmarkSomething/SubBenchmark3-12",
					N:        10000,
					NsPerOp:  10000,
					Measured: 1,
				},
			},
		},
		{
			name:  "SkipsNonBenchmarkLines",
			input: strings.NewReader("this is not a benchmark"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			benchmarks, err := readBenchmarks(test.input)
			if err != nil {
				t.Errorf("Error reading benchmarks - error: %v", err)
			}
			if !reflect.DeepEqual(test.want, benchmarks) {
				t.Errorf("want %v, got %v", test.want, benchmarks)
			}
		})
	}
}

var benchmarksRead []parse.Benchmark

func BenchmarkReadBenchmarks(b *testing.B) {
	benchmarks := []struct {
		name  string
		input io.Reader
	}{
		{
			name:  "Simple",
			input: strings.NewReader("BenchmarkSomething/SubBenchmark-12   	   10000	     10000 ns/op"),
		},
		{
			name: "MultipleLines",
			input: strings.NewReader(`BenchmarkSomething/SubBenchmark1-12   	   10000	     10000 ns/op
BenchmarkSomething/SubBenchmark2-12   	   10000	     10000 ns/op
BenchmarkSomething2/SubBenchmark3-12   	   10000	     10000 ns/op`),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var err error
				benchmarksRead, err = readBenchmarks(bm.input)
				if err != nil {
					b.Errorf("Could not read benchmark - error: %v", err)
				}
			}
		})
	}
}

// ===== ReadAndSeparateBenchmarks Tests =====

func TestReadAndSeparateBenchmarks(t *testing.T) {
	tests := []struct {
		name  string
		input io.Reader
		want  map[string][]parse.Benchmark
	}{
		{
			name:  "SingleLine",
			input: strings.NewReader("BenchmarkSomething/SubBenchmark-12   	   10000	     10000 ns/op"),
			want: map[string][]parse.Benchmark{
				"BenchmarkSomething": {
					{
						Name:     "BenchmarkSomething/SubBenchmark-12",
						N:        10000,
						NsPerOp:  10000,
						Measured: 1,
					},
				},
			},
		},
		{
			name: "MultipleLines",
			input: strings.NewReader(`BenchmarkSomething/SubBenchmark1-12   	   10000	     10000 ns/op
BenchmarkSomething/SubBenchmark2-12   	   10000	     10000 ns/op
BenchmarkSomething2/SubBenchmark3-12   	   10000	     10000 ns/op`),
			want: map[string][]parse.Benchmark{
				"BenchmarkSomething": {
					{
						Name:     "BenchmarkSomething/SubBenchmark1-12",
						N:        10000,
						NsPerOp:  10000,
						Measured: 1,
					},
					{
						Name:     "BenchmarkSomething/SubBenchmark2-12",
						N:        10000,
						NsPerOp:  10000,
						Measured: 1,
					},
				},
				"BenchmarkSomething2": {
					{
						Name:     "BenchmarkSomething2/SubBenchmark3-12",
						N:        10000,
						NsPerOp:  10000,
						Measured: 1,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			benchmarks, err := ReadAndSeparateBenchmarks(test.input)
			if err != nil {
				t.Errorf("Error reading benchmarks - error: %v", err)
			}
			if !reflect.DeepEqual(test.want, benchmarks) {
				t.Errorf("want %v, got %v", test.want, benchmarks)
			}
		})
	}
}

var separatedBenchmarksRead map[string][]parse.Benchmark

func BenchmarkReadAndSeparateBenchmarks(b *testing.B) {
	benchmarks := []struct {
		name  string
		input io.Reader
	}{
		{
			name:  "Simple",
			input: strings.NewReader("BenchmarkSomething/SubBenchmark-12   	   10000	     10000 ns/op"),
		},
		{
			name: "MultipleLines",
			input: strings.NewReader(`BenchmarkSomething/SubBenchmark1-12   	   10000	     10000 ns/op
BenchmarkSomething/SubBenchmark2-12   	   10000	     10000 ns/op
BenchmarkSomething2/SubBenchmark3-12   	   10000	     10000 ns/op`),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var err error
				separatedBenchmarksRead, err = ReadAndSeparateBenchmarks(bm.input)
				if err != nil {
					b.Errorf("Could not read benchmark - error: %v", err)
				}
			}
		})
	}
}