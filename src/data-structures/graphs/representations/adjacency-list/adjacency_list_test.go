//go:build contract

package adjacency_list

import (
	"errors"
	"testing"

	graphcontract "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func TestAdjacencyList(t *testing.T) {
	values := []string{"a", "b", "c"}
	list := NewAdjacencyList[string](false)
	var _ graphcontract.Graph[string] = list
	var _ graphcontract.UndirectedEdgeGraph[string] = list

	nodes := make([]*graphcontract.Node[string], 0, len(values))
	for index := range values {
		node, err := list.AddVertex(&values[index])
		if err != nil {
			t.Fatalf("AddVertex() error = %v", err)
		}
		nodes = append(nodes, node)
	}
	added, err := list.AddEdge(nodes[0].Key, nodes[1].Key, 4)
	has_reverse, err := list.HasEdge(nodes[1].Key, nodes[0].Key)
	duplicate, err := list.AddEdge(nodes[0].Key, nodes[1].Key, 4)
	if err != nil || !added || !has_reverse || duplicate {
		t.Fatal("undirected edges must mirror once and reject duplicates")
	}

	var neighbors []*graphcontract.Node[string]
	if complete, err := list.Neighbors(nodes[0].Key, func(node *graphcontract.Node[string], weight int64) bool {
		neighbors = append(neighbors, node)
		return true
	}); err != nil || !complete || len(neighbors) != 1 || neighbors[0] != nodes[1] {
		t.Fatal("Neighbors must visit deterministic outgoing nodes")
	}
	var edges []graphcontract.Edge[string]
	if complete := list.Edges(func(edge graphcontract.Edge[string]) bool { edges = append(edges, edge); return true }); !complete || len(edges) != 1 || edges[0] != (graphcontract.Edge[string]{From: nodes[0], To: nodes[1], Weight: 4}) {
		t.Fatal("undirected Edges must report one logical edge")
	}
	if _, err := list.AddEdge(-1, nodes[0].Key, 1); !errors.Is(err, graphcontract.ErrInvalidKey) || list.NodeCount() != 3 {
		t.Fatal("invalid keys must preserve graph")
	}
	if got, ok, err := list.NodeByKey(nodes[2].Key); err != nil || !ok || got != nodes[2] || got.Value != &values[2] {
		t.Fatal("NodeByKey must retain the caller value pointer")
	}
}
