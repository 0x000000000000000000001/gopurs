# 🐹 gopurs

<img height="160" alt="gopurs" src="https://github.com/user-attachments/assets/b013e7c3-fac6-4ee8-9d4c-f39ac8c2c921" />

An experimental **PureScript-to-Go backend**. The compiler is written in
PureScript, with a Go runtime, a Go parser compiled to WebAssembly for FFI
signatures, and JavaScript build tooling.

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
- **Native FFI bridges.** A Go parser compiled to WebAssembly reads foreign declarations; generated wrappers adapt supported Go signatures to PureScript calls.
- **Go runtime and library ports.** Generated programs include the runtime. The Go `Aff` port uses goroutines, and generated entrypoints wait for work registered with the runtime event loop.

## Benchmarks

The [altbak README](https://github.com/0x000000000000000000001/altbak.pub#go)
contains the reference results, workloads, and benchmark context. Compare
changes against those baselines; results depend on the workload and toolchain.
The sequential core campaign does not establish multicore scaling.

## Getting started

### Prerequisites

- Node.js and npm to build and run the backend. Node.js `24.8.0` is the documented development reference; a minimum Node version is not declared in `package.json`.
- A TAST-capable `purs` from the compiler fork for applications, plus Spago with YAML configuration support.
- Go to compile generated applications. Their `go.mod` declares Go 1.22; the walkthrough was checked with Go `1.27.0`, without establishing the minimum working version across library ports.
- Git and Bash for the checkout and helper scripts.

Go `1.27.0` is specifically required to rebuild the FFI parser. Ordinary backend builds use its checked-in WASM and JavaScript runtime.

### Build the backend

The source build currently depends on local checkouts. Keep this layout, or
adjust the paths in [spago.yaml](spago.yaml) explicitly:

```text
workspace/
├── purescript/                         # TAST compiler fork, if built locally
├── purescript-backend-optimizer-gopurs/ # optimizer fork, edge-gopurs branch
├── gopurs/
│   ├── gopurs/                         # this repository
│   ├── gopurs-prelude/
│   ├── gopurs-effect/
│   ├── gopurs-console/
│   ├── gopurs-st/
│   ├── gopurs-unsafe-coerce/
│   ├── gopurs-assert/
│   └── ...                            # other core Go library forks
└── hello/                             # example application below
```

The optimizer checkout comes from the
[Go branch of our optimizer fork](https://github.com/0x000000000000000000001/purescript-backend-optimizer/tree/edge-gopurs).
Its path is `../../purescript-backend-optimizer-gopurs` relative to this
repository. `st` and `unsafe-coerce` are also local build dependencies.
From this repository, `./bin/setup` installs the library siblings listed in
[bin/pkg](bin/pkg); it does not install the optimizer or the typed compiler.

```bash
# From an empty workspace directory:
git clone --branch edge-gopurs https://github.com/0x000000000000000000001/purescript-backend-optimizer.git purescript-backend-optimizer-gopurs
mkdir gopurs
cd gopurs
git clone https://github.com/0x000000000000000000001/gopurs.git
cd gopurs
./bin/setup
npm ci
```

`npm ci` installs the locked npm dependencies and its `prepare` hook runs
`npm run build`. Later edits need only:

```bash
npm run build
```

### Choose the library checkouts

```bash
./bin/setup --core --list  # 22 core libraries + lazy, ordered-collections, random
./bin/setup --core         # default; install missing checkouts, preserve existing ones
./bin/setup --all --list   # complete inventory of 50 libraries
./bin/setup --all          # requires the local Go-adapted QuickCheck checkout
```

The **22 `CORE_PACKAGES`** are the default dependencies of compiler fixtures.
The **three `CORE_CHECKOUT_DEPENDENCIES`** are additional local overrides used
when developing those core libraries themselves. `ADDITIONAL_PACKAGES` completes
the list of 50. All three lists live in `bin/pkg`; setup never changes the
fixture dependency list. An unknown option, an invalid existing destination or
a failed clone stops setup with a nonzero status.

QuickCheck is a specific local prerequisite for the complete library suite.
The available checkout is upstream **v8.0.1**, with the Go adaptation supplied
by `spago.yaml` and `src/Test/QuickCheck/Gen.go`. An upstream clone alone does
not supply these two files. Keep the adapted checkout, including those files,
at `../gopurs-quickcheck`; `--all` checks it before cloning anything. No remotely
installable Go fork was established in this review, so setup does not invent
one. Publishing that adaptation is separate from local setup.

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

A GitHub source install is not a standalone installation recipe with these
relative build dependencies. To install an already built backend in another
project, create a local archive after `npm run build`:

```bash
# From the backend repository:
npm pack --ignore-scripts

# From the application directory, using the archive just produced:
npm install --save-dev --ignore-scripts /absolute/path/to/gopurs-0.1.0.tgz
```

The archive includes the backend bundle and FFI parser artifacts. It does not
install the application compiler or the Go library overrides. The local
archive installation was checked offline in an empty npm project; rebuilding
the backend there is unnecessary.

### Experimental native compiler

```bash
npm run build:native
# Override discovery of the local TAST compiler fork when needed:
GOPURS_PURS=/absolute/path/to/typed/purs npm run build:native
# Preserve the successful bootstrap workspace and its logs for inspection:
npm run build:native -- --keep-workspace
```

This bootstraps `bin/gopurs-native` with the existing Node backend. It requires
the npm dependencies above, Go, the local `purescript-backend-optimizer-gopurs`
checkout and the sibling Go library checkouts. The Node backend is rebuilt with
this repository's npm toolchain. The isolated TAST build uses a separately selected
typed `purs`: `GOPURS_PURS` when set, otherwise the newest compiler binary under
`../../purescript/.stack-work/dist/*/*/build/purs/purs`. Its path is printed, and
every generated module is checked for `typeTable`, `dataDecls` and `classDecls`
before Go generation. A stock compiler with the same version number is rejected.
The command then generates Go and links the FFI parser into the executable. Its package set is
`77.10.1`; sibling `gopurs-*` packages with `spago.yaml` provide the native FFI.
This includes the native persistent Map from `gopurs-ordered-collections`.

The existing native binary is replaced only after Go compilation succeeds.
Failures retain the isolated workspace and print the failed stage and log paths.
The resulting executable accepts the backend's usual arguments, such as
`--main Main`, from a project containing typed `output/<Module>/corefn.json`.
Its backend and FFI parser run without Node. The PureScript frontend that creates
the TAST and Go compilation of the generated application remain separate steps.

Native bootstrapping is experimental. Its PBO implementation cache currently
retains immutable modules in memory for one build, without the JavaScript
backend's disk spill or memory budget. The unused legacy JSON BackendModule
cache is unsupported: reads miss and an attempted write fails explicitly.
The native JSON parser currently replaces isolated UTF-16 surrogates with
U+FFFD, so literals containing those code units do not yet have JavaScript
parity. Ordinary Unicode strings, including valid surrogate pairs, are unaffected
by this specific limitation.
The bootstrap was validated on 2026-09-17 by running the native backend on its
own 448 TAST modules: all 543 generated Go files matched the Node backend byte
for byte, and rebuilding them produced an identical second-generation binary.
That binary also generated identical Go for the Hello and NativeArrayReboxing
fixtures (80 and 193 files), whose applications compiled and ran successfully.
This establishes functional parity on these inputs, not general equivalence.
The native backend was slower on those validation workloads. TAST loading and
decoding support optional bounded parallel batches. Go emission also supports
bounded parallel batches; optimization remains sequential.

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
| `--rewrite-limit number` | Set the optimizer rewrite limit; default 10000. |

The shared optimizer argument parser also recognizes options such as `--output`
and `--bundle`, but `Main` does not use them to change gopurs output behavior.

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

`Gopurs.Metrics` uses a monotonic clock, following altbak's `Bench` approach:
`performance.now()` under Node and `time.Since` in the native Go compiler.
No flag is needed; both compiler builds report the same phases. These are real
elapsed times for the current invocation, not warm-up or repeated benchmarks.

TAST loading and decoding remain sequential by default. Set `GOPURS_JOBS` from
1 to 64 to select a worker count (for example, `GOPURS_JOBS=4 b -n` in b8x).
Invalid values use the default. Batches preserve input order before
dependency sorting. Native workers run on goroutines; Node overlaps file I/O
but still decodes JSON on its JavaScript thread. Other compiler passes are
unaffected by this setting.

Parallel loading is experimental: on 2026-09-18, a sample of 133 b8x modules
(7.1 MiB of TAST) took 1.26 s with one worker, 1.70 s with four and 2.39 s with
eight (two runs per setting, without race instrumentation). Limiting
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
the CPU work concurrently. PBO optimization and directive propagation retain
their original sequential order. See [parallel emission](docs/parallel-emission.md)
for measurements and validation.

## Develop one library locally

Each library keeps its own Spago configuration and repository. The
[51-directory map](todo.md#dossiers-api-et-commandes-actuelles) identifies its
API, test entrypoint and review owner. The Go development section in each
library README links back to this procedure.

1. Build gopurs as above and prepare the required sibling checkouts. `--core`
   covers the 22 core libraries and their supporting dependencies; use `--all`
   for the complete family, including the adapted QuickCheck prerequisite.
2. Put the TAST compiler and **Spago 1.0.4** on `PATH`. Check both with
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
   Start with an empty local `output` when validating a removed or renamed
   module: gopurs does not purge obsolete Go files. These direct commands do
   not run the sibling cleanup performed by many existing `bin/test` scripts.
   They describe the test path; they do not establish that every library's
   tests are already passing. See the [coverage limits](docs/testing.md).

`assert` has no executable test suite: use `spago build`, then
`../gopurs/bin/gopurs`, then `(cd output && go mod tidy && go build ./...)`.
QuickCheck has a local package configuration but no Go runner or `package.test`
declaration yet; use `spago build` for its package, and the consuming package's
Go tests to exercise it. `node-net` also lacks a `package.test` declaration;
its test entrypoint must be reconciled with the suite in the later test review.

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
Argument values containing spaces are not supported by that parser, including
quoted `--ffi` paths. There is no dedicated `--help` handler.

## Foreign function interface

Place a `.go` file beside the corresponding `.purs` file, or provide a lookup
directory with `--ffi`. For a foreign import named `returnInt64`, export a Go
function named `ReturnInt64`. The backend prefixes package-level declarations
and their references, merges generated modules into the `purescript` Go package,
and creates bridge functions using the parsed signatures and TAST types.

Supported bridges handle primitive values, slices, selected records and ADTs,
and callbacks. Their conversion rules are implemented in
[FfiBridge](src/Gopurs/FfiBridge.purs); see the executable
[integer FFI fixture](tests/passing/FFIIntegerReturns.purs) and its
[Go implementation](tests/passing/FFIIntegerReturns.go) for native `int`, `int64`,
slice, and dynamic return examples. Effect bindings must follow the thunk
convention used by the Go library ports.

The Go parser is not a general binding generator for every Go API. Confirm that
an imported signature is supported before using it. A missing binding can
produce a panic stub, so successful Go compilation alone does not establish FFI
coverage. Parser and decoding failures report the PureScript module and FFI
path and stop generation.

## Development and testing

Start with [the architecture map](docs/architecture.md) to locate the owner of
a change. [Testing and validation](docs/testing.md) describes targeted fixture
commands, snapshot review, cache behavior, the runner, exclusions and remaining
coverage gaps.

```bash
./bin/test --list
./bin/test FFIIntegerReturns -c
npm run test:runner
```

For a broader integration check, run `bin/go/run -c` from an `altbak.pub`
checkout. It rebuilds the backend and application, compiles Go, and executes
the 14 cases of the default `pure` campaign. See [testing and validation](docs/testing.md)
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
See [validation and remaining limits](docs/testing.md) for the scope of the checks.

### Editing the Go runtime

The canonical runtime source is [`runtime/runtime.go`](runtime/runtime.go).
Edit that file, then rebuild the backend:

```bash
npm run build
```

The build first runs `tools/embed-runtime.mjs`, which writes the ignored
`src/Gopurs/Runtime.js` FFI module. `Gopurs.Runtime.runtimeGoCode` remains a
PureScript `String`; Spago and esbuild embed it in `bin/gopurs.js`. The installed
backend does not read the Go source at execution time, and the generated
`output/gopurs_runtime/runtime.go` contains its exact text.

`npm run build:runtime` regenerates only the FFI module. Run it before a direct
`spago build` after changing the Go source or starting from a fresh checkout.
When the source is unchanged, regeneration preserves the FFI file's timestamp
so Spago can reuse its compiled output. The npm `prepare` hook runs the full
build and includes this step automatically. Generated FFI and bundle files are
build artifacts; commit the Go source and build tooling.

## Architecture

1. **Load and prepare:** [Main](src/Main.purs) loads enriched CoreFn, builds type and constructor metadata, and applies partial monomorphization.
2. **Optimize and lower:** the optimizer produces backend modules; [CodeGen](src/Gopurs/CodeGen.purs) and the specialized `Gopurs` modules lower them to Go representations and expressions.
3. **Print and bridge:** [Printer](src/Gopurs/Printer.purs) emits Go source; [FfiSupport](src/Gopurs/FfiSupport.purs) prepares foreign declarations for [FfiBridge](src/Gopurs/FfiBridge.purs).
4. **Assemble and execute:** the backend writes modules, the embedded runtime, `go.mod`, and entrypoints. Go's tools resolve dependencies and compile the application.

The [architecture map](docs/architecture.md) gives a detailed guide to module responsibilities.

## Current status and limitations

The backend remains experimental. Native representations coexist with a tagged
`Value` runtime; general unboxing and complete library compatibility are not
promised. Go implementations are needed for every reachable foreign binding.

The fixture runner records known exclusions for integer overflow at 32-bit
boundaries, floating-point serialization, isolated UTF-16 surrogates, and some
compiler-feature fixtures. Go's native integer and string representations need
care when porting code that depends on those JavaScript edge cases. See
[the exclusions and validation gaps](docs/testing.md#exclusions-et-modules-frères).

Targeted fixtures and the core benchmark campaign have recorded successful
checks; a complete green run of all `passing` fixtures and sibling-library
suites is not established by those results. The [development plan](todo.md)
tracks remaining work.

## License

MIT License. See [LICENSE](LICENSE) for details.
