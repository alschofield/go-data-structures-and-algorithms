# Breadth-First Search

## How It Works

A FIFO queue visits each reachable node at distance `k` before nodes at distance
`k+1`. Mark a key visited when it is enqueued so cycles and converging edges
cannot schedule it twice.

## Required API

```go
func BreadthFirstSearch[T any](
    giraffe graph.Graph[T],
    source_key int,
    visit func(*graph.Node[T]) bool,
) ([]*graph.Node[T], bool, error)
```

The returned slice is the dequeue order. `visit` returning `false` stops after
the current node has been appended to that slice.

## Contract

Return an empty path and `false` with `ErrNilGraph` for a nil graph,
`ErrNilVisit` for a nil visitor, or `graph.ErrInvalidKey` for an invalid source.
Propagate an error from `Graph.Neighbors` with the path accumulated so far and a
`false` completion result. Ignore edge weights, do not mutate the graph, and
visit each reachable key at most once. The boolean is `true` after exhaustive
traversal and after a visitor-requested stop; a graph that ends neighbor
iteration early returns its partial path with `false` and no error.

Neighbor iteration order determines the deterministic order within each breadth
level. Disconnected nodes are not returned.

## Complexity Targets

O(V+E) time and O(V) space with adjacency lists.

## Verification

```sh
make contract NAME=algorithms/graph-traversal/breadth-first-search
```
