//go:build contract

package dijkstra

import "testing"

func TestDijkstra(t *testing.T) {
	graph := fixture{{{1, 4}, {2, 1}}, {{3, 1}}, {{1, 2}, {3, 5}}, nil}
	result, ok := Dijkstra(graph, 0)
	if !ok {
		t.Fatal("valid nonnegative graph rejected")
	}
	for vertex, want := range map[int]int64{0: 0, 1: 3, 2: 1, 3: 4} {
		if got, present := result.Distance(vertex); !present || got != want {
			t.Fatalf("Distance(%d) = (%d, %t), want (%d, true)", vertex, got, present, want)
		}
	}
	if _, ok := Dijkstra(graph, 4); ok {
		t.Fatal("invalid source must fail")
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
