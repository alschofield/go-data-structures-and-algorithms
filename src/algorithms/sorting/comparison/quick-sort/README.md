# Quick Sort

## How It Works

Choose a median-of-three pivot, partition the slice into less-than, equal-to,
and greater-than regions, then recursively sort only the outer regions. Keeping
the equal region out of recursion handles duplicate-heavy input efficiently.

## Required API

```go
func QuickSort[T any](items []T, compare func(T, T) int) (bool, error)
```

`compare` orders values below zero, equal values at zero, and later values above
zero.

## Contract

Sort `items` ascending in place and return `(true, nil)`, including for nil,
empty, singleton, duplicate, and all-equal slices. A nil comparator returns
`(false, ErrNilComparator)` before modifying `items`. Partition swaps can
reorder equal items, so stability is not guaranteed.

## Complexity Targets

Best and average O(n log n), worst O(n^2) time. The current direct recursion
uses O(log n) stack space for balanced partitions and O(n) in the worst case.

## Verification

```sh
make contract NAME=algorithms/sorting/comparison/quick-sort
```
