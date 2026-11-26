import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, singleton } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Adt from '../output/Gopurs.AdtExprs/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { coerceGoExpr, generateReboxFunctions } from '../output/Gopurs.GoConversions/index.js';
import { exprTypeToGoType } from '../output/Gopurs.GoTypes/index.js';
import { printGoExpr, printGoDecl } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { withReboxFields } from './codegen-metadata.mjs';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

const put = (key, value, map = emptyMap) => insert(ordString)(key)(value)(map);
const array = new C.Array(C.Int.value);
const maybe = element => new C.ADT('Data.Maybe.Maybe', ['Data', 'Maybe', 'Maybe'], [element]);
const sourceType = maybe(array);
const metadata = withReboxFields({
  pointerAdtPaths: put('Data.Maybe.Maybe', { ctorName: 'Just', arity: 1 }),
  ctorTypes: put('Data_Maybe.Just', { vars: ['a'], fields: [new C.TypeVar('a')] },
    put('Data_Maybe.Nothing', { vars: ['a'], fields: [] })),
  pointerAdtNodes: singleton('Data_Data_Maybe_Just'),
  pointerAdtLeaves: put('Data_Data_Maybe_Nothing', { nodeBaseStruct: 'Data_Data_Maybe_Just', nodeCtor: 'Just' }),
  enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet,
  classDeclsFields: emptyMap, globalTypes: emptyMap, globalFunctions: emptyMap,
});
const nativeType = exprTypeToGoType(metadata.sharedRecords)(metadata.pointerAdtPaths)(emptySet)(emptySet)('MaybeArrays')(sourceType);
const state = () => Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();
const prepare = (name, type = sourceType, expected = new Just(nativeType)) =>
  Adt.prepareSaturatedAs(metadata)('MaybeArrays')(new Just('Data.Maybe'))(name)(type)(expected);
const construct = (name, type = sourceType, expected = new Just(nativeType)) => {
  const prepared = prepare(name, type, expected);
  return Adt.saturated(state())(prepared)({ exprs: name === 'Just' ? [new Go.GoVar('values')] : [],
    exprTypes: name === 'Just' ? [Adt.saturatedFieldType(prepared)(0).goType] : [] });
};

test('explicit canonical Maybe Array Int construction keeps the array and typed nil leaf', () => {
  assert.deepEqual(Adt.saturatedFieldType(prepare('Just'))(0).goType, new Go.TypeNativeArray(Go.TypeInt64.value));
  const just = construct('Just'), nothing = construct('Nothing');
  assert.deepEqual(just.exprType, nativeType); assert.deepEqual(nothing.exprType, nativeType);
  assert.ok(just.expr instanceof Go.GoConstructor);
  assert.match(printGoExpr(nothing.expr), /Constructor_Data_Maybe_Just\[\[\]int64\].*nil/);
  assert.doesNotMatch(printGoExpr(just.expr), /gopurs_runtime.Array|boxed|unboxed/);
});

test('default sum layouts and mismatching type or constructor requests retain their old ABI', () => {
  for (const expected of [Nothing.value, new Just(Go.TypeValue.value),
    new Just(new Go.TypeStructValue('Data.Maybe.Maybe', [Go.TypeValue.value, Go.TypeBool.value]))]) {
    assert.ok(construct('Just', sourceType, expected).exprType instanceof Go.TypeStructValue);
    assert.deepEqual(Adt.saturatedFieldType(prepare('Just', sourceType, expected))(0).goType, Go.TypeValue.value);
  }
  for (const type of [maybe(C.Int.value), maybe(new C.Array(C.String.value)), maybe(C.Any.value)]) {
    assert.ok(construct('Just', type).exprType instanceof Go.TypeStructValue);
  }
  for (const changed of [{ baseStructName: 'Data_Other_Just' }, { fullPath: 'Constructor_Other_Just' },
    { fullName: 'Other.Maybe' }, { symbol: { ...nativeType.value0.symbol, owner: 'generated/Other' } },
    { typeArgs: [Go.TypeValue.value] }]) {
    const expected = new Just(new Go.TypeStructPointer({ ...nativeType.value0, ...changed }));
    assert.ok(construct('Just', sourceType, expected).exprType instanceof Go.TypeStructValue);
  }
  assert.ok(construct('Other').exprType instanceof Go.TypeStructValue);
});

test('native and boxed Maybe arrays preserve every element, Just empty, Nothing and immutable reuse', t => {
  const just = construct('Just'), nothing = construct('Nothing');
  const ref = state();
  const box = coerceGoExpr(ref)('MaybeArrays')(new Go.GoVar('input'))(nativeType)(Go.TypeValue.value);
  const unbox = coerceGoExpr(ref)('MaybeArrays')(new Go.GoVar('input'))(Go.TypeValue.value)(nativeType);
  const reboxes = generateReboxFunctions(metadata)(ref)('MaybeArrays')();
  const code = expression => printGoExpr(expression).replaceAll('gopurs_runtime.', 'r.');
  const nativeName = Go.goTypeToStr(nativeType);
  const directory = goModule(t, 'gopurs-native-maybe-arrays-', { runtime: runtimeGoCode });
  writeGoFile(directory, 'maybe_test.go', `package main
import("testing";"reflect";"unsafe";r "gopurs/output/gopurs_runtime")
type Constructor_Data_Maybe_Just[A any]struct{Rc uint32;V0 A}
func just(values []int64)${nativeName}{return ${code(just.expr)}}
func nothing()${nativeName}{return ${code(nothing.expr)}}
func box(input ${nativeName})r.Value{return ${code(box)}}
func unbox(input r.Value)${nativeName}{return ${code(unbox)}}
${reboxes.map(printGoDecl).join('\n').replaceAll('gopurs_runtime.', 'r.')}
func TestMaybeArrays(t *testing.T){
 if nothing()!=nil||unbox(box(nothing()))!=nil{t.Fatal("Nothing changed")}
 for _,values:=range [][]int64{{},{0},{2147483647,-2147483648,17}}{
  before:=append([]int64{},values...);input:=just(values);if input==nil{t.Fatal("Just empty became Nothing")}
  for replay:=0;replay<3;replay++{boxed:=box(input);out:=unbox(boxed);if out==nil||!reflect.DeepEqual(out.V0,values)||!reflect.DeepEqual(values,before){t.Fatal("payload or source changed")}}
 }
}
`);
  runGo(directory, ['test', '-count=1', '.']);
});
