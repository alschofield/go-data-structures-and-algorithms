# Counting Sort

## How It Works

Count occurrences of each key in `[0, keyLimit)`, then rewrite `items` in
ascending key order by consuming those counts.

## Required API

```go
func CountingSort(items []uint32, keyLimit uint32) (bool, error)
```

## Contract

Sort `items` in place and return `(true, nil)`. Every value must be less than
`keyLimit`; validate the complete input before rewriting it. If any value is out
of range, return `(false, ErrKeyOutOfRange)` and leave `items` unchanged. A zero
`keyLimit` succeeds for empty input and rejects any non-empty input. This API
sorts bare `uint32` values, so it has no distinct equal payloads whose order
could be observed.

## Complexity Targets

O(n+k) time and O(k) auxiliary space, where `k` is `keyLimit`.

## Verification

```sh
just contract algorithms/sorting/non-comparison/counting-sort
```
