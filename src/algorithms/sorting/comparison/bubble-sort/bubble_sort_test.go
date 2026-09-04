//go:build contract

package bubble_sort

import "testing"

func TestBubbleSort(t *testing.T) { testComparisonSort(t, BubbleSort[int]) }

func testComparisonSort(t *testing.T, sort func([]int, func(int, int) int) bool) {
	for _, test := range [][]int{{}, {1}, {3, 1, 2}, {3, 2, 1}, {2, 2, 1}} {
		if !sort(test, func(a, b int) int { return a - b }) {
			t.Fatal("sort reported failure")
		}
		for i := 1; i < len(test); i++ {
			if test[i-1] > test[i] {
				t.Fatalf("not sorted: %v", test)
			}
		}
	}
}
