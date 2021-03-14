package go_benchpress

import (
	"errors"
	"reflect"
	"testing"
)

// ===== RenderType tests =====

func TestRenderType_String(t *testing.T) {
	tests := []struct {
		name  string
		input RenderType
		want  string
	}{
		{
			name:  "png",
			input: PNG,
			want:  "PNG",
		},
		{
			name:  "svg",
			input: SVG,
			want:  "SVG",
		},
		{
			name:  "json",
			input: JSON,
			want:  "JSON",
		},
		{
			name:  "csv",
			input: CSV,
			want:  "CSV",
		},
		{
			name:  "xml",
			input: XML,
			want:  "XML",
		},
		{
			name:  "unknown",
			input: RenderType(1000),
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

func TestRenderTypeFromString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    RenderType
		wantErr error
	}{
		{
			name:  "png",
			input: "PNG",
			want:  PNG,
		},
		{
			name:  "svg",
			input: "SVG",
			want:  SVG,
		},
		{
			name:  "json",
			input: "JSON",
			want:  JSON,
		},
		{
			name:  "csv",
			input: "CSV",
			want:  CSV,
		},
		{
			name:  "xml",
			input: "XML",
			want:  XML,
		},
		{
			name:    "unknown",
			input:   "abc123",
			want:    RenderType(-1),
			wantErr: ErrUnknownRenderType,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := RenderTypeFromString(test.input)
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

func TestRenderType_Renderer(t *testing.T) {

	type wantCmpFunc func(t *testing.T, got Renderer)

	rasterWantCmp := func(format RenderType) wantCmpFunc {
		return func(t *testing.T, got Renderer) {
			raster, ok := got.(*RasterRenderer)
			if !ok {
				t.Fatal("Could not convert renderer to RasterRenderer")
			}
			want := *NewRasterRenderer("title", format)
			// Set render funcs to nil to make comparable.
			want.barChartRenderFunc = nil
			raster.barChartRenderFunc = nil

			if !reflect.DeepEqual(want, *raster) {
				t.Errorf("Wanted %v, got %v", want, *raster)
			}
		}
	}

	tests := []struct {
		name  string
		input RenderType
		// wantCmp compares the want with the got, calling Error or Fatal on testing.T if failure.
		wantCmp wantCmpFunc
		wantErr error
	}{
		{
			name: "png",
			input: PNG,
			wantCmp: rasterWantCmp(PNG),
		},
		{
			name: "svg",
			input: SVG,
			wantCmp: rasterWantCmp(SVG),
		},
		{
			name: "json",
			input: JSON,
			wantCmp: func(t *testing.T, got Renderer) {
				raster, ok := got.(*JSONRenderer)
				if !ok {
					t.Fatal("Could not convert renderer to JSONRenderer")
				}
				want := JSONRenderer{}

				if !reflect.DeepEqual(want, *raster) {
					t.Errorf("Wanted %v, got %v", want, *raster)
				}
			},
		},
		{
			name: "csv",
			input: CSV,
			wantCmp: func(t *testing.T, got Renderer) {
				raster, ok := got.(*CSVRenderer)
				if !ok {
					t.Fatal("Could not convert renderer to CSVRenderer")
				}
				want := CSVRenderer{}

				if !reflect.DeepEqual(want, *raster) {
					t.Errorf("Wanted %v, got %v", want, *raster)
				}
			},
		},
		{
			name:  "xml",
			input: XML,
			wantCmp: func(t *testing.T, got Renderer) {
				raster, ok := got.(*XMLRenderer)
				if !ok {
					t.Fatal("Could not convert renderer to XMLRenderer")
				}
				want := XMLRenderer{}

				if !reflect.DeepEqual(want, *raster) {
					t.Errorf("Wanted %v, got %v", want, *raster)
				}
			},
		},
		{
			name: "unknown",
			input: RenderType(1000),
			wantErr: ErrUnknownRenderType,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.input.Renderer("title")
			if err != nil {
				if !errors.Is(err, test.wantErr) {
					t.Errorf("Wanted error '%v', got error '%v'", test.wantErr, err)
				}
				return
			}

			test.wantCmp(t, got)
		})
	}
}

func TestRenderType_FileExtension(t *testing.T) {
	tests := []struct {
		name string
		input RenderType
		want string
	}{
		{
			name: "png",
			input: PNG,
			want: ".png",
		},
		{
			name: "svg",
			input: SVG,
			want: ".svg",
		},
		{
			name: "json",
			input: JSON,
			want: ".json",
		},
		{
			name: "csv",
			input: CSV,
			want: ".csv",
		},
		{
			name: "xml",
			input: XML,
			want: ".xml",
		},
		{
			name: "unknown",
			input: RenderType(1000),
			want: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.input.FileExtension()
			if test.want != got {
				t.Errorf("Wanted %q, got %q", test.want, got)
			}
		})
	}
}

// ===== RenderDimension tests =====

func TestRenderDimension_String(t *testing.T) {
	tests := []struct {
		name  string
		input RenderDimension
		want  string
	}{
		{
			name:  "ns per op",
			input: RenderNsPerOp,
			want:  "NS_PER_OP",
		},
		{
			name:  "bytes per op",
			input: RenderBytesPerOp,
			want:  "BYTES_PER_OP",
		},
		{
			name:  "allocs per op",
			input: RenderAllocsPerOp,
			want:  "ALLOCS_PER_OP",
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
			name:  "ns per op",
			input: "NS_PER_OP",
			want:  RenderNsPerOp,
		},
		{
			name:  "bytes per op",
			input: "BYTES_PER_OP",
			want:  RenderBytesPerOp,
		},
		{
			name:  "allocs per op",
			input: "ALLOCS_PER_OP",
			want:  RenderAllocsPerOp,
		},
		{
			name:    "unknown",
			input:   "abc123",
			want:    RenderDimension(-1),
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
