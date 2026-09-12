package binary_search_tree

import (
	"errors"

	graph "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

var ErrNilComparator = errors.New("BST needs a valid compare function.")

type BinarySearchTree[T any] struct {
	size    int
	compare func(T, T) int
	root    *graph.Node[T]
}

func NewBinarySearchTree[T any](compare func(T, T) int) (*BinarySearchTree[T], error) {
	if compare == nil {
		return nil, ErrNilComparator
	}

	return &BinarySearchTree[T]{
		size:    0,
		compare: compare,
		root:    nil,
	}, nil
}

func (b *BinarySearchTree[T]) Insert(value T) (graph.Node[T], bool) {
	if b.root == nil {
		b.root = &graph.Node[T]{
			Value:      &value,
			Occurences: 1,
			Right:      nil,
			Left:       nil,
		}

		b.size++
		return *b.root, true
	}

	var candidate *graph.Node[T]
	var comparison int
	for candidate, comparison = b.root, b.compare(value, *candidate.Value); comparison < 0 && candidate.Left != nil || comparison > 0 && candidate.Right != nil || comparison == 0; comparison = b.compare(value, *candidate.Value) {
		if comparison < 0 {
			candidate = candidate.Left
		} else if comparison > 0 {
			candidate = candidate.Right
		}
	}

	if comparison != 0 {
		new_node := &graph.Node[T]{
			Value:      &value,
			Occurences: 1,
			Right:      nil,
			Left:       nil,
		}

		if comparison < 0 {
			candidate.Left = new_node
		} else if comparison > 0 {
			candidate.Right = new_node
		}

		b.size++
		return *new_node, true
	} else {
		candidate.Occurences++
		b.size++
		return *candidate, false
	}
}

func (b BinarySearchTree[T]) Find(value T) *graph.Node[T] {
	if b.size == 0 {
		return nil
	}

	var candidate *graph.Node[T] = nil
	var comparison int = 0
	for candidate, comparison = b.root, b.compare(value, *candidate.Value); comparison < 0 && candidate.Left != nil || comparison > 0 && candidate.Right != nil || comparison == 0; comparison = b.compare(value, *candidate.Value) {
		if comparison < 0 {
			candidate = candidate.Left
		} else if comparison > 0 {
			candidate = candidate.Right
		}
	}

	if comparison == 0 {
		return candidate
	} else {
		return nil
	}
}

func (b BinarySearchTree[T]) Contains(value T) bool {
	if b.size == 0 {
		return false
	}

	return b.Find(value) != nil
}

func (b BinarySearchTree[T]) Remove(value T) (*graph.Node[T], bool) {
	if b.size == 0 {
		return nil, false
	}

	var parent *graph.Node[T] = nil
	var candidate *graph.Node[T] = nil
	var comparison int = 0
	for candidate, comparison = b.root, b.compare(value, *candidate.Value); comparison < 0 && candidate.Left != nil || comparison > 0 && candidate.Right != nil || comparison == 0; comparison = b.compare(value, *candidate.Value) {
		if comparison < 0 {
			parent = candidate
			candidate = candidate.Left
		} else if comparison > 0 {
			parent = candidate
			candidate = candidate.Right
		}
	}

	if comparison == 0 {
		comparison = b.compare(value, *parent.Value)
		var replacement *graph.Node[T] = candidate
		var replacement_parent *graph.Node[T] = nil
		if replacement.Right != nil {
			for replacement.Right != nil {
				replacement_parent = replacement
				replacement = replacement.Right
			}

			replacement_parent.Right = nil
			replacement.Right = candidate.Right
		} else {
			replacement = replacement.Left
		}

		if comparison < 0 {
			parent.Left = replacement
		} else if comparison > 0 {
			parent.Right = replacement
		}

		return candidate, true
	} else {
		return nil, false
	}
}

func recurse[T any](node *graph.Node[T], visit func(*graph.Node[T]) bool) bool {
	if node.Left != nil {
		if !recurse(node.Left, visit) {
			return false
		}
	}

	if !visit(node) {
		return false
	}

	if node.Right != nil {
		if !recurse(node.Right, visit) {
			return false
		}
	}

	return true
}

func (b *BinarySearchTree[T]) InOrder(visit func(*graph.Node[T]) bool) bool {
	if b.size == 0 {
		return true
	}

	return recurse(b.root, visit)
}

func (b *BinarySearchTree[T]) Len() int {
	return b.size
}

func (b *BinarySearchTree[T]) IsEmpty() bool {
	return b.size == 0
}

func (b *BinarySearchTree[T]) Directed() bool {
	return true
}

func (b *BinarySearchTree[T]) NodeCount() int {
	return b.Len()
}

func (b *BinarySearchTree[T]) NodeByKey(key int) (*graph.Node[T], bool, error) {
	if key >= b.Len() {
		return nil, false, graph.ErrInvalidKey
	}

	var n int = 0
	var node *graph.Node[T] = nil
	b.InOrder(func(visitor *graph.Node[T]) bool {
		n++

		if n == key {
			node = visitor
			return false
		} else {
			return true
		}
	})

	return node, node == nil, nil
}

func (b *BinarySearchTree[T]) Neighbors(key int, visit func(*graph.Node[T], int64) bool) (bool, error) {
	if key >= b.Len() {
		return false, graph.ErrInvalidKey
	}

	node, status, err := b.NodeByKey(key)

	if !status {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if node.Left != nil {
		if !visit(node.Left, 1) {
			return true, nil
		}
	}

	if node.Right != nil {
		if !visit(node.Right, 1) {
			return true, nil
		}
	}

	return true, nil
}
