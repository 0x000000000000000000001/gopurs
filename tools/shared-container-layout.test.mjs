import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, singleton } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { buildPointerAdtMetadata, buildEnumAdtMetadata } from '../output/Gopurs.AdtMetadata/index.js';
import { buildConstructorTypes, collectElidedConstructors } from '../output/Gopurs.ConstructorMetadata/index.js';
import * as Layout from '../output/Gopurs.ConstructorLayout/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { exprTypeToGoType, exprTypeToGenericGoType } from '../output/Gopurs.GoTypes/index.js';
import { canonicalContainerArguments } from '../output/Gopurs.SharedContainerLayout/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const adt = (name, args = []) => new C.ADT(name, name.split('.'), args);
const variable = name => new C.TypeVar(name);
const row = (fields, tail = Nothing.value) => new C.Record(new C.Row(
    fields.map(([name, type]) => new Tuple(name, type)), tail));
const pair = (record, enumType = adt('Acme.Phase')) => adt('Data.Tuple.Tuple', [record, enumType]);
const list = element => adt('Data.List.Types.List', [element]);
const modules = [
    { name: 'Data.Tuple', dataDecls: [{ name: 'Tuple', vars: ['a', 'b'], constructors: [
        { name: 'Tuple', fields: [variable('a'), variable('b')] },
    ] }] },
    { name: 'Data.List.Types', dataDecls: [{ name: 'List', vars: ['a'], constructors: [
        { name: 'Nil', fields: [] },
        { name: 'Cons', fields: [variable('a'), list(variable('a'))] },
    ] }] },
    { name: 'Acme', dataDecls: [
        { name: 'Phase', vars: [], constructors: [
            { name: 'Start', fields: [] }, { name: 'Middle', fields: [] }, { name: 'End', fields: [] },
        ] },
        { name: 'Payload', vars: [], constructors: [
            { name: 'Small', fields: [C.Int.value] }, { name: 'Large', fields: [C.String.value] },
        ] },
    ] },
    { name: 'Vendor', dataDecls: [{ name: 'Status', vars: [], constructors: [
        { name: 'Active', fields: [] }, { name: 'Inactive', fields: [] },
    ] }] },
].map(module => ({ ...module, classDecls: [] }));
const baseMetadata = withReboxFields({
    ctorTypes: buildConstructorTypes(modules),
    elidedCtors: collectElidedConstructors(modules),
    ...buildPointerAdtMetadata(modules), ...buildEnumAdtMetadata(modules),
    classDeclsFields: emptyMap, globalTypes: emptyMap, globalFunctions: emptyMap,
});
const valueType = (type, meta = metadata) => exprTypeToGoType(meta.sharedRecords)(meta.pointerAdtPaths)(meta.enumAdts)(meta.elidedCtors)('Consumer')(type);
const genericType = (type, variables = [], meta = metadata) => exprTypeToGenericGoType(meta.sharedRecords)(meta.pointerAdtPaths)(meta.enumAdts)(meta.elidedCtors)(variables)('Consumer')(type);
const prepare = (form, name, type, meta = metadata) => Layout.prepare(form)(meta)('Consumer')('Fallback')(name)(type);
const argsOf = type => type.value0.typeArgs;
const canonical = (name, source, mapped, meta = metadata) => canonicalContainerArguments(meta.pointerAdtPaths)(meta.enumAdts)(meta.elidedCtors)(name)(source)(mapped);
const position = row([['offset', C.Int.value], ['row', C.Int.value]]);
const registeredRecords = [position, row([['label', C.String.value]]),
    row([['enabled', C.Boolean.value], ['weight', C.Number.value], ['initial', C.Char.value]])];
