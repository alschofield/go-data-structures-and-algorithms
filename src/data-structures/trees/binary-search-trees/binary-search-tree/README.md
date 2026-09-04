# Binary Search Tree

## How It Works
Nodes place smaller values left and larger values right; in-order walking returns sorted values.

## Required API
Generic `type BinarySearchTree[T any]` with
`NewBinarySearchTree(compare func(T,T) int) (*BinarySearchTree[T], error)`,
`Insert(T) bool`, `Find(T) (T,bool)`, `Contains(T) bool`, `Remove(T) (T,bool)`,
`InOrder(func(T) bool) bool`, `Len`, and `IsEmpty`.

## Contract
`NewBinarySearchTree` returns `ErrNilComparator` for a nil comparator. Comparator
equality rejects duplicates and retains the first item; missing `Find` and
`Remove` results use `ok=false`, not an error. Remove supports leaf, one-child,
two-child, and root nodes. InOrder stops on false. Implement nodes, not an
ordered library container.

## Complexity Targets
Balanced core operations O(log n), unbalanced O(n), traversal O(n); O(n) nodes plus O(height) work space.
