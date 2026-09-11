# Binary Search Tree

## How It Works

Nodes place smaller values left and larger values right; in-order walking returns sorted values. Each inserted node also has a stable graph key and exposes a pointer to its stored value.

## Required API

Generic `type BinarySearchTree[T any]` with
`NewBinarySearchTree(compare func(T, T) int) (*BinarySearchTree[T], error)`,
`Insert(value T) (*graph.Node[T], bool)`, `Find(value T) (*graph.Node[T], bool)`,
`Contains(value T) bool`, `Remove(value T) (*graph.Node[T], bool)`,
`InOrder(func(*graph.Node[T]) bool) bool`, `Len() int`, and `IsEmpty() bool`.

It directly implements `graph.Graph[T]`: `Directed() bool`, `NodeCount() int`,
`NodeByKey(key int) (*graph.Node[T], bool, error)`, and `Neighbors(key int,
func(*graph.Node[T], int64) bool) (bool, error)`.

## Contract

`NewBinarySearchTree` returns `ErrNilComparator` for a nil comparator.
Comparator equality rejects duplicates and retains the first node; missing `Find`
and `Remove` results use `ok=false`, not an error. Graph keys are stable for a
node's lifetime, may have gaps after removal, and `Node.Value` points to the
value stored in that node. The graph view is directed and reports left then
right child edges at weight `1`. Remove supports leaf, one-child, two-child,
and root nodes. InOrder stops on false. Implement nodes, not an ordered library
container.

## Complexity Targets

Balanced core operations O(log n), unbalanced O(n), traversal O(n); O(n) nodes plus O(height) work space. Graph key lookup may use the tree's own node structure; do not bolt on a separate external lookup structure.
