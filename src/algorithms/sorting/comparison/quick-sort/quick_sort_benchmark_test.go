//go:build contract

package quick_sort

import (
	"fmt"
	"testing"
)

var quickSortBenchmarkItems []int
var quickSortBenchmarkOK bool

func BenchmarkQuickSort(b *testing.B) {
	for _, size := range []int{1_024, 8_192} {
		for _, test := range []struct {
			name  string
			items []int
		}{
			{name: "random", items: quickSortRandomInput(size)},
			{name: "sorted", items: quickSortSortedInput(size)},
			{name: "reverse", items: quickSortReverseInput(size)},
			{name: "all-equal", items: quickSortAllEqualInput(size)},
		} {
			b.Run(fmt.Sprintf("%s/Size=%d", test.name, size), func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					items := append([]int(nil), test.items...)
					ok, err := QuickSort(items, compareInts)
					if err != nil || !ok {
						b.Fatalf("QuickSort() = (%t, %v), want (true, nil)", ok, err)
					}
					quickSortBenchmarkItems = items
					quickSortBenchmarkOK = ok
				}
				b.StopTimer()

				if !quickSortBenchmarkOK || !quickSortIsSorted(quickSortBenchmarkItems) {
					b.Fatal("QuickSort() left an unsuccessful or unsorted result")
				}
			})
		}
	}
}

func quickSortSortedInput(size int) []int {
	items := make([]int, size)
	for i := range items {
		items[i] = i
	}
	return items
}

func quickSortReverseInput(size int) []int {
	items := quickSortSortedInput(size)
	for i := range items {
		items[i] = size - 1 - i
	}
	return items
}

func quickSortAllEqualInput(size int) []int {
	return make([]int, size)
}

func quickSortRandomInput(size int) []int {
	items := quickSortSortedInput(size)
	state := uint64(1)
	for i := len(items) - 1; i > 0; i-- {
		state = state*6_364_136_223_846_793_005 + 1
		j := int(state % uint64(i+1))
		items[i], items[j] = items[j], items[i]
	}
	return items
}

func quickSortIsSorted(items []int) bool {
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			return false
		}
	}
	return true
}
