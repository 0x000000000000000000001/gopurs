# 0.1.1 release checklist

Final release `0.1.1` is in preparation. Its compiler inputs and binaries come
from signed RC20; original successful checks retain their identities. The
bounded recovery must finish all affected or missing fixtures and verify the
final packages on all supported platforms before publication.

The earlier npm channel observations and checked source-preparation items
below retain their original scope. They do not qualify RC20's changed compiler.
No historical benchmark or local RC19 equality proof substitutes for a fresh
RC20 campaign. Recheck public registry state at publication.

Check an item only after recording its exact source, artifact identity and
result in the [quality record](quality-0.1.0.md). The [release procedure](release.md)
defines the commands and evidence. Historical RC14 qualification and local
development results retain their original scope.

## Source preparation

- [x] Implement and document the ordinary Go FFI contract, including explicit
  limits for callbacks, ownership, native `Foreign` and undefined.
- [x] Verify the targeted frontend diagnostic locally: missing TAST metadata is
  rejected; incompatible old callback-shaped metadata gets actionable guidance;
  nominal actions and explicit callback implementations retain their contracts.
  The rebuilt JavaScript host passed the five focused tests and the existing
  31-test FFI suite, including executable documentation examples.
- [ ] Freeze and review the exact candidate commit, release version, source
  manifest, Nix/npm/Spago locks, generated parser/runtime artifacts, changelog
  and public guide. Confirm every dependency pin is publicly retrievable.
- [ ] Verify the clean RC20 compiler builds and complete core-gate inventory
  on all three platforms, then bind their unchanged compiler bytes and inputs
  to the final package through the bounded recovery procedure.

## Candidate qualification

- [ ] Qualify the signed final tag on `linux/amd64`, `linux/arm64` and
  `darwin/arm64`: preserve original RC20 core results, complete affected or
  missing fixture shards, and verify the final packages. Compiler hosts,
  library suites, fixtures, race checks, cache behavior and output identity
  must all be covered with their actual execution identities.
- [ ] Run the FFI end-to-end and installed-package regressions against the
  freshly built candidate hosts, including Release/Dev, official JavaScript
  reference behavior, empty npm cache, cold/warm compilation and error recovery.
- [x] Build and exercise the representative application with the installed
  RC20 CI compiler, including its migrated FFI and service/resource tests:
  b8x passed 287/287. Retain the exact inputs and artifacts; final-package
  binding remains part of artifact qualification below.
- [x] Run the mandatory [cold/warm benchmark protocol](benchmark-protocol.md)
  with RC20, compare against the
  [altbak reference workloads and baselines](https://github.com/0x000000000000000000001/altbak.pub#go),
  and retain the raw results as required by the [release procedure](release.md).
  Record compiler and generated-program measurements separately. The completed
  51-case campaign retains its RC20 identity; final reuse requires unchanged
  executed compiler bytes and inputs, not a claim of new measurements.

## Artifacts and public installation

- [ ] Review the exact npm archive and three native executables, matching
  checksums, cross-platform npm archive identity, source identity, licenses,
  third-party inventory, SBOM and signed provenance. Verify packaged guides and
  helpers from the installed archive rather than the source checkout.
- [ ] Bind the observed b8x run and benchmark to the final compiler by verified
  equality of executed bytes and inputs, retaining their original identities.
- [ ] Publish the qualified final release through the release procedure, then
  verify downloaded GitHub/npm bytes, signatures and intended dist-tags.
- [ ] In fresh consumer projects, follow the published quickstart and migration
  guide with only published artifacts and pinned sources. Exercise ordinary
  FFI examples with both compiler hosts; record platform and toolchain scope.
- [ ] Validate a retained fallback for the affected consumer before any channel
  rollback. RC14 predates the new FFI contract; RC7's historical Hello rollback
  does not establish a fallback for applications using the ordinary Go FFI contract.

## Final release and announcement

- [ ] Resolve qualification failures and documentation deviations, and have
  the retained evidence reviewed against the declared compatibility scope.
- [ ] Prepare the final `0.1.1` identity and qualify its exact artifacts through
  the same release process. Earlier results do not authorize substituting rebuilt
  or modified final artifacts without validation.
- [ ] Publish and verify `0.1.1`, update the active installation commands and
  status pages, then announce the qualified scope with links to the public
  artifacts, Go FFI guide, limitations and quality record.
