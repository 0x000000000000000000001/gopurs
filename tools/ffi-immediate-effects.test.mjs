import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { withReboxFields } from './codegen-metadata.mjs';
import * as Map from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { workerInputs, translateWithFunctions } from '../output/Gopurs.CodeGen/index.js';
import { ffiFunctionInfos, ffiValueWorkers, generateFfiBridge } from '../output/Gopurs.FfiBridge/index.js';
import { key } from '../output/Gopurs.FfiBridge.ImmediateEffects/index.js';
import { prepareFfi } from '../output/Gopurs.FfiSupport/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { StmtEmpty, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import { translate } from '../output/Gopurs.ImmediateFfiEffects/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const owner = 'RenamedBoundary';
const int = C.Int.value;
const variable = new C.TypeVar('a');
const foreign = new C.ADT('Foreign.Foreign', ['Foreign', 'Foreign'], []);
const effect = result => new C.ADT('Effect.Effect', ['Effect', 'Effect'], [result]);
const fn = (args, result) => new C.Func(args, result);
const foreigns = [
  ['identity', new C.ForAll(['a'], fn([variable], effect(variable)))],
  ['native', fn([int], effect(int))],
  ['foreign', fn([foreign], effect(foreign))],
  ['curriedForeign', fn([int], fn([foreign], effect(foreign)))],
  ['callback', fn([fn([int], int), int], effect(int))],
  ['returned', fn([int], effect(fn([int], int)))],
  ['panic', fn([int], effect(int))],
  ['ordinary', fn([int], fn([int], int))],
  ['saved', fn([int], effect(int))],
].map(([name, type]) => new Tuple(name, new Just(type)));
const prepared = prepareFfi({ moduleName: owner, path: owner + '.go' })(owner + '_')(`
var Trace []string
var Marker = &struct{}{}
func Identity(x any) func() any { Trace = append(Trace, "factory"); return func() any { Trace = append(Trace, "effect"); return x } }
func Native(x int64) func() any { return func() any { return x + 1 } }
func Foreign(x any) func() any { return func() any { return x } }
func CurriedForeign(x int64, value any) func() any { return func() any { return value } }
func Callback(f func(int64) int64, x int64) func() int64 { return func() int64 { return f(x) } }
func Returned(x int64) func() func(int64) int64 { return func() func(int64) int64 { return func(y int64) int64 { return x + y } } }
func Panic(x int64) func() int64 { Trace = append(Trace, "factory"); return func() int64 { Trace = append(Trace, "effect"); panic(Marker) } }
func Ordinary(x int64) func(int64) int64 { return func(y int64) int64 { return x + y } }
var Saved = func(x int64) func() int64 { return func() int64 { return x } }
`)();
const infos = ffiFunctionInfos(owner)(foreigns)(prepared.decls);
const workers = ffiValueWorkers(owner)([])(foreigns)(prepared.decls);
const wrappers = generateFfiBridge(owner)([])(prepared.decls)(foreigns);
const lookup = name => Map.lookup(ordString)(name)(infos);
const ir = syntax => new TcoExpr(null, syntax);
const local = (name, level = 0) => ir(new S.Local(new Just(name), level));
const global = (name, module = owner) => ir(new S.Var(new C.Qualified(module === null ? Nothing.value : new Just(module), name)));
const app = (target, ...args) => ir(new S.App(target, args));
const pure = value => ir(new S.EffectPure(value));
const arg = name => ir(new S.Lit(new C.LitString(name)));
const bind = continuation => ir(new S.EffectBind(new Just('answer'), 1, local('action'), continuation));
const context = {
  metadata: { ffiFunctions: infos, globalFunctions: Map.empty }, modNameStr: owner,
  bound: Map.empty, depth: 0, options: { isTail: true, inEffectBlock: true },
  tcoIdent: Nothing.value, loopCtx: [], mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
};
const translator = () => id => expr => ({
  stmts: StmtEmpty.value, nextId: id + 1, exprType: Go.TypeValue.value,
  expr: expr.value1 instanceof S.Lit ? Go.rawGo(`capture("${expr.value1.value0.value0}")`) : Go.rawGo('gopurs_runtime.Int(42)'),
});
const emit = (factory, body = bind(pure(arg('result'))), ctx = context) =>
  translate(translator)(ctx)(0)(new Just('action'))(0)(factory)(body);

