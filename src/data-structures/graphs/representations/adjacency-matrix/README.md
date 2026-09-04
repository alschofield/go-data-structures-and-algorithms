# Adjacency Matrix

## How It Works
A dynamically grown contiguous N by N grid stores weighted edge presence for node-handle indexes.

## Required API
`type AdjacencyMatrix` with `NewAdjacencyMatrix(directed bool)`, `AddVertex(value
any) int`, `VertexAt(index int) (any, bool, error)`, `AddEdge(from, to int,
weight int64) (bool, error)`, `RemoveEdge(from, to int) (bool, error)`,
`HasEdge(from, to int) (bool, error)`, `Neighbors(vertex int, func(to int,
weight int64) bool) (bool, error)`, `VertexCount() int`, `EdgeCount() int`, and
`AsGraph() graph.Graph`.

## Contract
The constructor creates an empty graph; `AddVertex` returns its stable dense
index and `VertexAt` uses insertion order. An invalid vertex or edge endpoint
returns `ErrInvalidVertex` or `ErrInvalidEdge` and preserves graph state.
Undirected mutations update symmetric cells with the same
weight. Duplicate add and absent remove are clean no-ops. Neighbor walking
scans an entire row and stops when its visitor returns false. `AsGraph`
preserves direction, weights, and early-stop behavior; do not substitute
another representation.

## Complexity Targets
AddVertex O(N^2); Add/Remove/HasEdge O(1); neighbors O(N); full traversal O(N^2); O(N^2) space.
