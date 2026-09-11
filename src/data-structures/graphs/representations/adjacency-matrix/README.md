# Adjacency Matrix

## Required API

`type AdjacencyMatrix[T any]` with `NewAdjacencyMatrix[T](directed bool)`,
`AddVertex(value *T) (*graph.Node[T], error)`, `NodeByKey(key int)
(*graph.Node[T], bool, error)`, `AddEdge(from_key, to_key int, weight int64)
(bool, error)`, `RemoveEdge(from_key, to_key int) (bool, error)`,
`HasEdge(from_key, to_key int) (bool, error)`, `Neighbors(key int,
func(*graph.Node[T], int64) bool) (bool, error)`, `NodeCount() int`, and
`EdgeCount() int`. It implements `graph.Graph[T]`; undirected instances also
implement `graph.UndirectedEdgeGraph[T]` with `Edges(func(graph.Edge[T]) bool)
bool`.

## Contract

`AddVertex` retains the caller-provided value pointer in a node with a stable,
unique key. Invalid or removed keys return `graph.ErrInvalidKey` and preserve
state. Undirected mutations update symmetric cells while `Edges` reports each
logical edge once in row-major key order. Duplicate add and absent remove are
clean no-ops. Neighbor walking scans a row and stops when its visitor returns
false; do not substitute another representation.

## Complexity Targets

AddVertex O(N^2); Add/Remove/HasEdge O(1); neighbors O(N); full traversal O(N^2); O(N^2) space.
