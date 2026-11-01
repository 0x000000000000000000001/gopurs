# Tests and snapshots

This directory contains the compiler fixtures, their golden Go output and the
host-level transition tests.

## Layout

| Path | Contents |
| --- | --- |
| `passing/*.purs` | 401 selected PureScript fixtures. Numeric names are upstream issue numbers and are kept as-is. |
| `passing/<Name>/` | multi-module fixtures (`Lib.purs`, `Main.purs`, …). |
| `passing-snapshots/` | Release golden output, one `.go` file per fixture. |
| `passing-snapshots-dev/` | Dev-profile golden output. |
| `passing-package-snapshots/` | package-level outputs and plans (`.txt`). |
| `host/` | cache and preparation transition tests (`CacheTransitions`, `PreparationTransitions`, host number checks). |

The fixture selection, exclusions and runner behaviour are documented in
[../docs/testing.md](../docs/testing.md). The `b8x` corpus used for the larger
campaigns is an external checkout; it is not part of this repository.

## Commands

Run from the repository root, with Go, Node, Spago and the TAST `purs` on
`PATH`:

```bash
./bin/test --list                 # 401 fixtures, no build
./bin/test ClassRefSyntax         # one fixture
./bin/test                        # all fixtures, Release
./bin/test -d                     # all fixtures, Dev profile
./bin/test --keep-workspace       # keep successful workspaces for inspection
UPDATE_SNAPSHOTS=1 ./bin/test ... # update golden files for the selected fixtures
```

`bin/test` uses `bin/gopurs` (native by default); `GOPURS_JS=1` selects the
JavaScript bundle and `GOPURS_RUST=1` the Rust host. A compiler change needs a
rebuild first:

```bash
npm run build:native -- --keep-workspace
```

The full suite is expensive. For a change, select the narrowest fixtures and
tests listed in [../docs/testing.md](../docs/testing.md), then run the full
suite before merge.

## Snapshots

A snapshot is the formatted Go output of one fixture. Updating it is a review
event, not a reflex:

1. Run the fixture without `UPDATE_SNAPSHOTS=1`; confirm the failure is the
   intended change.
2. Re-run with `UPDATE_SNAPSHOTS=1` for that fixture only.
3. Inspect the diff in `passing-snapshots/` **and** `passing-snapshots-dev/`;
   both profiles usually change together.
4. Re-run the fixture normally, then the full suite.

Never update snapshots to make a structural cleanup pass. A cleanup that
changes generated Go is a behaviour change and must be reviewed as such.

## Adding a fixture

- Put it in `tests/passing/` (or a directory for multiple modules).
- Name it after the feature: `NativeRecordReturns.purs`, not `issue-1234.purs`.
  Keep upstream numeric names when importing an existing test.
- Add the smallest program that fails before the fix and passes after it.
- Generate both snapshots, then run the fixture in both profiles.
- If the fixture is excluded, add the reason to the exclusion list and to
  [../docs/testing.md](../docs/testing.md).

## Conventions

- Do not rename existing snapshots for cosmetic reasons; the runner and
  package snapshots reference them by fixture name.
- `*-dev` denotes the Dev profile, not a separate feature. Do not add new
  `*-dev` directories.
- Keep fixtures self-contained; do not depend on sibling checkouts.
- A fixture that needs FFI must include the Go side in the fixture and keep it
  compatible with the three hosts.
