import { withReboxFields } from './codegen-metadata.mjs';
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { Nothing } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { optimizeImmediateApplications } from '../output/Gopurs.ImmediateApplications/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const int = C.Int.value, bool = C.Boolean.value;
const unary = new C.Func([int], int);
const higher = new C.Func([unary], int);
const ref = (level, type = int) => typed(type, expr(new S.Local(Nothing.value, level)));
const literal = n => typed(int, expr(new S.Lit(new C.LitInt(n))));
const lambda = (level, body, type = unary) => typed(type, expr(new S.Abs([new Tuple(Nothing.value, level)], body)));
const call = (head, arg, type = int) => typed(type, expr(new S.App(head, [arg])));
const add = (a, b) => typed(int, expr(new S.PrimOp(new S.Op2(new S.OpIntNum(S.OpAdd.value), a, b))));
const letIn = (level, value, body, type = int) => typed(type, expr(new S.Let(Nothing.value, level, value, body)));
const branch = (condition, yes, no, type = int) => typed(type, expr(new S.Branch([new S.Pair(condition, yes)], no)));
const metadata = withReboxFields({
  elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
  pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
  enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap,
  classDeclsFields: emptyMap,
});
const moduleOf = (body, name = 'Immediate') => ({
  name, bindings: [{recursive: false, bindings: [new Tuple('probe', body)]}],
  comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
  dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
  implementations: emptyMap, directives: emptyMap,
});
const bodyOf = module => module.bindings[0].bindings[0].value1;
function optimize(body) {
  const source = moduleOf(body), before = JSON.stringify(source);
  const result = optimizeImmediateApplications(source);
  assert.equal(JSON.stringify(source), before, 'input IR remains immutable');
  return bodyOf(result);
}
function count(value, ctor) {
  if (!value || typeof value !== 'object') return 0;
  return Number(value instanceof ctor) + Object.values(value).reduce((n, child) => n + count(child, ctor), 0);
}

test('ordinary immediate identity application is reduced', () => {
  const result = optimize(call(lambda(0, ref(0)), literal(19)));
  assert.equal(count(result, S.Abs), 0);
  assert.equal(count(result, S.App), 0);
});

test('branch closures consume an immediate callback without retaining closures', () => {
  const returned = branch(ref(0, bool),
    lambda(1, literal(-1), higher),
    lambda(1, call(ref(1, unary), literal(7)), higher), higher);
  const result = optimize(call(returned, lambda(1, add(ref(1), literal(2)))));
  assert.equal(count(result, S.Abs), 0);
  assert.equal(count(result, S.App), 0);
  assert.equal(count(result, S.Branch), 1);
});

test('unknown/effectful application arguments retain evaluation and ordering', () => {
  const effectful = call(ref(0, unary), literal(1));
  const original = call(lambda(1, literal(9)), effectful);
  assert.deepEqual(optimize(original), original, 'unused calls must still run');
});

test('a callback used twice is kept once, without duplicating its closure body', () => {
  const original = call(lambda(0,
    add(call(ref(0, unary), literal(1)), call(ref(0, unary), literal(2))), higher),
    lambda(1, call(ref(9, unary), ref(1))));
  assert.deepEqual(optimize(original), original);
});

test('a branch returning an opaque function is not distributed speculatively', () => {
  const original = call(branch(ref(0, bool), lambda(1, literal(1), higher), ref(1, higher), higher),
    lambda(1, add(ref(1), literal(2))));
  assert.deepEqual(optimize(original), original);
});

test('recursive local initialization scopes are preserved exactly', () => {
  const original = expr(new S.LetRec(0, [new Tuple('loop', call(lambda(1, ref(1)), literal(4)))], ref(0)));
  assert.deepEqual(optimize(original), original);
});

test('partially applied multi-argument lambdas remain unchanged', () => {
  const fn = typed(new C.Func([int, int], int), expr(new S.Abs([
    new Tuple(Nothing.value, 0), new Tuple(Nothing.value, 1),
  ], add(ref(0), ref(1)))));
  const original = call(fn, literal(4), unary);
  assert.deepEqual(optimize(original), original);
});

test('large callback bodies are not duplicated across branches', () => {
  let body = ref(1);
  for (let i = 0; i < 80; i++) body = add(body, literal(i));
  const original = call(branch(ref(0, bool), lambda(1, call(ref(1, unary), literal(1)), higher),
    lambda(1, call(ref(1, unary), literal(2)), higher), higher), lambda(1, body));
  assert.deepEqual(optimize(original), original);
});

