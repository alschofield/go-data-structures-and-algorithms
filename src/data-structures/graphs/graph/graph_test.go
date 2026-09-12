//go:build contract

package graph

import "testing"

func TestGraphContract(t *testing.T) {
	values := []string{"a", "b", "c"}
	nodes := []*Node[string]{{Key: 10, Value: &values[0]}, {Key: 20, Value: &values[1]}, {Key: 30, Value: &values[2]}}
	graph := graphFixture[string]{directed: true, nodes: nodes, edges: map[int][]edge[string]{10: {{to: nodes[1], weight: 7}, {to: nodes[2], weight: -3}}}}
	var _ Graph[string] = graph

	if !graph.Directed() || graph.NodeCount() != 3 {
		t.Fatal("Graph must expose direction and stable-key node count")
	}
	if node, ok := graph.NodeByKey(20); !ok || node != nodes[1] || node.Value != &values[1] {
		t.Fatal("NodeByKey must preserve key and underlying value pointer")
	}
	if _, ok := graph.NodeByKey(99); ok {
		t.Fatal("missing NodeByKey lookup must return ok=false")
	}

	var got []Edge[string]
	if complete, err := graph.Neighbors(10, func(to *Node[string], weight int64) bool {
		got = append(got, Edge[string]{From: nodes[0], To: to, Weight: weight})
		return true
	}); err != nil || !complete || len(got) != 2 || got[0] != (Edge[string]{From: nodes[0], To: nodes[1], Weight: 7}) || got[1] != (Edge[string]{From: nodes[0], To: nodes[2], Weight: -3}) {
		t.Fatal("Neighbors must visit each outgoing weighted node edge in order")
	}

	visits := 0
	if complete, err := graph.Neighbors(10, func(*Node[string], int64) bool {
		visits++
		return false
	}); err != nil || complete || visits != 1 {
		t.Fatalf("early stop = (complete=%t, visits=%d, err=%v), want (false, 1, nil)", complete, visits, err)
	}
}

func TestUndirectedEdgeGraphContract(t *testing.T) {
	values := []string{"a", "b"}
	nodes := []*Node[string]{{Key: 10, Value: &values[0]}, {Key: 20, Value: &values[1]}}
	graph := undirectedFixture[string]{nodes: nodes, edges: []Edge[string]{{From: nodes[0], To: nodes[1], Weight: 7}}}
	var _ UndirectedEdgeGraph[string] = graph

	var got []Edge[string]
	if complete, err := graph.Edges(func(edge Edge[string]) bool {
		got = append(got, edge)
		return true
	}); err != nil || !complete || len(got) != 1 || got[0] != (Edge[string]{From: nodes[0], To: nodes[1], Weight: 7}) {
		t.Fatal("Edges must report each logical undirected edge once in deterministic order")
	}
}

type edge[T any] struct {
	to     *Node[T]
	weight int64
}

type graphFixture[T any] struct {
	directed bool
	nodes    []*Node[T]
	edges    map[int][]edge[T]
}

type undirectedFixture[T any] struct {
	nodes []*Node[T]
	edges []Edge[T]
}

func (g graphFixture[T]) Directed() bool { return g.directed }
func (g graphFixture[T]) NodeCount() int { return len(g.nodes) }
func (g graphFixture[T]) NodeByKey(key int) (*Node[T], bool) {
	for _, node := range g.nodes {
		if node.Key == key {
			return node, true
		}
	}
	return nil, false
}
func (g graphFixture[T]) Neighbors(key int, visit func(*Node[T], int64) bool) (bool, error) {
	if _, ok := g.NodeByKey(key); !ok {
		return false, ErrInvalidKey
	}
	for _, edge := range g.edges[key] {
		if !visit(edge.to, edge.weight) {
			return false, nil
		}
	}
	return true, nil
}

func (g undirectedFixture[T]) Directed() bool { return false }
func (g undirectedFixture[T]) NodeCount() int { return len(g.nodes) }
func (g undirectedFixture[T]) NodeByKey(key int) (*Node[T], bool) {
	for _, node := range g.nodes {
		if node.Key == key {
			return node, true
		}
	}
	return nil, false
}
func (g undirectedFixture[T]) Neighbors(key int, visit func(*Node[T], int64) bool) (bool, error) {
	if _, ok := g.NodeByKey(key); !ok {
		return false, ErrInvalidKey
	}
	for _, edge := range g.edges {
		if edge.From.Key == key && !visit(edge.To, edge.Weight) {
			return false, nil
		}
		if edge.To.Key == key && !visit(edge.From, edge.Weight) {
			return false, nil
		}
	}
	return true, nil
}
func (g undirectedFixture[T]) Edges(visit func(Edge[T]) bool) (bool, error) {
	for _, edge := range g.edges {
		if !visit(edge) {
			return false, nil
		}
	}
	return true, nil
}
