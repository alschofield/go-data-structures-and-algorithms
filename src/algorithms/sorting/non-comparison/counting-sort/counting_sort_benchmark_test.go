//go:build contract

package counting_sort

import (
	"fmt"
	"testing"
)

func BenchmarkCountingSort(b *testing.B) {
	for _, test := range []struct {
		size, keyLimit int
	}{
		{256, 64},
		{1_024, 256},
		{4_096, 1_024},
	} {
		b.Run(fmt.Sprintf("Size=%d/Keys=%d", test.size, test.keyLimit), func(b *testing.B) {
			input := countingRandom(test.size, test.keyLimit)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				items := append([]uint32(nil), input...)
				if ok, err := CountingSort(items, uint32(test.keyLimit)); !ok || err != nil {
					b.Fatalf("CountingSort() = (%t, %v), want (true, nil)", ok, err)
				}
			}
		})
	}
}

func countingRandom(size, keyLimit int) []uint32 {
	items := make([]uint32, size)
	state := uint64(1)
	for i := range items {
		state = state*6_364_136_223_846_793_005 + 1
		items[i] = uint32(state % uint64(keyLimit))
	}
	return items
}
