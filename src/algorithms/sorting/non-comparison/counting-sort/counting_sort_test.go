//go:build contract

package counting_sort

import "testing"

func TestCountingSort(t *testing.T) {
	for _, test := range []struct {
		items []uint32
		limit uint32
		ok    bool
	}{{[]uint32{3, 1, 2, 1}, 4, true}, {[]uint32{3, 4}, 4, false}} {
		original := append([]uint32(nil), test.items...)
		got := CountingSort(test.items, test.limit)
		if got != test.ok {
			t.Fatalf("got %t, want %t", got, test.ok)
		}
		if !got && (test.items[0] != original[0] || test.items[len(test.items)-1] != original[len(original)-1]) {
			t.Fatal("invalid input must not mutate")
		}
		for i := 1; got && i < len(test.items); i++ {
			if test.items[i-1] > test.items[i] {
				t.Fatalf("not sorted: %v", test.items)
			}
		}
	}
}
