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
| Linux arm64 Nix backend and pinned TAST frontend builds | Passed with Nix 2.35.4 |
| Linux arm64 `nix flake check -L` | Passed: installed backend, TAST schema, Go FFI application in both layouts |
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
inputs. These results do not establish Darwin/amd64 support or a sandboxed CI
build; those remain independent platform gates.

The full Node campaign is running against the fresh source copy and retained
native bootstrap output. Release/Dev fixture campaigns, release benchmarks,
remote CI and publication are not yet complete.

## Environment

- Clean source checkout with no `node_modules`, `output`, `.spago` or compiler
  binary before `npm ci`.
- Node.js `24.8.0`.
- Go `1.27.0`.
- Typed `purs` fork first on `PATH` for application and fixture builds.
- Local `gopurs-*` sibling checkouts for fixture and runner tests.
- Native compiler from the 25 October build; native source revision matches the
  checkout used by the test runs.

## Passed in the clean checkout

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

## Fixes made during the gate

- The output-barrier cache test now copies the `c1c` module required by
  `build-cache/go.mod`.
- Cache transition tests pin `GOPURS_GO_PACKAGES=0` and
  `GOPURS_PACK_PUBLICATION=0` when they inspect individual cache objects.
- The context-transition check now accounts for prepared fingerprints reused
  within a run instead of requiring a full recomputation on every state.

## Not completed

- **Full fixture campaign.** `./bin/test --all --keep-going` reached the 23rd
  fixture and stopped because the temporary volume filled while writing
  `results.json`. This is an environment failure, not a fixture assertion.
- **Native parallel cache-source campaign.** The run hit the same disk limit.
  Native sequential and JavaScript passed.
- **Cold/warm release benchmark.** A small smoke run was executed, but it is not
  a substitute for the protocol in [benchmark-protocol.md](benchmark-protocol.md).
- **Release candidate/tag and green CI.** The workflow exists but has not run on
  GitHub for a release tag.

## Next gate actions

1. Free enough temporary space and rerun the complete fixture campaign for both
   Release and Dev.
2. Rerun native parallel cache-source and context campaigns.
3. Execute the published cold/warm protocol on the release corpus.
4. Tag the release candidate, run CI on the tag, and record the artifact
   checksums.
