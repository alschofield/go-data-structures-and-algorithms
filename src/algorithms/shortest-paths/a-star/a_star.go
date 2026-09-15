package a_star

import (
	"errors"
	"slices"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
	binary_heap "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/trees/heaps/binary-heap"
)

var ErrNilGraph error = errors.New("function requires a valid graph.")
var ErrNilHeuristic error = errors.New("function requires a valid heuristic function.")
var ErrNegativeWeight error = errors.New("function requires non negative weights.")
var ErrNegativeHeuristic error = errors.New("function requires non negative heuristics.")

type Heuristic func(key int) int64

type AStarNode[T any] struct {
	node    *graph.Node[T]
	g_score int64
	f_score int64
}

func AStar[T any](giraffe graph.Graph[T], source_key, goal_key int, heuristic Heuristic) ([]*graph.Node[T], error) {
	// result is assembled only after the goal is removed as the best candidate.
	result := []*graph.Node[T]{}
	// distances records the best actual source-to-node cost, independent of the heuristic.
	distances := map[int]int64{source_key: 0}
	// parents retains the predecessor needed to reconstruct the source-to-goal route.
	parents := map[int]*graph.Node[T]{}

	if giraffe == nil {
		return []*graph.Node[T]{}, ErrNilGraph
	}

	if heuristic == nil {
		return []*graph.Node[T]{}, ErrNilHeuristic
	}

	if heuristic(source_key) < 0 {
		return []*graph.Node[T]{}, ErrNegativeHeuristic
	}

	// BinaryHeap is a max-heap, so lower f and deterministic tie-break values win.
	heap, err := binary_heap.NewBinaryHeap(func(left *AStarNode[T], right *AStarNode[T]) int {
		if left.f_score < right.f_score {
			return 1
		} else if left.f_score > right.f_score {
			return -1
		} else {
			if left.g_score < right.g_score {
				return 1
			} else if left.g_score > right.g_score {
				return -1
			} else {
				if left.node.Key < right.node.Key {
					return 1
				} else if left.node.Key > right.node.Key {
					return -1
				} else {
					return 0
				}
			}
		}
	})

	if err != nil {
		return []*graph.Node[T]{}, err
	}

	source_node, found := giraffe.NodeByKey(source_key)

	if !found {
		return []*graph.Node[T]{}, graph.ErrInvalidKey
	}

	_, found = giraffe.NodeByKey(goal_key)

	if !found {
		return []*graph.Node[T]{}, graph.ErrInvalidKey
	}

	// The source starts at g=0; its initial priority is only the goal estimate.
	if !heap.Push(&AStarNode[T]{node: source_node, g_score: 0, f_score: heuristic(source_key)}) {
		return []*graph.Node[T]{}, nil
	}

	for !heap.IsEmpty() {
		node, found := heap.Pop()
		if !found {
			break
		}

		// A newer candidate may have improved this node's g score after this push.
		if distances[node.node.Key] != node.g_score {
			continue
		}

		if node.node.Key == goal_key {
			// Parents point goal-to-source, so reverse the separate reconstructed path.
			trav_key := goal_key
			result = append(result, node.node)
			for trav_key != source_key {
				result = append(result, parents[trav_key])
				trav_key = parents[trav_key].Key
			}

			slices.Reverse(result)
			return result, nil
		}

		status, err := giraffe.Neighbors(node.node.Key, func(neighbor *graph.Node[T], weight int64) bool {
			// Relax a neighbor only when this route improves its known actual cost.
			if weight < 0 {
				return false
			}

			if heuristic(neighbor.Key) < 0 {
				return false
			}

			next_g := node.g_score + weight
			prev_g, present := distances[neighbor.Key]

			if !present || next_g < prev_g {
				distances[neighbor.Key] = next_g
				parents[neighbor.Key] = node.node

				// A* ranks by actual cost plus the estimate from neighbor to the goal.
				f_score := next_g + heuristic(neighbor.Key)

				// Push an immutable score snapshot; older candidates are skipped later.
				if !heap.Push(&AStarNode[T]{
					node:    neighbor,
					g_score: next_g,
					f_score: f_score,
				}) {
					return false
				}
			}

			return true
		})

		if err != nil {
			return []*graph.Node[T]{}, err
		}

		if !status {
			return []*graph.Node[T]{}, nil
		}
	}

	return result, nil
}
