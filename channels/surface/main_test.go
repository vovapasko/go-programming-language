package main

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"testing"
)

func TestSVG(t *testing.T) {
	t.Run("run sync svg function", func(t *testing.T) {
		buff := &bytes.Buffer{}
		SVG(buff)
		approvals.VerifyString(t, buff.String())
	})
	t.Run("run concurrent svg function", func(t *testing.T) {
		buff := &bytes.Buffer{}
		SVG2(buff)
		approvals.VerifyString(t, buff.String())
	})
}

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
