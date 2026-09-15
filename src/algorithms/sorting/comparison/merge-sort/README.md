# Merge Sort

## How It Works

Recursively sort two halves, then merge them into a buffer. On a tie, take the
left value first so equal values retain their input order.

## Required API

```go
func MergeSort[T any](items []T, compare func(T, T) int) (bool, error)
```

`compare` orders values below zero, equal values at zero, and later values above
zero.

## Contract

Sort `items` ascending in place and return `(true, nil)`, including for nil,
empty, singleton, and unevenly divided slices. The merge must take the left
half on equality, making the result stable. A nil comparator returns `(false,
ErrNilComparator)` before modifying `items`.

## Complexity Targets

O(n log n) time and O(n) merge-buffer space, plus O(log n) recursion space.

## Verification

```sh
make contract NAME=algorithms/sorting/comparison/merge-sort
```
