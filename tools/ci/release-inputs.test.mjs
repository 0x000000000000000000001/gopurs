import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { existsSync, mkdtempSync, readFileSync, rmSync, statSync, symlinkSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { extractArtifact, fileSha256, flattenPages, githubContext, selectLatestJobs, selectQualifiedInputs, selectFixtureInputs, selectPreviewInputs, validateReportProducers, validateEnvironmentIdentity, validateFinalizationContext, verifyArtifactFile } from './release-inputs.mjs';

function fixture() {
  const context = { repository: 'owner/gopurs', commit: 'a'.repeat(40), tag: 'v0.1.0-rc.7', runId: '42', runAttempt: '2', target: 'linux-amd64' };
  const jobs = ['linux-amd64', 'linux-arm64', 'darwin-arm64'].flatMap(target => [
    'build and test (' + target + ')', ...[0, 1, 2, 3].map(shard => 'fixtures (' + target + ', shard ' + shard + ')'),
  ]).map((name, index) => ({ id: index + 1, name, run_id: 42, run_attempt: 1, head_sha: context.commit,
    status: 'completed', conclusion: 'success', started_at: '2026-10-30T00:00:00Z', completed_at: '2026-10-30T00:10:00Z' }));
  const artifacts = ['candidate-linux-amd64', ...[0, 1, 2, 3].map(shard => 'fixtures-linux-amd64-' + shard)].map((name, index) => ({
    id: 100 + index, name: name + '-attempt-1', size_in_bytes: 128, digest: 'sha256:' + 'b'.repeat(64), expired: false,
    created_at: '2026-10-30T00:05:00Z', updated_at: '2026-10-31T00:00:00Z', workflow_run: { id: 42, head_sha: context.commit },
  }));
  artifacts.push({ ...artifacts[0], id: 9999, name: 'fixture-environment-linux-amd64-attempt-1' });
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
    GOPURS_RELEASE_TARGET: context.target, GITHUB_WORKFLOW_SHA: context.commit,
    GITHUB_WORKFLOW_REF: context.repository + '/.github/workflows/ci.yml@refs/tags/' + context.tag };
}

function identity(selected) {
  return { source: { commit: selected.source.commit, dirty: false }, target: selected.target,
    compiler: { javascriptSha256: 'd'.repeat(64), nativeSha256: 'e'.repeat(64) },
    environment: { profile: '/nix/store/' + 'a'.repeat(32) + '-gopurs-env', profileSha256: 'b'.repeat(64), closureSha256: 'c'.repeat(64), shell: '/nix/store/' + 'd'.repeat(32) + '-bash/bin/bash' },
    github: { repository: selected.github.repository, runId: selected.github.runId, runAttempt: selected.jobs.find(job => job.name === 'build and test (' + selected.target + ')').attempt } };
}

test('finalization attempt two reuses the exact successful qualification from attempt one', () => {
  const data = fixture(), selected = selectQualifiedInputs(data), core = identity(selected);
  assert.ok(selected.jobs.every(job => job.attempt === '1'));
  assert.equal(selected.github.finalizationAttempt, '2');
  assert.equal(selected.jobs.length, 15); assert.equal(selected.artifacts.length, 5);
  assert.deepEqual(selected.artifacts.map(artifact => artifact.producerJobId), [1, 2, 3, 4, 5]);
  for (const status of ['queued', 'completed']) {
    assert.deepEqual(selectQualifiedInputs({ ...data, run: { ...data.run, status } }), selected);
  }
  selected.qualificationIdentity = structuredClone(core);
  assert.equal(validateFinalizationContext(selected, core, environment(data.context)), selected);
  assert.deepEqual(githubContext(environment(data.context)), data.context);
  // Updated metadata is harmless; only creation binds the immutable archive to its producer.
  assert.equal(selectQualifiedInputs(data).artifacts[0].id, 100);
});

