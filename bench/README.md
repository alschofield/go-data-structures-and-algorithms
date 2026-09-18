# Go Benchmarks

Benchmarks use Go's standard `testing.B` runner. They are regression evidence
for one machine and toolchain, not portable speed claims or proof of Big-O.
Never replace a learning implementation with a standard-library container,
`container/heap`, `sort`, `slices`, or a graph algorithm to improve a result.
Go slices are the native dynamic-sequence baseline, not a dynamic-array
exercise, and therefore have no standalone benchmark plan.

## Running And Comparing

```sh
just benchmark data-structures/linear/stacks/stack
go test -tags=contract -run '^$' -bench . -benchmem -count=10 ./src/algorithms/sorting/comparison/quick-sort > before.txt
go test -tags=contract -run '^$' -bench . -benchmem -count=10 ./src/algorithms/sorting/comparison/quick-sort > after.txt
just benchmark-compare before.txt after.txt
```

`benchstat` is the optional comparison tool from `golang.org/x/perf/cmd/benchstat`.
Keep setup outside `b.ResetTimer`, validate results after `b.StopTimer`, call
`b.ReportAllocs`, and use `b.SetBytes` where byte throughput is meaningful.
Use `b.Run` for size and workload subcases. Do not report a result until the
matching `contract` tests pass.

An invalid-input benchmark may measure rejection only after the normal contract
passes, but it is not a substitute for tagged error assertions. Never benchmark
or recover from a panic as an invalid-input outcome.

## Leaf Plans

| Leaf | Benchmark subcases |
| --- | --- |
| stack | push growth; pop; peek |
| queue | enqueue/dequeue wraparound; peek |
| singly-linked-list | front operations; tail operations; middle lookup |
| doubly-linked-list | both-end operations; middle lookup |
| separate-chaining | uniform keys; colliding keys; resizing inserts; lookup/remove |
| binary-search-tree | random insert/find/remove; sorted adversarial insert |
| prefix-trie | insert/contains/remove by key length and shared prefix |
| binary-heap | push; pop; mixed priority workload |
| graph | neighbor delegation for sparse and dense adapters |
| adjacency-list | add edge; has edge; neighbor walk on sparse graph |
| adjacency-matrix | add edge; has edge; neighbor walk on dense graph |
| union-find | find after compression; union; connected |
| linear-search | hit at first/middle/last; miss |
| binary-search | hit; miss; varying sorted input sizes |
| bubble-sort | sorted early exit; random; reverse |
| selection-sort | random; reverse; duplicate-heavy |
| insertion-sort | sorted; nearly sorted; reverse |
| merge-sort | random; duplicate-heavy; uneven sizes |
| quick-sort | random; sorted; reverse; all equal |
| heap-sort | random; reverse; duplicate-heavy |
| counting-sort | compact versus wide valid ranges; invalid-key rejection |
| radix-sort | random values; repeated values; high-byte variation |
| breadth-first-search | sparse/dense graph traversal; disconnected component |
| depth-first-search | sparse/dense graph traversal; deep chain |
| dijkstra | sparse/dense nonnegative weighted graphs; stale frontier entries |
| a-star | zero versus admissible heuristic; reachable versus unreachable goal |
| kruskal | sparse/dense undirected graphs; connected tree versus forest |

Add `BenchmarkXxx` functions beside each leaf's contract tests. Use the plan
above to keep workload construction and validation consistent.
`benchmark_template_test.go` is the minimal reusable `testing.B` shape; copy it
into the named leaf, replace `Skip` with its documented workload, and retain
the `-benchmem` command above. Contract correctness is a prerequisite for every
benchmark, and all 28 current leaf contract suites pass under `-tags=contract`.

## Shortest-Path Results

Run the A* leaf benchmark with:

```sh
go test -tags=contract -run '^$' -bench=AStar -benchmem ./src/algorithms/shortest-paths/a-star
```

It contrasts a zero heuristic, which behaves as Dijkstra, with an exact
remaining-distance estimate on a weighted chain. The comparison demonstrates
frontier prioritization only; both workloads must return the same optimal path.
Record machine-specific results with the Go version, CPU, and command output
when establishing a baseline. Do not compare those values directly to C's
nanosecond results: the harnesses, allocation models, compilers, and runtimes
are different.
