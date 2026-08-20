# Binary Heap

## How It Works
A contiguous complete tree uses `2i+1` and `2i+2` child indexes, holding the comparison extreme at the root.

## Required API
Generic `type BinaryHeap[T any]` with `NewBinaryHeap(compare func(T,T) int)`, `Push(T) bool`, `Pop() (T,bool)`, `Peek() (T,bool)`, `Len`, and `IsEmpty`.

## Contract
Push appends then sifts up; Pop moves the tail to root then sifts down. Equal priorities have no stable ordering. Grow geometrically; use no node allocation or library heap.

## Complexity Targets
Push/Pop O(log n), Peek/Len/IsEmpty O(1), bottom-up heapify O(n), O(n) contiguous space.
