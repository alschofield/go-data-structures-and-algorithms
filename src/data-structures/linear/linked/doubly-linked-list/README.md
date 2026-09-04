# Doubly Linked List

## How It Works
Previous and next node links plus held head and tail make both-end operations constant time and indexed walks start near the target.

## Required API
Generic `type DoublyLinkedList[T any]` with `NewDoublyLinkedList[T]()`,
`PushFront`, `PushBack`, `PopFront`, `PopBack`, `Get`, `Insert`, `Remove`,
`Len`, and `IsEmpty`. Empty pops use `(T, bool)`; indexed `Get` and `Remove`
use `(T, bool, error)`, and `Insert` uses `(bool, error)`.

## Contract
The zero value, `DoublyLinkedList[T]{}`, is an empty usable list; the constructor
returns the same state. `PushFront` and `PushBack` return `true` after adding a
value. `PopFront` and `PopBack` return the removed value and `true`; on an empty
list they return the zero value of `T` and `false` without mutation.

`Get` and `Remove` accept indexes in `[0, Len())`. `Insert` also accepts
`Len()`, which appends. A negative or out-of-range index returns the exported
`ErrInvalidIndex` sentinel, a zero value/`false` result where applicable, and
does not mutate the list. Successful indexed operations return a nil error.

## Invariants
An empty list has `head == nil`, `tail == nil`, and length zero. A non-empty
list has non-nil ends, `head.prev == nil`, and `tail.next == nil`. For every
adjacent pair, `node.next.prev == node` and `node.prev.next == node`. The
forward and backward chains contain exactly `Len()` nodes and describe the same
values in opposite directions.

## Complexity Targets

| Operation | Time | Why |
| --- | --- | --- |
| `PushFront`, `PushBack`, `PopFront`, `PopBack`, `Len`, `IsEmpty` | O(1) | Held head and tail avoid a walk to either end. |
| `Get`, `Insert`, `Remove` | O(n) worst case, at most about n/2 link steps | Traversal starts at the nearer end. |
| Space | O(n) | One node stores each value and its two links. |

Unlike a singly linked list with only a head, this representation's tail makes
`PushBack` and `PopBack` O(1). A singly linked list must walk from its head to
append; it must also walk to the penultimate node before removing its tail.
The doubly linked list pays for that benefit with a tail pointer and one extra
link per node.

## Tests And Benchmarks

```sh
make contract NAME=data-structures/linear/linked/doubly-linked-list
make test-all
make benchmark NAME=data-structures/linear/linked/doubly-linked-list
```

The tagged contract suite covers empty, one-node, and many-node end operations;
front, back, and middle indexed mutations; invalid-index failure non-mutation;
and forward/backward reciprocal-link invariants after every mutation.

Benchmarks use deterministic 1,024- and 65,536-node lists where traversal is
relevant, package-level sinks, and `-benchmem`. They measure both end round
trips, `Get` near each end and in the middle, and middle insert/remove round
trips. Results are machine- and Go-version-specific; repeat with `-count=10`
and compare with `benchstat` for before/after analysis.

Measured once with Go 1.25.5 on Windows/amd64 (11th Gen Intel Core i9-11900K):

| Operation | Size | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: | ---: |
| Front push/pop | 1 | 40.86 | 24 | 1 |
| Back push/pop | 1 | 38.07 | 24 | 1 |
| Front push/pop | 1,024 | 43.10 | 24 | 1 |
| Back push/pop | 1,024 | 35.30 | 24 | 1 |
| Front push/pop | 65,536 | 43.71 | 24 | 1 |
| Back push/pop | 65,536 | 47.56 | 24 | 1 |
| Get near head | 1,024 | 2.299 | 0 | 0 |
| Get middle | 1,024 | 513.5 | 0 | 0 |
| Get near tail | 1,024 | 1.949 | 0 | 0 |
| Get near head | 65,536 | 2.211 | 0 | 0 |
| Get middle | 65,536 | 54,999 | 0 | 0 |
| Get near tail | 65,536 | 2.467 | 0 | 0 |
| Middle insert/remove | 1,024 | 1,269 | 24 | 1 |
| Middle insert/remove | 65,536 | 102,975 | 24 | 1 |