const registeredPairs = registeredRecords.flatMap(record => [adt('Acme.Phase'), adt('Vendor.Status')].map(tag => [record, tag]));
const withRegistry = (meta, pairs) => ({
    ...meta,
    pointerAdtPaths: ['Data.Tuple.Tuple', 'Data.List.Types.List'].reduce((paths, name) => {
        const entry = lookup(ordString)(name)(paths);
        return entry instanceof Just ? insert(ordString)(name)({ ...entry.value0, sharedPairs: pairs })(paths) : paths;
    }, meta.pointerAdtPaths),
});
const metadata = withRegistry(baseMetadata, registeredPairs);

const mappedPair = [new Go.TypeRecord([new Tuple('offset', Go.TypeInt64.value), new Tuple('row', Go.TypeInt64.value)]), Go.TypeUint32.value];

test('closed scalar records and proven enums share the container ABI independently of application names', () => {
    for (const record of registeredRecords) {
        for (const enumType of [adt('Acme.Phase'), adt('Vendor.Status')]) {
            assert.deepEqual(argsOf(valueType(pair(record, enumType))), [Go.TypeValue.value, Go.TypeValue.value]);
            assert.deepEqual(argsOf(valueType(list(pair(record, enumType)))), [Go.TypeValue.value]);
            assert.deepEqual(argsOf(genericType(pair(record, enumType))), [Go.TypeValue.value, Go.TypeValue.value]);
            assert.deepEqual(argsOf(genericType(list(pair(record, enumType)))), [Go.TypeValue.value]);
        }
    }
    assert.deepEqual(valueType(position), mappedPair[0], 'the standalone record keeps its native fields');
    assert.equal(valueType(adt('Acme.Phase')), Go.TypeUint32.value, 'standalone enums stay native');
    const reordered = row([['row', C.Int.value], ['offset', C.Int.value]]);
    assert.deepEqual(argsOf(valueType(pair(reordered))), [Go.TypeValue.value, Go.TypeValue.value],
        'record label order does not change the selected family');
});

test('an unselected pair retains native arguments even when its structure is eligible', () => {
    const native = withRegistry(metadata, []);
    assert.equal(canonical('Data.Tuple.Tuple', [position, adt('Acme.Phase')], mappedPair, native), mappedPair);
    assert.deepEqual(argsOf(valueType(pair(position), native)), mappedPair);
    const listArgument = argsOf(valueType(list(pair(position)), native))[0];
    assert.deepEqual(argsOf(listArgument), mappedPair);
    const unrelated = withRegistry(metadata, [[position, adt('Vendor.Status')]]);
    assert.deepEqual(argsOf(valueType(pair(position), unrelated)), mappedPair);
});

test('constructors and value signatures agree in ordinary and nested TypeApp forms', () => {
    const tuple = pair(position);
    const stack = list(tuple);
    for (const [type, ctor, expected] of [
        [tuple, 'Tuple', [Go.TypeValue.value, Go.TypeValue.value]],
        [stack, 'Cons', [Go.TypeValue.value]],
        [stack, 'Nil', [Go.TypeValue.value]],
    ]) {
        for (const form of [Layout.Definition.value, Layout.Saturated.value]) {
            const result = prepare(form, ctor, type);
            assert.deepEqual(result.typeArgs, expected);
            assert.deepEqual(result.fieldTypeArgs, expected);
            assert.deepEqual(result.pointerType.value0.typeArgs, expected);
        }
    }
    const appliedPair = new C.TypeApp(new C.TypeApp(adt('Data.Tuple.Tuple'), [position]), [adt('Acme.Phase')]);
    const appliedList = new C.TypeApp(adt('Data.List.Types.List'), [appliedPair]);
    assert.deepEqual(argsOf(valueType(appliedPair)), [Go.TypeValue.value, Go.TypeValue.value]);
    assert.deepEqual(argsOf(genericType(appliedList)), [Go.TypeValue.value]);
    assert.deepEqual(prepare(Layout.Definition.value, 'Tuple', appliedPair).typeArgs,
        [Go.TypeValue.value, Go.TypeValue.value]);
    assert.deepEqual(prepare(Layout.Definition.value, 'Cons', appliedList).typeArgs, [Go.TypeValue.value]);
});

