package binary_search_tree

import (
	"errors"
	"math"

	graph "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

// ErrNilComparator reports that a BST cannot order values without a comparator.
var ErrNilComparator = errors.New("BST needs a valid compare function.")

// BinarySearchTree stores distinct ordered values with stable graph-node keys.
type BinarySearchTree[T any] struct {
	size     int
	next_key int
	compare  func(T, T) int
	root     *graph.Node[T]
}

// NewBinarySearchTree creates an empty tree after validating its comparator.
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

// Insert adds a structural node or records a duplicate occurrence on an existing node.
func (b *BinarySearchTree[T]) Insert(value T) (*graph.Node[T], bool) {
	if b.next_key >= math.MaxInt {
		return nil, false
	}

	// Stable keys are never reused after removal.
	if b.root == nil {
		b.root = &graph.Node[T]{
			Key:         b.next_key,
			Value:       &value,
			Occurrences: 1,
			Right:       nil,
			Left:        nil,
		}

		b.next_key++
		b.size++
		return b.root, true
	}

	// Walk until a matching value or the missing child that receives a new node.
	candidate, comparison := b.root, 0
	for candidate, comparison = b.root, b.compare(value, *candidate.Value); ((comparison < 0 && candidate.Left != nil) || (comparison > 0 && candidate.Right != nil)) && comparison != 0; comparison = b.compare(value, *candidate.Value) {
		if comparison < 0 {
			candidate = candidate.Left
		} else if comparison > 0 {
			candidate = candidate.Right
		}
	}

	if comparison != 0 {
		new_node := &graph.Node[T]{
			Key:         b.next_key,
			Value:       &value,
			Occurrences: 1,
			Right:       nil,
			Left:        nil,
		}

		if comparison < 0 {
			candidate.Left = new_node
		} else if comparison > 0 {
			candidate.Right = new_node
		}

		b.next_key++
		b.size++
		return new_node, true
	} else {
		// Duplicates retain their structural node and only update its metric.
		candidate.Occurrences++
		return candidate, false
	}
}

// Find returns the structural node whose value compares equal to value.
func (b *BinarySearchTree[T]) Find(value T) *graph.Node[T] {
	if b.size == 0 {
		return nil
	}

	candidate, comparison := b.root, 0
	for candidate, comparison = b.root, b.compare(value, *candidate.Value); ((comparison < 0 && candidate.Left != nil) || (comparison > 0 && candidate.Right != nil)) && comparison != 0; comparison = b.compare(value, *candidate.Value) {
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

func (b *BinarySearchTree[T]) Contains(value T) bool {
	if b.size == 0 {
		return false
	}

	return b.Find(value) != nil
}

// Remove unlinks a matching structural node regardless of its occurrence metric.
func (b *BinarySearchTree[T]) Remove(value T) (*graph.Node[T], bool) {
	if b.size == 0 {
		return nil, false
	}

	var parent *graph.Node[T] = nil
	candidate, comparison := b.root, 0
	for candidate, comparison = b.root, b.compare(value, *candidate.Value); ((comparison < 0 && candidate.Left != nil) || (comparison > 0 && candidate.Right != nil)) && comparison != 0; comparison = b.compare(value, *candidate.Value) {
		if comparison < 0 {
			parent = candidate
			candidate = candidate.Left
		} else if comparison > 0 {
			parent = candidate
			candidate = candidate.Right
		}
	}

	if comparison == 0 {
		// Replace a two-child node with the smallest node in its right subtree.
		replacement := candidate
		if replacement.Right != nil {
			replacement_parent := replacement
			replacement = replacement.Right
			for replacement.Left != nil {
				replacement, replacement_parent = replacement.Left, replacement
			}

			if replacement != candidate.Right {
				replacement_parent.Left = replacement.Right
				replacement.Right = candidate.Right
			}

			replacement.Left = candidate.Left
		} else {
			replacement = replacement.Left
		}

		if parent != nil {
			comparison = b.compare(value, *parent.Value)
			if comparison < 0 {
				parent.Left = replacement
			} else if comparison > 0 {
				parent.Right = replacement
			}
		} else {
			b.root = replacement
		}

		b.size--
		return candidate, true
	} else {
		return nil, false
	}
}

// recurse visits a subtree in comparator order and honors early visitor stop.
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

// InOrder visits nodes from smallest to largest value.
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

// NodeByKey finds a stable graph key through the current tree structure.
func (b *BinarySearchTree[T]) NodeByKey(key int) (*graph.Node[T], bool, error) {
	var node *graph.Node[T] = nil
	b.InOrder(func(visitor *graph.Node[T]) bool {
		if key == visitor.Key {
			node = visitor
			return false
		} else {
			return true
		}
	})

	if node == nil {
		return nil, false, graph.ErrInvalidKey
	} else {
		return node, true, nil
	}
}

// Neighbors exposes left and right child links as directed unit-weight edges.
func (b *BinarySearchTree[T]) Neighbors(key int, visit func(*graph.Node[T], int64) bool) (bool, error) {
	node, status, err := b.NodeByKey(key)

	if err != nil {
		return false, err
	}

	if !status {
		return false, nil
	}

	if node.Left != nil {
		if !visit(node.Left, 1) {
			return false, nil
		}
	}

	if node.Right != nil {
		if !visit(node.Right, 1) {
			return false, nil
		}
	}

	return true, nil
}
