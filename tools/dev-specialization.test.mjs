// A small native-code gate for dev builds without specialization.
// GOPURS_TEST_MODES=js or native-parallel selects an individual compiler host.
import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { spawnSync } from 'node:child_process';
import { mkdirSync, readFileSync, readdirSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './native-workspace.mjs';
import { hash as wireHash } from '../output/Gopurs.Cache.Wire/index.js';

const root = fileURLToPath(new URL('../', import.meta.url));
const hash = value => createHash('sha256').update(value).digest('hex');
const modes = [{ name: 'js', GOPURS_JS: '1' }, { name: 'native-parallel', GOPURS_JS: '0' }]
  .filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));
const helpers = ['Profile.Helpers.roundTrip', 'Profile.Helpers.consume'];

function generated(directory, prefix = '', result = {}) {
  for (const entry of readdirSync(join(directory, prefix), { withFileTypes: true })) {
    if (entry.name.startsWith('.')) continue;
    const name = prefix + entry.name;
    if (entry.isDirectory()) generated(directory, name + '/', result);
    else if (name.endsWith('.go')) result[name] = hash(readFileSync(join(directory, name)));
  }
  return result;
}

function mainSource(seed, extra) {
  return `module Main where
-- @dependencies: prelude effect console refs
import Prelude
import Effect (Effect)
import Effect.Console (logShow)
import Effect.Ref as Ref
import Profile.Box (Box(..))
import Profile.Relay as Relay

main :: Effect Unit
main = do
  cell <- Ref.new ${seed}
  n <- Ref.read cell
  let ints = Box n { value: n + 1 }
  let bools = Box (n > 0) { value: false }
  case Relay.pass ints of
    Box left fields -> logShow (left + fields.value)
  case Relay.pass bools of
    Box left fields -> logShow ((if left then 1 else 0) + (if fields.value then 2 else 0))
  logShow (Relay.score ints)
  logShow (Relay.score bools)
  case Relay.pass (Box { count: n } { value: { count: n + 1 } }) of
    Box left fields -> logShow (left.count + fields.value.count)
${extra ? '  case Relay.pass (Box "left" { value: "right" }) of\n    Box left fields -> logShow (if left == "right" && fields.value == "left" then 17 else 0)\n' : ''}`;
}

