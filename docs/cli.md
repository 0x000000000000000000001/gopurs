# Command-line reference

`gopurs` is a backend driver. It reads the typed CoreFn/TAST files already
written under `output/<Module>/corefn.json` and writes Go sources, a runtime and
`go.mod` under `output/`. It does not invoke the PureScript frontend and does
not build the final application binary.

## Synopsis

```text
gopurs [options]
```

Run it from the application workspace containing `output/`. The current
directory is the workspace boundary.

## Options

| Option | Behavior |
| --- | --- |
| `--main Module` | Select the entry module explicitly. Without it, every loaded module exporting `main` is discovered. |
| `--ffi directory` | Add an FFI lookup directory. Paths containing spaces are not supported by the shared argument parser. |
| `--dev` | Skip monomorphization and reduce imported inlining. Keeps incremental builds. |
| `--release` | Select the optimized behavior. This is the default. |
| `--rewrite-limit number` | Set the optimizer rewrite limit; default `10000`. |
| `--force` | Recompute preparation, optimization and emission; publish changed Go files. Also available as `GOPURS_FORCE_BUILD=1`. |

`--dev` and `--release` are mutually exclusive. Passing both is an error.

The shared optimizer parser also recognizes `--output` and `--bundle`, but they
do not change gopurs output behavior in `0.1.0`. They are not part of the
supported interface.

There is no dedicated `--help` or `--version` handler in `0.1.0`. Use this
document and the package version from `package.json`.

## Exit status and errors

- `0`: compilation succeeded.
- `1`: the driver raised an error.

After cleanup, the original error is written once to stderr with the prefix
`[gopurs] error:`. Progress and timing diagnostics also go to stderr.

## Compiler hosts

The launcher selects a host in this order:

1. `GOPURS_RUST=1`: `bin/gopurs-rust`.
2. `GOPURS_JS=1`: `bin/gopurs.js` through Node.
3. `bin/gopurs-native` when it exists.
4. Otherwise the launcher falls back to `bin/gopurs.js` and prints a notice.

`GOPURS_JS=1` and `GOPURS_RUST=1` are mutually exclusive. A missing requested
host is an error; the implicit fallback applies only when no host was requested.

All hosts run the same compiler pipeline and report the same phase names.

## Environment variables

### Supported build controls

| Variable | Range / default | Effect |
| --- | --- | --- |
| `GOPURS_JOBS` | 1–64; compiler default | TAST loading and decoding workers. |
| `GOPURS_PREPARE_JOBS` | 1–8; default 2 | Preparation workers, including transitive specialization. |
| `GOPURS_PBO_JOBS` | 1–64; default 1 | Optimizer workers. |
| `GOPURS_EMIT_JOBS` | 1–64; default 8 | Go emission workers. |
| `GOPURS_PIPELINE` | `0` or `1` | Overlap emission with preparation/optimization. |
| `GOPURS_GO_PACKAGES` | `0` or `1` | Select the package layout instead of the monolithic layout. The native launcher defaults this to `1`. |
| `GOPURS_GO_PACKAGE_OVERLAY` | `0` or `1` | Use the Go package overlay path. |
| `GOPURS_BUILD_MANIFEST` | `0` or `1` | Publish the build manifest used by the package layout. |
| `GOPURS_FORCE_BUILD` | `0` or `1` | Same effect as `--force`. |
| `GOPURS_ALLOC_PROFILE` | `0` or `1` | Enable allocation profiling and force the measured passes. Experimental. |

The `bin/gopurs` launcher raises preparation and PBO workers to `8` on machines
with at least 32 GiB of memory for all three hosts, unless the variables are
set explicitly. TAST loading defaults to eight workers for native Go and Rust,
and one for JavaScript. Set worker counts explicitly for reproducible comparisons.

`GOPURS_GO_PACKAGES`, `GOPURS_GO_PACKAGE_OVERLAY` and
`GOPURS_BUILD_MANIFEST` interact with the incremental cache. Keep the chosen
layout across builds; switching layouts requires a full regeneration.

### Experimental controls

The daemon and resident modes are not part of the `0.1.0` public contract:

- `GOPURS_DAEMON`
- `GOPURS_DAEMON_SERVER`
- `GOPURS_DAEMON_DIR`
- `GOPURS_RESIDENT`
- `GOPURS_RESIDENT_TRIGGER`
- `GOPURS_RESIDENT_STATUS`

The same applies to the Rust host and PGO controls, the native bootstrap and
test-only variables documented in the build guides.

