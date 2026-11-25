import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { optimize } from '../output/Gopurs.ClosedScalar/index.js';

const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const int = C.Int.value;
const lit = n => expr(new S.Lit(new C.LitInt(n)));
const scalar = value => typed(int, value);
const ref = (name, level = 0) => new Tuple(name === null ? Nothing.value : new Just(name), level);
const local = (name, level = 0) => expr(new S.Local(ref(name, level).value0, level));
const global = (name, owner = 'Renamed') => expr(new S.Var(new C.Qualified(owner === null ? Nothing.value : new Just(owner), name)));
const lambda = (body, name = 'x', level = 0) => expr(new S.Abs([ref(name, level)], body));
const letIn = (name, value, body, level = 0) => expr(new S.Let(ref(name, level).value0, level, value, body));
const app = (fn, ...args) => expr(new S.App(fn, args));
const add = (a, b) => expr(new S.PrimOp(new S.Op2(new S.OpIntNum(S.OpAdd.value), a, b)));
const record = fields => expr(new S.Lit(new C.LitRecord(fields.map(([name, value]) => new C.Prop(name, value)))));
const field = (value, name = 'run') => expr(new S.Accessor(value, new S.GetProp(name)));
const ctor = fields => expr(new S.CtorSaturated(new C.Qualified(new Just('Evidence'), 'Token'), C.ProductType.value,
  'Token', 'Token', fields.map((value, i) => new Tuple(`value${i}`, value))));
const unknown = app(global('opaque', 'Foreign'), lit(17));
const group = (entries, recursive = false) => ({recursive, bindings: entries.map(([name, value]) => new Tuple(name, value))});
const moduleOf = (groups, foreign = emptyMap) => ({
  name: 'Renamed', bindings: groups, comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
  dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign, implementations: emptyMap, directives: emptyMap,
});
const body = (module, name = 'result') => module.bindings.flatMap(group => group.bindings).find(binding => binding.value0 === name).value1;
function run(value, entries = [], options = {}) {
  const source = moduleOf([group(entries), group([['result', value]], options.recursive)], options.foreign);
  const before = JSON.stringify(source);
  const result = optimize(source);
  assert.equal(JSON.stringify(source), before, 'optimization must not mutate its input');
  return body(result);
}
function unchanged(value, entries = [], options = {}) {
  assert.deepEqual(run(value, entries, options), value);
}

test('evaluates renamed immutable dictionary methods with closed constructor arguments', () => {
  const entries = [
    ['zero', record([['run', lambda(lit(0))]])],
    ['step', lambda(record([['run', lambda(add(lit(1), app(field(local('prior')), ctor([]))), 'ignored', 1)]]), 'prior')],
  ];
  for (const count of [0, 1, 3, 5, 8]) {
    let chain = global('zero');
    for (let i = 0; i < count; i++) chain = app(global('step'), chain);
    assert.deepEqual(run(scalar(app(field(chain), ctor([]))), entries), scalar(lit(count)));
  }
});

test('preserves effect sequencing while replacing only the closed pure result', () => {
  const effect = expr(new S.EffectBind(new Just('token'), 3, unknown,
    expr(new S.EffectPure(scalar(add(lit(12), lit(30)))))));
  assert.deepEqual(run(effect), expr(new S.EffectBind(new Just('token'), 3, unknown,
    expr(new S.EffectPure(scalar(lit(42)))))));
});

test('evaluates curry, partial application, overapplication and anonymous binders exactly', () => {
  const two = expr(new S.Abs([ref('x', 0), ref('y', 1)], add(local('x'), local('y', 1))));
  const nested = lambda(lambda(add(local('x'), local('y', 1)), 'y', 1));
  for (const value of [app(two, lit(5), lit(8)), app(app(two, lit(5)), lit(8)), app(nested, lit(5), lit(8))]) {
    assert.deepEqual(run(scalar(value)), scalar(lit(13)));
  }
  assert.deepEqual(run(scalar(app(lambda(local(null), null), lit(6)))), scalar(lit(6)));
  unchanged(scalar(app(two, lit(5))));
  unchanged(scalar(app(lambda(lit(3)), lit(5), lit(8))));
});

test('closure capture survives shadowing at both matching and different binder levels', () => {
  for (const level of [0, 3]) {
    const value = letIn('x', lit(7), letIn('f', lambda(local('x'), 'unused', 2),
      letIn('x', lit(9), app(local('f', 1), lit(0)), level), 1));
    assert.deepEqual(run(scalar(value)), scalar(lit(7)));
  }
  unchanged(scalar(app(lambda(local('x', 2), 'x', 1), lit(7))));
  unchanged(scalar(local('outside', 4)));
});

