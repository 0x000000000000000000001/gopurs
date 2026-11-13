# Release backend benchmark

`tools/release/release-benchmark.mjs` compares the release JavaScript host, native
Release host and native Dev host on the frozen 25 October 2026 corpus. It does
not regenerate TAST, rebuild a compiler, compile a final application or measure
application execution. The primary metric is the driver's internal
`backend total`; process wall time is retained as a separate raw measurement.

The reference is the **Go compilation table in the altbak.pub README**, checked
against its
[25 October report](https://github.com/0x000000000000000000001/altbak.pub/blob/main/docs/benchmark-results/2026-10-25-gopurs-compilation-refresh.md)
and JSON. Its 51 corpora total **812,478 ms JS Release**, **332,715 ms native
Release**, and **265,633 ms native Dev**, summing the per-corpus cold medians.
Rust values in that historical report are estimates and are not used here.
There is no historical warm-cache baseline in that table.

## Completed RC20 campaign — 13 November 2026 (UTC)

All **51 corpora completed** with the locally built RC20 compiler from clean
commit `2202368d4b19849fdc365afacaffa4a617b0ea43`, signed tag object
`9a80c154eff82589275d6790e158e0636066197c`. The JavaScript SHA-256 is
`26ee3b1a8f40fa4ed9675f478f89c39464b3aa13e1b726587f9c12832e709d9a`;
the Darwin arm64 native SHA-256 is
`858b204ba61cf6691ef4638ccfd17541effaa55472e352663aef0eda7710d10e`.
These remain RC20 measurements if a later release reuses the same executables.

The audit verifies **1,224 retained invocations**: 918 measured samples,
306 warmups, 612 cold misses, 612 exact warm hits and 102 matching output
oracles. Each median uses three measured fresh processes after one warmup.
The settings remain those of the recorded protocol: eight workers, native
`GOGC=off` and `GOMEMLIMIT=16GiB`, packages/overlays enabled and daemon disabled.

The initial process started at 11:00:04 UTC and stopped at 11:47:12 after
19 complete cases when the 2 GiB disk floor prevented the next invocation.
No compiler invocation failed. The sole successful `js-date` cold JavaScript
warmup from that incomplete attempt is **excluded**, with its log and oracle
retained. After retiring private task caches and archiving verified old
workspaces, the official `--retry-failed` path restarted the whole case at
12:06:53, preserving the 19 completed cases. The second process exited 0
after the last case finished at 13:08:51. Both real process identities and
terminals are retained; the audit checks 1,225 recorded invocations in total.
The runner's disk floor and measurement settings were unchanged.

The primary comparison uses the **47 historically identical corpora**.
The other four retain the previously disclosed five added Go FFI files;
all 51 original PureScript and TAST inputs remain unchanged. Values below
sum per-corpus medians in **milliseconds**, not elapsed campaign duration.
Warm results have no historical baseline.

| Host/profile | Official cold, 47 unchanged | RC20 cold | Change | RC20 warm |
| --- | ---: | ---: | ---: | ---: |
| JavaScript Release | 718,784 | 749,080 | +4.21% | 14,259 |
| Native Release | 296,270 | 282,920 | −4.51% | 5,388 |
| Native Dev | 237,978 | 214,609 | −9.82% | 5,275 |

The following groups are context only because they include corrected FFI
inputs. They must not be presented as identical-input historical comparisons.

| Group | JS cold / warm | Native Release cold / warm | Native Dev cold / warm |
| --- | ---: | ---: | ---: |
| Four corrected corpora | 92,934 / 1,233 | 32,113 / 473 | 23,312 / 458 |
| All 51 corpora | 842,014 / 15,492 | 315,033 / 5,861 | 237,921 / 5,733 |

Compared with the separate RC18 campaign on the same derived inputs, the
51 cold totals increased by 3.94% JS, 3.93% native Release and 4.78% native Dev.
Three samples per cell, environmental variability and the disclosed pause do
not establish a causal attribution to a compiler change. These are backend
generation measurements, not frontend/build or application runtime results.

The public [JSON samples](benchmark-results/2026-11-13-rc20.json) and
[CSV samples](benchmark-results/2026-11-13-rc20.csv) contain all 1,225 recorded
invocations, their timing columns and the explicit warmup exclusion. The
per-cell medians and the 47/4/51 totals were recomputed against the audit.
These exports contain no private source files or local filesystem paths.

Detailed records and verified archives remain under
`.gopurs-evidence/publication-rc20-20261113/benchmark-audit` at the local workspace
root; private b8x inputs are not uploaded. The case archive contains all 2,553
retained/excluded files, and the frozen-input archive contains 32,176 entries.

| Evidence | SHA-256 |
| --- | --- |
| Terminal audit | `2f20b55188679fc21296a52e35e375a9df8beb8f5a882f66df58e75add539de9` |
| Archive checksum list | `7a8bfd6c649d1319f3d73b0631570ae5e589d9d400bac03121f0bfe251bf8e59` |
| Case records | `8b2cda309e1b8b6a10d53ca3c7a9b35ad78064261320045a20d9eb7c3802dfe7` |
| Frozen inputs | `2f6948365efb994ef75353421885ce4f0962d55e41f1a0f73f5ab2e0a80895e0` |

The two native helpers were preserved and hashed after completion; their
hashes were not sampled per invocation. The successful audit and benchmark
do not substitute for fixture, artifact or public-consumer qualification.

## Completed RC18 campaign — 12 November 2026 (UTC)

All **51 corpora passed on their first attempt** with the locally built
`0.1.0-rc.18` compiler from clean commit
`c6c68d1486362ec2ba6ec23534a10be60f571f68`. Its signed tag
`v0.1.0-rc.18` has annotated object
`69482ff422b3961751716b5c2821b033b3878edb`. The runner started at
`2026-11-12T21:50:42.697Z` and exited successfully after the last case finished
at `2026-11-12T23:35:20.397Z`. The independent terminal audit completed at
`2026-11-12T23:36:14.991261Z`. The campaign retains this RC18 identity even if
its evidence is later referenced by another release.

The primary historical comparison uses **47 corpora with identical inputs**.
These include b8x and 46 libraries. The other four library corpora received
explicit Go FFI additions and are shown separately below. All 51 retain their
original PureScript and TAST bytes, including b8x's 2,707 modules. No frontend
ran during this measurement. `gopurs-quickcheck` remains a bootstrap source
pin without a standalone case in the official 50-library table.

The tables sum per-corpus medians in **milliseconds**, each based on three
measured fresh processes after one excluded warmup. They do not give the
campaign's elapsed duration, complete application build time or runtime.
The warm column measures exact unchanged-input cache hits and has no historical
baseline. The official reference remains the altbak.pub README table above.

### Principal comparison: 47 unchanged corpora

| Host/profile | Official 25 October cold | RC18 cold | Change | RC18 warm |
| --- | ---: | ---: | ---: | ---: |
| JavaScript Release | 718,784 | 716,655 | −0.30% | 15,186 |
| Native Release | 296,270 | 270,492 | −8.70% | 5,447 |
| Native Dev | 237,978 | 203,360 | −14.55% | 5,397 |

The recorded native Release and Dev cold medians were lower than their
historical counterparts for each of these 47 cases. This is an observed
comparison with the historical protocol, not a controlled attribution to one
compiler change. JavaScript moved in both directions: across all 51 cases,
29 medians were lower and 22 higher. Its largest absolute increases were
avar (15,097 → 15,473 ms), strings-extra (13,757 → 14,086 ms) and partial
(9,519 → 9,827 ms); all three belong to the unchanged group. The full samples
and per-case comparisons remain available in the local evidence.

For b8x, native Release was 76,016 ms against the historical 78,037 ms
(−2.59%), and native Dev was 47,044 ms against 53,611 ms (−12.25%). Its warm
medians were 531 ms JS, 321 ms native Release and 307 ms native Dev. These are
backend-generation measurements from the historical TAST, separate from the
installed-artifact application build and execution gate.

### Context only: four corrected corpora

| Host/profile | Official 25 October cold | RC18 cold | Change | RC18 warm |
| --- | ---: | ---: | ---: | ---: |
| JavaScript Release | 93,694 | 93,433 | −0.28% | 1,269 |
| Native Release | 36,445 | 32,635 | −10.45% | 485 |
| Native Dev | 27,655 | 23,715 | −14.25% | 474 |

These four cases are `gopurs-node-http`, `gopurs-node-net`,
`gopurs-node-streams` and `gopurs-spec`. Their historical comparison is **not
at identical inputs**. The derived archive adds exactly five Go FFI files;
all original source files, existing FFI, PureScript and TAST remain unchanged:

| Derived corpus | Added Go FFI | Public source commit |
| --- | --- | --- |
| node-http | `test/Main.go` (`stdout`) | `9b11034d49163590634801a85219f765402b1fc7` in gopurs-node-http |
| node-net and node-streams | `Node/Buffer/Immutable.go` in each frozen node-buffer dependency | `f154e7d6ded317ed7559540f4fa0dff6d466d3f7` in gopurs-node-buffer |
| spec | `src/Node/ChildProcess.go` and `src/Node/ChildProcess/Types.go` | `de7f8d143d4a8468d9ec358598c24bb610a21bdc` in gopurs-spec |

The historical archive was not modified. The derived archive has 31,908 source
files versus 31,903 originally; the derivation manifest records every addition,
its public source identity and file hash. Current library configuration changes
are documented in the [source audit](library-source-audit.md); they must not
be confused with regeneration of this frozen TAST. The aggregate of all 51
cases is contextual only: cold JS/native Release/native Dev totals are
810,088 / 303,127 / 227,075 ms, against 812,478 / 332,715 / 265,633 ms in the
historical table; warm totals are 16,455 / 5,932 / 5,871 ms.

### Validation, environment and limits

The audit independently recomputed the results from **1,224 invocations:
918 measured samples and 306 warmups**. It checked all 612 explicit cold
misses and 612 exact warm hits, invocation order, exits, timers and output
fingerprints. The 51 Release oracles compare JS and native Release output;
51 separate Dev oracles check stability across its repetitions. All 102
per-file manifests are retained. No invocation failed, timed out, used a full
fallback or silently substituted JS for a missing native executable. No
sample was trimmed or retried. This checks generated output agreement, not
application execution or the correctness of every possible program.

The machine was an Apple M4 Pro with 14 logical CPUs and 48 GiB RAM, Darwin
25.1.0, Node 24.8.0 and Go 1.27.0 darwin/arm64. Source, runtime, compiler
files, lockfiles, pinned dependencies, tool hashes, baseline files and input
identities are retained in the audit. The frozen Go module cache was used with
`GOPROXY=off`, `GOTOOLCHAIN=local` and `GOWORK=off`. Eight prepare/PBO/emission
workers, pipelining, packages, overlays and build manifests were enabled;
daemon and resident modes were disabled. Native invocations used `GOGC=off`
and `GOMEMLIMIT=16GiB`; JS did not inherit these GC settings. The configured
1 GiB Go cache eviction policy never triggered; its largest recorded
post-case size was 100,019,199 bytes.

Cold means fresh backend/output caches, with a warm filesystem and Go export
cache. Every cold call was immediately followed by a fresh-process warm call.
The variants rotated between rounds. The historical JS v2 measurements were
retained before the native v3 measurements and did not include paired warm
calls; the compiler binaries also differ. The recorded host and Node version
match the historical reference, but software and scheduling/thermal conditions
are not a controlled identical environment. Frontend compilation, compiler
bootstrap, final Go compilation, application execution and process startup/exit
are excluded from the internal `backend total` metric.

One-minute load before invocations ranged from 2.28 to 9.15, with median 4.24;
it includes earlier work in the campaign and does not establish an otherwise
idle machine. The largest cold sample range was JS bigints:
4,748 / 4,763 / 5,958 ms, with the prescribed median 4,763 ms (25.40% range
relative to that median). All samples remain in the data. Three observations
per cell do not establish statistical confidence or a latency bound.

The measured compiler hashes are:

| Artifact | SHA-256 |
| --- | --- |
| JavaScript compiler | `5ad71cd10cdf4ed0f18e7bbeb393f35fac1fa1d11534f5b89d982c953195377c` |
| Native compiler, Darwin arm64 | `3e6cccfb97792003549f35b3283956a04fba0d57b383323c734fcbb8a9d47f62` |
| Benchmark runner | `274caeb10f7377b2be6f042674ebed0922527bfc8b8502d8cc96a3b3832a131d` |

### Retained evidence and release scope

Durable local evidence is in
`.gopurs-evidence/publication-rc18-20261112/benchmark-audit` at the workspace
root. It contains `analysis.json`, `cold-comparison.csv`, campaign/summary and
source manifests, runner and completion identities, baseline README/report,
input derivation, all raw records/logs and the 102 output oracles. Both
compressed archives were read back and their contents verified. The
`frozen-inputs.tar.gz` archive includes private b8x inputs and remains local;
this document publishes identities and aggregate results, not those sources.

| Evidence | SHA-256 |
| --- | --- |
| Terminal `audit-completion.json` | `79338162c35291e98ecb378beff9416007c62dbaa43e3e3e74da895d16304424` |
| Complete `SHA256SUMS` list | `054f3fa1f18141bf89253377deaa3bca9c46cf5f1684e2641eaedf3082cf89ee` |
| `case-records.tar.gz` | `cddb86ab047ffbb803ba095b729831c7d5511440c0e669d905f80ae020ad2101` |
| `frozen-inputs.tar.gz` | `e5c1f8629ea40189e98315fc14edd58c4c5abf3d94231af411e3942857c4a599` |

The actual build-cache and go-packages helpers were captured after the
campaign, with binary/source/tool identities and build information. Their
hashes were not recorded at each invocation. The separate helper receipt is
retained under the same RC18 publication-evidence directory with SHA-256
`0e7ce253b7b080c5146bb678e322f3064f0c084fcaec8739315c033dc85fdbd4`.
No resolver or rebuild was used for that capture.

RC18's CI qualification was interrupted by a separate test-fixture failure
and RC18 remains unpublished; see the [quality record](quality-0.1.0.md).
This completed local benchmark does not qualify a public package, another
platform, RC19 or the final release. Referring a later candidate to this
campaign requires an explicit proof of equality for all executed compiler,
launcher, helper and runtime bytes, as well as its inputs, tools and protocol.
That proof is pending. Any such reference must retain the words **RC18
campaign on identical bytes**, not claim a new run. Each later version still
requires its own CI, installed-artifact application and public consumer gates;
an executed-byte difference requires a new strict campaign.

## RC17 interrupted campaign — 12 November 2026

The RC17 local build completed 27 corpora before the frozen `node-http` entry
failed validation: its reachable `stdout` foreign binding had no Go
implementation in the 25 October archive. The implementation already present
in the current library port fixes generation when added alone to an isolated
copy; the compiler and the 312 TAST modules are unchanged.

The original archive, partial measurements and failure are retained. They are
not a completed release benchmark. Corrected entries must be identified as
derived inputs, and their added FFI must be disclosed when comparing with the
official historical baseline. See the
[qualification record](quality-0.1.0.md#rc17-qualification-interruption-and-rc18-preparation--12-november-2026)
for the diagnostic evidence.

## Completed local campaign — 27 October 2026

All **51 corpora passed**, each on its first attempt, between 08:37:40 and
10:19:22 UTC. These are the official table's **b8x + 50 library corpora**;
`gopurs-quickcheck`, one of the 51 bootstrap source pins, has no standalone
case in that table. The campaign used frozen 25 October TAST/FFI inputs,
including b8x's 2,707 modules, and the untagged compiler snapshot
`0c9a6a1f69d64473ae240f547312fd97f6dd45cb` prepared on 26 October.

The table below sums the 51 per-corpus medians, in milliseconds. Each median
uses three measured fresh processes after one excluded warmup. The warm
column measures an unchanged-input cache hit; it has no historical reference.
These totals are neither the campaign's elapsed duration nor application
build/execution times.

| Host/profile | Official 25 October cold | Current cold | Change | Current warm |
| --- | ---: | ---: | ---: | ---: |
| JavaScript Release | 812,478 | 809,532 | −0.36% | 13,111 |
| Native Release | 332,715 | 288,082 | −13.41% | 5,808 |
| Native Dev | 265,633 | 221,643 | −16.56% | 5,708 |

All 51 native Release and all 51 native Dev cold medians were below their
respective historical values. JavaScript was lower in 37 cases and higher in
14: its aggregate differs by only −0.36%, while individual cases move in both
directions. The largest
JS increases by absolute time were b8x (146,389 → 147,829 ms, +0.98%),
argonaut-core (15,672 → 16,647 ms, +6.22%) and aff (16,023 → 16,802 ms,
+4.86%). The full per-case comparisons and all samples are retained; these
increases are not hidden by the aggregate.

For b8x specifically, native Release was 70,353 ms versus 78,037 ms
(−9.85%), and native Dev was 47,227 ms versus 53,611 ms (−11.91%). Its warm
medians were 479 ms JS, 324 ms native Release and 308 ms native Dev. This is
generation from the archived TAST, separate from the complete application
build qualification in the
[quality record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/quality-0.1.0.md).

The completion audit independently recomputed all medians from **1,224 raw
records and logs: 918 measured samples and 306 warmups**. It verified 612
explicit cold misses, 612 exact warm hits, the required variant/round order,
and every recorded exit status. No invocation failed, timed out, used a full
fallback or silently selected JS because a native executable was missing.
For each corpus, all JS/native Release output fingerprints matched the
current JS oracle; Dev matched its separate oracle across all repetitions.
The 102 per-file oracle manifests are retained. No output was regenerated or
application executed by this audit.

Measurements ran on the Apple M4 Pro (14 logical CPUs, 48 GiB RAM), Darwin
25.1.0, Node 24.8.0 and Go 1.27.0. The source/FFI/TAST definitions and official
README/report hashes matched their preparation records. Workers, feature flags
and native GC policy matched the protocol below. The 1 GiB Go cache eviction
policy was configured but **never triggered**: its maximum recorded size
after a case was 99,801,914 bytes. No timings were trimmed or retried.

The new campaign interleaves JS, native Release and native Dev, with an
immediate warm invocation after each cold invocation. The historical JS v2
measurements preceded the native v3 measurements, and had no paired warm
invocations. Both campaigns enable packages, overlays and build manifests;
the older 4–5 October protocol differences are not the comparison used here.
The measured compiler binaries also differ from the historical binaries.
Consequently, these results show observed differences on this machine, not a
controlled attribution of gains to a particular compiler change.

The recorded one-minute load before invocations ranged from 2.59 to 10.57
(median 4.79); it includes the campaign's own prior work and does not prove an
otherwise idle machine or stable thermal conditions. Variability remains:
native Release's three js-uri cold samples were 6,783 / 2,788 / 2,826 ms.
The slow sample recorded 3,620 ms in `cache lookup + freeze`, compared with
86 / 106 ms in the others; its cause was not established. That successful
sample remains in the data, and the prescribed median is 2,826 ms. Three
samples per cell do not establish statistical confidence or a bound on latency.

The measured binary hashes are:

| Artifact | SHA-256 |
| --- | --- |
| JavaScript compiler | `e7c761c44aef7276b335265459c03cfe6551b7ed8ec00f76f8233aa6fbe30b3b` |
| Native compiler | `eb97922721f15c3e742ff14335a8294b6005ca321ae2e9a57bb8847776e46298` |

Both still matched the local checkout after its package metadata was advanced
to `0.1.0-rc.1`. This is evidence for those exact binaries and frozen inputs,
not for a tagged CI rebuild, a public install or another platform. The
benchmark does not establish application runtime gains.

Durable local evidence is under
`.gopurs-evidence/release-preparation-20261026/benchmark` at the local workspace root.
Its 72 files occupy 98,766,767 bytes. They include the completed state and
summary, independent analysis, per-case CSV, baseline documents, runner,
definitions, manifests, raw logs/records and exact frozen inputs. The input
archive retains 32,123 files and 43 symlinks, with equal contents deduplicated
as tar hardlinks. Both compressed archives were read back and verified.
Go module/build caches and backend caches are excluded; the Go module-cache
manifest and pinned dependency files remain. The original inputs and timings were not
modified.

| Evidence | SHA-256 |
| --- | --- |
| Completed `campaign.json` | `e83af80e05f3ceaf04bbc96495d7d4e2bcc4148dd1279b08f3718572a9dd872e` |
| Completed `summary.json` | `9361f216e2de0b7cf9aafcbda81aedaf00eb1671a8c50575d4a483af9c30be40` |
| `case-records.tar.gz` | `e97e32431a0915161ed77273605c03629bdab573940f6895cc96291fd261c5d5` |
| `frozen-inputs.tar.gz` | `126c4180687147379210d2446e2797ee7bfbdedcb5414648b12d23d941e99c85` |
| Complete `SHA256SUMS` list | `38cacf4f552662689fb7cb6a4a72a0ac90ddcd1af3a11d906e2e77fd094681fa` |

## Before a measurement

Finish other builds, tests, Nix work and benchmarks on the machine first.
Keep power/thermal conditions stable. The runner records machine information,
timestamps and load averages; those observations do not prove that the machine
was otherwise idle. A validation pilot must not be presented as a performance
measurement.

Build the release candidate first, with JavaScript and native binaries produced
from the same source. `prepare` freezes its npm distribution files and native
executable, source/runtime/lockfile hashes, and the benchmark runner hash. It
also copies the archive's Go module cache into its own directory. Archive
inputs and historical results are never used as writable workspaces.

The compiler's source checkout needs the release tools and normal package
allowlist. The frozen corpus archive contains `cases/*/definition.json`,
`library-corpus/packages`, b8x's case inputs, and `go-mod-cache`. No original
temporary Go export cache is required.

## Functional pilot

Prepare a small corpus without starting any compiler:

```bash
node tools/release/release-benchmark.mjs prepare \
  --archive /path/to/.gopurs-evidence/compilation-readme-20261025 \
  --altbak /path/to/altbak.pub \
  --out /tmp/gopurs-benchmark-pilot \
  --cases gopurs-assert --pilot
```

Run the functional pilot:

```bash
node tools/release/release-benchmark.mjs run --out /tmp/gopurs-benchmark-pilot
node tools/release/release-benchmark.mjs status --out /tmp/gopurs-benchmark-pilot
```

The pilot checks old TAST compatibility, both executable hosts, cold misses,
warm hits and generated-output equality. Its metadata permanently identifies
it as `pilot-functional-validation`. Use a new directory for the real campaign.
A functional pilot may run while other builds or tests are active when its
only purpose is to validate this machinery. Record that concurrent work with
the pilot evidence, and exclude **all** pilot times from performance
comparisons, including the rounds labelled measured by the runner. Load
averages and successful functional checks do not make those times comparable.
The full 51-corpus performance campaign must wait until other work on the
machine has stopped, as required above.

## Full campaign

```bash
node tools/release/release-benchmark.mjs prepare \
  --archive /path/to/.gopurs-evidence/compilation-readme-20261025 \
  --altbak /path/to/altbak.pub \
  --out /path/with/free/space/gopurs-release-benchmark

node tools/release/release-benchmark.mjs run \
  --out /path/with/free/space/gopurs-release-benchmark
```

Omitting `--cases` prepares all 51 corpus definitions. `--compiler DIR` selects
another candidate checkout; otherwise the checkout containing the script is
used. The official README and reference JSON must agree before preparation.
The frozen definitions, TAST, FFI and their sibling source files are checked
against their existing SHA-256 manifests. b8x retains the original pinned
`go.mod` and `go.sum`; it uses the copied module cache with `GOPROXY=off`.

For each corpus there are four rounds: one warmup and three measured rounds.
Each round rotates the variant order. Each variant gets a **cold → warm pair**,
with a separate process for each invocation:

1. Recreate the output from frozen TAST and remove backend caches. Keep the
   case's immutable source/FFI family in the same private location.
2. Invoke the selected compiler and require an explicit backend cache miss.
3. Invoke the same compiler again, preserving that output and cache. Require
   `inputs and outputs verified; 0 modules optimized/emitted`.
4. Verify input hashes and generated Go/module-file hashes after each process.

This produces 24 invocations per corpus: 6 warmup invocations and 18 measured
invocations. A complete campaign has 1,224 invocations, including 918 measured
samples. Every cold and warm cell uses the median of exactly three successful
samples. Warmups and failed attempts never enter medians.

The first JavaScript cold warmup establishes the current Release output oracle.
Every other JS/native Release output, cold or warm, must match its per-file
SHA-256 hashes and sizes. Dev establishes its own oracle and must remain stable
across repetitions. Historical generated code is not imposed as an oracle:
the compiler/runtime may legitimately have changed since the baseline. This
campaign verifies generation equality, not application semantics.

All variants explicitly enable Go packages, overlays and build manifests, eight
preparation/PBO/emission workers and pipelining. Native Release/Dev set
`GOGC=off` and `GOMEMLIMIT=16GiB`; `GOMAXPROCS` is unset. JavaScript does not
inherit those Go GC overrides. The exact selected Node and Go binaries and
their hashes are recorded. Daemon/resident modes are disabled.

“Cold” means a cold **backend/output cache**, with warm operating-system caches.
The Go export/compiler cache is shared within the campaign and warmed by the
excluded warmups. It is trimmed only between complete cases if it exceeds
1 GiB; that event is recorded and the next case still runs all its warmups.
The historical campaign likewise warmed Go export work, but retained its shared
cache across all cases. Include any trimming and machine differences when
assessing comparability.

## Progress, interruption and recovery

```bash
node tools/release/release-benchmark.mjs status --out /path/to/campaign
node tools/release/release-benchmark.mjs run --out /path/to/campaign
```

`campaign.json` is updated atomically after each invocation. It contains each
case's state, raw samples, command, selected environment, wall time, phase
timers and output fingerprint. `cases/NAME/attempt-N/` holds logs, individual
records and oracle manifests. `summary.json` is refreshed after each case;
aggregate medians are summed only when all 51 cases have passed.

Completed cases are retained on resume. An interrupted or failed case is
restarted from its first warmup with:

```bash
node tools/release/release-benchmark.mjs run --out /path/to/campaign --retry-failed
```

Its previous attempt and logs remain archived and excluded. `--cases NAME,...`
can select prepared cases for diagnosis; `--keep-going` continues to other
cases after a failure. There is no retry inside a measured cell. Runner,
compiler, toolchain or baseline changes require a new campaign directory.

One private source/output family is kept at a time. Successful case outputs
are deleted after their manifests are saved; the latest failed workspace stays
under `active/` until a retry or another case replaces it. Canonical generated
Go trees are not duplicated per sample. Go dependencies and the frozen compiler
are copied only once. The default free-space floor is 2 GiB, configurable during
preparation with `--min-free-gib`; hitting it stops before another measured
process starts. A PID lock prevents simultaneous runners in the same campaign.

## Reporting

Keep raw samples, failed/interrupted attempts, hashes, machine/configuration and
both baseline documents with the release evidence. Compare the cold medians to
the official 25 October rows under the stated conditions. Present warm medians
as a separate cache-reuse measurement, with no implied historical speedup.
The earlier baseline had its JS and native measurements collected at different
times; this runner interleaves all three variants, which is another protocol
difference to report. Do not infer application runtime gains from these backend
timings or describe unmeasured platforms as validated.
