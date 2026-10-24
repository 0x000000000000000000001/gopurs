// A generic tailRec callback must not copy its growing state on every step.
// Measure allocation growth, not wall time, so a busy machine cannot mask or
// spuriously report the quadratic representation-conversion regression.
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, readFileSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const modes = [{ name: 'js', GOPURS_JS: '1' }, { name: 'native-parallel', GOPURS_JS: '0' }]
  .filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));

test('generic Dev loops retain linear allocation and complete 100000 steps', async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Dev loop workspace: ${workspace}`);
  t.after(() => {
    if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, { recursive: true, force: true });
  });
  const typedBin = join(workspace, 'typed-bin');
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, 'purs'));
  const env = { ...process.env, PATH: [typedBin, join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter),
    GOWORK: 'off', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0', GOPURS_RUST: '0',
    GOPURS_GO_PACKAGES: '1', GOPURS_GO_PACKAGE_OVERLAY: '1', GOPURS_BUILD_MANIFEST: '1',
    GOPURS_PREPARE_JOBS: '2', GOPURS_PBO_JOBS: '2', GOPURS_EMIT_JOBS: '2', GOPURS_PIPELINE: '1' };
  for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_FORCE_BUILD', 'GOPURS_COMPILER_SNAPSHOT', 'PPROF']) delete env[key];
  const packages = corePackages(root);
  for (const [index, mode] of modes.entries()) await t.test(mode.name, () => {
    const { directory } = prepareFixture(root, workspace, join(root, 'tools/fixtures/dev-loop-state.purs'), index, packages);
    const output = join(directory, 'output');
    let sequence = 0;
    const run = (label, command, args, cwd = directory) => {
      const result = spawnSync(command, args, { cwd, env: { ...env, GOPURS_JS: mode.GOPURS_JS },
        encoding: 'utf8', timeout: 120000, maxBuffer: 16 * 1024 * 1024 });
      writeFileSync(join(directory, 'logs', `${sequence++}-${label}.log`), (result.stdout ?? '') + (result.stderr ?? ''));
      assert.ifError(result.error);
      assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
      return result.stdout;
    };
    run('frontend', 'spago', ['build', '-q', '--offline']);
    const reference = run('javascript', 'node', ['--input-type=module', '-e', 'import { main } from "./output/Main/index.js"; main();']);
    assert.equal(reference, '100000\n100000\n');
    run('dev', process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs'), ['--dev', '--main', 'Main']);
    const proof = JSON.parse(JSON.parse(readFileSync(join(output, '.gopurs-cache/profiles/dev/preparation-literals-v1/proof.json'))).payload);
    assert.deepEqual(proof.index, [], 'linear Dev execution must not depend on monomorphization');
    mkdirSync(join(output, 'allocation-probe'));
    writeFileSync(join(output, 'allocation-probe/main.go'), `package main
import (
  "encoding/json"
  "os"
  "runtime"
  p "gopurs/output/purescript/p_root"
  _ "gopurs/output/purescript/init_all"
)
type measurement struct { N int64; Bytes uint64; Count int64; Sum int64 }
func measure(n int64) measurement {
  runtime.GC()
  var before, after runtime.MemStats
  runtime.ReadMemStats(&before)
  result := p.Call_Main_buildList(n)
  runtime.ReadMemStats(&after)
  m := measurement{N:n, Bytes:after.TotalAlloc-before.TotalAlloc}
  for cell := result; cell != nil; cell = cell.V1 { m.Count++; m.Sum += cell.V0 }
  runtime.KeepAlive(result)
  return m
}
func main() {
  measure(10)
  // Fail fast on the old quadratic code before attempting the large case.
  small, large := measure(1000), measure(2000)
  json.NewEncoder(os.Stdout).Encode([]measurement{small, large})
  if large.Bytes > small.Bytes*5/2 { os.Exit(1) }
  json.NewEncoder(os.Stdout).Encode([]measurement{measure(100000)})
}
`);
    run('probe-build', 'go', ['build', '-o', 'allocation-probe-bin', './allocation-probe'], output);
    const measurements = run('probe', join(output, 'allocation-probe-bin'), [], output).trim().split('\n').flatMap(line => JSON.parse(line));
    for (const m of measurements) {
      assert.equal(m.Count, m.N);
      assert.equal(m.Sum, m.N);
    }
    assert.equal(measurements.length, 3);
    assert.ok(measurements[2].Bytes < 128 * 1024 * 1024, JSON.stringify(measurements));
    assert.equal(run('go-reference-parity', 'go', ['run', './main'], output), reference);
    console.log(`${mode.name}: ${JSON.stringify(measurements)}`);
    completed++;
  });
});
