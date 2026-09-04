//go:build contract

package queue

import (
	"errors"
	"testing"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name        string
		enqueue     []int
		dequeue     int
		enqueueMore []int
		want        []int
	}{
		{
			name:    "FIFO order",
			enqueue: []int{1, 2, 3},
			want:    []int{1, 2, 3},
		},
		{
			name:        "wraparound after dequeue and enqueue",
			enqueue:     []int{1, 2},
			dequeue:     1,
			enqueueMore: []int{3},
			want:        []int{2, 3},
		},
		{
			name:        "resize while wrapped preserves FIFO order",
			enqueue:     []int{1, 2, 3, 4},
			dequeue:     2,
			enqueueMore: []int{5, 6, 7},
			want:        []int{3, 4, 5, 6, 7},
		},
		{
			name: "multiple grows preserve FIFO order",
			enqueue: func() []int {
				values := make([]int, 65)
				for i := range values {
					values[i] = i
				}
				return values
			}(),
			want: func() []int {
				values := make([]int, 65)
				for i := range values {
					values[i] = i
				}
				return values
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			queue := NewQueue[int]()
			for _, value := range test.enqueue {
				if !queue.Enqueue(value) {
					t.Fatal("Enqueue() = false, want true")
				}
			}
			for range test.dequeue {
				if _, err := queue.Dequeue(); err != nil {
					t.Fatalf("setup Dequeue() error = %v, want nil", err)
				}
			}
			for _, value := range test.enqueueMore {
				if !queue.Enqueue(value) {
					t.Fatal("Enqueue() = false, want true")
				}
			}

			assertQueueContents(t, queue, test.want)
		})
	}
}

func TestQueuePeekDoesNotMutate(t *testing.T) {
	queue := NewQueue[int]()
	for _, value := range []int{10, 20, 30} {
		queue.Enqueue(value)
	}

	beforeLen, beforeHead, beforeTail := queue.Len(), queue.head, queue.tail
	for range 2 {
		if got, err := queue.Peek(); err != nil || got != 10 {
			t.Fatalf("Peek() = (%d, %v), want (10, nil)", got, err)
		}
	}
	if queue.Len() != beforeLen || queue.head != beforeHead || queue.tail != beforeTail {
		t.Fatal("Peek() mutated queue state")
	}
	assertQueueContents(t, queue, []int{10, 20, 30})
}

func TestQueueLenIsEmptyAndEmptyErrors(t *testing.T) {
	queue := NewQueue[int]()
	if queue.Len() != 0 || !queue.IsEmpty() {
		t.Fatalf("new queue state = Len() %d, IsEmpty() %t; want 0, true", queue.Len(), queue.IsEmpty())
	}

	beforeHead, beforeTail := queue.head, queue.tail
	for _, operation := range []struct {
		name string
		call func() (int, error)
	}{
		{name: "Dequeue", call: queue.Dequeue},
		{name: "Peek", call: queue.Peek},
	} {
		t.Run(operation.name, func(t *testing.T) {
			if got, err := operation.call(); got != 0 || !errors.Is(err, ErrEmptyQueue) {
				t.Fatalf("%s() = (%d, %v), want (0, ErrEmptyQueue)", operation.name, got, err)
			}
		})
	}
	if queue.Len() != 0 || !queue.IsEmpty() || queue.head != beforeHead || queue.tail != beforeTail {
		t.Fatal("empty reads mutated queue state")
	}

	queue.Enqueue(1)
	if queue.Len() != 1 || queue.IsEmpty() {
		t.Fatalf("populated queue state = Len() %d, IsEmpty() %t; want 1, false", queue.Len(), queue.IsEmpty())
	}
	assertQueueContents(t, queue, []int{1})
}

func TestQueueDequeueClearsReference(t *testing.T) {
	queue := NewQueue[*int]()
	value := new(int)
	queue.Enqueue(value)
	removedSlot := queue.head

	if got, err := queue.Dequeue(); err != nil || got != value {
		t.Fatalf("Dequeue() = (%v, %v), want (%v, nil)", got, err, value)
	}
	if queue.items[removedSlot] != nil {
		t.Fatal("Dequeue() retained the removed reference in its backing buffer")
	}
}

func assertQueueContents(t *testing.T, queue *Queue[int], want []int) {
	t.Helper()
	if got := queue.Len(); got != len(want) {
		t.Fatalf("Len() = %d, want %d", got, len(want))
	}
	for _, expected := range want {
		if got, err := queue.Dequeue(); err != nil || got != expected {
			t.Fatalf("Dequeue() = (%d, %v), want (%d, nil)", got, err, expected)
		}
	}
	if !queue.IsEmpty() || queue.Len() != 0 {
		t.Fatal("draining queue must restore empty state")
	}
}
