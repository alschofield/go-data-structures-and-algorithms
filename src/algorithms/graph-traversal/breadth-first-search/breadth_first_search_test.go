//go:build contract

package breadth_first_search

import (
	"errors"
	"testing"

	graphcontract "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func TestBreadthFirstSearch(t *testing.T) {
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
		complete   bool
		err        error
	}{{nodes[0].Key, []int{10, 20, 30, 40}, true, nil}, {nodes[4].Key, []int{50}, true, nil}, {99, nil, false, graphcontract.ErrInvalidKey}} {
		got, complete, err := BreadthFirstSearch(graph, test.source_key, func(*graphcontract.Node[string]) bool { return true })
		if !errors.Is(err, test.err) || complete != test.complete || !same_keys(got, test.want) {
			t.Fatalf("source %d: got (%v, %t, %v), want keys (%v, %t, %v)", test.source_key, got, complete, err, test.want, test.complete, test.err)
		}
	}

	stopped, complete, err := BreadthFirstSearch(graph, nodes[0].Key, func(node *graphcontract.Node[string]) bool { return node.Key != nodes[2].Key })
	if err != nil || !complete || !same_keys(stopped, []int{10, 20, 30}) {
		t.Fatalf("early stop = (%v, %t, %v), want ([10 20 30], true, nil)", stopped, complete, err)
	}
	if _, complete, err := BreadthFirstSearch[string](nil, nodes[0].Key, func(*graphcontract.Node[string]) bool { return true }); complete || !errors.Is(err, ErrNilGraph) {
		t.Fatalf("nil graph = (complete=%t, err=%v), want (false, ErrNilGraph)", complete, err)
	}
	if _, complete, err := BreadthFirstSearch(graph, nodes[0].Key, nil); complete || !errors.Is(err, ErrNilVisit) {
		t.Fatalf("nil visitor = (complete=%t, err=%v), want (false, ErrNilVisit)", complete, err)
	}
}

type fixture struct {
	nodes []*graphcontract.Node[string]
	edges map[int][]*graphcontract.Node[string]
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
