# Queue

## How It Works
A resizable ring buffer uses wrapping head and tail indexes so FIFO operations never shift values. The buffer grows geometrically when full; during growth, values are copied from `head` in logical FIFO order into a contiguous new buffer.

## Required API
Generic `type Queue[T any]` with `NewQueue[T]()`, `Enqueue(T) bool`, `Dequeue() (T, error)`, `Peek() (T, error)`, `Len() int`, and `IsEmpty() bool`.

## Contract
`Enqueue` adds a value at the back and returns `true`. `Dequeue` and `Peek` return the oldest value; only `Dequeue` removes it. On an empty queue, both return the zero value of `T` and the exported sentinel `ErrEmptyQueue` directly, without mutation. Successful `Dequeue` and `Peek` calls return a nil error. Do not use a library queue/container.

## Ring-Buffer Invariants

- `size` is the number of logical values and satisfies `0 <= size <= len(items)`.
- `head` identifies the oldest value when `size > 0`.
- `tail` identifies the next writable slot; both indexes wrap with modulo `len(items)`.
- `head == tail` is ambiguous without `size`: the queue may be empty or full.
- After a resize, the FIFO sequence occupies `items[0:size]`, `head` is zero, and `tail` is `size`.

`Dequeue` assigns the zero value of `T` to its removed backing-buffer slot before advancing `head`. This releases retained references so their referents are eligible for garbage collection while the buffer capacity remains reusable.

## Complexity Targets
`Enqueue` is amortized O(1), including occasional O(n) growth copies. `Dequeue`, `Peek`, `Len`, and `IsEmpty` are O(1). The queue uses O(n) contiguous space, including retained backing-buffer capacity.

## Benchmarks

```sh
make benchmark NAME=data-structures/linear/queues/queue
```

The deterministic benchmark reports allocations with `-benchmem`, validates final state, and writes operation results to package-level sinks to prevent compiler elimination. It covers `Enqueue`, `Dequeue`, and `Peek` at sizes 0 or 1, 1,024, and 65,536 as applicable, plus steady-state wrapped dequeue/enqueue and a growth from a full wrapped buffer.

Run the command on the target machine before comparing changes. Record results with the Go version, OS/architecture, and CPU model because benchmark numbers vary by environment.

Measured once with Go 1.25.5 on Windows/amd64 (11th Gen Intel Core i9-11900K @ 3.50GHz):

| Operation | Size | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: | ---: |
| Enqueue | 0 | 15.84 | 31 | 0 |
| Enqueue | 1,024 | 12.19 | 22 | 0 |
| Enqueue | 65,536 | 13.20 | 25 | 0 |
| Dequeue | 0 | 6.049 | 0 | 0 |
| Dequeue | 1,024 | 5.520 | 0 | 0 |
| Dequeue | 65,536 | 5.032 | 0 | 0 |
| Peek | 1 | 0.7731 | 0 | 0 |
| Peek | 1,024 | 0.6958 | 0 | 0 |
| Peek | 65,536 | 0.9313 | 0 | 0 |
| Wraparound | 2 | 8.133 | 0 | 0 |
| Grow wrapped | 1,024 | 7,081 | 16,384 | 1 |
