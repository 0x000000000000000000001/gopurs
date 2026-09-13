import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, toUnfoldable } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { coerceGoExpr, generateReboxFunctions, unboxGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { exprTypeToGoType, exprTypeToGenericGoType } from '../output/Gopurs.GoTypes/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { coerceLiteralField } from '../output/Gopurs.RecordExprs/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as Core from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const goType = exprTypeToGoType(emptyMap)(emptySet)(emptySet)('Test');
const genericGoType = exprTypeToGenericGoType(emptyMap)(emptySet)(emptySet)([])('Test');
const record = tail => new Core.Record(new Core.Row([
    new Tuple('size', Core.Int.value), new Tuple('newSeed', Core.Int.value),
], tail));

for (const [name, convert] of [['ordinary', goType], ['generic', genericGoType]]) {
    test(`${name} records with unknown or polymorphic tails retain their full payload`, () => {
        for (const tail of [Core.Any.value, new Core.TypeVar('r')]) {
            assert.deepEqual(convert(record(new Just(tail))), Go.TypeValue.value);
        }
    });

    test(`${name} closed records still use sorted native fields`, () => {
        assert.deepEqual(convert(record(Nothing.value)), new Go.TypeRecord([
            new Tuple('newSeed', Go.TypeInt64.value), new Tuple('size', Go.TypeInt64.value),
        ]));
    });
}

test('boxed record literals box their native scalar and array fields', () => {
    const ref = Ref.new({ reboxPairs: emptySet })();
    for (const [type, expected] of [
        [Go.TypeInt64.value, 'gopurs_runtime.Int(field)'],
        [new Go.TypeNativeArray(Go.TypeValue.value), 'gopurs_runtime.Array(field)'],
    ]) {
        const value = { expr: new Go.GoVar('field'), exprType: type };
        const boxed = coerceLiteralField(ref)('Test')('key')(Go.TypeValue.value)(value);
        assert.equal(printGoExpr(boxed.value1), expected);
        const nativeRecord = new Go.TypeRecord([new Tuple('key', type)]);
        const native = coerceLiteralField(ref)('Test')('key')(nativeRecord)(value);
        assert.deepEqual(native.value1, value.expr);
    }
});

test('native Tuple payloads are converted before accessing a typed pointer', () => {
    const ref = Ref.new({ reboxPairs: emptySet })();
    const native = new Go.TypeStructValue('Data.Tuple.Tuple', [Go.TypeValue.value, Go.TypeValue.value]);
    const typed = new Go.TypeStructPointer(
        'Data_Data_Tuple_Tuple', 'Data.Tuple.Tuple',
        'Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]',
        [Go.TypeInt64.value, Go.TypeValue.value],
    );
    const result = coerceGoExpr(ref)('Test')(new Go.GoVar('tuple'))(native)(typed);
    const pairs = toUnfoldable(unfoldableArray)(Ref.read(ref)().reboxPairs);

    assert.equal(pairs.length, 1, 'conversion must register a field-wise rebox');
    assert.deepEqual(pairs[0].value0.value3, [Go.TypeValue.value, Go.TypeValue.value]);
    assert.deepEqual(pairs[0].value1, typed);
    assert.ok(result instanceof Go.GoCall);
    assert.match(result.value0.value0, /^Rebox_Test_/);
    const code = printGoExpr(result);
    assert.match(code, /CoerceToStruct\[Constructor_Data_Tuple_Tuple\[gopurs_runtime.Value, gopurs_runtime.Value\]\]/);
    assert.doesNotMatch(code, /CoerceToStruct\[Constructor_Data_Tuple_Tuple\[int64,/);
});

test('arrays of boxed Tuple values preserve fields when converted to native tuples', t => {
    const ref = Ref.new({
        reboxPairs: emptySet, pointerAdtPaths: emptyMap, enumAdts: emptySet,
        elidedCtors: emptySet, classDeclsFields: emptyMap,
        ctorTypes: insert(ordString)('Data_Tuple.Tuple')({
            vars: ['a', 'b'], fields: [new Core.TypeVar('a'), new Core.TypeVar('b')],
        })(emptyMap),
    })();
    const tuple = new Go.TypeStructPointer(
        'Data_Data_Tuple_Tuple', 'Data.Tuple.Tuple',
        'Constructor_Data_Tuple_Tuple[int64, int64]',
        [Go.TypeInt64.value, Go.TypeInt64.value],
    );
    const expression = coerceGoExpr(ref)('Test')(new Go.GoVar('input'))
        (Go.TypeValue.value)(new Go.TypeNativeArray(tuple));
    const nativeExpression = unboxGoExpr(ref)('Test')(new Go.GoVar('native'))
        (new Go.TypeNativeArray(Go.TypeValue.value))(new Go.TypeNativeArray(tuple));
    const helpers = generateReboxFunctions(ref)('Test')();
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-tuple-array-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "unsafe"
    "gopurs/output/gopurs_runtime"
)
type Constructor_Data_Tuple_Tuple[A, B any] struct { Rc uint32; V0 A; V1 B }
${helpers.join('\n')}
func main() {
    boxed := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{
        Rc: 1, V0: gopurs_runtime.Int(4), V1: gopurs_runtime.Int(9),
    }
    native := []gopurs_runtime.Value{
        {Type: gopurs_runtime.TypeConstructor, UnsafePtr: unsafe.Pointer(boxed)},
        {Type: gopurs_runtime.TypeConstructor},
    }
    input := gopurs_runtime.Array(native)
    result := ${printGoExpr(expression)}
    nativeResult := ${printGoExpr(nativeExpression)}
    fmt.Printf("%d %d %t / %d %d %t", result[0].V0, result[0].V1, result[1] == nil,
        nativeResult[0].V0, nativeResult[0].V1, nativeResult[1] == nil)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, '4 9 true / 4 9 true');
});
