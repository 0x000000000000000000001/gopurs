# Quality-gate status — 0.1.0

This is a working record for the `0.1.0` release gates. It is not a release
certificate. The gate list is in [`todo.md`](../todo.md).

## Retained RC7 recovery check — 1 November 2026

The 45 retained RC7 assets still match the public GitHub release by name,
size and SHA-256. All three checksum manifests verify, and the npm registry
tarball is byte-identical to the retained archive. The observed registry has
`latest` and `next` at `0.1.0-rc.7`; its only other version is npm's
`0.0.0-stage` staging placeholder, which is not a recovery version.

The archive and macOS arm64 native executable passed provenance signature
verification, constrained to the CI workflow, immutable tag `v0.1.0-rc.7`
and source commit `fe79595b49a2d1ed81b034533ab7c643cc69438a`.
Their SHA-256 values are:

| Artifact | SHA-256 |
| --- | --- |
| `gopurs-0.1.0-rc.7.tgz` | `32e8888b46c9bfb448bc7767fd357942f99b1f27d69640dc4ac94d0472802bca` |
| `gopurs-0.1.0-rc.7-darwin-arm64` | `eb6948f46c55c294144efe56aee403d9b0ee9eeebe1244e574ab63aff63d313d` |

`tools/release-smoke.mjs` from main
`8a40ad11272ab2e149f8e050d735634ab488945c` reinstalled that archive offline
in a fresh directory with empty npm and Go caches. Both the installed JavaScript
entry point and the launcher using the retained native executable compiled and
ran the minimal TAST/Go FFI application, each printing
`gopurs release smoke: 42`. Exit code was 0; total measured duration was
8.015 seconds. Node 24.8.0, npm 11.6.0 and Go 1.27.0 were used on macOS arm64
with the local TAST development frontend identified in the clean-command
record below, not a certified build of the release frontend pin.

Evidence is retained in `.gopurs-evidence/rollback-readiness-20261101` at the
workspace root: the inventory and current dist-tags in `state.json`, provenance
verification outputs, and `smoke.json` with SHA-256
`fb24bac4bf2d6817c58890134f3d84c8759fb2020147ab57a83b4c1860c69bd7`.
The original archives and earlier qualification evidence are unchanged.

This establishes local reinstallation readiness, not a public rollback or
multi-platform recovery qualification. No registry tag was moved. RC7's own
bootstrap lacks the later library fixes; the minimal smoke does not test those
libraries or establish suitability for an application that needs those fixes.
Before choosing RC7 as a fallback, verify the affected application's exact
dependency pins and behavior. There is currently no previous public stable
version, and the real channel restoration gate remains open.

## Local clean-command checks — 1 November 2026

The requested `bin/test -c` and `bin/modtest -c` paths were exercised on macOS
arm64 from main commit `70c3fb8dfba2108303c27fa9f33f121693623ebc`, in an isolated
checkout. Node 24.8.0, Spago 0.93.45 and Go 1.27.0 were used with the existing
local TAST development frontend, whose binary SHA-256 was
`efac5816a482d5581cfc990892f8b24f1d3343f8df1798572ea1a0a96bd07991`.
That frontend is not claimed as a verified release-pin build.

All **401 fixtures** and **51 libraries** passed in the default Release profile,
with the following interruptions and exact resumes preserved:

| Campaign | Completed coverage and recovery | Executed time across attempts |
| --- | --- | ---: |
| `bin/test -c`, then `--resume-failed` | 119 successes before a missing shared Go cache entry; 220 more with a private cache before disk exhaustion; 62 remaining successes after deleting only that regenerable cache. No compiler or fixture source changed. | 6,971.475 s |
| `bin/modtest -c`, then `-c --resume-failed` | 26 successes before the missing HTTP test stdout adaptation; 25 remaining successes after the test-only correction below. | 958.552 s |

The failed attempts returned 1; the final resumes returned 0. The successful
target sets are disjoint and cover the complete selections. Timings include
failed targets and native rebuilds, exclude pauses between attempts, and are
neither uninterrupted-run timings nor a performance comparison. Three native
rebuilds actually ran: the initial fixture campaign and both module attempts.
Each produced exactly SHA-256
`50c9e42da5eb7e9bc696fe5a7d486df70ffa008ec4f21d047644a8e9b4e04f29`.
No published RC7 binary was reused, compiler/layout defaults were retained,
and no snapshots were updated.

