package tools

import (
	"fmt"
	"os"
	"testing"
)

// Test
func Benchmark(b *testing.B) {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
