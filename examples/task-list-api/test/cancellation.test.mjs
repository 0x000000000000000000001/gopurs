import assert from 'node:assert/strict';
import { spawn, spawnSync } from 'node:child_process';
import { cpSync, mkdirSync, mkdtempSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { request } from 'node:http';
import { tmpdir } from 'node:os';
import { basename, delimiter, join, resolve } from 'node:path';
import test from 'node:test';
import { setTimeout as delay } from 'node:timers/promises';
import { fileURLToPath, pathToFileURL } from 'node:url';

const root = fileURLToPath(new URL('../', import.meta.url));
const compiler = resolve(process.env.GOPURS_COMPILER_ROOT || join(root, '../..'));
const libraries = resolve(process.env.GOPURS_LIBRARY_ROOT || join(compiler, '..'));
const { corePackages } = await import(pathToFileURL(join(compiler, 'tools/test-workspace.mjs')));
const { findTypedCompiler } = await import(pathToFileURL(join(compiler, 'tools/build/native-workspace.mjs')));

test('compiled killFiber drains the HTTP handler before releasing its outer resource', { timeout: 240000 }, async t => {
  const workspace = mkdtempSync(join(tmpdir(), 'gopurs-http-cancellation-'));
  let passed = false;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, { recursive: true, force: true });
    else t.diagnostic(`Cancellation workspace retained: ${workspace}`);
  });
  for (const path of ['bin', 'logs', 'src/Test', 'src/Example', '.spago/ports']) mkdirSync(join(workspace, path), { recursive: true });
  symlinkSync(findTypedCompiler(compiler, process.env.GOPURS_PURS), join(workspace, 'bin/purs'));
  for (const extension of ['purs', 'go', 'js']) {
    cpSync(join(root, 'test/cancellation.' + extension), join(workspace, 'src/Test/Cancellation.' + extension));
    cpSync(join(root, 'src/Example/HTTP.' + extension), join(workspace, 'src/Example/HTTP.' + extension));
  }
  const overrides = [...new Set([...corePackages(compiler), 'node-process'])];
  const snapshots = join(workspace, '.spago/ports');
  const excluded = new Set(['.git', '.spago', '.cache', '.build', 'node_modules', 'output', 'output.bak']);
  for (const name of overrides) cpSync(join(libraries, 'gopurs-' + name), join(snapshots, 'gopurs-' + name), {
    recursive: true, verbatimSymlinks: true, filter: path => !excluded.has(basename(path)),
  });
  const dependencies = ['aff', 'avar', 'console', 'effect', 'either', 'exceptions', 'node-process', 'prelude'];
  writeFileSync(join(workspace, 'spago.yaml'),
    'package:\n  name: http-cancellation-test\n  dependencies:\n' + dependencies.map(name => `    - ${name}\n`).join('') +
    'workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n' +
    overrides.map(name => `    ${name}:\n      path: ${JSON.stringify(join(snapshots, 'gopurs-' + name))}\n`).join(''));
  const env = {
    ...process.env, PATH: [join(workspace, 'bin'), join(compiler, 'node_modules/.bin'), process.env.PATH].join(delimiter),
    GOPURS_JS: '1', GOPURS_RUST: '0', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0',
    GOPURS_PREPARE_JOBS: '2', GOPURS_PBO_JOBS: '2', GOPURS_EMIT_JOBS: '2', GOWORK: 'off',
  };
  for (const name of ['GOPURS_ALLOC_PROFILE', 'GOPURS_COMPILER_SNAPSHOT', 'GOPURS_FORCE_BUILD', 'PPROF']) delete env[name];
  function run(label, command, args, cwd = workspace, extraEnv = {}) {
    console.log(`HTTP cancellation: ${label}`);
    const result = spawnSync(command, args, { cwd, env: { ...env, ...extraEnv }, encoding: 'utf8', timeout: 180000, maxBuffer: 16 * 1024 * 1024 });
    const log = (result.stdout || '') + (result.stderr || '');
    writeFileSync(join(workspace, 'logs', label + '.log'), log);
    assert.ifError(result.error);
    assert.equal(result.status, 0, `${label}\n${log}`);
  }
  run('frontend', 'spago', ['build', '-q', ...(process.env.GOPURS_OFFLINE === '1' ? ['--offline'] : [])]);
  run('dependencies', process.execPath, [join(compiler, 'tools/library-go-dependencies.mjs'), join(workspace, 'output'),
    ...overrides.map(name => join(snapshots, 'gopurs-' + name))]);
  run('generate', join(compiler, 'bin/gopurs'), ['--release', '--main', 'Test.Cancellation']);
  const binary = join(workspace, 'cancellation-test');
  run('build-race', 'go', ['build', '-race', '-mod=readonly', '-o', binary, './main'], join(workspace, 'output'),
    { GOPROXY: 'off', GONOPROXY: 'none', GOSUMDB: 'off', GOTOOLCHAIN: 'local' });

  let completed = 0;
  for (const disconnected of [false, true]) await t.test(disconnected ? 'disconnected client' : 'connected client', { timeout: 15000 }, async () => {
    const child = spawn(binary, [], { env, stdio: ['pipe', 'pipe', 'pipe'] });
    let output = '';
    let errors = '';
    let spawnError;
    child.stdout.on('data', chunk => { output += chunk; });
    child.stderr.on('data', chunk => { errors += chunk; });
    child.on('error', error => { spawnError = error; });
    const exited = new Promise(resolve => child.once('close', (code, signal) => resolve({ code, signal })));
    async function waitFor(pattern) {
      for (let attempt = 0; attempt < 400; attempt++) {
        assert.ifError(spawnError);
        const match = output.match(pattern);
        if (match) return match;
        assert.equal(child.exitCode, null, `premature exit\n${output}\n${errors}`);
        await delay(25);
      }
      assert.fail(`missing ${pattern}\n${output}\n${errors}`);
    }
    let client;
    try {
      const listening = await waitFor(/LISTENING (http:\/\/127\.0\.0\.1:\d+)/);
      const received = new Promise(resolve => {
        client = request(listening[1] + '/pending', response => {
          let body = '';
          response.on('data', chunk => { body += chunk; });
          response.on('end', () => resolve({ status: response.statusCode, body }));
          response.on('error', error => resolve({ error }));
        });
        client.on('error', error => resolve({ error }));
        client.setTimeout(5000, () => client.destroy(new Error('HTTP cancellation response timed out')));
        client.end();
      });
      await waitFor(/CANCEL REQUESTED/);
      // The handler's stdin gate is still closed. A premature Aff finalizer
      // would appear here, before any test-controlled release of that handler.
      await delay(25);
      assert.doesNotMatch(output, /HANDLER DONE|RESOURCE RELEASED|CANCELED/);
      if (disconnected) {
        client.destroy(new Error('test client disconnect'));
        assert.ok((await received).error);
      }
      child.stdin.end('\n');
      if (!disconnected) assert.deepEqual(await received, { status: 200, body: '{"done":true}' });
      const timer = setTimeout(() => child.kill('SIGKILL'), 10000);
      let result;
      try { result = await exited; } finally { clearTimeout(timer); }
      assert.deepEqual(result, { code: 0, signal: null }, output + errors);
      assert.doesNotMatch(errors, /DATA RACE|panic:|handler failed/);
      const journal = output.trim().split(/\r?\n/).filter(line => !line.startsWith('LISTENING '));
      assert.deepEqual(journal, ['RESOURCE ACQUIRED', 'HANDLER STARTED', 'CANCEL REQUESTED', 'HANDLER DONE', 'RESOURCE RELEASED', 'CANCELED']);
      completed++;
      t.diagnostic(`${disconnected ? 'disconnected' : 'connected'}: ${journal.join(' -> ')}`);
    } finally {
      client?.destroy();
      if (child.exitCode === null && child.signalCode === null) child.kill('SIGKILL');
      await exited;
      writeFileSync(join(workspace, 'logs', disconnected ? 'disconnected.log' : 'connected.log'), output + errors);
    }
  });
  passed = completed === 2;
});
