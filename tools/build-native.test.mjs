import assert from 'node:assert/strict';
import { spawn, spawnSync } from 'node:child_process';
import {
  copyFileSync, existsSync, mkdirSync, mkdtempSync, readFileSync,
  realpathSync, rmSync, statSync, utimesSync, writeFileSync,
} from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join } from 'node:path';
import test from 'node:test';

function fixture(t) {
  const temporary = mkdtempSync(join(tmpdir(), 'gopurs-native-helper-'));
  t.after(() => rmSync(temporary, { recursive: true, force: true }));
  const root = join(temporary, 'checkout/gopurs/gopurs');
  const localBin = join(root, 'node_modules/.bin');
  const log = join(temporary, 'calls.jsonl');
  const put = (file, contents, mode) => {
    mkdirSync(dirname(file), { recursive: true });
    writeFileSync(file, contents, { mode });
  };
  mkdirSync(join(root, 'src'), { recursive: true });
  mkdirSync(join(temporary, 'tmp'));
  put(join(root, 'spago.yaml'), `package:
  name: gopurs
  dependencies:
    - aff
    - backend-optimizer
  test:
    main: Test.Main
    dependencies: []
workspace:
  packageSet:
    registry: 77.10.1
`);
  const optimizerRevision = '0123456789abcdef0123456789abcdef01234567';
  put(join(root, 'spago.lock'), JSON.stringify({ packages: {
    'backend-optimizer': { type: 'git', rev: optimizerRevision },
  } }));
  put(join(root, `.spago/p/backend-optimizer/${optimizerRevision}/spago.yaml`),
    'package:\n  name: backend-optimizer\n');
  for (const name of ['aff', 'argonaut-codecs', 'ordered-collections', 'node-process']) {
    put(join(temporary, `checkout/gopurs/gopurs-${name}/spago.yaml`),
      `package:\n  name: ${name === 'node-process' ? 'gopurs-node-process' : name}\n`);
  }
  mkdirSync(join(root, 'tools/build'), { recursive: true });
  mkdirSync(join(root, 'tools/lib'), { recursive: true });
  for (const [source, target] of [
    ['./build/build-native.mjs', 'tools/build/build-native.mjs'],
    ['./build/native-workspace.mjs', 'tools/build/native-workspace.mjs'],
    ['./lib/process.mjs', 'tools/lib/process.mjs'],
    ['./lib/fs.mjs', 'tools/lib/fs.mjs'],
  ]) {
    copyFileSync(new URL(source, import.meta.url), join(root, target));
  }
  // Every external build command is replaced; this test never compiles sources.
  const mockTool = `#!/usr/bin/env node
const fs = require('node:fs'), path = require('node:path');
const name = path.basename(process.argv[1]);
const purs = process.env.PATH.split(path.delimiter).map(dir => path.join(dir, 'purs')).find(fs.existsSync);
fs.appendFileSync(process.env.MOCK_LOG, JSON.stringify({ name, args: process.argv.slice(2), cwd: process.cwd(), purs: fs.realpathSync(purs) }) + '\\n');
if (process.env.MOCK_FAIL === name) process.exit(42);
if (name === 'spago') {
  fs.mkdirSync('output/Main', { recursive: true });
  fs.mkdirSync('output/main', { recursive: true });
  fs.writeFileSync('output/main/main.go', 'package main');
  fs.writeFileSync('output/Main/corefn.json', JSON.stringify(process.env.MOCK_UNTYPED ? {} : {
    typeTable: ['Int'], dataDecls: [], classDecls: [],
  }));
}
if (process.env.MOCK_PAUSE === name) {
  const child = require('node:child_process').spawn(process.execPath, ['-e', 'setInterval(() => {}, 1000)'], { stdio: 'ignore' });
  child.once('spawn', () => fs.writeFileSync(process.env.MOCK_READY, JSON.stringify({ parent: process.pid, child: child.pid })));
  setInterval(() => {}, 1000);
} else if (name === 'go') {
  fs.writeFileSync(process.argv[process.argv.indexOf('-o') + 1], 'new-native-binary');
}
`;
  for (const name of ['npm', 'spago', 'purs', 'go']) put(join(localBin, name), mockTool, 0o755);
  const fork = join(temporary, 'checkout/purescript/.stack-work/dist/mock/ghc-test/build/purs/purs');
  const staleFork = join(temporary, 'checkout/purescript/.stack-work/dist/other/ghc-old/build/purs/purs');
  const explicitFork = join(temporary, 'explicit-typed-purs');
  for (const file of [fork, staleFork, explicitFork, join(root, 'bin/gopurs'), join(root, 'bin/gopurs.js')]) put(file, mockTool, 0o755);
  utimesSync(staleFork, 1, 1);
  utimesSync(fork, 2, 2);
  const binary = join(root, 'bin/gopurs-native');
  put(binary, 'old-native-binary', 0o755);
  put(join(root, 'tools/build/prepare-native-output.mjs'), `
import { appendFileSync } from 'node:fs';
appendFileSync(process.env.MOCK_LOG, JSON.stringify({ name: 'parser' }) + '\\n');
`);
  const ready = join(temporary, 'ready.json');
  const environmentFor = overrides => ({
    ...process.env, TMPDIR: join(temporary, 'tmp'), MOCK_LOG: log, MOCK_READY: ready,
    MOCK_FAIL: '', MOCK_UNTYPED: '', MOCK_PAUSE: '', GOPURS_PURS: '', GOPURS_PBO_DIR: '',
    GOPURS_LIBRARY_ROOT: join(temporary, 'checkout/gopurs'), ...overrides,
  });
  const summarize = result => ({
    ...result,
    workspace: result.stdout.match(/Native bootstrap workspace: (.*)/)?.[1],
    calls: existsSync(log) ? readFileSync(log, 'utf8').trim().split('\n').map(JSON.parse) : [],
  });
  const run = (environment = {}, args = ['--keep-workspace']) => {
    const result = spawnSync(process.execPath, [join(root, 'tools/build/build-native.mjs'), ...args], {
      cwd: root, encoding: 'utf8', timeout: 15_000, env: environmentFor(environment),
    });
    assert.ifError(result.error);
    return summarize(result);
  };
  const start = environment => {
    const child = spawn(process.execPath, [join(root, 'tools/build/build-native.mjs'), '--keep-workspace'], {
      cwd: root, env: environmentFor(environment), stdio: ['ignore', 'pipe', 'pipe'],
    });
    let stdout = '', stderr = '';
    child.stdout.on('data', chunk => stdout += chunk);
    child.stderr.on('data', chunk => stderr += chunk);
    const completed = new Promise((resolve, reject) => {
      child.once('error', reject);
      child.once('close', (status, signal) => resolve(summarize({ status, signal, stdout, stderr })));
    });
    t.after(async () => {
      child.kill('SIGTERM');
      await completed;
    });
    return { child, completed };
  };
  return { run, start, ready, root, put, fork, explicitFork, localBin, binary };
}

