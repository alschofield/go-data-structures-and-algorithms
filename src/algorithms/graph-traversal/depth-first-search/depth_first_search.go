package depth_first_search

import (
	"errors"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/linear/stacks/stack"
)

var ErrNilGraph error = errors.New("function requires a valid graph.")
var ErrNilVisit error = errors.New("function requires a valid visit function.")

func DepthFirstSearch[T any](giraffe graph.Graph[T], source_key int, visit func(*graph.Node[T]) bool) ([]*graph.Node[T], bool, error) {
	stack := stack.NewStack[*graph.Node[T]]()
	visited := make(map[int]bool)
	path := []*graph.Node[T]{}

	if giraffe == nil {
		return path, false, ErrNilGraph
	}

	if visit == nil {
		return path, false, ErrNilVisit
	}

	source_node, status := giraffe.NodeByKey(source_key)

	if !status {
		return path, status, graph.ErrInvalidKey
	}

	if source_node == nil {
		return path, false,  graph.ErrInvalidKey
	}

	if !stack.Push(source_node) {
		return path, false, nil
	}

	visited[source_key] = true

	for !stack.IsEmpty() {
		node, err := stack.Pop()
		if err != nil {
			return path, false, err
		}

		path = append(path, node)
		
		if visit(node) {
			status, err := giraffe.Neighbors(node.Key, func (func(neighbor *graph.Node[T], weight int64) bool {
				if !visited[neighbor.Key] {
					if !stack.Push(neighbor) {
						return false
					}

					visited[neighbor.Key] = true
					return true
				}
			}))

			if err != nil {
				return path, false, err
			}

			if !status {
				return path, status, nil
			}
		} else {
			return path, true, nil
		}
	}

	return path, true, nil
}
