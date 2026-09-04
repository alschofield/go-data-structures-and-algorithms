//go:build contract

package doubly_linked_list

import (
	"errors"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.PushBack(2)
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
	for _, want := range []int{1, 3, 9, 2} {
		var got int
		var ok bool
		if want == 1 || want == 9 {
			got, ok = list.PopFront()
		} else {
			got, ok = list.PopBack()
		}
		if !ok || got != want {
			t.Fatalf("end removal = (%d, %t), want (%d, true)", got, ok, want)
		}
	}
	if !list.IsEmpty() {
		t.Fatal("final removal must clear both ends")
	}
	if ok, err := list.Insert(-1, 0); ok || !errors.Is(err, ErrInvalidIndex) {
		t.Fatalf("invalid index = (%t, %v), want (false, ErrInvalidIndex)", ok, err)
	}
}
