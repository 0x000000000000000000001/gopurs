# Quickstart

This walkthrough compiles a small PureScript program to Go and runs the
resulting binary. It was tested on 26 October 2026 with a local gopurs `0.1.0` preview,
Node.js `24.8.0`, Go `1.27.0`, Spago `1.0.x` and the typed compiler fork.

The example uses the Go port of `console`. Stock PureScript libraries compile,
but their JavaScript FFI is not a Go implementation.

## Prerequisites

- Node.js `24.8.0` for the npm package and JavaScript compiler host.
- Go `1.27.0` to build the generated application.
- A TAST-capable `purs` from the compiler fork, first on `PATH`.
- Spago `1.0.x` with YAML configuration.
- Git.

A version number alone does not prove that `purs` is the typed fork. After
`spago build`, inspect `output/Main/corefn.json` for `typeTable`, `dataDecls`
and `classDecls`.

## Workspace

Keep the application and the Go library ports next to each other:

```text
workspace/
├── gopurs-console/
└── hello/
```

## Install the compiler

```bash
mkdir hello
cd hello
npm init -y
npm install --save-dev --ignore-scripts gopurs@0.1.0-rc.4
```

The command above requires the candidate to have been published. Until then,
install its reviewed local release archive instead:

```bash
npm install --save-dev --ignore-scripts /absolute/path/to/gopurs-0.1.0-rc.4.tgz
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

`spago.yaml`:

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

`src/Main.purs`:

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
cd output
go build -o hello ./main
./hello
```

Expected output:

```text
Hello from gopurs
```

`npx gopurs` runs the JavaScript compiler from the npm package. The compiler
writes Go sources, the runtime and `go.mod` under `output/`; `go build` produces
the final application binary.

If `purs` is not the typed fork, gopurs stops with a TAST decoding error. See
the [troubleshooting guide](troubleshooting.md).

## Minimal check without FFI

To verify the compiler without a Go library port, remove `console` and use:

```purescript
module Main where

import Prelude
import Effect (Effect)

main :: Effect Unit
main = pure unit
```

Then run the same `spago build`, `npx gopurs --main Main` and `go build` steps.
