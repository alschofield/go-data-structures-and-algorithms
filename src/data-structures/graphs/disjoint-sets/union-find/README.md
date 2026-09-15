# Union-Find

## How It Works

Parent-pointer trees represent disjoint sets. Union by rank keeps trees shallow
and `Find` applies path halving while walking toward a representative.

## Required API

```go
var ErrInvalidCapacity error
var ErrInvalidIndex error

type UnionFind struct

func NewUnionFind(count int) (*UnionFind, error)
func (uf *UnionFind) Find(index int) (int, bool, error)
func (uf *UnionFind) Union(thingOne, thingTwo int) (bool, error)
func (uf *UnionFind) Connected(thingOne, thingTwo int) (bool, error)
func (uf *UnionFind) SetCount() int
```

## Contract

Construction rejects a negative count with `ErrInvalidCapacity`; elements are
`[0, count)` and initially form singleton sets. Invalid elements return
`ErrInvalidIndex`. A successful `Find` returns a representative key and
`ok=true`. `Union` returns `true` only when it merges distinct sets; an already
connected pair returns `false, nil` and preserves the set count. `Connected`
compares representatives. Each element is a `graph.Node[struct{}]` whose `Key`
is its element ID, `Parent` is its set parent, and `Rank` is meaningful only for
roots.

## Complexity Targets

Construction is O(n); `Find`, `Union`, and `Connected` are amortized
O(alpha(n)); space is O(n).

## Verification

```sh
make contract NAME=data-structures/graphs/disjoint-sets/union-find
```
