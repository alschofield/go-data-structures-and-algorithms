package heap_sort

import "errors"

// ErrNilComparator reports that HeapSort cannot order values without a comparator.
var ErrNilComparator = errors.New("function needs a valid compare function.")

// HeapSort sorts items ascending in place with an implicit max heap.
func HeapSort[T any](items []T, compare func(T, T) int) (bool, error) {
	if compare == nil {
		return false, ErrNilComparator
	}

	// Heapify each parent from the last parent back to the root.
	for i := len(items)/2 - 1; i >= 0; i-- {
		for n := i; n < len(items)-1; {
			left_index := n*2 + 1
			right_index := n*2 + 2
			var child_index int

			if left_index >= len(items) {
				break
			} else if right_index >= len(items) {
				child_index = left_index
			} else if compare(items[left_index], items[right_index]) > 0 {
				child_index = left_index
			} else {
				child_index = right_index
			}

			if compare(items[n], items[child_index]) < 0 {
				items[child_index], items[n] = items[n], items[child_index]
				n = child_index
			} else {
				break
			}
		}
	}

	// Move the maximum to the tail, then restore the smaller heap prefix.
	for i := len(items) - 1; i >= 0; i-- {
		items[0], items[i] = items[i], items[0]

		// Sift the new root down without touching the sorted suffix.
		for n := 0; n < i; {
			left_index := n*2 + 1
			right_index := n*2 + 2
			var child_index int

			if left_index >= i {
				break
			} else if right_index >= i {
				child_index = left_index
			} else if compare(items[left_index], items[right_index]) > 0 {
				child_index = left_index
			} else {
				child_index = right_index
			}

			if compare(items[n], items[child_index]) < 0 {
				items[child_index], items[n] = items[n], items[child_index]
				n = child_index
			} else {
				break
			}
		}
	}

	return true, nil
}
