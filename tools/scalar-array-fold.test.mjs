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
import { emit } from '../output/Gopurs.ScalarArrayFold/index.js';
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
const ref = (module, name) => new TcoExpr(null, new S.Var(new C.Qualified(new Just(module), name)));
const fold = ref('Data.Foldable', 'foldlArray');
const add = ref('Data.Semiring', 'intAdd');
const seed = ref('Probe', 'seed'), source = ref('Probe', 'source');
test('scalar fold admits only saturated canonical callbacks and foldlArray', () => {
    for (const [fn, args] of [
        [ref('Other', 'foldlArray'), [add, seed, source]],
        [fold, [ref('Other', 'intAdd'), seed, source]],
        [fold, [ref('Probe', 'callback'), seed, source]],
        [fold, [add, seed]], [fold, [add, seed, source, source]],
    ]) assert.deepEqual(emit(() => assert.fail('translated rejected argument'))(context())(0)(fn)(args), Nothing.value);
});
test('scalar fold retains argument order, empty seeds, Int wrap and string append semantics', t => {
    const functions = [];
    for (const strings of [false, true]) for (const boxed of [false, true]) {
        const scalar = strings ? Go.TypeString.value : Go.TypeInt64.value;
        const callback = strings ? ref('Data.Semigroup', 'concatString') : add;
        const args = [callback, seed, source];
        const types = [Go.TypeValue.value, Go.TypeValue.value, boxed ? Go.TypeValue.value : new Go.TypeNativeArray(scalar)];
        const translation = _ctx => nextId => arg => {
            const i = args.indexOf(arg);
            if (i === 2) assert.deepEqual(_ctx.mbExpectedGoType, new Just(Go.TypeValue.value));
            return { stmts: StmtEmpty.value, expr: Go.rawGo(['callback()', 'initial()', 'input()'][i]), exprType: types[i], nextId };
        };
        const result = emit(translation)(context())(0)(fold)(args).value0;
        const body = [...flattenStmts(result.stmts), new Go.GoReturn(result.expr)].map(printGoExpr).join('\n');
        assert.doesNotMatch(body, /Apply2|make\(\[\]int64/);
        functions.push(`func ${strings ? 'concat' : 'sum'}${boxed ? 'Boxed' : 'Native'}(initial func() r.Value, input func() ${Go.goTypeToStr(types[2])}) ${Go.goTypeToStr(scalar)} {\n${body}\n}`);
    }
    const dir = mkdtempSync(join(tmpdir(), 'gopurs-scalar-fold-'));
    t.after(() => rmSync(dir, { recursive: true, force: true }));
    mkdirSync(join(dir, 'gopurs_runtime'));
    writeFileSync(join(dir, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    for (const file of ['runtime.go', 'simd_asm.go', 'simd_generic.go', 'simd_arm64.s', 'simd_amd64.s']) {
        writeFileSync(join(dir, 'gopurs_runtime', file), readFileSync(new URL(`../runtime/${file}`, import.meta.url)));
    }
    writeFileSync(join(dir, 'fold.go'), `package fold
import (r "gopurs/output/gopurs_runtime"; "gopurs/output/gopurs_runtime")
var trace []int
func callback() r.Value { trace=append(trace,0);return r.Value{} }
${functions.join('\n')}`);
    writeFileSync(join(dir, 'fold_test.go'), `package fold
import ("testing"; "reflect"; "math/big"; "math/rand"; r "gopurs/output/gopurs_runtime")
func TestFold(t *testing.T) {
 for _,xs:=range [][]int64{nil,{1,2,3},{2147483647,1},{-2147483648,-1}} {
  want:=int64(7);values:=make([]r.Value,len(xs));for i,v:=range xs {want=r.IntAdd(want,v);values[i]=r.Int(v)}
  for _,boxed:=range []bool{false,true} {
   trace=nil;seed:=func()r.Value{trace=append(trace,1);return r.Int(7)}
   var got int64
   if boxed {got=sumBoxed(seed,func()r.Value{trace=append(trace,2);return r.Array(values)})} else {got=sumNative(seed,func()[]int64{trace=append(trace,2);return xs})}
   if got!=want || !reflect.DeepEqual(trace,[]int{0,1,2}) {t.Fatalf("sum %d != %d trace %v",got,want,trace)}
   for i,v:=range xs {if values[i].IntVal!=v {t.Fatal("mutated input")}}
  }
 }
 for _,xs:=range [][]string{nil,{"a","b"},{"",string([]byte{237,176,128}),"x"}} {
  seedString:=string([]byte{237,160,128});want:=seedString;values:=make([]r.Value,len(xs));for i,v:=range xs {want=r.ConcatString(want,v);values[i]=r.Str(v)}
  for _,boxed:=range []bool{false,true} {
   trace=nil;seed:=func()r.Value{trace=append(trace,1);return r.Str(seedString)}
   var got string
   if boxed {got=concatBoxed(seed,func()r.Value{trace=append(trace,2);return r.Array(values)})} else {got=concatNative(seed,func()[]string{trace=append(trace,2);return xs})}
   if got!=want || !reflect.DeepEqual(trace,[]int{0,1,2}) {t.Fatalf("concat %q != %q trace %v",got,want,trace)}
  }
 }
}
func TestAssociativeIntegerSum(t *testing.T) {
 rng:=rand.New(rand.NewSource(77));count:=0
 for _,n:=range []int{0,1,2,3,4,5,6,7,8,9,15,16,17,31,32,33,63,64,65,999,1000,1001,10000} {
  for mode:=0;mode<5;mode++ {
   xs:=make([]int64,n);values:=make([]r.Value,n)
   for i:=range xs {switch mode {case 0:xs[i]=int64(i);case 1:xs[i]=2147483647;case 2:xs[i]=-2147483648;case 3:xs[i]=int64(rng.Uint64());default:xs[i]=int64(rng.Int31())-1073741824};values[i]=r.Int(xs[i])}
   for _,initial:=range []int64{0,1,-1,2147483647,-2147483648,4294967295,1<<62,-1<<63} {
    exact:=big.NewInt(initial);for _,v:=range xs {exact.Add(exact,big.NewInt(v))};exact.Mod(exact,big.NewInt(1<<32));want:=int64(int32(exact.Int64()));if n==0 {want=initial}
    for _,boxed:=range []bool{false,true} {
     trace=nil;seed:=func()r.Value{trace=append(trace,1);return r.Int(initial)};var got int64
     if boxed {got=sumBoxed(seed,func()r.Value{trace=append(trace,2);return r.Array(values)})}else{got=sumNative(seed,func()[]int64{trace=append(trace,2);return xs})}
     if got!=want || !reflect.DeepEqual(trace,[]int{0,1,2}) {t.Fatalf("n%d mode%d seed%d boxed%v got%d want%d trace%v",n,mode,initial,boxed,got,want,trace)}
     for i,v:=range xs {if values[i].IntVal!=v {t.Fatal("mutated input")}};count++
    }
   }
  }
 }
 t.Logf("%d generated Go / independent BigInt comparisons",count)
}
`);
    const run = spawnSync('go', ['test', '-v', '-count=1', './...'], { cwd: dir, encoding: 'utf8', env: { ...process.env, GOWORK: 'off' } });
    assert.equal(run.status, 0, run.stdout + run.stderr);
    assert.match(run.stdout, /1840 generated Go \/ independent BigInt comparisons/);
});
