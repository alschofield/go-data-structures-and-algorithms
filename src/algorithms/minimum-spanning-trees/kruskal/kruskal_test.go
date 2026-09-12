//go:build contract

package kruskal

import (
	"errors"
	"reflect"
	"testing"

	graphcontract "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func TestKruskalMinimumSpanningForest(t *testing.T) {
	for _, test := range []struct {
		name   string
		graph  graphFixture
		want   []graphcontract.Edge[string]
		weight int64
		err    error
	}{
		{"connected", undirectedFixture([]string{"a", "b", "c", "d"}, []edge{{0, 1, 4}, {0, 2, 1}, {1, 2, 2}, {1, 3, 5}, {2, 3, 3}}), []graphcontract.Edge[string]{{From: node("a", 10), To: node("c", 30), Weight: 1}, {From: node("b", 20), To: node("c", 30), Weight: 2}, {From: node("c", 30), To: node("d", 40), Weight: 3}}, 6, nil},
		{"forest", undirectedFixture([]string{"a", "b", "c", "d"}, []edge{{0, 1, -2}, {2, 3, 4}}), []graphcontract.Edge[string]{{From: node("a", 10), To: node("b", 20), Weight: -2}, {From: node("c", 30), To: node("d", 40), Weight: 4}}, 2, nil},
		{"directed", graphFixture{directed: true, nodes: nodes("a", "b")}, nil, 0, ErrDirectedGraph},
	} {
		t.Run(test.name, func(t *testing.T) {
			var _ graphcontract.UndirectedEdgeGraph[string] = test.graph
			got, weight, err := KruskalMinimumSpanningForest(test.graph)
			if !errors.Is(err, test.err) || weight != test.weight || !reflect.DeepEqual(got, test.want) {
				t.Fatalf("KruskalMinimumSpanningForest() = (%v, %d, %v), want (%v, %d, %v)", got, weight, err, test.want, test.weight, test.err)
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
	nodes    []*graphcontract.Node[string]
	logical  []graphcontract.Edge[string]
}

func nodes(values ...string) []*graphcontract.Node[string] {
	result := make([]*graphcontract.Node[string], len(values))
	for index := range values {
		result[index] = node(values[index], (index+1)*10)
	}
	return result
}
func node(value string, key int) *graphcontract.Node[string] {
	return &graphcontract.Node[string]{Key: key, Value: &value}
}
func undirectedFixture(values []string, input []edge) graphFixture {
	result := graphFixture{nodes: nodes(values...), logical: make([]graphcontract.Edge[string], 0, len(input))}
	for _, edge := range input {
		result.logical = append(result.logical, graphcontract.Edge[string]{From: result.nodes[edge.from], To: result.nodes[edge.to], Weight: edge.weight})
	}
	return result
}
func (g graphFixture) Directed() bool { return g.directed }
func (g graphFixture) NodeCount() int { return len(g.nodes) }
func (g graphFixture) NodeByKey(key int) (*graphcontract.Node[string], bool) {
	for _, node := range g.nodes {
		if node.Key == key {
			return node, true
		}
	}
	return nil, false
}
func (g graphFixture) Neighbors(key int, visit func(*graphcontract.Node[string], int64) bool) (bool, error) {
	if _, ok := g.NodeByKey(key); !ok {
		return false, graphcontract.ErrInvalidKey
	}
	for _, edge := range g.logical {
		if edge.From.Key == key && !visit(edge.To, edge.Weight) {
			return false, nil
		}
		if !g.directed && edge.To.Key == key && !visit(edge.From, edge.Weight) {
			return false, nil
		}
	}
	return true, nil
}
func (g graphFixture) Edges(visit func(graphcontract.Edge[string]) bool) (bool, error) {
	for _, edge := range g.logical {
		if !visit(edge) {
			return false, nil
		}
	}
	return true, nil
}
