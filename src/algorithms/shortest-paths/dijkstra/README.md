# Dijkstra

## How It Works
A min-priority queue settles the lowest tentative distance then relaxes its non-negative outgoing edges.

## Required API
`func Dijkstra(graph graph_view.GraphView, source graph_view.NodeHandle) (DijkstraResult, bool)`, where `DijkstraResult` exposes `Distance(NodeHandle) (uint64, bool)` and `Parent(NodeHandle) (NodeHandle, bool)`; unreachable nodes have no distance.

## Contract
Reject negative weights and invalid or foreign source handles. Settle and relax dynamic GraphView neighbor weights. Settled distances never change. Support cycles, parallel edges, and self-loops using decrease-key or stale-entry skipping. Parent handles reconstruct shortest paths. Do not use a library shortest-path routine.

## Complexity Targets
O((V+E) log V) time and O(V) space.
