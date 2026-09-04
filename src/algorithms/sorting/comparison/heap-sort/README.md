# Heap Sort

## How It Works
Bottom-up heapify creates a max heap, then root-to-tail swaps and sift-down sort the shrinking prefix.

## Required API
`func HeapSort[T any](items []T, compare func(T,T) int) error`.

## Contract
Sort ascending in place, build via O(n) bottom-up heapify with implicit indexes,
and do not claim stability. Empty/singleton input is a no-op. A nil comparator
returns `ErrNilComparator` without changing input; do not call `sort`.

## Complexity Targets
Best/average/worst O(n log n), O(1) iterative space.
