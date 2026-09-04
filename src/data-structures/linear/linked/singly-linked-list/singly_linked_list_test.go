//go:build contract

package singly_linked_list

import "testing"

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.PushFront(2)
	list.PushFront(1)
	list.PushBack(3)
	if !list.Insert(1, 9) {
		t.Fatal("valid insert failed")
	}
	for index, want := range []int{1, 9, 2, 3} {
		if got, ok := list.Get(index); !ok || got != want {
			t.Fatalf("Get(%d) = (%d, %t), want (%d, true)", index, got, ok, want)
		}
	}
	if got, ok := list.PopBack(); !ok || got != 3 {
		t.Fatal("PopBack must return tail")
	}
	if list.Insert(-1, 0) {
		t.Fatal("invalid insert must fail")
	}
}
