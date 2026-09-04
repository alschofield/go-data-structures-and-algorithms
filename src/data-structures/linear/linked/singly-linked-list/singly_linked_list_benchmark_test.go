//go:build contract

package singly_linked_list

import (
	"strconv"
	"testing"
)

var singlyLinkedListBenchmarkBool bool
var singlyLinkedListBenchmarkValue int

func BenchmarkSinglyLinkedListPushFront(b *testing.B) {
	for _, size := range []int{0, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			list := benchmarkListOfSize(size)
			b.ReportAllocs()
			b.ResetTimer()
			for value := 0; value < b.N; value++ {
				singlyLinkedListBenchmarkBool = list.PushFront(value)
			}
			b.StopTimer()
			if !singlyLinkedListBenchmarkBool || list.Len() != size+b.N {
				b.Fatalf("PushFront() produced Len() = %d, want %d", list.Len(), size+b.N)
			}
		})
	}
}

func BenchmarkSinglyLinkedListPopFront(b *testing.B) {
	for _, size := range []int{0, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			list := benchmarkListOfSize(size + b.N)
			b.ReportAllocs()
			b.ResetTimer()
			for range b.N {
				value, ok := list.PopFront()
				if !ok {
					b.Fatal("PopFront() = false, want true")
				}
				singlyLinkedListBenchmarkValue = value
			}
			b.StopTimer()
			if list.Len() != size || singlyLinkedListBenchmarkValue != b.N-1 {
				b.Fatalf("PopFront() left Len() = %d and returned %d last, want %d and %d", list.Len(), singlyLinkedListBenchmarkValue, size, b.N-1)
			}
		})
	}
}

func BenchmarkSinglyLinkedListBackRoundTrip(b *testing.B) {
	for _, size := range []int{1, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			list := benchmarkListOfSize(size)
			b.ReportAllocs()
			b.ResetTimer()
			for value := 0; value < b.N; value++ {
				singlyLinkedListBenchmarkBool = list.PushBack(value + size)
				value, ok := list.PopBack()
				if !ok {
					b.Fatal("PopBack() = false, want true")
				}
				singlyLinkedListBenchmarkValue = value
			}
			b.StopTimer()
			if !singlyLinkedListBenchmarkBool || list.Len() != size || singlyLinkedListBenchmarkValue != size+b.N-1 {
				b.Fatalf("back round trip left Len() = %d and returned %d last, want %d and %d", list.Len(), singlyLinkedListBenchmarkValue, size, size+b.N-1)
			}
		})
	}
}

func BenchmarkSinglyLinkedListGetMiddle(b *testing.B) {
	for _, size := range []int{1, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			list := benchmarkListOfSize(size)
			index := size / 2
			b.ReportAllocs()
			b.ResetTimer()
			for range b.N {
				value, ok, err := list.Get(index)
				if err != nil || !ok {
					b.Fatalf("Get(%d) = (%d, %t, %v), want a value", index, value, ok, err)
				}
				singlyLinkedListBenchmarkValue = value
			}
			b.StopTimer()
			if list.Len() != size || singlyLinkedListBenchmarkValue != index {
				b.Fatalf("Get(%d) left Len() = %d and returned %d last, want %d and %d", index, list.Len(), singlyLinkedListBenchmarkValue, size, index)
			}
		})
	}
}

func benchmarkListOfSize(size int) *SinglyLinkedList[int] {
	list := NewSinglyLinkedList[int]()
	// Build in linear time outside measurements; repeated PushBack setup is quadratic.
	for value := size - 1; value >= 0; value-- {
		list.PushFront(value)
	}
	return list
}
