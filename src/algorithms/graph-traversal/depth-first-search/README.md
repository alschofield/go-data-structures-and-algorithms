# Depth-First Search

## How It Works

An explicit or call-stack frontier explores one branch as far as possible before backtracking.

## Required API

`func DepthFirstSearch[T any](input_graph graph.Graph[T], source_key int,
visit func(*graph.Node[T]) bool) ([]*graph.Node[T], bool, error)`.

## Contract

Use a visited set of node keys and visit each reachable node once. Return
`ErrNilGraph` for a nil graph, `ErrNilVisit` for a nil visitor, and
`graph.ErrInvalidKey` for an invalid source key; propagate errors from
`Graph.Neighbors`. Handle cycles, self-loops, and disconnected graphs without
mutation; ignore edge weights. Nodes are pushed in neighbor iteration order, so
the most recently enumerated neighbor is visited first. Return nodes in visit
order with their retained value pointers. The boolean is true for full
exhaustion or a visitor-requested successful stop; otherwise it is false with
any partial path and available error. Do not use a library traversal.

## Complexity Targets

O(V+E) time and O(V) space.

## Verification

```sh
make contract NAME=algorithms/graph-traversal/depth-first-search
go test -tags=contract -run '^$' -bench=DepthFirstSearch -benchmem ./src/algorithms/graph-traversal/depth-first-search
```

Benchmarks cover a full wide depth traversal and a visitor-requested stop at 256
and 1,024 nodes. The adjacency-list graph is built before timing.
