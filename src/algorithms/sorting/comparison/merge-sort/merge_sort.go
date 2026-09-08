package merge_sort

import (
	"errors"
)

var ErrNilComparator = errors.New("function needs a non nil compare function")

func recursion[T any](items []T, compare func(T, T) int) ([]T, bool) {
	// A zero- or one-element slice is already sorted.
	if len(items) <= 1 {
		return items, true
	}

	// Divide into two smaller slices until the base case above is reached.
	var middle_index = len(items) / 2

	left_items, status := recursion(items[:middle_index], compare)

	if !status {
		return items, status
	}

	right_items, status := recursion(items[middle_index:], compare)

	if !status {
		return items, status
	}

	// Merge both sorted halves into a buffer sized for every input item.
	var new_array []T = make([]T, len(items))
	var left_index int = 0
	var right_index int = 0
	for i := 0; i < len(items); i++ {
		var candidate_index int
		var candidate_items []T
		if left_index >= len(left_items) {
			// The left half is exhausted, so only right values remain.
			candidate_index = right_index
			candidate_items = right_items
			right_index++
		} else if right_index >= len(right_items) {
			// The right half is exhausted, so only left values remain.
			candidate_index = left_index
			candidate_items = left_items
			left_index++
		} else {
			var comparison int = compare(left_items[left_index], right_items[right_index])
			// On ties take left first, preserving equal items' original order.
			if comparison <= 0 {
				candidate_index = left_index
				candidate_items = left_items
				left_index++
			} else if comparison > 0 {
				candidate_index = right_index
				candidate_items = right_items
				right_index++
			}
		}

		new_array[i] = candidate_items[candidate_index]
	}

	return new_array, true
}

func MergeSort[T any](items []T, compare func(T, T) int) (bool, error) {
	if compare == nil {
		return false, ErrNilComparator
	}

	// Recursive merges allocate buffers; copy the final buffer into the caller's slice.
	sorted_items, status := recursion(items, compare)

	copy(items, sorted_items)

	return status, nil
}
