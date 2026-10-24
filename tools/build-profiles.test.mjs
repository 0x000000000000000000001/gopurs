import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { spawnSync } from 'node:child_process';
import { cpSync, existsSync, mkdirSync, readFileSync, readdirSync, rmSync, statSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './native-workspace.mjs';
import { sourceVariants } from './cache-source-variants.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const hash = value => createHash('sha256').update(value).digest('hex');
const baseEnv = { ...process.env, GOWORK: 'off', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0', GOPURS_RUST: '0' };
for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_FORCE_BUILD', 'GOPURS_COMPILER_SNAPSHOT', 'PPROF']) delete baseEnv[key];
const modes = [
  { name: 'js', GOPURS_JS: '1', jobs: '1' },
  { name: 'native-parallel', GOPURS_JS: '0', jobs: '8' },
].filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));

function generated(output, prefix = '', files = {}) {
  for (const entry of readdirSync(join(output, prefix), { withFileTypes: true })) {
    if (entry.name.startsWith('.')) continue;
    const name = prefix + entry.name;
    if (entry.isDirectory()) generated(output, name + '/', files);
    else if (name.endsWith('.go')) files[name] = hash(readFileSync(join(output, name)));
  }
  return files;
}

test('dev and release preserve semantics, coexist, and match clean builds after source changes', async t => {
  assert.ok(modes.length > 0, 'select at least one supported host mode');
  const workspace = createWorkspace();
  let passed = false;
  let completedModes = 0;
  t.after(() => {
    if (passed && process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, { recursive: true, force: true });
    else console.log(`Build profiles workspace: ${workspace}`);
  });
  const typedBin = join(workspace, 'typed-bin');
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, 'purs'));
  const env = { ...baseEnv, PATH: [typedBin, join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter),
    GOPURS_GO_PACKAGES: '1', GOPURS_GO_PACKAGE_OVERLAY: '1', GOPURS_BUILD_MANIFEST: '1',
    GOPURS_PIPELINE: '1', GOGC: 'off', GOMEMLIMIT: '4GiB' };
  const fixtureRoot = join(root, 'tests/host/CacheTransitions');
  const states = sourceVariants(fixtureRoot);
  const { directory } = prepareFixture(root, workspace, join(fixtureRoot, 'Main.purs'), 0, corePackages(root));
  const output = join(directory, 'output');
  const logs = join(workspace, 'logs');
  mkdirSync(logs);
  console.log(`Build profiles workspace: ${workspace}`);
  let sequence = 0;
  function run(label, command, args, cwd = directory, extra = {}, expected = 0) {
    const result = spawnSync(command, args, { cwd, env: { ...env, ...extra }, encoding: 'utf8',
      timeout: 120000, maxBuffer: 16 * 1024 * 1024 });
    writeFileSync(join(logs, `${sequence++}-${label}.log`), (result.stdout ?? '') + (result.stderr ?? ''));
    assert.ifError(result.error);
    assert.equal(result.status, expected, `${label}\n${result.stdout}\n${result.stderr}`);
    return result;
  }
  function writeSources(state) {
    rmSync(join(directory, 'src'), { recursive: true, force: true });
    for (const [name, source] of Object.entries(state.sources)) {
      const target = join(directory, 'src', name);
      mkdirSync(dirname(target), { recursive: true });
      const text = name === 'Main.purs' ? source.replace('import Prelude', 'import Prelude\nimport Profile.FastMath as FastMath\nimport Profile.Normalization as Normalization')
        .replace('  log "Done"', '  logShow (FastMath.step seed)\n  logShow (Normalization.discarded seed)\n  logShow (Normalization.branch true seed)\n  effectResult <- Normalization.effectOrder\n  logShow effectResult\n  log "Done"') : source;
      writeFileSync(target, text);
    }
    mkdirSync(join(directory, 'src/Profile'), { recursive: true });
    writeFileSync(join(directory, 'src/Profile/FastMath.purs'), 'module Profile.FastMath where\nimport Prelude\nstep :: Int -> Int\nstep x = x * 3 + 1\n');
    writeFileSync(join(directory, 'src/Profile/Normalization.purs'), `module Profile.Normalization where
import Prelude
import Effect (Effect)
import Effect.Ref as Ref
import Profile.FastMath as FastMath

discarded :: Int -> Int
discarded x = let y = FastMath.step x in (\\a b -> 1) y y

branch :: Boolean -> Int -> Int
branch flag x =
  let f = if flag then (\\n -> { a: n + 1, b: n }) else (\\n -> { a: n, b: n + 1 })
  in (f x).a

effectOrder :: Effect Int
effectOrder = do
  ref <- Ref.new 0
  let change n = do
        before <- Ref.read ref
        Ref.write (before * 10 + n) ref
        pure n
  _ <- change 1
  _ <- change 2
  Ref.read ref
`);
  }
  writeSources(states[0]);
  run('frontend-initial', 'spago', ['build', '-q', '--offline']);
  const tast = join(workspace, 'tast');
  cpSync(output, tast, { recursive: true });
  for (const mode of modes) await t.test(mode.name, () => {
    rmSync(output, { recursive: true, force: true });
    cpSync(tast, output, { recursive: true });
    writeSources(states[0]);
    const host = { GOPURS_JS: mode.GOPURS_JS, GOPURS_PREPARE_JOBS: mode.jobs,
      GOPURS_PBO_JOBS: mode.jobs, GOPURS_EMIT_JOBS: mode.jobs };
    const compiler = process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs');
    const cachePaths = { release: join(output, '.gopurs-cache'), dev: join(output, '.gopurs-cache/profiles/dev') };
    let compilerIdentity;
    const build = (label, flags = [], extra = {}, expected = 0) => {
      const result = run(`${mode.name}-${label}`, compiler, [...flags, '--main', 'Main'], directory, { ...host, ...extra }, expected);
      if (expected === 0) {
        const cache = cachePaths[flags.includes('--dev') ? 'dev' : 'release'];
        const pointer = JSON.parse(readFileSync(join(cache, 'current.json')));
        const bytes = readFileSync(join(cache, 'objects', pointer.manifest.sha256));
        assert.equal(hash(bytes), pointer.manifest.sha256, `${label}: committed manifest integrity`);
        const manifest = JSON.parse(bytes);
        compilerIdentity ??= manifest.snapshot.compiler;
        assert.equal(manifest.snapshot.compiler, compilerIdentity,
          `${label}: compiler or helper changed during the test; rerun with a stable installation`);
        if (flags.includes('--force')) {
          assert.match(result.stderr, /cache miss: forced/);
          const reports = [...result.stderr.matchAll(/module cache: optimized=(\d+), reused=(\d+), emitted=(\d+), restored=(\d+)/g)];
          assert.equal(reports.length, 1, `${label}: one completed module-cache report`);
          const count = readdirSync(output, { withFileTypes: true }).filter(entry =>
            entry.isDirectory() && existsSync(join(output, entry.name, 'corefn.json'))).length;
          assert.ok(count > 0);
          assert.deepEqual(reports[0].slice(1).map(Number), [count, 0, count, 0],
            `${label}: the full oracle must optimize and emit every module without reuse`);
        }
      }
      return result;
    };
    const execute = (label, state) => {
      const expected = state.stdout.replace('Done\n', '16\n1\n6\n12\nDone\n');
      assert.equal(run(`${mode.name}-${label}-execute`, 'go', ['run', './main'], output).stdout, expected);
    };
    const rejected = build('conflicting-flags', ['--dev', '--release'], {}, 1);
    assert.match(rejected.stderr, /Choose either --dev or --release/);
    assert.ok(!existsSync(join(output, '.gopurs-cache/current.json')));
    build('default');
    const releaseBase = generated(output);
    const explicit = build('explicit-release', ['--release']);
    assert.match(explicit.stderr, /cache hit:/);
    assert.deepEqual(generated(output), releaseBase);
    execute('release-initial', states[0]);
    build('dev', ['--dev']);
    const devBase = generated(output);
    assert.notDeepEqual(devBase, releaseBase, 'dev must change optimization, not just select another cache');
    execute('dev-initial', states[0]);
    for (const path of Object.values(cachePaths)) assert.ok(existsSync(join(path, 'current.json')));
    const lockPath = join(output, '.gopurs-cache/write.lock');
    const lockInode = statSync(lockPath).ino;
    for (const [profile, expected] of [['release', releaseBase], ['dev', devBase], ['release', releaseBase]]) {
      const transition = build(`restore-${profile}`, [`--${profile}`]);
      assert.deepEqual(generated(output), expected);
      assert.match(transition.stderr, /module cache: optimized=0, reused=\d+, emitted=0, restored=\d+/);
      const timestamps = Object.fromEntries(Object.keys(expected).map(name => [name, statSync(join(output, name)).mtimeMs]));
      const noop = build(`noop-${profile}`, [`--${profile}`]);
      assert.match(noop.stderr, /cache hit:/);
      assert.deepEqual(Object.fromEntries(Object.keys(expected).map(name => [name, statSync(join(output, name)).mtimeMs])), timestamps);
      assert.equal(statSync(lockPath).ino, lockInode);
    }
    for (const state of states.slice(1)) {
      writeSources(state);
      run(`${mode.name}-${state.label}-frontend`, 'spago', ['build', '-q', '--offline']);
      for (const profile of ['dev', 'release']) {
        build(`${state.label}-${profile}`, [`--${profile}`]);
        const incremental = generated(output);
        execute(`${state.label}-${profile}`, state);
        build(`${state.label}-${profile}-full`, [`--${profile}`, '--force'], { GOPURS_GO_PACKAGE_OVERLAY: '0' });
        assert.deepEqual(generated(output), incremental, `${state.label}/${profile}: incremental must equal full generation`);
        const noop = build(`${state.label}-${profile}-noop`, [`--${profile}`]);
        assert.match(noop.stderr, /cache hit:/);
      }
    }
    completedModes += 1;
  });
  // node:test reports failed subtests without rejecting t.test(). Keep their
  // fixture and logs instead of mistaking the resolved promise for success.
  passed = completedModes === modes.length;
});
