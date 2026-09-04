//go:build contract

package graph

import "testing"

func TestGraphContract(t *testing.T) {
	var _ Graph = graphFixture{}
	graph := graphFixture{directed: true, edges: [][]edge{{{to: 1, weight: 7}, {to: 2, weight: -3}}, nil, nil}}
	if !graph.Directed() || graph.VertexCount() != 3 {
		t.Fatal("Graph must expose direction and dense vertex count")
	}

	var got []edge
	if !graph.Neighbors(0, func(to int, weight int64) bool {
		got = append(got, edge{to, weight})
		return true
	}) || len(got) != 2 || got[0] != (edge{1, 7}) || got[1] != (edge{2, -3}) {
		t.Fatal("Neighbors must visit each outgoing weighted edge in order")
	}
	if graph.Neighbors(3, func(int, int64) bool { return true }) {
		t.Fatal("Neighbors must reject an out-of-range vertex")
	}
}

type edge struct {
	to     int
	weight int64
}
type graphFixture struct {
	directed bool
	edges    [][]edge
}

func (g graphFixture) Directed() bool   { return g.directed }
func (g graphFixture) VertexCount() int { return len(g.edges) }
func (g graphFixture) Neighbors(vertex int, visit func(int, int64) bool) bool {
	if vertex < 0 || vertex >= len(g.edges) {
		return false
	}
	for _, edge := range g.edges[vertex] {
		if !visit(edge.to, edge.weight) {
			break
		}
	}
	return true
}
