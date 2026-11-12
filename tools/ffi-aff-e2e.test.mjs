import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, readFileSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './build/native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const modes = [{ name: 'js', GOPURS_JS: '1' }, { name: 'native-parallel', GOPURS_JS: '0' }]
  .filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));
const expected = '0\nhello Go\nhello Go\n2\nexpected error\nplain result\neffect without Unit\ncallback=typed\ncontext canceled\n0\n';

test('ordinary Go FFI preserves EffectFnAff semantics across hosts and profiles', async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Idiomatic Aff workspace: ${workspace}`);
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
  for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_COMPILER_SNAPSHOT', 'GOPURS_FORCE_BUILD', 'PPROF']) delete env[key];
  const packages = corePackages(root);
  const source = join(root, 'tools/fixtures/idiomatic-aff.purs');

  // The handwritten implementation remains a normal Go package with no backend
  // import. Exercise it separately before the generated adapter is involved.
  const ordinary = join(workspace, 'ordinary');
  mkdirSync(ordinary);
  copyFileSync(source.replace(/\.purs$/, '.go'), join(ordinary, 'fixture.go'));
  assert.doesNotMatch(readFileSync(join(ordinary, 'fixture.go'), 'utf8'), /gopurs_runtime/);
  writeFileSync(join(ordinary, 'go.mod'), 'module fixture\n\ngo 1.23\n');
  writeFileSync(join(ordinary, 'fixture_test.go'), `package fixture
import ("context"; "testing")
func TestOrdinary(t *testing.T) {
  if value, err := Hello("Go"); value != "hello Go" || err != nil { t.Fatal(value, err) }
  if _, err := Hello("error"); err == nil { t.Fatal("missing error") }
  var callback string; RunCallback(func(value string) { callback = value }); if callback != "typed" { t.Fatal(callback) }
  ctx, cancel := context.WithCancel(context.Background()); cancel(); if _, err := Pending(ctx); err != context.Canceled { t.Fatal(err) }
}
`);
  const direct = spawnSync('go', ['test', '-race', '-count=1', './...'], { cwd: ordinary, env, encoding: 'utf8', timeout: 60000 });
  assert.ifError(direct.error);
  assert.equal(direct.status, 0, direct.stdout + direct.stderr);

  const makeRunner = (directory, mode) => {
    let sequence = 0;
    return (label, command, args, cwd = directory) => {
      const result = spawnSync(command, args, { cwd, env: { ...env, GOPURS_JS: mode.GOPURS_JS },
        encoding: 'utf8', timeout: 120000, maxBuffer: 16 * 1024 * 1024 });
      writeFileSync(join(directory, 'logs', `${sequence++}-${label}.log`), (result.stdout ?? '') + (result.stderr ?? ''));
      assert.ifError(result.error);
      assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
      return result.stdout;
    };
  };
  // The Go Aff port deliberately has native-only internals. The JavaScript
  // reference uses the official package with the same public PureScript API.
  const reference = prepareFixture(root, workspace, source, 'reference', packages).directory;
  const referenceConfig = join(reference, 'spago.yaml');
  writeFileSync(referenceConfig, readFileSync(referenceConfig, 'utf8').replace(/^    aff:\n      path: .*\n/m, ''));
  const referenceRun = makeRunner(reference, { GOPURS_JS: '1' });
  referenceRun('frontend', 'spago', ['build', '-q', '--offline']);
  assert.equal(referenceRun('javascript-reference', 'node', ['--input-type=module', '-e', 'import { main } from "./output/Main/index.js"; main();']), expected);

  for (const [index, mode] of modes.entries()) await t.test(mode.name, () => {
    const { directory } = prepareFixture(root, workspace, source, index, packages);
    const run = makeRunner(directory, mode);
    run('frontend', 'spago', ['build', '-q', '--offline']);
    for (const profile of ['release', 'dev']) {
      run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs'), [`--${profile}`, '--main', 'Main']);
      assert.equal(run(`${profile}-go`, 'go', ['run', './main'], join(directory, 'output')), expected);
      run(`${profile}-cached`, process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs'), [`--${profile}`, '--main', 'Main']);
      assert.equal(run(`${profile}-cached-go`, 'go', ['run', './main'], join(directory, 'output')), expected);
    }
    completed++;
  });
});
