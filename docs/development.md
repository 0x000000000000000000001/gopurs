# Development guide

Use the [quickstart](quickstart.md) to build an application, or the
[bootstrap guide](bootstrap.md) to build the compiler itself. This guide covers
changes to library ports, Go FFI and the embedded runtime.

- [Libraries](#libraries)
- [Go FFI](#go-ffi)
- [FFI parser](#ffi-parser)
- [Embedded runtime](#embedded-runtime)

## Libraries

Each library keeps its own Spago configuration and repository. The
[repository map](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/architecture.md#carte-des-dépôts-et-des-consommateurs) describes
the source families and their consumers. `./bin/modtest --all --list` gives the
current executable library suite. The Go development sections in the library READMEs link to this procedure.

1. Build gopurs with the [bootstrap guide](bootstrap.md#clean-source-installation).
   From the gopurs checkout, run `./bin/setup --core` for the 22 core libraries,
   their supporting dependencies and the Argonaut codec support checkout, or
   `./bin/setup --all` for the complete family, including the adapted QuickCheck
   prerequisite.
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

## Go FFI

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

## FFI parser

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

## Embedded runtime

The canonical runtime source is [`runtime/runtime.go`](https://github.com/0x000000000000000000001/gopurs/blob/main/runtime/runtime.go).
Edit that file, then rebuild the compiler hosts you use:

```bash
npm run build          # JavaScript host
npm run build:native   # native Go host
npm run build:rust     # native Rust host
```

The build first runs `tools/build/embed-runtime.mjs`, which writes the ignored
`src/Gopurs/Runtime/Source.js`, `Source.go` and `Source.rs` FFI companions.
`Gopurs.Runtime.Source.runtimeGoCode` remains a PureScript `String`; all three compiler
hosts embed the same constant. The installed
backend does not read the Go source at execution time, and the generated
`output/gopurs_runtime/runtime.go` contains its exact text.

`npm run build:runtime` regenerates only the FFI companions. Run it before a direct
`spago build` after changing the Go source or starting from a fresh checkout.
When the source is unchanged, regeneration preserves the FFI file's timestamp
so Spago can reuse its compiled output. The npm `prepare` hook runs the JS
build and includes this step automatically. Generated FFI and bundle files are
build artifacts; commit the Go source and build tooling.
