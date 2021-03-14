package go_benchpress

import (
	"fmt"
	"golang.org/x/tools/benchmark/parse"
	"io"
)

// ===== RenderType =====

type RenderType int

const (
	PNG RenderType = iota
	SVG
	JSON
	CSV
)

func (r RenderType) String() string {
	switch r {
	case PNG:
		return "PNG"
	case SVG:
		return "SVG"
	case JSON:
		return "JSON"
	case CSV:
		return "CSV"
	default:
		return fmt.Sprintf("Unknown (%d)", r)
	}
}

// Renderer provides an instance of a Renderer for the RenderType - for instance, a JSONRenderer for JSON.
// If there is no matching Renderer for the RenderType, an ErrUnknownRenderType is returned.
func (r RenderType) Renderer(title string) (Renderer, error) {
	switch r {
	case PNG, SVG:
		return NewRasterRenderer(title, r), nil
	case JSON:
		return &JSONRenderer{}, nil
	case CSV:
		return &CSVRenderer{}, nil
	default:
		return nil, ErrUnknownRenderType
	}
}

func (r RenderType) FileExtension() string {
	switch r {
	case PNG:
		return ".png"
	case SVG:
		return ".svg"
	case JSON:
		return ".json"
	case CSV:
		return ".csv"
	default:
		return ""
	}
}

func RenderTypeFromString(str string) (RenderType, error) {
	switch str {
	case "PNG":
		return PNG, nil
	case "SVG":
		return SVG, nil
	case "JSON":
		return JSON, nil
	case "CSV":
		return CSV, nil
	default:
		return -1, fmt.Errorf("render type %q not supported: %w", str, ErrUnknownRenderType)
	}
}

// ===== RenderDimension =====

type RenderDimension int

const (
	RenderNsPerOp RenderDimension = iota
	RenderBytesPerOp
	RenderAllocsPerOp
)

func (r RenderDimension) String() string {
	switch r {
	case RenderNsPerOp:
		return "NS_PER_OP"
	case RenderBytesPerOp:
		return "BYTES_PER_OP"
	case RenderAllocsPerOp:
		return "ALLOCS_PER_OP"
	default:
		return fmt.Sprintf("Unknown (%d)", r)
	}
}

func RenderDimensionFromString(str string) (RenderDimension, error) {
	switch str {
	case "NS_PER_OP":
		return RenderNsPerOp, nil
	case "BYTES_PER_OP":
		return RenderBytesPerOp, nil
	case "ALLOCS_PER_OP":
		return RenderAllocsPerOp, nil
	default:
		return -1, fmt.Errorf("render dimension %q not supported: %w", str, ErrUnknownDimensionType)
	}
}

// ===== Renderer =====

type Renderer interface {
	Render(writer io.Writer, parentBenchmark string, dimension RenderDimension, benchmarks []parse.Benchmark) error
}
