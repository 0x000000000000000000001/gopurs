# Release procedure

The release consists of one npm archive and native executables for
`linux-amd64`, `linux-arm64`, and `darwin-arm64`. The archive supplies the
JavaScript compiler, launcher and runtime helpers; the matching native
executable can be installed alongside it. Rust is outside this release gate.

The workflows and scripts below implement the procedure. Their presence is not
evidence that a tag, platform, benchmark or public installation has passed.
Record actual results in the repository's
[quality record](https://github.com/0x000000000000000000001/gopurs/blob/main/docs/quality-0.1.0.md)
and the release notes. The remaining gates are tracked in the
[release checklist](https://github.com/0x000000000000000000001/gopurs/blob/main/todo.md).

## Build and validate a candidate

Set `package.json` and `package-lock.json` to the same release version. Use
`0.1.0-rc.1`, then `0.1.0-rc.2`, etc. for candidates; the final version is
`0.1.0`. Add the corresponding `## [VERSION]` section to `CHANGELOG.md`, with
the exact TAST revision from `bootstrap-sources.json`, compatibility limits,
platform results and benchmark evidence.

Before creating the tag, update the active version references in `README.md`,
`docs/quickstart.md`, `docs/publication-scope.md` and
`docs/library-source-audit.md`, while preserving historical versions and results.

Start from a fresh checkout and use the committed Nix and dependency locks:

```bash
nix flake check --no-update-lock-file --print-build-logs
nix develop
npm ci
./bin/setup --locked --all
npm run build:native -- --keep-workspace
```

The bootstrap prints its workspace. Set `GOPURS_NATIVE_OUTPUT` to its `output`
directory, then run:

```bash
node tools/release-quality.mjs --full --report /tmp/gopurs-quality
```

Keep this command inside the Nix shell. For an equivalent manual environment,
put the typed frontend's directory first on `PATH`: Spago's fixture builds
invoke `purs` through `PATH`. Setting `GOPURS_PURS` alone selects the frontend
for native-build helpers but does not change Spago's executable selection.

The full gate runs every Node test file, Go tests with the race detector,
native emission cancellation/concurrency tests, all pinned library test suites
through `bin/modtest`, and all selected fixtures. Libraries and fixtures run
under both Release and Dev. Both full and quick gates require the freshly
bootstrapped native output so native probes cannot silently skip. Host-parity
tests explicitly select JavaScript, native sequential and native parallel modes.
Experimental Rust tests remain disabled. Without `--full`, the command runs the
tooling/CLI/cache checks, native scanner and preparation probes, Go tests, and a
smaller executable fixture sample suitable for ordinary CI changes.

The full core phase runs each of the 51 libraries pinned in
`bootstrap-sources.json` separately in each profile: 102 module gates. It
reuses the freshly built native compiler without rebuilding it between
libraries. Before and after these gates, the installed library selection,
Git revisions and clean working trees must match the pins. Each library runs
in the disposable source family provided by `bin/modtest`; successful generated
files are removed, and the private Go cache is trimmed between libraries.
Successful campaign JSON and logs are retained under
`module-reports/<profile>/<library>.{json,log}`. The combined quality report
embeds each campaign and rejects missing libraries, wrong pins or profiles,
failed, skipped or resumed targets. These gates apply to future full campaigns
containing this change; they do not add coverage to existing immutable tags.

The scanner, preparation and emission functional probes compile the large
`gopurs/output/purescript` package with `-N -l` to reduce compilation memory.
Their assertions, execution deadlines, scanner stack limit and preparation/
emission race checks remain active. The release compiler bootstrap and the
Release/Dev application fixture profiles keep their own compilation settings.

The command creates its own Go build cache and never prunes the caller's
`GOCACHE`. It runs Node test files in separate sequential processes, then
trims the private cache to 2 GiB after each completed file or Go test command.
The full fixture selection runs once per profile in batches of at most five.
Concurrent batches are limited to one fewer than the available CPUs, with a
minimum of one and maximum of six (two batches on three CPUs, three on four).
Each batch uses `GOMAXPROCS=2`, `GOGC=100`, `GOMEMLIMIT=4GiB`, and backend
prepare/PBO/emission worker counts of 1/1/2. Snapshot updates are forced off.

Before concurrent fixture batches start, the controller runs `spago fetch`
once in a disposable fixture workspace, using the same shared registry cache.
This serial initialization prevents fresh Spago processes from racing to create
the registry. It does not compile fixtures. The controller records
`fixturePreparation` and `fixture-preparation.log`, stops if fetch fails, then
runs every fixture in its own fresh workspace as before.

The cache is trimmed only after every process in a wave has finished. It may
grow beyond 2 GiB during a command or wave. Before a fixture wave, the command
requires 3 GiB free plus 0.75 GiB per batch. A failed batch prevents later
waves from starting; already active batches finish and retain their reports.
Each successful batch must report the exact requested profile and targets,
with no exclusions, failures, pending targets or snapshot updates.

`--report` retains command logs, fixture workspaces/reports and `quality.json`,
including cache sizes before and after trimming. Use a fresh report directory;
without the option, a temporary evidence directory is created and printed.
The private Go cache is removed when the command finishes or reports failure;
failed fixture sources, generated output and logs remain for diagnosis.
This resource policy follows the completed local fixture campaign. The new
orchestration and its platform CI runs require their own recorded validation.

Run the [cold/warm benchmark protocol](benchmark-protocol.md), compare with the
official `altbak.pub` README baselines, and retain the raw results. Also validate
a complete b8x build against the installed candidate and verify the public
README with a new application. These are additional release gates; the small
application smoke below does not replace them.

## Prepare reviewable local artifacts

After building the JavaScript and native hosts from the same source tree:

```bash
node tools/release-prepare.mjs \
  --native bin/gopurs-native --out /tmp/gopurs-candidate --preview
```

The output directory must be empty. `--preview` permits an uncommitted tree
and records that fact in the build metadata. Without it, preparation requires
a clean tree at the matching annotated `vVERSION` tag and a final changelog
section. Preparation packages existing binaries; it does not rebuild them or
establish that an arbitrary supplied native binary matches the source. CI
always bootstraps a fresh binary immediately before preparation.

Each platform produces:

- `gopurs-VERSION.tgz`, the npm archive;
- `gopurs-VERSION-TARGET`, the native executable;
- `gopurs-VERSION-TARGET-SHA256SUMS`;
- a separate checksum list for the JavaScript bundle and FFI parser files
  inside the npm package;
- a CycloneDX 1.6 dependency inventory, build metadata and release notes;
- copies of the npm, Spago and Nix locks and the bootstrap source manifest.

The inventory includes npm build dependencies, PureScript packages, native
bootstrap libraries and reference Node/Go toolchains. It records source inputs,
not optimizer reachability; optional build tools need not appear in the final
executable. Transitive Nix/Haskell build dependencies are pinned by `flake.lock`
and the frontend derivation rather than expanded into this inventory. The
unsigned `build.json` records input hashes and Go binary build information;
cryptographic provenance is generated by the tag workflow.

Validate the prepared files on their own platform:

```bash
cd /tmp/gopurs-candidate
shasum -a 256 -c gopurs-VERSION-TARGET-SHA256SUMS
cd /path/to/gopurs
node tools/release-smoke.mjs \
  --package /tmp/gopurs-candidate/gopurs-VERSION.tgz \
  --native /tmp/gopurs-candidate/gopurs-VERSION-TARGET \
  --purs "$GOPURS_PURS" --report /tmp/gopurs-candidate-smoke.json
npm publish --dry-run --ignore-scripts --tag next /tmp/gopurs-candidate/gopurs-VERSION.tgz
```

The smoke installs the tarball offline, with an empty npm cache and lifecycle
scripts disabled. It creates an application with fresh TAST and Go output,
builds it and checks its printed result using both the installed JavaScript
host and the installed native host. The FFI covers `String`, `Int` through Go
`int32`, and `Effect`. Its minimal source modules require neither registry
packages nor sibling repositories. All temporary Go/npm caches are deleted on
success; failed workspaces are retained with their paths printed.

## Signed tag and CI artifacts

Commit the reviewed source, version and notes, then create a signed annotated
tag using the maintainer's configured Git signing key:

```bash
git tag -s vVERSION -m 'gopurs VERSION'
git verify-tag vVERSION
git push origin vVERSION
```

The tag's signature must also be verified by GitHub. The CI workflow checks
that status through the GitHub API. Configure the signing key on the maintainer
account before tagging; this process does not create or distribute a private
signing key.

The CI checkouts select tag refs explicitly so the checkout action preserves
the annotated object. Each job checks its source commit and tag immediately;
the build also compares the local tag object with the GitHub-verified object
before installing Nix or compiling. Ordinary PRs retain the checkout action's
default event-commit selection. The planning job also checks the release version,
matching tag, both package-lock version fields and changelog section before
starting any tagged build.

`.github/workflows/ci.yml` runs on branch changes, pull requests and version
tags. Ordinary PRs and branch pushes run 24 quick gates on Linux amd64,
including output-identity and documentation checker tests, embedded-runtime
consistency, compiler snapshot tests, both compiler hosts, the native race
probe and artifact smokes.
A small planning job always runs the classification, quality-inventory and
qualification-reuse tests before any build. It skips heavy builds only when every changed path is an explicitly
allowed regular Markdown document. Renames are examined as deletions and
additions; code under `docs/`, file type changes and unknown paths require a
build. Manual non-full runs also build Linux.

A push to `main` may reuse one successful, associated, same-repository PR
qualification. Its small `linux-quick-v2` receipt binds the 24 gates, each
Node file's actual passing test counts with no skips, both artifact smokes, the run attempt and
the actual tested merge commit. Receipts using the old `linux-quick-v1`
policy or missing any required gate cannot be reused.
The planner verifies the artifact ZIP against GitHub's SHA-256 and compares the
tested Git object's tree with the new `main` tree. The PR head alone is
insufficient. Missing, expired, ambiguous or inconsistent evidence, fork runs
and API failures all cause a new Linux build. The scope decision and reused
run identity are retained in `ci-plan.json`.

Initial tag campaigns and manual `full` runs build all three platforms and run the
complete qualification; documentation filtering and PR reuse never apply to
them. Each build uses the locked Nix environment and a fresh checkout. A successful tag
run packages the candidate, verifies checksums, runs the offline application
smoke and npm publication dry-run, and signs provenance and SBOM attestations
using GitHub OIDC/Sigstore. The portable signature bundles are retained beside
the artifacts. See the [GitHub attestation documentation](https://docs.github.com/en/actions/how-tos/secure-your-work/use-artifact-attestations/use-artifact-attestations).

Quick builds may reuse the pinned frontend tool's runtime closure as described
in [nix.md](nix.md#ci-frontend-cache). This optional cache never contains the
gopurs compiler being qualified and is disabled for tags and full campaigns.

On a release tag, the Linux amd64 job also exports
`gopurs-VERSION-linux-amd64-npm-provenance.sigstore.json`. This separate npm
bundle contains one subject (`pkg:npm/gopurs@VERSION`) and the exact archive's
SHA-512, using npm's SLSA v1 format. The helper uses the public Sigstore API
bundled with pinned npm 11.6.0. Generation requires the tagged GitHub-hosted
CI job, its OIDC endpoint, and matching repository, workflow, commit, version
and clean build metadata. It verifies the resulting signature before exporting
the bundle; CI includes it in the checksums and GitHub attestation. Local
environment variables or an unsigned statement are not a replacement for
that CI signature.

Only after every platform succeeds does CI compare the npm tarballs byte for
byte and create a **draft** GitHub release containing the tested artifacts.
Different tarball hashes block this step. The native binaries have separate
platform hashes; the process does not claim their cross-platform identity.
Quality/bootstrap logs are retained as workflow artifacts for 90 days. Copy
long-term release evidence into the release archive before that retention
period expires.

### Retrying finalization

After all three builds and twelve fixture jobs succeed, a transient signing
failure can be retried with `gh run rerun RUN_ID --failed`. This reuses the
successful qualification from that same run, commit and immutable annotated
tag. Finalization verifies each input artifact's ID, SHA-256, size and producer
job before extraction. All qualification jobs must still belong to one
successful attempt; a newer failure, mixed attempts or missing evidence blocks
reuse.

The quality report retains its original qualification attempt and compiler
hashes. Its separate `finalization` record binds the actual signing attempt to
those inputs. Final artifacts use `release-TARGET-attempt-N` names. Draft
creation selects the latest successful finalizer for each platform, allowing
an unchanged successful platform to coexist with a retried one.

If only draft creation/upload failed, rerun that job with
`gh run rerun RUN_ID --job JOB_ID`. Identical assets already uploaded are kept;
only missing assets are added. Different bytes, unexpected assets or an already
published release stop recovery. Do not rerun successful signing jobs after
starting a draft: new signatures would change its assets.

This recovery uses the workflow stored at the original tag. A source or workflow
fix still requires a new candidate and fresh qualification. Failed compiler or
fixture qualification is not eligible for this finalization-only recovery.

RC7 exposed a limit in this retry path: draft visibility failed immediately
after creation, and the retry metadata copied successful jobs into attempt 2
without new executions or artifacts. Its selector then requested a nonexistent
attempt-2 release archive. The workflow remains failed. The maintainer
authorized a specific operator recovery using the cryptographically verified
original attempt-1 assets, without rebuilding or changing the signed tag.
This exception and the completed public checks are recorded in the
[RC7 quality record](quality-0.1.0.md#rc7-qualification-and-operator-recovery--30-october-2026).

The post-RC7 selector recognizes copied successful jobs whose new creation time
is after their recorded completion. Their status, execution timestamps, runner
identity and complete steps must match the latest original execution. It keeps
that original job ID and artifact attempt; missing or inconsistent evidence is
rejected. A genuinely newer failed or pending execution still blocks recovery.
This also preserves the original qualification identity when only finalization
or draft preparation is retried.

The public npm verifier now accepts an absent DSSE `keyid` and an empty string
as equivalent. Both original bundles still undergo cryptographic verification;
payload type, signed payload, signature bytes, signature order and nonempty key
identifiers must match. Unknown signature fields are rejected. These fixes are
covered by local regressions, including a replay of RC7's recorded job/artifact
metadata; they do not modify RC7 or claim a new successful full tag campaign.

Finalization installs only pinned Node 24.8.0 (npm 11.6.0); it does not reinstall
the Nix compiler environment. Builds and fixture jobs keep their locked Nix
environment. The four fixture shards each mix Release and Dev batches while
preserving the complete inventory and existing concurrency limits.

Verify a downloaded executable or tarball against the repository, workflow and
source tag:

```bash
gh attestation verify gopurs-VERSION-TARGET \
  --repo 0x000000000000000000001/gopurs \
  --signer-workflow 0x000000000000000000001/gopurs/.github/workflows/ci.yml \
  --source-ref refs/tags/vVERSION --deny-self-hosted-runners
```

Use `--bundle FILE` to supply a downloaded Sigstore bundle for offline
verification. The [GitHub CLI verification reference](https://cli.github.com/manual/gh_attestation_verify)
describes the identity and digest checks.

## Publish the reviewed candidate

Before the first dispatch, integrate the reviewed release workflow into the
default branch (`main`) through the normal review/merge process. GitHub only
offers `workflow_dispatch` when the workflow exists there, even when `--ref`
selects a release tag. This prerequisite is not approval to merge or publish.
See the [GitHub dispatch requirements](https://docs.github.com/en/actions/how-tos/manage-workflow-runs/manually-run-a-workflow).

Review the CI draft, external benchmark/b8x evidence and rollback prerequisites.
Both normal paths below require the complete tag CI matrix to have succeeded
and the reviewed source to be checked out at that tag. RC7 follows the explicit
operator exception recorded above; it does not satisfy the green-workflow gate.
Use the pinned Nix environment
for npm 11.6.0. Keep downloaded files and local evidence outside the checkout.

For the **first publication**, an npm trusted publisher cannot be configured
on a package that does not yet exist. Publish the real CI archive from the
maintainer's already authenticated npm CLI, supplying the CI-generated bundle.
No placeholder package, npm token export, or npm secret in GitHub is needed.
During RC7's first publication, npm's staging process created its own
`0.0.0-stage` placeholder and then assigned both `next` and `latest` to RC7.
The registry rejected removal of `latest`; the maintainer explicitly accepted
this behavior. An explicit `--tag next` therefore did not establish a
next-only first publication. Verify the public registry state and document
any accepted exception before announcing the release.
First download and verify the draft, replacing `VERSION` with the reviewed
version:

```bash
version=VERSION
tag="v$version"
repo=0x000000000000000000001/gopurs
candidate=$(mktemp -d "${TMPDIR:-/tmp}/gopurs-first-publish.XXXXXX")
test "$(gh release view "$tag" --repo "$repo" --json isDraft --jq .isDraft)" = true
gh release download "$tag" --repo "$repo" --dir "$candidate"
commit=$(git rev-parse "$tag^{commit}")
test "$(git rev-parse HEAD)" = "$commit"
for target in linux-amd64 linux-arm64 darwin-arm64; do
  prefix="gopurs-$version-$target"
  gh attestation verify "$candidate/$prefix-SHA256SUMS" --repo "$repo" \
    --signer-workflow "$repo/.github/workflows/ci.yml" \
    --source-ref "refs/tags/$tag" --source-digest "$commit" --deny-self-hosted-runners
  (cd "$candidate" && shasum -a 256 -c "$prefix-SHA256SUMS")
done
prefix="gopurs-$version-linux-amd64"
nix develop --no-update-lock-file --command node tools/release-npm-provenance.mjs verify \
  --package "$candidate/gopurs-$version.tgz" --build "$candidate/$prefix-build.json" \
  --bundle "$candidate/$prefix-npm-provenance.sigstore.json" \
  > "$candidate/local-provenance-verification.json"
```

Record the registry state before publishing. Only the package's `E404`
establishes this first-package case; other errors are not evidence of absence.
If the package already exists, stop this first-publication path and use its
owner/trusted-publisher configuration instead.

```bash
export GOPURS_FIRST_PUBLISH_EVIDENCE="$candidate"
nix develop --no-update-lock-file --command node --input-type=module <<'JS'
import { spawnSync } from 'node:child_process';
import { writeFileSync } from 'node:fs';
const result = spawnSync('npm', ['view', 'gopurs', 'dist-tags', '--json', '--registry=https://registry.npmjs.org'], { encoding: 'utf8' });
const value = JSON.parse(result.stdout || '{}');
if (result.status === 0 || value.error?.code !== 'E404') throw new Error('First publication requires the confirmed absence of gopurs');
writeFileSync(process.env.GOPURS_FIRST_PUBLISH_EVIDENCE + '/registry-before.json', JSON.stringify(value, null, 2) + '\n');
writeFileSync(process.env.GOPURS_FIRST_PUBLISH_EVIDENCE + '/previous-dist-tags.json', '{}\n');
JS
```

After review authorizes this publication, use the existing local account
session (and complete any npm authentication prompt):

```bash
nix develop --no-update-lock-file --command npm whoami --registry=https://registry.npmjs.org
channel=latest
if [[ "$version" == *-rc.* ]]; then channel=next; fi
nix develop --no-update-lock-file --command npm publish --ignore-scripts --access public \
  --registry=https://registry.npmjs.org --tag "$channel" \
  --provenance-file "$candidate/$prefix-npm-provenance.sigstore.json" \
  "$candidate/gopurs-$version.tgz"
gh workflow run publish.yml --ref "$tag" -f tag="$tag" -f already_published=true
```

`--provenance-file` makes npm verify and attach the signed CI bundle; it does
not generate provenance on the local machine. See the
[npm publication option](https://docs.npmjs.com/cli/v11/commands/npm-publish/#provenance-file).
The explicit `already_published` mode never invokes `npm publish`. It requires
the registry's exact package/version/SHA-512, the same signed CI envelope served
by npm, the expected release channel, and the original tag's CI attestations.
It then downloads the public tarball, compares its bytes with the candidate,
runs the application smokes and publishes the matching GitHub release. Keep
the locally recorded previous tags with the workflow's verification evidence;
the tags observed after publishing are not a rollback baseline. A failure
leaves the npm version immutable and the GitHub release in draft for diagnosis.

This first-publication path must be qualified by a real CI OIDC signature and
the registry checks above; unit tests of guards or local previews do not prove
public publication. After the package exists, configure npm's trusted publisher
for this repository and `publish.yml`, with direct `npm publish` allowed. The
public package and repository URL must match. See
[npm trusted publishing](https://docs.npmjs.com/trusted-publishers/).

For **subsequent versions**, dispatch the publication workflow on the reviewed
release tag with its default mode:

```bash
gh workflow run publish.yml --ref vVERSION -f tag=vVERSION
```

The workflow requires the matching draft release and verified CI attestations.
It records previous npm dist-tags, publishes the exact tested tarball with npm
provenance (`next` for `-rc.N`, `latest` for a final version), fetches it again
from the registry and compares it byte for byte. It repeats the fresh offline
application smoke, then publishes the matching GitHub release. A failure after
npm publication requires inspecting the retained publication evidence; reruns
must not try to overwrite an existing npm version. The first-publication
recovery mode specifically requires the exported CI bundle; it does not accept
a different provenance bundle generated by a later trusted-publishing job.

After a successful publication, verify the README installation with the public
version, finish the b8x check using the published artifacts, archive the proof,
and announce the version/platforms/support channel. Keep the candidate and
final release versions immutable.

## Rollback and recovery

Before each final publication, retain the previous version's npm tarball,
platform executables, checksums and signatures, plus the current npm dist-tags.
Test recovery by running `release-smoke.mjs` against those retained artifacts
on the target platform. Record its output and the exact restored version.
A first release has no previous public version to restore: recovery requires
communicating the regression and publishing a corrected version through the
complete workflow. Until then, no validated public rollback is available.

RC7's retained assets and a fresh local macOS reinstallation were checked on
1 November 2026; see the [recovery evidence](quality-0.1.0.md#retained-rc7-recovery-check--1-november-2026).
This does not validate a public channel restoration. RC7 is the only real
published version at that date, and is a prerelease. Its bootstrap predates
the later library fixes, so select it as a fallback only after checking the
affected application's dependency pins and behavior. The npm staging
placeholder `0.0.0-stage` is not a fallback. Refresh the observed dist-tags
immediately before any future publication or restoration.

Removing `latest` does **not** block `npm install gopurs`. npm can fall back to
an eligible published version, including the affected version when it is the
only one. Deprecation produces a warning; it does not prevent installation.
See npm's [version selection implementation](https://github.com/npm/npm-pick-manifest/blob/v10.0.0/lib/index.js)
and [deprecation reference](https://docs.npmjs.com/cli/v11/commands/npm-deprecate/).

If a published version has a blocking regression:

1. If a suitable previously validated public version exists, restore the
   affected channel using a maintainer account: for a candidate,
   `npm dist-tag add gopurs@PREVIOUS next`; for an approved stable fallback,
   `npm dist-tag add gopurs@PREVIOUS latest`. Record the previous and resulting
   dist-tags and leave other channels unchanged. RC7's accepted first-publish
   `latest` exception is not blanket approval to move that channel later.
   Removing a dist-tag does not substitute for restoring a validated version.
2. Mark the affected version deprecated with a concrete explanation and the
   supported replacement, if one exists; otherwise state that no corrected
   public version is available yet. Update the GitHub release notes and remove its
   latest-release designation. Preserve the version/tag/artifacts for diagnosis.
3. When restoring a previous version, reinstall its retained npm archive and
   matching native executable in a
   fresh directory, verify their signatures/checksums, and rerun the smoke and
   affected application. Also check a fresh registry installation of the
   restored channel selects the expected version: `gopurs@next` for `next`,
   or unversioned `gopurs` for `latest`. Existing lockfiles and explicit version pins
   are not changed by moving a dist-tag. A native compiler is installed as
   `node_modules/gopurs/bin/gopurs-native` with executable permission; invoke
   `node_modules/gopurs/bin/gopurs` to select it. The `npx gopurs` entry stays JS.
4. Publish a new patch or candidate version through the complete workflow.

The selection behavior was exercised on 27 October 2026 with npm 11.6.0 and
its installed `npm-pick-manifest` 10.0.0, using a loopback mock registry, fresh
consumer directories and empty npm caches:

| Mock registry state | Fresh unversioned `npm install` result |
| --- | --- |
| Only `0.1.0`, after `npm dist-tag rm ... latest` | Installed `0.1.0` |
| Only `0.1.0`, deprecated and without `latest` | Installed `0.1.0` with a deprecation warning |
| Eligible `0.1.0` and `0.1.1`, without `latest` | Installed `0.1.1` |
| Both versions, after `npm dist-tag add ...@0.1.0 latest` | Installed `0.1.0` |

All six local npm commands succeeded. This qualifies tag selection locally;
it is not evidence of a public registry restoration, a signed-artifact check,
or a `gopurs` application smoke. No public registry was modified.

For interrupted publication, compare the registry tarball with the signed
candidate first. If npm publication succeeded but only GitHub publication
failed, repeat the public smoke and finish that GitHub release; do not publish
the same npm version again. For transient finalization failures, use the retry
procedure above. If a failure needs a source or workflow fix, issue a new
candidate tag instead of moving the signed tag.

### CI tool lifetime

Each build and fixture job keeps its development environment in
`$RUNNER_TEMP/gopurs-dev-profile`. Keep this profile for the entire job so its
Spago, frontend and language tools have an explicit Nix garbage-collection
root. An `ENOENT` diagnostic records the command search paths and working
directory; investigate it before a costly retry. The signed candidate tag is
immutable, including when its CI fails.
