# Radix Sort

## How It Works
Stable counting-sort passes process fixed-radix digits from least to most significant.

## Required API
`func RadixSort(items []uint32) bool`.

## Contract
Use stable per-digit counting sorts in LSD order and reuse auxiliary storage where
practical. Success returns `true`; allocation failure returns `false` and
preserves input. Do not compare keys or call `sort`.

## Complexity Targets
Best/average/worst O(d(n+k)), O(n+k) auxiliary space.
