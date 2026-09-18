# A-Star

## How It Works

A* is Dijkstra directed toward one goal. For each candidate node:

```text
g(node) = actual shortest cost discovered from source to node
h(node) = estimated remaining cost from node to goal
f(node) = g(node) + h(node)
```

The frontier pops the lowest `f` value first. This repository's `BinaryHeap` is
a max-heap, so its comparator must treat a smaller `f` value as higher priority.
Candidates carry immutable `g` and `f` snapshots; after a better route is found,
push a replacement candidate and skip the old candidate when popped.

## Required API

```go
type Heuristic func(key int) int64

func AStar[T any](
    inputGraph graph.Graph[T],
    sourceKey, goalKey int,
    heuristic Heuristic,
) ([]*graph.Node[T], error)
```

The returned slice is the source-to-goal node path, including both endpoints.

## Contract

Return `ErrNilGraph`, `ErrNilHeuristic`, `graph.ErrInvalidKey`,
`ErrNegativeWeight`, or `ErrNegativeHeuristic` for those respective invalid
inputs; propagate errors from `Graph.Neighbors`. A zero heuristic must behave
as Dijkstra. Resolve ties deterministically, reconstruct an optimal node path
with admissible heuristics, and report no path as a nil path with nil error
when exhausted. Do not use library pathfinding.

For graph search, use a **consistent** heuristic when stopping as soon as the
goal is popped:

```text
h(current) <= edgeWeight(current, neighbor) + h(neighbor)
```

Consistency implies admissibility and keeps `f` values nondecreasing along a
path. A zero heuristic is consistent, so it reduces A* to Dijkstra.

## Path Reconstruction

Store `parent[neighborKey] = currentKey` only when relaxation improves the
neighbor's `g` score. When the goal is reached, collect goal, its parents, and
the source into a separate key slice, reverse that slice, then resolve keys to
nodes. Do not reverse the `parents` or `g` score maps.

## Complexity Targets

Worst O((V+E) log V) time and O(V) space.

## Verification

```sh
just contract algorithms/shortest-paths/a-star
go test -tags=contract -run '^$' -bench=AStar -benchmem ./src/algorithms/shortest-paths/a-star
```

The contract covers the zero-heuristic Dijkstra-equivalent path and an
unreachable goal. Benchmarks compare a zero heuristic with an exact chain
heuristic at 256 and 1,024 nodes; graph construction occurs before timing.
