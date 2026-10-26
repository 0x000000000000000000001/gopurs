# Known limitations

gopurs is a `0.x` backend. This page states what is intentionally not promised.
The compatibility policy is in [publication-scope.md](publication-scope.md).

## Not a drop-in PureScript compiler

- gopurs does not run the PureScript frontend. It consumes typed TAST produced
  by the compiler fork.
- A stock `purs` with the same version number is not sufficient. It does not
  write `typeTable`, `dataDecls` and `classDecls`.
- gopurs has no `--help` or `--version` handler in `0.1.0`; see the
  [CLI reference](cli.md).
- The backend writes Go, not JavaScript. Existing JavaScript FFI modules are
  not reused.

## FFI and library coverage

- Every reachable foreign declaration needs a Go implementation. Missing FFI
  fails at runtime with `FFI not implemented: <name>`.
- The stock PureScript libraries contain JavaScript FFI. Using them without a
  Go port may compile but fail when the foreign function is called.
- The `gopurs-*` library ports are separate distributions. Their versions and
  licenses are not implied by the compiler version.
- The FFI parser supports the documented Go signatures, not arbitrary
  JavaScript FFI shapes.
- `--ffi` paths containing spaces are not supported by the shared argument
  parser.

## Generated code and representation

- Native representations coexist with a tagged `Value` runtime. General
  unboxing, complete specialization and optimal representation for every type
  are not promised.
- Generated identifiers, file layout beyond the documented entry points, Go
  package partitioning and import layout can change in a minor release.
- The `gopurs_runtime` API and the FFI ABI are internal interfaces used by
  generated code. They are not a public library API.
- `--dev` deliberately changes optimization and is not the production profile.

## Platform and toolchain limits

- `0.1.0` targets `linux/arm64`, `linux/amd64` and `darwin/arm64`.
- Windows and 32-bit targets are out of scope.
- Node.js `24.8.0` is the supported reference for the build and JavaScript
  host. Other Node versions are not validated.
- Go `1.27.0` is the reference toolchain. Generated `go.mod` files declare Go
  1.22, but older Go toolchains are not a support claim.
- The native bootstrap requires the pinned TAST frontend and Go library inputs
  listed in `bootstrap-sources.json`. The Rust host additionally requires the
  Purust toolchain and its library checkouts.
- The npm package includes Go helper sources and compiles a cache helper on
  first use. A Go toolchain is therefore required even when the compiler host is
  JavaScript.

## Native, Rust and experimental modes

The following are not production contracts in `0.1.0`:

- `bin/gopurs-rust`, Purust, and Rust PGO.
- daemon/resident modes;
- incremental cache formats and package overlays;
- allocation profiling;
- `--output` and `--bundle`, which are parsed but ignored.

## Performance

- Published timings are measurements of a specific corpus, host and protocol,
  not guarantees.
- Backend timing excludes the PureScript frontend, compiler bootstrap, final Go
  build, application execution and process startup.
- Results vary with cache state, worker limits, memory limits and toolchain.
- The project does not claim multicore scaling for workloads it has not
  measured. See the [benchmark protocol](benchmark-protocol.md).

## Support

- Only the latest `0.x` minor line receives fixes.
- There is no LTS promise before `1.0.0`.
- The project is experimental and does not promise source compatibility with
  future gopurs releases or with upstream PureScript tooling.
