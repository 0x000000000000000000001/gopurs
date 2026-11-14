# 🐹 gopurs

<img height="160" alt="gopurs" src="https://github.com/user-attachments/assets/b013e7c3-fac6-4ee8-9d4c-f39ac8c2c921" />

**Compile PureScript applications to Go.** gopurs generates Go source that you
can build into a native executable. Your PureScript code can call Go libraries
through the foreign function interface (FFI).

<!-- gopurs:source-version:start -->
For installation, start with the [latest published release](https://github.com/0x000000000000000000001/gopurs/releases/latest)
and the guides at its tag. This guide targets the version in this checkout:
`gopurs@0.1.1`.
<!-- gopurs:source-version:end -->

The supported targets are Linux amd64, Linux arm64 and macOS arm64.
Consult the [quality record](docs/quality-0.1.0.md) and
[release checklist](docs/release-checklist.md) for qualification evidence.

## Getting started

**[Build your first Hello World](docs/quickstart.md)**. Start in an empty
directory and follow the guide until your executable prints `Hello from gopurs`.
It includes the project files and the Go port of `console` used by the example.

You'll need the [typed PureScript compiler fork](https://github.com/0x000000000000000000001/purescript)
because gopurs uses the type information it emits. The quickstart shows how to
get that compiler, Spago, Node.js and Go together in a Nix shell, and lists the
versions for a manual setup.

### Compile an existing application

Follow the [JavaScript migration walkthrough](docs/migration.md#from-a-javascript-project)
to select the typed frontend, configure library ports and add Go implementations
of your application's FFI. The walkthrough includes a check of the emitted
type information and comparison with the application's JavaScript output.

Once those steps are complete, build from your Spago project directory.
This example assumes your entry module is `Main`.

<!-- gopurs:install:start -->
```bash
npm install --save-dev --ignore-scripts gopurs@0.1.1
spago build                 # Compile PureScript with the typed fork
npx gopurs --main Main      # Generate Go source and go.mod in output/
(cd output && go build -o app ./main && ./app)
```
<!-- gopurs:install:end -->

## What gopurs provides

- **Go library support.** The quickstart uses `gopurs-console` to implement
  `Effect.Console.log` in Go. For your own foreign functions, see the
  [Go FFI guide](docs/ffi-idiomatic.md). Existing ports and their
  revisions are listed in the [source manifest](https://github.com/0x000000000000000000001/gopurs/blob/main/bootstrap-sources.json).
- **Typed Go generation.** The backend uses PureScript type information when
  generating records, constructors, type-class dictionaries and specialized functions.
- **Build profiles.** Release is the default. Use `--dev` to spend less time
  optimizing while you iterate. Both profiles support incremental builds.

## Find the right guide

| I want to… | Read |
| --- | --- |
| Build my first application | [Hello World quickstart](docs/quickstart.md) |
| Set up the pinned toolchain | [Nix guide](docs/nix.md#typed-frontend-and-development-shell) |
| Use compiler flags or select a compiler host | [CLI reference](docs/cli.md) |
| Check whether a project will work with gopurs | [Known limitations](docs/limitations.md) and [migration guide](docs/migration.md) |
| Diagnose a build or runtime error | [Troubleshooting](docs/troubleshooting.md) |
| Understand how the compiler works | [Architecture overview](docs/architecture-overview.md) |

The [documentation index](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/README.md) also covers the runtime,
build cache and release process.

## Development and testing

To work on gopurs itself, clone the repository and build the JavaScript compiler.

```bash
git clone https://github.com/0x000000000000000000001/gopurs.git
cd gopurs
npm ci
```

`npm ci` builds `bin/gopurs.js`. The [bootstrap guide](docs/bootstrap.md)
covers the native Go and experimental Rust compiler hosts. All three hosts
generate Go applications. See [contributing](CONTRIBUTING.md) and
[testing and validation](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md#choisir-un-contrôle-court) for development work.

### Develop one library locally

The [library development guide](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/development.md#libraries) covers
checkout layout, Go overrides, lockfiles and isolated test runs.

## Benchmarks

The [altbak README](https://github.com/0x000000000000000000001/altbak.pub#go)
contains the reference workloads and results. The [benchmark protocol](docs/benchmark-protocol.md)
explains how to measure compiler time and generated-program performance separately.

## Credits and license

gopurs builds on [Arista's purescript-backend-optimizer](https://github.com/aristanetworks/purescript-backend-optimizer)
and draws inspiration from Andy Arvanitis's
[purescript-native](https://github.com/andyarvanitis/purescript-native).
The [development log](https://discourse.purescript.org/t/leveraging-a-blazing-fast-runtime-a-new-go-backend-for-purescript/5841)
records the earlier design work.

Released under the [MIT license](LICENSE). Third-party attributions are in
[NOTICE](NOTICE) and [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md).
See also the [security policy](SECURITY.md) and [code of conduct](CODE_OF_CONDUCT.md).
