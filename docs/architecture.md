# Backend architecture

This guide follows the current [Driver](../src/Gopurs/Driver.purs), entered
through [Main](../src/Main.purs). See the [overview](architecture-overview.md)
for the public pipeline, the [module map](module-map.md) for source locations
and direct dependencies, and [naming](naming.md) for terminology.

## Typed input and pass order

1. The typed PureScript fork emits TAST in `output/<Module>/corefn.json`.
   The cache inventories and verifies these files, then copies inputs needed
   for a build into private staging. PBO decodes the copies and orders modules
   by dependencies. `dataDecls`, `classDecls`, type annotations and `TypeApp`
   instantiations are available to the backend.
2. `Driver` loads directives, including on a cache hit.
   `Driver.Prepare.prepareModules` collects eliminated constructors,
   constructor types, global types and class fields. It adds synthetic data
   declarations for class dictionaries.
3. `Monomorphization.monomorphizeModulesWith` collects instantiations,
   propagates transitive requirements, filters candidates and specializes
   modules. Pointer metadata comes from the enriched modules before
   specialization; enums and global types come from the original TAST.
   `Driver.Cache.plan` resolves and freezes FFI inputs, including absent
   candidates and searched directories.
4. `Driver.Build` computes prepared-module keys. It restores valid PBO
   results or invokes the PBO builder to convert modules to optimized
   `BackendModule` values, using directives and `coreForeignSemantics`.
   Restored and freshly computed contributions become visible at the same
   ordered barriers. **gopurs monomorphization precedes PBO optimization.**
5. `Cache.Modules.emit` attempts verified emission reuse. Where preparation
   is needed, `CodeGen.prepareModule` applies `ThunkFusion`, `FunctionFusion`,
   `ImmediateApplications`, `ClosedDictionaries`, `BorrowedObjects`,
   `DecoderSchemas` and `Ownership`, in that order. External references in
   the rewritten module determine the imported worker signatures covered by
   its emission key, including absent signatures.
6. `Driver.Output.emitPreparedModule` adds frozen module-local FFI signatures
   and calls `CodeGen.translatePreparedWithFunctions`. `ModuleBindings`
   prepares tail-call analysis and native signatures; the emitters produce
   `GoDecl` and `GoExpr` values. `GoImports.collectImports` collects
   dependencies before `Printer.printGoFile` renders the result. Each
   parallel batch reads an immutable signature snapshot; results are merged
   in module order before the next batch uses them.
7. `Driver.Output` writes module source and FFI bridges into staging.
   The driver supplies the embedded runtime and entry points.
   `Driver.Cache.publish` verifies and publishes changed contents and removes
   obsolete owned outputs. Storage creates `go.mod` only when missing and
   treats `go.sum` as external.

`ForAll`, constraints, structural types, row-label order and row tails belong
in representation decisions. Do not reconstruct this information from printed
Go or assume the type erasure of standard CoreFn.

## Metadata provenance

`CodegenState` separates shared, immutable `CodegenMetadata` from the mutable
state of one translation: declarations, fresh identifiers and requested Rebox
pairs.

| Data | Source and availability |
| --- | --- |
| `ctorTypes`, `elidedCtors`, `enumAdts`, `enumCtors` | Original TAST ADTs, before synthetic classes are added. |
| `classDeclsFields` | Original classes; methods and superclasses sorted together by label. |
| `globalTypes` | Original bindings and FFI imports, keyed by qualified PureScript names. |
| `reboxFields` | Constructor and class metadata with their respective identity and precedence rules. |
| `pointerAdtPaths`, `pointerAdtNodes`, `pointerAdtLeaves` | Class-enriched modules before specialization. |
| `PreparedModules.modules` | Modules with source-usage facts invalidated and specialization applied, in loaded order. |
| `mainModules` | Original exports, or an explicit CLI entry module. |
| `globalFunctions` | Native signatures from previously emitted modules, published in canonical batch order. |
| `ffiFunctions` | Recognized FFI signatures for the module currently being emitted. |

`GlobalTypes` prefers a binding annotation, then its expression annotation.
An explicit `Any` stops this search. Its final fallback follows application
results through quantifier/constraint wrappers; it is not general inference.
An annotated FFI import takes precedence over a definition with the same name;
an untyped import does not erase an existing type.

## Specialization and preparation

`Monomorphization` invalidates source-usage identities and proofs before
constructing its global index, even when no specialization is selected.
Copied or substituted bindings cannot retain lexical usage proofs. Backend
and tail-call analyses establish their own facts on the IR they consume.
Source usage also does not prove exclusive ownership or permission to mutate.

