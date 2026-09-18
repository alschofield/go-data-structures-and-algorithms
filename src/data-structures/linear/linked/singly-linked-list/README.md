# Singly Linked List

## How It Works

Nodes link forward from a held head. Front operations are constant time; tail
and indexed operations walk the chain because the representation has no tail.

## Required API

```go
var ErrInvalidIndex error

type SinglyLinkedList[T any] struct

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T]
func (ll *SinglyLinkedList[T]) PushFront(value T) bool
func (ll *SinglyLinkedList[T]) PushBack(value T) bool
func (ll *SinglyLinkedList[T]) PopFront() (T, bool)
func (ll *SinglyLinkedList[T]) PopBack() (T, bool)
func (ll *SinglyLinkedList[T]) Get(index int) (T, bool, error)
func (ll *SinglyLinkedList[T]) Insert(index int, value T) (bool, error)
func (ll *SinglyLinkedList[T]) Remove(index int) (T, bool, error)
func (ll *SinglyLinkedList[T]) Len() int
func (ll *SinglyLinkedList[T]) IsEmpty() bool
```

## Contract

The zero value is a usable empty list. Push methods always return true. Empty
pops return the zero value of `T` and `false` without mutation. `Get` and
`Remove` accept `[0, Len())`; `Insert` also accepts `Len()` to append. Invalid
indexes return `ErrInvalidIndex`, a zero/false result, and preserve state.
Successful indexed operations return nil error.

The list uses `graph.Node[T]` through `Next`. Its forward chain contains exactly
`Len()` nodes, and removing the final node restores an empty list with a nil
head.

## Complexity Targets

`PushFront`, `PopFront`, `Len`, and `IsEmpty` are O(1). `PushBack`, `PopBack`,
`Get`, `Insert`, and `Remove` are O(n). Space is O(n).

## Verification

```sh
just contract data-structures/linear/linked/singly-linked-list
```
