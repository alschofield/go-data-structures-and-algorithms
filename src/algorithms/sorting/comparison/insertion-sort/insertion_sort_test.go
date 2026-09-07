//go:build contract

package insertion_sort

import (
	"errors"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, items := range [][]int{{}, {1}, {3, 1, 2}, {3, 2, 1}, {2, 2, 1}} {
		ok, err := InsertionSort(items, func(a, b int) int { return a - b })
		if !ok || err != nil {
			t.Fatalf("InsertionSort() = (%t, %v), want (true, nil)", ok, err)
		}
		for i := 1; i < len(items); i++ {
			if items[i-1] > items[i] {
				t.Fatalf("not sorted: %v", items)
			}
		}
	}

	items := []int{3, 1, 2}
	want := append([]int(nil), items...)
	ok, err := InsertionSort[int](items, nil)
	if ok || !errors.Is(err, ErrNilComparator) {
		t.Fatalf("InsertionSort(nil comparator) = (%t, %v), want (false, ErrNilComparator)", ok, err)
	}
	for i := range items {
		if items[i] != want[i] {
			t.Fatalf("nil comparator mutated items to %v, want %v", items, want)
		}
	}
}
