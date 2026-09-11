//go:build contract

package heap_sort

import (
	"errors"
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, items := range [][]int{{}, {1}, {3, 1, 2}, {3, 2, 1}, {2, 2, 1}, {4, 10, 3, 5, 1, 2, 8}} {
		ok, err := HeapSort(items, func(a, b int) int { return a - b })
		if !ok || err != nil {
			t.Fatalf("HeapSort() = (%t, %v), want (true, nil)", ok, err)
		}
		for i := 1; i < len(items); i++ {
			if items[i-1] > items[i] {
				t.Fatalf("not sorted: %v", items)
			}
		}
	}
}

func TestHeapSortRejectsNilComparatorWithoutMutation(t *testing.T) {
	items := []int{3, 1, 2}
	want := append([]int(nil), items...)

	ok, err := HeapSort(items, nil)
	if ok || !errors.Is(err, ErrNilComparator) {
		t.Fatalf("HeapSort(nil comparator) = (%t, %v), want (false, ErrNilComparator)", ok, err)
	}
	for i := range items {
		if items[i] != want[i] {
			t.Fatalf("HeapSort(nil comparator) mutated items to %v, want %v", items, want)
		}
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for _, size := range []int{1 << 10, 1 << 15} {
		for _, workload := range []string{"random", "reverse", "duplicates"} {
			b.Run(fmt.Sprintf("%s/%d", workload, size), func(b *testing.B) {
				source := make([]int, size)
				for index := range source {
					switch workload {
					case "reverse":
						source[index] = size - index
					case "duplicates":
						source[index] = index % 16
					default:
						source[index] = (index*7919 + 104729) % size
					}
				}

				items := make([]int, size)
				compare := func(left, right int) int { return left - right }
				b.ReportAllocs()
				b.ResetTimer()
				for iteration := 0; iteration < b.N; iteration++ {
					b.StopTimer()
					copy(items, source)
					b.StartTimer()
					ok, err := HeapSort(items, compare)
					if !ok || err != nil {
						b.Fatalf("HeapSort() = (%t, %v), want (true, nil)", ok, err)
					}
				}
				b.StopTimer()

				for index := 1; index < len(items); index++ {
					if items[index-1] > items[index] {
						b.Fatalf("not sorted: %v", items)
					}
				}
			})
		}
	}
}
