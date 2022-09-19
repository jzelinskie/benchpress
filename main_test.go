package main

import "testing"

func Double(x int) (result int) {
	result = x + 2
	return result
}

func BenchmarkDouble(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Double(10)
	}
}
