# Graph

## Required API

`type Node[T any]` with `Key`, `Value`, `Next`, `Prev`, `Left`, `Right`,
`Parent`, `Children`, and `Edges`; `type Edge[T any] struct { From, To
*Node[T]; Weight int64 }`; and `type Graph[T any] interface` with `Directed()
bool`, `NodeCount() int`, `NodeByKey(key int) (*Node[T], bool, error)`, and
`Neighbors(key int, visit func(*Node[T], int64) bool) (bool, error)`.

`type UndirectedEdgeGraph[T any] interface` extends `Graph[T]` with
`Edges(visit func(Edge[T]) bool) bool`.

## Contract

- `Node.Key` is a stable, unique graph identity. Keys may have gaps after node
  removal and algorithms use them as map keys rather than array indexes.
- `Node.Value` points to the value retained by the concrete data structure.
  Node-backed structures reuse this record and leave links irrelevant to their
  representation nil. Each representation owns its value lifetime and must not
  substitute a copied or unrelated payload.
- `Node.Occurrences` is an optional structure-owned observational metric. A
  structure that uses it must document whether it changes mutation behavior.
- `NodeByKey` and `Neighbors` return `ErrInvalidKey` for an absent key.
  `Neighbors` visits outgoing weighted node edges in deterministic order and
  returns false only when its visitor requests an early stop.
- Adjacency-list and adjacency-matrix structs implement `Graph[T]` directly.
  BSTs may implement it as a directed tree view. The interface never owns or
  mutates representation storage.
- Undirected representations implement `UndirectedEdgeGraph[T]` by reporting
  each logical edge once in deterministic order. Kruskal rejects directed
  graphs before reading edges.
- Edge weights use `int64`. Traversal ignores them; Dijkstra and A-star reject
  negative weights; Kruskal accepts signed weights for undirected graphs.

## Algorithm Consumers

- BFS and DFS accept `Graph[T]` and return visited nodes.
- Dijkstra and A-star accept `Graph[T]` and use `Node.Key` for their distance
  and parent state.
- Kruskal accepts `UndirectedEdgeGraph[T]` and returns `Edge[T]` values.

## Verification

`make contract NAME=data-structures/graphs/graph`
