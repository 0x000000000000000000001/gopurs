import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import * as Maps from '../output/Data.Map/index.js';
import * as Maybe from '../output/Data.Maybe/index.js';
import * as Ord from '../output/Data.Ord/index.js';
import * as Sets from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import * as Conversions from '../output/Gopurs.GoConversions/index.js';
import * as Functions from '../output/Gopurs.FunctionExprs/index.js';
import * as Calls from '../output/Gopurs.CallExprs/index.js';
import * as Records from '../output/Gopurs.RecordExprs/index.js';
import * as Statements from '../output/Gopurs.ExprContext/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import * as Tco from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as Core from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as Syntax from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const metadata = withReboxFields({
    pointerAdtPaths: Maps.empty, pointerAdtNodes: Sets.empty, pointerAdtLeaves: Maps.empty,
    enumAdts: Sets.empty, enumCtors: Sets.empty, elidedCtors: Sets.empty, ctorTypes: Maps.empty,
    classDeclsFields: Maps.empty, globalTypes: Maps.empty, globalFunctions: Maps.empty,
});
const context = () => ({
    metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: Sets.empty })(),
    depth: 0, modNameStr: 'Probe', recVars: [], moduleFunctions: Maps.empty, bound: Maps.empty,
    tcoIdent: Maybe.Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Maybe.Nothing.value, mbExpectedGoType: Maybe.Nothing.value,
});
const record = new Core.Record(new Core.Row([new Tuple('value', Core.Int.value)], Maybe.Nothing.value));
const nativeRecord = new Go.TypeRecord([new Tuple('value', Go.TypeInt64.value)]);
const variable = name => new Tco.TcoExpr(null, new Syntax.Var(new Core.Qualified(Maybe.Nothing.value, name)));
const typed = (type, expression) => new Tco.TcoExpr(null, new Syntax.Typed(type, expression));

function callback(arity, firstType = record) {
    const ctx = context();
    const params = Array.from({ length: arity }, (_, index) => new Tuple(new Maybe.Just(`arg${index}`), index));
    const body = variable('body');
    const expression = typed(new Core.Func([firstType, ...Array(arity - 1).fill(Core.Int.value)], Core.Int.value),
        new Tco.TcoExpr(null, new Syntax.Abs(params, body)));
    const seen = [];
    const translate = child => nextId => actualBody => {
        assert.equal(actualBody, body);
        const bindings = params.map((_, i) => Maps.lookup(Ord.ordString)(`arg${i}_${i}`)(child.bound).value0);
        seen.push(bindings.map(binding => binding.goType));
        assert.deepEqual(child.mbExpectedGoType, new Maybe.Just(Go.TypeValue.value));
        const native = bindings[0].goType instanceof Go.TypeRecord;
        let result = Records.getProp(metadata)(ctx.codegenStateRef)('Probe')('value')({
            expr: new Go.GoVar(bindings[0].name), exprType: bindings[0].goType,
        });
        let expr = Conversions.coerceGoExpr(ctx.codegenStateRef)('Probe')(result.expr)(result.exprType)(Go.TypeInt64.value);
        for (const binding of bindings.slice(1)) expr = new Go.GoBinOp('+', expr,
            Conversions.coerceGoExpr(ctx.codegenStateRef)('Probe')(new Go.GoVar(binding.name))(binding.goType)(Go.TypeInt64.value));
        expr = new Go.GoBinOp('+', expr, new Go.GoVar('captured'));
        return { stmts: new Statements.StmtLeaf(Go.rawGo(`entered("${native ? 'native' : 'fallback'}")`)),
            expr, exprType: Go.TypeInt64.value, nextId: nextId + 1 };
    };
    const result = Functions.abstraction(translate)(ctx)(0)(expression)(params)(body);
    return { source: printGoExpr(result.expr), seen, result };
}

function call(arity, { name = `call${arity}`, scalarType = Go.TypeInt64.value, literal = false } = {}) {
    const ctx = context();
    const fn = variable('callback');
    const args = Array.from({ length: arity }, (_, index) => typed(index === 0 ? record : Core.Int.value, variable(`arg${index}`)));
    const expressions = new Map([[fn, { name: 'fn', exprType: Go.TypeValue.value }],
        ...args.map((arg, index) => [arg, { name: `arg${index}`, exprType: index === 0 ? nativeRecord : scalarType }])]);
    const translated = [];
    const translate = _child => nextId => expression => {
        const item = expressions.get(expression);
        assert.ok(item);
        translated.push(item.name);
        return { stmts: new Statements.StmtLeaf(Go.rawGo(`mark("${item.name}:statement")`)),
            expr: Go.rawGo(`markExpression("${item.name}:expression", ${literal && item.name.startsWith('arg') && item.name !== 'arg0' ? '5' : item.name})`), exprType: item.exprType, nextId };
    };
    const result = Calls.application(translate)(ctx)(0)(new Tco.TcoExpr(null, new Syntax.App(fn, args)));
    assert.deepEqual(translated, ['fn', ...args.map((_, i) => `arg${i}`)], 'code generation visits each operand once');
    const body = [...Statements.flattenStmts(result.stmts), new Go.GoReturn(result.expr)].map(printGoExpr).join('\n');
    const params = ['fn gopurs_runtime.Value', `arg0 ${Go.goTypeToStr(nativeRecord)}`,
        ...args.slice(1).map((_, i) => `arg${i + 1} ${Go.goTypeToStr(scalarType)}`)];
    return `func ${name}(${params.join(', ')}) gopurs_runtime.Value {\n${body}\n}`;
}

