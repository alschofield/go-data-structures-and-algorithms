# Doubly Linked List

## How It Works

Each node links to its predecessor and successor. Held head and tail pointers
make both-end mutations constant time; indexed walks start from the nearer end.

## Required API

```go
var ErrInvalidIndex error

type DoublyLinkedList[T any] struct

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T]
func (ll *DoublyLinkedList[T]) PushFront(value T) bool
func (ll *DoublyLinkedList[T]) PushBack(value T) bool
func (ll *DoublyLinkedList[T]) PopFront() (T, bool)
func (ll *DoublyLinkedList[T]) PopBack() (T, bool)
func (ll *DoublyLinkedList[T]) Get(index int) (T, bool, error)
func (ll *DoublyLinkedList[T]) Insert(index int, value T) (bool, error)
func (ll *DoublyLinkedList[T]) Remove(index int) (T, bool, error)
func (ll *DoublyLinkedList[T]) Len() int
func (ll *DoublyLinkedList[T]) IsEmpty() bool
```

## Contract

The zero value is a usable empty list. Push methods always return true. Empty
pops return the zero value of `T` and `false` without mutation. `Get` and
`Remove` accept `[0, Len())`; `Insert` also accepts `Len()` to append. Invalid
indexes return `ErrInvalidIndex`, a zero/false result, and preserve state.
Successful indexed operations return nil error.

The list uses `graph.Node[T]` through `Next` and `Prev`. An empty list has nil
head and tail. A non-empty list has non-nil ends, `head.Prev == nil`, and
`tail.Next == nil`; adjacent nodes have reciprocal links, and each direction
contains exactly `Len()` nodes.

## Complexity Targets

`PushFront`, `PushBack`, `PopFront`, `PopBack`, `Len`, and `IsEmpty` are O(1).
`Get`, `Insert`, and `Remove` are O(n) worst-case, with at most roughly n/2
links traversed. Space is O(n).

## Verification

```sh
make contract NAME=data-structures/linear/linked/doubly-linked-list
```
