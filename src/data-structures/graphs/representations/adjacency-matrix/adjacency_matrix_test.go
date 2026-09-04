//go:build contract

package adjacency_matrix

import "testing"

func TestAdjacencyMatrix(t *testing.T) {
	graph := NewAdjacencyMatrix(false)
	for _, value := range []string{"a", "b", "c"} {
		graph.AddVertex(value)
	}
	if !graph.AddEdge(0, 1, 4) || !graph.HasEdge(1, 0) || graph.AddEdge(0, 1, 4) {
		t.Fatal("undirected edges must mirror once and reject duplicates")
	}
	if !graph.RemoveEdge(0, 1) || graph.HasEdge(0, 1) || graph.HasEdge(1, 0) {
		t.Fatal("removal must clear both matrix cells")
	}
	if graph.RemoveEdge(0, 1) || graph.AddEdge(3, 0, 1) {
		t.Fatal("absent or invalid mutations must fail cleanly")
	}
	if got, ok := graph.VertexAt(2); !ok || got != "c" {
		t.Fatal("VertexAt must retain insertion order")
	}
}
