import assert from 'node:assert/strict';
import * as S from '../../output/PureScript.Backend.Optimizer.Syntax/index.js';
import * as C from '../../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as M from '../../output/Data.Maybe/index.js';
import * as PM from '../../output/Data.Map/index.js';
import * as PS from '../../output/Data.Set/index.js';
import { Tuple } from '../../output/Data.Tuple/index.js';
import { optimizeThunkProducers } from '../../output/Gopurs.ThunkFusion/index.js';

const I = C.Int.value, U = C.Unit.value, B = C.Boolean.value;
const thunkType = new C.Func([U], I);
const typed = (type, expr) => new S.Typed(type, expr);
const lit = n => typed(I, new S.Lit(new C.LitInt(n)));
const local = (level, type = I) => typed(type, new S.Local(new M.Just('v' + level), level));
const ref = level => new Tuple(new M.Just('v' + level), level);
const variable = name => new S.Var(new C.Qualified(new M.Just('Audit'), name));
const unit = typed(U, new S.Var(new C.Qualified(new M.Just('Data.Unit'), 'unit')));
const app = (fn, args, type = I) => typed(type, new S.App(fn, args));
const abs = (levels, body, type) => typed(type, new S.Abs(levels.map(ref), body));
const op = (operator, left, right) => typed(I, new S.PrimOp(new S.Op2(operator, left, right)));
const num = (operator, left, right) => op(new S.OpIntNum(operator), left, right);
const add = (a, b) => num(S.OpAdd.value, a, b);
const sub = (a, b) => num(S.OpSubtract.value, a, b);
const mul = (a, b) => num(S.OpMultiply.value, a, b);
const div = (a, b) => num(S.OpDivide.value, a, b);
const unary = (operator, a) => typed(I, new S.PrimOp(new S.Op1(operator, a)));
const eq = (a, b) => typed(B, new S.PrimOp(new S.Op2(new S.OpIntOrd(S.OpEq.value), a, b)));
const branch = (condition, yes, no, type = I) => typed(type, new S.Branch([new S.Pair(condition, yes)], no));
const i64 = value => BigInt.asIntN(64, value);
const tag = value => value.constructor.name;

// Independent strict, lexical interpreter of the IR grammar. A closure captures
// an immutable environment; its body is evaluated only when applied.
function interpret(mod, name, args, budget = 3000) {
  const definitions = new Map(mod.bindings.flatMap(group => group.bindings.map(p => [p.value0, p.value1])));
  let fuel = budget;
  const evaluate = (expr, env) => {
    if (--fuel < 0) throw new Error('FUEL');
    if (expr instanceof S.Typed || expr instanceof S.TypeApp) return evaluate(expr instanceof S.Typed ? expr.value1 : expr.value0, env);
    if (expr instanceof S.Lit) return expr.value0 instanceof C.LitInt ? BigInt(expr.value0.value0) : expr.value0.value0;
    if (expr instanceof S.Local) {
      assert.ok(env.has(expr.value1), `Unbound local ${expr.value1}`);
      return env.get(expr.value1);
    }
    if (expr instanceof S.Var) {
      const q = expr.value0;
      if (q.value0.value0 === 'Data.Unit' && q.value1 === 'unit') return null;
      assert.ok(definitions.has(q.value1), `Unknown global ${q.value1}`);
      return evaluate(definitions.get(q.value1), new Map());
    }
    if (expr instanceof S.Abs) return { params: expr.value0.map(p => p.value1), body: expr.value1, env };
    if (expr instanceof S.App) {
      let fn = evaluate(expr.value0, env);
      let values = expr.value1.map(arg => evaluate(arg, env));
      while (values.length) {
        assert.ok(fn && Array.isArray(fn.params), 'Application of non-function');
        const nextEnv = new Map(fn.env);
        const count = Math.min(values.length, fn.params.length);
        for (let j = 0; j < count; j++) nextEnv.set(fn.params[j], values[j]);
        values = values.slice(count);
        fn = count < fn.params.length
          ? { params: fn.params.slice(count), body: fn.body, env: nextEnv }
          : evaluate(fn.body, nextEnv);
      }
      return fn;
    }
    if (expr instanceof S.Branch) {
      for (const pair of expr.value0) if (evaluate(pair.value0, env)) return evaluate(pair.value1, env);
      return evaluate(expr.value1, env);
    }
    if (expr instanceof S.PrimOp) {
      const operation = expr.value0;
      if (operation instanceof S.Op1) {
        const value = evaluate(operation.value1, env);
        if (operation.value0 instanceof S.OpIntNegate) return i64(-value);
        if (operation.value0 instanceof S.OpIntBitNot) return i64(~value);
        if (operation.value0 instanceof S.OpBooleanNot) return !value;
      }
      const left = evaluate(operation.value1, env), right = evaluate(operation.value2, env);
      const operator = operation.value0;
      if (operator instanceof S.OpIntNum) {
        if (operator.value0 instanceof S.OpAdd) return i64(left + right);
        if (operator.value0 instanceof S.OpSubtract) return i64(left - right);
        if (operator.value0 instanceof S.OpMultiply) return i64(left * right);
        if (operator.value0 instanceof S.OpDivide) return i64(left / right);
      }
      if (operator instanceof S.OpIntBitAnd) return i64(left & right);
      if (operator instanceof S.OpIntBitOr) return i64(left | right);
      if (operator instanceof S.OpIntBitXor) return i64(left ^ right);
      if (operator instanceof S.OpIntOrd) {
        const name = tag(operator.value0);
        return ({ OpEq: left === right, OpNotEq: left !== right, OpLt: left < right,
          OpLte: left <= right, OpGt: left > right, OpGte: left >= right })[name];
      }
    }
    throw new Error('Unsupported IR: ' + tag(expr));
  };
  try {
    const expression = app(variable(name), args.map(value => lit(value)));
    return { status: 'value', value: evaluate(expression, new Map()).toString() };
  } catch (error) {
    if (error.message === 'FUEL') return { status: 'fuel-exhausted' };
    throw error;
  }
}

