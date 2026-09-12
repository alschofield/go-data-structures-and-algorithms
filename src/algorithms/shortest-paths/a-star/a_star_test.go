//go:build contract

package a_star

import (
	"testing"

	graphcontract "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func TestAStar(t *testing.T) {
	values := []string{"a", "b", "c", "d", "e"}
	nodes := make([]*graphcontract.Node[string], len(values))
	for index := range values {
		nodes[index] = &graphcontract.Node[string]{Key: (index + 1) * 10, Value: &values[index]}
	}
	graph := fixture{nodes: nodes, edges: map[int][]edge{nodes[0].Key: {{to: nodes[1], weight: 4}, {to: nodes[2], weight: 1}}, nodes[1].Key: {{to: nodes[3], weight: 1}}, nodes[2].Key: {{to: nodes[1], weight: 2}, {to: nodes[3], weight: 5}}}}
	var _ graphcontract.Graph[string] = graph
	for _, test := range []struct {
		name      string
		heuristic Heuristic
		goal_key  int
		want      []int
	}{
		{"zero", func(int) int64 { return 0 }, nodes[3].Key, []int{10, 30, 20, 40}},
		{"unreachable", func(int) int64 { return 0 }, nodes[4].Key, nil},
	} {
		got, err := AStar(graph, nodes[0].Key, test.goal_key, test.heuristic)
		if err != nil || !same_keys(got, test.want) {
			t.Fatalf("%s: got (%v, %v), want keys (%v, nil)", test.name, got, err, test.want)
		}
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
func (f fixture) NodeByKey(key int) (*graphcontract.Node[string], bool) {
	for _, node := range f.nodes {
		if node.Key == key {
			return node, true
		}
	}
	return nil, false
}
func (f fixture) Neighbors(key int, visit func(*graphcontract.Node[string], int64) bool) (bool, error) {
	if _, ok := f.NodeByKey(key); !ok {
		return false, graphcontract.ErrInvalidKey
	}
	for _, edge := range f.edges[key] {
		if !visit(edge.to, edge.weight) {
			return false, nil
		}
	}
	return true, nil
}
func same_keys(got []*graphcontract.Node[string], want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for index := range got {
		if got[index].Key != want[index] {
			return false
		}
	}
	return true
}
