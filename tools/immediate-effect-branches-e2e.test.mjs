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


function emittedFunction(source, name) {
  const start = source.indexOf(`func ${name}(`);
  if (start < 0) return '';
  const end = source.indexOf('\nfunc ', start + 1);
  return source.slice(start, end < 0 ? undefined : end);
}

test('immediate effect branches preserve source traces, saved actions, nested effects and thrown conditions', async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  console.log(`Immediate effect branches workspace: ${workspace}`);
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
    const { directory } = prepareFixture(root, workspace, join(root, 'tools/fixtures/immediate-effect-branches/Main.purs'), index, packages);
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
    const expected = run('javascript', 'node', ['--input-type=module', '-e', 'import { main } from "./output/Main/index.js"; main();']);
    assert.match(expected, /true\nX\n/, "a throwing condition leaves only its own trace");
    assert.match(expected, /false\nXFZEZ\n/, "a nonthrowing false condition constructs and executes the fallback");
    for (const profile of ['release', 'dev']) {
      run(profile, process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs'), [`--${profile}`, '--main', 'Main']);
      const generated = join(directory, 'output/purescript');
      const source = readdirSync(generated, { recursive: true }).filter(name => name.endsWith('.go'))
        .map(name => readFileSync(join(generated, name), 'utf8')).join('\n');
      const callers = Object.fromEntries(['choose', 'nested', 'makeSaved', 'returned', 'failing'].map(name => [name,
        emittedFunction(source, `Get_Main_${name}`) + emittedFunction(source, `Call_Main_${name}`)]));
      assert.ok(callers.choose.length > 0);
      // Inspect the actual source worker: a selected pure branch should assign
      // its result; only the outer saved action itself remains a Func.
      assert.match(callers.choose, /(?:r_|gopurs_runtime)\.Int\(int64\(7\)\)/);
      const chooseWorker = emittedFunction(source, "Call_Main_choose");
      assert.equal((chooseWorker.match(/\.Func\(/g) ?? []).length, 1, "only the saved outer action remains a closure");
      writeFileSync(join(directory, 'logs', `${profile}-callers.json`), JSON.stringify(callers, null, 2) + '\n');
      assert.equal(run(`${profile}-go`, 'go', ['run', './main'], join(directory, 'output')), expected);
    }
    completed++;
  });
});
