import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, toUnfoldable } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import { buildPointerAdtMetadata, buildEnumAdtMetadata } from '../output/Gopurs.AdtMetadata/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { exprTypeToGoType, exprTypeToGenericGoType } from '../output/Gopurs.GoTypes/index.js';
import { collectSharedRecordShapes, isSharedRecord } from '../output/Gopurs.SharedRecordLayout/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const variable = name => new C.TypeVar(name);
const adt = (name, args = []) => new C.ADT(name, name.split('.'), args);
const record = (fields, tail = Nothing.value) => new C.Record(new C.Row(fields.map(([name, type]) => new Tuple(name, type)), tail));
const fieldsOf = type => type.value0.value0;
const fields = type => fieldsOf(type).map(({ value0, value1 }) => [value0, value1]);
const payload = argument => adt('Library.Payload', [argument]);
const phase = adt('Library.Phase');
const child = (argument = variable('a')) => record([
    ['contents', payload(argument)], ['phase', phase], ['count', C.Int.value],
]);
const nested = (argument = variable('a')) => record([
    ['ready', C.Boolean.value], ['part', child(argument)], ['contents', payload(argument)],
]);
const ann = type => ({ span: C.emptySpan, meta: Nothing.value, sourceUsage: Nothing.value, type: new Just(type) });
const ref = (type, name = 'value') => new C.ExprVar(ann(type), new C.Qualified(new Just('External'), name));
const binding = (name, type, expression = ref(type)) => new C.NonRec(new C.Binding(ann(type), name, expression));
const moduleOf = (name, decls = [], dataDecls = [], foreign = emptyMap) => ({ name, decls, dataDecls, foreign, classDecls: [] });
const declarations = [moduleOf('Library', [], [
    { name: 'Payload', vars: ['a'], constructors: [
        { name: 'Value', fields: [variable('a')] }, { name: 'Text', fields: [C.String.value] },
    ] },
    { name: 'Phase', vars: [], constructors: [{ name: 'Start', fields: [] }, { name: 'End', fields: [] }] },
    { name: 'Cell', vars: ['a'], constructors: [{ name: 'Cell', fields: [variable('a')] }] },
])];
const pointers = buildPointerAdtMetadata(declarations).pointerAdtPaths;
const enums = buildEnumAdtMetadata(declarations).enumAdts;
const callback = type => new C.Func([type], type);
const collect = modules => collectSharedRecordShapes(pointers)(enums)([...declarations, ...modules]);
const select = type => collect([moduleOf('Caller', [binding('stored', callback(type))])]);
const selected = (registry, type) => isSharedRecord(registry)(fieldsOf(type));
const valueType = (registry, type) => exprTypeToGoType(registry)(pointers)(enums)(emptySet)('Caller')(type);
const genericType = (registry, type) => exprTypeToGenericGoType(registry)(pointers)(enums)(emptySet)(['a'])('Caller')(type);

test('a function-value boundary selects the outer record while its flat child stays native', () => {
    const type = nested();
    const registry = select(type);
    assert.equal(toUnfoldable(unfoldableArray)(registry).length, 1);
    assert.equal(valueType(registry, type), Go.TypeValue.value);
    assert.equal(genericType(registry, type), Go.TypeValue.value);
    assert.ok(valueType(registry, child()) instanceof Go.TypeRecord);
    assert.ok(valueType(emptySet, type) instanceof Go.TypeRecord);
});

test('direct monomorphic calls do not select a record, but passing that function does', () => {
    const type = nested(C.Int.value);
    const fn = callback(type);
    const directWorker = binding('worker', fn, new C.ExprAbs(ann(fn), 'input', ref(type, 'input')));
    const directCall = new C.ExprApp(ann(type), ref(fn, 'worker'), ref(type, 'input'));
    const direct = collect([moduleOf('Caller', [directWorker, binding('result', type, directCall)])]);
    assert.equal(selected(direct, type), false);
    const pass = new C.ExprApp(ann(C.Int.value), ref(new C.Func([fn], C.Int.value), 'consume'), ref(fn, 'worker'));
    const indirect = collect([moduleOf('Caller', [directWorker, binding('result', C.Int.value, pass)])]);
    assert.equal(selected(indirect, type), true);
});

