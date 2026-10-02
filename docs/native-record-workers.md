# Shared native workers for read-only open records

An open-row reader can use a native worker shared across record shapes:

```purescript
score :: forall r. { id :: Int, name :: String, active :: Boolean | r } -> Int
```

If every use of the parameter reads one of those known fields, the worker
accepts one native struct containing `active bool`, `id int64`, and `name string`.
A caller with a larger native record projects those fields directly. The source
record is evaluated once, including its extra fields; the caller's record and
its immutable updates retain the full layout. The worker can return a native
scalar or an existing native `Maybe`, `Either`, or `Tuple` layout.

This is a **field projection**, not an interface wrapper or a new representation
for all open records. It copies known primitive fields and needs neither an
accessor chain nor a different function body for every caller's layout.

## Eligibility and fallback

`NativeRecordArgs.candidateToShare` proves eligibility on typed CoreFn before
monomorphization. Every quantified variable must occur among the argument row
tails, and at least one argument must have a projectable open row. Known fields
are primitive scalars; other arguments must be monomorphic. Constraints, unknown
annotations and arbitrary polymorphic fields do not enter this sharing path.
Results must be scalars or supported native ADT layouts; payload representations
remain governed by the existing conversion rules.

Labels must map to distinct usable Go field names. The first occurrence of a
duplicate source label determines its type; fields are then sorted by source
label. Unsupported labels or sanitization collisions retain the ordinary ABI.

Only direct reads of declared fields qualify. Returning the row, aliasing it,
passing it to another function, updating it, or capturing it in a closure rejects
the projection. Lexical shadowing is respected. The transformed TCO body is
checked independently by `workerArguments`, without relying on source usage
annotations. Published worker signatures are also used to emit their bodies
and all direct callers.

## Responsibilities and proof boundaries

- `NativeRecordArgs.Projection` recognizes projectable fields, constructs their
  native argument type and owns the common result-admission policy.
- `NativeRecordArgs.Source` owns the source signature and lexical read-only proof.
  The binding annotation takes precedence, including explicit `Any`; expression
  annotations are used only when it is absent. Leading lambdas must exactly match
  the signature arity and have distinct parameter names. Let initializers and
  case scrutinees are checked before their binders shadow the source parameter.
- `NativeRecordArgs.workerArguments` checks the transformed TCO body separately
  for each argument. It uses emitted local identities (sanitized name and level),
  unwraps `Typed` around a direct read, and rejects captures in closures or deferred
  effects. Missing names or failed proofs use the ordinary argument type.

The source proof controls the specialization filter in `Monomorphization`.
`ModuleBindings` owns publication of the final signature, which `ModuleWorkers`
and direct callers consume. No source-usage annotation authorizes a projection.

## Calling conventions

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

`tools/native-record-args.test.mjs` fixes the source and TCO proof boundaries,
annotation precedence, row/label policy, lexical shadowing and result admission.
`tools/native-record-workers.test.mjs` checks native signatures and reads,
multiple imported layouts, scalar arguments, dynamic/partial applications, full
record preservation, and unsupported/escaping cases. The PureScript fixture
`tests/passing/NativeRecordWorkers.purs` and its companion `Worker` module exercise
the actual TAST pipeline and immutable updates across three layouts.
`NativeRecordReturns` combines the shared projection with an `Either` result
carrying a closed record, including dynamic calls and retained results.

Use the typed PureScript fork in PATH when running `bin/test`. The native
bootstrap discovers that fork itself. Runtime benchmarking, separate from tests,
uses the generated caller/producer pipeline with an unchanged scalar oracle.
The 22 September experiment is recorded in altbak.pub's
`docs/benchmark-results/2026-09-22-native-record-workers.md`.

Projection applies only to proven read-only parameters. Whole-row returns,
arbitrary row updates and escaping or higher-order uses keep the ordinary ABI.
Closed record payloads in supported native results use the existing result
conversion rules. This validation establishes no new JSON/TAST or b8x timing.
