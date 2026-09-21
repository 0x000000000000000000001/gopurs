import { withReboxFields } from './codegen-metadata.mjs';
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, toUnfoldable } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { boxGoExpr, coerceGoExpr, generateReboxFunctions, unboxGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { exprTypeToGoType, exprTypeToGenericGoType } from '../output/Gopurs.GoTypes/index.js';
import { printGoDecl, printGoExpr } from '../output/Gopurs.Printer/index.js';
import { coerceLiteralField, getProp, prepareLiteral } from '../output/Gopurs.RecordExprs/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as Core from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { hashString } from '../output/PureScript.Backend.Optimizer.FfiSupport/index.js';

const emptyMetadata = withReboxFields({
    pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
    enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet,
    ctorTypes: emptyMap, classDeclsFields: emptyMap, globalTypes: emptyMap, globalFunctions: emptyMap,
});
const newState = () => Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();

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

test('duplicate row labels preserve the first field through native access and boxing', t => {
    const ref = newState();
    const cases = [
        { name: 'string', first: Core.String.value, second: Core.Int.value, input: 'gopurs_runtime.Str("left")', read: 'StrVal()', format: '%s', expected: 'left' },
        { name: 'int', first: Core.Int.value, second: Core.String.value, input: 'gopurs_runtime.Int(2)', read: 'IntVal', format: '%d', expected: '2' },
    ];
    const blocks = [];
    const expected = [];
    for (const [mode, convert] of [['ordinary', goType], ['generic', genericGoType]]) {
        for (const fixture of cases) {
            const row = new Core.Record(new Core.Row([
                new Tuple('z', Core.Boolean.value), new Tuple('y', fixture.first),
                new Tuple('x', Core.Int.value), new Tuple('y', fixture.second),
            ], Nothing.value));
            const type = convert(row);
            const native = unboxGoExpr(ref)('Test')(new Go.GoVar('input'))(Go.TypeValue.value)(type);
            const boxed = boxGoExpr(ref)('Test')(new Go.GoVar('native'))(type);
            const field = getProp(emptyMetadata)(ref)('Test')('y')({ expr: new Go.GoVar('native'), exprType: type });
            const fieldValue = boxGoExpr(ref)('Test')(field.expr)(field.exprType);
            blocks.push(`{
    input := gopurs_runtime.RecordDict3("x", "y", "z", gopurs_runtime.Int(1), ${fixture.input}, gopurs_runtime.Bool(true))
    native := ${printGoExpr(native)}
    boxed := ${printGoExpr(boxed)}
    field := ${printGoExpr(fieldValue)}
    fmt.Printf("${mode}/${fixture.name} %d ${fixture.format} %t ${fixture.format}\\n",
        gopurs_runtime.RecordGet(boxed, "x").IntVal, gopurs_runtime.RecordGet(boxed, "y").${fixture.read},
        gopurs_runtime.RecordGet(boxed, "z").IntVal != 0, field.${fixture.read})
}`);
            expected.push(`${mode}/${fixture.name} 1 ${fixture.expected} true ${fixture.expected}`);
        }
    }
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-duplicate-row-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "gopurs/output/gopurs_runtime")
func main() { ${blocks.join('\n')} }
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, expected.join('\n') + '\n');
});

test('record literals use the first duplicate label for field typing, including open rows', () => {
    for (const tail of [Nothing.value, new Just(new Core.TypeVar('r'))]) {
        const row = new Core.Record(new Core.Row([
            new Tuple('y', Core.String.value), new Tuple('y', Core.Int.value),
        ], tail));
        const literal = prepareLiteral(emptyMetadata)('Test')(row)(Nothing.value);
        assert.deepEqual(lookup(ordString)('y')(literal.fields), new Just(Core.String.value));
    }
});

