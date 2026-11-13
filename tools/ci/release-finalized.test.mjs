import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { selectFinalizedInputs, validateFinalizedQuality, verifyFinalizedDirectory } from './release-finalized.mjs';
import { recoveryBase, recoveryHash, validateRecoveryPackage, validateRecoverySmokeCommands } from './release-recovery.mjs';
import { qualityInventory } from '../release/release-quality-report.mjs';
import { fileURLToPath } from 'node:url';

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
    .map((name, index) => artifact(name + '-attempt-1', 100 + index, '2026-10-30T00:05:00Z'));
  targets.forEach((target, index) => {
    const attempt = index === 0 ? 2 : 1, hour = attempt === 2 ? '01' : '00';
    jobs.push({ id: 50 + index, name: `qualify and attest (${target})`, run_id: 42, run_attempt: attempt, head_sha: context.commit,
      status: 'completed', conclusion: 'success', started_at: `2026-10-30T${hour}:11:00Z`, completed_at: `2026-10-30T${hour}:20:00Z` });
    artifacts.push(artifact(`fixture-environment-${target}-attempt-1`, 300 + index, '2026-10-30T00:05:00Z'));
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

test('final signing binds the new package smoke and unchanged RC20 producer identities', () => {
  const convert = (data, context) => {
    data.context = context;
    Object.assign(data.run, { id: Number(context.runId), run_attempt: Number(context.runAttempt), head_sha: context.commit,
      head_branch: context.tag, repository: { full_name: context.repository }, head_repository: { full_name: context.repository } });
    data.tag.ref.ref = 'refs/tags/' + context.tag;
    data.tag.ref.object.sha = context.tag === recoveryBase.tag ? recoveryBase.tagObject : 'c'.repeat(40);
    Object.assign(data.tag.object, { sha: data.tag.ref.object.sha, tag: context.tag, object: { type: 'commit', sha: context.commit } });
    for (const job of data.jobs) Object.assign(job, { run_id: Number(context.runId), head_sha: context.commit });
    for (const artifact of data.artifacts) artifact.workflow_run = { id: Number(context.runId), head_sha: context.commit };
    return data;
  };
  const base = convert(fixture(), { ...recoveryBase, runAttempt: '2', target: targets[0] });
  base.run.status = 'completed';
  for (const target of ['linux-amd64', 'linux-arm64', 'darwin-arm64']) {
    const artifact = base.artifacts.find(row => row.name === `candidate-${target}-attempt-1`);
    if (artifact) base.artifacts.push({ ...artifact, id: artifact.id + 50000, name: `quality-${target}-attempt-1` });
  }
  const current = convert(fixture(), { repository: recoveryBase.repository, commit: 'd'.repeat(40), tag: 'v0.1.1', runId: '999', runAttempt: '2', target: targets[0] });
  current.jobs = current.jobs.filter(job => job.name.endsWith('shard 0)') || job.name.startsWith('qualify and attest')).map(job => ({ ...job, id: job.id + 10000 }));
  current.artifacts = current.artifacts.filter(artifact => artifact.name.startsWith('release-') || /fixtures-.*-0-attempt/.test(artifact.name)).map(artifact => ({ ...artifact, id: artifact.id + 10000 }));
  targets.forEach((target, index) => {
    current.jobs.push({ ...current.jobs[0], id: 20000 + index, name: `repackage and smoke (${target})` });
    current.artifacts.push({ ...current.artifacts[0], id: 21000 + index, name: `repackaged-${target}-attempt-1` });
  });
  const manifest = { schema: 1, enabled: true, base: recoveryBase, targetTag: 'v0.1.1', fixtureSources: [], goldens: [{ path: 'tests/passing-snapshots/1110.go', fixture: 'tests/passing/1110.purs', beforeSha256: 'a'.repeat(64), afterSha256: 'b'.repeat(64) }] };
  const recovery = { base, manifest, proof: { manifestSha256: recoveryHash(JSON.stringify(manifest)), changes: [] }, inventory: qualityInventory(fileURLToPath(new URL('../../', import.meta.url))) };
  const selected = selectFinalizedInputs({ ...current, recovery });
  for (const item of selected) {
    const report = quality(item);
    for (const phase of report.phases) {
      const name = phase.phase === 'core' ? `build and test (${item.target})` : `fixtures (${item.target}, shard ${phase.shard.index})`;
      const producer = item.finalization.jobs.find(job => job.name === name);
      phase.identity.source.commit = producer.commit;
      phase.identity.github.runId = producer.runId;
    }
    report.identity = structuredClone(report.phases[0].identity);
    report.finalization.qualificationIdentity = report.identity;
    report.publicationSource = item.finalization.source;
    report.gates = Array.from({ length: 384 }, () => ({})); report.phases[0].gates = 222;
    report.fixtures = { selectedPerProfile: 401, passed: { release: 401, dev: 401 } };
    validateFinalizedQuality(report, item, environment(current.context));
    const context = report.finalization, sourceManifest = { name: 'gopurs', version: '0.1.1' };
    const contents = { build: { schema: 1, preview: false, version: '0.1.1', target: item.target,
      source: { commit: current.context.commit, tag: 'v0.1.1', dirty: false },
      compilerOrigin: { base: recoveryBase, identity: report.identity, artifact: item.finalization.artifacts[0] } },
      nativeSha256: report.identity.compiler.nativeSha256, javascriptSha256: report.identity.compiler.javascriptSha256,
      manifest: sourceManifest, sourceManifest, packageSha256: 'e'.repeat(64),
      smoke: { version: '0.1.1', packageSHA256: 'e'.repeat(64), commands: ['javascript', 'native'].flatMap(host =>
        [0, 0, 0, 1, 0].map(status => ({ status, cwd: `/tmp/smoke/${host}`, args: ['--main', 'Main'],
          command: '/tmp/smoke/node_modules/' + (host === 'javascript' ? '.bin/gopurs' : 'gopurs/bin/gopurs') }))),
        tests: ['javascript', 'native'].map(host => ({ host, passed: true, stages: ['official-javascript', 'cold', 'warm', 'signature-change', 'invalid-signature', 'recovery'] })) } };
    validateRecoveryPackage(context, contents);
    assert.throws(() => validateRecoveryPackage(context, { ...contents, nativeSha256: 'a'.repeat(64) }), /bytes changed/);
    assert.throws(() => validateRecoveryPackage(context, { ...contents, manifest: { ...sourceManifest, version: '0.1.0-rc.20' } }));
    assert.throws(() => validateRecoveryPackage(context, { ...contents, packageSha256: 'f'.repeat(64) }), /different package/);
    assert.throws(() => validateRecoveryPackage(context, { ...contents, smoke: { ...contents.smoke, tests: contents.smoke.tests.slice(0, 1) } }));
    for (const change of [commands => { commands[3].status = 0; }, commands => { commands[2].status = 1; },
      commands => { commands[3].status = null; }, commands => { commands.splice(5); },
      commands => { commands.push({ command: 'go', args: ['test'], cwd: '/tmp/smoke/native', status: 1 }); },
      commands => { commands[3].command = 'unrelated-failure'; }]) {
      const commands = structuredClone(contents.smoke.commands); change(commands);
      assert.throws(() => validateRecoverySmokeCommands(commands), change.toString());
    }
  }
  current.jobs.find(job => job.name === 'repackage and smoke (linux-amd64)').conclusion = 'failure';
  assert.throws(() => selectFinalizedInputs({ ...current, recovery }), /did not pass/);
});
function quality(selected) {
  const identity = { source: { commit: selected.finalization.source.commit, dirty: false }, target: selected.target,
    compiler: { javascriptSha256: sha('javascript'), nativeSha256: sha('native') },
    environment: { profile: '/nix/store/' + 'a'.repeat(32) + '-gopurs-env', profileSha256: 'b'.repeat(64), closureSha256: 'c'.repeat(64), shell: '/nix/store/' + 'd'.repeat(32) + '-bash/bin/bash' },
    github: { repository: selected.finalization.github.repository, runId: selected.finalization.github.runId,
      runAttempt: selected.finalization.jobs.find(job => job.name === 'build and test (' + selected.target + ')').attempt } };
  const phases = [{ phase: 'core', shard: null, gates: 1, identity: structuredClone(identity) },
    ...[0, 1, 2, 3].map(index => ({ phase: 'fixtures', shard: { index, count: 4 }, gates: 1, identity: { ...structuredClone(identity),
      github: { ...identity.github, runAttempt: selected.finalization.artifacts[index + 1].attempt } } }))];
  return { phases, schema: 1, requestedFull: true, full: true, phase: 'all', shard: null, status: 'passed', identity,
    finalization: { ...structuredClone(selected.finalization), qualificationIdentity: structuredClone(identity) } };
}

test('draft retry accepts mixed successful finalization attempts while retaining qualification one', () => {
  const data = fixture(), selected = selectFinalizedInputs(data);
  assert.deepEqual(selected.map(item => item.job.attempt), ['2', '1', '1']);
  assert.deepEqual(selected.map(item => item.artifact.id), [200, 201, 202]);
  for (const status of ['queued', 'completed']) {
    assert.deepEqual(selectFinalizedInputs({ ...data, run: { ...data.run, status } }), selected);
  }
  for (const item of selected) {
    assert.ok(item.finalization.jobs.every(job => job.attempt === '1'));
    assert.equal(item.finalization.jobs.length, 15);
    assert.equal(validateFinalizedQuality(quality(item), item, environment(data.context)).identity.github.runAttempt, '1');
  }
  data.context.runAttempt = '3'; data.run.run_attempt = 3;
  assert.deepEqual(selectFinalizedInputs(data).map(item => item.artifact.id), [200, 201, 202]);
});

test('a repaired fixture shard retains original build and eleven other shard producers through signing', () => {
  const data = fixture(), shard = data.jobs.find(job => job.name === 'fixtures (linux-amd64, shard 2)');
  data.jobs.push({ ...shard, id: 999, run_attempt: 2, started_at: '2026-10-30T01:00:00Z', completed_at: '2026-10-30T01:10:00Z' });
  const artifact = data.artifacts.find(value => value.name === 'fixtures-linux-amd64-2-attempt-1');
  data.artifacts.push({ ...artifact, id: 999, name: 'fixtures-linux-amd64-2-attempt-2', created_at: '2026-10-30T01:05:00Z' });
  for (const job of data.jobs.filter(value => value.name.startsWith('qualify and attest ('))) {
    job.run_attempt = 2; job.started_at = '2026-10-30T01:11:00Z'; job.completed_at = '2026-10-30T01:20:00Z';
  }
  for (const release of data.artifacts.filter(value => value.name.startsWith('release-'))) {
    release.name = release.name.replace('-attempt-1', '-attempt-2'); release.created_at = '2026-10-30T01:19:00Z';
  }
  const selected = selectFinalizedInputs(data);
  assert.deepEqual(selected[0].finalization.artifacts.map(value => value.attempt), ['1', '1', '1', '2', '1']);
  for (const target of selected) validateFinalizedQuality(quality(target), target, environment(data.context));
  const report = quality(selected[0]); report.phases[3].identity.github.runAttempt = '1';
  assert.throws(() => validateFinalizedQuality(report, selected[0], environment(data.context)), /actual artifact producer/);
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
  for (const runStatus of ['in_progress', 'queued']) {
    for (const [status, conclusion] of [['completed', 'failure'], ['in_progress', null], ['queued', null]]) {
      const data = fixture(), previous = data.jobs.find(job => job.name === 'qualify and attest (linux-arm64)');
      data.run.status = runStatus;
      data.jobs.push({ ...previous, id: 999, run_attempt: 2, conclusion, status });
      assert.throws(() => selectFinalizedInputs(data));
    }
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
    data => { data.run.id = 43; },
    data => { data.tag.object.verification.verified = false; }, data => data.jobs.pop(),
  ];
  for (const status of ['in_progress', 'queued']) {
    mutations.forEach((mutate, index) => {
      const data = fixture(); data.run.status = status; mutate(data);
      assert.throws(() => selectFinalizedInputs(data), undefined, `binding ${index} during ${status}`);
    });
  }
});

test('all fifteen latest qualifiers must still pass together and precede every selected finalizer', () => {
  const mutations = [
    data => { data.jobs[0].conclusion = 'failure'; }, data => data.jobs.splice(0, 1),
    data => { data.jobs[0].status = 'queued'; data.jobs[0].conclusion = null; },
    data => data.jobs.push({ ...data.jobs[0], id: 999, run_attempt: 2, status: 'queued', conclusion: null }),
    data => data.jobs.push({ ...data.jobs[0], id: 999, run_attempt: 2, conclusion: 'failure' }),
    data => data.jobs.push({ ...data.jobs[0], id: 999, run_attempt: 2 }),
    data => { for (const job of data.jobs.slice(0, 15)) job.run_attempt = 2; },
    data => { data.jobs.at(-1).started_at = '2026-10-30T00:09:59Z'; },
  ];
  for (const status of ['in_progress', 'queued']) {
    for (const mutate of mutations) {
      const data = fixture(); data.run.status = status; mutate(data);
      assert.throws(() => selectFinalizedInputs(data));
    }
  }
});

test('full quality evidence cannot relabel qualification or substitute identities and input metadata', () => {
  const data = fixture(), selected = selectFinalizedInputs(data)[1], env = environment(data.context);
  const mutations = [
    q => { q.phases[0].shard = {}; }, q => { q.phases[1].shard.count = 5; },
    q => { q.identity.environment = {}; q.finalization.qualificationIdentity.environment = {}; },
    q => { q.phases[1].identity.github.runAttempt = '2'; }, q => { q.phases[1].identity.environment.closureSha256 = 'f'.repeat(64); },
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
  assert.ok(verifyFinalizedDirectory(directory, selected, env).producerAttempts.every(job => job.attempt === '1'));
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
