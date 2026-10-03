// Adapted from purust/purust/tools/test-native-type-table.mjs; see that file
// for the differential contract (native type-table resolution, error/fallback
// precedence, frozen module tables). The candidate is the gopurs PBO
// CoreFn/Json.rs injected at NATIVE_FFI; the oracles come from the retained
// Rust-hosted gopurs workspace. Only paths, logging and the shared cargo
// target directory differ. Like the Purust contract, this case needs a
// directory corpus (<Module>/corefn.json), not the packed file form.
//
// Usage: node tools/native-pbo/type-table.mjs GENERATED_RUST TAST_CORPUS
import assert from 'node:assert/strict';
import { existsSync, readFileSync, readdirSync, writeFileSync } from 'node:fs';
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
const cases = [];
const add = (mode, table) => cases.push(`${mode}\t${JSON.stringify(table)}`);
const constructors = [
  'Int', 'Number', 'String', 'Char', 'Boolean', 'Unit', 'Any',
  { type: 'TypeVar', name: 'a' }, { TypeVar: 'b' },
  { type: 'TypeLevelString', value: [0xd800, 0xdc00, 0xdfff] },
  { type: 'Adt', fqn: ['X', 'é😀'], args: [0, 2, 7] },
  { type: 'TypeApp', constructor: 10, args: [0, 7] },
  { type: 'Func', args: [0, 2], ret: 7 },
  { type: 'Array', element: 11 },
  { type: 'Row', fields: [{ label: [0xd800], type: 0 }, { label: 'x', type: 13 }], tail: null },
  { type: 'Row', fields: [], tail: 14 },
  { type: 'Record', row: 15 },
  { type: 'ForAll', vars: ['a', [0xdfff]], body: 16 },
  { type: 'ConstrainedType', constraints: [{ fqn: ['Data', 'Eq'], args: [7] }], body: 17 },
];
function remap(type, n) {
  if (typeof type !== 'object') return type;
  const copy = structuredClone(type), ref = id => n - 1 - id;
  for (const field of ['constructor', 'ret', 'element', 'row', 'body', 'tail']) {
    if (typeof copy[field] === 'number') copy[field] = ref(copy[field]);
  }
  if (copy.args) copy.args = copy.args.map(ref);
  if (copy.fields) copy.fields.forEach(field => { field.type = ref(field.type); });
  if (copy.constraints) copy.constraints.forEach(constraint => { constraint.args = constraint.args.map(ref); });
  return copy;
}
add('fast', []);
add('fast', constructors);
add('fast', constructors.map(type => remap(type, constructors.length)).reverse());
add('fast', [[73, 110, 116], { type: [65, 114, 114, 97, 121], element: 0 }]);
for (const type of constructors) {
  if (typeof type === 'object') {
    for (const key of Object.keys(type)) {
      const bad = structuredClone(type); delete bad[key];
      if (key !== 'tail') add('either', [...constructors, bad]);
    }
  }
}
const invalid = [null, false, 123, '', 'Unknown', {}, { type: 'Unknown' }, { TypeVar: 7 },
  { type: 'Array', element: -1 }, { type: 'Array', element: 999999 },
  { type: 'Array', element: 0.5 }, { type: 'Array', element: 2147483648 },
  { type: 'Array', element: 0 }, { type: 'TypeLevelString', value: [65536] },
  { type: 'Row', fields: [{ label: 'x', type: 0 }], tail: false },
  { type: 'TypeApp', constructor: 0, args: false },
  { type: 'ConstrainedType', constraints: [{ fqn: ['C'], args: [0] }], body: false }];
for (const bad of invalid) add('fallback', [bad]);
add('fallback', [{ type: 'Array', element: 1 }, { type: 'Record', row: 0 }]);
for (const first of invalid) for (const second of invalid.slice(0, 8)) add('either', [first, second]);
let seed = 42;
const next = () => (seed = (Math.imul(seed, 1664525) + 1013904223) >>> 0);
for (let sample = 0; sample < 200; sample++) {
  const table = ['Int', 'String', { type: 'TypeVar', name: 'a' }];
  for (let i = 3; i < 30; i++) {
    const ref = next() % i;
    switch (next() % 6) {
      case 0: table.push({ type: 'Array', element: ref }); break;
      case 1: table.push({ type: 'Func', args: [0, 2], ret: ref }); break;
      case 2: table.push({ type: 'ForAll', vars: ['a'], body: ref }); break;
      case 3: table.push({ type: 'Row', fields: [{ label: 'x', type: ref }] }); break;
      case 4: table.push({ type: 'Adt', fqn: ['M', `T${i}`], args: [ref] }); break;
      default: table.push({ type: 'ConstrainedType', constraints: [{ fqn: ['C'], args: [0, 2] }], body: ref });
    }
  }
  add('fast', table);
  add('fast', table.map(type => remap(type, table.length)).reverse());
}
let corpusTables = 0;
for (const module of readdirSync(corpus, { withFileTypes: true })) {
  if (!module.isDirectory()) continue;
  const path = join(corpus, module.name, 'corefn.json');
  if (!existsSync(path)) continue;
  const value = JSON.parse(readFileSync(path, 'utf8'));
  add('either', value.typeTable); corpusTables++;
}
assert(corpusTables > 0, `No frozen module tables in ${corpus}`);
writeFileSync(join(directory, 'cases.ndjson'), cases.join('\n') + '\n');
  return corpusTables;
}

await runCase("type-table", rust, async ctx => {
  const directory = createWorkspace("gopurs-native-pbo-type-table");
  ctx.state.workspace = directory;
  ctx.note(`fixture workspace ${directory}`);
  const corpusTables = buildInputs(directory);

  const modules = ['purust_core', 'perceus_ptr', 'Purs_Data_Argonaut_Core', 'Purs_Data_Argonaut_Decode_Error', 'Purs_Data_Either',
    'Purs_Data_Maybe', 'Purs_Data_Tuple', 'Purs_Foreign_Object', 'Purs_PureScript_Backend_Optimizer_CoreFn',
    'Purs_PureScript_Backend_Optimizer_CoreFn_TypeTable', 'Purs_PureScript_Backend_Optimizer_CoreFn_Json',
    'Purs_Data_Map_Internal', 'Purs_Data_Ord', 'Purs_Data_Foldable'];
  writeCrate(directory, 'purust_native_type_table_test', modules, rust);

  ctx.note(`injecting gopurs ${join(pboDir, 'src/PureScript/Backend/Optimizer/CoreFn/Json.rs')}`);
  const ffi = threadedRust(pboSource('PureScript/Backend/Optimizer/CoreFn/Json.rs'));
  writeFileSync(join(directory, 'src/main.rs'),
    fixtureSource('test-native-type-table.rs').replace(/^\s*\/\/ NATIVE_FFI$/m, () => ffi));
  ctx.note(`${corpusTables} frozen module tables from ${corpus}`);

  const executable = ctx.cargoBuild(directory, 'purust_native_type_table_test');
  ctx.spawn('run', executable, [join(directory, 'cases.ndjson')]);
});
