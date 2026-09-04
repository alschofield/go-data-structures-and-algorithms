//go:build contract

package doubly_linked_list

import (
	"errors"
	"testing"
)

func TestDoublyLinkedListEndOperationsEmptyOneAndManyNodes(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		for _, operation := range []struct {
			name string
			call func() (int, bool)
		}{
			{name: "PopFront", call: list.PopFront},
			{name: "PopBack", call: list.PopBack},
		} {
			t.Run(operation.name, func(t *testing.T) {
				beforeHead, beforeTail, beforeSize := list.head, list.tail, list.size
				if got, ok := operation.call(); got != 0 || ok {
					t.Fatalf("%s() = (%d, %t), want (0, false)", operation.name, got, ok)
				}
				assertUnchanged(t, list, beforeHead, beforeTail, beforeSize, nil)
			})
		}
	})

	t.Run("one node", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		if !list.PushFront(7) {
			t.Fatal("PushFront() = false, want true")
		}
		assertListState(t, list, []int{7})
		if got, ok := list.PopBack(); got != 7 || !ok {
			t.Fatalf("PopBack() = (%d, %t), want (7, true)", got, ok)
		}
		assertListState(t, list, nil)

		if !list.PushBack(8) {
			t.Fatal("PushBack() = false, want true")
		}
		assertListState(t, list, []int{8})
		if got, ok := list.PopFront(); got != 8 || !ok {
			t.Fatalf("PopFront() = (%d, %t), want (8, true)", got, ok)
		}
		assertListState(t, list, nil)
	})

	t.Run("many nodes", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		if !list.PushFront(2) || !list.PushFront(1) || !list.PushBack(3) || !list.PushBack(4) {
			t.Fatal("end insertion = false, want true")
		}
		assertListState(t, list, []int{1, 2, 3, 4})
		for _, test := range []struct {
			name  string
			call  func() (int, bool)
			want  int
			after []int
		}{
			{name: "PopFront", call: list.PopFront, want: 1, after: []int{2, 3, 4}},
			{name: "PopBack", call: list.PopBack, want: 4, after: []int{2, 3}},
			{name: "PopBack", call: list.PopBack, want: 3, after: []int{2}},
			{name: "PopFront", call: list.PopFront, want: 2, after: nil},
		} {
			t.Run(test.name, func(t *testing.T) {
				if got, ok := test.call(); got != test.want || !ok {
					t.Fatalf("%s() = (%d, %t), want (%d, true)", test.name, got, ok, test.want)
				}
				assertListState(t, list, test.after)
			})
		}
	})
}

func TestDoublyLinkedListInsertAndRemoveFrontBackAndMiddle(t *testing.T) {
	list := listOf(2, 4)
	for _, test := range []struct {
		name  string
		index int
		value int
		want  []int
	}{
		{name: "front", index: 0, value: 1, want: []int{1, 2, 4}},
		{name: "middle", index: 2, value: 3, want: []int{1, 2, 3, 4}},
		{name: "back", index: 4, value: 5, want: []int{1, 2, 3, 4, 5}},
	} {
		t.Run("insert "+test.name, func(t *testing.T) {
			if ok, err := list.Insert(test.index, test.value); !ok || err != nil {
				t.Fatalf("Insert(%d, %d) = (%t, %v), want (true, nil)", test.index, test.value, ok, err)
			}
			assertListState(t, list, test.want)
		})
	}

	for _, test := range []struct {
		name  string
		index int
		want  int
		after []int
	}{
		{name: "front", index: 0, want: 1, after: []int{2, 3, 4, 5}},
		{name: "middle", index: 1, want: 3, after: []int{2, 4, 5}},
		{name: "back", index: 2, want: 5, after: []int{2, 4}},
	} {
		t.Run("remove "+test.name, func(t *testing.T) {
			if got, ok, err := list.Remove(test.index); got != test.want || !ok || err != nil {
				t.Fatalf("Remove(%d) = (%d, %t, %v), want (%d, true, nil)", test.index, got, ok, err, test.want)
			}
			assertListState(t, list, test.after)
		})
	}
}

func TestDoublyLinkedListIndexedOperationsUseBothSides(t *testing.T) {
	list := listOf(0, 1, 2, 3, 4, 5, 6)
	for _, test := range []struct {
		name  string
		index int
		want  int
	}{
		{name: "near head", index: 1, want: 1},
		{name: "middle", index: 3, want: 3},
		{name: "near tail", index: 5, want: 5},
	} {
		t.Run("Get "+test.name, func(t *testing.T) {
			if got, ok, err := list.Get(test.index); got != test.want || !ok || err != nil {
				t.Fatalf("Get(%d) = (%d, %t, %v), want (%d, true, nil)", test.index, got, ok, err, test.want)
			}
			assertListState(t, list, []int{0, 1, 2, 3, 4, 5, 6})
		})
	}

	if ok, err := list.Insert(1, 10); !ok || err != nil {
		t.Fatalf("Insert(1, 10) = (%t, %v), want (true, nil)", ok, err)
	}
	assertListState(t, list, []int{0, 10, 1, 2, 3, 4, 5, 6})
	if ok, err := list.Insert(7, 20); !ok || err != nil {
		t.Fatalf("Insert(7, 20) = (%t, %v), want (true, nil)", ok, err)
	}
	assertListState(t, list, []int{0, 10, 1, 2, 3, 4, 5, 20, 6})
	if got, ok, err := list.Remove(1); got != 10 || !ok || err != nil {
		t.Fatalf("Remove(1) = (%d, %t, %v), want (10, true, nil)", got, ok, err)
	}
	assertListState(t, list, []int{0, 1, 2, 3, 4, 5, 20, 6})
	if got, ok, err := list.Remove(6); got != 20 || !ok || err != nil {
		t.Fatalf("Remove(6) = (%d, %t, %v), want (20, true, nil)", got, ok, err)
	}
	assertListState(t, list, []int{0, 1, 2, 3, 4, 5, 6})
}

