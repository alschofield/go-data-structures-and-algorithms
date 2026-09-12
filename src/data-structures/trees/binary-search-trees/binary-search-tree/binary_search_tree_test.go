//go:build contract

package binary_search_tree

import (
	"errors"
	"testing"

	graphcontract "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func TestBinarySearchTree(t *testing.T) {
	tree, err := NewBinarySearchTree(func(left, right int) int { return left - right })
	if err != nil {
		t.Fatalf("constructor error = %v", err)
	}
	var _ graphcontract.Graph[int] = tree

	nodes := make([]*graphcontract.Node[int], 0, 7)
	for _, value := range []int{4, 2, 6, 1, 3, 5, 7} {
		node, added := tree.Insert(value)
		if !added || node.Value == nil || *node.Value != value {
			t.Fatal("distinct insert must return a node pointing to its stored value")
		}
		nodes = append(nodes, node)
	}
	if node, added := tree.Insert(4); added || node != nodes[0] || node.Occurrences != 2 {
		t.Fatal("duplicate insert must retain the node and increment its occurrence metric")
	}
	if got, ok := tree.NodeByKey(nodes[0].Key); !ok || got != nodes[0] {
		t.Fatal("NodeByKey must retain stable node identity")
	}
	var neighbors []*graphcontract.Node[int]
	if complete, err := tree.Neighbors(nodes[0].Key, func(node *graphcontract.Node[int], weight int64) bool {
		neighbors = append(neighbors, node)
		return true
	}); err != nil || !complete || len(neighbors) != 2 || neighbors[0] != nodes[1] || neighbors[1] != nodes[2] {
		t.Fatal("root graph view must visit left and right child nodes")
	}

	if got, ok := tree.Remove(4); !ok || got.Value == nil || *got.Value != 4 || tree.Contains(4) {
		t.Fatal("Remove must structurally remove a root with two children regardless of Occurrences")
	}
	var got []int
	tree.InOrder(func(node *graphcontract.Node[int]) bool { got = append(got, *node.Value); return true })
	for index, want := range []int{1, 2, 3, 5, 6, 7} {
		if got[index] != want {
			t.Fatalf("in-order = %v", got)
		}
	}
	if _, err := NewBinarySearchTree[int](nil); !errors.Is(err, ErrNilComparator) {
		t.Fatalf("nil comparator error = %v, want ErrNilComparator", err)
	}
}
