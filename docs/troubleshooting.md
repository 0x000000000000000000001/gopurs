# Troubleshooting and FAQ

## `TAST decoding did not produce every inventoried module`

The frontend did not produce the typed format. Check that the `purs` first on
`PATH` is the typed fork, then rebuild from a clean output directory:

```bash
which purs
purs --version
rm -rf output .spago
spago build
```

Inspect `output/Main/corefn.json`: it must contain `typeTable`, `dataDecls` and
`classDecls`. A stock compiler with the same version number is not sufficient.

## `FFI not implemented: log`

The reachable foreign declaration has no Go implementation. For console output,
use the `gopurs-console` port and add it as an `extraPackages` override; see the
[quickstart](quickstart.md). Stock PureScript packages expose JavaScript FFI
and are not a Go implementation.

## `go: go.mod file not found`

Run `go build` from the generated `output/` directory:

```bash
cd output
go build -o hello ./main
```

gopurs writes `go.mod` but does not run the final application build.

## `Native compiler not found; using the JavaScript host`

This is informational. The source checkout has no `bin/gopurs-native`, so the
launcher uses `bin/gopurs.js`. Build the native host only if you need it:

```bash
npm run build:native
```

## `Error: 'node' is not installed or not in the PATH`

The JavaScript host needs Node.js `24.8.0`. Install it or select a built native
host. The native compiler does not need Node at runtime.

## `GOPURS_JS=1` or `GOPURS_RUST=1` does not select a host

The selectors are mutually exclusive, and the requested binary must exist.
Check that `bin/gopurs.js` or `bin/gopurs-rust` is present and executable. A
missing requested host is an error; only the implicit no-selector path falls
back to JavaScript.

## A cache error survives a rebuild

The cache is an optimization, not a stable format. Regenerate the frontend and
remove the backend cache:

```bash
rm -rf output .gopurs-daemon
spago build
gopurs --main Main
```

If you changed `GOPURS_GO_PACKAGES`, `GOPURS_GO_PACKAGE_OVERLAY` or
`GOPURS_BUILD_MANIFEST`, keep the new setting for subsequent builds and do a
full regeneration.

## The build is slow or runs out of memory

- Reduce `GOPURS_PREPARE_JOBS`, `GOPURS_PBO_JOBS` and `GOPURS_EMIT_JOBS`.
- Keep `GOPURS_JOBS` at or below the number of available cores.
- Use `--dev` only for iteration, not production output.
- Set an explicit `GOMEMLIMIT`/`GOMAXPROCS` for the host if the launcher
  defaults do not fit the machine.
- Build the final application separately from the compiler.

## `--ffi` path is rejected

The shared argument parser does not support spaces in argument values. Move the
FFI directory to a path without spaces.

## `--help` or `--version` does nothing

`0.1.0` has no dedicated handler. Use the [CLI reference](cli.md) and
`package.json` for the version. Unknown arguments may be ignored by the shared
parser rather than rejected.

## FAQ

### Is gopurs production-ready?

It is an experimental `0.x` backend. The supported surface and limitations are
in [publication-scope.md](publication-scope.md) and
[limitations.md](limitations.md). Use it where the Go output and library
coverage have been validated for your workload.

### Does it work with the stock PureScript compiler?

No. The backend requires the typed TAST fork. Stock `purs` output is rejected.

### Can I use stock PureScript libraries?

They may compile, but their JavaScript FFI is not a Go implementation. Use the
corresponding `gopurs-*` port when a foreign function is reachable.

### Does gopurs build the final Go binary?

No. It writes Go sources and `go.mod`; `go build` is a separate step.

### Why is there a `go build` during tests?

The test runners compile the selected entry point immediately before execution
so they never run a stale binary. That is separate from the backend compile-time
measurement.

### Does it support Windows or 32-bit targets?

No, not in `0.1.0`.

### Are benchmark numbers guarantees?

No. They depend on the corpus, cache state, host, memory limits and toolchain.
See the [benchmark protocol](benchmark-protocol.md).

### Where do I report a security issue?

Follow [SECURITY.md](../SECURITY.md). Do not open a public issue for a
vulnerability.