### Caller environment

`GOGC`, `GOMEMLIMIT` and `GOMAXPROCS` are ordinary Go runtime settings. The
native launcher may set defaults for the Go host; it preserves explicit caller
values. They are not compiler flags and are not exported to the generated
application.

## Examples

```bash
# Release build of the explicitly selected entry point
gopurs --main Main

# Development build with a lower optimizer budget
gopurs --dev --rewrite-limit 5000 --main Main

# Recompute everything and publish changed files
gopurs --force --main Main

# Select the JavaScript host
GOPURS_JS=1 gopurs --main Main

# Select the Rust host (must be built)
GOPURS_RUST=1 gopurs --main Main
```

Build the generated application separately:

```bash
cd output
go build -o hello ./main
./hello
```

See [build-cache.md](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cache.md) for cache layout, migration and
concurrency, and [publication-scope.md](publication-scope.md) for the stability
policy.

## Build profiles

The default profile remains `release`. Use `--dev` explicitly to reduce
compilation work. Passing both `--dev` and `--release` is an error;
neither flag overrides the other.

Dev skips heuristic inlining of ordinary imported functions and all
monomorphization: both direct and transitive specialization clones are omitted.
It keeps the generic functions and the types and layouts from the TAST. Local
simplification, explicit inlining directives, dictionary resolution and typed
normalization remain enabled; the optimizer still runs to convergence. The Go
generation passes, ordinary native workers and tail-call optimization remain
enabled.

This profile is experimental. It reduces preparation and code-generation work,
but this repository does not attach a public, reproducible benchmark to it. It
does not establish a faster incremental build or a speedup through final Go
compilation. Generated-program runtime performance was not measured here and
depends on the workload.

Pass the profile to the backend itself, from an application with typed TAST
already generated:

```bash
gopurs --dev --main Main
gopurs --release --main Main
```

When `workspace.backend` in `spago.yaml` already selects gopurs, forward the
option through Spago:

```bash
spago build --backend-args "--dev"
spago build --backend-args "--release"
```

Project-specific build wrappers must forward backend arguments to select a
profile; these options are gopurs options, not additional Spago build flags.

The test runners accept both `-d` and `--dev`:

```bash
./bin/test -d FFIIntegerReturns
./bin/modtest -d prelude
```

Fixtures compare the generated Go against snapshots for the selected profile,
then compile and execute it. `bin/test` defaults `GOPURS_GO_PACKAGES` to `0`,
using the complete legacy snapshot set: `tests/passing-snapshots` for Release
and `tests/passing-snapshots-dev` for Dev. Explicitly setting
`GOPURS_GO_PACKAGES=1` tests the package layout instead, using
`tests/passing-package-snapshots` or `tests/passing-package-snapshots-dev`.
Each invocation checks only its selected layout and profile. `--update-snapshots` (or
`UPDATE_SNAPSHOTS=1`) updates only the selected profile, after successful Go
compilation and execution. Campaign reports record the selected profile.
In sibling projects, `b8x/bin/t -d` builds in Dev before
running tests, and `altbak.pub/bin/go/run -d` selects Dev for the Go backend.
These wrappers retain Release as the default for builds without the option.
For both test runners, `-c` rebuilds the selected compiler host: native by
default, JavaScript with `GOPURS_JS=1`, or Rust with `GOPURS_RUST=1`.
`modtest` streams each library's output as it runs and also preserves its log.
Its private compiler launcher announces Dev mode for every library build.

In Dev, these runners also compile the generated application packages with
`-gcflags='gopurs/output/...=-N -l'`: Go optimization and inlining are disabled
for those packages to shorten compilation. The compiler host rebuilt by `-c`,
the Go standard library and Release builds keep their normal flags. Each Dev
application build logs the flags; generated programs can run more slowly.
Calling `bin/gopurs --dev` alone only generates Go sources. When compiling
those sources manually, use `go build -gcflags='gopurs/output/...=-N -l' ./main`
from the output directory to select the same Go compilation mode.

Generic callbacks construct records and ADTs in their consumer's representation.
This avoids copying a growing list merely to box and unbox loop state, including
when Dev omits specialization. `npm run test:profiles` checks linear allocation
growth and JavaScript/Go parity over 100,000 generic loop steps.

Both profiles retain incremental compilation with persistent disk caches and
require no daemon. Release and the default use `output/.gopurs-cache`; dev uses
`output/.gopurs-cache/profiles/dev`. Their reusable state coexists. They publish
to the same public Go output tree under one shared write lock, so switching
profiles may restore generated files even when the selected profile's modules
are cached. The first build of a profile needs to populate its own cache.
The `dev-v4` policy uses distinct build and module keys, so results from other
experimental dev policies are not reused.

