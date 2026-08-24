# Graph View

## Required API

`type GraphView interface` with `VertexCount() int` and `Neighbors(vertex int, visit func(neighbor int, weight uint64) bool) bool`.

## Contract

- Vertexes are dense indexes in `[0, VertexCount())`. `Neighbors` rejects out-of-range indexes, visits each outgoing weighted edge once in deterministic order, and stops when the visitor returns false.
- Weights are nonnegative. Adjacency-list, adjacency-matrix, and imported-graph adapters map their storage to vertex indexes. GraphView neither owns nor mutates graph storage and never exposes node handles or values.

## Complexity Targets

VertexCount is O(1); neighbor iteration matches the adapted representation.
