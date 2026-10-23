import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, mkdirSync, rmSync, writeFileSync, readFileSync, copyFileSync, statSync, utimesSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { Left } from '../output/Data.Either/index.js';
import { Cons } from '../output/Data.List.Types/index.js';
import * as Map from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import * as Set from '../output/Data.Set/index.js';
import * as Aff from '../output/Effect.Aff/index.js';
import * as Prepare from '../output/Gopurs.Driver.Prepare/index.js';
import * as Persistent from '../output/Gopurs.Driver.Prepare.Persistent/index.js';
import * as Wire from '../output/Gopurs.Cache.Wire/index.js';
import { findTypedCompiler } from './native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const compiler = findTypedCompiler(root, process.env.GOPURS_PURS);
const run = action => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(action)());
const array = list => { const values = []; for (; list instanceof Cons; list = list.value1) values.push(list.value0); return values; };
const wire = mod => Wire.write(Wire.codecModule(Wire.codecAnn))(mod)();
function setup(t, sources) {
  const workspace = mkdtempSync(join(tmpdir(), 'gopurs-compact-proof-'));
  t.after(() => rmSync(workspace, { recursive: true, force: true }));
  const src = join(workspace, 'src'), output = join(workspace, 'output'), stage = join(workspace, 'stage');
  mkdirSync(src); mkdirSync(join(stage, 'prior'), { recursive: true });
  const stamp = () => {
    const digest = Wire.hash(Object.keys(sources).sort().map(name => readFileSync(join(output, name, 'corefn.json'), 'utf8')).join('\n'));
    writeFileSync(join(stage, 'session.json'), JSON.stringify({ snapshot: { compiler: 'compiler-test-v1', inputs: [{ name: 'tast', kind: 'tast', digest }] } }));
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
  const candidate = () => run(Persistent.preparePersistent(session)(1)(Nothing.value)(Map.empty));
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

test('an empty specialization index preserves untyped private bindings and source exports', async t => {
  const secret = text => `module Secret (public) where\nprivateValue :: String\nprivateValue = "${text}"\npublic :: String\npublic = privateValue\n`;
  const f = setup(t, { Secret: secret('before'), Main: 'module Main (answer) where\nimport Secret (public)\nanswer :: String\nanswer = public\n' });
  const compileUntypedPrivate = text => {
    f.compile({ Secret: secret(text) });
    const path = join(f.output, 'Secret', 'corefn.json');
    const json = JSON.parse(readFileSync(path, 'utf8'));
    const binding = json.decls.find(binding => binding.identifier === 'privateValue');
    assert(binding, 'fixture retains its private binding');
    binding.annotation.type = null;
    binding.expression.annotation.type = null;
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
