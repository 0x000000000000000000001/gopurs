# Shared native workers for read-only open records

An open-row reader can use a native worker shared across record shapes:

```purescript
score :: forall r. { id :: Int, name :: String, active :: Boolean | r } -> Int
```

If every use of the parameter reads one of those known fields, the worker
accepts one native struct containing `active bool`, `id int64`, and `name string`.
A caller with a larger native record projects those fields directly. The source
record is evaluated once, including its extra fields; the caller's record and
its immutable updates retain the full layout. The worker returns a native scalar.

This is a **field projection**, not an interface wrapper or a new representation
for all open records. It copies known primitive fields and needs neither an
accessor chain nor a different function body for every caller's layout.

## Eligibility and fallback

`NativeRecordArgs.candidateToShare` proves eligibility on typed CoreFn before
monomorphization. Its accepted type variables occur only as open row tails;
known fields and the result are primitive scalar types. Other arguments must
be monomorphic. Constraints, unknown annotations, arbitrary polymorphic fields,
and non-scalar results do not enter this sharing path. Labels must also map to
distinct usable Go field names; unsupported labels retain the ordinary ABI.

Only direct reads of declared fields qualify. Returning the row, aliasing it,
passing it to another function, updating it, or capturing it in a closure rejects
the projection. Lexical shadowing is respected. The transformed TCO body is
checked independently by `workerArguments`, without relying on source usage
annotations. Published worker signatures are also used to emit their bodies
and all direct callers.

The existing `Get_*` wrapper remains the public boxed-function adapter. Dynamic
and partial applications can still receive a full `Value` record; the wrapper
extracts the worker's fields. Unknown or escaping uses keep the ordinary ABI.
Native record-to-record coercion also avoids a needless box/unbox when the
destination fields are all present in the source.

Sharing is selected before PBO. If a source candidate no longer qualifies after
PBO, the final check keeps it correct with the ordinary ABI, but the removed
specializations are not restored. That can affect performance. Likewise, sharing
can lose an inlining opportunity; comparisons must retain the ordinary optimized
baseline, including its inlined specializations.

## Validation

`tools/native-record-workers.test.mjs` checks native signatures and reads,
multiple imported layouts, scalar arguments, dynamic/partial applications, full
record preservation, and unsupported/escaping cases. The PureScript fixture
`tests/passing/NativeRecordWorkers.purs` and its companion `Worker` module exercise
the actual TAST pipeline and immutable updates across three layouts.

Use the typed PureScript fork in PATH when running `bin/test`. The native
bootstrap discovers that fork itself. Runtime benchmarking, separate from tests,
uses the generated caller/producer pipeline with an unchanged scalar oracle.
The 22 September experiment is recorded in altbak.pub's
`docs/benchmark-results/2026-09-22-native-record-workers.md`.

This is the first generated proof toward a native shared calling convention.
It does not yet cover records returned by decoders, Either/Maybe payloads,
arbitrary row updates, or general higher-order dictionary methods. It does not
establish a new JSON/TAST or b8x timing.