test('dev retains typed generic fallbacks across ADTs, records, dictionaries and source edits', async t => {
  assert.ok(modes.length > 0);
  const workspace = createWorkspace();
  let completed = 0;
  t.after(() => {
    if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, { recursive: true, force: true });
    else console.log(`Dev specialization workspace: ${workspace}`);
  });
  console.log(`Dev specialization workspace: ${workspace}`);
  const typedBin = join(workspace, 'typed-bin');
  mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, 'purs'));
  const fixture = join(workspace, 'Main.purs');
  writeFileSync(fixture, mainSource(5, false));
  const packages = corePackages(root);
  const env = { ...process.env, PATH: [typedBin, join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter),
    GOWORK: 'off', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0', GOPURS_RUST: '0',
    GOPURS_GO_PACKAGES: '1', GOPURS_GO_PACKAGE_OVERLAY: '1', GOPURS_BUILD_MANIFEST: '1',
    GOPURS_PREPARE_JOBS: '2', GOPURS_PBO_JOBS: '2', GOPURS_EMIT_JOBS: '2', GOPURS_PIPELINE: '1',
    GOGC: 'off', GOMEMLIMIT: '4GiB' };
  for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_FORCE_BUILD', 'GOPURS_COMPILER_SNAPSHOT', 'PPROF']) delete env[key];

  for (const [index, mode] of modes.entries()) await t.test(mode.name, () => {
    const { directory } = prepareFixture(root, workspace, fixture, index, packages);
    const output = join(directory, 'output'), source = join(directory, 'src');
    mkdirSync(join(source, 'Profile'));
    writeFileSync(join(source, 'Profile/Box.purs'), `module Profile.Box where
import Prelude
data Box a = Box a { value :: a }
class Mark a where
  mark :: a -> Int
instance markInt :: Mark Int where
  mark value = value
instance markBoolean :: Mark Boolean where
  mark value = if value then 10 else 0
`);
    writeFileSync(join(source, 'Profile/Helpers.purs'), `-- @inline export roundTrip never
-- @inline export consume never
module Profile.Helpers where
import Prelude
import Profile.Box (Box(..), class Mark, mark)
roundTrip :: forall a. Box a -> Box a
roundTrip (Box left fields) = Box fields.value { value: left }
consume :: forall a. Mark a => Box a -> Int
consume (Box left fields) = mark left + mark fields.value
`);
    writeFileSync(join(source, 'Profile/Relay.purs'), `-- @inline export pass never
-- @inline export score never
module Profile.Relay where
import Profile.Box (Box, class Mark)
import Profile.Helpers as Helpers
pass :: forall a. Box a -> Box a
pass value = Helpers.roundTrip value
score :: forall a. Mark a => Box a -> Int
score value = Helpers.consume value
`);
    let sequence = 0, compilerIdentity;
    const run = (label, command, args, cwd = directory) => {
      const result = spawnSync(command, args, { cwd, env: { ...env, GOPURS_JS: mode.GOPURS_JS },
        encoding: 'utf8', timeout: 120000, maxBuffer: 16 * 1024 * 1024 });
      writeFileSync(join(directory, 'logs', `${sequence++}-${label}.log`), (result.stdout ?? '') + (result.stderr ?? ''));
      assert.ifError(result.error);
      assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`);
      return result;
    };
    const cache = profile => join(output, '.gopurs-cache', ...(profile === 'dev' ? ['profiles', 'dev'] : []));
    const proof = profile => {
      const envelope = JSON.parse(readFileSync(join(cache(profile), 'preparation-literals-v1/proof.json')));
      assert.equal(wireHash(envelope.payload), envelope.hash);
      return JSON.parse(envelope.payload);
    };
    const build = (label, profile, force = false) => {
      const result = run(label, process.env.GOPURS_TEST_COMPILER ?? join(root, 'bin/gopurs'),
        [`--${profile}`, '--main', 'Main', ...(force ? ['--force'] : [])]);
      const pointer = JSON.parse(readFileSync(join(cache(profile), 'current.json')));
      const bytes = readFileSync(join(cache(profile), 'objects', pointer.manifest.sha256));
      assert.equal(hash(bytes), pointer.manifest.sha256);
      const manifest = JSON.parse(bytes);
      compilerIdentity ??= manifest.snapshot.compiler;
      assert.equal(manifest.snapshot.compiler, compilerIdentity, 'compiler/helpers changed while the gate was running');
      if (force) {
        const count = readdirSync(output, { withFileTypes: true }).filter(entry => entry.isDirectory()
          && !entry.name.startsWith('.') && entry.name !== 'main' && entry.name !== 'packages'
          && readdirSync(join(output, entry.name)).includes('corefn.json')).length;
        assert.match(result.stderr, new RegExp(`module cache: optimized=${count}, reused=0, emitted=${count}, restored=0`));
      }
      return result;
    };
    const execute = (label, seed, extra) => {
      const sum = seed * 2 + 1;
      assert.equal(run(label, 'go', ['run', './main'], output).stdout, [sum, 2, sum, 10, sum, ...(extra ? [17] : [])].join('\n') + '\n');
    };
    const helperKeys = state => Object.fromEntries(state.index.filter(entry => helpers.includes(entry.name))
      .map(entry => [entry.name, entry.keys.toSorted()]));

    run('frontend', 'spago', ['build', '-q', '--offline']);
    build('dev-initial', 'dev');
    execute('dev-initial-execute', 5, false);
    console.log(`${mode.name}: first dev Go compile/run passed (ADT Int/Boolean, record payload, class dictionaries)`);
    assert.deepEqual(proof('dev').index, [], 'dev must not contain direct or transitive specializations');
    build('release-initial', 'release');
    execute('release-initial-execute', 5, false);
    const releaseKeys = helperKeys(proof('release'));
    for (const name of helpers) {
      assert.ok(releaseKeys[name]?.length > 0, `${name}: release must retain specialization`);
    }

    for (const [label, seed, extra] of [['literal', 6, false], ['added-type', 6, true], ['removed-type', 6, false]]) {
      writeFileSync(join(source, 'Main.purs'), mainSource(seed, extra));
      run(`${label}-frontend`, 'spago', ['build', '-q', '--offline']);
      build(`${label}-dev`, 'dev');
      execute(`${label}-execute`, seed, extra);
      assert.deepEqual(proof('dev').index, [], `${label}: a source edit must not reintroduce dev specializations`);
      const incremental = generated(output);
      build(`${label}-full`, 'dev', true);
      assert.deepEqual(proof('dev').index, [], `${label}: forced preparation must also skip specialization`);
      assert.deepEqual(generated(output), incremental, `${label}: incremental must match a fully generated build`);
      assert.match(build(`${label}-noop`, 'dev').stderr, /cache hit:/);
    }
    completed++;
  });
});
