import assert from 'node:assert/strict';
import test from 'node:test';
import { optimize } from '../output/Gopurs.RefLoopRegions/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr as E } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';

const make = syntax => new E(syntax);
const typed = (type, value) => make(new S.Typed(type, value));
const int = value => typed(C.Int.value, make(new S.Lit(new C.LitInt(value))));
const local = level => typed(C.Int.value, make(new S.Local(Nothing.value, level)));
const variable = (owner, name) => make(new S.Var(new C.Qualified(new Just(owner), name)));
const app = (owner, name, args) => make(new S.App(variable(owner, name), args));
const pure = value => make(new S.EffectPure(value));
const bind = (level, value, body) => make(new S.EffectBind(Nothing.value, level, value, body));
const lambda = (levels, body) => make(new S.Abs(levels.map(level => new Tuple(Nothing.value, level)), body));
const binary = (operator, left, right) => make(new S.PrimOp(new S.Op2(operator, left, right)));
const num = (operator, left, right) => typed(C.Int.value, binary(new S.OpIntNum(operator), left, right));
const ord = (operator, left, right) => typed(C.Boolean.value, binary(new S.OpIntOrd(operator), left, right));
const branch = (condition, yes, no) => make(new S.Branch([new S.Pair(condition, yes)], no));
const step = (name, value) => pure(make(new S.CtorSaturated(new C.Qualified(new Just('Control.Monad.Rec.Class'), name), C.SumType.value, 'Step', name, [new Tuple('value0', value)])));
const read = (st, ref) => st ? make(new S.PrimEffect(new S.EffectRefRead(ref))) : app('Effect.Ref', 'read', [ref]);
const write = (st, ref, value) => st ? make(new S.PrimEffect(new S.EffectRefWrite(ref, value))) : app('Effect.Ref', 'write', [value, ref]);
const max = 2147483647, min = -2147483648;

function fixture({st = false, increment = true, customInstance = false, stride = 1, observed = false,
  resultRef = false, argumentCall = false, scalarUnknown = false, noAllocation = false, name = 'renamedWorker'} = {}) {
  // Callee and caller levels intentionally overlap and permute delta/ref slots.
  const seed = local(0), delta = local(1), ref = local(2), state = local(3);
  const add = num(S.OpAdd.value, local(4), delta);
  const mutation = increment ? (st ? bind(4, read(true, ref), write(true, ref, add))
    : app('Effect.Ref', 'modifyImpl', [lambda([4], make(new S.Lit(new C.LitRecord([
      new C.Prop('state', add), new C.Prop('value', add),
    ])))), ref])) : write(st, ref, state);
  const next = num(increment ? S.OpSubtract.value : S.OpAdd.value, state, int(stride));
  let repeated = bind(5, mutation, step('Loop', next));
  if (observed) repeated = bind(6, app('Observers', 'inspect', [ref]), repeated);
  const callback = lambda([3], branch(ord(increment ? S.OpLte.value : S.OpGt.value, state, increment ? int(0) : delta),
    step('Done', variable('Data.Unit', 'unit')), repeated));
  const dictionary = variable(customInstance ? 'Other' : st ? 'Control.Monad.ST.Internal' : 'Control.Monad.Rec.Class', st ? 'monadRecST' : 'monadRecEffect');
  const helper = lambda([0, 1, 2], make(new S.App(make(new S.Accessor(dictionary, new S.GetProp('tailRecM'))), [callback, seed])));
  const initial = local(1), count = argumentCall ? app('Observer', 'count', []) : scalarUnknown ? make(new S.Local(Nothing.value, 0)) : local(0);
  const callerRef = local(3), callerDelta = local(2);
  const body = bind(4, app('Renamed', name, [count, callerDelta, callerRef]), resultRef ? pure(callerRef) : read(st, callerRef));
  const region = noAllocation ? body : bind(3, st ? app('Control.Monad.ST.Internal', 'newImpl', [initial]) : app('Effect.Ref', '_new', [initial]), body);
  const entry = st ? app('Control.Monad.ST.Internal', 'run', [region]) : region;
  const module = {name:'Renamed', bindings:[{recursive:false, bindings:[new Tuple(name, helper), new Tuple('entry', entry)]}]};
  return {module, entry, helper};
}

