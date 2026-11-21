# Documentation

Start with the guide for your task, then use the reference pages as needed.
These pages describe the current source tree. When installing a published
release, use the documentation at its tag, as explained in the
[project README](../README.md).

## Using gopurs

For a first application, follow the [quickstart](quickstart.md) through to a
running executable. For an existing JavaScript project, follow the
[migration walkthrough](migration.md#from-a-javascript-project). To implement a
foreign import, start with the [Go FFI guide](ffi-idiomatic.md).

The [task list API](../examples/task-list-api/README.md) is a complete HTTP, JSON
and SQLite example. Its README explains the development checkouts it needs.

For everyday reference:

| Guide | Purpose |
| --- | --- |
| [CLI reference](cli.md) | Options, environment variables and exit status |
| [Limitations](limitations.md) | Unsupported features and compatibility gaps |
| [Troubleshooting](troubleshooting.md) | Diagnose build and runtime errors |
| [Publication scope](publication-scope.md) | Supported platforms and compatibility policy |

## Development and releases

Use these guides when changing the compiler or its library ports.

| Guide | Purpose |
| --- | --- |
| [Bootstrap](bootstrap.md) | Build the compiler hosts from pinned inputs |
| [Nix](nix.md) | Install the toolchain and use the development shell |
| [Development](development.md) | Work on library ports, Go FFI and the runtime |
| [Testing](testing.md) | Choose checks, run fixtures and inspect reports |
| [Benchmark protocol](benchmark-protocol.md) | Measure compiler and application performance |
| [Release benchmark runner](release-benchmark.md) | Compare compiler hosts on a fixed set of inputs |

For maintainers preparing a release:

| Guide | Purpose |
| --- | --- |
| [Release quickstart](release-quickstart.md) | Prepare, qualify and publish a new version |
| [Release procedure](release.md) | Qualify, package and publish a candidate |
| [Release qualification](quality-0.1.0.md) | Evidence for the published release and earlier candidates |
| [Release checklist](release-checklist.md) | Track qualification and publication gates |
| [Library source review](library-source-audit.md) | Review source pins and attribution |

## Compiler internals

Read the [architecture overview](architecture-overview.md) first. Use the
[module map](module-map.md) to find the code for a particular responsibility;
the remaining pages explain individual parts of the compiler.

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

For work on a specific optimization:

- [ADT reuse](adt-reuse.md), [array round-trips](array-roundtrip.md),
  [function fusion](function-fusion.md) and [Int32 loops](int32-loops.md).
- [JSON decoders](json-decoder-optimizations.md),
  [record construction](json-record-plans.md) and
  [JSON traversals](json-traversals.md).
- [Native record workers](native-record-workers.md),
  [packed strings](packed-strings.md) and [Rust PGO](pgo.md).

## Historical records

The [FFI baseline](ffi-baseline.md) and [FFI validation report](ffi-validation.md)
record the work done on 12 November 2026. Their measurements and intermediate
limits belong to that work. Use the [Go FFI guide](ffi-idiomatic.md) for the
current interface and the [quality record](quality-0.1.0.md) for dated release
results.

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
