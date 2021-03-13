package benchpress

import (
	"errors"
	"testing"
)

func TestRenderDimension_String(t *testing.T) {
	tests := []struct {
		name  string
		input RenderDimension
		want  string
	}{
		{
			name:  "ns per op",
			input: RenderNsPerOp,
			want:  "Ns Per Op",
		},
		{
			name:  "bytes per op",
			input: RenderBytesPerOp,
			want:  "Bytes Per Op",
		},
		{
			name:  "allocs per op",
			input: RenderAllocsPerOp,
			want:  "Allocs Per Op",
		},
		{
			name:  "unknown",
			input: RenderDimension(1000),
			want:  "Unknown (1000)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.String()
			if test.want != got {
				t.Errorf("want %q, got %q", test.want, got)
			}
		})
	}
}

func TestRenderDimensionFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    RenderDimension
		wantErr error
	}{
		{
			name: "ns per op",
			input: "NS_PER_OP",
			want: RenderNsPerOp,
		},
		{
			name: "bytes per op",
			input: "BYTES_PER_OP",
			want: RenderBytesPerOp,
		},
		{
			name: "allocs per op",
			input: "ALLOCS_PER_OP",
			want: RenderAllocsPerOp,
		},
		{
			name: "unknown",
			input: "abc123",
			want: RenderDimension(-1),
			wantErr: ErrUnknownDimensionType,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := RenderDimensionFromString(test.input)
			if err != nil {
				if !errors.Is(err, test.wantErr) {
					t.Errorf("Want error '%v', got error '%v'", test.wantErr, err)
				}
			}
			if test.want != got {
				t.Errorf("want %q, got %q", test.want, got)
			}
		})
	}
}