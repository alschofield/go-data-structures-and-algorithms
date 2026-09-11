# Depth-First Search

## How It Works

An explicit or call-stack frontier explores one branch as far as possible before backtracking.

## Required API

`func DepthFirstSearch[T any](graph graph.Graph[T], source_key int) ([]*graph.Node[T], error)`.

## Contract

Use a visited set of node keys and visit each reachable node once. Return
`ErrNilGraph` for a nil graph and `graph.ErrInvalidKey` for an invalid source
key; propagate errors from `Graph.Neighbors`. Handle cycles, self-loops, and
disconnected graphs without mutation; ignore edge weights. Return nodes in visit
order with their retained value pointers. Do not use a library traversal.

## Complexity Targets

O(V+E) time and O(V) space.
