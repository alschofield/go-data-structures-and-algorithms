# Bubble Sort

## How It Works

Each pass swaps adjacent, strictly out-of-order items, moving the largest
remaining item into the unsorted tail. A pass with no swaps proves the remaining
prefix is already sorted.

## Required API

```go
func BubbleSort[T any](items []T, compare func(T, T) int) (bool, error)
```

`compare` orders values below zero, equal values at zero, and later values above
zero.

## Contract

Sort `items` ascending in place and return `(true, nil)`, including for nil,
empty, and singleton slices. Swap only when `compare(left, right) > 0`; equal
items retain their input order, so the sort is stable. A nil comparator returns
`(false, ErrNilComparator)` before modifying `items`.

## Complexity Targets

Best O(n), average and worst O(n^2) time; O(1) auxiliary space.

## Verification

```sh
just contract algorithms/sorting/comparison/bubble-sort
```
