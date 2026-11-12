# Quickstart

Build and run a PureScript program that prints `Hello from gopurs`. It uses
`gopurs@0.1.0-rc.16` and the Go port of `console` for `Effect.Console.log`.
This guide targets the experimental RC16 candidate. Consult the live
[quality record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/quality-0.1.0.md)
for its qualification and publication status. The
[source bootstrap](bootstrap.md#clean-source-installation) covers local builds.

## Prerequisites

Start in an empty directory that will hold the example and its library port.
If you use Nix with flakes enabled, open the development shell for this release
and run the rest of the guide inside it.

```bash
nix develop github:0x000000000000000000001/gopurs/v0.1.0-rc.16
```

The first shell setup can download and compile the Haskell frontend and its
dependencies. Let that setup finish before entering the application commands.

The shell supplies the tools below. For a manual setup, put them on `PATH`
before continuing.

- The [pinned typed `purs` compiler](https://github.com/0x000000000000000000001/purescript/tree/ae6a6f644502bafcfa4fa94d8f81a1a1a1ca5080). Spago uses it to emit the type information gopurs needs.
- Spago `1.0.3` for the application build.
- Node.js `24.8.0` and npm for the gopurs JavaScript compiler host.
- Go `1.27.0` and Git.

See the [Nix guide](nix.md#typed-frontend-and-development-shell) for details of
the pinned toolchain. The [qualification record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/quality-0.1.0.md#public-publication-and-consumer-checks)
covers the public RC14 Hello tested on Linux arm64 on 5 November 2026. RC16
requires consumer checks with its own published artifacts.

## Workspace

Keep the application and the Go library ports next to each other:

```text
workspace/
├── gopurs-console/
└── hello/
```

## Install gopurs in the application

```bash
mkdir hello
cd hello
mkdir src
npm init -y
```

Install the selected candidate:

```bash
npm install --save-dev --ignore-scripts gopurs@0.1.0-rc.16
```

## Install the console port

```bash
cd ..
git clone https://github.com/0x000000000000000000001/gopurs-console.git
git -C gopurs-console checkout --detach d4931b94e7731581e761131c56844cf65c4d0b6d
cd hello
```

This is the console revision pinned by the candidate's bootstrap manifest.

## Create the project

In `hello/`, create `spago.yaml` with these contents.

```yaml
package:
  name: hello
  dependencies:
    - prelude
    - effect
    - console

workspace:
  packageSet:
    registry: 77.10.1
  extraPackages:
    console:
      path: ../gopurs-console
```

Save the program as `src/Main.purs`.

```purescript
module Main where

import Prelude
import Effect (Effect)
import Effect.Console (log)

main :: Effect Unit
main = log "Hello from gopurs"
```

## Compile and run

```bash
spago build
npx gopurs --main Main
(cd output && go build -o hello ./main && ./hello)
```

The executable prints the following line.

```text
Hello from gopurs
```

`npx gopurs` runs the JavaScript compiler from the npm package. The compiler
writes Go sources, the runtime and `go.mod` under `output/`; `go build` produces
the final application binary.

Use the pinned typed fork. gopurs rejects frontend output that lacks the TAST
metadata arrays. After switching frontend revisions, rebuild from clean
frontend output to replace older type annotations, preserving `go.mod` and
`go.sum`. The [migration guide](migration.md#from-a-javascript-project)
shows how to check the frontend selected on `PATH` and its generated output.

## Minimal check without FFI

From `hello/`, remove `console` from both `dependencies` and `extraPackages`
in `spago.yaml` to verify the compiler without a Go library port. Then use
the following `src/Main.purs`.

```purescript
module Main where

import Prelude
import Effect (Effect)

main :: Effect Unit
main = pure unit
```

Then run the same `spago build`, `npx gopurs --main Main` and `go build` steps.
