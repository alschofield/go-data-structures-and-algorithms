# Dijkstra

## How It Works

A min-priority queue settles the lowest tentative distance then relaxes its non-negative outgoing edges.

## Required API

`func Dijkstra[T any](graph graph.Graph[T], source_key int) (DijkstraResult, error)`, where `DijkstraResult` exposes `Distance(key int) (int64, bool)` and `Parent(key int) (int, bool)`; unreachable keys have no distance.

## Contract

Return `ErrNilGraph` for a nil graph, `graph.ErrInvalidKey` for an invalid source, and `ErrNegativeWeight` for a negative edge; propagate errors from `Graph.Neighbors`. Settle and relax neighbor weights by stable node key. Settled distances never change. Support cycles, parallel edges, and self-loops using decrease-key or stale-entry skipping. Parent keys reconstruct shortest paths. Do not use a library shortest-path routine.

## Complexity Targets

O((V+E) log V) time and O(V) space.
