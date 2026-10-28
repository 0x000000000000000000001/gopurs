# 🐹 gopurs

<img height="160" alt="gopurs" src="https://github.com/user-attachments/assets/b013e7c3-fac6-4ee8-9d4c-f39ac8c2c921" />

An experimental **PureScript-to-Go backend**. The compiler is written in
PureScript, with JavaScript, native Go and native Rust compiler builds, a shared Go parser
for FFI signatures, a Go runtime, and JavaScript build tooling.

`gopurs` consumes the enriched TAST (`tcorefn`) produced by our
[PureScript fork](https://github.com/0x000000000000000000001/purescript).
It uses preserved types and partial monomorphization to choose native Go
representations, with a tagged `Value` representation where needed. Go FFI
bridges adapt native function signatures, and the Go library forks provide
implementations of effects and `Aff` using goroutines.

The project builds on [Arista's purescript-backend-optimizer](https://github.com/aristanetworks/purescript-backend-optimizer)
and draws inspiration from Andy Arvanitis's
[purescript-native](https://github.com/andyarvanitis/purescript-native).
The [development log](https://discourse.purescript.org/t/leveraging-a-blazing-fast-runtime-a-new-go-backend-for-purescript/5841)
records the earlier design work.

## Features

- **Type-guided Go generation.** TAST expression types, ADT layouts, and type-class declarations guide native records, constructors, dictionaries, and specialized functions. Partial monomorphization reduces generic representations where supported.
- **Optimization before code generation.** The backend combines the shared optimizer with Go-specific handling of function calls, tail recursion, records, arrays, and effect thunks.
- **Native FFI bridges.** The same Go parser reads foreign declarations through WebAssembly in the JS compiler and direct calls in the native compiler; generated wrappers adapt supported Go signatures to PureScript calls.
- **Go runtime and library ports.** Generated programs include the runtime. The Go `Aff` port uses goroutines, and generated entrypoints wait for work registered with the runtime event loop.

## Documentation

- [Quickstart](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/quickstart.md): compile and run a small application.
- [CLI reference](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/cli.md): options, hosts, environment variables and exit status.
- [Architecture overview](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/architecture-overview.md): input format, pipeline and generated output.
- [Known limitations](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/limitations.md) and [publication scope](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/publication-scope.md).
- [Migration and compatibility](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/migration.md).
- [Troubleshooting and FAQ](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/troubleshooting.md).
- [Benchmark protocol](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/benchmark-protocol.md).
- [Bootstrap chain](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/bootstrap.md).
- [Nix installation and development](docs/nix.md).
- [Release procedure](docs/release.md).
- [Release qualification, 27 October 2026](docs/quality-0.1.0.md): 802 Release/Dev fixture checks, initial three-platform CI, completed cold/warm benchmarks and remaining publication gates.
- [Contributing](CONTRIBUTING.md), [security policy](SECURITY.md), [changelog](CHANGELOG.md) and [code of conduct](CODE_OF_CONDUCT.md).

## Benchmarks

The [altbak README](https://github.com/0x000000000000000000001/altbak.pub#go)
contains the reference results, workloads, and benchmark context. Compare
changes against those baselines; results depend on the workload and toolchain.
The sequential core campaign does not establish multicore scaling.
The [24 October cold-backend campaign](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cold-20261024.md) records
three alternating pairs, incremental checks and the experiments not retained.
Its empty-backend-cache protocol is separate from the historical altbak timings.
See the [benchmark protocol](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/benchmark-protocol.md)
for measurement boundaries and reporting requirements.

## Getting started

### Prerequisites

- Node.js and npm to build the backend and run its JavaScript version. The native executable runs without Node. Node.js `24.8.0` is the supported reference and is declared in `package.json` and `.nvmrc`.
- A TAST-capable `purs` from the compiler fork for applications, plus Spago with YAML configuration support.
- Go to compile generated applications. Their `go.mod` declares Go 1.22; the walkthrough was checked with Go `1.27.0`, without establishing the minimum working version across library ports.
- Git and Bash for the checkout and helper scripts.

Go `1.27.0` is specifically required to rebuild the FFI parser. Ordinary backend builds use its checked-in WASM and JavaScript runtime.

### Build the backend

The optimizer fork and the `st`/`unsafe-coerce` overrides are pinned as git
dependencies in [spago.yaml](https://github.com/0x000000000000000000001/gopurs/blob/main/spago.yaml); Spago downloads them on first build,
so no manual optimizer checkout is required. The typed compiler fork remains a
separate prerequisite for compiling applications and for the optional native
bootstrap. `./bin/setup` installs the library checkouts used by the fixture and
native paths; it does not install the typed compiler.

```text
workspace/
├── purescript/        # TAST compiler fork, if built locally
├── gopurs/
│   ├── gopurs/        # this repository
│   └── gopurs-*/      # optional library checkouts from ./bin/setup
└── hello/             # example application below
```

```bash
# From an empty workspace directory:
mkdir gopurs
cd gopurs
git clone https://github.com/0x000000000000000000001/gopurs.git
cd gopurs
npm ci
```

`npm ci` installs the locked npm dependencies and its `prepare` hook runs
`npm run build`, producing the JavaScript compiler. `./bin/setup --core` is
optional: it installs the library checkouts used by tests and the native
bootstrap. After later edits, choose the build matching the launcher you use:

```bash
npm run build          # refresh bin/gopurs.js (JavaScript host)
npm run build:native   # optional: bootstrap bin/gopurs-native
npm run build:rust     # optional: bootstrap bin/gopurs-rust; select with GOPURS_RUST=1
```

`bin/gopurs` uses `bin/gopurs-native` when it exists. Otherwise it falls back to
the JavaScript compiler and prints a notice, so a clean `npm ci` checkout is
usable without running the native bootstrap. See the
[bootstrap chain](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/bootstrap.md) for the pinned inputs, artifacts and
checksums.

The incremental Go package layout is selected with `GOPURS_GO_PACKAGES=1` when
building an application. Keep that setting across builds; `GOPURS_GO_PACKAGES=0`
selects the monolithic layout. See the [build and migration guide](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cache.md#sélectionner-le-découpage-go-g06)
and the [G06 measurements](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-go-measurements.md) for the delivered
1/8192 partition, compiler identities, measured cycle gains and backend costs.

`npm run build:rust` profiles the Rust host by default: it trains on the
compiler's own Go-runtime modules and rebuilds with the merged profile. Set
`GOPURS_RUST_PGO=0` for the plain O3 ThinLTO binary, and
`GOPURS_KEEP_WORKSPACE=1` to retain logs, raw profiles and `pgo-profile.json`
(see [docs/pgo.md](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/pgo.md)).

### Nix environment

A [Nix flake](https://github.com/0x000000000000000000001/gopurs/blob/main/flake.nix) builds the JavaScript backend and the pinned TAST
frontend. It provides packages, executable apps, application checks and a
development shell for Linux amd64/arm64 and macOS arm64:

```bash
nix build .#gopurs
nix build .#purs-tast
nix develop
nix flake check
```

The development shell provides the typed `purs` fork, Go **1.27.0**, Node.js
**24.8.0** and Spago **1.0.3**. The backend's own build uses its separately
pinned upstream PureScript compiler and Spago; application builds use the TAST
frontend. The Nix package prebuilds its Go helpers so it can run from the
immutable Nix store. See [the Nix guide](docs/nix.md) for application commands
and [the qualification record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/quality-0.1.0.md) for validations actually
completed on each platform.

### Choose the library checkouts

```bash
./bin/setup --core --list  # 22 core libraries + lazy, ordered-collections, random
./bin/setup --core         # default; install missing checkouts, preserve existing ones
./bin/setup --all --list   # complete inventory of 51 libraries
./bin/setup --all          # includes the public Go-adapted QuickCheck fork
./bin/setup --locked --all # require the exact release revisions and clean sources
```

The **22 `CORE_PACKAGES`** are the default dependencies of compiler fixtures.
The **three `CORE_CHECKOUT_DEPENDENCIES`** are additional local overrides used
when developing those core libraries themselves. `ADDITIONAL_PACKAGES` completes
the list of 51, including `js-uri`. All three lists live in `bin/pkg`; setup never
changes the fixture dependency list. An unknown option, an invalid existing destination or
a failed clone stops setup with a nonzero status.

New checkouts use the public commits in [bootstrap-sources.json](https://github.com/0x000000000000000000001/gopurs/blob/main/bootstrap-sources.json),
including the Go-adapted QuickCheck fork. Ordinary setup preserves existing
checkouts; `--locked` additionally verifies their revisions and rejects local
changes. Set `GOPURS_LIBRARY_ROOT` to an empty directory for an isolated release
installation. The same variable selects libraries for the bootstrap and tests.

The build embeds the Go runtime, compiles PureScript, and bundles `Main` into
`bin/gopurs.js`. The checked-in FFI WASM and its matching JavaScript runtime are
used as-is. Node.js 24.8.0, local Spago 0.93.45 and the npm-provided `purs` 0.15.16
were used for the backend build recorded below.

For **application, library and fixture compilation**, put a TAST-capable `purs` on
`PATH`. A version number alone does not establish that it is the typed fork.
The example below was rechecked on 14 September 2026 with the locally built
fork and Spago 1.0.3; earlier integration checks used Spago 1.0.4; `bin/test` inherits the caller's `PATH`. npm build commands prioritize
this repository's `node_modules/.bin`, so inspect both toolchains when diagnosing
compiler differences. The fork currently writes the enriched format to
`output/<Module>/corefn.json`; the name `tcorefn` describes the format, not a
separate filename consumed by gopurs. Check a generated file for `dataDecls`,
`classDecls`, and, with the current fork, `typeTable`. If they are absent, select
the fork explicitly on `PATH` and rebuild in a fresh output directory.

A GitHub source install builds the JavaScript compiler from pinned dependencies.
Compiling applications still requires the typed compiler fork and Spago on
`PATH`. To install an already built backend in another project, create a local
archive after `npm run build`:

```bash
# From the backend repository:
npm pack --ignore-scripts

# From the application directory, using the archive just produced:
npm install --save-dev --ignore-scripts /absolute/path/to/gopurs-0.1.0-rc.3.tgz
```

The archive includes the backend bundle and FFI parser artifacts. It does not
install the application compiler or the Go library overrides. The local
archive installation was checked offline in an empty npm project; rebuilding
the backend there is unnecessary.

### Native compiler bootstrap

```bash
npm run build:native
# Override discovery of the local TAST compiler fork when needed:
GOPURS_PURS=/absolute/path/to/typed/purs npm run build:native
# Preserve the successful bootstrap workspace and its logs for inspection:
npm run build:native -- --keep-workspace
```

For the complete native validation path, run `./bin/setup --locked --all`
before bootstrapping. The Argonaut codec adaptation is included, with its
license, in `support/argonaut-codecs`. The bootstrap always selects this
versioned adaptation, so it does not depend on an unversioned sibling copy.

This bootstraps `bin/gopurs-native` with the existing Node backend. It requires
the npm dependencies above, Go, the typed compiler fork and the sibling Go
library checkouts installed by `./bin/setup --core`. The optimizer is resolved
from the pinned Spago checkout; set `GOPURS_PBO_DIR` to use another checkout.
The Node backend is rebuilt with this repository's npm toolchain. The isolated
TAST build uses a separately selected typed `purs`: `GOPURS_PURS` when set,
otherwise the newest compiler binary under
`../../purescript/.stack-work/dist/*/*/build/purs/purs`. Its path is printed, and
every generated module is checked for `typeTable`, `dataDecls` and `classDecls`
before Go generation. A stock compiler with the same version number is rejected.
`npm run test:cli` uses the same typed-compiler discovery; set `GOPURS_PURS` there
as well when the fork is supplied outside `../../purescript`.
The command then generates Go and links the FFI parser into the executable. Its package set is
`77.10.1`; sibling `gopurs-*` packages with `spago.yaml` provide the native FFI.
This includes the native persistent Map from `gopurs-ordered-collections`.

The existing native binary is replaced only after Go compilation succeeds.
Failures retain the isolated workspace and print the failed stage and log paths.
SIGINT/SIGTERM also stop the active command's descendants and return status
130/143. See the [tooling map](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/architecture.md#outils-de-build-et-de-validation)
and [bootstrap checks](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md#bootstrap-natif-et-processus-des-runners).
The resulting executable accepts the backend's usual arguments, such as
`--main Main`, from a project containing typed `output/<Module>/corefn.json`.
Its backend and FFI parser run without Node. The PureScript frontend that creates
the TAST and Go compilation of the generated application remain separate steps.

Native bootstrapping is experimental. Its PBO implementation cache currently
retains immutable modules in memory for one build, without the JavaScript
backend's disk spill or memory budget. The unused legacy JSON BackendModule
cache is unsupported: reads miss and an attempted write fails explicitly.
The native TAST decoder preserves integer magnitudes and isolated UTF-16
surrogates as WTF-8, while combining valid surrogate pairs. Literal preservation
does not establish every JavaScript string operation's semantics; the dated
[validation record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md#exclusions-et-modules-frères) documents the
string corrections and fixture reintegrations.

The 6 October 2026 reconstruction covers **500 modules and 288,603 types**.
Two fresh JS builds are byte-identical; the rebuilt native compiler also matches
the active native binary exactly.
TAST loading, transitive preparation, PBO optimization and Go emission each have
their own bounded parallelism controls. Native sequential, native parallel and
JS generation are compared on frozen inputs; see [the validation record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md)
for corpus sizes, execution campaigns and their limits. These checks establish
functional results, not performance gains.

### Content-verified build reuse

The driver checks typed inputs, FFI resolution, compiler resources and generated
outputs before decoding the TAST. An identical build reports a cache hit and
skips preparation, optimization and emission. After an input change, the driver
prepares global metadata and specializations, then reuses verified module-level
optimization and emission results. It works from private copies of its inputs
and publishes only changed files.
`go.mod` is created only when absent; existing dependency declarations and
`go.sum` are preserved. Continue to run `go build` to validate the application.

Use `--force` or `GOPURS_FORCE_BUILD=1` to force a complete generation.
`GOPURS_ALLOC_PROFILE` also forces the measured passes. Cache data lives under
`output/.gopurs-cache`; the separate `output/.gopurs-ownership.json` journal
tracks generated files and must survive a cache reset. A historical output is
adopted only when its generated files match a fresh generation. See the
[cache usage and maintenance commands](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cache.md#usage-quotidien-et-maintenance)
for forcing generation, resetting the reusable cache and reading diagnostics,
the [cache contract](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cache.md) for migration and concurrency,
and the [incremental-build results](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-incremental.md) for the
measured edit-time gains, cold-build overhead, cache size and validation evidence.
The [24 October follow-up](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cold-20261024.md) records the subsequent
allocation and array-intrinsic changes, with cold and incremental timings.
The completed b8x cycle campaign includes [all 24 measurements](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cycle-lot06.csv):
the common Go package still dominates rebuilds after edits, motivating the next
package-granularity work. The report includes regressions and pair-to-pair variation.

### Rust-hosted gopurs

```bash
npm run build:rust
GOPURS_RUST=1 ../gopurs/bin/gopurs --main Main # from a project with typed output/
```

`build:rust` compiles **gopurs itself** through the sibling `../../purust/purust`
backend and Cargo. The resulting `bin/gopurs-rust` generates Go, just like the
Go and JavaScript executables. It uses this repository's PureScript sources,
`purescript-backend-optimizer-gopurs`, and the `purust-*` libraries for its host
runtime. The same Go FFI parser is linked through a C archive into the executable.

The bootstrap requires the npm dependencies, typed `purs`, Go with cgo and a C
toolchain, Cargo, Purust and its sibling library ports. It builds threaded Rust
with release O3 and ThinLTO, then compiles and executes a fresh Go/FFI smoke test
before atomically installing the compiler. Failures retain the workspace/logs;
`--keep-workspace` retains successful builds too.
On macOS it uses the Mach-O linker bundled with the Rust toolchain, keeping
the linker compatible with rustc's LLVM bitcode version.

`GOPURS_PURS` overrides the frontend, `PURUST_DIR` the Purust checkout,
`GOPURS_NATIVE_TMPDIR` the workspace parent, and `GOPURS_RUST_OUTPUT` the binary
destination. `GOPURS_KEEP_WORKSPACE=1` also retains successful builds invoked by
library runners. Cargo honors `CARGO_TARGET_DIR`. `PURUST_JS=1` selects the JavaScript
Purust executable for this bootstrap; it has no effect on an installed gopurs
compiler's host selection. `GOPURS_JS=1` and `GOPURS_RUST=1` are mutually exclusive.

In `../gopurs-aff`, `GOPURS_RUST=1 ./bin/test` generates, builds and executes Go;
adding `-c` rebuilds this Rust-hosted compiler first.
On machines with at least 32 GiB of RAM, the launcher selects eight preparation
and PBO workers for Rust as well as Go and JavaScript. Explicit worker settings
take precedence; the automatic Go GC policy applies only to the Go compiler host.

With all three compiler executables built, `npm run test:rust` checks exact Go
output and CLI failures in sequential/parallel modes, then compiles and executes
the string, FFI, record-worker, owned-tree and JSON regression fixtures on each
host. `GOPURS_TEST_KEEP_WORKSPACE=1` retains successful fixture workspaces.

### Compile and run an application

For `workspace/hello`, create `src/Main.purs`:

```purescript
module Main where

import Prelude (Unit)
import Effect (Effect)
import Effect.Console (log)

main :: Effect Unit
main = log "Hello from gopurs"
```

Use this `spago.yaml` for the example:

```yaml
package:
  name: hello
  dependencies:
    - prelude
    - effect
    - console
workspace:
  packageSet:
    registry: 77.10.1
  extraPackages:
    prelude:
      path: "../gopurs/gopurs-prelude"
    effect:
      path: "../gopurs/gopurs-effect"
    console:
      path: "../gopurs/gopurs-console"
```

Add the corresponding Go overrides for any other packages with FFI used by your
application. Ordinary registry packages may supply JavaScript FFI only.
With the typed compiler and Spago on `PATH`, run from `workspace/hello`:

```bash
spago build
../gopurs/gopurs/bin/gopurs --main Main
cd output
go mod tidy
go run ./main
```

This prints `Hello from gopurs`. If the archive is installed in the application,
replace the backend invocation with `./node_modules/.bin/gopurs --main Main`.
`output/go.mod` is generated by gopurs; there is no `go mod init` step.
The entry file is `output/main/main.go`. To build an executable from `output`,
use `go build -o hello ./main`. Generated modules declare Go 1.22; the integration
checks used Go 1.27.0, without establishing the oldest working Go toolchain.

### Compiler options

The backend reads and writes `output` in its current directory. Options used
by the current gopurs entrypoint are:

| Option | Behavior |
| --- | --- |
| `--main Module` | Select the entry module explicitly. Without it, discover all loaded modules exporting `main`; each gets `output/<Module>/main/main.go`, and the shared `output/main/main.go` belongs to the last discovered target. |
| `--ffi directory` | Supply an additional FFI lookup directory. |
| `--dev` | Skip monomorphization and reduce imported inlining; keep incremental builds. |
| `--release` | Select the existing optimized behavior explicitly; this is also the default. |
| `--rewrite-limit number` | Set the optimizer rewrite limit; default 10000. |
| `--force` | Recompute preparation, optimization and emission; publish only changed Go files. Also available as `GOPURS_FORCE_BUILD=1`. |

The shared optimizer argument parser also recognizes options such as `--output`
and `--bundle`, but `Driver.Config` does not use them to change gopurs output behavior.
Argument values containing spaces are not supported by this parser, including
quoted `--ffi` paths. There is no dedicated `--help` handler.

The CLI exits with status **0** on successful compilation and **1** when the
driver raises an error. After compilation cleanup, it reports the original
message once on stderr, prefixed with `[gopurs] error:`. This contract applies
to all three compiler hosts selected by the launcher.

### Build profiles

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
../gopurs/gopurs/bin/gopurs --dev --main Main
../gopurs/gopurs/bin/gopurs --release --main Main
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

### Backend compilation timings

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

`Gopurs.Metrics` uses a monotonic clock, following altbak's `Bench` approach:
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

On 2026-09-21, after removing whole-array conversions from individual indexed
reads, an alternating comparison on the same b8x TAST and native binary measured
21.077/21.174 s with one worker and 13.659/13.536 s with eight: about 36% less
elapsed loading time, with similar CPU time. Peak RSS rose from about 1.59 GiB
to 2.24 GiB. These measurements cover loading and sorting, not a complete build.

For historical context, on 2026-09-18 a sample of 133 b8x modules
(7.1 MiB of TAST) took 1.26 s with one worker, 1.70 s with four and 2.39 s with
eight with the earlier implementation (two runs per setting, without race instrumentation). Limiting
`GOMAXPROCS` to four did not reverse this regression. These measurements cover
loading and sorting only, not a complete b8x build or altbak runtime benchmarks.
The native race check passed on this sample, and Node tests compare complete
decoded modules and dependency order between sequential and parallel loading.

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

## Develop one library locally

Each library keeps its own Spago configuration and repository. The
[repository map](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/architecture.md#carte-des-dépôts-et-des-consommateurs) describes
the source families and their consumers. `./bin/modtest --all --list` gives the
current executable library suite. The Go development section in each
library README links back to this procedure.

1. Build gopurs as above and prepare the required sibling checkouts. `--core`
   covers the 22 core libraries and their supporting dependencies; use `--all`
   for the complete family, including the adapted QuickCheck prerequisite.
2. Put the TAST compiler and YAML-capable Spago on `PATH` (**1.0.3** in the
   2 October campaign; earlier checks also used 1.0.4). Check both with
   `command -v purs spago` and `purs --version; spago --version`. The interactive
   shell, npm and altbak can select different tools. A stock compiler's version
   number alone does not prove that it emits TAST.
3. From the library root, inspect the resolved packages and build:

   ```bash
   spago ls deps --transitive --json
   spago build
   ../gopurs/bin/gopurs --main Test.Main
   (cd output && go mod tidy && go run ./main)
   ```

   Add `--offline` to the Spago commands when all dependencies are cached.
    Remove stale frontend TAST when validating a removed or renamed module, or
    compile the frontend into a fresh output. gopurs removes obsolete Go files
    owned by its journal; module discovery follows the TAST that remain present.
    Library `bin/test` scripts
   clean only their own build directories. For an isolated campaign, run
   `./bin/modtest prelude strings --keep-going` from gopurs: each target gets a
   fresh copy of the library family and private Spago temporaries.
   They describe the test path; they do not establish that every library's
   tests are already passing. See the [coverage limits](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md).

Both `bin/test` and `bin/modtest` print a `Results:` JSON report path. Reports
and per-target logs are retained; `--keep-workspace` also retains successful
builds. Resume unsuccessful or unattempted targets with
`./bin/modtest --resume-failed /path/to/results.json --keep-going` (or `bin/test`
for fixtures). The retry creates a new report linked to the original campaign.
See [isolation and reports](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md#isolation-logs-et-caches).

`assert` has no executable test suite: use `spago build`, then
`../gopurs/bin/gopurs`, then `(cd output && go mod tidy && go build ./...)`.
QuickCheck has a standalone deterministic `Test.Go` suite: use
`./bin/modtest quickcheck` from gopurs, or `./bin/test` from its checkout. The
runner builds and executes JavaScript and Go, then compares their transcripts.
`node-net` has a `test/Test/Main.purs` entrypoint that
the current Spago command compiles despite the absent `package.test` stanza;
its runner executes that Go program. See [the module campaign](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md).

### Configuration and lockfiles

The Go configuration is `spago.yaml`. In 42 libraries it is a **tracked** link
to `spago.go.yaml`, which is the file to edit; the other configurations are
ordinary files. Setup preserves these links. Library `workspace.extraPackages`
lists only overrides used by that workspace's resolved graph, including its
test dependencies. A dependency's own `workspace` does not supply overrides to
the consuming application: add the necessary Go overrides to that application's
configuration as well. Keep required transitive overrides; retaining only the
direct imports can silently select registry JavaScript implementations.

Most libraries use registry set **77.7.0**; `assert` retains **73.3.0**, while
the backend uses **77.10.1**. Package names in four historical configurations
(`js-promise`, `js-promise-aff`, `node-path`, `node-process`) include `gopurs-`;
their existing resolution aliases are preserved. No package versions or public
module names were changed by the configuration cleanup.

Spago lockfiles record registry versions and local **paths**, not commits of
sibling repositories. Keep related checkouts at the intended revisions when
reproducing a build. Update a tracked lockfile with the same Spago version as
the workspace after changing its configuration. Seven libraries currently
ignore their local Spago lockfile: `aff`, `argonaut-core`, `avar`, `js-date`,
`now`, `nullable` and `strings-extra`. This existing policy is preserved;
dependency resolution was also checked without these local files.

The backend's `package-lock.json` is the source for `npm ci`; its npm `prepare`
hook builds the bundle. Library npm manifests, Bower files and the eleven
Dhall configuration pairs still serve their JavaScript or upstream workflows.
In particular, Spago 0.20/0.21 reads Dhall while Spago 0.93/1.x reads YAML.
Follow a library's existing npm/CI commands for those workflows; `npm test`
does not universally run Go. Preserve npm lockfiles used by those commands.
The Go parser has the only maintained `go.mod`, in `tools/ffi-gen`; the backend
creates each application's Go module under `output`.

## Foreign function interface

Place a `.go` file beside the corresponding `.purs` file, or provide a lookup
directory with `--ffi`. For a foreign import named `returnInt64`, export a Go
function named `ReturnInt64`. The backend prefixes package-level declarations
and their references, merges generated modules into the `purescript` Go package,
and creates bridge functions using the parsed signatures and TAST types.

Supported bridges handle primitive values, slices, selected records and ADTs,
and callbacks. Their conversion rules are implemented in
[FfiBridge](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Gopurs/FfiBridge.purs); see the executable
[integer FFI fixture](https://github.com/0x000000000000000000001/gopurs/blob/main/tests/passing/FFIIntegerReturns.purs) and its
[Go implementation](https://github.com/0x000000000000000000001/gopurs/blob/main/tests/passing/FFIIntegerReturns.go) for native `int`, `int64`,
slice, and dynamic return examples. Effect bindings must follow the thunk
convention used by the Go library ports.

The Go parser is not a general binding generator for every Go API. Confirm that
an imported signature is supported before using it. A missing binding can
produce a panic stub, so successful Go compilation alone does not establish FFI
coverage. Parser and decoding failures report the PureScript module and FFI
path and stop generation.

## Development and testing

Start with [the architecture map](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/architecture.md) to locate the owner of
a change. [Testing and validation](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md) describes targeted fixture
commands, snapshot review, cache behavior, the runner, exclusions and remaining
coverage gaps.

```bash
./bin/test --list
npm run build:native -- --keep-workspace
./bin/test FFIIntegerReturns
./bin/test -d FFIIntegerReturns
npm run test:runner
npm run test:profiles
```

Run the complete fixture suite separately for each profile:

```bash
./bin/test --all --keep-going
./bin/test -d --all --keep-going
```

`test:profiles` checks the JS and native parallel hosts after rebuilding them:
profile selection, conflicting flags, runtime behavior, cache coexistence and
incremental output against full rebuilds after source changes.
It also checks generic ADT, record and type-class calls with no specialization
clones, including edits and forced rebuilds.

`bin/test -c` rebuilds the selected compiler host: native by default,
JavaScript with `GOPURS_JS=1`, or Rust with `GOPURS_RUST=1`.

For a broader integration check, run `bin/go/run -c` from an `altbak.pub`
checkout. It rebuilds the backend and application, compiles Go, and executes
the 14 cases of the default `pure` campaign. See [testing and validation](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md)
for dated results and the exact scope of completed checks. Performance
comparisons use the altbak baselines separately from functional validation.

### Rebuilding the FFI parser

The Go parser lives in `tools/ffi-gen`. `parser.go` analyzes declarations and
returns the JSON contract; `types.go` defines that contract, and
`main_js_wasm.go` exposes it to JavaScript. `tools/ffi-runner.mjs` reads Go source
from stdin and writes the declarations as JSON to stdout. Valid input with no
retained declarations returns `[]`. Invalid Go or a runner failure writes a
diagnostic to stderr and exits unsuccessfully. The backend also rejects invalid
JSON responses, reporting the PureScript module and the FFI file path.

The generator's `go.mod` pins **Go 1.27.0**. With that version on `PATH`, run:

```bash
npm run build:ffi
```

This explicitly rebuilds `tools/ffi_gen.wasm` for `js/wasm` and copies
`tools/wasm_exec.js` from the same Go installation. The command checks the exact
Go version, disables automatic toolchain switching, and uses `-trimpath`,
`-buildvcs=false` and an empty build ID. Both artifacts are prepared before
replacement, so a compilation failure preserves the existing pair. Commit the
WASM and its matching JavaScript runtime together when regenerating them.

The parser's contract tests run natively, without WASM:

```bash
cd tools/ffi-gen
go test ./...
```

After rebuilding the WASM and the backend with `npm run build:ffi` and
`npm run build`, run the Node integration tests from the repository root:

```bash
npm run test:ffi
```

These exercise the typed `FfiSupport` API, syntax and JSON errors, and missing
or corrupt runner/WASM files in temporary directories. `Main` receives decoded
declarations; extraction and decoding failures propagate as exceptions. A build
can have written some output files before failing.

The checked-in WASM and runtime are shipped with `tools/ffi-runner.mjs` in the
npm package. `npm run build` and the installation `prepare` hook build the
PureScript backend using those artifacts; rebuilding the parser is a separate
maintainer command. After editing `FfiSupport.js`, run `npm run build` to update
the backend bundle. Runner paths are resolved relative to the source, Spago
output or installed bundle, independently of the current directory.

The parser workflow was validated with Node.js 24.8.0 and Go 1.27.0.
See [validation and remaining limits](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md) for the scope of the checks.

### Editing the Go runtime

The canonical runtime source is [`runtime/runtime.go`](https://github.com/0x000000000000000000001/gopurs/blob/main/runtime/runtime.go).
Edit that file, then rebuild the compiler hosts you use:

```bash
npm run build:native
npm run build:rust
```

The build first runs `tools/embed-runtime.mjs`, which writes the ignored
`src/Gopurs/Runtime.js`, `Runtime.go` and `Runtime.rs` FFI companions.
`Gopurs.Runtime.runtimeGoCode` remains a PureScript `String`; all three compiler
hosts embed the same constant. The installed
backend does not read the Go source at execution time, and the generated
`output/gopurs_runtime/runtime.go` contains its exact text.

`npm run build:runtime` regenerates only the FFI companions. Run it before a direct
`spago build` after changing the Go source or starting from a fresh checkout.
When the source is unchanged, regeneration preserves the FFI file's timestamp
so Spago can reuse its compiled output. The npm `prepare` hook runs the JS
build and includes this step automatically. Generated FFI and bundle files are
build artifacts; commit the Go source and build tooling.

## Architecture

[`Main`](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Main.purs) launches [`Driver`](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Gopurs/Driver.purs), which owns phase ordering:

1. **Load and prepare:** [Driver.Prepare](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Gopurs/Driver/Prepare.purs) loads enriched CoreFn, builds type and constructor metadata, and applies partial monomorphization.
2. **Optimize and lower:** [Driver.Build](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Gopurs/Driver/Build.purs) runs the optimizer and publishes module signatures; [CodeGen](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Gopurs/CodeGen.purs) and the specialized `Gopurs` modules lower backend modules to Go representations and expressions.
3. **Print and bridge:** [Printer](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Gopurs/Printer.purs) emits Go source; [FfiSupport](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Gopurs/FfiSupport.purs) prepares foreign declarations for [FfiBridge](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Gopurs/FfiBridge.purs).
4. **Assemble and execute:** [Driver.Output](https://github.com/0x000000000000000000001/gopurs/blob/main/src/Gopurs/Driver/Output.purs) writes modules, the embedded runtime, `go.mod`, and entrypoints. Go's tools resolve dependencies and compile the application.

The [architecture map](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/architecture.md) gives a detailed guide to module responsibilities.
The [runtime/FFI contracts](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/runtime-ffi-contracts.md) cover shared storage,
closure lifetime, and the per-render mutable buffer owned by `Printer.Builder`.

## Current status and limitations

The backend remains experimental. Native representations coexist with a tagged
`Value` runtime; general unboxing and complete library compatibility are not
promised. Go implementations are needed for every reachable foreign binding.

As of **7 October 2026**, the fixture runner has **no exclusions**. The v2 work
resolved the integer, number-display and string cases, then restored the four
derived-instance fixtures by importing the instance modules needed for isolated
compilation. Their derived methods are exercised in JS and Go. See
[the reintegration evidence](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md#exclusions-et-modules-frères).

The **6–7 October final validation** passes all **400 fixtures** with strict
snapshots and Go execution, plus **51 sibling-library runners**: 50 with Go
execution and the compile-only `assert` check. QuickCheck includes a deterministic
JavaScript/Go transcript comparison. `spec` passes **74 tests**, including nine
integration cases, and retains **3 intentional pending fixtures**; its two
initialization tests also pass.

The JS/native builds from fresh checkouts, offline npm archive installation and
**435 Node/Go checks** pass. On **2,683 frozen b8x TAST inputs**, all **2,987 Go
files plus `go.mod`** are byte-identical across the previous JS bundle, rebuilt
JS, native sequential and native parallel hosts. `nix flake check`, `nix develop`
and the legacy shell evaluation pass on **aarch64-linux**; see
[the Nix section](#nix-environment) for the toolchain scope.

The [final validation record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md#plan-v2--lot-08--installation-et-validation-finale-7-octobre-2026)
identifies the local dependencies, retained evidence and disk-space retries.
The [v1 maintainability plan](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md#consolidation-finale--lot-15-2-octobre-2026)
is complete at **15/15 lots, 100/100 points**, and the
[v2 reliability and reproducibility plan](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md#plan-v2--lot-08--installation-et-validation-finale-7-octobre-2026)
is complete at **8/8 lots, 100/100 points**. Performance work remains paused.

## License

MIT License. See [LICENSE](LICENSE) for details. Third-party attributions are in
[NOTICE](NOTICE) and [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md). The
publication scope and compatibility policy is in
[docs/publication-scope.md](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/publication-scope.md).
