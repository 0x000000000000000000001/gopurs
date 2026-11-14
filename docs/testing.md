# Testing

Run these commands from the gopurs repository root with Node, Go, Spago and
the typed PureScript frontend available. See [bootstrap](bootstrap.md) for
the toolchain and library checkouts.

`bin/test` and `bin/modtest` use `bin/gopurs`: the native Go compiler when
installed, otherwise the JavaScript bundle. Set `GOPURS_JS=1` to select JS or
`GOPURS_RUST=1` to select the experimental Rust host. `-c` rebuilds the selected
host once before the campaign; it does not clear global caches. After compiler
changes, rebuild the host you intend to exercise. Tests that import generated
JavaScript modules also require an up-to-date `output/` from `npm run build`.

## Choose a targeted check

The commands below select regression suites; they do not assert that a campaign
has already passed on your checkout. Many code-generation suites compile and
execute Go programs, so they require Go even when invoked through Node.

| Changed area | Targeted checks |
| --- | --- |
| Documentation | `npm run check:docs` |
| Bootstrap and process management | `node --test tools/build-native.test.mjs tools/test-runner.test.mjs` |
| CLI and host selection | `npm run test:cli`; `npm run test:launcher` |
| Driver and task lifetime | `node --test tools/emission.test.mjs`; native emission tests below |
| TAST/PBO boundary and specialization | `node --test tools/monomorphization.test.mjs tools/global-types.test.mjs tools/preparation.test.mjs tools/foreign-forwarders.test.mjs` |
| Layouts and instantiated types | `node --test tools/representation-contract.test.mjs tools/mixed-constructor-tags.test.mjs tools/elided-constructor-payloads.test.mjs tools/record-tuple-conversions.test.mjs` |
| Native records and sums | `node --test tools/native-record-args.test.mjs tools/native-record-workers.test.mjs tools/boxed-record-arguments.test.mjs tools/native-sum-results.test.mjs` |
| Rebox conversions | `node --test tools/rebox-generation.test.mjs tools/rebox-metadata.test.mjs tools/struct-pointer-boxing.test.mjs tools/value-array-unboxing.test.mjs` |
| Composite literals and constructors | `node --test tools/composite-expressions.test.mjs tools/nullary-constructor-bindings.test.mjs`; `./bin/test ConstructorReuse NativeArrayReboxing` |
| Bindings, closures and recursion | `node --test tools/binding-contracts.test.mjs tools/recursive-initialization.test.mjs tools/local-native-returns.test.mjs tools/zero-arity-functions.test.mjs`; `./bin/test TCO TCOMutRec CurriedLambdas` |
| Function fusion | `node --test tools/thunk-fusion.test.mjs tools/counted-functions.test.mjs tools/immediate-applications.test.mjs`; `./bin/test ThunkFusion ThunkFusionNewtype CountedFunctions` |
| Int32 loop fast paths | `node --test tools/int32-loops.test.mjs tools/integer-boundaries.test.mjs`; `./bin/test Int32Loops TCO ThunkFusion 2136` |
| Array intrinsics and evaluation order | `node --test tools/array-intrinsics.test.mjs tools/array-intrinsic-evaluation.test.mjs tools/array-safe-index.test.mjs tools/array-unsafe-index.test.mjs` |
| Either traversals | `node --test tools/array-traverse-either.test.mjs tools/object-traverse-either.test.mjs`; `./bin/test NativeTraverseCallback ArrayTraverseEither ObjectTraverseEither ArrayRoundtrip` |
| Tree ownership | `node --test tools/owned-trees.test.mjs`; `./bin/test OwnedTrees RBTree ConstructorReuse` |
| Dictionaries and JSON decoders | `node --test tools/closed-dictionaries.test.mjs tools/borrowed-objects.test.mjs tools/decoder-schemas.test.mjs`; `./bin/test StaticDictionary JsonRecordPlan` |
| Go FFI bridge | `node --test tools/ffi-bridge.test.mjs tools/ffi-generics.test.mjs`; `./bin/test FFIIntegerReturns` |
| Go parser and qualified references | `node --test tools/go-source.test.mjs tools/module-codec.test.mjs tools/emitter-worker-keys.test.mjs`; `go -C tools/ffi-gen test -race -count=1 ./...` |
| FFI errors and WASM transport | `npm run test:ffi` |
| Compiler FFI and runtime embedding | `node --test tools/native-ffi.test.mjs tools/embed-runtime.test.mjs tools/go-imports.test.mjs`; native Go scanner tests below |
| Runtime storage and closure lifetime | `node --test tools/runtime-contracts.test.mjs tools/closure-lifetime.test.mjs tools/apply-arity.test.mjs tools/function-data.test.mjs` |
| Persistent cache and codecs | `npm run test:cache`; `npm run test:cache-codecs`; `go -C tools/build-cache test -race -count=1 ./...` |
| Cache integration and invalidation | `npm run test:cache-driver`; `npm run test:cache-sources`; `npm run test:cache-context` |
| Release/Dev behavior | `npm run test:profiles` |
| Runner selection, isolation and errors | `npm run test:runner` |

