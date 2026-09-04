//go:build contract

package singly_linked_list

import (
	"errors"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.PushFront(2)
	list.PushFront(1)
	list.PushBack(3)
	if ok, err := list.Insert(1, 9); err != nil || !ok {
		t.Fatal("valid insert failed")
	}
	for index, want := range []int{1, 9, 2, 3} {
		if got, ok, err := list.Get(index); err != nil || !ok || got != want {
			t.Fatalf("Get(%d) = (%d, %t, %v), want (%d, true, nil)", index, got, ok, err, want)
		}
	}
	if got, ok := list.PopBack(); !ok || got != 3 {
		t.Fatal("PopBack must return tail")
	}
	if ok, err := list.Insert(-1, 0); ok || !errors.Is(err, ErrInvalidIndex) {
		t.Fatal("invalid insert must fail")
	}
}
