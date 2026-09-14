# gopurs

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
The [development log](https://discourse.purescript.org/t/leveraging-a-blazing-fast-runtime-a-new-go-backend-for-purescript)
records the earlier design work.

## Benchmarks

The [altbak README](https://github.com/0x000000000000000000001/altbak.pub#go)
contains the reference results and benchmark context. Use those baselines when
comparing performance; successful cleanup checks alone do not demonstrate a
speed or memory improvement.

## Build the backend locally

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
repository. `st`, `unsafe-coerce` and `assert` are also local build dependencies.
From this repository, `./bin/setup` clones the core library siblings listed in
[bin/pkg](bin/pkg); it does not install the optimizer or the typed compiler.

```bash
# From workspace/gopurs/gopurs, after installing the optimizer checkout:
./bin/setup
npm ci
```

`npm ci` installs the locked npm dependencies and its `prepare` hook runs
`npm run build`. Later edits need only:

```bash
npm run build
```

The build embeds the Go runtime, compiles PureScript, and bundles `Main` into
`bin/gopurs.js`. The checked-in FFI WASM and its matching JavaScript runtime are
used as-is. Node.js 24.8.0, local Spago 0.93.45 and the npm-provided `purs` 0.15.16
were used for the backend build recorded below.

For **application and fixture compilation**, put a TAST-capable `purs` on
`PATH`. A version number alone does not establish that it is the typed fork.
The validated application toolchain uses the locally built fork and Spago
1.0.4; `bin/test` inherits the caller's `PATH`. npm build commands prioritize
this repository's `node_modules/.bin`, so inspect both toolchains when diagnosing
compiler differences. The fork currently writes the enriched format to
`output/<Module>/corefn.json`; the name `tcorefn` describes the format, not a
separate filename consumed by gopurs.

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

## Compile and run an application

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

The backend reads and writes `output` in its current directory. Options used
by the current gopurs entrypoint are:

| Option | Behavior |
| --- | --- |
| `--main Module` | Select the entry module explicitly. Without it, discover all loaded modules exporting `main`; each gets `output/<Module>/main/main.go`, and the shared `output/main/main.go` belongs to the last discovered target. |
| `--ffi directory` | Supply an additional FFI lookup directory. |
| `--rewrite-limit number` | Set the optimizer rewrite limit; default 10000. |

The shared optimizer argument parser also recognizes options such as `--output`
and `--bundle`, but `Main` does not use them to change gopurs output behavior.

## Development and validation

Start with [the architecture map](docs/architecture.md) to locate the owner of
a change. [Testing and validation](docs/testing.md) describes targeted fixture
commands, snapshot review, cache behavior, the runner, exclusions and remaining
coverage gaps.

```bash
./bin/test --list
./bin/test FFIIntegerReturns -c
npm run test:runner
```

For the recent cleanup batches, the agreed integration checkpoint is
`bin/go/run -c` from `altbak.pub`. It rebuilds the backend and application,
compiles Go and executes the 14 cases of the default `pure` campaign. The
recorded comparisons preserved all 387 generated Go files and their functional
outputs. This does not establish a green full `passing` or sibling-module suite.
The [cleanup plan](todo.md) tracks the current wave; [testing and validation](docs/testing.md) records the completed checks.

## Rebuilding the FFI parser

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

## Editing the Go runtime

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

## License

MIT License. See [LICENSE](LICENSE) for details.
