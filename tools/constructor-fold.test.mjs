import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { existsSync, mkdirSync, readFileSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './build/native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const modes = [
  {name: 'js', GOPURS_JS: '1'}, {name: 'native-parallel', GOPURS_JS: '0'},
].filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));
const expected = [
  [7, 7, 7, 21, 51, 163, 663, 3327, 19975, 139839, 1118727],
  [7, 7, 7, 10, 3, -4, -11, -18, -25, -32, -39],
  [69280, -80641, 2147403007, -2147414368],
  1118688, 3, 46165, 72, [1, 2, 3, 4],
];
const parse = stdout => stdout.trim().split('\n').map(line => JSON.parse(line));

test('constructor folding preserves nested trees, arithmetic, strict inputs and opaque fallbacks', async t => {
  assert.ok(modes.length);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Constructor-fold fixture workspace: ${workspace}`);
  t.after(() => {
    if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, {recursive: true, force: true});
  });
  const typedBin = join(workspace, 'typed-bin');
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, 'purs'));
  const env = {
    ...process.env, PATH: [typedBin, join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter),
    GOWORK: 'off', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0', GOPURS_RUST: '0',
    GOPURS_GO_PACKAGES: '1', GOPURS_GO_PACKAGE_OVERLAY: '1', GOPURS_BUILD_MANIFEST: '1',
    GOPURS_PREPARE_JOBS: '2', GOPURS_PBO_JOBS: '2', GOPURS_EMIT_JOBS: '2',
    GOPURS_PIPELINE: '1', GOPURS_FORCE_BUILD: '1',
  };
  for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_COMPILER_SNAPSHOT', 'PPROF']) delete env[key];
  for (const [index, mode] of modes.entries()) await t.test(mode.name, () => {
    const {directory} = prepareFixture(root, workspace, join(root, 'tools/fixtures/constructor-fold.purs'), index, corePackages(root));
    let sequence = 0;
    const run = (label, command, args, cwd = directory) => {
      console.log(`${mode.name}: ${label}`);
      const result = spawnSync(command, args, {cwd, env: {...env, GOPURS_JS: mode.GOPURS_JS}, encoding: 'utf8', timeout: 120000, maxBuffer: 16 * 1024 * 1024});
      writeFileSync(join(directory, 'logs', `${sequence++}-${label}.log`), (result.stdout ?? '') + (result.stderr ?? ''));
      assert.ifError(result.error);
      assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
      return result.stdout;
    };
    run('frontend', 'spago', ['build', '-q', '--offline']);
    assert.deepEqual(parse(run('javascript', 'node', ['--input-type=module', '-e', 'import { main } from "./output/Main/index.js"; main();'])), expected);
    for (const profile of ['release', 'dev']) {
      run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs'), [`--${profile}`, '--main', 'Main']);
      const candidates = ['purescript/p_root/0_gopurs_package.go', 'purescript/Main.go'].map(path => join(directory, 'output', path));
      const generated = readFileSync(candidates.find(existsSync), 'utf8');
      assert.match(generated, /func Call_Main_plant__gopurs_constructor_fold_/, 'fixture must exercise the new scalar workers');
      assert.doesNotMatch(generated, /func Call_Main_plantWith__gopurs_constructor_fold_/, 'opaque producer must retain its original path');
      assert.deepEqual(parse(run(`${profile}-go`, 'go', ['run', './main'], join(directory, 'output'))), expected);
    }
    completed++;
  });
});
