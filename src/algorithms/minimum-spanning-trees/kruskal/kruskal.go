package kruskal

import (
	"errors"

	quick_sort "github.com/alschofield/go-data-structures-and-algorithms/src/algorithms/sorting/comparison/quick-sort"
	union_find "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/disjoint-sets/union-find"
	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

var ErrNilGraph error = errors.New("function requires a valid graph.")

func Kruskal[T any](giraffe graph.UndirectedEdgeGraph[T]) ([]graph.Edge[T], int64, error) {
	result := []graph.Edge[T]{}
	var total int64 = 0

	if giraffe == nil {
		return nil, 0, ErrNilGraph
	}

	if giraffe.Directed() {
		return nil, 0, graph.ErrDirectedGraph
	}

	node_count := giraffe.NodeCount()
	edges := []graph.Edge[T]{}

	status, err := giraffe.Edges(func(edge graph.Edge[T]) bool {
		edges = append(edges, edge)
		return true
	})

	if err != nil {
		return nil, 0, err
	}

	if !status {
		return nil, 0, nil
	}

	// Kruskal is greedy by weight: consider every logical edge cheapest-first.
	quick_sort.QuickSort(edges, func(left graph.Edge[T], right graph.Edge[T]) int {
		if left.Weight > right.Weight {
			return 1
		} else if left.Weight < right.Weight {
			return -1
		} else {
			return 0
		}
	})

	union, err := union_find.NewUnionFind(node_count)

	if err != nil {
		return nil, 0, err
	}

	// Graph keys are sparse (10, 20, ...), but union-find accepts only dense
	// indices 0..count-1, so each key is assigned the next unused index once.
	index_by_key := map[int]int{}
	index_of := func(key int) int {
		index, present := index_by_key[key]
		if !present {
			index = len(index_by_key)
			index_by_key[key] = index
		}
		return index
	}

	for i := 0; i < len(edges); i++ {
		from_key, status, err := union.Find(index_of(edges[i].From.Key))
		if err != nil {
			return nil, 0, err
		}

		if !status {
			return nil, 0, nil
		}

		to_key, status, err := union.Find(index_of(edges[i].To.Key))
		if err != nil {
			return nil, 0, err
		}

		if !status {
			return nil, 0, nil
		}

		// Distinct roots mean this edge joins two components; shared roots
		// mean it would close a cycle, so it is skipped.
		if from_key != to_key {
			result = append(result, edges[i])
			total += edges[i].Weight
			_, err := union.Union(index_of(edges[i].From.Key), index_of(edges[i].To.Key))

			if err != nil {
				return nil, 0, err
			}

			// A spanning tree holds exactly V-1 edges; further edges only cycle.
			if len(result) == node_count-1 {
				break
			}
		}
	}

	return result, total, nil
}
