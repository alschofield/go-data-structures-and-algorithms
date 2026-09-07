package bubble_sort

import (
	"errors"
)

var ErrNilComparator = errors.New("function requires a non nil compare function.")

func BubbleSort[T any](items []T, compare func(T, T) int) (bool, error) {
	if compare == nil {
		return false, ErrNilComparator
	}

	for end := len(items) - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if compare(items[i], items[i+1]) > 0 {
				items[i+1], items[i] = items[i], items[i+1]
			}
		}
	}

	return true, nil
}
