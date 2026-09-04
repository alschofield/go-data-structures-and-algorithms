//go:build contract

package dynamic_array

import (
	"errors"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	a := NewDynamicArray[int]()
	for _, value := range []int{1, 3} {
		if ok, err := a.Insert(a.Len(), value); err != nil || !ok {
			t.Fatal("append insert failed")
		}
	}
	if ok, err := a.Insert(1, 2); err != nil || !ok || a.Len() != 3 {
		t.Fatal("insert must shift and grow length")
	}
	for index, want := range []int{1, 2, 3} {
		if got, ok, err := a.Get(index); err != nil || !ok || got != want {
			t.Fatalf("Get(%d) = (%d, %t, %v), want (%d, true, nil)", index, got, ok, err, want)
		}
	}
	if got, ok, err := a.Remove(1); err != nil || !ok || got != 2 {
		t.Fatal("Remove must return and remove indexed value")
	}
	if ok, err := a.Insert(-1, 0); ok || !errors.Is(err, ErrInvalidIndex) || a.Len() != 2 {
		t.Fatal("invalid insert must preserve state")
	}
}
