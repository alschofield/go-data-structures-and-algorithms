//go:build contract

package radix_sort

import "testing"

func TestRadixSort(t *testing.T) {
	for _, items := range [][]uint32{{}, {1}, {329, 457, 657, 839, 436, 720, 355}, {2, 2, 1}} {
		if !RadixSort(items) {
			t.Fatal("sort reported failure")
		}
		for i := 1; i < len(items); i++ {
			if items[i-1] > items[i] {
				t.Fatalf("not sorted: %v", items)
			}
		}
	}
}
