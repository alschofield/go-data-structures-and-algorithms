//go:build contract

package a_star

import "testing"

func TestAStar(t *testing.T) {
	graph := fixture{{{1, 4}, {2, 1}}, {{3, 1}}, {{1, 2}, {3, 5}}, nil}
	for _, test := range []struct {
		name      string
		heuristic Heuristic
		want      []int
		ok        bool
	}{{"zero", func(int) int64 { return 0 }, []int{0, 2, 1, 3}, true}, {"unreachable", func(int) int64 { return 0 }, nil, false}} {
		goal := 3
		if test.name == "unreachable" {
			goal = 4
		}
		got, ok := AStar(graph, 0, goal, test.heuristic)
		if ok != test.ok || !same(got, test.want) {
			t.Fatalf("%s: got (%v, %t), want (%v, %t)", test.name, got, ok, test.want, test.ok)
		}
	}
}

type edge struct {
	to     int
	weight int64
}
type fixture [][]edge

func (f fixture) Directed() bool   { return true }
func (f fixture) VertexCount() int { return len(f) }
func (f fixture) Neighbors(vertex int, visit func(int, int64) bool) bool {
	if vertex < 0 || vertex >= len(f) {
		return false
	}
	for _, edge := range f[vertex] {
		if !visit(edge.to, edge.weight) {
			break
		}
	}
	return true
}
func same(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
