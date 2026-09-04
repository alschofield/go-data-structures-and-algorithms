//go:build contract

package singly_linked_list

import (
	"errors"
	"testing"
)

func TestSinglyLinkedListPushAndPopOrder(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	for _, value := range []int{2, 3} {
		if !list.PushBack(value) {
			t.Fatal("PushBack() = false, want true")
		}
	}
	if !list.PushFront(1) {
		t.Fatal("PushFront() = false, want true")
	}
	assertListContents(t, list, []int{1, 2, 3})

	for _, want := range []int{1, 2, 3} {
		if got, ok := list.PopFront(); !ok || got != want {
			t.Fatalf("PopFront() = (%d, %t), want (%d, true)", got, ok, want)
		}
	}
	assertListContents(t, list, nil)
}

func TestSinglyLinkedListEmptyPopsDoNotMutate(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	for _, operation := range []struct {
		name string
		call func() (int, bool)
	}{
		{name: "PopFront", call: list.PopFront},
		{name: "PopBack", call: list.PopBack},
	} {
		t.Run(operation.name, func(t *testing.T) {
			if got, ok := operation.call(); got != 0 || ok {
				t.Fatalf("%s() = (%d, %t), want (0, false)", operation.name, got, ok)
			}
			assertListContents(t, list, nil)
		})
	}
}

func TestSinglyLinkedListPopBackOneAndManyNodes(t *testing.T) {
	t.Run("one node", func(t *testing.T) {
		list := listOf(7)
		if got, ok := list.PopBack(); !ok || got != 7 {
			t.Fatalf("PopBack() = (%d, %t), want (7, true)", got, ok)
		}
		assertListContents(t, list, nil)
	})

	t.Run("many nodes", func(t *testing.T) {
		list := listOf(1, 2, 3)
		for _, want := range []int{3, 2, 1} {
			if got, ok := list.PopBack(); !ok || got != want {
				t.Fatalf("PopBack() = (%d, %t), want (%d, true)", got, ok, want)
			}
		}
		assertListContents(t, list, nil)
	})
}

func TestSinglyLinkedListInsertHeadMiddleAndEnd(t *testing.T) {
	list := listOf(2, 4)
	for _, test := range []struct {
		name  string
		index int
		value int
		want  []int
	}{
		{name: "head", index: 0, value: 1, want: []int{1, 2, 4}},
		{name: "middle", index: 2, value: 3, want: []int{1, 2, 3, 4}},
		{name: "end", index: 4, value: 5, want: []int{1, 2, 3, 4, 5}},
	} {
		t.Run(test.name, func(t *testing.T) {
			if ok, err := list.Insert(test.index, test.value); !ok || err != nil {
				t.Fatalf("Insert(%d, %d) = (%t, %v), want (true, nil)", test.index, test.value, ok, err)
			}
			assertListContents(t, list, test.want)
		})
	}
}

func TestSinglyLinkedListRemoveHeadMiddleAndTail(t *testing.T) {
	list := listOf(1, 2, 3, 4, 5)
	for _, test := range []struct {
		name  string
		index int
		want  int
		after []int
	}{
		{name: "head", index: 0, want: 1, after: []int{2, 3, 4, 5}},
		{name: "middle", index: 1, want: 3, after: []int{2, 4, 5}},
		{name: "tail", index: 2, want: 5, after: []int{2, 4}},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got, ok, err := list.Remove(test.index); !ok || err != nil || got != test.want {
				t.Fatalf("Remove(%d) = (%d, %t, %v), want (%d, true, nil)", test.index, got, ok, err, test.want)
			}
			assertListContents(t, list, test.after)
		})
	}
}

func TestSinglyLinkedListInvalidIndexesDoNotMutate(t *testing.T) {
	list := listOf(10, 20, 30)
	for _, test := range []struct {
		name string
		call func() error
	}{
		{name: "Get negative", call: func() error { _, _, err := list.Get(-1); return err }},
		{name: "Get out of range", call: func() error { _, _, err := list.Get(list.Len()); return err }},
		{name: "Insert negative", call: func() error { _, err := list.Insert(-1, 0); return err }},
		{name: "Insert out of range", call: func() error { _, err := list.Insert(list.Len()+1, 0); return err }},
		{name: "Remove negative", call: func() error { _, _, err := list.Remove(-1); return err }},
		{name: "Remove out of range", call: func() error { _, _, err := list.Remove(list.Len()); return err }},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := test.call(); !errors.Is(err, ErrInvalidIndex) {
				t.Fatalf("error = %v, want ErrInvalidIndex", err)
			}
			assertListContents(t, list, []int{10, 20, 30})
		})
	}
}

func TestSinglyLinkedListLenAndIsEmpty(t *testing.T) {
	var zeroValue SinglyLinkedList[int]
	assertListContents(t, &zeroValue, nil)

	list := NewSinglyLinkedList[int]()
	if !list.PushFront(1) {
		t.Fatal("PushFront() = false, want true")
	}
	assertListContents(t, list, []int{1})
}

func listOf(values ...int) *SinglyLinkedList[int] {
	list := NewSinglyLinkedList[int]()
	for _, value := range values {
		list.PushBack(value)
	}
	return list
}

func assertListContents(t *testing.T, list *SinglyLinkedList[int], want []int) {
	t.Helper()
	if got := list.Len(); got != len(want) {
		t.Fatalf("Len() = %d, want %d", got, len(want))
	}
	if got := list.IsEmpty(); got != (len(want) == 0) {
		t.Fatalf("IsEmpty() = %t, want %t", got, len(want) == 0)
	}
	for index, expected := range want {
		if got, ok, err := list.Get(index); err != nil || !ok || got != expected {
			t.Fatalf("Get(%d) = (%d, %t, %v), want (%d, true, nil)", index, got, ok, err, expected)
		}
	}
}
