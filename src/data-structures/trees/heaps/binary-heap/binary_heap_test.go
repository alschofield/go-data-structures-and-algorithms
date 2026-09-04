//go:build contract

package binary_heap

import (
	"errors"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	heap, err := NewBinaryHeap(func(a, b int) int { return a - b })
	if err != nil {
		t.Fatalf("constructor error = %v", err)
	}
	for _, value := range []int{3, 1, 4, 2} {
		if !heap.Push(value) {
			t.Fatal("push failed")
		}
	}
	for _, want := range []int{4, 3, 2, 1} {
		if got, ok := heap.Pop(); !ok || got != want {
			t.Fatalf("Pop() = (%d, %t), want (%d, true)", got, ok, want)
		}
	}
	if _, err := NewBinaryHeap[int](nil); !errors.Is(err, ErrNilComparator) {
		t.Fatalf("nil comparator error = %v, want ErrNilComparator", err)
	}
}
