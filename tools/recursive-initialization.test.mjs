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

const metadata = {
    elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap,
    classDeclsFields: emptyMap,
};
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (name, level = 1) => expr(new S.Local(new Just(name), level));
const literal = value => expr(new S.Lit(new C.LitInt(value)));
const lambda = (name, level, body) => expr(new S.Abs([new Tuple(new Just(name), level)], body));
const app = fn => expr(new S.App(fn, [literal(0)]));
const prop = (record, key) => expr(new S.Accessor(record, new S.GetProp(key)));
const record = fields => typed(
    new C.Record(new C.Row(fields.map(([key, type]) => new Tuple(key, type)), Nothing.value)),
    expr(new S.Lit(new C.LitRecord(fields.map(([key, , value]) => new Tuple(key, value))))));
const rec = (bindings, body) => expr(new S.LetRec(1,
    bindings.map(([name, value]) => new Tuple(name, value)), body));

test('recursive native values reject early reads and publish complete initializers in order', t => {
    const self = local('self');
    const fixtures = [
        ['eagerRecord', rec([['self', record([
            ['a', C.Int.value, literal(1)],
            ['b', C.Int.value, prop(self, 'a')],
        ])]], self)],
        ['indirectRecord', rec([
            ['read', lambda('ignored', 2, prop(self, 'a'))],
            ['self', record([['a', C.Int.value, app(local('read'))]])],
        ], self)],
        ['eagerInt', rec([['self', typed(C.Int.value, self)]], self)],
        ['delayedRecord', rec([['self', record([
            ['a', C.Int.value, literal(1)],
            ['backref', C.Any.value, lambda('ignored', 2, prop(self, 'a'))],
        ])]], app(prop(self, 'backref')))],
        ['orderedInitializers', rec([
            ['alpha', record([
                ['a', C.Int.value, literal(7)],
                ['backref', C.Any.value, lambda('ignored', 2, local('bravo'))],
            ])],
            ['bravo', typed(C.Int.value, expr(new S.Let(new Just('value'), 2,
                prop(local('alpha'), 'a'), local('value', 2))))],
        ], app(prop(local('alpha'), 'backref')))],
    ];
    const code = CodeGen.translate(metadata)({
        name: 'Recursive', bindings: fixtures.map(([name, body]) => ({
            recursive: false, bindings: [new Tuple(name, lambda('ignored', 0, body))],
        })),
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-recursive-initialization-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/Recursive.go'), code);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "strings"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
func report(name string, fn func(gopurs_runtime.Value) gopurs_runtime.Value) {
    defer func() {
        if err := recover(); err != nil {
            if !strings.Contains(fmt.Sprint(err), "nil pointer") { panic(err) }
            fmt.Println(name, "uninitialized")
        }
    }()
    result := fn(gopurs_runtime.Value{})
    fmt.Println(name, result.IntVal)
}
func main() {
    ${fixtures.map(([name]) => `report("${name}", purescript.Call_Recursive_${name})`).join('\n    ')}
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, [
        'eagerRecord uninitialized',
        'indirectRecord uninitialized',
        'eagerInt uninitialized',
        'delayedRecord 1',
        'orderedInitializers 7',
        '',
    ].join('\n'));
});