for (const explicit of [false, true]) {
  test(`native helper separates Node and TAST compilers (${explicit ? 'GOPURS_PURS' : 'local discovery'})`, t => {
    const f = fixture(t);
    const result = f.run(explicit ? { GOPURS_PURS: f.explicitFork } : {});
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.deepEqual(result.calls.map(call => call.name), ['npm', 'spago', 'parser', 'gopurs.js', 'go']);
    const selected = realpathSync(explicit ? f.explicitFork : f.fork);
    for (const call of result.calls.filter(call => call.name !== 'parser')) {
      assert.equal(call.purs, call.name === 'spago' ? selected : realpathSync(join(f.localBin, 'purs')));
    }
    assert.equal(realpathSync(join(result.workspace, 'typed-bin/purs')), selected);
    const config = readFileSync(join(result.workspace, 'spago.yaml'), 'utf8');
    assert.doesNotMatch(config, /^  test:/m);
    assert.match(config, /registry: 77\.10\.1/);
    for (const name of ['backend-optimizer', 'aff', 'argonaut-codecs', 'ordered-collections', 'node-process']) {
      assert.ok(config.includes(`    ${name}:\n`), `missing local override ${name}`);
    }
    assert.doesNotMatch(config, /    gopurs-node-process:/);
    assert.equal(readFileSync(f.binary, 'utf8'), 'new-native-binary');
    assert.ok(statSync(f.binary).mode & 0o100, 'published binary must be executable');
  });
}

