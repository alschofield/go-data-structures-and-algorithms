//go:build contract

package merge_sort

import (
	"fmt"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	for _, size := range []int{256, 1_024, 4_096} {
		b.Run(fmt.Sprintf("random/Size=%d", size), func(b *testing.B) {
			input := mergeRandom(size)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				items := append([]int(nil), input...)
				if ok, err := MergeSort(items, func(a, b int) int { return a - b }); !ok || err != nil {
					b.Fatalf("MergeSort() = (%t, %v), want (true, nil)", ok, err)
				}
			}
		})
	}
}

func mergeRandom(size int) []int {
	items := make([]int, size)
	state := uint64(1)
	for i := range items {
		state = state*6_364_136_223_846_793_005 + 1
		items[i] = int(state % uint64(size))
	}
	return items
}
