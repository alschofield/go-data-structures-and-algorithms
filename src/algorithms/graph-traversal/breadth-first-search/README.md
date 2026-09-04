# Breadth-First Search

## How It Works
A FIFO frontier explores every vertex at distance k before distance k+1.

## Required API
`func BreadthFirstSearch(graph graph.Graph, source int) ([]int, bool)`.

## Contract
Mark vertices visited when enqueueing, visit each reachable vertex exactly once,
reject an invalid source, handle cycles/self-loops/disconnected graphs, and
never mutate the graph. Traverse through `Graph.Neighbors` but ignore every
edge weight. Return vertex indexes in visit order. Do not use a library graph
traversal.

## Complexity Targets
O(V+E) time and O(V) space with adjacency lists.
