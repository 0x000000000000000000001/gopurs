import assert from 'node:assert/strict';
import test from 'node:test';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Map from '../output/Data.Map/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { fromMap } from '../output/Gopurs.Cache.LookupTable/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { rewrite, optimize, helperName, helperType } from '../output/Gopurs.RegexMatchProjection/index.js';
import { buildGlobalTypes } from '../output/Gopurs.GlobalTypes/index.js';

const expr = syntax => new NeutralExpr(syntax);
const global = (owner, name) => expr(new S.Var(new C.Qualified(owner === null ? Nothing.value : new Just(owner), name)));
const local = level => expr(new S.Local(new Just('sameName'), level));
const apply = (fn, ...args) => expr(new S.App(fn, args));
const integer = n => expr(new S.Lit(new C.LitInt(n)));
const letIn = (level, value, body) => expr(new S.Let(new Just('sameName'), level, value, body));
const just = new C.Qualified(new Just('Data.Maybe'), 'Just');
const isJust = level => expr(new S.PrimOp(new S.Op1(new S.OpIsTag(just), local(level))));
const field = (level, index = 0) => expr(new S.Accessor(local(level), new S.GetCtorField(just, C.SumType.value, 'Maybe', 'Just', 'value0', index)));
const branch = (condition, yes, no = integer(-1)) => expr(new S.Branch([new S.Pair(condition, yes)], no));
const match = (...args) => apply(global('Data.String.Regex', 'match'), ...args);
const groupHead = (value = field(2), owner = 'Data.Array.NonEmpty', name = 'head') => apply(global(owner, name), value);

function consumer({ call = match(local(0), local(1)), head = groupHead(), success = field(3), innerFailure = integer(-2), outerFailure = integer(-1), innerCondition = isJust(3), outerCondition = isJust(2) } = {}) {
  return letIn(2, call, branch(outerCondition, letIn(3, head, branch(innerCondition, success, innerFailure)), outerFailure));
}

test('whole-match projection preserves the call arguments and both failure branches', () => {
  for (const head of [groupHead(), apply(global('Data.Array.NonEmpty', 'adaptMaybe'), global('Data.Array', 'head'), field(2))]) {
    const transformed = rewrite(consumer({ head }));
    assert.ok(transformed instanceof Just);
    const encoded = JSON.stringify(transformed.value0);
    assert.match(encoded, /matchWholeImpl/);
    assert.match(encoded, /matched/);
    assert.match(encoded, /text/);
    assert.match(encoded, /"value0":-1/);
    assert.match(encoded, /"value0":-2/);
    assert.doesNotMatch(encoded, /GetCtorField|adaptMaybe/);
  }
  const wrapped = expr(new S.Typed(C.Int.value, expr(new S.TypeApp(consumer(), C.Int.value))));
  assert.ok(rewrite(wrapped) instanceof Just);
});

test('other observations and unknown match/head implementations retain the original', () => {
  const rejected = [
    ['array escape', { success: local(2) }],
    ['first Maybe escape', { success: local(3) }],
    ['array capture', { success: expr(new S.Abs([new Tuple(Nothing.value, 8)], local(2))) }],
    ['first Maybe alias', { success: letIn(4, local(3), field(4)) }],
    ['other array consumer', { head: apply(global('Probe', 'head'), field(2)) }],
    ['other adapt callback', { head: apply(global('Data.Array.NonEmpty', 'adaptMaybe'), global('Probe', 'head'), field(2)) }],
    ['other match implementation', { call: apply(global('Probe', 'match'), local(0), local(1)) }],
    ['unqualified match', { call: apply(global(null, 'match'), local(0), local(1)) }],
    ['partial match', { call: match(local(0)) }],
    ['extra match operand', { call: match(local(0), local(1), integer(1)) }],
    ['outer failure observes array', { outerFailure: local(2) }],
    ['inner failure observes Maybe', { innerFailure: local(3) }],
    ['unguarded first payload', { innerCondition: integer(1) }],
    ['other payload index', { success: field(3, 1) }],
    ['wrong outer test', { outerCondition: isJust(3) }],
  ];
  for (const [name, options] of rejected) assert.equal(rewrite(consumer(options)), Nothing.value, name);
});

test('only a compatible port opts in; missing or incompatible helpers keep the expression', () => {
  const module = { bindings: [{ recursive: false, bindings: [new Tuple('probe', consumer())] }] };
  const metadata = type => ({ globalTypes: fromMap(type ? Map.insert(ordString)(helperName)(type)(Map.empty) : Map.empty) });
  assert.equal(optimize(metadata(null))(module), module);
  assert.equal(optimize(metadata(C.Int.value))(module), module);
  assert.notEqual(optimize(metadata(helperType))(module), module);
});

