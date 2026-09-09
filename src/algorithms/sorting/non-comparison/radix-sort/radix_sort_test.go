//go:build contract

package radix_sort

import "testing"

func TestRadixSort(t *testing.T) {
	for _, items := range [][]uint32{nil, {}, {1}, {329, 457, 657, 839, 436, 720, 355}, {2, 2, 1}, {0xffff_ffff, 0, 0x0100_0000, 0x0000_00ff}} {
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

func TestRadixCountPreservesEqualByteOrder(t *testing.T) {
	items := []uint32{0x0201, 0x0101, 0x0302}
	sorted, ok := RadixCount(items, 0)
	if !ok || sorted[0] != 0x0201 || sorted[1] != 0x0101 || sorted[2] != 0x0302 {
		t.Fatalf("RadixCount() = %v, want stable low-byte order", sorted)
	}
}
