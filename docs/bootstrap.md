# Bootstrap chain

This document records how the `gopurs` compiler is built from a clean source
checkout, which inputs are trusted, and which artifacts a release must publish.
It covers the compiler distribution; it does not cover application builds.

## Pinned inputs

| Input | Pin |
| --- | --- |
| Node.js | `24.8.0` (`.nvmrc`) |
| npm dependencies | `package-lock.json` |
| PureScript compiler (npm `purescript`) | `0.15.16` |
| Spago | `0.93.45` |
| esbuild | `0.28.1` |
| PureScript package set | registry `77.10.1` |
| `purescript-backend-optimizer` | commit `4e7de48ffd6e19480ed3de6bfe8ca8227133be13` |
| `gopurs-st` | commit `b915b18e254f57994f88827a4172f6022ff47166` |
| `gopurs-unsafe-coerce` | commit `0b2ab4359a40ed0d0036f3f8d2bcdc29f3e57cd9` |
| Go (FFI parser and native bootstrap) | `1.27.0` |
| Typed PureScript compiler | commit `105f5752b39cc0f6c99a40064295338bf57ad223` |
| Go library ports | 51 exact commits in [`bootstrap-sources.json`](https://github.com/0x000000000000000000001/gopurs/blob/main/bootstrap-sources.json) |
| Go Argonaut codecs support checkout | commit `bdf31cc0f7d74d510dcf53f2a8fbca3b550d8ff2` in [`bootstrap-sources.json`](https://github.com/0x000000000000000000001/gopurs/blob/main/bootstrap-sources.json) |

The npm and Spago lockfiles and `bootstrap-sources.json` are authoritative for a source installation. A
lockfile change requires regenerating the third-party inventory; see
[`THIRD_PARTY_NOTICES.md`](../THIRD_PARTY_NOTICES.md).

## Artifacts

| Artifact | Built by | Required |
| --- | --- | --- |
| `bin/gopurs.js` | `npm ci` → `npm run build` | Yes |
| `tools/ffi_gen.wasm` | Checked in; `npm run build:ffi` with Go 1.27.0 | Yes |
| `tools/wasm_exec.js` | Checked in with the FFI parser | Yes |
| `bin/gopurs-native` | `npm run build:native` | Optional native host |
| `bin/gopurs-rust` | `npm run build:rust` | Optional experimental host |

`bin/gopurs.js` is the bootstrap root. It is built from the checked-in
PureScript sources and the pinned npm/Spago dependencies; no prebuilt compiler
binary is used to produce it. The native binary is then bootstrapped by the
JavaScript compiler, the typed `purs` fork and the Go toolchain. The Rust binary
is bootstrapped separately by the Purust toolchain.

The source install does not require the Rust toolchain or a prebuilt native
compiler. If `bin/gopurs-native` is absent, the launcher falls back to the
JavaScript host and prints a notice. Run `npm run build:native` to build the
native host.

The npm tarball is restricted by the `files` allowlist in `package.json` to the
launcher, the JavaScript bundle, the runtime helpers needed by that bundle, the
public documentation, and the license and community files. Tests, generated
snapshots, private benchmark evidence, source trees and platform-specific
native binaries are excluded. `npm run test:package` fails if the surface grows
beyond the expected file or size budget.

## Clean source installation

```bash
git clone https://github.com/0x000000000000000000001/gopurs.git
cd gopurs
npm ci
```

`npm ci` installs the locked dependencies and runs the `prepare` script
(`npm run build`). It builds the runtime, compiles the compiler with the pinned
Spago/PureScript toolchain, and writes `bin/gopurs.js`.

To install into an application without building the backend there:

```bash
npm pack --ignore-scripts
# In the application directory, use the archive printed by npm pack:
npm install --save-dev --ignore-scripts /absolute/path/to/the-generated-archive.tgz
```

The archive contains the JavaScript compiler bundle and the FFI parser
artifacts. It does not contain the typed application compiler or the Go library
ports.

## Native bootstrap

```bash
./bin/setup --locked --all
GOPURS_PURS=/absolute/path/to/typed/purs npm run build:native
```

The native bootstrap uses the JavaScript compiler to compile the compiler's own
source, verifies that the typed `purs` output contains `typeTable`, `dataDecls`
and `classDecls`, generates Go, and runs `go build -trimpath` for the native
executable. It retains a failed workspace and logs; a successful workspace is
removed unless `--keep-workspace` is passed.

`bin/setup` clones the manifest's exact commits, including the published Go
QuickCheck adaptation. Existing development checkouts are preserved. For a
release, `--locked` first checks every existing checkout's commit and working
tree and rejects mismatches before cloning anything. It never resets local
work. Set `GOPURS_LIBRARY_ROOT` to an empty directory to install a separate
release library family; use the same value for the native build.

The native path resolves the optimizer from the pinned Spago checkout. A local
development optimizer is selected only through explicit `GOPURS_PBO_DIR`.
The Go bootstrap exports each library and support revision from
`bootstrap-sources.json` into its temporary workspace using the local Git
objects. Newer branches and uncommitted changes in sibling libraries remain
untouched. An explicit `GOPURS_LIBRARY_ROOT` selects those development sources
instead; they must match the compiler's runtime and FFI conventions.
The Go Argonaut codecs adaptation is part of the source release, including its
MIT license and tests, and takes precedence over an unversioned sibling copy.

The TAST compiler is a separate prerequisite, at the manifest's exact revision.
Build it from the public fork or use the Nix toolchain described in
[`nix.md`](nix.md). `GOPURS_PURS` selects its executable; without that variable,
development builds retain discovery under `../../purescript/.stack-work`.
The TAST metadata check is necessary but does not establish a compiler's source
revision; record the source revision when producing a release.

The native PBO implementation cache retains immutable modules in memory for
one build; it does not use the JavaScript host's disk spill or memory budget.
This is separate from the driver's persistent incremental cache.

## Rust bootstrap

```bash
npm run build:rust
GOPURS_RUST=1 ../gopurs/bin/gopurs --main Main # from a project with typed output/
```

`build:rust` compiles **gopurs itself** through the sibling `../../purust/purust`
backend and Cargo. The resulting `bin/gopurs-rust` generates Go, just like the
Go and JavaScript executables. It uses this repository's PureScript sources,
`purescript-backend-optimizer-gopurs`, and the `purust-*` libraries for its host
runtime. The same Go FFI parser is linked through a C archive into the executable.

The bootstrap requires the npm dependencies from the source installation,
typed `purs`, Go with cgo and a C toolchain, Cargo, Purust and its sibling
library ports. It builds threaded Rust
with release O3 and ThinLTO, then compiles and executes a fresh Go/FFI smoke test
before atomically installing the compiler. Failures retain the workspace/logs;
`--keep-workspace` retains successful builds too.
On macOS it uses the Mach-O linker bundled with the Rust toolchain, keeping
the linker compatible with rustc's LLVM bitcode version.

`GOPURS_PURS` overrides the frontend, `PURUST_DIR` the Purust checkout,
`GOPURS_NATIVE_TMPDIR` the workspace parent, and `GOPURS_RUST_OUTPUT` the binary
destination. `GOPURS_KEEP_WORKSPACE=1` also retains successful builds invoked by
library runners. Cargo honors `CARGO_TARGET_DIR`. `PURUST_JS=1` selects the JavaScript
Purust executable for this bootstrap; it has no effect on an installed gopurs
compiler's host selection. `GOPURS_JS=1` and `GOPURS_RUST=1` are mutually exclusive.

In `../gopurs-aff`, `GOPURS_RUST=1 ./bin/test` generates, builds and executes Go;
adding `-c` rebuilds this Rust-hosted compiler first.
On machines with at least 32 GiB of RAM, the launcher selects eight preparation
and PBO workers for Rust as well as Go and JavaScript. Explicit worker settings
take precedence; the automatic Go GC policy applies only to the Go compiler host.

With all three compiler executables built, `npm run test:rust` checks exact Go
output and CLI failures in sequential/parallel modes, then compiles and executes
the string, FFI, record-worker, owned-tree and JSON regression fixtures on each
host. `GOPURS_TEST_KEEP_WORKSPACE=1` retains successful fixture workspaces.

The Rust build uses PGO by default. Set `GOPURS_RUST_PGO=0` for a plain O3
ThinLTO build. See the [PGO guide](pgo.md) for training and retained profiles.

## Checksums

Generate checksums for the artifacts present in the checkout:

```bash
npm run checksums
```

The output is a sorted, machine-readable list of `sha256  path` lines. CI
uploads one `bootstrap-checksums.txt` per platform. A release must publish the
checksums together with:

- the source tag and commit;
- the npm tarball;
- the JavaScript bundle and FFI parser artifacts;
- the native compiler for each release-blocking platform, when built.

Use `node tools/bootstrap-checksums.mjs --require-all` for a release that
includes both native and Rust compilers. The flag fails if either binary is
missing.

## Reproducibility and input checks

Build from a fresh checkout using the committed package, Spago and Nix locks.
Use `npm ci` to select the pinned compiler for the backend's JavaScript bundle;
an unpinned global `purs` can produce different output. Applications and the
native bootstrap require the typed compiler fork.

The native build uses `-trimpath`. Its typed compiler, library sources and Go
toolchain must match the recorded release provenance. Rust PGO is experimental;
profile data and LLVM versions can change the resulting binary.

`./bin/setup --locked --all` installs the exact library and support revisions
from `bootstrap-sources.json`. The Argonaut codec adaptation is a separate
pinned support checkout, described above. Local sibling checkouts are not a
substitute for verifying the manifest-selected inputs.

Run the bootstrap orchestration checks with:

```bash
node --test tools/bootstrap-inputs.test.mjs tools/build-native.test.mjs
```

These cover pinned installation, refusal to overwrite changed inputs, optimizer
and codec selection, build failures and cancellation. They do not replace a
complete compiler build. For a release, repeat the clean build from its tag and
compare artifact checksums as described in [release.md](release.md). The
[quality record](quality-0.1.0.md) states the scope of the published candidate's
checks.