test('nominal effect signatures retain polymorphism, native any, Foreign and callback results', () => {
  for (const name of ['identity', 'native', 'foreign', 'curriedForeign', 'callback', 'returned', 'panic']) {
    assert.ok(lookup(key(name)) instanceof Just, name);
    assert.equal(lookup(key(name)).value0.fRet, Go.TypeValue.value);
  }
  for (const name of ['ordinary', 'saved']) assert.equal(lookup(key(name)), Nothing.value);
  assert.match(workers, /return gopurs_runtime.Box\(go_result\)/);
  assert.match(workers, /gopurs_runtime.ForeignValue\(arg0\)/);
  assert.match(workers, /return gopurs_runtime.Any\(go_result\)/);
});

test('only a saturated adjacent bind with no remaining action use can select a worker', () => {
  const factory = app(global('identity'), arg('first'));
  assert.ok(emit(factory) instanceof Just);
  for (const value of [global('identity'), app(global('identity'), arg('first'), arg('extra')),
    app(local('alias'), arg('first')), app(global('ordinary'), arg('first'))]) {
    assert.equal(emit(value), Nothing.value);
  }
  assert.equal(emit(factory, pure(local('action'))), Nothing.value);
  assert.equal(emit(factory, bind(pure(local('action')))), Nothing.value);
  const capture = ir(new S.Abs([new Tuple(new Just('unused'), 2)], local('action')));
  assert.equal(emit(factory, bind(pure(capture))), Nothing.value);
  assert.equal(emit(factory, ir(new S.EffectBind(new Just('answer'), 1, local('action', 3), pure(arg('result'))))), Nothing.value);
  assert.equal(emit(factory, undefined, { ...context, options: { isTail: false, inEffectBlock: false } }), Nothing.value);
});

test('uncurried calls retain their evaluation boundaries while type applications stay transparent', () => {
  const target = global('identity');
  const value = arg('first');
  for (const factory of [
    ir(new S.UncurriedApp(target, [value])),
    app(ir(new S.UncurriedApp(target, [])), value),
    ir(new S.UncurriedApp(app(target, value), [])),
    app(ir(new S.Typed(fn([int], effect(int)), ir(new S.UncurriedApp(target, [])))), value),
  ]) assert.equal(emit(factory), Nothing.value);
  assert.ok(emit(app(ir(new S.TypeApp(target, int)), value)) instanceof Just);
  assert.ok(emit(ir(new S.Typed(effect(int), app(target, value)))) instanceof Just);
});

test('imported workers and absent workers both participate in the cache dependency snapshot', () => {
  const fullKey = owner + '.' + key('identity');
  const available = Map.singleton(fullKey)(lookup(key('identity')).value0);
  const backend = {
    name: 'Consumer', bindings: [{ bindings: [new Tuple('result', new NeutralExpr(new S.Var(new C.Qualified(new Just(owner), 'identity'))))] }],
    foreign: Map.empty, imports: emptySet,
  };
  for (const functions of [Map.empty, available]) {
    const snapshot = workerInputs(functions)(backend);
    const requested = Map.lookup(ordString)(fullKey)(snapshot);
    assert.ok(requested instanceof Just);
    assert.equal(requested.value0 instanceof Just, functions === available);
  }
  assert.ok(emit(app(global('identity'), arg('first')), undefined, {
    ...context, modNameStr: 'Consumer', metadata: { ffiFunctions: Map.empty, globalFunctions: available },
  }) instanceof Just);
});

test('module emission selects the effect worker without exposing its reserved key as an ordinary function', () => {
  const e = syntax => new NeutralExpr(syntax);
  const target = name => e(new S.Var(new C.Qualified(new Just(owner), name)));
  const literal = e(new S.Typed(int, e(new S.Lit(new C.LitInt(41)))));
  const application = name => e(new S.App(target(name), [literal]));
  const collision = Go.sanitizeName(key('identity'));
  const foreignMap = foreigns.reduce((acc, entry) => Map.insert(ordString)(entry.value0)(entry.value1)(acc), Map.empty);
  const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: Map.empty, pointerAdtPaths: Map.empty,
    pointerAdtNodes: emptySet, pointerAdtLeaves: Map.empty, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: Map.empty, globalFunctions: Map.empty,
    classDeclsFields: Map.empty, ffiFunctions: infos,
  });
  const module = value => ({
    name: owner, bindings: [{ recursive: false, bindings: [new Tuple('action', value)] }],
    comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
    dataTypes: Map.empty, dataDecls: [], classDecls: [], foreign: foreignMap,
    implementations: Map.empty, directives: Map.empty,
  });
  const action = e(new S.Let(new Just('pending'), 0, application('identity'),
    e(new S.EffectBind(new Just('answer'), 1, e(new S.Local(new Just('pending'), 0)),
      e(new S.EffectPure(e(new S.Local(new Just('answer'), 1))))))));
  const emitted = translateWithFunctions(metadata)(module(action));
  assert.match(emitted.code, /RenamedBoundary_Identity_runEffectWorker\(/);
  assert.ok(Map.lookup(ordString)(owner + '.' + key('identity'))(emitted.functions) instanceof Just);
  assert.doesNotMatch(translateWithFunctions(metadata)(module(application(collision))).code,
    /RenamedBoundary_Identity_runEffectWorker\(/);
});

