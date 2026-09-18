# Parallel Go emission

`GOPURS_EMIT_JOBS` selects the maximum number of modules emitted together
(default 2; set 1 for sequential emission).
It is independent of `GOPURS_JOBS`, which only controls TAST loading.

The sequential PBO builder enqueues optimized modules. A full batch, a module
depending on a queued module, or the end of the build flushes the queue. The
dependency check uses optimized imports, including references introduced by
inlining. Each batch receives one immutable function-signature snapshot;
signatures are merged in source traversal order after all workers finish.
The Go translation itself runs inside deferred Aff computations.

PBO optimization stays sequential. Its current builder passes the preceding
module's exported directives to the next module, so scheduling optimization
solely by the original import graph would change this behavior.

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
