// Package linear_search contains a generic sequential search implementation.
package linear_search

// errors constructs the stable sentinel returned for an invalid comparator.
import (
	"errors"
)

// ErrNilComparator reports that LinearSearch cannot compare the supplied values.
var ErrNilComparator = errors.New("function needs a compare function.")

// LinearSearch returns the first index whose item compares equal to key.
func LinearSearch[T any](items []T, key T, compare func(T, T) int) (int, bool, error) {
	// A comparison function is required because generic values have no built-in order.
	if compare == nil {
		return 0, false, ErrNilComparator
	}

	// Visit items in input order so the first matching duplicate is returned.
	for i, item := range items {
		// A zero comparison denotes equality under the caller's comparison rule.
		if compare(key, item) == 0 {
			// The index and true distinguish a match at index zero from a miss.
			return i, true, nil
		}
	}

	// A completed scan means no item matched; absence is an expected, non-error result.
	return 0, false, nil
}
