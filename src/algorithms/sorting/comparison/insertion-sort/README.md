# Insertion Sort

## How It Works
Grow a sorted prefix, shifting strictly greater items right before each insertion.

## Required API
`func InsertionSort[T any](items []T, compare func(T,T) int) (bool, error)`.

## Contract
Sort ascending in place and remain stable by inserting after equals.
Empty/singleton input is a successful no-op. Success returns `(true, nil)`. A
nil comparator returns `(false, ErrNilComparator)` without changing input. Do
not call `sort`.

## Complexity Targets
Best O(n), average/worst O(n^2), O(1) space.
