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
import { buildPointerAdtMetadata } from '../output/Gopurs.AdtMetadata/index.js';
import { buildConstructorTypes } from '../output/Gopurs.ConstructorMetadata/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { hashString } from '../output/PureScript.Backend.Optimizer.FfiSupport/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

test('boxed constructor tests inspect the tag without copying recursive payloads', t => {
    const expr = syntax => new NeutralExpr(syntax);
    const typed = (type, value) => expr(new S.Typed(type, value));
    const list = item => new C.ADT('Data.List.Types.List', ['Data', 'List', 'Types', 'List'], [item]);
    const entry = new C.Record(new C.Row([
        new Tuple('count', C.Int.value), new Tuple('label', C.String.value),
    ], Nothing.value));
    const listType = list(entry);
    const modules = [{ name: 'Data.List.Types', dataDecls: [{
        name: 'List', vars: ['a'], constructors: [
            { name: 'Nil', fields: [] },
            { name: 'Cons', fields: [new C.TypeVar('a'), list(new C.TypeVar('a'))] },
        ],
    }] }];
    const pointerMetadata = buildPointerAdtMetadata(modules);
    const bindings = [];
    for (const observe of [false, true]) {
        for (const tag of ['Nil', 'Cons']) {
            let operand = expr(new S.Local(new Just('value'), 0));
            if (observe) operand = expr(new S.App(expr(new S.Var(
                new C.Qualified(new Just('Probe'), 'observe'))), [operand]));
            // The parameter stays boxed, as in Builder's polymorphic local loop.
            // These annotations describe its payload; a tag test needs neither
            // a native List nor conversion of every record in its suffix.
            operand = typed(listType, typed(listType, operand));
            bindings.push({ recursive: false, bindings: [new Tuple(
                `${observe ? 'observed' : 'is'}${tag}`,
                expr(new S.Abs([new Tuple(new Just('value'), 0)],
                    typed(C.Boolean.value, expr(new S.PrimOp(new S.Op1(
                        new S.OpIsTag(new C.Qualified(new Just('Data.List.Types'), tag)), operand,
                    )))))),
            )] });
        }
    }
    const code = CodeGen.translate({
        ...pointerMetadata, elidedCtors: emptySet, ctorTypes: buildConstructorTypes(modules),
        enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap,
        globalFunctions: emptyMap, classDeclsFields: emptyMap,
    })({ name: 'TagChecks', bindings, comments: [], imports: emptySet,
        exports: emptySet, reExports: emptySet, dataTypes: emptyMap,
        dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-boxed-constructor-tags-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/TagChecks.go'), code);
    writeFileSync(join(directory, 'purescript/Probe.go'), `package purescript
import "gopurs/output/gopurs_runtime"
type Constructor_Data_List_Types_Cons[A any] struct { Rc uint32; V0 A; V1 *Constructor_Data_List_Types_Cons[A] }
var Observations int
func Get_Probe_observe() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        Observations++
        return value
    })
}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "encoding/json"
    "os"
    "testing"
    "unsafe"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
const tag = ${hashString('Data_Data_List_Types_Cons')}
type node = purescript.Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
func boxed(value *node) gopurs_runtime.Value {
    return gopurs_runtime.Value{Type: 9, IntVal: tag, UnsafePtr: unsafe.Pointer(value)}
}
func main() {
    isNil := purescript.Get_TagChecks_isNil()
    isCons := purescript.Get_TagChecks_isCons()
    observedNil := purescript.Get_TagChecks_observedNil()
    observedCons := purescript.Get_TagChecks_observedCons()
    item := gopurs_runtime.RecordDict2("label", "count", gopurs_runtime.Str("kept"), gopurs_runtime.Int(7))
    values := []gopurs_runtime.Value{boxed(nil), boxed(&node{Rc: 1, V0: item})}
    for index, value := range values {
        for testIndex, check := range []gopurs_runtime.Value{isNil, isCons, observedNil, observedCons} {
            actual := gopurs_runtime.Apply(check, value).IntVal != 0
            expected := index == testIndex % 2
            if actual != expected { panic("wrong constructor") }
        }
    }
    if purescript.Observations != 4 { panic("operand evaluated more than once") }
    measurements := make(map[int]float64)
    for _, size := range []int{32, 64, 128, 256} {
        var head *node
        for i := 0; i < size; i++ { head = &node{Rc: 1, V0: item, V1: head} }
        measurements[size] = testing.AllocsPerRun(3, func() {
            for current := head; current != nil; current = current.V1 {
                value := boxed(current)
                if gopurs_runtime.Apply(isNil, value).IntVal != 0 ||
                    gopurs_runtime.Apply(isCons, value).IntVal == 0 { panic("wrong suffix tag") }
            }
        })
        if head.V0.UnsafePtr != item.UnsafePtr || gopurs_runtime.RecordGet(head.V0, "count").IntVal != 7 {
            panic("input mutated")
        }
    }
    json.NewEncoder(os.Stdout).Encode(measurements)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    const allocations = JSON.parse(result.stdout);
    t.diagnostic(`allocations while walking suffixes: ${JSON.stringify(allocations)}`);
    assert.ok(allocations[256] <= allocations[32] * 9 + 16,
        'tag checks must have constant cost per node, independent of the suffix length');
});
