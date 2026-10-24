import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, mkdirSync, rmSync, writeFileSync, readFileSync, copyFileSync, statSync, utimesSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import fs from 'node:fs';
import { syncBuiltinESMExports } from 'node:module';
import test from 'node:test';
import { Left } from '../output/Data.Either/index.js';
import { Cons, Nil } from '../output/Data.List.Types/index.js';
import * as Map from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import * as Set from '../output/Data.Set/index.js';
import * as Aff from '../output/Effect.Aff/index.js';
import * as LookupTable from '../output/Gopurs.Cache.LookupTable/index.js';
import * as Capsule from '../output/Gopurs.Driver.Prepare.ProjectionCapsule/index.js';
import * as Lazy from '../output/Data.Lazy/index.js';
import { Right } from '../output/Data.Either/index.js';
import * as F from '../output/Data.Foldable/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import * as Prepare from '../output/Gopurs.Driver.Prepare/index.js';
import * as Persistent from '../output/Gopurs.Driver.Prepare.Persistent/index.js';
import * as Projection from '../output/Gopurs.Driver.Prepare.Projection/index.js';
import * as Wire from '../output/Gopurs.Cache.Wire/index.js';
import * as Manifest from '../output/Gopurs.Driver.Manifest/index.js';
import { findTypedCompiler } from './native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const compiler = findTypedCompiler(root, process.env.GOPURS_PURS);
const run = action => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(action)());
const array = list => { const values = []; for (; list instanceof Cons; list = list.value1) values.push(list.value0); return values; };
const wire = mod => Wire.write(Wire.codecModule(Wire.codecAnn))(mod)();
const metadataEntries = metadata => Object.fromEntries(Object.entries(metadata).map(([name, values]) =>
  [name, Map.toUnfoldable(unfoldableArray)(name === 'globalTypes' ? LookupTable.completeOrFail(values) : values).map(pair => [pair.value0, pair.value1])]));
function setup(t, sources) {
  const workspace = mkdtempSync(join(tmpdir(), 'gopurs-compact-proof-'));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  const src = join(workspace, 'src'), output = join(workspace, 'output'), stage = join(workspace, 'stage');
  mkdirSync(src); mkdirSync(join(stage, 'prior'), { recursive: true });
  const stamp = () => {
    const digest = Wire.hash(Object.keys(sources).sort().map(name => readFileSync(join(output, name, 'corefn.json'), 'utf8')).join('\n'));
    writeFileSync(join(stage, 'session.json'), JSON.stringify({ snapshot: { compiler: 'compiler-test-v1', inputs: [{ name: 'tast', kind: 'tast', digest }] } }));
    writeFileSync(join(stage, 'source-hashes.json'), JSON.stringify(Object.keys(sources).sort().map(name => ({ name, hash: Wire.hash(readFileSync(join(output, name, 'corefn.json'), 'utf8')) }))));
  };
  const compile = changes => {
    Object.assign(sources, changes);
    for (const [name, text] of Object.entries(sources)) writeFileSync(join(src, name + '.purs'), text);
    const result = spawnSync(compiler, ['compile', '--codegen', 'corefn', '--output', output, ...Object.keys(sources).map(name => join(src, name + '.purs'))], { cwd: workspace, encoding: 'utf8', timeout: 30000 });
    assert.ifError(result.error); assert.equal(result.status, 0, result.stdout + result.stderr);
    mkdirSync(join(output, '.gopurs-cache'), { recursive: true });
    stamp();
  };
  const session = { root: output, input: output, stage, context: 'context-test-v1', count: Object.keys(sources).length };
  const full = () => run(Prepare.prepareModulesWithCache(output)(session.count)(1)(Nothing.value)(Map.empty));
  const candidate = (jobs = 1) => run(Persistent.preparePersistent(session)(jobs)(Nothing.value)(Map.empty));
  const publish = async (cache, jobs = 1) => {
    await run(Persistent.publishPersistent(session)(jobs)(cache));
    copyFileSync(join(stage, 'session.json'), join(stage, 'prior', 'manifest.json'));
  };
  return { sources, workspace, output, stage, session, compile, stamp, full, candidate, publish };
}

const poly = 'module Poly (ignore) where\nignore :: forall a. a -> Int\nignore _ = 42\n';
const leaf = (text, arg = '1') => `module Leaf (value, number) where\nimport Poly (ignore)\nvalue :: String\nvalue = "${text}"\nnumber :: Int\nnumber = ignore ${arg}\n`;
const main = 'module Main (answer) where\nimport Leaf (value)\nanswer :: String\nanswer = value\n';