func TestDoublyLinkedListInvalidIndexesDoNotMutate(t *testing.T) {
	list := listOf(10, 20, 30)
	for _, test := range []struct {
		name string
		call func() (int, bool, error)
	}{
		{name: "Get negative", call: func() (int, bool, error) { return list.Get(-1) }},
		{name: "Get out of range", call: func() (int, bool, error) { return list.Get(list.Len()) }},
		{name: "Remove negative", call: func() (int, bool, error) { return list.Remove(-1) }},
		{name: "Remove out of range", call: func() (int, bool, error) { return list.Remove(list.Len()) }},
	} {
		t.Run(test.name, func(t *testing.T) {
			beforeHead, beforeTail, beforeSize := list.head, list.tail, list.size
			if got, ok, err := test.call(); got != 0 || ok || !errors.Is(err, ErrInvalidIndex) {
				t.Fatalf("result = (%d, %t, %v), want (0, false, ErrInvalidIndex)", got, ok, err)
			}
			assertUnchanged(t, list, beforeHead, beforeTail, beforeSize, []int{10, 20, 30})
		})
	}

	for _, test := range []struct {
		name string
		call func() (bool, error)
	}{
		{name: "Insert negative", call: func() (bool, error) { return list.Insert(-1, 0) }},
		{name: "Insert out of range", call: func() (bool, error) { return list.Insert(list.Len()+1, 0) }},
	} {
		t.Run(test.name, func(t *testing.T) {
			beforeHead, beforeTail, beforeSize := list.head, list.tail, list.size
			if ok, err := test.call(); ok || !errors.Is(err, ErrInvalidIndex) {
				t.Fatalf("result = (%t, %v), want (false, ErrInvalidIndex)", ok, err)
			}
			assertUnchanged(t, list, beforeHead, beforeTail, beforeSize, []int{10, 20, 30})
		})
	}
}

func TestDoublyLinkedListLenAndIsEmpty(t *testing.T) {
	var zeroValue DoublyLinkedList[int]
	assertListState(t, &zeroValue, nil)

	list := NewDoublyLinkedList[int]()
	if !list.PushBack(1) {
		t.Fatal("PushBack() = false, want true")
	}
	assertListState(t, list, []int{1})
}

func listOf(values ...int) *DoublyLinkedList[int] {
	list := NewDoublyLinkedList[int]()
	for _, value := range values {
		list.PushBack(value)
	}
	return list
}

func assertUnchanged(t *testing.T, list *DoublyLinkedList[int], head, tail *Node[int], size int, want []int) {
	t.Helper()
	if list.head != head || list.tail != tail || list.size != size {
		t.Fatal("operation mutated list metadata on failure")
	}
	assertListState(t, list, want)
}

func assertListState(t *testing.T, list *DoublyLinkedList[int], want []int) {
	t.Helper()
	if got := list.Len(); got != len(want) {
		t.Fatalf("Len() = %d, want %d", got, len(want))
	}
	if got := list.IsEmpty(); got != (len(want) == 0) {
		t.Fatalf("IsEmpty() = %t, want %t", got, len(want) == 0)
	}
	if len(want) == 0 {
		if list.head != nil || list.tail != nil {
			t.Fatal("empty list must have nil head and tail")
		}
		return
	}
	if list.head == nil || list.tail == nil {
		t.Fatal("non-empty list must have head and tail")
	}
	if list.head.prev != nil || list.tail.next != nil {
		t.Fatal("head.prev and tail.next must be nil")
	}

	forward := list.head
	for index, value := range want {
		if forward == nil {
			t.Fatalf("forward traversal ended at index %d", index)
		}
		if forward.value != value {
			t.Fatalf("forward value at index %d = %d, want %d", index, forward.value, value)
		}
		if forward.next != nil && forward.next.prev != forward {
			t.Fatalf("forward link at index %d is not reciprocal", index)
		}
		forward = forward.next
	}
	if forward != nil {
		t.Fatal("forward traversal contains more nodes than Len()")
	}

	backward := list.tail
	for index := len(want) - 1; index >= 0; index-- {
		if backward == nil {
			t.Fatalf("backward traversal ended at index %d", index)
		}
		if backward.value != want[index] {
			t.Fatalf("backward value at index %d = %d, want %d", index, backward.value, want[index])
		}
		if backward.prev != nil && backward.prev.next != backward {
			t.Fatalf("backward link at index %d is not reciprocal", index)
		}
		backward = backward.prev
	}
	if backward != nil {
		t.Fatal("backward traversal contains more nodes than Len()")
	}
}
