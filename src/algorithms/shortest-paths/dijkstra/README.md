# Dijkstra

## How It Works
A min-priority queue settles the lowest tentative distance then relaxes its non-negative outgoing edges.

## Required API
`func Dijkstra(graph graph.Graph, source int) (DijkstraResult, error)`, where
`DijkstraResult` exposes `Distance(vertex int) (int64, bool)` and
`Parent(vertex int) (int, bool)`; unreachable vertices have no distance.

## Contract
Return `ErrNilGraph` for a nil graph, `ErrInvalidVertex` for an invalid source,
and `ErrNegativeWeight` for a negative edge; propagate malformed-neighbor
errors from `Graph.Neighbors`. Settle and relax neighbor weights. Settled
distances never change. Support cycles, parallel edges, and self-loops using
decrease-key or stale-entry skipping. Parent indexes reconstruct shortest paths.
Do not use a library shortest-path routine.

## Complexity Targets
O((V+E) log V) time and O(V) space.
