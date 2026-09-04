# Bubble Sort

## How It Works
Adjacent swaps move the largest unsorted item tailward; a no-swap pass exits early.

## Required API
`func BubbleSort[T any](items []T, compare func(T,T) int) error`.

## Contract
Sort ascending in place, swap only strictly out-of-order neighbors to remain
stable, and implement early exit. Empty/singleton input is a no-op. A nil
comparator returns `ErrNilComparator` without changing input. Do not call `sort`.

## Complexity Targets
Best O(n), average/worst O(n^2), O(1) space.
