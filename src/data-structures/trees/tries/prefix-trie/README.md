# Prefix Trie

## How It Works
Tree edges are characters; root-to-node paths are prefixes and an end marker distinguishes keys from waypoints.

## Required API
`type PrefixTrie` with `NewPrefixTrie()`, `Insert(string) bool`, `Contains(string) bool`, `StartsWith(string) bool`, `Remove(string) bool`, and `Len() int`. Internal trie nodes use `graph.Node[string]` and its `Children` map; graph-unrelated links remain nil.

## Contract
Duplicate insertion is idempotent. Contains matches complete keys, StartsWith accepts empty prefix, and Remove rejects a path that is not a stored word. Removal clears only the terminal marker, preserves longer shared-prefix words, and prunes only nodes no remaining key needs. Trie traversal is rune-aware. Do not use a library trie/map.

## Complexity Targets
Core operations O(m) for key length independent of key count; O(total stored characters) space.

## Verification

```sh
make contract NAME=data-structures/trees/tries/prefix-trie
go test -tags=contract -run '^$' -bench=PrefixTrie -benchmem ./src/data-structures/trees/tries/prefix-trie
```

Benchmarks cover full insertion, exact contains, shared-prefix lookup, and
structural removal at 256 and 1,024 keys. Insert and remove workloads include
fresh trie construction; read workloads prebuild a trie before timing.
