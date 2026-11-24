import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { optimize } from '../output/Gopurs.ListFusion/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const int = C.Int.value, bool = C.Boolean.value;
const expression = syntax => new NeutralExpr(syntax);
const typed = (ty, value) => expression(new S.Typed(ty, value));
const integer = n => typed(int, expression(new S.Lit(new C.LitInt(n))));
const local = (name, level, ty) => typed(ty, expression(new S.Local(new Just(name), level)));
const qualified = (module, name) => new C.Qualified(new Just(module), name);
const variable = (module, name) => expression(new S.Var(qualified(module, name)));
const apply = (ty, fn, ...args) => typed(ty, expression(new S.App(fn, args)));
const lambda = (ty, args, body) => typed(ty, expression(new S.Abs(args.map(([name, level]) => new Tuple(new Just(name), level)), body)));
const binary = (ty, op, a, b) => typed(ty, expression(new S.PrimOp(new S.Op2(op, a, b))));
const number = (op, a, b) => binary(int, new S.OpIntNum(op), a, b);
const order = (op, a, b) => binary(bool, new S.OpIntOrd(op), a, b);
const branch = (ty, condition, yes, no) => typed(ty, expression(new S.Branch([new S.Pair(condition, yes)], no)));

function fixture(options = {}) {
    const module = options.module ?? 'FusionProbe';
    const cons = options.cons ?? 'Link', nil = options.nil ?? 'End';
    const foldName = options.foldName ?? 'reduce', filterName = options.filterName ?? 'select';
    const list = new C.ADT(`${module}.Chain`, [module, 'Chain'], [int]);
    const generic = new C.ADT(`${module}.Chain`, [module, 'Chain'], [new C.TypeVar('a')]);
    const callbackType = new C.Func([int, int], int);
    const folderType = new C.Func([callbackType, int, list], int);
    const filterType = new C.Func([list], list);
    const traversalType = new C.Func([list, list], list);
    const producerType = new C.Func([int, list], list);
    const ctor = (name, ...fields) => typed(list, expression(new S.CtorSaturated(
        qualified(module, name), C.SumType.value, 'Chain', name,
        fields.map((value, index) => new Tuple(`value${index}`, value)))));
    const field = (input, index) => typed(index === 0 ? int : list, expression(new S.Accessor(input,
        new S.GetCtorField(qualified(module, cons), C.SumType.value, 'Chain', cons, `value${index}`, index))));
    const tag = (name, input) => typed(bool, expression(new S.PrimOp(new S.Op1(new S.OpIsTag(qualified(module, name)), input))));
    const cases = (ty, input, yes, no) => typed(ty, expression(new S.Branch([
        new S.Pair(tag(nil, input), yes), new S.Pair(tag(cons, input), no),
    ], expression(new S.Fail('incomplete')))));
    const add = variable('Data.Semiring', 'intAdd');
    const acc = local('acc', 1, int), source = local('source', 2, list);
    const foldStep = number(options.foldOp ?? S.OpAdd.value, acc, field(source, 0));
    const fold = lambda(options.foldType?.(folderType) ?? folderType,
        [['ignored', 0], ['acc', 1], ['source', 2]],
        cases(int, source, acc, apply(int, variable(module, foldName), options.recursiveCallback ?? add,
            foldStep, field(source, 1))));
    const input = local('input', 0, list), item = local('item', 2, list), gathered = local('gathered', 3, list);
    const self = local('scan', 1, traversalType);
    const head = field(item, 0), tail = field(item, 1);
    const predicate = options.predicate?.(head) ?? order(S.OpEq.value,
        number(S.OpMod.value, head, integer(2)), integer(0));
    const choose = branch(list, predicate,
        apply(list, self, tail, ctor(cons, head, gathered)), apply(list, self, tail, gathered));
    const filterLoop = lambda(options.filterLoopType?.(traversalType) ?? traversalType,
        [['item', 2], ['gathered', 3]], cases(list, item, gathered, choose));
    const filter = lambda(filterType, [['input', 0]], typed(list, expression(new S.LetRec(1,
        [new Tuple('scan', filterLoop)], apply(list, self, input, ctor(nil))))));
    const current = local('current', 2, int), built = local('built', 3, list);
    const recur = local('build', 1, producerType);
    const producerBody = branch(list, order(S.OpLt.value, current, options.bound ?? integer(1)), built,
        apply(list, recur, number(S.OpSubtract.value, current, integer(options.stride ?? 1)),
            ctor(cons, options.element?.(current) ?? current, built)));
    const producer = typed(list, expression(new S.LetRec(1,
        [new Tuple('build', lambda(options.producerType?.(producerType) ?? producerType,
            [['current', 2], ['built', 3]], producerBody))],
        apply(list, recur, local('n', 0, int), options.initialList ?? ctor(nil)))));
    const pipeline = apply(int, variable(options.folderOwner ?? module, foldName), options.callback ?? add,
        options.seed ?? integer(0), apply(list, variable(module, filterName), options.produced ?? producer));
    const consumerBody = options.consumer?.(pipeline) ?? pipeline;
    return {
        name: module,
        bindings: [
            { recursive: true, bindings: [new Tuple(foldName, fold)] },
            { recursive: false, bindings: [new Tuple(filterName, filter)] },
            { recursive: false, bindings: [new Tuple('consume', lambda(new C.Func([int], int), [['n', 0]], consumerBody))] },
        ],
        dataDecls: [{ name: 'Chain', vars: ['a'], constructors: options.constructors ?? [
            { name: nil, fields: [] }, { name: cons, fields: [new C.TypeVar('a'), generic] },
        ] }],
        classDecls: [], dataTypes: emptyMap, foreign: emptyMap, implementations: emptyMap,
        directives: emptyMap, comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
    };
}