async function oracle(candidate, expected) {
  const expectedModules = array(expected.prepared.modules);
  assert(candidate.materialize instanceof Just);
  assert.equal(candidate.prepared.metadataUnchanged, true);
  assert.equal(expected.prepared.metadataUnchanged, false);
  assert(candidate.compactSourceDelta instanceof Just, 'compact delta is backed by the accepted committed proof');
  assert(expected.compactSourceDelta instanceof Nothing, 'full preparation carries no shortcut delta');
  const changed = Map.toUnfoldable(unfoldableArray)(candidate.sourceHashes).map(pair => pair.value0)
    .filter(name => !Set.toUnfoldable(unfoldableArray)(candidate.reusable).includes(name));
  assert.deepEqual(changed, [candidate.compactSourceDelta.value0], 'the delta identifies the sole changed source');
  assert.deepEqual(metadataEntries(candidate.prepared.metadata), metadataEntries(expected.prepared.metadata), 'literal proof preserves every source metadata table');
  const known = Manifest.buildManifestWithInterfaces(candidate.sourceInterfaces)(Nothing.value)('')('ctx')(Map.empty)(array(candidate.sourceModules))(candidate.sourceHashes)();
  const complete = Manifest.buildManifest('')('ctx')(Map.empty)(array(expected.sourceModules))(expected.sourceHashes)();
  assert.deepEqual(known, complete, 'persisted interfaces remain exact without the optional manifest sidecar');
  for (const mod of expectedModules) {
    const result = await run(candidate.materialize.value0(mod.name));
    assert(result instanceof Just, 'payload must materialize: ' + mod.name);
    assert.equal(wire(result.value0), wire(mod), 'compact materialization equals full: ' + mod.name);
  }
}

test('first publication has identical proof, payloads and fragment order with one or eight workers', async t => {
  const sources = { Poly: poly, Leaf: leaf('before'), Main: main };
  for (let i = 0; i < 6; i++) sources[`Extra${i}`] = `module Extra${i} (label) where\nlabel :: String\nlabel = "${i}"\n`;
  const f = setup(t, sources);
  f.compile({});
  const initial = await f.full();
  assert(initial.proof instanceof Nothing, 'the publication must construct a new proof');
  const directory = join(f.output, '.gopurs-cache', 'preparation-literals-v1');
  const snapshot = () => {
    const bytes = readFileSync(join(directory, 'proof.json'));
    const envelope = JSON.parse(bytes.toString('utf8'));
    assert.equal(Wire.hash(envelope.payload), envelope.hash, 'the complete proof envelope is valid');
    const proof = JSON.parse(envelope.payload);
    assert.equal(proof.schema, 5);
    assert.deepEqual(proof.modules.map(mod => mod.name), array(initial.sourceModules).map(mod => mod.name),
      'module proof order follows the original source order');
    assert(proof.payloads.length > 0, 'the fixture exercises specialization payloads');
    assert.equal(proof.projection.length, 8, 'the fixture exercises every projection fragment');
    assert.equal(new globalThis.Set(proof.projection).size, 8, 'projection references are distinct');
    const hashes = [...new globalThis.Set([...proof.payloads.map(entry => entry.hash), ...proof.projection,
      ...(proof.capsule ?? [])])].sort();
    const objects = hashes.map(hash => {
      const bytes = readFileSync(join(directory, hash + '.json'));
      assert.equal(Wire.hash(bytes.toString('utf8')), hash, 'every referenced object is content-addressed');
      return [hash, bytes];
    });
    assert.deepEqual(fs.readdirSync(directory).sort(), ['proof.json', ...hashes.map(hash => hash + '.json')].sort(),
      'first publication contains exactly the referenced objects and proof');
    return { bytes, objects, projection: proof.projection };
  };
  assert.equal(fs.existsSync(directory), false);
  await f.publish(initial, 1);
  const sequential = snapshot();
  // Remove this fixture's entire private proof directory. Reusing a valid
  // previous proof would skip encoding and would not test first publication.
  rmSync(directory, { recursive: true });
  assert.equal(fs.existsSync(directory), false);
  assert(initial.proof instanceof Nothing);
  await f.publish(initial, 8);
  assert.deepEqual(snapshot(), sequential, 'worker count preserves every byte and reference order');
  f.compile({ Leaf: leaf('after') });
  await oracle(await f.candidate(8), await f.full());
});

test('literal edit reuses anchored compact preparation and materializes every reader exactly', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); const initial = await f.full(); await f.publish(initial);
  f.compile({ Leaf: leaf('a longer label') });
  const candidate = await f.candidate();
  assert.deepEqual(Set.toUnfoldable(unfoldableArray)(candidate.reusable), ['Main', 'Poly']);
  await oracle(candidate, await f.full());
  await f.publish(candidate);
  const resident = await run(Prepare.prepareModulesDelta(f.output)(f.session.count)(1)(Nothing.value)(Map.empty)(candidate));
  assert(resident.compactSourceDelta instanceof Nothing, 'a resident no-change run cannot reuse the previous run delta');
  f.compile({ Leaf: leaf('third') });
  await oracle(await f.candidate(), await f.full());
});

test('literal metadata proof preserves ADT layouts, classes and foreign types', async t => {
  const layout = text => `module Layout (Token(..), Chain(..), Wrap(..), class Select, select, foreignValue, label) where
data Token = LeftToken | RightToken
data Chain = End | Next Int Chain
newtype Wrap = Wrap Int
class Select a where
  select :: a -> Int
instance selectInt :: Select Int where
  select x = x
foreign import foreignValue :: Int
label :: String
label = "${text}"
`;
  const f = setup(t, { Poly: poly, Leaf: leaf('unchanged'), Layout: layout('before') });
  f.compile({});
  const initial = await f.full();
  assert.equal(initial.prepared.metadataUnchanged, false);
  await f.publish(initial);
  f.compile({ Layout: layout('after') });
  const candidate = await f.candidate();
  assert.equal(candidate.prepared.metadataUnchanged, true);
  assert.deepEqual(metadataEntries(candidate.prepared.metadata), metadataEntries(initial.prepared.metadata));
  await oracle(candidate, await f.full());
});

