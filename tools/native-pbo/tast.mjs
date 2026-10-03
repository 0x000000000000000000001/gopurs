// Adapted from purust/purust/tools/test-native-tast.mjs; see that file for the
// differential contract (decodeArrayImpl callback/error order, native
// decodeAnnWithUsageImpl vs the PureScript oracle, frozen annotation corpus).
// The candidate is the gopurs PBO CoreFn/Json.rs injected at NATIVE_FFI; the
// oracles come from the retained Rust-hosted gopurs workspace. Only paths,
// logging and the shared cargo target directory differ.
//
// Usage: node tools/native-pbo/tast.mjs GENERATED_RUST [TAST_CORPUS]
import assert from 'node:assert/strict';
import { existsSync, readFileSync, readdirSync, statSync, writeFileSync } from 'node:fs';
import { join } from 'node:path';
import {
  createWorkspace, fixtureSource, loadThreadedRust, pboDir, pboSource,
  resolveCaseArgs, runCase, writeCrate,
} from './shared.mjs';

const { rust, corpus } = resolveCaseArgs({ corpus: true });
const threadedRust = await loadThreadedRust();

// The case tables below are kept verbatim from the Purust harness; only the
// output directory is supplied by this adapter.
function buildInputs(directory) {

// decodeArrayImpl cases. The decoder succeeds on every value except null, so
// an input null is the deterministic first callback error at that index.
const arrays = [];
const arrayCase = (mode, fail, value) => arrays.push(`${mode}\t${fail}\t${JSON.stringify(value)}`);
arrayCase('fast', '-', []);
arrayCase('fast', '-', [0, -1, 2.5, -0.5, 'x', '', true, false, [], [[1], ['a']], {}, { a: 1, b: ['x'] }]);
arrayCase('fast', '-', ['\ud800\udfff']);
arrayCase('fast', '-', Array.from({ length: 257 }, (_, i) => (i % 2 === 0 ? i : `s${i}`)));
arrayCase('error', 0, [null]);
arrayCase('error', 1, [1, null, 2]);
arrayCase('error', 3, [1, 2, 3, null]);
arrayCase('error', 0, [null, 1, 2]);
arrayCase('error', 2, ['a', 'b', null, 'd']);
// Deterministic pseudo-random arrays; the first null, when present, is the
// expected callback error index.
let seed = 0x5eed1234;
const next = () => (seed = (Math.imul(seed, 1664525) + 1013904223) >>> 0);
for (let sample = 0; sample < 64; sample++) {
  const length = next() % 40;
  const values = Array.from({ length }, () => {
    switch (next() % 6) {
      case 0: return null;
      case 1: return next() % 1000 - 500;
      case 2: return (next() % 100) / 4;
      case 3: return `s${next() % 1000}é`;
      case 4: return (next() & 1) === 0;
      default: return [next() % 10, `x${next() % 10}`];
    }
  });
  const first = values.indexOf(null);
  if (first < 0) arrayCase('fast', '-', values);
  else arrayCase('error', first, values);
}
for (const value of [null, 0, 1.5, -1, '', 'x', true, false, { a: 1 }]) {
  arrays.push(`nonarray\t-\t${JSON.stringify(value)}`);
}
writeFileSync(join(directory, 'arrays.ndjson'), arrays.join('\n') + '\n');

// Annotation cases against a small shared type table. "fast" is a valid shape
// that must stay native, "fallback" must delegate exactly once.
const typeTable = ['Int', 'String', { type: 'Array', element: 0 }, { type: 'Adt', fqn: ['M', 'T'], args: [0] }];
const annotations = [];
const ann = (mode, input, module = 'Test.Module') =>
  annotations.push(`${mode}\t${module}\t${JSON.stringify(typeTable)}\t${JSON.stringify(input)}`);
ann('fast', {});
ann('fast', { meta: null, type: null, bindingUsage: null, variableUse: null });
ann('fast', { sourceSpan: { start: [1, 1], end: [2, 3] } });
ann('fast', { meta: { metaType: 'IsNewtype' } });
ann('fast', { meta: { metaType: 'IsTypeClassConstructor' } });
ann('fast', { meta: { metaType: 'IsForeign' } });
ann('fast', { meta: { metaType: 'IsWhere' } });
ann('fast', { meta: { metaType: 'IsSyntheticApp' } });
ann('fast', { meta: { metaType: 'IsConstructor', constructorType: 'ProductType', identifiers: [] } });
ann('fast', { meta: { metaType: 'IsConstructor', constructorType: 'SumType', identifiers: ['a', 'é', '😀', '\ud800'] } });
for (let i = 0; i < typeTable.length; i++) ann('fast', { type: i });
ann('fast', { type: 0, meta: { metaType: 'IsNewtype' } });
ann('fast', { type: -0 });
ann('fast', { type: -1 });
ann('fast', { type: -2147483648 });
ann('fast', { type: 4 });
ann('fast', { type: 2147483648 });
ann('fallback', { type: -2147483649 });
ann('fallback', { type: 2147483649 });
ann('fallback', { type: 1.5 });
ann('fallback', { type: 1e300 });
ann('fallback', { type: '0' });
ann('fallback', { type: true });
ann('fallback', { type: {} });
ann('fallback', { type: [] });
ann('fallback', { meta: { metaType: 'Bogus' } });
ann('fallback', { meta: {} });
ann('fallback', { meta: { metaType: 7 } });
ann('fallback', { meta: { metaType: null } });
ann('fallback', { meta: { metaType: 'IsConstructor' } });
ann('fallback', { meta: { metaType: 'IsConstructor', constructorType: 'Bogus', identifiers: [] } });
ann('fallback', { meta: { metaType: 'IsConstructor', constructorType: 'ProductType', identifiers: 'x' } });
ann('fallback', { meta: { metaType: 'IsConstructor', constructorType: 'ProductType', identifiers: [7] } });
ann('fallback', { meta: 'x' });
ann('fallback', { meta: [] });
ann('fallback', { meta: { metaType: 'IsConstructor', constructorType: 'ProductType', identifiers: [null] } });
ann('fast', { bindingUsage: { bindingId: 0 } });
ann('fast', { bindingUsage: { bindingId: 2147483647, maxUses: 0, hasEscapingUseContext: true } });
ann('fast', { bindingUsage: { bindingId: 0, maxUses: null, hasEscapingUseContext: null } });
ann('fast', { bindingUsage: { bindingId: 0, maxUses: 2147483647 } });
ann('fast', { bindingUsage: { bindingId: 0, maxUses: 2147483648 } });
ann('fast', { bindingUsage: { bindingId: 0, maxUses: 1e300 } });
ann('fast', { variableUse: { bindingId: 1 } });
ann('fast', { variableUse: { bindingId: 1, lastLocalUse: null } });
ann('fast', { variableUse: { bindingId: 1, lastLocalUse: true } });
ann('fast', { bindingUsage: { bindingId: 0, hasEscapingUseContext: false }, variableUse: { bindingId: 1, lastLocalUse: true } });
ann('fallback', { bindingUsage: {} });
ann('fallback', { bindingUsage: { bindingId: null } });
ann('fallback', { bindingUsage: { bindingId: -1 } });
ann('fallback', { bindingUsage: { bindingId: 2147483648 } });
ann('fallback', { bindingUsage: { bindingId: 1.5 } });
ann('fallback', { bindingUsage: { bindingId: '0' } });
ann('fallback', { bindingUsage: { bindingId: 0, maxUses: -1 } });
ann('fallback', { bindingUsage: { bindingId: 0, maxUses: 0.5 } });
ann('fallback', { bindingUsage: { bindingId: 0, maxUses: 'x' } });
ann('fallback', { bindingUsage: { bindingId: 0, hasEscapingUseContext: 'yes' } });
ann('fallback', { variableUse: {} });
ann('fallback', { variableUse: { bindingId: 0, lastLocalUse: false } });
ann('fallback', { variableUse: { bindingId: 0, lastLocalUse: 0 } });
ann('fallback', { bindingUsage: 3 });
ann('fallback', { variableUse: [] });
ann('fallback', { bindingUsage: { bindingId: 0 }, variableUse: {} });
ann('fallback', { meta: { metaType: 'Bogus' }, type: 1.5 });
ann('fallback', { type: 1.5, bindingUsage: {} });
ann('fallback', { bindingUsage: {}, variableUse: {} });
ann('fast', { bindingUsage: { bindingId: 3 } }, 'M.é😀');
ann('fast', { variableUse: { bindingId: 9, lastLocalUse: true } }, 'M.surrogate');
writeFileSync(join(directory, 'annotations.ndjson'), annotations.join('\n') + '\n');

// Frozen corpus: every real annotation is paired with its module type table.
// Both a directory of <Module>/corefn.json trees and the packed
// [{ name, contents }] diagnostic corpus are accepted.
function corpusEntries() {
  if (statSync(corpus).isFile()) {
    return JSON.parse(readFileSync(corpus, 'utf8')).map(entry => ({ name: entry.name, value: JSON.parse(entry.contents) }));
  }
  const entries = [];
  for (const entry of readdirSync(corpus, { withFileTypes: true })) {
    if (!entry.isDirectory()) continue;
    const file = join(corpus, entry.name, 'corefn.json');
    if (!existsSync(file)) continue;
    entries.push({ name: entry.name, value: JSON.parse(readFileSync(file, 'utf8')) });
  }
  return entries;
}
const isAnnotation = value => value && typeof value === 'object' && !Array.isArray(value)
  && Object.prototype.hasOwnProperty.call(value, 'meta')
  && Object.prototype.hasOwnProperty.call(value, 'sourceSpan');
function collectAnnotations(value, out) {
  if (Array.isArray(value)) {
    for (const item of value) collectAnnotations(item, out);
    return;
  }
  if (value && typeof value === 'object') {
    if (isAnnotation(value)) out.push(value);
    for (const key of Object.keys(value)) collectAnnotations(value[key], out);
  }
}
const corpusLines = [];
let corpusModules = 0, corpusAnnotations = 0;
for (const entry of corpusEntries()) {
  const annotations = [];
  collectAnnotations(entry.value, annotations);
  if (annotations.length === 0) continue;
  const name = Array.isArray(entry.value.moduleName) ? entry.value.moduleName.join('.') : entry.name;
  corpusLines.push([name, JSON.stringify(entry.value.typeTable ?? []), JSON.stringify(annotations)].join('\t'));
  corpusModules++;
  corpusAnnotations += annotations.length;
}
assert(corpusModules > 0, `no corefn modules found in ${corpus}`);
writeFileSync(join(directory, 'corpus.ndjson'), corpusLines.join('\n') + '\n');

  return { corpusModules, corpusAnnotations };
}

