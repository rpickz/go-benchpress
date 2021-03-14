package go_benchpress

import (
	"bytes"
	"golang.org/x/tools/benchmark/parse"
	"testing"
)

func TestJSONRenderer_Render(t *testing.T) {
	tests := []struct {
		name            string
		parentBenchmark string
		benchmarks      []parse.Benchmark
		want            string
	}{
		{
			name:            "single benchmark",
			parentBenchmark: "BenchmarkOne",
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
			want: `{"ParentBenchmark":"BenchmarkOne","Benchmarks":[{"Name":"BenchmarkOne/SubBenchmark","N":100,"NsPerOp":1000,"AllocedBytesPerOp":10000,"AllocsPerOp":100000,"MBPerS":1000000,"Measured":10000000,"Ord":100000000}]}`,
		},
		{
			name:            "multiple benchmarks",
			parentBenchmark: "BenchmarkOne",
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
			want: `{"ParentBenchmark":"BenchmarkOne","Benchmarks":[{"Name":"BenchmarkOne/SubBenchmarkOne","N":100,"NsPerOp":1000,"AllocedBytesPerOp":10000,"AllocsPerOp":100000,"MBPerS":1000000,"Measured":10000000,"Ord":100000000},{"Name":"BenchmarkOne/SubBenchmarkTwo","N":100000000,"NsPerOp":10000000,"AllocedBytesPerOp":1000000,"AllocsPerOp":100000,"MBPerS":10000,"Measured":1000,"Ord":100}]}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer
			renderer := JSONRenderer{}
			err := renderer.Render(&output, test.parentBenchmark, RenderNsPerOp, test.benchmarks)
			if err != nil {
				t.Fatalf("Error rendering JSON: %v", err)
			}
			got := output.String()
			if test.want != got {
				t.Errorf("want %q, got %q", test.want, got)
			}
		})
	}
}

func BenchmarkJSONRenderer_Render(b *testing.B) {

	benchmark := parse.Benchmark{
		Name:              "BenchmarkOne/SubBenchmarkOne",
		N:                 100,
		NsPerOp:           1000,
		AllocedBytesPerOp: 10000,
		AllocsPerOp:       100000,
		MBPerS:            1000000,
		Measured:          10000000,
		Ord:               100000000,
	}

	benchmarks := []struct {
		name            string
		parentBenchmark string
		benchmarks      []parse.Benchmark
	}{

		{
			name:            "single benchmark",
			parentBenchmark: "BenchmarkOne",
			benchmarks: []parse.Benchmark{benchmark},
		},
		{
			name:            "2 benchmarks",
			parentBenchmark: "BenchmarkOne",
			benchmarks: []parse.Benchmark{benchmark, benchmark},
		},
		{
			name:            "4 benchmarks",
			parentBenchmark: "BenchmarkOne",
			benchmarks: []parse.Benchmark{benchmark, benchmark, benchmark, benchmark},
		},
		{
			name:            "8 benchmarks",
			parentBenchmark: "BenchmarkOne",
			benchmarks: []parse.Benchmark{benchmark, benchmark, benchmark, benchmark, benchmark, benchmark, benchmark, benchmark},
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var output bytes.Buffer
				renderer := JSONRenderer{}
				_ = renderer.Render(&output, bm.parentBenchmark, RenderNsPerOp, bm.benchmarks)
			}
		})
	}
}
