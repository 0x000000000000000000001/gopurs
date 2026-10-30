import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { selectFinalizedInputs, validateFinalizedQuality, verifyFinalizedDirectory } from './release-finalized.mjs';

const targets = ['linux-amd64', 'linux-arm64', 'darwin-arm64'];
const sha = bytes => createHash('sha256').update(bytes).digest('hex');
function fixture() {
  const context = { repository: 'owner/gopurs', commit: 'a'.repeat(40), tag: 'v0.1.0-rc.7', runId: '42', runAttempt: '2', target: targets[0] };
  const jobs = targets.flatMap(target => [`build and test (${target})`, ...[0, 1, 2, 3].map(shard => `fixtures (${target}, shard ${shard})`)])
    .map((name, index) => ({ id: index + 1, name, run_id: 42, run_attempt: 1, head_sha: context.commit,
      status: 'completed', conclusion: 'success', started_at: '2026-10-30T00:00:00Z', completed_at: '2026-10-30T00:10:00Z' }));
  const artifact = (name, id, created_at) => ({ id, name, size_in_bytes: 128, digest: 'sha256:' + 'b'.repeat(64), expired: false,
    created_at, workflow_run: { id: 42, head_sha: context.commit } });
  const artifacts = targets.flatMap(target => [`candidate-${target}`, ...[0, 1, 2, 3].map(shard => `fixtures-${target}-${shard}`)])
    .map((name, index) => artifact(name, 100 + index, '2026-10-30T00:05:00Z'));
  targets.forEach((target, index) => {
    const attempt = index === 0 ? 2 : 1, hour = attempt === 2 ? '01' : '00';
    jobs.push({ id: 50 + index, name: `qualify and attest (${target})`, run_id: 42, run_attempt: attempt, head_sha: context.commit,
      status: 'completed', conclusion: 'success', started_at: `2026-10-30T${hour}:11:00Z`, completed_at: `2026-10-30T${hour}:20:00Z` });
    artifacts.push(artifact(`release-${target}-attempt-${attempt}`, 200 + index, `2026-10-30T${hour}:19:00Z`));
  });
  const tag = { ref: { ref: 'refs/tags/' + context.tag, object: { type: 'tag', sha: 'c'.repeat(40) } },
    object: { sha: 'c'.repeat(40), tag: context.tag, object: { type: 'commit', sha: context.commit }, verification: { verified: true } } };
  const run = { id: 42, run_attempt: 2, head_sha: context.commit, head_branch: context.tag, event: 'push', path: '.github/workflows/ci.yml',
    status: 'in_progress', repository: { full_name: context.repository }, head_repository: { full_name: context.repository } };
  return { context, jobs, artifacts, tag, run };
}
function environment(context) {
  return { CI: 'true', GITHUB_ACTIONS: 'true', RUNNER_ENVIRONMENT: 'github-hosted', GITHUB_SERVER_URL: 'https://github.com',
    GITHUB_EVENT_NAME: 'push', GITHUB_REF_TYPE: 'tag', GITHUB_REPOSITORY: context.repository, GITHUB_SHA: context.commit,
    GITHUB_REF_NAME: context.tag, GITHUB_REF: 'refs/tags/' + context.tag, GITHUB_RUN_ID: context.runId, GITHUB_RUN_ATTEMPT: context.runAttempt,
    GITHUB_WORKFLOW_SHA: context.commit, GITHUB_WORKFLOW_REF: context.repository + '/.github/workflows/ci.yml@refs/tags/' + context.tag };
}
function quality(selected) {
  const identity = { source: { commit: selected.finalization.source.commit, dirty: false }, target: selected.target,
    compiler: { javascriptSha256: sha('javascript'), nativeSha256: sha('native') },
    github: { repository: selected.finalization.github.repository, runId: selected.finalization.github.runId,
      runAttempt: selected.finalization.github.qualificationAttempt } };
  return { schema: 1, requestedFull: true, full: true, phase: 'all', shard: null, status: 'passed', identity,
    finalization: { ...structuredClone(selected.finalization), qualificationIdentity: structuredClone(identity) } };
}

