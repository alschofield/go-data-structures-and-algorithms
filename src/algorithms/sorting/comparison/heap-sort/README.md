# Heap Sort

## How It Works
Bottom-up heapify creates a max heap, then root-to-tail swaps and sift-down sort the shrinking prefix. Equal values may reorder; HeapSort is not stable.

## Required API
`func HeapSort[T any](items []T, compare func(T,T) int) (bool, error)`.

## Contract
Sort ascending in place, build via O(n) bottom-up heapify with implicit indexes,
and do not claim stability. Empty/singleton input is a no-op. A nil comparator
returns `ErrNilComparator` without changing input. Successful calls return
`(true, nil)` and invalid calls return `(false, ErrNilComparator)`; do not call
`sort`.

## Complexity Targets
Best/average/worst O(n log n), O(1) iterative space.

## Verification

```sh
make contract NAME=algorithms/sorting/comparison/heap-sort
go test -tags=contract -run '^$' -bench=HeapSort -benchmem ./src/algorithms/sorting/comparison/heap-sort
```

The benchmark covers random, reverse, and duplicate-heavy integer inputs at
1,024 and 32,768 items. It excludes input copying from the timed region and
checks the final output after timing stops.
