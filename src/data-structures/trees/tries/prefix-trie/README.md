# Prefix Trie

## How It Works

Rune-keyed child edges form shared prefixes. A terminal marker distinguishes a
stored word from a node that is only a prefix.

## Required API

```go
type PrefixTrie struct

func NewPrefixTrie() *PrefixTrie
func (pt *PrefixTrie) Insert(str string) bool
func (pt *PrefixTrie) Contains(str string) bool
func (pt *PrefixTrie) StartsWith(str string) bool
func (pt *PrefixTrie) Remove(str string) bool
func (pt *PrefixTrie) Len() int
```

## Contract

`Insert` returns true only for a newly stored word; duplicate insertion is
idempotent. `Contains` requires a terminal word marker, while `StartsWith`
returns true for every existing prefix, including the empty string. The trie is
rune-aware, so multibyte Unicode characters occupy one edge each.

`Remove` returns false for an empty string, a missing path, or a path that is
not a stored word. A successful removal clears only that terminal marker and
prunes nodes only when no stored word requires them. It preserves longer words
and shared prefixes. `Len` is the number of stored words, not trie nodes.
Internal nodes use `graph.Node[string]`, `Children`, and `IsEndOfWord`.

## Complexity Targets

All key operations are O(m) for rune length m; space is O(total stored runes).

## Verification

```sh
just contract data-structures/trees/tries/prefix-trie
```