test('native record callbacks have an exact Value-returning entry point within the complete runtime arity', () => {
    for (const arity of [1, 2, 5]) {
        const result = callback(arity);
        assert.match(result.source, /WithFunctionData/);
        assert.match(result.source, /arg0_0\.value/);
        assert.match(result.source, /RecordGet/);
        assert.equal(result.seen.length, 2);
        assert.ok(result.seen[0].every(type => type instanceof Go.TypeValue));
        assert.deepEqual(result.seen[1], [nativeRecord, ...Array(arity - 1).fill(Go.TypeInt64.value)]);
        assert.equal(result.result.nextId, 2);
    }
    for (const [arity, firstType] of [[6, record], [1, Core.Any.value],
        [1, new Core.Record(new Core.Row([new Tuple('value', Core.Int.value)], new Maybe.Just(Core.Any.value)))]]) {
        const result = callback(arity, firstType);
        assert.doesNotMatch(result.source, /WithFunctionData/);
        assert.equal(result.seen.length, 1);
    }
});

test('generated dispatch preserves captures, fallback, partial calls and evaluation order while avoiding record allocations', t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-native-record-callback-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    const factories = [1, 2, 5].map(arity => `func callback${arity}(captured int64) gopurs_runtime.Value { return ${callback(arity).source} }`);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "slices"; "testing"; "gopurs/output/gopurs_runtime")
var trace []string
var tracing bool
var fastCalls, fallbackCalls int
var sink gopurs_runtime.Value
func mark(label string) { if tracing { trace = append(trace, label) } }
func markExpression[T any](label string, value T) T { mark(label); return value }
func entered(kind string) { mark(kind); if kind == "native" { fastCalls++ } else { fallbackCalls++ } }
${factories.join('\n')}
${[1, 2, 5].map(arity => call(arity)).join('\n')}
${[['Int',Go.TypeInt64.value],['Float',Go.TypeFloat64.value],['Enum',Go.TypeUint32.value]].map(([kind,scalarType]) => call(2,{name:`callLiteral${kind}`,scalarType,literal:true})).join('\n')}
func check(got gopurs_runtime.Value, want int64) { if got.IntVal != want { panic(fmt.Sprintf("got %v want %v",got.IntVal,want)) } }
func main() {
    input := ${Go.goTypeToStr(nativeRecord)}{value: 7}
    unary, binary, five := callback1(30), callback2(30), callback5(30)
    tracing = true
    check(call1(unary, input),37)
    if !slices.Equal(trace, []string{"fn:statement","fn:expression","arg0:statement","arg0:expression","native"}) { panic(fmt.Sprint(trace)) }
    trace = nil
    check(call2(binary, input, 5),42)
    if !slices.Equal(trace, []string{"fn:statement","fn:expression","arg0:statement","arg0:expression","arg1:statement","arg1:expression","native"}) { panic(fmt.Sprint(trace)) }
    check(call5(five,input,1,2,3,4),47)
    tracing = false
    check(gopurs_runtime.Apply(unary,gopurs_runtime.RecordDict1("value",gopurs_runtime.Int(7))),37)
    ordinary := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value { entered("fallback"); return gopurs_runtime.Int(gopurs_runtime.RecordGet(value,"value").IntVal+30) })
    wrong := gopurs_runtime.WithFunctionData(ordinary,func(_ int64) gopurs_runtime.Value { panic("wrong signature") })
    check(call1(ordinary,input),37)
    check(call1(wrong,input),37)
    partial := call1(binary,input)
    if _,ok := gopurs_runtime.FunctionData[func(int64) gopurs_runtime.Value](partial); ok { panic("partial kept full-call metadata") }
    check(gopurs_runtime.Apply(partial,gopurs_runtime.Int(5)),42)
    if fastCalls != 3 || fallbackCalls != 4 { panic(fmt.Sprintf("fast=%d fallback=%d",fastCalls,fallbackCalls)) }
    ${[['Int','int64'],['Float','float64'],['Enum','uint32']].map(([kind,scalar]) => `{
        fn := gopurs_runtime.WithFunctionData(gopurs_runtime.Func2(func(a,b gopurs_runtime.Value) gopurs_runtime.Value { panic("literal fallback") }),
            func(record ${Go.goTypeToStr(nativeRecord)}, scalar ${scalar}) gopurs_runtime.Value { return gopurs_runtime.Int(record.value + int64(scalar)) })
        check(callLiteral${kind}(fn,input,0),12)
    }`).join('\n')}
    fast := testing.AllocsPerRun(1000,func(){ sink = call1(unary,input) })
    slow := testing.AllocsPerRun(1000,func(){ sink = call1(ordinary,input) })
    if fast != 0 || slow <= fast { panic(fmt.Sprintf("allocations fast=%g fallback=%g",fast,slow)) }
    fmt.Printf("fast=%g fallback=%g\\n",fast,slow)
}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 60000,
        env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(result.error);
    assert.equal(result.status, 0, `${result.stdout}\n${result.stderr}`);
    assert.match(result.stdout, /^fast=0 fallback=[1-9]/);
});