The global index retains monomorphic bodies, dictionaries, FFI wrappers and
record readers because they can reveal transitive generic calls. The exact
`Data.Ring.negate` binding is excluded to preserve its dictionary for the
signed-zero intrinsic; its typed calls remain discoverable.

After the PBO fixed point, specialization barriers exclude FFI globals,
recognized foreign forwarders and record readers admitted for sharing.
Eligibility uses original global types: a missing type, `Any`, or a type
without an eligible variable does not justify a copy.
`Monomorphization.ForeignForwarders` requires distinct lambda parameters,
forwarded exactly once and in order to a same-module FFI. Captures, extra
computation, changed arity, reordered arguments and shadowed FFI references
reject the candidate.

`Preparation.runPreparationJobs` executes one ordered round of deferred `Aff`
jobs, in contiguous chunks with concurrency bounded at eight. PBO owns
snapshots, per-call caches, ordered merging and barriers between fixed-point
rounds. See [parallel preparation](parallel-preparation.md) and
[parallel emission](parallel-emission.md).

## Driver and cache boundaries

`Host.Snapshot` copies and re-executes the compiler from private storage so its
cache identity describes the code actually executing. `Main` starts the driver
through `runAff_`; after cleanup, a failure writes its original diagnostic and
exits with status 1. The JS host lets Node drain output; the Go host exits after
its synchronous diagnostic. `Driver.Output` explicitly converts synchronous
FFI preparation exceptions into `Aff` errors so they receive the same cleanup.

`Driver.Config` reads options once. `Driver.Cache.begin` can return early when
input, context, object and output proofs match. On a miss, a bracket owns the
session and staging through publication. Consumed inputs are frozen copies;
live inputs are checked again before publication. Failure preserves the last
committed manifest. Rebox diagnostics make a manifest non-reusable so they
are reproduced on the next computation.

`Cache.Modules` manages prepared, PBO and emission reuse. A compiler/input
identity proof can permit reuse of prepared fingerprints when options and
directives match. Changed TAST requires recomputing relevant preparation,
including specialization providers. PBO observations can validate unchanged
lookup results and absences across a changed conservative prefix; unfinished
predecessors cannot supply finalized contributions.

The persistence protocol, publication ordering and invalidation rules are in
[build-cache.md](build-cache.md). Cache keys must cover transformations and
all metadata they consult, including absence observations.

## Representations and conversions

`GoTypes` chooses native representations from TAST metadata. Closed records
have no row tail; their first occurrence of each label wins and fields are
sorted by label. Open records retain a general representation unless a
separate projection proof admits a worker ABI.

An enum has one or more constructors, all without fields. Empty class
dictionaries are excluded. A pointer ADT has one payload constructor and at
most one fieldless alternative, represented by `nil`; additional fieldless
alternatives need distinct tags. Synthetic class declarations are added after
constructor-elimination metadata, so a one-method class keeps its wrapper.
`ConstructorLayout` resolves constructor identity, generic arguments and
fields together; ADT fields take precedence over class fields.

`GoAst.TypeStructPointer` retains constructor identity, qualified PureScript
type and generic arguments. Substitution reconstructs the structured type;
it does not split printed names. Unknown generic arguments use `Value` where
required by the selected layout. Applied type arguments and declared layout
arity remain distinct inputs to `GoTypes` and `ConstructorLayout`.

`GoConversions.coerceGoExpr` first tries layout-aware conversions: identity,
record projection, native ADT slots and pointers to the same constructor.
Other conversions pass through `Value`. A record projection evaluates its
source once, including extra fields, before converting selected fields.
Boxing a generic pointer converts its payload arguments to `Value`; typed
unboxing reverses that conversion. Array helpers capture their source once;
`[]Value` can share immutable backing storage, while other representations
convert their elements. Integer arrays retain dedicated AST nodes so the
[round-trip optimization](array-roundtrip.md) runs before printing.

`GoConversions.NativeAdts` owns native `Maybe`, `Either` and `Tuple` layouts,
tags and slots. Closed-record payload slots may remain native; other payloads
use `Value`. Composite payload conversions call the ordinary converter.

`GoConversions.Rebox` records directed source/destination pairs, derives
field conversions from `ReboxMetadata`, and emits helpers to a fixed point.
Requests can introduce further requests, including recursive pairs. Helpers
preserve `nil`; identical immutable layouts can share a pointer, while field
conversions allocate a new value. `Rc` does not establish this equivalence.
Missing metadata retains the `ERROR: Rebox missing!` diagnostic. Constructor
metadata takes precedence over class metadata.

## Bindings, captures and tail calls

`ModuleBindings.prepare` analyzes recursive groups before publishing native
signatures. Arity follows adjacent lambdas: computation between lambdas keeps
a boxed function result. A recursive group with non-function values keeps its
boxed getters; a singleton recursive function can receive a module loop.
`ModuleWorkers` emits the exact prepared ABI and its boxed wrapper.

