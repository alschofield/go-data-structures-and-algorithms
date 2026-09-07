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
		key  int
		ok   bool
		name string
	}{
		{name: "first", key: 1, ok: true},
		{name: "duplicate", key: 3, ok: true},
		{name: "last", key: 8, ok: true},
		{name: "middle miss", key: 4, ok: false},
		{name: "below range", key: 0, ok: false},
		{name: "above range", key: 9, ok: false},
	} {
		t.Run(test.name, func(t *testing.T) {
			index, ok, err := BinarySearch(items, test.key, compare)
			if err != nil || ok != test.ok || (ok && items[index] != test.key) {
				t.Fatalf("key %d: got (%d, %t, %v)", test.key, index, ok, err)
			}
		})
	}
	if _, ok, err := BinarySearch([]int(nil), 1, compare); err != nil || ok {
		t.Fatal("empty input must be absent")
	}
	if _, _, err := BinarySearch(items, 1, nil); !errors.Is(err, ErrNilComparator) {
		t.Fatalf("nil comparator error = %v, want ErrNilComparator", err)
	}
}

func TestBinarySearchGenericComparatorDoesNotMutate(t *testing.T) {
	type record struct {
		name string
		age  int
	}
	items := []record{{"Grace", 28}, {"Ada", 36}, {"Linus", 54}}
	compareByAge := func(left, right record) int { return left.age - right.age }

	index, ok, err := BinarySearch(items, record{age: 36}, compareByAge)
	if err != nil || !ok || index != 1 {
		t.Fatalf("BinarySearch() = (%d, %t, %v), want (1, true, nil)", index, ok, err)
	}
	if items[0].name != "Grace" || items[1].name != "Ada" || items[2].name != "Linus" {
		t.Fatal("BinarySearch() mutated input")
	}
}
