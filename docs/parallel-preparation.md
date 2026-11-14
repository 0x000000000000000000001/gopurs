# Parallel preparation

`GOPURS_PREPARE_JOBS` controls preparation concurrency, independently of TAST
loading and Go emission. The compiler default is 2, bounded to 1–8; the launcher
selects 8 on machines with at least 32 GiB unless explicitly overridden.
A value of 1 runs jobs sequentially.

Each transitive monomorphization round reads an immutable specialization table
and cache. A job handles one qualified name and specialization key. PBO merges
its contribution and cache update in input order before starting the next round,
preserving first-payload priority, caller unions and the stopping condition.

[`Gopurs.Preparation`](../src/Gopurs/Preparation.purs) uses a bounded worker pool.
Workers claim small contiguous slices from a per-run counter; results are sorted
by their starting index before publication. `runPreparationJobs` defers pure
jobs inside Aff, with slices of at most 16 jobs. `mapPreparationJobs` schedules
effectful jobs similarly, with slices of at most 8. Native Aff runs workers in
Go goroutines; this JavaScript path does not provide CPU multithreading.

Constructing a pure-job Aff must not run jobs, even for sequential or empty
input. Reusing the action recomputes every result. The dispatcher owns neither
PBO's cache nor its merge and stopping rules. Specialization barriers apply only
after complete transitive collection. See [metadata provenance](architecture.md#metadata-provenance)
for the distinction between source preparation and signatures published during
emission.

## Validation

After building the compiler:

```sh
node --test tools/monomorphization.test.mjs tools/preparation.test.mjs
GOPURS_NATIVE_OUTPUT=/path/to/bootstrap/output node --test tools/preparation-native.test.mjs
```

The native test checks deferred execution, reruns, exactly-once execution,
overlap, concurrency bounds and result order under the race detector. Its Go
fixture is installed temporarily in the retained bootstrap and removed afterward.
