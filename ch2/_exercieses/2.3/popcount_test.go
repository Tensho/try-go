// $ go test -cpu=4 -bench=. popcount_test.go

package popcount_test

import (
	"testing"

	"./popcountmap"
	"./popcountloop"
)

func BenchmarkPopCountMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountmap.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcountloop.PopCount(0x1234567890ABCDEF)
	}
}
