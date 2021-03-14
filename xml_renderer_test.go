package go_benchpress

import (
	"bytes"
	"golang.org/x/tools/benchmark/parse"
	"testing"
)

func TestXMLRenderer_Render(t *testing.T) {
	tests := []struct {
		name       string
		benchmarks []parse.Benchmark
		want       string
	}{
		{
			name:            "single benchmark",
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmark",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
			},
			want: `<xmlBenchmarkRecord><ParentBenchmark>ParentBenchmark</ParentBenchmark><Benchmarks><Name>BenchmarkOne/SubBenchmark</Name><N>100</N><NsPerOp>1000</NsPerOp><AllocedBytesPerOp>10000</AllocedBytesPerOp><AllocsPerOp>100000</AllocsPerOp><MBPerS>1e+06</MBPerS><Measured>10000000</Measured><Ord>100000000</Ord></Benchmarks></xmlBenchmarkRecord>`,
		},
		{
			name:            "multiple benchmarks",
			benchmarks: []parse.Benchmark{
				{
					Name:              "BenchmarkOne/SubBenchmarkOne",
					N:                 100,
					NsPerOp:           1000,
					AllocedBytesPerOp: 10000,
					AllocsPerOp:       100000,
					MBPerS:            1000000,
					Measured:          10000000,
					Ord:               100000000,
				},
				{
					Name:              "BenchmarkOne/SubBenchmarkTwo",
					N:                 100000000,
					NsPerOp:           10000000,
					AllocedBytesPerOp: 1000000,
					AllocsPerOp:       100000,
					MBPerS:            10000,
					Measured:          1000,
					Ord:               100,
				},
			},
			want: `<xmlBenchmarkRecord><ParentBenchmark>ParentBenchmark</ParentBenchmark><Benchmarks><Name>BenchmarkOne/SubBenchmarkOne</Name><N>100</N><NsPerOp>1000</NsPerOp><AllocedBytesPerOp>10000</AllocedBytesPerOp><AllocsPerOp>100000</AllocsPerOp><MBPerS>1e+06</MBPerS><Measured>10000000</Measured><Ord>100000000</Ord></Benchmarks><Benchmarks><Name>BenchmarkOne/SubBenchmarkTwo</Name><N>100000000</N><NsPerOp>1e+07</NsPerOp><AllocedBytesPerOp>1000000</AllocedBytesPerOp><AllocsPerOp>100000</AllocsPerOp><MBPerS>10000</MBPerS><Measured>1000</Measured><Ord>100</Ord></Benchmarks></xmlBenchmarkRecord>`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer
			renderer := XMLRenderer{}
			err := renderer.Render(&output, "ParentBenchmark", RenderNsPerOp, test.benchmarks)
			if err != nil {
				t.Fatalf("Error rendering XML - error: %v", err)
			}

			got := output.String()
			if test.want != got {
				t.Errorf("Want %q, got %q", test.want, got)
			}
		})
	}
}
