import { withReboxFields } from './codegen-metadata.mjs';
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
import { unsafeIndex } from '../output/Gopurs.ArrayIntrinsics/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as Core from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as Syntax from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const metadata = withReboxFields({
    pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
    enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet,
    ctorTypes: emptyMap, classDeclsFields: emptyMap, globalTypes: emptyMap, globalFunctions: emptyMap,
});
const variable = name => new TcoExpr(null, new Syntax.Var(new Core.Qualified(Nothing.value, name)));
const typed = (type, value) => new TcoExpr(null, new Syntax.Typed(type, value));
const recordType = new Core.Record(new Core.Row([new Tuple('x', Core.Int.value)], Nothing.value));
const nativeRecord = new Go.TypeRecord([new Tuple('x', Go.TypeInt64.value)]);

// Exercise the public compiler entry point too, so a disconnected helper cannot
// satisfy this regression. The parameters stay boxed, as in the decoder's ST loop.
function primitiveModule() {
    const expr = syntax => new NeutralExpr(syntax);
    const local = (name, level) => expr(new Syntax.Local(new Just(name), level));
    const observe = (name, value) => expr(new Syntax.App(expr(new Syntax.Var(
        new Core.Qualified(new Just('IndexProbe'), name))), [value]));
    const annotation = (type, value) => expr(new Syntax.Typed(type, value));
    const bindings = [false, true].map(observed => {
        let source = local('source', 0), index = local('index', 1);
        if (observed) {
            source = observe('source', source);
            index = observe('index', index);
        }
        source = annotation(new Core.Array(Core.Int.value), annotation(new Core.Array(Core.Int.value), source));
        const body = annotation(Core.Int.value, expr(new Syntax.PrimOp(new Syntax.Op2(
            Syntax.OpArrayIndex.value, source, annotation(Core.Int.value, index),
        ))));
        return { recursive: false, bindings: [new Tuple(observed ? 'observed' : 'plain',
            expr(new Syntax.Abs([new Tuple(new Just('source'), 0), new Tuple(new Just('index'), 1)], body)))] };
    });
    return CodeGen.translate(metadata)({ name: 'UnsafePrimitive', bindings, comments: [], imports: emptySet,
        exports: emptySet, reExports: emptySet, dataTypes: emptyMap,
        dataDecls: [], classDecls: [], foreign: emptyMap, implementations: emptyMap, directives: emptyMap });
}