test('large callback bodies move into exactly one live branch without code growth', () => {
  let body = ref(1);
  for (let i = 0; i < 80; i++) body = add(body, literal(i));
  const original = call(branch(ref(0, bool), lambda(1, literal(-1), higher),
    lambda(1, call(ref(1, unary), literal(7)), higher), higher), lambda(1, body));
  const result = optimize(original);
  assert.equal(count(result, S.Abs), 0);
  assert.equal(count(result, S.App), 0);
  assert.equal(count(result, S.PrimOp), 80, 'the large callback body exists once');
});

test('large unused callbacks are discarded without evaluating their bodies', () => {
  let body = call(ref(3, unary), literal(99));
  for (let i = 0; i < 80; i++) body = add(body, literal(i));
  const original = call(branch(ref(0, bool), lambda(1, literal(1), higher),
    lambda(1, literal(2), higher), higher), lambda(1, body));
  assert.deepEqual(optimize(original), branch(ref(0, bool), literal(1), literal(2)));
});

test('duplication budget sums all syntactic leaves in nested branches', () => {
  let body = ref(1);
  for (let i = 0; i < 12; i++) body = add(body, literal(i));
  const leaf = lambda(1, call(ref(1, unary), literal(1)), higher);
  const nested = branch(ref(0, bool), leaf, leaf, higher);
  const original = call(branch(ref(0, bool), nested, nested, higher), lambda(1, body));
  assert.deepEqual(optimize(original), original, 'four copies exceed the extra-syntax budget');
});

// The callback's parameter (3) and the selected branch's payload (3) occupy
// sibling scopes before rewriting. Its captured outer value at level 2 must
// remain distinct when the callback is transplanted under that payload bind.
function executionFixture(name) {
  const returned = branch(ref(0, bool),
    letIn(3, call(ref(1, unary), literal(11)), lambda(4, literal(-1), higher), higher),
    letIn(3, call(ref(1, unary), literal(20)), lambda(4, call(ref(4, unary), ref(3)), higher), higher), higher);
  let callbackBody = call(ref(1, unary), add(ref(3), ref(2)));
  for (let i = 0; i < 80; i++) callbackBody = add(callbackBody, literal(0));
  const result = letIn(2, literal(100), call(returned, lambda(3, callbackBody)));
  return moduleOf(typed(new C.Func([bool, unary], int), expr(new S.Abs([
    new Tuple(Nothing.value, 0), new Tuple(Nothing.value, 1),
  ], result))), name);
}

test('generated optimized Go matches the capture, effect, failure, and reuse oracle', t => {
  const optimized = optimizeImmediateApplications(executionFixture('Immediate'));
  assert.equal(count(bodyOf(optimized), S.Abs), 1, 'only the public abstraction remains');
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-immediate-apps-'));
  t.after(() => rmSync(directory, {recursive: true, force: true}));
  mkdirSync(join(directory, 'purescript'));
  mkdirSync(join(directory, 'gopurs_runtime'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
  writeFileSync(join(directory, 'purescript', 'Immediate.go'), CodeGen.translate(metadata)(optimized));
  writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; rt "gopurs/output/gopurs_runtime"; p "gopurs/output/purescript")
func run(fn func(bool,rt.Value)int64, condition, fail bool) (out string) {
  trace:=[]int64{}
  defer func(){ if e:=recover();e!=nil {out=fmt.Sprintf("panic %v %v",e,trace)} }()
  callback:=rt.Func(func(n rt.Value)rt.Value {trace=append(trace,n.IntVal);if fail&&n.IntVal==140 {panic("chosen")};return rt.Int(n.IntVal*2)})
  result:=fn(condition,callback)
  return fmt.Sprintf("%d %v",result,trace)
}
func main(){
  fn:=p.Call_Immediate_probe
  fmt.Println(run(fn,true,false));fmt.Println(run(fn,false,false));fmt.Println(run(fn,false,true));fmt.Println(run(fn,false,false))
}
`);
  const result = spawnSync('go', ['run', '.'], {cwd: directory, encoding: 'utf8', timeout: 30_000, env: {...process.env, GOWORK: 'off'}});
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stderr);
  const expected = '-1 [11]\n280 [20 140]\npanic chosen [20 140]\n280 [20 140]\n';
  assert.equal(result.stdout, expected);
});
