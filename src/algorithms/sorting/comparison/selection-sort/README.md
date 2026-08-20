# Selection Sort

## How It Works
Each pass selects the remaining minimum and swaps it into the sorted prefix.

## Required API
`func SelectionSort[T any](items []T, compare func(T,T) int) bool`.

## Contract
Sort ascending in place using at most n-1 swaps. Classic selection sort is not stable. Empty/singleton input is a no-op; do not call `sort`.

## Complexity Targets
Best/average/worst O(n^2), O(1) space.
