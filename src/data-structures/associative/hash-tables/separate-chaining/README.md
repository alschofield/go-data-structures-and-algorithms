# Hash Table

## How It Works
A hash selects a bucket and collisions form linked chains. `Set` preserves the
chosen fixed capacity; `SetResizing` doubles and rehashes buckets before a new
entry would exceed a 0.75 load factor.

## Required API
Generic `type HashTable[K any, V any]` with
`NewHashTable(initialCapacity int, hash func(K) uint, equal func(K, K) bool)
(*HashTable[K, V], error)`,
`Set(K, V) (V, bool)`, `SetResizing(K, V) (V, bool)`, `Get(K) (V, bool)`,
`Remove(K) (V, bool)`, `Contains(K) bool`, `Len() int`, `Cap() int`, and
`IsEmpty() bool`.

## Contract
- `NewHashTable` requires a positive initial capacity and non-nil `hash` and
  `equal` functions. It returns `ErrInvalidCapacity`, `ErrNilHash`, or
  `ErrNilEqual` without creating a table. Standard callers use `10`.
- Both set methods insert a new key or replace an equal key's value while
  retaining the first stored key. Their boolean reports whether a prior value
  was returned.
- `Set` never changes capacity.
- `SetResizing` checks whether adding a new key would exceed a 0.75 load
  factor. If so, it doubles capacity and rehashes every entry with
  `hash(key) % newCapacity` before insertion. A failed allocation preserves
  the table, capacity, and result.
- Absent lookups and removals return `ok=false`, not an error, and do not
  mutate. Different keys with equal hashes
  remain correct. Do not use Go maps.

## Complexity Targets
`Set`, `Get`, `Remove`, and `Contains` are expected O(1) with short chains,
O(n / capacity) as fixed chains grow, and O(n) worst case. `SetResizing` is
expected amortized O(1) and O(n) when resizing. `Len`, `Cap`, and `IsEmpty`
are O(1); space is O(entries + capacity).
