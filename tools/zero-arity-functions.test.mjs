import { withReboxFields } from './codegen-metadata.mjs';
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
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap,
    classDeclsFields: emptyMap,
});
const expr = syntax => new NeutralExpr(syntax);
const local = (name, level) => expr(new S.Local(new Just(name), level));
const variable = name => expr(new S.Var(new C.Qualified(new Just('Probe'), name)));
const app = (fn, args) => expr(new S.App(fn, args));
const literal = value => expr(new S.Lit(new C.LitInt(value)));
const lambda = (name, body) => expr(new S.Abs([new Tuple(new Just(name), 0)], body));

for (const { effect, returnsLambda } of [
    { effect: false, returnsLambda: false },
    { effect: true, returnsLambda: false },
    { effect: false, returnsLambda: true },
]) {
    const label = returnsLambda ? 'lambda-returning' : effect ? 'effectful' : 'pure';
    test(`zero-argument ${label} functions defer their complete body and return closures intact`, t => {
        const Abs = effect ? S.UncurriedEffectAbs : S.UncurriedAbs;
        const App = effect ? S.UncurriedEffectApp : S.UncurriedApp;
        // Mirrors a wildcard case alternative: UncurriedAbs [] (Let ... (App ...)).
        const body = returnsLambda
            ? expr(new S.Abs([new Tuple(new Just('value'), 2)],
                app(variable('tick'), [local('value', 2)])))
            : expr(new S.Let(new Just('value'), 2,
            app(variable('tick'), [literal(42)]),
            app(variable('finish'), [local('value', 2)])));
        const deferred = expr(new S.Let(new Just('deferred'), 1,
            expr(new Abs([], body)), local('deferred', 1)));
        const code = CodeGen.translate(metadata)({
            name: 'Consumer', bindings: [
                ['createDeferred', lambda('ignored', deferred)],
                // Worker extraction must also stop at a zero-argument boundary.
                ['makeDirect', lambda('ignored', expr(new Abs([], body)))],
                ['topDeferred', expr(new Abs([], body))],
                ['call', lambda('fn', expr(new App(local('fn', 0), [])))],
            ].map(([name, value]) => ({
                recursive: false, bindings: [new Tuple(name, value)],
            })),
            comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
            dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
            implementations: emptyMap, directives: emptyMap,
        });
        const directory = mkdtempSync(join(tmpdir(), `gopurs-zero-arity-${label}-`));
        t.after(() => rmSync(directory, { recursive: true, force: true }));
        mkdirSync(join(directory, 'purescript'));
        mkdirSync(join(directory, 'gopurs_runtime'));
        writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
        writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
        writeFileSync(join(directory, 'purescript/Consumer.go'), code);
        writeFileSync(join(directory, 'purescript/Probe.go'), `package purescript
import "gopurs/output/gopurs_runtime"
var TickCalls, FinishCalls, EffectCalls, InnerCalls int
func Get_Probe_tick() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        TickCalls++
        return value
    })
}


func Get_Probe_finish() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        FinishCalls++
        result := gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
            InnerCalls++
            return value
        })
        ${effect ? `return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
            EffectCalls++
            return result
        })` : 'return result'}
    })
}
`);
        writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
func counters() {
    fmt.Printf("%d %d %d %d\\n", purescript.TickCalls, purescript.FinishCalls,
        purescript.EffectCalls, purescript.InnerCalls)
}
func main() {
    deferred := purescript.Call_Consumer_createDeferred(gopurs_runtime.Value{})
    _ = purescript.Call_Consumer_createDeferred(gopurs_runtime.Value{}) // unused function
    counters()
    first := purescript.Call_Consumer_call(deferred)
    second := purescript.Call_Consumer_call(deferred)
    ${effect ? `counters() // Creating Effect actions also leaves their bodies deferred.
    first = gopurs_runtime.Apply(first, gopurs_runtime.Value{})
    second = gopurs_runtime.Apply(second, gopurs_runtime.Value{})` : ''}
    counters()
    fmt.Printf("%d %d\\n", gopurs_runtime.Apply(first, gopurs_runtime.Int(42)).IntVal,
        gopurs_runtime.Apply(second, gopurs_runtime.Int(42)).IntVal)
    counters()
    direct := purescript.Call_Consumer_makeDirect(gopurs_runtime.Value{})
    counters()
    result := purescript.Call_Consumer_call(direct)
    ${effect ? 'result = gopurs_runtime.Apply(result, gopurs_runtime.Value{})' : ''}
    fmt.Printf("direct %d\\n", gopurs_runtime.Apply(result, gopurs_runtime.Int(42)).IntVal)
    counters()
    top := purescript.Get_Consumer_topDeferred()
    _ = purescript.Get_Consumer_topDeferred()
    counters()
    result = purescript.Call_Consumer_call(top)
    ${effect ? 'result = gopurs_runtime.Apply(result, gopurs_runtime.Value{})' : ''}
    fmt.Printf("top %d\\n", gopurs_runtime.Apply(result, gopurs_runtime.Int(42)).IntVal)
    counters()
}
`);
        const result = spawnSync('go', ['run', '.'], {
            cwd: directory, encoding: 'utf8', timeout: 30_000,
            env: { ...process.env, GOWORK: 'off' },
        });
        assert.ifError(result.error);
        assert.equal(result.status, 0, result.stdout + result.stderr);
        const effectCalls = effect ? 2 : 0;
        const expected = returnsLambda
            ? '0 0 0 0\n0 0 0 0\n42 42\n2 0 0 0\n2 0 0 0\ndirect 42\n3 0 0 0\n3 0 0 0\ntop 42\n4 0 0 0\n'
            : `0 0 0 0\n${effect ? '0 0 0 0\n' : ''}2 2 ${effectCalls} 0\n42 42\n2 2 ${effectCalls} 2\n`
                + `2 2 ${effectCalls} 2\ndirect 42\n3 3 ${effect ? 3 : 0} 3\n`
                + `3 3 ${effect ? 3 : 0} 3\ntop 42\n4 4 ${effect ? 4 : 0} 4\n`;
        assert.equal(result.stdout, expected);
    });
}

test('a recursive zero-argument local function starts its loop only when called', t => {
    const self = local('loop', 1);
    const condition = expr(new S.PrimOp(new S.Op2(new S.OpIntOrd(S.OpLt.value),
        app(variable('tick'), [literal(0)]), literal(3))));
    const body = expr(new S.Branch([
        new S.Pair(condition, expr(new S.UncurriedApp(self, []))),
    ], literal(42)));
    const recursive = expr(new S.LetRec(1, [
        new Tuple('loop', expr(new S.UncurriedAbs([], body))),
    ], self));
    const code = CodeGen.translate(metadata)({
        name: 'Recursive', bindings: [
            ['createDeferred', lambda('ignored', recursive)],
            ['call', lambda('fn', expr(new S.UncurriedApp(local('fn', 0), [])))],
        ].map(([name, value]) => ({ recursive: false, bindings: [new Tuple(name, value)] })),
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-zero-arity-recursive-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/Recursive.go'), code);
    writeFileSync(join(directory, 'purescript/Probe.go'), `package purescript
import "gopurs/output/gopurs_runtime"
var Calls int64
func Get_Probe_tick() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
        Calls++
        return gopurs_runtime.Int(Calls)
    })
}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
func main() {
    deferred := purescript.Call_Recursive_createDeferred(gopurs_runtime.Value{})
    fmt.Printf("created %d\\n", purescript.Calls)
    first := purescript.Call_Recursive_call(deferred)
    fmt.Printf("first %d calls %d\\n", first.IntVal, purescript.Calls)
    second := purescript.Call_Recursive_call(deferred)
    fmt.Printf("second %d calls %d\\n", second.IntVal, purescript.Calls)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout, 'created 0\nfirst 42 calls 3\nsecond 42 calls 4\n');
});
