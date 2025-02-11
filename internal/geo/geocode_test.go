package geo

import (
	"testing"
)

func BenchmarkNeighbors(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Neighbors("dry1", 5)
	}
}
