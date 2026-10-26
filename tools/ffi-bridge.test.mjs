import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { ffiFunctionInfos, ffiValueWorkers, generateFfiBridge } from '../output/Gopurs.FfiBridge/index.js';
import { prepareFfi } from '../output/Gopurs.FfiSupport/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const moduleName = 'BridgeContract';
const int = C.Int.value;
const fn = (args, result) => new C.Func(args, result);
const callback = fn([int], int);
const record = new C.Record(new C.Row([new Tuple('count', int)], Nothing.value));
const effect = new C.ADT('Effect.Effect', ['Effect', 'Effect'], [int]);
const foreigns = [
    ['choose', fn([int], int)], ['suffix', fn([int], int)],
    ['callback', fn([callback, int], int)], ['record', fn([int], record)],
    ['unsupported', fn([callback, record], int)], ['noResult', fn([int], C.Unit.value)],
    ['effectful', fn([int], effect)], ['constant', int], ['zero', int],
    ['untyped', null], ['missing', fn([int], int)],
].map(([name, type]) => new Tuple(name, type === null ? Nothing.value : new Just(type)));

const prepared = prepareFfi({ moduleName, path: 'BridgeContract.go' })(moduleName + '_')(`
var Counter int64
var Constant int64 = 41
func Choose(x int64) int64 { return x + 1 }
func Choose_(x int64) int64 { return -1 }
func Suffix_(x int64) int64 { return x + 2 }
func Callback(f func(int64) int64, x int64) any { return f(x) }
func Record(x int64) any { return map[string]any{"count": x} }
func Unsupported(f func(int64) int64, fields map[string]any) any { return f(int64(len(fields))) }
func NoResult(x int64) { Counter = x }
func Effectful(x int64) func() any { return func() any { Counter = x; return x } }
func Zero() int64 { return 42 }
func Untyped(x int64) any { return x }
`)();
const signatures = ffiFunctionInfos(moduleName)(foreigns)(prepared.decls);
const wrappers = generateFfiBridge(moduleName)([])(prepared.decls)(foreigns);
const workers = ffiValueWorkers(moduleName)([])(foreigns)(prepared.decls);

test('Go int32 FFI preserves signed values through scalar, array and callback bridges', t => {
    const name = 'Int32Bridge';
    const array = new C.Array(int);
    const foreigns = [
        new Tuple('identity', new Just(fn([int], int))),
        new Tuple('array', new Just(fn([array], array))),
        new Tuple('callback', new Just(fn([callback, int], int))),
    ];
    const parsed = prepareFfi({ moduleName: name, path: name + '.go' })(name + '_')(`
func Identity(value int32) int32 { return value }
func Array(values []int32) []int32 { return values }
func Callback(callback func(int32) int32, value int32) int32 { return callback(value) }
`)();
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-int32-ffi-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'bridge.go'), 'package main\nimport "gopurs/output/gopurs_runtime"\n'
        + parsed.content + '\n' + generateFfiBridge(name)([])(parsed.decls)(foreigns));
    writeFileSync(join(directory, 'bridge_test.go'), `package main
import (
    "testing"
    rt "gopurs/output/gopurs_runtime"
)
func TestSignedInt32Bridges(t *testing.T) {
    if got := rt.Unbox[int32](rt.Any(int32(-42))); got != -42 {
        t.Fatalf("opaque int32 unboxing: %d", got)
    }
    values := []rt.Value{rt.Int(-2147483648), rt.Int(-1), rt.Int(0), rt.Int(2147483647)}
    callback := rt.Func(func(value rt.Value) rt.Value { return value })
    for _, value := range values {
        for _, result := range []rt.Value{
            rt.Apply(_Gopurs_Int32Bridge_Identity, value),
            rt.Apply2(_Gopurs_Int32Bridge_Callback, callback, value),
        } {
            if result.Type != rt.TypeInt || result.IntVal != value.IntVal {
                t.Fatalf("int32 round trip: got %+v, want %+v", result, value)
            }
        }
    }
    result := rt.Apply(_Gopurs_Int32Bridge_Array, rt.Array(values))
    returned := rt.Unbox[[]any](result)
    if len(returned) != len(values) { t.Fatalf("array length: %d", len(returned)) }
    for index, value := range returned {
        got := rt.Unbox[int64](value)
        if got != values[index].IntVal { t.Fatalf("array[%d]: %d", index, got) }
    }
}
`);
    const result = spawnSync('go', ['test', '-count=1', '-v', './...'], {
        cwd: directory, encoding: 'utf8', timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /--- PASS: TestSignedInt32Bridges/);
});

