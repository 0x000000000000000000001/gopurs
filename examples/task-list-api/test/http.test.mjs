import assert from 'node:assert/strict';
import { spawn } from 'node:child_process';
import { once } from 'node:events';
import { mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { createServer } from 'node:net';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { setTimeout as delay } from 'node:timers/promises';
import { fileURLToPath } from 'node:url';

const binary = fileURLToPath(new URL('../.build/task-list-api', import.meta.url));

async function start(t, database) {
  const child = spawn(binary, [], { env: { ...process.env, TASK_LIST_ADDRESS: '127.0.0.1:0', TASK_LIST_DATABASE: database }, stdio: ['ignore', 'pipe', 'pipe'] });
  let output = '';
  let errors = '';
  child.stdout.on('data', chunk => { output += chunk; });
  child.stderr.on('data', chunk => { errors += chunk; });
  const exited = once(child, 'exit');
  t.after(() => { if (child.exitCode === null && child.signalCode === null) child.kill('SIGKILL'); });
  let url;
  for (let attempt = 0; attempt < 400; attempt++) {
    const match = output.match(/LISTENING (http:\/\/127\.0\.0\.1:\d+)/);
    if (match) { url = match[1]; break; }
    if (child.exitCode !== null || child.signalCode !== null) assert.fail(`server exited before listening\n${output}\n${errors}`);
    await delay(25);
  }
  assert.ok(url, `server did not start\n${output}\n${errors}`);
  return { url, async stop() {
    child.kill('SIGTERM');
    const timer = setTimeout(() => child.kill('SIGKILL'), 10000);
    try {
      const [code, signal] = await exited;
      assert.equal(signal, null, errors);
      assert.equal(code, 0, errors);
      assert.match(output, /STOPPED/);
      assert.doesNotMatch(errors, /DATA RACE|panic:/);
    } finally { clearTimeout(timer); }
  } };
}

async function request(url, path, method = 'GET', body) {
  const response = await fetch(url + path, { method,
    ...(body === undefined ? {} : { body: typeof body === 'string' ? body : JSON.stringify(body), headers: { 'content-type': 'application/json' } }),
    signal: AbortSignal.timeout(5000) });
  const text = await response.text();
  return { status: response.status, body: text ? JSON.parse(text) : null };
}

test('compiled PureScript Task list API persists data, validates input and stops cleanly', { timeout: 45000 }, async t => {
  const workspace = mkdtempSync(join(tmpdir(), 'gopurs-task-http-'));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  const database = join(workspace, 'tasks.db');
  let server = await start(t, database);
  assert.deepEqual(await request(server.url, '/health'), { status: 200, body: { status: 'ok' } });
  assert.deepEqual(await request(server.url, '/tasks'), { status: 200, body: [] });
  for (const body of ['{', {}, { title: '   ' }, { title: 123 }, { title: 'x'.repeat(201) }]) {
    assert.equal((await request(server.url, '/tasks', 'POST', body)).status, 400);
  }
  const created = await request(server.url, '/tasks', 'POST', { title: "Learn gopurs — '); DROP TABLE tasks; --" });
  assert.equal(created.status, 201);
  assert.match(created.body.id, /^\d+$/);
  assert.equal(created.body.completed, false);
  const path = '/tasks/' + created.body.id;
  assert.deepEqual((await request(server.url, path)).body, created.body);
  assert.equal((await request(server.url, path, 'PATCH', { completed: 'yes' })).status, 400);
  assert.deepEqual((await request(server.url, path, 'PATCH', { completed: true })).body, { ...created.body, completed: true });
  assert.equal((await request(server.url, '/tasks/999999999')).status, 404);
  assert.equal((await request(server.url, '/missing')).status, 404);
  assert.equal((await request(server.url, path, 'PUT', {})).status, 405);

  const batch = await Promise.all(Array.from({ length: 8 }, (_, index) => request(server.url, '/tasks', 'POST', { title: 'Concurrent ' + index })));
  assert.ok(batch.every(value => value.status === 201));
  assert.equal(new Set(batch.map(value => value.body.id)).size, 8);
  await server.stop();
  server = await start(t, database);
  assert.deepEqual((await request(server.url, path)).body, { ...created.body, completed: true });
  assert.equal((await request(server.url, '/tasks')).body.length, 9);
  assert.equal((await request(server.url, path, 'DELETE')).status, 204);
  assert.equal((await request(server.url, path)).status, 404);
  assert.equal((await request(server.url, path, 'DELETE')).status, 404);
  await server.stop();
});

test('startup failures report an error and exit unsuccessfully', { timeout: 15000 }, async t => {
  const workspace = mkdtempSync(join(tmpdir(), 'gopurs-task-startup-'));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  const blocker = join(workspace, 'not-a-directory');
  writeFileSync(blocker, 'file');
  const occupied = createServer();
  occupied.listen(0, '127.0.0.1');
  await once(occupied, 'listening');
  t.after(() => occupied.close());
  const cases = [
    { TASK_LIST_DATABASE: join(blocker, 'tasks.db'), TASK_LIST_ADDRESS: '127.0.0.1:0' },
    { TASK_LIST_DATABASE: ':memory:', TASK_LIST_ADDRESS: `127.0.0.1:${occupied.address().port}` },
  ];
  for (const environment of cases) {
    const child = spawn(binary, [], { env: { ...process.env, ...environment }, stdio: ['ignore', 'pipe', 'pipe'] });
    let output = '';
    let errors = '';
    child.stdout.on('data', chunk => { output += chunk; });
    child.stderr.on('data', chunk => { errors += chunk; });
    const timer = setTimeout(() => child.kill('SIGKILL'), 5000);
    t.after(() => { if (child.exitCode === null && child.signalCode === null) child.kill('SIGKILL'); });
    try {
      const [code, signal] = await once(child, 'close');
      assert.equal(signal, null, output + errors);
      assert.equal(code, 1, output + errors);
      assert.ok(errors.trim(), 'startup error must be logged');
      assert.doesNotMatch(output, /LISTENING|STOPPED/);
      assert.doesNotMatch(errors, /DATA RACE|panic:/);
    } finally { clearTimeout(timer); }
  }
});
