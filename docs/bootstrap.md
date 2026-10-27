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
| Go Argonaut codecs adaptation | versioned source in [`support/argonaut-codecs`](https://github.com/0x000000000000000000001/gopurs/blob/main/support/argonaut-codecs) |

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
npm install --save-dev --ignore-scripts /absolute/path/to/gopurs-0.1.0-rc.1.tgz
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
The Go Argonaut codecs adaptation is part of the source release, including its
MIT license and tests, and takes precedence over an unversioned sibling copy.

The TAST compiler is a separate prerequisite, at the manifest's exact revision.
Build it from the public fork or use the Nix toolchain described in
[`nix.md`](nix.md). `GOPURS_PURS` selects its executable; without that variable,
development builds retain discovery under `../../purescript/.stack-work`.
The TAST metadata check is necessary but does not establish a compiler's source
revision; record the source revision when producing a release.

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

## Reproducibility status

The 26 October 2026 clean-install check used a checkout containing only tracked
and new source files, with no `node_modules`, `output`, `.spago` or compiler
binary. `npm ci` completed and built `bin/gopurs.js` from the pinned git
dependencies; the gopurs source build did not use the local sibling checkouts.
Three clean checkouts at different paths produced the same SHA-256
`0ce3e26633fe8d81fba8db605adefec34a7579a5b4432b135056d16ee2df1b39`, and an
immediate `npm run build` in each checkout kept that hash. The release pin uses
the npm `purescript` `0.15.16` compiler; a build that used an unpinned global
`purs` produced a different bundle, so the compiler pin is part of the
reproducibility contract.

Limits of that check:

- It covers the JavaScript bundle, not the native or Rust binaries.
- The native build uses `-trimpath`. Its manifest-pinned typed compiler, library
  sources and Go toolchain must match the recorded release provenance.
- Rust PGO is experimental; profile data and LLVM versions can change the
  binary.
- The application walkthrough still requires the typed compiler fork; the
  npm-provided compiler is used to build the backend, not to compile
  applications.

The release process must repeat the clean build from the tag and compare the
published checksums. Any difference must be explained in the release notes
before publication.

## Bootstrap input verification — 26 October 2026

The manifest's 51 library commits, the optimizer commit and the TAST compiler
commit were each checked against the public repositories with `git ls-remote`.
`./bin/setup --locked --all` also verified the 51 local checkouts at those
commits with clean working trees. This establishes source availability and
identity. A subsequent `--locked --all` installation into an empty temporary
directory cloned all 51 exact revisions successfully. The complete native
compilation remains a separate release gate.

The 21 source, test and configuration files in the versioned Argonaut codecs
adaptation were initially checked byte for byte against the previous working
input. The relative compiler paths in its two JavaScript test runners were
then adjusted to the new location. The configuration declares its own workspace
to prevent discovery as part of the JavaScript compiler build. The source and
test cases are unchanged. Its original upstream MIT license is retained separately.

`node --test tools/bootstrap-inputs.test.mjs tools/build-native.test.mjs` passes
11 checks covering pinned installation, refusal to overwrite changed inputs,
the optimizer and codec selection, native build failure handling and process
cancellation. These tests exercise bootstrap orchestration; they do not stand
in for compiling the full compiler.