test('copied successful jobs retain the original identity only with matching execution evidence', () => {
  const data = fixture(), original = { ...data.jobs[0], created_at: '2026-10-29T23:59:59Z',
    runner_id: 123, runner_name: 'GitHub Actions 123',
    steps: [{ name: 'Build', number: 1, status: 'completed', conclusion: 'success',
      started_at: data.jobs[0].started_at, completed_at: data.jobs[0].completed_at }] };
  const copy = { ...structuredClone(original), id: 1000, run_attempt: 2, created_at: '2026-10-30T01:00:00Z' };
  assert.deepEqual(selectLatestJobs([copy, original], data.context).get(original.name), original);
  for (const mutate of [
    job => { job.steps[0].name = 'Different execution'; },
    job => { job.runner_id++; }, job => { job.runner_name = 'different'; },
    job => { job.completed_at = '2026-10-30T00:11:00Z'; },
    job => { job.steps = []; }, job => { delete job.steps; },
  ]) {
    const changed = structuredClone(copy); mutate(changed);
    assert.throws(() => selectLatestJobs([original, changed], data.context), /copied job/i);
  }
  assert.throws(() => selectLatestJobs([copy], data.context), /copied job/i);
  const failed = { ...structuredClone(copy), status: 'completed', conclusion: 'failure',
    created_at: '2026-10-30T00:59:59Z', started_at: '2026-10-30T01:00:00Z', completed_at: '2026-10-30T01:10:00Z' };
  assert.equal(selectLatestJobs([original, failed], data.context).get(original.name).id, failed.id);
  const laterCopy = { ...copy, id: 2000, run_attempt: 3, created_at: '2026-10-30T02:00:00Z' };
  assert.throws(() => selectLatestJobs([original, failed, laterCopy], { ...data.context, runAttempt: '3' }), /copied job/i);
});

test('latest actual qualification jobs must all succeed without hiding newer failures', () => {
  const mutations = [
    data => data.jobs.push({ ...data.jobs[0], id: 1000, run_attempt: 2, conclusion: 'failure' }),
    data => data.jobs.push({ ...data.jobs[0], id: 1000, run_attempt: 2, status: 'in_progress', conclusion: null }),
    data => data.jobs.push({ ...data.jobs[0], id: 1000, run_attempt: 2, status: 'queued', conclusion: null }),
    data => { data.jobs[0].status = 'queued'; data.jobs[0].conclusion = null; },
    data => { data.jobs[0].conclusion = 'failure'; },
    data => data.jobs.push({ ...data.jobs[0], id: 1000, run_attempt: 2 }),
    data => data.jobs.push({ ...data.jobs[0], id: 1000 }),
    data => data.jobs.pop(),
    data => data.jobs.push({ ...data.jobs[0], id: 1000, name: 'fixtures (linux-amd64, shard 4)' }),
    data => { data.jobs[0].run_attempt = 3; },
    data => { data.jobs[0].run_id = 43; },
    data => { data.jobs[0].head_sha = 'f'.repeat(40); },
  ];
  for (const status of ['in_progress', 'queued']) {
    mutations.forEach((mutate, index) => {
      const data = fixture(); data.run.status = status; mutate(data);
      assert.throws(() => selectQualifiedInputs(data), undefined, 'job mutation ' + index + ' during ' + status);
    });
  }
  const newer = fixture();
  newer.jobs = [...newer.jobs.map(job => ({ ...job, conclusion: 'failure' })),
    ...newer.jobs.map(job => ({ ...job, id: job.id + 1000, run_attempt: 2, started_at: '2026-10-30T01:00:00Z', completed_at: '2026-10-30T01:10:00Z' }))];
  assert.throws(() => selectQualifiedInputs(newer), /Expected one artifact/);
  newer.artifacts.forEach(artifact => { artifact.name = artifact.name.replace('-attempt-1', '-attempt-2'); });
  assert.throws(() => selectQualifiedInputs(newer), /Artifact was not created/);
  newer.artifacts.forEach(artifact => { artifact.created_at = '2026-10-30T01:05:00Z'; });
  assert.ok(selectQualifiedInputs(newer).jobs.every(job => job.attempt === '2'));
});

test('a retry selects only the repaired shard from its real new attempt and retains original producers', () => {
  const data = fixture(), original = data.jobs[3];
  data.jobs.push({ ...original, id: 999, run_attempt: 2, started_at: '2026-10-30T01:00:00Z', completed_at: '2026-10-30T01:10:00Z' });
  data.artifacts.push({ ...data.artifacts[3], id: 999, name: 'fixtures-linux-amd64-2-attempt-2', created_at: '2026-10-30T01:05:00Z' });
  const selected = selectQualifiedInputs(data), core = identity(selected);
  assert.deepEqual(selected.artifacts.map(value => value.attempt), ['1', '1', '1', '2', '1']);
  assert.equal(selected.artifacts[3].producerJobId, 999);
  selected.qualificationIdentity = structuredClone(core);
  validateFinalizationContext(selected, core, environment(data.context));
  const reports = [0, 1, 2, 3].map(index => ({ phase: 'fixtures', shard: { index, count: 4 },
    identity: { ...core, github: { ...core.github, runAttempt: index === 2 ? '2' : '1' } } }));
  validateReportProducers(selected, { phase: 'core', identity: core }, reports);
  reports[2].identity.github.runAttempt = '1';
  assert.throws(() => validateReportProducers(selected, { phase: 'core', identity: core }, reports), /actual artifact producer/);
  data.jobs.at(-1).conclusion = 'failure';
  assert.throws(() => selectQualifiedInputs(data), /latest qualification job/);
});

