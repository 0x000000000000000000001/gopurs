import { withReboxFields } from './codegen-metadata.mjs';
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { buildPointerAdtMetadata } from '../output/Gopurs.AdtMetadata/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { hashString } from '../output/PureScript.Backend.Optimizer.FfiSupport/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const declaration = (name, leaves) => ({ name, vars: [], constructors: [
    { name: `${name}Payload`, fields: [C.String.value] },
    ...leaves.map(name => ({ name, fields: [] })),
] });
const declarations = [declaration('Family', ['IPv4', 'IPv6']),
    declaration('Optional', ['Absent']), declaration('Product', [])];
const pointerMetadata = buildPointerAdtMetadata([{ name: 'Producer', dataDecls: declarations }]);

test('pointer ADTs permit only one nullary alternative', () => {
    const path = name => lookup(ordString)(`Producer.${name}`)(pointerMetadata.pointerAdtPaths);
    assert.ok(path('Family') instanceof Nothing, 'two distinct nullary tags cannot share a nil pointer');
    assert.ok(path('Optional') instanceof Just, 'one payload plus one nullary constructor remains native');
    assert.ok(path('Product') instanceof Just, 'a single payload constructor remains native');
    assert.ok(lookup(ordString)('Data_Producer_Absent')(pointerMetadata.pointerAdtLeaves) instanceof Just);
});

test('mixed ADTs retain every nullary tag and their payload through generated workers', t => {
    const expr = syntax => new NeutralExpr(syntax);
    const typed = (type, value) => expr(new S.Typed(type, value));
    const family = new C.ADT('Producer.Family', ['Producer', 'Family'], []);
    const unary = new C.Func([family], C.Boolean.value);
    const bindings = ['IPv4', 'IPv6', 'FamilyPayload'].map(tag => ({
        recursive: false,
        bindings: [new Tuple(`is${tag}`, typed(unary,
            expr(new S.Abs([new Tuple(new Just('value'), 0)],
                typed(C.Boolean.value, expr(new S.PrimOp(new S.Op1(
                    new S.OpIsTag(new C.Qualified(new Just('Producer'), tag)),
                    typed(family, expr(new S.Local(new Just('value'), 0))),
                ))))))))],
    }));
    const code = CodeGen.translate(withReboxFields({
        ...pointerMetadata, elidedCtors: emptySet, ctorTypes: emptyMap,
        enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap,
        globalFunctions: emptyMap, classDeclsFields: emptyMap,
    }))({ name: 'Consumer', bindings, comments: [], imports: emptySet,
        exports: emptySet, reExports: emptySet, dataTypes: emptyMap,
        dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-mixed-constructor-tags-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/Consumer.go'), code);
    writeFileSync(join(directory, 'purescript/Producer.go'), `package purescript
type Constructor_Producer_FamilyPayload struct { Rc uint32; V0 string }
`);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "unsafe"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
func main() {
    payload := &purescript.Constructor_Producer_FamilyPayload{Rc: 1, V0: "custom"}
    values := []gopurs_runtime.Value{
        {Type: 9, IntVal: ${hashString('Data_Producer_IPv4')}},
        {Type: 9, IntVal: ${hashString('Data_Producer_IPv6')}},
        {Type: 9, IntVal: ${hashString('Data_Producer_FamilyPayload')}, UnsafePtr: unsafe.Pointer(payload)},
    }
    for _, value := range values {
        for _, check := range []gopurs_runtime.Value{
            purescript.Get_Consumer_isIPv4(), purescript.Get_Consumer_isIPv6(),
            purescript.Get_Consumer_isFamilyPayload(),
        } {
            fmt.Printf("%t ", gopurs_runtime.Apply(check, value).IntVal != 0)
        }
        fmt.Println()
    }
    fmt.Print(payload.V0)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, 'true false false \nfalse true false \nfalse false true \ncustom');
});
