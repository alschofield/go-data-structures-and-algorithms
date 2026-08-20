# Dijkstra

## How It Works
A min-priority queue settles the lowest tentative distance then relaxes its non-negative outgoing edges.

## Required API
`type WeightedGraph` and `func Dijkstra(graph *WeightedGraph, source int) (distances []uint64, parents []int, ok bool)`; unreachable distance is `math.MaxUint64`.

## Contract
Reject negative weights and invalid source. Settled distances never change. Support cycles, parallel edges, and self-loops using decrease-key or stale-entry skipping. Parents reconstruct shortest paths. Do not use a library shortest-path routine.

## Complexity Targets
O((V+E) log V) time and O(V) space.
