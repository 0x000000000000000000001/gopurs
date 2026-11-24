# Development guide

For application work, use the [quickstart](quickstart.md). To build the compiler,
use [bootstrap](bootstrap.md). This page covers library ports, FFI tooling and
the runtime embedded in generated programs.

| Change | Start here |
| --- | --- |
| Library port | [Libraries](#libraries) |
| Application Go binding | [Writing Go FFI](ffi-idiomatic.md) |
| FFI bridge/parser | [Go FFI](#go-ffi), [parser](#ffi-parser) |
| Generated application's runtime | [Embedded runtime](#embedded-runtime) |

## Libraries

Each port has its own repository and Spago configuration. From the compiler
checkout:

1. Build gopurs using [bootstrap](bootstrap.md#clean-source-installation).
2. Run `./bin/setup --core` for the 22 core ports and supporting dependencies,
   or `./bin/setup --all` for the whole family, including adapted QuickCheck.
3. Put the [typed frontend and Spago](nix.md#typed-frontend-and-development-shell)
   on `PATH`. Check `command -v purs spago`: a version number alone does not
   prove that `purs` emits TAST.
4. Run the selected ports in isolated workspaces:

```sh
./bin/modtest --all --list
./bin/modtest prelude strings --keep-going
./bin/modtest -d prelude strings --keep-going
```

For a manual run from a library checkout:

```sh
spago ls deps --transitive --json
spago build
# Prepare any external Go dependencies before the next command.
../gopurs/bin/gopurs --main Test.Main
(cd output && go mod tidy && go run ./main)
```

Ports declaring `go-dependencies.json` need [Go dependency preparation](ffi-idiomatic.md#external-go-modules)
after `spago build` and before gopurs. Check the port's `bin/test` for additional
setup. Use Spago's `--offline` only when dependencies are cached.

After removing or renaming a module, rebuild frontend output cleanly: gopurs
still discovers any old TAST left there. Its ownership journal removes obsolete
Go output, not stale frontend modules.

Both runners print a `Results:` JSON path and retain logs. `--keep-workspace`
also keeps successful generated workspaces. Resume failures with:

```sh
./bin/modtest --resume-failed /path/to/results.json --keep-going
```

See [testing](testing.md#library-tests) for selection, reports and isolation.
Special cases: `assert` is build-only (`spago build`, gopurs without `--main`,
then `go build ./...` under `output/`); QuickCheck runs deterministic `Test.Go`
and compares JS/Go transcripts; `node-net` uses `test/Test/Main.purs` despite
having no `package.test` stanza.

### Configuration and lockfiles

- Edit `spago.go.yaml` when `spago.yaml` is a tracked link to it.
- Keep the package set, package names and resolution aliases used by the port.
  A consuming application's `extraPackages` must include its required Go
  overrides, including transitive ones; dependency workspaces do not supply them.
- Spago locks record local paths, not sibling Git commits. Reproduce both the
  lock and checkout revisions. Update tracked locks with the workspace's Spago
  version; preserve its policy for ignored locks.
- npm/Bower/Dhall files may support upstream JS workflows. `npm test` does not
  universally mean Go tests. Spago 0.20/0.21 uses Dhall; 0.93/1.x uses YAML.
- The compiler uses `npm ci` and its tracked npm lock; `prepare` builds the bundle.
  Ports with external Go imports pin them in `go-dependencies.json`. Keep the
  prepared application's `output/go.mod` and `go.sum` between builds.

## Go FFI

Implement `returnInt64` as `ReturnInt64` in a `.go` companion, or supply an FFI
lookup directory with `--ffi`. The bridge prefixes copied declarations and
adapts parsed Go signatures using TAST types. It supports both generated package
layouts selected by `GOPURS_GO_PACKAGES`.

Use the [FFI guide](ffi-idiomatic.md) for signatures, effects and ownership.
Implementation starts in [FfiBridge](../src/Gopurs/FfiBridge.purs);
[FFIIntegerReturns](../tests/passing/FFIIntegerReturns.purs) and its
[Go companion](../tests/passing/FFIIntegerReturns.go) provide executable examples.

An explicit `--main` checks reachable missing bindings before output publication.
Batch generation without it, older RC14 builds and dynamic calls inside Go FFI
can still reach panic stubs. Exercise application FFI in tests. The parser is
not a universal Go binding generator; unsupported signatures must be adapted.

## FFI parser

The parser lives in `tools/ffi-gen`: `parser.go` reads declarations, `types.go`
defines the JSON contract and `main_js_wasm.go` exposes it to JS. The runner
`tools/ffi-runner.mjs` accepts Go on stdin and returns JSON on stdout. Empty
valid input yields `[]`; parsing, transport and JSON failures report diagnostics.

After changing it, use the pinned **Go 1.27.0** and run, from the compiler root:

```sh
go -C tools/ffi-gen test ./...
npm run build:ffi
npm run build
npm run test:ffi
```

`build:ffi` rebuilds `tools/ffi_gen.wasm` and copies `tools/wasm_exec.js` from the
same Go installation. Commit the pair together. It checks the exact toolchain,
disables automatic switching and prepares both artifacts before replacing them.
The ordinary build/install uses the checked-in pair; it does not rebuild the parser.

After editing `FfiSupport.js`, rebuild the backend bundle too. Integration tests
cover typed decoding, invalid input and absent/corrupt runner files. Failures can
occur after some output files have been written. See [testing](testing.md).

## Embedded runtime

Edit the canonical [runtime/runtime.go](../runtime/runtime.go), then rebuild
each compiler host you use:

```sh
npm run build          # JavaScript
npm run build:native   # Go
npm run build:rust     # experimental Rust
```

Each build embeds the same Go source through generated `Runtime/Source` FFI
companions. Installed compilers do not read `runtime.go` at execution time.
`npm run build:runtime` regenerates only those companions; run it before a direct
`spago build` after changing the runtime or starting fresh. Unchanged source
preserves timestamps for incremental builds.

Commit the Go source and build tools, not generated FFI or bundles.
