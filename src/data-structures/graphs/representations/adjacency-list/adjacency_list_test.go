//go:build contract

package adjacency_list

import (
	"errors"
	"testing"
)

func TestAdjacencyList(t *testing.T) {
	graph := NewAdjacencyList(false)
	for _, value := range []string{"a", "b", "c"} {
		graph.AddVertex(value)
	}
	added, err := graph.AddEdge(0, 1, 4)
	hasReverse, err := graph.HasEdge(1, 0)
	duplicate, err := graph.AddEdge(0, 1, 4)
	if err != nil || !added || !hasReverse || duplicate {
		t.Fatal("undirected edges must mirror once and reject duplicates")
	}
	var neighbors []int
	if complete, err := graph.Neighbors(0, func(to int, weight int64) bool { neighbors = append(neighbors, to); return true }); err != nil || !complete || len(neighbors) != 1 || neighbors[0] != 1 {
		t.Fatal("Neighbors must visit deterministic outgoing edges")
	}
	if _, err := graph.AddEdge(-1, 0, 1); !errors.Is(err, ErrInvalidEdge) || graph.VertexCount() != 3 {
		t.Fatal("invalid indexes must preserve graph")
	}
	if got, ok, err := graph.VertexAt(2); err != nil || !ok || got != "c" {
		t.Fatal("VertexAt must retain insertion order")
	}
}
