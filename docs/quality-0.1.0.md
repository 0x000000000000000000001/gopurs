# Quality-gate status — 0.1.0

This is a working record for the `0.1.0` release gates. It is not a release
certificate. The gate list is in [`todo.md`](../todo.md).

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
Linux amd64, macOS arm64 and sandboxed CI builds remain independent platform
gates. The native bootstrap runs inside `nix develop`, rather than as a pure
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
`/Users/0x1/Documents/htdocs/.gopurs-evidence/release-preparation-20261026`:
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

The evidence is
`/private/tmp/gopurs-b8x-release-build-20261026/plan.json`, its stage logs and
executable hashes. Sources, TAST, generated Go and the 12 executables remain
available. Only that workspace's backend and Go build caches were removed
after success, freeing 1,860,024,370 bytes; their inventories and backend
manifests were retained. This is complete local build evidence, not a benchmark
or a service execution check. The gate requiring b8x to build from publicly
published artifacts remains open.

### Full benchmark corpus prepared, not measured — 26 October

`tools/release-benchmark.mjs prepare` verified all 51 frozen 25 October corpora
and the baseline table in the official `altbak.pub/README.md`. It prepared
`/private/tmp/gopurs-release-benchmark-20261026` using the compiler from commit
`0c9a6a1f69d64473ae240f547312fd97f6dd45cb`; only this quality record was dirty.
The copied JavaScript/native hashes match the previews listed above. The
prepared snapshot occupies 115,436,835 bytes before its small review record,
below the 1 GiB preparation budget.

All 51 cases are `pending`, with zero attempts. No compiler invocation or
benchmark measurement was launched; fixture campaigns were still running.
This is preparation of an untagged local candidate, not a performance result
or a completed release gate.

`campaign.json` records each corpus, compiler file, source/runtime input,
toolchain, baseline and runner hash. `preparation-review.json` records the
scope and checksums of the prepared manifests. Initial file SHA-256 values:

| Evidence | SHA-256 |
| --- | --- |
| `campaign.json` before any run | `7f54ed917d88948efeb2ca2ba6297940809f33e8175f235a2149bb86f99a3300` |
| `source-manifest.json` | `036ef7ebf0e0b8ddb1e6e1fe55f190940df438face28533a448905abd90bee8b` |
| `module-cache-manifest.json` | `011403b41d4065348d81cf75178e6fb3f30ca4f971c4a0f4c51e46b92c0a749c` |

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

## Next gate actions

1. Execute the published cold/warm protocol on the release corpus.
2. Tag the release candidate, run CI on the tag, and record the artifact
   checksums.
