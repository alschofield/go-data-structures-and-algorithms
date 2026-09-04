//go:build contract

package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	compare := func(a, b int) int { return a - b }
	items := []int{1, 3, 3, 5, 8}
	for _, test := range []struct {
		key int
		ok  bool
	}{{1, true}, {3, true}, {8, true}, {4, false}} {
		index, ok := BinarySearch(items, test.key, compare)
		if ok != test.ok || (ok && items[index] != test.key) {
			t.Fatalf("key %d: got (%d, %t)", test.key, index, ok)
		}
	}
	if _, ok := BinarySearch([]int(nil), 1, compare); ok {
		t.Fatal("empty input must be absent")
	}
}
