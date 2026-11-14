# Build cache

The cache reuses generated Go and compiler artifacts only after checking their
inputs and integrity. JavaScript, Go and Rust hosts use the same Go storage
implementation. A cache hit does not replace frontend compilation or Go's own
dependency and build checks.

## Everyday use

Run from the application directory with its typed PureScript compiler and
Spago on `PATH`. If `spago.yaml` configures gopurs as its backend, `spago build`
invokes it automatically. Otherwise invoke gopurs after the frontend build,
using the same entry and FFI options on each run. See the
[quickstart](quickstart.md).

Caching is enabled by default. To force a complete backend computation:

```sh
GOPURS_FORCE_BUILD=1 spago build
# Or, when invoking the backend directly:
/path/to/gopurs/bin/gopurs --force --main Main
```

Forcing recomputes preparation, optimization and emission, publishes the result,
and preserves timestamps of identical Go files. It does not force the frontend
or Go toolchain to rebuild their own valid artifacts.

Release uses `output/.gopurs-cache`. Dev uses
`output/.gopurs-cache/profiles/dev`; its current semantic key is
`optimization-profile=dev-v5`. Dev skips direct and transitive specialization,
retains typed generic modules, and uses PBO's Dev optimization policy. Ordinary
Go code generation and tail-call handling remain active. Older recognized dev
policy versions share the profile namespace but have distinct semantic keys.
The profiles share the publication lock and public-output ownership journal;
switching profiles restores the selected profile's outputs.

### Go packages and persistent preparation

The native Go launcher enables `GOPURS_GO_PACKAGES=1`,
`GOPURS_GO_PACKAGE_OVERLAY=1` and `GOPURS_BUILD_MANIFEST=1` unless explicitly
set otherwise. Other hosts and direct compiler execution require explicit
selection. Set `GOPURS_GO_PACKAGES=0` for monolithic output.

The package planner partitions generated declarations into Go packages. Its
incremental overlay reuses the published plan only when the required signature,
declaration-edge, import and initialization checks pass. Otherwise the complete
planner recomputes the layout. Output migration uses the ownership journal;
external files, `go.mod`, `go.sum` and the stable lock are preserved.

[`Driver.Prepare.Persistent`](../src/Gopurs/Driver/Prepare/Persistent.purs)
can reuse a committed compact preparation proof for a single admissible literal
edit. It checks compiler/options context, directives, the committed snapshot,
source inventory, interface/body-read constraints and projection integrity.
Required TAST modules are loaded on demand. A missing fact, unsupported edit,
new dependency or corrupt artifact falls back to full preparation. This is
available across ordinary one-shot processes; it does not require a daemon.

Module artifacts are packed into 64 private groups with an index by default.
Public Go files remain individual files. `GOPURS_PACK_PUBLICATION=0` selects
individual private objects. Packing changes storage, not the integrity proof.

### Optional daemon

`GOPURS_DAEMON=1 /path/to/gopurs/bin/gopurs --main Main` sends a build request
to a retained compiler process. The first request starts it; later requests
can reuse prepared/PBO state. There is no filesystem watcher.

The launcher stores `pid`, `log` and `status` in the application's
`.gopurs-daemon/` directory, or `GOPURS_DAEMON_DIR` when set. It restarts the
process when the compiler, arguments or tracked settings change.
`GOPURS_DAEMON=0` bypasses the daemon. Resident state is profile-sensitive;
source equality alone never authorizes reuse across profiles.

### Diagnostics

A global hit finishes before TAST decoding and reports zero optimized/emitted
modules. After an edit, `optimized`/`reused` and `emitted`/`restored` distinguish
computation from restoration. Missing or corrupt cache objects produce misses;
a conflicting external output produces an error.

Pipeline pass times may overlap. Backend timing excludes Spago, startup and
`go build`; measure the complete workflow separately when assessing iteration
time. Follow the [benchmark protocol](benchmark-protocol.md).

## Implementation boundaries

| Component | Responsibility |
| --- | --- |
| [BuildCache.purs](../src/Gopurs/BuildCache.purs) | String-based storage and session FFI |
| [Driver.Cache](../src/Gopurs/Driver/Cache.purs) | Driver adapter, semantic options and session lifetime |
| [tools/build-cache](../tools/build-cache/) | Keys, validation, frozen inputs, locking, objects, packages and publication |
| [Cache.Modules](../src/Gopurs/Cache/Modules.purs) | Typed module reuse decisions and canonical restoration |
| [Cache.Wire](../src/Gopurs/Cache/Wire.purs) | Portable business codecs |
| [Driver.Prepare.Persistent](../src/Gopurs/Driver/Prepare/Persistent.purs) | Durable preparation proofs and projection reuse |

