import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { optimize } from '../output/Gopurs.ConstructorFold/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const int = C.Int.value, bool = C.Boolean.value;
const expr = node => new NeutralExpr(node);
const typed = (ty, value) => expr(new S.Typed(ty, value));
const integer = n => typed(int, expr(new S.Lit(new C.LitInt(n))));
const ref = (name, level) => new Tuple(name === null ? Nothing.value : new Just(name), level);
const local = (name, level, ty = int) => typed(ty, expr(new S.Local(ref(name, level).value0, level)));
const q = (owner, name) => new C.Qualified(new Just(owner), name);
const variable = (owner, name, ty) => typed(ty, expr(new S.Var(q(owner, name))));
const app = (ty, fn, ...args) => typed(ty, expr(new S.App(fn, args)));
const lambda = (ty, refs, body) => typed(ty, expr(new S.Abs(refs.map(([name, level]) => ref(name, level)), body)));
const binary = (ty, op, a, b) => typed(ty, expr(new S.PrimOp(new S.Op2(op, a, b))));
const number = (op, a, b) => binary(int, new S.OpIntNum(op), a, b);
const branch = (ty, condition, yes, no) => typed(ty, expr(new S.Branch([new S.Pair(condition, yes)], no)));
const group = (name, value, recursive = true) => ({recursive, bindings: [new Tuple(name, value)]});

