# Adjacency Matrix

## How It Works
A dynamically grown contiguous N by N grid stores weighted edge presence for node-handle indexes.

## Required API
`type AdjacencyMatrix` with `NewAdjacencyMatrix(directed bool)`, `AddNode(value any) NodeHandle`, `FindNode(value any) (NodeHandle, bool)`, `NodeAt(index int) (NodeHandle, bool)`, `NodeValue(NodeHandle) (any, bool)`, `AddEdge(from, to NodeHandle, weight int64)`, `RemoveEdge`, `HasEdge`, `Neighbors(NodeHandle, func(NodeHandle, int64) bool) bool`, `NodeCount`, `EdgeCount`, and `AsGraphView() GraphView`.

## Contract
The constructor creates an empty graph; `AddNode` returns a stable graph-local handle. `NodeAt` uses insertion order, and `FindNode` locates a node by value. Reject foreign or invalid handles cleanly. Undirected mutations update symmetric cells with the same weight. Duplicate add and absent remove are clean no-ops. Neighbor walking scans an entire row and stops when its visitor returns false. `AsGraphView` preserves dynamic node lookup, edge direction, weights, and early-stop behavior; do not substitute another representation.

## Complexity Targets
AddNode O(N^2); Add/Remove/HasEdge O(1); neighbors O(N); full traversal O(N^2); O(N^2) space.
