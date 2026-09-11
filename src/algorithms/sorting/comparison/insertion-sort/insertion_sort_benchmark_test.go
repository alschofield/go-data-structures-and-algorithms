//go:build contract

package insertion_sort

import (
	"fmt"
	"testing"
)

func BenchmarkInsertionSort(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		for _, input := range []struct {
			name  string
			items []int
		}{
			{"sorted", insertionSorted(size)},
			{"nearly-sorted", insertionNearlySorted(size)},
			{"reverse", insertionReverse(size)},
		} {
			b.Run(fmt.Sprintf("%s/Size=%d", input.name, size), func(b *testing.B) {
				items := make([]int, len(input.items))
				compare := func(left, right int) int { return left - right }
				b.ReportAllocs()
				b.ResetTimer()
				for iteration := 0; iteration < b.N; iteration++ {
					copy(items, input.items)
					ok, err := InsertionSort(items, compare)
					if !ok || err != nil {
						b.Fatalf("InsertionSort() = (%t, %v), want (true, nil)", ok, err)
					}
				}
				for index := 1; index < len(items); index++ {
					if items[index-1] > items[index] {
						b.Fatalf("not sorted: %v", items)
					}
				}
			})
		}
	}
}

func insertionSorted(size int) []int {
	items := make([]int, size)
	for index := range items {
		items[index] = index
	}
	return items
}

func insertionNearlySorted(size int) []int {
	items := insertionSorted(size)
	for index := 32; index < len(items); index += 32 {
		items[index], items[index-1] = items[index-1], items[index]
	}
	return items
}

func insertionReverse(size int) []int {
	items := insertionSorted(size)
	for index := range items {
		items[index] = size - 1 - index
	}
	return items
}
