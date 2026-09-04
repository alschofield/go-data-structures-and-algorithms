# Stack

## How It Works
An array-backed LIFO collection keeps its top at length minus one and grows geometrically.

## Required API
Generic `type Stack[T any]` with `NewStack[T]()`, `Push(T) bool`, `Pop() (T, error)`, `Peek() (T, error)`, `Len() int`, and `IsEmpty() bool`.

## Contract
`Push` adds a value at the top and returns `true`. `Pop` and `Peek` return the newest value; only `Pop` removes it. On an empty stack, both return the zero value of `T` and the exported sentinel `ErrEmptyStack` directly, without changing the stack. A successful `Pop` or `Peek` returns a nil error.

`Pop` assigns the zero value of `T` to the removed backing-slice slot before shortening the slice. This releases any reference retained by that slot so its referent is eligible for garbage collection; the backing array capacity remains available for later `Push` calls. Do not use a library stack/container.

## Complexity Targets
`Push` is amortized O(1), including occasional slice growth. `Pop`, `Peek`, `Len`, and `IsEmpty` are O(1). The stack uses O(n) contiguous space, including retained backing-slice capacity.

## Benchmarks

```sh
make benchmark NAME=data-structures/linear/stacks/stack
```

The benchmark covers deterministic `Push`, `Pop`, and `Peek` subcases at sizes 0 or 1, 1,024, and 65,536 as applicable. It reports allocations with `-benchmem`, validates the final state, and stores operation results in package-level sinks to prevent compiler elimination.

Measured once with Go 1.25.5 on Windows/amd64 (11th Gen Intel Core i9-11900K):

| Operation | Size | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: | ---: |
| Push | 0 | 5.632 | 40 | 0 |
| Push | 1,024 | 4.330 | 43 | 0 |
| Push | 65,536 | 5.760 | 44 | 0 |
| Pop | 0 | 1.489 | 0 | 0 |
| Pop | 1,024 | 1.732 | 0 | 0 |
| Pop | 65,536 | 1.566 | 0 | 0 |
| Peek | 1 | 0.6920 | 0 | 0 |
| Peek | 1,024 | 0.7067 | 0 | 0 |
| Peek | 65,536 | 0.8199 | 0 | 0 |
