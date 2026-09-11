# Singly Linked List

## How It Works
Nodes link forward from a head, making front operations constant time while tail and indexed operations walk the chain.
The list uses the shared `graph.Node[T]` record and its `Next` link; graph-only
links remain unused.

## Required API
Generic `type SinglyLinkedList[T any]` with `NewSinglyLinkedList[T]()`,
`PushFront`, `PushBack`, `PopFront`, `PopBack`, `Get`, `Insert`, `Remove`,
`Len`, and `IsEmpty`. Empty pops use `(T, bool)`; indexed `Get` and `Remove`
use `(T, bool, error)`, and `Insert` uses `(bool, error)`.

## Contract
The zero value, `SinglyLinkedList[T]{}`, is an empty usable list; the constructor
returns the same state. `PushFront` and `PushBack` return `true` after adding a
value. `PopFront` and `PopBack` return the removed value and `true`; on an empty
list they return the zero value of `T` and `false` without mutation.

`Get` and `Remove` accept indexes in `[0, Len())`. `Insert` also accepts
`Len()`, which appends. A negative or out-of-range index returns the exported
`ErrInvalidIndex` sentinel, a zero value/`false` result where applicable, and
does not mutate the list. Successful indexed operations return a nil error.
Removing the final node restores a valid empty state. Implement nodes directly.

## Complexity Targets
`PushFront`, `PopFront`, `Len`, and `IsEmpty` are O(1). `PushBack`, `PopBack`,
`Get`, `Insert`, and `Remove` are O(n) in the worst case. The list consumes O(n)
space for its nodes.

This implementation intentionally stores only a head pointer. That keeps the
representation minimal and makes front operations constant time, but reaching,
adding, or removing the tail requires walking from head. A tail pointer would
make append O(1), but would add representation state and would not make
`PopBack` O(1) in a singly linked list because its predecessor is still needed.

## Benchmarks

```sh
make benchmark NAME=data-structures/linear/linked/singly-linked-list
```

The benchmark uses deterministic size subcases (0 or 1, 1,024, and 65,536 as
applicable), `-benchmem`, package-level sinks, and post-run state checks. Front
operations measure a growing or prefilled list. The tail benchmark pairs
`PushBack` with `PopBack` so every timed iteration starts with the same list
size while exposing both tail traversals. `Get` measures the middle index.
Results are machine- and Go-version-specific; use `-count=10` and `benchstat`
for a meaningful before/after comparison.

Measured once with Go 1.25.5 on Windows/amd64 (11th Gen Intel Core i9-11900K):

| Operation | Size | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: | ---: |
| PushFront | 0 | 56.22 | 16 | 1 |
| PushFront | 1,024 | 52.40 | 16 | 1 |
| PushFront | 65,536 | 44.42 | 16 | 1 |
| PopFront | 0 | 1.887 | 0 | 0 |
| PopFront | 1,024 | 2.041 | 0 | 0 |
| PopFront | 65,536 | 2.107 | 0 | 0 |
| PushBack + PopBack | 1 | 29.65 | 16 | 1 |
| PushBack + PopBack | 1,024 | 2,796 | 16 | 1 |
| PushBack + PopBack | 65,536 | 146,685 | 16 | 1 |
| Get middle | 1 | 1.381 | 0 | 0 |
| Get middle | 1,024 | 592.0 | 0 | 0 |
| Get middle | 65,536 | 42,653 | 0 | 0 |
