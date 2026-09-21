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
const int = C.Int.value;
const unary = new C.Func([int], int);
const binary = new C.Func([int, int], int);
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (name, level) => typed(int, expr(new S.Local(new Just(name), level)));
const variable = (module, name, type) => typed(type,
    expr(new S.Var(new C.Qualified(new Just(module), name))));
const literal = value => typed(int, expr(new S.Lit(new C.LitInt(value))));
const app = (fn, args, type) => typed(type, expr(new S.App(fn, args)));
const lambda = (params, body, type) => typed(type,
    expr(new S.Abs(params.map((name, level) => new Tuple(new Just(name), level)), body)));
const moduleOf = (name, bindings) => ({
    name, bindings: bindings.map(([name, value]) => ({
        recursive: false, bindings: [new Tuple(name, value)],
    })),
    comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
    dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
    implementations: emptyMap, directives: emptyMap,
});

test('imported calls use emitted workers and their actual signatures', t => {
    const add = variable('Producer', 'add', binary);
    const producer = moduleOf('Producer', [
        ['add', lambda(['x', 'y'], typed(int, expr(new S.PrimOp(
            new S.Op2(new S.OpIntNum(S.OpAdd.value), local('x', 0), local('y', 1)),
        ))), binary)],
        // A function-valued expression has a getter, but no direct worker.
        ['alias', app(add, [literal(10)], unary)],
        // Only one lambda is emitted; the second argument belongs to its result.
        ['partial', lambda(['x'], app(add, [local('x', 0)], unary), binary)],
    ]);
    const generated = CodeGen.translateWithFunctions(metadata)(producer);
    const consumer = moduleOf('Consumer', [
        ['viaAlias', lambda(['x'], app(variable('Producer', 'alias', unary),
            [local('x', 0)], int), unary)],
        ['viaWorker', lambda(['x'], app(add, [local('x', 0), literal(20)], int), unary)],
        ['viaPartial', lambda(['x'], app(variable('Producer', 'partial', binary),
            [literal(10), local('x', 0)], int), unary)],
    ]);
    const code = CodeGen.translate({ ...metadata, globalFunctions: generated.functions })(consumer);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-imported-workers-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/Producer.go'), generated.code);
    writeFileSync(join(directory, 'purescript/Consumer.go'), code);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/purescript"
)
func main() {
    fmt.Printf("%d %d %d", purescript.Call_Consumer_viaAlias(32),
        purescript.Call_Consumer_viaWorker(22), purescript.Call_Consumer_viaPartial(32))
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, '42 42 42');
    assert.match(code, /Call_Producer_add\(/, 'ordinary imported workers stay direct');
    assert.match(code, /Call_Producer_partial\(/, 'workers returning closures stay direct');
    assert.doesNotMatch(code, /Call_Producer_alias\(/);
});
