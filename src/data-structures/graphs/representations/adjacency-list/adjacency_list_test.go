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
		node := list.AddVertex(&values[index])
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
	if complete, err := list.Edges(func(edge graphcontract.Edge[string]) bool { edges = append(edges, edge); return true }); err != nil || !complete || len(edges) != 1 || edges[0] != (graphcontract.Edge[string]{From: nodes[0], To: nodes[1], Weight: 4}) {
		t.Fatal("undirected Edges must report one logical edge")
	}
	if _, err := list.AddEdge(-1, nodes[0].Key, 1); !errors.Is(err, graphcontract.ErrInvalidKey) || list.NodeCount() != 3 {
		t.Fatal("invalid keys must preserve graph")
	}
	if got, ok := list.NodeByKey(nodes[2].Key); !ok || got != nodes[2] || got.Value != &values[2] {
		t.Fatal("NodeByKey must retain the caller value pointer")
	}
	if _, ok := list.NodeByKey(-1); ok {
		t.Fatal("invalid NodeByKey lookup must return ok=false")
	}
	if removed, err := list.RemoveEdge(nodes[0].Key, nodes[1].Key); err != nil || !removed || list.EdgeCount() != 0 {
		t.Fatal("RemoveEdge must remove one logical undirected edge")
	}
	if removed, err := list.RemoveEdge(nodes[0].Key, nodes[1].Key); err != nil || removed || list.EdgeCount() != 0 {
		t.Fatal("removing an absent edge must be a non-mutating no-op")
	}
}

func TestAdjacencyListRejectsDirectedEdgesView(t *testing.T) {
	value_a, value_b := "a", "b"
	list := NewAdjacencyList[string](true)
	from := list.AddVertex(&value_a)
	to := list.AddVertex(&value_b)
	if added, err := list.AddEdge(from.Key, to.Key, 1); err != nil || !added {
		t.Fatal("directed edge setup failed")
	}
	if _, err := list.Edges(func(graphcontract.Edge[string]) bool { return true }); !errors.Is(err, graphcontract.ErrDirectedGraph) {
		t.Fatalf("directed Edges() error = %v, want ErrDirectedGraph", err)
	}
}
