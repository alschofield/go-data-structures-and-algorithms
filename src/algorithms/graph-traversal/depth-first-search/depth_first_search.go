package depth_first_search

import (
	"errors"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/linear/stacks/stack"
)

// ErrNilGraph reports that traversal requires a graph instance.
var ErrNilGraph error = errors.New("function requires a valid graph.")

// ErrNilVisit reports that traversal requires a visitor callback.
var ErrNilVisit error = errors.New("function requires a valid visit function.")

// DepthFirstSearch explores the most recently discovered branch before backtracking.
func DepthFirstSearch[T any](input_graph graph.Graph[T], source_key int, visit func(*graph.Node[T]) bool) ([]*graph.Node[T], bool, error) {
	// A LIFO stack drives depth-first order.
	stack := stack.NewStack[*graph.Node[T]]()
	// Stable node keys prevent cycles and duplicate pushes.
	visited := make(map[int]bool)
	// Path contains every node popped from the traversal stack.
	path := []*graph.Node[T]{}

	if input_graph == nil {
		return path, false, ErrNilGraph
	}

	if visit == nil {
		return path, false, ErrNilVisit
	}

	// Resolve the source before starting a traversal for an invalid key.
	source_node, status := input_graph.NodeByKey(source_key)

	if !status {
		return path, status, graph.ErrInvalidKey
	}

	if source_node == nil {
		return path, false, graph.ErrInvalidKey
	}

	if !stack.Push(source_node) {
		return path, false, nil
	}

	// Mark on push so a later edge cannot schedule the source again.
	visited[source_key] = true

	for !stack.IsEmpty() {
		node, err := stack.Pop()
		if err != nil {
			return path, false, err
		}

		// A pop defines the observable depth-first visit order.
		path = append(path, node)

		if visit(node) {
			// Neighbors are pushed in representation order; the latest is popped first.
			status, err := input_graph.Neighbors(node.Key, func(neighbor *graph.Node[T], weight int64) bool {
				if !visited[neighbor.Key] {
					if !stack.Push(neighbor) {
						return false
					}

					visited[neighbor.Key] = true
				}

				return true
			})

			if err != nil {
				return path, false, err
			}

			if !status {
				return path, status, nil
			}
		} else {
			// Caller-defined success, commonly a matching search result.
			return path, true, nil
		}
	}

	return path, true, nil
}