test('global initializers cannot capture caller locals, and qualified globals remain lexical', () => {
  const caller = scalar(letIn('x', lit(7), app(global('f'), lit(0))));
  unchanged(caller, [['f', lambda(local('x'), 'ignored', 1)]]);
  assert.deepEqual(run(caller, [['f', lambda(global('x'), 'ignored', 1)], ['x', lit(23)]]), scalar(lit(23)));
  assert.deepEqual(run(scalar(global('x', null)), [['x', lit(23)]]), scalar(lit(23)));
});

test('ignored strict arguments, let initializers and unprojected fields retain opaque work', () => {
  for (const strict of [unknown, expr(new S.Fail('strict argument must fail'))]) {
    for (const value of [
      app(lambda(lit(5)), strict),
      letIn('unused', strict, lit(5)),
      field(record([['run', lit(5)], ['unused', strict]])),
      app(lambda(lit(5)), ctor([strict])),
      app(lambda(lit(5)), lit(0), strict),
    ]) unchanged(scalar(value));
  }
});

test('closures containing unknown work are harmless until invoked', () => {
  const deferred = lambda(unknown);
  assert.deepEqual(run(scalar(app(lambda(lit(5)), deferred))), scalar(lit(5)));
  assert.deepEqual(run(scalar(field(record([['run', lit(5)], ['deferred', deferred]])))), scalar(lit(5)));
  unchanged(scalar(app(deferred, lit(0))));
  unchanged(scalar(app(field(record([['run', deferred]])), lit(0))));
});

test('effects, foreign and imported bindings, unknown branches and unsafe calls are rejected', () => {
  for (const value of [
    expr(new S.EffectPure(lit(5))),
    expr(new S.EffectDefer(lit(5))),
    expr(new S.UncurriedEffectApp(global('f'), [lit(0)])),
    expr(new S.Branch([new S.Pair(unknown, lit(5))], lit(9))),
    app(global('unsafePerformEffect', 'Effect.Unsafe'), lambda(lit(5))),
    global('known', 'Other'),
  ]) unchanged(scalar(value), [['known', lit(5)]]);
  const foreign = insert(ordString)('known')(int)(emptyMap);
  unchanged(scalar(global('known')), [['known', lit(5)]], {foreign});
});

test('nonterminating calls, direct initializers and recursive lets are rejected', () => {
  const value = scalar(app(global('f'), lit(0)));
  unchanged(value, [['f', lambda(app(global('f'), local('x')))]]);
  unchanged(value, [['f', lambda(app(global('g'), local('x')))], ['g', lambda(app(global('f'), local('x')))]]);
  unchanged(scalar(global('a')), [['a', global('b')], ['b', global('a')]]);
  const source = moduleOf([group([['f', lambda(lit(5))]], true), group([['result', value]])]);
  assert.deepEqual(body(optimize(source)), scalar(lit(5)));
  unchanged(scalar(expr(new S.LetRec(1, [new Tuple('loop', lambda(lit(5)))], app(local('loop', 1), lit(0))))));
});

test('Int addition preserves Int32 overflow at both boundaries', () => {
  for (const [a, b, expected] of [[2147483647, 1, -2147483648], [-2147483648, -1, 2147483647], [2147483647, 2147483647, -2], [-2147483648, -2147483648, 0]]) {
    assert.deepEqual(run(scalar(add(lit(a), lit(b)))), scalar(lit(expected)));
  }
});

test('the output domain requires a TAST Int and never coerces other runtime values', () => {
  for (const type of [C.Any.value, C.Number.value, C.String.value, new C.TypeVar('a')]) unchanged(typed(type, add(lit(1), lit(2))));
  for (const value of [expr(new S.Lit(new C.LitNumber(3))), expr(new S.Lit(new C.LitString('3'))), record([['run', lit(3)]]), ctor([]), lambda(lit(3))]) unchanged(scalar(value));
  unchanged(add(lit(1), lit(2)));
  unchanged(scalar(add(lit(1), expr(new S.Lit(new C.LitNumber(2))))));
});

test('fuel is shared across every strict child and depth also counts type wrappers', () => {
  const large = record(Array.from({length: 8200}, (_, i) => [`field${i}`, lit(i)]));
  unchanged(scalar(app(lambda(lit(5)), large)));
  let wrapped = lit(5);
  for (let i = 0; i < 1030; i++) wrapped = typed(C.Any.value, wrapped);
  unchanged(scalar(wrapped));
  let typedApplication = lit(5);
  for (let i = 0; i < 1030; i++) typedApplication = expr(new S.TypeApp(typedApplication, int));
  unchanged(scalar(typedApplication));
});

test('invalid duplicate record fields are rejected conservatively', () => {
  unchanged(scalar(field(record([['run', lit(1)], ['run', lit(2)]]))));
});

