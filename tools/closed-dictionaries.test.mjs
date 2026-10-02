import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { cacheClosedDictionaries } from '../output/Gopurs.ClosedDictionaries/index.js';
import { translate } from '../output/Gopurs.CodeGen/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
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

test('reuse selects the first admissible source binding, even when the use precedes it', () => {
  const source = [
    group([['use', lambda(construction)], ['wrongType', typed(dict(C.String.value), construction)]]),
    group([['recursive', typed(dictionary, construction)]], true),
    group([['preferred', typed(dictionary, construction)]]),
    group([['later', typed(dictionary, construction)]]),
  ];
  const result = optimize(source);
  assert.equal(result.bindings.length, source.length);
  assert.deepEqual(body(result, 'use'), lambda(typed(dictionary, global('preferred', 'Example'))));
  assert.deepEqual(result.bindings.slice(1), source.slice(1));
  assert.deepEqual(body(result, 'wrongType'), typed(dict(C.String.value), construction));
});

test('annotated lifting retains its broader type and module policy than annotation-lost reuse', () => {
  for (const type of [dict(variable), dict(C.Any.value), new C.ForAll(['a'], dict(variable)),
    new C.ConstrainedType([new Tuple(['Classes', 'C'], [int])], dictionary),
    new C.TypeApp(dict(variable), [int])]) {
    const annotated = typed(type, construction);
    const result = optimize([group([['candidate', annotated], ['untyped', lambda(construction)], ['typed', lambda(annotated)]])]);
    assert.deepEqual(body(result, 'untyped'), lambda(construction));
    assert.deepEqual(body(result, 'typed'), lambda(typed(type, global('__cached_dict_0', 'Example'))));
    assert.deepEqual(body(result, '__cached_dict_0'), annotated);
  }
  const annotated = typed(dictionary, app(global('make'), global('intDict', 'Example')));
  const result = optimize([group([['use', lambda(annotated)]])]);
  assert.deepEqual(body(result, '__cached_dict_0'), annotated);
  const unknownClass = {...metadataOf(), classDeclsFields:emptyMap};
  assert.deepEqual(body(optimize([group([['use', lambda(annotated)]])], unknownClass), 'use'), lambda(annotated));
});

test('imported type recovery substitutes across arguments while application spines stay distinct', () => {
  const make = new C.ForAll(['a'], new C.Func([dict(variable), dict(variable)], dict(new C.Array(variable))));
  const metadata = metadataOf([['Instances.makePair', make], ['Instances.stringDict', dict(C.String.value)]]);
  const forms = [
    (a, b) => app(global('makePair'), a, b),
    (a, b) => app(app(global('makePair'), a), b),
    (a, b) => uncurried(global('makePair'), a, b),
  ];
  for (const form of forms) {
    const candidate = form(global('intDict'), global('intDict'));
    const result = optimize([group([['shared', typed(dictionary, candidate)], ['use', lambda(candidate)]])], metadata);
    assert.deepEqual(body(result, 'use'), lambda(sharedReference));
    for (const other of forms.filter(other => other !== form)) {
      const use = lambda(other(global('intDict'), global('intDict')));
      assert.deepEqual(body(optimize([group([['shared', typed(dictionary, candidate)], ['use', use]])], metadata), 'use'), use);
    }
    const mismatched = form(global('intDict'), global('stringDict'));
    const use = lambda(mismatched);
    assert.deepEqual(body(optimize([group([['shared', typed(dictionary, mismatched)], ['use', use]])], metadata), 'use'), use);
  }
});

