# Data Structures and Algorithms in Go

This repository mirrors the canonical C curriculum's 27-leaf taxonomy. Each
leaf contains a Go API contract and a generated test scaffold. Implement every
exercise from first principles; do not substitute standard-library containers,
sorts, searches, or graph algorithms.

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
src/data-structures/graphs/graph-view
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
```

## Commands

```sh
go test ./...
```

The test scaffolds intentionally fail until their production API is written.
