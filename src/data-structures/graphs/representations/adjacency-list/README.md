# Adjacency List

## How It Works
Each vertex owns a list of outgoing neighbors, so storage follows actual edges.

## Required API
`type AdjacencyList` with `NewAdjacencyList(vertexCount int, directed bool)`, `AddEdge`, `HasEdge`, `Neighbors(vertex int, visit func(int) bool) bool`, `VertexCount`, and `EdgeCount`.

## Contract
Vertices are `[0,V)`. Undirected graphs record both directions. Reject duplicate edges, permit self-loops, and visit out-edges once in deterministic insertion order. Do not use a library graph type.

## Complexity Targets
AddEdge amortized O(1); HasEdge/neighbors O(deg(u)); full traversal O(V+E); O(V+E) space.