test('a shared result cannot hide dynamic, open-row or higher-rank imported arguments', () => {
  const candidate = app(global('build'), global('argument'));
  const use = lambda(candidate);
  const rows = tail => new C.Record(new C.Row([new Tuple('value', int)], tail));
  const rejected = [C.Any.value, variable, new C.Array(C.Any.value), dict(variable),
    new C.ForAll(['a'], new C.Func([variable], variable)),
    new C.ConstrainedType([new Tuple(['Classes', 'C'], [int])], int),
    new C.Func([C.Any.value], int), new C.TypeApp(dict(int), [C.Any.value]), rows(new Just(C.Any.value))];
  for (const argument of [...rejected, rows(Nothing.value), new C.Func([int], int)]) {
    const metadata = metadataOf([
      ['Instances.build', new C.Func([argument], dictionary)], ['Instances.argument', argument],
    ]);
    const result = optimize([group([['shared', typed(dictionary, candidate)], ['use', use]])], metadata);
    assert.deepEqual(body(result, 'use'), rejected.includes(argument) ? use : lambda(sharedReference));
  }
});

test('lifting is outermost and source-ordered, reserves all binding names, and keeps fresh sites distinct', () => {
  const inner = typed(dictionary, construction);
  const outer = typed(dictionary, app(global('combine'), inner));
  const array = values => expr(new S.Lit(new C.LitArray(values)));
  const source = [
    group([['__cached_dict_0', global('intDict')]], true),
    group([['outer', lambda(outer)], ['siblings', lambda(array([inner, inner]))]]),
    group([['__cached_dict_2', global('intDict')]]),
  ];
  const result = optimize(source);
  assert.equal(result.bindings.length, source.length + 1);
  assert.deepEqual(result.bindings[0], source[0]);
  assert.deepEqual(result.bindings[2], source[2]);
  assert.deepEqual(body(result, 'outer'), lambda(typed(dictionary, global('__cached_dict_1', 'Example'))));
  assert.deepEqual(body(result, 'siblings'), lambda(array([3, 4].map(n => typed(dictionary, global(`__cached_dict_${n}`, 'Example'))))));
  assert.deepEqual(result.bindings.at(-1), group([
    ['__cached_dict_1', outer], ['__cached_dict_3', inner], ['__cached_dict_4', inner],
  ]));
});

test('a rejected annotation protects its root but still visits independently liftable children', () => {
  const inner = typed(dictionary, construction);
  const wrappedRoot = typed(C.Any.value, typed(dictionary, expr(new S.TypeApp(construction, int))));
  const container = typed(C.Any.value, app(global('combine'), inner));
  const result = optimize([shared(), group([['root', lambda(wrappedRoot)], ['nested', lambda(container)]])]);
  assert.deepEqual(body(result, 'root'), lambda(wrappedRoot));
  assert.deepEqual(body(result, 'nested'), lambda(typed(C.Any.value,
    app(global('combine'), typed(dictionary, global('__cached_dict_0', 'Example'))))));
  assert.deepEqual(body(result, '__cached_dict_0'), inner);
});

test('closure proofs distinguish let initializers, bound anonymous locals and mismatched binder identities', () => {
  const letIn = (value, next) => expr(new S.Let(new Just('x'), 1, value, next));
  const anonymous = expr(new S.Abs([new Tuple(Nothing.value, 1)], expr(new S.Local(Nothing.value, 1))));
  for (const closed of [letIn(global('intDict'), local('x', 1)), anonymous]) {
    const candidate = typed(dictionary, app(global('make'), closed));
    assert.deepEqual(body(optimize([group([['use', lambda(candidate)]])]), '__cached_dict_0'), candidate);
  }
  for (const captured of [letIn(local('x', 1), local('x', 1)), lambda(local('other', 1), 'x', 1),
    lambda(local('x_prime_', 1), "x'", 1),
    expr(new S.Abs([new Tuple(Nothing.value, 1)], local('x', 1)))]) {
    const use = lambda(typed(dictionary, app(global('make'), captured)));
    assert.deepEqual(body(optimize([group([['use', use]])]), 'use'), use);
  }
});

