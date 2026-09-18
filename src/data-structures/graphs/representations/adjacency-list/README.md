# Adjacency List

## How It Works

Each node owns its outgoing weighted edges. Undirected graphs store reciprocal
arcs while tracking and exposing one logical edge.

## Required API

```go
type AdjacencyList[T any] struct

func NewAdjacencyList[T any](directed bool) *AdjacencyList[T]
func (al *AdjacencyList[T]) AddVertex(value *T) *graph.Node[T]
func (al *AdjacencyList[T]) NodeByKey(key int) (*graph.Node[T], bool)
func (al *AdjacencyList[T]) AddEdge(fromKey, toKey int, weight int64) (bool, error)
func (al *AdjacencyList[T]) RemoveEdge(fromKey, toKey int) (bool, error)
func (al *AdjacencyList[T]) HasEdge(fromKey, toKey int) (bool, error)
func (al *AdjacencyList[T]) Neighbors(key int, visit func(*graph.Node[T], int64) bool) (bool, error)
func (al *AdjacencyList[T]) NodeCount() int
func (al *AdjacencyList[T]) EdgeCount() int
func (al *AdjacencyList[T]) Directed() bool
func (al *AdjacencyList[T]) Edges(visit func(graph.Edge[T]) bool) (bool, error)
```

It implements `graph.Graph[T]`; undirected instances also provide
`graph.UndirectedEdgeGraph[T]`.

## Contract

`AddVertex` retains the supplied value pointer and returns a stable, unique key.
`NodeByKey` returns `ok=false` when absent. Edge operations validate both keys
and return `graph.ErrInvalidKey` without mutation when either is invalid.
Duplicate adds and absent removals return `false, nil`; self-loops and signed
weights are permitted.

`EdgeCount` counts logical edges. Undirected non-self-loop edges store reciprocal
arcs; `Edges` reports only the lower-key orientation, in node and adjacency
insertion order. `Edges` on a directed list returns `graph.ErrDirectedGraph`.
`Neighbors` visits outgoing edges in insertion order and returns `false, nil`
on visitor early stop.

## Complexity Targets

`AddVertex` and `NodeByKey` are O(1). `Neighbors`, `HasEdge`, `AddEdge`, and
`RemoveEdge` are O(deg(u)); `Edges` and full traversal are O(V+E); space is
O(V+E).

## Verification

```sh
just contract data-structures/graphs/representations/adjacency-list
```
