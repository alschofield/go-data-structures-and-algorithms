//go:build contract

package counting_sort

import (
	"errors"
	"testing"
)

func TestCountingSort(t *testing.T) {
	for _, test := range []struct {
		items []uint32
		limit uint32
		valid bool
	}{{[]uint32{3, 1, 2, 1}, 4, true}, {[]uint32{3, 4}, 4, false}} {
		original := append([]uint32(nil), test.items...)
		err := CountingSort(test.items, test.limit)
		if test.valid && err != nil {
			t.Fatalf("valid input error = %v", err)
		}
		if !test.valid && !errors.Is(err, ErrKeyOutOfRange) {
			t.Fatalf("invalid key error = %v, want ErrKeyOutOfRange", err)
		}
		if !test.valid && (test.items[0] != original[0] || test.items[len(test.items)-1] != original[len(original)-1]) {
			t.Fatal("invalid input must not mutate")
		}
		for i := 1; test.valid && i < len(test.items); i++ {
			if test.items[i-1] > test.items[i] {
				t.Fatalf("not sorted: %v", test.items)
			}
		}
	}
}
