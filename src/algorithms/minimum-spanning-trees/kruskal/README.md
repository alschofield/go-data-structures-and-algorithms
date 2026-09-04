# Kruskal

## How It Works

Kruskal collects each undirected edge once, considers edges in nondecreasing
weight order, and uses union-find to accept only edges connecting two distinct
components. The result is a minimum spanning forest.

## Required API

`type Edge struct { From, To int; Weight int64 }` and
`func KruskalMinimumSpanningForest(graph graph.Graph) ([]Edge, int64, bool)`.

## Contract

- Reject a directed graph without partial output.
- Consider every logical undirected edge exactly once, even though each graph
  representation stores it in both directions.
- Return a minimum spanning forest for disconnected input, with no cycles and
  its exact summed weight. Negative weights are valid.
- Empty and one-vertex graphs succeed with an empty edge slice and zero total
  weight. Resolve equal-weight ties deterministically by `(From, To)` after
  normalizing endpoints to `From < To`.
- Collect and sort edges and implement disjoint-set work directly; do not use a
  library graph algorithm, `sort`, or a library union-find substitute.

## Complexity Targets

`O(E log E + E alpha(V))` time and `O(E + V)` space.

## Verification

```sh
make contract NAME=algorithms/minimum-spanning-trees/kruskal
make benchmark NAME=algorithms/minimum-spanning-trees/kruskal
```