Native Go calls the storage package directly. JavaScript uses
`tools/build-cache.mjs`, compiling the Go helper on first use and caching it by
source contents and platform. Rust links the same package into its Go FFI
support archive. The lock protocol targets Linux and macOS.

Storage messages are UTF-8 JSON with `schema: 1`. The low-level operations are
`snapshot`, `lookup` and `publish`; responses use `snapshot`, `hit`, `miss`,
`published` or `error`. Session operations additionally manage frozen inputs
and driver publication. Protocol errors are JSON responses; process-launch
failures cross the Effect/Aff error boundary.

## Inputs and keys

Keys use SHA-256 over length-prefixed UTF-8 components. The prefix is
`gopurs-build-cache`, `1`, then a domain. Compiler resources, options and
observations have unique logical names and deterministic ordering. Neither
filesystem timestamps nor JSON property order proves validity.

| Observation | Covered data |
| --- | --- |
| `file` | Presence, resolved path and SHA-256 of contents |
| `directory` | Presence, resolved path and immediate entry names/types/link targets |
| `tast` | Exact immediate module/CoreFn inventory, resolved paths and content hashes |

A directory observation is not recursive. The FFI resolver records every
searched directory, higher-priority absent candidate and selected file. Search
priority covers the source companion, `.spago`, `spago.d`, packages/versions,
`--ffi` and local roots. For a root, candidates are `src/A/B.go`, `src/A.B.go`
and `A.B.go`. Adding a higher-priority file must invalidate the prior result.
The recorded recipe is only a hint: a hit still requires a matching verified
manifest.

Semantic options include entries, FFI location, rewrite limit, profile, layout
and worker/pipeline settings. Embedded directives belong to the compiler
identity. Any future external directive file must become an observed input,
including its absence.

### Compiler identity

The executed compiler and all relevant resources are hashed. Native identity
covers the executable; JavaScript includes its bundle, helper resources and
runtime versions. Rebuilt local changes count even without a new Git commit.
Unbuilt source edits do not describe the compiler being executed.

`Host.Snapshot` copies the executable or bundle and resources into a private
directory, then reexecutes that copy before cache lookup. Logical names and
contents determine the identity; temporary paths do not. A replaced installed
binary therefore cannot silently attach its identity to already loaded old code.
Different host binaries invalidate conservatively despite sharing the format.

## Per-module artifacts

`current.json` points to a versioned immutable manifest containing the verified
snapshot, sorted outputs, canonical module order and optional semantic
diagnostics. Each module memo includes prepared, optimizer-environment,
emitter-environment and caller-demand keys, plus these business artifacts:

| Codec | Contents |
| --- | --- |
| `pbo/backend-v2` | Backend body and optional PBO read observations |
| `pbo/implementations-v1` | Analyses and complete external implementations |
| `pbo/directives-v1` | Inline directive contributions |
| `gopurs/workers-v2` | Owned symbols, structural Go types, signatures and arities |

Older backend/worker versions can remain readable for output recovery without
qualifying for typed restoration. Storage integrity checks never replace the
business codecs' constructor, arity and invariant checks.

Prepared keys describe the actual specialized module. Optimizer keys cover
implementations and absences, inlined bodies, directives and analysis settings.
An unchanged provider may be invalidated by new caller specialization demands.
Emission keys cover source metadata, layouts, actual worker signatures and
resolved FFI contents, including absent signatures.

PBO read observations can prove reuse when a conservative predecessor-prefix
key changes. Observations belong to individual attempts; a predecessor not yet
finalized means wait, not absence. Restored contributions and worker signatures
are published in canonical order, never completion order. If every optimizer
memo is reusable, verified backend and worker artifacts can feed emission
without starting the optimizer; consumed contributions still require decoding.

An early emitter lookup is valid only when its input references cover all
references introduced by local transformations. Otherwise lookup occurs after
preparation. Rendering always receives the exact transformed signature view.

### Portable values

The wire grammar uses explicit constructors rather than host layouts:

| Value | Encoding |
| --- | --- |
| Unit | `["unit"]` |
| Boolean | `["boolean", true]` |
| Int | `["int", "-42"]`, canonical signed 64-bit decimal |
| Number | `["number", "8000000000000000"]`, 16 lowercase IEEE-754 hex digits |
| String | `["string", "0061d800"]`, four hex digits per UTF-16 unit |
| Array | `["array", [value, ...]]` |
| Constructor | `["ctor", "StableName", [argument, ...]]` |
| Record | `["record", [["labelUTF16hex", value], ...]]` |

