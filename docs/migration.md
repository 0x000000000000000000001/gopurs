# Migration and compatibility

## From a JavaScript project

Start with a working Spago application and record its JavaScript output before
changing its frontend or library configuration. The following example uses
Spago 1.0.3, registry package set 77.10.1 and the dependencies `prelude`,
`effect`, `console`, `arrays` and `strings`. Its `src/Main.purs` is:

```purescript
module Main where

import Prelude
import Data.Array as Array
import Data.String.Common (joinWith, toUpper)
import Effect (Effect)
import Effect.Console (log)

foreign import prefix :: String -> String

main :: Effect Unit
main = do
  let names = Array.sort [ "beta", "alpha", "gamma" ]
  let loud = map toUpper names
  log $ prefix $ joinWith " / " loud
  log $ "entries=" <> show (Array.length names)
```

The existing `src/Main.js` supplies its application FFI:

```javascript
export const prefix = value => "report: " + value;
```

With the upstream PureScript compiler on `PATH`, run from the project root:

```bash
spago build
node --input-type=module -e 'import { main } from "./output/Main/index.js"; main();' > expected.txt
```

The output should contain `report: ALPHA / BETA / GAMMA` followed by `entries=3`.
Use your existing application's tests and expected results when migrating a
different program.

### Select the frontend and backend

