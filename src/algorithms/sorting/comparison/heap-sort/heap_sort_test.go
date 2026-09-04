//go:build contract

package heap_sort

import "testing"

func TestHeapSort(t *testing.T) {
	for _, items := range [][]int{{}, {1}, {3, 1, 2}, {3, 2, 1}, {2, 2, 1}} {
		if err := HeapSort(items, func(a, b int) int { return a - b }); err != nil {
			t.Fatalf("sort error = %v", err)
		}
		for i := 1; i < len(items); i++ {
			if items[i-1] > items[i] {
				t.Fatalf("not sorted: %v", items)
			}
		}
	}
}