test('frozen source hashes require a complete inventory with distinct module names', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); await f.publish(await f.full());
  f.compile({ Leaf: leaf('after') });
  const entries = Object.keys(f.sources).sort().map(name => ({ name, hash: Wire.hash(readFileSync(join(f.output, name, 'corefn.json'), 'utf8')) }));
  const expected = await f.full();
  const path = join(f.stage, 'source-hashes.json');
  for (const inventory of [
    entries,
    entries.slice(1),
    entries.map((entry, i) => i === 0 ? { ...entry, name: 'Unknown' } : entry),
    entries.map((entry, i) => i === 0 ? { ...entry, name: entries[1].name } : entry),
  ]) {
    writeFileSync(path, JSON.stringify(inventory));
    const candidate = await f.candidate();
    assert.deepEqual(Map.toUnfoldable(unfoldableArray)(candidate.sourceHashes).map(pair => ({ name: pair.value0, hash: pair.value1 })), entries);
    if (inventory === entries) await oracle(candidate, expected);
    else {
      assert(candidate.materialize instanceof Nothing, 'invalid freeze inventory falls back to full loading');
      assert.deepEqual(array(candidate.prepared.modules).map(wire), array(expected.prepared.modules).map(wire));
    }
  }
});

test('an empty specialization index preserves untyped private bindings and source exports', async t => {
  const secret = text => `module Secret (public) where\nprivateValue :: String\nprivateValue = "${text}"\npublic :: String\npublic = privateValue\n`;
  const entry = text => `module Main (answer, label) where\nimport Secret (public)\nanswer :: String\nanswer = public\nlabel :: String\nlabel = "${text}"\n`;
  const f = setup(t, { Secret: secret('before'), Main: entry('first') });
  const compileUntypedPrivate = (text, changes = {}) => {
    f.compile({ Secret: secret(text), ...changes });
    const path = join(f.output, 'Secret', 'corefn.json');
    const json = JSON.parse(readFileSync(path, 'utf8'));
    for (const name of ['privateValue', 'public']) {
      const binding = json.decls.find(binding => binding.identifier === name);
      assert(binding, `fixture retains ${name}`);
      binding.annotation.type = null;
      binding.expression.annotation.type = null;
    }
    writeFileSync(path, JSON.stringify(json));
    f.stamp();
  };
  compileUntypedPrivate('before');
  const initial = await f.full();
  assert.equal(Map.size(initial.mono.instantiations), 0);
  const mod = array(initial.prepared.modules).find(mod => mod.name === 'Secret');
  assert.deepEqual(mod.exports, ['public']);
  const binding = mod.decls.find(decl => decl.value0.value1 === 'privateValue').value0;
  assert(binding.value0.type instanceof Nothing);
  assert(binding.value2.value0.type instanceof Nothing);
  await f.publish(initial);
  compileUntypedPrivate('after');
  const candidate = await f.candidate();
  assert.deepEqual(Set.toUnfoldable(unfoldableArray)(candidate.reusable), ['Main']);
  await oracle(candidate, await f.full());
  await f.publish(candidate);
  compileUntypedPrivate('after', { Main: entry('second') });
  const next = await f.candidate();
  assert.deepEqual(Set.toUnfoldable(unfoldableArray)(next.reusable), ['Secret']);
  const header = array(next.prepared.modules).find(mod => mod.name === 'Secret');
  assert.deepEqual(header.exports, [], 'inert parallel headers have already filtered exported bindings');
  assert.deepEqual(header.decls.map(bind => bind.value0.value1), ['privateValue']);
  assert(header.decls[0].value0.value0.type instanceof Nothing);
  assert(header.decls[0].value0.value2.value0.type instanceof Nothing);
  await oracle(next, await f.full());
});

test('a type-instantiation edit falls back, including same-interface Generic.ignore Int to String', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); await f.publish(await f.full());
  f.compile({ Leaf: leaf('before', '"hello"') });
  const candidate = await f.candidate();
  assert.equal(Set.size(candidate.reusable), 0);
  assert(candidate.materialize instanceof Nothing);
  assert.deepEqual(array(candidate.prepared.modules).map(wire), array((await f.full()).prepared.modules).map(wire));
});

test('a dictionary body read blocks literal-only reuse', async t => {
  const dict = n => `module Dict (dict) where\ndict :: { method :: Int -> Int }\ndict = { method: \\_ -> ${n} }\n`;
  const apply = 'module Apply (apply) where\napply :: forall a. { method :: a -> Int } -> a -> Int\napply d x = d.method x\n';
  const caller = 'module Caller (answer) where\nimport Dict (dict)\nimport Apply (apply)\nanswer :: Int\nanswer = apply dict 1\n';
  const f = setup(t, { Dict: dict(1), Apply: apply, Caller: caller });
  f.compile({}); const initial = await f.full();
  assert(Set.toUnfoldable(unfoldableArray)(initial.mono.bodyReads).includes('Dict.dict'));
  await f.publish(initial); f.compile({ Dict: dict(2) });
  const candidate = await f.candidate();
  assert.equal(Set.size(candidate.reusable), 0);
  assert(candidate.materialize instanceof Nothing);
});

