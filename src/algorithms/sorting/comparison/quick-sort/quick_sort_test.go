//go:build contract

package quick_sort

import (
	"errors"
	"reflect"
	"testing"
)

func TestQuickSortIntegers(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  []int
	}{
		{name: "empty", items: []int{}, want: []int{}},
		{name: "singleton", items: []int{42}, want: []int{42}},
		{name: "sorted", items: []int{-3, -1, 0, 2, 7}, want: []int{-3, -1, 0, 2, 7}},
		{name: "reverse", items: []int{7, 2, 0, -1, -3}, want: []int{-3, -1, 0, 2, 7}},
		{name: "random", items: []int{8, 1, 6, 3, 7, 2, 5, 4}, want: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{name: "duplicates", items: []int{4, 1, 4, 2, 4, 1, 3, 2}, want: []int{1, 1, 2, 2, 3, 4, 4, 4}},
		{name: "all equal", items: []int{5, 5, 5, 5, 5}, want: []int{5, 5, 5, 5, 5}},
		{name: "negative", items: []int{-4, -1, -9, -2, -9}, want: []int{-9, -9, -4, -2, -1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			items := append([]int{}, test.items...)
			ok, err := QuickSort(items, compareInts)
			if err != nil || !ok {
				t.Fatalf("QuickSort() = (%t, %v), want (true, nil)", ok, err)
			}
			if !reflect.DeepEqual(items, test.want) {
				t.Fatalf("QuickSort() items = %v, want %v", items, test.want)
			}
		})
	}
}

func TestQuickSortNilComparatorDoesNotMutate(t *testing.T) {
	items := []int{3, 1, 2}
	want := append([]int(nil), items...)

	ok, err := QuickSort(items, nil)
	if ok || !errors.Is(err, ErrNilComparator) {
		t.Fatalf("QuickSort() = (%t, %v), want (false, ErrNilComparator)", ok, err)
	}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("QuickSort() mutated items to %v, want %v", items, want)
	}
}

func TestQuickSortGenericComparator(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	items := []person{{"Ada", 36}, {"Grace", 28}, {"Linus", 54}, {"Ken", 28}}
	want := []person{{"Grace", 28}, {"Ken", 28}, {"Ada", 36}, {"Linus", 54}}
	compareByAge := func(left, right person) int {
		return left.age - right.age
	}

	ok, err := QuickSort(items, compareByAge)
	if err != nil || !ok {
		t.Fatalf("QuickSort() = (%t, %v), want (true, nil)", ok, err)
	}
	for index, person := range items {
		if person.age != want[index].age {
			t.Fatalf("QuickSort() items = %v, want ages %v", items, want)
		}
	}
}

func compareInts(left, right int) int {
	return left - right
}
