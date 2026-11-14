# Publish a gopurs release

Use this guide to publish a new gopurs version. Run the commands from the
repository root, using the same shell so the version variable remains available.
The [full release procedure](release.md) defines the qualification requirements
and recovery paths.

The sequence is: prepare and merge the release changes, push a signed tag,
review the successful CI draft, then dispatch publication.

## 1. Prepare the version

Start with a clean working tree on your release preparation branch. Choose an
unused version: `X.Y.Z` for a final release or `X.Y.Z-rc.N` for a candidate.
Replace `X.Y.Z` below with that version:

```bash
git status --short
release_version=X.Y.Z
npm version "$release_version" --no-git-tag-version
```

If `git status --short` lists changes, finish or set aside that work before
running `npm version`. The version command updates `package.json` and
`package-lock.json`; its hook regenerates and stages the marked documentation
blocks. It does not create a release tag.

Add a `## [VERSION]` section, using the chosen version, to
[CHANGELOG.md](../CHANGELOG.md), describing the changes,
the exact TAST revision, compatibility limits and qualification evidence.
Check that `bootstrap-sources.json` and the dependency locks select the fork,
PBO and libraries intended for distribution. Local sibling checkouts alone
do not select the sources used by release CI.

If you changed source pins, run `npm run docs:sync`, then review and check:

```bash
git diff HEAD
npm run check:docs
```

## 2. Review and merge the release changes

Commit the reviewed version, lockfile, changelog, generated documentation and
any intended source or dependency updates. Push and merge those changes into
`main` through the project's normal review process. Resolve failing checks
before continuing.

Confirm that the release workflow is present on `main`. GitHub requires it
on the default branch even when publication is dispatched from a tag.

## 3. Push the signed release tag

Update your local `main`. Confirm that it contains the reviewed release
changes, that its package version equals `$release_version`, and that the
working tree is clean before creating the tag:

```bash
git switch main
git pull --ff-only
git status --short
git show HEAD:package.json
git tag -s "v$release_version" -m "gopurs $release_version"
git verify-tag "v$release_version"
git push origin "v$release_version"
```

The tag must be annotated and signed with a key whose signature GitHub
recognizes as verified. Configure that key on the maintainer account first.
CI requires the tag, package version, both lockfile version fields and
changelog heading to agree. Keep published tags immutable.

## 4. Wait for qualification and review the draft

The tag triggers the [CI workflow](../.github/workflows/ci.yml) with complete
qualification on Linux amd64, Linux arm64 and macOS arm64. Wait for every
required job to succeed. CI then creates a **draft** GitHub release containing
the tested npm archive, native executables, checksums and provenance.

Before publication, review that draft and retain the required b8x validation,
fresh-consumer checks and [cold/warm benchmark evidence](benchmark-protocol.md),
including comparison with the official `altbak.pub` README baselines.
Record results in the [quality record](quality-0.1.0.md) and release notes,
preserving the actual candidate and artifact identities. A successful local
build alone does not qualify a release.

## 5. Publish the reviewed artifacts

For an existing npm package, configure its trusted publisher for this
repository and `publish.yml`. Authenticate the GitHub CLI as a maintainer.
If the package has never been published, follow the
[first-publication procedure](release.md#publish-the-reviewed-candidate)
instead. For subsequent versions:

```bash
git switch --detach "v$release_version"
gh workflow run publish.yml --ref "v$release_version" -f tag="v$release_version"
```

Both `--ref` and `tag` must name the reviewed tag. The
[publication workflow](../.github/workflows/publish.yml) verifies the draft,
publishes the tested archive to npm (`latest` for a final version, `next` for
a candidate), downloads it to verify its bytes
and installation, then makes the GitHub release public. Wait for that workflow
to succeed before announcing the release. Verify the published installation
and b8x with the public artifacts as described in the full procedure.

Return to your development branch afterwards with `git switch <branch>`.

## Candidates and failed runs

- Release candidates use the same procedure and full qualification as final
  versions. Changing from a candidate version to a final version creates a
  new release identity and requires qualification under the release procedure.
- If CI fails, inspect the failing job. An infrastructure-only failure can
  use the [targeted retry procedure](release.md#reusing-successful-jobs-after-a-targeted-retry).
  A source correction requires a new candidate and qualification.
- If publication fails after npm accepted the version, inspect its retained
  evidence before retrying. npm versions are immutable; use the applicable
  [publication recovery procedure](release.md#publish-the-reviewed-candidate).
  Do not blindly rerun publication or move the tag.
