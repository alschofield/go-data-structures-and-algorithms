# Queue

## How It Works

A resizable ring buffer wraps head and tail indexes, so FIFO operations do not
shift values. When full, it doubles and copies values in logical FIFO order.

## Required API

```go
var ErrEmptyQueue error

type Queue[T any] struct

func NewQueue[T any]() *Queue[T]
func (q *Queue[T]) Enqueue(value T) bool
func (q *Queue[T]) Dequeue() (T, error)
func (q *Queue[T]) Peek() (T, error)
func (q *Queue[T]) Len() int
func (q *Queue[T]) IsEmpty() bool
```

## Contract

The zero value is a usable empty queue. `Enqueue` appends and returns true.
`Dequeue` and `Peek` return the oldest value; only `Dequeue` removes it. Empty
reads return the zero value of `T` and `ErrEmptyQueue` without mutation.
Successful reads return nil error. Dequeue clears its removed backing slot before
advancing head so references are not retained by reusable capacity.

`size` is the number of logical values, `head` identifies the oldest value, and
`tail` identifies the next writable slot. Both indexes wrap modulo buffer length;
`size` distinguishes a full buffer from an empty one when head equals tail. After
growth, values occupy the leading contiguous slots, with head zero and tail size.

## Complexity Targets

`Enqueue` is amortized O(1), including O(n) growth copies. `Dequeue`, `Peek`,
`Len`, and `IsEmpty` are O(1). Space is O(n), including retained capacity.

## Verification

```sh
make contract NAME=data-structures/linear/queues/queue
```
