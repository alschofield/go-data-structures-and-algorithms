# Depth-First Search

## How It Works
An explicit or call-stack frontier explores one branch as far as possible before backtracking.

## Required API
`func DepthFirstSearch(graph graph.Graph, source int) ([]int, bool)`.

## Contract
Use a visited set of vertex indexes, visit each reachable vertex once, reject an
invalid source, handle cycles/self-loops/disconnected graphs, and never mutate
the graph. Traverse through `Graph.Neighbors` but ignore every edge weight.
Return indexes in visit order. Understand both recursive and explicit stack
approaches. Do not use a library traversal.

## Complexity Targets
O(V+E) time and O(V) space.
