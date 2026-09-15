# Radix Sort

## How It Works

Four stable counting passes process `uint32` bytes from least to most
significant at shifts `0`, `8`, `16`, and `24`. Each pass emits a separate
buffer, and the final buffer is copied back to `items`.

## Required API

```go
func RadixCount(items []uint32, shift uint) ([]uint32, bool)
func RadixSort(items []uint32) bool
```

`RadixCount` returns a newly allocated stable ordering by the byte at `shift`.

## Contract

`RadixSort` orders `items` ascending in place and returns `true`; nil, empty,
and singleton slices are successful no-ops. `RadixCount` returns `true` with a
new output slice whose equal-byte values retain input order. `RadixSort` relies
on that stability across all four byte passes. Neither function reports an
error or accepts a comparator.

## Complexity Targets

For `uint32`, O(n+256) time and O(n+256) auxiliary space: four passes make the
constant digit count independent of input size.

## Verification

```sh
make contract NAME=algorithms/sorting/non-comparison/radix-sort
```
