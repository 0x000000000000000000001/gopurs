# Documentation map

This directory contains both normative documentation and historical engineering
records. Read the section that matches your question; do not treat a benchmark
campaign as a contract.

## Start here

| Document | Audience | Status |
| --- | --- | --- |
| [README](../README.md) | users, newcomers | project overview, install, quickstart |
| [quickstart.md](quickstart.md) | users | end-to-end first program |
| [cli.md](cli.md) | users | CLI options, environment, exit status |
| [limitations.md](limitations.md) | users | known gaps and unsupported features |
| [migration.md](migration.md) | users | compatibility between releases |
| [troubleshooting.md](troubleshooting.md) | users | common failures |
| [publication-scope.md](publication-scope.md) | users | support and compatibility policy |

## Internals

| Document | Subject | Status |
| --- | --- | --- |
| [architecture-overview.md](architecture-overview.md) | pipeline in one page | normative |
| [architecture.md](architecture.md) | detailed pass order, metadata provenance, responsibility table | normative, being split |
| [naming.md](naming.md) | glossary, naming rules and historical names | normative |
| [module-map.md](module-map.md) | phase, responsibility and dependencies of every module | normative, source of truth for the cleanup moves |
| [build-cache.md](build-cache.md) | persistent cache protocol | normative |
| [runtime-ffi-contracts.md](runtime-ffi-contracts.md) | runtime and host FFI ownership | normative |
| [go-ast-printer.md](go-ast-printer.md) | boundary between Go AST and printer | normative |
| [parallel-preparation.md](parallel-preparation.md) | preparation jobs and scheduling | normative |
| [parallel-emission.md](parallel-emission.md) | emission batching | normative |
| [runtime-closures.md](runtime-closures.md) | closure lifetime and heap escape | normative |
| [adt-reuse.md](adt-reuse.md) | tree-cell reuse | reference |
| [array-roundtrip.md](array-roundtrip.md) | integer-array round-trip rule | reference |
| [function-fusion.md](function-fusion.md) | counted function producers | reference |
| [int32-loops.md](int32-loops.md) | signed 32-bit loop fast path | reference |
| [json-decoder-optimizations.md](json-decoder-optimizations.md) | JSON decoder codegen | reference |
| [json-record-plans.md](json-record-plans.md) | shared record construction | reference |
| [json-traversals.md](json-traversals.md) | fused Either traversals | reference |
| [native-record-workers.md](native-record-workers.md) | native read-only record workers | reference |
| [packed-strings.md](packed-strings.md) | packed String payload | reference |
| [pgo.md](pgo.md) | profile-guided optimization for Rust | reference |

## Build, release and benchmarks

| Document | Subject |
| --- | --- |
| [bootstrap.md](bootstrap.md) | bootstrap chain and pinned inputs |
| [benchmark-protocol.md](benchmark-protocol.md) | measurement protocol and reporting rules |
| [release.md](release.md) | release procedure |
| [release-benchmark.md](release-benchmark.md) | release backend benchmark |
| [quality-0.1.0.md](quality-0.1.0.md) | quality-gate status for 0.1.0 |
| [nix.md](nix.md) | Nix installation and development |
| [library-source-audit.md](library-source-audit.md) | library source pins and licenses |

## Testing

[testing.md](testing.md) is the test entry point: command selection, contracts
and historical campaign logs. It is intentionally large for now; the plan
[cleanup.md](../cleanup.md) separates the command table and contracts from the
historical lots.

## History and evidence

The `build-*.md`, `build-*.csv`, `build-*.json`, `quality-*.md`,
`release-benchmark.md` and `plan-archive-*.md` files are dated engineering
records. They document what was measured, not what is guaranteed. Keep them
for provenance; do not use them as the current contract.

- Build campaigns: `build-*.md`
- Cost profiles and measurements: `build-cost-*`, `build-preparation-*`
- Plans and archives: `plan-archive-20261019.md`, [../todo.md](../todo.md)
- Cleanup baseline: [history/cleanup-baseline.md](history/cleanup-baseline.md)

## Conventions

- **Normative** documents describe current behaviour and must be updated with
  the code.
- **Reference** documents explain one rule or optimisation and may lag behind
  the code; link the tests that enforce them.
- **Historical** documents are frozen. Add a date in the title or the first
  paragraph.
- Public documentation is written in English. Internal architecture and
  historical records may be in French; do not mix languages within one page.
- New documents must be linked from this page, and relative links must resolve
  before merge. Run `node tools/check-docs.mjs` (or `npm run check:docs`) to
  check relative links, anchors, packaged files and module-map coverage.
