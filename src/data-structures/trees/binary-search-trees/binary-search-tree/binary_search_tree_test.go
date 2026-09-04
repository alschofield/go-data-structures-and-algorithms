//go:build contract

package binary_search_tree

import "testing"

func TestBinarySearchTree(t *testing.T) {
	tree := NewBinarySearchTree(func(a, b int) int { return a - b })
	for _, value := range []int{4, 2, 6, 1, 3, 5, 7} {
		if !tree.Insert(value) {
			t.Fatal("distinct insert failed")
		}
	}
	if tree.Insert(4) {
		t.Fatal("duplicate insert must fail")
	}
	if got, ok := tree.Remove(4); !ok || got != 4 || tree.Contains(4) {
		t.Fatal("root with two children must be removable")
	}
	var got []int
	tree.InOrder(func(value int) bool { got = append(got, value); return true })
	for index, want := range []int{1, 2, 3, 5, 6, 7} {
		if got[index] != want {
			t.Fatalf("in-order = %v", got)
		}
	}
}
