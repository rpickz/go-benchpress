package go_benchpress

import (
	"bytes"
	"errors"
	"golang.org/x/tools/benchmark/parse"
	"testing"
)

func TestCSVRenderer_Render(t *testing.T) {
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
			want: `Name,N,NsPerOp,AllocedBytesPerOp,AllocsPerOp,MBPerS,Measured,Ord
BenchmarkOne/SubBenchmark,100,1000.000000000000,10000,100000,1000000.000000000000,10000000,100000000
`,
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
			want: `Name,N,NsPerOp,AllocedBytesPerOp,AllocsPerOp,MBPerS,Measured,Ord
BenchmarkOne/SubBenchmarkOne,100,1000.000000000000,10000,100000,1000000.000000000000,10000000,100000000
BenchmarkOne/SubBenchmarkTwo,100000000,10000000.000000000000,1000000,100000,10000.000000000000,1000,100
`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer
			renderer := CSVRenderer{}
			err := renderer.Render(&output, test.parentBenchmark, RenderNsPerOp, test.benchmarks)
			if err != nil {
				t.Fatalf("Error rendering CSV: %v", err)
			}
			got := output.String()
			if test.want != got {
				t.Errorf("want %q, got %q", test.want, got)
			}
		})
	}
}

func TestCSVRenderer_Render_WriteErrors(t *testing.T) {

	wantErr := errors.New("something went wrong")

	benchmarks := []parse.Benchmark{
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
	}

	writer := errWriter{
		replyWith:  wantErr,
	}
	renderer := CSVRenderer{}

	err := renderer.Render(&writer, "ParentBenchmark", RenderNsPerOp, benchmarks)
	if err == nil {
		t.Fatal("Despite writer error, no error reported")
	}

	if !errors.Is(err, wantErr) {
		t.Errorf("Wanted error '%v', got error '%v'", wantErr, err)
	}
}

func BenchmarkCSVRenderer_Render(b *testing.B) {

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
				renderer := CSVRenderer{}
				_ = renderer.Render(&output, bm.parentBenchmark, RenderNsPerOp, bm.benchmarks)
			}
		})
	}
}

// ===== errWriter =====

type errWriter struct {
	errAfter int
	writeCount int
	replyWith error
}

func (e errWriter) Write(p []byte) (n int, err error) {
	e.writeCount++
	if e.writeCount > e.errAfter {
		return 0, e.replyWith
	}
	return 0, nil
}
