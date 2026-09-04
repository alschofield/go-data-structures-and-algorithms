# Depth-First Search

## How It Works
An explicit or call-stack frontier explores one branch as far as possible before backtracking.

## Required API
`func DepthFirstSearch(graph graph.Graph, source int) ([]int, error)`.

## Contract
Use a visited set of vertex indexes and visit each reachable vertex once. Return
`ErrNilGraph` for a nil graph and `ErrInvalidVertex` for an invalid source;
propagate malformed-neighbor errors from `Graph.Neighbors`. Handle cycles,
self-loops, and disconnected graphs without mutating the graph; ignore edge
weights. Return indexes in visit order. Understand both recursive and explicit
stack approaches. Do not use a library traversal.

## Complexity Targets
O(V+E) time and O(V) space.
