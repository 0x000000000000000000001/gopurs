// A recursive worker returning a function must retain the function's remaining
// arguments. Its result is not the final scalar codomain of that function.
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const modes = [{ name: 'js', GOPURS_JS: '1' }, { name: 'native-parallel', GOPURS_JS: '0' }]
  .filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));

test('recursive function lookup preserves callable results in Release and Dev', async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Recursive function workspace: ${workspace}`);
  t.after(() => {
    if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, { recursive: true, force: true });
  });
  const typedBin = join(workspace, 'typed-bin');
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, 'purs'));
  const env = { ...process.env, PATH: [typedBin, join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter),
    GOWORK: 'off', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0', GOPURS_RUST: '0',
    GOPURS_GO_PACKAGES: '1', GOPURS_GO_PACKAGE_OVERLAY: '1', GOPURS_BUILD_MANIFEST: '1',
    GOPURS_PREPARE_JOBS: '2', GOPURS_PBO_JOBS: '2', GOPURS_EMIT_JOBS: '2', GOPURS_PIPELINE: '1',
    GOPURS_FORCE_BUILD: '1' };
  for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_COMPILER_SNAPSHOT', 'PPROF']) delete env[key];
  const packages = corePackages(root);
  for (const [index, mode] of modes.entries()) await t.test(mode.name, () => {
    const { directory } = prepareFixture(root, workspace, join(root, 'tools/fixtures/recursive-function-result.purs'), index, packages);
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
    assert.equal(reference, 'true\nfalse\ntrue\n"42"\n"value=42"\nLT\nGT\n');
    for (const profile of ['release', 'dev']) {
      run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs'), [`--${profile}`, '--main', 'Main']);
      assert.equal(run(`${profile}-go`, 'go', ['run', './main'], join(directory, 'output')), reference);
    }
    completed++;
  });
});
