# Depth-First Search

## How It Works
An explicit or call-stack frontier explores one branch as far as possible before backtracking.

## Required API
`func DepthFirstSearch(graph *adjacency_list.AdjacencyList, source int) ([]int, bool)`.

## Contract
Use a visited set, visit each reachable vertex once, reject invalid source, handle cycles/self-loops/disconnected graphs, and never mutate the graph. Understand both recursive and explicit stack approaches. Do not use library traversal.

## Complexity Targets
O(V+E) time and O(V) space.
