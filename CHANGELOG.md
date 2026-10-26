# Changelog

All notable changes to gopurs are recorded here. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and the project uses
[Semantic Versioning](https://semver.org/).

## [Unreleased]

### Added

- First public `0.1.0` release preparation.
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