test('draft retry accepts mixed successful finalization attempts while retaining qualification one', () => {
  const data = fixture(), selected = selectFinalizedInputs(data);
  assert.deepEqual(selected.map(item => item.job.attempt), ['2', '1', '1']);
  assert.deepEqual(selected.map(item => item.artifact.id), [200, 201, 202]);
  for (const item of selected) {
    assert.equal(item.finalization.github.qualificationAttempt, '1');
    assert.equal(item.finalization.jobs.length, 15);
    assert.equal(validateFinalizedQuality(quality(item), item, environment(data.context)).identity.github.runAttempt, '1');
  }
  data.context.runAttempt = '3'; data.run.run_attempt = 3;
  assert.deepEqual(selectFinalizedInputs(data).map(item => item.artifact.id), [200, 201, 202]);
});

test('draft-only retry keeps original producers when GitHub copies completed jobs into a later attempt', () => {
  const data = fixture();
  for (const job of data.jobs) {
    job.created_at = new Date(Date.parse(job.started_at) - 1000).toISOString();
    job.runner_id = 123; job.runner_name = 'GitHub Actions 123';
    job.steps = [{ name: 'Produce artifact', number: 1, status: job.status, conclusion: job.conclusion,
      started_at: job.started_at, completed_at: job.completed_at }];
  }
  const original = selectFinalizedInputs(data);
  const copies = data.jobs.map(job => ({ ...structuredClone(job), id: job.id + 1000, run_attempt: 3,
    created_at: '2026-10-30T02:00:00Z' }));
  data.jobs = [...copies, ...data.jobs].reverse();
  data.context.runAttempt = '3'; data.run.run_attempt = 3;
  assert.deepEqual(selectFinalizedInputs(data), original);
});

test('newer failed or pending finalizers cannot fall back to successful older artifacts', () => {
  for (const conclusion of ['failure', null]) {
    const data = fixture(), previous = data.jobs.find(job => job.name === 'qualify and attest (linux-arm64)');
    data.jobs.push({ ...previous, id: 999, run_attempt: 2, conclusion, status: conclusion === null ? 'in_progress' : 'completed' });
    assert.throws(() => selectFinalizedInputs(data));
  }
  const data = fixture(); data.jobs.push({ ...data.jobs.at(-1), id: 999 });
  assert.throws(() => selectFinalizedInputs(data), /Ambiguous/);
});

test('each release archive must bind to its exact successful job, source, run and attempt', () => {
  const mutations = [
    data => data.artifacts.pop(), data => data.artifacts.push({ ...data.artifacts.at(-1), id: 999 }),
    data => { data.artifacts.at(-1).name = 'release-darwin-arm64-attempt-2'; },
    data => { data.artifacts.at(-1).workflow_run.id = 43; }, data => { data.artifacts.at(-1).workflow_run.head_sha = 'f'.repeat(40); },
    data => { data.artifacts.at(-1).created_at = '2026-10-30T00:10:59Z'; },
    data => { data.artifacts.at(-1).created_at = '2026-10-30T00:20:01Z'; },
    data => { data.artifacts.at(-1).expired = true; }, data => { data.artifacts.at(-1).digest = 'sha256:bad'; },
    data => { data.artifacts.at(-1).size_in_bytes = 0; }, data => { data.artifacts.at(-1).id = data.artifacts.at(-2).id; },
    data => { data.jobs.at(-1).run_attempt = 3; }, data => { data.run.head_sha = 'f'.repeat(40); },
    data => { data.tag.object.verification.verified = false; }, data => data.jobs.pop(),
  ];
  mutations.forEach((mutate, index) => { const data = fixture(); mutate(data); assert.throws(() => selectFinalizedInputs(data), undefined, `binding ${index}`); });
});

test('all fifteen latest qualifiers must still pass together and precede every selected finalizer', () => {
  const mutations = [
    data => { data.jobs[0].conclusion = 'failure'; }, data => data.jobs.splice(0, 1),
    data => data.jobs.push({ ...data.jobs[0], id: 999, run_attempt: 2, conclusion: 'failure' }),
    data => data.jobs.push({ ...data.jobs[0], id: 999, run_attempt: 2 }),
    data => { for (const job of data.jobs.slice(0, 15)) job.run_attempt = 2; },
    data => { data.jobs.at(-1).started_at = '2026-10-30T00:09:59Z'; },
  ];
  for (const mutate of mutations) { const data = fixture(); mutate(data); assert.throws(() => selectFinalizedInputs(data)); }
});

