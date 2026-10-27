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
native emission cancellation/concurrency tests, and all selected fixtures
under both Release and Dev. It requires the freshly
bootstrapped native output so the native preparation and code-generation tests
cannot silently skip. Host-parity tests explicitly select JavaScript, native
sequential and native parallel modes. Experimental Rust tests remain disabled. Without
`--full`, the command runs the tooling/CLI/cache checks, Go tests, and a smaller
executable fixture sample suitable for ordinary CI changes.

The command creates its own Go build cache and never prunes the caller's
`GOCACHE`. It runs Node test files in separate sequential processes, then
trims the private cache to 2 GiB after each completed file or Go test command.
The full fixture selection runs once per profile in batches of at most five.
Concurrent batches are limited to one fewer than the available CPUs, with a
minimum of one and maximum of six (two batches on three CPUs, three on four).
Each batch uses `GOMAXPROCS=2`, `GOGC=100`, `GOMEMLIMIT=4GiB`, and backend
prepare/PBO/emission worker counts of 1/1/2. Snapshot updates are forced off.

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
npm publish --dry-run --ignore-scripts /tmp/gopurs-candidate/gopurs-VERSION.tgz
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

`.github/workflows/ci.yml` runs on branch changes, pull requests and version
tags. All three platforms build inside the locked Nix environment from a fresh
checkout without a restored project cache. Public Nix dependency substitutes
are allowed. Tags automatically select the full quality gate. A successful tag
run packages the candidate, verifies checksums, runs the offline application
smoke and npm publication dry-run, and signs provenance and SBOM attestations
using GitHub OIDC/Sigstore. The portable signature bundles are retained beside
the artifacts. See the [GitHub attestation documentation](https://docs.github.com/en/actions/how-tos/secure-your-work/use-artifact-attestations/use-artifact-attestations).

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
Both paths below require the complete tag CI matrix to have succeeded and the
reviewed source to be checked out at that tag. Use the pinned Nix environment
for npm 11.6.0. Keep downloaded files and local evidence outside the checkout.

For the **first publication**, an npm trusted publisher cannot be configured
on a package that does not yet exist. Publish the real CI archive from the
maintainer's already authenticated npm CLI, supplying the CI-generated bundle.
No placeholder package, npm token export, or npm secret in GitHub is needed.
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

Removing `latest` does **not** block `npm install gopurs`. npm can fall back to
an eligible published version, including the affected version when it is the
only one. Deprecation produces a warning; it does not prevent installation.
See npm's [version selection implementation](https://github.com/npm/npm-pick-manifest/blob/v10.0.0/lib/index.js)
and [deprecation reference](https://docs.npmjs.com/cli/v11/commands/npm-deprecate/).

If a published version has a blocking regression:

1. If a previously validated public version exists, restore `latest` with
   `npm dist-tag add gopurs@PREVIOUS latest`, using a maintainer account. Remove
   the affected `next` tag if it points at a broken candidate. Record the
   resulting dist-tags; do not substitute removal of `latest` for restoration.
2. Mark the affected version deprecated with a concrete explanation and the
   supported replacement, if one exists; otherwise state that no corrected
   public version is available yet. Update the GitHub release notes and remove its
   latest-release designation. Preserve the version/tag/artifacts for diagnosis.
3. When restoring a previous version, reinstall its retained npm archive and
   matching native executable in a
   fresh directory, verify their signatures/checksums, and rerun the smoke and
   affected application. Also check a fresh unversioned registry installation
   selects the restored version. Existing lockfiles and explicit version pins
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
the same npm version again. If a CI tag build failed, fix the source and issue
a new candidate tag instead of moving the signed tag.
