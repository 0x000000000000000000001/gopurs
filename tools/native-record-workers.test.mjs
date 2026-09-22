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
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
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
const string = C.String.value;
const bool = C.Boolean.value;
const fields = [['id', int], ['name', string], ['active', bool]];
const record = (entries, tail = Nothing.value) => new C.Record(new C.Row(
    entries.map(([name, type]) => new Tuple(name, type)), tail));
const openRecord = record(fields, new Just(new C.TypeVar('r')));
const firstLayout = record([...fields, ['email', string]]);
const secondLayout = record([...fields, ['rank', int], ['zone', string]]);
const readType = new C.ForAll(['r'], new C.Func([openRecord], int));
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (name, level, type) => typed(type, expr(new S.Local(new Just(name), level)));
const variable = (module, name, type) => typed(type,
    expr(new S.Var(new C.Qualified(new Just(module), name))));
const integer = value => typed(int, expr(new S.Lit(new C.LitInt(value))));
const text = value => typed(string, expr(new S.Lit(new C.LitString(value))));
const field = (value, name, type) => typed(type, expr(new S.Accessor(value, new S.GetProp(name))));
const app = (fn, args, type) => typed(type, expr(new S.App(fn, args)));
const lambda = (params, body, type) => typed(type, expr(new S.Abs(
    params.map(([name, level]) => new Tuple(new Just(name), level)), body)));
const letIn = (name, level, value, body, type) => typed(type,
    expr(new S.Let(new Just(name), level, value, body)));
const add = (left, right) => typed(int, expr(new S.PrimOp(
    new S.Op2(new S.OpIntNum(S.OpAdd.value), left, right))));
const choose = (condition, yes, no) => typed(int,
    expr(new S.Branch([new S.Pair(condition, yes)], no)));
const row = () => local('row', 0, openRecord);
const readBody = () => choose(field(row(), 'active', bool),
    add(field(row(), 'id', int), choose(typed(bool, expr(new S.PrimOp(
        new S.Op2(new S.OpStringOrd(S.OpEq.value), field(row(), 'name', string), text('alpha')),
    ))), integer(1), integer(2))), add(field(row(), 'id', int), integer(-3)));
const moduleOf = (name, bindings) => ({
    name, bindings: bindings.map(([name, value]) => ({
        recursive: false, bindings: [new Tuple(name, value)],
    })),
    comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
    dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
    implementations: emptyMap, directives: emptyMap,
});
const info = (generated, module, name) => {
    const found = lookup(ordString)(`${module}.${name}`)(generated.functions);
    assert.ok(found instanceof Just, `${module}.${name} must publish its actual worker signature`);
    return found.value0;
};

function readerModule() {
    return CodeGen.translateWithFunctions(metadata)(moduleOf('NativeRecordReader', [
        ['read', lambda([['row', 0]], readBody(), readType)],
        ['withOffset', lambda([['row', 0], ['offset', 1]],
            add(readBody(), local('offset', 1, int)),
            new C.ForAll(['r'], new C.Func([openRecord, int], int)))],
    ]));
}

