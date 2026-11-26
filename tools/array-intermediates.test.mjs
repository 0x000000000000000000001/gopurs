import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { Just } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { optimize } from '../output/Gopurs.ClosedIndexedFolds/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const e = syntax => new NeutralExpr(syntax);
const typed = (type, value) => e(new S.Typed(type, value));
const int = value => e(new S.Lit(new C.LitInt(value)));
const array = new C.Array(C.Int.value);
const foldType = new C.Func([array], C.Int.value);
const callbackType = new C.Func([C.Int.value, C.Int.value, C.Int.value], C.Int.value);
const local = (name, level, type) => {
  const value = e(new S.Local(new Just(name), level));
  return type ? typed(type, value) : value;
};
const ref = (module, name) => e(new S.Var(new C.Qualified(new Just(module), name)));
const app = (fn, ...args) => e(new S.App(fn, args));
const abs = (names, body) => e(new S.Abs(names.map(([name, level]) => new Tuple(new Just(name), level)), body));
const num = (operation, left, right) => typed(C.Int.value,
  e(new S.PrimOp(new S.Op2(new S.OpIntNum(operation), left, right))));
const add = (a, b) => num(S.OpAdd.value, a, b);
const fold = ref('Data.FoldableWithIndex', 'foldlWithIndexImpl');
const callback = typed(callbackType, abs([['index', 0], ['acc', 1], ['value', 2]],
  add(num(S.OpMultiply.value, local('acc', 1), int(3)), add(local('index', 0), local('value', 2)))));
const partial = (cb = callback, seed = int(17), target = fold) => typed(foldType, app(target, cb, seed));
const moduleOf = (bindings, recursive = false) => ({ name: 'IntermediateFixture',
  bindings: [{ recursive, bindings: bindings.map(([name, value]) => new Tuple(name, value)) }],
  comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
  dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap, implementations: emptyMap, directives: emptyMap });
const transformed = (value, recursive = false) => optimize(moduleOf([['consume', value]], recursive)).bindings[0].bindings[0].value1;
const metadata = withReboxFields({ elidedCtors: emptySet, ctorTypes: emptyMap,
  pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
  enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap, classDeclsFields: emptyMap });

test('closed indexed partial folds expose only a proved native array parameter', () => {
  const original = app(fold, callback, int(17));
  const generic = new C.Func([new C.Array(new C.TypeVar('a'))], new C.TypeVar('b'));
  const arrayAdt = new C.Func([new C.ADT('Prim.Array', ['Prim', 'Array'], [C.Int.value])], C.Int.value);
  const realAnnotations = typed(foldType, typed(arrayAdt, typed(arrayAdt, typed(generic, original))));
  for (const value of [partial(), typed(C.Any.value, partial()), realAnnotations]) {
    const result = transformed(value);
    assert.ok(result instanceof S.Typed);
    assert.deepEqual(result.value0, foldType);
    assert.ok(result.value1 instanceof S.Abs);
    assert.equal(result.value1.value0[0].value1, 3, 'fresh input level follows every callback local');
  }
  const captured = typed(callbackType, abs([['i', 0], ['a', 1], ['v', 2]], add(local('a', 1), local('capture', 99, C.Int.value))));
  const unknown = typed(callbackType, abs([['i', 0], ['a', 1], ['v', 2]], app(ref('Foreign', 'run'), local('v', 2))));
  const staged = typed(callbackType, abs([['i', 0]], e(new S.Let(new Just('captured'), 9, int(1), abs([['a', 1], ['v', 2]], local('a', 1))))));
  for (const value of [partial(captured), partial(unknown), partial(staged), partial(callback, ref('Foreign', 'seed')),
    partial(callback, typed(C.Boolean.value, int(17))), typed(C.Int.value, partial()),
    typed(new C.Func([array], C.Boolean.value), partial()), typed(new C.Func([new C.TypeVar('a')], new C.TypeVar('a')), partial()),
    partial(callback, int(17), ref('Other', 'foldlWithIndexImpl')),
    typed(foldType, app(fold, callback)), typed(foldType, app(fold, callback, int(17), local('x', 4, array)))]) {
    assert.deepEqual(transformed(value), value);
  }
  assert.deepEqual(transformed(partial(), true), partial(), 'recursive groups remain unchanged');
});

function generated() {
  const input = local('input', 4, C.Any.value);
  const right = local('right', 5, C.Any.value);
  const consume = source => typed(C.Int.value, app(ref('IntermediateFixture', 'consume'), source));
  // Retain an inner residual annotation, as in the real optimized TAST.
  const mapped = typed(array, typed(new C.TypeVar('mappedResult'), app(ref('Data.Functor', 'arrayMap'),
    typed(new C.Func([C.Int.value], C.Int.value), abs([['item', 6]], add(local('item', 6), int(1)))), typed(array, input))));
  const concat = rhs => typed(array, app(ref('Data.Semigroup', 'concatArray'), typed(array, input), rhs));
  const wrapper = (names, result) => typed(new C.Func(names.map(() => C.Any.value), C.Int.value), abs(names, result));
  return CodeGen.translate(metadata)(moduleOf([
    ['consume', partial()],
    ['mapped', wrapper([['input', 4]], consume(mapped))],
    ['joined', wrapper([['input', 4], ['right', 5]], consume(concat(typed(array, right))))],
    ['aliasRight', wrapper([['input', 4]], consume(concat(typed(array, app(ref('Foreign', 'right'), int(0))))))],
  ]));
}

