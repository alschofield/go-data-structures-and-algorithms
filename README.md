# Data Structures and Algorithms in Go

This repository mirrors the canonical C curriculum's 28-leaf taxonomy. Each
leaf contains an API contract, opt-in table-driven TDD tests, and a benchmark
plan. Implement every exercise from first principles; do not substitute
standard-library containers, maps, heaps, sorts, searches, or graph algorithms.

## Taxonomy

```text
src/data-structures/linear/arrays/dynamic-array
src/data-structures/linear/stacks/stack
src/data-structures/linear/queues/queue
src/data-structures/linear/linked/singly-linked-list
src/data-structures/linear/linked/doubly-linked-list
src/data-structures/associative/hash-tables/separate-chaining
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
matrix retain ownership of values and storage, then expose their dense indexed
view through `AsGraph()`. Traversal, shortest-path, and spanning-tree
algorithms operate only on that `Graph` interface.

## Error Handling

All leaf APIs follow the [error-handling contract](ERROR-HANDLING.md): final
`error` values report only invalid arguments or invalid state; expected
absence and empty lookups remain ordinary boolean or zero/empty results.
Contract tests are opt-in and encode these signatures without requiring any
unfinished production declaration during default validation.

## Status

All 28 leaves include user-owned package-only production source scaffolds under
`src/`. These files intentionally contain only their package declaration;
implementations remain the learner's responsibility. A leaf becomes available
to contract tests only after its documented API has been implemented.

## TDD Workflow

Default tests deliberately exclude unfinished contract tests, so a fresh clone
remains usable. After implementing a leaf, compile and run its contracts:

```sh
make contract NAME=data-structures/linear/stacks/stack
make contract NAME=algorithms/graph-traversal/breadth-first-search
```

The `contract` build tag is intentional, not an implementation escape hatch:
remove no tags from the tests. The learner's production API makes the test
package compile; the behavioral assertions then drive the implementation.

## Commands

```sh
make test NAME=data-structures/linear/stacks/stack
make test-all
make contract NAME=data-structures/linear/stacks/stack
make benchmark NAME=data-structures/linear/stacks/stack
make race
make benchmark-compare OLD=before.txt NEW=after.txt
```

`make test` and `make test-all` are safe before any exercise is implemented.
`make contract` and `make benchmark` require the named leaf's production API.
See [bench/README.md](bench/README.md) for benchmark design and all leaf plans.
