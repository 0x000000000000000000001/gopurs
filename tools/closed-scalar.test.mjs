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

test('recursive groups and direct or indirect global cycles are rejected', () => {
  const value = scalar(app(global('f'), lit(0)));
  unchanged(value, [['f', lambda(app(global('f'), local('x')))]]);
  unchanged(value, [['f', lambda(app(global('g'), local('x')))], ['g', lambda(app(global('f'), local('x')))]]);
  unchanged(scalar(global('a')), [['a', global('b')], ['b', global('a')]]);
  const source = moduleOf([group([['f', lambda(lit(5))]], true), group([['result', value]])]);
  assert.deepEqual(body(optimize(source)), value);
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
  const large = record(Array.from({length: 520}, (_, i) => [`field${i}`, lit(i)]));
  unchanged(scalar(app(lambda(lit(5)), large)));
  let wrapped = lit(5);
  for (let i = 0; i < 80; i++) wrapped = typed(C.Any.value, wrapped);
  unchanged(scalar(wrapped));
  let typedApplication = lit(5);
  for (let i = 0; i < 80; i++) typedApplication = expr(new S.TypeApp(typedApplication, int));
  unchanged(scalar(typedApplication));
});

test('invalid duplicate record fields are rejected conservatively', () => {
  unchanged(scalar(field(record([['run', lit(1)], ['run', lit(2)]]))));
});
