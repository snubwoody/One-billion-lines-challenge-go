package main

import "testing"

func BenchmarkMeasurments(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseMeasurements("dummy.txt")
	}
}