test('unsafe indexing converts one element, preserves evaluation and allocates constant bytes per access', t => {
    const primitiveCode = primitiveModule();
    assert.match(primitiveCode, /arrayUnsafe_value_/, 'OpArrayIndex must use the representation-preserving emitter');
    const fixtures = [];
    for (const element of [
        { name: 'Int', core: Core.Int.value, go: Go.TypeInt64.value, value: 'int64(i + 7)', check: 'result.IntVal != int64(index + 7)' },
        { name: 'String', core: Core.String.value, go: Go.TypeString.value, value: 'fmt.Sprint(i + 7)', check: 'result.StrVal() != fmt.Sprint(index + 7)' },
        { name: 'Record', core: recordType, go: nativeRecord, value: `${Go.goTypeToStr(nativeRecord)}{x: int64(i + 7)}`, check: 'gopurs_runtime.RecordGet(result, "x").IntVal != int64(index + 7)' },
    ]) {
        for (const mode of ['native', 'boxedSlice', 'boxed']) {
            fixtures.push({ ...element, mode, functionName: `${mode}${element.name}`,
                sourceType: mode === 'native' ? new Go.TypeNativeArray(element.go)
                    : mode === 'boxedSlice' ? new Go.TypeNativeArray(Go.TypeValue.value) : Go.TypeValue.value });
        }
    }
    const declarations = fixtures.map(fixture => {
        const context = {
            metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
            depth: 0, modNameStr: 'IndexProbe', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
            tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
            mbExpectedExprType: Nothing.value,
        };
        const translate = _child => nextId => expression => {
            assert.ok(expression.value1 instanceof Syntax.Var, 'array Typed wrappers must be moved onto the selected element');
            const name = expression.value1.value0.value1;
            assert.ok(name === 'source' || name === 'index');
            return {
                stmts: new StmtLeaf(Go.rawGo(`mark("${name}:statement")`)),
                expr: Go.rawGo(`observed("${name}:expression", ${name})`),
                exprType: name === 'source' ? fixture.sourceType : Go.TypeInt64.value, nextId,
            };
        };
        const array = typed(new Core.Array(fixture.core), typed(new Core.Array(fixture.core), variable('source')));
        const result = unsafeIndex(translate)(context)(0)(array)(variable('index'));
        const body = [...flattenStmts(result.stmts), new Go.GoReturn(result.expr)].map(printGoExpr).join('\n');
        return `func ${fixture.functionName}(source ${Go.goTypeToStr(fixture.sourceType)}, index int64) gopurs_runtime.Value {\n${body}\n}`;
    });
    const checks = fixtures.map(fixture => {
        const box = fixture.name === 'Int' ? 'gopurs_runtime.Int(native[i])'
            : fixture.name === 'String' ? 'gopurs_runtime.Str(native[i])'
                : 'gopurs_runtime.RecordDict1("x", gopurs_runtime.Int(native[i].x))';
        const source = fixture.mode === 'native' ? 'native' : fixture.mode === 'boxedSlice' ? 'boxedValues' : 'gopurs_runtime.Array(boxedValues)';
        return `t.Run("${fixture.functionName}", func(t *testing.T) {
            costs := make([]float64, 0, 2)
            for _, size := range []int{16, 16384} {
                native := make([]${Go.goTypeToStr(fixture.go)}, size)
                boxedValues := make([]gopurs_runtime.Value, size)
                for i := range native { native[i] = ${fixture.value}; boxedValues[i] = ${box} }
                beforeNative := append([]${Go.goTypeToStr(fixture.go)}(nil), native...)
                beforeBoxed := append([]gopurs_runtime.Value(nil), boxedValues...)
                source := ${source}
                traceEnabled = true
                for _, index := range []int{0, size / 2, size - 1} {
                    trace = nil
                    result := ${fixture.functionName}(source, int64(index))
                    if ${fixture.check} { t.Fatalf("wrong value at %d: %v", index, result) }
                    expected := []string{"source:statement", "source:expression", "index:statement", "index:expression"}
                    if !reflect.DeepEqual(trace, expected) { t.Fatalf("evaluation order/count: %v", trace) }
                }
                traceEnabled = false
                if !reflect.DeepEqual(native, beforeNative) || !reflect.DeepEqual(boxedValues, beforeBoxed) { t.Fatal("source mutated") }
                costs = append(costs, allocatedBytes(func() { indexSink = ${fixture.functionName}(source, int64(size / 2)) }))
                for i := range native {
                    if !reflect.DeepEqual(native[i], beforeNative[i]) || boxedValues[i] != beforeBoxed[i] { t.Fatal("source mutated during allocation measurement") }
                }
            }
            if costs[1] > costs[0] + 128 || costs[1] > 1024 { t.Fatalf("bytes per access grow with array length: %v", costs) }
            t.Logf("bytes per access, lengths 16/16384: %v", costs)
        })`;
    });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-array-unsafe-index-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    mkdirSync(join(directory, 'purescript'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/UnsafePrimitive.go'), primitiveCode);
    writeFileSync(join(directory, 'purescript/index_test.go'), `package purescript
import ("fmt"; "reflect"; "runtime"; "testing"; "gopurs/output/gopurs_runtime")
var traceEnabled bool
var trace []string
var indexSink gopurs_runtime.Value
func mark(name string) { if traceEnabled { trace = append(trace, name) } }
func observed[T any](name string, value T) T { mark(name); return value }
func Get_IndexProbe_source() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value { mark("source"); return value })
}
func Get_IndexProbe_index() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value { mark("index"); return value })
}
// Measure bytes, not merely allocation count: one copied slice is one allocation
// at every length, but its bytes grow linearly with the entire input array.
func allocatedBytes(f func()) float64 {
    f()
    var before, after runtime.MemStats
    runtime.ReadMemStats(&before)
    const repetitions = 2000
    for i := 0; i < repetitions; i++ { f() }
    runtime.ReadMemStats(&after)
    return float64(after.TotalAlloc - before.TotalAlloc) / repetitions
}
${declarations.join('\n')}
func TestUnsafeIndexMatrix(t *testing.T) {
    ${checks.join('\n')}
}
func TestActualPrimitive(t *testing.T) {
    plain := Get_UnsafePrimitive_plain()
    observer := Get_UnsafePrimitive_observed()
    costs := make([]float64, 0, 2)
    for _, size := range []int{16, 16384} {
        values := make([]gopurs_runtime.Value, size)
        for i := range values { values[i] = gopurs_runtime.Int(int64(i + 7)) }
        original := append([]gopurs_runtime.Value(nil), values...)
        source := gopurs_runtime.Array(values)
        index := gopurs_runtime.Int(int64(size - 1))
        traceEnabled = true
        trace = nil
        result := gopurs_runtime.Apply2(observer, source, index)
        if result.IntVal != int64(size + 6) || !reflect.DeepEqual(trace, []string{"source", "index"}) { t.Fatalf("primitive value/order: %v %v", result, trace) }
        traceEnabled = false
        costs = append(costs, allocatedBytes(func() { indexSink = gopurs_runtime.Apply2(plain, source, index) }))
        if !reflect.DeepEqual(values, original) { t.Fatal("primitive source mutated") }
    }
    if costs[1] > costs[0] + 128 || costs[1] > 1024 { t.Fatalf("primitive bytes per access grow with length: %v", costs) }
    t.Logf("primitive bytes per access, lengths 16/16384: %v", costs)
}
`);
    const result = spawnSync('go', ['test', '-v', '-count=1', './purescript'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /--- PASS: TestActualPrimitive/);
    for (const fixture of fixtures) assert.match(result.stdout, new RegExp(`--- PASS: TestUnsafeIndexMatrix/${fixture.functionName}`));
});