const consume = mod => mod.bindings[2].bindings[0].value1;
function inspect(node, visit) {
    if (!node || typeof node !== 'object') return;
    visit(node);
    for (const value of Object.values(node)) inspect(value, visit);
}

test('fusion recognizes renamed ADTs and preserves their public producer/consumer bindings', () => {
    for (const options of [{}, { module: 'Renamed', cons: 'Push', nil: 'Done', foldName: 'aggregate', filterName: 'retain' }]) {
        const source = fixture(options), before = JSON.stringify(source);
        const result = optimize(source);
        assert.equal(JSON.stringify(source), before, 'input is immutable');
        assert.deepEqual(result.bindings.slice(0, 2), source.bindings.slice(0, 2));
        assert.notDeepEqual(consume(result), consume(source));
        const constructors = [], loops = [];
        inspect(consume(result), node => {
            if (node instanceof S.CtorSaturated) constructors.push(node);
            if (node instanceof S.LetRec) loops.push(node);
        });
        assert.equal(constructors.length, 0, 'the complete temporary lists disappear');
        assert.equal(loops.length, 1);
    }
});

for (const [name, options] of [
    ['opaque predicates', { predicate: head => apply(bool, variable('Foreign', 'predicate'), head) }],
    ['effect predicates', { predicate: head => expression(new S.EffectPure(head)) }],
    ['captured predicate values', { predicate: head => order(S.OpEq.value, head, local('captured', 99, int)) }],
    ['zero divisors', { predicate: head => order(S.OpEq.value, number(S.OpMod.value, head, integer(0)), integer(0)) }],
    ['division', { predicate: head => order(S.OpEq.value, number(S.OpDivide.value, head, integer(2)), integer(0)) }],
    ['quantified folders', { foldType: ty => new C.ForAll(['a'], ty) }],
    ['quantified traversal loops', { filterLoopType: ty => new C.ForAll(['a'], ty) }],
    ['quantified producer loops', { producerType: ty => new C.ForAll(['a'], ty) }],
    ['noncommutative folds', { foldOp: S.OpSubtract.value }],
    ['foreign folders with the same name', { folderOwner: 'Unrelated' }],
    ['opaque callbacks', { callback: apply(new C.Func([int, int], int), variable('Foreign', 'callback'), integer(0)) }],
    ['recursive callback evaluations', { recursiveCallback: apply(int, variable('Foreign', 'callback'), integer(0)) }],
    ['unevaluated seeds', { seed: apply(int, variable('Foreign', 'seed'), integer(0)) }],
    ['shared input lists', { produced: local('shared', 1, new C.ADT('FusionProbe.Chain', ['FusionProbe', 'Chain'], [int])) }],
    ['nonempty producer seeds', { initialList: local('shared', 1, new C.ADT('FusionProbe.Chain', ['FusionProbe', 'Chain'], [int])) }],
    ['different recurrences', { stride: 2 }],
    ['transformed elements', { element: current => number(S.OpAdd.value, current, integer(1)) }],
    ['captured range bounds', { bound: local('lower', 98, int) }],
    ['unproved constructor layouts', { constructors: [{ name: 'End', fields: [] }, { name: 'Link', fields: [int, int] }] }],
]) {
    test(`fusion rejects ${name}`, () => {
        const source = fixture(options);
        assert.deepEqual(optimize(source), source);
    });
}

test('each fusion has distinct local levels beyond all binders in its containing function', () => {
    const source = fixture({ consumer: pipeline => number(S.OpAdd.value, pipeline, pipeline) });
    const result = optimize(source), levels = [];
    inspect(consume(result), node => {
        if (node instanceof S.LetRec) {
            levels.push(node.value0);
            inspect(node.value1, child => {
                if (child instanceof S.Abs) levels.push(...child.value0.map(pair => pair.value1));
            });
        }
    });
    assert.equal(levels.length, 6);
    assert.equal(new Set(levels).size, 6);
    assert.ok(levels.every(level => level > 3));
});