test('monomorphized adapters require the instantiated type and cannot impersonate source names', () => {
  const maybeString = new C.ADT('Data.Maybe.Maybe', ['Data', 'Maybe', 'Maybe'], [C.String.value]);
  const maybeMaybeString = new C.ADT('Data.Maybe.Maybe', ['Data', 'Maybe', 'Maybe'], [maybeString]);
  const array = new C.Array(maybeString);
  const adapterType = new C.Func([new C.Func([array], maybeMaybeString), array], maybeString);
  const candidate = (name = 'adaptMaybe__2113979814', type = adapterType, result = maybeString) => consumer({
    head: expr(new S.Typed(result, apply(type ? expr(new S.Typed(type, global('Data.Array.NonEmpty', name))) : global('Data.Array.NonEmpty', name), global('Data.Array', 'head'), field(2)))),
  });
  const known = Map.insert(ordString)(helperName)(helperType)(Map.insert(ordString)('Data.Array.NonEmpty.adaptMaybe')(adapterType)(Map.empty));
  const run = (source, types = known) => {
    const module = { bindings: [{ recursive: false, bindings: [new Tuple('probe', source)] }] };
    return JSON.stringify(optimize({ globalTypes: fromMap(types) })(module)).includes('matchWholeImpl');
  };
  assert.equal(run(candidate()), true);
  assert.equal(run(candidate('adaptMaybe__2113979814', null)), true, 'PBO may annotate only the application');
  assert.equal(run(candidate('adaptMaybe__2113979814', null, C.Any.value)), false);
  for (const name of ['adaptMaybe__', 'adaptMaybe__custom', 'adaptMaybe__12_suffix']) assert.equal(run(candidate(name)), false);
  assert.equal(run(candidate('adaptMaybe__123', C.Any.value)), false);
  assert.equal(run(candidate(), Map.insert(ordString)(helperName)(helperType)(Map.empty)), false);
  assert.equal(run(candidate(), Map.insert(ordString)('Data.Array.NonEmpty.adaptMaybe')(C.Any.value)(known)), false);
  assert.equal(run(candidate(), Map.insert(ordString)('Data.Array.NonEmpty.adaptMaybe__2113979814')(adapterType)(known)), false);

  const none = { span: C.emptySpan, meta: Nothing.value, sourceUsage: Nothing.value, type: Nothing.value };
  const source = { name: 'Data.Array.NonEmpty', foreign: Map.empty, decls: [new C.NonRec(new C.Binding(none, 'adaptMaybe__123', new C.ExprVar(none, new C.Qualified(new Just('Probe'), 'unknown'))))] };
  const originals = buildGlobalTypes([source]);
  const types = Map.insert(ordString)(helperName)(helperType)(Map.insert(ordString)('Data.Array.NonEmpty.adaptMaybe')(adapterType)(originals));
  assert.equal(run(candidate('adaptMaybe__123'), types), false, 'unannotated source declarations also reserve their names');
});

test('strict operand order, selected branches and captured text agree before and after', () => {
  const original = consumer({
    call: match(apply(global('Probe', 'regex'), integer(1)), apply(global('Probe', 'input'), integer(2))),
    success: apply(global('Probe', 'use'), expr(new S.Abs([new Tuple(Nothing.value, 8)], field(3)))),
    outerFailure: apply(global('Probe', 'fail'), integer(7)),
  });
  const transformed = rewrite(original).value0;
  const evaluate = (source, matched, text) => {
    const trace = [];
    const globals = {
      'Probe.regex': value => (trace.push(['regex', value]), 'regex'),
      'Probe.input': value => (trace.push(['input', value]), 'input'),
      'Probe.use': fn => (trace.push(['use']), fn(0)),
      'Probe.fail': value => (trace.push(['fail', value]), value),
      'Data.String.Regex.match': (regex, input) => {
        trace.push(['match', regex, input]);
        return matched ? { tag: 'Just', field: [{ tag: 'Just', field: text }] } : { tag: 'Nothing' };
      },
      'Data.String.Regex.matchWholeImpl': (regex, input) => {
        trace.push(['match', regex, input]);
        return { matched, text: matched ? text : '' };
      },
      'Data.Array.NonEmpty.head': array => array[0],
    };
    const run = (expression, scope = new globalThis.Map()) => {
      const syntax = expression;
      if (syntax instanceof S.Typed) return run(syntax.value1, scope);
      if (syntax instanceof S.TypeApp) return run(syntax.value0, scope);
      if (syntax instanceof S.Local) return scope.get(syntax.value1);
      if (syntax instanceof S.Var) return globals[syntax.value0.value0.value0 + '.' + syntax.value0.value1];
      if (syntax instanceof S.Lit) return syntax.value0.value0;
      if (syntax instanceof S.App) return run(syntax.value0, scope)(...syntax.value1.map(arg => run(arg, scope)));
      if (syntax instanceof S.Let) {
        const value = run(syntax.value2, scope);
        const next = new globalThis.Map(scope); next.set(syntax.value1, value);
        return run(syntax.value3, next);
      }
      if (syntax instanceof S.Abs) return (...values) => {
        const next = new globalThis.Map(scope);
        syntax.value0.forEach((binding, index) => next.set(binding.value1, values[index]));
        return run(syntax.value1, next);
      };
      if (syntax instanceof S.Branch) {
        for (const arm of syntax.value0) if (run(arm.value0, scope)) return run(arm.value1, scope);
        return run(syntax.value1, scope);
      }
      if (syntax instanceof S.PrimOp && syntax.value0 instanceof S.Op1 && syntax.value0.value0 instanceof S.OpIsTag)
        return run(syntax.value0.value1, scope).tag === syntax.value0.value0.value0.value1;
      if (syntax instanceof S.Accessor) {
        const value = run(syntax.value0, scope), access = syntax.value1;
        if (access instanceof S.GetCtorField) {
          assert.equal(value.tag, 'Just', 'payload access stays guarded');
          return value.field;
        }
        if (access instanceof S.GetProp) return value[access.value0];
      }
      throw new Error('unexpected test syntax ' + syntax.constructor.name);
    };
    return { result: run(source), trace };
  };
  for (const [matched, text] of [[false, 'unused'], [true, ''], [true, 'λ😀']]) {
    assert.deepEqual(evaluate(transformed, matched, text), evaluate(original, matched, text));
  }
});
