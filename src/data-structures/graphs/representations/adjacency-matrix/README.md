# Adjacency Matrix

## How It Works
A single contiguous V by V grid stores edge presence at `u*V+v`.

## Required API
`type AdjacencyMatrix` with `NewAdjacencyMatrix(vertexCount int, directed bool)`, `AddEdge`, `RemoveEdge`, `HasEdge`, `Neighbors(vertex int, visit func(int) bool) bool`, `VertexCount`, and `EdgeCount`.

## Contract
Vertices are `[0,V)` and invalid vertices fail cleanly. Undirected mutations update symmetric cells. Duplicate add and absent remove are clean no-ops. Neighbor walking scans an entire row; do not substitute another representation.

## Complexity Targets
Add/Remove/HasEdge O(1); neighbors O(V); full traversal O(V^2); O(V^2) space.