for (const hasExtraPackages of [false, true]) {
  test(`native helper uses explicit PBO for the Node seed and TAST (${hasExtraPackages ? 'existing' : 'new'} overrides)`, t => {
    const f = fixture(t);
    const optimizer = join(dirname(f.root), 'local-pbo');
    f.put(join(optimizer, 'spago.yaml'), 'package:\n  name: backend-optimizer\n');
    const configPath = join(f.root, 'spago.yaml');
    if (hasExtraPackages) {
      writeFileSync(configPath, readFileSync(configPath, 'utf8') + `  extraPackages:
    backend-optimizer:
      git: https://example.invalid/pbo.git
      ref: pinned-pbo
    st:
      git: https://example.invalid/st.git
      ref: pinned-st
`);
    }
    const originalConfig = readFileSync(configPath, 'utf8');
    const originalLock = readFileSync(join(f.root, 'spago.lock'), 'utf8');
    const result = f.run({ GOPURS_PBO_DIR: optimizer });
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.deepEqual(result.calls.map(call => call.name), ['npm', 'spago', 'spago', 'spago', 'parser', 'gopurs.js', 'go']);
    const nodeWorkspace = join(result.workspace, '.spago/node-bootstrap');
    const nodeConfig = readFileSync(join(nodeWorkspace, 'spago.yaml'), 'utf8');
    const nativeConfig = readFileSync(join(result.workspace, 'spago.yaml'), 'utf8');
    const override = `    backend-optimizer:\n      path: ${JSON.stringify(optimizer)}`;
    assert.ok(nodeConfig.includes(override));
    assert.ok(nativeConfig.includes(override));
    assert.equal(nodeConfig.match(/^    backend-optimizer:/gm).length, 1);
    assert.doesNotMatch(nodeConfig, /pinned-pbo|^  test:|^    argonaut-codecs:/m);
    if (hasExtraPackages) assert.match(nodeConfig, /ref: pinned-st/);
    assert.equal(readFileSync(configPath, 'utf8'), originalConfig);
    assert.equal(readFileSync(join(f.root, 'spago.lock'), 'utf8'), originalLock);
    assert.equal(realpathSync(join(nodeWorkspace, 'src')), realpathSync(join(f.root, 'src')));
    assert.deepEqual(result.calls[0].args, ['run', 'build:runtime']);
    for (const call of result.calls.slice(1, 3)) {
      assert.equal(call.cwd, realpathSync(nodeWorkspace));
      assert.equal(call.purs, realpathSync(join(f.localBin, 'purs')));
    }
    assert.deepEqual(result.calls[2].args, [
      'bundle', '--module', 'Main', '--platform', 'node', '--outfile', join(realpathSync(f.root), 'bin/gopurs.js'), '--bundle-type', 'app',
    ]);
    assert.equal(result.calls[3].cwd, realpathSync(result.workspace));
    assert.equal(result.calls[3].purs, realpathSync(f.fork));
    assert.equal(readFileSync(f.binary, 'utf8'), 'new-native-binary');
  });
}

test('native helper rejects untyped CoreFn before Go generation', t => {
  const f = fixture(t);
  const result = f.run({ MOCK_UNTYPED: '1' });
  assert.equal(result.status, 1);
  assert.match(result.stderr, /failed during verify-tast/);
  assert.match(result.stderr, /lacks TAST type metadata/);
  assert.deepEqual(result.calls.map(call => call.name), ['npm', 'spago']);
  assert.equal(readFileSync(f.binary, 'utf8'), 'old-native-binary');
  assert.ok(existsSync(join(result.workspace, 'typed-corefn.log')), 'failure logs must remain available');
});

test('native helper preserves the published binary when Go compilation fails', t => {
  const f = fixture(t);
  const result = f.run({ MOCK_FAIL: 'go' });
  assert.equal(result.status, 1);
  assert.match(result.stderr, /failed during go-build/);
  assert.deepEqual(result.calls.map(call => call.name), ['npm', 'spago', 'parser', 'gopurs.js', 'go']);
  assert.equal(readFileSync(f.binary, 'utf8'), 'old-native-binary');
  assert.ok(existsSync(join(result.workspace, 'go-build.log')), 'failure logs must remain available');
});

test('native helper removes a successful workspace unless retention is requested', t => {
  const f = fixture(t);
  const result = f.run({}, []);
  assert.equal(result.status, 0, result.stdout + result.stderr);
  assert.ok(result.workspace);
  assert.equal(existsSync(result.workspace), false);
  assert.equal(readFileSync(f.binary, 'utf8'), 'new-native-binary');
});

async function waitFor(check) {
  for (let i = 0; i < 200; i++) {
    if (check()) return;
    await new Promise(resolve => setTimeout(resolve, 25));
  }
  assert.fail('Timed out waiting for subprocess state');
}

for (const [signal, code] of [['SIGINT', 130], ['SIGTERM', 143]]) {
  test(`native build ${signal} stops descendants and preserves the previous binary`, async t => {
    const f = fixture(t);
    const running = f.start({ MOCK_PAUSE: 'go' });
    await waitFor(() => existsSync(f.ready));
    const pids = JSON.parse(readFileSync(f.ready, 'utf8'));
    running.child.kill(signal);
    const result = await running.completed;
    assert.equal(result.status, code, result.stdout + result.stderr);
    assert.equal(result.signal, null);
    assert.match(result.stderr, /failed during go-build: Interrupted by SIG/);
    assert.ok(existsSync(join(result.workspace, 'go-build.log')));
    assert.equal(readFileSync(f.binary, 'utf8'), 'old-native-binary');
    for (const pid of [pids.parent, pids.child]) await waitFor(() => {
      try { process.kill(pid, 0); return false; }
      catch (error) { return error.code === 'ESRCH'; }
    });
  });
}
