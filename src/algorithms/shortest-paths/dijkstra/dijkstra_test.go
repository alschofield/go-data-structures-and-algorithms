//go:build contract

package dijkstra

import (
	"errors"
	"testing"

	graphcontract "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func TestDijkstra(t *testing.T) {
	values := []string{"a", "b", "c", "d"}
	nodes := make([]*graphcontract.Node[string], len(values))
	for index := range values {
		nodes[index] = &graphcontract.Node[string]{Key: (index + 1) * 10, Value: &values[index]}
	}
	graph := fixture{nodes: nodes, edges: map[int][]edge{nodes[0].Key: {{to: nodes[1], weight: 4}, {to: nodes[2], weight: 1}}, nodes[1].Key: {{to: nodes[3], weight: 1}}, nodes[2].Key: {{to: nodes[1], weight: 2}, {to: nodes[3], weight: 5}}}}
	var _ graphcontract.Graph[string] = graph

	result, err := Dijkstra(graph, nodes[0].Key)
	if err != nil {
		t.Fatalf("valid nonnegative graph rejected: %v", err)
	}
	for key, want := range map[int]int64{10: 0, 20: 3, 30: 1, 40: 4} {
		if got, present := result.Distance(key); !present || got != want {
			t.Fatalf("Distance(%d) = (%d, %t), want (%d, true)", key, got, present, want)
		}
	}
	if _, err := Dijkstra(graph, 99); !errors.Is(err, graphcontract.ErrInvalidKey) {
		t.Fatalf("invalid source error = %v, want ErrInvalidKey", err)
	}
}

type edge struct {
	to     *graphcontract.Node[string]
	weight int64
}
type fixture struct {
	nodes []*graphcontract.Node[string]
	edges map[int][]edge
}

func (f fixture) Directed() bool { return true }
func (f fixture) NodeCount() int { return len(f.nodes) }
func (f fixture) NodeByKey(key int) (*graphcontract.Node[string], bool, error) {
	for _, node := range f.nodes {
		if node.Key == key {
			return node, true, nil
		}
	}
	return nil, false, graphcontract.ErrInvalidKey
}
func (f fixture) Neighbors(key int, visit func(*graphcontract.Node[string], int64) bool) (bool, error) {
	if _, ok, err := f.NodeByKey(key); err != nil || !ok {
		return false, graphcontract.ErrInvalidKey
	}
	for _, edge := range f.edges[key] {
		if !visit(edge.to, edge.weight) {
			return false, nil
		}
	}
	return true, nil
}
