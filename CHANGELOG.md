# Changelog

All notable changes to gopurs are recorded here. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and the project uses
[Semantic Versioning](https://semver.org/).

## [Unreleased]

## [0.1.0-rc.2]

Second candidate for `0.1.0`, retaining the frontend, bootstrap dependencies,
features and known limitations listed for `0.1.0-rc.1` below. The first candidate
was not published to npm or GitHub Releases: its full macOS gate timed out while
compiling the native scanner test, and both Linux jobs reached their six-hour
limit during the fixture campaign. The signed first-candidate tag is retained.

### Fixed

- Compile the native scanner, preparation and emission functional probes without
  Go optimizations for their large compiler package, keeping the existing memory
  budget, assertions, execution deadlines, scanner stack limit and race checks.
  Release compiler binaries retain their optimized bootstrap profile.
- Include the scanner and preparation probes in the quick CI campaign so their
  compilation and execution are checked before creating another release tag.
- Publish build manifests, FFI digests and global-function sidecars from each
  cache session's private stage. Concurrent builds no longer share a temporary
  sidecar path that another writer can rename, causing an `ENOENT` failure.
- Separate full CI into core checks and four fixture shards per platform,
  reusing the exact candidate compiler. Require complete Release/Dev coverage
  and matching source, compiler and run identities before signing artifacts.
- Upload stable quality logs and reports instead of live generated workspaces,
  avoiding missing-file errors while interrupted processes clean up temporary
  compiler snapshots.

## [0.1.0-rc.1]

First release candidate for `0.1.0`, using the typed PureScript compiler at
`105f5752b39cc0f6c99a40064295338bf57ad223` and the bootstrap libraries pinned in
`bootstrap-sources.json`. The [qualification record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/quality-0.1.0.md)
distinguishes completed local checks from the remaining tag and public-install
gates. The [benchmark protocol](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/release-benchmark.md) defines the cold and
warm measurements; the historical reference is the official altbak.pub table.

### Added

- First public release candidate preparation for `0.1.0`.
- PureScript-to-Go backend driven by typed TAST.
- JavaScript, native Go and experimental Rust compiler hosts.
- Content-verified incremental cache and optional Go package layout.
- Go runtime, FFI bridge and event loop.
- Compiler bootstrap documentation, checksums and clean-install checks.
- Publication scope, limitations, migration, CLI, quickstart and troubleshooting
  documentation.
- npm package allowlist and package-surface check.
- CI for Linux arm64/amd64 and macOS arm64.
- Nix packages, apps, development shell and application checks with pinned
  frontend, Node.js and Go toolchains.
- Public, commit-pinned bootstrap library inventory and `bin/setup --locked`.
- Versioned Argonaut codec adaptation for the native bootstrap.

### Fixed

- Go `int32` FFI arguments and results preserve signed integer values through
  scalar, array and callback bridges instead of falling through to pointer
  unboxing.

- The private output-barrier cache test now copies the `c1c` module resolved by
  `build-cache/go.mod`.
- Cache transition tests pin the monolithic layout and disable pack publication
  when they inspect individual cache objects.
- The context-transition check accounts for prepared fingerprints that are
  legitimately reused within a run.

### Known limitations

- The backend remains experimental and is not a drop-in replacement for the
  upstream PureScript compiler.
- The typed compiler fork and Go library ports are separate prerequisites.
- Stock PureScript libraries expose JavaScript FFI and may fail at runtime when
  a foreign declaration is reached.
- Windows, 32-bit targets, daemon/resident modes and the Rust host are outside
  the `0.1.x` production contract.
- Performance figures are measurements, not guarantees.

### Security

- No public security release has been made yet. See `SECURITY.md`.
