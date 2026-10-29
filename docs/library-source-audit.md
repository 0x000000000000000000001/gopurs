# Audit of the library source pins

Audit date: 26 October 2026. This records the 51 repositories selected by
[`bootstrap-sources.json`](../bootstrap-sources.json), whose SHA-256 was
`e3db111252325962279ec3de313d843110c2525caa408526a9e29ceed30855ff`.
All 51 inspected checkouts matched their exact commits and had clean working
trees before and after the 26 October audit. That audit modified no library
and reran no library build. The notice corrections published on 27 October
and their new source pins are recorded separately below.

## What is being distributed

| Surface | Identity and release requirement |
| --- | --- |
| Compiler npm package | `gopurs@0.1.0-rc.5`, then the final `0.1.0`; its own version, dependency lockfile, file allowlist and installed-artifact tests apply. |
| Library sources used by bootstrap/applications | Git URL and exact commit from the bootstrap manifest; validate this combination with the selected compiler and package set. The compiler npm archive does not contain these repositories. |
| Independent future npm/Spago library publication | Review that package's name, destination, version, dependency bounds, files and notices before publishing it. This is a separate operation. |

These independent libraries do not need identical version numbers. Their
current release identity is the compatible set of source commits. A version
copied from an upstream package is not a version assigned to the Go port.

## Manifest and file checks

- 49 repositories have `package.json`: 46 say `private: true`; `spec` and
  `yoga-json` have neither npm name nor version; `uuid` retains the upstream
  name `purescript-uuid` and version `1.0.0`. `js-bigints` and
  `node-event-emitter` have no npm manifest. None declares an npm `files`
  allowlist. These trees are not 51 prepared npm packages named `gopurs-*`.
- All 51 have Spago configurations. 50 select registry package set `77.7.0`;
  `assert` selects `73.3.0`. The compiler bootstrap constructs its own workspace
  using `77.10.1`, as documented in [bootstrap.md](bootstrap.md).
- 42 checked-in Spago lockfiles contain 882 registry entries with exact
  versions and integrity hashes, and 673 local entries with existing paths.
  There are no Git dependencies in those resolved entries. Counts include
  repeated dependencies across workspaces. All 42 package-set addresses agree
  with their configurations. The remaining nine configurations still select
  an exact registry package set; their resolved closure is not recorded in a
  checked-in lockfile.
- `node-streams`, `spec` and `yoga-json` declare 73 dependency ranges. All 33
  resolved registry versions satisfy them. The other 40 resolve to local
  sources: the corresponding upstream package-set versions satisfy the ranges,
  but compatibility of the Go commits must be established by the library and
  application tests, not by that version comparison.
- Those same three configurations retain upstream `publish` destinations and
  versions: `purescript-node/purescript-node-streams` (`9.0.1`),
  `purescript-spec/purescript-spec` (`8.1.2`) and
  `rowtype-yoga/purescript-yoga-json` (`5.2.1`). They are not publication
  instructions for the Go forks.
- The tracked source trees contain 1,355 files, approximately 6.3 MB. No tracked
  `output`, `.spago`, `node_modules`, `.env` or `.DS_Store` path was found, and
  no file exceeds 1 MiB. Git source distributions include tests, scripts and
  examples; the compiler's npm allowlist is a separate boundary.

`js-uri` declares two unavailable local overrides, `simple-json` and `test`.
Neither is a core/test dependency or part of its resolved 35-package closure
in the frozen 25 October corpus. They are unused configuration entries, not
missing inputs to the current library build.

## Confirmed notice and provenance gaps

