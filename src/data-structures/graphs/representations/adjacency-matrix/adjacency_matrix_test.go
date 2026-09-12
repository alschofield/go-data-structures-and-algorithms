//go:build contract

package adjacency_matrix

import (
	"errors"
	"testing"

	graphcontract "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func TestAdjacencyMatrix(t *testing.T) {
	values := []string{"a", "b", "c"}
	matrix := NewAdjacencyMatrix[string](false)
	var _ graphcontract.Graph[string] = matrix
	var _ graphcontract.UndirectedEdgeGraph[string] = matrix

	nodes := make([]*graphcontract.Node[string], 0, len(values))
	for index := range values {
		node := matrix.AddVertex(&values[index])
		nodes = append(nodes, node)
	}
	added, err := matrix.AddEdge(nodes[0].Key, nodes[1].Key, 4)
	has_reverse, err := matrix.HasEdge(nodes[1].Key, nodes[0].Key)
	duplicate, err := matrix.AddEdge(nodes[0].Key, nodes[1].Key, 4)
	if err != nil || !added || !has_reverse || duplicate {
		t.Fatal("undirected edges must mirror once and reject duplicates")
	}

	var edges []graphcontract.Edge[string]
	if complete, err := matrix.Edges(func(edge graphcontract.Edge[string]) bool { edges = append(edges, edge); return true }); err != nil || !complete || len(edges) != 1 || edges[0] != (graphcontract.Edge[string]{From: nodes[0], To: nodes[1], Weight: 4}) {
		t.Fatal("undirected Edges must report one logical edge")
	}
	removed, err := matrix.RemoveEdge(nodes[0].Key, nodes[1].Key)
	has_forward, err := matrix.HasEdge(nodes[0].Key, nodes[1].Key)
	has_reverse, err = matrix.HasEdge(nodes[1].Key, nodes[0].Key)
	if err != nil || !removed || has_forward || has_reverse {
		t.Fatal("removal must clear both matrix cells")
	}
	if removed, err := matrix.RemoveEdge(nodes[0].Key, nodes[1].Key); err != nil || removed {
		t.Fatal("absent mutation must be a normal no-op")
	}
	if _, err := matrix.AddEdge(999, nodes[0].Key, 1); !errors.Is(err, graphcontract.ErrInvalidKey) {
		t.Fatal("absent or invalid mutations must fail cleanly")
	}
	if got, ok := matrix.NodeByKey(nodes[2].Key); !ok || got != nodes[2] || got.Value != &values[2] {
		t.Fatal("NodeByKey must retain the caller value pointer")
	}
}