test('boxed ADT instantiations share a key while identity and complete arity remain significant', () => {
    const registry = select(nested(variable('a$scope123')));
    for (const argument of [variable('b'), C.Int.value, C.Any.value, record([['local', C.String.value]])]) {
        assert.equal(selected(registry, nested(argument)), true);
    }
    const applied = new C.TypeApp(adt('Library.Payload'), [C.String.value]);
    const appliedChild = record([['contents', applied], ['phase', phase], ['count', C.Int.value]]);
    const appliedOuter = record([['ready', C.Boolean.value], ['part', appliedChild], ['contents', applied]]);
    assert.equal(selected(registry, appliedOuter), true);
    for (const bad of [adt('Library.Payload'), adt('Library.Payload', [C.Int.value, C.String.value]), adt('Missing.Payload', [C.Int.value])]) {
        const type = record([['part', record([['contents', bad]])]]);
        assert.equal(selected(select(type), type), false);
        const altered = record([['ready', C.Boolean.value], ['part', child()], ['contents', bad]]);
        assert.equal(selected(registry, altered), false);
    }
});

test('open rows, bare unknowns, arrays, functions and deeper native records are not selected', () => {
    const base = nested();
    const registry = select(base);
    for (const extra of [C.Any.value, variable('a'), new C.Array(C.Int.value), new C.Func([C.Int.value], C.Int.value)]) {
        const type = record([...fields(base), ['extra', extra]]);
        assert.equal(selected(select(type), type), false);
    }
    const openChild = record([['contents', payload(C.Int.value)]], new Just(variable('row')));
    const deepChild = record([['contents', payload(C.Int.value)], ['deeper', record([['scalar', C.Int.value]])]]);
    for (const part of [openChild, deepChild, record([])]) {
        const type = record([['part', part], ['ready', C.Boolean.value]]);
        assert.equal(selected(select(type), type), false);
    }
    for (const tail of [new Just(variable('row')), new Just(C.Any.value)]) {
        const open = record(fields(base), tail);
        assert.ok(valueType(registry, open) === Go.TypeValue.value, 'open rows keep the existing generic ABI');
        assert.equal(toUnfoldable(unfoldableArray)(select(open)).length, 0);
    }
});

test('scalar and enum-only children and native pointer payloads retain their normal layout', () => {
    for (const element of [C.Int.value, phase, adt('Library.Cell', [C.Int.value])]) {
        const type = record([['part', record([['contents', element]])]]);
        assert.equal(selected(select(type), type), false);
    }
    const malformedEnum = record([['part', record([['contents', payload(C.Int.value)], ['phase', adt('Library.Phase', [C.Int.value])]])]]);
    assert.equal(selected(select(malformedEnum), malformedEnum), false);
});

test('selection has no application labels and keys preserve label identity, order and visible duplicates', () => {
    const original = nested();
    const reordered = record(fields(original).reverse().map(([name, type]) => [name, name === 'part' ? record(fields(type).reverse()) : type]));
    const duplicate = record([...fields(reordered), ['contents', C.Any.value]]);
    const registry = select(original);
    assert.equal(selected(registry, reordered), true);
    assert.equal(selected(registry, duplicate), true);
    assert.deepEqual(select(duplicate), registry);
    const hiddenValid = record([['contents', C.Any.value], ...fields(original)]);
    assert.equal(selected(select(hiddenValid), hiddenValid), false);
    const renamed = record([['🧩:part;', record([['x:y', payload(C.Int.value)], ['β', C.Int.value]])], ['other', C.Boolean.value]]);
    assert.equal(selected(select(renamed), renamed), true);
    assert.equal(selected(registry, renamed), false);
    const distinct = record([['🧩:part', record([[';x:y', payload(C.Int.value)], ['β', C.Int.value]])], ['other', C.Boolean.value]]);
    assert.equal(selected(select(renamed), distinct), false);
});

