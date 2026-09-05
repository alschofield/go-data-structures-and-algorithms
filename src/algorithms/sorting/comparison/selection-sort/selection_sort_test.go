//go:build contract

package selection_sort

import (
	"errors"
	"testing"
)

func TestSelectionSortAscending(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name  string
		items []int
		want  []int
	}{
		{name: "empty", items: []int{}, want: []int{}},
		{name: "singleton", items: []int{7}, want: []int{7}},
		{name: "already sorted", items: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
		{name: "reverse sorted", items: []int{4, 3, 2, 1}, want: []int{1, 2, 3, 4}},
		{name: "duplicates", items: []int{3, 1, 3, 2, 1}, want: []int{1, 1, 2, 3, 3}},
		{name: "negative values", items: []int{0, -3, 5, -1, 2}, want: []int{-3, -1, 0, 2, 5}},
	} {
		t.Run(test.name, func(t *testing.T) {
			items := append([]int(nil), test.items...)

			ok, err := SelectionSort(items, compareInts)

			if err != nil || !ok {
				t.Fatalf("SelectionSort() = (%t, %v), want (true, nil)", ok, err)
			}
			if !selectionSortEqual(items, test.want) {
				t.Errorf("SelectionSort() items = %v, want %v", items, test.want)
			}
		})
	}
}

func TestSelectionSortNilComparatorDoesNotMutate(t *testing.T) {
	t.Parallel()

	items := []int{3, 1, 2}
	want := append([]int(nil), items...)

	ok, err := SelectionSort(items, nil)

	if ok {
		t.Error("SelectionSort() ok = true, want false")
	}
	if !errors.Is(err, ErrNilComparator) {
		t.Errorf("SelectionSort() error = %v, want ErrNilComparator", err)
	}
	if !selectionSortEqual(items, want) {
		t.Errorf("SelectionSort() mutated items to %v, want %v", items, want)
	}
}

func compareInts(left, right int) int {
	if left < right {
		return -1
	}
	if left > right {
		return 1
	}
	return 0
}

func selectionSortEqual(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
