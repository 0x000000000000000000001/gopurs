# Release backend benchmark

`tools/release-benchmark.mjs` compares the release JavaScript host, native
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
node tools/release-benchmark.mjs prepare \
  --archive /path/to/.gopurs-evidence/compilation-readme-20261025 \
  --altbak /path/to/altbak.pub \
  --out /tmp/gopurs-benchmark-pilot \
  --cases gopurs-assert --pilot
```

Run the functional pilot:

```bash
node tools/release-benchmark.mjs run --out /tmp/gopurs-benchmark-pilot
node tools/release-benchmark.mjs status --out /tmp/gopurs-benchmark-pilot
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
node tools/release-benchmark.mjs prepare \
  --archive /path/to/.gopurs-evidence/compilation-readme-20261025 \
  --altbak /path/to/altbak.pub \
  --out /path/with/free/space/gopurs-release-benchmark

node tools/release-benchmark.mjs run \
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
node tools/release-benchmark.mjs status --out /path/to/campaign
node tools/release-benchmark.mjs run --out /path/to/campaign
```

`campaign.json` is updated atomically after each invocation. It contains each
case's state, raw samples, command, selected environment, wall time, phase
timers and output fingerprint. `cases/NAME/attempt-N/` holds logs, individual
records and oracle manifests. `summary.json` is refreshed after each case;
aggregate medians are summed only when all 51 cases have passed.

Completed cases are retained on resume. An interrupted or failed case is
restarted from its first warmup with:

```bash
node tools/release-benchmark.mjs run --out /path/to/campaign --retry-failed
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