function optimized(options) {
  const input = fixture(options), before = JSON.stringify(input.module);
  const output = optimize(input.module);
  assert.equal(JSON.stringify(input.module), before, 'source module remains unchanged');
  assert.deepEqual(output.bindings[0].bindings[0].value1, input.helper, 'public helper is unchanged');
  return {...input, result:output.bindings[0].bindings[1].value1};
}

function evaluate(expression, env, fallback) {
  if (expression === fallback) return 'fallback';
  const s = expression;
  if (s instanceof S.Typed) return evaluate(s.value1, env, fallback);
  if (s instanceof S.EffectPure) return evaluate(s.value0, env, fallback);
  if (s instanceof S.Local) {
    assert.ok(env.has(s.value1), 'no captured reference or free callee parameter');
    return env.get(s.value1);
  }
  if (s instanceof S.Lit && s.value0 instanceof C.LitInt) return s.value0.value0;
  if (s instanceof S.Branch) {
    for (const pair of s.value0) if (evaluate(pair.value0, env, fallback)) return evaluate(pair.value1, env, fallback);
    return evaluate(s.value1, env, fallback);
  }
  if (s instanceof S.PrimOp && s.value0 instanceof S.Op2) {
    const op=s.value0, a=evaluate(op.value1,env,fallback), b=evaluate(op.value2,env,fallback);
    const kind=op.value0.value0;
    if(kind instanceof S.OpAdd)return (a+b)|0;
    if(kind instanceof S.OpMultiply)return Math.imul(a,b);
    if(kind instanceof S.OpSubtract)return (a-b)|0;
    if(kind instanceof S.OpGt)return a>b;
    if(kind instanceof S.OpLt)return a<b;
    if(kind instanceof S.OpLte)return a<=b;
  }
  throw new Error('Unexpected optimized node '+JSON.stringify(expression));
}

test('private ST and Effect increments preserve caller levels, arbitrary seeds and Int32 overflow', () => {
  for(const st of [false,true])for(const name of ['renamedWorker','anotherHelper']){
    const {entry,result}=optimized({st,name}); assert.notDeepEqual(result,entry);
    for(const count of [min,-1,0,1,17,max])for(const initial of [min,-9,0,31,max])for(const delta of [min,-3,0,1,max]){
      const expected=count>0?(initial+Math.imul(count,delta))|0:initial;
      assert.equal(evaluate(result,new Map([[0,count],[1,initial],[2,delta]]),entry),expected);
    }
  }
});

test('private upward writes select original state for zero iterations and retain maxBound fallback', () => {
  for(const st of [false,true]){
    const {entry,result}=optimized({st,increment:false});assert.notDeepEqual(result,entry);
    for(const start of [min,-7,0,1,max-1,max])for(const limit of [min,-1,0,17,max-1,max])for(const initial of [min,7,max]){
      assert.equal(evaluate(result,new Map([[0,start],[1,initial],[2,limit]]),entry),limit===max?'fallback':start>limit?initial:limit);
    }
  }
});

test('unknown effects, escaping/shared refs, nonstandard instances, strides and unevaluated inputs are refused', () => {
  for(const st of [false,true])for(const flag of ['customInstance','observed','resultRef','argumentCall','scalarUnknown','noAllocation']){
    const {entry,result}=optimized({st,[flag]:true});assert.deepEqual(result,entry,`${st}:${flag}`);
  }
  for(const st of [false,true])for(const increment of [false,true]){
    const {entry,result}=optimized({st,increment,stride:2});assert.deepEqual(result,entry);
  }
});
