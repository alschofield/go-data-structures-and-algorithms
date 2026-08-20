# Binary Search

## How It Works
Compare sorted input's middle element and discard one candidate half per step.

## Required API
`func BinarySearch[T any](items []T, key T, compare func(T,T) int) (int, bool)`.

## Contract
Assume ascending input but never sort/validate it. Use overflow-safe midpoint arithmetic. Absence and empty input return `ok=false`; any duplicate match is valid. Never modify input or use library search.

## Complexity Targets
Best O(1), average/worst O(log n), O(1) iterative space.
