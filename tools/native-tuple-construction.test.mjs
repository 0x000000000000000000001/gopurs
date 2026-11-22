import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Adt from '../output/Gopurs.AdtExprs/index.js';
import * as Constructors from '../output/Gopurs.ConstructorExprs/index.js';
import { StmtEmpty } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { boxGoExpr, coerceGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { exprTypeToGoType } from '../output/Gopurs.GoTypes/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import * as TypedExprs from '../output/Gopurs.TypedExprs/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { hashString } from '../output/PureScript.Backend.Optimizer.FfiSupport/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

const tuple = fields => new C.ADT('Data.Tuple.Tuple', ['Data', 'Tuple', 'Tuple'], fields);
const list = new C.ADT('Data.List.Types.List', ['Data', 'List', 'Types', 'List'], [C.Int.value]);
const position = new C.Record(new C.Row([
    new Tuple('column', C.Int.value), new Tuple('line', C.Int.value),
], Nothing.value));
const metadata = withReboxFields({
    pointerAdtPaths: insert(ordString)('Data.Tuple.Tuple')({ ctorName: 'Tuple', arity: 2 })(
        insert(ordString)('Data.List.Types.List')({ ctorName: 'Cons', arity: 1 })(emptyMap)),
    ctorTypes: insert(ordString)('Data_Tuple.Tuple')({
        vars: ['a', 'b'], fields: [new C.TypeVar('a'), new C.TypeVar('b')],
    })(emptyMap),
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, elidedCtors: emptySet, classDeclsFields: emptyMap,
    globalTypes: emptyMap, globalFunctions: emptyMap,
});
const goType = exprTypeToGoType(metadata.pointerAdtPaths)(emptySet)(emptySet)('NativeTuples');
const state = () => Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();
const prepare = (type, expected) => Adt.prepareSaturatedAs(metadata)('NativeTuples')(
    new Just('Data.Tuple'))('Tuple')(type)(expected);
const construct = (prepared, names = ['left', 'right']) => Adt.saturated(state())(prepared)({
    exprs: names.map(name => new Go.GoVar(name)),
    exprTypes: names.map((_, index) => Adt.saturatedFieldType(prepared)(index).goType),
});

function annotatedConstruction(fields, expected = Nothing.value) {
    const type = tuple(fields);
    const context = {
        metadata, codegenStateRef: state(), depth: 0, modNameStr: 'NativeTuples',
        recVars: [], moduleFunctions: emptyMap, bound: emptyMap, tcoIdent: Nothing.value,
        loopCtx: [], options: { isTail: false, inEffectBlock: false },
        mbExpectedExprType: Nothing.value, mbExpectedGoType: expected,
    };
    const properties = ['left', 'right'].map((name, index) => new Tuple(`value${index}`,
        new TcoExpr(null, new S.Local(new Just(name), index))));
    const syntax = new TcoExpr(null, new S.CtorSaturated(
        new C.Qualified(new Just('Data.Tuple'), 'Tuple'), C.ProductType.value, 'Tuple', 'Tuple', properties));
    const events = [];
    let constructorContext;
    const translateField = child => nextId => value => {
        const { value0: name, value1: index } = value.value1;
        events.push({ name: name.value0, expected: child.mbExpectedGoType });
        return { stmts: StmtEmpty.value, expr: new Go.GoVar(name.value0),
            exprType: goType(fields[index]), nextId: nextId + 1 };
    };
    const translate = child => nextId => value => {
        constructorContext = child;
        assert.equal(value, syntax, 'the original TAST/TCO node must reach constructor translation');
        return Constructors.saturated(translateField)(child)(nextId)(value)(new Just('Data.Tuple'))('Tuple')(properties);
    };
    const result = TypedExprs.annotated(translate)(context)(3)(type)(syntax);
    assert.equal(constructorContext.mbExpectedExprType.value0, type, 'keep the complete source annotation');
    assert.deepEqual(events.map(event => event.name), ['left', 'right']);
    assert.equal(result.nextId, 5);
    return { result, constructorContext, events };
}

test('annotated Tuple construction receives its final native layout before translating composite fields', () => {
    for (const fields of [[position, list], [list, new C.Array(C.Int.value)]]) {
        const { result, constructorContext, events } = annotatedConstruction(fields);
        const target = goType(tuple(fields));
        assert.deepEqual(constructorContext.mbExpectedGoType, new Just(target));
        assert.deepEqual(events.map(event => event.expected), fields.map(field => new Just(goType(field))));
        assert.deepEqual(result.exprType, target);
        assert.ok(result.expr instanceof Go.GoConstructor);
        assert.doesNotMatch(printGoExpr(result.expr), /Rebox_|CoerceToStruct|RecordShape|\.Array\(/);
    }
});

test('annotated Tuple construction respects explicit boxed and scalar sum consumer representations', () => {
    const sum = new Go.TypeStructValue('Data.Tuple.Tuple', [Go.TypeValue.value, Go.TypeValue.value]);
    for (const expected of [Go.TypeValue.value, sum]) {
        const { result, constructorContext, events } = annotatedConstruction([C.Int.value, C.Int.value], new Just(expected));
        assert.deepEqual(constructorContext.mbExpectedGoType, new Just(expected));
        assert.deepEqual(events.map(event => event.expected), [new Just(Go.TypeValue.value), new Just(Go.TypeValue.value)]);
        assert.deepEqual(result.exprType, sum, 'the existing native scalar sum ABI remains available');
    }
});

test('explicit native Tuple consumers retain record, list and array field representations', () => {
    for (const fields of [[position, list], [list, new C.Array(C.Int.value)]]) {
        const type = tuple(fields);
        const target = goType(type);
        const prepared = prepare(type, new Just(target));
        for (let i = 0; i < fields.length; i++) {
            const field = Adt.saturatedFieldType(prepared)(i);
            assert.deepEqual(field.goType, goType(fields[i]));
            assert.deepEqual(field.exprType, new C.TypeVar(i === 0 ? 'a' : 'b'),
                'source declaration annotations must not be replaced by representation choices');
        }
        const result = construct(prepared);
        assert.ok(result.expr instanceof Go.GoConstructor);
        assert.deepEqual(result.exprType, target);
        assert.doesNotMatch(printGoExpr(result.expr), /Rebox_|RecordShape|CoerceToStruct|\.Array\(/);
    }
});

test('default, boxed and struct-value Tuple consumers keep their existing ABI', () => {
    const type = tuple([new C.Array(C.Int.value), new C.Array(C.Int.value)]);
    const sum = new Go.TypeStructValue('Data.Tuple.Tuple', [Go.TypeValue.value, Go.TypeValue.value]);
    for (const expected of [Nothing.value, new Just(Go.TypeValue.value), new Just(sum)]) {
        const prepared = prepare(type, expected);
        assert.deepEqual(Adt.saturatedFieldType(prepared)(0).goType, Go.TypeValue.value);
        assert.deepEqual(Adt.saturatedFieldType(prepared)(1).goType, Go.TypeValue.value);
        assert.deepEqual(construct(prepared).exprType, sum);
    }
    assert.ok(construct(prepare(tuple([C.Int.value, C.Int.value]), Nothing.value)).exprType instanceof Go.TypeStructValue);
});

test('only a matching explicit Tuple pointer selects native construction', () => {
    const type = tuple([position, new C.Array(C.Int.value)]);
    const target = goType(type).value0;
    const invalid = [
        { ...target, baseStructName: 'Data_Other_Pair' },
        { ...target, fullPath: 'Constructor_Other_Pair' },
        { ...target, symbol: { ...target.symbol, owner: 'generated/Other' } },
        { ...target, typeArgs: [target.typeArgs[0]] },
    ];
    for (const pointer of invalid) {
        assert.ok(construct(prepare(type, new Just(new Go.TypeStructPointer(pointer)))).exprType instanceof Go.TypeStructValue);
    }
});

test('native Tuple construction preserves values, immutable aliases and the boxed envelope', t => {
    const positionList = tuple([position, list]);
    const listArray = tuple([list, new C.Array(C.Int.value)]);
    const first = construct(prepare(positionList, new Just(goType(positionList))));
    const second = construct(prepare(listArray, new Just(goType(listArray))));
    const boxed = construct(prepare(positionList, new Just(Go.TypeValue.value)));
    const boxedExpression = boxGoExpr(state())('NativeTuples')(boxed.expr)(boxed.exprType);
    const directory = goModule(t, 'gopurs-native-tuple-construction-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'construction_test.go', `package main
import ("testing"; "unsafe"; r "gopurs/output/gopurs_runtime")
type Constructor_Data_Tuple_Tuple[A, B any] struct { Rc uint32; V0 A; V1 B }
type Constructor_Data_List_Types_Cons[A any] struct { Rc uint32; V0 A; V1 *Constructor_Data_List_Types_Cons[A] }
func positionList(left ${Go.goTypeToStr(goType(position))}, right ${Go.goTypeToStr(goType(list))}) ${Go.goTypeToStr(first.exprType)} {
    return ${printGoExpr(first.expr).replaceAll('gopurs_runtime.', 'r.')}
}
func listArray(left ${Go.goTypeToStr(goType(list))}, right []int64) ${Go.goTypeToStr(second.exprType)} {
    return ${printGoExpr(second.expr).replaceAll('gopurs_runtime.', 'r.')}
}
func boxedPair(left r.Value, right r.Value) r.Value {
    return ${printGoExpr(boxedExpression).replaceAll('gopurs_runtime.', 'r.')}
}
func TestNativeAndBoxed(t *testing.T) {
    pos := ${Go.goTypeToStr(goType(position))}{column: 3, line: 8}
    tail := &Constructor_Data_List_Types_Cons[int64]{V0: 9}
    values := []int64{4, 7}
    first := positionList(pos, tail)
    second := listArray(tail, values)
    if first.V0 != pos || first.V1 != tail || second.V0 != tail || &second.V1[0] != &values[0] { t.Fatal("native field copy or alias changed") }
    if first == positionList(pos, tail) { t.Fatal("Tuple allocation unexpectedly memoized") }
    if positionList(pos, nil).V1 != nil || listArray(nil, nil).V1 != nil { t.Fatal("nil field changed") }
    generic := &Constructor_Data_List_Types_Cons[r.Value]{V0: r.Int(9)}
    l := r.RecordDict2("column", "line", r.Int(3), r.Int(8))
    rr := r.Value{Type: 9, IntVal: ${hashString('Data_Data_List_Types_Cons')}, UnsafePtr: unsafe.Pointer(generic)}
    boxed := boxedPair(l, rr)
    if boxed.Type != 9 || boxed.IntVal != ${hashString('Data_Data_Tuple_Tuple')} { t.Fatal("Tuple tag changed") }
    fields := (*Constructor_Data_Tuple_Tuple[r.Value,r.Value])(boxed.UnsafePtr)
    if fields.V0 != l || fields.V1 != rr || generic.V0.IntVal != 9 || values[0] != 4 || pos.line != 8 { t.Fatal("boxed payload or inputs changed") }
}
`);
    runGo(directory, ['test', '-count=1', '.']);
});


test('native constructors retain their payloads and sum tags across worker result conversions', t => {
    const convert = (module, name, types, args, slots) => {
        const base = `Data_Data_${module}_${name}`;
        const symbol = { owner: `generated/Data_${module}`, name: `Constructor_Data_${module}_${name}` };
        const source = Go.structPointer({ baseStructName: base, fullName: `Data.${module}.${module}`,
            structName: symbol.name, symbol })(types);
        const expression = new Go.GoConstructor(hashString(base), symbol, types, args.map(Go.rawGo));
        const target = new Go.TypeStructValue(`Data.${module}.${module}`, slots);
        const result = coerceGoExpr(state())('NativeTuples')(expression)(source)(target);
        assert.ok(result instanceof Go.GoStructValue);
        return { expression: printGoExpr(result).replaceAll('gopurs_runtime.', 'r.'),
            type: Go.goTypeToStr(target).replaceAll('gopurs_runtime.', 'r.') };
    };
    const scalar = convert('Tuple', 'Tuple', [Go.TypeString.value, Go.TypeInt64.value],
        ['nextString()', 'nextInt()'], [Go.TypeValue.value, Go.TypeValue.value]);
    assert.match(scalar.expression, /nextString\(\)/);
    assert.match(scalar.expression, /nextInt\(\)/);
    const composite = convert('Tuple', 'Tuple', [goType(position), new Go.TypeNativeArray(Go.TypeInt64.value)],
        ['nextPosition()', 'nextArray()'], [goType(position), Go.TypeValue.value]);
    const just = convert('Maybe', 'Just', [Go.TypeInt64.value], ['nextInt()'],
        [Go.TypeValue.value, Go.TypeBool.value]);
    const left = convert('Either', 'Left', [Go.TypeInt64.value, Go.TypeString.value], ['nextInt()'],
        [Go.TypeValue.value, Go.TypeValue.value, Go.TypeBool.value]);
    const right = convert('Either', 'Right', [Go.TypeInt64.value, Go.TypeString.value], ['nextString()'],
        [Go.TypeValue.value, Go.TypeValue.value, Go.TypeBool.value]);
    const directory = goModule(t, 'gopurs-native-constructor-conversion-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'conversion_test.go', `package main
import ("testing"; "unsafe"; r "gopurs/output/gopurs_runtime")
var order string
func nextString() string { order += "s"; return "kept" }
func nextInt() int64 { order += "i"; return 42 }
func nextPosition() ${Go.goTypeToStr(goType(position))} {
    order += "p"
    return ${Go.goTypeToStr(goType(position))}{column: 3, line: 8}
}
func nextArray() []int64 { order += "a"; return []int64{5, 9} }
func scalarWorker() ${scalar.type} { return ${scalar.expression} }
func compositeWorker() ${composite.type} { return ${composite.expression} }
func justWorker() ${just.type} { return ${just.expression} }
func leftWorker() ${left.type} { return ${left.expression} }
func rightWorker() ${right.type} { return ${right.expression} }
func TestWorkerResults(t *testing.T) {
    scalar := scalarWorker()
    if scalar.V0.StrVal() != "kept" || scalar.V1.IntVal != 42 || order != "si" {
        t.Fatal("scalar payload or evaluation order lost", scalar, order)
    }
    order = ""
    composite := compositeWorker()
    values := *(*[]r.Value)(unsafe.Pointer(composite.V1.UnsafePtr))
    if composite.V0.column != 3 || composite.V0.line != 8 || len(values) != 2 || values[0].IntVal != 5 || values[1].IntVal != 9 || order != "pa" {
        t.Fatal("composite payload or evaluation order lost", composite, values, order)
    }
    order = ""
    just := justWorker()
    left := leftWorker()
    right := rightWorker()
    if !just.V1 || just.V0.IntVal != 42 || left.V2 || left.V0.IntVal != 42 || !right.V2 || right.V1.StrVal() != "kept" || order != "iis" {
        t.Fatal("sum payload, discriminator or evaluation count lost", just, left, right, order)
    }
}
`);
    runGo(directory, ['test', '-count=1', '.']);
});