test('workers preserve factory/thunk order, saved action replay, native fallback and callback values', t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-immediate-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, 'gopurs_runtime'));
  copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  const emitted = emit(app(global('identity'), arg('first'))).value0;
  const optimized = flattenStmts(emitted.stmts).map(printGoExpr).join('\n') + '\nreturn ' + printGoExpr(emitted.expr);
  writeFileSync(join(directory, 'bridge.go'), 'package main\nimport "gopurs/output/gopurs_runtime"\n' + prepared.content + '\n' + wrappers + '\n' + workers);
  writeFileSync(join(directory, 'bridge_test.go'), `package main
import("reflect";"testing";"gopurs/output/gopurs_runtime")
func capture(label string) gopurs_runtime.Value { ${owner}_Trace = append(${owner}_Trace,label); return gopurs_runtime.Int(41) }
func immediate() gopurs_runtime.Value { ${optimized} }
func TestImmediate(t *testing.T) {
 for replay:=0;replay<2;replay++ {
  ${owner}_Trace=nil
  if immediate().IntVal!=42 || !reflect.DeepEqual(${owner}_Trace,[]string{"first","factory","effect"}) { t.Fatal(${owner}_Trace) }
 }
 ${owner}_Trace=nil
 saved:=gopurs_runtime.Apply(_Gopurs_${owner}_Identity,gopurs_runtime.Int(41))
 if !reflect.DeepEqual(${owner}_Trace,[]string{"factory"}) { t.Fatal("eager saved action",${owner}_Trace) }
 for i:=0;i<2;i++ { if gopurs_runtime.Apply(saved,gopurs_runtime.Value{}).IntVal!=41 {t.Fatal("replay")} }
 if !reflect.DeepEqual(${owner}_Trace,[]string{"factory","effect","effect"}) {t.Fatal(${owner}_Trace)}
 if ${owner}_Native_runEffectWorker(gopurs_runtime.Int(41)).IntVal!=42 {t.Fatal("native any fallback")}
 payload:=map[string]any{"count":42}
 if !reflect.DeepEqual(gopurs_runtime.ForeignValue(${owner}_Foreign_runEffectWorker(gopurs_runtime.Any(payload))),payload){t.Fatal("Foreign payload")}
 if !reflect.DeepEqual(gopurs_runtime.ForeignValue(${owner}_CurriedForeign_runEffectWorker(gopurs_runtime.Int(0),gopurs_runtime.Any(payload))),payload){t.Fatal("curried Foreign payload")}
 callback:=gopurs_runtime.Func(func(x gopurs_runtime.Value)gopurs_runtime.Value{return gopurs_runtime.Int(x.IntVal+1)})
 if ${owner}_Callback_runEffectWorker(callback,gopurs_runtime.Int(41)).IntVal!=42 {t.Fatal("callback")}
 returned:=${owner}_Returned_runEffectWorker(gopurs_runtime.Int(40))
 if gopurs_runtime.Apply(returned,gopurs_runtime.Int(2)).IntVal!=42 {t.Fatal("Effect returning callback")}
 if gopurs_runtime.Apply(${owner}_Identity_runEffectWorker(callback),gopurs_runtime.Int(41)).IntVal!=42 {t.Fatal("polymorphic callback result")}
 ${owner}_Trace=nil
 func(){defer func(){if recover()!=${owner}_Marker{t.Fatal("panic identity")}}();${owner}_Panic_runEffectWorker(gopurs_runtime.Int(0))}()
 if !reflect.DeepEqual(${owner}_Trace,[]string{"factory","effect"}) {t.Fatal("panic timing",${owner}_Trace)}
}
`);
  const result = spawnSync('go', ['test', '.'], { cwd: directory, encoding: 'utf8', env: { ...process.env, GOWORK: 'off' } });
  assert.equal(result.status, 0, result.stderr || result.stdout);
});
