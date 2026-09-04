# A-Star

## How It Works
A min-priority frontier orders candidates by `f = g + h`: cost so far plus an admissible goal estimate.

## Required API
`type Heuristic func(vertex int) int64` and `func AStar(graph graph.Graph,
source, goal int, heuristic Heuristic) ([]int, bool)`.

## Contract
Reject invalid source and goal indexes, negative edge weights, and a negative
heuristic result. A zero heuristic must behave as Dijkstra. Resolve ties
deterministically, reconstruct an optimal index path with admissible heuristics,
and report no path when exhausted. Do not use library pathfinding.

## Complexity Targets
Worst O((V+E) log V) time and O(V) space.
