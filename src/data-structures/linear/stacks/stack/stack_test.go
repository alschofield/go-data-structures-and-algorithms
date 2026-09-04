//go:build contract

package stack

import (
	"errors"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	if _, err := stack.Pop(); !errors.Is(err, ErrEmptyStack) {
		t.Fatalf("empty Pop() error = %v, want ErrEmptyStack", err)
	}
	for _, value := range []int{1, 2, 3} {
		if !stack.Push(value) {
			t.Fatal("push failed")
		}
	}
	for _, want := range []int{3, 2, 1} {
		if got, err := stack.Pop(); err != nil || got != want {
			t.Fatalf("Pop() = (%d, %v), want (%d, nil)", got, err, want)
		}
	}
	if !stack.IsEmpty() || stack.Len() != 0 {
		t.Fatal("popping all values must restore empty state")
	}
}
