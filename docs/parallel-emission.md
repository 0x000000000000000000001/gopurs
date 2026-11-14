# Parallel Go emission

`GOPURS_EMIT_JOBS` selects the maximum number of modules emitted together
(default 8; set 1 for sequential emission).
It is independent of `GOPURS_JOBS`, which only controls TAST loading.

The PBO coordinator enqueues finalized modules in canonical order, with either
the sequential or parallel optimizer. By default it can optimize the next
modules while preceding batches wait or emit Go. `GOPURS_PIPELINE=0` disables
this overlap; `GOPURS_EMIT_JOBS=1` also retains immediate sequential emission.

At most one emission batch executes at a time. The pipeline tracks up to
`max 2 (GOPURS_EMIT_JOBS * 2)` launched batches, including those waiting for
their predecessor, plus one partially filled batch. At capacity, the producer
joins the oldest batch before launching another. A full batch, a module
depending on a queued module, or the end of the build flushes the queue. The
dependency check uses optimized imports, including references introduced by
inlining. A new batch starts only after the preceding batch finishes. Its
immutable function-signature snapshot is read at that point; signatures are
merged in source traversal order after all module workers finish. Go translation
runs inside deferred Aff computations on native Aff goroutines.

`Driver.Build` uses `Emission.withEmitter` to own the build lifetime. The scope
flushes and joins the final batch on success. Its `supervise` wraps both the
producer and emission, inside the emitter's `bracket`: all remaining child
fibers are cancelled and joined before emitter cleanup runs. This includes PBO
workers suspended while publishing a result that a failed coordinator will no
longer consume. Errors propagate through `enqueue`, `finish`, or the producer.

The low-level `createPipelinedEmitter` API also exposes `cancel`. On its own,
this method stops new admissions, skips the pending batch, and drains the
launched chain without interrupting it. It does not supervise producer workers;
the compiler uses `withEmitter` for that. Joining Aff children does not join
detached filesystem callbacks whose `nonCanceler` cannot stop the OS operation.

`GOPURS_PBO_JOBS` selects optimizer concurrency (compiler default 1; the launcher
selects 8 on machines with at least 32 GiB for Go, Rust and JS unless overridden).
An explicit Go GC policy does not change this worker selection. The parallel
builder preserves sequential directive visibility through ranked snapshots and
replays attempts whose predecessors are not finalized yet. Original imports
alone are insufficient for scheduling. `Driver.Build` owns the Aff scheduler
and publishes generated signatures in module order.

## Validation

After building the compiler:

```sh
node --test tools/emission.test.mjs
```

The tests cover sequential emission, bounded and partial batches, dependency
barriers, producer overlap, backpressure, ordered signature snapshots, failure
propagation and low-level draining. Managed-scope tests include producer and
emission failures with workers blocked on unconsumed results.

Native checks are in [emission-native_test.go](../tools/emission-native_test.go).
Follow its header instructions to run against a retained bootstrap under the
Go race detector. For output comparisons and performance measurements, use the
[benchmark protocol](benchmark-protocol.md).
