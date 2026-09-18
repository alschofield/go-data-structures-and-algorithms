# Data Structures and Algorithms in Go

This repository follows the canonical C curriculum with a 27-leaf Go taxonomy.
Each leaf contains an API contract, opt-in table-driven TDD tests, and a
benchmark plan. Implement every exercise from first principles; do not
substitute standard-library containers, maps, heaps, sorts, searches, or graph
algorithms.

## Dynamic Sequence Policy

The C curriculum implements a dynamic array. Higher-level language curricula
use their native dynamic sequence as the baseline rather than duplicate a
runtime container: Go slice, C++ vector, Java ArrayList, Rust Vec, Python list,
TypeScript array, C# List<T>, and Kotlin MutableList. Therefore, Go has no
dynamic-array exercise; slices are the baseline for exercises that require
dynamic contiguous storage.

## Taxonomy

```text
src/data-structures/linear/stacks/stack
src/data-structures/linear/queues/queue
src/data-structures/linear/linked/singly-linked-list
src/data-structures/linear/linked/doubly-linked-list
src/data-structures/associative/hash-table
src/data-structures/trees/binary-search-trees/binary-search-tree
src/data-structures/trees/tries/prefix-trie
src/data-structures/trees/heaps/binary-heap
src/data-structures/graphs/graph
src/data-structures/graphs/representations/adjacency-list
src/data-structures/graphs/representations/adjacency-matrix
src/data-structures/graphs/disjoint-sets/union-find
src/algorithms/searching/linear-search
src/algorithms/searching/binary-search
src/algorithms/sorting/comparison/bubble-sort
src/algorithms/sorting/comparison/selection-sort
src/algorithms/sorting/comparison/insertion-sort
src/algorithms/sorting/comparison/merge-sort
src/algorithms/sorting/comparison/quick-sort
src/algorithms/sorting/comparison/heap-sort
src/algorithms/sorting/non-comparison/counting-sort
src/algorithms/sorting/non-comparison/radix-sort
src/algorithms/graph-traversal/breadth-first-search
src/algorithms/graph-traversal/depth-first-search
src/algorithms/shortest-paths/dijkstra
src/algorithms/shortest-paths/a-star
src/algorithms/minimum-spanning-trees/kruskal
```

`graph` defines the representation-neutral graph contract. Adjacency list and
matrix retain ownership of values and storage while directly implementing its
stable-key node view. Traversal and shortest-path algorithms operate on `Graph`;
Kruskal operates on `UndirectedEdgeGraph` so it receives each logical edge once.

## Error Handling

All leaf APIs follow the [error-handling contract](ERROR-HANDLING.md): final
`error` values report only invalid arguments or invalid state; expected
absence and empty lookups remain ordinary boolean or zero/empty results.
Contract tests are opt-in and encode these signatures without requiring any
unfinished production declaration during default validation.

## Status

All 28 taxonomy leaves compile and pass their opt-in contract suites:

```sh
go test -tags=contract ./...
```

This includes every linear, associative, tree, graph, search, sort, traversal,
shortest-path, and minimum-spanning-tree leaf. Each leaf README is the
authoritative API and behavioral contract; benchmark coverage is tracked in
`bench/README.md` and each completed graph algorithm's own README.

## TDD Workflow

Default tests deliberately exclude contract tests, so a fresh clone remains
usable. Run an individual leaf contract while implementing or reviewing it:

```sh
just contract data-structures/linear/stacks/stack
just contract algorithms/graph-traversal/breadth-first-search
```

The `contract` build tag is intentional, not an implementation escape hatch.
Keep it on all contract suites so the default test path remains fast and the
full contract gate remains explicit.

## Commands

```sh
just test data-structures/linear/stacks/stack
just test-all
just contract data-structures/linear/stacks/stack
just benchmark data-structures/linear/stacks/stack
go test -tags=contract ./...
just race
just benchmark-compare before.txt after.txt
```

`just test` and `just test-all` remain the safe default suite. `just contract`
and `just benchmark` target a named completed leaf. See
[bench/README.md](bench/README.md) for benchmark design and all leaf plans.
