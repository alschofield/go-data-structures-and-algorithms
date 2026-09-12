# Adjacency List

## Required API

`type AdjacencyList[T any]` with `NewAdjacencyList[T](directed bool)`,
`AddVertex(value *T) *graph.Node[T]`,
`NodeByKey(key int) (*graph.Node[T], bool)`, `AddEdge(from_key, to_key
int, weight int64) (bool, error)`, `RemoveEdge(from_key, to_key int) (bool,
error)`, `HasEdge(from_key, to_key int) (bool, error)`, `Neighbors(key int,
func(*graph.Node[T], int64) bool) (bool, error)`, `NodeCount() int`, and
`EdgeCount() int`. It implements `graph.Graph[T]`; undirected instances also
implement `graph.UndirectedEdgeGraph[T]` with `Edges(func(graph.Edge[T]) bool)
(bool, error)`.

## Contract

`AddVertex` retains the caller-provided value pointer in a node with a stable,
unique key. It cannot fail. `NodeByKey` reports an absent key with `ok=false`;
operations supplied invalid keys return `graph.ErrInvalidKey` and preserve state.
Undirected graphs store mirrored adjacency but `Edges` reports each
logical edge once in insertion order and rejects directed instances with
`graph.ErrDirectedGraph`. Reject duplicate edges, permit self-loops
and negative weights, and visit outgoing edges once in insertion order. Do not
use a library graph type.

## Complexity Targets

`AddVertex` and `NodeByKey` are O(1). `Neighbors`, `HasEdge`, duplicate-aware
`AddEdge`, and `RemoveEdge` are O(deg(u)); `Edges` is O(V+E). Full traversal is
O(V+E) with O(V+E) space.

## Verification

```sh
make contract NAME=data-structures/graphs/representations/adjacency-list
go test -tags=contract -run '^$' -bench=AdjacencyList -benchmem ./src/data-structures/graphs/representations/adjacency-list
```

Benchmarks cover edge insertion, present-edge lookup, and full neighbor walking
at 256 and 1,024 nodes. Setup constructs the needed graph before timed lookup
and traversal workloads; insertion measures a fresh graph construction plus its
final edge insertion.
