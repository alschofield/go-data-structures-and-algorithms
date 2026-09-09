//go:build contract

package radix_sort

import (
	"fmt"
	"testing"
)

func BenchmarkRadixSort(b *testing.B) {
	for _, size := range []int{256, 1_024, 4_096} {
		b.Run(fmt.Sprintf("random/Size=%d", size), func(b *testing.B) {
			input := radixRandom(size)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				items := append([]uint32(nil), input...)
				if !RadixSort(items) {
					b.Fatal("RadixSort() = false, want true")
				}
			}
		})
	}
}

func radixRandom(size int) []uint32 {
	items := make([]uint32, size)
	state := uint64(1)
	for i := range items {
		state = state*6_364_136_223_846_793_005 + 1
		items[i] = uint32(state)
	}
	return items
}
