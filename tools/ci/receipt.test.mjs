import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { execFileSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { readQualificationArtifact, reuseQualification, validateQualification, validateReuse } from './receipt.mjs';

const hash = bytes => createHash('sha256').update(bytes).digest('hex');
const encode = value => Buffer.from(JSON.stringify(value));
function fixture() {
  const repository = 'owner/gopurs', commit = 'a'.repeat(40), tree = 'b'.repeat(40), head = 'c'.repeat(40), base = 'd'.repeat(40);
  const names = ['release-quality-resources', 'release-npm-provenance', 'ffi-errors', 'launcher', 'test-runner',
    'build-native', 'bootstrap-inputs', 'cli', 'build-profiles', 'manifest-prepared-cache', 'driver-cache', 'native-go-code', 'preparation-native'];
  const quality = encode({ schema: 1, status: 'passed', full: false, requestedFull: false, phase: 'all', shard: null,
    identity: { source: { commit, dirty: false }, target: 'linux-amd64', github: { repository, runId: '123', runAttempt: '1' } },
    gates: [...names.map(name => `Node ${name}.test.mjs`), 'Go build-cache', 'Go c1c', 'Go ffi-gen', 'Representative executable fixtures', 'npm package surface']
      .map((name, index) => ({ name, index: index + 1, passed: true, exitCode: 0 })) });
  const smoke = encode({ version: '0.1.0-test', tests: [{ host: 'javascript', passed: true }, { host: 'native', passed: true }] });
  const counts = [12, 8, 14, 9, 96, 7, 5, 19, 3, 7, 4, 2, 1];
  const receipt = { schema: 1, policy: 'linux-quick-v1', repository, ref: 'refs/pull/7/merge', commit, tree,
    runId: '123', runAttempt: '1', version: '0.1.0-test', parents: [base, head], qualitySha256: hash(quality), smokeSha256: hash(smoke),
    nodeResults: names.map((name, index) => ({ name: `Node ${name}.test.mjs`, tests: counts[index], pass: counts[index], fail: 0, cancelled: 0, skipped: 0, todo: 0 })),
    nodeTotals: { tests: 187, pass: 187, fail: 0, cancelled: 0, skipped: 0, todo: 0 } };
  const context = { repository, tree, pr: { number: 7, merged: true, head: { sha: head, repo: { full_name: repository } }, base: { ref: 'main', repo: { full_name: repository } } },
    run: { id: 123, run_attempt: 1, event: 'pull_request', status: 'completed', conclusion: 'success', path: '.github/workflows/ci.yml',
      repository: { full_name: repository }, head_repository: { full_name: repository }, head_sha: head, pull_requests: [{ number: 7, head: { sha: head } }] },
    tested: { sha: commit, tree: { sha: tree }, parents: [{ sha: base }, { sha: head }] } };
  return { receipt, quality, smoke, context };
}

test('reuse requires the tested merge tree, not the PR head', () => {
  const f = fixture();
  validateReuse(f.receipt, f.quality, f.smoke, f.context);
  f.context.tested.sha = f.context.pr.head.sha;
  assert.throws(() => validateReuse(f.receipt, f.quality, f.smoke, f.context));
});

test('a cleared post-merge PR list is accepted only with the other identity links intact', () => {
  const f = fixture();
  f.context.run.pull_requests = [];
  validateReuse(f.receipt, f.quality, f.smoke, f.context);
  f.context.tested.tree.sha = 'e'.repeat(40);
  assert.throws(() => validateReuse(f.receipt, f.quality, f.smoke, f.context));
});

test('changed tree, parents, source, run, attempt, PR, repository and fork are rejected', () => {
  const mutations = [
    f => { f.context.tree = 'e'.repeat(40); },
    f => { f.context.tested.tree.sha = 'e'.repeat(40); },
    f => { f.context.tested.parents.reverse(); },
    f => { f.context.pr.head.sha = 'e'.repeat(40); },
    f => { f.context.pr.head.repo.full_name = 'fork/gopurs'; },
    f => { f.context.pr.merged = false; },
    f => { f.context.pr.base.ref = 'edge'; },
    f => { f.context.pr.number++; },
    f => { f.context.run.id++; },
    f => { f.context.run.run_attempt++; },
    f => { f.context.run.conclusion = 'failure'; },
    f => { f.context.run.status = 'in_progress'; },
    f => { f.context.run.event = 'push'; },
    f => { f.context.run.path = '.github/workflows/other.yml'; },
    f => { f.context.run.head_repository.full_name = 'fork/gopurs'; },
    f => { delete f.context.run.pull_requests; },
    f => { f.context.run.pull_requests = null; },
    f => { f.context.run.pull_requests = [{ number: 8, head: { sha: f.context.pr.head.sha } }]; },
    f => { f.context.run.pull_requests = [{ number: 7, head: { sha: 'e'.repeat(40) } }]; },
    f => { f.receipt.parents = [f.context.pr.head.sha]; },
    f => { f.receipt.nodeTotals.skipped = 1; },
    f => { f.receipt.nodeResults[0].tests = 0; f.receipt.nodeResults[0].pass = 0; },
    f => { f.receipt.nodeResults[0].pass--; },
    f => { f.receipt.nodeResults[0].skipped = 1; },
    f => { f.receipt.nodeResults[0].tests += 0.5; },
    f => { f.receipt.nodeResults.pop(); },
  ];
  for (const mutate of mutations) {
    const f = fixture(); mutate(f);
    assert.throws(() => validateReuse(f.receipt, f.quality, f.smoke, f.context));
  }
});

test('additional successful tests are accepted with coherent per-file and total counts', () => {
  const f = fixture();
  f.receipt.nodeResults[0].tests++; f.receipt.nodeResults[0].pass++;
  f.receipt.nodeTotals.tests++; f.receipt.nodeTotals.pass++;
  validateReuse(f.receipt, f.quality, f.smoke, f.context);
  f.receipt.nodeTotals.tests--;
  assert.throws(() => validateReuse(f.receipt, f.quality, f.smoke, f.context));
});

test('quality must include every gate, exact source/run, zero failures and both smokes', () => {
  for (const mutate of [q => { q.gates.pop(); }, q => { q.gates[0].exitCode = 1; },
    q => { q.gates[0].passed = false; }, q => { q.gates[0].index = 2; },
    q => { q.full = true; }, q => { q.identity.source.dirty = true; },
    q => { q.identity.source.commit = 'f'.repeat(40); }, q => { q.identity.github.runAttempt = '2'; },
    q => { q.identity.target = 'darwin-arm64'; }]) {
    const f = fixture(), q = JSON.parse(f.quality); mutate(q); f.quality = encode(q); f.receipt.qualitySha256 = hash(f.quality);
    assert.throws(() => validateQualification(f.receipt, f.quality, f.smoke));
  }
  const f = fixture(), s = JSON.parse(f.smoke); s.tests[1].passed = false; f.smoke = encode(s); f.receipt.smokeSha256 = hash(f.smoke);
  assert.throws(() => validateQualification(f.receipt, f.quality, f.smoke));
});

test('receipt ZIP digest, byte count and exact member inventory are enforced', t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-receipt-test-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  const f = fixture();
  for (const [name, bytes] of [['receipt.json', encode(f.receipt)], ['quality.json', f.quality], ['smoke.json', f.smoke]]) writeFileSync(join(directory, name), bytes);
  const zip = join(directory, 'receipt.zip');
  execFileSync('zip', ['-q', zip, 'receipt.json', 'quality.json', 'smoke.json'], { cwd: directory });
  const bytes = readFileSync(zip), meta = { size_in_bytes: bytes.length, digest: `sha256:${hash(bytes)}` };
  const actual = readQualificationArtifact(bytes, meta);
  validateReuse(actual.receipt, actual.quality, actual.smoke, f.context);
  assert.throws(() => readQualificationArtifact(bytes, { ...meta, digest: `sha256:${'0'.repeat(64)}` }));
  assert.throws(() => readQualificationArtifact(bytes, { ...meta, size_in_bytes: bytes.length + 1 }));
  writeFileSync(join(directory, 'extra.json'), '{}');
  execFileSync('zip', ['-q', zip, 'extra.json'], { cwd: directory });
  const extra = readFileSync(zip);
  assert.throws(() => readQualificationArtifact(extra, { size_in_bytes: extra.length, digest: `sha256:${hash(extra)}` }));
});