test('a dictionary alias records its underlying body and rejects a changed scalar', async t => {
  const base = n => `module Base (dict) where\ndict :: { method :: Int -> Int }\ndict = { method: \\_ -> ${n} }\n`;
  const alias = 'module Alias (dict) where\nimport Base as Base\ndict :: { method :: Int -> Int }\ndict = Base.dict\n';
  const apply = 'module Apply (apply) where\napply :: forall a. { method :: a -> Int } -> a -> Int\napply d x = d.method x\n';
  const caller = 'module Caller (answer) where\nimport Alias (dict)\nimport Apply (apply)\nanswer :: Int\nanswer = apply dict 1\n';
  const f = setup(t, { Base: base(1), Alias: alias, Apply: apply, Caller: caller });
  f.compile({});
  const initial = await f.full();
  const reads = Set.toUnfoldable(unfoldableArray)(initial.mono.bodyReads);
  assert(reads.includes('Alias.dict'), 'dictionary resolution reads the alias');
  assert(reads.includes('Base.dict'), 'dictionary resolution also reads its underlying body');
  await f.publish(initial);
  f.compile({ Base: base(2) });
  const candidate = await f.candidate();
  assert.equal(Set.size(candidate.reusable), 0);
  assert(candidate.materialize instanceof Nothing);
  const expected = await f.full();
  assert.deepEqual(array(candidate.prepared.modules).map(wire), array(expected.prepared.modules).map(wire));
  const provider = cache => wire(array(cache.prepared.modules).find(mod => mod.name === 'Apply'));
  assert.notEqual(provider(candidate), provider(initial), 'the underlying dictionary scalar changes another module’s prepared specialization');
});

test('proof from a different committed generation is rejected', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); await f.publish(await f.full());
  writeFileSync(join(f.stage, 'prior', 'manifest.json'), JSON.stringify({ snapshot: { compiler: 'compiler-test-v1', inputs: [{ name: 'tast', kind: 'tast', digest: 'wrong-generation' }] } }));
  f.compile({ Leaf: leaf('after') });
  assert.equal(Set.size((await f.candidate()).reusable), 0);
});

test('missing or corrupted own payload fails materialization without creating a placeholder hit', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); await f.publish(await f.full());
  f.compile({ Leaf: leaf('after') }); const candidate = await f.candidate();
  assert(candidate.materialize instanceof Just);
  const envelope = JSON.parse(readFileSync(join(f.output, '.gopurs-cache', 'preparation-literals-v1', 'proof.json'), 'utf8'));
  const proof = JSON.parse(envelope.payload), own = proof.payloads.find(value => value.name === 'Poly');
  assert(own);
  writeFileSync(join(f.output, '.gopurs-cache', 'preparation-literals-v1', own.hash + '.json'), '{}');
  assert((await run(candidate.materialize.value0('Poly'))) instanceof Nothing);
});

test('literal length changes preserve static-argument source locations in prepared providers', async t => {
  const apply = 'module Apply (apply) where\napply :: forall a. String -> (Int -> a) -> a\napply _ f = f 1\n';
  const helper = 'module Helper (fun) where\nfun :: Int -> Int\nfun _ = 42\n';
  const caller = text => `module Caller (answer) where\nimport Apply (apply)\nimport Helper (fun)\nanswer :: Int\nanswer = apply "${text}" fun\n`;
  const f = setup(t, { Apply: apply, Helper: helper, Caller: caller('x') });
  f.compile({}); await f.publish(await f.full());
  f.compile({ Caller: caller('a much longer literal') });
  const candidate = await f.candidate(), expected = await f.full();
  if (candidate.materialize instanceof Just) await oracle(candidate, expected);
  else assert.deepEqual(array(candidate.prepared.modules).map(wire), array(expected.prepared.modules).map(wire));
});

test('a corrupted proof envelope falls back to complete preparation', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); await f.publish(await f.full());
  const path = join(f.output, '.gopurs-cache', 'preparation-literals-v1', 'proof.json');
  const envelope = JSON.parse(readFileSync(path, 'utf8'));
  envelope.payload = envelope.payload.replace('context-test-v1', 'context-test-v2');
  writeFileSync(path, JSON.stringify(envelope));
  f.compile({ Leaf: leaf('after') });
  assert.equal(Set.size((await f.candidate()).reusable), 0);
});


test('an unchanged full build reuses verified proof payloads and repairs corruption', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); const initial = await f.full(); await f.publish(initial);
  const proofPath = join(f.output, '.gopurs-cache', 'preparation-literals-v1', 'proof.json');
  const proof = JSON.parse(JSON.parse(readFileSync(proofPath, 'utf8')).payload);
  const own = proof.payloads.find(entry => entry.name === 'Poly'); assert(own);
  const payload = join(f.output, '.gopurs-cache', 'preparation-literals-v1', own.hash + '.json');
  const expected = readFileSync(payload, 'utf8');
  utimesSync(payload, new Date(0), new Date(0));
  await f.publish(initial);
  assert.equal(statSync(payload).mtimeMs, 0, 'valid unchanged payload is not re-encoded or rewritten');
  writeFileSync(payload, '{}');
  await f.publish(initial);
  assert.equal(readFileSync(payload, 'utf8'), expected, 'corruption is repaired even when the full TAST is unchanged');
});

