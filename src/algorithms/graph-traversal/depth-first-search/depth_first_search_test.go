//go:build contract

package depth_first_search

import "testing"

func TestDepthFirstSearch(t *testing.T) {
	graph := fixture{{1, 2}, {3}, {3}, nil, nil}
	for _, test := range []struct {
		source int
		want   []int
		ok     bool
	}{{0, []int{0, 1, 3, 2}, true}, {4, []int{4}, true}, {5, nil, false}} {
		got, ok := DepthFirstSearch(graph, test.source)
		if ok != test.ok || !same(got, test.want) {
			t.Fatalf("source %d: got (%v, %t), want (%v, %t)", test.source, got, ok, test.want, test.ok)
		}
	}
}

type fixture [][]int

func (f fixture) Directed() bool   { return true }
func (f fixture) VertexCount() int { return len(f) }
func (f fixture) Neighbors(vertex int, visit func(int, int64) bool) bool {
	if vertex < 0 || vertex >= len(f) {
		return false
	}
	for _, to := range f[vertex] {
		if !visit(to, 1) {
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
