# Hash Table

## How It Works

The table uses separate chaining: `hash(key) % Cap()` selects one bucket and
each bucket holds a singly linked collision chain. Equal keys share a single
node. Inserting an equal key replaces only its value, retaining the original
stored key.

`Set` always retains the current bucket count. `SetResizing` doubles the bucket
array and rehashes every stored key before adding a *new* entry that would make
the load factor exceed 0.75. It does not resize at exactly 0.75, and a
replacement does not resize because it does not add an entry.

## Required API
Generic `type HashTable[K any, V any]` with
`NewHashTable(initialCapacity int, hash func(K) uint, equal func(K, K) bool)
(*HashTable[K, V], error)`,
`Set(K, V) (V, bool)`, `SetResizing(K, V) (V, bool)`, `Get(K) (V, bool)`,
`Remove(K) (V, bool)`, `Contains(K) bool`, `Len() int`, `Cap() int`, and
`IsEmpty() bool`.

## Invariants

- `capacity` is positive and `len(items) == capacity`.
- `size` equals the total number of nodes across every chain.
- Each node belongs to exactly one chain at `hash(node.key) % capacity`.
- A chain contains no pair of keys for which `equal` returns true.
- Replacing a value retains the original node and key; removing detaches the
  removed node and decrements `size` exactly once.

## Errors And Results

`NewHashTable` rejects a non-positive capacity with `ErrInvalidCapacity`, a nil
hash with `ErrNilHash`, and a nil equality function with `ErrNilEqual`; every
failure returns a nil table. Standard callers use capacity `10`.

`Set` and `SetResizing` return `(oldValue, true)` for replacements and the zero
value of `V` with `false` for new entries. `Get` and `Remove` return the zero
value and `false` for an absent key. These absence paths, plus `Contains`, do
not mutate the table. Different keys with equal hashes remain independent. Do
not use Go maps.

## Load Factor And Rehashing

The load factor is `Len() / Cap()`. `SetResizing` evaluates the prospective
load as `(Len()+1) / Cap()` only after confirming the key is new. When that
value is greater than `3/4`, it allocates `2 * Cap()` buckets, moves every node
to `hash(key) % newCap`, then inserts the requested key. Rehashing can change
chain order but preserves all key/value associations. `Set` deliberately never
performs this work, even when its collision chains grow long.

## Complexity Targets
| Operation | Expected time | Worst case | Notes |
| --- | --- | --- | --- |
| `Set`, `Get`, `Remove`, `Contains` | O(1) with short chains | O(n) | A pathological hash places all entries in one chain. |
| `SetResizing` | Amortized O(1) | O(n) | A threshold-crossing insertion rehashes every entry. |
| `Len`, `Cap`, `IsEmpty` | O(1) | O(1) | Counters and capacity are stored fields. |
| Space | O(entries + capacity) | O(entries + capacity) | Nodes plus bucket-head slice. |

## Tests And Benchmarks

```sh
make contract NAME=data-structures/associative/hash-table
make test-all
make benchmark NAME=data-structures/associative/hash-table
go test -tags=contract -run '^$' -bench . -benchmem -count=10 ./src/data-structures/associative/hash-table
```

Tagged contracts cover constructor errors, collision chains, replacement-key
identity, fixed-capacity inserts, resize thresholds and rehash preservation,
present and absent queries, every chain removal position, and state preservation
after misses. Deterministic benchmarks use integer keys and both uniform and
constant hashes. They measure `Set`, `Get`, and `Remove`; fixed versus resizing
insertion; and a threshold-crossing rehash workload. Package-level sinks retain
results, and `-benchmem` reports allocation behavior.

Benchmark numbers are machine- and Go-version-specific. Run the commands above
on an otherwise idle system, keep the output with the Go version and hardware,
then repeat with `-count=10` and compare samples through `make benchmark-compare
OLD=before.txt NEW=after.txt`. Treat a regression as a measured comparison, not
a single-run result.

Measured once with Go 1.24 on Windows/amd64 (11th Gen Intel Core i9-11900K):

| Workload | Size | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: | ---: |
| `Set`, uniform hash | n/a | 29.69 | 24 | 1 |
| `Set`, constant hash | n/a | 110,073 | 24 | 1 |
| `Get`, uniform hash | 1,024 | 6.108 | 0 | 0 |
| `Get`, uniform hash | 65,536 | 6.250 | 0 | 0 |
| `Get`, constant hash | 1,024 | 1,158 | 0 | 0 |
| `Get`, constant hash | 65,536 | 116,773 | 0 | 0 |
| `Remove`, uniform hash | n/a | 6.054 | 0 | 0 |
| `Remove`, constant hash | n/a | 150,737 | 0 | 0 |
| Fixed-capacity insertion | n/a | 122,442 | 24 | 1 |
| Resizing insertion | n/a | 42.73 | 46 | 1 |
| Threshold rehash | 1,024 | 7,699 | 18,456 | 2 |
| Threshold rehash | 65,536 | 192,934 | 1,048,600 | 2 |
