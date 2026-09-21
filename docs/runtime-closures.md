# Closure lifetime and heap escape

`runtime/runtime.go` represents closures through `Value.UnsafePtr`. The captured
environment must remain valid after its creator returns or its goroutine exits,
including when Go grows or shrinks stacks and collects garbage.

`forceEscape` makes the closure visible to Go escape analysis through a global
store guarded by `escapeAlwaysFalse`. This follows the variable-guarded store
used by Go's `internal/abi.Escape`. The guard must remain a variable, never a
constant, and must never be assigned. The store remains visible to the compiler
but does not execute, so closure creation needs no shared write or mutex.

Do not replace this with `runtime.KeepAlive` alone: preserving reachability at
one program point does not itself force the captured environment onto the heap.

Run the focused lifetime and concurrency checks with:

```sh
node --test tools/closure-lifetime.test.mjs
```

The tests use an isolated copy of the actual runtime, exercising returned
closures, partial applications, exited creator goroutines, stack churn and
forced GC, both normally and under Go's race detector. When changing the unsafe
representation or upgrading the Go toolchain, also inspect escape analysis and
validate a generated application.