Rebuild the FFI parser with `npm run build:ffi` when its sources change; keep
its WASM artifact and JavaScript runtime paired as described in
[development](development.md). Cache integration and CLI suites require the
compiler artifacts specified by those tests; `GOPURS_TEST_RUST=1` adds Rust
coverage where supported.

## Compiler contracts

Keep the targeted suite and an end-to-end fixture when changing an optimization.
The suites test admission and rejection cases as well as the generated Go.

| Contract | What the suites protect |
| --- | --- |
| TAST/PBO boundary | Invalidate source usage facts before collection; preserve useful bodies, type provenance, module order and the distinction between early intrinsic barriers and later exclusions. Pure and `Aff` paths agree. |
| Representations | Original and enriched metadata agree with Go declarations; arities, ordered `TypeApp` instantiations, dictionary fields and imported alternatives remain consistent. |
| Projected record arguments | Source proofs respect lexical scope, TCO identities, row quantifiers and field collisions. Final verification rejects invalid proofs independently for each argument; wrappers preserve extra fields. |
| Closed dictionaries | Cache admission preserves annotations, lexical capture, recursion barriers and effects. Lazy construction and per-call captures remain observable. |
| Borrowed objects | Alias and scope checks distinguish borrowing from ownership; owned paths copy the container, and unknown decoders keep the fallback. |
| Decoder schemas | ABI, proxy and symbol checks bound admission. Generated workers preserve missing/null distinctions, error payloads, ordered choices, duplicate labels and string ownership. |
| Consuming workers | Mutation requires disjoint paths and no observable aliases. Fixed-point rejection propagates through callers and recursive families; persistent tree versions remain intact. |
| Intrinsics and traversals | Exact names and arities control admission. Argument order, callbacks after an initial `Left`, captured values, bounds checks and fresh result storage match the reference behavior. |
| Fusion and immediate application | Producers are forced once; partial applications, effects, failure order, recursion scopes, capture substitution and name reservation remain correct. |
| Bindings and functions | Recursive signatures are available before assignments. TCO argument permutations, per-iteration closures, deferred initialization and native return types preserve evaluation order. |
| Composite expressions | Child contexts, statements and Rebox state remain isolated; arrays and record fields retain their required conversion order. |
| Rebox | Helper dependencies reach a fixed point in both directions without duplicates. Nested payloads, `nil`, phantom pointers, native arrays and single evaluation are preserved. |
| FFI bridge | Parsed signatures match emitted workers and direct-call eligibility. Callbacks, records, deferred effects, foreign variables and generic fallback paths remain valid. |

For driver or scheduling changes, compare generated files on the **same TAST
inputs** across JS, native sequential and native parallel hosts. Include runtime,
FFI bridges and executable entry points. See the [benchmark protocol](benchmark-protocol.md)
for performance measurements; a functional test run alone does not establish a
performance improvement.

### Native bootstrap tests

Retain the generated compiler workspace:

```bash
npm run build:native -- --keep-workspace
GOPURS_NATIVE_OUTPUT=/path/to/bootstrap/output node --test tools/preparation-native.test.mjs
GOPURS_NATIVE_OUTPUT=/path/to/bootstrap/output node --test tools/native-go-code.test.mjs
```

The preparation suite exercises the compiled `Preparation` implementation under
`go test -race`: deferred work, repeated execution, bounded overlap, result order
and sequential modes. Its cold Go compilation timeout is separate from the
30-second test execution timeout. Without `GOPURS_NATIVE_OUTPUT`, the native
parts of these Node suites are skipped.

For emission task lifetime, copy `tools/emission-native_test.go` into the
retained workspace's `output/purescript/`, then run from its `output/` directory:

```bash
go test -race -run '^TestPipelineNative' -count=1 -timeout 30s ./purescript
```

Bootstrap tooling tests use simulated external commands to verify frontend
selection, untyped-input rejection, preservation of the previous binary on
failure, workspace cleanup and signal propagation. A change to real bootstrap
behavior also needs a build with the actual toolchain.

### TAST strings and labels

The typed frontend represents isolated UTF-16 surrogates in `PSString` as arrays
of code units. `TypeLevelString.value` and `Row.fields.label` must accept these
arrays, combine valid pairs and preserve isolated surrogates as WTF-8 in Go.
`StringEdgeCases`, `CompilerHostStrings` and `record-tuple-conversions.test.mjs`
cover these boundaries in gopurs.

With the PBO checkout at the path used by the bootstrap, its own decoder
regressions can run against the current `output/`:

```bash
node ../../purescript-backend-optimizer-gopurs/test/type-table-strings.mjs output
```

For the retained native bootstrap, set `GOPURS_NATIVE_OUTPUT` and run:

```bash
cp ../../purescript-backend-optimizer-gopurs/test/type-table-strings_test.go \
  "$GOPURS_NATIVE_OUTPUT/purescript/type_table_strings_test.go"
go -C "$GOPURS_NATIVE_OUTPUT" test -race ./purescript -run '^TestTypeTablePSString' -count=1 -v
```

