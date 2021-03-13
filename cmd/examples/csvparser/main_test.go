package main

import (
	"crypto/rand"
	"reflect"
	"testing"
)

func TestParseCSVLine(t *testing.T) {
	tests := []struct{
		name string
		input string
		wanted []string
	}{
		{
			name: "Simple Example",
			input: "a,b,c,d",
			wanted: []string{"a", "b", "c", "d"},
		},
		{
			name: "Fields Within Quotes",
			input: `"a,b",c,d`,
			wanted: []string{`"a,b"`, "c", "d"},
		},
		{
			name: "Fields Within Multiply Nested Quotes",
			input: `"a,"b"",c,d`,
			wanted: []string{`"a,"b""`, "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseCSVLine(tc.input)
			if !reflect.DeepEqual(tc.wanted, got) {
				t.Errorf("wanted %v, got %v", tc.wanted, got)
			}
		})
	}
}

var result []string
func BenchmarkParseCSVLineFields(b *testing.B) {
	benchmarks := []struct {
		name string
		input string
	}{
		{
			name: "10 Fields",
			input: generateCSVLineFields(10),
		},
		{
			name: "20 Fields",
			input: generateCSVLineFields(20),
		},
		{
			name: "40 Fields",
			input: generateCSVLineFields(40),
		},
		{
			name: "80 Fields",
			input: generateCSVLineFields(80),
		},
		{
			name: "160 Fields",
			input: generateCSVLineFields(160),
		},
		{
			name: "320 Fields",
			input: generateCSVLineFields(320),
		},
		{
			name: "640 Fields",
			input: generateCSVLineFields(640),
		},
		{
			name: "1280 Fields",
			input: generateCSVLineFields(1280),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = parseCSVLine(bm.input)
			}
		})
	}
}

func generateCSVLineFields(numFields int) string {
	field := "aaaaa"
	var result string
	for i := 0; i < numFields; i++ {
		result += field
		if i == numFields - 1 {
			result += ","
		}
	}
	return result
}

func BenchmarkParseCSVLineFieldLength(b *testing.B) {
	benchmarks := []struct {
		name string
		input string
	}{
		{
			name: "Length 10",
			input: generateCSVLineFieldLen(b, 10),
		},
		{
			name: "Length 20",
			input: generateCSVLineFieldLen(b, 20),
		},
		{
			name: "Length 40",
			input: generateCSVLineFieldLen(b, 40),
		},
		{
			name: "Length 80",
			input: generateCSVLineFieldLen(b, 80),
		},
		{
			name: "Length 160",
			input: generateCSVLineFieldLen(b, 160),
		},
		{
			name: "Length 320",
			input: generateCSVLineFieldLen(b, 320),
		},
		{
			name: "Length 640",
			input: generateCSVLineFieldLen(b, 640),
		},
		{
			name: "Length 1280",
			input: generateCSVLineFieldLen(b, 1280),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = parseCSVLine(bm.input)
			}
		})
	}
}

func generateCSVLineFieldLen(b *testing.B, fieldLen int) string {
	numFields := 10

	buf := make([]byte, fieldLen)
	_, err := rand.Read(buf)
	if err != nil {
		b.Fatalf("Could not read random data of length %d - error: %v", fieldLen, err)
	}

	var result string
	for i := 0; i < numFields; i++ {
		result += string(buf)
		if i == numFields - 1 {
			result += ","
		}
	}

	return result
}
