# Hash Table

## How It Works

Separate chaining stores colliding entries in singly linked bucket chains.
Equal keys share one entry; replacement changes its value but retains its
original stored key.

## Required API

```go
var ErrInvalidCapacity error
var ErrNilHash error
var ErrNilEqual error

type HashTable[K any, V any] struct

func NewHashTable[K any, V any](initialCapacity int, hash func(K) uint, equal func(K, K) bool) (*HashTable[K, V], error)
func (ht *HashTable[K, V]) Set(key K, value V) (V, bool)
func (ht *HashTable[K, V]) SetResizing(key K, value V) (V, bool)
func (ht *HashTable[K, V]) Get(key K) (V, bool)
func (ht *HashTable[K, V]) Remove(key K) (V, bool)
func (ht *HashTable[K, V]) Contains(key K) bool
func (ht *HashTable[K, V]) Cap() int
func (ht *HashTable[K, V]) Len() int
func (ht *HashTable[K, V]) IsEmpty() bool
```

## Contract

Construction rejects a non-positive capacity, nil hash, or nil equality function
with `ErrInvalidCapacity`, `ErrNilHash`, or `ErrNilEqual`, respectively, and
returns no table. `Set` and `SetResizing` return `(oldValue, true)` for a
replacement and `(zeroValue, false)` for a new key. `Get` and `Remove` return
`(zeroValue, false)` for an absent key without mutation; `Contains` is a lookup.

`Set` never changes capacity. `SetResizing` checks whether a new entry would
make the load strictly greater than 0.75; if so, it doubles capacity and rehashes
existing entries before insertion. It does not resize at exactly 0.75 or for a
replacement. Rehashing may reorder collision chains but preserves associations.

## Invariants

Capacity is positive, each entry belongs to `hash(key) % Cap()`, `Len` equals
the total chain-node count, and no chain contains two equal keys. Removing an
entry detaches its node and decrements length once.

## Complexity Targets

`Set`, `Get`, `Remove`, and `Contains` are expected O(1) and worst-case O(n).
`SetResizing` is amortized O(1) with O(n) threshold-crossing rehashes. `Len`,
`Cap`, and `IsEmpty` are O(1); space is O(entries + buckets).

## Verification

```sh
make contract NAME=data-structures/associative/hash-table
```
