import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, readFileSync, readdirSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './build/native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const modes = [{ name: 'js', GOPURS_JS: '1' }, { name: 'native-parallel', GOPURS_JS: '0' }]
  .filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));
// The frozen pre-change Go compiler already interleaves array callback factory
// construction and execution; JS constructs all three factories first. Keep
// the complete Go trace stable, and compare thunk effects across both backends.
const trace = 'FBTA'.repeat(3) + 'FBA' + 'FBTA'.repeat(10) + 'FBA' + 'TT';
const expected = ['42', '42', '42', '7', '9', '3', '5', '7', '3', '5', '7', '11', '21', '42', 'hello', '42', '42', trace, ''].join('\n');
const withoutFactories = output => {
  const lines = output.trimEnd().split('\n');
  lines[lines.length - 1] = lines[lines.length - 1].replaceAll('F', '');
  return lines.join('\n');
};

test('native Effect Int callbacks preserve aliases, retained captures and all ordinary fallback entries', async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Native effect callback workspace: ${workspace}`);
  t.after(() => {
    if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, { recursive: true, force: true });
  });
  const typedBin = join(workspace, 'typed-bin');
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, 'purs'));
  const env = { ...process.env, PATH: [typedBin, join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter),
    GOWORK: 'off', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0', GOPURS_RUST: '0',
    GOPURS_GO_PACKAGES: '1', GOPURS_GO_PACKAGE_OVERLAY: '1', GOPURS_BUILD_MANIFEST: '1',
    GOPURS_PREPARE_JOBS: '2', GOPURS_PBO_JOBS: '2', GOPURS_EMIT_JOBS: '2', GOPURS_PIPELINE: '1', GOPURS_FORCE_BUILD: '1' };
  for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_COMPILER_SNAPSHOT', 'PPROF']) delete env[key];
  const packages = corePackages(root);
  for (const [index, mode] of modes.entries()) await t.test(mode.name, () => {
    const { directory } = prepareFixture(root, workspace, join(root, 'tools/fixtures/ffi-native-effect-callbacks.purs'), index, packages);
    let sequence = 0;
    const run = (label, command, args, cwd = directory) => {
      console.log(`${mode.name}: ${label}`);
      const result = spawnSync(command, args, { cwd, env: { ...env, GOPURS_JS: mode.GOPURS_JS },
        encoding: 'utf8', timeout: 120000, maxBuffer: 16 * 1024 * 1024 });
      writeFileSync(join(directory, 'logs', `${sequence++}-${label}.log`), (result.stdout ?? '') + (result.stderr ?? ''));
      assert.ifError(result.error);
      assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
      return result.stdout;
    };
    run('frontend', 'spago', ['build', '-q', '--offline']);
    const javascript = run('javascript', 'node', ['--input-type=module', '-e', 'import { main } from "./output/Main/index.js"; main();']);
    assert.equal(withoutFactories(javascript), withoutFactories(expected));
    for (const profile of ['release', 'dev']) {
      run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs'), [`--${profile}`, '--main', 'Main']);
      const generated = join(directory, 'output/purescript');
      const source = readdirSync(generated, { recursive: true }).filter(name => name.endsWith('.go'))
        .map(name => readFileSync(join(generated, name), 'utf8')).join('\n');
      assert.ok(/cache_Main_cached = \w+\.WithFunctionData\(/.test(source), `${profile}: cached action carries native metadata`);
      assert.ok(/FunctionData\[func\(\) int64\]/.test(source), `${profile}: exact native adapter emitted`);
      assert.equal(run(`${profile}-go`, 'go', ['run', './main'], join(directory, 'output')), expected);
    }
    completed++;
  });
});
