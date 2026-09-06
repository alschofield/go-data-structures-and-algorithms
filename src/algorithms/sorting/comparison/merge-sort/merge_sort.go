package merge_sort

import (
	"errors"
)

var ErrNilComparator = errors.New("function needs a non nil compare function")

func recursion[T any](items []T, compare func(T, T) int) ([]T, bool) {
	if len(items) <= 1 {
		return items, true
	}

	var middle_index = len(items) / 2

	left_items, status := recursion(items[:middle_index], compare)

	if !status {
		return items, status
	}

	right_items, status := recursion(items[middle_index:], compare)

	if !status {
		return items, status
	}

	var new_array []T
	var left_index int = 0
	var right_index int = 0
	for i := 0; i < len(items); i++ {
		var comparison int = compare(left_items[left_index], right_items[right_index])
		if comparison < 0 {
			new_array[i] = left_items[left_index]
			left_index++
		} else if comparison >= 0 {
			new_array[i] = right_items[right_index]
			right_index++
		}
	}

	return items, true
}

func MergeSort[T any](items []T, compare func(T, T) int) (bool, error) {
	if compare == nil {
		return false, ErrNilComparator
	}

	// i think using the new slices MIGHT cause excess memory consumption during runtime
	// 		but slices should be pointer to value so i may be wrong here as long as we pass and return slices
	sorted_items, status := recursion(items, compare)

	items = sorted_items

	return status, nil
}
