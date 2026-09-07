import assert from "node:assert/strict";
import * as C from "../../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import * as S from "../../output/PureScript.Backend.Optimizer.Syntax/index.js";
import { Just } from "../../output/Data.Maybe/index.js";
import { Tuple } from "../../output/Data.Tuple/index.js";
import { empty } from "../../output/Data.Map/index.js";
import { optimizeThunkProducers } from "../../output/Gopurs.ThunkFusion/index.js";

// Exercise the actual pass before frontend simplification can erase LetRec.
const int = C.Int.value;
const thunk = new C.Func([C.Unit.value], int);
const producerType = new C.Func([int, thunk, C.Unit.value], int);
const typed = (type, expr) => new S.Typed(type, expr);
const literal = n => new S.Lit(new C.LitInt(n));
const local = (name, level, type) => typed(type, new S.Local(new Just(name), level));
const binder = (name, level) => new Tuple(new Just(name), level);
const variable = (module, name) => new S.Var(new C.Qualified(new Just(module), name));
const unit = variable("Data.Unit", "unit");
const binary = (op, a, b) => typed(int, new S.PrimOp(new S.Op2(op, a, b)));
const add = S.OpIntNum.create(S.OpAdd.value);
const subtract = S.OpIntNum.create(S.OpSubtract.value);
const n = local("n", 0, int);
const acc = local("acc", 1, thunk);
const producer = variable("RecursiveSeedRegression", "chain");
const nextThunk = typed(thunk, new S.Abs([binder("unused", 2)],
  binary(add, new S.App(acc, [unit]), literal(1))));
const binding = typed(producerType, new S.Abs([binder("n", 0), binder("acc", 1)],
  new S.Branch([
    new S.Pair(new S.PrimOp(new S.Op2(new S.OpIntOrd(S.OpEq.value), n, literal(0))), acc),
  ], new S.App(producer, [binary(subtract, n, literal(1)), nextThunk]))));
const consume = seed => typed(int, new S.App(producer, [
  literal(3), typed(thunk, new S.Abs([binder("unused", 1)], seed)), unit,
]));
const recursiveValue = local("x", 0, int);
const duringInitialization = new S.LetRec(0, [
  new Tuple("x", consume(recursiveValue)),
], recursiveValue);
const inRecursiveScope = new S.LetRec(0, [
  new Tuple("x", literal(7)),
], consume(recursiveValue));
const ordinaryParameter = typed(new C.Func([int], int),
  new S.Abs([binder("seed", 0)], consume(local("seed", 0, int))));
const module = {
  name: "RecursiveSeedRegression",
  foreign: empty,
  bindings: [
    { recursive: true, bindings: [new Tuple("chain", binding)] },
    { recursive: false, bindings: [
      new Tuple("ordinary", ordinaryParameter),
      new Tuple("initializing", duringInitialization),
      new Tuple("recursiveScope", inRecursiveScope),
    ] },
  ],
};
const optimized = optimizeThunkProducers(module);
const bindings = new Map(optimized.bindings.flatMap(group =>
  group.bindings.map(pair => [pair.value0, pair.value1])));
assert.match(JSON.stringify(bindings.get("ordinary")), /chain__gopurs_strict_thunk_/,
  "An already evaluated function parameter should still allow fusion");
assert.deepEqual(bindings.get("initializing"), duringInitialization,
  "Fusion must not advance a recursive value read during initialization");
assert.deepEqual(bindings.get("recursiveScope"), inRecursiveScope,
  "The conservative barrier must include the entire LetRec scope");
console.log("Thunk fusion scope: ordinary parameter fused, recursive scopes preserved");
