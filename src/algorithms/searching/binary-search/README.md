# Binary Search

## How It Works
Compare sorted input's middle element and discard one candidate half per step.

## Required API
`func BinarySearch[T any](items []T, key T, compare func(T,T) int) (int, bool, error)`.

## Verified Behavior
`BinarySearch` requires input sorted in ascending order according to `compare`.
It does not sort or validate `items`, so results for unsorted input are not
defined by this API. It never modifies input and does not use a standard-library
search routine.

A match returns `(index, true, nil)`. For duplicates, any index containing a
matching item is valid. A missing key, including an empty or nil slice, returns
`(0, false, nil)`. A nil `compare` function returns
`(0, false, ErrNilComparator)`; callers can identify that sentinel with
`errors.Is`.

## Complexity Targets
Best O(1), average/worst O(log n), O(1) iterative space.

## Benchmarks
The benchmark uses deterministic ascending integer slices of 1,000, 16,000,
and 1,000,000 elements, looking up the final element. It checks each result and
writes it to package-level sinks so the compiler cannot remove the work.
Allocation reporting is enabled.

Run it with:

```sh
go test -tags=contract -run '^$' -bench BenchmarkBinarySearchPresentLast -benchmem ./src/algorithms/searching/binary-search
```

No local timing figures are recorded here because benchmark results depend on
the machine, Go version, and current system load.
