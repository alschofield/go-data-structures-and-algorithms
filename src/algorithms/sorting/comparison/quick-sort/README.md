# Quick Sort

## How It Works
Partition around a defended pivot then sort the lower and greater partitions.

## Required API
`func QuickSort[T any](items []T, compare func(T,T) int) (bool, error)`.

## Contract
Sort ascending in place, without stability. Success returns `(true, nil)`; a
nil comparator returns `(false, ErrNilComparator)` without changing input. Use
median-of-three or randomized pivots; handle duplicate/all-equal input without
unbounded recursion. Do not call `sort`.

## Recursion Behavior
The implementation uses a median-of-three pivot and three-way partitioning, so
items equal to the pivot are excluded from recursive calls. It recurses into the
lower partition first and then the greater partition. Recursing on the smaller
partition and iterating over the larger one is an optional stack-space
optimization, not a first-pass requirement.

## Complexity Targets
Best/average O(n log n), worst O(n^2). The current direct two-sided recursion
uses O(log n) stack space for balanced partitions and can use O(n) in the worst
case; smaller-side recursion would bound stack space to O(log n).

## Benchmarks
Run `make benchmark NAME=algorithms/sorting/comparison/quick-sort` to measure
the deterministic random, sorted, reverse, and all-equal inputs at 1,024 and
8,192 items. Results below were collected with `-benchmem`; each operation
copies a fixed input before sorting, so allocation figures include that copy.
They were measured with Go 1.25.5 on Windows/amd64 on an 11th Gen Intel Core
i9-11900K @ 3.50GHz.

| Input | Size | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: | ---: |
| Random | 1,024 | 102,217 | 8,192 | 1 |
| Sorted | 1,024 | 117,449 | 8,192 | 1 |
| Reverse | 1,024 | 98,395 | 8,192 | 1 |
| All equal | 1,024 | 5,613 | 8,192 | 1 |
| Random | 8,192 | 1,075,814 | 65,536 | 1 |
| Sorted | 8,192 | 2,084,686 | 65,536 | 1 |
| Reverse | 8,192 | 1,573,726 | 65,536 | 1 |
| All equal | 8,192 | 32,020 | 65,536 | 1 |

Benchmark timings vary with Go version, CPU, system load, and compiler
optimizations. They are a reproducible workload baseline, not a cross-machine
performance guarantee.