test('foreign callbacks and stored functions participate, and their removal retracts the selection', () => {
    const type = nested(C.Int.value);
    const fn = callback(type);
    const foreign = insert(ordString)('use')(new Just(new C.Func([fn], C.Int.value)))(emptyMap);
    assert.equal(selected(collect([moduleOf('Foreign', [], [], foreign)]), type), true);
    const stored = record([['handler', fn]]);
    assert.equal(selected(collect([moduleOf('Caller', [binding('saved', stored)])]), type), true);
    assert.equal(selected(collect([moduleOf('Caller', [binding('saved', C.Int.value)])]), type), false);
});

const directBody = type => {
    if (type instanceof C.ForAll || type instanceof C.ConstrainedType) return directBody(type.value1);
    if (type instanceof C.Func && type.value0.length > 0) {
        const rest = type.value0.length === 1 ? type.value1 : new C.Func(type.value0.slice(1), type.value1);
        return new C.ExprAbs(ann(type), 'argument', directBody(rest));
    }
    return ref(type, 'result');
};
const recursive = (name, type) => new C.Rec([new C.Binding(ann(type), name, directBody(type))]);

test('a scalar-controlled recursive worker vetoes the shared family for reductions and returned accumulators', () => {
    const type = nested(C.Int.value);
    const stored = moduleOf('Callback', [binding('rare', callback(type))]);
    assert.equal(selected(collect([stored]), type), true);
    for (const scalar of [C.Int.value, C.Number.value, C.Boolean.value, C.String.value, C.Char.value]) {
        for (const result of [C.Int.value, type]) {
            const signature = new C.ForAll(['a'], new C.ConstrainedType([], new C.Func([scalar], new C.Func([type], result))));
            const worker = moduleOf('Worker', [recursive('loop', signature)]);
            for (const modules of [[stored, worker], [worker, stored]]) {
                const registry = collect(modules);
                assert.equal(selected(registry, type), false);
                assert.ok(valueType(registry, type) instanceof Go.TypeRecord);
            }
        }
    }
});

test('local recursive groups also protect native record accumulators', () => {
    const type = nested(C.String.value);
    const signature = new C.Func([C.Int.value, type], type);
    const expression = new C.ExprLet(ann(type), [recursive('loop', signature)], ref(type));
    const registry = collect([moduleOf('Caller', [binding('rare', callback(type)), binding('result', type, expression)])]);
    assert.equal(selected(registry, type), false);
});

test('recursive record and ADT parameters without a direct scalar keep the shared representation', () => {
    const type = nested(C.Int.value);
    const signature = new C.Func([type, adt('Library.Cell', [C.Int.value])], type);
    const registry = collect([moduleOf('Caller', [binding('stored', callback(type)), recursive('realign', signature)])]);
    assert.equal(selected(registry, type), true);
    assert.equal(valueType(registry, type), Go.TypeValue.value);
});

test('the loop guard does not inspect records hidden inside arrays or ADT arguments', () => {
    const type = nested(C.Int.value);
    for (const container of [new C.Array(type), adt('Library.Cell', [type])]) {
        const signature = new C.Func([C.Int.value, container], C.Int.value);
        const registry = collect([moduleOf('Caller', [binding('stored', callback(type)), recursive('loop', signature)])]);
        assert.equal(selected(registry, type), true);
    }
    const direct = new C.Func([C.Int.value, type], type);
    const registry = collect([moduleOf('Caller', [binding('stored', callback(type)), binding('once', direct, directBody(direct))])]);
    assert.equal(selected(registry, type), true, 'nonrecursive direct workers do not trigger the loop guard');
});
