# Breadth-First Search

## How It Works

A FIFO frontier explores every node at distance k before distance k+1.

## Required API

`func BreadthFirstSearch[T any](graph graph.Graph[T], source_key int) ([]*graph.Node[T], error)`.

## Contract

Mark node keys visited when enqueueing and visit each reachable node once.
Return `ErrNilGraph` for a nil graph and `graph.ErrInvalidKey` for an invalid
source key; propagate errors from `Graph.Neighbors`. Handle cycles, self-loops,
and disconnected graphs, never mutate the graph, and ignore edge weights.
Return nodes in visit order with their retained value pointers. Do not use a
library graph traversal.

## Complexity Targets

O(V+E) time and O(V) space with adjacency lists.
