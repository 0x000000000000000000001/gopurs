#!/usr/bin/env node
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { createHash } from 'node:crypto';
import { closeSync, constants, fstatSync, lstatSync, openSync, readSync, readdirSync } from 'node:fs';
import { resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const repository = '0x000000000000000000001/gopurs';
const maxAssetBytes = 512 * 1024 ** 2;
const safeName = /^[A-Za-z0-9][A-Za-z0-9._-]*$/;
const shaPattern = /^sha256:[a-f0-9]{64}$/;

export function releaseVersion(tag) {
  assert.match(tag, /^v\d+\.\d+\.\d+(?:-rc\.\d+)?$/);
  return { version: tag.slice(1), prerelease: /-rc\.\d+$/.test(tag) };
}

function validateLocalContext({ tag, env, git }) {
  releaseVersion(tag);
  assert.equal(env.GITHUB_ACTIONS, 'true');
  assert.equal(env.RUNNER_ENVIRONMENT, 'github-hosted');
  assert.equal(env.GITHUB_SERVER_URL, 'https://github.com');
  assert.equal(env.GITHUB_REPOSITORY, repository);
  assert.ok(['push', 'workflow_dispatch'].includes(env.GITHUB_EVENT_NAME));
  assert.equal(env.GITHUB_REF, `refs/tags/${tag}`);
  assert.equal(env.GITHUB_REF_TYPE, 'tag');
  assert.equal(env.GITHUB_REF_NAME, tag);
  assert.match(env.GITHUB_SHA, /^[a-f0-9]{40}$/);
  assert.equal(env.GITHUB_WORKFLOW_SHA, env.GITHUB_SHA);
  assert.equal(env.GITHUB_WORKFLOW_REF, `${repository}/.github/workflows/ci.yml@refs/tags/${tag}`);
  assert.equal(git.head, env.GITHUB_SHA);
  assert.equal(git.trackedStatus, '');
  assert.equal(git.tagType, 'tag');
  assert.equal(git.tagCommit, git.head);
  assert.match(git.tagObject, /^[a-f0-9]{40}$/);
  return { repository, tag, commit: git.head, ...releaseVersion(tag) };
}

export function validateDraftContext({ tag, env, git, ref, annotated }) {
  const context = validateLocalContext({ tag, env, git });
  assert.equal(ref.ref, `refs/tags/${tag}`);
  assert.deepEqual({ type: ref.object?.type, sha: ref.object?.sha }, { type: 'tag', sha: git.tagObject });
  assert.equal(annotated.sha, git.tagObject);
  assert.equal(annotated.tag, tag);
  assert.equal(annotated.verification?.verified, true);
  assert.deepEqual({ type: annotated.object?.type, sha: annotated.object?.sha }, { type: 'commit', sha: git.head });
  return context;
}

function fileDigest(path) {
  const descriptor = openSync(path, constants.O_RDONLY | constants.O_NOFOLLOW);
  try {
    const stat = fstatSync(descriptor);
    assert.ok(stat.isFile() && stat.size <= maxAssetBytes, 'Expected a bounded regular asset file');
    const hash = createHash('sha256'), buffer = Buffer.alloc(1024 * 1024);
    for (let count; (count = readSync(descriptor, buffer, 0, buffer.length, null)) > 0;) hash.update(buffer.subarray(0, count));
    return { size: stat.size, digest: `sha256:${hash.digest('hex')}` };
  } finally { closeSync(descriptor); }
}

export function readAssets(directory) {
  assert.ok(lstatSync(directory).isDirectory() && !lstatSync(directory).isSymbolicLink(), 'Expected an asset directory, not a symlink');
  const names = readdirSync(directory).sort();
  assert.ok(names.length > 0 && names.length <= 200, 'Invalid asset inventory');
  return names.map(name => {
    assert.match(name, safeName);
    const path = resolve(directory, name);
    return { name, path, ...fileDigest(path) };
  });
}

export function planDraft({ tag, assets, release }) {
  const { version, prerelease } = releaseVersion(tag);
  assert.ok(Array.isArray(assets) && assets.length > 0 && assets.length <= 200);
  const expected = new Map();
  for (const asset of assets) {
    assert.match(asset.name, safeName);
    assert.ok(asset.name.startsWith(`gopurs-${version}.`) || asset.name.startsWith(`gopurs-${version}-`), 'Asset belongs to another release');
    assert.ok(!expected.has(asset.name), 'Duplicate local asset');
    assert.ok(Number.isSafeInteger(asset.size) && asset.size >= 0 && asset.size <= maxAssetBytes);
    assert.match(asset.digest, shaPattern);
    expected.set(asset.name, asset);
  }
  if (release === null) return { create: true, upload: [...assets], verify: [] };
  assert.ok(Number.isSafeInteger(release.id) && release.id > 0);
  assert.equal(release.tag_name, tag, 'Draft tag differs');
  assert.equal(release.draft, true, 'Published releases cannot be modified');
  assert.equal(release.published_at, null, 'Published releases cannot be modified');
  assert.equal(release.prerelease, prerelease, 'Draft prerelease policy differs');
  assert.ok(Array.isArray(release.assets));
  const names = new Set(), ids = new Set(), verify = [];
  for (const remote of release.assets) {
    assert.ok(Number.isSafeInteger(remote.id) && remote.id > 0 && !ids.has(remote.id), 'Invalid or repeated remote asset ID');
    assert.ok(!names.has(remote.name), 'Duplicate remote asset name');
    names.add(remote.name); ids.add(remote.id);
    const local = expected.get(remote.name);
    assert.ok(local, `Unexpected existing asset: ${remote.name}`);
    assert.equal(remote.state, 'uploaded', 'An existing asset is not fully uploaded');
    assert.equal(remote.size, local.size, `Existing asset size differs: ${remote.name}`);
    if (remote.digest === undefined || remote.digest === null) verify.push({ id: remote.id, name: remote.name, size: local.size, digest: local.digest });
    else {
      assert.match(remote.digest, shaPattern);
      assert.equal(remote.digest, local.digest, `Existing asset digest differs: ${remote.name}`);
    }
  }
  return { create: false, upload: assets.filter(asset => !names.has(asset.name)), verify };
}

async function verifyDownloads(plan, io, verified) {
  for (const asset of plan.verify) {
    const key = `${asset.id}:${asset.size}:${asset.digest}`;
    if (verified.has(key)) continue;
    const actual = await io.digest(asset.id, asset.size);
    assert.equal(actual.size, asset.size, `Downloaded asset size differs: ${asset.name}`);
    assert.equal(actual.digest, asset.digest, `Downloaded asset digest differs: ${asset.name}`);
    verified.add(key);
  }
}

// No deletion, overwrite, clobber, publication or waiting occurs here. A failed
// upload leaves a draft whose exact matching assets can be reused next time.
export async function reconcileDraft({ tag, assets, notes, io }) {
  const verified = new Set();
  let release = await io.getRelease(tag);
  let plan = planDraft({ tag, assets, release });
  if (plan.create) {
    await io.create(tag, notes);
    release = await io.getRelease(tag);
    assert.ok(release, 'Created draft is not available yet');
    plan = planDraft({ tag, assets, release });
  }
  const id = release.id;
  await verifyDownloads(plan, io, verified);
  if (plan.upload.length) {
    // Check again immediately before mutation. Uploads address this release ID,
    // never resolve the tag to a potentially replaced release or overwrite data.
    const current = await io.getRelease(tag);
    assert.equal(current?.id, id, 'Draft identity changed');
    plan = planDraft({ tag, assets, release: current });
    await verifyDownloads(plan, io, verified);
    if (plan.upload.length) await io.upload(id, plan.upload);
  }
  const final = await io.getRelease(tag);
  assert.equal(final?.id, id, 'Draft identity changed');
  const checked = planDraft({ tag, assets, release: final });
  assert.equal(checked.upload.length, 0, 'The draft is still missing assets');
  await verifyDownloads(checked, io, verified);
  return { status: 'draft-verified', releaseId: id, tag, assets: assets.length };
}

export function githubIO(context, run = execFileSync) {
  const gh = (args, options = {}) => run('gh', args, { timeout: 120000, maxBuffer: 32 * 1024 ** 2, ...options });
  const api = path => JSON.parse(gh(['api', path]));
  const pages = path => {
    const value = JSON.parse(gh(['api', path, '--paginate', '--slurp']));
    assert.ok(Array.isArray(value) && value.every(Array.isArray));
    return value.flat();
  };
  const base = `repos/${context.repository}`;
  return {
    api,
    getRelease(tag) {
      // Listing authenticated releases includes drafts; the by-tag endpoint
      // is documented for published releases and is insufficient for recovery.
      const matches = pages(`${base}/releases?per_page=100`).filter(release => release.tag_name === tag);
      assert.ok(matches.length <= 1, 'Ambiguous release tag');
      if (!matches.length) return null;
      const release = api(`${base}/releases/${matches[0].id}`);
      return { ...release, assets: pages(`${base}/releases/${release.id}/assets?per_page=100`) };
    },
    create(tag, notes) {
      gh(['release', 'create', tag, '--repo', context.repository, '--verify-tag', '--draft',
        '--target', context.commit, '--title', `gopurs ${context.version}`, '--notes-file', notes,
        ...(context.prerelease ? ['--prerelease'] : [])]);
    },
    upload(id, assets) {
      for (const asset of assets) {
        assert.deepEqual(fileDigest(asset.path), { size: asset.size, digest: asset.digest }, 'Local asset changed before upload');
        gh(['api', `https://uploads.github.com/${base}/releases/${id}/assets?name=${encodeURIComponent(asset.name)}`,
          '--method', 'POST', '--header', 'Content-Type: application/octet-stream', '--input', asset.path]);
      }
    },
    digest(id, size) {
      const bytes = gh(['api', `${base}/releases/assets/${id}`, '--header', 'Accept: application/octet-stream'], { maxBuffer: size + 1024 });
      return { size: bytes.length, digest: `sha256:${createHash('sha256').update(bytes).digest('hex')}` };
    },
  };
}

async function main(args) {
  const options = {};
  for (let index = 0; index < args.length; index += 2) {
    assert.ok(['--tag', '--assets', '--notes'].includes(args[index]) && args[index + 1], 'Expected --tag TAG --assets DIR --notes FILE');
    assert.ok(!options[args[index]], 'Duplicate option');
    options[args[index]] = args[index + 1];
  }
  assert.ok(options['--tag'] && options['--assets'] && options['--notes'], 'Expected --tag TAG --assets DIR --notes FILE');
  const tag = options['--tag']; releaseVersion(tag);
  const git = (...args) => execFileSync('git', args, { encoding: 'utf8' }).trim();
  const env = process.env, local = { head: git('rev-parse', 'HEAD'), trackedStatus: git('status', '--porcelain', '--untracked-files=no'),
    tagType: git('cat-file', '-t', `refs/tags/${tag}`), tagCommit: git('rev-parse', `refs/tags/${tag}^{commit}`), tagObject: git('rev-parse', `refs/tags/${tag}`) };
  const context = validateLocalContext({ tag, env, git: local });
  const io = githubIO(context);
  validateDraftContext({ tag, env, git: local,
    ref: io.api(`repos/${repository}/git/ref/tags/${tag}`), annotated: io.api(`repos/${repository}/git/tags/${local.tagObject}`) });
  const notes = resolve(options['--notes']);
  assert.ok(lstatSync(notes).isFile() && !lstatSync(notes).isSymbolicLink(), 'Expected a regular notes file');
  console.log(JSON.stringify(await reconcileDraft({ tag, assets: readAssets(resolve(options['--assets'])), notes, io }), null, 2));
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) await main(process.argv.slice(2));
