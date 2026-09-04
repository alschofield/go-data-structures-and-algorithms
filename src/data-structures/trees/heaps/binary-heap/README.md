# Binary Heap

## How It Works
A contiguous complete tree uses `2i+1` and `2i+2` child indexes, holding the comparison extreme at the root.

## Required API
Generic `type BinaryHeap[T any]` with
`NewBinaryHeap(compare func(T,T) int) (*BinaryHeap[T], error)`, `Push(T) bool`,
`Pop() (T,bool)`, `Peek() (T,bool)`, `Len`, and `IsEmpty`.

## Contract
`NewBinaryHeap` returns `ErrNilComparator` for a nil comparator. Push appends
then sifts up; Pop and Peek on an empty heap return `ok=false`, not an error.
Equal priorities have no stable ordering. Grow geometrically; use no node
allocation or library heap.

## Complexity Targets
Push/Pop O(log n), Peek/Len/IsEmpty O(1), bottom-up heapify O(n), O(n) contiguous space.
