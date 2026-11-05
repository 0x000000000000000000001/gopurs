# TAST loading measurements — September 2026

These historical measurements were originally recorded in the project README.
They cover TAST loading and sorting, not complete builds or generated-program
performance. The [benchmark protocol](../benchmark-protocol.md) defines the
measurement boundaries; use the [altbak baselines](https://github.com/0x000000000000000000001/altbak.pub#go)
for application performance comparisons.

On 2026-09-21, after removing whole-array conversions from individual indexed
reads, an alternating comparison on the same b8x TAST and native binary measured
21.077/21.174 s with one worker and 13.659/13.536 s with eight: about 36% less
elapsed loading time, with similar CPU time. Peak RSS rose from about 1.59 GiB
to 2.24 GiB. These measurements cover loading and sorting, not a complete build.

For historical context, on 2026-09-18 a sample of 133 b8x modules
(7.1 MiB of TAST) took 1.26 s with one worker, 1.70 s with four and 2.39 s with
eight with the earlier implementation (two runs per setting, without race instrumentation). Limiting
`GOMAXPROCS` to four did not reverse this regression. These measurements cover
loading and sorting only, not a complete b8x build or altbak runtime benchmarks.
The native race check passed on this sample, and Node tests compare complete
decoded modules and dependency order between sequential and parallel loading.
