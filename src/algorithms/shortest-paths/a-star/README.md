# A-Star

## How It Works

A min-priority frontier orders candidates by `f = g + h`: cost so far plus an admissible goal estimate.

## Required API

`type Heuristic func(key int) int64` and `func AStar[T any](graph graph.Graph[T], source_key, goal_key int, heuristic Heuristic) ([]*graph.Node[T], error)`.

## Contract

Return `ErrNilGraph`, `ErrNilHeuristic`, `graph.ErrInvalidKey`,
`ErrNegativeWeight`, or `ErrNegativeHeuristic` for those respective invalid
inputs; propagate errors from `Graph.Neighbors`. A zero heuristic must behave
as Dijkstra. Resolve ties deterministically, reconstruct an optimal node path
with admissible heuristics, and report no path as a nil path with nil error
when exhausted. Do not use library pathfinding.

## Complexity Targets

Worst O((V+E) log V) time and O(V) space.
