# Troubleshooting and FAQ

## Check the frontend when generation fails or behaves unexpectedly

gopurs requires the typed compiler fork. RC14 can accept upstream CoreFn and
generate Go, so a successful backend command alone does not establish that the
frontend supplied the required type information. Check the `purs` first on
`PATH` and follow the [migration walkthrough](migration.md#from-a-javascript-project)
to regenerate and validate the typed output.

Source builds after RC14 check every input module before optimization. An error
containing `requires TAST metadata arrays` identifies the module and file
missing `typeTable`, `dataDecls` or `classDecls`. Rebuild the frontend output
with the pinned fork. If invoking Spago through an npm script, check for an
upstream `purs` in `node_modules/.bin`, which npm prepends to `PATH`.

## `TAST decoding did not produce every inventoried module`

The error `TAST decoding did not produce every inventoried module` means that
some inventoried inputs failed to decode. Check the frontend and regenerate
its output. From the project root, choose an unused backup directory name:

```bash
which purs
purs --version
mv output output-before-rebuild
spago build
```

The backup retains any application-managed `go.mod` and `go.sum`. Copy those
back into the new `output/` before invoking gopurs when the application has
external Go dependencies. Inspect `output/Main/corefn.json` for the
`typeTable`, `dataDecls` and `classDecls` arrays.

## `missing Go implementation for ...`

Source builds after RC14 check Go bindings referenced from the entry point
selected with `--main` before publishing generated output. The diagnostic identifies the
PureScript module, foreign symbol and expected Go source file.

Add the binding in the companion `.go` file or the configured `--ffi`
directory. For a library, select its Go port in `workspace.extraPackages`,
including transitive dependencies when needed. Functions implemented by
compiler intrinsics and unused foreign imports do not require a port.

Use `--main Your.Module` to validate an executable. Without it, gopurs retains
its batch generation behavior: modules can include entry points for other
targets, so missing bindings remain runtime stubs. Dynamic calls within native
Go FFI can also reach those stubs; exercise the application's FFI in tests.

## `FFI not implemented: log`

The reachable foreign declaration has no Go implementation. For console output,
use the `gopurs-console` port and add it as an `extraPackages` override; see the
[quickstart](quickstart.md). Stock PureScript packages expose JavaScript FFI
and are not a Go implementation.

RC14 can generate and build a binary containing this failure stub. Run the
application to exercise its FFI. Check transitive packages as well as direct
dependencies. For example, `ordStringImpl` belongs to `prelude`, which needs
its Go port when the application sorts strings. Application-owned foreign
imports need corresponding `.go` files alongside their `.purs` sources.

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
mv output output-before-rebuild
spago build
```

Choose an unused backup name and restore application-managed `go.mod` and
`go.sum` into the new `output/` before running gopurs.

```bash
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

### Which PureScript compiler should I use?

Use the pinned typed TAST fork. RC14 may accept stock `purs` output, but its
type information does not satisfy the supported compiler contract. Validate
the frontend output as shown in the migration walkthrough.

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
