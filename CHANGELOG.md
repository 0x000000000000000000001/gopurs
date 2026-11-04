# Changelog

All notable changes to gopurs are recorded here. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and the project uses
[Semantic Versioning](https://semver.org/).

## [Unreleased]

## [0.1.0-rc.12]

Twelfth candidate for `0.1.0`, correcting the Nix cache inventory check that
blocked RC11 before Linux quality tests. Compiler/runtime, frontend, optimizer,
library pins and test corpus remain unchanged. Qualification and publication
are pending.

### Fixed

- Accept the empty `realisations` directory created by Nix 2.28.3 during a
  runtime-path cache export. Continue rejecting nonempty auxiliary directories,
  symlinks and unknown cache entries; no metadata is silently discarded.

## [0.1.0-rc.11]

Eleventh candidate for `0.1.0`, grouping CI recovery and documentation after
RC10's infrastructure failures. Compiler, runtime, TAST/frontend, optimizer
and library pins are unchanged. Both Linux jobs failed at the runtime-cache
export inventory check; RC11 remains unqualified and unpublished.

### Changed

- Reuse the build's exact locked Nix runtime environment in fixture shards,
  with bounded preparation and diagnostics, instead of rebuilding that
  environment in every shard.
- Retain the actual producer attempt for each qualification artifact so a
  targeted retry can reuse successful work from the same run and source.
  Exact compiler and toolchain identities remain required; historical tags
  and their qualification results are unchanged.

## [0.1.0-rc.10]

Tenth candidate for `0.1.0`, carrying the compiler/runtime and platform choices
of RC9 with the corrected spec source pin. Both qualification attempts failed
in environment/bootstrap preparation; the complete retry passed all three
builds and eleven fixture shards but timed out before the remaining shard's
tests. RC10 is unqualified and unpublished; RC7 remains public.

### Fixed

- Pin `gopurs-spec` at `c88da162df5d9ef85d672f5f2e819bb287e9647b`:
  await cancellation of compiler subprocesses and their process groups before
  reusing an integration directory; emit native child `close` only after
  stdout/stderr delivery and `exit`. Include native lifecycle and Aff
  cancellation regressions. Keep fixture source and golden output unchanged.
- Give a whole integration case, including cold compilation, a finite default
  budget of 300 seconds before parsing explicit timeout options. Log each
  preparation/runtime phase so future failures identify the expensive command.

Local corrected-source spec tests pass in Release and Dev (74 passing,
3 pending, all nine integration cases per profile). They use a development
frontend on macOS arm64 and do not replace this candidate's full qualification.

## [0.1.0-rc.9]

Ninth candidate for `0.1.0`, carrying the same compiler, runtime and dependency
pins as RC8. Qualification failed in the macOS spec integration tests; RC9
is unpublished and its tag remains immutable. RC7 remains public.

### Fixed

- Move both the macOS build and fixture jobs to the ARM64 `macos-15` runner.
  GitHub cancelled RC8's macOS job before any steps ran during the macOS 14
  retirement schedule. The new qualification does not establish compatibility
  with macOS 14. Keep the build and fixture runner matrices consistent through
  the CI policy regression test.

## [0.1.0-rc.8]

Eighth candidate for `0.1.0`, prepared from the post-RC7 fixes and reviewed
cleanup. Its macOS job was cancelled before any steps ran because GitHub
retired that runner image; the remaining campaign was cancelled. RC8 is
unqualified and unpublished. Its signed tag remains immutable, and RC7 remains
the published candidate.

The TAST frontend remains pinned to
`105f5752b39cc0f6c99a40064295338bf57ad223`, and the backend optimizer to
`4e7de48ffd6e19480ed3de6bfe8ca8227133be13`. The supported targets remain Linux
amd64/arm64 and macOS arm64. Rust stays experimental, and the documented
compatibility limits still apply. The existing benchmark record predates this
candidate; it does not establish RC8 performance results.

### Fixed

- Keep the original producer job and artifact attempt when GitHub copies a
  completed execution into retry metadata. Require matching execution evidence
  and keep newer failed or pending executions blocking recovery.
- Accept npm's absent-to-empty DSSE `keyid` serialization without changing
  signed payloads, signature bytes or cryptographic verification requirements.
- Pin the node-fs synchronous/asynchronous `chmod` and file stream fixes, the
  native node-streams pipe implementation, and the node-http/event-emitter
  response-body fixes. Preserve the targeted regressions and add the missing
  Go `stdout` adaptation used by the HTTP tests.
- Complete the identified uuid and js-bigints source metadata and attribution
  corrections without changing their library code or existing licenses.

### Changed

- Include all 51 pinned library suites through `bin/modtest` in future full
  release qualification, in Release and Dev profiles. Reuse the native compiler,
  trim the private Go cache between libraries, retain campaign reports and logs,
  and require exact library revisions and complete coverage before finalization.
  Ordinary quick CI and existing release tags keep their current scope.
- Replace external HTTP test endpoints with a local HTTPS fixture, retaining
  certificate verification and refusal tests. Remove the unused `go get`
  resolution of `strcase` from strings-extra.
- Derive the optional quick-CI frontend cache key from its Nix derivation and
  platform. Full candidate qualification continues to bypass this cache.
- Group compiler host services under `Gopurs.Host`, move the embedded runtime
  source under `Gopurs.Runtime.Source`, and document the module map and naming.
  Add output-identity and documentation checks; quick receipts now require all
  24 gates under `linux-quick-v2`. Check test inventories before CI builds.

The next candidate must validate its own exact artifacts, 371 gates per
platform (including 51 libraries and 401 fixtures in each profile), signatures
and installation. Historical RC7 and local post-RC7 results remain supporting
evidence only; they do not close those candidate gates. RC7 reinstallation has
been checked locally, but application suitability and public rollback remain
open in the release checklist.

## [0.1.0-rc.7]

Seventh candidate for `0.1.0`. RC6 was not published: its Linux amd64 core
campaign exposed an intermittent asynchronous emission failure. The macOS and
Linux arm64 core jobs passed; fixture shards did not run. Earlier tags remain
immutable.

### Fixed

- Return pipelined worker failures as explicit results and propagate them at
  dependency joins and producer boundaries. A worker error no longer escapes
  globally while its producer is suspended, including under supervision.
  Failures remain sticky and prevent queued dependent batches from starting.
- Include emission lifecycle regression tests in quick PR qualification.
- Retry signing and draft upload using the exact successful qualification of
  the same immutable tag, preserving the original test and compiler identities.
- Use pinned Node directly for finalization, avoiding a redundant Nix setup,
  and balance Release/Dev batches across the four existing fixture shards.

## [0.1.0-rc.6]

Sixth candidate for `0.1.0`, with the compiler and dependency pins of RC5.
RC5 passed all three core jobs and all twelve fixture shards, but final
attestation failed. It was not published; its signed tag remains immutable.

### Fixed

- Add a reproducible UUID serial number to the CycloneDX dependency inventory,
  required by the GitHub attestation action.
- Recover an existing Rekor entry when the signing service accepted a request
  but its retry returned HTTP 409. Continue to verify the complete signature,
  transparency-log proof, artifact digest and GitHub workflow identity before
  accepting the recovered entry.
- Publish using the CI provenance bundle already verified for the candidate,
  preserving the same attestation for publication recovery.

## [0.1.0-rc.5]

Fifth candidate for `0.1.0`, retaining the compiler, dependencies and feature
scope of RC4. RC4 was not published: all three core jobs and ten fixture shards
passed, but a Linux runner lost communication and two macOS fixtures could not
start Spago (`ENOENT`). The signed RC4 tag remains immutable.

### Changed

- Keep CI development tools reachable through a persistent Nix profile for the
  lifetime of each job. This adds an explicit garbage-collection root instead
  of relying on platform-dependent process inspection. Garbage collection has
  not been established as the cause of the RC4 macOS failure.
- Include the working directory and searched executable paths, with existence
  checks, when a fixture command cannot start. Preserve the original failure;
  do not silently retry or skip it.

## [0.1.0-rc.4]

Fourth candidate for `0.1.0`, retaining the compiler, pinned dependencies,
features and limitations below. RC3 was not published: all three build/core
jobs passed, but fixture qualification failed. A local reproduction exposed a
race while initializing Spago's shared registry cache. The previous signed
candidate tags remain immutable.

### Fixed

- Initialize fixture dependencies once with `spago fetch` before starting
  parallel batches on a fresh runner. Keep fresh fixture workspaces, both build
  profiles, the existing resource limits and all 401 fixtures per profile.
- Preserve Spago diagnostics on a failed fixture by removing its quiet flag.
- Reuse qualified PR results after merge even when GitHub clears the run's PR
  association list, while checking the actual tested tree, parents, receipt,
  repository, run and artifact identity.
- Import cached frontend tools using their complete Nix store paths and verify
  the imported paths before recording success.

### Changed

- Run ordinary PR checks on Linux amd64 and skip heavy builds for documentation
  changes. Tags and explicit full campaigns retain all three supported targets.
- Cache the pinned frontend runtime closure for quick runs; missing or unusable
  cache entries fall back to a fresh build. Full release runs remain independent
  of this cache.
- Align active public installation examples and candidate tables with this
  version, and check tag/package/lockfile/changelog agreement before heavy builds.

## [0.1.0-rc.3]

Third candidate for `0.1.0`, retaining the compiler, pinned dependencies,
features and limitations of the candidates below. RC2 was not published:
its Linux core gates passed, but the checkout action replaced the local
annotated tag with its peeled commit and artifact preparation rejected it.
The signed RC1 and RC2 tags are retained.

### Fixed

- Select release tags explicitly in every CI checkout, preserving their
  annotated objects while retaining the default commit selection for PRs.
- Check source commit and annotated tag identity immediately after checkout.
  Verify that the local tag object matches GitHub's verified signature before
  starting the release build.

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