test('tags, full campaigns and other events never attempt reuse', () => {
  for (const env of [{ GITHUB_EVENT_NAME: 'push', GITHUB_REF: 'refs/tags/v1' },
    { GITHUB_EVENT_NAME: 'push', GITHUB_REF: 'refs/heads/main', CI_FULL: 'true' },
    { GITHUB_EVENT_NAME: 'pull_request', GITHUB_REF: 'refs/pull/7/merge' },
    { GITHUB_EVENT_NAME: 'push', GITHUB_REF: 'refs/heads/edge' }]) {
    assert.throws(() => reuseQualification(env, 'b'.repeat(40)));
  }
});

test('receipt CLI records a real clean merge commit and validates all Node summaries', t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-receipt-git-')), checkout = join(directory, 'checkout');
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(checkout);
  const git = (...args) => execFileSync('git', args, { cwd: checkout, encoding: 'utf8', stdio: ['ignore', 'pipe', 'pipe'] }).trim();
  git('init', '-b', 'main'); git('config', 'user.name', 'CI test'); git('config', 'user.email', 'ci@example.invalid');
  git('config', 'commit.gpgsign', 'false');
  writeFileSync(join(checkout, 'package.json'), '{"version":"0.1.0-test"}\n');
  git('add', '.'); git('commit', '-m', 'base'); const base = git('rev-parse', 'HEAD');
  git('checkout', '-b', 'feature'); writeFileSync(join(checkout, 'source.txt'), 'changed\n');
  git('add', '.'); git('commit', '-m', 'feature'); const head = git('rev-parse', 'HEAD');
  git('checkout', 'main'); git('merge', '--no-ff', '--no-edit', 'feature');
  const commit = git('rev-parse', 'HEAD'), tree = git('rev-parse', 'HEAD^{tree}');
  const f = fixture(), quality = JSON.parse(f.quality), counts = [12, 8, 14, 9, 96, 7, 5, 19, 3, 7, 4, 2, 1];
  quality.identity.source.commit = commit;
  quality.gates.slice(0, 13).forEach((gate, index) => {
    gate.log = join(directory, `${index}.log`);
    writeFileSync(gate.log, `# tests ${counts[index]}\n# pass ${counts[index]}\n# fail 0\n# cancelled 0\n# skipped 0\n# todo 0\n`);
  });
  const qualityFile = join(directory, 'quality.json'), smokeFile = join(directory, 'smoke.json'), destination = join(directory, 'receipt');
  writeFileSync(qualityFile, encode(quality)); writeFileSync(smokeFile, f.smoke);
  const env = { ...process.env, GITHUB_EVENT_NAME: 'pull_request', GITHUB_REF: 'refs/pull/7/merge', GITHUB_SHA: commit,
    GITHUB_REPOSITORY: 'owner/gopurs', GITHUB_RUN_ID: '123', GITHUB_RUN_ATTEMPT: '1' };
  execFileSync(process.execPath, [fileURLToPath(new URL('receipt.mjs', import.meta.url)), qualityFile, smokeFile, destination], { cwd: checkout, env });
  const receipt = JSON.parse(readFileSync(join(destination, 'receipt.json')));
  assert.equal(receipt.commit, commit); assert.equal(receipt.tree, tree); assert.deepEqual(receipt.parents, [base, head]);
  validateQualification(receipt, readFileSync(qualityFile), f.smoke);
  writeFileSync(quality.gates[0].log, '# tests 13\n# pass 13\n# fail 0\n# cancelled 0\n# skipped 0\n# todo 0\n');
  const additional = join(directory, 'receipt-additional-test');
  execFileSync(process.execPath, [fileURLToPath(new URL('receipt.mjs', import.meta.url)), qualityFile, smokeFile, additional], { cwd: checkout, env });
  const more = JSON.parse(readFileSync(join(additional, 'receipt.json')));
  assert.equal(more.nodeTotals.tests, receipt.nodeTotals.tests + 1);
  validateQualification(more, readFileSync(qualityFile), f.smoke);
  writeFileSync(quality.gates[0].log, '# tests 12\n# pass 11\n# fail 0\n# cancelled 0\n# skipped 1\n# todo 0\n');
  assert.throws(() => execFileSync(process.execPath,
    [fileURLToPath(new URL('receipt.mjs', import.meta.url)), qualityFile, smokeFile, join(directory, 'rejected')],
    { cwd: checkout, env, stdio: 'pipe' }));
});
