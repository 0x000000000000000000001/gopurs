// Compare the optimized collector with a frozen compiled pre-change collector.
// Both use the same constructors and foreign dependencies from the candidate.
import assert from 'node:assert/strict';
import { existsSync, mkdirSync, readFileSync, writeFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { pathToFileURL } from 'node:url';
import { test } from 'node:test';

const [outputArg, referenceArg, outArg] = process.argv.slice(2);
assert(outputArg && referenceArg && outArg);
const output = resolve(outputArg), out = resolve(outArg);
assert(!existsSync(out)); mkdirSync(out, { recursive: true });
const reference = readFileSync(resolve(referenceArg), 'utf8').replace(/from (["'])([^"']+)\1/g, (_, quote, path) => {
  assert(path.startsWith('../') || path === './foreign.js');
  const target = path.startsWith('../') ? join(output, path.slice(3)) : join(output, 'Gopurs.GoImports/foreign.js');
  return 'from ' + JSON.stringify(pathToFileURL(target).href);
});
const referenceModule = join(out, 'reference.mjs'); writeFileSync(referenceModule, reference);
const [before, after, G, T] = await Promise.all([
  referenceModule, join(output, 'Gopurs.GoImports/index.js'), join(output, 'Gopurs.GoAst/index.js'), join(output, 'Data.Tuple/index.js'),
].map(path => import(pathToFileURL(path))));
const pair = (a, b) => new T.Tuple(a, b), fields = values => values.map((value, i) => pair('v' + i, value));
const raw = path => new G.GoRaw({ text: '', imports: [path] });
const leaves = [raw('math'), raw('sync'), raw('unsafe'), raw('gopurs/output/gopurs_runtime'),
  raw('α'), raw('😀'), raw('\ud800'), new G.GoString('math.Mod'), new G.GoInt(7)];
const types = [G.TypeValue.value, G.TypeInt64.value, G.TypeFloat64.value, G.TypeString.value, G.TypeBool.value, G.TypeUint32.value,
  new G.TypeInterface('unsafe.Pointer'), new G.TypeGenericParam('T')];
types.push(new G.TypeNativeArray(types[0]), new G.TypeRecord(fields(types.slice(0, 7))),
  new G.TypeFunc(types.slice(0, 7), types[6]), new G.TypeStructValue('Pair', types.slice(0, 7)),
  new G.TypeStructPointer({ baseStructName: 'D', fullName: 'D.D', fullPath: 'D', structName: 'D', typeArgs: types.slice(0, 7) }));
const wrap = (a, b, c, ty) => [
  new G.GoVar('math.Pi'), new G.GoCall(a, [b, c]), new G.GoSelector(new G.GoVar('math'), 'Mod'), new G.GoSelector(a, 'field'),
  new G.GoBlock([a, b, c]), new G.GoReturn(a), new G.GoAssign('x', a), new G.GoRecordDict(ty, fields([a, b])),
  new G.GoRecordUpdateDict(a, fields([b, c])), new G.GoRecordUpdateStatic(a, 2, [pair(0, b)], fields([c])),
  new G.GoRecordUpdateNative(ty, a, fields([b, c])), new G.GoIIFE('x', a, b), new G.GoLetRec(fields([a, b]), c),
  new G.GoRecordAccess(a, 'field'), new G.GoStructAccess(a, 'field'), new G.GoRecordAccessStatic(a, 0, 2),
  new G.GoConstructor('D', 'C', [ty], [a, b]), new G.GoConstructor('D', 'C', [ty], []),
  new G.GoConstructorDict('C', [a, b]), new G.GoConstructorAccess(a, 'C', [ty], 0, false),
  new G.GoConstructorAccess(a, 'C', [ty], 0, true), new G.GoBranch([pair(a, b)], c), new G.GoBinOp('+', a, b),
  new G.GoPrefixOp('-', a), new G.GoTypeAssertion(a, 'unsafe.Pointer'), new G.GoIndex(a, b),
  new G.GoBoxStructPointer('D', a), new G.GoBoxIntArray(a), new G.GoUnboxIntArray(a), new G.GoFreshFilterArray(a),
  new G.GoFor('i', [a, b]), new G.GoForRange('gopurs_runtime.Value', [a, b]), new G.GoContinue('loop'),
  new G.GoMutate('x', a), new G.GoIfElse(a, [b], [c]), new G.GoFuncBlock(fields([ty]), [a, b], ty),
  new G.GoFuncLit(fields([ty]), [a, b], c, ty), new G.GoStructValue('D', [ty], [a, b]),
];
let checks = 0;
function check(expr, ty = G.TypeInt64.value) {
  const groups = [[new G.GoFunctionDecl({ name: 'f', params: fields([ty]), result: ty, body: expr }),
    new G.GoInitDecl(expr), new G.GoCachedValue({ identifier: 'x', goType: ty, expression: expr })],
    [new G.GoStructDecl({ name: 'D', typeParams: fields([ty]), fields: fields([ty]) }),
      new G.GoForeignGetter({ name: 'ffi', value: 'ffi' })]];
  // Check each declaration separately so one cannot conceal another's imports.
  for (const declarations of [...groups.flat().map(item => [[item]]), groups]) {
    assert.deepEqual(after.collectImports(declarations), before.collectImports(declarations)); checks++;
  }
}
test('all structured nodes and type forms retain exactly their prior dependencies', () => {
  for (const ty of types) for (const expr of wrap(leaves[0], leaves[1], leaves[2], ty)) check(expr, ty);
  for (const ty of types) check(new G.GoAssign('nilBinding', new G.GoConstructor('D', 'C', [ty], [])));
});
test('nested, wide and duplicate-rich expressions retain sorted unique imports', () => {
  let seed = 0x6a09e667;
  const pick = values => { seed ^= seed << 13; seed ^= seed >>> 17; seed ^= seed << 5; return values[(seed >>> 0) % values.length]; };
  const pool = [...leaves];
  for (let i = 0; i < 500; i++) {
    const expr = pick(wrap(pick(pool), pick(leaves), pick(leaves), pick(types)));
    check(expr); pool.push(expr); if (pool.length > 48) pool.splice(leaves.length, 1);
  }
  check(new G.GoBlock(Array.from({ length: 2048 }, () => pick(leaves))));
  assert.deepEqual(after.collectImports([]), []);
  console.log(`Import collector differential: ${checks} declaration/group comparisons passed`);
});
