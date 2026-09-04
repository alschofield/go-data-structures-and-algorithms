# Doubly Linked List

## How It Works
Previous and next node links plus held head and tail make both-end operations constant time and indexed walks start near the target.

## Required API
Generic `type DoublyLinkedList[T any]` with `NewDoublyLinkedList[T]()`,
`PushFront`, `PushBack`, `PopFront`, `PopBack`, `Get`, `Insert`, `Remove`,
`Len`, and `IsEmpty`. Empty pops use `(T, bool)`; indexed `Get` and `Remove`
use `(T, bool, error)`, and `Insert` uses `(bool, error)`.

## Contract
Indexes are `[0, Len())`; Insert accepts Len. An invalid index returns
`ErrInvalidIndex`; empty pops are normal `ok=false` results. Keep reciprocal
links correct and clear both ends after final removal. Failures preserve state.
Implement nodes directly.

## Complexity Targets
Both-end push/pop, Len, IsEmpty O(1); indexed operations O(n), at most n/2 steps; O(n) nodes.
