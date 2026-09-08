# Counting Sort

## How It Works
Count keys in `[0,keyLimit)`, then rewrite the input in key order using each
frequency count.

## Required API
`func CountingSort(items []uint32, keyLimit uint32) (bool, error)`.

## Contract
Use no comparisons. Success returns `(true, nil)`. Validate all keys before
mutation; a key outside `[0,keyLimit)` returns `(false, ErrKeyOutOfRange)` and
leaves input unchanged. A zero `keyLimit` is valid only for empty input. This
in-place `[]uint32` form has no distinct equal payloads to preserve. Do not call
`sort`.

## Complexity Targets
Best/average/worst O(n+k), O(k) auxiliary space.

## Verification
`make contract NAME=algorithms/sorting/non-comparison/counting-sort`

`go test -tags=contract -bench=. ./src/algorithms/sorting/non-comparison/counting-sort`
