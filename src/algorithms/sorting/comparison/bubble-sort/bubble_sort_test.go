//go:build contract

package bubble_sort

import (
	"errors"
	"testing"
)

func TestBubbleSort(t *testing.T) { testComparisonSort(t, BubbleSort[int], ErrNilComparator) }

func TestBubbleSortPreservesEqualOrder(t *testing.T) {
	type item struct{ key, order int }
	items := []item{{2, 0}, {1, 1}, {2, 2}}
	_, err := BubbleSort(items, func(left, right item) int { return left.key - right.key })
	if err != nil || items[1].order != 0 || items[2].order != 2 {
		t.Fatalf("BubbleSort() did not preserve equal order: %#v", items)
	}
}

func testComparisonSort(t *testing.T, sort func([]int, func(int, int) int) (bool, error), nilComparatorError error) {
	t.Helper()

	for _, test := range [][]int{{}, {1}, {3, 1, 2}, {3, 2, 1}, {2, 2, 1}} {
		ok, err := sort(test, func(a, b int) int { return a - b })
		if !ok || err != nil {
			t.Fatalf("sort() = (%t, %v), want (true, nil)", ok, err)
		}
		for i := 1; i < len(test); i++ {
			if test[i-1] > test[i] {
				t.Fatalf("not sorted: %v", test)
			}
		}
	}

	items := []int{3, 1, 2}
	want := append([]int(nil), items...)
	ok, err := sort(items, nil)
	if ok || !errors.Is(err, nilComparatorError) {
		t.Fatalf("sort(nil comparator) = (%t, %v), want (false, ErrNilComparator)", ok, err)
	}
	for i := range items {
		if items[i] != want[i] {
			t.Fatalf("nil comparator mutated items to %v, want %v", items, want)
		}
	}
}
