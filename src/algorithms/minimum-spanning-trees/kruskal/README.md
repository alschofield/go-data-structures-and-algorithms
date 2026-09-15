# Kruskal

## How It Works

Kruskal considers logical undirected edges by nondecreasing weight and accepts
an edge only when union-find shows that its endpoints are in different
components. The accepted edges form a minimum spanning forest.

## Required API

```go
func KruskalMinimumSpanningForest[T any](
    inputGraph graph.UndirectedEdgeGraph[T],
) ([]graph.Edge[T], int64, error)
```

The result is the selected forest and its total weight.

## Contract

The tagged contract requires a connected graph to return its minimum spanning
tree, and a disconnected graph to return a minimum spanning forest. Negative
edge weights are valid. A directed graph returns `(nil, 0, ErrDirectedGraph)`.
The returned edges must not form a cycle, and their weights must sum exactly to
the reported total.

Use `Edges` as the logical-edge source rather than reconstructing undirected
edges from neighbors. The current package has no implementation; behavior not
exercised by the contract, including nil input and equal-weight tie ordering,
is not yet specified.

## Complexity Targets

O(E log E + E alpha(V)) time and O(E+V) space.

## Verification

```sh
make contract NAME=algorithms/minimum-spanning-trees/kruskal
```
