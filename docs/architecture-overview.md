# Architecture overview

gopurs turns typed PureScript compiler output into Go source. It uses the
preserved types to choose native representations, optimize calls and generate
adapters for foreign functions. The Go toolchain then builds the application.

Start here to understand that flow and the roles of the compiler, its hosts and
the generated runtime. For implementation work, use the detailed
[architecture guide](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/architecture.md),
the [module map](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/module-map.md)
for source locations, and the
[naming guide](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/naming.md)
for the terms used in the code.

## Terminology

The phase names describe what happens to the program: **prepare** loads typed
input and derives metadata; **specialize** adapts generic code to particular
uses, a process called monomorphization; **rewrite** applies Go-specific
transformations; **emit** creates Go declarations in memory; **output** writes
files; and **publish** commits the cache.

Two other distinctions matter when reading the implementation. A **worker** is
a generated Go function, while a **job** is a unit of parallel compiler work.
A **host** runs the compiler (JS, Go or Rust); the **runtime** supports the
generated Go application.

The source modules implementing these phases include `Driver.Prepare`,
`Monomorphization` and `CodeGen`. The
[module map](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/module-map.md)
lists their current names, source paths and dependencies.

## Hosts and driver

The compiler is written in PureScript and shipped in three hosts:

- `bin/gopurs.js`: JavaScript/Node bundle; the default source-install fallback.
- `bin/gopurs-native`: Go-hosted executable, bootstrapped from the JavaScript
  compiler.
- `bin/gopurs-rust`: Rust-hosted executable, built through Purust; experimental.

All hosts run the same driver and backend. Choosing a host changes how the
compiler itself executes; the generated Go contract stays the same.

## Input: typed TAST

gopurs consumes a TAST (typed abstract syntax tree) written by the typed
PureScript compiler fork. This enriched CoreFn format is called `tcorefn`,
but the file is still named `output/<Module>/corefn.json`.

The backend requires the typed fields, including `typeTable`, `dataDecls` and
`classDecls`. They describe types and the layouts of data constructors and
type-class dictionaries. Node annotations and `TypeApp` also retain the types
chosen at calls; the backend does not have to reconstruct erased types. A stock
`purs` produces ordinary CoreFn and is rejected before Go generation. The typed
frontend runs before gopurs:

```text
purs + Spago -> output/<Module>/corefn.json -> gopurs -> output/*.go
```

## Pipeline

1. **Load and freeze.** The driver inventories the typed output, verifies file
   contents, copies a consistent input set into a private staging area and
   decodes it in dependency order.
2. **Prepare.** It collects constructors, class members, global types, FFI
   candidates, record fields and other metadata needed for native layouts.
3. **Specialize.** It specializes supported higher-rank/type-class uses and
   propagates the transitive specialization requirements.
4. **Build and rewrite.** PureScript Backend Optimizer (PBO) builds backend
   modules; the Go rewrite passes then handle calls, tail recursion, records, arrays,
   effects, JSON plans and workers.
5. **Emit.** Emission chooses native Go representations where the types allow
   it, falls back to a tagged `Value` representation otherwise, and prints
   modules, imports and FFI bridges.
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

Canonical `Data.Array.snoc` workers, including specialized copies, use a fresh
native slice and a shallow copy. Admission checks the TAST-derived worker
signature and the complete thaw/push/freeze body before replacing it. Elements
keep their existing representation and references; partial applications retain
the ordinary wrappers. Unrecognized bodies use the normal translation path.
In the current Dev profile, the `ST.run (withArray (push x) xs)` body remains
unexpanded and uses that fallback; the native worker optimization applies to
the recognized Release form.

Composite annotations do not force an allocation: an already boxed immutable
producer can retain that representation through local bindings until a native
consumer requires conversion. Literal record keys are shared before output
layout and cache planning. Small record storage is private to the runtime;
generated accessors also handle the compact representation of two-Int records.

The foreign function interface (FFI) bridge reads Go declarations and generates
wrappers that adapt Go signatures to PureScript calls. The JavaScript host runs the parser as
WebAssembly; native hosts call the same parser directly.

With `--main`, gopurs checks for missing Go implementations reachable from the
selected entry point and rejects the build before publishing output. Batch
generation without `--main` and dynamic calls inside native Go FFI can still
reach stubs that fail at runtime. There is no fallback to JavaScript: foreign
functions need Go implementations, supplied by the separate Go library ports
or the application.

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
