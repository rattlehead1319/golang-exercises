package iodocs

import (
	"runtime"
	"testing"
)
func BenchmarkSequential(b *testing.B) {
	docs := generateList(1e3)
	for i := 0; i < b.N; i++ {
		find("test", docs)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	docs := generateList(1e3)
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "test", docs)
	}
}
