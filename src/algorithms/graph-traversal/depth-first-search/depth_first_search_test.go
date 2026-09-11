//go:build contract

package depth_first_search

import (
	"errors"
	"testing"

	graphcontract "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func TestDepthFirstSearch(t *testing.T) {
	values := []string{"a", "b", "c", "d", "e"}
	nodes := make([]*graphcontract.Node[string], len(values))
	for index := range values {
		nodes[index] = &graphcontract.Node[string]{Key: (index + 1) * 10, Value: &values[index]}
	}
	graph := fixture{nodes: nodes, edges: map[int][]*graphcontract.Node[string]{nodes[0].Key: {nodes[1], nodes[2]}, nodes[1].Key: {nodes[3]}, nodes[2].Key: {nodes[3]}}}
	var _ graphcontract.Graph[string] = graph
	for _, test := range []struct {
		source_key int
		want       []int
		err        error
	}{{nodes[0].Key, []int{10, 20, 40, 30}, nil}, {nodes[4].Key, []int{50}, nil}, {99, nil, graphcontract.ErrInvalidKey}} {
		got, err := DepthFirstSearch(graph, test.source_key)
		if !errors.Is(err, test.err) || !same_keys(got, test.want) {
			t.Fatalf("source %d: got (%v, %v), want keys (%v, %v)", test.source_key, got, err, test.want, test.err)
		}
	}
}

type fixture struct {
	nodes []*graphcontract.Node[string]
	edges map[int][]*graphcontract.Node[string]
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
	for _, node := range f.edges[key] {
		if !visit(node, 1) {
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
