import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { sanitizeName } from '../output/Gopurs.GoAst/index.js';
import { optimizeThunkProducers } from '../output/Gopurs.ThunkFusion/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const moduleName = 'ThunkProbe';
const int = C.Int.value, unit = C.Unit.value;
const thunk = new C.Func([unit], int);
const producerType = new C.Func([int, thunk], thunk);
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (level, type = int) => typed(type, expr(new S.Local(Nothing.value, level)));
const integer = value => typed(int, expr(new S.Lit(new C.LitInt(value))));
const variable = (name, module = moduleName) => expr(new S.Var(new C.Qualified(module === null ? Nothing.value : new Just(module), name)));
const unitValue = typed(unit, variable('unit', 'Data.Unit'));
const apply = (type, fn, ...args) => typed(type, expr(new S.App(fn, args)));
const lambda = (type, levels, body) => typed(type, expr(new S.Abs(levels.map(level => new Tuple(Nothing.value, level)), body)));
const binary = (type, operator, left, right) => typed(type, expr(new S.PrimOp(new S.Op2(operator, left, right))));
const add = (left, right) => binary(int, new S.OpIntNum(S.OpAdd.value), left, right);
const force = () => apply(int, local(1, thunk), unitValue);
const condition = binary(C.Boolean.value, new S.OpIntOrd(S.OpEq.value), local(0), integer(0));
const branch = (yes, no, test = condition) => typed(thunk, expr(new S.Branch([new S.Pair(test, yes)], no)));
const allBindings = source => source.bindings.flatMap(group => group.bindings);

function fixture(options = {}) {
    const name = options.name ?? 'suspend';
    const type = options.type ?? producerType;
    const self = options.self ?? variable(name);
    const update = options.update ?? lambda(thunk, [2], add(force(), local(0)));
    const counter = options.counter ?? binary(int, new S.OpIntNum(S.OpSubtract.value), local(0), integer(1));
    const recursion = apply(thunk, typed(producerType, self), counter, update);
    const body = options.body ?? branch(local(1, thunk), recursion, options.condition ?? condition);
    const producer = lambda(type, [0, 1], body);
    const seed = options.seed ?? lambda(thunk, [2], integer(7));
    const built = apply(thunk, typed(producerType, variable(name)), options.depth ?? integer(3), seed);
    const consumer = options.consumer ? options.consumer(built) : apply(int, built, options.unit ?? unitValue);
    return {
        name: moduleName, bindings: [
            { recursive: options.recursive ?? true, bindings: [new Tuple(name, producer)] },
            { recursive: false, bindings: [new Tuple('consume', consumer)] },
        ],
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    };
}
function optimize(source) {
    const before = JSON.stringify(source);
    const result = optimizeThunkProducers(source);
    assert.equal(JSON.stringify(source), before, 'the input IR remains immutable');
    return result;
}
function workerNames(source, result) {
    const originals = new Set(allBindings(source).map(pair => pair.value0));
    return allBindings(result).map(pair => pair.value0).filter(name => !originals.has(name));
}

test('strict thunk workers preserve the producer and are inserted only for immediate consumers', () => {
    for (const options of [{}, { unit: typed(unit, expr(S.PrimUndefined.value)) },
        { seed: lambda(thunk, [2], add(local(9), integer(1))) }]) {
        const source = fixture(options);
        const result = optimize(source);
        const [worker] = workerNames(source, result);
        assert.equal(workerNames(source, result).length, 1);
        assert.equal(result.bindings[0].bindings[0].value0, worker);
        assert.deepEqual(result.bindings[1], source.bindings[0], 'the original lazy producer is retained');
        assert.notDeepEqual(result.bindings[2], source.bindings[1]);
    }
    const escaped = fixture({ consumer: built => built });
    assert.deepEqual(optimize(escaped), escaped, 'no unused worker is emitted for an escaping thunk');
});

for (const [label, options] of [
    ['a non-recursive producer', { recursive: false }],
    ['an unqualified recursive call', { self: variable('suspend', null) }],
    ['a foreign recursive call', { self: variable('suspend', 'Other') }],
    ['a Number parameter', { type: new C.Func([C.Number.value, thunk], thunk) }],
    ['a quantified producer type', { type: new C.ForAll(['a'], producerType) }],
    ['a predecessor demanded zero times', { update: lambda(thunk, [2], local(0)) }],
    ['a predecessor demanded twice', { update: lambda(thunk, [2], add(force(), force())) }],
    ['a conditional predecessor demand', { update: lambda(thunk, [2], typed(int,
        expr(new S.Branch([new S.Pair(condition, force())], integer(0))))) }],
    ['division moved out of a thunk', { update: lambda(thunk, [2], binary(int,
        new S.OpIntNum(S.OpDivide.value), force(), integer(2))) }],
    ['use of the ignored Unit parameter', { update: lambda(thunk, [2], add(force(), local(2))) }],
    ['a condition that forces the predecessor', { condition: binary(C.Boolean.value,
        new S.OpIntOrd(S.OpEq.value), force(), integer(0)) }],
    ['a scalar argument that forces the predecessor', { counter: force() }],
    ['an unknown call in the seed', { seed: lambda(thunk, [2], apply(int, variable('opaque'), integer(7))) }],
    ['an unknown scalar argument', { depth: apply(int, variable('opaque'), integer(3)) }],
    ['an untyped undefined Unit', { unit: expr(S.PrimUndefined.value) }],
    ['a producer without recursion', { body: local(1, thunk) }],
]) {
    test(`strict thunk fusion rejects ${label}`, () => {
        const source = fixture(options);
        assert.deepEqual(optimize(source), source);
    });
}

test('recursive scopes block rewriting both initializers and their body', () => {
    for (const insideInitializer of [false, true]) {
        const source = fixture({ consumer: built => {
            const immediate = apply(int, built, unitValue);
            return expr(new S.LetRec(10, [new Tuple('loop', insideInitializer ? immediate : integer(0))],
                insideInitializer ? local(10) : immediate));
        } });
        assert.deepEqual(optimize(source), source);
    }
});

for (const collision of ['binding', 'sanitized binding', 'foreign']) {
    test(`strict thunk worker names avoid a ${collision} collision`, () => {
        const source = fixture({ name: collision === 'sanitized binding' ? 'suspend$' : 'suspend' });
        const [first] = workerNames(source, optimize(source));
        assert.ok(first);
        const reserved = collision === 'sanitized binding' ? sanitizeName(first) : first;
        if (collision === 'foreign') source.foreign = insert(ordString)(reserved)(int)(emptyMap);
        else source.bindings.push({ recursive: false, bindings: [new Tuple(reserved, integer(42))] });
        const result = optimize(source);
        const [worker] = workerNames(source, result);
        assert.equal(workerNames(source, result).length, 1);
        assert.notEqual(sanitizeName(worker), sanitizeName(reserved));
        assert.deepEqual(result.foreign, source.foreign);
        if (collision !== 'foreign') assert.deepEqual(result.bindings.at(-1), source.bindings.at(-1));
    });
}
