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
