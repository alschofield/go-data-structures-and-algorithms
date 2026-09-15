# Linear Search

## How It Works

Scan `items` from index zero until an item compares equal to the key. The first
match ends the scan.

## Required API

```go
func LinearSearch[T any](
    items []T,
    key T,
    compare func(T, T) int,
) (int, bool, error)
```

## Contract

Accept sorted or unsorted input and do not mutate `items`. Return the first
index whose item compares equal to `key` as `(index, true, nil)`. A missing key,
including on a nil or empty slice, returns `(0, false, nil)`.

A nil comparator returns `(0, false, ErrNilComparator)` before examining the
slice. The boolean distinguishes a match at index zero from a miss.

## Complexity Targets

Best O(1), average and worst O(n) time; O(1) auxiliary space.

## Verification

```sh
make contract NAME=algorithms/searching/linear-search
```
