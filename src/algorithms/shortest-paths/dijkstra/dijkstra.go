package dijkstra

import (
	"errors"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
	binary_heap "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/trees/heaps/binary-heap"
)

var ErrNilGraph error = errors.New("function requires a valid graph.")
var ErrNegativeWeight error = errors.New("function requires non negative weights.")

type DijkstraResult struct {
	Distance func(key int) (int64, bool)
	Parent   func(key int) (int, bool)
}

type DijkstraNode[T any] struct {
	node     *graph.Node[T]
	distance int64
}

func Dijkstra[T any](giraffe graph.Graph[T], source_key int) (DijkstraResult, error) {
	// distances stores the best total edge weight currently known from source_key.
	distances := map[int]int64{source_key: 0}
	// parents stores the predecessor key used to reach each node's best distance.
	parents := map[int]int{}
	result := DijkstraResult{
		Distance: func(key int) (int64, bool) {
			distance, presence := distances[key]
			return distance, presence
		},
		Parent: func(key int) (int, bool) {
			parent, presence := parents[key]
			return parent, presence
		},
	}

	if giraffe == nil {
		return result, ErrNilGraph
	}

	source_node, status := giraffe.NodeByKey(source_key)
	if !status || source_node == nil {
		return result, graph.ErrInvalidKey
	}

	// BinaryHeap is a max-heap, so a smaller tentative distance receives higher priority.
	heap, err := binary_heap.NewBinaryHeap[*DijkstraNode[T]](func(left *DijkstraNode[T], right *DijkstraNode[T]) int {
		if left.distance > right.distance {
			return -1
		} else if left.distance < right.distance {
			return 1
		} else {
			return 0
		}
	})

	if err != nil {
		return result, err
	}

	if !heap.Push(&DijkstraNode[T]{node: source_node, distance: distances[source_key]}) {
		return result, nil
	}

	for !heap.IsEmpty() {
		node, status := heap.Pop()
		if !status {
			break
		}

		// A later relaxation may have queued a better candidate for this key.
		if distances[node.node.Key] != node.distance {
			continue
		}

		success, err := giraffe.Neighbors(node.node.Key, func(neighbor *graph.Node[T], weight int64) bool {
			// Relaxation asks whether traveling through the popped node improves this neighbor.
			new_distance := distances[node.node.Key] + weight
			if weight < 0 {
				return false
			}

			distance, presence := distances[neighbor.Key]

			if !presence || new_distance < distance {
				// Record both the cheaper total cost and the edge that achieved it.
				distances[neighbor.Key] = new_distance
				parents[neighbor.Key] = node.node.Key
				// Push a snapshot; the older, more expensive snapshot becomes stale.
				return heap.Push(&DijkstraNode[T]{node: neighbor, distance: new_distance})
			}

			return true
		})

		if err != nil {
			return result, err
		}

		if !success {
			return result, ErrNegativeWeight
		}
	}

	return result, nil
}
