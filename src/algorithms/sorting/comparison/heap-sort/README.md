# Heap Sort

## How It Works

Bottom-up heapify builds an implicit max heap. Repeatedly swap its root with the
last unsorted item, then sift the new root down within the shrinking heap.

## Required API

```go
func HeapSort[T any](items []T, compare func(T, T) int) (bool, error)
```

`compare` orders values below zero, equal values at zero, and later values above
zero.

## Contract

Sort `items` ascending in place and return `(true, nil)`, including for nil,
empty, singleton, and duplicate slices. A nil comparator returns `(false,
ErrNilComparator)` before modifying `items`. Heap swaps can reorder equal
items, so stability is not guaranteed.

## Complexity Targets

O(n log n) time in every case and O(1) iterative auxiliary space.

## Verification

```sh
just contract algorithms/sorting/comparison/heap-sort
```