test('fixture inputs select exact candidate and environment from the original build, never a copied rerun label', () => {
  const data = fixture();
  const selected = selectFixtureInputs(data);
  assert.deepEqual(selected.artifacts.map(value => value.logicalName), ['candidate-linux-amd64', 'fixture-environment-linux-amd64']);
  assert.deepEqual(selected.artifacts.map(value => value.attempt), ['1', '1']);
  assert.equal(selected.jobs.length, 1);
  // GitHub may report the aggregate run queued while this shard imports a completed build.
  data.run.status = 'queued';
  data.jobs = data.jobs.map((job, index) => index === 0 ? job : {
    ...job, status: 'queued', conclusion: null, started_at: null, completed_at: null,
  });
  assert.deepEqual(selectFixtureInputs(data), selected);
  for (const mutate of [
    changed => { changed.jobs[0].status = 'queued'; changed.jobs[0].conclusion = null; },
    changed => { changed.jobs[0].conclusion = 'failure'; },
    changed => changed.jobs.push({ ...changed.jobs[0], id: 1000, run_attempt: 2, status: 'queued', conclusion: null }),
    changed => changed.jobs.push({ ...changed.jobs[0], id: 1000, run_attempt: 2, conclusion: 'failure' }),
    changed => { changed.artifacts[0].workflow_run.id = 43; },
    changed => { changed.artifacts[0].workflow_run.head_sha = 'f'.repeat(40); },
    changed => { changed.artifacts[0].digest = 'sha256:bad'; },
    changed => { changed.artifacts.at(-1).expired = true; },
  ]) {
    const changed = structuredClone(data); mutate(changed);
    assert.throws(() => selectFixtureInputs(changed));
  }
  data.artifacts.find(a => a.name.startsWith('fixture-environment-')).name = 'fixture-environment-linux-amd64-attempt-2';
  assert.throws(() => selectFixtureInputs(data), /Expected one artifact/);
});

test('unsigned preview validates its branch workflow_dispatch identity without asserting tag qualification', () => {
  const data = fixture(); data.context.ref = 'refs/heads/edge'; delete data.context.tag;
  data.run.event = 'workflow_dispatch'; data.run.head_branch = 'edge'; delete data.tag;
  const selected = selectPreviewInputs(data);
  assert.deepEqual(selected.source, { commit: data.context.commit, ref: 'refs/heads/edge' });
  data.run.status = 'queued';
  assert.deepEqual(selectPreviewInputs(data), selected);
  assert.equal(selectFixtureInputs(data, true).jobs.length, 1);
  data.run.event = 'pull_request';
  assert.throws(() => selectPreviewInputs(data));
});

test('run, workflow, source and verified annotated tag identities cannot be substituted', () => {
  const mutations = [
    data => { data.run.head_sha = 'f'.repeat(40); }, data => { data.run.head_branch = 'main'; },
    data => { data.run.repository.full_name = 'other/gopurs'; }, data => { data.run.head_repository.full_name = 'other/gopurs'; },
    data => { data.run.event = 'workflow_dispatch'; }, data => { data.run.path = '.github/workflows/other.yml'; },
    data => { data.run.id = 43; }, data => { data.run.run_attempt = 1; },
    data => { data.tag.ref.object.type = 'commit'; }, data => { data.tag.object.verification.verified = false; },
    data => { data.tag.object.object.sha = 'f'.repeat(40); }, data => { data.tag.object.sha = 'f'.repeat(40); },
    data => { data.tag.object.tag = 'v0.1.0-rc.6'; }, data => { data.tag.ref.ref = 'refs/heads/main'; },
  ];
  for (const select of [selectQualifiedInputs, selectFixtureInputs]) {
    for (const status of ['in_progress', 'queued']) {
      mutations.forEach((mutate, index) => {
        const data = fixture(); data.run.status = status; mutate(data);
        assert.throws(() => select(data), undefined, 'source mutation ' + index + ' during ' + status);
      });
    }
    for (const status of ['requested', 'waiting', 'pending', 'unknown', undefined, null]) {
      const data = fixture(); data.run.status = status;
      assert.throws(() => select(data), undefined, 'unsupported run status ' + status);
    }
  }
});

