# Architecture overview

This is the public overview of the gopurs data flow. The detailed,
implementation-oriented map remains in [architecture.md](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/architecture.md).

## Hosts and driver

The compiler is written in PureScript and shipped in three hosts:

- `bin/gopurs.js`: JavaScript/Node bundle; the default source-install fallback.
- `bin/gopurs-native`: Go-hosted executable, bootstrapped from the JavaScript
  compiler.
- `bin/gopurs-rust`: Rust-hosted executable, built through Purust; experimental.

All hosts run the same driver and backend. The host changes the runtime that
executes the compiler, not the generated Go contract.

## Input: typed TAST

gopurs consumes the enriched CoreFn written by the typed PureScript compiler
fork. The file is still named `output/<Module>/corefn.json`; the `tcorefn` name
describes the format, not a different filename.

The backend requires the typed fields, including `typeTable`, `dataDecls` and
`classDecls`. A stock `purs` produces ordinary CoreFn and is rejected before Go
generation. The compile-time frontend is therefore a separate step:

```text
purs + Spago -> output/<Module>/corefn.json -> gopurs -> output/*.go
```

## Pipeline

1. **Load and freeze.** The driver inventories the typed output, verifies file
   contents, copies a consistent input set into a private staging area and
   decodes it in dependency order.
2. **Prepare.** It collects constructors, class members, global types, FFI
   candidates, record fields and other metadata needed for native layouts.
3. **Monomorphize.** It specializes supported higher-rank/type-class uses and
   propagates the transitive specialization requirements.
4. **Optimize.** The shared backend optimizer runs with Go-specific handling of
   calls, tail recursion, records, arrays, effects, JSON plans and workers.
5. **Emit.** Code generation chooses native Go representations where the types
   allow it, falls back to a tagged `Value` representation otherwise, and
   prints modules, imports and FFI bridges.
6. **Publish.** Generated files are written through the content-verified cache.
   The runtime, `go.mod` and entry points are added when the application is
   assembled.

The exact pass order and metadata provenance are documented in
[architecture.md](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/architecture.md).

## Generated output

| Path | Contents |
| --- | --- |
| `output/go.mod` | Module definition for the generated application. |
| `output/gopurs_runtime/runtime.go` | Go runtime embedded by the compiler. |
| `output/purescript/*.go` | Generated application modules and FFI bridges. |
| `output/<Module>/main/main.go` | Per-target entry point. |
| `output/main/main.go` | Shared entry point for the last discovered target. |

Applications are built with the normal Go toolchain:

```bash
cd output
go build -o hello ./main
```

gopurs does not run the final `go build`; keeping application compilation
outside the backend is deliberate and also keeps backend timing measurements
separate.

## Runtime and FFI

Generated code uses the `gopurs_runtime` package for values, closures, records,
constructors, arrays and the event loop. Where the type information permits it,
values use native Go representations; the tagged `Value` path remains the
general fallback.

The FFI bridge reads Go foreign declarations and generates wrappers that adapt
Go signatures to PureScript calls. The JavaScript host runs the parser as
WebAssembly; native hosts call the same parser directly. A reachable foreign
declaration without a Go implementation fails at runtime rather than silently
falling back to JavaScript. This is why the Go library ports are separate
packages.

## Caches

The content-verified cache lives under `output/.gopurs-cache` and may use the
package layout selected by `GOPURS_GO_PACKAGES`. Cache entries validate source
content, compiler context, FFI resolution and generated output. They are an
optimization, not a stable file format: a new minor release may invalidate or
change them. `--force` recomputes preparation, optimization and emission.

## Boundaries

- The PureScript frontend and the typed fork are outside this repository.
- The final Go application build, execution and process startup are outside the
  backend timer.
- The Go library ports are separate distributions.
- Performance results depend on the corpus, cache state, host and toolchain;
  see the [benchmark protocol](benchmark-protocol.md).
