# Radix Sort

## How It Works
Four stable counting-sort passes process `uint32` bytes from least to most
significant: shifts `0`, `8`, `16`, and `24`.

## Required API
`func RadixSort(items []uint32) bool`.

## Contract
Use stable per-byte counting sorts in LSD order. Each pass counts byte keys,
prefix-sums them into output offsets, and moves complete values into a buffer.
Success returns `true`; nil, empty, and singleton input are successful no-ops.
Do not compare keys or call `sort`.

## Complexity Targets
Best/average/worst O(d(n+k)), O(n+k) auxiliary space.

For `uint32`, `d` is four and `k` is 256.

## Verification
`make contract NAME=algorithms/sorting/non-comparison/radix-sort`

`go test -tags=contract -bench=. ./src/algorithms/sorting/non-comparison/radix-sort`
