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
import * as Workers from '../output/Gopurs.LocalWorkers/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import * as T from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { localId } from '../output/PureScript.Backend.Optimizer.FreeVars/index.js';
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
const binary = (op, left, right) => typed(C.Int.value,
    expr(new S.PrimOp(new S.Op2(new S.OpIntNum(op), left, right))));

test('wrapper admission counts saturated calls but preserves escaping, partial and opaque uses', () => {
    // Direct TCO fixtures keep the test independent of optimizer inlining.
    const e = syntax => new T.TcoExpr(null, syntax);
    const target = e(new S.Local(new Just('renamedHelper'), 7));
    const arg = e(new S.Lit(new C.LitInt(3)));
    const call = (fn, ...args) => e(new S.App(fn, args));
    const required = value => Workers.needsWrapper(localId(new Just('renamedHelper'))(7))(2)(value);
    assert.equal(required(call(target, arg, arg)), false);
    assert.equal(required(call(call(target, arg), arg)), false);
    assert.equal(required(call(e(new S.Typed(new C.Func([C.Int.value, C.Int.value], C.Int.value), target)), arg, arg)), false);
    assert.equal(required(call(e(new S.TypeApp(target, C.Int.value)), arg, arg)), false);
    assert.equal(required(call(target, arg, arg, arg)), false);
    assert.equal(required(arg), false);
    assert.equal(required(e(new S.Local(new Just('renamedHelper'), 8))), false);
    assert.equal(required(e(new S.Local(new Just('anotherHelper'), 7))), false);
    for (const value of [
        target,
        call(target, arg),
        call(target, arg, target),
        call(e(new S.Var(new C.Qualified(new Just('External'), 'consume'))), target),
        e(new S.Abs([new Tuple(new Just('ignored'), 9)], target)),
        e(new S.Let(new Just('alias'), 9, target, call(e(new S.Local(new Just('alias'), 9)), arg, arg))),
        e(new S.UncurriedApp(target, [arg, arg])),
        call(e(new S.UncurriedApp(target, [arg])), arg),
        call(e(new S.TypeApp(e(new S.UncurriedApp(target, [arg])), C.Int.value)), arg),
    ]) assert.equal(required(value), true);
});

function helper(name) {
    const count = local('remaining', 2), accumulator = local('accumulator', 3);
    const done = typed(C.Boolean.value, expr(new S.PrimOp(new S.Op2(
        new S.OpIntOrd(S.OpEq.value), count, integer(0)))));
    const step = typed(C.Int.value, app(local(name, 1),
        binary(S.OpSubtract.value, count, integer(1)),
        binary(S.OpAdd.value, accumulator, integer(7))));
    return typed(new C.Func([C.Int.value, C.Int.value], C.Int.value),
        lambda([['remaining', 2], ['accumulator', 3]],
            typed(C.Int.value, expr(new S.Branch([new S.Pair(done, accumulator)], step)))));
}

function fixture(name, use, peer = false) {
    const bindings = [new Tuple(name, helper(name))];
    if (peer) bindings.push(new Tuple('returnPeer', typed(new C.Func([C.Int.value],
        new C.Func([C.Int.value, C.Int.value], C.Int.value)),
    lambda([['ignored', 4]], local(name, 1)))));
    return lambda([['input', 0]], expr(new S.LetRec(1, bindings, use(local(name, 1)))));
}

function generatedModule() {
    const direct = name => fixture(name, fn => typed(C.Int.value, app(fn, local('input', 0), integer(3))));
    return CodeGen.translate(metadata)({
        name: 'WorkerWrappers', bindings: [
            ['direct', direct('countSteps')],
            ['renamed', direct('countElsewhere')],
            ['partial', fixture('partialSteps', fn => app(fn, local('input', 0)))],
            ['returned', fixture('returnedSteps', fn => fn)],
            ['peer', fixture('peerSteps', () => app(local('returnPeer', 1), integer(0)), true)],
            ['alias', fixture('aliasSteps', fn => letIn('saved', 4, fn,
                typed(C.Int.value, app(local('saved', 4), local('input', 0), integer(3)))))],
        ].map(([name, value]) => ({ recursive: false, bindings: [new Tuple(name, value)] })),
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
}

test('unused curried wrappers do not make native recursive workers escape', t => {
    const code = generatedModule();
    for (const name of ['countSteps', 'countElsewhere']) {
        assert.match(code, new RegExp(`(var Call_local_WorkerWrappers_${name}_\\w+ func\\(int64, int64\\) int64|func Call_WorkerWrappers___gopurs_hot_local_\\d+\\()`));
        assert.doesNotMatch(code, new RegExp(`\\n${name}_\\w+ =`), `${name} needs no boxed wrapper`);
    }
    for (const name of ['partialSteps', 'returnedSteps', 'peerSteps', 'aliasSteps']) {
        assert.match(code, new RegExp(`\\n${name}_\\w+ =`), `${name} must preserve its first-class wrapper`);
    }
    const directory = process.env.GOPURS_LOCAL_WORKER_TEST_OUTPUT ?? mkdtempSync(join(tmpdir(), 'gopurs-local-worker-wrappers-'));
    if (!process.env.GOPURS_LOCAL_WORKER_TEST_OUTPUT) t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'), { recursive: true });
    mkdirSync(join(directory, 'gopurs_runtime'), { recursive: true });
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/WorkerWrappers.go'), code);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "testing"
    r "gopurs/output/gopurs_runtime"
    p "gopurs/output/purescript"
)
var sink r.Value
func expect(label string, actual r.Value, expected int64) {
    if actual.IntVal != expected { panic(fmt.Sprintf("%s: %d != %d", label, actual.IntVal, expected)) }
}
func main() {
    direct, renamed := p.Get_WorkerWrappers_direct(), p.Get_WorkerWrappers_renamed()
    for _, count := range []int64{0, 1, 4, 1000} {
        expect("direct", r.Apply(direct, r.Int(count)), 7*count+3)
        expect("renamed", r.Apply(renamed, r.Int(count)), 7*count+3)
        partial := r.Apply(p.Get_WorkerWrappers_partial(), r.Int(count))
        expect("partial", r.Apply(partial, r.Int(5)), 7*count+5)
        expect("partial reused", r.Apply(partial, r.Int(9)), 7*count+9)
        returned := r.Apply(p.Get_WorkerWrappers_returned(), r.Int(count))
        expect("returned", r.Apply2(returned, r.Int(count), r.Int(6)), 7*count+6)
        peer := r.Apply(p.Get_WorkerWrappers_peer(), r.Int(count))
        expect("peer", r.Apply2(peer, r.Int(count), r.Int(8)), 7*count+8)
        expect("alias", r.Apply(p.Get_WorkerWrappers_alias(), r.Int(count)), 7*count+3)
    }
    for _, action := range []r.Value{direct, renamed} {
        allocations := testing.AllocsPerRun(1000, func() { sink = r.Apply(action, r.Int(100)) })
        if allocations != 0 { panic(fmt.Sprintf("direct native worker allocates: %g", allocations)) }
    }
    fmt.Println("values, reusable wrappers and zero-allocation native workers verified")
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout, 'values, reusable wrappers and zero-allocation native workers verified\n');
});