function fixture(options = {}) {
  const owner = options.owner ?? 'Garden', typeName = options.typeName ?? 'Tree';
  const leaf = options.leaf ?? 'Seed', fork = options.fork ?? 'Join';
  const build = options.build ?? 'grow', fold = options.fold ?? 'weigh';
  const tree = new C.ADT(`${owner}.${typeName}`, [owner, typeName], options.generic ? [int] : []);
  const generic = new C.ADT(`${owner}.${typeName}`, [owner, typeName], options.generic ? [new C.TypeVar('a')] : []);
  const buildType = new C.Func([int, int], tree), foldType = new C.Func([tree], int);
  const ctor = (name, ...fields) => typed(tree, expr(new S.CtorSaturated(q(owner, name), C.SumType.value,
    typeName, name, fields.map((value, index) => new Tuple(`value${index}`, value)))));
  const field = (name, value, index, ty) => typed(ty, expr(new S.Accessor(value,
    new S.GetCtorField(q(owner, name), C.SumType.value, typeName, name, `value${index}`, index))));
  const tag = (name, value) => typed(bool, expr(new S.PrimOp(new S.Op1(new S.OpIsTag(q(owner, name)), value))));
  const root = local('tree', 0, tree), n = local('depth', 0), seed = local('seed', 1);
  const recurse = value => app(int, variable(owner, fold, foldType), value);
  const opaque = value => app(int, variable('Opaque', 'touch', new C.Func([int], int)), value);
  const left = recurse(field(fork, root, options.wrongIndex ? 2 : 0, tree));
  const right = recurse(field(fork, root, 1, tree));
  let forkResult = number(options.operation ?? S.OpAdd.value, left, right);
  if (options.duplicateField) forkResult = number(S.OpAdd.value, left, left);
  if (options.discardField) forkResult = left;
  if (options.reverseFields) forkResult = number(S.OpSubtract.value, right, left);
  if (options.opaqueConsumer) forkResult = opaque(forkResult);
  let cases = [new S.Pair(tag(leaf, root), field(leaf, root, 0, int)), new S.Pair(tag(fork, root), forkResult)];
  if (options.missingCase) cases = cases.slice(0, 1);
  if (options.duplicateCase) cases.push(cases[0]);
  const foldBody = typed(int, expr(new S.Branch(cases, expr(new S.Fail('not a constructor')))));
  const consumer = lambda(foldType, [['tree', 0]], foldBody);
  const next = number(S.OpSubtract.value, n, integer(1));
  const recursive = app(tree, variable(options.importedProducer ? 'Elsewhere' : owner, build, buildType),
    options.opaqueRecursion ? opaque(next) : next, seed);
  let continuation = ctor(fork, ctor(leaf, n), recursive);
  if (options.opaqueField) continuation = ctor(fork, ctor(leaf, opaque(n)), recursive);
  if (options.freeCapture) continuation = ctor(fork, ctor(leaf, local('captured', 99)), recursive);
  if (options.wrongArity) continuation = ctor(fork, ctor(leaf, n));
  if (options.borrowed) continuation = local('borrowed', 99, tree);
  if (options.effectInProducer) continuation = typed(tree, expr(new S.EffectPure(continuation)));
  if (options.scalarLet) continuation = typed(tree, expr(new S.Let(new Just('depth'), 0,
    number(S.OpAdd.value, n, integer(1)), continuation)));
  if (options.duplicateCall) {
    const leftBranch = ctor(fork, ctor(leaf, n), recursive);
    const rightBranch = ctor(fork, recursive, ctor(leaf, integer(1)));
    continuation = ctor(fork, leftBranch, rightBranch);
  }
  if (options.shadowedCall) {
    const shadowed = typed(tree, expr(new S.Let(new Just('depth'), 0,
      number(S.OpAdd.value, n, integer(1)), recursive)));
    continuation = ctor(fork, recursive, shadowed);
  }
  if (options.siblingBranches) {
    const pick = binary(bool, new S.OpIntOrd(S.OpEq.value), seed, integer(0));
    const branched = branch(tree, pick, recursive, ctor(leaf, integer(1)));
    continuation = ctor(fork, branched, recursive);
  }
  const condition = binary(bool, new S.OpIntOrd(S.OpLte.value), n, integer(0));
  const producer = lambda(buildType, [['depth', 0], ['seed', 1]], branch(tree, condition, ctor(leaf, seed), continuation));
  let input = app(tree, variable(owner, build, buildType), options.opaqueInput ? opaque(integer(3)) : integer(3), integer(7));
  if (options.partial) input = app(tree, variable(owner, build, buildType), integer(3));
  if (options.overapplied) input = app(tree, variable(owner, build, buildType), integer(3), integer(7), integer(9));
  if (options.escaped) input = local('saved', 4, tree);
  const pipeline = app(int, variable(owner, fold, foldType), input);
  let action = expr(new S.EffectBind(new Just('observed'), 2, opaque(integer(101)), expr(new S.EffectPure(pipeline))));
  if (options.escaped) action = typed(int, expr(new S.Let(new Just('saved'), 4,
    app(tree, variable(owner, build, buildType), integer(3), integer(7)), pipeline)));
  if (options.outerScope) action = lambda(new C.Func([int], int), [[build, 0]], action);
  const bindings = [group(fold, consumer), group(build, producer), group('action', action, false)];
  if (options.collision) bindings.unshift(group(`${build}__gopurs_constructor_fold_0`, integer(0), false));
  if (options.mutual) bindings[1].bindings.push(new Tuple('other', producer));
  const module = {name: owner, bindings, dataDecls: [{name: typeName, vars: options.generic ? ['a'] : [], constructors: [
    {name: leaf, fields: [options.wrongLayout ? C.Number.value : options.generic ? new C.TypeVar('a') : int]},
    {name: fork, fields: [generic, generic]},
  ]}], classDecls: [], dataTypes: emptyMap, foreign: emptyMap, implementations: emptyMap,
  directives: emptyMap, imports: emptySet, exports: emptySet, reExports: emptySet, comments: []};
  return {module, build, fold};
}

const bindings = module => module.bindings.flatMap(group => group.bindings);
const body = (module, name) => bindings(module).find(binding => binding.value0 === name).value1;
const helpers = module => bindings(module).filter(binding => binding.value0.includes('__gopurs_constructor_fold_'));
const text = value => JSON.stringify(value);
function nodesOf(value, constructor) {
  const found = [];
  const visit = value => {
    if (value && typeof value === 'object') {
      if (value instanceof constructor) found.push(value);
      for (const child of Object.values(value)) visit(child);
    }
  };
  visit(value);
  return found;
}
function unchanged(options) {
  const {module} = fixture(options);
  assert.deepEqual(optimize(module), module);
}

