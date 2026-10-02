import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import { Curried, Uncurried, recognize, emitCurried, emitUncurried } from '../output/Gopurs.ArrayIntrinsics/index.js';
import { StmtEmpty } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { boxGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const metadata = withReboxFields({
    pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
    enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet, ctorTypes: emptyMap,
    classDeclsFields: emptyMap, globalTypes: emptyMap, globalFunctions: emptyMap,
});
const context = () => ({
    metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 2, modNameStr: 'Probe', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
    tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Nothing.value,
});
const ref = (module, name) => new TcoExpr(null, new S.Var(new C.Qualified(new Just(module), name)));
const target = (module, name) => new Just({ mbMod: module === null ? Nothing.value : new Just(module), name });
const typed = (type, expr) => new TcoExpr(null, new S.Typed(type, expr));
const int = value => new TcoExpr(null, new S.Lit(new C.LitInt(value)));
const select = (convention, name, count) => recognize(convention)('Probe')(target('Data.Array', name))(count).value0;

test('map/fold admission and convention-specific filter qualification retain their existing boundaries', () => {
    for (const convention of [Curried.value, Uncurried.value]) {
        for (const [name, arity] of [['arrayMap', 2], ['foldlArray', 3]]) {
            for (const module of [null, 'Data.Array', 'Other']) {
                const admit = count => recognize(convention)('Probe')(target(module, name))(count);
                assert.deepEqual(admit(arity - 1), Nothing.value);
                assert.ok(admit(arity) instanceof Just);
                assert.ok(admit(arity + 1) instanceof Just);
            }
        }
        const name = convention === Curried.value ? 'filter' : 'filterImpl';
        const other = convention === Curried.value ? 'filterImpl' : 'filter';
        const admit = (mod, candidate, count = 2) => recognize(convention)(mod)(candidate)(count);
        assert.ok(admit('Probe', target('Data.Array', name)) instanceof Just);
        assert.ok(admit('Data.Array', target(null, name)) instanceof Just);
        for (const candidate of [Nothing.value, target(null, name), target('Other', name), target('Data.Array', other)]) {
            assert.deepEqual(admit('Probe', candidate), Nothing.value);
        }
        assert.deepEqual(admit('Data_Array', target(null, name)), Nothing.value);
        assert.deepEqual(admit('Probe', target('Data.Array', name), 1), Nothing.value);
    }
});

test('fresh integer normalization requires a resolved scalar fold, literal seed and immediate fresh filter', () => {
    const fn = ref('Data.Foldable', 'foldlArray');
    const callback = typed(new C.Func([C.Int.value, C.Int.value], C.Int.value), ref('Probe', 'fold'));
    const source = new Go.GoFreshFilterArray(Go.rawGo('freshFilter()'));
    const runtimeArray = value => new Go.GoCall(new Go.GoSelector(new Go.GoVar('gopurs_runtime'), 'Array'), [value]);
    const roundtrip = value => new Go.GoBoxIntArray(new Go.GoUnboxIntArray(runtimeArray(value)));
    const render = (callee, args, arrayExpr) => printGoExpr(emitCurried(context())(callee)(args)
        (select(Curried.value, 'foldlArray', 3))({ stmts: StmtEmpty.value, nextId: 7,
            exprs: [Go.rawGo('callback'), Go.rawGo('seed'), arrayExpr],
            exprTypes: [Go.TypeValue.value, Go.TypeValue.value, Go.TypeValue.value] }).expr);
    const args = [callback, int(0), ref('Probe', 'array')];
    assert.match(render(fn, args, roundtrip(source)), /source_int_array_/);
    for (const [callee, supplied, value] of [
        [ref('Other', 'foldlArray'), args, roundtrip(source)],
        [fn, [ref('Probe', 'unknown'), ...args.slice(1)], roundtrip(source)],
        [fn, [callback, typed(C.Int.value, ref('Probe', 'seed')), args[2]], roundtrip(source)],
        [fn, [...args, int(1)], roundtrip(source)],
        [fn, args, roundtrip(Go.rawGo('storedFilter'))],
        [fn, args, Go.rawGo('opaqueArray()')],
    ]) assert.doesNotMatch(render(callee, supplied, value), /source_int_array_/);
});

test('generated map/filter/fold loops preserve callbacks, left-fold order, native workers and fresh arrays', t => {
    const declarations = [];
    const kinds = [['map', 'arrayMap'], ['filter', 'filter'], ['fold', 'foldlArray']];
    for (const mode of ['curried', 'boxed', 'native', 'worker']) {
        for (const [kind, intrinsicName] of kinds) {
            const ctx = context();
            const native = mode === 'native' || mode === 'worker';
            const worker = mode === 'worker';
            const arity = kind === 'fold' ? 2 : 1;
            const resultType = kind === 'filter' ? Go.TypeBool.value : Go.TypeInt64.value;
            if (worker) ctx.moduleFunctions = insert(ordString)(`Probe.${kind}`)({
                arity, fArgs: Array(arity).fill(Go.TypeInt64.value), fRet: resultType,
            })(emptyMap);
            const callbackType = worker ? new Go.TypeFunc(Array(arity).fill(Go.TypeInt64.value), resultType) : Go.TypeValue.value;
            const arrayType = native ? new Go.TypeNativeArray(Go.TypeInt64.value) : Go.TypeValue.value;
            // The fold IIFE returns Value; direct workers adapt its boxed accumulator.
            const seedType = Go.TypeValue.value;
            const seed = Go.rawGo('gopurs_runtime.Int(1)');
            const args = [ref('Probe', kind), ...(kind === 'fold' ? [int(1)] : []), ref('Probe', 'values')];
            const translated = { stmts: StmtEmpty.value, nextId: 0,
                exprs: [Go.rawGo(`callback_${kind}`), ...(kind === 'fold' ? [seed] : []), Go.rawGo('values')],
                exprTypes: [callbackType, ...(kind === 'fold' ? [seedType] : []), arrayType] };
            const convention = mode === 'curried' ? Curried.value : Uncurried.value;
            const name = kind === 'filter' && mode !== 'curried' ? 'filterImpl' : intrinsicName;
            const intrinsic = select(convention, name, args.length);
            const result = mode === 'curried'
                ? emitCurried(ctx)(ref('Data.Array', name))(args)(intrinsic)(translated)
                : emitUncurried(ctx)(args)(intrinsic)(translated);
            const boxed = boxGoExpr(ctx.codegenStateRef)('Probe')(result.expr)(result.exprType);
            declarations.push(`func ${mode}_${kind}(values ${native ? '[]int64' : 'gopurs_runtime.Value'}) gopurs_runtime.Value { return ${printGoExpr(boxed)} }`);
        }
    }
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-array-intrinsics-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "slices"; "gopurs/output/gopurs_runtime")
var trace []int64
func Call_Probe_map(v int64) int64 { trace=append(trace,v); return v*2 }
func Call_Probe_filter(v int64) bool { trace=append(trace,v); return v%2==0 }
func Call_Probe_fold(a,v int64) int64 { trace=append(trace,v); return a*10+v }
var callback_map = gopurs_runtime.Func(func(v gopurs_runtime.Value) gopurs_runtime.Value {return gopurs_runtime.Int(Call_Probe_map(v.IntVal))})
var callback_filter = gopurs_runtime.Func(func(v gopurs_runtime.Value) gopurs_runtime.Value {return gopurs_runtime.Bool(Call_Probe_filter(v.IntVal))})
var callback_fold = gopurs_runtime.Func2(func(a,v gopurs_runtime.Value) gopurs_runtime.Value {return gopurs_runtime.Int(Call_Probe_fold(a.IntVal,v.IntVal))})
func values(v gopurs_runtime.Value) []int64 { a:=*(*[]gopurs_runtime.Value)(v.UnsafePtr);out:=make([]int64,len(a));for i,x:=range a{out[i]=x.IntVal};return out }
${declarations.join('\n')}
func main() {
    for _,input:=range [][]int64{nil,{-1,0,2,3}} {
        boxed:=make([]gopurs_runtime.Value,len(input));for i,v:=range input{boxed[i]=gopurs_runtime.Int(v)}
        array:=gopurs_runtime.Array(boxed)
        for _,kind:=range []string{"map","filter","fold"} {
            var expected []int64;fold:=int64(1)
            for _,v:=range input{if kind=="map"{expected=append(expected,v*2)};if kind=="filter"&&v%2==0{expected=append(expected,v)};fold=fold*10+v}
            probes:=map[string][]func()gopurs_runtime.Value{
                ${kinds.map(([kind]) => `"${kind}":{func()gopurs_runtime.Value{return curried_${kind}(array)},func()gopurs_runtime.Value{return boxed_${kind}(array)},func()gopurs_runtime.Value{return native_${kind}(input)},func()gopurs_runtime.Value{return worker_${kind}(input)}}`).join(',\n')},
            }
            for mode,probe:=range probes[kind]{
                trace=nil;result:=probe();if !slices.Equal(trace,input){panic(fmt.Sprintf("%s/%d callback order: %v",kind,mode,trace))}
                if kind=="fold"{if result.IntVal!=fold{panic("left fold")};continue}
                if !slices.Equal(values(result),expected){panic(fmt.Sprintf("%s/%d values: %v",kind,mode,values(result)))}
                repeated:=probe();if len(expected)>0{(*(*[]gopurs_runtime.Value)(result.UnsafePtr))[0]=gopurs_runtime.Int(99)}
                if !slices.Equal(values(repeated),expected)||!slices.Equal(values(array),input){panic("array alias")}
            }
        }
    }
    fmt.Println("map, filter, left fold, callback order and fresh arrays: ok")
}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /callback order and fresh arrays: ok/);
});
