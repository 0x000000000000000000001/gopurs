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
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as Control from '../output/Gopurs.ControlExprs/index.js';
import { StmtEmpty, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap,
    classDeclsFields: emptyMap,
});

test('branch joins adapt mixed producers only when the consumer fixes a representation', () => {
    const context = {
        metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
        depth: 0, modNameStr: 'Branches', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
        tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
        mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
    };
    const node = syntax => new TcoExpr(null, syntax);
    const condition = node(new S.Lit(new C.LitBoolean(true)));
    const native = node(new S.Lit(new C.LitInt(42)));
    const boxed = node(new S.Lit(new C.LitInt(7)));
    const translate = child => nextId => value => {
        if (value === condition) {
            assert.deepEqual(child.mbExpectedGoType, Nothing.value);
            assert.equal(child.options.isTail, false);
        }
        return {
            stmts: StmtEmpty.value,
            expr: value === condition ? Go.rawGo('true')
                : value === native ? new Go.GoInt(42) : Go.rawGo('boxed'),
            exprType: value === condition ? Go.TypeBool.value
                : value === native ? Go.TypeInt64.value : Go.TypeValue.value,
            nextId,
        };
    };
    const emit = expected => Control.branch(translate)({ ...context, mbExpectedGoType: expected })(0)
        ([new S.Pair(condition, native)])(boxed);
    const fixed = emit(new Just(Go.TypeInt64.value));
    assert.deepEqual(fixed.exprType, Go.TypeInt64.value);
    const fixedCode = flattenStmts(fixed.stmts).map(printGoExpr).join('\n');
    assert.match(fixedCode, /var __t0 int64/);
    assert.match(fixedCode, /__t0 = 42/);
    assert.match(fixedCode, /__t0 = boxed.IntVal/);
    assert.doesNotMatch(fixedCode, /gopurs_runtime.Int\(/);
    for (const expected of [Nothing.value, new Just(Go.TypeValue.value)]) {
        assert.deepEqual(emit(expected).exprType, Go.TypeValue.value);
    }
    const failure = Control.failure({ ...context, mbExpectedGoType: new Just(Go.TypeInt64.value) })
        (0)(node(new S.Fail('unreachable')))('unreachable');
    assert.deepEqual(failure.exprType, Go.TypeInt64.value);
    assert.match(printGoExpr(failure.expr), /^func\(\) int64/);
});

const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const integer = value => typed(C.Int.value, expr(new S.Lit(new C.LitInt(value))));
const local = (name, level) => expr(new S.Local(new Just(name), level));
const call = (name, value, type) => typed(type, expr(new S.App(
    expr(new S.Var(new C.Qualified(new Just('Probe'), name))), [value])));
const entry = new C.Record(new C.Row([
    new Tuple('count', C.Int.value), new Tuple('label', C.String.value),
], Nothing.value));
const record = value => typed(entry, expr(new S.Lit(new C.LitRecord([
    new C.Prop('count', call('tick', integer(value), C.Int.value)),
    new C.Prop('label', expr(new S.Lit(new C.LitString('kept')))),
]))));
const branch = (yes, no) => expr(new S.Branch([new S.Pair(
    call('condition', local('choose', 0), C.Boolean.value), yes)], no));
const lambda = body => expr(new S.Abs([new Tuple(new Just('choose'), 0)], body));
const letIn = (name, level, value, body) => expr(new S.Let(new Just(name), level, value, body));

function generate(name, body, result = entry) {
    return CodeGen.translate(metadata)({
        name, bindings: [{ recursive: false, bindings: [new Tuple('run',
            typed(new C.Func([C.Boolean.value], result), lambda(body)))] }],
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
}

test('native results survive nested branches and identity lets without changing evaluation', t => {
    const mixed = () => branch(record(11), call('boxed', integer(22), entry));
    const codes = {
        Direct: generate('Direct', mixed()),
        Identity: generate('Identity', letIn('answer', 1, mixed(), local('answer', 1))),
        Nested: generate('Nested', branch(mixed(), call('boxed', integer(33), entry))),
        // The outer native result must not leak into an unrelated binding.
        Unrelated: generate('Unrelated', letIn('ignored', 1,
            branch(integer(1), call('tick', integer(2), C.Int.value)), record(44))),
        Generic: generate('Generic', mixed(), C.Any.value),
        Failing: generate('Failing', branch(record(55), expr(new S.Fail('selected failure')))),
    };
    for (const name of ['Direct', 'Identity', 'Nested', 'Failing']) {
        assert.match(codes[name], /var __t\d+ struct\{/,
            `${name} must join in the worker's native record layout`);
        assert.doesNotMatch(codes[name], /var __t\d+ gopurs_runtime\.Value/);
    }
    assert.match(codes.Unrelated, /var __t\d+ int64/);
    assert.match(codes.Generic, /var __t\d+ gopurs_runtime\.Value/);

    const directory = mkdtempSync(join(tmpdir(), 'gopurs-native-branch-results-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    for (const [name, code] of Object.entries(codes)) {
        writeFileSync(join(directory, 'purescript', `${name}.go`), code);
    }
    writeFileSync(join(directory, 'purescript/Probe.go'), `package purescript
import "gopurs/output/gopurs_runtime"
var Calls []int64
func Get_Probe_condition() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        Calls = append(Calls, 100 + value.IntVal)
        return value
    })
}
func Get_Probe_tick() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        Calls = append(Calls, value.IntVal)
        return value
    })
}
func Get_Probe_boxed() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        Calls = append(Calls, value.IntVal)
        return gopurs_runtime.Record(map[string]gopurs_runtime.Value{
            "count": value, "label": gopurs_runtime.Str("kept"),
        })
    })
}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    r "gopurs/output/gopurs_runtime"
    p "gopurs/output/purescript"
)
func show(name string, fn r.Value, choose bool) {
    p.Calls = nil
    value := r.Apply(fn, r.Bool(choose))
    fmt.Println(name, r.RecordGet(value, "count").IntVal, r.RecordGet(value, "label").StrVal(), p.Calls)
}
func fail() {
    defer func() { fmt.Println("failure", recover(), p.Calls) }()
    p.Calls = nil
    r.Apply(p.Get_Failing_run(), r.Bool(false))
}
func main() {
    for _, choose := range []bool{true, false} {
        show("direct", p.Get_Direct_run(), choose)
        show("identity", p.Get_Identity_run(), choose)
        show("nested", p.Get_Nested_run(), choose)
        show("unrelated", p.Get_Unrelated_run(), choose)
        show("generic", p.Get_Generic_run(), choose)
    }
    show("success", p.Get_Failing_run(), true)
    fail()
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000, env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout,
        'direct 11 kept [101 11]\nidentity 11 kept [101 11]\nnested 11 kept [101 101 11]\n'
        + 'unrelated 44 kept [101 44]\ngeneric 11 kept [101 11]\n'
        + 'direct 22 kept [100 22]\nidentity 22 kept [100 22]\nnested 33 kept [100 33]\n'
        + 'unrelated 44 kept [100 2 44]\ngeneric 22 kept [100 22]\n'
        + 'success 55 kept [101 55]\nfailure selected failure [100]\n');
});
