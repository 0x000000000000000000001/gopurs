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

test('a dynamic consumer receives the original boxed record without rebuilding its layout', t => {
    const metadata = {
        elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
        pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
        enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap,
        classDeclsFields: emptyMap,
    };
    const expr = syntax => new NeutralExpr(syntax);
    const local = (name, level) => expr(new S.Local(new Just(name), level));
    const entry = new C.Record(new C.Row([
        new Tuple('count', C.Int.value), new Tuple('label', C.String.value),
    ], Nothing.value));
    const body = expr(new S.Abs([
        new Tuple(new Just('consume'), 0), new Tuple(new Just('record'), 1),
    ], expr(new S.App(local('consume', 0), [
        expr(new S.Typed(entry, local('record', 1))),
    ]))));
    const code = CodeGen.translate(metadata)({
        name: 'RecordArgument', bindings: [{
            recursive: false, bindings: [new Tuple('forward', body)],
        }],
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-boxed-record-argument-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/RecordArgument.go'), code);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
func main() {
    original := gopurs_runtime.RecordDict2("label", "count", gopurs_runtime.Str("alpha"), gopurs_runtime.Int(5))
    calls := 0
    consumer := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        calls++
        fields := (*gopurs_runtime.RecordData2)(value.UnsafePtr)
        fmt.Printf("%t %s|%s %d:%s\\n", value.UnsafePtr == original.UnsafePtr,
            fields.K0, fields.K1, gopurs_runtime.RecordGet(value, "count").IntVal,
            gopurs_runtime.RecordGet(value, "label").StrVal())
        return gopurs_runtime.Value{}
    })
    purescript.Call_RecordArgument_forward(consumer, original)
    fmt.Println("calls", calls)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, 'true label|count 5:alpha\ncalls 1\n');
});
