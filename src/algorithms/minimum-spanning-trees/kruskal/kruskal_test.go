//go:build contract

package kruskal

import (
	"reflect"
	"testing"
)

func TestKruskalMinimumSpanningForest(t *testing.T) {
	for _, test := range []struct {
		name   string
		graph  graphFixture
		want   []Edge
		weight int64
		ok     bool
	}{
		{"connected", undirectedFixture(4, []edge{{0, 1, 4}, {0, 2, 1}, {1, 2, 2}, {1, 3, 5}, {2, 3, 3}}), []Edge{{0, 2, 1}, {1, 2, 2}, {2, 3, 3}}, 6, true},
		{"forest", undirectedFixture(4, []edge{{0, 1, -2}, {2, 3, 4}}), []Edge{{0, 1, -2}, {2, 3, 4}}, 2, true},
		{"directed", graphFixture{directed: true, vertices: 2}, nil, 0, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, weight, ok := KruskalMinimumSpanningForest(test.graph)
			if ok != test.ok || weight != test.weight || !reflect.DeepEqual(got, test.want) {
				t.Fatalf("KruskalMinimumSpanningForest() = (%v, %d, %t), want (%v, %d, %t)", got, weight, ok, test.want, test.weight, test.ok)
			}
		})
	}
}

type edge struct {
	from, to int
	weight   int64
}
type graphFixture struct {
	directed bool
	vertices int
	edges    [][]edge
}

func undirectedFixture(vertices int, edges []edge) graphFixture {
	g := graphFixture{vertices: vertices, edges: make([][]edge, vertices)}
	for _, edge := range edges {
		g.edges[edge.from] = append(g.edges[edge.from], edge)
		g.edges[edge.to] = append(g.edges[edge.to], edge{edge.to, edge.from, edge.weight})
	}
	return g
}
func (g graphFixture) Directed() bool   { return g.directed }
func (g graphFixture) VertexCount() int { return g.vertices }
func (g graphFixture) Neighbors(vertex int, visit func(int, int64) bool) bool {
	if vertex < 0 || vertex >= g.vertices {
		return false
	}
	for _, edge := range g.edges[vertex] {
		if !visit(edge.to, edge.weight) {
			break
		}
	}
	return true
}
