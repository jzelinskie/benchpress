package tools

import (
	"fmt"
	"os"
	"testing"
)

func Benchmark(b *testing.B) {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
