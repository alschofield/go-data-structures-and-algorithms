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
	if got, ok, err := LinearSearch([]int(nil), 4, compare); got != 0 || ok || err != nil {
		t.Fatalf("nil input = (%d, %t, %v), want (0, false, nil)", got, ok, err)
	}
	before := append([]int(nil), items...)
	if got, ok, err := LinearSearch(items, 4, nil); got != 0 || ok || !errors.Is(err, ErrNilComparator) {
		t.Fatalf("nil comparator = (%d, %t, %v), want (0, false, ErrNilComparator)", got, ok, err)
	}
	for index := range items {
		if items[index] != before[index] {
			t.Fatal("LinearSearch must not modify input")
		}
	}
}

func TestLinearSearchGenericComparator(t *testing.T) {
	type record struct {
		id   int
		name string
	}
	items := []record{{id: 1, name: "Ada"}, {id: 2, name: "Grace"}, {id: 1, name: "Linus"}}
	compareByID := func(left, right record) int { return left.id - right.id }

	if got, ok, err := LinearSearch(items, record{id: 1}, compareByID); got != 0 || !ok || err != nil {
		t.Fatalf("LinearSearch() = (%d, %t, %v), want (0, true, nil)", got, ok, err)
	}
}
