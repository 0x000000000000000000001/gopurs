# Terminology and naming

Use these terms when describing the current compiler. Source links and module
names must match the checkout; the [module map](module-map.md) lists the
actual modules and dependencies.

## Glossary

| Term | Meaning |
| --- | --- |
| `compile` | Complete driver invocation, from options through publication. |
| `driver` | Orchestration, phase ordering and final result. |
| `prepare` | TAST loading, metadata, specialization and preparation reuse. |
| `specialize` | Monomorphization of a module or binding. |
| `rewrite` | A transformation of the typed/backend IR before emission. |
| `build` | PBO construction of a `BackendModule`; qualify final application builds as `go build`. |
| `emit` | Production of Go declarations in memory. |
| `output` | Writing generated files into staging or public output. |
| `publish` | Verified cache commit, including removal of obsolete owned outputs. |
| `worker` | A generated Go function, local or at module scope. |
| `job` | A parallel unit of work scheduled by the driver or PBO. |
| `protocol` | Persistent storage formats and transactions. |
| `session` | One compilation's lifetime and staging area. |
| `snapshot` | An immutable capture of inputs or state. |
| `metadata` | Immutable facts derived from TAST or published native signatures. |
| `state` | Mutable data owned by a pass or operation. |
| `host` | JS, Go or Rust environment executing the compiler. |
| `runtime` | The Go runtime embedded in the generated application. |

`Driver.Build` builds backend modules; `Driver.compile` orchestrates the whole
compilation. Emission returns a `GoFile`; `Driver.Output` writes it and the
cache session publishes it. A worker signature describes generated code, while
a job describes scheduling. Keep these distinctions in names and prose.

## Modules and files

- A module declaration matches its path under `src/` without `.purs`.
  `src/Gopurs/Host/Snapshot.purs` declares `Gopurs.Host.Snapshot`.
- Names describe a responsibility and use `PascalCase` path segments.
  Foreign `.go`, `.js` and `.rs` companions share the PureScript basename.
- Use existing module names when discussing the code: `Driver.Prepare`,
  `Monomorphization`, `CodeGen`, `FfiBridge` and `BuildCache` are current names.
  The conceptual phases in the module map do not imply different source paths.
- New names should identify their domain instead of using `Utils`, `Common`,
  `Helpers`, `Manager` or `Misc`. Prefer full words when practical.
- `Analysis` describes read-only facts; `Rewrite` or `Pass` describes a
  transformation; `Check` describes validation.
- `State` is mutable, `Metadata` is immutable and `Config` comes from options.
  `Snapshot`, `Session` and `Protocol` retain the meanings in the glossary.

## Identifiers

- Functions use verbs such as `prepare`, `specialize`, `rewrite`, `emit` and
  `publish`. Injectable variants use `With...`; FFI implementations use `...Impl`.
- Booleans use `is*`, `has*` or `should*`. Prefer a named mode over a new
  positional boolean in a public API.
- Proofs name the established property, such as `covered`, `complete`,
  `permitsLiteralEdit` or `sameTast`.
- Prefer `FFI`, `ADT`, `AST`, `JSON`, `Go`, `JS` and `Rust` in new identifiers,
  while preserving existing public and source names.
- Generated Go identifiers are part of the output contract. A documentation
  cleanup or internal module move must not rename them incidentally.

## Documentation

Documentation is written in English and describes current behavior. Update
source paths, commands, anchors and the module map when their implementation
changes. Keep dated experiments and completed development plans in Git history
rather than presenting them as current contracts. Link new guides from the
[documentation index](README.md) and run `npm run check:docs`.
