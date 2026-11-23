import assert from 'node:assert/strict';
import test from 'node:test';
import { Cons, Nil } from '../output/Data.List.Types/index.js';
import { empty as emptyMap, lookup, toUnfoldable } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import { completeOrFail } from '../output/Gopurs.Cache.LookupTable/index.js';
import * as Wire from '../output/Gopurs.Cache.Wire/index.js';
import { sourceMetadata } from '../output/Gopurs.Driver.Prepare/index.js';
import * as Projection from '../output/Gopurs.Driver.Prepare.Projection/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const adt = (name, args = []) => new C.ADT(name, name.split('.'), args);
const variable = name => new C.TypeVar(name);
const position = new C.Record(new C.Row([
    new Tuple('offset', C.Int.value), new Tuple('row', C.Int.value),
], Nothing.value));
const phase = adt('Fixture.Phase');
const pair = adt('Data.Tuple.Tuple', [position, phase]);
const stack = adt('Data.List.Types.List', [pair]);
const callback = new C.Func([stack], C.Int.value);
const ann = type => ({ span: C.emptySpan, meta: Nothing.value, sourceUsage: Nothing.value, type: new Just(type) });
const ref = (type, name) => new C.ExprVar(ann(type), new C.Qualified(new Just('External'), name));
const integer = n => new C.ExprLit(ann(C.Int.value), new C.LitInt(n));
const binding = (name, type, expression) => new C.NonRec(new C.Binding(ann(type), name, expression));
const moduleOf = (name, dataDecls = [], decls = []) => ({ name, dataDecls, decls, classDecls: [], foreign: emptyMap });
const asList = modules => modules.reduceRight((rest, module) => new Cons(module, rest), Nil.value);
const declarations = [
    moduleOf('Data.Tuple', [{ name: 'Tuple', vars: ['a', 'b'], constructors: [
        { name: 'Tuple', fields: [variable('a'), variable('b')] },
    ] }]),
    moduleOf('Data.List.Types', [{ name: 'List', vars: ['a'], constructors: [
        { name: 'Nil', fields: [] },
        { name: 'Cons', fields: [variable('a'), adt('Data.List.Types.List', [variable('a')])] },
    ] }]),
    moduleOf('Fixture', [{ name: 'Phase', vars: [], constructors: [
        { name: 'Before', fields: [] }, { name: 'After', fields: [] },
    ] }]),
    moduleOf('External', [], [binding('handler', callback, new C.ExprAbs(ann(callback), 'input', integer(1)))]),
];
const modules = passesCallback => {
    const signature = new C.Func([C.Int.value], C.Int.value);
    const body = passesCallback
        ? new C.ExprApp(ann(C.Int.value), ref(new C.Func([C.Any.value], C.Int.value), 'consume'), ref(callback, 'handler'))
        : integer(0);
    return [...declarations, moduleOf('Caller', [], [binding('run', signature,
        new C.ExprAbs(ann(signature), 'seed', body))])];
};
const prepare = modules => sourceMetadata(modules)(asList(modules));
const selected = metadata => lookup(ordString)('Data.Tuple.Tuple')(metadata.pointerAdtPaths).value0.sharedPairs;
const registryCodec = Wire.codecArray(Wire.codecArray(Wire.codecExprType));

test('nonempty shared-container metadata survives projection fragments and portable type serialization', () => {
    const sources = modules(true);
    const metadata = prepare(sources);
    assert.deepEqual(selected(metadata), [[position, phase]]);
    const projection = {
        modules: asList(sources), interfaces: emptyMap, mainModules: ['Caller'],
        metadata: { ...metadata, globalTypes: completeOrFail(metadata.globalTypes) },
    };
    const fragments = Projection.splitMaps(projection);
    assert.equal(fragments.length, sources.length);
    const tupleFragments = fragments.filter(fragment =>
        lookup(ordString)('Data.Tuple.Tuple')(fragment.value.metadata.pointerAdtPaths) instanceof Just);
    assert.equal(tupleFragments.length, 1, 'the selected registry travels with exactly one pointer declaration');
    assert.deepEqual(selected(tupleFragments[0].value.metadata), [[position, phase]]);
    const merged = Projection.mergeMaps(fragments);
    assert.ok(merged instanceof Just);
    assert.deepEqual(merged.value0.modules, projection.modules);
    assert.deepEqual(merged.value0.mainModules, projection.mainModules);
    assert.deepEqual(toUnfoldable(unfoldableArray)(merged.value0.metadata.pointerAdtPaths),
        toUnfoldable(unfoldableArray)(projection.metadata.pointerAdtPaths));
    const encoded = Wire.write(registryCodec)(selected(merged.value0.metadata))();
    const decoded = Wire.read(registryCodec)(encoded);
    assert.ok(decoded instanceof Just);
    assert.deepEqual(decoded.value0, [[position, phase]]);
    assert.equal(Wire.write(registryCodec)(decoded.value0)(), encoded);
});

test('a body-only removal of the callback use retracts selection without changing declared interfaces', () => {
    const beforeSources = modules(true), afterSources = modules(false);
    assert.deepEqual(beforeSources.map(module => module.dataDecls), afterSources.map(module => module.dataDecls));
    const before = prepare(beforeSources), after = prepare(afterSources);
    assert.deepEqual(before.ctorTypes, after.ctorTypes);
    assert.deepEqual(before.enumAdts, after.enumAdts);
    assert.deepEqual(completeOrFail(before.globalTypes), completeOrFail(after.globalTypes));
    assert.deepEqual(selected(before), [[position, phase]]);
    assert.deepEqual(selected(after), []);
    const fingerprint = metadata => Wire.hash(Wire.write(registryCodec)(selected(metadata))());
    assert.notEqual(fingerprint(before), fingerprint(after),
        'a declaration-only cache key cannot authorize reuse of this metadata');
    assert.deepEqual(prepare(beforeSources).pointerAdtPaths, before.pointerAdtPaths,
        'reintroducing the callback reconstructs the same selection without stale state');
});
