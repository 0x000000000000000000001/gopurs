# Documentation

These guides describe the current source tree. Start with the
[project README](../README.md) or choose a guide below.

## Using gopurs

| Guide | Purpose |
| --- | --- |
| [Quickstart](quickstart.md) | Build and run your first application |
| [CLI reference](cli.md) | Options, environment variables and exit status |
| [Migration](migration.md) | Move a JavaScript project to Go |
| [Go FFI](ffi-idiomatic.md) | Implement foreign imports with ordinary Go functions |
| [Limitations](limitations.md) | Unsupported features and compatibility gaps |
| [Troubleshooting](troubleshooting.md) | Diagnose build and runtime errors |
| [Publication scope](publication-scope.md) | Supported platforms and compatibility policy |

## Development and releases

| Guide | Purpose |
| --- | --- |
| [Bootstrap](bootstrap.md) | Build the compiler hosts from pinned inputs |
| [Nix](nix.md) | Install the toolchain and use the development shell |
| [Development](development.md) | Work on library ports, Go FFI and the runtime |
| [Testing](testing.md) | Choose checks, run fixtures and inspect reports |
| [Release procedure](release.md) | Qualify, package and publish a candidate |
| [Release qualification](quality-0.1.0.md) | Evidence for the published release and earlier candidates |
| [Release checklist](release-checklist.md) | Track qualification and publication gates |
| [Library source review](library-source-audit.md) | Review source pins and attribution |
| [Benchmark protocol](benchmark-protocol.md) | Measure and report performance |
| [Release benchmark runner](release-benchmark.md) | Compare compiler hosts on the frozen corpus |

## Compiler internals

| Guide | Purpose |
| --- | --- |
| [Architecture overview](architecture-overview.md) | Follow the compilation pipeline |
| [Architecture](architecture.md) | Pass order, metadata and module boundaries |
| [Module map](module-map.md) | Locate implementations and their dependencies |
| [Naming](naming.md) | Understand compiler terminology |
| [Build cache](build-cache.md) | Cache keys, storage, publication and reuse |
| [Go AST and printer](go-ast-printer.md) | Separate code generation from formatting |
| [Runtime and FFI contracts](runtime-ffi-contracts.md) | Distinguish compiler hosts from generated programs |
| [Parallel preparation](parallel-preparation.md) | Schedule preparation jobs |
| [Parallel emission](parallel-emission.md) | Schedule optimization and emission |
| [Runtime closures](runtime-closures.md) | Understand closure lifetime and heap escape |

Optimization references:

- [ADT reuse](adt-reuse.md), [array round-trips](array-roundtrip.md),
  [function fusion](function-fusion.md) and [Int32 loops](int32-loops.md).
- [JSON decoders](json-decoder-optimizations.md),
  [record construction](json-record-plans.md) and
  [JSON traversals](json-traversals.md).
- [Native record workers](native-record-workers.md),
  [packed strings](packed-strings.md) and [Rust PGO](pgo.md).

## Keeping these guides current

Write documentation in English, describe implemented behavior and link to the
source or tests that enforce a contract. Keep campaign logs, abandoned designs
and completed development plans in Git history. The
[documentation before this cleanup](https://github.com/0x000000000000000000001/gopurs/tree/7589ec34aed08ec3b7dfc52014a4253b007b2fdc/docs)
retains those records, including their raw measurements and probe scripts.

Link new guides from this index. Run `npm run check:docs` after editing: it
checks relative links and anchors, packaged files and module-map coverage.
Use the [altbak README](https://github.com/0x000000000000000000001/altbak.pub#go)
as the official performance baseline.
