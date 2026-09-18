# Kruskal

## How It Works

Kruskal considers logical undirected edges by nondecreasing weight and accepts
an edge only when union-find shows that its endpoints are in different
components. The accepted edges form a minimum spanning forest. Unlike Dijkstra,
there is no source or path accumulation, so negative edge weights are valid.

## Required API

```go
func Kruskal[T any](
    inputGraph graph.UndirectedEdgeGraph[T],
) ([]graph.Edge[T], int64, error)
```

The result is the selected forest and its total weight.

## Contract

Return `ErrNilGraph` for a nil graph and `graph.ErrDirectedGraph` with a nil
forest for a directed graph; propagate errors from `Graph.Edges`. A connected
graph returns its minimum spanning tree; a disconnected graph returns a minimum
spanning forest with one tree per component. Negative edge weights are valid.
The returned edges must not form a cycle, and their weights must sum exactly to
the reported total.

Use `Edges` as the logical-edge source rather than reconstructing undirected
edges from neighbors. Equal-weight tie ordering follows the quick-sort order of
the collected edge slice and is not otherwise specified.

## Key Mapping

Union-find accepts only dense indices `0..count-1`, while graph node keys are
sparse. Each key is mapped to the next unused dense index on first sight; all
`Find`/`Union` calls use mapped indices, never raw keys. Capacity is the graph
node count, not the edge count. Acceptance stops early once `V-1` edges are
selected, since any further edge could only close a cycle.

## Complexity Targets

O(E log E + E alpha(V)) time and O(E+V) space.

## Verification

```sh
just contract algorithms/minimum-spanning-trees/kruskal
go test -tags=contract -run '^$' -bench=Kruskal -benchmem ./src/algorithms/minimum-spanning-trees/kruskal
```

The contract covers a connected minimum spanning tree, a disconnected forest
with a negative weight, and directed-graph rejection. Benchmarks cover a
connected weighted cycle and a disconnected pair forest at 256 and 1,024 nodes;
graph construction occurs before timing.
