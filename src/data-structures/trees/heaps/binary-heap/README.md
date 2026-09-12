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
Equal priorities have no stable ordering. Grow geometrically using a contiguous
slice of shared `*graph.Node[T]` records; implicit child positions use `2i+1`
and `2i+2`, so explicit Left and Right links remain unused. Do not use a library heap.

## Complexity Targets
Push/Pop O(log n), Peek/Len/IsEmpty O(1), bottom-up heapify O(n), O(n) contiguous space.

## Verification

```sh
make contract NAME=data-structures/trees/heaps/binary-heap
go test -tags=contract -run '^$' -bench=BinaryHeap -benchmem ./src/data-structures/trees/heaps/binary-heap
```

Benchmarks cover random and ascending full-heap construction, full removal from
a random heap, and a steady mixed push-pop workload at 256 and 1,024 nodes.
Full-build and full-removal workloads include their required setup; the mixed
workload prebuilds one heap before timing.
