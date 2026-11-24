# Documentation

Choose a guide for your task. These pages describe this checkout; for a published
release, use the documentation at its tag.

## Using gopurs

| Task | Guide |
| --- | --- |
| Build a first application | [Quickstart](quickstart.md) |
| Move an existing JS project to Go | [Migration](migration.md#from-a-javascript-project) |
| Write a Go foreign function | [Go FFI](ffi-idiomatic.md) |
| Follow a complete HTTP/JSON/SQLite example | [Task list API](../examples/task-list-api/README.md) |
| Find a command or setting | [CLI reference](cli.md) |
| Fix an error or check support | [Troubleshooting](troubleshooting.md), [limitations](limitations.md) |
| Install with Nix | [Nix guide](nix.md) |

## Development and releases

For contributors changing the compiler or its library ports:

| Task | Guide |
| --- | --- |
| Build a compiler host | [Bootstrap](bootstrap.md) |
| Develop a library, FFI or runtime change | [Development](development.md) |
| Choose and run checks | [Testing](testing.md) |
| Measure a change | [Benchmark protocol](benchmark-protocol.md), [release benchmark runner](release-benchmark.md) |
| Publish a version | [Release quickstart](release-quickstart.md) |
| Check publication requirements or recover a release | [Detailed release reference](release.md) |
| Review supported interfaces and source pins | [Publication scope](publication-scope.md), [library source review](library-source-audit.md) |

The [release checklist](release-checklist.md) and [quality record](quality-0.1.0.md)
contain dated qualification evidence, not a guarantee about an untested checkout.

## Compiler internals

These are **technical references for compiler maintainers**. Start with the
[architecture overview](architecture-overview.md); open a contract page when
changing that part of the implementation.

| Area | Reference |
| --- | --- |
| Pipeline and code locations | [Architecture](architecture.md), [module map](module-map.md), [naming](naming.md) |
| Caching and scheduling | [Build cache](build-cache.md), [parallel preparation](parallel-preparation.md), [parallel emission](parallel-emission.md) |
| Go output and runtime boundary | [Go AST/printer](go-ast-printer.md), [runtime/FFI contracts](runtime-ffi-contracts.md), [closure lifetime](runtime-closures.md) |
| Loops and allocation | [ADT reuse](adt-reuse.md), [array round-trips](array-roundtrip.md), [function fusion](function-fusion.md), [Int32 loops](int32-loops.md) |
| JSON | [Decoders](json-decoder-optimizations.md), [record plans](json-record-plans.md), [traversals](json-traversals.md) |
| Representations | [Record workers](native-record-workers.md), [shared containers](shared-container-layouts.md), [packed strings](packed-strings.md) |
| Experimental Rust host | [PGO](pgo.md) |

## Historical records

[FFI baseline](ffi-baseline.md) and [FFI validation](ffi-validation.md) describe
the work of 12 November 2026. Use [Go FFI](ffi-idiomatic.md) for the current
interface. Older designs and campaign logs remain in
[Git history](https://github.com/0x000000000000000000001/gopurs/tree/7589ec34aed08ec3b7dfc52014a4253b007b2fdc/docs).

## Keeping these guides current

Keep user guides short and task-oriented. Name the audience of technical
references; keep dated evidence distinct from current instructions. Preserve
commands, contracts and proof data when simplifying prose.

Link new pages here and run `npm run check:docs` to check links, anchors, packaged
files and module-map coverage. Use the [altbak README](https://github.com/0x000000000000000000001/altbak.pub#go)
as the official performance baseline.
