# Error-Handling Contracts

This curriculum uses a final `error` result only when a caller supplies an
invalid argument or invokes an operation in an invalid state. Every other
outcome is represented by the operation's ordinary result.

## Policy

- A successful call returns `nil` as its final error.
- An expected miss or empty read is not an error: use the documented `ok=false`
  or zero/empty result. Examples include a missing search key, missing hash
  key, empty stack or queue, absent trie key, and an unreachable path.
- Invalid indexes, invalid constructor parameters, nil required callbacks,
  malformed graph inputs, unsupported directed input, and rejected numeric
  constraints return an error and leave caller-owned input and receiver state
  unchanged unless the leaf explicitly documents otherwise.
- APIs that can distinguish a normal miss from invalid input place `error`
  last, for example `(T, bool, error)` or `(int, bool, error)`.
- Define exported sentinel errors only when callers can act on a stable failure
  class. Name them `Err` plus the condition, such as `ErrInvalidIndex`,
  `ErrInvalidCapacity`, `ErrNilComparator`, `ErrInvalidVertex`, and
  `ErrNegativeWeight`. Wrap a sentinel only with context and let callers use
  `errors.Is`; tests must not compare error strings.
- Do not add generic nil checks for values of a type parameter: a generic `T`
  is not necessarily nil-capable. Validate only explicitly required callback
  functions and interface values. A typed-nil value stored in a non-nil
  interface is outside this curriculum's nil contract unless a leaf says
  otherwise.
- Panics are inappropriate for ordinary invalid input, absence, capacity,
  index, graph, comparator, or range errors. Panic only for an unrecoverable
  internal invariant breach that cannot be caused by a documented public input;
  curriculum implementations should normally avoid even that.

## Opt-In Verification

Contract tests use `//go:build contract`. They intentionally do not compile
until the learner implements the documented production declarations. Run one
implemented leaf with `make contract NAME=<taxonomy-leaf>`. Default tests and
benchmarks do not require unfinished APIs; benchmark a leaf only after its
contract tests pass.
