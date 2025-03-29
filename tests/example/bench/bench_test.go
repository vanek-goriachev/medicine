package bench_test

import (
	"strconv"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	b.Log("This is a bench example test")

	for range b.N {
		if x := strconv.Itoa(0); x != "0" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}