The real pipe implementation exposed an inherited test gap: `Test.Main.stdout`
had JavaScript FFI but no Go FFI, so the generated destination was a panic
placeholder. `node-http` source revision
`9b11034d49163590634801a85219f765402b1fc7`
([PR #2](https://github.com/0x000000000000000000001/gopurs-node-http/pull/2))
adds the native test constant carrying `os.Stdout` in an EventEmitter and a
regression check. Only three test files change. The original compiler,
production library sources and 26 earlier module successes remain valid;
the bootstrap now pins the corrected HTTP tests. The real HTTP application
and seven native regressions with the race detector also pass.

Evidence is in `.gopurs-evidence/clean-commands-20261031` at the workspace root.
`verification.json` and `state.json` record the five pass counts (119, 220, 62, 26, 25), exact
report/log hashes, the inventory hash of 2,457 successful target logs and source
checks. This is local macOS Release coverage with an explicit test-source pin
change; it does not replace a complete multi-platform release qualification.

## Post-RC7 integration and local library results — 31 October 2026

The following changes are merged into `main`. Each PR passed quick Linux amd64
qualification and its merged tree was checked against the tested tree.
These integrations do not change the published RC7 tag, assets or CI conclusion.

| Change | Integration and scope |
| --- | --- |
| Library qualification | [PR #13](https://github.com/0x000000000000000000001/gopurs/pull/13), merge `94dd04c88e5e4cd1556d60c50daab48c855d2acd`: future full campaigns include `bin/modtest` for 51 libraries in Release and Dev. Quick CI does not run this library campaign. |
| Finalization fixes | [PR #14](https://github.com/0x000000000000000000001/gopurs/pull/14), merge `c2356d0f976c18845b0eeb38f6a3ce96ea449b6d`: copied job metadata no longer masquerades as a new artifact producer; npm provenance accepts only the documented absent-to-empty DSSE `keyid` difference. Local validation passed 52 CI tests, 10 npm tests and replay of the recorded RC7 metadata. A new tag's finalization remains untested. |
| node-fs bootstrap | [PR #15](https://github.com/0x000000000000000000001/gopurs/pull/15), merge `66f34168ae76c5b4bbbaefb9415533f68ac99ccf`: pins the published sync/async `chmod` fix `2a8653066a61f9d2ff106668a187ba2f3bc66a9d`, merged through [node-fs PR #1](https://github.com/0x000000000000000000001/gopurs-node-fs/pull/1). The cleanup fixture now supplies the new permission pretest; all 96 runner tests passed locally. |
| Frontend cache | [PR #16](https://github.com/0x000000000000000000001/gopurs/pull/16), merge `23e3858f054ebe26491f3c008b174bdbcbfc49fd`: quick CI keys its optional frontend cache by the Nix derivation and runner platform. Actual Nix evaluations show unrelated library/PBO pins preserve the key, while frontend recipe or architecture changes invalidate it. |

The local library campaign passed **102 passages: 51 Release and 51 Dev** on
Linux arm64 using the unchanged public RC7 native binary (SHA-256
`8c31a8c65709176d80cf673698ea6c9704d43854aab39627a00fd88f09f28c5c`)
and runner commit `021635908285124713f1de1bf2873ba94b6e01ef`.
It retained 25 earlier Release successes after auditing that their compiled
closures did not depend on Node.FS, then ran 77 new passages with the corrected
node-fs pin above. All 102 logs and native reports, their hashes, the 51 pins per
profile and unchanged campaign source family were verified. The last 77-case
attempt took 1,929.790 seconds; this is not a fresh 102-case timing or a
performance comparison.

This is local Linux arm64 coverage with a changed library pin, not qualification
of the immutable RC7 bootstrap or a complete multi-platform release campaign.
`strings-extra` resolved `github.com/iancoleman/strcase` to `v0.3.0` in both
profiles via an unversioned `go get`; that observed version is not a hermetic
pin. `node-http` used public HTTP endpoints and passed both profiles. The two
unused `js-uri` overrides did not enter compilation. No b8x or altbak runner
was added or executed as part of this campaign.

Local evidence is retained under
`.gopurs-evidence/release-preparation-20261026` at the workspace root:
`modtest-integration-state.json`, `finalization-fixes-integration-state.json`,
`node-fs-bootstrap-integration-state.json` and
`modtest-libraries-20261030/attempt4/evidence/verification.json`.
The campaign report SHA-256 is
`9708affbe8f1f05c1a159d8dfad73ac2d74ae6f33511e15354149d2168190a32`.

The cache producer [run 41216211875](https://github.com/0x000000000000000000001/gopurs/actions/runs/41216211875)
passed quick Linux CI on the PR #16 merge. Its export/save steps and the `main`
cache inventory confirm key
`nix-purs-v2-Linux-X64-6rz07jhjyaj9601s4klbav1pqy9cj0s6-purescript-0.15.16.drv`,
containing frontend `/nix/store/ywn7kkp49z4awi2199syxr3crps1vma4-purescript-0.15.16`.
The saved archive is 27,971,509 bytes. Evidence is in
`.gopurs-evidence/ci-frontend-cache-20261031/state.json` and
`seed-cache-steps.log`. No warm consumer has been measured; no CI speedup is
claimed. Full/tag campaigns continue to bypass this cache.

### strings-extra targeted check — 31 October 2026

The historical `strcase` resolution above was traced to an unused `go get` in
`strings-extra/bin/test`; no source among the 51 pinned libraries imports it.
Revision `b448bb4a91086146fd430ae1cc0a370fab78cce5` removes only that command.
The public RC7 native compiler and the same runner implementation as commit
`021635908285124713f1de1bf2873ba94b6e01ef` then passed
`bin/modtest strings-extra --keep-workspace` and its `--dev` variant in fresh
isolated workspaces on Linux arm64. No `-c` compiler rebuild was needed.

Both profiles used a newly created Go module cache, `GOPROXY=off` and
`GOSUMDB=off`; each generated project reported only `gopurs/output` for
`go list -m all`. Release passed in 31.001 seconds and Dev in 20.495 seconds.
These are two targeted checks, not another 102-case or multi-platform campaign,
and the durations are not a performance comparison. PureScript dependency
metadata/toolchain caches were prepared; no fully offline bootstrap is claimed.

The native campaign reports, log hashes, tested script hash and unchanged
source-family check are retained in
`.gopurs-evidence/strings-extra-20261031/evidence/report.json` at the workspace
root, alongside the two logs and native reports. The earlier campaign evidence
and immutable RC7 artifacts remain unchanged.

### Native pipe targeted checks — 31 October 2026

The bootstrap pairs the published source revisions `node-streams`
`cf406cc9f84fbc82826157ca203cc4a2ff16d674` ([PR #1](https://github.com/0x000000000000000000001/gopurs-node-streams/pull/1))
and `node-fs` `8cd86e81d8d71d219d6bc1894e2153b78f956554`
([PR #2](https://github.com/0x000000000000000000001/gopurs-node-fs/pull/2)).
Native pipe now forwards data events to the real destination writer instead of
copying a fixed fixture. It preserves existing Net/HTTPS producers, handles
`end: false`, and reports write/close errors without a successful finish.
Default file adapters retain the opened file and its errors.

Targeted `bin/modtest` runs used the unchanged public RC7 native compiler and
runner implementation `021635908285124713f1de1bf2873ba94b6e01ef` on Linux arm64:

| Library | Release | Dev |
| --- | ---: | ---: |
| node-streams | 44.940 s | 26.491 s |
| node-fs | 35.193 s | 22.894 s |

The Go module cache was fresh with `GOPROXY=off` and `GOSUMDB=off`; each graph
contained only `gopurs/output`. Prepared PureScript registry/toolchain caches
were reused. These timings are not a performance comparison or a new full
campaign. Source families stayed unchanged during their runs, and implementation
and test files match the committed inputs; README/CI metadata was finalized
afterward and the FS lockfile was copied from its successful run.

The native regression failed before the fix and passes with the race detector.
It covers binary chunks, empty/multichunk real files, `end: false`, short writes,
opening/read/write/close errors, duplex half-close and concurrent completion.
The FS test waits for destination finish and requires two exact success markers.
Replacing only pipe with a no-op in a separate generated application causes its
timeout to fail the test. Five upstream JavaScript streams suites and the changed
FS `Test.Streams` pass. The full upstream FS suite stops at its existing W_OK
assertion when run as root in Docker; it is not reported as passing.

Evidence is retained under `.gopurs-evidence/stream-pipe-20261031` at the workspace
root: `initial-evidence` records the two streams successes and an initial FS test
typing error; `final-evidence` records the corrected FS successes and negative
probe. Log/native-report hashes were verified. The cleanup fixture's 52 checks
also pass with the new native pretest.

General unpipe, pause/resume, read/pipeline, transforms and options/fd file
constructors are outside this correction. A native file reader may continue to
EOF after detachment. This does not requalify RC7 or the full 51-library matrix.

### Local HTTPS targeted checks — 31 October 2026

The public Pursuit/httpbin requests in `node-http` are replaced by an independent
Node HTTPS fixture on an ephemeral loopback port. Its public test certificate
is trusted only in the child process; TLS verification remains enabled. The
PureScript tests assert status 200, the complete body and both Set-Cookie values,
and report explicit failures. Altered bodies, cookies and status codes, plus an
untrusted certificate, each produced a nonzero exit in native negative checks.

The new body assertion exposed a real Go port defect: response headers arrived,
but no body reader produced `data` or `end`. The HTTPS wrapper now starts one
reader after both the response callback and registration of a data listener,
copies retained chunks, reports read errors and closes the body. The optional
listener notification runs after the EventEmitter lock is released. The two
public source revisions must be used together:

- `node-http`: `fe20f37626de1c6fafe904ceb1499cd705f749e1` ([PR #1](https://github.com/0x000000000000000000001/gopurs-node-http/pull/1)).
- `node-event-emitter`: `be06532c41d348cede5a185b2c32faab2009caa0` ([PR #1](https://github.com/0x000000000000000000001/gopurs-node-event-emitter/pull/1)).

Targeted `bin/modtest` runs on Linux arm64 used the unchanged public RC7 native
compiler and runner implementation `021635908285124713f1de1bf2873ba94b6e01ef`.
The final results were:

| Library | Release | Dev |
| --- | ---: | ---: |
| node-event-emitter | 41.442 s | 25.071 s |
| node-http | 39.243 s | 24.288 s |

These are measured test durations, not a performance comparison. Sources stayed
unchanged during each run; report and log hashes were checked. Go module
acquisition was disabled with fresh module caches; each generated module graph
contained only `gopurs/output`. Existing PureScript registry/toolchain caches
were reused. No compiler rebuild or full 102-passage campaign was required.
The emitter's later CI-only commit changes no implementation or test source.

Native regressions cover all four listener-add variants, ordering, once and
reentry behavior, plus six response cases including empty bodies, read errors,
retained chunks and concurrent startup. They also passed with Go race detection.
Both libraries' JavaScript tests passed separately with upstream registry
77.7.0 dependencies; the Go-specific Aff FFI is not a JavaScript runtime.
The cleanup fixture was updated for the two new pretests; all 96 runner tests
passed on the host, and 11 bootstrap/module-quality tests passed. Four process
termination checks did not complete in the dedicated container; the host run
is the complete runner result, not that container attempt.

This qualifies the targeted library behavior locally, not the complete Node
stream API: pause/resume, piping and pull-based reads remain outside this fix.
It also does not qualify the pre-existing scheduling of initial request
`response/error` events. The historical 102 results and immutable published RC7
remain unchanged; this is not a multi-platform release qualification.
Evidence, identities and hashes are retained in
`.gopurs-evidence/node-http-local-20261031/state.json` and `verification.json`
at the workspace root.

## RC7 qualification and operator recovery — 30 October 2026

The original attempt of tagged run
[41041948153](https://github.com/0x000000000000000000001/gopurs/actions/runs/41041948153)
passed all three builds, twelve fixture shards and three qualification/signing
jobs. These results belong to clean source commit
`fe79595b49a2d1ed81b034533ab7c643cc69438a`, referenced by the verified,
immutable tag object `6c673452f4ff239e55294f7c6f6bcd40ef3cda92`
(`v0.1.0-rc.7`).

Each of Linux amd64, Linux arm64 and macOS arm64 passed 267 gates: the
100-file Node inventory, three Go race gates, native emission race checks,
package checks and all 401 Release plus 401 Dev fixtures. The experimental
Rust-host test remains the intentional exclusion. The three npm archives
match SHA-256 `32e8888b46c9bfb448bc7767fd357942f99b1f27d69640dc4ac94d0472802bca`.
The original-attempt audit verified 22 workflow archives. The recovery
preflight verified the exact 45 release assets, including 36 checksummed
files, with 16 cryptographic checks of provenance and SBOM attestations.

The **workflow conclusion remains `failure`**. Only draft preparation failed:
its first attempt created the draft, then failed because the draft was not
immediately visible. Retrying that job copied 19 successful jobs into GitHub's
attempt-2 metadata without executing them again or producing new archives.
The retry selector mistook those copies for new producers and requested a
nonexistent `release-linux-amd64-attempt-2` artifact. Neither failure invalidates
the audited original qualification; neither is a successful complete CI run.

After reviewing this distinction, the maintainer explicitly authorized an
operator-recovery exception for the original RC7 assets, on npm `next` and as
a GitHub prerelease. The tag, source and signed assets remain unchanged.
The 45 assets were uploaded to release `429140696` and its remote inventory
reverified. RC7 is now published on npm and as a
[GitHub prerelease](https://github.com/0x000000000000000000001/gopurs/releases/tag/v0.1.0-rc.7)
(30 October 2026, 15:47 UTC). A public npm download is byte-identical to the
approved archive. Both the original and registry-served provenance pass the
pinned Sigstore verifier. npm adds an empty `keyid` field to DSSE signatures;
the operator check accepts only that absent-to-empty serialization difference,
with the signed payload and signature bytes required to match exactly. Four
negative cases reject modified payloads, signatures, key identifiers and fields.

For this first publication npm also assigned `latest` to RC7 despite the
explicit `--tag next`, and rejected its removal after authentication. The
maintainer explicitly accepted this second exception: `next` and `latest`
both resolve to `0.1.0-rc.7`, so unqualified `npm install gopurs` installs RC7.
The version remains a candidate and GitHub marks it as a prerelease, not its
latest release. No stable version or validated public rollback is claimed.
The public quickstart passed on Linux arm64 with fresh npm, Spago and Go caches.
It installed `gopurs@0.1.0-rc.7` from npm, used the README-linked quickstart's
literal configuration and source, resolved real registry dependencies and the
pinned public console port, then compiled and ran the example. Its exact output
was `Hello from gopurs\n`. The backend host was the published JavaScript package;
the prerequisites were the pinned typed frontend 0.15.16, Node 24.8.0, Spago
1.0.3 and Go 1.27.0. This verifies a fresh application workspace with prepared
toolchain prerequisites, not their installation on a bare operating system.
The public b8x compilation also passed on Linux arm64: all 17 stages completed,
the typed frontend produced 2,707 modules, all generated Go packages built and
12 executables linked. It used the public RC7 npm archive and native binary,
with their published checksums and attestation verified, against the frozen
application source `3adb29c7acddf6799872bd98c95d64ae284dd34e`. The 3,242 input
files and 138 dependency source directories were checked against their retained
hashes; those dependency archives were not downloaded again. The build used
private caches and two workers. No b8x executable, test program or service ran.
This verifies compilation and linking, not application runtime behavior.

Local evidence is under
`.gopurs-evidence/release-preparation-20261026/publication-rc7-20261030/draft-recovery-diagnosis`:
`original-attempt1-audit/audit.json` retains the failed global status alongside
the successful qualification inspection; `recovery-preflight.json` records
the artifact-only checks before authorization, and `draft-recovery-result.json`
records the restored draft; `npm-public-verification.json` records public bytes
and provenance, and `github-published.json` records the public prerelease.
The later explicit approvals are recorded in
`rc7-candidate-state.json` in the enclosing evidence archive.
The public acquisition and quickstart receipts are under the same publication
directory: `public-artifacts/linux-arm64/public-acquisition.json` and
`public-quickstart-evidence/report.json` (SHA-256
`4da1fb13aef58fb882d52b08bab839c05fe1a270ba08f027d4aca8347fb19180`).
The b8x result is `b8x-public-evidence/plan.json` (SHA-256
`f731cef0abdfc1818f7b0272f1ddac22214914145681b26c5ae2d1b895f0b927`),
with stage logs, input hashes and executable hashes retained beside it.

The post-RC7 `bin/modtest` integration adds Release/Dev campaigns for all 51
library ports to future full qualification. The local Linux arm64 campaign and
its corrected node-fs pin are recorded above; a full multi-platform library
campaign has not run. Neither those local results nor the quick PR checks
extend the immutable RC7 tag's 267-gate qualification.

## RC7 CI finalization recovery — 30 October 2026

RC7 also separates qualification identity from finalization attempt. A later
signing attempt in the same tagged run may reuse exact successful build and
fixture artifacts, verified against GitHub metadata and their hashes. The
original quality identity is retained. Draft uploads can resume without
overwriting matching assets. This does not reuse results across source changes
or repair an immutable tagged workflow in place.

The RC5 jobs spent 12–16 minutes setting up Nix during finalization, which only
needed Node and npm. RC7 uses pinned Node 24.8.0/npm 11.6.0 for that phase.
The same four fixture shards now rotate their Release/Dev batches instead of
isolating profiles. Replaying recorded macOS batch durations through the
existing two-worker scheduling model estimates a reduction from 127.44 to
119.93 minutes for the slowest shard. This is a scheduling estimate, not a
measured new CI duration or a compiler performance claim. Platform coverage,
fixture inventory and resource limits are unchanged.

## RC6 worker-error race and RC7 preparation — 30 October 2026

The corrected RC6 PR #10 passed all 19 quick gates and 190 Node tests with no
skips. Its exact source tree was merged and tagged. Full tag campaign
[40956246171](https://github.com/0x000000000000000000001/gopurs/actions/runs/40956246171)
passed the macOS and Linux arm64 core jobs but failed one Linux amd64 emission
test. The fixture phase did not run and RC6 was not published.

A bounded local reproduction using unchanged emission sources and dependencies
failed once in 30 executions of the original 19-test suite. Delaying the next
producer action by 10 milliseconds also exposed the failure. A separate probe
through the production `withEmitter` API confirmed that supervision alone did
not consume the worker's error: a global exception occurred while the producer
was suspended, followed later by the expected parent rejection.

RC7 keeps worker outcomes as explicit `Either` values and rethrows failures at
controlled joins. It retains sticky failure detection, dependency ordering and
cancellation/draining. A regression exercises worker failure while the producer
is suspended, and all emission tests now run in quick PR qualification.
Dependency pins, fixture inventory and resource limits are unchanged. Full
coverage remains 100 Node files, 162 fixture batches and 267 gates; quick
coverage becomes 20 gates and 210 Node tests. This preparation required fresh
CI, artifact verification and public consumer checks; the qualification and
remaining publication checks are recorded above.

## RC5 attestation failures and RC6 preparation — 29 October 2026

Signed RC5 source `30a98ecaed56f3e0200667cb6cdd315828176ddd` was merged
through PR #9 without changing its tested tree. Full tag campaign
[40845932937](https://github.com/0x000000000000000000001/gopurs/actions/runs/40845932937)
passed all three build/core jobs and all twelve fixture shards. Finalization
failed after those tests; no RC5 package or GitHub release was published.

The npm attestation request returned HTTP 409 for an equivalent entry already
present in Rekor. That public entry was recorded seconds before the error.
Sigstore 3.1's high-level API disables conflict recovery. RC6 composes the
pinned SDK's public signing APIs with conflict recovery enabled, then retains
the existing full cryptographic and exact release/workflow identity checks.
Its local regression test exercises the actual SDK transport against a local
HTTP server and rejects altered payloads, signatures and log evidence.
The publication workflow also submits that verified CI bundle instead of
generating a different signature, so its recovery check uses the same evidence.

The GitHub action rejected the CycloneDX inventory because its format detector
requires `serialNumber`. RC6 supplies a deterministic UUIDv5 derived from the
complete inventory. Regression tests cover the action's format discriminator,
a separately calculated UUID and identity changes when dependencies change.

Compiler sources, dependency pins, fixtures and resource limits are unchanged.
The added SBOM regression file raises the full inventory to 100 Node files and
267 final gates, retaining 162 fixture batches and 401 fixtures per profile.
These targeted checks do not replace complete tagged CI, artifact verification
or the public quickstart and b8x compilation checks. RC6 is not yet qualified.

## RC4 runner failures and RC5 preparation — 29 October 2026

Signed RC4 source `ae6b715efad686905ca633e6a2d6ee0bf5fd00c5` was qualified
by PR #8 and merged without changing its tested tree. Full tag campaign
[40763342099](https://github.com/0x000000000000000000001/gopurs/actions/runs/40763342099)
passed all three build/core jobs and ten of twelve fixture shards. Linux amd64
shard 0 lost its hosted runner connection, leaving no downloadable job log or
fixture artifact. macOS shard 0 passed serial dependency preparation and most
fixtures, then two commands failed with `spawn spago ENOENT`. Finalization was
skipped. RC4 is not fully qualified and was not published.

The macOS failures occurred before Spago started, for
`TypeClassesWithOverlappingTypeVariables` and `TypeSynonyms`. Their two reduced
frontend builds passed concurrently in 7.349 seconds on Linux arm64 with the
same frontend/tool versions and fixture preparation. This does not reproduce
or explain the macOS failure. The runner opened each fixture log immediately
before the failed spawn; inspection found no shared workspace cleanup that
would account for a missing command.

A separate isolated-store experiment, without garbage collection or a build,
verified that `nix print-dev-env --profile` creates a persistent root while
leaving the returned environment unchanged. Nix 2.35.4 uses process environment
inspection on Linux and `lsof` on other systems; a profile provides an explicit
root independently of that difference. The hosted Determinate version and a
collection during the failed job have not been established. RC5 therefore
adds explicit profiles as a tool-lifetime safeguard and adds command-start
failure diagnostics; it does not claim a proven garbage-collection root cause.

Compiler sources, dependency pins, fixtures, snapshots, resource limits and
full gate coverage remain unchanged. A successful complete RC5 tag campaign,
artifact verification and public consumer checks are still required.

## RC3 fixture-cache race and RC4 preparation — 28 October 2026

Tag campaign [40727256711](https://github.com/0x000000000000000000001/gopurs/actions/runs/40727256711)
for signed RC3 commit `367163a040c90ba12f8d34c909b4c24541ea1592`
completed with all three build/core jobs successful and all twelve fixture
shards failed. The core jobs passed the annotated-tag checks, optimized native
bootstrap, partial core gates and installed-artifact smokes. Finalization and
draft creation were skipped. RC3 was not published and is not fully qualified.

The examined failures occurred at `spago build -q`, before backend generation.
The same fixture could pass elsewhere, so this was not evidence of a deterministic
backend regression. A bounded Linux arm64 reproduction used the exact RC3
fixture preparation, pinned library sources, public TAST frontend, Node and
Spago tools. One cold fixture passed in 15 seconds. Starting the first three
frontends concurrently against a new shared XDG cache reproduced two failures:
Spago could not find its registry package-sets directory while another process
was initializing it. The third frontend passed.

A separate experiment used another entirely new XDG cache: serial `spago fetch`
completed in 24.608 seconds without creating compiled output, then three new
fixture workspaces passed in parallel in 7.399 seconds. It executed neither the
backend nor generated Go applications. The retained experiment summary has
SHA-256 `b8b1b98025fdc61a5d833e87ae831a704453347ac748695bdb58b814e64d0923`.
This proves the targeted initialization fix, not full RC4 qualification.

RC4 fetches the shared fixture dependencies before launching full fixture waves,
records the preparation result and a stable log, and stops if preparation fails.
It removes Spago's quiet flag so failed phases expose their existing captured
diagnostics. Core and quick phases, fixture coverage, snapshots, profiles and
resource limits are unchanged. The complete inventory remains 99 Node files,
162 fixture batches and 266 final gates; the planning job separately runs the
CI policy tests. A new successful tag campaign and public installations are
still required for RC4. Historical sections below retain their original scope
and status; they do not qualify this candidate.

## RC2 tag checkout failure and RC3 preparation — 28 October 2026

PR campaign [40671552573](https://github.com/0x000000000000000000001/gopurs/actions/runs/40671552573)
passed on all three platforms for head
`a75c984db2a8c829470cd852b4b3bf11b7fbabbd`, tested through merge
`95dbb23b4f9bbbc8a0c5b477508621f3e11e8d0d`. Each platform passed 18 quick
gates and 187 Node entries with no skips. The artifact audits verified clean
sources, optimized bootstraps, checksums and both installed-host smokes. All
three npm archives have SHA-256
`973abc2ea12b543f8e3fe512475e57ab9629a0e4161dab104edce897298ff234`.
The macOS artifact download required a retry after a TLS timeout; the original
failure is retained separately from the successful recovery audit. The
262-file evidence manifest has SHA-256
`d797123f1c5a369e6a4a0cf91b7fbf69039a2fb76939fe62f6246c6de8af827f`.

PR #3 was merged as `bb5656f38a893327b1f16abf8371cf32cdf55fc6` with the same
source tree. Signed tag `v0.1.0-rc.2` points to the reviewed head; its annotated
object `326d48776ec8dd222a667f350babd9af09593b39` is verified by GitHub.
In tag campaign [40683735647](https://github.com/0x000000000000000000001/gopurs/actions/runs/40683735647),
both Linux platforms passed their 104 partial core gates, the remote tag
signature check and the clean-source check. Artifact preparation then failed
with `HEAD must be the annotated tag v0.1.0-rc.2`. This does not establish full
fixture qualification or publication; RC2 remains unpublished.

The checkout logs show the cause directly: after fetching the annotated
object, `actions/checkout@v4` at
`11d5960a326750d5838078e36cf38b85af677262` force-fetched the event commit into
the local tag ref. That made the local ref a lightweight tag while the remote
signed tag remained correct. Its [ref-selection implementation](https://github.com/actions/checkout/blob/11d5960a326750d5838078e36cf38b85af677262/src/ref-helper.ts)
compares the annotated object with the event commit, then uses that commit in
the fallback refspec. An isolated Git experiment reproduced the failing
refspec and verified that explicit tag selection preserves the original tag
object and passes the existing identity checks. The original checkout and tag
were unchanged by the experiment.

RC3 selects tag refs explicitly in each CI checkout, preserves default PR
selection, and checks the source commit and local annotation immediately. The
build also compares the local object with GitHub's verified object before Nix
or compilation. Compiler sources and full qualification requirements are
unchanged. The diagnostic logs, failed-run monitor and Git experiment are
retained under `publication-rc2-20261028` within
`.gopurs-evidence/release-preparation-20261026` at the workspace root. A new
green tag campaign and verified public installations remain required for RC3.

## RC2 quick CI and full-campaign split — 28 October 2026

PR campaign [40547270158](https://github.com/0x000000000000000000001/gopurs/actions/runs/40547270158)
passed on Linux amd64/arm64 and macOS arm64. It tested head
`738265985438d9dfb1301236c9b2900f911b4980` through merge commit
`09792f8cc802dd8d39311a28062bbb0606f655cf`. Each platform passed 18 quick
gates: 13 Node files with 183 passing entries and no skips, three Go tool
commands, five representative fixtures, and the package-surface check. The
seven-entry sidecar regression, two scanner tests and native preparation test
with race detection ran on every platform. Nix checks, fresh optimized native
bootstraps, clean-source checks and both installed-artifact host smokes passed.

The three npm archives are identical, with SHA-256
`6bf4b90184a1cee74c1c1e97c8e99a4e0bf75cac08157a7b974833a216e08896`.
All six artifact ZIP digests and 12 release checksum subjects per platform were
verified. The metadata identifies clean preview builds at the tested merge,
with the public TAST frontend `105f5752b39cc0f6c99a40064295338bf57ad223` and
identical pinned inputs. The recovered completion audit re-read the existing
Linux evidence and downloaded only the missing macOS ZIPs. Its 222-file archive
is `ci-rc2-pr-40547270158` under
`.gopurs-evidence/release-preparation-20261026` at the workspace root; its
`SHA256SUMS` digest is
`b782ef4d2334487112c9577f5cb85dddc89c1792ad5a3882280e8a073a999f1a`.
This establishes quick PR qualification of that commit, before the workflow
changes below; it does not establish full tagged qualification or publication.

The completed logs of the earlier RC1 tag campaign
[40503372835](https://github.com/0x000000000000000000001/gopurs/actions/runs/40503372835)
also establish a separate Linux duration problem. Both jobs reached the
six-hour job limit during fixtures: amd64 had started wave 37 of 54, while
arm64 had started the final wave. Reaching fixtures means the controller had
successfully completed all 99 Node-file commands, the three Go tool race gates
and the native emission race gate. It does not establish individual Node
pass/skip totals or completion of the interrupted fixture wave. Both evidence
uploads then failed with `ENOENT` while archiving live workspace paths under
temporary compiler snapshots. The workflow logs and
`ci-rc1-tag/final-linux-diagnosis.json` are retained, but the missing Linux
quality artifacts prevent an independent audit of their per-test reports.
Those incomplete jobs do not qualify RC1.

The revised full workflow keeps the fresh build, all Node files, three Go race
gates, native emission race gate, package check and installed-artifact smokes in
the build/core job. Four disjoint fixture jobs per platform then consume that
job's candidate artifact, verify its checksums and reuse its native compiler
without another bootstrap. They retain the existing fixture resource limits,
snapshot comparisons and application executions. The aggregate requires every
selected fixture exactly once in each profile: currently 401 Release and 401
Dev fixtures. It checks the complete core inventory, all four shard identities,
stable per-batch reports, clean source commit, target, compiler hashes and GitHub
run identity against the original build metadata. Partial reports explicitly
use `requestedFull: true` and `full: false`; only the validated aggregate records
`full: true`.

The final quality report replaces the partial report and its checksum before
provenance and SBOM signing. Draft-release creation depends on successful
finalization of all three platforms. Evidence uploads include only stable logs
and reports, excluding generated workspaces and caches. At this stage, the
identity included one `GITHUB_RUN_ATTEMPT` for the entire campaign, so partial
retries were rejected. RC7 supersedes that finalization restriction with the
separate, verified attempt binding described above; qualification itself still
requires one consistent attempt.

Actionlint 1.7.7 passes both workflows. The focused resource/report suite passes
12/12 tests, including a real aggregation CLI invocation over 162 synthetic
batch reports and negative identity/coverage cases. A real fixture-phase pilot
then passed the first five Release fixtures in 121.4 seconds, including TAST
compilation, native generation, strict snapshots, Go builds and execution. It
used the recorded local frontend and compiler, and remains explicitly partial
(`full: false`, one of 162 diagnostic shards); it is not a full qualification.
The stable batch report is retained outside its generated workspace. Hosted
validation of this workflow revision, the full tag campaign and public-artifact
installation remain required.

## RC1 tag failure and RC2 preparation — 27 October 2026

The clean PR campaign [40491438925](https://github.com/0x000000000000000000001/gopurs/actions/runs/40491438925)
passed on all three platforms, followed by the main campaign
[40503318512](https://github.com/0x000000000000000000001/gopurs/actions/runs/40503318512)
at merge commit `c91d2f6ad6b833e44170794fa38ec74062fd26d2`. These are quick gates.
The signed tag `v0.1.0-rc.1` points to
`8361c05f0a0517d55be7ce8be7600cec198c33ee`, with the same source tree.

The full tag campaign [40503372835](https://github.com/0x000000000000000000001/gopurs/actions/runs/40503372835)
cannot qualify that candidate: macOS passed Nix, its fresh native bootstrap and
65 Node files, then failed `native-go-code.test.mjs`. Its `go build -trimpath`
timed out after 900 seconds, before the scanner executable ran. The bootstrap's
Go build had taken 709.9 seconds. No quality-cache entries had been evicted before
the failure. The signed tag and original failure evidence are retained; no npm
package or public GitHub Release was created for RC1.

Three local diagnostic builds used Go 1.27.0, the same retained native source,
new caches, two Go threads and a 120-second limit per invocation. RSS was sampled
every two seconds and Go GC traces were recorded:

| Compiler package profile | Memory setting | Build time | Peak compiler RSS | Final cumulative GC share |
| --- | --- | ---: | ---: | ---: |
| Optimized | `GOMEMLIMIT=4GiB` | 108.00 s | 7.55 GiB | 43% |
| Optimized | Default limit (`off`) | 50.95 s | 10.52 GiB | 12% |
| `-N -l` for `gopurs/output/purescript` only | `GOMEMLIMIT=4GiB` | 36.75 s | 3.96 GiB | 14% |

The first build's live Go heap exceeded its soft memory limit. This establishes
local GC pressure, not a complete diagnosis of the slower GitHub runner. The
third profile kept the library's assertions intact: the scanner's original two
tests passed in 35.9 seconds from a new cache, including JavaScript parity and
the explicit one-MiB stack limit. This is a diagnostic comparison, not a backend
performance benchmark.

A separate cold-cache experiment compiled the actual preparation and emission
Go tests together with `-race`. The optimized build reached a sampled 7.64-GiB
compiler RSS and was stopped at 120.01 seconds. The targeted `-N -l` build
finished in 67.81 seconds at 3.96 GiB; both real test groups then passed with race
detection and their original 30-second runtime deadlines. The optimized attempt
is an interrupted diagnostic, not a completed failing test campaign.

RC2 applies that profile only to the scanner, preparation and emission functional
probes. Scanner and preparation checks are included in quick CI as well, so the
next PR exercises the corrected compilations. The optimized release bootstrap
and the generated application's Release/Dev gates remain separate checks. The
new candidate still requires its own CI, signed-tag gate, provenance verification
and installation from public artifacts.

The first RC2 PR campaign [40528983021](https://github.com/0x000000000000000000001/gopurs/actions/runs/40528983021)
tested head `23e452010ea7a3f5af587f74bb967f6bc75b2ac2` through merge commit
`5c87178d163fd9faf48c54418c2185eb94f67a56`. Both Linux platforms passed all
17 quick gates, with 176 Node entries each and no skips, including the scanner
and preparation probes. Their installed-artifact smokes passed and their npm
archives were identical. Their optimized native binaries, JavaScript compiler
bundle and embedded runtime/helper bytes matched the preceding clean RC1
artifacts. These comparisons describe the source before the sidecar fix below.

macOS passed Nix, its fresh native bootstrap and nine Node files, then failed
gate 10, `driver-cache.test.mjs`, in the `native-sequential` subtest. Publishing
`ffi-digests-v1.json` failed because its shared `.next` file had already been
renamed by another writer. The failing backend invocation took 3.685 seconds;
this was an `ENOENT` race, not a compilation timeout. The JavaScript and native
parallel subtests passed. The scanner and preparation probes, subsequent gates,
clean-source check, packaging and smokes were not reached on macOS. The failed
campaign and its original logs are retained in `ci-rc2-pr` under the evidence
root, with `SHA256SUMS` digest
`902e87dbaf2e08ec0cf92e1bb077e3ca250c626001794f44e4bc882e8e939062`.

A focused reproduction called the real JavaScript-hosted manifest writer with
20 concurrent publishers: one succeeded and 19 failed with `ENOENT` in about
2 ms. Build manifests, FFI digests and global-function sidecars now write their
temporary files in the owning cache session's private stage before renaming
them into the shared cache. The regression runs 20 concurrent real Aff writers
for each sidecar, requires every write to succeed, decodes a complete published
payload, and checks that the stages are empty and no shared temporary remains.
The updated JavaScript build and the regression file pass: seven Node entries,
including its three concurrent-write subtests. Quick CI now also runs this file
before the end-to-end driver-cache check. A fresh optimized native bootstrap
passes with 522 typed modules and 326,579 types. This local diagnostic uses the
recorded frontend binary with SHA-256
`efac5816a482d5581cfc990892f8b24f1d3343f8df1798572ea1a0a96bd07991`;
it does not establish identity with the public frontend used by CI. The complete
driver-cache scenario then passed on JavaScript, native sequential and native
parallel hosts: four Node entries, zero failures or skips, in 275.7 seconds.
Each host ran the existing pair of simultaneous forced compilations and checked
that their generated output still matched clean generation. The subsequent
hosted quick validation is recorded above. Historical timings, compiler hashes
and frontend identities below retain their pre-fix provenance.

The local diagnostic scripts, source copy, command reports, RSS samples and
logs are archived under `.gopurs-evidence/release-preparation-20261026/rc1-go-timeout-probes`
at the workspace root. Its `SHA256SUMS` digest is
`b3e3c93d1cce235fd37a04ff21b08d31ab2d78cc7ba4d39140e302b625aaa5f7`.
The tag failure archive is `ci-rc1-tag` under the same evidence root.
The separate race experiment is archived in `rc1-race-memory-probes`, including
the initial sandbox-rejected attempt and the completed second attempt. Its
`SHA256SUMS` digest is
`f839fa30bb46d769eb46bbcc9aa5113608186e5e5a5fa739fc483f2a55547d63`.

## Initial three-platform CI — 27 October 2026

PR CI run [40476860720](https://github.com/0x000000000000000000001/gopurs/actions/runs/40476860720)
passed on Linux amd64/arm64 and macOS arm64. It tested merge commit
`1157e8778595ed7e44d83ad41f14a6120c9de22b`, combining branch head
`8bbcf324b61b161f1edaa085c8a7c3994a641202` with main
`9b6e147868e2b2accda9e398ac2c739e3447263f`.

Each platform passed the Nix package/application checks, fresh native bootstrap
(522 typed modules, 326,573 types), quick quality gate, both installed-artifact
host smokes, and npm publication dry-run. The quick gate used `full: false`:
157 Node entries passed, three Go test commands ran without `-race`, and five
fixtures built and executed. Tag-only signatures and draft-release creation
were skipped. The macOS bootstrap took 34 min 38 s, including 29 min 47 s in
`go build`; these logs do not establish the hardware cause of that duration.

The three npm archives are byte-identical, with SHA-256
`a2b2cffb3c600294ed357822a371261b3be94b2bbc0fe653b95ea76f0ed1eba1`.
Their JavaScript compiler and the Darwin native binary match the previously
qualified local compiler hashes below. Six artifact ZIPs matched their GitHub
digests; 36 release checksums and nine internal bootstrap checksums passed.

All build metadata records `preview: true`, no tag and `source.dirty: true`.
The tested merge differs from the branch only by a tracked `node_modules`
symlink inherited from main; npm logs record its replacement. The four
source/lock manifests match the branch exactly. This run therefore does not
establish a clean-tree release gate. Local merge `3925dc1` integrates main and
removes that inherited tracked path; its source tree otherwise exactly matches
the preceding release commit. The working checkout's ignored dependency
directory was preserved. CI now explicitly rejects a dirty tree after builds
and tests; the updated branch still requires its own CI run.
Full tag quality, signing/provenance and public publication remain separate gates.

Metadata, original ZIPs, complete workflow logs and repeatable checksum audits
are archived at the local workspace root under
`.gopurs-evidence/release-preparation-20261026/ci-initial`: 129 files,
228,932,462 bytes. The verified `SHA256SUMS` list has SHA-256
`85a850f0342313ea12a2ee22714080ddb5ca3e76ea7790c6fe706f2b5dc70e9a`.

## Candidate tooling — 27 October 2026

The candidate version is `0.1.0-rc.1`. The bounded quality controller and the
first-publication provenance helper have 23 passing focused tests, including
the native-build orchestration with CI log streaming enabled. A failed
digest-mutation test exposed a shared subject object in the generated
statement; copying the subject fixed it, and all 23 tests then passed with
zero skips. The fixture-report verifier also accepted all 162 archived real
reports, covering the exact 401 Release and 401 Dev fixtures.

Both workflow YAML files parse, and Bash syntax checks pass for their 22
shell blocks. These checks do not exercise GitHub's OIDC service or npm's
acceptance of a provenance bundle; those remain tag/publication gates.

The quick local checks were completed across three recorded invocations:
the initial controller stopped after nine successful gates during a session
interruption; a continuation passed the driver-cache test and three Go modules,
then failed its first fixture because the manual environment selected upstream
`purs` 0.15.15 through `PATH`. The retained `corefn.json` lacks the TAST fields.
Prepending the selected typed frontend's directory corrected that invocation;
all five executable fixtures then passed with unchanged snapshots, followed by
the package-surface check (59 files, approximately 3.4 MB packed).

Together these runs cover all 15 quick gates: ten Node files with 173 passing
entries and no skips, three Go modules, five native fixtures, and package
contents. This is combined evidence, not one uninterrupted green command.
The original interrupted/failed reports and the corrected results are retained
at the local workspace root under
`.gopurs-evidence/release-preparation-20261026/rc1-tooling`.

The clean local RC preview installs offline with an empty npm cache and passes
the TAST-to-Go FFI smoke on both JavaScript and native hosts. All ten packaged
checksums pass. npm 11.6.0 rejected the initial prerelease dry-run without an
explicit distribution tag; it passed with `--tag next`. CI now chooses `next`
for candidates and `latest` for final versions in its dry-run as well as in
the publication workflow. No public version was created by these checks.

## Release preparation follow-up — 26 October 2026

The following checks were executed during the reproducible-bootstrap/Nix work.
They do not replace the full release gate on the final tagged artifacts.

| Check | Observed result |
| --- | --- |
| Public TAST, optimizer and 51 Go library source revisions | All pinned revisions verified remotely |
| Locked library installation into an empty directory | 51 repositories installed; second `--locked --all` verification passed |
| Bootstrap orchestration and pinned-input tests | 11 tests passed |
| Vendored codec borrowed-object and decoder tests | 44 tests passed; codec record-plan Go race test passed |
| Native parallel cache-source transitions | Passed (2 Node test entries including the host subtest) |
| Native sequential/parallel context, option and entry transitions | Passed |
| Actual `int32` FFI smoke | Reproduced nil-pointer panic, then produced the expected output after the runtime fix |
| Signed `int32` scalar, array and callback bridges, including min/max values | Passed; complete FFI bridge file 3/3 |
| Fresh source JavaScript build with vendored codec workspace | Passed after isolating the codec workspace to prevent duplicate `Main` modules |
| Runner selection and cleanup after excluding backend worktrees | 96/96; external locked library-root cleanup 52/52 |
| Native bootstrap from fresh sources and locked public libraries | Passed: 522 typed modules, 326,573 types; no previous native binary used |
| Clean JavaScript bundle reproducibility | Two source paths match SHA-256 `e7c761c44aef7276b335265459c03cfe6551b7ed8ec00f76f8233aa6fbe30b3b` after removing stale generated output |
| Final preview archive installation, empty npm cache and offline install | JavaScript and native TAST → Go → `int32` FFI smoke passed |
| Preview from source commit `0c9a6a1f69d64473ae240f547312fd97f6dd45cb` | Offline installation and both host smokes passed; compiler bytes match the preceding preview |
| Complete isolated b8x build from local preview artifacts | 17 stages passed: 2,707 TAST modules, all Go packages and 12 linked executables; no application executed |
| `npm publish --dry-run` on the preview archive | Passed: 59 files, approximately 3.4 MB |
| Preview SBOM against the official CycloneDX 1.6 schema | Passed: 351 unique `bom-ref` values, 350 resolved dependency references and 224 valid purls |
| npm rollback selection against a local mock registry | Six commands passed; removing `latest` and deprecation do not prevent installation, while restoring `latest` selects the retained version |
| Real package consumers, JavaScript and native sequential/parallel | Targeted rerun passed: all nine fixtures and their parent (10/10 entries), 321,939 ms |
| Local worker native-return assertion after record-label metadata update | Targeted rerun passed (1/1) |
| Go race checks in `tools/build-cache`, `tools/c1c` and `tools/ffi-gen` | All three passed with `go test -race -count=1 ./...` |
| Native emission pipeline against retained bootstrap output | All four `TestPipelineNative` tests passed with `-race -count=1 -timeout 30s` |
| Complete native fixture campaigns, Release and Dev | 401/401 passed in each profile; 802 frontend/backend/Go builds and executions, strict snapshots, no updates or exclusions |
| Linux arm64 Nix backend and pinned TAST frontend builds | Passed with Nix 2.35.4 |
| Linux arm64 `nix flake check -L` | Passed: installed backend, TAST schema, Go FFI application in both layouts |
| Linux arm64 `nix run /src -- --main Main` | Passed from a fresh application directory, followed by Go build and exact stdout verification |
| Nix evaluation of all three target systems, with IFD disabled | `nix flake check --all-systems --no-build --option allow-import-from-derivation false --option eval-cache false` passed; evaluation only |
| Nix development shell toolchain selection | Node 24.8.0, npm 11.6.0, Go 1.27.0, Spago 1.0.3 and TAST `purs` 0.15.16 |

The runtime now boxes Go `int32` values as integer payloads and unboxes them
through the existing integer conversion, while retaining opaque `Any(int32)`
values. Both hosts were rebuilt before artifact installation was checked.

The runner cleanup suite exposed a sibling backend worktree being mistaken for
a library; selection now checks the package identity and uses the configured
library root. Native test compiler discovery also respects `GOPURS_PURS` when
the caller does not provide an explicit argument (12 bootstrap/input tests).

The Nix checks above ran in an isolated Linux arm64 container whose Nix daemon
has `sandbox = false`. The backend build used Spago's offline mode and locked
inputs. Evaluation passed for `aarch64-linux`, `x86_64-linux` and
`aarch64-darwin`; actual builds and execution were validated on Linux arm64.
The other platform builds subsequently passed in the initial CI recorded above;
the final tagged builds remain a separate gate. The native bootstrap runs inside
`nix develop`, rather than as a pure
Nix derivation. The local evidence is
`/private/tmp/gopurs-nix-validation-20261026/validation.json`. Its durable copy,
build/check/run logs and recorded source inputs are under `nix/` in the evidence
archive described below. The 27 October audit found all 11 recorded file hashes
unchanged and all six protected Nix outputs present, with no active build.
This audit did not rebuild or extend the validated platform set.

### Full Node campaign — 26 October

The fresh source copy and retained native bootstrap output produced **613 test
entries: 607 passed, 5 failed, 1 skipped**, in 1,760,977 ms. The command exited
with status 1; this campaign is not a green release gate. The complete local
log is `/private/tmp/gopurs-release-20261026/all-node-tests.log`.

| Entry | Classification and follow-up |
| --- | --- |
| `NativeTraverseCallback`, `tools/driver-package-consumers.test.mjs` | The frontend passed, then the explicit free-space assertion (`> 8 GiB`) prevented monolithic generation |
| `JsonRecordPlan` and `B8xNativeGeneration`, same file | The same free-space assertion prevented the frontend from starting |
| `real package consumers preserve the monolithic ABI and execution across compiler hosts` | Aggregate parent failure for the three entries above; not a fourth independent defect |
| `local workers return native records while first-class and partial calls retain wrappers`, `tools/local-native-returns.test.mjs` | Outdated assertion excluded the emitted record-label metadata comment, although the result remained a concrete native struct; assertion corrected and focused check subsequently passed |
| `Rust-hosted gopurs preserves Go sources and execution on native regression fixtures`, `tools/rust-host.test.mjs` | The only skipped test; the experimental Rust host requires its separate build and `npm run test:rust` |

The first six package-consumer fixtures passed with the JavaScript host. Their
retained command inventory confirms that this campaign did not exercise the
two native package-consumer modes. A targeted rerun of the complete consumer
file with `GOPURS_TEST_MODES=js,native-sequential,native-parallel` first hit a
sandbox restriction on Spago's cache database (`SQLITE_CANTOPEN`) before any
consumer compilation; the corrected local-return test passed in that run.
The consumer rerun with cache access permitted then passed all nine fixtures
on all three hosts (10/10 test entries including their parent), in 321,939 ms,
recorded in
`/private/tmp/gopurs-release-20261026/package-consumers-three-hosts-unrestricted.log`.
The 8 GiB assertion is an environment guard, not an observed `ENOSPC` or
generated-code failure. No additional failing tests were found in the full
Node log.

Together, the full campaign and these focused reruns resolve every failure
from the original Node report; the experimental Rust skip remains. This is
combined evidence, not a newly executed full Node campaign with all entries
passing in one run.

### Go race checks — 26 October

The three tool modules were tested sequentially from the fresh source copy,
followed by `go test -race -count=1 -run '^TestPipelineNative' -timeout 30s
./purescript` in the retained native bootstrap output. All four commands
exited with status 0. The Go-reported package test times were 26.864 s
(`build-cache`), 1.603 s (`c1c`), 10.396 s (`ffi-gen`) and 1.531 s (native
emission); the native command additionally needed about 229 s for compilation
and setup.

Each command used `GOMAXPROCS=2`, `GOGC=100`, `GOMEMLIMIT=4GiB`, `GOWORK=off`
and the shared task `GOCACHE`. The free-space check passed before each command
(minimum observed at those checks: 11.11 GiB; required floor: 8.5 GiB).
`tools/emission-native_test.go` was copied with exclusive creation into the
bootstrap's `purescript` package, then removed in cleanup; its absence was
verified afterward. No backend source was changed. Command arguments,
environment, timestamps, exit statuses, source hash and individual logs are
preserved in `/private/tmp/gopurs-release-20261026/go-race-gates/status.json`
and the adjacent `.log` files.

Native parallel cache-source transitions and native sequential/parallel
context, option and entry transitions passed on 26 October, including in this
full Node campaign. The older incomplete cache campaigns below are historical.
The complete Release/Dev fixture campaigns subsequently passed, as recorded
below. Release benchmarks, remote CI and publication remain separate gates.

### Complete Release and Dev fixtures — 26 October, audited 27 October

All **401 Release fixtures and 401 Dev fixtures passed**, with zero failures,
pending targets or exclusions. Each profile covered the complete selection
exactly once, in 81 batches. Every fixture compiled through the pinned TAST
frontend and the freshly bootstrapped native backend, compared the generated
Go against its profile's golden files, then built and executed the Go program.
Dev builds used `-gcflags=gopurs/output/...=-N -l`. Snapshot updates were disabled
throughout. This campaign used the monolithic fixture layout; the package-layout
consumer checks are recorded separately above.

The first fixture started at 15:46:24 UTC and the campaign completed at
17:42:36 UTC on 26 October. This interval includes pilot runs and resource
coordination pauses. After measured pilots, the controller used six independent
batches of at most five fixtures, with `GOMAXPROCS=2`, `GOGC=100` and
`GOMEMLIMIT=4GiB`. It trimmed only its dedicated Go cache between completed
waves, never while a fixture build was active. These qualification timings are
not backend performance measurements.

The completion audit re-read all 162 batch reports, checked the recorded 324
report/log hashes and retained all 4,812 phase logs. It independently rehashed
the 1,357 recorded source, runner, runtime and snapshot inputs, including all
810 golden files: every tested input was unchanged. Both compiler hashes also
matched. No campaign process remained. Compared with the current checkout,
only the vendored `support/argonaut-codecs/LICENSE` had changed afterward;
the archive retains its tested version and records both hashes. The other
1,356 inputs, including every fixture, snapshot and runtime source, matched.

| Verified evidence | SHA-256 |
| --- | --- |
| Native compiler | `eb97922721f15c3e742ff14335a8294b6005ca321ae2e9a57bb8847776e46298` |
| TAST compiler | `efac5816a482d5581cfc990892f8b24f1d3343f8df1798572ea1a0a96bd07991` |
| Runtime source | `24764a5840c20841be5dd892dddc6b56ab44f4a9897f61b775db28c4a38c9388` |
| Fixture `inputs.json` | `54e2343880d7fb1016c390db85a54864fe3727d65a5917ef69d6393ba62ec2e6` |
| Completed fixture `progress.json` | `dc694ea8d94323abb56a13d47ae6980b47f822b5fb63132647173bebbb3860a1` |
| Nix `validation.json` | `62f1e51e320f2ec3b3c1965cf4bc0b2089421a1c09ac7723d72da4d4d526cb3a` |

The original fixture evidence is in
`/private/tmp/gopurs-release-20261026/fixture-campaign`. A durable local archive
is in
`.gopurs-evidence/release-preparation-20261026` at the local workspace root:
`fixtures/` contains the reports, selection, telemetry and compressed source
inputs/phase logs; `nix/` contains the Nix evidence and inputs.
`fixtures-nix-audit.json` records the audit, and `fixtures-nix-SHA256SUMS`
covers the 391 archived files (9,800,100 bytes before the checksum list).
The compressed members were also read back and verified. Caches, generated
workspaces, compiler binaries and Nix store closures are excluded. The checksum
list itself has SHA-256
`5aaa506455abe7bdcae326b8a4796e498751ac90666e49fcda24d780d5a05926`.
This is completed local qualification of the recorded compiler bytes; the
release tag, published artifacts and three-platform CI remain distinct gates.

### Local artifact and b8x qualification — 26 October

The newer preview in
`/private/tmp/gopurs-release-review-0c9a6a1-20261026` records clean source commit
`0c9a6a1f69d64473ae240f547312fd97f6dd45cb`, `preview: true` and no release tag.
Its `smoke.json` reports an offline installation with an empty npm cache and
lifecycle scripts disabled. Both JavaScript and native hosts compiled and ran
the TAST/Go FFI example, producing exactly `gopurs release smoke: 42`.

Comparison with `/private/tmp/gopurs-release-final-review-20261026` established
the following identical compiler hashes, without rebuilding either compiler:

| Artifact | SHA-256 in both previews |
| --- | --- |
| `bin/gopurs.js` inside the npm tarball | `e7c761c44aef7276b335265459c03cfe6551b7ed8ec00f76f8233aa6fbe30b3b` |
| Native Darwin arm64 compiler | `eb97922721f15c3e742ff14335a8294b6005ca321ae2e9a57bb8847776e46298` |

All archive entries were compared: only `NOTICE`, `THIRD_PARTY_NOTICES.md`
and `docs/nix.md` differ. There are no added or removed entries. The older npm
tarball SHA-256 is
`05c7aa62674b59e8cd216b25eb4597a39e43d68509d41c7c0b1b99dd359b6310`;
the newer one is
`b3a6b19958c946b96c47a88d9c03cb1e52d7ae4d36d32aa1737ae4d1e4cb04d9`.
Release metadata additionally records the newer source commit. This is an
untagged local candidate, not an installation from publicly published artifacts.
All listed checksum entries were verified in both previews (12 in the older
directory and 10 in the newer one).

The b8x qualification installed the older preview tarball offline into a new
temporary workspace and copied its native artifact there. It used a clean
b8x source snapshot at commit `3adb29c7acddf6799872bd98c95d64ae284dd34e`,
138 dependency source directories, and the pinned TAST frontend. Its 17 stages
all passed: artifact installation, frontend, private Go dependency download,
native Release generation, `go build -mod=readonly -trimpath -p 2 ./...`, and
12 individual executable links. The frontend produced 2,707 typed modules
(215,934,561 bytes). The executables include `Inter.Api.Main`, `Inter.Ui.Main`,
the CLI entry points, `Test.Ffi.Native` and `Test.Main`.

The build used Go 1.27.0, Node 24.8.0, `GOMAXPROCS=2`, `GOMEMLIMIT=12GiB` and
private npm/Go caches. The source snapshot includes the application's test
modules; neither the application binaries nor any service/database test was
executed. The original b8x checkout remained clean. Registry dependency sources
were copied from the existing application cache and hashed; their archives
were not downloaded again. The copied application `go.mod` and `go.sum`
provided the dependency pins.

The original temporary workspace was lost during the interruption on
28 October; its `plan.json` has not been recovered. The previously archived
generated sources and 12 executables remain in
`b8x-generated-20261026.tar.gz` under
`.gopurs-evidence/release-preparation-20261026` at the workspace root. The
3,242 input files were recovered separately and their exact hashes checked
against the retained input manifest, as recorded in
`task-inputs/b8x-release-20261026/recovery-audit.json`. This proves input
identity, not a resumed execution of the lost plan. The local build result
above remains historical. The separate public RC7 build on 30 October closes
the public-artifact compilation gate, as recorded at the top of this document.

### Complete local cold/warm benchmark — 27 October

All **51 official corpora passed on their first attempt** between 08:37:40
and 10:19:22 UTC: b8x plus 50 library corpora, on the frozen 25 October inputs.
`gopurs-quickcheck` has no standalone case in that historical table. The
campaign used the untagged compiler snapshot
`0c9a6a1f69d64473ae240f547312fd97f6dd45cb` prepared on 26 October; only this
quality record was dirty at preparation. Its exact JS/native hashes match the
previews above and the local binaries after the package metadata update to
`0.1.0-rc.1`. This does not qualify a tagged CI rebuild or public installation.

An independent completion audit verified **1,224 invocations: 918 measured
samples and 306 warmups**, 612 cold cache misses and 612 exact warm hits. Each
cell contains the specified three measurements after one warmup. All exits,
variant orders, raw timers and 102 output-oracle manifests were checked.
JS/native Release outputs matched within each corpus; Dev remained stable
against its own oracle. There were no failed/retried invocations, full or host
fallbacks, or Go-cache evictions.

The following values are sums of the 51 per-corpus medians of `backend total`,
in milliseconds. The reference is the official `altbak.pub/README.md` table,
verified against its 25 October report and JSON.

| Host/profile | Official cold | Current cold | Change | Current warm |
| --- | ---: | ---: | ---: | ---: |
| JavaScript Release | 812,478 | 809,532 | −0.36% | 13,111 |
| Native Release | 332,715 | 288,082 | −13.41% | 5,808 |
| Native Dev | 265,633 | 221,643 | −16.56% | 5,708 |

All native cold medians were lower than their historical counterparts. JS was
lower in 37 cases and higher in 14; its aggregate differs by only −0.36%.
Warm values measure unchanged-input cache reuse and have no
historical baseline. The full
[benchmark record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/release-benchmark.md)
retains the JS increases, individual samples, load observations and protocol
limits: the new campaign interleaves all hosts and cold/warm pairs, whereas
the historical JS measurements preceded the native reruns. No causal speedup,
application runtime gain or cross-platform performance claim is established.

The audit started only after the runner exited with code 0 and its lock was
removed. Its durable evidence is in
`.gopurs-evidence/release-preparation-20261026/benchmark` at the local workspace root:
72 files, 98,766,767 bytes, including raw logs/records, manifests, baseline
documents, analysis and compressed frozen sources/TAST/compiler inputs.
The compressed members were read back and verified. Caches are excluded.
No old input or timing was changed.

| Evidence | SHA-256 |
| --- | --- |
| Completed `campaign.json` | `e83af80e05f3ceaf04bbc96495d7d4e2bcc4148dd1279b08f3718572a9dd872e` |
| Completed `summary.json` | `9361f216e2de0b7cf9aafcbda81aedaf00eb1671a8c50575d4a483af9c30be40` |
| Complete archive `SHA256SUMS` | `38cacf4f552662689fb7cb6a4a72a0ac90ddcd1af3a11d906e2e77fd094681fa` |

## Historical clean-checkout campaign — 25 October

### Environment

- Clean source checkout with no `node_modules`, `output`, `.spago` or compiler
  binary before `npm ci`.
- Node.js `24.8.0`.
- Go `1.27.0`.
- Typed `purs` fork first on `PATH` for application and fixture builds.
- Local `gopurs-*` sibling checkouts for fixture and runner tests.
- Native compiler from the 25 October build; native source revision matches the
  checkout used by the test runs.

### Passed in the clean checkout

| Check | Result |
| --- | --- |
| Clean `npm ci` and JavaScript bundle build | Pass |
| `npm run test:ffi` | 14/14 |
| `npm run test:launcher` | 9/9 |
| `npm run test:cli` (`GOPURS_GO_PACKAGES=0`) | 19/19 |
| `npm run test:runner` | 95/95 |
| `npm run test:profiles` | 12/12 |
| `npm run test:cache` | 6/6 |
| `npm run test:cache-codecs` | 15/15 |
| `npm run test:cache-driver` | 4/4 |
| `npm run test:cache-sources`, JavaScript and native sequential | Pass |
| `npm run test:cache-context`, JavaScript | Pass |
| `tools/build-native.test.mjs` | 7/7 |
| Go tests in `tools/build-cache`, `tools/c1c`, `tools/ffi-gen` | Pass |
| Representative fixtures: FFI, arrays, TCO, thunk fusion, JSON/records, native workers, strings, int32 loops, owned trees | 9/9, snapshots and Go execution |
| `npm run test:package` | Pass, 57 files, 3.39 MB packed |

### Fixes made during that gate

- The output-barrier cache test now copies the `c1c` module required by
  `build-cache/go.mod`.
- Cache transition tests pin `GOPURS_GO_PACKAGES=0` and
  `GOPURS_PACK_PUBLICATION=0` when they inspect individual cache objects.
- The context-transition check now accounts for prepared fingerprints reused
  within a run instead of requiring a full recomputation on every state.

### Incomplete at the end of 25 October

- **Full fixture campaign.** `./bin/test --all --keep-going` reached the 23rd
  fixture and stopped because the temporary volume filled while writing
  `results.json`. This is an environment failure, not a fixture assertion.
- **Native parallel cache-source campaign.** The run hit the same disk limit.
  Native sequential and JavaScript passed. The missing native parallel check
  subsequently passed on 26 October, as recorded above.
- **Cold/warm release benchmark.** A small smoke run was executed, but it is not
  a substitute for the protocol in [benchmark-protocol.md](benchmark-protocol.md).
- **Release candidate/tag and green CI.** The workflow exists but has not run on
  GitHub for a release tag.

## Remaining gates before stable — updated 1 November 2026

1. Keep the corrected library pins and notices in the frozen candidate inputs,
   as recorded in [the audit](library-source-audit.md). The identified issues
   are integrated; independent library publication still needs its own review.
2. Select a fallback suitable for the affected applications and verify the
   documented [rollback procedure](release.md#rollback-and-recovery).
   RC7's retained assets and local reinstallation are verified above; neither
   this check nor the mock registry exercise is a validated public rollback.
3. Once the candidate inputs are frozen, qualify their exact tag on every
   supported platform, including the library campaigns, artifact finalization
   and public installation checks. Preserve all earlier RC tags and record the
   new results separately; quick CI and local evidence do not close these gates.
