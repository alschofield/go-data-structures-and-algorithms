package binary_heap

import (
	"errors"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

// ErrNilComparator reports that a heap cannot order values without a comparator.
var ErrNilComparator error = errors.New("BinaryHeap requires a valid compare function.")

// BinaryHeap stores a max heap in a contiguous slice of shared nodes.
type BinaryHeap[T any] struct {
	size     int
	next_key int
	compare  func(T, T) int
	heap     []*graph.Node[T]
}

// NewBinaryHeap creates an empty heap after validating its comparator.
func NewBinaryHeap[T any](compare func(T, T) int) (*BinaryHeap[T], error) {
	if compare == nil {
		return nil, ErrNilComparator
	}

	return &BinaryHeap[T]{
		size:    0,
		compare: compare,
	}, nil
}

// Push appends a node then sifts it toward the root until heap order holds.
func (bh *BinaryHeap[T]) Push(value T) bool {
	bh.heap = append(bh.heap, &graph.Node[T]{Key: bh.next_key, Value: &value})
	bh.next_key++
	bh.size++

	// Each swap moves the newly appended value one parent level upward.
	for i := bh.size - 1; bh.compare(*bh.heap[i].Value, *bh.heap[(i-1)/2].Value) > 0; i = (i - 1) / 2 {
		bh.heap[i], bh.heap[(i-1)/2] = bh.heap[(i-1)/2], bh.heap[i]
	}

	return true
}

// Pop removes the maximum root, replaces it with the final node, then sifts down.
func (bh *BinaryHeap[T]) Pop() (T, bool) {
	if bh.size == 0 {
		var zero T
		return zero, false
	}

	popped := bh.heap[0]

	// Move the final node to the root and clear its old slot for garbage collection.
	bh.heap[0] = bh.heap[bh.size-1]
	bh.heap[bh.size-1] = nil
	bh.heap = bh.heap[:bh.size-1]

	bh.size--

	if bh.size == 0 {
		return *popped.Value, true
	}

	i := 0
	for {
		// Stop before indexing when the current node has no children.
		var child_index int
		if i*2+1 >= bh.size {
			break
		} else if i*2+2 >= bh.size {
			child_index = i*2 + 1
		} else {
			comparison := bh.compare(*bh.heap[i*2+1].Value, *bh.heap[i*2+2].Value)

			if comparison > 0 {
				child_index = i*2 + 1
			} else {
				child_index = i*2 + 2
			}
		}

		// Restore the max-heap invariant by swapping with the larger child.
		comparison := bh.compare(*bh.heap[i].Value, *bh.heap[child_index].Value)
		if comparison < 0 {
			bh.heap[i], bh.heap[child_index] = bh.heap[child_index], bh.heap[i]
			i = child_index
		} else {
			break
		}
	}

	return *popped.Value, true
}

// Peek returns the maximum root without changing heap structure.
func (bh *BinaryHeap[T]) Peek() (T, bool) {
	if bh.size == 0 {
		var zero T
		return zero, false
	}

	return *bh.heap[0].Value, true
}

func (bh *BinaryHeap[T]) Len() int {
	return bh.size
}

func (bh *BinaryHeap[T]) IsEmpty() bool {
	return bh.size == 0
}
