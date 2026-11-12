# RC19 release checklist

Candidate: `0.1.0-rc.19`, in preparation on 13 November 2026. Its commit,
signed tag, full CI and artifact qualification are pending. RC18 remains
unpublished after its cleanup-fixture failure; its tag and failed evidence
are preserved. The planned first final release is `0.1.0`.

The previously recorded npm channel state is historical; recheck the public
registry before publication. Checked source-preparation items below describe
existing targeted evidence, not a successful RC19 matrix.

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
- [ ] Rebuild from a clean checkout with the matching typed frontend and run
  the complete declared quick-gate inventory without skipped tests.

## Candidate qualification

- [ ] Qualify the same signed tag on `linux/amd64`, `linux/arm64` and
  `darwin/arm64` using the full workflow: compiler hosts, library suites,
  fixtures, race checks, cache behavior and output identity.
- [ ] Run the FFI end-to-end and installed-package regressions against the
  freshly built candidate hosts, including Release/Dev, official JavaScript
  reference behavior, empty npm cache, cold/warm compilation and error recovery.
- [ ] Rebuild and exercise the representative application with this candidate,
  including its migrated FFI and service/resource tests; retain the source and
  service versions with the results.
- [ ] Run the mandatory [cold/warm benchmark protocol](benchmark-protocol.md)
  with this candidate, compare against the
  [altbak reference workloads and baselines](https://github.com/0x000000000000000000001/altbak.pub#go),
  and retain the raw results as required by the [release procedure](release.md).
  Record compiler and generated-program measurements separately; do not reuse
  an older candidate's numbers as new evidence.

## Artifacts and public installation

- [ ] Review the exact npm archive and three native executables, matching
  checksums, cross-platform npm archive identity, source identity, licenses,
  third-party inventory, SBOM and signed provenance. Verify packaged guides and
  helpers from the installed archive rather than the source checkout.
- [ ] Publish the qualified candidate through the release procedure, then
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
- [ ] Prepare the final `0.1.0` identity and qualify its exact artifacts through
  the same release process. RC19 results do not authorize substituting rebuilt
  or modified final artifacts without validation.
- [ ] Publish and verify `0.1.0`, update the active installation commands and
  status pages, then announce the qualified scope with links to the public
  artifacts, Go FFI guide, limitations and quality record.
