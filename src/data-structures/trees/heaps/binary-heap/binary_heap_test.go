//go:build contract

package binary_heap

import "testing"

func TestBinaryHeap(t *testing.T) {
	heap := NewBinaryHeap(func(a, b int) int { return a - b })
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
}
