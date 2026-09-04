# Breadth-First Search

## How It Works
A FIFO frontier explores every vertex at distance k before distance k+1.

## Required API
`func BreadthFirstSearch(graph graph.Graph, source int) ([]int, error)`.

## Contract
Mark vertices visited when enqueueing and visit each reachable vertex exactly
once. Return `ErrNilGraph` for a nil graph and `ErrInvalidVertex` for an invalid
source; propagate malformed-neighbor errors from `Graph.Neighbors`. Handle
cycles/self-loops/disconnected graphs, never mutate the graph, and ignore edge
weights. Return vertex indexes in visit order. Do not use a library graph
traversal.

## Complexity Targets
O(V+E) time and O(V) space with adjacency lists.
