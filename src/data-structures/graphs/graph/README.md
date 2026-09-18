# Graph

## How It Works

`Graph` is the read-only weighted-graph boundary shared by graph algorithms.
Concrete representations retain their storage and expose nodes through stable keys.

## Required API

```go
var ErrInvalidKey error
var ErrDirectedGraph error

type Node[T any] struct {
    Key int; Value *T
    Occurrences int; IsEndOfWord bool; Rank int
    Next, Prev, Left, Right, Parent *Node[T]
    Children map[rune]*Node[T]
    Edges []Edge[T]
}

type Edge[T any] struct { From, To *Node[T]; Weight int64 }

type Graph[T any] interface {
    Directed() bool
    NodeCount() int
    NodeByKey(key int) (*Node[T], bool)
    Neighbors(key int, visit func(*Node[T], int64) bool) (bool, error)
}

type UndirectedEdgeGraph[T any] interface {
    Graph[T]
    Edges(visit func(Edge[T]) bool) (bool, error)
}
```

## Contract

`Node.Key` is a stable identity, not a required array index; removed nodes may
leave gaps. `Node.Value` points to the value retained by its owning structure.
Structures use only their applicable node links and leave unrelated fields nil.

`NodeByKey` returns `ok=false` when absent. `Neighbors` returns
`ErrInvalidKey` for an invalid key, visits outgoing weighted edges in a
deterministic representation order, and returns `false, nil` only when the
visitor requests an early stop. `Edges` likewise stops early; an undirected
view reports each logical edge once in deterministic order. Directed graph
representations return `ErrDirectedGraph` from `Edges`.

## Complexity Targets

The interfaces impose no storage or operation complexity. Each concrete
representation documents its own targets.

## Verification

```sh
just contract data-structures/graphs/graph
```
