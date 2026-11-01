import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { execFileSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const policy = 'linux-quick-v2';
const nodeGates = ['release-quality-resources', 'release-npm-provenance', 'release-sbom', 'emission', 'ffi-errors',
  'launcher', 'test-runner', 'build-native', 'bootstrap-inputs', 'cli',
  'build-profiles', 'manifest-prepared-cache', 'driver-cache', 'native-go-code', 'preparation-native',
  'check-output-identity', 'check-docs', 'embed-runtime', 'compiler-snapshot']
  .map(name => `Node ${name}.test.mjs`);
const gates = [...nodeGates,
  'Go build-cache', 'Go c1c', 'Go ffi-gen', 'Representative executable fixtures', 'npm package surface'];
const sha = bytes => createHash('sha256').update(bytes).digest('hex');
const git = (...args) => execFileSync('git', args, { encoding: 'utf8' }).trim();
const emptyTotals = () => ({ tests: 0, pass: 0, fail: 0, cancelled: 0, skipped: 0, todo: 0 });

function totalsOfNodeResults(results) {
  assert.deepEqual(results.map(result => result.name), nodeGates);
  const totals = emptyTotals();
  for (const result of results) {
    for (const field of Object.keys(totals)) {
      assert.ok(Number.isSafeInteger(result[field]) && result[field] >= 0);
      if (!['tests', 'pass'].includes(field)) assert.equal(result[field], 0);
      totals[field] += result[field];
    }
    assert.ok(result.tests > 0); assert.equal(result.pass, result.tests);
  }
  return totals;
}

export function validateQualification(receipt, qualityBytes, smokeBytes) {
  assert.equal(receipt.schema, 1); assert.equal(receipt.policy, policy);
  assert.match(receipt.commit, /^[a-f0-9]{40}$/); assert.match(receipt.tree, /^[a-f0-9]{40}$/);
  assert.equal(receipt.qualitySha256, sha(qualityBytes)); assert.equal(receipt.smokeSha256, sha(smokeBytes));
  const quality = JSON.parse(qualityBytes), smoke = JSON.parse(smokeBytes);
  assert.equal(quality.schema, 1);
  assert.equal(quality.status, 'passed'); assert.equal(quality.full, false);
  assert.equal(quality.requestedFull, false); assert.equal(quality.phase, 'all'); assert.equal(quality.shard, null);
  assert.deepEqual(quality.identity.source, { commit: receipt.commit, dirty: false });
  assert.deepEqual(quality.identity.github, { repository: receipt.repository, runId: receipt.runId, runAttempt: receipt.runAttempt });
  assert.equal(quality.identity.target, 'linux-amd64');
  assert.deepEqual(quality.gates.map(gate => gate.name), gates);
  assert.ok(quality.gates.every((gate, index) => gate.index === index + 1 && gate.passed === true && gate.exitCode === 0));
  assert.deepEqual(receipt.nodeTotals, totalsOfNodeResults(receipt.nodeResults));
  assert.equal(smoke.version, receipt.version);
  assert.equal(smoke.tests.length, 2);
  assert.deepEqual(smoke.tests.map(test => test.host).sort(), ['javascript', 'native']);
  assert.ok(smoke.tests.every(test => test.passed === true));
}

export function validateReuse(receipt, quality, smoke, { repository, tree, pr, run, tested }) {
  validateQualification(receipt, quality, smoke);
  assert.equal(pr.merged, true); assert.equal(pr.base.ref, 'main');
  assert.equal(pr.head.repo.full_name, repository); assert.equal(pr.base.repo.full_name, repository);
  assert.equal(run.event, 'pull_request'); assert.equal(run.path, '.github/workflows/ci.yml');
  assert.equal(run.status, 'completed'); assert.equal(run.conclusion, 'success');
  assert.equal(run.repository.full_name, repository); assert.equal(run.head_repository.full_name, repository);
  assert.equal(run.head_sha, pr.head.sha);
  // GitHub clears this list after a PR is merged. The associated merged PR,
  // receipt ref and actual tested Git tree/parents remain mandatory below.
  assert.ok(Array.isArray(run.pull_requests));
  if (run.pull_requests.length) assert.ok(run.pull_requests.some(item => item.number === pr.number && item.head.sha === pr.head.sha));
  assert.equal(receipt.repository, repository); assert.equal(receipt.ref, `refs/pull/${pr.number}/merge`);
  assert.equal(receipt.runId, String(run.id)); assert.equal(receipt.runAttempt, String(run.run_attempt));
  assert.equal(receipt.commit, tested.sha); assert.equal(receipt.tree, tree); assert.equal(tested.tree.sha, tree);
  assert.equal(tested.parents.length, 2);
  assert.deepEqual(receipt.parents, tested.parents.map(parent => parent.sha));
  assert.ok(receipt.parents.includes(pr.head.sha));
}

export function readQualificationArtifact(bytes, artifact) {
  assert.ok(artifact.size_in_bytes > 0 && artifact.size_in_bytes < 512 * 1024);
  assert.equal(bytes.length, artifact.size_in_bytes); assert.equal(`sha256:${sha(bytes)}`, artifact.digest);
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-ci-receipt-'));
  try {
    const file = join(directory, 'receipt.zip'); writeFileSync(file, bytes);
    const unzip = (...args) => execFileSync('unzip', args, { timeout: 10000, maxBuffer: 128 * 1024 });
    assert.deepEqual(unzip('-Z1', file).toString().trim().split('\n').sort(), ['quality.json', 'receipt.json', 'smoke.json']);
    const receipt = JSON.parse(unzip('-p', file, 'receipt.json'));
    assert.match(receipt.commit, /^[a-f0-9]{40}$/);
    return { receipt, quality: unzip('-p', file, 'quality.json'), smoke: unzip('-p', file, 'smoke.json') };
  } finally { rmSync(directory, { recursive: true, force: true }); }
}

export function reuseQualification(env, tree) {
  assert.equal(env.GITHUB_EVENT_NAME, 'push'); assert.equal(env.GITHUB_REF, 'refs/heads/main');
  assert.notEqual(env.CI_FULL, 'true');
  const repository = env.GITHUB_REPOSITORY;
  const request = endpoint => execFileSync('gh', ['api', '--hostname', 'github.com', `repos/${repository}/${endpoint}`],
    { timeout: 30000, maxBuffer: 1024 * 1024, env: { ...env, GH_PROMPT_DISABLED: '1' } });
  const api = endpoint => JSON.parse(request(endpoint));
  const prs = api(`commits/${env.GITHUB_SHA}/pulls?per_page=100`);
  assert.equal(prs.length, 1, 'Expected one associated PR');
  const pr = api(`pulls/${prs[0].number}`);
  assert.equal(pr.merge_commit_sha, env.GITHUB_SHA, 'PR is not this push');
  assert.equal(pr.head.repo.full_name, repository, 'Fork qualifications are not reused');
  const runs = api(`actions/workflows/ci.yml/runs?event=pull_request&head_sha=${pr.head.sha}&status=success&per_page=100`);
  assert.equal(runs.total_count, 1, 'Expected one successful PR run');
  const run = runs.workflow_runs[0];
  const listed = api(`actions/runs/${run.id}/artifacts?per_page=100`);
  assert.ok(listed.total_count <= 100, 'Artifact listing is incomplete');
  const artifacts = listed.artifacts.filter(item => item.name === 'qualification-linux-amd64');
  assert.equal(artifacts.length, 1, 'Expected one qualification artifact');
  const artifact = artifacts[0];
  assert.equal(artifact.expired, false); assert.equal(artifact.workflow_run.id, run.id);
  assert.equal(artifact.workflow_run.head_sha, pr.head.sha);
  const bytes = request(`actions/artifacts/${artifact.id}/zip`);
  const { receipt, quality, smoke } = readQualificationArtifact(bytes, artifact);
  validateReuse(receipt, quality, smoke, { repository, tree, pr, run, tested: api(`git/commits/${receipt.commit}`) });
  return { runId: run.id, runAttempt: run.run_attempt, testedCommit: receipt.commit, tree,
    artifactId: artifact.id, artifactDigest: artifact.digest, policy };
}

function createReceipt(qualityFile, smokeFile, destination) {
  const env = process.env;
  assert.equal(env.GITHUB_EVENT_NAME, 'pull_request');
  assert.match(env.GITHUB_REF, /^refs\/pull\/\d+\/merge$/);
  assert.equal(git('rev-parse', 'HEAD'), env.GITHUB_SHA);
  assert.equal(git('status', '--porcelain', '--untracked-files=all'), '');
  const qualityBytes = readFileSync(qualityFile), smokeBytes = readFileSync(smokeFile);
  const quality = JSON.parse(qualityBytes);
  const nodeResults = quality.gates.filter(gate => gate.name.startsWith('Node ')).map(gate => {
    const text = readFileSync(gate.log, 'utf8');
    const result = { name: gate.name, ...emptyTotals() };
    for (const field of Object.keys(emptyTotals())) {
      const matches = [...text.matchAll(new RegExp(`^(?:ℹ|#) ${field} (\\d+)\\s*$`, 'gm'))];
      assert.equal(matches.length, 1); result[field] = Number(matches[0][1]);
    }
    return result;
  });
  const receipt = { schema: 1, policy, repository: env.GITHUB_REPOSITORY, ref: env.GITHUB_REF,
    version: JSON.parse(readFileSync('package.json', 'utf8')).version,
    runId: env.GITHUB_RUN_ID, runAttempt: env.GITHUB_RUN_ATTEMPT, commit: env.GITHUB_SHA,
    tree: git('rev-parse', 'HEAD^{tree}'), parents: git('show', '-s', '--format=%P', 'HEAD').split(' '),
    qualitySha256: sha(qualityBytes), smokeSha256: sha(smokeBytes), nodeResults, nodeTotals: totalsOfNodeResults(nodeResults) };
  validateQualification(receipt, qualityBytes, smokeBytes);
  mkdirSync(destination);
  writeFileSync(join(destination, 'quality.json'), qualityBytes); writeFileSync(join(destination, 'smoke.json'), smokeBytes);
  writeFileSync(join(destination, 'receipt.json'), JSON.stringify(receipt, null, 2) + '\n');
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) createReceipt(...process.argv.slice(2));
