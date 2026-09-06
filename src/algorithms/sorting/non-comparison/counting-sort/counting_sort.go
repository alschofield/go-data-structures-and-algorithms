package counting_sort

import (
	"errors"
)

var ErrKeyOutOfRange error = errors.New("function needs a key that is within range.")

func CountingSort(items []uint32, keyLimit uint32) (bool, error) {
	if len(items) <= int(keyLimit) {
		return false, ErrKeyOutOfRange
	}

	var keys []int
	for i := 0; i < len(items); i++ {
		keys[items[i]]++
	}

	for i, j := 0, 0; i < len(keys); {
		if keys[i] != 0 {
			items[j] = uint32(i)
			j++
		} else {
			i++
		}
	}

	return true, nil
}
