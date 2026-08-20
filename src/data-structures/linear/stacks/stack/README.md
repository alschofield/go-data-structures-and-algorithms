# Stack

## How It Works
An array-backed LIFO collection keeps its top at length minus one and grows geometrically.

## Required API
Generic `type Stack[T any]` with `NewStack[T]()`, `Push(T) bool`, `Pop() (T, bool)`, `Peek() (T, bool)`, `Len() int`, and `IsEmpty() bool`.

## Contract
Push adds at the top; Pop and Peek return the newest item, only Pop removes it. Empty reads return `ok=false` without mutation. Do not use a library stack/container.

## Complexity Targets
Push amortized O(1); Pop, Peek, Len, IsEmpty O(1); O(n) contiguous space.
