//go:build contract

package binary_search

import (
	"errors"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	compare := func(a, b int) int { return a - b }
	items := []int{1, 3, 3, 5, 8}
	for _, test := range []struct {
		key int
		ok  bool
	}{{1, true}, {3, true}, {8, true}, {4, false}} {
		index, ok, err := BinarySearch(items, test.key, compare)
		if err != nil || ok != test.ok || (ok && items[index] != test.key) {
			t.Fatalf("key %d: got (%d, %t, %v)", test.key, index, ok, err)
		}
	}
	if _, ok, err := BinarySearch([]int(nil), 1, compare); err != nil || ok {
		t.Fatal("empty input must be absent")
	}
	if _, _, err := BinarySearch(items, 1, nil); !errors.Is(err, ErrNilComparator) {
		t.Fatalf("nil comparator error = %v, want ErrNilComparator", err)
	}
}