async function recordTastReads(action) {
  const original = fs.readFile, reads = [];
  fs.readFile = function (path, ...args) {
    if (String(path).endsWith('/corefn.json')) reads.push(String(path).split('/').at(-2));
    return original.call(this, path, ...args);
  };
  syncBuiltinESMExports();
  try { return { value: await action(), reads }; }
  finally { fs.readFile = original; syncBuiltinESMExports(); }
}

test('compact loading reads only the changed TAST and materializes other sources on demand', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); await f.publish(await f.full());
  f.compile({ Leaf: leaf('after') });
  const { value: candidate, reads } = await recordTastReads(f.candidate);
  assert(candidate.materialize instanceof Just);
  assert.deepEqual(reads, ['Leaf']);
  const second = await recordTastReads(() => run(candidate.materialize.value0('Main')));
  assert(second.value instanceof Just);
  assert.deepEqual(second.reads, ['Main']);
  const repeated = await recordTastReads(() => run(candidate.materialize.value0('Main')));
  assert(repeated.value instanceof Just);
  assert.deepEqual(repeated.reads, [], 'already loaded TAST is shared across materializations');
  await oracle(candidate, await f.full());
});

test('dictionary aliases are resolved by demand replay with exact prepared output', async t => {
  const base = 'module Base (dict) where\ndict :: { method :: Int -> Int }\ndict = { method: \\_ -> 42 }\n';
  const alias = 'module Alias (dict) where\nimport Base as Base\ndict :: { method :: Int -> Int }\ndict = Base.dict\n';
  const apply = 'module Apply (apply) where\napply :: forall a. { method :: a -> Int } -> a -> Int\napply d x = d.method x\n';
  const caller = 'module Caller (answer) where\nimport Alias (dict)\nimport Apply (apply)\nanswer :: Int\nanswer = apply dict 1\n';
  const label = text => `module Label (label) where\nlabel :: String\nlabel = "${text}"\n`;
  const f = setup(t, { Base: base, Alias: alias, Apply: apply, Caller: caller, Label: label('before') });
  f.compile({}); await f.publish(await f.full());
  f.compile({ Label: label('after') });
  const { value: candidate, reads } = await recordTastReads(f.candidate);
  assert.deepEqual(reads, ['Label']);
  assert(candidate.materialize instanceof Just);
  const expected = await f.full();
  const fresh = await recordTastReads(() => run(candidate.materialize.value0('Apply')));
  assert(fresh.value instanceof Just);
  assert.equal(wire(fresh.value.value0), wire(array(expected.prepared.modules).find(mod => mod.name === 'Apply')));
  assert(fresh.reads.includes('Alias'), 'missed alias lookup loads its owner');
  assert(fresh.reads.includes('Base'), 'replay discovers and loads the underlying dictionary');
  await oracle(candidate, expected);
  const path = join(f.output, '.gopurs-cache', 'preparation-literals-v1', 'proof.json');
  const proof = JSON.parse(JSON.parse(readFileSync(path, 'utf8')).payload);
  proof.bodyReads = proof.bodyReads.filter(name => name !== 'Alias.dict');
  const payload = JSON.stringify(proof);
  writeFileSync(path, JSON.stringify({ hash: Wire.hash(payload), payload }));
  const incomplete = await f.candidate();
  assert(incomplete.materialize instanceof Just);
  assert((await run(incomplete.materialize.value0('Apply'))) instanceof Nothing, 'a body lookup missing from the committed proof cannot be assumed absent');
});

test('a corrupt projection falls back and a full publication repairs it', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); const initial = await f.full(); await f.publish(initial);
  const proofPath = join(f.output, '.gopurs-cache', 'preparation-literals-v1', 'proof.json');
  const proof = JSON.parse(JSON.parse(readFileSync(proofPath, 'utf8')).payload);
  const projection = join(f.output, '.gopurs-cache', 'preparation-literals-v1', proof.projection[0] + '.json');
  const exact = readFileSync(projection, 'utf8');
  writeFileSync(projection, '{}');
  f.compile({ Leaf: leaf('after') });
  const fallback = await f.candidate();
  assert(fallback.materialize instanceof Nothing);
  f.compile({ Leaf: leaf('before') });
  await f.publish(initial);
  assert.equal(readFileSync(projection, 'utf8'), exact, 'the full path repairs the projection even with identical TAST');
  f.compile({ Leaf: leaf('final') });
  await oracle(await f.candidate(), await f.full());
});

test('projection fragments reject missing, repeated, reordered, and incompatible manifests', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); const initial = await f.full(); await f.publish(initial);
  const directory = join(f.output, '.gopurs-cache', 'preparation-literals-v1');
  const path = join(directory, 'proof.json');
  const original = readFileSync(path, 'utf8');
  const proof = JSON.parse(JSON.parse(original).payload);
  assert.equal(proof.schema, 5);
  assert.equal(proof.projection.length, 3);
  const first = join(directory, proof.projection[0] + '.json');
  const fragment = readFileSync(first, 'utf8');
  const save = value => { const payload = JSON.stringify(value); writeFileSync(path, JSON.stringify({ hash: Wire.hash(payload), payload })); };
  f.compile({ Leaf: leaf('after') });
  rmSync(first);
  assert((await f.candidate()).materialize instanceof Nothing, 'a missing private fragment forces full preparation');
  writeFileSync(first, fragment);
  for (const digests of [proof.projection.slice(1), [proof.projection[0], ...proof.projection.slice(0, -1)], proof.projection.toReversed()]) {
    save({ ...proof, projection: digests });
    assert((await f.candidate()).materialize instanceof Nothing, 'the manifest must describe every distinct fragment in original order');
  }
  save({ ...proof, schema: 3 });
  assert((await f.candidate()).materialize instanceof Nothing, 'older proof layouts take the full path');
  writeFileSync(path, original);
  await oracle(await f.candidate(), await f.full());
});

