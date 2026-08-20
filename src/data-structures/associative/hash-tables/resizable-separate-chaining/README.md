# Resizable Separate-Chaining Hash Table

## How It Works
Colliding keys chain within buckets; before insertion exceeds load 0.75, bucket capacity doubles and every entry rehashes.

## Required API
Generic `type ResizableHashTable[K any, V any]` with constructor hash/equal functions and `Set`, `Get`, `Remove`, `Contains`, `Len`, `Cap`, and `IsEmpty`; value-returning operations use `(V, bool)`.

## Contract
Start with ten buckets. Retain the first equal key on replacement. Rehash with `hash(key) % newCapacity` before exceeding 0.75; allocation failure preserves table and capacity. Never shrink automatically. Do not use Go maps.

## Complexity Targets
Core operations expected amortized O(1), O(n) worst case; resize O(n) amortized; Len/Cap/IsEmpty O(1); O(buckets+entries) space.