## Timing and worker diagnostics

Every invocation reports phase durations and `[gopurs] backend total: … ms`
on stderr. The PureScript entrypoint measures loading and sorting the TAST,
preparation and monomorphization, runtime output, optimization and emission,
and entrypoint output. The total includes these phases; do not add it to them.
It excludes the preceding `purs` compilation, process startup, and subsequent
`go build`. Failed actions report elapsed time with `(failed)` and rethrow the
original error.

Optimization and emission also report the current module every 100 modules,
starting with the first, so long builds show progress before the phase completes.
The initial `[gopurs] workers: prepare=…, pbo=…, emit=…, pipeline=…` line reports
the effective worker limits and emission pipeline state after configuration is
resolved. Preparation is capped at eight, and PBO/emission at 64. A limit of one
selects the sequential path; emission at one also disables pipeline overlap.

`Gopurs.Host.Metrics` uses a monotonic clock, following altbak's `Bench` approach:
`performance.now()` under Node, `time.Since` in Go and `Instant` in Rust.
No flag is needed; all three compiler builds report the same phases. These are real
elapsed times for the current invocation, not warm-up or repeated benchmarks.

The `bin/gopurs` launcher runs `bin/gopurs-native` when it exists, unless
`GOPURS_JS=1` selects the Node bundle `bin/gopurs.js` or `GOPURS_RUST=1` selects
`bin/gopurs-rust`. Without a native build it falls back to the Node bundle and
prints a notice.
All three hosts report identical phases, so
`GOPURS_JS=1 ../gopurs/bin/gopurs --main Main` in a project with a typed `output`
measures the JavaScript compiler with the same instrumentation as the native one.
`npm run build` refreshes the bundle; `npm run build:native` refreshes the bundle
and Go executable; `npm run build:rust` refreshes the Rust executable.

TAST loading and decoding use eight workers by default in the native Go and Rust compilers;
the JavaScript compiler defaults to one. Set `GOPURS_JOBS` from 1 to 64 to select
a worker count (for example, `GOPURS_JOBS=1 b` in b8x).
Missing or invalid values use the backend's default. Batches preserve input order before
dependency sorting. Native workers run on goroutines; Node overlaps file I/O
but still decodes JSON on its JavaScript thread. Other compiler passes are
unaffected by this setting.

Dated [TAST loading measurements](history/tast-loading-202609.md) are retained
separately from this command reference.

`GOPURS_EMIT_JOBS` controls Go emission separately (1 to 64, defaulting to 8;
set it to 1 for sequential emission). The compiler batches consecutive independent modules, using the
optimizer's effective imports to wait for generated function signatures before
emitting a dependent module. Workers receive immutable metadata snapshots and
publish their signatures in the original order after the batch completes.
Translation is deferred until its Aff worker runs, so native workers execute
the CPU work concurrently. `GOPURS_PBO_JOBS` independently controls optimization
concurrency while preserving sequential directive visibility and publication
order. See [parallel emission](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/parallel-emission.md) for the lifecycle
contract, defaults, measurements and validation.

`GOPURS_PREPARE_JOBS` controls preparation workers, including release transitive
specialization rounds and the final per-module rewrite (compiler default 2,
clamped to 1–8). Each transitive round merges ordered results before starting
the next. The launcher selects eight preparation and PBO workers on machines
with at least 32 GiB for every host, unless those variables are explicitly set.
This choice is independent of `GOGC` and `GOMEMLIMIT`. Smaller machines retain
the compiler defaults (preparation 2, PBO 1); emission defaults to 8 everywhere.
Set the environment variables explicitly for reproducible comparisons. See
[parallel preparation](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/parallel-preparation.md).

For the native Go host, the launcher also selects `GOGC=off` with a soft GC
memory target (`GOMEMLIMIT`): 10 GiB on machines with 32–47 GiB of RAM, and 16 GiB from 48 GiB
upwards. An explicit `GOGC` or `GOMEMLIMIT` preserves the caller's entire GC
policy. `GOMAXPROCS` is left to Go unless the caller sets it. This target is not
a strict RSS ceiling. It governs the compiler process and does not configure
the subsequent application build.
Initial persistent preparation proofs use the same bounded preparation workers,
then join in source order before publishing cache objects. Cache formats and
incremental invalidation rules are unchanged.