test('artifact selection rejects omissions, ambiguity and wrong producer, run or digest', () => {
  const mutations = [
    data => data.artifacts.pop(), data => data.artifacts.splice(4, 1), data => data.artifacts.push({ ...data.artifacts[0], id: 999 }),
    data => { data.artifacts[4].name = 'fixtures-linux-amd64-4'; }, data => { data.artifacts[0].expired = true; },
    data => { data.artifacts[0].workflow_run.id = 43; }, data => { data.artifacts[0].workflow_run.head_sha = 'f'.repeat(40); },
    data => { data.artifacts[0].digest = 'sha256:short'; }, data => { data.artifacts[0].size_in_bytes = 0; },
    data => { data.artifacts[0].size_in_bytes = 1.5; }, data => { data.artifacts[0].id = -1; },
    data => { data.artifacts[1].id = data.artifacts[0].id; }, data => { delete data.artifacts[0].created_at; },
    data => { data.artifacts[0].created_at = '2026-10-29T23:59:59Z'; }, data => { data.artifacts[0].created_at = '2026-10-30T00:10:01Z'; },
  ];
  for (const status of ['in_progress', 'queued']) {
    mutations.forEach((mutate, index) => {
      const data = fixture(); data.run.status = status; mutate(data);
      assert.throws(() => selectQualifiedInputs(data), undefined, 'artifact mutation ' + index + ' during ' + status);
    });
  }
});

test('merged reports retain qualification identity while finalization remains bound to the current attempt', () => {
  const mutations = [
    (selected, core) => { core.environment = {}; selected.qualificationIdentity.environment = {}; },
    (selected, core) => { core.environment.profileSha256 = 'bad'; selected.qualificationIdentity.environment.profileSha256 = 'bad'; },
    (selected, core) => { core.compiler.nativeSha256 = 'f'.repeat(64); },
    (selected, core) => { core.compiler.javascriptSha256 = 'f'.repeat(64); },
    (selected, core) => { core.source.dirty = true; }, (selected, core) => { core.source.commit = 'f'.repeat(40); },
    (selected, core) => { core.target = 'darwin-arm64'; }, (selected, core) => { core.github.runAttempt = '2'; },
    selected => { selected.github.finalizationAttempt = '1'; }, selected => { selected.jobs[0].attempt = '2'; },
    selected => { selected.environmentArtifact.attempt = '2'; }, selected => { selected.environmentArtifact.producerJobId = 2; },
    selected => { selected.artifacts[0].producerJobId = 2; }, selected => { selected.source.tag = 'v0.1.0-rc.6'; },
    (selected, core, env) => { env.GITHUB_WORKFLOW_SHA = 'f'.repeat(40); },
    (selected, core, env) => { env.GITHUB_REF = 'refs/heads/main'; },
    (selected, core, env) => { env.GITHUB_REPOSITORY = 'other/gopurs'; },
    (selected, core, env) => { env.GITHUB_RUN_ID = '43'; },
  ];
  mutations.forEach((mutate, index) => {
    const data = fixture(), selected = selectQualifiedInputs(data), core = identity(selected), env = environment(data.context);
    selected.qualificationIdentity = structuredClone(core); mutate(selected, core, env);
    assert.throws(() => validateFinalizationContext(selected, core, env), undefined, 'context mutation ' + index);
  });
});

test('paginated API responses must be complete and internally consistent', () => {
  assert.deepEqual(flattenPages([{ total_count: 2, jobs: [1] }, { total_count: 2, jobs: [2] }], 'jobs'), [1, 2]);
  assert.throws(() => flattenPages([{ total_count: 2, jobs: [1] }], 'jobs'));
  assert.throws(() => flattenPages([{ total_count: 2, jobs: [1] }, { total_count: 3, jobs: [2] }], 'jobs'));
});

