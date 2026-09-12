//go:build contract

package binary_heap

import (
	"fmt"
	"testing"
)

func BenchmarkBinaryHeap(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		for _, input := range []struct {
			name   string
			values []int
		}{
			{"push-random", heapRandomValues(size)},
			{"push-ascending", heapAscendingValues(size)},
		} {
			b.Run(fmt.Sprintf("%s/Size=%d", input.name, size), func(b *testing.B) {
				b.ReportAllocs()
				for iteration := 0; iteration < b.N; iteration++ {
					heap := heapFromValues(b, nil)
					for _, value := range input.values {
						if !heap.Push(value) {
							b.Fatal("Push() failed")
						}
					}
				}
			})
		}

		b.Run(fmt.Sprintf("pop-all/Size=%d", size), func(b *testing.B) {
			values := heapRandomValues(size)
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				heap := heapFromValues(b, values)
				for count := 0; count < size; count++ {
					if _, ok := heap.Pop(); !ok {
						b.Fatal("Pop() failed before the heap was empty")
					}
				}
			}
		})

		b.Run(fmt.Sprintf("mixed-push-pop/Size=%d", size), func(b *testing.B) {
			values := heapRandomValues(size)
			heap := heapFromValues(b, values)
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				if !heap.Push(values[iteration%size]) {
					b.Fatal("Push() failed")
				}
				if _, ok := heap.Pop(); !ok {
					b.Fatal("Pop() failed")
				}
			}
		})
	}
}

func heapFromValues(b *testing.B, values []int) *BinaryHeap[int] {
	heap, err := NewBinaryHeap(compareHeapInts)
	if err != nil {
		b.Fatal(err)
	}
	for _, value := range values {
		if !heap.Push(value) {
			b.Fatal("heap setup Push() failed")
		}
	}
	return heap
}

func compareHeapInts(left, right int) int { return left - right }

func heapAscendingValues(size int) []int {
	values := make([]int, size)
	for index := range values {
		values[index] = index
	}
	return values
}

func heapRandomValues(size int) []int {
	values := make([]int, size)
	for index := range values {
		values[index] = (index*7919 + 104729) % size
	}
	return values
}
