//go:build contract

package stack

import (
	"errors"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	for _, operation := range []struct {
		name string
		call func() (int, error)
	}{
		{name: "Pop", call: stack.Pop},
		{name: "Peek", call: stack.Peek},
	} {
		t.Run("empty "+operation.name, func(t *testing.T) {
			if got, err := operation.call(); got != 0 || !errors.Is(err, ErrEmptyStack) {
				t.Fatalf("%s() = (%d, %v), want (0, ErrEmptyStack)", operation.name, got, err)
			}
			if stack.Len() != 0 || !stack.IsEmpty() {
				t.Fatalf("empty %s() mutated stack state", operation.name)
			}
		})
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

func TestStackPeekDoesNotMutate(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(10)
	stack.Push(20)

	for range 2 {
		if got, err := stack.Peek(); err != nil || got != 20 {
			t.Fatalf("Peek() = (%d, %v), want (20, nil)", got, err)
		}
	}
	if stack.Len() != 2 || stack.IsEmpty() {
		t.Fatal("Peek() mutated stack state")
	}
	if got, err := stack.Pop(); err != nil || got != 20 {
		t.Fatalf("Pop() after Peek() = (%d, %v), want (20, nil)", got, err)
	}
}

func TestStackZeroValueAndPopClearsReference(t *testing.T) {
	var stack Stack[*int]
	if stack.Len() != 0 || !stack.IsEmpty() {
		t.Fatal("zero-value stack must be empty")
	}

	value := new(int)
	stack.Push(value)
	if got, err := stack.Pop(); err != nil || got != value {
		t.Fatalf("Pop() = (%v, %v), want (%v, nil)", got, err, value)
	}
	if len(stack.items) != 0 || cap(stack.items) == 0 || stack.items[:cap(stack.items)][0] != nil {
		t.Fatal("Pop() retained the removed reference in its backing array")
	}
}
