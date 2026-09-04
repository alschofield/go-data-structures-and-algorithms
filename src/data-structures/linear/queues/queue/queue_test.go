//go:build contract

package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()
	if _, ok := queue.Dequeue(); ok {
		t.Fatal("empty dequeue must fail")
	}
	for _, value := range []int{1, 2, 3} {
		if !queue.Enqueue(value) {
			t.Fatal("enqueue failed")
		}
	}
	for _, want := range []int{1, 2, 3} {
		if got, ok := queue.Dequeue(); !ok || got != want {
			t.Fatalf("Dequeue() = (%d, %t), want (%d, true)", got, ok, want)
		}
	}
	if !queue.IsEmpty() || queue.Len() != 0 {
		t.Fatal("draining queue must restore empty state")
	}
}
