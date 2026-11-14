# Quality-gate status — 0.1.0

This page retains the qualification and public consumer evidence for the
`0.1.0-rc.14` prerelease. It does not qualify later source changes or a final
`0.1.0` release. The procedure is in [release.md](release.md).

## RC14 qualification and artifact review — 5 November 2026

The exact signed tag `v0.1.0-rc.14` is qualified at commit
`efcbf91bdd120ec1a33e915abd769c1d853b366e`, tree
`96a7d7e8e216e96fc099a1dd379d70772350f46a`, annotated tag object
`b7c80719ffed431f457cdbf4ba742070fae1d906`.
[Run 42001425870](https://github.com/0x000000000000000000001/gopurs/actions/runs/42001425870)
completed successfully in attempt 2.

| Platform | Gates | Fixtures Release / Dev | Libraries Release / Dev |
| --- | --- | --- | --- |
| Linux amd64 | 371 | 401 / 401 | 51 / 51 |
| Linux arm64 | 371 | 401 / 401 | 51 / 51 |
| macOS arm64 | 371 | 401 / 401 | 51 / 51 |

All three builds, twelve fixture shards, three signed finalizations and the
scope job are original successful attempt-1 producers. The only retried job
prepared the GitHub draft. Its first execution created the draft but could not
find it immediately afterwards; the precise remote visibility cause is not
established. The explicitly authorized retry reused the same draft and original
artifacts. No build, test or signature operation was repeated.

The full audit retained 25 original CI artifacts, 306 library JSON reports and
306 nonempty library logs. Repository frontend-cache imports were skipped on
all three builders; normal Nix substitutions remained available. Clean source
installation, first application build and the Nix application checks passed on
all three platforms. This does not replace installation from the public npm
package after publication.

The draft review verified all 45 assets and passed 16 cryptographic checks,
including npm provenance. The same npm archive bytes were produced on all
three platforms. Its SHA-256 is
`fba16760fb889d822b64c12bbeea9cfcea3d03969a0bb68c8adaf29d66eea845`.
Licenses, notices and the allowed package surface were reviewed. A targeted
scan of the exact assets, all 59 npm archive members and seven decoded signed
payloads found no configured secret or private-path patterns; this supplements
source and file-list review, rather than proving absence of arbitrary unknown
secrets. The explicit verification procedure remains in
[the release guide](release.md).

Retained audit identities:

- Full CI audit: `edf245c9b8122ceb44e2743cc3e31d6f4d9cdccd8223e875e1b6b7216e7e395d`.
- Draft verification: `620d999e7255ff347e56d400b429dbe16776e9adc2e26692a9467f1bde1dd67e`.

The signed archive still contains obsolete active RC2, RC9 and RC10 installation
examples. This missed the pre-tag documentation requirement in the release
procedure. The post-tag documentation cleanup is a separate source change.
The operator explicitly accepted this deviation and authorized publication
with an erratum giving the RC14 commands. That erratum has been applied and
verified on the now-public GitHub prerelease, with all 45 asset identities
unchanged. The signed archive and tag remain unchanged. The corrected source
documentation is a separate commit and is linked from the public erratum.

### Public publication and consumer checks

The exact archive is published as `gopurs@0.1.0-rc.14`. Its downloaded bytes,
npm integrity and registry provenance match the reviewed candidate.
[Publication run 42085073187](https://github.com/0x000000000000000000001/gopurs/actions/runs/42085073187)
succeeded, including JavaScript and native public-archive smoke tests. The
[GitHub prerelease](https://github.com/0x000000000000000000001/gopurs/releases/tag/v0.1.0-rc.14)
contains the same 45 reviewed assets and the documentation erratum.

On Linux arm64, a fresh public npm installation completed the documented Hello
with new npm, Spago and Go caches: PureScript compilation, Go generation, Go
build and exact output `Hello from gopurs`. The prepared prerequisites were
Node.js 24.8.0, Go 1.27.0, Spago 1.0.3 and the pinned TAST frontend; this test
does not claim to install those prerequisites on an empty operating system.
The unchanged YAML, PureScript example and console pin came from the packaged
guide; the RC14 installation command follows the public erratum.

The retained b8x source snapshot (3,242 files, 2,707 TAST modules) also passed
all 17 compilation stages and linked 12 executables with the public RC14
artifacts. No b8x executable or service was run; this establishes compilation
and linking only.

### Actual public rollback and restoration

The authorized exercise moved the real npm `next` channel from RC14 to RC7,
then restored RC14. Each phase installed `gopurs@next` in a separate empty
project with new npm, Spago and Go caches, verified the selected version and
lockfile integrity against its verified public archive, and produced exactly
`Hello from gopurs`. Complete dist-tag snapshots matched before and after
each installation and after each Hello execution.

| Phase | `next` | `latest` | Fresh channel installation and Hello |
| --- | --- | --- | --- |
| Fallback | `0.1.0-rc.7` | `0.1.0-rc.7` | Passed |
| Restored | `0.1.0-rc.14` | `0.1.0-rc.7` | Passed |

RC7 is validated as a fallback for this unchanged Hello consumer only. It does
not contain later library fixes and is not a blanket fallback for applications
that depend on them. Git tags and release assets were unchanged; no version
was deprecated or unpublished. Channel guards checked for concurrent changes
before each mutation; they are observations, not an atomic registry transaction.

The retained evidence includes public acquisition and provenance receipts,
original publication workflow artifacts, consumer reports and logs, both
channel-transition receipts, fresh lockfiles and source/toolchain hashes. No
stable release or new performance result is claimed.

## Earlier candidates

The [full qualification history](https://github.com/0x000000000000000000001/gopurs/blob/7589ec34aed08ec3b7dfc52014a4253b007b2fdc/docs/quality-0.1.0.md)
records earlier candidates, failed attempts and targeted library checks. Those
dated records remain in Git rather than in the current release guide.
