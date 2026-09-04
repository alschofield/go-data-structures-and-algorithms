//go:build contract

package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	if _, ok := stack.Pop(); ok {
		t.Fatal("empty pop must fail")
	}
	for _, value := range []int{1, 2, 3} {
		if !stack.Push(value) {
			t.Fatal("push failed")
		}
	}
	for _, want := range []int{3, 2, 1} {
		if got, ok := stack.Pop(); !ok || got != want {
			t.Fatalf("Pop() = (%d, %t), want (%d, true)", got, ok, want)
		}
	}
	if !stack.IsEmpty() || stack.Len() != 0 {
		t.Fatal("popping all values must restore empty state")
	}
}
