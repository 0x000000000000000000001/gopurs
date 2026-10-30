import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { existsSync, mkdtempSync, readFileSync, rmSync, statSync, symlinkSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { extractArtifact, fileSha256, flattenPages, githubContext, selectLatestJobs, selectQualifiedInputs, validateFinalizationContext, verifyArtifactFile } from './release-inputs.mjs';

function fixture() {
  const context = { repository: 'owner/gopurs', commit: 'a'.repeat(40), tag: 'v0.1.0-rc.7', runId: '42', runAttempt: '2', target: 'linux-amd64' };
  const jobs = ['linux-amd64', 'linux-arm64', 'darwin-arm64'].flatMap(target => [
    'build and test (' + target + ')', ...[0, 1, 2, 3].map(shard => 'fixtures (' + target + ', shard ' + shard + ')'),
  ]).map((name, index) => ({ id: index + 1, name, run_id: 42, run_attempt: 1, head_sha: context.commit,
    status: 'completed', conclusion: 'success', started_at: '2026-10-30T00:00:00Z', completed_at: '2026-10-30T00:10:00Z' }));
  const artifacts = ['candidate-linux-amd64', ...[0, 1, 2, 3].map(shard => 'fixtures-linux-amd64-' + shard)].map((name, index) => ({
    id: 100 + index, name, size_in_bytes: 128, digest: 'sha256:' + 'b'.repeat(64), expired: false,
    created_at: '2026-10-30T00:05:00Z', updated_at: '2026-10-31T00:00:00Z', workflow_run: { id: 42, head_sha: context.commit },
  }));
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
    github: { repository: selected.github.repository, runId: selected.github.runId, runAttempt: selected.github.qualificationAttempt } };
}

test('finalization attempt two reuses the exact successful qualification from attempt one', () => {
  const data = fixture(), selected = selectQualifiedInputs(data), core = identity(selected);
  assert.equal(selected.github.qualificationAttempt, '1');
  assert.equal(selected.github.finalizationAttempt, '2');
  assert.equal(selected.jobs.length, 15); assert.equal(selected.artifacts.length, 5);
  assert.deepEqual(selected.artifacts.map(artifact => artifact.producerJobId), [1, 2, 3, 4, 5]);
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

test('latest qualification jobs must all succeed in one attempt without hiding newer failures', () => {
  const mutations = [
    data => data.jobs.push({ ...data.jobs[0], id: 1000, run_attempt: 2, conclusion: 'failure' }),
    data => data.jobs.push({ ...data.jobs[0], id: 1000, run_attempt: 2, status: 'in_progress', conclusion: null }),
    data => data.jobs.push({ ...data.jobs[0], id: 1000, run_attempt: 2 }),
    data => data.jobs.push({ ...data.jobs[0], id: 1000 }),
    data => data.jobs.pop(),
    data => data.jobs.push({ ...data.jobs[0], id: 1000, name: 'fixtures (linux-amd64, shard 4)' }),
    data => { data.jobs[0].run_attempt = 3; },
    data => { data.jobs[0].run_id = 43; },
    data => { data.jobs[0].head_sha = 'f'.repeat(40); },
  ];
  mutations.forEach((mutate, index) => { const data = fixture(); mutate(data); assert.throws(() => selectQualifiedInputs(data), undefined, 'job mutation ' + index); });
  const newer = fixture();
  newer.jobs = [...newer.jobs.map(job => ({ ...job, conclusion: 'failure' })),
    ...newer.jobs.map(job => ({ ...job, id: job.id + 1000, run_attempt: 2, started_at: '2026-10-30T01:00:00Z', completed_at: '2026-10-30T01:10:00Z' }))];
  assert.throws(() => selectQualifiedInputs(newer), /Artifact was not created/);
  newer.artifacts.forEach(artifact => { artifact.created_at = '2026-10-30T01:05:00Z'; });
  assert.equal(selectQualifiedInputs(newer).github.qualificationAttempt, '2');
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
  mutations.forEach((mutate, index) => { const data = fixture(); mutate(data); assert.throws(() => selectQualifiedInputs(data), undefined, 'source mutation ' + index); });
});

test('artifact selection rejects omissions, ambiguity and wrong producer, run or digest', () => {
  const mutations = [
    data => data.artifacts.pop(), data => data.artifacts.push({ ...data.artifacts[0], id: 999 }),
    data => { data.artifacts[4].name = 'fixtures-linux-amd64-4'; }, data => { data.artifacts[0].expired = true; },
    data => { data.artifacts[0].workflow_run.id = 43; }, data => { data.artifacts[0].workflow_run.head_sha = 'f'.repeat(40); },
    data => { data.artifacts[0].digest = 'sha256:short'; }, data => { data.artifacts[0].size_in_bytes = 0; },
    data => { data.artifacts[0].size_in_bytes = 1.5; }, data => { data.artifacts[0].id = -1; },
    data => { data.artifacts[1].id = data.artifacts[0].id; }, data => { delete data.artifacts[0].created_at; },
    data => { data.artifacts[0].created_at = '2026-10-29T23:59:59Z'; }, data => { data.artifacts[0].created_at = '2026-10-30T00:10:01Z'; },
  ];
  mutations.forEach((mutate, index) => { const data = fixture(); mutate(data); assert.throws(() => selectQualifiedInputs(data), undefined, 'artifact mutation ' + index); });
});

test('merged reports retain qualification identity while finalization remains bound to the current attempt', () => {
  const mutations = [
    (selected, core) => { core.compiler.nativeSha256 = 'f'.repeat(64); },
    (selected, core) => { core.compiler.javascriptSha256 = 'f'.repeat(64); },
    (selected, core) => { core.source.dirty = true; }, (selected, core) => { core.source.commit = 'f'.repeat(40); },
    (selected, core) => { core.target = 'darwin-arm64'; }, (selected, core) => { core.github.runAttempt = '2'; },
    selected => { selected.github.finalizationAttempt = '1'; }, selected => { selected.jobs[0].attempt = '2'; },
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