function makeModule(formula, thunkIndex, options = {}) {
  const levels = [10, 20];
  levels.splice(thunkIndex, 0, 30);
  const types = levels.map(level => level === 30 ? thunkType : I);
  const producerType = new C.Func(types, thunkType);
  const n = local(10), step = local(20), acc = local(30, thunkType);
  const forced = app(acc, [unit]);
  const thunkBody = formula(forced, n, step);
  const next = new Map([[10, options.divergent ? n : sub(n, lit(1))], [20, add(step, n)],
    [30, abs([40], thunkBody, thunkType)]]);
  const recursive = app(typed(producerType, variable('produce')), levels.map(level => next.get(level)), thunkType);
  const producer = abs(levels, branch(eq(n, lit(0)), acc, recursive, thunkType), producerType);
  const seed = options.seedDivision ? div(local(102), lit(2)) : local(102);
  const consumerArgs = new Map([[10, local(100)], [20, local(101)], [30, abs([140], seed, thunkType)]]);
  const produced = app(typed(producerType, variable('produce')), levels.map(level => consumerArgs.get(level)), thunkType);
  const consumer = abs([100, 101, 102], app(produced, [unit]), new C.Func([I, I, I], I));
  return { name: 'Audit', comments: [], imports: PS.empty, dataTypes: PM.empty,
    bindings: [{ recursive: true, bindings: [new Tuple('produce', producer)] },
      { recursive: false, bindings: [new Tuple('consume', consumer)] }],
    exports: PS.empty, reExports: PS.empty, dataDecls: [], classDecls: [],
    foreign: PM.empty, implementations: PM.empty, directives: PM.empty };
}

const formulas = [
  ['add', (a,n,s) => add(a,s)],
  ['left-subtract', (a,n,s) => sub(sub(n,a),s)],
  ['right-subtract', (a,n,s) => sub(a,add(n,s))],
  ['noncommutative', (a,n,s) => add(mul(lit(2),a),n)],
  ['multiply-capture', (a,n,s) => sub(mul(a,add(n,s)),lit(3))],
  ['xor', (a,n,s) => op(S.OpIntBitXor.value,a,mul(n,s))],
  ['and', (a,n,s) => op(S.OpIntBitAnd.value,a,sub(s,n))],
  ['or', (a,n,s) => op(S.OpIntBitOr.value,sub(s,n),a)],
  ['negate', (a,n,s) => sub(unary(S.OpIntNegate.value,a),s)],
  ['bitnot', (a,n,s) => add(unary(S.OpIntBitNot.value,a),n)],
  ['mixed', (a,n,s) => op(S.OpIntBitXor.value,sub(s,mul(n,a)),unary(S.OpIntBitNot.value,n))],
];
const seeds = [0,1,-3,7,2147483647,-2147483648,9223372036854775807n,-9223372036854775808n];
let comparisons = 0, transformedModules = 0, rejectedModules = 0, fuelPairs = 0;
const workers = mod => mod.bindings.flatMap(group => group.bindings).filter(p => p.value0.includes('__gopurs_strict_thunk_'));
const references = (value,name) => value instanceof S.Var ? value.value0.value1===name
  : value!==null && typeof value==='object' && Object.values(value).some(child=>references(child,name));
