# Quick Sort

## How It Works
Partition around a defended pivot then sort the lower and greater partitions.

## Required API
`func QuickSort[T any](items []T, compare func(T,T) int) bool`.

## Contract
Sort ascending in place, without stability. Use median-of-three or randomized pivots; handle duplicate/all-equal input without unbounded recursion and recurse on the smaller side. Do not call `sort`.

## Complexity Targets
Best/average O(n log n), worst O(n^2), expected O(log n) recursion space.
