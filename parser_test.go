package go_benchpress

import (
	"errors"
	"golang.org/x/tools/benchmark/parse"
	"io"
	"reflect"
	"strings"
	"testing"
)

// ===== ReadBenchmarks Tests =====

func TestReadBenchmarks(t *testing.T) {
	tests := []struct {
		name    string
		input   io.Reader
		want    []parse.Benchmark
		wantErr error
	}{
		{
			name:  "single line",
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
			name: "multiple lines",
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
			name:  "skips non benchmark lines",
			input: strings.NewReader("this is not a benchmark"),
		},
		{
			name:  "malformed benchmark lines error",
			input: strings.NewReader("Benchmark$123tlekgjb13rdjasldjv12e;2'"),
			wantErr: ErrCouldNotParseLine,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			benchmarks, err := readBenchmarks(test.input)
			if err != nil {
				if !errors.Is(err, test.wantErr) {
					t.Errorf("Wanted error '%v', got error '%v'", test.wantErr, err)
				}
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
			name:  "single benchmark",
			input: strings.NewReader("BenchmarkSomething/SubBenchmark-12   	   10000	     10000 ns/op"),
		},
		{
			name: "multiple lines",
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
		want  BenchmarkSets
		wantErr error
	}{
		{
			name:  "single benchmark",
			input: strings.NewReader("BenchmarkSomething/SubBenchmark-12   	   10000	     10000 ns/op"),
			want: BenchmarkSets{
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
			name: "multiple lines",
			input: strings.NewReader(`BenchmarkSomething/SubBenchmark1-12   	   10000	     10000 ns/op
BenchmarkSomething/SubBenchmark2-12   	   10000	     10000 ns/op
BenchmarkSomething2/SubBenchmark3-12   	   10000	     10000 ns/op`),
			want: BenchmarkSets{
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
		{
			name:  "malformed benchmark lines error",
			input: strings.NewReader("Benchmark$123tlekgjb13rdjasldjv12e;2'"),
			wantErr: ErrCouldNotParseLine,
		},
		{
			name:  "single benchmark without parent",
			input: strings.NewReader("Benchmark   	   10000	     10000 ns/op"),
			want: BenchmarkSets{
				"Benchmark": {
					{
						Name:     "Benchmark",
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
				if !errors.Is(err, test.wantErr) {
					t.Errorf("Wanted error '%v', got error '%v'", test.wantErr, err)
				}
			}
			if !reflect.DeepEqual(test.want, benchmarks) {
				t.Errorf("want %v, got %v", test.want, benchmarks)
			}
		})
	}
}

var separatedBenchmarksRead BenchmarkSets

func BenchmarkReadAndSeparateBenchmarks(b *testing.B) {
	benchmarks := []struct {
		name  string
		input io.Reader
	}{
		{
			name:  "single benchmark",
			input: strings.NewReader("BenchmarkSomething/SubBenchmark-12   	   10000	     10000 ns/op"),
		},
		{
			name: "multiple lines",
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
