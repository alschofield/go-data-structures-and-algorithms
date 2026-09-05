//go:build contract

package selection_sort

import (
	"fmt"
	"testing"
)

var selectionSortBenchmarkItems []int
var selectionSortBenchmarkOK bool

func BenchmarkSelectionSort(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		for _, test := range []struct {
			name  string
			items []int
		}{
			{name: "random", items: selectionSortRandomInput(size)},
			{name: "sorted", items: selectionSortSortedInput(size)},
			{name: "reverse", items: selectionSortReverseInput(size)},
		} {
			b.Run(fmt.Sprintf("%s/Size=%d", test.name, size), func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					items := append([]int(nil), test.items...)
					ok, err := SelectionSort(items, compareInts)
					if err != nil || !ok {
						b.Fatalf("SelectionSort() = (%t, %v), want (true, nil)", ok, err)
					}
					selectionSortBenchmarkItems = items
					selectionSortBenchmarkOK = ok
				}
				b.StopTimer()

				if !selectionSortBenchmarkOK || !selectionSortIsSorted(selectionSortBenchmarkItems) {
					b.Fatal("SelectionSort() left an unsuccessful or unsorted result")
				}
			})
		}
	}
}

func selectionSortSortedInput(size int) []int {
	items := make([]int, size)
	for i := range items {
		items[i] = i
	}
	return items
}

func selectionSortReverseInput(size int) []int {
	items := selectionSortSortedInput(size)
	for i := range items {
		items[i] = size - 1 - i
	}
	return items
}

func selectionSortRandomInput(size int) []int {
	items := selectionSortSortedInput(size)
	state := uint64(1)
	for i := len(items) - 1; i > 0; i-- {
		state = state*6_364_136_223_846_793_005 + 1
		j := int(state % uint64(i+1))
		items[i], items[j] = items[j], items[i]
	}
	return items
}

func selectionSortIsSorted(items []int) bool {
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			return false
		}
	}
	return true
}
