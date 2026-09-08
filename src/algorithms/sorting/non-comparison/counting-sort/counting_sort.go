package counting_sort

import (
	"errors"
)

var ErrKeyOutOfRange error = errors.New("function needs a key that is within range.")

func CountingSort(items []uint32, keyLimit uint32) (bool, error) {
	// Each valid key gets one counter; keyLimit bounds both memory and values.
	var keys []int = make([]int, keyLimit)
	for i := 0; i < len(items); i++ {
		// Validate before rewriting so invalid input remains unchanged.
		if items[i] >= keyLimit {
			return false, ErrKeyOutOfRange
		}

		// Count how many times this value appears.
		keys[items[i]]++
	}

	// Rewrite the input in key order, consuming every recorded occurrence.
	for i, j := 0, 0; i < len(keys); {
		if keys[i] != 0 {
			items[j] = uint32(i)
			keys[i]--
			j++
		} else {
			// This key has been fully written; advance to the next value.
			i++
		}
	}

	return true, nil
}
