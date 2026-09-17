import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import { Curried, recognize, safeIndex } from '../output/Gopurs.ArrayIntrinsics/index.js';
import { uncurriedApplication } from '../output/Gopurs.CallExprs/index.js';
import { StmtEmpty, StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { coerceGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { exprTypeToGoType } from '../output/Gopurs.GoTypes/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as Core from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as Syntax from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const metadata = {
    pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
    enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet,
    ctorTypes: emptyMap, classDeclsFields: emptyMap, globalTypes: emptyMap, globalFunctions: emptyMap,
};
const context = () => ({
    metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 0, modNameStr: 'Probe', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
    tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Nothing.value,
});
const variable = (name, module = Nothing.value) => new TcoExpr(null, new Syntax.Var(new Core.Qualified(module, name)));
const typed = (type, value) => new TcoExpr(null, new Syntax.Typed(type, value));
const args = [variable('just'), variable('nothing'), typed(new Core.Array(Core.Int.value), variable('xs')), variable('i')];
const target = new Just({ mbMod: new Just('Data.Array'), name: 'indexImpl' });

test('safe indexing requires exactly Data.Array.indexImpl with four uncurried arguments', () => {
    const ctx = context();
    const forbidden = () => { throw new Error('unexpected translation'); };
    for (const candidate of [Nothing.value, new Just({ mbMod: new Just('Other'), name: 'indexImpl' }),
        new Just({ mbMod: Nothing.value, name: 'indexImpl' }),
        new Just({ mbMod: new Just('Data.Array'), name: 'unsafeIndex' })]) {
        assert.deepEqual(safeIndex(forbidden)(ctx)(0)(candidate)(args), Nothing.value);
    }
    for (const actual of [args.slice(0, 3), [...args, variable('extra')]]) {
        assert.deepEqual(safeIndex(forbidden)(ctx)(0)(target)(actual), Nothing.value);
    }
    assert.deepEqual(recognize(Curried.value)('Probe')(target)(4), Nothing.value);
    const translation = _ctx => nextId => _arg => ({
        stmts: StmtEmpty.value, expr: new Go.GoVar('value'), exprType: Go.TypeValue.value, nextId,
    });
    assert.ok(safeIndex(translation)({ ...ctx, modNameStr: 'Data_Array' })(0)
        (new Just({ mbMod: Nothing.value, name: 'indexImpl' }))(args) instanceof Just);
});

test('safe indexing preserves bounds, argument order, typed elements and constant allocation cost', t => {
    const recordType = new Core.Record(new Core.Row([new Tuple('x', Core.Int.value)], Nothing.value));
    const fixtures = [
        { name: 'native', sourceType: new Go.TypeNativeArray(Go.TypeInt64.value), parameter: '[]int64', element: Core.Int.value },
        { name: 'boxed', sourceType: Go.TypeValue.value, parameter: 'gopurs_runtime.Value', element: Core.Int.value },
        { name: 'boxedSlice', sourceType: new Go.TypeNativeArray(Go.TypeValue.value), parameter: '[]gopurs_runtime.Value', element: Core.Int.value },
        { name: 'record', sourceType: new Go.TypeNativeArray(Go.TypeValue.value), parameter: '[]gopurs_runtime.Value', element: recordType },
        { name: 'nativeCallback', sourceType: new Go.TypeNativeArray(Go.TypeInt64.value), parameter: '[]int64', element: Core.Int.value, nativeCallback: true },
    ];
    const declarations = fixtures.map(fixture => {
        const ctx = context();
        const inputArgs = [args[0], args[1], typed(new Core.Array(fixture.element), variable('xs')), args[3]];
        const types = {
            just: fixture.nativeCallback ? new Go.TypeFunc([Go.TypeInt64.value], Go.TypeInt64.value) : Go.TypeValue.value,
            nothing: Go.TypeValue.value, xs: fixture.sourceType, i: Go.TypeInt64.value,
        };
        const translate = child => nextId => expression => {
            const syntax = expression.value1;
            if (syntax instanceof Syntax.Typed) {
                const inner = translate(child)(nextId)(syntax.value1);
                const type = exprTypeToGoType(emptyMap)(emptySet)(emptySet)('Probe')(syntax.value0);
                return { ...inner, expr: coerceGoExpr(ctx.codegenStateRef)('Probe')(inner.expr)(inner.exprType)(type), exprType: type };
            }
            assert.ok(syntax instanceof Syntax.Var);
            const name = syntax.value0.value1;
            assert.ok(Object.hasOwn(types, name), `unexpected translated variable ${name}`);
            return {
                stmts: new StmtLeaf(Go.rawGo(`markStatement("${name}")`)),
                expr: Go.rawGo(`markExpression("${name}", ${name === 'nothing' ? 'gopurs_runtime.Int(-1)' : name})`),
                exprType: types[name], nextId,
            };
        };
        const fn = variable('indexImpl', new Just('Data.Array'));
        const result = uncurriedApplication(translate)(ctx)(0)(new TcoExpr(null, new Syntax.UncurriedApp(fn, inputArgs)))(fn)(inputArgs);
        const body = [...flattenStmts(result.stmts), new Go.GoReturn(result.expr)].map(printGoExpr).join('\n');
        const callback = fixture.nativeCallback ? 'func(int64) int64' : 'gopurs_runtime.Value';
        return `func ${fixture.name}(xs ${fixture.parameter}, i int64, just ${callback}) gopurs_runtime.Value {\n${body}\n}`;
    });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-array-safe-index-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "slices"; "testing"; "gopurs/output/gopurs_runtime")
var tracing bool
var trace []string
var sink gopurs_runtime.Value
func markStatement(name string) { if tracing { trace = append(trace, name + ":statement") } }
func markExpression[T any](name string, value T) T { if tracing { trace = append(trace, name + ":expression") }; return value }
func called() { if tracing { trace = append(trace, "called") } }
${declarations.join('\n')}
func main() {
    just := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value { called(); return gopurs_runtime.Int(value.IntVal + 100) })
    justRecord := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value { called(); return gopurs_runtime.Int(gopurs_runtime.RecordGet(value, "x").IntVal + 100) })
    justNative := func(value int64) int64 { called(); return value + 100 }
    values := []int64{7, 9}
    boxedValues := []gopurs_runtime.Value{gopurs_runtime.Int(7), gopurs_runtime.Int(9)}
    boxedArray := gopurs_runtime.Array(boxedValues)
    records := []gopurs_runtime.Value{gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(7)), gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(9))}
    probes := []func(int64) gopurs_runtime.Value{
        func(i int64) gopurs_runtime.Value { return native(values, i, just) },
        func(i int64) gopurs_runtime.Value { return boxed(boxedArray, i, just) },
        func(i int64) gopurs_runtime.Value { return boxedSlice(boxedValues, i, just) },
        func(i int64) gopurs_runtime.Value { return record(records, i, justRecord) },
        func(i int64) gopurs_runtime.Value { return nativeCallback(values, i, justNative) },
    }
    tracing = true
    for probeIndex, probe := range probes { for _, i := range []int64{-1, 0, 1, 2, 9223372036854775807} {
        trace = nil
        result := probe(i)
        expectedTrace := []string{"just:statement", "just:expression", "nothing:statement", "nothing:expression", "xs:statement", "xs:expression", "i:statement", "i:expression"}
        expected := int64(-1)
        if i >= 0 && i < 2 { expected = values[i] + 100; expectedTrace = append(expectedTrace, "called") }
        if result.IntVal != expected || !slices.Equal(trace, expectedTrace) { panic(fmt.Sprintf("probe=%d index=%d result=%v trace=%v", probeIndex, i, result, trace)) }
    } }
    tracing = false
    if native(nil, 0, just).IntVal != -1 || boxed(gopurs_runtime.Array(nil), 0, just).IntVal != -1 { panic("empty array") }
    for _, mode := range []string{"native", "boxed", "boxedSlice"} {
        costs := []float64{}
        for _, size := range []int{16, 16384} {
            xs := make([]int64, size); ys := make([]gopurs_runtime.Value, size)
            for i := range xs { xs[i] = int64(i); ys[i] = gopurs_runtime.Int(int64(i)) }
            array := gopurs_runtime.Array(ys)
            costs = append(costs, testing.AllocsPerRun(20, func() {
                switch mode { case "native": sink = native(xs, 0, just); case "boxed": sink = boxed(array, 0, just); case "boxedSlice": sink = boxedSlice(ys, 0, just) }
            }))
        }
        if costs[1] > costs[0] + 2 || costs[1] > 32 { panic(fmt.Sprintf("%s allocations grow with length: %v", mode, costs)) }
        fmt.Printf("%s allocations=%v\\n", mode, costs)
    }
    fmt.Println("bounds, order, typed elements: ok")
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.match(result.stdout, /bounds, order, typed elements: ok/);
});