test('full quality evidence cannot relabel qualification or substitute identities and input metadata', () => {
  const data = fixture(), selected = selectFinalizedInputs(data)[1], env = environment(data.context);
  const mutations = [
    q => { q.full = false; }, q => { q.status = 'failed'; }, q => { q.phase = 'core'; },
    q => { q.identity.github.runAttempt = '2'; }, q => { q.identity.source.commit = 'f'.repeat(40); },
    q => { q.identity.compiler.nativeSha256 = 'f'.repeat(64); }, q => { q.identity.target = 'linux-amd64'; },
    q => { q.finalization.github.finalizationAttempt = '2'; }, q => { q.finalization.jobs[0].id = 999; },
    q => { q.finalization.artifacts[0].digest = 'sha256:' + 'f'.repeat(64); }, q => { q.finalization.source.tagObject = 'f'.repeat(40); },
  ];
  for (const mutate of mutations) { const q = quality(selected); mutate(q); assert.throws(() => validateFinalizedQuality(q, selected, env)); }
});

function directoryFixture(t) {
  const data = fixture(), selected = selectFinalizedInputs(data)[0], root = mkdtempSync(join(tmpdir(), 'gopurs-finalized-test-'));
  t.after(() => rmSync(root, { recursive: true, force: true }));
  const directory = join(root, 'release'); mkdirSync(directory);
  const version = '0.1.0-rc.7', prefix = `gopurs-${version}-${selected.target}`;
  const files = { [prefix]: 'native', [`gopurs-${version}.tgz`]: 'npm', [`${prefix}-release-notes.md`]: 'notes',
    [`${prefix}-bootstrap-checksums.txt`]: `${sha('javascript')}  bin/gopurs.js\n`,
    [`${prefix}-quality.json`]: JSON.stringify(quality(selected)),
    [`${prefix}-provenance.sigstore.json`]: '{}', [`${prefix}-sbom.sigstore.json`]: '{}' };
  files[`${prefix}-build.json`] = JSON.stringify({ schema: 1, version, target: selected.target, preview: false,
    source: { commit: data.context.commit, tag: data.context.tag, dirty: false },
    artifacts: [prefix, `${prefix}-bootstrap-checksums.txt`].map(name => ({ name, sha256: sha(files[name]) })) });
  const checksums = Object.entries(files).filter(([name]) => !name.endsWith('.sigstore.json')).map(([name, bytes]) => `${sha(bytes)}  ${name}\n`).join('');
  files[`${prefix}-SHA256SUMS`] = checksums;
  for (const [name, bytes] of Object.entries(files)) writeFileSync(join(directory, name), bytes);
  return { data, selected, directory, prefix, checksums };
}

test('downloaded file checksums, build metadata and qualified compiler bytes are checked', t => {
  const { data, selected, directory, prefix, checksums } = directoryFixture(t), env = environment(data.context);
  assert.equal(verifyFinalizedDirectory(directory, selected, env).qualificationAttempt, '1');
  writeFileSync(join(directory, prefix), 'wrong');
  assert.throws(() => verifyFinalizedDirectory(directory, selected, env), /checksum differs/);
  writeFileSync(join(directory, prefix), 'native');
  writeFileSync(join(directory, `${prefix}-SHA256SUMS`), checksums + checksums.split('\n')[0] + '\n');
  assert.throws(() => verifyFinalizedDirectory(directory, selected, env), /checksum inventory/);
  writeFileSync(join(directory, `${prefix}-SHA256SUMS`), checksums);
  writeFileSync(join(directory, `${prefix}-unexpected`), 'extra');
  assert.throws(() => verifyFinalizedDirectory(directory, selected, env), /Incomplete checksum inventory/);
  rmSync(join(directory, `${prefix}-unexpected`));
  const nativeHash = sha('native'), wrongNative = 'tampered';
  writeFileSync(join(directory, prefix), wrongNative);
  writeFileSync(join(directory, `${prefix}-SHA256SUMS`), checksums.replace(`${nativeHash}  ${prefix}\n`, `${sha(wrongNative)}  ${prefix}\n`));
  assert.throws(() => verifyFinalizedDirectory(directory, selected, env), /Qualified native compiler differs/);
});
