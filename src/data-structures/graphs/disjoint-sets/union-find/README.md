# Union-Find

## How It Works
Parent-pointer trees represent sets; path compression and union by rank keep them nearly flat.

## Required API
`type UnionFind` with `NewUnionFind(elementCount int)`, `Find(int) (int,bool)`, `Union(int,int) (bool,bool)`, `Connected(int,int) (bool,bool)`, and `SetCount() int`.

## Contract
Elements are `[0,n)` and start singleton. Find compresses paths; Union uses rank/size and does not alter rank/count for existing connections. Only representative equality is observable. Do not use a library disjoint-set type.

## Complexity Targets
Find/Union/Connected amortized O(alpha(n)); construction O(n); O(n) space.
