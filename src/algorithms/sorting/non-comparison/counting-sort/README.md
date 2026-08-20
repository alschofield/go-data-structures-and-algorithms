# Counting Sort

## How It Works
Count keys in `[0,keyLimit)`, prefix-sum those counts, then place values into an output buffer.

## Required API
`func CountingSort(items []uint32, keyLimit uint32) bool`.

## Contract
Use no comparisons. Validate all keys before mutation; invalid input or allocation failure leaves input unchanged. Place in reverse or equivalently to remain stable. Do not call `sort`.

## Complexity Targets
Best/average/worst O(n+k), O(n+k) auxiliary space.
