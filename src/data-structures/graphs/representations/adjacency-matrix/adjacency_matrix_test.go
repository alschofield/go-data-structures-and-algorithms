//go:build contract

package adjacency_matrix

import (
	"errors"
	"testing"
)

func TestAdjacencyMatrix(t *testing.T) {
	graph := NewAdjacencyMatrix(false)
	for _, value := range []string{"a", "b", "c"} {
		graph.AddVertex(value)
	}
	added, err := graph.AddEdge(0, 1, 4)
	hasReverse, err := graph.HasEdge(1, 0)
	duplicate, err := graph.AddEdge(0, 1, 4)
	if err != nil || !added || !hasReverse || duplicate {
		t.Fatal("undirected edges must mirror once and reject duplicates")
	}
	removed, err := graph.RemoveEdge(0, 1)
	hasForward, err := graph.HasEdge(0, 1)
	hasReverse, err = graph.HasEdge(1, 0)
	if err != nil || !removed || hasForward || hasReverse {
		t.Fatal("removal must clear both matrix cells")
	}
	if removed, err := graph.RemoveEdge(0, 1); err != nil || removed {
		t.Fatal("absent mutation must be a normal no-op")
	}
	if _, err := graph.AddEdge(3, 0, 1); !errors.Is(err, ErrInvalidEdge) {
		t.Fatal("absent or invalid mutations must fail cleanly")
	}
	if got, ok, err := graph.VertexAt(2); err != nil || !ok || got != "c" {
		t.Fatal("VertexAt must retain insertion order")
	}
}
