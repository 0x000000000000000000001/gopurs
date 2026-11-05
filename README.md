# 🐹 gopurs

<img height="160" alt="gopurs" src="https://github.com/user-attachments/assets/b013e7c3-fac6-4ee8-9d4c-f39ac8c2c921" />

**Compile PureScript applications to Go.** Build native executables, call Go
libraries through FFI, and use Go implementations of PureScript effects and
`Aff`.

gopurs uses a [typed PureScript compiler fork](https://github.com/0x000000000000000000001/purescript).
The fork preserves the types the backend needs to generate Go. Applications
must be compiled with this fork; the standard PureScript compiler does not
produce the required input.

The project is **experimental**. The published prerelease is
[`0.1.0-rc.7`](https://github.com/0x000000000000000000001/gopurs/releases/tag/v0.1.0-rc.7);
no stable release has been published. See the [known limitations](docs/limitations.md)
and [release qualification record](docs/quality-0.1.0.md) before adopting it.

## Getting started

Start with the **[quickstart guide](docs/quickstart.md)**. It walks through the
toolchain, library setup and a complete Hello World application.

You'll need Node.js, Go, Spago and the typed `purs` fork. The guide specifies
the supported versions. [Nix](docs/nix.md) provides a pinned development
environment for Linux amd64/arm64 and macOS arm64.

Install the published backend in your application:

```bash
npm install --save-dev --ignore-scripts gopurs@0.1.0-rc.7
```

This installs the backend; the typed frontend and Go library ports are set up
separately in the quickstart. With that project configured, compilation has
three steps:

```bash
spago build                   # PureScript → typed intermediate files
npx gopurs --main Main         # Typed intermediate files → Go sources
cd output
go build -o hello ./main       # Go sources → executable
./hello
```

gopurs writes the Go sources, runtime and `go.mod` under `output/`. You do not
need to run `go mod init`.

## What gopurs provides

- **Typed Go generation.** Preserved types and data layouts guide records,
  constructors, type-class dictionaries and specialized functions.
- **Release and development profiles.** Release is the default; `--dev`
  reduces optimization work. Both support incremental builds.
- **Go FFI.** Write Go implementations beside PureScript foreign declarations.
  The backend generates bridges for supported signatures.
- **Library ports.** Go implementations of effects, `Aff` and other libraries
  are pinned in the [source manifest](bootstrap-sources.json).

A PureScript package with JavaScript FFI needs a Go implementation before it
can run through gopurs. The [quickstart](docs/quickstart.md) demonstrates library
overrides; the [FFI guide](docs/development.md#go-ffi) explains native bindings.

The compiler itself can run as JavaScript, native Go or experimental native
Rust. All three generate Go applications. See [compiler hosts](docs/cli.md#compiler-hosts)
and the [bootstrap guide](docs/bootstrap.md) for their setup and scope.

## Find the right guide

| I want to… | Read |
| --- | --- |
| Build my first application | [Quickstart](docs/quickstart.md) |
| Use flags, profiles or compiler hosts | [CLI reference](docs/cli.md) |
| Set up Nix | [Nix guide](docs/nix.md) |
| Understand compatibility or diagnose a failure | [Limitations](docs/limitations.md), [migration](docs/migration.md), [troubleshooting](docs/troubleshooting.md) |
| Work on the compiler or a library | [Development guide](docs/development.md), [contributing](CONTRIBUTING.md) |
| Understand the compiler pipeline | [Architecture overview](docs/architecture-overview.md) |

The [documentation index](docs/README.md) also links to the cache design,
runtime contracts, release procedure and dated engineering records.

## Development and testing

To build the JavaScript compiler from source:

```bash
git clone https://github.com/0x000000000000000000001/gopurs.git
cd gopurs
npm ci
```

`npm ci` installs the locked dependencies and builds `bin/gopurs.js`. The
launcher uses that bundle when no native compiler is present. Application
compilation still needs the typed frontend and Go library ports.

Use the [bootstrap guide](docs/bootstrap.md) for native builds, and
[testing and validation](docs/testing.md#choisir-un-contrôle-court) to select a
targeted check. The [development guide](docs/development.md) covers library
work, the FFI parser and runtime changes.

### Develop one library locally

Follow the [library development guide](docs/development.md#libraries) for
checkout layout, Go overrides, lockfiles and isolated test runs.

## Benchmarks

The [altbak README](https://github.com/0x000000000000000000001/altbak.pub#go)
contains the reference workloads and results. Use the
[benchmark protocol](docs/benchmark-protocol.md) when comparing changes:
compiler time and generated-program performance are separate measurements.

## Credits and license

gopurs builds on [Arista's purescript-backend-optimizer](https://github.com/aristanetworks/purescript-backend-optimizer)
and draws inspiration from Andy Arvanitis's
[purescript-native](https://github.com/andyarvanitis/purescript-native).
The [development log](https://discourse.purescript.org/t/leveraging-a-blazing-fast-runtime-a-new-go-backend-for-purescript/5841)
records the earlier design work.

Released under the [MIT license](LICENSE). Third-party attributions are in
[NOTICE](NOTICE) and [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md).
See also the [security policy](SECURITY.md) and [code of conduct](CODE_OF_CONDUCT.md).
