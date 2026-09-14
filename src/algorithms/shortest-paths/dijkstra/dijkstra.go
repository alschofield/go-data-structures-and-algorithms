package dijkstra

import (
	"errors"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/linear/queues/queue"
)

var ErrNilGraph error = errors.New("function requires a valid graph.")
var ErrNegativeWeight error = errors.New("function requires non negative weights.")

type DijkstraResult struct {
	Distance func(key int) (int64, bool)
	Parent   func(key int) (int, bool)
}

func Dijkstra[T any](giraffe graph.Graph[T], source_key int) (DijkstraResult, error) {
	result := DijkstraResult{}
	// is there supposed to be like a weights array or something?

	if giraffe == nil {
		return result, ErrNilGraph
	}

	source_node, status := giraffe.NodeByKey(source_key)
	if !status || source_node == nil {
		return result, graph.ErrInvalidKey
	}

	queue := queue.NewQueue[*graph.Node[T]]()

	// we should be storing like scores on these nodes
	if !queue.Enqueue(source_node) {
		return result, nil
	}

	for !queue.IsEmpty() {
		node, err := queue.Dequeue()
		if err != nil {
			return result, err
		}

		// dijkstra stuff
		// something like adding to the final path based on the total weight
		// maybe parents and weights are an array of the final shortest path that is accessed by the returned DR methods?
		// distance is the total weight of the path?
		// parent is the pointer at the provided key?
		// youd get the path by traversing Parent with the source key?

		success, err := giraffe.Neighbors(node.Key, func(neighbor *graph.Node[T], weight int64) bool {
			// this should like add to the score for the node that is added
			// maybe it needs to be a copy of the node so it doesnt screw up past nodes?
			return !queue.Enqueue(neighbor)
		})

		if err != nil {
			return result, err
		}

		if !success {
			return result, nil
		}
	}

	return DijkstraResult{}, nil
}
