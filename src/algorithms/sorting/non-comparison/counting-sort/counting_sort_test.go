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
	}{{[]uint32{}, 0, true}, {[]uint32{3, 1, 2, 1}, 4, true}, {[]uint32{3, 3, 3}, 4, true}, {[]uint32{0, 3, 1}, 4, true}, {[]uint32{3, 4}, 4, false}, {[]uint32{0}, 0, false}} {
		original := append([]uint32(nil), test.items...)
		ok, err := CountingSort(test.items, test.limit)
		if test.valid && (!ok || err != nil) {
			t.Fatalf("CountingSort() = (%t, %v), want (true, nil)", ok, err)
		}
		if !test.valid && (ok || !errors.Is(err, ErrKeyOutOfRange)) {
			t.Fatalf("CountingSort() = (%t, %v), want (false, ErrKeyOutOfRange)", ok, err)
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
