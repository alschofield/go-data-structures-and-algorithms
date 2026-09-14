# Breadth-First Search

## How It Works

A FIFO frontier explores every node at distance k before distance k+1.

## Required API

`func BreadthFirstSearch[T any](input_graph graph.Graph[T], source_key int,
visit func(*graph.Node[T]) bool) ([]*graph.Node[T], bool, error)`.

## Contract

Mark node keys visited when enqueueing and visit each reachable node once.
Return `ErrNilGraph` for a nil graph and `graph.ErrInvalidKey` for an invalid
source key; return `ErrNilVisit` for a nil visitor; and propagate errors from
`Graph.Neighbors`. Handle cycles, self-loops, and disconnected graphs, never
mutate the graph, and ignore edge weights. Return nodes in visit order with
their retained value pointers. The boolean is true for full exhaustion or a
visitor-requested successful stop; otherwise it is false with any partial path
and available error. Do not use a library graph traversal.

## Complexity Targets

O(V+E) time and O(V) space with adjacency lists.

## Verification

```sh
make contract NAME=algorithms/graph-traversal/breadth-first-search
go test -tags=contract -run '^$' -bench=BreadthFirstSearch -benchmem ./src/algorithms/graph-traversal/breadth-first-search
```

Benchmarks cover a full wide breadth traversal and a visitor-requested stop at
256 and 1,024 nodes. The adjacency-list graph is built before timing.
