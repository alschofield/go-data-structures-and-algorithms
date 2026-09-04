# A-Star

## How It Works
A min-priority frontier orders candidates by `f = g + h`: cost so far plus an admissible goal estimate.

## Required API
`type Heuristic func(vertex int) int64` and `func AStar(graph graph.Graph,
source, goal int, heuristic Heuristic) ([]int, error)`.

## Contract
Return `ErrNilGraph`, `ErrNilHeuristic`, `ErrInvalidVertex`,
`ErrNegativeWeight`, or `ErrNegativeHeuristic` for those respective invalid
inputs; propagate malformed-neighbor errors from `Graph.Neighbors`. A zero
heuristic must behave as Dijkstra. Resolve ties deterministically, reconstruct
an optimal index path with admissible heuristics, and report no path as a nil
path with nil error when exhausted. Do not use library pathfinding.

## Complexity Targets
Worst O((V+E) log V) time and O(V) space.
