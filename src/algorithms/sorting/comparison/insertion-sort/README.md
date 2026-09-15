# Insertion Sort

## How It Works

Grow a sorted prefix one item at a time, shifting only strictly greater prefix
items right until the next item can be inserted.

## Required API

```go
func InsertionSort[T any](items []T, compare func(T, T) int) (bool, error)
```

`compare` orders values below zero, equal values at zero, and later values above
zero.

## Contract

Sort `items` ascending in place and return `(true, nil)`, including for nil,
empty, and singleton slices. Do not move equal items past one another, so the
sort is stable. A nil comparator returns `(false, ErrNilComparator)` before
modifying `items`.

## Complexity Targets

Best O(n), average and worst O(n^2) time; O(1) auxiliary space.

## Verification

```sh
make contract NAME=algorithms/sorting/comparison/insertion-sort
```
