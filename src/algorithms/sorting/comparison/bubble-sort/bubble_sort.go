package bubble_sort

import (
	"errors"
)

var ErrNilComparator = errors.New("function requires a non nil compare function.")

func BubbleSort[T any](items []T, compare func(T, T) int) (bool, error) {
	if compare == nil {
		return false, ErrNilComparator
	}

	// Each pass moves the largest remaining value into the unsorted tail.
	for end := len(items) - 1; end > 0; end-- {
		// A complete pass with no swaps proves the remaining prefix is sorted.
		swapped := false
		for i := 0; i < end; i++ {
			// Strict comparison preserves the original order of equal values.
			if compare(items[i], items[i+1]) > 0 {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return true, nil
}
