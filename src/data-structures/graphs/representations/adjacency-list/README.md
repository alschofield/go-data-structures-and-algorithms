# Adjacency List

## Required API

`type AdjacencyList[T any]` with `NewAdjacencyList[T](directed bool)`,
`AddVertex(value *T) (*graph.Node[T], error)`,
`NodeByKey(key int) (*graph.Node[T], bool, error)`, `AddEdge(from_key, to_key
int, weight int64) (bool, error)`, `RemoveEdge(from_key, to_key int) (bool,
error)`, `HasEdge(from_key, to_key int) (bool, error)`, `Neighbors(key int,
func(*graph.Node[T], int64) bool) (bool, error)`, `NodeCount() int`, and
`EdgeCount() int`. It implements `graph.Graph[T]`; undirected instances also
implement `graph.UndirectedEdgeGraph[T]` with `Edges(func(graph.Edge[T]) bool)
bool`.

## Contract

`AddVertex` retains the caller-provided value pointer in a node with a stable,
unique key. Invalid or removed keys return `graph.ErrInvalidKey` and preserve
state. Undirected graphs store mirrored adjacency but `Edges` reports each
logical edge once in insertion order. Reject duplicate edges, permit self-loops
and negative weights, and visit outgoing edges once in insertion order. Do not
use a library graph type.

## Complexity Targets

AddVertex and AddEdge amortized O(1); HasEdge/neighbors O(deg(u)); full traversal O(V+E); O(V+E) space.