test('effect syntax rejects enclosing lifts while independent dictionary children can still be cached', () => {
  const scopes = [
    value => expr(new S.PrimEffect(new S.EffectRefNew(value))),
    value => expr(new S.EffectBind(Nothing.value, 1, value, expr(S.PrimUndefined.value))),
    value => expr(new S.EffectPure(value)),
    value => expr(new S.EffectDefer(value)),
    value => expr(new S.UncurriedEffectApp(global('effect'), [value])),
    value => expr(new S.UncurriedEffectAbs([new Tuple(Nothing.value, 1)], value)),
  ];
  for (const scope of scopes) {
    const enclosing = value => typed(dictionary, app(global('make'), scope(value)));
    const rejected = lambda(enclosing(construction));
    assert.deepEqual(body(optimize([group([['use', rejected]])]), 'use'), rejected);
    const child = typed(dictionary, construction);
    const result = optimize([group([['use', lambda(enclosing(child))]])]);
    assert.deepEqual(body(result, 'use'), lambda(enclosing(typed(dictionary, global('__cached_dict_0', 'Example')))));
    assert.deepEqual(result.bindings.at(-1), group([['__cached_dict_0', child]]));
  }
});

test('generated getters stay lazy, reuse proven roots, cache each lifted site and retain call-local captures', t => {
  const source = moduleOf([shared(), group([
    ['use', lambda(construction)],
    ['firstLift', lambda(typed(dictionary, construction))],
    ['secondLift', lambda(typed(dictionary, construction))],
    ['captured', lambda(typed(dictionary, app(global('make'), local('x'))))],
  ])]);
  const code = translate(metadataOf())(source);
  const work = mkdtempSync(join(tmpdir(), 'gopurs-closed-dictionaries-'));
  t.after(() => rmSync(work, {recursive:true, force:true}));
  mkdirSync(join(work, 'gopurs_runtime'));
  mkdirSync(join(work, 'purescript'));
  writeFileSync(join(work, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(work, 'gopurs_runtime/runtime.go'), runtimeGoCode);
  writeFileSync(join(work, 'purescript/Example.go'), code);
  writeFileSync(join(work, 'purescript/cache_test.go'), `package purescript
import ("testing"; "gopurs/output/gopurs_runtime")
var constructions int
func dictionary(value int64) gopurs_runtime.Value {
  return gopurs_runtime.RecordDict1("value", gopurs_runtime.Int(value))
}
func Get_Instances_intDict() gopurs_runtime.Value { return dictionary(7) }
func Get_Instances_go__make() gopurs_runtime.Value {
  return gopurs_runtime.Func(func(value gopurs_runtime.Value) gopurs_runtime.Value {
    constructions++
    return value
  })
}
func TestDictionaryCaching(t *testing.T) {
  check := func(value gopurs_runtime.Value, want int64, calls int) {
    t.Helper()
    if got := gopurs_runtime.RecordGet(value, "value").IntVal; got != want || constructions != calls {
      t.Fatalf("value %d (want %d), constructions %d (want %d)", got, want, constructions, calls)
    }
  }
  if constructions != 0 { t.Fatal("getter evaluated before first use") }
  for i := 0; i < 3; i++ { check(Call_Example_use(dictionary(0)), 7, 1) }
  check(Get_Example_shared(), 7, 1)
  for i := 0; i < 3; i++ { check(Call_Example_firstLift(dictionary(0)), 7, 2) }
  for i := 0; i < 3; i++ { check(Call_Example_secondLift(dictionary(0)), 7, 3) }
  check(Call_Example_captured(dictionary(11)), 11, 4)
  check(Call_Example_captured(dictionary(22)), 22, 5)
}
`);
  const result = spawnSync('go', ['test', '-v', '-race', '-count=1', './purescript'], {
    cwd:work, encoding:'utf8', timeout:60_000, env:{...process.env, GOWORK:'off', GOMAXPROCS:'2'},
  });
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stdout + result.stderr);
  assert.match(result.stdout, /--- PASS: TestDictionaryCaching/);
});
