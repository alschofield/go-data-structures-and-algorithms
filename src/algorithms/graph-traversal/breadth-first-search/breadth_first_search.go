package breadth_first_search

import (
	"errors"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/linear/queues/queue"
)

// ErrNilGraph reports that traversal requires a graph instance.
var ErrNilGraph error = errors.New("function requires a valid graph.")

// ErrNilVisit reports that traversal requires a visitor callback.
var ErrNilVisit error = errors.New("function requires a valid visit function.")

// BreadthFirstSearch visits nodes by increasing unweighted distance from source_key.
func BreadthFirstSearch[T any](giraffe graph.Graph[T], source_key int, visit func(*graph.Node[T]) bool) ([]*graph.Node[T], bool, error) {
	// A FIFO queue preserves breadth-first order.
	queue := queue.NewQueue[*graph.Node[T]]()
	// Stable node keys prevent cycles and duplicate enqueueing.
	visited := make(map[int]bool)
	// Path records every node actually dequeued, including an early-stop node.
	path := []*graph.Node[T]{}

	if giraffe == nil {
		return path, false, ErrNilGraph
	}

	if visit == nil {
		return path, false, ErrNilVisit
	}

	// Resolve the source before allocating traversal state around an invalid key.
	source_node, status := giraffe.NodeByKey(source_key)
	if !status {
		return path, false, graph.ErrInvalidKey
	}

	if !queue.Enqueue(source_node) {
		return path, false, nil
	}

	// Mark on enqueue so another parent cannot queue the same node again.
	visited[source_node.Key] = true

	for !queue.IsEmpty() {
		node, err := queue.Dequeue()
		if err != nil {
			return path, false, err
		}

		// The dequeue order is the observable breadth-first traversal order.
		path = append(path, node)

		if visit(node) {
			// Discover the next breadth level only after visiting this node.
			status, err := giraffe.Neighbors(node.Key, func(neighbor *graph.Node[T], weight int64) bool {
				if !visited[neighbor.Key] {
					queue.Enqueue(neighbor)
					visited[neighbor.Key] = true
				}

				return true
			})

			if err != nil {
				return path, false, err
			}

			if !status {
				return path, false, nil
			}
		} else {
			// Caller-defined success, commonly a matching search result.
			return path, true, nil
		}
	}

	return path, true, nil
}
