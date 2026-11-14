# Library source review

[`bootstrap-sources.json`](../bootstrap-sources.json) selects the typed
compiler, optimizer and Go library ports by repository URL and exact commit.
Review that manifest together with the dependency locks and
[third-party notices](../THIRD_PARTY_NOTICES.md) when preparing a release.

## Distribution boundaries

| Surface | Review |
| --- | --- |
| Compiler npm archive and native executables | The actual bundled dependency closures, package allowlist, licenses, notices and installed-artifact checks |
| Library sources used by bootstrap or applications | The exact Git pins, dependency resolution, source notices and tests with the selected compiler |
| A separately published library | Its own package name, destination, version, dependency bounds, files and notices |

The compiler npm archive does not distribute the library repositories as
standalone packages. A version inherited from an upstream package does not
establish a published version of its Go port. The current compiler release
identity and qualification are documented in [publication scope](publication-scope.md)
and the [quality record](quality-0.1.0.md).

## Source configuration for 0.1.1

The [0.1.1 manifest](https://github.com/0x000000000000000000001/gopurs/blob/v0.1.1/bootstrap-sources.json)
retains RC20's exact typed frontend, support inputs and 51 library pins. In
particular, `gopurs-foreign` at
`8c326898e821a9151af35b2a647fc9b2354ba92e` supplies native Go reads that preserve
shared and cyclic child values. The release's compiler and runtime come from
RC20; the recovery CLI, packaging and fixture qualification are tracked
separately in the [quality record](quality-0.1.0.md).

Use the checkout's manifest for current pins and the
[latest release](https://github.com/0x000000000000000000001/gopurs/releases/latest)
for published installation instructions. The
[release source audit](https://github.com/0x000000000000000000001/gopurs/blob/e691356/docs/library-source-audit.md)
retains the dated RC18–RC20 pin changes and their original validation scope.

## Reviewing a pin change

1. Compare the old and new commits, including source, tests, configuration and
   notices. Preserve upstream attribution and identify the Go adaptation.
2. Check the library's resolved dependencies. Local Spago paths do not pin the
   sibling repositories; the bootstrap manifest supplies those identities.
3. Run the relevant library and consumer checks using the intended pins and
   compiler. Follow [development.md](development.md#libraries) and
   [testing.md](testing.md). Record the commits and scope with the results.
4. Review both JavaScript and native compiler dependency closures. A library
   absent from one closure may still be included in another. Update
   `THIRD_PARTY_NOTICES.md` for the components actually distributed.
5. Apply the [release procedure](release.md) to the resulting candidate.
   Earlier library checks do not qualify a different compiler or package.

## Earlier audit

The [October–November 2026 audit](https://github.com/0x000000000000000000001/gopurs/blob/7589ec34aed08ec3b7dfc52014a4253b007b2fdc/docs/library-source-audit.md)
records the original 51-library inventory, source hashes, notice corrections
and targeted checks. Its counts and results belong to those exact revisions;
use the current manifest and distribution contents for a new review.
