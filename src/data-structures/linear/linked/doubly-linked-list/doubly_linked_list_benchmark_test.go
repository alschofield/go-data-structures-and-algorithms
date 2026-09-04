//go:build contract

package doubly_linked_list

import (
	"strconv"
	"testing"
)

var doublyLinkedListBenchmarkBool bool
var doublyLinkedListBenchmarkValue int

func BenchmarkDoublyLinkedListEndRoundTrip(b *testing.B) {
	for _, size := range []int{1, 1_024, 65_536} {
		for _, test := range []struct {
			name string
			push func(*DoublyLinkedList[int], int) bool
			pop  func(*DoublyLinkedList[int]) (int, bool)
		}{
			{name: "front", push: (*DoublyLinkedList[int]).PushFront, pop: (*DoublyLinkedList[int]).PopFront},
			{name: "back", push: (*DoublyLinkedList[int]).PushBack, pop: (*DoublyLinkedList[int]).PopBack},
		} {
			b.Run(test.name+"/Size="+strconv.Itoa(size), func(b *testing.B) {
				list := benchmarkListOfSize(size)
				b.ReportAllocs()
				b.ResetTimer()
				for value := 0; value < b.N; value++ {
					doublyLinkedListBenchmarkBool = test.push(list, size+value)
					popped, ok := test.pop(list)
					if !ok {
						b.Fatal("pop = false, want true")
					}
					doublyLinkedListBenchmarkValue = popped
				}
				b.StopTimer()
				if !doublyLinkedListBenchmarkBool || list.Len() != size || doublyLinkedListBenchmarkValue != size+b.N-1 {
					b.Fatalf("round trip left Len() = %d and returned %d last", list.Len(), doublyLinkedListBenchmarkValue)
				}
			})
		}
	}
}

func BenchmarkDoublyLinkedListGet(b *testing.B) {
	for _, size := range []int{1_024, 65_536} {
		list := benchmarkListOfSize(size)
		for _, test := range []struct {
			name  string
			index int
		}{
			{name: "near-head", index: 1},
			{name: "middle", index: size / 2},
			{name: "near-tail", index: size - 2},
		} {
			b.Run(test.name+"/Size="+strconv.Itoa(size), func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for range b.N {
					value, ok, err := list.Get(test.index)
					if err != nil || !ok {
						b.Fatalf("Get(%d) = (%d, %t, %v), want a value", test.index, value, ok, err)
					}
					doublyLinkedListBenchmarkValue = value
				}
				b.StopTimer()
				if list.Len() != size || doublyLinkedListBenchmarkValue != test.index {
					b.Fatalf("Get(%d) left Len() = %d and returned %d last", test.index, list.Len(), doublyLinkedListBenchmarkValue)
				}
			})
		}
	}
}

func BenchmarkDoublyLinkedListInsertRemoveMiddle(b *testing.B) {
	for _, size := range []int{1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			list := benchmarkListOfSize(size)
			index := size / 2
			b.ReportAllocs()
			b.ResetTimer()
			for value := 0; value < b.N; value++ {
				doublyLinkedListBenchmarkBool, _ = list.Insert(index, value)
				removed, ok, err := list.Remove(index)
				if err != nil || !ok {
					b.Fatalf("Remove(%d) = (%d, %t, %v), want a value", index, removed, ok, err)
				}
				doublyLinkedListBenchmarkValue = removed
			}
			b.StopTimer()
			if !doublyLinkedListBenchmarkBool || list.Len() != size || doublyLinkedListBenchmarkValue != b.N-1 {
				b.Fatalf("middle round trip left Len() = %d and returned %d last", list.Len(), doublyLinkedListBenchmarkValue)
			}
		})
	}
}

func benchmarkListOfSize(size int) *DoublyLinkedList[int] {
	list := NewDoublyLinkedList[int]()
	for value := 0; value < size; value++ {
		list.PushBack(value)
	}
	return list
}
