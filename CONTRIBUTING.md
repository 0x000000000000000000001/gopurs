# Contributing

Thanks for helping with gopurs. The project is experimental; small, tested
changes with clear evidence are preferred over broad refactors.

By participating, you agree to follow the [Code of Conduct](CODE_OF_CONDUCT.md).

## Development setup

Requirements:

- Node.js `24.8.0` (see `.nvmrc`);
- Go `1.27.0`;
- the typed PureScript compiler fork on `PATH`;
- Spago `1.0.x` for application and fixture builds;
- Git and Bash.

Install the compiler dependencies:

```bash
npm ci
```

The source build downloads the pinned optimizer and `st`/`unsafe-coerce`
overrides through `spago.lock`. The typed compiler fork and the optional
`gopurs-*` library checkouts are still separate:

```bash
./bin/setup --core
```

`bin/setup` preserves existing checkouts. It does not install the typed
compiler.

## Checks

Start with the checks that do not need a full library corpus:

```bash
npm run test:ffi
npm run test:launcher
npm run test:package
```

With the typed compiler and library checkouts available, run the broader suites
listed in `package.json` (`test:cli`, `test:profiles`, `test:cache`,
`test:runner`, and the Go tests under `tools/build-cache` and `tools/c1c`).

For compiler changes, compare generated Go on a frozen TAST corpus and run the
relevant fixture suite. A change that alters generated output must update the
corresponding snapshots and explain the intended difference.

## Vocabulary and module map

- [docs/naming.md](docs/naming.md) is the normative glossary: one concept, one
  word. Phases are verbs, data are nouns, proofs are adjectives. Read it before
  naming a module, a function or a field.
- [docs/module-map.md](docs/module-map.md) is the source of truth for module
  responsibilities, target phases and allowed dependencies. Add or update the
  matching row in the same commit as a module change.
- The target phases are **Driver**, **Prepare**, **Rewrite**, **Emit**,
  **Cache**, **Host** and **Runtime**. The current flat names are temporary:
  `Driver.Prepare` prepares, the rewrite passes transform, `CodeGen` emits and
  `Driver.Output` writes. Do not introduce a new ambiguous name while the
  cleanup lots are in progress.
- A move is a separate pull request from a rename or a behaviour change. Move
  the `.purs` file and its `.go`/`.js`/`.rs` companions together, update every
  import and path, and never update a snapshot to justify a move.
- Generated Go identifiers are a public contract. Do not rename them during
  cleanup.

## Code and documentation

- Follow the style of the surrounding module and keep responsibilities in the
  phase described by the module map.
- Keep public documentation in English. Historical campaign notes may use their
  original language. Use the vocabulary from `docs/naming.md` in normative
  documents and link to `docs/module-map.md` instead of duplicating the module
  list.
- Do not commit generated output, caches, benchmark workspaces, logs or
  platform binaries.
- Do not add a runtime dependency without updating the license inventory.
- Update `CHANGELOG.md` for user-visible behavior.
- Update `docs/publication-scope.md` when a public compatibility commitment
  changes.

## Benchmarks

Read the [benchmark protocol](docs/benchmark-protocol.md) before publishing a
number. Include the host, toolchain, corpus, cache state, worker settings,
sample count and raw data. Do not present a backend-only result as an
application runtime result.

## Pull requests

A pull request should include:

- the problem and the chosen scope;
- the commands and results used to validate it;
- generated-output or snapshot changes;
- compatibility or migration impact;
- documentation and changelog updates when relevant.

Security reports must follow [SECURITY.md](SECURITY.md), not the public issue
tracker.

## License

Contributions are accepted under the repository's MIT License. By submitting a
change, you confirm that you have the right to license it under those terms.