test('projection splitting preserves order and every metadata entry, rejecting overlapping tables', async t => {
  const sources = { Poly: poly, Leaf: leaf('before'), Main: main };
  for (let i = 0; i < 6; i++) sources[`Extra${i}`] = `module Extra${i} (label) where\nlabel :: String\nlabel = "${i}"\n`;
  const f = setup(t, sources); f.compile({});
  const initial = await f.full();
  const value = { modules: initial.prepared.modules, metadata: initial.prepared.metadata, interfaces: initial.sourceInterfaces, mainModules: ['Main', 'Extra0'] };
  const fragments = Projection.split(value);
  assert.equal(fragments.length, 8);
  const merged = Projection.merge(fragments);
  assert(merged instanceof Just);
  assert.deepEqual(array(merged.value0.modules).map(mod => mod.name), array(value.modules).map(mod => mod.name));
  assert.deepEqual(metadataEntries(merged.value0.metadata), metadataEntries(value.metadata));
  assert.deepEqual(merged.value0.mainModules, value.mainModules);
  for (const count of [0, 1, 7, 8, 9]) {
    const modules = array(value.modules).slice(0, count).reduceRight((tail, head) => new Cons(head, tail), Nil.value);
    const parts = Projection.split({ ...value, modules });
    assert.equal(parts.length, Math.max(1, Math.min(8, count)));
    const restored = Projection.merge(parts);
    assert(restored instanceof Just);
    assert.deepEqual(array(restored.value0.modules), array(modules));
    assert.deepEqual(metadataEntries(restored.value0.metadata), metadataEntries(value.metadata));
  }
  for (const field of Object.keys(value.metadata)) {
    const donor = fragments.findIndex(fragment => Map.size(field === 'globalTypes' ? LookupTable.completeOrFail(fragment.value.metadata[field]) : fragment.value.metadata[field]) > 0);
    if (donor < 0) continue;
    const recipient = (donor + 1) % fragments.length;
    const overlap = fragments.map((fragment, index) => index !== recipient ? fragment : { ...fragment, value: { ...fragment.value, metadata: { ...fragment.value.metadata, [field]: fragments[donor].value.metadata[field] } } });
    assert(Projection.merge(overlap) instanceof Nothing, `duplicate ${field} keys cannot disappear during union`);
  }
  const setOverlap = fragments.map((fragment, index) => index > 1 ? fragment : { ...fragment, value: { ...fragment.value, metadata: { ...fragment.value.metadata, enumAdts: Set.singleton('SharedEnum') } } });
  assert(Projection.merge(setOverlap) instanceof Nothing, 'overlapping set entries are rejected');
  const interfaceOverlap = fragments.map((fragment, index) => index > 1 ? fragment : { ...fragment, value: { ...fragment.value, interfaces: Map.singleton('Shared')('interface-hash') } });
  assert(Projection.merge(interfaceOverlap) instanceof Nothing, 'overlapping source interfaces are rejected');
  await f.publish(initial);
  f.compile({ Leaf: leaf('after') });
  const expected = await f.full();
  for (const jobs of [1, 8]) {
    const candidate = await f.candidate(jobs);
    assert(candidate.materialize instanceof Just);
    await oracle(candidate, expected);
  }
});

test('demand replay preserves proved missing bindings and absent owners', async t => {
  for (const owner of ['Base', 'Absent']) {
    const label = text => `module Label (label) where\nlabel :: String\nlabel = "${text}"\n`;
    const f = setup(t, {
      Base: 'module Base (dict) where\ndict :: { method :: Int -> Int }\ndict = { method: \\_ -> 42 }\n',
      Alias: 'module Alias (dict) where\nimport Base as Base\ndict :: { method :: Int -> Int }\ndict = Base.dict\n',
      Apply: 'module Apply (apply) where\napply :: forall a. { method :: a -> Int } -> a -> Int\napply d x = d.method x\n',
      Caller: 'module Caller (answer) where\nimport Alias (dict)\nimport Apply (apply)\nanswer :: Int\nanswer = apply dict 1\n',
      Label: label('before'),
    });
    const compileMissing = text => {
      f.compile({ Label: label(text) });
      const path = join(f.output, 'Alias', 'corefn.json');
      const source = JSON.parse(readFileSync(path, 'utf8'));
      const target = source.decls.find(binding => binding.identifier === 'dict').expression.value;
      target.moduleName = [owner]; target.identifier = 'missing';
      writeFileSync(path, JSON.stringify(source));
      f.stamp();
    };
    compileMissing('before');
    const initial = await f.full();
    assert(Set.toUnfoldable(unfoldableArray)(initial.mono.bodyReads).includes(`${owner}.missing`));
    await f.publish(initial);
    compileMissing('after');
    const candidate = await f.candidate();
    assert(candidate.materialize instanceof Just);
    const fresh = await recordTastReads(() => run(candidate.materialize.value0('Apply')));
    assert(fresh.value instanceof Just);
    assert.equal(fresh.reads.filter(name => name === owner).length, owner === 'Base' ? 1 : 0);
    await oracle(candidate, await f.full());
  }
});