Maps and sets use deterministic PureScript key order and reject duplicates.
Strings preserve lone surrogates; numbers preserve negative zero, infinities
and NaN payloads. Unsupported values fail decoding. Producers emit compact JSON
without trailing newlines. Schema or grammar changes require a new version.

`WithGraph`, `WithTypes` and `TypeRef` preserve sharing through artifact-local
references. Object edges and type dependencies must refer backward; missing,
cyclic or forward references produce misses. Interning uses canonical wire
fragments. Identity memoization is private to one encoding and cannot affect
persistent bytes.

## Frozen inputs and publication

1. The driver checks directives, recipe, blobs and outputs. A verified global
   hit returns before decoding or preparation.
2. A miss creates a private stage and freezes the inventoried TAST bytes.
   Their digests must match the initial snapshot. The driver consumes those
   bytes rather than rereading live files.
3. FFI resolution records observations and freezes the exact consumed bytes.
   Decoding, FFI and emission failures prevent publication.
4. Generated runtime, module, bridge and entry files remain private until
   publication verifies the expected input key under the stable output lock.
5. Publication verifies objects, writes an ownership union journal, replaces
   changed files atomically, synchronizes their directories, removes only
   obsolete owned files, rechecks inputs, and commits the manifest and pointer.
   The journal is then reduced to the final inventory.

A source edit that remains at publication causes failure and a retry request.
An edit-and-revert during computation cannot attribute intermediate bytes to a
different key because only frozen inputs are consumed. Cache stages and PBO
`.purmeta` directories are private to each invocation.

The versioned, hashed `output/.gopurs-ownership.json` journal records ownership
independently of disposable cache objects. Duplicate paths, destinations outside
the output root and symlinks at generated destinations are rejected. Unknown
files are not removed by a blanket `*.go` scan. Modified stale owned files and
different external files at new generated paths cause explicit conflicts.

`go.mod` is created only when absent; an existing module graph and `go.sum` are
preserved. They are not generated-Go output hashes. Deleting `go.mod` causes a
miss, and Go remains responsible for validating dependencies.

The lock is `flock` on `.gopurs-cache/write.lock`; its inode stays stable.
Object/public writers are bounded and joined before deletion or commit.
Directories are synchronized after writes and before advancing the pointer.
A failed or interrupted publication leaves the union journal and previous
pointer for recovery. Publication is coherent for protocol clients, but is not
an atomic rename of the entire output tree: external readers must coordinate
with the build workflow.

## Recovery and maintenance

Missing or invalid cache objects, unknown versions, changed inputs and changed
outputs cause a miss. A missing journal may be recovered from a verified prior
manifest; a corrupt journal is an explicit error. Interrupted writes require
the journal. Preserve it when removing reusable cache state.

Adopting legacy generated output requires a fresh generation with matching paths
and bytes. Unknown or different files are not guessed to be owned. If no trusted
inventory can be recovered, build into a fresh frontend output directory.

Use `--force` for routine recomputation. If manually resetting reusable cache
objects, stop project builds and hold the existing publication lock; preserve
its file and the ownership journal. The journal belongs to its physical output
root and must not be transplanted elsewhere.

Multipackage maintenance removes only unreferenced hash-named regular objects
after rereading a consistent committed manifest and ownership journal under the
lock. It preserves referenced artifacts and leaves private stages/control files
alone. Interrupted publication suspends maintenance; maintenance failure does
not invalidate an otherwise successful build.

Allocation profiling and semantic diagnostic markers force computation rather
than replaying pass timings. In particular, `Rebox missing` marks the manifest
to reproduce the diagnostic on the next build. Cache misses remain observable
through their reasons and reuse counters.

## Validation

After building the required compiler hosts:

```sh
npm run test:cache
npm run test:cache-driver
npm run test:cache-sources
npm run test:cache-context
npm run test:cache-codecs
node --test tools/preparation-persistent.test.mjs
# Storage in isolation:
go -C tools/build-cache test -race -count=1 -v -timeout 90s ./...
```

The suites cover content edits with preserved timestamps, input inventories,
FFI priority and absence, compiler replacement/reexecution, options/profiles,
symlinks, wire limits, corrupt artifacts, stale/external outputs, concurrency,
process death and recovery. Driver tests compare incremental output with fresh
output on identical TAST and execute the resulting programs. Persistent
preparation tests enforce proof invalidation and full-preparation fallbacks.
`GOPURS_TEST_RUST=1` adds Rust where supported;
`GOPURS_TEST_KEEP_WORKSPACE=1` retains diagnostic workspaces. See
[testing](testing.md) for selecting a suitable validation scope.
