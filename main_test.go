package main

import (
	"strings"
	"testing"
)

func BenchmarkMeasurments(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseMeasurements("text-files/test.txt")
	}
}

func BenchmarkStringSplit(b *testing.B) {
	for i := 0; i < b.N; i++{
		strings.Split("Vancouver;90.0",";")
	}
}
func BenchmarkStringIndex(b *testing.B) {
	for i := 0; i < b.N; i++{
		strings.Index("Vancouver;90.0",";")
	}
}