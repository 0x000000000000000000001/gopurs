# Command-line reference

`gopurs` turns typed PureScript compiler output into Go source. It is the middle
step of a build: run the PureScript frontend first, then `gopurs`, then
`go build` to create the application executable. The [quickstart](quickstart.md)
walks through all three steps.

The input files are typed abstract syntax trees (TAST), stored in
`output/<Module>/corefn.json`. gopurs writes Go sources, its runtime and `go.mod`
under `output/`.

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
| `--dev` | Use generic functions and less inlining to reduce compiler work. Keeps incremental builds; see [build profiles](#build-profiles). |
| `--release` | Select the optimized behavior. This is the default. |
| `--rewrite-limit number` | Set the optimizer rewrite limit; default `10000`. |
| `--force` | Recompute preparation, optimization and emission; publish changed Go files. Also available as `GOPURS_FORCE_BUILD=1`. |

`--dev` and `--release` are mutually exclusive. Passing both is an error.

The shared optimizer parser also recognizes `--output` and `--bundle`, but they
do not change gopurs output behavior in this release. They are not part of the
supported interface.

There is no dedicated `--help` or `--version` handler in this release. Use this
document and the package version from `package.json`.

## Exit status and errors

- `0`: compilation succeeded.
- `1`: the driver raised an error.

After cleanup, the original error is written once to stderr with the prefix
`[gopurs] error:`. Progress and timing diagnostics also go to stderr.

## Compiler hosts

A compiler host is the executable that runs gopurs. Choosing a host changes how
the compiler runs; all hosts produce Go applications.

The launcher selects a host in this order:

1. `GOPURS_RUST=1`: `bin/gopurs-rust`.
2. `GOPURS_JS=1`: `bin/gopurs.js` through Node.
3. `bin/gopurs-native` when it exists.
4. Otherwise the launcher falls back to `bin/gopurs.js` and prints a notice.

`GOPURS_JS=1` and `GOPURS_RUST=1` are mutually exclusive. A missing requested
host is an error; the implicit fallback applies only when no host was requested.

All hosts run the same compiler pipeline and report the same phase names.

## Environment variables

The defaults are a starting point for ordinary builds. Set worker counts when
tuning resource use or reproducing a benchmark; use the
[timing diagnostics](#timing-and-worker-diagnostics) to see the effective settings.

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
| `GOPURS_ALLOC_PROFILE` | Output file path; unset by default | Write a native Go allocation profile and force the measured passes. Unsupported by JS/Rust. Experimental; unset to disable (`0` is a filename). |

The `bin/gopurs` launcher raises preparation and PBO workers to `8` on machines
with at least 32 GiB of memory for all three hosts, unless the variables are
set explicitly. TAST loading defaults to eight workers for native Go and Rust,
and one for JavaScript. Set worker counts explicitly for reproducible comparisons.

`GOPURS_GO_PACKAGES`, `GOPURS_GO_PACKAGE_OVERLAY` and
`GOPURS_BUILD_MANIFEST` interact with the incremental cache. Keep the chosen
layout across builds; switching layouts requires a full regeneration.

### Experimental controls

The daemon and resident modes are not part of the public compatibility contract:

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

See the [cache guide](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cache.md)
for cache layout and concurrent builds, and the
[compatibility policy](publication-scope.md) for stability commitments.

## Build profiles

The default profile remains `release`. Use `--dev` explicitly to reduce
compilation work. Passing both `--dev` and `--release` is an error;
neither flag overrides the other.

Dev keeps functions generic instead of creating versions for specific types
(monomorphization). It skips both direct and transitive specialization, along
with automatic inlining of ordinary imported functions. The types and layouts
from the TAST remain available.

Local simplification, explicit inlining directives, dictionary resolution and
typed normalization remain enabled; the optimizer still runs to convergence.
Go generation, ordinary native workers and tail-call optimization also remain
enabled.

The optimization policy of this profile is experimental. It reduces preparation
and code-generation work. A faster complete build, including `go build`, needs
measurement on your workload; no public benchmark establishes that gain here.
The generated program's runtime performance also depends on the workload and
was not measured here.

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

### Profiles in test runners

The test runners accept both `-d` and `--dev`:

```bash
./bin/test -d FFIIntegerReturns
./bin/modtest -d prelude
```

Fixtures compare the generated Go against saved snapshots, then compile and
execute it. Each invocation checks one layout and profile:

| Layout | Release snapshots | Dev snapshots |
| --- | --- | --- |
| Legacy (`GOPURS_GO_PACKAGES=0`, the fixture default) | `tests/passing-snapshots` | `tests/passing-snapshots-dev` |
| Packages (`GOPURS_GO_PACKAGES=1`) | `tests/passing-package-snapshots` | `tests/passing-package-snapshots-dev` |

`--update-snapshots` (or `UPDATE_SNAPSHOTS=1`) updates only the selected profile
after successful Go compilation and execution. Reports record that profile.
In sibling projects, `b8x/bin/t -d` and `altbak.pub/bin/go/run -d` also select
dev, with release as their default.

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

`npm run test:profiles` also checks JavaScript/Go parity and linear allocation
growth over 100,000 generic loop steps. Generic callbacks construct records and
algebraic data types in the representation expected by their consumer, avoiding
copies of a growing list when converting loop state.

### Caches when switching profiles

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

Timings use a monotonic clock and cover the current invocation. All three hosts
report the same phases, so the [host selectors](#compiler-hosts) let you compare
them with the same instrumentation. For repeatable measurements, follow the
[benchmark protocol](benchmark-protocol.md).

The [build controls table](#supported-build-controls) lists worker settings and
defaults. `GOPURS_JOBS` affects only TAST loading and decoding; missing or invalid
values use the host's default. Native workers run on goroutines. Node can
overlap file reads, but decodes JSON on its JavaScript thread.

Preparation, optimization and emission have independent worker limits. Their
ordering and resource-lifetime rules are explained in
[parallel preparation](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/parallel-preparation.md)
and [parallel emission](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/parallel-emission.md).

### Native compiler memory defaults

For the native Go host, the launcher also selects `GOGC=off` with a soft GC
memory target (`GOMEMLIMIT`): 10 GiB on machines with 32–47 GiB of RAM, and 16 GiB from 48 GiB
upwards. An explicit `GOGC` or `GOMEMLIMIT` preserves the caller's entire GC
policy. `GOMAXPROCS` is left to Go unless the caller sets it. This target is not
a strict RSS ceiling. It governs the compiler process and does not configure
the subsequent application build.