This walkthrough targets experimental `0.1.0-rc.17`. Consult the live
[quality record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/quality-0.1.0.md)
for its qualification and publication status. Open the pinned shell in the
application directory and install the candidate with the commands below.
For local builds, use the [source bootstrap](bootstrap.md#clean-source-installation).
See the [quickstart prerequisites](quickstart.md#prerequisites) for manual
toolchain requirements.

```bash
nix develop github:0x000000000000000000001/gopurs/v0.1.0-rc.17
command -v purs
purs --version
npm install --save-dev --ignore-scripts gopurs@0.1.0-rc.17
```

Use `spago build` directly in this shell for the steps below. An existing
`npm run build` script prepends `node_modules/.bin` to `PATH`, which can select
an upstream `purs` installed in the project. The frontend-output check below
detects this. To adapt an npm script, select the typed compiler inside that
script. The Nix shell supplies its executable path in `GOPURS_PURS`.

### Configure the Go library ports

Add this `extraPackages` block to the application's existing `workspace` in
`spago.yaml`. These revisions come from the candidate's
[source manifest](https://github.com/0x000000000000000000001/gopurs/blob/main/bootstrap-sources.json).

```yaml
  extraPackages:
    prelude:
      git: https://github.com/0x000000000000000000001/gopurs-prelude.git
      ref: aa220e3571b6f0ae3743861d04153a20b10ad0d0
    console:
      git: https://github.com/0x000000000000000000001/gopurs-console.git
      ref: d4931b94e7731581e761131c56844cf65c4d0b6d
    arrays:
      git: https://github.com/0x000000000000000000001/gopurs-arrays.git
      ref: e68970644734ba8420fc811337dd45beb4b2acea
    strings:
      git: https://github.com/0x000000000000000000001/gopurs-strings.git
      ref: a4189eb3ee1dc5e18d6962359fb43430ea81efa4
```

Configure overrides at the application workspace level. Include transitive
dependencies whose FFI the application reaches. In this example, sorting
strings calls `prelude`'s `ordStringImpl`, so overriding only `arrays`,
`strings` and `console` still produces a runtime failure. Use
`spago ls deps --transitive` to inspect a larger application's dependency graph
and select additional ports from the source manifest.

### Port the application's foreign functions

Save the Go implementation of `prefix` beside `src/Main.purs` as `src/Main.go`:

```go
package main

func Prefix(value string) string {
    return "report: " + value
}
```

The [Go FFI guide](ffi-idiomatic.md) covers signatures and representations
for other types. The `.js` implementation serves the JavaScript target, while
gopurs reads the `.go` implementation. RC17 checks foreign implementations in
the selected entry-point scope before emitting Go. Exercise the application's
foreign calls when checking behavior, errors and resource cleanup.

### Rebuild and compare

From the project root, move the previous output to an unused backup name, then
compile with the typed frontend:

```bash
mv output output-before-gopurs
spago build
```

The backup retains any application-managed Go dependency manifests. For a
project that already has `go.mod` and `go.sum`, restore them to the new
`output/` before invoking gopurs. The example uses only the Go standard library.

gopurs rejects frontend output without the TAST metadata arrays. Check the
frontend output explicitly when diagnosing a frontend or cache mismatch:

```bash
node --input-type=module <<'NODE'
import { readFileSync } from "node:fs";
const module = JSON.parse(readFileSync("output/Main/corefn.json", "utf8"));
for (const field of ["typeTable", "dataDecls", "classDecls"]) {
  if (!Array.isArray(module[field])) {
    throw new Error(`Missing ${field}. Select the typed purs and rebuild output.`);
  }
}
console.log("Typed frontend output verified");
NODE
npx gopurs --main Main
(cd output && go build -o app ./main && ./app) > actual.txt
diff -u expected.txt actual.txt
```

A matching transcript validates this example. Run your project's tests and
check errors, resource handling and asynchronous operations for its own FFI.

## Current baseline

RC17 is a candidate for the first final release, `0.1.0`.
The compatibility commitments and the `0.x` policy are in
[publication-scope.md](publication-scope.md).

## Version policy

- Patch releases (`0.1.x`) do not intentionally break documented CLI behavior,
  exit status, error prefix or documented entry points.
- Minor releases may change generated layout, internal caches, FFI details,
  experimental flags and defaults.
- `1.0.0` is reserved for a stable public contract across minor releases.
- The version source is `package.json`; the changelog records user-visible
  changes.

## Migrating from a local edge build

A local checkout may contain generated output and caches from a different
compiler revision. When moving to a release tag:

1. Install the pinned typed compiler fork revision for that release.
2. Regenerate the typed output with `spago build`.
3. For a cache or TAST mismatch, move `output/` aside and rebuild. Restore
   application-managed `go.mod` and `go.sum` into the new output directory
   before invoking the backend.
4. Run `gopurs --main <Module>`.
5. Build the generated application with the supported Go toolchain.

If a cache error persists after a clean frontend build, remove
`output/.gopurs-cache` and rerun. The cache is not a compatibility surface.

## Generated output

Treat generated Go as a build artifact:

- Do not edit it by hand; regenerate it.
- Do not rely on exact filenames, identifiers or package partitioning beyond
  the documented entry points.
- Rebuild it after changing the compiler, `GOPURS_GO_PACKAGES`,
  `GOPURS_GO_PACKAGE_OVERLAY` or `GOPURS_BUILD_MANIFEST`.
- Keep `go.mod` and `go.sum` under the normal Go workflow; gopurs preserves
  external `go.sum` content.

## Library ports

The `gopurs-*` libraries are separate packages. A compiler release does not
automatically make every stock PureScript library work. Use the library
revision associated with the release notes or the repository branch documented
for that release. RC17 reports missing Go FFI in the selected entry-point scope;
library behavior still needs application tests.

## Toolchain pins

Record these with any reproduced build:

- gopurs version and source commit;
- typed `purs` fork revision;
- Node.js and npm versions for the JavaScript host;
- Go version used for generated applications;
- Spago version and package set;
- the `gopurs-*` library revisions;
- cache/package-layout settings.

The [bootstrap chain](bootstrap.md) records the compiler-side pins.

## Removing experimental settings

`0.1.0` accepts experimental daemon, resident, Rust and profiling settings.
They are not required for a normal build. When upgrading:

- unset `GOPURS_DAEMON`, `GOPURS_RESIDENT` and related variables;
- prefer the default `release` profile;
- rebuild the native host with the release tag rather than reusing an older
  `bin/gopurs-native`;
- remove stale daemon directories such as `.gopurs-daemon` if they are not in
  use.
