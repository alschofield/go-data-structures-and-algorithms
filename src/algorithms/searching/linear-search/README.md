# Linear Search

## How It Works
Scan arbitrary input from front to back until equality matches the key.

## Required API
`func LinearSearch[T any](items []T, key T, compare func(T,T) int) (int, bool)`.

## Contract
Works unsorted, returns the first duplicate, reports absence distinctly, and never changes input. Do not use a standard-library search routine.

## Complexity Targets
Best O(1), average/worst O(n), O(1) space.