test('native producer buffers reach the fold while its generic wrapper keeps the original Value view', t => {
  const code = generated();
  assert.match(code, /func Call_IntermediateFixture_consume\([^\n]*\[\]int64/);
  const getter = code.slice(code.indexOf('func Get_IntermediateFixture_consume()')).split('\nfunc ')[0];
  assert.match(getter, /\*\[\]gopurs_runtime.Value/);
  assert.doesNotMatch(getter, /make\(|GoUnboxIntArray|unboxed|Call_IntermediateFixture_consume/);
  for (const name of ['mapped', 'joined', 'aliasRight']) {
    const worker = code.slice(code.indexOf(`func Call_IntermediateFixture_${name}(`)).split('\nfunc ')[0];
    assert.match(worker, /Call_IntermediateFixture_consume\(/);
    assert.doesNotMatch(worker, /boxed := make|unboxed := make|\[\]gopurs_runtime.Value, len/);
    if (name !== 'mapped') assert.match(worker, /concatInt_output_/);
  }
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-array-intermediates-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, 'gopurs_runtime'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
  writeFileSync(join(directory, 'fixture.go'), code.replace('package purescript', 'package main'));
  writeFileSync(join(directory, 'main.go'), `package main
import("fmt";"reflect";"testing";r "gopurs/output/gopurs_runtime")
var leftAlias []r.Value
var factories int
var sink int64
func Get_Foreign_right()r.Value{return r.Func(func(_ r.Value)r.Value{factories++;if len(leftAlias)>0{leftAlias[0]=r.Int(91)};return r.Array([]r.Value{r.Int(11)})})}
func oracle(xs []int64)int64{acc:=int64(17);for i,v:=range xs{acc=r.IntAdd(r.IntMul(acc,3),r.IntAdd(int64(i),v))};return acc}
func check(ok bool){if !ok{panic("value, order, alias, replay or input failure")}}
func main(){cases:=0;for _,n:=range []int{0,1,2,7,17,1000}{for _,seed:=range []int64{0,-1,2147483647,-2147483648}{
xs:=make([]r.Value,n);ns:=make([]int64,n);mapped:=make([]int64,n);for i:=range xs{ns[i]=r.IntAdd(seed,int64(i));xs[i]=r.Int(ns[i]);mapped[i]=r.IntAdd(ns[i],1)};before:=append([]r.Value{},xs...);boxed:=r.Array(xs)
for replay:=0;replay<3;replay++{check(r.Apply(Get_IntermediateFixture_consume(),boxed).IntVal==oracle(ns));check(Call_IntermediateFixture_consume(ns)==oracle(ns));check(Call_IntermediateFixture_mapped(boxed)==oracle(mapped));check(reflect.DeepEqual(xs,before));cases+=3}
for _,m:=range []int{0,1,3}{ys:=make([]r.Value,m);joined:=append([]int64{},ns...);for i:=range ys{ys[i]=r.Int(int64(i-8));joined=append(joined,ys[i].IntVal)};check(Call_IntermediateFixture_joined(boxed,r.Array(ys))==oracle(joined));check(reflect.DeepEqual(xs,before));cases++}
leftAlias=xs;factories=0;want:=append([]int64{},ns...);if n>0{want[0]=91};want=append(want,11);check(Call_IntermediateFixture_aliasRight(boxed)==oracle(want));check(factories==1);cases++
}}
xs:=make([]r.Value,1000);boxed:=r.Array(xs);getter:=Get_IntermediateFixture_consume();allocs:=testing.AllocsPerRun(10,func(){sink=r.Apply(getter,boxed).IntVal});check(allocs==0)
fmt.Printf("%d native and boxed intermediate checks: ok\\n",cases)}
`);
  const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 60000, env: { ...process.env, GOWORK: 'off' } });
  assert.ifError(result.error); assert.equal(result.status, 0, result.stdout + result.stderr);
  assert.match(result.stdout, /312 native and boxed intermediate checks: ok/);
});

test('array representation demand stays within unary calls, including surplus arguments', () => {
  const input = local('input', 4, C.Any.value);
  const joined = typed(array, app(ref('Data.Semigroup', 'concatArray'), typed(array, input), typed(array, input)));
  const multi = typed(new C.Func([array, C.Int.value], C.Int.value), abs([['input', 4], ['ignored', 5]],
    typed(C.Int.value, app(ref('IntermediateFixture', 'consume'), typed(array, input)))));
  const run = target => typed(new C.Func([C.Any.value], C.Int.value), abs([['input', 4]],
    typed(C.Int.value, app(ref('IntermediateFixture', target), joined, int(0)))));
  const code = CodeGen.translate(metadata)(moduleOf([['consume', partial()], ['multi', multi],
    ['multiCaller', run('multi')], ['surplusCaller', run('consume')]]));
  for (const name of ['multiCaller', 'surplusCaller']) {
    const worker = code.slice(code.indexOf(`func Call_IntermediateFixture_${name}(`)).split('\nfunc ')[0];
    assert.doesNotMatch(worker, /concatInt_|nativeArray_arg_/);
    assert.match(worker, /Get_Data_Semigroup_concatArray/);
  }
});
