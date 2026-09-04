//go:build contract

package adjacency_list

import "testing"

func TestAdjacencyList(t *testing.T) {
	graph := NewAdjacencyList(false)
	for _, value := range []string{"a", "b", "c"} {
		graph.AddVertex(value)
	}
	if !graph.AddEdge(0, 1, 4) || !graph.HasEdge(1, 0) || graph.AddEdge(0, 1, 4) {
		t.Fatal("undirected edges must mirror once and reject duplicates")
	}
	var neighbors []int
	if !graph.Neighbors(0, func(to int, weight int64) bool { neighbors = append(neighbors, to); return true }) || len(neighbors) != 1 || neighbors[0] != 1 {
		t.Fatal("Neighbors must visit deterministic outgoing edges")
	}
	if graph.AddEdge(-1, 0, 1) || graph.VertexCount() != 3 {
		t.Fatal("invalid indexes must preserve graph")
	}
	if got, ok := graph.VertexAt(2); !ok || got != "c" {
		t.Fatal("VertexAt must retain insertion order")
	}
}
