# Dynamic Array

## How It Works
Contiguous storage tracks length separately from capacity and grows geometrically when needed.

## Required API
Generic `type DynamicArray[T any]` with `NewDynamicArray[T]()`, `Get(int) (T, bool)`, `Set(int, T) (T, bool)`, `Insert(int, T) bool`, `Remove(int) (T, bool)`, `Len() int`, `Cap() int`, and `IsEmpty() bool`.

## Contract
Indexes are `[0, Len())`; Insert accepts Len. Failed operations preserve contents. Capacity is always at least length and grows geometrically. Implement owned contiguous storage, not a wrapper around a library container.

## Complexity Targets
Get, Set, Len, Cap, IsEmpty O(1); append amortized O(1); other inserts/remove O(n); O(n) contiguous space.
