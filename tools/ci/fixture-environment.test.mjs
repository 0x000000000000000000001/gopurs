import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, mkdirSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { cacheInventory, closureIdentity, environmentScript, printEnvironmentCommand, validateCoreEnvironment, validateManifest } from './fixture-environment.mjs';
import { validateEnvironmentIdentity } from './release-inputs.mjs';

const profile = '/nix/store/' + 'a'.repeat(32) + '-gopurs-env';
const tool = '/nix/store/' + 'b'.repeat(32) + '-purs';
const hash = value => createHash('sha256').update(value).digest('hex');
const closure = () => [
  { path: profile, narHash: 'sha256-' + 'A'.repeat(43) + '=', narSize: 120, references: [tool] },
  { path: tool, narHash: 'sha256-' + 'B'.repeat(43) + '=', narSize: 500, references: [] },
];
function sample() {
  const context = { source: { commit: 'a'.repeat(40), tree: 'b'.repeat(40) }, target: 'linux-amd64',
    compiler: { javascriptSha256: 'c'.repeat(64), nativeSha256: 'd'.repeat(64) },
    github: { repository: 'owner/repo', runId: '10', runAttempt: '2' } };
  return { context, manifest: { schema: 1, ...structuredClone(context),
    environment: { profile, profileSha256: 'e'.repeat(64), closureSha256: hash(JSON.stringify(closure())), shell: tool + '/bin/bash' },
    closure: closure(), cache: [{ path: 'nix-cache-info', bytes: 21, sha256: 'f'.repeat(64) }] } };
}

test('transferred environment is tied to the original source, platform, binaries and producer', () => {
  const { context, manifest } = sample();
  assert.equal(validateManifest(manifest, context), manifest);
  validateEnvironmentIdentity(manifest.environment);
  for (const [section, field, value] of [
    ['source', 'commit', 'f'.repeat(40)], ['source', 'tree', 'e'.repeat(40)],
    ['compiler', 'nativeSha256', '0'.repeat(64)], ['compiler', 'javascriptSha256', '1'.repeat(64)],
    ['github', 'repository', 'other/repo'], ['github', 'runId', '11'], ['github', 'runAttempt', '3'],
  ]) {
    const changed = structuredClone(manifest); changed[section][field] = value;
    assert.throws(() => validateManifest(changed, context));
  }
  assert.throws(() => validateManifest({ ...manifest, target: 'linux-arm64' }, context));
});

test('runtime closure requires every reference and rejects build derivations', () => {
  assert.deepEqual(closureIdentity(Object.fromEntries(closure().map(({ path, ...entry }) => [path, entry]))), closure());
  assert.throws(() => closureIdentity(closure().slice(0, 1)), /Incomplete runtime closure/);
  assert.throws(() => closureIdentity([...closure(), closure()[0]]));
  assert.throws(() => closureIdentity([{ ...closure()[1], path: tool + '.drv' }]), /runtime outputs/);
  const { context, manifest } = sample();
  manifest.closure[1].narHash = 'sha256-' + 'C'.repeat(43) + '=';
  assert.throws(() => validateManifest(manifest, context), /Runtime closure identity differs/);
});

test('fixture preparation rejects an environment different from the successful core before running tests', () => {
  const { manifest } = sample();
  const core = { source: { commit: manifest.source.commit, dirty: false }, compiler: manifest.compiler,
    target: manifest.target, github: manifest.github, environment: manifest.environment };
  validateCoreEnvironment(manifest, core);
  for (const key of ['profileSha256', 'closureSha256']) {
    const changed = structuredClone(core); changed.environment[key] = '0'.repeat(64);
    assert.throws(() => validateCoreEnvironment(manifest, changed), /Core qualification/);
  }
  const wrongAttempt = structuredClone(core); wrongAttempt.github.runAttempt = '3';
  assert.throws(() => validateCoreEnvironment(manifest, wrongAttempt), /Core qualification/);
});

test('cache inventory detects corrupted bytes and rejects links or extra files', () => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-fixture-cache-'));
  try {
    mkdirSync(join(directory, 'nar'));
    const auxiliary = ['log', 'build-trace-v2', 'realisations'];
    for (const name of auxiliary) mkdirSync(join(directory, name));
    writeFileSync(join(directory, 'nix-cache-info'), 'StoreDir: /nix/store\n');
    writeFileSync(join(directory, 'a'.repeat(32) + '.narinfo'), 'test narinfo');
    const archive = join(directory, 'nar/test.nar.zst'); writeFileSync(archive, 'original runtime bytes');
    const before = cacheInventory(directory);
    for (const name of auxiliary) {
      const metadata = join(directory, name, 'must-not-be-ignored.json');
      writeFileSync(metadata, '{}');
      assert.throws(() => cacheInventory(directory), /Unexpected contents in runtime cache directory/);
      rmSync(metadata);
    }
    mkdirSync(join(directory, 'unrecognized'));
    assert.throws(() => cacheInventory(directory));
    rmSync(join(directory, 'unrecognized'), { recursive: true });
    writeFileSync(archive, 'tampered runtime bytes');
    assert.notDeepEqual(cacheInventory(directory), before);
    symlinkSync(archive, join(directory, 'nar/link.nar.zst'));
    assert.throws(() => cacheInventory(directory), /links/);
    rmSync(join(directory, 'nar/link.nar.zst'));
    writeFileSync(join(directory, 'token'), 'must not be exported');
    assert.throws(() => cacheInventory(directory));
  } finally { rmSync(directory, { recursive: true, force: true }); }
});

test('fixture commands enter only the transferred profile with every build route disabled', () => {
  const args = printEnvironmentCommand(profile);
  assert.deepEqual(args, ['print-dev-env', profile, '--offline', '--no-substitute', '--max-jobs', '0', '--builders', '']);
  for (const input of ['.', '.#default', '/tmp/profile', profile + '.drv', profile + '\n']) {
    assert.throws(() => printEnvironmentCommand(input));
  }
});

test('rendered environment runs the hook once and stops before the command if the hook fails', () => {
  const script = environmentScript('export GOPURS_HOOK_COUNT=0\nshellHook=\'export GOPURS_HOOK_COUNT=$(( GOPURS_HOOK_COUNT + 1 ))\'\neval "$shellHook"');
  const success = spawnSync('bash', ['-c', script, 'environment', 'bash', '-c', 'printf "%s:%s" "$GOPURS_HOOK_COUNT" "$1"', 'command', 'literal $VALUE'], { encoding: 'utf8' });
  assert.equal(success.status, 0); assert.equal(success.stdout, '1:literal $VALUE');
  const failure = spawnSync('bash', ['-c', environmentScript('shellHook=false\neval "$shellHook"'), 'environment', 'echo', 'must-not-run'], { encoding: 'utf8' });
  assert.notEqual(failure.status, 0); assert.equal(failure.stdout, '');
});
