# Linear Search

## How It Works
Scan arbitrary input from front to back until equality matches the key.

## Required API
`func LinearSearch[T any](items []T, key T, compare func(T,T) int) (int, bool, error)`.

## Verified Behavior
`LinearSearch` accepts sorted or unsorted input and scans it from index zero.
When duplicates compare equal to `key`, it returns the first matching index.
It never modifies `items` and does not use a standard-library search routine.

A match returns `(index, true, nil)`. A missing key, including an empty or nil
slice, returns `(0, false, nil)`. A nil `compare` function returns
`(0, false, ErrNilComparator)`; callers can identify that sentinel with
`errors.Is`.

## Complexity Targets
Best O(1), average/worst O(n), O(1) space.

## Benchmarks
The benchmark uses deterministic ascending integer slices of 1,000, 16,000,
and 1,000,000 elements, looking up the final element to exercise the complete
linear scan. It checks each result and writes it to package-level sinks so the
compiler cannot remove the work. Allocation reporting is enabled.

Run it with:

```sh
go test -tags=contract -run '^$' -bench BenchmarkLinearSearchPresentLast -benchmem ./src/algorithms/searching/linear-search
```

No local timing figures are recorded here because benchmark results depend on
the machine, Go version, and current system load.