test('native-call metadata names the emitted worker and retains its calling convention', () => {
    for (const [name, fullName, fArgs, fRet] of [
        ['choose', 'BridgeContract_Choose', [Go.TypeInt64.value], Go.TypeInt64.value],
        ['suffix', 'BridgeContract_Suffix_', [Go.TypeInt64.value], Go.TypeInt64.value],
        ['callback', 'BridgeContract_Callback_nativeWorker', [Go.TypeValue.value, Go.TypeInt64.value], Go.TypeValue.value],
        ['noResult', 'BridgeContract_NoResult_nativeWorker', [Go.TypeInt64.value], Go.TypeValue.value],
    ]) {
        const found = lookup(ordString)(name)(signatures);
        assert.ok(found instanceof Just, name);
        const { reference, ...signature } = found.value0;
        assert.deepEqual(reference.value0, { name: fullName, owner: 'ffi/' + moduleName });
        assert.deepEqual(signature, { fArgs, fRet, arity: fArgs.length });
    }
    for (const name of ['record', 'unsupported', 'effectful', 'constant', 'zero', 'untyped', 'missing']) {
        assert.equal(lookup(ordString)(name)(signatures), Nothing.value, name);
    }
    // Emission precedes signature admissibility: record returns and unsupported
    // native arguments still have workers, although direct callers cannot use them.
    assert.deepEqual([...workers.matchAll(/^func (\w+)\(/gm)].map(match => match[1]), [
        'BridgeContract_Callback_nativeWorker', 'BridgeContract_Record_nativeWorker',
        'BridgeContract_Unsupported_nativeWorker', 'BridgeContract_NoResult_nativeWorker',
    ]);
});

test('boxed and native entries preserve callbacks, records, effects and declaration fallbacks', t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-bridge-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
    writeFileSync(join(directory, 'bridge.go'), 'package main\nimport "gopurs/output/gopurs_runtime"\n'
        + prepared.content + '\n' + wrappers + '\n' + workers);
    writeFileSync(join(directory, 'bridge_test.go'), `package main
import (
    "testing"
    rt "gopurs/output/gopurs_runtime"
)
func TestBridgeEntries(t *testing.T) {
    calls := 0
    callback := rt.Func(func(value rt.Value) rt.Value { calls++; return rt.Int(value.IntVal + 1) })
    direct := BridgeContract_Callback_nativeWorker(callback, 41)
    boxed := rt.Apply2(_Gopurs_BridgeContract_Callback, callback, rt.Int(41))
    if direct.IntVal != 42 || boxed.IntVal != 42 || calls != 2 {
        t.Fatalf("callback: direct=%+v boxed=%+v calls=%d", direct, boxed, calls)
    }
    rec := rt.Apply(_Gopurs_BridgeContract_Record, rt.Int(42))
    if rt.RecordGet(rec, "count").IntVal != 42 { t.Fatal("record wrapper lost its field") }
    fromMap := rt.Apply2(_Gopurs_BridgeContract_Unsupported, callback, rec)
    if fromMap.IntVal != 2 || calls != 3 { t.Fatal("map callback fallback") }
    BridgeContract_NoResult_nativeWorker(7)
    if BridgeContract_Counter != 7 { t.Fatal("void worker did not execute") }
    rt.Apply(_Gopurs_BridgeContract_NoResult, rt.Int(8))
    if BridgeContract_Counter != 8 { t.Fatal("void wrapper did not execute") }
    delayed := rt.Apply(_Gopurs_BridgeContract_Effectful, rt.Int(9))
    if BridgeContract_Counter != 8 { t.Fatal("effect executed before its thunk was called") }
    result := rt.Apply(delayed, rt.Value{})
    if BridgeContract_Counter != 9 || result.IntVal != 9 { t.Fatal("effect result") }
    for name, got := range map[string]rt.Value{
        "primary": rt.Apply(_Gopurs_BridgeContract_Choose, rt.Int(41)),
        "suffix": rt.Apply(_Gopurs_BridgeContract_Suffix, rt.Int(40)),
        "zero": rt.Apply(_Gopurs_BridgeContract_Zero, rt.Value{}),
        "untyped": rt.Apply(_Gopurs_BridgeContract_Untyped, rt.Int(42)),
    } {
        if got.IntVal != 42 { t.Fatalf("%s: %+v", name, got) }
    }
    if _Gopurs_BridgeContract_Constant.IntVal != 41 { t.Fatal("foreign variable") }
}
func TestMissingDeclarationFailsAtUse(t *testing.T) {
    defer func() {
        if got := recover(); got != "FFI not implemented: missing" { t.Fatalf("panic: %v", got) }
    }()
    rt.Apply(_Gopurs_BridgeContract_Missing, rt.Int(1))
    t.Fatal("missing foreign declaration succeeded")
}
`);
    const result = spawnSync('go', ['test', '-count=1', '-v', './...'], {
        cwd: directory, encoding: 'utf8', timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /--- PASS: TestBridgeEntries/);
    assert.match(result.stdout, /--- PASS: TestMissingDeclarationFailsAtUse/);
});