for (const [label, formula] of formulas) for (let index = 0; index < 3; index++) {
  const original = makeModule(formula,index), optimized = optimizeThunkProducers(original);
  assert.equal(workers(optimized).length,1, `Expected worker for ${label}/position${index}`);
  const optimizedConsumer = optimized.bindings.flatMap(group=>group.bindings).find(p=>p.value0==='consume').value1;
  assert.ok(references(optimizedConsumer,workers(optimized)[0].value0), 'Consumer does not call emitted worker');
  assert.ok(!references(optimizedConsumer,'produce'), 'Consumer still calls original producer');
  transformedModules++;
  for (let depth = 0; depth <= 8; depth++) for (const step of [-3,0,2]) for (const seed of seeds) {
    const args = [depth,step,seed];
    const before = interpret(original,'consume',args), after = interpret(optimized,'consume',args);
    assert.equal(before.status,'value');
    assert.deepEqual(after,before,`${label}/position${index}/${args}`);
    comparisons++;
  }
  const divergent = makeModule(formula,index,{divergent:true});
  const divergentOptimized = optimizeThunkProducers(divergent);
  assert.equal(workers(divergentOptimized).length,1);
  for (const budget of [100,500,1500]) {
    assert.deepEqual(interpret(divergent,'consume',[1,2,7],budget),{status:'fuel-exhausted'});
    assert.deepEqual(interpret(divergentOptimized,'consume',[1,2,7],budget),{status:'fuel-exhausted'});
    fuelPairs++;
  }
}

for (const [label,formula,options] of [
  ['conditional-force',(a,n,s)=>branch(eq(n,lit(2)),n,add(a,lit(1))),{}],
  ['division',(a,n,s)=>div(a,lit(2)),{}],
  ['double-force',(a,n,s)=>add(a,a),{}],
  ['zero-force',(a,n,s)=>n,{}],
  ['seed-division',(a,n,s)=>add(a,n),{seedDivision:true}],
]) for (let index=0; index<3; index++) {
  const original = makeModule(formula,index,options), optimized=optimizeThunkProducers(original);
  assert.equal(workers(optimized).length,0,`Expected refusal: ${label}/position${index}`);
  for (const depth of [0,1,2,5]) {
    assert.deepEqual(interpret(optimized,'consume',[depth,-3,7]),interpret(original,'consume',[depth,-3,7]));
    comparisons++;
  }
  rejectedModules++;
}

// Structural probe only: the interpreter intentionally does not invent a
// runtime initialization policy for recursive value bindings. The binding
// probe uses an Int-annotated forward reference, which need not be initialized
// when the seed would move out of its closure. The body probe ensures the
// proposed conservative guard blocks traversal of both LetRec children.
const letrecProbes = [];
for (const placement of ['binding','body']) {
  const original = makeModule((a,n,s)=>add(a,n),2,{divergent:true});
  const producerType = new C.Func([I,I,thunkType],thunkType);
  const callWithForwardSeed = app(app(typed(producerType,variable('produce')),
    [local(100),local(101),abs([240],local(201),thunkType)],thunkType),[unit]);
  const entries = [new Tuple('v200',placement==='binding'?callWithForwardSeed:lit(0)),
    new Tuple('v201',lit(7))];
  const letrec = typed(I,new S.LetRec(200,entries,placement==='body'?callWithForwardSeed:local(200)));
  original.bindings[1].bindings[0] = new Tuple('consume',abs([100,101,102],letrec,new C.Func([I,I,I],I)));
  const optimized = optimizeThunkProducers(original);
  letrecProbes.push({placement,workers:workers(optimized).length});
}
const expectedLetrecWorkers = 0;
console.log(JSON.stringify({ passed:letrecProbes.every(probe=>probe.workers===expectedLetrecWorkers),
  realPass:'optimizeThunkProducers', transformedModules,
  formulas:formulas.length, thunkPositions:3, differentialComparisons:comparisons,
  rejectedModules, boundedNonterminationPairs:fuelPairs, letrecProbes,
  arithmetic:'signed wrapping int64; bitwise int64, matching current Go emitter',
  limitations:['finite IR subset, no frontend/codegen invocation','fuel exhaustion is not a divergence proof',
    'synthetic well-scoped IR; does not validate frontend annotations or lexical level uniqueness',
    'LetRec probes are structural only; no claim of a confirmed Go initialization bug'] },null,2));
for (const probe of letrecProbes) assert.equal(probe.workers,expectedLetrecWorkers,
  `LetRec ${probe.placement}: current expectation is ${expectedLetrecWorkers} worker(s)`);
