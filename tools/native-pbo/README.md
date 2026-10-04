# Native-PBO differential qualification

Runs the eight native FFI differential contracts of the Rust-hosted gopurs
compiler against a **retained** `npm run build:rust` workspace. The candidate
native sources are taken from the local gopurs PBO checkout; the PureScript
oracles and the embedded FFI under test come from the generated crates.

| Case | Contract | Injected from gopurs PBO | Needs corpus |
| --- | --- | --- | --- |
| `memo` | `BoundedMemo.rs` inner-tree keys, FIFO, reentrancy | generated crate only | no |
| `directives` | `Directives.rs` ASCII fast path and exact fallback | generated crate only | no |
| `qualified` | `CoreFn.rs` borrowed `Qualified Ident` compare/eq vs the PS oracle, wrappers and generic `Ord` | `CoreFn.rs` | no |
| `source-usage` | `CoreFn/Usage.rs` validator vs `validateSourceUsageModulePS` | `CoreFn/Usage.rs` | yes |
| `tast` | `CoreFn/Json.rs` array loop + annotation decoder | `CoreFn/Json.rs` | yes |
| `tast-module` | `CoreFn/Json.rs` module decoder + usage validation | `CoreFn/Json.rs` | yes |
| `type-table` | `CoreFn/Json.rs` type-table resolution | `CoreFn/Json.rs` | yes |
| `maps` | Persistent native/generic map interoperability, insertWith merge order and Unicode keys | `NativeMaps.rs` | no |

The fixture programs (`test-native-*.rs`) and the TAST-module boundary table are
reused from `PURUST_DIR/tools`; the `qualified` fixture ships next to its case
script under `tools/native-pbo/fixtures/`. Before Cargo runs, the orchestrator
compares the full threaded FFI source with the code embedded in
`GENERATED_RUST` and records its SHA-256. A mismatched source or stale generated
compiler fails the check.

The local `maps-insert-with.rs` supplement checks 4,000 updates against both
the generated map implementation and a key/value model, including callback
counts, existing/incoming argument order and persistent versions.

## Prerequisites

- Retained Rust-hosted gopurs build. `build-rust.mjs` deletes the workspace
  unless keep is requested:

  ```sh
  GOPURS_KEEP_WORKSPACE=1 npm run build:rust
  # or: npm run build:rust -- --keep-workspace
  # prints: Workspace retained: <workspace>
  ```

  `GENERATED_RUST` is `<workspace>/rust`.

- Purust checkout for the reused fixtures and `threadedRust` transform
  (default `../../purust/purust`, override `PURUST_DIR`).

- The local gopurs PBO checkout (default
  `../../purescript-backend-optimizer-gopurs`, override `GOPURS_PBO_DIR`).

- The TAST corpus defaults to `<GENERATED_RUST>/../output` (the typed output of
  the retained workspace). Pass a packed `[{ name, contents }]` corpus file or a
  directory of `<Module>/corefn.json` trees explicitly to bound it. The
  `type-table` case requires the directory form, as in Purust.

## Commands

Full campaign (all eight cases, sequential, one shared target):

```sh
node tools/test-native-pbo.mjs <workspace>/rust <log-dir>
```

Explicit corpus and target directory:

```sh
node tools/test-native-pbo.mjs <workspace>/rust <log-dir> <frozen-output> \
  --target-dir <workspace>/target
```

Re-run selected cases in a new log directory:

```sh
node tools/test-native-pbo.mjs <workspace>/rust <log-dir> --only memo,directives
```

Run a single contract standalone (no aggregate log, stdout only):

```sh
node tools/native-pbo/memo.mjs <workspace>/rust
node tools/native-pbo/source-usage.mjs <workspace>/rust <frozen-output>
```

## Output and disk behavior

- `<log-dir>/summary.json`, `<log-dir>/environment.json`: campaign result and
  resolved paths/provenance.
- `<log-dir>/<case>.log`: progressive per-case log (written as the case runs,
  so interrupted campaigns keep partial evidence).
- `<log-dir>/<case>.console.log`: captured child output written after the case.
- `<log-dir>/workspaces/`: fixture workspaces; failures are retained there,
  successes are removed unless `--keep` (`PBO_NATIVE_KEEP=1`).
- One shared cargo target directory: default `<GENERATED_RUST>/../target`
  (the retained build's own target, reuse intended), override `--target-dir`.
- Every fixture build forces `profile.release {opt-level=3, debug=false,
  lto=false}` via `cargo --config`, independently of the fixture manifest.
- The orchestrator is sequential; builds/tests are serialized by the caller.

## Environment overrides

`GOPURS_PBO_DIR`, `PURUST_DIR`, `PBO_NATIVE_FIXTURES`,
`PBO_NATIVE_TARGET_DIR`, `PBO_NATIVE_TMPDIR`, `PBO_NATIVE_KEEP=1`.
The orchestrator sets the per-case values, so standalone case runs follow the
same defaults.

## Provenance check

The complete FFI source is checked against the generated compiler, including
the memo and directives cases that call their generated crates directly.
When a corpus is present the orchestrator reads `modulePath` from the corefn
modules and fails if the corpus names the Purust fork and no gopurs fork path.
Pass `TAST_CORPUS` for a corpus outside the retained workspace.

## Standalone text gate (`text.mjs`)

Final gate for the native CoreFn JSON text decoder (not part of the eight-case
`test-native-pbo.mjs` list, which keeps the older ABI and does not expect the
cursor path for every case). Run it standalone:

```sh
node tools/native-pbo/text.mjs --workspace PATH --corpus PATH \
  --expect-cases 4030 --expect-corpus 238 --log DIR
```

- Candidate FFI defaults to the live gopurs PBO
  `src/PureScript/Backend/Optimizer/CoreFn/Json/Text.rs` through the shared
  `pboSource` helper; `--text FILE` overrides it.
- Reference: **full PureScript oracle**. The runner copies the generated
  Json/Usage crates into `*Oracle` packages (Cargo paths rewritten absolute,
  `JsonOracle` depending on `UsageOracle`), rewrites the FFI entry points
  (`decodeAnnWithUsageImpl`/`decodeArrayImpl`/`decodeModuleImpl` to their PS
  fallbacks, `decodeTypeTableImpl` to `decodeTypeTablePS`,
  `validateSourceUsageModuleImpl` to its PS fallback) and drops the dead native
  helpers. Static checks assert unique anchors, exact PS-only bodies, zero
  native calls and no fallback loops.
- Cases: frozen corpus + deterministic contract mutations + boundary matrix
  (4030 with the gopurs-aff corpus). Errors compare the printed `Left` byte for
  byte; successes compare a full structural dump (types, annotations, usage,
  expressions, cold declarations); `valid-sharing` keeps the type-table alias
  check.
- Artifacts: `<log>/report.json`, the first-failure snapshot, and
  `<log>/oracle-report.json` plus the rewritten oracle sources and original/
  rewritten hashes.
- The gate verifies `PurustJsonCursor::materialize` is present in the workspace
  before building; JSON syntax errors keep the Argonaut parser as the byte
  source of truth, and the fixture workspace is retained on failure.
