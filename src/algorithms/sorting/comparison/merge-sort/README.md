# Merge Sort

## How It Works
Recursively sort halves and merge them into a buffer, taking left values on ties.

## Required API
`func MergeSort[T any](items []T, compare func(T,T) int) error`.

## Contract
Sort ascending and stably. A nil comparator returns `ErrNilComparator` without
changing input. Handle empty, singleton, and uneven halves. Do not call `sort`.

## Complexity Targets
Best/average/worst O(n log n), O(n) buffer plus O(log n) recursion space.
