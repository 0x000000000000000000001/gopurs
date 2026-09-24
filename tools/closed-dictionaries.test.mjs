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
import { cacheClosedDictionaries } from '../output/Gopurs.ClosedDictionaries/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const expr = syntax => new NeutralExpr(syntax);
const typed = (ty, body) => expr(new S.Typed(ty, body));
const global = (name, module = 'Instances') => expr(new S.Var(new C.Qualified(module === null ? Nothing.value : new Just(module), name)));
const local = (name, level = 0) => expr(new S.Local(new Just(name), level));
const lambda = (body, name = 'x', level = 0) => expr(new S.Abs([new Tuple(new Just(name), level)], body));
const app = (fn, ...args) => expr(new S.App(fn, args));
const uncurried = (fn, ...args) => expr(new S.UncurriedApp(fn, args));
const dict = t => new C.ADT('Classes.C', ['Classes', 'C'], [t]);
const int = C.Int.value, variable = new C.TypeVar('a');
const dictionary = dict(new C.Array(int));
const construction = app(global('make'), global('intDict'));
const mapOf = entries => entries.reduce((map, [key, value]) => insert(ordString)(key)(value)(map), emptyMap);
const metadataOf = (globals = []) => withReboxFields({
  elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
  pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
  enumCtors: emptySet, globalFunctions: emptyMap,
  globalTypes: mapOf([
    ['Instances.make', new C.ForAll(['a'], new C.ConstrainedType([new Tuple(['Classes', 'C'], [variable])], dict(new C.Array(variable))))],
    ['Instances.intDict', dict(int)], ...globals,
  ]),
  classDeclsFields: mapOf([['Classes.C', {vars: ['a'], fields: []}]]),
});
const group = (bindings, recursive = false) => ({recursive, bindings: bindings.map(([name, body]) => new Tuple(name, body))});
const moduleOf = groups => ({
  name: 'Example', bindings: groups,
  comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
  dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
  implementations: emptyMap, directives: emptyMap,
});
const bindings = mod => mod.bindings.flatMap(g => g.bindings);
const body = (mod, name) => bindings(mod).find(b => b.value0 === name).value1;
function optimize(groups, metadata = metadataOf()) {
  const mod = moduleOf(groups), before = JSON.stringify(mod);
  const result = cacheClosedDictionaries(metadata)(mod);
  assert.equal(JSON.stringify(mod), before, 'source IR remains immutable');
  return result;
}
const shared = () => group([['shared', typed(dictionary, construction)]]);
const sharedReference = typed(dictionary, global('shared', 'Example'));

test('reuses an annotation-lost imported dictionary even without any new lifts', () => {
  const result = optimize([shared(), group([['use', lambda(app(construction, local('x')))]])]);
  assert.equal(bindings(result).length, 2);
  assert.deepEqual(body(result, 'shared'), typed(dictionary, construction));
  assert.deepEqual(body(result, 'use'), lambda(app(sharedReference, local('x'))));
});

test('does not lift a top-level root through annotation or type-application wrappers', () => {
  const root = typed(dictionary, typed(dictionary, expr(new S.TypeApp(construction, int))));
  const result = optimize([group([['root', root]])]);
  assert.equal(bindings(result).length, 1);
  assert.deepEqual(body(result, 'root'), root);
});

test('closed annotated applications are lifted; names never collide', () => {
  for (const construct of [construction, uncurried(global('make'), global('intDict'))]) {
    const result = optimize([group([
      ['__cached_dict_0', global('intDict')],
      ['use', lambda(typed(dictionary, construct))],
    ])]);
    assert.deepEqual(body(result, 'use'), lambda(typed(dictionary, global('__cached_dict_1', 'Example'))));
    assert.deepEqual(body(result, '__cached_dict_1'), typed(dictionary, construct));
  }
});

test('exact qualifications, application convention and annotations are respected', () => {
  for (const candidate of [
    app(global('make', 'Other'), global('intDict')),
    app(global('make', null), global('intDict')),
    app(local('make'), global('intDict')),
    uncurried(global('make'), global('intDict')),
    app(typed(C.Any.value, global('make')), global('intDict')),
    expr(new S.TypeApp(construction, int)),
    typed(C.Any.value, construction),
  ]) {
    const original = lambda(candidate);
    assert.deepEqual(body(optimize([shared(), group([['use', original]])]), 'use'), original);
  }
});

test('only fully determined monomorphic dictionary results may justify untyped reuse', () => {
  for (const make of [
    new C.ForAll(['a'], new C.Func([dict(int)], dict(new C.Array(variable)))),
    new C.Func([C.String.value], dictionary),
    new C.Func([C.Any.value], dictionary),
    C.Any.value,
  ]) {
    const use = lambda(construction);
    const result = optimize([shared(), group([['use', use]])], metadataOf([['Instances.make', make]]));
    assert.deepEqual(body(result, 'use'), use);
  }
  const wrongType = dict(C.String.value);
  const use = lambda(construction);
  const result = optimize([group([['shared', typed(wrongType, construction)], ['use', use]])]);
  assert.deepEqual(body(result, 'use'), use, 'the binding annotation must match independently recovered type');
});

test('same-module references cannot create cached-getter initialization dependencies', () => {
  const construct = app(global('make'), global('intDict', 'Example'));
  const use = lambda(construct);
  const result = optimize([group([['shared', typed(dictionary, construct)], ['use', use]])], metadataOf([['Example.intDict', dict(int)]]));
  assert.deepEqual(body(result, 'use'), use);
});

test('captured locals remain call-local, including shadowed names at different levels', () => {
  for (const captured of [local('x'), local('x', 4)]) {
    const original = lambda(typed(dictionary, app(global('make'), lambda(captured, 'x', 1))));
    assert.deepEqual(body(optimize([group([['use', original]])]), 'use'), original);
  }
  const closed = typed(dictionary, app(global('make'), lambda(local('x', 1), 'x', 1)));
  const result = optimize([group([['use', lambda(closed)]])]);
  assert.deepEqual(body(result, '__cached_dict_0'), closed, 'internally bound locals are not captures');
});

test('recursive groups and local recursive initialization remain untouched', () => {
  const use = lambda(app(construction, typed(dictionary, construction)));
  const recursive = group([['recursive', use]], true);
  const rec = expr(new S.LetRec(2, [new Tuple('loop', use)], construction));
  const wrapped = lambda(typed(dictionary, app(global('make'), rec)));
  const result = optimize([shared(), recursive, group([['use', rec], ['wrapped', wrapped]])]);
  assert.deepEqual(result.bindings[1], recursive);
  assert.deepEqual(body(result, 'use'), rec);
  assert.deepEqual(body(result, 'wrapped'), wrapped);
});

test('effect syntax, ordinary values and literal dictionaries are not shared', () => {
  const effects = [expr(new S.EffectPure(global('intDict'))), expr(new S.EffectDefer(global('intDict'))),
    expr(new S.UncurriedEffectApp(global('make'), []))];
  for (const effect of effects) {
    const original = lambda(typed(dictionary, app(global('make'), effect)));
    assert.deepEqual(body(optimize([group([['use', original]])]), 'use'), original);
  }
  for (const value of [typed(int, construction), typed(dictionary, expr(new S.Lit(new C.LitRecord([]))))]) {
    const original = lambda(value);
    assert.deepEqual(body(optimize([group([['use', original]])]), 'use'), original);
  }
});