const keysOf = keys => Set.fromFoldable(F.foldableArray)(ordString)(keys);
const readState = f => JSON.parse(JSON.parse(readFileSync(join(f.output, '.gopurs-cache/preparation-literals-v1/proof.json'), 'utf8')).payload);
const globalWire = table => Wire.write(LookupTable.codecLookupTable(Wire.codecExprType))(table)();

test('projection capsules are proof-anchored, retain exact metadata and survive an intervening full oracle', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('first'), Main: main });
  f.compile({}); await f.publish(await f.full());
  assert.equal(readState(f).capsule, null, 'full preparation does not precompute a capsule');
  f.compile({ Leaf: leaf('second') });
  const first = await f.candidate(8);
  assert(first.metadataCoverage instanceof Nothing);
  first.prepared.metadataDemands.value0(keysOf(['Leaf.value', 'Known.absent']))();
  await f.publish(first);
  const primed = readState(f);
  assert.equal(primed.schema, 5);
  assert.equal(primed.capsule.length, primed.projection.length);
  assert.equal(new globalThis.Set(primed.capsule).size, primed.capsule.length);
  const complete = await f.full(); await f.publish(complete);
  assert.deepEqual(readState(f).capsule, primed.capsule, 'full carry preserves already validated capsule references');
  f.compile({ Leaf: leaf('third') });
  const steady = await f.candidate(8), expected = await f.full();
  assert(steady.metadataCoverage instanceof Just);
  assert.deepEqual(Set.toUnfoldable(unfoldableArray)(steady.metadataCoverage.value0), ['Known.absent', 'Leaf.value']);
  assert.deepEqual(LookupTable.lookup('Known.absent')(steady.prepared.metadata.globalTypes), Nothing.value);
  assert.deepEqual(LookupTable.lookup('Leaf.value')(steady.prepared.metadata.globalTypes), LookupTable.lookup('Leaf.value')(expected.prepared.metadata.globalTypes));
  for (const key of ['Poly.ignore', 'New.absent']) {
    assert.deepEqual(LookupTable.lookup(key)(steady.prepared.metadata.globalTypes), LookupTable.lookup(key)(expected.prepared.metadata.globalTypes), 'uncovered lookup uses the complete table: ' + key);
  }
  assert.equal(globalWire(steady.prepared.metadata.globalTypes), globalWire(expected.prepared.metadata.globalTypes), 'opaque encoding forces exact complete types');
  await oracle(steady, expected);
  steady.prepared.metadataDemands.value0(keysOf(['Poly.ignore', 'New.absent']))();
  await f.publish(steady);
  assert.notDeepEqual(readState(f).capsule, primed.capsule, 'new demands expand the persisted capsule');
  f.compile({ Leaf: leaf('fourth') });
  const expanded = await f.candidate();
  assert.deepEqual(Set.toUnfoldable(unfoldableArray)(expanded.metadataCoverage.value0), ['Known.absent', 'Leaf.value', 'New.absent', 'Poly.ignore']);
});

test('a capsule remains lazy for covered answers and preserves disjoint fragments and all eleven full tables', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('first'), Main: main });
  f.compile({}); const full = await f.full();
  const projection = { modules: full.prepared.modules, metadata: full.prepared.metadata, interfaces: Map.empty, mainModules: ['Main'] };
  const coverage = keysOf(['Leaf.value', 'Known.absent']);
  const fragments = Capsule.create(coverage)(projection);
  let forced = 0;
  const delayed = Lazy.defer(_ => { forced++; return new Right(LookupTable.completeOrFail(full.prepared.metadata.globalTypes)); });
  const opened = Capsule.open(fragments)(delayed);
  assert(opened instanceof Just);
  const table = opened.value0.projection.metadata.globalTypes;
  assert(LookupTable.lookup('Leaf.value')(table) instanceof Just);
  assert(LookupTable.lookup('Known.absent')(table) instanceof Nothing);
  assert.equal(forced, 0, 'covered presence and absence do not decode complete metadata');
  assert(LookupTable.lookup('Poly.ignore')(table) instanceof Just);
  assert(LookupTable.lookup('Uncovered.absent')(table) instanceof Nothing);
  assert.equal(forced, 1, 'all uncovered requests share the complete decoder');
  assert.deepEqual(metadataEntries(opened.value0.projection.metadata), metadataEntries(projection.metadata));
  assert.deepEqual(array(opened.value0.projection.modules), array(projection.modules));
  assert.deepEqual(opened.value0.projection.mainModules, ['Main']);
  for (const mutation of [fragments.slice(1), fragments.toReversed(), fragments.map((fragment, i) => i ? fragment : { ...fragment, schema: 2 })]) {
    assert(Capsule.open(mutation)(delayed) instanceof Nothing);
  }
  const overlap = fragments.map((fragment, i) => i > 1 ? fragment : { ...fragment, coverage: keysOf(['same']) });
  assert(Capsule.open(overlap)(delayed) instanceof Nothing, 'duplicate coverage cannot disappear during union');
  const outside = fragments.map((fragment, i) => i ? fragment : { ...fragment, values: Map.singleton('Outside.coverage')(LookupTable.lookup('Leaf.value')(table).value0) });
  assert(Capsule.open(outside)(delayed) instanceof Nothing, 'values outside coverage are invalid');
});

