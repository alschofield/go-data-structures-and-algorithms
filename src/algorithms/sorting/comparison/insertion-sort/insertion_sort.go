package insertion_sort

import (
	"errors"
)

var ErrNilComparator error = errors.New("function needs a non nil compare function.")

func InsertionSort[T any](items []T, compare func(T, T) int) (bool, error) {
	if compare == nil {
		return false, ErrNilComparator
	}

	for i := 0; i < len(items); i++ {
		candidate_index := i
		for j := i + 1; j < len(items); j++ {
			if compare(items[candidate_index], items[j]) < 0 {
				candidate_index = j
			}
		}

		items[candidate_index], items[i] = items[i], items[candidate_index]
	}

	return true, nil
}
