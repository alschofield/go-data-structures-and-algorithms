# Depth-First Search

## How It Works

A LIFO stack explores the most recently discovered branch before backtracking.
Mark a key visited when it is pushed so cycles and converging edges cannot push
it again.

## Required API

```go
func DepthFirstSearch[T any](
    input_graph graph.Graph[T],
    source_key int,
    visit func(*graph.Node[T]) bool,
) ([]*graph.Node[T], bool, error)
```

The returned slice is the pop order. `visit` returning `false` stops after the
current node has been appended to that slice.

## Contract

Return an empty path and `false` with `ErrNilGraph` for a nil graph,
`ErrNilVisit` for a nil visitor, or `graph.ErrInvalidKey` for an invalid or nil
source node. Propagate an error from `Graph.Neighbors` with the path accumulated
so far and a `false` completion result. Ignore edge weights, do not mutate the
graph, and visit each reachable key at most once. The boolean is `true` after
exhaustive traversal and after a visitor-requested stop; a graph that ends
neighbor iteration early returns its partial path with `false` and no error.

Neighbors are pushed in graph iteration order, so the last enumerated neighbor
is popped first. Disconnected nodes are not returned.

## Complexity Targets

O(V+E) time and O(V) space with adjacency lists.

## Verification

```sh
make contract NAME=algorithms/graph-traversal/depth-first-search
```