const subtract = (a, b) => expr(new S.PrimOp(new S.Op2(new S.OpIntNum(S.OpSubtract.value), a, b)));
const equals = (a, b) => expr(new S.PrimOp(new S.Op2(new S.OpIntOrd(S.OpEq.value), a, b)));
const branch = (condition, yes, no) => expr(new S.Branch([new S.Pair(condition, yes)], no));
const imported = name => global(name, 'Imported');
const retained = (names, value, base = 0) => names.reduceRight((body, name, index) => scalar(letIn(null, imported(name), body, base + index)), scalar(lit(value)));
function recursiveProjection(count, seed = 0, delta = 1, extra = imported('token')) {
  const number = local('count', 0), state = local('state', 1);
  const finished = lambda(record([['state', state], ['token', extra]]), 'state', 1);
  const step = lambda(app(global('walk'), subtract(number, lit(1)), add(state, lit(delta))), 'state', 1);
  const walk = lambda(branch(equals(number, lit(0)), finished, step), 'count', 0);
  const result = scalar(field(app(global('walk'), typeof count === 'number' ? lit(count) : count, lit(seed)), 'state'));
  return moduleOf([group([['walk', walk]], true), group([['result', result]])]);
}

test('closed recursive record projections fold while retaining imported initialization', () => {
  for (const count of [0, 1, 2, 7, 60]) for (const [seed, delta] of [[0, 1], [8, -3], [2147483647, 1]]) {
    const source = recursiveProjection(count, seed, delta);
    const before = JSON.stringify(source);
    assert.deepEqual(body(optimize(source)), retained(['token'], (seed + count * delta) | 0));
    assert.equal(JSON.stringify(source), before);
  }
  const runtime = recursiveProjection(local('opaqueCount', 5));
  assert.deepEqual(body(optimize(runtime)), body(runtime), 'opaque repetition count is never substituted');
  for (const count of [-1, -2147483648, 5000]) {
    const source = recursiveProjection(count);
    assert.deepEqual(body(optimize(source)), body(source), 'divergence or exhausted budget leaves source intact');
  }
});

test('strict global reads keep first-observed order through recursion, even in ignored fields', () => {
  const strict = record([['first', imported('a')], ['again', imported('a')], ['second', imported('b')]]);
  const source = recursiveProjection(60, 0, 1, strict);
  assert.deepEqual(body(optimize(source)), retained(['a', 'b'], 60));
  for (const [value, base] of [
    [scalar(app(lambda(lit(5)), imported('token'))), 1],
    [scalar(field(record([['run', lit(5)], ['other', imported('token')]]))), 0],
  ]) assert.deepEqual(run(value), retained(['token'], 5, base));
  unchanged(scalar(app(imported('function'), lit(2))));
  unchanged(scalar(add(imported('integer'), lit(2))));
});

test('known integer branches evaluate only the selected branch and preserve strict conditions', () => {
  const failure = expr(new S.Fail('selected failure must survive'));
  assert.deepEqual(run(scalar(branch(equals(lit(3), lit(3)), lit(8), failure))), scalar(lit(8)));
  assert.deepEqual(run(scalar(branch(equals(lit(3), lit(4)), failure, lit(9)))), scalar(lit(9)));
  unchanged(scalar(branch(equals(imported('integer'), lit(4)), lit(8), lit(9))));
  unchanged(scalar(branch(equals(lit(3), lit(3)), failure, lit(9))));
  unchanged(scalar(branch(equals(lit(3), lit(4)), lit(8), failure)));
  const source = recursiveProjection(7, 0, 1, unknown);
  assert.deepEqual(body(optimize(source)), body(source), 'opaque strict call cannot disappear behind a projection');
});

test('closed subtraction preserves Int32 overflow and re-entering globals still spends fuel', () => {
  for (const [left, right, value] of [[-2147483648, 1, 2147483647], [2147483647, -1, -2147483648]]) {
    assert.deepEqual(run(scalar(subtract(lit(left), lit(right)))), scalar(lit(value)));
  }
  const recursive = lambda(app(global('again'), local('x')));
  unchanged(scalar(app(global('again'), lit(0))), [['again', recursive]]);
  unchanged(scalar(global('bad')), [['bad', app(lambda(global('bad')), lit(0))]]);
});

test('strict-child stack depth is bounded independently of tail evaluation steps', () => {
  let value = lit(0);
  for (let i = 0; i < 1000; i++) value = add(lit(1), value);
  assert.doesNotThrow(() => unchanged(scalar(value)));
});

test('residual bindings receive distinct fresh levels across the original binding', () => {
  const candidate = scalar(field(record([['run', lit(3)], ['a', imported('first')], ['b', imported('second')]])));
  const source = expr(new S.Abs([ref(null, 0)], record([['left', candidate], ['right', candidate]])));
  const rewritten = run(source);
  const levels = [];
  function walk(value) {
    if (!value || typeof value !== 'object') return;
    if (value instanceof S.Let && value.value0 === Nothing.value) levels.push(value.value1);
    for (const child of Object.values(value)) Array.isArray(child) ? child.forEach(walk) : walk(child);
  }
  walk(rewritten);
  assert.deepEqual(levels, [1, 2, 3, 4]);
});

