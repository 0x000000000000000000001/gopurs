# Publication scope and compatibility policy

Status: decision record for the first public release of `gopurs`. It defines
what users may rely on. It is not evidence that the release gates in
[`todo.md`](https://github.com/0x000000000000000000001/gopurs/blob/main/todo.md) have passed.

## Release identity

| Item | Decision |
| --- | --- |
| Package | `gopurs` |
| Current candidate | `0.1.0-rc.5` |
| First final version | `0.1.0` |
| Version source | [`package.json`](../package.json) |
| Versioning | [Semantic Versioning 2.0.0](https://semver.org/) |
| Git tag | `vVERSION` (annotated and signed; verified by GitHub before release packaging) |

During `0.y.z`, a minor release may change or remove documented behavior. Patch
releases do not intentionally break the compatibility commitments below.
`1.0.0` is reserved for the point where the CLI, generated-project contract and
input-format policy are stable enough to make those commitments across minor
releases.

Pre-releases use `-rc.N`. A release candidate is published through the same
artifact and installation path as the final release.

## Compatibility commitments for `0.1.x`

The following surfaces are supported and will not change intentionally in a
patch release:

- The command name `gopurs` and the launcher at `bin/gopurs`.
- The documented compiler options: `--main`, `--ffi`, `--dev`, `--release`,
  `--rewrite-limit` and `--force`.
- Exit status `0` for a successful compilation and `1` for a driver error.
- The error contract on stderr: a single original message prefixed with
  `[gopurs] error:` after cleanup.
- The documented generated-project entry points:
  `output/<Module>/main/main.go`, the shared `output/main/main.go`, and
  building them with `go build ./main` or `go run ./main`.
- The npm package name `gopurs` and the `bin/gopurs.js` executable entry.
- The JavaScript compiler host shipped by npm and Nix, and the native Go
  compiler binaries supplied as release artifacts. Both consume the same typed
  frontend output and support the documented CLI. Each distribution must pass
  its installation and application checks before publication.
- The documented `release` and `dev` profile names and the rule that passing
  both is an error.

Improvements to generated code, diagnostics, performance and the exact set of
emitted files are allowed in patch releases. The commands documented in the
README remain valid.

## Experimental or internal surfaces

The following are not compatibility commitments for `0.1.x`:

- The exact generated file layout beyond the documented entry points, Go
  package partitioning, identifiers, import layout and emitted comments.
- The TAST/`tcorefn` schema and its compatibility with any compiler other than
  the typed fork revision pinned for the release.
- Incremental caches and state, including `output/.gopurs-cache`,
  `.gopurs-daemon`, build manifests, overlays and package caches.
- The daemon and resident modes (`GOPURS_DAEMON`, `GOPURS_RESIDENT`).
- The Rust compiler host (`GOPURS_RUST=1`) and Rust PGO. They are supported for
  development but are not required for a production installation.
- The FFI ABI, the `gopurs_runtime` Go API, closure lifetime rules and other
  interfaces used by generated code.
- Optimizer behavior, generated-code performance and benchmark results.
- `--output` and `--bundle`, which the shared argument parser recognizes but
  which do not currently change gopurs output behavior.
- Environment tuning knobs such as `GOPURS_JOBS`, `GOPURS_EMIT_JOBS`,
  `GOPURS_PBO_JOBS`, `GOPURS_PREPARE_JOBS`, `GOPURS_GO_PACKAGES`,
  `GOPURS_GO_PACKAGE_OVERLAY`, `GOPURS_BUILD_MANIFEST`,
  `GOPURS_FORCE_BUILD` and `GOPURS_ALLOC_PROFILE`. Their documented meaning is
  supported, but names, defaults and interaction may change in a minor release.

The `gopurs-*` library ports are separate distributions. Each repository's
version, license and compatibility policy governs its package.

## Supported platforms and toolchains

The target matrix for `0.1.0` is:

| Target | Compiler | Generated Go applications | Release status |
| --- | --- | --- | --- |
| `linux/arm64` | Tier 1 | Tier 1 | Release-blocking |
| `linux/amd64` | Tier 1 | Tier 1 | Release-blocking |
| `darwin/arm64` | Tier 1 | Tier 1 | Release-blocking |
| Windows | Not supported | Not supported | Out of scope for `0.1.0` |
| 32-bit targets | Not supported | Not supported | Out of scope for `0.1.0` |

Release-blocking means that the platform must pass the release checks in blocks
2 and 5. It is not a claim that those checks have already passed.

Toolchain reference:

| Tool | Supported reference | Notes |
| --- | --- | --- |
| Node.js | `24.8.0` | Build tooling and JavaScript host; other versions are not supported in `0.1.0` |
| Go | `1.27.0` | Reference for the FFI parser and application builds |
| Spago (backend build) | `0.93.45` | Matches the npm lockfile |
| Spago (applications) | `1.0.x` with YAML configuration | The tested revisions are documented in the README |
| PureScript compiler | Typed fork, exact revision pinned in release notes | Upstream `purs` is not a substitute |

Generated `go.mod` files declare Go 1.22. That declaration is not a support
claim for older toolchains: only the reference Go version is validated for
`0.1.0`. The initial CI validated the Nix packages, application checks and
development shell on `aarch64-linux`, `x86_64-linux` and `aarch64-darwin`.
The [quality record](quality-0.1.0.md) distinguishes those results from the
remaining clean-tree and tagged release gates. The development shell is a
development environment, not a production artifact.

## Licensing and attribution

`gopurs` is distributed under the MIT License. See [`LICENSE`](../LICENSE).

Third-party components in the compiler distribution:

| Component | License | Distribution |
| --- | --- | --- |
| PureScript compiler fork | BSD-3-Clause | Build input and source dependency; not vendored in this repository |
| `purescript-backend-optimizer` fork | MIT | Bundled into the compiler |
| Go `wasm_exec.js` | BSD-3-Clause | Checked in at `tools/wasm_exec.js` |
| Go runtime in `tools/ffi_gen.wasm` | BSD-3-Clause | Checked-in WebAssembly parser |
| PureScript registry packages | MIT, BSD-3-Clause, Apache-2.0 | Bundled into `bin/gopurs.js` |

The detailed inventory and license texts are in
[`THIRD_PARTY_NOTICES.md`](../THIRD_PARTY_NOTICES.md). The short attribution
summary is in [`NOTICE`](../NOTICE).

The `gopurs-*` library ports are separate distributions and retain the licenses
of their upstream PureScript packages. For example, `gopurs-prelude` is
BSD-3-Clause and `gopurs-aff` is Apache-2.0. Their own `LICENSE` files govern.

A new runtime dependency must not be added without updating
`THIRD_PARTY_NOTICES.md` and the package metadata in the same change.

## Support and deprecation

- During `0.x`, only the latest minor line receives fixes. There is no LTS
  promise before `1.0.0`.
- Security fixes may be released outside the normal cadence.
- A public surface is deprecated in the changelog and documentation before
  removal when a compatible migration exists. Removal may be immediate when
  keeping the old behavior would be unsafe.
- Support requests and defect reports use the repository issue tracker.
  Security reports follow `SECURITY.md`, added before publication.
- `gopurs` is not a drop-in replacement for the upstream PureScript compiler.
  It requires the typed compiler fork and the documented Go library ports.

## Evidence required before publication

- CI results for every release-blocking target.
- A fresh-checkout installation that uses only published artifacts.
- A regenerated third-party inventory that matches the release bundle.
- A release candidate installed and exercised through the documented npm path.
- A changelog entry that states the compatibility impact of the release.
