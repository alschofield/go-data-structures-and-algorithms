package merge_sort

import (
	"errors"
)

var ErrNilComparator = errors.New("function needs a non nil compare function")

func recursion[T any](items []T, compare func(T, T) int) bool {
	if len(items) <= 1 {
		return true
	}

	var middle_index = len(items) / 2

	if !recursion(items[:middle_index], compare) {
		return false
	}

	if !recursion(items[middle_index:], compare) {
		return false
	}

	// combine the two sorted halves
	// might need to change this to use left and right indexs
	// or recursion needs to return a slice

	return true
}

func MergeSort[T any](items []T, compare func(T, T) int) (bool, error) {
	if compare == nil {
		return false, ErrNilComparator
	}

	return recursion(items, compare), nil
}
