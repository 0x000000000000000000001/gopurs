# Parallel Go emission

`GOPURS_EMIT_JOBS` selects the maximum number of modules emitted together
(default 8; set 1 for sequential emission).
It is independent of `GOPURS_JOBS`, which only controls TAST loading.

The sequential PBO builder enqueues optimized modules. By default it can optimize
the next modules while the preceding batch emits Go. `GOPURS_PIPELINE=0` disables
this overlap; `GOPURS_EMIT_JOBS=1` also retains immediate sequential emission.

At most one emission batch is active and one is pending, in addition to the
module currently being optimized by the producer. A full batch, a module
depending on a queued module, or the end of the build flushes the queue. The
dependency check uses optimized imports, including references introduced by
inlining. A new batch starts only after the preceding batch finishes. Its
immutable function-signature snapshot is read at that point; signatures are
merged in source traversal order after all module workers finish. Go translation
runs inside deferred Aff computations on native Aff goroutines.

`finish` waits for emission. On producer failure, `cancel` stops new admissions,
discards the pending batch, and drains the active batch naturally. It does not
interrupt filesystem effects: their `nonCanceler` does not guarantee that the
underlying OS callback has stopped. Main brackets the builder with this cleanup.
An emission error still propagates through `enqueue` or `finish`.

PBO optimization stays sequential. Its current builder passes the preceding
module's exported directives to the next module, so scheduling optimization
solely by the original import graph would change this behavior.

## Worker counts after forceEscape correction, 2026-09-21

The default is now **8**. On the corrected native runtime, three alternating
passes per setting on 304 modules compared 1/2/4/8 emission workers, keeping
preparation at 2. Mean optimize/emit times were 6.915 / 5.438 / 5.517 / 5.115 s.
All 399 Go files matched. Moving from 1 to 2 also enables producer overlap.

On b8x, two isolated controls at 2 workers took 124.229 and 132.346 s for
optimize/emit, versus 113.532 s at 8. Backend times were 203.089 / 208.951 s
versus 186.326 s. Loading and preparation varied too, so the full difference
cannot be attributed to emission. Peak RSS at 8 remained close to the controls.

A real `b -c -n` rebuilt the compiler with the new default, without any
`GOPURS_*` environment overrides: 120.331 s optimize/emit, 197.565 s backend,
304.738 s for the command. This validates the default and provides a second
observation at 8, under a different workflow from the isolated backend runs.
All 2,655 TAST and 2,959 Go files matched; the 10 emission tests passed.
The complete command did not improve over the previous historical profile;
that profile used pprof/gctrace and is not a controlled whole-workflow comparison.

Preparation remains at 2: increasing it to 8 did not establish a clear benefit
beyond run variation. These measurements use GOMAXPROCS=14 on this machine;
8 is a measured choice, not a universal optimum. PBO optimization remains
sequential. Earlier worker-count conclusions below predate the runtime fix.

[Detailed measurements](../../../scratch/gopurs-workers-20260921/rapport.md).

## Pipeline measurements on 2026-09-20

A 304-module altbak corpus was compiled in four isolated native processes with
the same binary, toggling the pipeline in the order 0 / 1 / 1 / 0. Emission used
two workers throughout. Mean optimize/emit time decreased from 8.219 to 7.714 s
(6.1%); mean total backend time decreased from 13.627 to 13.049 s (4.2%). All
399 generated Go files matched the preceding compiler byte for byte.

On the real b8x `b -c -n`, optimize/emit decreased from 289.795 to 256.177 s
(11.6%), backend time from 424.522 to 388.011 s (8.6%), and the whole command from
528.775 to 490.771 s (7.2%). All 2,655 TAST inputs and 2,959 generated Go files
matched. This is one full pair, supported by the smaller alternating runs;
variation in unchanged preparation/bootstrap phases is not attributed to the
pipeline. The existing sourcemap/config failures remain unchanged.

The full profile preceded the final producer-failure cleanup adjustment from
interrupting to draining the active batch. That branch was not exercised by the
successful backend run; its successful emission path is unchanged. Final cleanup
is separately covered by JS/native tests and a rebuilt native compiler.

## Worker count with the pipeline on 2026-09-20

