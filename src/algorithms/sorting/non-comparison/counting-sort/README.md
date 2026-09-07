# Counting Sort

## How It Works
Count keys in `[0,keyLimit)`, prefix-sum those counts, then place values into an output buffer.

## Required API
`func CountingSort(items []uint32, keyLimit uint32) (bool, error)`.

## Contract
Use no comparisons. Success returns `(true, nil)`. Validate all keys before
mutation; a key outside `[0,keyLimit)` returns `(false, ErrKeyOutOfRange)` and
leaves input unchanged. A zero `keyLimit` is valid only for empty input. Place
in reverse or equivalently to remain stable. Do not call `sort`.

## Complexity Targets
Best/average/worst O(n+k), O(n+k) auxiliary space.
