# Breadth-First Search

## How It Works
A FIFO frontier explores every vertex at distance k before distance k+1.

## Required API
`func BreadthFirstSearch(graph graph_view.GraphView, source graph_view.NodeHandle) ([]graph_view.NodeHandle, bool)`.

## Contract
Mark handles visited when enqueueing, visit each reachable node exactly once, reject invalid or foreign source handles, handle cycles/self-loops/disconnected graphs, and never mutate the graph. Traverse through dynamic GraphView neighbors but ignore every edge weight. Return handles in visit order. Do not use a library graph traversal.

## Complexity Targets
O(V+E) time and O(V) space with adjacency lists.