await runCase("tast", rust, async ctx => {
  const directory = createWorkspace("gopurs-native-pbo-tast");
  ctx.state.workspace = directory;
  ctx.note(`fixture workspace ${directory}`);
  const counts = buildInputs(directory);

  const modules = ['purust_core', 'perceus_ptr', 'Purs_Data_Argonaut_Core', 'Purs_Data_Argonaut_Decode_Error',
    'Purs_Data_Either', 'Purs_Data_Maybe', 'Purs_Data_Tuple', 'Purs_Foreign_Object', 'Purs_PureScript_Backend_Optimizer_CoreFn',
    'Purs_PureScript_Backend_Optimizer_CoreFn_Json', 'Purs_PureScript_Backend_Optimizer_CoreFn_TypeTable',
    'Purs_Data_Map_Internal', 'Purs_Data_Ord', 'Purs_Data_Foldable'];
  writeCrate(directory, 'purust_native_tast_test', modules, rust);

  ctx.note(`injecting gopurs ${join(pboDir, 'src/PureScript/Backend/Optimizer/CoreFn/Json.rs')}`);
  const ffi = threadedRust(pboSource('PureScript/Backend/Optimizer/CoreFn/Json.rs'));
  writeFileSync(join(directory, 'src/main.rs'),
    fixtureSource('test-native-tast.rs').replace(/^\s*\/\/ NATIVE_FFI$/m, () => ffi));
  ctx.note(`${counts.corpusModules} corpus modules / ${counts.corpusAnnotations} annotations from ${corpus}`);

  const executable = ctx.cargoBuild(directory, 'purust_native_tast_test');
  ctx.spawn('run', executable, [join(directory, 'arrays.ndjson'), join(directory, 'annotations.ndjson'), join(directory, 'corpus.ndjson')]);
});
