# Graph

## Required API

`type Graph interface` with `Directed() bool`, `VertexCount() int`, and
`Neighbors(vertex int, visit func(neighbor int, weight int64) bool) (bool, error)`.

## Contract

- Vertices are dense indexes in `[0, VertexCount())`. `Neighbors` returns
  `ErrInvalidVertex` for an out-of-range vertex, visits each outgoing weighted
  edge once in deterministic order, and stops when its visitor returns false.
  Its boolean reports whether iteration completed rather than an input error.
- `Directed` describes the graph rather than its adapter. Adjacency-list and
  adjacency-matrix implementations expose `AsGraph() Graph`; the interface
  neither owns nor mutates their storage and never exposes values or handles.
- Edge weights use `int64`. Traversal ignores them; Dijkstra and A-star reject
  negative weights; Kruskal accepts signed weights for an undirected graph.

## Complexity Targets

VertexCount is O(1); neighbor iteration matches the adapted representation.
