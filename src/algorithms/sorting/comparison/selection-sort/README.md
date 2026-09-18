# Selection Sort

## How It Works

For each prefix position, find the minimum item in the remaining suffix and
swap it into place. After pass `i`, positions `[0, i]` contain the smallest
items in ascending order.

## Required API

```go
func SelectionSort[T any](items []T, compare func(T, T) int) (bool, error)
```

`compare` orders values below zero, equal values at zero, and later values above
zero.

## Contract

Sort `items` ascending in place and return `(true, nil)`, including for nil,
empty, and singleton slices. A nil comparator returns `(false,
ErrNilComparator)` before modifying `items`. The final minimum swap can reorder
equal values, so stability is not guaranteed.

## Complexity Targets

O(n^2) time in every case and O(1) auxiliary space.

## Verification

```sh
just contract algorithms/sorting/comparison/selection-sort
```
