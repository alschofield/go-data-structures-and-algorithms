# Adjacency Matrix

## How It Works

An N by N matrix stores an edge record at each source-key and destination-key
cell. Undirected graphs mirror each logical edge in its reciprocal cell.

## Required API

```go
type AdjacencyMatrix[T any] struct

func NewAdjacencyMatrix[T any](directed bool) *AdjacencyMatrix[T]
func (am *AdjacencyMatrix[T]) AddVertex(value *T) *graph.Node[T]
func (am *AdjacencyMatrix[T]) NodeByKey(key int) (*graph.Node[T], bool)
func (am *AdjacencyMatrix[T]) AddEdge(fromKey, toKey int, weight int64) (bool, error)
func (am *AdjacencyMatrix[T]) RemoveEdge(fromKey, toKey int) (bool, error)
func (am *AdjacencyMatrix[T]) HasEdge(fromKey, toKey int) (bool, error)
func (am *AdjacencyMatrix[T]) Neighbors(key int, visit func(*graph.Node[T], int64) bool) (bool, error)
func (am *AdjacencyMatrix[T]) NodeCount() int
func (am *AdjacencyMatrix[T]) EdgeCount() int
func (am *AdjacencyMatrix[T]) Directed() bool
func (am *AdjacencyMatrix[T]) Edges(visit func(graph.Edge[T]) bool) (bool, error)
```

It implements `graph.Graph[T]`; undirected instances also provide
`graph.UndirectedEdgeGraph[T]`.

## Contract

`AddVertex` retains the supplied value pointer and returns a node with a stable,
unique key. `NodeByKey` returns `ok=false` when absent. Edge operations validate
both keys and return `graph.ErrInvalidKey` without mutation when either is
invalid. Duplicate adds and absent removals return `false, nil`.

`EdgeCount` counts logical edges. An undirected mutation updates both cells;
`Edges` reports each logical edge once in row-major key order. `Edges` on a
directed matrix returns `graph.ErrDirectedGraph`. `Neighbors` scans a row in
ascending key order and returns `false, nil` on visitor early stop. Self-loops
and signed weights are stored without restriction.

## Complexity Targets

`AddVertex` is O(N^2); `AddEdge`, `RemoveEdge`, and `HasEdge` are O(1);
`Neighbors` is O(N); `Edges` is O(N^2); space is O(N^2).

## Verification

```sh
just contract data-structures/graphs/representations/adjacency-matrix
```
