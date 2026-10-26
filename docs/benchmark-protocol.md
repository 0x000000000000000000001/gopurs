# Benchmark protocol

This page defines what the published gopurs measurements mean and what a new
measurement must record. It is not a performance guarantee.

## Two different benchmarks

- **Backend compilation time**: the gopurs driver's internal `backend total`,
  covering loading, preparation, optimization, emission and cache publication.
- **Application runtime**: execution of the generated Go binary, measured by the
  workload's own harness.

Do not compare the two or present a backend gain as an application gain without
a corresponding execution measurement.

## Reference conditions

The compilation table linked from the README was measured on an Apple M4 Pro
with 10 performance cores and 4 efficiency cores, without explicit CPU
affinity. The host ran the pinned Go, Node.js and gopurs revisions recorded with
the campaign.

Every published row must record:

- host CPU/OS and memory;
- gopurs version/commit and compiler host (`js`, `native` or `rust`);
- typed `purs` fork revision;
- Go toolchain version;
- Node.js version when the JavaScript host or npm tooling is involved;
- corpus name, module/type counts and TAST hash or frozen input path;
- cache state (cold, warm or incremental), worker limits and memory settings;
- number of warmups and measured repetitions;
- median and raw samples;
- what is included and excluded from the timer.

## Backend protocol

The reference protocol is:

1. Freeze the typed TAST inputs and record their hashes.
2. Start from a fresh generated output and empty backend cache for a cold
   measurement.
3. Run one warmup, then three measured fresh processes per cell.
4. Report the median per cell and keep the raw samples.
5. Use the internal `backend total` timer.
6. Exclude the PureScript frontend, compiler bootstrap, final `go build`,
   application execution and process startup/exit.
7. Run variants in alternating order when comparing compilers to avoid a
   one-sided warm-cache advantage.
8. Publish failures and excluded attempts separately; they must not enter a
   median.

For the 25 October 2026 refresh published in altbak, the measured native
variants used explicit `GOGC=off` and `GOMEMLIMIT=16GiB` after a launcher check
showed that the sandbox could deny memory discovery. The JavaScript variant
retained its own launcher environment. See
[the refresh record](https://github.com/0x000000000000000000001/altbak.pub/blob/main/docs/benchmark-results/2026-10-25-gopurs-compilation-refresh.md)
and the local [cold campaign](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cold-20261024.md) and
[G06 measurements](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-go-measurements.md) for campaign settings and
sources.

## Application protocol

The application table in the altbak README is separate. It records execution
time for a fixed workload and identifies the runtime and machine. A result is
only comparable when the workload, input data, compiler output, runtime flags
and host are the same.

Parallelism results must state the number of threads and the Go runtime
settings. A multicore result does not establish single-core performance or
scaling for another workload.

## Recorded commands and scripts

The protocol is implemented by the campaign records, not by an informal
stopwatch:

- [Cold backend campaign, 24 October 2026](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-cold-20261024.md)
  and its CSV.
- [G06 package-layout measurements](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/build-go-measurements.md).
- [Parallel emission](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/parallel-emission.md)
  and [parallel preparation](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/parallel-preparation.md).
- The altbak refresh record linked above, including the raw JSON and campaign
  scripts.

A backend invocation has this shape:

```bash
GOPURS_GO_PACKAGES=1 \
GOPURS_PBO_JOBS=8 GOPURS_PREPARE_JOBS=8 GOPURS_EMIT_JOBS=8 \
bin/gopurs --main Main
```

Use the driver's phase output for `backend total`; do not substitute process
wall time. For a controlled comparison, alternate the variants and keep the
cache policy fixed.

## Reproducing a result

Use the release tag, not an arbitrary working tree:

```bash
git clone https://github.com/0x000000000000000000001/gopurs.git
cd gopurs
git checkout v0.1.0
npm ci
```

Use the typed compiler revision and Go version recorded for that release.
Generate the corpus, then run the compiler with the documented worker and cache
settings. Record the exact commands, hashes and raw output.

The benchmark scripts under `tools/` are development tools. They are not part of
the npm runtime surface; use the versions in the tagged repository.

## Reporting

When publishing a comparison:

- link the campaign record and raw JSON/CSV;
- state whether the result is cold or warm;
- state the number of samples and the statistic used;
- include excluded and failed attempts;
- do not describe an unmeasured platform or corpus as validated;
- do not describe a backend-only gain as a runtime gain.
