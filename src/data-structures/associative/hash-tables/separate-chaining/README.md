# Hash Table

## How It Works
A hash selects one of ten fixed buckets and collisions form linked chains.

## Required API
Generic `type HashTable[K any, V any]` with `NewHashTable(hash func(K) uint, equal func(K,K) bool)`, `Set(K,V) (V,bool)`, `Get(K) (V,bool)`, `Remove(K) (V,bool)`, `Contains(K) bool`, `Len() int`, and `IsEmpty() bool`.

## Contract
Set inserts or replaces an equal key's value while retaining the first key. Different keys with equal hashes remain correct. Use exactly ten fixed buckets and separate chaining, not Go maps.

## Complexity Targets
Expected O(1) with short chains, O(n/10) as chains grow, O(n) worst case; Len/IsEmpty O(1); O(entries+10) space.