Nine sequential native processes compared 2 / 4 / 8 workers with the pipeline
enabled, using the same binary and frozen inputs from the 304-module corpus.
After one excluded warmup, the order was 2/4/8, 8/2/4, 4/8/2. All 399 generated
Go files matched the historical reference in every run.

| Emission workers | Mean optimize + emit | Mean backend total |
| --- | --- | --- |
| 2 | 7.482 s | 12.671 s |
| 4 | 7.613 s | 12.751 s |
| 8 | 7.737 s | 12.886 s |

Neither higher setting improved compilation time on this corpus; medians led to
the same decision. The default remains 2. No new full b8x run was warranted by
this experiment, and no further b8x gain is claimed.

## Pipeline wait diagnosis on 2026-09-20

Scratch-only probes on two runs of the frozen 304-module corpus measured a mean
7.783 s optimize/emit phase. Synchronous PBO conversion occupied 6.554 s (84.2%),
including 2.700 s overlapping emission. Producer joins occupied 1.210 s (15.5%);
the final join was only 1.56 ms. An emission batch was active for 3.914 s (50.3%).
These wall-time intervals overlap and must not be added as CPU times.

Both captures observed 304 PBO conversions, 304 translations and 175 ordered
batches, with at most two translations in flight. All output matched two
uninstrumented controls. Instrumented phase time averaged 2.64% above controls;
measurement overhead and run variation are not separated by this small sample.
The result prioritizes studying independent work inside the sequential PBO
producer. It is not extrapolated to b8x and does not establish that module
optimization can safely ignore preceding directives.

## Measurements on 2026-09-18

Go 1.27.0, darwin/arm64. Existing TAST inputs, isolated output directories,
separate processes, sequential TAST loading, no simultaneous benchmark runs.
These are backend compilation measurements, separate from the official
program-execution baselines in `altbak.pub/README.md`.

For altbak's 300 modules, two runs per setting:

| Emission workers | Optimize + emit | Backend total |
| --- | --- | --- |
| 1 | 12.735 / 13.230 s | 18.055 / 18.564 s |
| 2 | 12.395 / 12.227 s | 17.816 / 17.557 s |
| 4 | 12.251 / 12.305 s | 17.729 / 17.752 s |

All 392 generated Go files matched the pre-change compiler byte for byte in
all six runs. Two workers reduced the average optimize/emit duration by about
5%, and total backend time by about 3%. These results do not imply the same
percentage for the complete `b -c -n` command, which also rebuilds the compiler
and runs other tools.

The full b8x reference contains 2,657 modules (204.6 MiB of TAST). Its
pre-change native backend took 614.509 s: 34.722 s loading/sorting,
133.664 s preparing/monomorphizing, and 446.120 s optimizing/emitting.
With two emission workers, the backend took 581.421 s: 33.306 s loading/sorting,
128.409 s preparing/monomorphizing, and 419.704 s optimizing/emitting.
All 2,962 Go files matched byte for byte. This full before/after pair measured
26.416 s saved in the changed phase (5.9%) and 33.088 s overall (5.4%). The
unchanged loading and preparation phases also varied; their difference is not
attributed to parallel emission. This is one full b8x pair, supported by the
repeated smaller altbak comparison above, not a guarantee for every build.

## Validation

After building, run `node --test tools/emission.test.mjs`. The tests cover
immediate sequential emission, bounded batches, partial final batches,
dependency barriers with asynchronous completion, and error propagation.
The integration comparison hashes every generated Go file against the
pre-change compiler on identical TAST inputs.

A native compiler built with `go build -race` also completed emission with two
workers on a 98-module corpus (the dependency closure of `Data.Array`,
`Data.List`, and `Data.Map`) without a race report. This check is separate from
the uninstrumented timing runs.

The pipeline adds tests for producer overlap, bounded backpressure, ordered
snapshots, failure propagation, and non-preemptive cleanup. Native tests are in
`tools/emission-native_test.go`: copy them into a retained bootstrap's
`output/purescript` directory and run the command in the file header. The native
pipeline also completed the full 304-module corpus under `go build -race` with
no race report and identical output; race timings are excluded from benchmarks.
