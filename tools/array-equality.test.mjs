import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import { emit } from '../output/Gopurs.ArrayEquality/index.js';
import { StmtEmpty, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const context = () => ({
    metadata: withReboxFields({ pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
        enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet, ctorTypes: emptyMap,
        classDeclsFields: emptyMap, globalTypes: emptyMap, globalFunctions: emptyMap }),
    codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 0, modNameStr: 'Probe', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
    tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
});
const variable = (module, name) => new TcoExpr(null, new S.Var(new C.Qualified(module === null ? Nothing.value : new Just(module), name)));
const app = (fn, ...args) => new TcoExpr(null, new S.App(fn, args));
const property = (value, name) => new TcoExpr(null, new S.Accessor(value, new S.GetProp(name)));
const dictionary = scalar => app(variable('Data.Eq', 'eqArray'), variable('Data.Eq', scalar));

test('native array equality requires the canonical scalar dictionary and a saturated call', () => {
    const left = variable('Probe', 'left'), right = variable('Probe', 'right');
    const translate = () => () => () => assert.fail('unsupported equality translated an argument');
    const reject = (fn, args, ctx = context()) => assert.deepEqual(emit(translate)(ctx)(0)(fn)(args), Nothing.value);
    reject(property(dictionary('eqInt'), 'eq'), [left]);
    reject(property(dictionary('eqInt'), 'eq'), [left, right, left]);
    reject(property(dictionary('eqInt'), 'notEq'), [left, right]);
    reject(property(dictionary('eqCustom'), 'eq'), [left, right]);
    reject(property(app(variable('Other', 'eqArray'), variable('Data.Eq', 'eqInt')), 'eq'), [left, right]);
    reject(property(app(variable('Data.Eq', 'eqArray'), variable('Other', 'eqInt')), 'eq'), [left, right]);
    reject(variable('Other', 'eq'), [dictionary('eqInt'), left, right]);
    reject(property(app(variable(null, 'eqArray'), variable(null, 'eqInt')), 'eq'), [left, right]);
});

test('generated equality preserves scalar semantics, input order and lengths', t => {
    const definitions = [];
    for (const [scalar, type] of [
        ['eqInt', Go.TypeInt64.value], ['eqNumber', Go.TypeFloat64.value],
        ['eqString', Go.TypeString.value], ['eqChar', Go.TypeString.value], ['eqBoolean', Go.TypeBool.value],
    ]) for (const boxed of [false, true]) for (const accessor of [false, true]) {
        const dict = dictionary(scalar), left = variable('Probe', 'left'), right = variable('Probe', 'right');
        const fn = accessor ? property(dict, 'eq') : variable('Data.Eq', 'eq');
        const args = accessor ? [left, right] : [dict, left, right];
        const arrayType = boxed ? Go.TypeValue.value : new Go.TypeNativeArray(type);
        const translate = () => nextId => expr => ({ stmts: StmtEmpty.value, nextId,
            expr: Go.rawGo(expr === dict ? 'dict()' : expr === left ? 'left()' : 'right()'),
            exprType: expr === dict ? Go.TypeValue.value : arrayType });
        const result = emit(translate)(context())(0)(fn)(args);
        assert.ok(result instanceof Just);
        const body = [...flattenStmts(result.value0.stmts), new Go.GoReturn(result.value0.expr)].map(printGoExpr).join('\n');
        const name = `${scalar}_${boxed ? 'boxed' : 'native'}_${accessor ? 'accessor' : 'method'}`;
        const goType = boxed ? 'gopurs_runtime.Value' : `[]${Go.goTypeToStr(type)}`;
        definitions.push(`func ${name}(left, right func() ${goType}) bool {\n${body}\n}`);
    }
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-array-equality-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), readFileSync(new URL('../runtime/runtime.go', import.meta.url)));
    writeFileSync(join(directory, 'equality.go'), `package equality
import "gopurs/output/gopurs_runtime"
var trace []int
func dict() gopurs_runtime.Value { trace = append(trace, 0); return gopurs_runtime.Value{} }
${definitions.join('\n')}
`);
    writeFileSync(join(directory, 'equality_test.go'), `package equality
import ("math"; "reflect"; "testing"; r "gopurs/output/gopurs_runtime")
func boxed[A any](xs []A) r.Value { values:=make([]r.Value,len(xs)); for i,x:=range xs {values[i]=r.Box(x)}; return r.Array(values) }
func check[A any](t *testing.T, xs,ys []A,want bool, native func(func()[]A,func()[]A)bool, dynamic func(func()r.Value,func()r.Value)bool) {
  t.Helper(); trace=nil
  got:=native(func()[]A{trace=append(trace,1);return xs},func()[]A{trace=append(trace,2);return ys})
  if got!=want || !reflect.DeepEqual(trace,[]int{0,1,2}) {t.Fatalf("native got %v trace %v",got,trace)}
  trace=nil; got=dynamic(func()r.Value{trace=append(trace,1);return boxed(xs)},func()r.Value{trace=append(trace,2);return boxed(ys)})
  if got!=want || !reflect.DeepEqual(trace,[]int{0,1,2}) {t.Fatalf("boxed got %v trace %v",got,trace)}
}
func TestEquality(t *testing.T) {
 for _,mode:=range []bool{false,true} {
  ints, boxedInts:=eqInt_native_method,eqInt_boxed_method
  numbers, boxedNumbers:=eqNumber_native_method,eqNumber_boxed_method
  strs, boxedStrs:=eqString_native_method,eqString_boxed_method
  chars, boxedChars:=eqChar_native_method,eqChar_boxed_method
  bools, boxedBools:=eqBoolean_native_method,eqBoolean_boxed_method
  if mode {ints,boxedInts=eqInt_native_accessor,eqInt_boxed_accessor;numbers,boxedNumbers=eqNumber_native_accessor,eqNumber_boxed_accessor;strs,boxedStrs=eqString_native_accessor,eqString_boxed_accessor;chars,boxedChars=eqChar_native_accessor,eqChar_boxed_accessor;bools,boxedBools=eqBoolean_native_accessor,eqBoolean_boxed_accessor}
  check(t,[]int64{},[]int64{},true,ints,boxedInts)
  check(t,[]int64{1},[]int64{},false,ints,boxedInts)
  check(t,[]int64{1,2,3},[]int64{1,2,3},true,ints,boxedInts)
  check(t,[]int64{1,2,3},[]int64{1,2,4},false,ints,boxedInts)
  check(t,[]float64{math.NaN()},[]float64{math.NaN()},false,numbers,boxedNumbers)
  check(t,[]float64{0},[]float64{math.Copysign(0,-1)},true,numbers,boxedNumbers)
  check(t,[]string{"😀","é","\\xed\\xa0\\x80"},[]string{"😀","é","\\xed\\xa0\\x80"},true,strs,boxedStrs)
  check(t,[]string{"é"},[]string{"e"},false,chars,boxedChars)
  check(t,[]bool{true,false},[]bool{true,false},true,bools,boxedBools)
  check(t,[]bool{true,false},[]bool{true,true},false,bools,boxedBools)
 }
}
`);
    const result = spawnSync('go', ['test', '.'], { cwd: directory, encoding: 'utf8', timeout: 30_000, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr + result.stdout);
});

test('canonical Int equality reads an explicit scalar view without materializing either array', () => {
    const dict=dictionary('eqInt'),left=variable('Probe','left'),right=variable('Probe','right');
    const translate=()=>nextId=>expr=>({stmts:StmtEmpty.value,nextId,
        expr:expr===dict?Go.rawGo('dictionary()'):new Go.GoUnboxIntArray(Go.rawGo(expr===left?'left()':'right()')),
        exprType:expr===dict?Go.TypeValue.value:new Go.TypeNativeArray(Go.TypeInt64.value)});
    const result=emit(translate)(context())(0)(property(dict,'eq'))([left,right]);
    assert.ok(result instanceof Just);
    const code=[...flattenStmts(result.value0.stmts),result.value0.expr].map(printGoExpr).join('\n');
    assert.doesNotMatch(code,/unboxed := make/);
    assert.match(code,/left\(\)/);assert.match(code,/right\(\)/);
});