`BindingExprs` owns lexical scope and initialization order. `LocalWorkers`
publishes provisional signatures for recognized recursive loops before their
bodies, then refines result signatures in source order. Recursive values use
pointer cells: an early read fails, while a delayed closure can observe the
value after its complete initialization.

Closure contexts remove inherited loop targets. `GoFunctions` separates
mutable loop slots from source parameters recreated on every iteration, so
closures capture the values from their own iteration. `CallExprs` evaluates
and adapts tail-call arguments before assigning slots and continuing; argument
permutations therefore read the old iteration values.

`Int32Loops` can version local or module loops after proving signed writes
and checking entry values. The fast path uses `int32` slots and widens at
call/return boundaries. Out-of-range entries retain the original `int64`
loop, including zero-iteration results. Arbitrary calls, division and unsigned
shifts do not prove a signed result. See [int32-loops.md](int32-loops.md).

`NativeRecordArgs` separately proves pre-specialization sharing and the final
post-PBO worker ABI. Only read-only field access admits projection; whole-record
uses and delayed captures keep the ordinary argument. Boxed wrappers preserve
dynamic and partial calls. See [native-record-workers.md](native-record-workers.md).

## Rewrite and intrinsic contracts

| Area | Contract and reference |
| --- | --- |
| Function and thunk fusion | Admit recognized producer shapes, preserve evaluation boundaries, reserve collision-free worker names and keep original fallback behavior. [Function fusion](function-fusion.md). |
| Immediate applications | Scope-aware substitution and renaming; no recursive scopes; bounded code growth and preserved evaluation order. `ImmediateApplications` and its `Scope` module own admission. |
| Closed dictionaries | Separate reuse of an existing closed binding from lifting a new one. Effects and recursion reject moving a construction; source order determines winners. `ClosedDictionaries.Admission` and `.Scope` own the proofs. |
| Borrowed JSON objects | Prove all uses read-only before replacing a decoder with a borrowed view. Delayed or escaping payload uses reject admission; runtime helpers retain a fallback. [JSON decoder optimizations](json-decoder-optimizations.md). |
| Decoder schemas | Derive plans from recognized source dictionaries and bodies, not result types alone. Preserve read order, first error and opaque callback behavior. [JSON record plans](json-record-plans.md). |
| Tree ownership | Establish eligible layouts and a fixed point of consuming workers, then rewrite calls only with fresh tree arguments. Source usage counts and `Rc` do not replace ownership proofs. [ADT reuse](adt-reuse.md). |
| Array intrinsics | Recognize exact calls before emitting loops; capture each argument at its original evaluation boundary. [Array round trips](array-roundtrip.md). |
| Either traversals | Keep the first `Left` while executing subsequent callbacks, preserve source iteration order and allocate a result per invocation. [JSON traversals](json-traversals.md). |

## Runtime, FFI and printing

Emitters own type selection, boxing, evaluation order, captures and calls.
`Printer` renders the chosen AST without reading or mutating translation state.
Structured references retain ownership through package qualification; raw Go
fragments carry explicit dependency information. See the
[AST/printer boundary](go-ast-printer.md).

`runtime/runtime.go` is the canonical application runtime. The build embeds it
in `Gopurs.Runtime.Source`; a built compiler does not read that Go file at
startup. `FfiSupport` transports requests to the Go parser, `FfiBridge` adapts
typed calls, and the driver owns diagnostics and cleanup. JS uses WebAssembly;
native Go and Rust hosts use native transports. The
[runtime and FFI contracts](runtime-ffi-contracts.md) describe memory sharing,
primitive semantics, asynchronous work and companion implementations.

## Repositories and consumers

| Repository or component | Responsibility |
| --- | --- |
| Typed PureScript fork | Produces TAST, including layouts, class declarations and instantiated types. |
| `purescript-backend-optimizer` | Decodes the typed format and supplies the shared optimizer, builder and IR. |
| `gopurs` | Specialization, Go rewrites, emission, compiler hosts and cache. |
| `gopurs-*` library ports | Go FFI implementations and library-specific tests. Pins are recorded in `bootstrap-sources.json`. |
| `altbak.pub` | Reference generated-program benchmarks and historical baselines. |
| Go toolchain | Compiles and runs the generated application after gopurs finishes. |

Use [development.md](development.md) for the checkout layout and library ports,
[bootstrap.md](bootstrap.md) for compiler hosts, [testing.md](testing.md) for
validation commands and [benchmark-protocol.md](benchmark-protocol.md) for
measurement boundaries. Foreign companions, runtime embedding, package file
lists and dynamically selected test fixtures remain uses even when they have
no PureScript import edge.
