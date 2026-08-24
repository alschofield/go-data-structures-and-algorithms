# A-Star

## How It Works
A min-priority frontier orders candidates by `f = g + h`: cost so far plus an admissible goal estimate.

## Required API
`type Heuristic func(graph_view.NodeHandle) uint64` and `func AStar(graph graph_view.GraphView, source, goal graph_view.NodeHandle, heuristic Heuristic) ([]graph_view.NodeHandle, bool)`.

## Contract
Reject invalid or foreign source and goal handles and require non-negative dynamic GraphView edge weights. A zero heuristic must behave as Dijkstra. Resolve ties deterministically, reconstruct an optimal handle path with admissible heuristics, and report no path when exhausted. Do not use library pathfinding.

## Complexity Targets
Worst O((V+E) log V) time and O(V) space.