| Port | Evidence | Remaining work for its source distribution |
| --- | --- | --- |
| `effect` | All seven inherited `.purs`/`.js` files exactly match upstream `v4.0.0`, commit `a192ddb923027d426d6ea3d8deb030c9aa7c7dda`. Its [BSD-3-Clause license](https://github.com/purescript/purescript-effect/blob/a192ddb923027d426d6ea3d8deb030c9aa7c7dda/LICENSE) names Copyright 2018 PureScript. The originally audited pin had the identifier in `bower.json`, but no license text or replacement source headers. | Corrected by the published and pinned notice commit below, which also explicitly licenses the Go adaptation. |
| `js-bigints` | Both inherited files exactly match upstream `v2.2.1`, commit `36a7d8ac75a7230043ae511f3145f9ed130954a9`. Its [license file](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/LICENSE) is MIT, Copyright (c) 2022 Jan Schulte. The originally audited pin had no license text or source headers. | The published and pinned notice commit below restores the notice and explicitly licenses the Go adaptation. The inherited metadata/example questions remain documented. |
| `uuid` | Its npm manifest declares ISC, while the checked-in `LICENSE` and Bower manifest declare MIT. The npm name/repository still identify the upstream project. | Resolve the metadata before an autonomous package publication; do not infer an intended npm release from this manifest. |

For `js-bigints`, upstream [package.json](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/package.json)
says ISC while Bower/Dhall and the license file say MIT; the registry records
`MIT AND ISC`. Its [JavaScript FFI](https://github.com/purescript-contrib/purescript-js-bigints/blob/36a7d8ac75a7230043ae511f3145f9ed130954a9/src/JS/BigInt.js)
also identifies a copied [Stack Overflow example](https://stackoverflow.com/questions/55646698/base-36-to-bigint).
That comment is an additional provenance point to clarify, not evidence for
assigning a new license to the whole port. The Go implementation uses
`math/big.Int.SetString` instead of that JavaScript algorithm.

The official license files were fetched at the commits above and matched the
registry copies: SHA-256 `d1ceeb89b7cce8e527167e969128601a80628f229f29db8a1f0111a9dd35375f`
for `effect`, and `4a56dd0010e318fc1ebb942ee91a858ddb2a4de76f8085ffc5d84c58d1eadd26`
for `js-bigints`. The source pins audited on 26 October do not contain these
license files.

## Published notice corrections — 27 October 2026

The maintainer selected the same upstream license terms for each Go
adaptation: BSD-3-Clause for `effect`, and MIT for `js-bigints`. After explicit
maintainer approval, the following notice-only commits were pushed as fast
forwards to `gopurs-effect/master` and `gopurs-js-bigints/main` respectively.
The bootstrap manifest now selects these public revisions:

| Library | Source pin used for the original audit/tests | Published source pin |
| --- | --- | --- |
| `effect` | `9f72d40a510e37e53b49699e91e372345bfa386a` | `d4312e08a73769f2b41b89866717393192fe103e` |
| `js-bigints` | `e431af881ec743ff79242e09fe608a09b2703f5c` | `11222d0fea160322fbf08682109febb78284b033` |

Both commits add only `LICENSE` and `NOTICE.md`. The license files exactly
match the upstream hashes above and preserve the original copyright notices.
The `effect` notice applies BSD-3-Clause to `src/Effect.go`,
`src/Effect/Uncurried.go` and `src/Effect/Unsafe.go`; the `js-bigints` notice
applies MIT to `src/JS/BigInt.go`. The SPDX identifiers and file scopes are
recorded in the notices; no header was added to the compiled source files.
Neither notice attributes the Go adaptation to the upstream copyright holder.

The seven/two inherited source files were rechecked byte for byte against the
exact upstream revisions and registry copies before restoration. Every tracked
`.purs`, `.js` and `.go` file, including tests, was also checked byte for byte
against the initial source pins: 11 files in `effect` and four in `js-bigints`,
all identical. Both committed worktrees are clean, and their diffs pass
`git diff --check`. No library build was run for these notice-only changes.
Review patches, upstream copies and hash verification records are retained in
`/private/tmp/gopurs-library-notice-fixes-20261026/`.

The explicit MIT choice applies to the Go adaptation. It does not resolve the
MIT/ISC metadata discrepancy or the Stack Overflow attribution/licensing
questions in inherited JavaScript. Those facts remain documented in the
`js-bigints` notice; the `uuid` metadata issue also remains open.

The isolated library checkouts used by the test campaigns remain clean at
their original source pins. The notice-only differences above preserve every
compiled source byte; the recorded tests and frozen benchmark inputs retain
their original identities instead of being relabelled with the new commits.

## Compiler release boundary

The native snapshot corresponding to binary SHA-256
`eb97922721f15c3e742ff14335a8294b6005ca321ae2e9a57bb8847776e46298`
contains 71 resolved packages: 38 local and 33 registry packages, producing 522
typed modules. Its Spago lockfile SHA-256 is
`022cab0e951e68020a3af88afb33b6773861d5cbee17e1ba80bfee741b3e268c`.

`effect` is in that native closure; its upstream copyright and full BSD text
are already in [THIRD_PARTY_NOTICES.md](../THIRD_PARTY_NOTICES.md).
`js-bigints` is absent from both the resolved compiler lockfile and native
output, so its source notice gap is not evidence of a missing component in
the compiler artifacts.

The native closure adds `assert` through `nullable` compared with the JavaScript
compiler lockfile. `Test_Assert.go` and `Test_Assert_ffi.go` are generated. The
port's license is BSD-3-Clause, Copyright 2018 PureScript. The missing inventory
entry was added to `THIRD_PARTY_NOTICES.md` during this review, with its exact
source pin and license link; the common BSD license text is already reproduced.
No other missing package notice was identified in the resolved compiler
closures. This inventory check does not claim every generated function
survives Go linker elimination.

The compiler release must account for its actual JavaScript/native closures
and ship their notices. The missing source notices above have been restored;
the inherited metadata and JavaScript provenance questions remain open for
the public library source family.
Do not describe the 51 ports as separately published npm packages or as having
completed autonomous publication checks.
