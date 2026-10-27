# Migration and compatibility

## Current baseline

`0.1.0` is the first public release. There is no earlier public version to
migrate from. The compatibility commitments and the `0.x` policy are in
[publication-scope.md](publication-scope.md).

## Version policy

- Patch releases (`0.1.x`) do not intentionally break documented CLI behavior,
  exit status, error prefix or documented entry points.
- Minor releases may change generated layout, internal caches, FFI details,
  experimental flags and defaults.
- `1.0.0` is reserved for a stable public contract across minor releases.
- The version source is `package.json`; the changelog records user-visible
  changes.

## Migrating from a local edge build

A local checkout may contain generated output and caches from a different
compiler revision. When moving to a release tag:

1. Install the pinned typed compiler fork revision for that release.
2. Regenerate the typed output with `spago build`.
3. Rebuild the application from a clean `output/` when the compiler reports a
   cache or TAST mismatch.
4. Run `gopurs --main <Module>`.
5. Build the generated application with the supported Go toolchain.

If a cache error persists after a clean frontend build, remove
`output/.gopurs-cache` and rerun. The cache is not a compatibility surface.

## Generated output

Treat generated Go as a build artifact:

- Do not edit it by hand; regenerate it.
- Do not rely on exact filenames, identifiers or package partitioning beyond
  the documented entry points.
- Rebuild it after changing the compiler, `GOPURS_GO_PACKAGES`,
  `GOPURS_GO_PACKAGE_OVERLAY` or `GOPURS_BUILD_MANIFEST`.
- Keep `go.mod` and `go.sum` under the normal Go workflow; gopurs preserves
  external `go.sum` content.

## Library ports

The `gopurs-*` libraries are separate packages. A compiler release does not
automatically make every stock PureScript library work. Use the library
revision associated with the release notes or the repository branch documented
for that release. A library without Go FFI can compile but fail at runtime.

## Toolchain pins

Record these with any reproduced build:

- gopurs version and source commit;
- typed `purs` fork revision;
- Node.js and npm versions for the JavaScript host;
- Go version used for generated applications;
- Spago version and package set;
- the `gopurs-*` library revisions;
- cache/package-layout settings.

The [bootstrap chain](bootstrap.md) records the compiler-side pins.

## Removing experimental settings

`0.1.0` accepts experimental daemon, resident, Rust and profiling settings.
They are not required for a normal build. When upgrading:

- unset `GOPURS_DAEMON`, `GOPURS_RESIDENT` and related variables;
- prefer the default `release` profile;
- rebuild the native host with the release tag rather than reusing an older
  `bin/gopurs-native`;
- remove stale daemon directories such as `.gopurs-daemon` if they are not in
  use.
