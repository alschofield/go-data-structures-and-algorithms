# Breadth-First Search

## How It Works
A FIFO frontier explores every vertex at distance k before distance k+1.

## Required API
`func BreadthFirstSearch(graph *adjacency_list.AdjacencyList, source int) ([]int, bool)`.

## Contract
Mark visited when enqueueing, visit each reachable vertex exactly once, reject invalid source, handle cycles/self-loops/disconnected graphs, and never mutate the graph. Do not use a library graph traversal.

## Complexity Targets
O(V+E) time and O(V) space with adjacency lists.
