# Adjacency List

## How It Works
Each dynamically added node owns a list of weighted outgoing edges, so storage follows actual edges.

## Required API
`type AdjacencyList` with `NewAdjacencyList(directed bool)`, `AddNode(value any) NodeHandle`, `FindNode(value any) (NodeHandle, bool)`, `NodeAt(index int) (NodeHandle, bool)`, `NodeValue(NodeHandle) (any, bool)`, `AddEdge(from, to NodeHandle, weight int64)`, `HasEdge`, `Neighbors(NodeHandle, func(NodeHandle, int64) bool) bool`, `NodeCount`, `EdgeCount`, and `AsGraphView() GraphView`.

## Contract
The constructor creates an empty graph; `AddNode` returns a stable graph-local handle. `NodeAt` uses insertion order, and `FindNode` locates a node by value. Reject foreign or invalid handles cleanly. Undirected graphs record both directions with the same weight. Reject duplicate edges, permit self-loops and negative weights, and visit weighted out-edges once in deterministic insertion order. `AsGraphView` preserves dynamic node lookup, edge direction, weights, and early-stop behavior. Do not use a library graph type.

## Complexity Targets
AddNode and AddEdge amortized O(1); HasEdge/neighbors O(deg(u)); full traversal O(V+E); O(V+E) space.
