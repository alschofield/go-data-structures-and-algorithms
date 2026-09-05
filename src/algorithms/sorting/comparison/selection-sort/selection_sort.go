package selection_sort

import (
	"errors"
)

// ErrNilComparator reports that SelectionSort cannot order items without a comparator.
var ErrNilComparator = errors.New("function needs a compare function.")

// SelectionSort orders items in place according to compare.
func SelectionSort[T any](items []T, compare func(T, T) int) (bool, error) {
	// A missing comparator is invalid because values of an arbitrary type have no inherent order.
	if compare == nil {
		// Return before touching items so invalid input is always a no-op.
		return false, ErrNilComparator
	}

	// Slices with fewer than two items are already sorted.
	if len(items) == 0 || len(items) == 1 {
		return true, nil
	}

	// Grow the sorted prefix one position at a time.
	for i := 0; i < len(items); i++ {
		// Start by treating the first unsorted item as the smallest candidate.
		var candidate int = i
		// Find the smallest value in the remaining unsorted suffix.
		for n := i; n < len(items); n++ {
			// Replace the candidate whenever the current item sorts before it.
			if compare(items[candidate], items[n]) > 0 {
				candidate = n
			}
		}

		// Move the minimum into the prefix only when it is not already there.
		if candidate != i {
			items[i], items[candidate] = items[candidate], items[i]
		}
	}

	// Every position is now part of the sorted prefix.
	return true, nil
}
