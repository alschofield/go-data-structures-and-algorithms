# Singly Linked List

## How It Works
Nodes link forward from a head, making front operations constant time while tail and indexed operations walk the chain.

## Required API
Generic `type SinglyLinkedList[T any]` with `NewSinglyLinkedList[T]()`, `PushFront`, `PushBack`, `PopFront`, `PopBack`, `Get`, `Insert`, `Remove`, `Len`, and `IsEmpty`; returned items use `(T, bool)`.

## Contract
Indexes are `[0, Len())`; Insert accepts Len. Failures preserve the list. Removing the final node restores a valid empty state. Implement nodes directly.

## Complexity Targets
PushFront, PopFront, Len, IsEmpty O(1); remaining operations O(n); O(n) nodes.
