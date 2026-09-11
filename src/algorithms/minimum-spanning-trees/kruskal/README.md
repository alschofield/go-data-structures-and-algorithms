# Kruskal

## How It Works

Kruskal collects each undirected edge once, considers edges in nondecreasing weight order, and uses union-find to accept only edges connecting two distinct components. The result is a minimum spanning forest.

## Required API

`func KruskalMinimumSpanningForest[T any](graph graph.UndirectedEdgeGraph[T]) ([]graph.Edge[T], int64, error)`.

## Contract

- Return `ErrNilGraph` for a nil graph and `ErrDirectedGraph` for a directed graph without partial output.
- Consume each logical undirected edge exactly once through `Edges`; never reconstruct edges from mirrored neighbor storage.
- Return a minimum spanning forest for disconnected input, with no cycles and its exact summed weight. Negative weights are valid.
- Empty and one-node graphs succeed with an empty edge slice and zero total weight. Resolve equal-weight ties deterministically by normalized `(From.Key, To.Key)`.
- Collect and sort edges and implement disjoint-set work directly; do not use a library graph algorithm, `sort`, or a library union-find substitute.

## Complexity Targets

`O(E log E + E alpha(V))` time and `O(E + V)` space.
