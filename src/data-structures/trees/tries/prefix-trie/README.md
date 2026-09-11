# Prefix Trie

## How It Works
Tree edges are characters; root-to-node paths are prefixes and an end marker distinguishes keys from waypoints.

## Required API
`type PrefixTrie` with `NewPrefixTrie()`, `Insert(string) bool`, `Contains(string) bool`, `StartsWith(string) bool`, `Remove(string) bool`, and `Len() int`. Internal trie nodes use `graph.Node[string]` and its `Children` map; graph-unrelated links remain nil.

## Contract
Duplicate insertion is idempotent. Contains matches complete keys, StartsWith accepts empty prefix, and Remove prunes only nodes no key needs. Do not use a library trie/map.

## Complexity Targets
Core operations O(m) for key length independent of key count; O(total stored characters) space.