The PBO suites `monomorphize-transitive`, `transitive-parallel`,
`monomorphize-callsite`, `monomorphize-cache` and `source-usage` use the same
`node ../../purescript-backend-optimizer-gopurs/test/<suite>.mjs output` pattern.
After editing PBO's `Json.go`, run `python3 bin/json-text/generate.py`, then the
same command with `--check`, from that repository.

## Selection and snapshots

```bash
./bin/test --list
./bin/test TCOMutRec ThunkFusion --list
./bin/test --skip-before ThunkFusion --list
./bin/test FFIIntegerReturns --keep-workspace
./bin/test -d FFIIntegerReturns --keep-workspace
```

With no target, `bin/test` selects all non-excluded fixtures in `tests/passing`.
The exclusion set in [test-selection.mjs](../tools/test-selection.mjs) is
currently empty. Explicit names retain their order; `--skip-before` includes
its target. Unknown names or options fail before building. `--list` does not
build or create a workspace.

Snapshot verification is the default. Each fixture is compiled with the typed
frontend, translated to Go, compared with its selected snapshots, then built
and executed. Missing or changed snapshots fail the run: inspect the diff and
the responsible compiler phase before intentionally updating them.

| Profile | Single-package layout (`GOPURS_GO_PACKAGES=0`) | Package layout (`GOPURS_GO_PACKAGES=1`) |
| --- | --- | --- |
| Release, the default | `tests/passing-snapshots` | `tests/passing-package-snapshots` |
| Dev (`-d`, `--dev`) | `tests/passing-snapshots-dev` | `tests/passing-package-snapshots-dev` |

The fixture runner defaults to `GOPURS_GO_PACKAGES=0`; set it explicitly to `1`
to test the package layout. One invocation covers one layout and one profile.
Package snapshots are available for a subset of fixtures; selecting another
fixture does not silently create its missing reference.

Single-package snapshots use `<Fixture>.go`, plus `<Fixture>_ffi.go` when the
source declares `-- @snapshot-ffi`. To update an intentionally changed snapshot:

```bash
./bin/test FFIIntegerReturns --update-snapshots
./bin/test -d FFIIntegerReturns --update-snapshots
```

Only the selected profile's snapshots are updated, after successful Go build
and execution. `UPDATE_SNAPSHOTS=1` is also supported. Execution must exit zero
and must not print `Fail`. Updates from earlier successful fixtures remain
written if a later fixture fails.

To run the complete fixture selection in both profiles:

```bash
./bin/test --all --keep-going
./bin/test -d --all --keep-going
```

## Isolation, reports and resume

Each fixture has its own temporary sources, Spago configuration, lockfile,
`.spago`, `output` and `tmp`. `TMPDIR`, `TMP` and `TEMP` point into that workspace.
Companion `.go`, `.js` files and fixture directories are copied with the source.
`-- @dependencies: assert prelude effect console` can narrow dependencies;
otherwise the shared list in `bin/pkg` applies. Core library checkouts are
required, and Spago's global package cache is shared.

Logs separate frontend compilation, Go generation, formatting, snapshots,
dependency resolution, Go compilation and execution. Reports and logs are
always kept. Successful generated workspaces are removed unless
`--keep-workspace` is set; failed or interrupted workspaces are retained.
SIGINT/SIGTERM stop the active command and its descendants. Failed frontend
compilations are not automatically retried.

Both runners print `Results: /path/to/results.json`. Reports are updated
atomically around each target, with selection, status, paths, timestamps and
diagnostics. `--keep-going` attempts subsequent targets after an individual
failure and returns nonzero if any target fails.

```bash
./bin/test --resume-failed /path/to/results.json --list
./bin/test --resume-failed /path/to/results.json --keep-going
```

`--resume-failed` selects all unsuccessful targets, including interrupted,
pending or abandoned `running` entries. It checks the campaign type and
repository root, then creates a new report linked by `resumedFrom`. It cannot
be combined with explicit names, `--all` or `--skip-before`. Previous reports
remain intact, and snapshots are still checked strictly.

Fresh fixture workspaces are independent of `-c`. That option rebuilds the
compiler; it does not clear global caches. The backend's own reuse and
invalidation rules are described in [build-cache.md](build-cache.md).

## Library tests

`bin/modtest` selects library checkouts under the repository's parent directory,
or under `GOPURS_LIBRARY_ROOT` when set. Each must have a matching package name
in `spago.yaml` and an executable `bin/test`; backend worktrees are excluded.

```bash
./bin/modtest --all --list
./bin/modtest --skip-before strings --list
./bin/modtest prelude strings
./bin/modtest -d prelude strings
```

The full selection is the default. Names accept the `gopurs-` prefix optionally.
Each target receives a private copy of the library source family, without
checkout metadata, dependencies or generated outputs. Relative configuration
links stay relative. The prebuilt compiler is shared read-only; `-c` rebuilds
it once before the campaign. In Dev, a private launcher forwards `--dev`.

The report, resume and retention options above also apply to library campaigns.
Each target streams stdout/stderr and retains it in `logs/test.log`. See
[development](development.md#libraries) for local library setup.