test('a callback with a local recursive worker retains each factory capture in both entry points', async t => {
    const [{ NeutralExpr }, CodeGen] = await Promise.all([
        import('../output/PureScript.Backend.Optimizer.Semantics/index.js'),
        import('../output/Gopurs.CodeGen/index.js'),
    ]);
    const expr = syntax => new NeutralExpr(syntax);
    const ann = (type, value) => expr(new Syntax.Typed(type, value));
    const local = (name, level, type = Core.Int.value) => ann(type, expr(new Syntax.Local(new Maybe.Just(name), level)));
    const integer = value => ann(Core.Int.value, expr(new Syntax.Lit(new Core.LitInt(value))));
    const binary = (type, op, left, right) => ann(type, expr(new Syntax.PrimOp(new Syntax.Op2(op, left, right))));
    const plus = (left, right) => binary(Core.Int.value, new Syntax.OpIntNum(Syntax.OpAdd.value), left, right);
    const loopType = new Core.Func([Core.Int.value], Core.Int.value);
    const invoke = value => ann(Core.Int.value, expr(new Syntax.App(local('go', 3, loopType), [value])));
    const n = local('n', 4);
    const value = ann(Core.Int.value, expr(new Syntax.Accessor(local('token', 2, record), new Syntax.GetProp('value'))));
    const loopBody = expr(new Syntax.Branch([new Syntax.Pair(
        binary(Core.Boolean.value, new Syntax.OpIntOrd(Syntax.OpEq.value), n, integer(0)),
        plus(local('saved', 1), value))],
        invoke(binary(Core.Int.value, new Syntax.OpIntNum(Syntax.OpSubtract.value), n, integer(1)))));
    const loop = ann(loopType, expr(new Syntax.Abs([new Tuple(new Maybe.Just('n'), 4)], loopBody)));
    const callbackType = new Core.Func([record], Core.Int.value);
    const callbackBody = expr(new Syntax.LetRec(3, [new Tuple('go', loop)], invoke(integer(4))));
    const callback = ann(callbackType, expr(new Syntax.Abs([new Tuple(new Maybe.Just('token'), 2)], callbackBody)));
    // The computation between the two lambdas is a real currying boundary.
    const body = expr(new Syntax.Let(new Maybe.Just('saved'), 1, plus(local('base', 0), integer(1)), callback));
    const factory = ann(new Core.Func([Core.Int.value], callbackType),
        expr(new Syntax.Abs([new Tuple(new Maybe.Just('base'), 0)], body)));
    const code = CodeGen.translate(metadata)({
        name: 'CallbackLetRec', bindings: [{ recursive: false, bindings: [new Tuple('create', factory)] }],
        comments: [], imports: Sets.empty, exports: Sets.empty, reExports: Sets.empty,
        dataTypes: Maps.empty, dataDecls: [], classDecls: [], foreign: Maps.empty,
        implementations: Maps.empty, directives: Maps.empty,
    });
    assert.match(code, /WithFunctionData/);
    assert.match(code, /Call_local_CallbackLetRec_go/);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-native-callback-letrec-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    mkdirSync(join(directory, 'purescript'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/CallbackLetRec.go'), code);
    writeFileSync(join(directory, 'purescript/Probe.go'), `package purescript
import "gopurs/output/gopurs_runtime"
func NativeInvoke(fn gopurs_runtime.Value, value int64) int64 {
    native,ok := gopurs_runtime.FunctionData[func(struct{value int64}) gopurs_runtime.Value](fn)
    if !ok { panic("missing native entry") }
    return native(struct{value int64}{value:value}).IntVal
}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "gopurs/output/gopurs_runtime"; "gopurs/output/purescript")
func main() {
    first,second := gopurs_runtime.Apply(purescript.Get_CallbackLetRec_create(),gopurs_runtime.Int(30)),gopurs_runtime.Apply(purescript.Get_CallbackLetRec_create(),gopurs_runtime.Int(80))
    for i:=0;i<5;i++ {
        for _,item := range []struct{callback gopurs_runtime.Value; want int64}{{first,34},{second,84}} {
            native := purescript.NativeInvoke(item.callback,3)
            boxed := gopurs_runtime.Apply(item.callback,gopurs_runtime.RecordDict1("value",gopurs_runtime.Int(3))).IntVal
            if native != item.want || boxed != item.want { panic(fmt.Sprintf("native=%d boxed=%d want=%d",native,boxed,item.want)) }
        }
    }
    fmt.Println("ok")
}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 60000,
        env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(result.error);
    assert.equal(result.status, 0, `${result.stdout}\n${result.stderr}`);
    assert.equal(result.stdout, 'ok\n');
});
