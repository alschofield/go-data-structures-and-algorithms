# Stack

## How It Works

An array-backed LIFO collection keeps its top at the final slice element and
grows through Go slice capacity expansion.

## Required API

```go
var ErrEmptyStack error

type Stack[T any] struct

func NewStack[T any]() *Stack[T]
func (s *Stack[T]) Push(value T) bool
func (s *Stack[T]) Pop() (T, error)
func (s *Stack[T]) Peek() (T, error)
func (s *Stack[T]) Len() int
func (s *Stack[T]) IsEmpty() bool
```

## Contract

The zero value is a usable empty stack. `Push` places a value on top and returns
true. `Pop` and `Peek` return the newest value; only `Pop` removes it. Empty
reads return the zero value of `T` and `ErrEmptyStack` without mutation.
Successful reads return nil error. `Pop` zeroes its removed backing slot before
shortening the slice so references are not retained by reusable capacity.

## Complexity Targets

`Push` is amortized O(1). `Pop`, `Peek`, `Len`, and `IsEmpty` are O(1). Space
is O(n), including retained backing-slice capacity.

## Verification

```sh
just contract data-structures/linear/stacks/stack
```
