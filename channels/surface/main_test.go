package main

import (
	"bytes"
	"testing"
)

func BenchmarkSVG(b *testing.B) {
	buff := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SVG(buff)
	}
}
