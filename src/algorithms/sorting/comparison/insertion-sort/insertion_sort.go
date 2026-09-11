package insertion_sort

import "errors"

// ErrNilComparator reports that InsertionSort cannot order values without a comparator.
var ErrNilComparator error = errors.New("function needs a non nil compare function.")

// InsertionSort sorts items ascending in place by growing a stable sorted prefix.
func InsertionSort[T any](items []T, compare func(T, T) int) (bool, error) {
	if compare == nil {
		return false, ErrNilComparator
	}

	// Insert each later item into the sorted prefix before it.
	for i := 1; i < len(items); i++ {
		for j := i; j > 0 && compare(items[j-1], items[j]) > 0; j-- {
			// Move a larger prefix item right, leaving equals in their original order.
			items[j-1], items[j] = items[j], items[j-1]
		}
	}

	return true, nil
}
