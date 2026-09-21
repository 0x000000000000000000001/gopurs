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
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (name, level) => expr(new S.Local(new Just(name), level));
const integer = value => typed(C.Int.value, expr(new S.Lit(new C.LitInt(value))));
const app = (fn, ...args) => expr(new S.App(fn, args));
const lambda = (params, body) => expr(new S.Abs(
    params.map(([name, level]) => new Tuple(new Just(name), level)), body));
const letIn = (name, level, value, body) => expr(new S.Let(new Just(name), level, value, body));
const add = (left, right) => typed(C.Int.value,
    expr(new S.PrimOp(new S.Op2(new S.OpIntNum(S.OpAdd.value), left, right))));
const entry = new C.Record(new C.Row([
    new Tuple('count', C.Int.value), new Tuple('label', C.String.value),
], Nothing.value));
const record = count => typed(entry, expr(new S.Lit(new C.LitRecord([
    new C.Prop('count', count),
    new C.Prop('label', expr(new S.Lit(new C.LitString('kept')))),
]))));
const tick = value => typed(C.Int.value, app(
    expr(new S.Var(new C.Qualified(new Just('Probe'), 'tick'))), value));

function recordHelper(name, partial, use) {
    const params = partial ? [['first', 2], ['second', 3]] : [['offset', 2]];
    const offset = partial ? add(local('first', 2), local('second', 3)) : local('offset', 2);
    const helper = typed(new C.Func(params.map(() => C.Int.value), entry),
        lambda(params, record(add(local('seed', 0), offset))));
    // Construct the Let directly after PBO, so inlining cannot erase the worker.
    return lambda([['seed', 0]], letIn(name, 1, helper, use(local(name, 1))));
}

function generatedModule() {
    const staged = typed(new C.Func([C.Int.value, C.Int.value], entry),
        lambda([['first', 2]], letIn('observed', 3, tick(local('first', 2)),
            lambda([['second', 4]], letIn('finished', 5, tick(local('second', 4)),
                record(add(local('observed', 3), local('finished', 5))))))));
    return CodeGen.translate(metadata)({
        name: 'LocalReturns', bindings: [
            ['direct', recordHelper('makeDirect', false, helper => app(helper, integer(2)))],
            ['firstClass', recordHelper('makeBoxed', false, helper => helper)],
            ['partial', recordHelper('makePartial', true, helper => app(helper, integer(2)))],
            ['staged', lambda([['ignored', 0]], letIn('makeStaged', 1, staged, local('makeStaged', 1)))],
        ].map(([name, value]) => ({ recursive: false, bindings: [new Tuple(name, value)] })),
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
}

test('local workers return native records while first-class and partial calls retain wrappers', t => {
    const code = generatedModule();
    for (const [name, params] of [
        ['makeDirect', 'int64'], ['makeBoxed', 'int64'], ['makePartial', 'int64, int64'],
    ]) {
        assert.match(code, new RegExp(
            `var Call_local_LocalReturns_${name}_\\w+ func\\(${params}\\) struct\\{\\s*count int64\\s*label string\\s*\\}`),
        `${name} must keep its native record result`);
    }
    // The computation separating the lambdas must remain at arity one. Its
    // result is a closure, so this worker legitimately retains a Value return.
    assert.match(code, /var Call_local_LocalReturns_makeStaged_\w+ func\(int64\) gopurs_runtime\.Value/);

    const directory = mkdtempSync(join(tmpdir(), 'gopurs-local-native-returns-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/LocalReturns.go'), code);
    writeFileSync(join(directory, 'purescript/Probe.go'), `package purescript
import "gopurs/output/gopurs_runtime"
var Calls []int64
func Get_Probe_tick() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        Calls = append(Calls, value.IntVal)
        return value
    })
}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
func show(label string, result gopurs_runtime.Value) {
    fmt.Printf("%s %d:%s\\n", label, gopurs_runtime.RecordGet(result, "count").IntVal,
        gopurs_runtime.RecordGet(result, "label").StrVal())
}
func main() {
    seed := gopurs_runtime.Int(40)
    show("direct", gopurs_runtime.Apply(purescript.Get_LocalReturns_direct(), seed))
    maker := gopurs_runtime.Apply(purescript.Get_LocalReturns_firstClass(), seed)
    show("first-class", gopurs_runtime.Apply(maker, gopurs_runtime.Int(5)))
    show("first-class-again", gopurs_runtime.Apply(maker, gopurs_runtime.Int(-1)))
    partial := gopurs_runtime.Apply(purescript.Get_LocalReturns_partial(), seed)
    show("partial", gopurs_runtime.Apply(partial, gopurs_runtime.Int(3)))
    staged := gopurs_runtime.Apply(purescript.Get_LocalReturns_staged(), gopurs_runtime.Value{})
    fmt.Println("created", purescript.Calls)
    next := gopurs_runtime.Apply(staged, seed)
    fmt.Println("first-argument", purescript.Calls)
    show("completed", gopurs_runtime.Apply(next, gopurs_runtime.Int(2)))
    fmt.Println("second-argument", purescript.Calls)
    show("completed-again", gopurs_runtime.Apply(next, gopurs_runtime.Int(3)))
    fmt.Println("reused-closure", purescript.Calls)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout,
        'direct 42:kept\nfirst-class 45:kept\nfirst-class-again 39:kept\npartial 45:kept\n'
        + 'created []\nfirst-argument [40]\ncompleted 42:kept\nsecond-argument [40 2]\n'
        + 'completed-again 43:kept\nreused-closure [40 2 3]\n');
});