function fallbackModule() {
    const callback = new C.Func([openRecord], int);
    const closure = new C.Func([int], int);
    const unknown = record([['id', C.Any.value]], new Just(new C.TypeVar('r')));
    const element = new C.TypeVar('a');
    const generic = record([['id', element]], new Just(new C.TypeVar('r')));
    return CodeGen.translateWithFunctions(metadata)(moduleOf('NativeRecordFallback', [
        ['returned', lambda([['row', 0]], row(),
            new C.ForAll(['r'], new C.Func([openRecord], openRecord)))],
        ['forwarded', lambda([['row', 0], ['consume', 1]],
            app(local('consume', 1, callback), [row()], int),
            new C.ForAll(['r'], new C.Func([openRecord, callback], int)))],
        ['updated', lambda([['row', 0]], typed(openRecord, expr(new S.Update(row(), [
            new C.Prop('id', integer(99)),
        ]))), new C.ForAll(['r'], new C.Func([openRecord], openRecord)))],
        ['unknownField', lambda([['row', 0]], field(local('row', 0, unknown), 'id', C.Any.value),
            new C.ForAll(['r'], new C.Func([unknown], C.Any.value)))],
        ['polymorphicField', lambda([['row', 0]], field(local('row', 0, generic), 'id', element),
            new C.ForAll(['a', 'r'], new C.Func([generic], element)))],
        ['missingField', lambda([['row', 0]], field(row(), 'extra', int),
            new C.ForAll(['r'], new C.Func([openRecord], int)))],
        // Keep an intervening Let so the nested function remains a closure
        // capturing row, rather than becoming another saturated parameter.
        ['captured', lambda([['row', 0]], letIn('saved', 1, field(row(), 'id', int),
            lambda([['offset', 2]], add(field(row(), 'id', int),
                add(local('saved', 1, int), local('offset', 2, int))), closure), closure),
            new C.ForAll(['r'], new C.Func([openRecord], closure)))],
    ]));
}

