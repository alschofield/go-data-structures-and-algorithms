//go:build contract

package breadth_first_search

import (
	"errors"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	graph := fixture{{1, 2}, {3}, {3}, nil, nil}
	for _, test := range []struct {
		source int
		want   []int
		err    error
	}{{0, []int{0, 1, 2, 3}, nil}, {4, []int{4}, nil}, {5, nil, ErrInvalidVertex}} {
		got, err := BreadthFirstSearch(graph, test.source)
		if !errors.Is(err, test.err) || !same(got, test.want) {
			t.Fatalf("source %d: got (%v, %v), want (%v, %v)", test.source, got, err, test.want, test.err)
		}
	}
}

type fixture [][]int

func (f fixture) Directed() bool   { return true }
func (f fixture) VertexCount() int { return len(f) }
func (f fixture) Neighbors(vertex int, visit func(int, int64) bool) (bool, error) {
	if vertex < 0 || vertex >= len(f) {
		return false, ErrInvalidVertex
	}
	for _, to := range f[vertex] {
		if !visit(to, 1) {
			break
		}
	}
	return true, nil
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