test('open, unknown and composite record payloads retain their ordinary representations', () => {
    const rejected = [
        row([]),
        row([['offset', C.Int.value]], new Just(variable('tail'))),
        row([['offset', C.Int.value]], new Just(C.Any.value)),
        row([['offset', C.Any.value]]),
        row([['offset', variable('a')]]),
        row([['items', new C.Array(C.Int.value)]]),
        row([['nested', position]]),
        row([['callback', new C.Func([C.Int.value], C.Int.value)]]),
    ];
    for (const record of rejected) {
        const registered = withRegistry(metadata, [[record, adt('Acme.Phase')]]);
        assert.equal(canonical('Data.Tuple.Tuple', [record, adt('Acme.Phase')], mappedPair, registered), mappedPair);
        const mappedList = [valueType(pair(record), registered)];
        assert.equal(canonical('Data.List.Types.List', [pair(record)], mappedList, registered), mappedList);
    }
    for (const enumType of [variable('a'), C.Any.value, adt('Acme.Payload'), adt('Unknown.Phase'), adt('Acme.Phase', [C.Int.value])]) {
        assert.equal(canonical('Data.Tuple.Tuple', [position, enumType], mappedPair, withRegistry(metadata, [[position, enumType]])), mappedPair);
    }
    assert.equal(canonical('Other.Tuple', [position, adt('Acme.Phase')], mappedPair), mappedPair);
    assert.equal(canonical('Other.List', [pair(position)], [Go.TypeInt64.value])[0], Go.TypeInt64.value);
});

test('arity, metadata and constructor elision remain admission boundaries', () => {
    const args = [position, adt('Acme.Phase')];
    for (const supplied of [[], [position], [...args, C.Int.value]]) {
        assert.equal(canonical('Data.Tuple.Tuple', supplied, mappedPair), mappedPair);
    }
    const incomplete = [Go.TypeInt64.value];
    const excessive = [...mappedPair, Go.TypeString.value];
    assert.equal(canonical('Data.Tuple.Tuple', args, incomplete), incomplete);
    assert.equal(canonical('Data.Tuple.Tuple', args, excessive), excessive);
    const missingPointers = { ...metadata, pointerAdtPaths: emptyMap };
    assert.equal(canonical('Data.Tuple.Tuple', args, mappedPair, missingPointers), mappedPair);
    const wrongArity = { ...metadata, pointerAdtPaths: insert(ordString)('Data.Tuple.Tuple')({ ctorName: 'Tuple', arity: 1 })(metadata.pointerAdtPaths) };
    assert.equal(canonical('Data.Tuple.Tuple', args, mappedPair, wrongArity), mappedPair);
    assert.equal(canonical('Data.Tuple.Tuple', args, mappedPair, { ...metadata, enumAdts: emptySet }), mappedPair);
    const historicalElision = collectElidedConstructors([{
        name: 'Data.Tuple', dataDecls: [{ name: 'Tuple', vars: ['a', 'b'], constructors: [
            { name: 'Tuple', fields: [variable('a')] },
        ] }], classDecls: [],
    }]);
    const normalizedElision = singleton(Go.constructorNames('Data_Tuple')('Tuple').structName);
    for (const elidedCtors of [historicalElision, normalizedElision]) {
        assert.equal(canonical('Data.Tuple.Tuple', args, mappedPair, { ...metadata, elidedCtors }), mappedPair);
    }
    const generic = pair(variable('a'), variable('b'));
    assert.deepEqual(argsOf(genericType(generic, ['a', 'b'])), [new Go.TypeGenericParam('a'), new Go.TypeGenericParam('b')]);
    assert.deepEqual(argsOf(valueType(adt('Data.Tuple.Tuple', [C.Int.value, C.String.value]))), [Go.TypeInt64.value, Go.TypeString.value]);
    assert.deepEqual(argsOf(valueType(list(C.Int.value))), [Go.TypeInt64.value]);
});
