package addnumbers

import (
	"math/rand"
	"testing"
	"runtime"
	"time"
)

const N = 10000000
var cache = make([]int, N)

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		cache[i] = rand.Intn(N)
	}
}

func BenchmarkSequential(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(cache)
    }
}

func BenchmarkConcurrent(b *testing.B) {
    for i := 0; i < b.N; i++ {
        addConcurrent(runtime.NumCPU(), cache)
    }
}