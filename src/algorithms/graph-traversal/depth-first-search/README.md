# Depth-First Search

## How It Works
An explicit or call-stack frontier explores one branch as far as possible before backtracking.

## Required API
`func DepthFirstSearch(graph graph_view.GraphView, source graph_view.NodeHandle) ([]graph_view.NodeHandle, bool)`.

## Contract
Use a visited set of handles, visit each reachable node once, reject invalid or foreign source handles, handle cycles/self-loops/disconnected graphs, and never mutate the graph. Traverse through dynamic GraphView neighbors but ignore every edge weight. Return handles in visit order. Understand both recursive and explicit stack approaches. Do not use library traversal.

## Complexity Targets
O(V+E) time and O(V) space.
