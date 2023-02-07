package main

import (
	"bytes"
	"testing"
)

func BenchmarkSVG(b *testing.B) {
	b.Run("benchmark syncron svg", func(b *testing.B) {
		buff := &bytes.Buffer{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			SVG(buff)
		}
	})
	b.Run("benchmark concurrent svg", func(b *testing.B) {
		buff := &bytes.Buffer{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			SVG2(buff)
		}
	})
}
