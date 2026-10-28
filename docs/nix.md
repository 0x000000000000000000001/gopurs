# Nix installation and development

The flake exposes packages, executable apps, a development shell and checks for
`x86_64-linux`, `aarch64-linux` and `aarch64-darwin`. Its inputs are fixed by
`flake.lock`; it does not read a sibling compiler checkout or an existing
gopurs binary.

## Install or run the backend

From a checkout at the release revision:

```sh
nix build .#gopurs
```

From an application directory containing typed `output/*/corefn.json`, run the
flake by its checkout path:

```sh
nix run /path/to/gopurs -- --main Main
```

The installed `result/bin/gopurs` is the JavaScript compiler host. Its Node
runtime, FFI parser and Go cache/package helpers are included. The helpers are
compiled during the derivation: running gopurs never needs to write into its
read-only installation in `/nix/store`.

The backend consumes typed `output/*/corefn.json` from the fork below. Building
the generated application still requires Go; the development shell supplies it.

## Typed frontend and development shell

```sh
nix build .#purs-tast
nix run .#purs -- --version
nix develop
```

The shell provides the pinned TAST-capable `purs`, Go `1.27.0`, Node `24.8.0`
and application Spago `1.0.3`. It exports `GOPURS_PURS` to the Nix-built frontend
and disables automatic Go toolchain downloads with `GOTOOLCHAIN=local`.

The TAST frontend is built from the exact source revision in the `purs-tast`
flake input using GHC 9.8.4. Nix locks its Haskell dependency set independently
of the fork's Stack development plan. Its generated Cabal dependency description
is checked in as `nix/purs-tast-package.nix`, so evaluating another platform
does not need to execute a generator for that platform. As with nixpkgs' upstream PureScript
package, the package derivation relaxes Cabal version bounds to use this locked
set. The end-to-end application check verifies that the result emits TAST with
`typeTable`, `dataDecls` and `classDecls` and that gopurs accepts it.

The backend's own JavaScript build uses upstream PureScript `0.15.16` and Spago
`0.93.45`. These tools are deliberately separate from the application frontend.
Inside the shell, `npm ci` selects the build tools from `package-lock.json`.
The Nix derivation uses the same versions directly and fetches PureScript
dependencies from `spago.lock` before its offline build.

Node and Go are fetched from official release archives with per-platform
SHA-256 values in `nix/toolchains.nix`. They do not float with the default
versions in nixpkgs. The shell also includes the compiler and linker needed
for native builds.

## Checks

```sh
nix flake check -L
```

The checks build the installed backend, verify its error outside a project,
then compile a small PureScript program with the pinned TAST frontend. That
program uses a Go FFI function and runs in both monolithic and Go-package
layouts. This exercises the installed resources and immutable helper caches,
in addition to the frontend schema and generated executable.

The complete compiler fixture suite and the release benchmark gates are
separate checks described in [testing.md](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/testing.md) and
[benchmark-protocol.md](benchmark-protocol.md).

## CI frontend cache

Quick PR and branch builds can restore an Actions cache containing only the
runtime closure of the pinned `purs-tast` tool. A quick push or manual run on
`main` populates it before checking or building gopurs. Manual quick runs can
therefore seed the cache even when a merged PR's qualification was reused.
The exact key includes the runner OS/architecture and the flake, Nix and
bootstrap pins; there is no fallback prefix or cross-platform restoration.

The workflow exports that closure with `nix copy` into a dedicated directory
under `RUNNER_TEMP`. It never archives the complete `/nix` directory, installer
receipts, GHC build dependencies or gopurs outputs. Importing this repository's
Actions cache uses `--no-check-sigs`; its trust comes from the repository's
cache scope rather than a public Nix signing key. PRs only restore; they do not
publish caches for subsequent builds.

Cache failures fall back to the normal pinned build. Set the repository variable
`GOPURS_NIX_TOOLS_CACHE=false` to disable the optimization. Tags and manual full
campaigns always bypass it and bootstrap the current gopurs source afresh.
No warm-run CI speedup is claimed until a producer and consumer have been
measured. The compiler's functional test caches remain private to each run.

## Native bootstrap

The native Go compiler is built from source with the same pinned tools:

```sh
nix develop
npm ci
./bin/setup --locked --all
npm run build:native
```

`bootstrap-sources.json` records the library revisions for this workflow.
`bin/setup --locked --all` installs them; it requires network access for the
initial clones. The native bootstrap is not itself a pure Nix derivation.
The npm and native release archives remain separate from the Nix package.

## Updating pins

Update `flake.lock` explicitly and review the resulting versions. When changing
the TAST input, also update the compiler revision in `bootstrap-sources.json`
and the release notes, and regenerate `nix/purs-tast-package.nix` with
`cabal2nix` from that pinned source (keeping `src` as its function argument).
When changing Node or Go, update all three platform
archive hashes, the npm/build toolchain declarations and the release contract,
then rerun the Nix checks on every supported system.

The legacy `nix-shell` entry point uses the same locked flake through
`shell.nix`; it does not define a separate environment.
