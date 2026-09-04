package queue

import (
	"errors"
)

var ErrEmptyQueue = errors.New("queue is empty")

// Queue stores its logical FIFO sequence in a circular backing buffer.
type Queue[T any] struct {
	// size distinguishes an empty queue from a full one when head equals tail.
	size int
	// head identifies the next value returned by Dequeue or Peek.
	head int
	// tail identifies the unused slot that receives the next enqueued value.
	tail int
	// items retains allocated capacity so later enqueues can reuse it.
	items []T
}

// NewQueue creates an empty queue whose backing buffer is allocated on demand.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue appends value after every existing value in FIFO order.
func (q *Queue[T]) Enqueue(value T) bool {
	// A full buffer has no unused tail slot, so grow it before writing.
	if len(q.items) == q.size {
		var new_capacity int
		if q.size == 0 {
			// The first allocation establishes the minimum circular buffer size.
			new_capacity = 2
		} else {
			// Doubling makes occasional copies amortize to constant-time enqueues.
			new_capacity = q.size * 2
		}

		var items []T = make([]T, new_capacity)
		// Copy in logical order because a full ring may be physically wrapped.
		for index := range q.items {
			items[index] = q.items[(q.head+index)%len(q.items)]
		}

		// The copied sequence is contiguous, so reset the circular indexes.
		q.items = items
		q.head = 0
		q.tail = q.size
	}

	// tail is always an available slot after the full-buffer check.
	q.items[q.tail] = value
	q.size++
	// Modulo wraps tail to the start without shifting stored values.
	q.tail = (q.tail + 1) % len(q.items)
	return true
}

// Dequeue removes and returns the oldest queued value.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.size == 0 {
		// Return the generic zero value alongside the stable empty-queue sentinel.
		var empty T
		return empty, ErrEmptyQueue
	}

	var zeroed T
	item := q.items[q.head]
	// Clear references before advancing head so removed values are collectible.
	q.items[q.head] = zeroed
	q.size--
	// Modulo keeps head in bounds after it reaches the end of the buffer.
	q.head = (q.head + 1) % len(q.items)
	return item, nil
}

// Peek returns the oldest queued value without changing queue state.
func (q *Queue[T]) Peek() (T, error) {
	if q.size == 0 {
		var empty T
		return empty, ErrEmptyQueue
	}

	return q.items[q.head], nil
}

// Len reports the number of logical values rather than backing-buffer capacity.
func (q *Queue[T]) Len() int {
	return q.size
}

// IsEmpty reports whether no values remain available to dequeue.
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
