import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { toUnfoldable } from '../output/Data.Set/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { collectSharedContainerPairs } from '../output/Gopurs.SharedContainerUsage/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const adt = (name, args = []) => new C.ADT(name, name.split('.'), args);
const record = fields => new C.Record(new C.Row(fields.map(([name, type]) => new Tuple(name, type)), Nothing.value));
const position = record([['offset', C.Int.value], ['row', C.Int.value]]);
const phase = adt('Fixture.Phase');
const pair = adt('Data.Tuple.Tuple', [position, phase]);
const stack = adt('Data.List.Types.List', [pair]);
const callback = new C.Func([stack], C.Int.value);
const ann = type => ({ span: C.emptySpan, meta: Nothing.value, sourceUsage: Nothing.value, type: new Just(type) });
const ref = (type, name = 'value') => new C.ExprVar(ann(type), new C.Qualified(new Just('External'), name));
const binding = (name, type, expression = ref(type)) => new C.NonRec(new C.Binding(ann(type), name, expression));
const moduleOf = (decls, foreign = emptyMap) => ({ name: 'Probe', decls, foreign, dataDecls: [], classDecls: [] });
const collect = modules => toUnfoldable(unfoldableArray)(collectSharedContainerPairs(modules));
const selected = [[position, phase]];
const directBody = type => {
    if (type instanceof C.ForAll || type instanceof C.ConstrainedType) return directBody(type.value1);
    if (type instanceof C.Func && type.value0.length > 0) {
        const next = type.value0.length === 1 ? type.value1 : new C.Func(type.value0.slice(1), type.value1);
        return new C.ExprAbs(ann(type), 'argument', directBody(next));
    }
    return ref(type);
};


test('direct and curried monomorphic worker signatures do not select shared containers', () => {
    const direct = new C.Func([stack], stack);
    const curried = new C.Func([C.Int.value], new C.Func([stack], stack));
    const producer = new C.Func([C.Int.value], stack);
    const dataResult = new C.Func([C.Int.value], new C.Array(stack));
    const directCall = new C.ExprApp(ann(stack), ref(direct, 'worker'), ref(stack, 'data'));
    const modules = [moduleOf([
        binding('direct', direct, directBody(direct)), binding('curried', curried, directBody(curried)),
        binding('producer', producer, directBody(producer)), binding('dataResult', dataResult, directBody(dataResult)),
        binding('directCall', stack, directCall), binding('plainData', stack),
    ])];
    assert.deepEqual(collect(modules), []);
});

test('callback parameters and stored functions select their list or tuple payload family', () => {
    const cases = [
        new C.Func([callback, C.Int.value], C.Int.value),
        record([['handler', callback]]),
        new C.Array(callback),
        adt('Effect.Ref.Ref', [callback]),
        new C.Func([new C.Func([pair], C.Int.value)], C.Int.value),
    ];
    for (const type of cases) {
        const modules = [moduleOf([binding('callbacks', type)])];
        const before = JSON.stringify(modules);
        assert.deepEqual(collect(modules), selected);
        assert.equal(JSON.stringify(modules), before, 'selection must not rewrite the source TAST');
    }
});

test('passing a function as an argument is distinct from directly calling that function', () => {
    const direct = new C.ExprApp(ann(C.Int.value), ref(callback, 'handler'), ref(stack, 'input'));
    const passed = new C.ExprApp(ann(C.Int.value), ref(new C.Func([C.Any.value], C.Int.value), 'consume'), ref(callback, 'handler'));
    assert.deepEqual(collect([moduleOf([binding('direct', C.Int.value, direct)])]), []);
    assert.deepEqual(collect([moduleOf([binding('passed', C.Int.value, passed)])]), selected);
});

test('function aliases and pointfree partial applications are function values', () => {
    const partial = new C.ExprApp(ann(callback), ref(new C.Func([C.Int.value], callback)),
        new C.ExprLit(ann(C.Int.value), new C.LitInt(0)));
    for (const expression of [ref(callback, 'alias'), partial]) {
        assert.deepEqual(collect([moduleOf([binding('stored', callback, expression)])]), selected);
    }
});

test('collection elements and local or recursive callbacks participate in selection', () => {
    const array = new C.ExprLit(ann(new C.Array(callback)), new C.LitArray([ref(callback)]));
    const local = new C.ExprLet(ann(C.Int.value), [binding('stored', record([['handler', callback]]))], ref(C.Int.value));
    const recursive = new C.Rec([new C.Binding(ann(new C.Func([callback], C.Int.value)), 'recursive', ref(new C.Func([callback], C.Int.value)))]);
    for (const decl of [binding('array', new C.Array(callback), array), binding('local', C.Int.value, local), recursive]) {
        assert.deepEqual(collect([moduleOf([decl])]), selected);
    }
});

test('quantifiers, constraints and nested type applications retain callback payload identity', () => {
    const appliedPair = new C.TypeApp(new C.TypeApp(adt('Data.Tuple.Tuple'), [position]), [phase]);
    const appliedStack = new C.TypeApp(adt('Data.List.Types.List'), [appliedPair]);
    const quantified = new C.ForAll(['a'], new C.Func([new C.Func([appliedStack], C.Int.value), new C.TypeVar('a')], C.Int.value));
    const constrained = new C.ConstrainedType([new Tuple(['Fixture', 'Class'], [C.Int.value])], quantified);
    assert.deepEqual(collect([moduleOf([binding('nested', constrained)])]), selected);
    const directQuantified = new C.ForAll(['a'], new C.Func([C.Int.value], new C.Func([stack], C.Int.value)));
    assert.deepEqual(collect([moduleOf([binding('direct', directQuantified, directBody(directQuantified))])]), []);
});

test('foreign callback annotations select families and unrelated data annotations do not', () => {
    const foreign = insert(ordString)('callback')(new Just(new C.Func([callback], C.Int.value)))(
        insert(ordString)('missing')(Nothing.value)(emptyMap));
    assert.deepEqual(collect([moduleOf([], foreign)]), selected);
    const plain = insert(ordString)('data')(new Just(new C.Func([stack], stack)))(emptyMap);
    assert.deepEqual(collect([moduleOf([], plain)]), []);
});

test('selection is deterministic and duplicate callback uses do not duplicate a family', () => {
    const first = moduleOf([binding('stored', record([['handler', callback]]))]);
    const second = { ...moduleOf([binding('apply', new C.Func([callback], C.Int.value))]), name: 'Second' };
    assert.deepEqual(collect([first, second, first]), selected);
    assert.deepEqual(collect([second, first]), collect([first, second]));
});


test('record field order does not duplicate the same selected family', () => {
    const reordered = record([['row', C.Int.value], ['offset', C.Int.value]]);
    const alternative = new C.Func([adt('Data.List.Types.List', [adt('Data.Tuple.Tuple', [reordered, phase])])], C.Int.value);
    const modules = [moduleOf([
        binding('first', record([['handler', callback]])),
        binding('reordered', record([['handler', alternative]])),
    ])];
    assert.deepEqual(collect(modules), selected);
});
