// Package binary_search contains a generic iterative binary search implementation.
package binary_search

// errors constructs the stable sentinel returned for an invalid comparator.
import (
	"errors"
)

var (
	// ErrNilComparator reports that BinarySearch cannot compare the supplied values.
	ErrNilComparator = errors.New("function needs a compare function.")
)

// BinarySearch returns an index whose item compares equal to key in ascending input.
func BinarySearch[T any](items []T, key T, compare func(T, T) int) (int, bool, error) {
	// A comparison function is required because generic values have no built-in order.
	if compare == nil {
		return 0, false, ErrNilComparator
	}

	// These inclusive bounds initially cover every candidate index.
	left, right := 0, len(items)-1

	// Continue while the remaining candidate range is non-empty.
	for left <= right {
		// This midpoint form avoids overflowing when left and right are large.
		middle := left + (right-left)/2
		// The comparison determines which half can still contain the key.
		comparison := compare(key, items[middle])

		// Equality identifies a valid matching index; duplicates need not return first.
		if comparison == 0 {
			return middle, true, nil
		} else if comparison < 0 {
			// The key is smaller, so discard the midpoint and upper half.
			right = middle - 1
		} else if comparison > 0 {
			// The key is larger, so discard the midpoint and lower half.
			left = middle + 1
		}
	}

	// An exhausted candidate range means no item matched; absence is not an error.
	return 0, false, nil
}