const zipPython = String.raw`
import json, struct, sys, warnings, zipfile
path, entries, encrypted = sys.argv[1], json.loads(sys.argv[2]), sys.argv[3] == 'true'
with warnings.catch_warnings():
    warnings.simplefilter('ignore', UserWarning)
    with zipfile.ZipFile(path, 'w') as output:
        for entry in entries:
            info = zipfile.ZipInfo(entry['name'])
            info.create_system = 3
            info.external_attr = entry.get('mode', 0o100644) << 16
            output.writestr(info, entry.get('data', 'ok'))
if encrypted:
    data = bytearray(open(path, 'rb').read())
    for marker, offset in [(b'PK\x03\x04', 6), (b'PK\x01\x02', 8)]:
        index = data.index(marker) + offset
        struct.pack_into('<H', data, index, struct.unpack_from('<H', data, index)[0] | 1)
    open(path, 'wb').write(data)
`;
function archive(root, entries, encrypted = false) {
  const path = join(root, 'archive.zip');
  execFileSync('python3', ['-c', zipPython, path, JSON.stringify(entries), String(encrypted)]);
  return { path, artifact: { id: 123, size: statSync(path).size, digest: 'sha256:' + fileSha256(path) } };
}

test('ZIP bytes are verified before extraction into a fresh destination', t => {
  const root = mkdtempSync(join(tmpdir(), 'gopurs-release-inputs-'));
  t.after(() => rmSync(root, { recursive: true, force: true }));
  const { path, artifact } = archive(root, [{ name: 'logs/quality.json', data: '{}' }]), destination = join(root, 'unpacked');
  assert.throws(() => verifyArtifactFile({ ...artifact, digest: 'sha256:' + '0'.repeat(64) }, path));
  assert.throws(() => extractArtifact({ ...artifact, size: artifact.size + 1 }, path, destination));
  assert.equal(existsSync(destination), false);
  assert.equal(extractArtifact(artifact, path, destination), 2);
  assert.equal(readFileSync(join(destination, 'logs/quality.json'), 'utf8'), '{}');
  assert.throws(() => extractArtifact(artifact, path, destination));
  const link = join(root, 'linked'); symlinkSync(destination, link);
  assert.throws(() => extractArtifact(artifact, path, link));
});

test('hostile ZIP members, encryption and expansion over budget are rejected before creating files', t => {
  const root = mkdtempSync(join(tmpdir(), 'gopurs-release-inputs-hostile-'));
  t.after(() => rmSync(root, { recursive: true, force: true }));
  const cases = [
    [{ name: '/absolute' }], [{ name: '../escape' }], [{ name: 'nested/../escape' }], [{ name: 'a\\escape' }], [{ name: 'C:/escape' }],
    [{ name: 'a' }, { name: 'a' }], [{ name: 'a' }, { name: 'A' }], [{ name: 'a', mode: 0o120777 }], [{ name: 'a', mode: 0o020644 }],
    [{ name: 'a' }, { name: 'a/file' }], [{ name: 'a//' }],
  ];
  cases.forEach((entries, index) => {
    const { path, artifact } = archive(root, entries), destination = join(root, 'case-' + index);
    assert.throws(() => extractArtifact(artifact, path, destination), undefined, 'ZIP mutation ' + index);
    assert.equal(existsSync(destination), false);
  });
  const encrypted = archive(root, [{ name: 'encrypted' }], true);
  assert.throws(() => extractArtifact(encrypted.artifact, encrypted.path, join(root, 'encrypted')));
  const oversized = archive(root, [{ name: 'large', data: 'ab' }]);
  assert.throws(() => extractArtifact(oversized.artifact, oversized.path, join(root, 'oversized'), 1));
  assert.equal(existsSync(join(root, 'oversized')), false);
});


test('qualification environments require the exact retained Nix bash as well as profile and closure', () => {
  const value = identity(selectQualifiedInputs(fixture())).environment;
  validateEnvironmentIdentity(value);
  for (const shell of [undefined, '/bin/bash', '/nix/store/' + 'd'.repeat(32) + '-bash/bin/sh',
    '/nix/store/' + 'd'.repeat(32) + '-bash/bin/../bin/bash',
    '/nix/store/' + 'e'.repeat(32) + '-bash/bin/bash']) {
    const changed = { ...value, shell };
    if (shell === undefined) delete changed.shell;
    assert.throws(() => validateEnvironmentIdentity(changed), undefined, String(shell));
  }
});