test('integer primitive operands supply missing annotations without assigning types to unknown values', () => {
  const source = recursiveProjection(60);
  const closed = body(source).value1;
  const value = scalar(add(local('opaqueAccumulator', 9), closed));
  source.bindings.at(-1).bindings[0].value1 = value;
  const result = body(optimize(source));
  assert.ok(result instanceof S.Typed);
  assert.ok(result.value1 instanceof S.PrimOp);
  assert.deepEqual(result.value1.value0.value1, local('opaqueAccumulator', 9));
  const folded = result.value1.value0.value2;
  assert.ok(folded.value1 instanceof S.Let);
  assert.deepEqual(folded.value1.value2, imported('token'));
  assert.deepEqual(folded.value1.value3, scalar(lit(60)));
  unchanged(scalar(add(local('opaqueAccumulator', 9), unknown)));
});

function affineFixture(count, seed, delta, options = {}) {
  const name = options.name ?? 'advanceBy';
  const counter = local('remaining', 0), accumulator = local('total', 1);
  const ordered = values => options.swapped ? values.toReversed() : values;
  const update = options.update ?? (options.commuted ? add(lit(delta), accumulator) : add(accumulator, lit(delta)));
  const recursion = app(global(options.target ?? name, options.owner ?? 'Renamed'),
    ...ordered([subtract(counter, lit(options.decrement ?? 1)), update]));
  const condition = options.condition ?? (options.commuted ? equals(lit(0), counter) : equals(counter, lit(0)));
  const expression = branch(condition, options.base ?? accumulator, options.next ?? recursion);
  const fn = typed(options.type ?? new C.Func([int, int], int),
    expr(new S.Abs(ordered([ref('remaining', 0), ref('total', 1)]), expression)));
  const argument = value => typeof value === 'number' ? lit(value) : value;
  return { name, fn, value: scalar(app(global(name), ...ordered([argument(count), argument(seed)]))) };
}

test('proven closed affine recurrences fold without enumerating iterations and preserve exact Int32 products', () => {
  const counts = [0, 1, 2, 1000, 65535, 65536, 2147483647];
  const values = [-2147483648, -2147483647, -3, 0, 1, 2147483647];
  for (const count of counts) for (const seed of values) for (const delta of values) {
    const {name, fn, value} = affineFixture(count, seed, delta,
      {swapped: delta < 0, commuted: seed < 0, name: 'renamedLoop'});
    const expected = Number(BigInt.asIntN(32, BigInt(seed) + BigInt(count) * BigInt(delta)));
    assert.deepEqual(run(value, [[name, fn]]), scalar(lit(expected)), `${count}, ${seed}, ${delta}`);
  }
  const {name, fn, value} = affineFixture(0, 4294967299, 1);
  assert.deepEqual(run(value, [[name, fn]]), scalar(lit(4294967299)), 'zero iterations return the seed without narrowing');
  for (const [count, seed, expected] of [[-2147483648, 0, -2147483648], [5000, 4294967296, 5000]]) {
    const {name, fn, value} = affineFixture(count, seed, 1);
    assert.deepEqual(run(value, [[name, fn]]), scalar(lit(expected)),
      'the bounded interpreter may first take an exact wrapping step into the proven domain');
  }
});

test('affine summaries retain strict argument reads and reject unknown work and unsupported recurrence shapes', () => {
  const count = app(lambda(lit(2147483647)), imported('first'));
  const seed = app(lambda(lit(0)), imported('second'));
  const fixture = affineFixture(count, seed, 2147483647);
  assert.deepEqual(run(fixture.value, [[fixture.name, fixture.fn]]), retained(['first', 'second'], 1, 1));
  for (const [count, seed, options] of [
    [local('opaqueCount', 8), 0, {}], [-1, 0, {}],
    [4294967296, 0, {}], [5000, 0, {decrement: 2}],
    [5000, 0, {target: 'other'}], [5000, 0, {owner: 'Imported'}],
    [5000, 0, {base: add(local('total', 1), lit(1))}],
    [5000, 0, {type: new C.Func([C.Number.value, int], int)}],
    [5000, 0, {update: add(local('total', 1), unknown)}],
    [5000, 0, {next: letIn('ignored', unknown, lit(1), 4)}],
    [expr(new S.Fail('strict count must fail')), 0, {}], [5000, unknown, {}],
  ]) {
    const {name, fn, value} = affineFixture(count, seed, 1, options);
    unchanged(value, [[name, fn]]);
  }
});
