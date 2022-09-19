package main

import (
	"fmt"
	"os"
	"testing"
)

func Double(x int) (result int) {
	result = x + 2
	return result
}

func BenchmarkDouble(b *testing.B) {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Double(10)
	}
}
