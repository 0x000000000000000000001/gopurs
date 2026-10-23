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
  [name, Map.toUnfoldable(unfoldableArray)(values).map(pair => [pair.value0, pair.value1])]));
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
  const publish = async cache => {
    await run(Persistent.publishPersistent(session)(cache));
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

test('literal edit reuses anchored compact preparation and materializes every reader exactly', async t => {
  const f = setup(t, { Poly: poly, Leaf: leaf('before'), Main: main });
  f.compile({}); const initial = await f.full(); await f.publish(initial);
  f.compile({ Leaf: leaf('a longer label') });
  const candidate = await f.candidate();
  assert.deepEqual(Set.toUnfoldable(unfoldableArray)(candidate.reusable), ['Main', 'Poly']);
  await oracle(candidate, await f.full());
  await f.publish(candidate);
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
  assert.equal(proof.schema, 4);
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
    const donor = fragments.findIndex(fragment => Map.size(fragment.value.metadata[field]) > 0);
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
