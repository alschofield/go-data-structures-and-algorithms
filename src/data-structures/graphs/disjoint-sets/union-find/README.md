# Union-Find

## How It Works
Parent-pointer trees represent sets; path compression and union by rank keep them nearly flat.

## Required API
`type UnionFind` with `NewUnionFind(elementCount int) (*UnionFind, error)`,
`Find(int) (int,bool,error)`, `Union(int,int) (bool,error)`, `Connected(int,int)
(bool,error)`, and `SetCount() int`.

## Contract
`NewUnionFind` returns `ErrInvalidCapacity` for a negative element count.
Elements are `[0,n)` and start singleton; an out-of-range element returns
`ErrInvalidIndex`. Find compresses paths; Union returns `false, nil` for an
existing connection and does not alter rank/count. Only representative equality
is observable. Each element uses a shared `graph.Node[struct{}]`: `Key` is its
element ID and `Parent` is the disjoint-set parent link. Do not use a library
disjoint-set type.

## Complexity Targets
Find/Union/Connected amortized O(alpha(n)); construction O(n); O(n) space.
