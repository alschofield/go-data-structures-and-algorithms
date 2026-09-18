# Binary Search Tree

## How It Works

Each node places comparator-smaller values left and larger values right.
In-order traversal therefore visits structural nodes in sorted order.

## Required API

```go
var ErrNilComparator error

type BinarySearchTree[T any] struct

func NewBinarySearchTree[T any](compare func(T, T) int) (*BinarySearchTree[T], error)
func (b *BinarySearchTree[T]) Insert(value T) (*graph.Node[T], bool)
func (b *BinarySearchTree[T]) Find(value T) *graph.Node[T]
func (b *BinarySearchTree[T]) Contains(value T) bool
func (b *BinarySearchTree[T]) Remove(value T) (*graph.Node[T], bool)
func (b *BinarySearchTree[T]) InOrder(visit func(*graph.Node[T]) bool) bool
func (b *BinarySearchTree[T]) Len() int
func (b *BinarySearchTree[T]) IsEmpty() bool
func (b *BinarySearchTree[T]) Directed() bool
func (b *BinarySearchTree[T]) NodeCount() int
func (b *BinarySearchTree[T]) NodeByKey(key int) (*graph.Node[T], bool)
func (b *BinarySearchTree[T]) Neighbors(key int, visit func(*graph.Node[T], int64) bool) (bool, error)
```

It implements `graph.Graph[T]`.

## Contract

Construction rejects a nil comparator with `ErrNilComparator`. Distinct inserts
return a new node with `added=true`; comparator-equal inserts retain the first
structural node, increment its `Occurrences`, and return it with `added=false`.
`Occurrences` is observational only: `Remove` structurally removes the matching
node regardless of its count. Missing `Find` returns nil; missing `Remove`
returns `nil, false`.

Keys are stable for a node lifetime and may have gaps after removal. `NodeByKey`
returns `ok=false` when absent. The graph view is directed: `Neighbors` returns
`graph.ErrInvalidKey` for an absent key, otherwise visits left then right child
edges at weight `1`, honoring early stop. Removal supports leaf, one-child,
two-child, and root cases. `InOrder` honors visitor early stop.

## Complexity Targets

Core operations are O(log n) on a balanced tree and O(n) in the worst case;
in-order traversal and `NodeByKey` are O(n), and `Neighbors` is O(n) plus
visited children; space is O(n).

## Verification

```sh
just contract data-structures/trees/binary-search-trees/binary-search-tree
```
