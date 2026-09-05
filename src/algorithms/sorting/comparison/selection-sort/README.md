# Selection Sort

## API And Errors

`func SelectionSort[T any](items []T, compare func(T, T) int) (bool, error)`
sorts `items` in place. `compare(left, right)` must return a value below zero
when `left` sorts before `right`, zero when they are equivalent, and a value
above zero when `left` sorts after `right`.

Successful calls return `(true, nil)`. Empty and singleton slices are successful
no-ops. A nil comparator returns `(false, ErrNilComparator)` and leaves the
input unchanged.

## Selection Invariant

Before pass `i`, positions `[0, i)` are sorted and contain the `i` smallest
items. The pass finds the minimum in `[i, n)` and swaps it into position `i` if
needed. That establishes the invariant for the next pass and leaves the whole
slice ascending after the final pass.

## Properties

- Selection sort is unstable: swapping the selected minimum can reorder equal
  values.
- It performs at most `n - 1` swaps because each of the first `n - 1` passes
  makes at most one conditional swap. The public API does not expose swap
  counts, so this property is documented rather than instrumented in tests.
- It uses no standard-library sorting routine.

## Complexity

The algorithm makes quadratic comparisons in every case: `O(n^2)` best,
average, and worst time. It uses `O(1)` auxiliary space apart from the input
slice.

## Benchmarks

The tagged benchmarks use deterministic random, sorted, and reverse inputs at
256 and 1,024 items. Each iteration copies its source input because the API
sorts in place; allocation results therefore include that required fresh-input
copy.

```sh
make contract NAME=algorithms/sorting/comparison/selection-sort
make benchmark NAME=algorithms/sorting/comparison/selection-sort
go test -tags=contract -run '^$' -bench BenchmarkSelectionSort -benchmem ./src/algorithms/sorting/comparison/selection-sort
```

Measured on September 4, 2026 with Go on Windows/amd64 (i9-11900K); these are
local regression evidence, not portable performance claims:

| Input | 256 items | 1,024 items |
| --- | ---: | ---: |
| Random | 84.3 us/op, 2,048 B/op, 1 alloc/op | 1.14 ms/op, 8,195 B/op, 1 alloc/op |
| Sorted | 76.4 us/op, 2,048 B/op, 1 alloc/op | 1.19 ms/op, 8,192 B/op, 1 alloc/op |
| Reverse | 85.5 us/op, 2,048 B/op, 1 alloc/op | 936 us/op, 8,192 B/op, 1 alloc/op |
