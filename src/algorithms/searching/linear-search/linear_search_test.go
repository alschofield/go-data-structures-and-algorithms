//go:build contract

package linear_search

import (
	"errors"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	compare := func(a, b int) int { return a - b }
	items := []int{4, 2, 4, 1}
	for _, test := range []struct {
		key, index int
		ok         bool
	}{{4, 0, true}, {1, 3, true}, {9, 0, false}} {
		got, ok, err := LinearSearch(items, test.key, compare)
		if err != nil || got != test.index || ok != test.ok {
			t.Fatalf("key %d: got (%d, %t, %v), want (%d, %t, nil)", test.key, got, ok, err, test.index, test.ok)
		}
	}
	if _, _, err := LinearSearch(items, 4, nil); !errors.Is(err, ErrNilComparator) {
		t.Fatalf("nil comparator error = %v, want ErrNilComparator", err)
	}
	if items[0] != 4 {
		t.Fatal("LinearSearch must not modify input")
	}
}