test('boxed record literals box their native scalar and array fields', () => {
    const ref = newState();
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
    const ref = newState();
    const native = new Go.TypeStructValue('Data.Tuple.Tuple', [Go.TypeValue.value, Go.TypeValue.value]);
    const typed = Go.structPointer({ baseStructName: 'Data_Data_Tuple_Tuple', fullName: 'Data.Tuple.Tuple', structName: 'Constructor_Data_Tuple_Tuple' })([Go.TypeInt64.value, Go.TypeValue.value]);
    const result = coerceGoExpr(ref)('Test')(new Go.GoVar('tuple'))(native)(typed);
    const pairs = toUnfoldable(unfoldableArray)(Ref.read(ref)().reboxPairs);

    assert.equal(pairs.length, 1, 'conversion must register a field-wise rebox');
    assert.deepEqual(pairs[0].value0.value0.typeArgs, [Go.TypeValue.value, Go.TypeValue.value]);
    assert.deepEqual(pairs[0].value1, typed);
    assert.ok(result instanceof Go.GoCall);
    assert.match(result.value0.value0, /^Rebox_Test_/);
    const code = printGoExpr(result);
    assert.match(code, /CoerceToStruct\[Constructor_Data_Tuple_Tuple\[gopurs_runtime.Value, gopurs_runtime.Value\]\]/);
    assert.doesNotMatch(code, /CoerceToStruct\[Constructor_Data_Tuple_Tuple\[int64,/);
});

test('arrays of boxed Tuple values preserve fields when converted to native tuples', t => {
    const ref = newState();
    const metadata = withReboxFields({
        ...emptyMetadata,
        ctorTypes: insert(ordString)('Data_Tuple.Tuple')({
            vars: ['a', 'b'], fields: [new Core.TypeVar('a'), new Core.TypeVar('b')],
        })(emptyMap),
    });
    const tuple = Go.structPointer({ baseStructName: 'Data_Data_Tuple_Tuple', fullName: 'Data.Tuple.Tuple', structName: 'Constructor_Data_Tuple_Tuple' })([Go.TypeInt64.value, Go.TypeInt64.value]);
    const expression = coerceGoExpr(ref)('Test')(new Go.GoVar('input'))
        (Go.TypeValue.value)(new Go.TypeNativeArray(tuple));
    const nativeExpression = unboxGoExpr(ref)('Test')(new Go.GoVar('native'))
        (new Go.TypeNativeArray(Go.TypeValue.value))(new Go.TypeNativeArray(tuple));
    const helpers = generateReboxFunctions(metadata)(ref)('Test')().map(printGoDecl);
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

test('generic class properties retain the tag and fields of native ADT values', t => {
    const dateType = Go.structPointer({ baseStructName: 'Data_Fixture_Date', fullName: 'Fixture.Date', structName: 'Constructor_Fixture_Date' })([]);
    const boundedType = Go.structPointer({ baseStructName: 'Data_Fixture_Bounded', fullName: 'Fixture.Bounded', structName: 'Constructor_Fixture_Bounded' })([dateType]);
    const ref = newState();
    const metadata = withReboxFields({
        ...emptyMetadata,
        classDeclsFields: insert(ordString)('Fixture.Bounded')({
            vars: ['a'], fields: [{ name: 'bottom', type: new Core.TypeVar('a') }],
        })(emptyMap),
    });
    const property = getProp(metadata)(ref)('Test')('bottom')({ expr: new Go.GoVar('dictionary'), exprType: boundedType });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-class-adt-field-'));
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
type Constructor_Fixture_Date struct { Rc uint32; Year int64; Month int64; Day int64 }
type Constructor_Fixture_Bounded[A any] struct { Rc uint32; V0 A }
func main() {
    dictionary := &Constructor_Fixture_Bounded[*Constructor_Fixture_Date]{
        Rc: 1, V0: &Constructor_Fixture_Date{Rc: 1, Year: 2026, Month: 9, Day: 13},
    }
    value := ${printGoExpr(property.expr)}
    if value.Type != gopurs_runtime.TypeConstructor {
        fmt.Printf("unexpected type %d", value.Type)
        return
    }
    date := (*Constructor_Fixture_Date)(unsafe.Pointer(value.UnsafePtr))
    fmt.Printf("%t %d %d %d", value.IntVal == ${hashString('Data_Fixture_Date')}, date.Year, date.Month, date.Day)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, 'true 2026 9 13');
});

test('boxed record fields convert generic ADT payloads to their native layout', t => {
    const ref = newState();
    const metadata = withReboxFields({
        ...emptyMetadata,
        ctorTypes: insert(ordString)('Data_Tuple.Tuple')({
            vars: ['a', 'b'], fields: [new Core.TypeVar('a'), new Core.TypeVar('b')],
        })(emptyMap),
    });
    const tuple = Go.structPointer({ baseStructName: 'Data_Data_Tuple_Tuple', fullName: 'Data.Tuple.Tuple', structName: 'Constructor_Data_Tuple_Tuple' })([Go.TypeInt64.value, Go.TypeInt64.value]);
    const recordType = new Go.TypeRecord([new Tuple('payload', tuple)]);
    const expression = unboxGoExpr(ref)('Test')(new Go.GoVar('input'))(Go.TypeValue.value)(recordType);
    const helpers = generateReboxFunctions(metadata)(ref)('Test')().map(printGoDecl);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-record-adt-field-'));
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
    input := gopurs_runtime.RecordDict1("payload", gopurs_runtime.Value{
        Type: gopurs_runtime.TypeConstructor, UnsafePtr: unsafe.Pointer(boxed),
    })
    result := ${printGoExpr(expression)}
    fmt.Printf("%d %d", result.payload.V0, result.payload.V1)
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, '4 9');
});
