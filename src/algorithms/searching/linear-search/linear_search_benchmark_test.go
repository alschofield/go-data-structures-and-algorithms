package linear_search

import (
	"strconv"
	"testing"
)

var linearSearchBenchmarkIndex int
var linearSearchBenchmarkFound bool

func BenchmarkLinearSearchPresentLast(b *testing.B) {
	for _, size := range []int{1_000, 16_000, 1_000_000} {
		b.Run("N="+strconv.Itoa(size), func(b *testing.B) {
			items := benchmarkItems(size)
			key := items[len(items)-1]

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				index, found, err := LinearSearch(items, key, compareInts)
				if err != nil || !found || index != len(items)-1 {
					b.Fatalf("LinearSearch() = (%d, %t, %v), want (%d, true, nil)", index, found, err, len(items)-1)
				}
				linearSearchBenchmarkIndex = index
				linearSearchBenchmarkFound = found
			}
		})
	}
}

func benchmarkItems(size int) []int {
	items := make([]int, size)
	for i := range items {
		items[i] = i * 2
	}
	return items
}

func compareInts(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