test('missing or corrupt optional capsules load complete local fragments and are repaired without global TAST loading', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('first'), Main: main });
  f.compile({}); await f.publish(await f.full());
  f.compile({ Leaf: leaf('second') });
  const first = await f.candidate();
  first.prepared.metadataDemands.value0(keysOf(['Leaf.value']))(); await f.publish(first);
  for (const kind of ['missing', 'corrupt']) {
    const state = readState(f), path = join(f.output, '.gopurs-cache/preparation-literals-v1', state.capsule[0] + '.json');
    if (kind === 'missing') rmSync(path); else writeFileSync(path, '{}');
    f.compile({ Leaf: leaf(kind) });
    const observed = await recordTastReads(() => f.candidate(8));
    const candidate = observed.value;
    assert(candidate.materialize instanceof Just, 'capsule failure uses the verified full projection locally');
    assert(candidate.metadataCoverage instanceof Nothing);
    assert.deepEqual(observed.reads, ['Leaf'], 'unchanged TAST owners remain unloaded');
    await oracle(candidate, await f.full());
    candidate.prepared.metadataDemands.value0(keysOf(['Leaf.value']))(); await f.publish(candidate);
    const repaired = readState(f);
    assert.equal(Wire.hash(readFileSync(join(f.output, '.gopurs-cache/preparation-literals-v1', repaired.capsule[0] + '.json'), 'utf8')), repaired.capsule[0]);
  }
});


test('capsule encoders preserve proof references at jobs 1/8 and publish nothing on encoding failure', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('first'), Main: main });
  f.compile({}); await f.publish(await f.full());
  const directory = join(f.output, '.gopurs-cache/preparation-literals-v1');
  const proofPath = join(directory, 'proof.json');
  const previous = readFileSync(proofPath, 'utf8');
  const inventory = () => fs.readdirSync(directory).sort();
  const initialObjects = inventory();
  f.compile({ Leaf: leaf('second') });
  const candidate = await f.candidate(8);
  candidate.prepared.metadataDemands.value0(keysOf(['Leaf.value', 'Known.absent']))();
  const source = candidate.metadataProjection.value0;
  let calls = 0;
  const faulty = { ...source, modules: array(source.modules).map((mod, i) => i === 1 ? { ...mod, get path() { calls++; throw new Error('injected capsule Wire failure'); } } : mod).reduceRight((tail, head) => new Cons(head, tail), Nil.value) };
  for (const jobs of [1, 8]) {
    await f.publish({ ...candidate, metadataProjection: new Just(faulty) }, jobs);
    assert.equal(readFileSync(proofPath, 'utf8'), previous, 'failed encoders cannot publish a new proof');
    assert.deepEqual(inventory(), initialObjects, 'no object is written before all encoders join');
  }
  assert.equal(calls, 2, 'both configurations exercised Wire failure');
  await f.publish(candidate, 1);
  const sequential = readState(f);
  const bytes = sequential.capsule.map(hash => readFileSync(join(directory, hash + '.json'), 'utf8'));
  writeFileSync(proofPath, previous);
  await f.publish(candidate, 8);
  assert.deepEqual(readState(f), sequential, 'worker count does not change proof or reference order');
  assert.deepEqual(readState(f).capsule.map(hash => readFileSync(join(directory, hash + '.json'), 'utf8')), bytes);
});


test('speculative projection requires the exact accepted proof envelope and session identity', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); await f.publish(await f.full());
  const directory = join(f.output, '.gopurs-cache/preparation-literals-v1');
  const proofPath = join(directory, 'proof.json');
  const envelope = JSON.parse(readFileSync(proofPath, 'utf8'));
  const proof = JSON.parse(envelope.payload);
  const pending = await run(Persistent.preloadProjection(8)(Map.empty)(JSON.stringify({root: f.session.root, context: f.session.context})));
  assert(pending instanceof Just);
  f.compile({ Leaf: leaf('after') });
  const candidate = preload => run(Persistent.preparePersistentWithProjection(preload)(f.session)(8)(Nothing.value)(Map.empty));
  const expected = await f.full();
  // Removing a fragment establishes whether the decoded projection was reused.
  const path = join(directory, proof.projection[0] + '.json');
  const bytes = readFileSync(path);
  rmSync(path);
  await oracle(await candidate(pending), expected);
  for (const altered of [
    {...pending.value0, root: pending.value0.root + '/other'},
    {...pending.value0, context: 'another-context'},
    {...pending.value0, count: pending.value0.count + 1},
  ]) assert.equal(Set.size((await candidate(new Just(altered))).reusable), 0);
  const payload = JSON.stringify(proof, null, 1);
  assert.notEqual(Wire.hash(payload), envelope.hash);
  writeFileSync(proofPath, JSON.stringify({hash: Wire.hash(payload), payload}));
  assert.equal(Set.size((await candidate(pending)).reusable), 0, 'equivalent JSON with a new envelope hash rejects the old decoded result');
  writeFileSync(path, bytes);
  await oracle(await candidate(pending), expected);
  assert(await run(Persistent.preloadProjection(8)(Map.empty)('not json')) instanceof Nothing);
  assert(await run(Persistent.preloadProjection(8)(Map.empty)(JSON.stringify({root: f.session.root, context: 'wrong'}))) instanceof Nothing);
});