test('open-row readers share native field signatures across wider record layouts', t => {
    const generated = readerModule();
    const expected = new Go.TypeRecord([
        new Tuple('active', Go.TypeBool.value), new Tuple('id', Go.TypeInt64.value),
        new Tuple('name', Go.TypeString.value),
    ]);
    assert.deepEqual(info(generated, 'NativeRecordReader', 'read').fArgs, [expected]);
    assert.deepEqual(info(generated, 'NativeRecordReader', 'read').fRet, Go.TypeInt64.value);
    assert.deepEqual(info(generated, 'NativeRecordReader', 'withOffset').fArgs,
        [expected, Go.TypeInt64.value], 'other arguments retain their native types');
    const header = generated.code.match(/func Call_NativeRecordReader_read\([\s\S]*?\) int64 \{/);
    assert.ok(header, 'the emitted worker must retain its native return signature');
    // Local anonymous structs also contain closing braces at column zero.
    // This fixture has no braces in string literals, so count nested braces.
    const start = header.index + header[0].length;
    let end = start;
    for (let depth = 1; depth > 0 && end < generated.code.length; end++) {
        if (generated.code[end] === '{') depth++;
        if (generated.code[end] === '}') depth--;
    }
    const body = generated.code.slice(start, end - 1);
    for (const name of ['active', 'id', 'name']) {
        assert.match(body, new RegExp(`\\.${name}\\b`), `${name} is read natively`);
    }
    assert.doesNotMatch(body, /gopurs_runtime\.Record(?:Get|Dict)/,
        'typed open-row annotations must not rebox the native worker parameter');
    const consumer = CodeGen.translate({ ...metadata, globalFunctions: generated.functions })(
        moduleOf('NativeRecordConsumer', [firstLayout, secondLayout].map((layout, index) => [
            `from${index}`, lambda([['input', 0]], app(
                variable('NativeRecordReader', 'read', readType),
                [local('input', 0, layout)], int), new C.Func([layout], int)),
        ])));
    assert.equal([...generated.code.matchAll(/func Call_NativeRecordReader_read\(/g)].length, 1,
        'one worker serves both layouts');
    assert.equal([...consumer.matchAll(/Call_NativeRecordReader_read\(/g)].length, 2,
        'both imported calls use the same published worker');
    assert.doesNotMatch(consumer, /Get_NativeRecordReader_read\(/,
        'known imported calls must not fall back to dynamic application');
    const fallback = fallbackModule();

    const directory = mkdtempSync(join(tmpdir(), 'gopurs-native-record-workers-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/Reader.go'), generated.code);
    writeFileSync(join(directory, 'purescript/Consumer.go'), consumer);
    writeFileSync(join(directory, 'purescript/Fallback.go'), fallback.code);
    // Anonymous native record fields are private to package purescript.
    writeFileSync(join(directory, 'purescript/Probe.go'), `package purescript
import (
    "fmt"
    "gopurs/output/gopurs_runtime"
)
func RunRecordWorkerProbe() {
    first := struct { active bool; email string; id int64; name string }{true, "kept", 40, "alpha"}
    second := struct { active bool; id int64; name string; rank int64; zone string }{false, 45, "beta", 7, "west"}
    firstBefore, secondBefore := first, second
    fmt.Println("native", Call_NativeRecordConsumer_from0(first), Call_NativeRecordConsumer_from1(second))
    fmt.Println("unchanged", first == firstBefore, second == secondBefore)
    // Deliberately different property order, with a field unknown to the reader.
    original := gopurs_runtime.RecordDict4("name", "extra", "id", "active",
        gopurs_runtime.Str("beta"), gopurs_runtime.Int(7), gopurs_runtime.Int(40), gopurs_runtime.Bool(true))
    pointer := original.UnsafePtr
    read := gopurs_runtime.Apply(Get_NativeRecordReader_read(), original)
    offset := gopurs_runtime.Apply2(Get_NativeRecordReader_withOffset(), original, gopurs_runtime.Int(8))
    partial := gopurs_runtime.Apply(Get_NativeRecordReader_withOffset(), original)
    fmt.Println("wrapped", read.IntVal, offset.IntVal, gopurs_runtime.Apply(partial, gopurs_runtime.Int(1)).IntVal)
    returned := gopurs_runtime.Apply(Get_NativeRecordFallback_returned(), original)
    fmt.Println("returned", returned.UnsafePtr == pointer, gopurs_runtime.RecordGet(returned, "extra").IntVal)
    calls := 0
    callback := gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
        calls++
        if value.UnsafePtr != pointer { panic("callback lost the original row") }
        return gopurs_runtime.RecordGet(value, "extra")
    })
    forwarded := gopurs_runtime.Apply2(Get_NativeRecordFallback_forwarded(), original, callback)
    fmt.Println("callback", calls, forwarded.IntVal)
    updated := gopurs_runtime.Apply(Get_NativeRecordFallback_updated(), original)
    fmt.Println("updated", gopurs_runtime.RecordGet(updated, "id").IntVal, gopurs_runtime.RecordGet(updated, "extra").IntVal)
    captured := gopurs_runtime.Apply(Get_NativeRecordFallback_captured(), original)
    fmt.Println("captured", gopurs_runtime.Apply(captured, gopurs_runtime.Int(2)).IntVal,
        gopurs_runtime.Apply(captured, gopurs_runtime.Int(3)).IntVal)
    unknown := gopurs_runtime.Apply(Get_NativeRecordFallback_unknownField(), original)
    polymorphic := gopurs_runtime.Apply(Get_NativeRecordFallback_polymorphicField(), original)
    missing := gopurs_runtime.Apply(Get_NativeRecordFallback_missingField(), original)
    fmt.Println("fallback-fields", unknown.IntVal, polymorphic.IntVal, missing.IntVal)
    fmt.Println("boxed-unchanged", original.UnsafePtr == pointer,
        gopurs_runtime.RecordGet(original, "id").IntVal,
        gopurs_runtime.RecordGet(original, "name").StrVal(),
        gopurs_runtime.RecordGet(original, "active").IntVal != 0,
        gopurs_runtime.RecordGet(original, "extra").IntVal)
}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import "gopurs/output/purescript"
func main() { purescript.RunRecordWorkerProbe() }
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout, 'native 41 42\nunchanged true true\nwrapped 42 50 43\n'
        + 'returned true 7\ncallback 1 7\nupdated 99 7\ncaptured 82 83\n'
        + 'fallback-fields 40 40 7\nboxed-unchanged true 40 beta true 7\n');
});

test('escaping or unsupported open-row uses keep the complete Value argument', () => {
    const generated = fallbackModule();
    for (const name of [
        'returned', 'forwarded', 'updated', 'unknownField', 'polymorphicField', 'missingField', 'captured',
    ]) {
        assert.deepEqual(info(generated, 'NativeRecordFallback', name).fArgs[0], Go.TypeValue.value,
            `${name} cannot safely project away the row tail`);
    }
});
