# Binary Search

## How It Works

Compare the key with the middle item of an ascending slice, then discard the
half that cannot contain a match. Repeat until finding an equal item or
exhausting the candidate range.

## Required API

```go
func BinarySearch[T any](
    items []T,
    key T,
    compare func(T, T) int,
) (int, bool, error)
```

## Contract

`items` must already be ascending according to `compare`; the function does not
sort or validate that precondition. It does not mutate `items`. A match returns
`(index, true, nil)`. When duplicates match, any matching index is valid. A
missing key, including on a nil or empty slice, returns `(0, false, nil)`.

A nil comparator returns `(0, false, ErrNilComparator)` before examining the
slice. The boolean distinguishes a match at index zero from a miss.

## Complexity Targets

Best O(1), average and worst O(log n) time; O(1) auxiliary space.

## Verification

```sh
just contract algorithms/searching/binary-search
```
