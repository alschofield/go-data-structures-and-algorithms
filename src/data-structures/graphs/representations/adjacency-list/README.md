# Adjacency List

## How It Works
Each dynamically added node owns a list of weighted outgoing edges, so storage follows actual edges.

## Required API
`type AdjacencyList` with `NewAdjacencyList(directed bool)`, `AddVertex(value any)
int`, `VertexAt(index int) (any, bool, error)`, `AddEdge(from, to int, weight
int64) (bool, error)`, `RemoveEdge(from, to int) (bool, error)`, `HasEdge(from,
to int) (bool, error)`, `Neighbors(vertex int, func(to int, weight int64) bool)
(bool, error)`, `VertexCount() int`, `EdgeCount() int`, and `AsGraph() graph.Graph`.

## Contract
The constructor creates an empty graph; `AddVertex` returns its stable dense
index and `VertexAt` uses insertion order. An invalid vertex or edge endpoint
returns `ErrInvalidVertex` or `ErrInvalidEdge` and preserves graph state.
Undirected graphs record both directions with the same
weight. Reject duplicate edges, permit self-loops and negative weights, and
visit out-edges once in deterministic insertion order. `AsGraph` preserves
direction, weights, and early-stop behavior. Do not use a library graph type.

## Complexity Targets
AddVertex and AddEdge amortized O(1); HasEdge/neighbors O(deg(u)); full traversal O(V+E); O(V+E) space.
