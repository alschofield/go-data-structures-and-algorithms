package bubble_sort

import (
	"errors"
)

var ErrNilComparator = errors.New("function requires a non nil compare function.")

func BubbleSort[T any](items []T, compare func(T, T) int) (bool, error) {
	if compare == nil {
		return false, ErrNilComparator
	}

	for i := 0; i < len(items); i++ {
		for j := i; j < len(items); j++ {
			if compare(items[j], items[j+1]) < 0 {
				items[j+1], items[j] = items[j], items[j+1]
			}
		}
	}

	return true, nil
}
