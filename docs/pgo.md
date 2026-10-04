# Profile-guided optimization for the Rust host

`npm run build:rust` produces a profile-guided `bin/gopurs-rust` by default.
The profile is generated during every build from the compiler's own modules:
no archived profile, absolute path, or benchmark workload is needed.

Set `GOPURS_RUST_PGO=0` to install the unprofiled O3 ThinLTO bootstrap
instead. `GOPURS_KEEP_WORKSPACE=1` (or `--keep-workspace`) retains the build
directory, including logs, raw profiles, and `pgo-profile.json`.

## Pipeline

1. **Bootstrap.** The generated Rust source is linked as usual (O3, ThinLTO,
   `profile.release.debug=false`, bundled `ld64.lld` on macOS) and the
   resulting binary is kept as the generation oracle.
2. **Training workspace.** A second, separate workspace is created from
   `nativeWorkspaceConfig(root, { runtime: "go" })` and built by the typed
   `purs` frontend. It compiles this repository's `src/` against the Go library
   checkouts; the Rust workspace and the live `src/` tree are untouched.
3. **Freeze.** Every `output/<Module>/corefn.json` is copied into
   `pgo/training` with its sibling `.purs`, `.go`, and `.js` sources and a
   relative `modulePath` (`sources/<Module>/<last-segment>.purs`), so the
   workload is self-contained. `Test.Main` is excluded and the presence of
   `Main` is required.
4. **Oracle.** The unprofiled bootstrap binary compiles the frozen workload and
   its generated Go sources (`*.go`, `go.mod`) become the comparison baseline.
5. **Instrumented build.** The exact same Rust source is rebuilt with
   `CARGO_ENCODED_RUSTFLAGS="-C\x1fprofile-generate=<raw>"` into a dedicated
   target directory. Profile paths travel through the environment, never
   through a shell command line.
6. **Training passes.** Three passes run with `GOPURS_JOBS=8`,
   `GOPURS_PREPARE_JOBS=8`, `GOPURS_PBO_JOBS=8`, `GOPURS_EMIT_JOBS=8`,
   `GOPURS_PIPELINE=1`; the last pass uses one prepare/PBO/emit worker. Every
   pass must regenerate byte-identical `*.go` and `go.mod` files, and the full
   frozen training manifest (including `Main/corefn.json`) is asserted before
   and after each pass. Cleanup removes only generated `*.go`, `go.mod` and
    `go.sum` files: it never removes the `main` directory, because `output/main`
    aliases `output/Main` on case-insensitive APFS. The generated PBO `.cache`
    and `.purmeta` directories are also reset between processes.
7. **Merge.** The `llvm-profdata` from the active `rustc` sysroot merges every
   `.profraw` into `training.profdata`.
8. **Profile use.** The same ThinLTO/linker configuration is rebuilt with
   `-C\x1fprofile-use=<merged>`, and that binary goes through the regular
   smoke test and installation.

## Metadata

The build writes `<workspace>/pgo-profile.json` with the qualification data
needed to compare a local build against the isolated experiment:

- `status`, `started_at`, `finished_at`, `stage`, `failure` on errors
- `source.manifest` / `source.sha256`: generated Rust sources and their digest
- `training.modules`: module names, original relative source paths, frozen
  sources, and CoreFn hashes; `training.excluded_modules` /
  `excluded_present`; `training.required_modules`; `training.manifest` /
  `training.sha256`: every frozen file; `training.policy`: library-overlap and
  held-out policy
- `training.oracle`: the generated Go manifest the passes must match
- `profile.tool`, `profile.raw`, `profile.merged.sha256`
- `passes`: per-round worker settings and generated digests
- `binary.sha256`, `binary.bootstrap_sha256`, `binary.instrumented_sha256`
- `builds`: bootstrap and instrumented binaries with their digests

The final lines printed by the build (`[pgo] source sha256: …`,
`[pgo] training sha256: …`, `[pgo] profile sha256: …`,
`[pgo] binary sha256: …`, `[pgo] metadata: …`) carry the same values for
automated qualification.

## Interruption and failures

All commands run through the existing `CommandRunner`, so `SIGINT`/`SIGTERM`
reach the active process group, the previous `bin/gopurs-rust` stays in place,
and the workspace is retained. A failed PGO stage writes `status: "failed"`
with the failing stage and stack before the build reports the error.

`node --test tools/pgo.test.mjs` covers the cleanup/Main preservation,
encoded Rust flags with spaces, and module inclusion/exclusion helpers without
running any compiler.
