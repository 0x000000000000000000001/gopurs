import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { mkdirSync, mkdtempSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { githubIO, planDraft, readAssets, reconcileDraft, releaseVersion, validateDraftContext } from './release-draft.mjs';

const tag = 'v0.1.0-rc.7', repository = '0x000000000000000000001/gopurs';
const digest = bytes => `sha256:${createHash('sha256').update(bytes).digest('hex')}`;
const assets = [
  { name: 'gopurs-0.1.0-rc.7.tgz', path: '/synthetic/package.tgz', size: 3, digest: digest('npm') },
  { name: 'gopurs-0.1.0-rc.7-linux-amd64', path: '/synthetic/native', size: 6, digest: digest('native') },
];
const remote = (asset, id, hash = asset.digest) => ({ id, name: asset.name, size: asset.size, digest: hash, state: 'uploaded' });
const draft = (items = []) => ({ id: 27, tag_name: tag, draft: true, published_at: null, prerelease: true, assets: items });
function context() {
  const commit = 'a'.repeat(40), object = 'b'.repeat(40);
  return { tag, env: { GITHUB_ACTIONS: 'true', RUNNER_ENVIRONMENT: 'github-hosted', GITHUB_SERVER_URL: 'https://github.com',
    GITHUB_REPOSITORY: repository, GITHUB_EVENT_NAME: 'push', GITHUB_REF: `refs/tags/${tag}`, GITHUB_REF_TYPE: 'tag', GITHUB_REF_NAME: tag,
    GITHUB_SHA: commit, GITHUB_WORKFLOW_SHA: commit, GITHUB_WORKFLOW_REF: `${repository}/.github/workflows/ci.yml@refs/tags/${tag}` },
  git: { head: commit, trackedStatus: '', tagType: 'tag', tagCommit: commit, tagObject: object },
  ref: { ref: `refs/tags/${tag}`, object: { type: 'tag', sha: object } },
  annotated: { sha: object, tag, verification: { verified: true }, object: { type: 'commit', sha: commit } } };
}

test('draft creation requires the exact hosted CI source and verified annotated tag', () => {
  assert.equal(validateDraftContext(context()).commit, 'a'.repeat(40));
  const mutations = [
    c => { c.env.GITHUB_REPOSITORY = 'fork/gopurs'; }, c => { c.env.GITHUB_EVENT_NAME = 'pull_request'; },
    c => { c.env.GITHUB_WORKFLOW_SHA = 'c'.repeat(40); }, c => { c.env.GITHUB_WORKFLOW_REF = `${repository}/publish.yml@main`; },
    c => { c.env.GITHUB_REF = 'refs/heads/main'; }, c => { c.env.RUNNER_ENVIRONMENT = 'self-hosted'; },
    c => { c.git.head = 'c'.repeat(40); }, c => { c.git.trackedStatus = ' M source'; }, c => { c.git.tagType = 'commit'; },
    c => { c.git.tagCommit = 'c'.repeat(40); }, c => { c.ref.object.sha = 'c'.repeat(40); },
    c => { c.annotated.verification.verified = false; }, c => { c.annotated.object.sha = 'c'.repeat(40); },
    c => { c.annotated.tag = 'v0.1.0'; }, c => { c.ref.ref = 'refs/tags/v0.1.0'; },
  ];
  for (const mutate of mutations) { const value = context(); mutate(value); assert.throws(() => validateDraftContext(value)); }
  for (const value of ['main', 'v1.0.0/../../escape', 'v1.0.0;echo', 'v1.0.0\n', 'v1.0.0-beta.1', '-v1.0.0']) assert.throws(() => releaseVersion(value));
  assert.deepEqual(releaseVersion('v1.0.0'), { version: '1.0.0', prerelease: false });
});

test('draft planning preserves exact existing assets and selects only missing uploads', () => {
  assert.deepEqual(planDraft({ tag, assets, release: null }), { create: true, upload: assets, verify: [] });
  const partial = planDraft({ tag, assets, release: draft([remote(assets[0], 1)]) });
  assert.deepEqual(partial, { create: false, upload: [assets[1]], verify: [] });
  assert.equal(planDraft({ tag, assets, release: draft(assets.map((asset, i) => remote(asset, i + 1))) }).upload.length, 0);
  const fallback = planDraft({ tag, assets, release: draft([remote(assets[0], 1, null)]) });
  assert.deepEqual(fallback.verify, [{ id: 1, name: assets[0].name, size: assets[0].size, digest: assets[0].digest }]);
});

test('published, divergent, extra, duplicate and incomplete remote assets are refused', () => {
  const mutations = [
    r => { r.draft = false; }, r => { r.published_at = '2026-10-30T00:00:00Z'; },
    r => { r.tag_name = 'v0.1.0'; }, r => { r.prerelease = false; },
    r => { r.assets[0].size++; }, r => { r.assets[0].digest = digest('bad'); },
    r => { r.assets[0].digest = 'sha512:invalid'; }, r => { r.assets[0].state = 'starter'; },
    r => { r.assets[0].name = '../private'; }, r => { r.assets.push({ ...r.assets[0] }); },
    r => { r.assets.push(remote(assets[1], r.assets[0].id)); }, r => { r.assets[0].id = -1; },
  ];
  for (const mutate of mutations) { const release = draft([remote(assets[0], 1)]); mutate(release); assert.throws(() => planDraft({ tag, assets, release })); }
  for (const name of ['../escape', '/absolute', 'file#label', 'file\n', 'gopurs-0.2.0.tgz']) {
    assert.throws(() => planDraft({ tag, assets: [{ ...assets[0], name }], release: null }));
  }
  assert.throws(() => planDraft({ tag, assets: [...assets, assets[0]], release: null }));
});

function fakeIO(initial) {
  let current = initial === null ? null : structuredClone(initial);
  const calls = [], io = {
    async getRelease() { calls.push('get'); return structuredClone(current); },
    async create(actualTag, notes) { assert.equal(actualTag, tag); assert.equal(notes, '/synthetic/notes'); calls.push('create'); current = draft(); },
    async digest(id, size) { calls.push(`digest:${id}`); assert.equal(size, assets[0].size); return { size, digest: assets[0].digest }; },
    async upload(id, selected) {
      assert.equal(id, 27); calls.push(`upload:${selected.map(a => a.name).join(',')}`);
      current.assets.push(...selected.map((asset, i) => remote(asset, 100 + i)));
    },
  };
  return { io, calls, current: () => current };
}

test('absent draft is created once and verified after adding its exact assets', async () => {
  const { io, calls } = fakeIO(null);
  assert.deepEqual(await reconcileDraft({ tag, assets, notes: '/synthetic/notes', io }), { status: 'draft-verified', releaseId: 27, tag, assets: 2 });
  assert.equal(calls.filter(c => c === 'create').length, 1);
  assert.deepEqual(calls.filter(c => c.startsWith('upload:')), [`upload:${assets.map(a => a.name).join(',')}`]);
});

test('partial-draft recovery hashes missing API digests before uploading and never overwrites', async () => {
  const { io, calls } = fakeIO(draft([remote(assets[0], 1, null)]));
  await reconcileDraft({ tag, assets, notes: '/synthetic/notes', io });
  assert.equal(calls.includes('create'), false);
  assert.deepEqual(calls.filter(c => c.startsWith('upload:')), [`upload:${assets[1].name}`]);
  assert.equal(calls.filter(c => c === 'digest:1').length, 1, 'Immutable asset digest is checked once per invocation');
  assert.ok(calls.indexOf('digest:1') < calls.findIndex(c => c.startsWith('upload:')));
});

test('a corrupt downloaded asset or publication during recovery prevents mutation', async () => {
  const bad = fakeIO(draft([remote(assets[0], 1, null)]));
  bad.io.digest = async () => ({ size: assets[0].size, digest: digest('bad') });
  await assert.rejects(reconcileDraft({ tag, assets, notes: '/synthetic/notes', io: bad.io }));
  assert.equal(bad.calls.some(c => c.startsWith('upload:')), false);
  const race = fakeIO(draft([remote(assets[0], 1)]));
  const get = race.io.getRelease; let count = 0;
  race.io.getRelease = async () => { const value = await get(); if (++count === 2) value.draft = false; return value; };
  await assert.rejects(reconcileDraft({ tag, assets, notes: '/synthetic/notes', io: race.io }));
  assert.equal(race.calls.some(c => c.startsWith('upload:')), false);
});

test('post-upload verification rejects missing, changed and published final assets', async () => {
  for (const change of [r => { r.assets.pop(); }, r => { r.assets[0].digest = digest('changed'); }, r => { r.draft = false; }]) {
    const fake = fakeIO(draft()); const upload = fake.io.upload;
    fake.io.upload = async (...args) => { await upload(...args); change(fake.current()); };
    await assert.rejects(reconcileDraft({ tag, assets, notes: '/synthetic/notes', io: fake.io }));
  }
});

test('filesystem inventory hashes regular files and refuses symlinks, directories and unsafe names', t => {
  const root = mkdtempSync(join(tmpdir(), 'gopurs-draft-test-')); t.after(() => rmSync(root, { recursive: true, force: true }));
  const directory = join(root, 'assets'); mkdirSync(directory);
  const name = assets[0].name; writeFileSync(join(directory, name), 'npm');
  assert.deepEqual(readAssets(directory), [{ name, path: join(directory, name), size: 3, digest: digest('npm') }]);
  symlinkSync(directory, join(root, 'linked-directory')); assert.throws(() => readAssets(join(root, 'linked-directory')));
  symlinkSync(join(directory, name), join(directory, 'linked')); assert.throws(() => readAssets(directory)); rmSync(join(directory, 'linked'));
  mkdirSync(join(directory, 'nested')); assert.throws(() => readAssets(directory)); rmSync(join(directory, 'nested'), { recursive: true });
  writeFileSync(join(directory, 'unsafe#label'), 'x'); assert.throws(() => readAssets(directory));
});

test('GitHub adapter creates only drafts and uploads immutable assets by release ID', t => {
  const root = mkdtempSync(join(tmpdir(), 'gopurs-draft-adapter-')); t.after(() => rmSync(root, { recursive: true, force: true }));
  const path = join(root, assets[0].name); writeFileSync(path, 'npm');
  const calls = [], source = validateDraftContext(context());
  const run = (command, args) => { assert.equal(command, 'gh'); calls.push(args); return Buffer.from('npm'); };
  const io = githubIO(source, run);
  io.create(tag, '/notes.md');
  const create = calls[0];
  assert.deepEqual(create.slice(0, 3), ['release', 'create', tag]);
  assert.ok(create.includes('--verify-tag') && create.includes('--draft') && create.includes('--prerelease'));
  assert.equal(create[create.indexOf('--target') + 1], source.commit);
  assert.equal(create[create.indexOf('--repo') + 1], repository);
  io.upload(27, [{ ...assets[0], path }]);
  assert.deepEqual(calls[1], ['api', `https://uploads.github.com/repos/${repository}/releases/27/assets?name=${assets[0].name}`,
    '--method', 'POST', '--header', 'Content-Type: application/octet-stream', '--input', path]);
  assert.deepEqual(io.digest(41, 3), { size: 3, digest: assets[0].digest });
  assert.ok(calls[2].includes(`repos/${repository}/releases/assets/41`));
  writeFileSync(path, 'changed');
  assert.throws(() => io.upload(27, [{ ...assets[0], path }]));
  assert.equal(calls.length, 3, 'Changed local bytes must be rejected before network access');
  githubIO({ ...source, tag: 'v1.0.0', ...releaseVersion('v1.0.0') }, run).create('v1.0.0', '/notes.md');
  assert.equal(calls[3].includes('--prerelease'), false);
});
