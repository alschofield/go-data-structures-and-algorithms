//go:build contract

package linear_search

import "testing"

func TestLinearSearch(t *testing.T) {
	compare := func(a, b int) int { return a - b }
	items := []int{4, 2, 4, 1}
	for _, test := range []struct {
		key, index int
		ok         bool
	}{{4, 0, true}, {1, 3, true}, {9, 0, false}} {
		got, ok := LinearSearch(items, test.key, compare)
		if got != test.index || ok != test.ok {
			t.Fatalf("key %d: got (%d, %t), want (%d, %t)", test.key, got, ok, test.index, test.ok)
		}
	}
	if items[0] != 4 {
		t.Fatal("LinearSearch must not modify input")
	}
}
