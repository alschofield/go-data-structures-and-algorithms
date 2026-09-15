# Dijkstra

## How It Works

A min-priority queue settles the lowest tentative distance then relaxes its non-negative outgoing edges.

## Required API

`func Dijkstra[T any](graph graph.Graph[T], source_key int) (DijkstraResult, error)`, where `DijkstraResult` exposes `Distance(key int) (int64, bool)` and `Parent(key int) (int, bool)`; unreachable keys have no distance.

## Contract

Return `ErrNilGraph` for a nil graph, `graph.ErrInvalidKey` for an invalid source, and `ErrNegativeWeight` for a negative edge; propagate errors from `Graph.Neighbors`. Settle and relax neighbor weights by stable node key. Settled distances never change. Support cycles, parallel edges, and self-loops using decrease-key or stale-entry skipping. Parent keys reconstruct shortest paths. Do not use a library shortest-path routine.

## Complexity Targets

O((V+E) log V) time and O(V) space.

## Path Reconstruction

`Distance(key)` returns the shortest total weight from the source to `key`.
`Parent(key)` returns the predecessor key selected by the final relaxation; the
source has no parent. Do not reverse either map. To build a source-to-target
path, collect the target and each parent through the source into a separate key
slice, then reverse that slice.

## Verification

```sh
make contract NAME=algorithms/shortest-paths/dijkstra
go test -tags=contract -run '^$' -bench=Dijkstra -benchmem ./src/algorithms/shortest-paths/dijkstra
```

The benchmark builds weighted chains of 256 and 1,024 nodes before timing.
