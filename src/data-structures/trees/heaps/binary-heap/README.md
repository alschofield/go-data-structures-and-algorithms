# Binary Heap

## How It Works

A contiguous complete binary tree keeps the comparison maximum at its root.
Children of index `i` are at `2i+1` and `2i+2`; pushes sift up and pops sift down.

## Required API

```go
var ErrNilComparator error

type BinaryHeap[T any] struct

func NewBinaryHeap[T any](compare func(T, T) int) (*BinaryHeap[T], error)
func (bh *BinaryHeap[T]) Push(value T) bool
func (bh *BinaryHeap[T]) Pop() (T, bool)
func (bh *BinaryHeap[T]) Peek() (T, bool)
func (bh *BinaryHeap[T]) Len() int
func (bh *BinaryHeap[T]) IsEmpty() bool
```

## Contract

Construction rejects a nil comparator with `ErrNilComparator`. The comparator
must return a positive result when its first value has higher priority; this is
a max-heap. For example, `func(a, b int) int { return a - b }` pops largest
integers first. A comparator that reverses that relation produces a min-priority
heap. Equal priorities have no stable pop order.

`Push` always returns true. `Pop` and `Peek` return the zero value of `T` and
`ok=false` when empty without mutation; otherwise `Peek` preserves the heap and
`Pop` removes the root. The backing slice holds `graph.Node[T]` records with
stable insertion keys, but node positions change as values sift. Explicit child
links are unused.

## Complexity Targets

`Push` and `Pop` are O(log n); `Peek`, `Len`, and `IsEmpty` are O(1); space is
O(n).

## Verification

```sh
just contract data-structures/trees/heaps/binary-heap
```
