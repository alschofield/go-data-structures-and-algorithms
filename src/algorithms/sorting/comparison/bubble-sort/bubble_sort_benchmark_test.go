//go:build contract

package bubble_sort

import (
	"fmt"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		for _, input := range []struct {
			name  string
			items []int
		}{
			{"sorted", bubbleSorted(size)},
			{"reverse", bubbleReverse(size)},
		} {
			b.Run(fmt.Sprintf("%s/Size=%d", input.name, size), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					items := append([]int(nil), input.items...)
					if ok, err := BubbleSort(items, func(a, b int) int { return a - b }); !ok || err != nil {
						b.Fatalf("BubbleSort() = (%t, %v), want (true, nil)", ok, err)
					}
				}
			})
		}
	}
}

func bubbleSorted(size int) []int {
	items := make([]int, size)
	for i := range items {
		items[i] = i
	}
	return items
}

func bubbleReverse(size int) []int {
	items := bubbleSorted(size)
	for i := range items {
		items[i] = size - 1 - i
	}
	return items
}
