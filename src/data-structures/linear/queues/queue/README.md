# Queue

## How It Works
A resizable ring buffer uses wrapping head and tail indexes so FIFO operations never shift values.

## Required API
Generic `type Queue[T any]` with `NewQueue[T]()`, `Enqueue(T) bool`, `Dequeue() (T, bool)`, `Peek() (T, bool)`, `Len() int`, and `IsEmpty() bool`.

## Contract
Enqueue adds at the back; Dequeue and Peek return the oldest item, only Dequeue removes it. Empty reads preserve state. Do not use a library queue/container.

## Complexity Targets
Enqueue amortized O(1); Dequeue, Peek, Len, IsEmpty O(1); O(n) contiguous space.
