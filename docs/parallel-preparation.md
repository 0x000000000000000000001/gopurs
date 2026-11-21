# Parallel preparation

Preparation discovers which generic bindings need specialized versions and
collects the information used to build them. Independent jobs can run together,
but later rounds must see the same merged results as a sequential run. This
page describes that ordering contract and the scheduling implementation.

`GOPURS_PREPARE_JOBS` controls preparation concurrency, independently of TAST
loading and Go emission. The compiler default is 2, bounded to 1–8; the launcher
selects 8 on machines with at least 32 GiB unless explicitly overridden.
A value of 1 runs jobs sequentially.

Specializing one binding can reveal further calls that need specialization,
so collection proceeds in rounds. Each round reads an immutable specialization
table and cache. A job handles one qualified name and specialization key.

PBO merges contributions and cache updates in input order before starting the
next round. When contributions overlap, the first payload keeps priority and
caller sets are combined. Parallel execution therefore preserves both the
merge rules and the condition for stopping collection.

[`Gopurs.Preparation`](../src/Gopurs/Preparation.purs) uses a bounded worker pool.
Workers claim small contiguous slices from a per-run counter; results are sorted
by their starting index before publication. `runPreparationJobs` defers pure
jobs inside Aff, with slices of at most 16 jobs. `mapPreparationJobs` schedules
effectful jobs similarly, with slices of at most 8. Native Aff runs workers in
Go goroutines; this JavaScript path does not provide CPU multithreading.

Creating the Aff action for pure jobs must defer their execution, even for
sequential or empty input. Running the same action again recomputes every
result. The dispatcher controls scheduling; PBO still owns its cache, merge
rules and stopping condition. Specialization barriers apply only
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
