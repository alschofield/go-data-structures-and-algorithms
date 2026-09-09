//go:build contract

package graph

import (
	"errors"
	"testing"
)

func TestGraphContract(t *testing.T) {
	var _ Graph = graphFixture{}
	graph := graphFixture{directed: true, edges: [][]edge{{{to: 1, weight: 7}, {to: 2, weight: -3}}, nil, nil}}
	if !graph.Directed() || graph.VertexCount() != 3 {
		t.Fatal("Graph must expose direction and dense vertex count")
	}

	var got []edge
	if complete, err := graph.Neighbors(0, func(to int, weight int64) bool {
		got = append(got, edge{to, weight})
		return true
	}); err != nil || !complete || len(got) != 2 || got[0] != (edge{1, 7}) || got[1] != (edge{2, -3}) {
		t.Fatal("Neighbors must visit each outgoing weighted edge in order")
	}
	if _, err := graph.Neighbors(3, func(int, int64) bool { return true }); !errors.Is(err, ErrInvalidVertex) {
		t.Fatalf("invalid vertex error = %v, want ErrInvalidVertex", err)
	}
	if _, err := graph.Neighbors(-1, func(int, int64) bool { return true }); !errors.Is(err, ErrInvalidVertex) {
		t.Fatalf("negative vertex error = %v, want ErrInvalidVertex", err)
	}

	visits := 0
	if complete, err := graph.Neighbors(0, func(int, int64) bool {
		visits++
		return false
	}); err != nil || complete || visits != 1 {
		t.Fatalf("early stop = (complete=%t, visits=%d, err=%v), want (false, 1, nil)", complete, visits, err)
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
func (g graphFixture) Neighbors(vertex int, visit func(int, int64) bool) (bool, error) {
	if vertex < 0 || vertex >= len(g.edges) {
		return false, ErrInvalidVertex
	}
	for _, edge := range g.edges[vertex] {
		if !visit(edge.to, edge.weight) {
			return false, nil
		}
	}
	return true, nil
}
