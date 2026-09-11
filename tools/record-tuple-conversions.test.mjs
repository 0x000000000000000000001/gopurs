import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { empty as emptySet, toUnfoldable } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { coerceGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { exprTypeToGoType, exprTypeToGenericGoType } from '../output/Gopurs.GoTypes/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { coerceLiteralField } from '../output/Gopurs.RecordExprs/index.js';
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
