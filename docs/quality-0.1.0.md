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
| `npm publish --dry-run` on the preview archive | Passed: 59 files, approximately 3.4 MB |
| Real package consumers, JavaScript and native sequential/parallel | Targeted rerun passed: all nine fixtures and their parent (10/10 entries), 321,939 ms |
| Local worker native-return assertion after record-label metadata update | Targeted rerun passed (1/1) |
| Go race checks in `tools/build-cache`, `tools/c1c` and `tools/ffi-gen` | All three passed with `go test -race -count=1 ./...` |
| Native emission pipeline against retained bootstrap output | All four `TestPipelineNative` tests passed with `-race -count=1 -timeout 30s` |
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
`/private/tmp/gopurs-nix-validation-20261026/validation.json`.

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
Release/Dev fixture campaigns, release benchmarks, remote CI and publication
are not yet complete.

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

1. Free enough temporary space and rerun the complete fixture campaign for both
   Release and Dev.
2. Execute the published cold/warm protocol on the release corpus.
3. Tag the release candidate, run CI on the tag, and record the artifact
   checksums.
