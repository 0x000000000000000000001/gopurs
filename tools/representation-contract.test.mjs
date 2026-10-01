import assert from 'node:assert/strict';
import test from 'node:test';
import { empty, insert, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { member, singleton } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { buildPointerAdtMetadata, buildEnumAdtMetadata } from '../output/Gopurs.AdtMetadata/index.js';
import { addClassDataDeclarations, buildClassFields } from '../output/Gopurs.ClassMetadata/index.js';
import { buildConstructorTypes, collectElidedConstructors } from '../output/Gopurs.ConstructorMetadata/index.js';
import * as Layout from '../output/Gopurs.ConstructorLayout/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { exprTypeToGoType, exprTypeToGenericGoType } from '../output/Gopurs.GoTypes/index.js';
import { constructors } from '../output/Gopurs.ModuleDeclarations/index.js';
import { printGoDecl } from '../output/Gopurs.Printer/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const a = new C.TypeVar('a');
const original = {
    name: 'Fixture.Model',
    dataDecls: [
        { name: 'Pair', vars: ['a', 'b'], constructors: [{ name: 'Pair', fields: [a, new C.TypeVar('b')] }] },
        { name: 'Wrapped', vars: ['a'], constructors: [{ name: 'Payload', fields: [a] }] },
        { name: 'Choice', vars: [], constructors: [{ name: 'Left', fields: [] }, { name: 'Right', fields: [] }] },
        { name: 'Optional', vars: ['a'], constructors: [{ name: 'Some', fields: [a] }, { name: 'None', fields: [] }] },
    ],
    classDecls: [
        { name: 'C', vars: ['a'], superclasses: [new Tuple(['Base', 'Ord'], [a])],
            methods: [new Tuple('z', C.Boolean.value), new Tuple('alpha', a)] },
        { name: 'Single', vars: ['a'], superclasses: [], methods: [new Tuple('only', a)] },
        { name: 'Empty', vars: [], superclasses: [], methods: [] },
    ],
};
const enriched = addClassDataDeclarations(original);
const metadata = withReboxFields({
    ctorTypes: buildConstructorTypes([original]),
    elidedCtors: collectElidedConstructors([original]),
    classDeclsFields: buildClassFields([original]),
    ...buildPointerAdtMetadata([enriched]),
    ...buildEnumAdtMetadata([original]),
    globalTypes: empty, globalFunctions: empty,
});
const get = (map, key) => lookup(ordString)(key)(map);
const adt = (name, args = []) => new C.ADT(`Fixture.Model.${name}`, ['Fixture', 'Model', name], args);
const valueType = (type, meta = metadata) => exprTypeToGoType(meta.pointerAdtPaths)(meta.enumAdts)(meta.elidedCtors)('Consumer')(type);
const genericType = (type, vars, meta = metadata) => exprTypeToGenericGoType(meta.pointerAdtPaths)(meta.enumAdts)(meta.elidedCtors)(vars)('Consumer')(type);
const prepare = (form, name, type) => Layout.prepare(form)(metadata)('Consumer')('Fallback')(name)(type);
const argsOf = type => type.value0.typeArgs.map(Go.goTypeToStr);
const preparedArgs = prepared => prepared.typeArgs.map(Go.goTypeToStr);
const boxed = 'gopurs_runtime.Value';

test('original constructor metadata and enriched ADT metadata retain their distinct identities', () => {
    assert.ok(get(metadata.ctorTypes, 'Fixture_Model.Pair') instanceof Just);
    assert.ok(get(metadata.ctorTypes, 'Fixture.Model.Pair') instanceof Nothing);
    assert.ok(get(metadata.ctorTypes, 'Fixture_Model.Single') instanceof Nothing,
        'synthetic dictionaries must not enter the original constructor table');
    assert.equal(member(ordString)('Constructor_Fixture.Model_Payload')(metadata.elidedCtors), true);
    assert.equal(member(ordString)('Constructor_Fixture.Model_Single')(metadata.elidedCtors), false);
    assert.equal(Go.goTypeToStr(valueType(adt('Single', [C.Int.value]))), '*Constructor_Fixture_Model_Single[int64]');
    assert.equal(Go.goTypeToStr(valueType(adt('Choice'))), 'uint32');
    assert.equal(member(ordString)('Data_Fixture_Model_Left')(metadata.enumCtors), true);
    assert.equal(member(ordString)('Fixture.Model.Empty')(metadata.enumAdts), false);
    assert.equal(member(ordString)('Data_Fixture_Model_Empty')(metadata.enumCtors), false);
    assert.equal(valueType(adt('Empty')), Go.TypeValue.value);
    assert.deepEqual(get(metadata.pointerAdtLeaves, 'Data_Fixture_Model_None'),
        new Just({ nodeBaseStruct: 'Data_Fixture_Model_Some', nodeCtor: 'Some' }));
});

test('dictionary declarations, getters and instantiated fields agree on superclass and method order', () => {
    const fields = get(metadata.classDeclsFields, 'Fixture.Model.C').value0.fields;
    assert.deepEqual(fields.map(field => field.name), ['Ord0', 'alpha', 'z']);
    const declarations = constructors(metadata)('Fixture_Model')(enriched);
    const index = declarations.findIndex(decl => decl instanceof Go.GoStructDecl && decl.value0.name === 'Constructor_Fixture_Model_C');
    const structure = declarations[index].value0;
    assert.deepEqual(structure.fields.map(field => [field.value0, Go.goTypeToStr(field.value1)]),
        [['Rc', 'uint32'], ['V0', boxed], ['V1', 'T_a'], ['V2', 'bool']]);
    const getter = printGoDecl(declarations[index + 1]);
    assert.match(getter, /case "Ord0": return gopurs_runtime.Box\(c.V0\)/);
    assert.match(getter, /case "alpha": return gopurs_runtime.Box\(c.V1\)/);
    assert.match(getter, /case "z": return gopurs_runtime.Box\(c.V2\)/);

    const prepared = prepare(Layout.Definition.value, 'C', adt('C', [C.Int.value]));
    assert.equal(prepared.layout.identity.moduleName, 'Fixture_Model');
    assert.equal(prepared.layout.identity.fullName, 'Fixture.Model.C');
    assert.ok(prepared.layout.constructorFields instanceof Nothing);
    assert.deepEqual(fields.map((_, index) => Go.goTypeToStr(
        Layout.fieldType(metadata)('Consumer')(prepared.layout.fields)(prepared.fieldTypeArgs)(index))),
    [boxed, 'int64', 'bool']);
});

test('partial annotations preserve the different value, generic-field and constructor arity contracts', () => {
    const cases = [
        { args: [], value: [boxed, boxed], generic: ['T_a', 'T_b'], constructor: [boxed, boxed] },
        { args: [C.Int.value], value: ['int64', boxed], generic: ['T_a', 'T_b'], constructor: [boxed, boxed] },
        { args: [C.Int.value, C.String.value], value: ['int64', 'string'], generic: ['int64', 'string'], constructor: ['int64', 'string'] },
        { args: [C.Int.value, C.String.value, C.Boolean.value], value: ['int64', 'string'], generic: ['T_a', 'T_b'], constructor: ['int64', 'string'] },
    ];
    for (const entry of cases) {
        const type = adt('Pair', entry.args);
        assert.deepEqual(argsOf(valueType(type)), entry.value);
        assert.deepEqual(argsOf(genericType(type, ['a', 'b'])), entry.generic);
        for (const form of [Layout.Definition.value, Layout.Saturated.value]) {
            const prepared = prepare(form, 'Pair', type);
            assert.deepEqual(preparedArgs(prepared), entry.constructor);
            assert.deepEqual(prepared.fieldTypeArgs, prepared.typeArgs);
        }
    }
    assert.deepEqual(argsOf(genericType(adt('Pair', [C.Int.value]), ['a'])), [boxed, boxed]);
    assert.equal(Go.goTypeToStr(genericType(new C.Array(a), ['a'])), `[]${boxed}`,
        'generic arrays retain the value representation of their elements');
});

test('nested applications keep inner argument order and the definition-only instantiation policy', () => {
    const applied = new C.TypeApp(new C.TypeApp(adt('Pair', [C.Int.value]), [C.String.value]), []);
    assert.deepEqual(argsOf(valueType(applied)), ['int64', 'string']);
    assert.deepEqual(argsOf(genericType(applied, ['a', 'b'])), ['int64', 'string']);
    assert.deepEqual(preparedArgs(prepare(Layout.Definition.value, 'Pair', applied)), ['int64', 'string']);
    assert.deepEqual(preparedArgs(prepare(Layout.Saturated.value, 'Pair', applied)), [boxed, boxed]);

    const incomplete = new C.TypeApp(adt('Pair'), [C.Int.value]);
    const definition = prepare(Layout.Definition.value, 'Pair', incomplete);
    assert.deepEqual(preparedArgs(definition), ['int64'], 'definition TypeApps do not pad the missing argument');
    assert.equal(Go.goTypeToStr(Layout.fieldType(metadata)('Consumer')(definition.layout.fields)(definition.fieldTypeArgs)(1)), boxed);
    assert.ok(definition.adtFullName instanceof Nothing, 'only a plain ADT publishes this constructor hint');
    assert.equal(valueType(new C.TypeApp(a, [C.Int.value])), Go.TypeValue.value);
    assert.equal(valueType(new C.ForAll(['a'], adt('Pair', [a, a]))), Go.TypeValue.value);
});

test('pointer lookup prefers the exact type to its dictionary alias and respects context-specific elision', () => {
    const alias = insert(ordString)('Fixture.Model.C$Dict')({ ctorName: 'Fallback', arity: 0 })(empty);
    const type = adt('C', [C.Int.value]);
    const fallback = { ...metadata, pointerAdtPaths: alias };
    assert.equal(Go.goTypeToStr(valueType(type, fallback)), '*Constructor_Fixture_Model_Fallback');
    assert.equal(Go.goTypeToStr(genericType(type, ['a'], fallback)), '*Constructor_Fixture_Model_Fallback');
    const exact = { ...metadata, pointerAdtPaths: insert(ordString)('Fixture.Model.C')({ ctorName: 'C', arity: 1 })(alias) };
    assert.equal(Go.goTypeToStr(valueType(type, exact)), '*Constructor_Fixture_Model_C[int64]');
    assert.equal(Go.goTypeToStr(genericType(type, ['a'], exact)), '*Constructor_Fixture_Model_C[int64]');

    const optional = adt('Optional', [C.Int.value]);
    const elidedPayload = { ...metadata, elidedCtors: singleton('Constructor_Fixture_Model_Some') };
    assert.equal(Go.goTypeToStr(valueType(optional, elidedPayload)), '*Constructor_Fixture_Model_Some[int64]');
    assert.equal(genericType(optional, ['a'], elidedPayload), Go.TypeValue.value);
    const elidedAnnotation = { ...metadata, elidedCtors: singleton('Constructor_Fixture_Model_Optional') };
    assert.equal(valueType(optional, elidedAnnotation), Go.TypeValue.value);
    assert.equal(Go.goTypeToStr(genericType(optional, ['a'], elidedAnnotation)), '*Constructor_Fixture_Model_Some[int64]');
});

test('imported nil alternatives keep the runtime identity while each construction form resolves its worker', () => {
    const type = adt('Optional', [C.Int.value]);
    const definition = prepare(Layout.Definition.value, 'None', type);
    const saturated = prepare(Layout.Saturated.value, 'None', type);
    assert.equal(Go.goTypeToStr(definition.leafPointerType.value0), '*Constructor_Consumer_Some[int64]');
    assert.equal(Go.goTypeToStr(saturated.leafPointerType.value0), '*Constructor_Fixture_Model_Some[int64]');
    for (const prepared of [definition, saturated]) {
        assert.equal(prepared.leafPointerType.value0.value0.baseStructName, 'Data_Fixture_Model_Some');
        assert.equal(prepared.leafPointerType.value0.value0.fullName, 'Fixture.Model.Optional');
        assert.equal(prepared.layout.identity.key, 'Fixture_Model.None');
    }
    const fieldAccess = Layout.layout(metadata)('Explicit.Module')('Pair')(Nothing.value);
    assert.equal(fieldAccess.identity.moduleName, 'Explicit_Module', 'field access keeps its explicit qualifier');
});