test('fuses renamed constructor producers and integer folds, leaving their public bodies and effects intact', () => {
  const {module, build, fold} = fixture({owner: 'Other', typeName: 'Grove', leaf: 'Single', fork: 'Pair', build: 'plant', fold: 'total'});
  const before = text(module), result = optimize(module), worker = helpers(result);
  assert.equal(text(module), before, 'input must be immutable');
  assert.equal(worker.length, 1);
  assert.deepEqual(body(result, build), body(module, build));
  assert.deepEqual(body(result, fold), body(module, fold));
  assert.notEqual(text(body(result, 'action')), text(body(module, 'action')));
  assert.equal(text(body(result, 'action')).match(/101/g).length, 1, 'surrounding effect retained exactly once');
  assert.equal(nodesOf(worker[0].value1, S.CtorSaturated).length, 0);
  assert.equal(nodesOf(worker[0].value1, S.GetCtorField).length, 0);
  assert.equal(text(worker[0].value1).includes(worker[0].value0), true, 'recursive calls select the scalar worker');
});

test('supports arbitrary integer algebras and instantiated TAST layouts', () => {
  for (const operation of [S.OpAdd.value, S.OpSubtract.value, S.OpMultiply.value]) {
    assert.equal(helpers(optimize(fixture({operation, generic: true}).module)).length, 1);
  }
});

test('keeps opaque input evaluation once, and global identity independent of outer local names', () => {
  for (const outerScope of [false, true]) {
    const {module} = fixture({opaqueInput: true, outerScope});
    const result = optimize(module);
    assert.equal(helpers(result).length, 1);
    assert.equal(text(body(result, 'action')).match(/touch/g).length, 2, 'one input call plus one enclosing effect');
  }
});

test('scalar lexical lets retain their binding and generated levels avoid original scopes', () => {
  const {module} = fixture({scalarLet: true});
  const worker = helpers(optimize(module))[0].value1;
  const lets = nodesOf(worker, S.Let);
  assert.equal(lets.filter(node => node.value1 === 0).length, 1, 'original scalar let retained');
  assert.ok(lets.filter(node => node.value1 >= 2).length >= 3, 'new field bindings start beyond original parameters');
});

test('reserves fresh worker names against module declarations', () => {
  const {module} = fixture({collision: true});
  assert.equal(helpers(optimize(module)).filter(binding => binding.value0.endsWith('_1')).length, 1);
});

for (const option of ['missingCase', 'duplicateCase', 'wrongLayout', 'wrongIndex', 'wrongArity',
  'duplicateField', 'discardField', 'reverseFields', 'opaqueConsumer', 'opaqueField', 'freeCapture',
  'opaqueRecursion', 'importedProducer', 'borrowed', 'effectInProducer', 'partial', 'overapplied', 'escaped', 'mutual']) {
  test(`rejects unproved fusion: ${option}`, () => unchanged({[option]: true}));
}

test('rejects division and modulo instead of moving potentially failing arithmetic', () => {
  for (const operation of [S.OpDivide.value, S.OpMod.value]) unchanged({operation});
});

test('deduplicates identical pure recursive worker calls within the same branch', () => {
  const {module} = fixture({duplicateCall: true});
  const worker = helpers(optimize(module))[0].value1;
  assert.equal(nodesOf(worker, S.App).length, 1, 'identical sibling self-calls should share one binding');
});

test('does not share cached self-calls across shadowed lets or sibling branch arms', () => {
  const shadowedWorker = helpers(optimize(fixture({shadowedCall: true}).module))[0].value1;
  assert.equal(nodesOf(shadowedWorker, S.App).length, 2, 'shadowed let must invalidate cached self-call');
  const branchedWorker = helpers(optimize(fixture({siblingBranches: true}).module))[0].value1;
  assert.equal(nodesOf(branchedWorker, S.App).length, 2, 'call inside branch arm must not be reused outside it');
});
