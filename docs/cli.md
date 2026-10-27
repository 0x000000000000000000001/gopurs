# Command-line reference

`gopurs` is a backend driver. It reads the typed CoreFn/TAST files already
written under `output/<Module>/corefn.json` and writes Go sources, a runtime and
`go.mod` under `output/`. It does not invoke the PureScript frontend and does
not build the final application binary.

## Synopsis

```text
gopurs [options]
```

Run it from the application workspace containing `output/`. The current
directory is the workspace boundary.

## Options

| Option | Behavior |
| --- | --- |
| `--main Module` | Select the entry module explicitly. Without it, every loaded module exporting `main` is discovered. |
| `--ffi directory` | Add an FFI lookup directory. Paths containing spaces are not supported by the shared argument parser. |
| `--dev` | Skip monomorphization and reduce imported inlining. Keeps incremental builds. |
| `--release` | Select the optimized behavior. This is the default. |
| `--rewrite-limit number` | Set the optimizer rewrite limit; default `10000`. |
| `--force` | Recompute preparation, optimization and emission; publish changed Go files. Also available as `GOPURS_FORCE_BUILD=1`. |

`--dev` and `--release` are mutually exclusive. Passing both is an error.

The shared optimizer parser also recognizes `--output` and `--bundle`, but they
do not change gopurs output behavior in `0.1.0`. They are not part of the
supported interface.

There is no dedicated `--help` or `--version` handler in `0.1.0`. Use this
document and the package version from `package.json`.

## Exit status and errors

- `0`: compilation succeeded.
- `1`: the driver raised an error.

After cleanup, the original error is written once to stderr with the prefix
`[gopurs] error:`. Progress and timing diagnostics also go to stderr.

## Compiler hosts

The launcher selects a host in this order:

1. `GOPURS_RUST=1`: `bin/gopurs-rust`.
2. `GOPURS_JS=1`: `bin/gopurs.js` through Node.
3. `bin/gopurs-native` when it exists.
4. Otherwise the launcher falls back to `bin/gopurs.js` and prints a notice.

`GOPURS_JS=1` and `GOPURS_RUST=1` are mutually exclusive. A missing requested
host is an error; the implicit fallback applies only when no host was requested.

All hosts run the same compiler pipeline and report the same phase names.

## Environment variables

### Supported build controls

| Variable | Range / default | Effect |
| --- | --- | --- |
| `GOPURS_JOBS` | 1–64; compiler default | TAST loading and decoding workers. |
| `GOPURS_PREPARE_JOBS` | 1–8; default 2 | Preparation workers, including transitive specialization. |
| `GOPURS_PBO_JOBS` | 1–64; default 1 | Optimizer workers. |
| `GOPURS_EMIT_JOBS` | 1–64; default 8 | Go emission workers. |
| `GOPURS_PIPELINE` | `0` or `1` | Overlap emission with preparation/optimization. |
| `GOPURS_GO_PACKAGES` | `0` or `1` | Select the package layout instead of the monolithic layout. The native launcher defaults this to `1`. |
| `GOPURS_GO_PACKAGE_OVERLAY` | `0` or `1` | Use the Go package overlay path. |
| `GOPURS_BUILD_MANIFEST` | `0` or `1` | Publish the build manifest used by the package layout. |
| `GOPURS_FORCE_BUILD` | `0` or `1` | Same effect as `--force`. |
| `GOPURS_ALLOC_PROFILE` | `0` or `1` | Enable allocation profiling and force the measured passes. Experimental. |

The native launcher raises preparation and PBO workers to `8` on machines with
at least 32 GiB of memory unless the variables are set explicitly. The
JavaScript and Rust hosts keep their own defaults.

`GOPURS_GO_PACKAGES`, `GOPURS_GO_PACKAGE_OVERLAY` and
`GOPURS_BUILD_MANIFEST` interact with the incremental cache. Keep the chosen
layout across builds; switching layouts requires a full regeneration.

### Experimental controls

The daemon and resident modes are not part of the `0.1.0` public contract:

- `GOPURS_DAEMON`
- `GOPURS_DAEMON_SERVER`
- `GOPURS_DAEMON_DIR`
- `GOPURS_RESIDENT`
- `GOPURS_RESIDENT_TRIGGER`
- `GOPURS_RESIDENT_STATUS`

The same applies to the Rust host and PGO controls, the native bootstrap and
test-only variables documented in the build guides.

### Caller environment

`GOGC`, `GOMEMLIMIT` and `GOMAXPROCS` are ordinary Go runtime settings. The
native launcher may set defaults for the Go host; it preserves explicit caller
values. They are not compiler flags and are not exported to the generated
application.

## Examples

```bash
# Release build of the explicitly selected entry point
gopurs --main Main

# Development build with a lower optimizer budget
gopurs --dev --rewrite-limit 5000 --main Main

# Recompute everything and publish changed files
gopurs --force --main Main

# Select the JavaScript host
GOPURS_JS=1 gopurs --main Main

# Select the Rust host (must be built)
GOPURS_RUST=1 gopurs --main Main
```

Build the generated application separately:

```bash
cd output
go build -o hello ./main
./hello
```

See [build-cache.md](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cache.md) for cache layout, migration and
concurrency, and [publication-scope.md](publication-scope.md) for the stability
policy.
