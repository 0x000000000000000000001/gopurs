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
import { key, factoryKey, statelessFactoryKey } from '../output/Gopurs.FfiBridge.ImmediateEffects/index.js';
import { prepareFfi } from '../output/Gopurs.FfiSupport/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { StmtEmpty, StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import { translate, direct, bindThunk, invokeBoundThunk } from '../output/Gopurs.ImmediateFfiEffects/index.js';
import { printGoDecl, printGoExpr } from '../output/Gopurs.Printer/index.js';
import { literal as primitiveLiteral } from '../output/Gopurs.PrimitiveExprs/index.js';
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
  ['stateless', new C.ForAll(['a'], fn([variable], effect(variable)))],
  ['choose', new C.ForAll(['a'], fn([int, variable, variable], effect(variable)))],
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
var Last any
func Identity(x any) func() any { Last = x; Trace = append(Trace, "factory"); return func() any { Trace = append(Trace, "effect"); return x } }
func Stateless(x any) func() any { return func() any { Trace = append(Trace, "stateless"); return x } }
func Choose(count int64, x any, y interface{}) func() interface{} { Trace = append(Trace, "factory"); return func() interface{} { Trace = append(Trace, "effect"); if count == 0 { return x }; return y } }
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
const emitDirect = (factory, continuation = pure(arg('result')), ctx = context) =>
  direct(translator)(ctx)(0)(new Just('answer'))(1)(factory)(continuation);
const cacheContext = (module = owner) => ({ ...context, modNameStr: module,
  codegenStateRef: { value: { declarations: [], globalId: 0, reboxPairs: emptySet } },
  metadata: module === owner ? context.metadata : {
    ffiFunctions: Map.empty, globalFunctions: Map.singleton(owner + '.' + key('identity'))(lookup(key('identity')).value0),
  },
});
const scalarTranslator = () => id => expr => {
  let syntax = expr.value1;
  while (syntax instanceof S.Typed) syntax = syntax.value1.value1;
  const primitive = syntax instanceof S.Lit ? primitiveLiteral(syntax.value0) : Nothing.value;
  return { stmts: StmtEmpty.value, nextId: id + 1,
    ...(primitive instanceof Just ? primitive.value0 : syntax instanceof S.Local ?
      { expr: Go.rawGo('input'), exprType: Go.TypeInt64.value } :
      { expr: Go.rawGo('gopurs_runtime.Int(42)'), exprType: Go.TypeValue.value }),
  };
};
const emitArgument = (argument, ctx, translateArgument = scalarTranslator, target = 'identity') =>
  translate(translateArgument)(ctx)(0)(new Just('action'))(0)(app(global(target), argument))(bind(pure(local('answer', 1)))).value0;
const renderBody = emitted => flattenStmts(emitted.stmts).map(printGoExpr).join('\n') + '\nreturn ' + printGoExpr(emitted.expr);

test('nominal effect signatures retain polymorphism, native any, Foreign and callback results', () => {
  for (const name of ['identity', 'native', 'foreign', 'curriedForeign', 'callback', 'returned', 'panic']) {
    assert.ok(lookup(key(name)) instanceof Just, name);
    assert.equal(lookup(key(name)).value0.fRet, Go.TypeValue.value);
  }
  for (const name of ['ordinary', 'saved']) assert.equal(lookup(key(name)), Nothing.value);
  assert.match(workers, /return gopurs_runtime.Box\(go_result\)/);
  assert.match(workers, /gopurs_runtime.ForeignValue\(arg0\)/);
  assert.match(workers, /return gopurs_runtime.Any\(go_result\)/);
  assert.deepEqual(lookup(key('identity')).value0.fArgs, [new Go.TypeInterface('interface{}')]);
  assert.deepEqual(lookup(key('choose')).value0.fArgs, [Go.TypeValue.value, new Go.TypeInterface('interface{}'), new Go.TypeInterface('interface{}')]);
  for (const name of ['native', 'foreign', 'curriedForeign', 'callback']) {
    assert.ok(lookup(key(name)).value0.fArgs.every(type => type === Go.TypeValue.value), name);
  }
});

test('only statement-free scalar literals cache their canonical Value-to-interface conversion', () => {
  for (const [literal, nativeCall] of [
    [new C.LitInt(41), /gopurs_runtime.Int\(int64\(41\)\)/],
    [new C.LitBoolean(true), /gopurs_runtime.Bool\(true\)/],
    [new C.LitString('hello'), /gopurs_runtime.StrWithCodeUnits\("hello", 5\)/],
    [new C.LitChar('x'), /gopurs_runtime.StrWithCodeUnits\("x", 1\)/],
    [new C.LitNumber(-0), /gopurs_runtime.NegativeZero\(\)/],
    [new C.LitNumber(NaN), /math.NaN\(\)/],
    [new C.LitNumber(Infinity), /math.Inf\(1\)/],
  ]) {
    const ctx = cacheContext();
    const emitted = emitArgument(ir(new S.Typed(variable, ir(new S.Lit(literal)))), ctx);
    const [declaration] = ctx.codegenStateRef.value.declarations;
    assert.ok(declaration instanceof Go.GoCachedValue);
    assert.deepEqual(declaration.value0.goType, new Go.TypeInterface('interface{}'));
    const code = printGoDecl(declaration);
    assert.match(code, nativeCall);
    assert.doesNotMatch(code, /Identity|Apply|runEffect/);
    assert.match(renderBody(emitted), /Get___gopurs_ffi_argument_RenamedBoundary_0\(\)/);
  }
  for (const argument of [local('changing'), global('changing'), ir(new S.Lit(new C.LitArray([]))),
    ir(new S.Lit(new C.LitRecord([]))), app(global('identity'), arg('computed'))]) {
    const ctx = cacheContext();
    emitArgument(argument, ctx);
    assert.deepEqual(ctx.codegenStateRef.value.declarations, []);
  }
  const ctx = cacheContext();
  const withStatements = () => id => expression => ({ ...scalarTranslator()(id)(expression), stmts: new StmtLeaf(Go.rawGo('observe()')) });
  emitArgument(ir(new S.Lit(new C.LitInt(41))), ctx, withStatements);
  emitArgument(ir(new S.Lit(new C.LitInt(41))), ctx, scalarTranslator, 'native');
  assert.deepEqual(ctx.codegenStateRef.value.declarations, []);
});

test('cache symbols are unique across functions, literal types and modules', () => {
  const names = [];
  for (const module of [owner, 'OtherModule']) {
    const ctx = cacheContext(module);
    // Each call starts with nextId=0, as separate function bodies can do.
    for (const literal of [new C.LitInt(41), new C.LitInt(99), new C.LitString('same site')]) {
      emitArgument(ir(new S.Lit(literal)), ctx);
    }
    names.push(...ctx.codegenStateRef.value.declarations.map(declaration => declaration.value0.identifier.name));
  }
  assert.equal(new Set(names).size, 6);
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

test('direct EffectBind factories use the same certified worker and strict argument captures', () => {
  const factory = app(global('choose'), arg('count'), arg('left'), arg('right'));
  const result = emitDirect(factory);
  assert.ok(result instanceof Just);
  assert.equal(renderBody(result.value0), renderBody(emit(factory).value0));
  for (const value of [global('identity'), app(global('identity'), arg('first'), arg('extra')),
    app(local('saved'), arg('first')), local('saved'), app(global('ordinary'), arg('first')),
    app(global('saved'), arg('first')), ir(new S.UncurriedApp(global('identity'), [arg('first')])),
    app(ir(new S.UncurriedApp(global('identity'), [])), arg('first')),
    ir(new S.EffectDefer(app(global('identity'), arg('first'))))]) {
    assert.equal(emitDirect(value), Nothing.value);
  }
  assert.equal(emitDirect(factory, undefined, { ...context, options: { isTail: false, inEffectBlock: false } }), Nothing.value);
  assert.ok(emitDirect(app(ir(new S.TypeApp(global('identity'), int)), arg('first'))) instanceof Just);
  const functions = Map.singleton(owner + '.' + key('identity'))(lookup(key('identity')).value0);
  assert.ok(emitDirect(app(global('identity'), arg('first')), undefined,
    { ...context, modNameStr: 'Consumer', metadata: { ffiFunctions: Map.empty, globalFunctions: functions } }) instanceof Just);
  assert.equal(emitDirect(app(global('identity'), arg('first')), undefined,
    { ...context, modNameStr: 'Consumer', metadata: { ffiFunctions: Map.empty, globalFunctions: Map.empty } }), Nothing.value);
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
    name: owner, bindings: [{ recursive: false, bindings: [new Tuple('ffi_argument_0', literal), new Tuple('action', value)] }],
    comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
    dataTypes: Map.empty, dataDecls: [], classDecls: [], foreign: foreignMap,
    implementations: Map.empty, directives: Map.empty,
  });
  const action = e(new S.Let(new Just('pending'), 0, application('identity'),
    e(new S.EffectBind(new Just('answer'), 1, e(new S.Local(new Just('pending'), 0)),
      e(new S.EffectPure(e(new S.Local(new Just('answer'), 1))))))));
  const emitted = translateWithFunctions(metadata)(module(action));
  const directAction = e(new S.EffectBind(new Just('answer'), 1, application('identity'),
    e(new S.EffectPure(e(new S.Local(new Just('answer'), 1))))));
  const directCode = translateWithFunctions(metadata)(module(directAction)).code;
  assert.match(directCode, /gopurs_runtime\.Box\(RenamedBoundary_Identity\(/);
  assert.match(directCode, /Get___gopurs_ffi_argument_RenamedBoundary_\d+\(/);
  assert.doesNotMatch(directCode, /RenamedBoundary_Identity_runEffectWorker\(/);
  assert.doesNotMatch(directCode, /Apply2\(Get_RenamedBoundary_identity/);

  assert.match(emitted.code, /gopurs_runtime\.Box\(RenamedBoundary_Identity\(/);
  assert.match(emitted.code, /func Get_RenamedBoundary_ffi_argument_0\(/);
  assert.match(emitted.code, /func Get___gopurs_ffi_argument_RenamedBoundary_\d+\(/);
  assert.ok(Map.lookup(ordString)(owner + '.' + key('identity'))(emitted.functions) instanceof Just);
  assert.ok(Map.lookup(ordString)(owner + '.' + factoryKey('identity'))(emitted.functions) instanceof Just);
  assert.doesNotMatch(translateWithFunctions(metadata)(module(application(collision))).code,
    /RenamedBoundary_Identity_runEffectWorker\(/);
  assert.doesNotMatch(translateWithFunctions(metadata)(module(application(Go.sanitizeName(factoryKey('identity'))))).code,
    /RenamedBoundary_Identity\(/);

  // Complex adapters (Foreign) still route through their certified _runEffectWorker.
  const foreignAction = e(new S.EffectBind(new Just('answer'), 1, application('foreign'),
    e(new S.EffectPure(e(new S.Local(new Just('answer'), 1))))));
  assert.match(translateWithFunctions(metadata)(module(foreignAction)).code, /RenamedBoundary_Foreign_runEffectWorker\(/);

  // Unused result skips Box while still invoking factory and thunk.
  const discardedAction = e(new S.EffectBind(new Just('unused'), 1, application('identity'),
    e(new S.EffectPure(literal))));
  const discardedCode = translateWithFunctions(metadata)(module(discardedAction)).code;
  assert.match(discardedCode, /_ = RenamedBoundary_Identity\(__ffi_effect_arg_\d+\)\(\)/);
  assert.doesNotMatch(discardedCode, /gopurs_runtime\.Box\(RenamedBoundary_Identity/);

  // Let-bound Effect factory hoisted across EffectPure / replay keeps its native thunk without Apply.
  const hoistedAction = e(new S.Let(new Just('pending'), 0, application('identity'),
    e(new S.EffectPure(
      e(new S.EffectBind(new Just('first'), 1, e(new S.Local(new Just('pending'), 0)),
        e(new S.EffectBind(new Just('second'), 2, e(new S.Local(new Just('pending'), 0)),
          e(new S.EffectPure(e(new S.Local(new Just('second'), 2))))))))))));
  const hoistedCode = translateWithFunctions(metadata)(module(hoistedAction)).code;
  assert.match(hoistedCode, /pending_0_0 := RenamedBoundary_Identity\(/);
  assert.match(hoistedCode, /_ = pending_0_0\(\)/);
  assert.match(hoistedCode, /second_2_\d+ := gopurs_runtime\.Box\(pending_0_0\(\)\)/);
  assert.doesNotMatch(hoistedCode, /gopurs_runtime\.Apply/);

  // Escaping Let-bound Effect factory retains its boxed Apply representation.
  const escapedAction = e(new S.Let(new Just('pending'), 0, application('identity'),
    e(new S.EffectBind(new Just('first'), 1, e(new S.Local(new Just('pending'), 0)),
      e(new S.EffectPure(e(new S.Local(new Just('pending'), 0))))))));
  const escapedCode = translateWithFunctions(metadata)(module(escapedAction)).code;
  assert.match(escapedCode, /gopurs_runtime\.Apply\(Get_RenamedBoundary_identity\(\)/);

  // Stateless Effect factory sinks into EffectBind without hoisting a thunk closure.
  assert.ok(lookup(statelessFactoryKey('stateless')) instanceof Just);
  assert.equal(lookup(statelessFactoryKey('identity')), Nothing.value);
  const hoistedStateless = e(new S.Let(new Just('pending'), 0, application('stateless'),
    e(new S.EffectPure(
      e(new S.EffectBind(new Just('first'), 1, e(new S.Local(new Just('pending'), 0)),
        e(new S.EffectBind(new Just('second'), 2, e(new S.Local(new Just('pending'), 0)),
          e(new S.EffectPure(e(new S.Local(new Just('second'), 2))))))))))));
  const hoistedStatelessCode = translateWithFunctions(metadata)(module(hoistedStateless)).code;
  assert.doesNotMatch(hoistedStatelessCode, /pending_0_0 := RenamedBoundary_Stateless\(/);
  assert.match(hoistedStatelessCode, /_ = RenamedBoundary_Stateless\(__ffi_effect_arg_\d+\)\(\)/);
  assert.match(hoistedStatelessCode, /second_2_\d+ := gopurs_runtime\.Box\(RenamedBoundary_Stateless\(__ffi_effect_arg_\d+\)\(\)\)/);
  assert.doesNotMatch(hoistedStatelessCode, /gopurs_runtime\.Apply/);
});

test('hoisted Let-bound native thunks execute factory once and replay thunk without runtime Apply', t => {
  const e = syntax => new NeutralExpr(syntax);
  const target = name => e(new S.Var(new C.Qualified(new Just(owner), name)));
  const literal = n => e(new S.Typed(int, e(new S.Lit(new C.LitInt(n)))));
  const foreignMap = foreigns.reduce((acc, entry) => Map.insert(ordString)(entry.value0)(entry.value1)(acc), Map.empty);
  const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: Map.empty, pointerAdtPaths: Map.empty,
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMapOrEmpty(), enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: Map.empty, globalFunctions: Map.empty,
    classDeclsFields: Map.empty, ffiFunctions: infos,
  });
  function emptyMapOrEmpty() { return Map.empty; }
  const hoisted = e(new S.Let(new Just('pending'), 0, e(new S.App(target('identity'), [literal(41)])),
    e(new S.EffectPure(
      e(new S.EffectBind(new Just('unused'), 1, e(new S.Local(new Just('pending'), 0)),
        e(new S.EffectBind(new Just('answer'), 2, e(new S.Local(new Just('pending'), 0)),
          e(new S.EffectPure(e(new S.Local(new Just('answer'), 2))))))))))));
  const hoistedStateless = e(new S.Let(new Just('pending'), 0, e(new S.App(target('stateless'), [literal(99)])),
    e(new S.EffectPure(
      e(new S.EffectBind(new Just('unused'), 1, e(new S.Local(new Just('pending'), 0)),
        e(new S.EffectBind(new Just('answer'), 2, e(new S.Local(new Just('pending'), 0)),
          e(new S.EffectPure(e(new S.Local(new Just('answer'), 2))))))))))));
  const panicUnused = e(new S.EffectBind(new Just('unused'), 1, e(new S.App(target('panic'), [literal(0)])),
    e(new S.EffectPure(literal(7)))));
  const mod = {
    name: owner,
    bindings: [{ recursive: false, bindings: [new Tuple('hoisted', hoisted), new Tuple('hoistedStateless', hoistedStateless), new Tuple('panicUnused', panicUnused)] }],
    comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
    dataTypes: Map.empty, dataDecls: [], classDecls: [], foreign: foreignMap,
    implementations: Map.empty, directives: Map.empty,
  };
  const emitted = translateWithFunctions(metadata)(mod);
  assert.doesNotMatch(emitted.code, /gopurs_runtime\.Apply/);
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-bound-thunk-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, 'gopurs_runtime'));
  mkdirSync(join(directory, 'purescript'));
  copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(directory, 'purescript/mod.go'), emitted.code);
  writeFileSync(join(directory, 'purescript/ffi.go'), 'package purescript\nimport "gopurs/output/gopurs_runtime"\n' + prepared.content + '\n' + wrappers + '\n' + workers);
  writeFileSync(join(directory, 'main_test.go'), `package main
import("reflect";"testing";"gopurs/output/gopurs_runtime";p "gopurs/output/purescript")
func TestHoistedThunkReplay(t *testing.T) {
  p.${owner}_Trace = nil
  inner := gopurs_runtime.Apply(p.Get_${owner}_hoisted(), gopurs_runtime.Value{})
  if !reflect.DeepEqual(p.${owner}_Trace, []string{"factory"}) { t.Fatal("factory timing", p.${owner}_Trace) }
  for i := 0; i < 2; i++ {
    if got := gopurs_runtime.Apply(inner, gopurs_runtime.Value{}).IntVal; got != 41 { t.Fatal("replay value", got) }
  }
  if !reflect.DeepEqual(p.${owner}_Trace, []string{"factory", "effect", "effect", "effect", "effect"}) { t.Fatal("replay trace", p.${owner}_Trace) }
  p.${owner}_Trace = nil
  innerStateless := gopurs_runtime.Apply(p.Get_${owner}_hoistedStateless(), gopurs_runtime.Value{})
  if len(p.${owner}_Trace) != 0 { t.Fatal("stateless outer timing", p.${owner}_Trace) }
  for i := 0; i < 2; i++ {
    if got := gopurs_runtime.Apply(innerStateless, gopurs_runtime.Value{}).IntVal; got != 99 { t.Fatal("stateless replay value", got) }
  }
  if !reflect.DeepEqual(p.${owner}_Trace, []string{"stateless", "stateless", "stateless", "stateless"}) { t.Fatal("stateless replay trace", p.${owner}_Trace) }
  p.${owner}_Trace = nil
  func() {
    defer func() { if recover() != p.${owner}_Marker { t.Fatal("discarded panic identity") } }()
    gopurs_runtime.Apply(p.Get_${owner}_panicUnused(), gopurs_runtime.Value{})
  }()
  if !reflect.DeepEqual(p.${owner}_Trace, []string{"factory", "effect"}) { t.Fatal("discarded panic trace", p.${owner}_Trace) }
}
`);
  const result = spawnSync('go', ['test', '-count=1', '.'], { cwd: directory, encoding: 'utf8', env: { ...process.env, GOWORK: 'off' } });
  assert.equal(result.status, 0, result.stderr || result.stdout);
});

test('workers preserve factory/thunk order, saved action replay, native fallback and callback values', t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-immediate-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, 'gopurs_runtime'));
  copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  const emitted = emit(app(global('identity'), arg('first'))).value0;
  const optimized = renderBody(emitted);
  const ordered = renderBody(emit(app(global('choose'), arg('count'), arg('left'), arg('right'))).value0);
  const directOrdered = renderBody(emitDirect(app(global('choose'), arg('count'), arg('left'), arg('right'))).value0);
  const ctx = cacheContext();
  const literals = [new C.LitInt(41), new C.LitInt(99), new C.LitString('retained'), new C.LitNumber(-0), new C.LitNumber(NaN)];
  const cached = literals.map(literal => renderBody(emitArgument(ir(new S.Lit(literal)), ctx)));
  const dynamic = renderBody(emitArgument(local('changing'), ctx));
  const declarations = ctx.codegenStateRef.value.declarations;
  const getters = declarations.map(declaration => 'Get_' + declaration.value0.identifier.name + '()');
  const printedDeclarations = declarations.map(printGoDecl).join('\n');
  const syncImport = printedDeclarations.includes('sync.Once') ? '"sync";' : '';
  writeFileSync(join(directory, 'bridge.go'), `package main\nimport(${syncImport}"math";"gopurs/output/gopurs_runtime")\n` +
    prepared.content + '\n' + wrappers + '\n' + workers + '\n' + printedDeclarations + '\n' +
    cached.map((body, index) => `func cached${index}() gopurs_runtime.Value { ${body} }`).join('\n') +
    `\nfunc dynamic(input int64) gopurs_runtime.Value { ${dynamic} }`);
  writeFileSync(join(directory, 'bridge_test.go'), `package main
import("math";"reflect";"testing";"gopurs/output/gopurs_runtime")
func capture(label string) gopurs_runtime.Value { ${owner}_Trace = append(${owner}_Trace,label); return gopurs_runtime.Int(41) }
func immediate() gopurs_runtime.Value { ${optimized} }
func ordered() gopurs_runtime.Value { ${ordered} }
func directOrdered() gopurs_runtime.Value { ${directOrdered} }
func TestImmediate(t *testing.T) {
 for replay:=0;replay<2;replay++ {
  ${owner}_Trace=nil
  if immediate().IntVal!=42 || !reflect.DeepEqual(${owner}_Trace,[]string{"first","factory","effect"}) { t.Fatal(${owner}_Trace) }
 }
 ${owner}_Trace=nil;ordered()
 if !reflect.DeepEqual(${owner}_Trace,[]string{"count","left","right","factory","effect"}) {t.Fatal("argument order",${owner}_Trace)}
 ${owner}_Trace=nil;directOrdered()
 if !reflect.DeepEqual(${owner}_Trace,[]string{"count","left","right","factory","effect"}) {t.Fatal("direct argument order",${owner}_Trace)}
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
func TestCachedArguments(t *testing.T) {
 ${owner}_Trace=nil
 initial:=[]any{${getters.join(',')}}
 if len(${owner}_Trace)!=0 {t.Fatal("cache construction executed a factory",${owner}_Trace)}
 for index,call:=range []func()gopurs_runtime.Value{cached0,cached1,cached2,cached3,cached4} {
  value,ok:=initial[index].(gopurs_runtime.Value)
  if !ok {t.Fatalf("literal %d lost canonical Value ABI: %T",index,initial[index])}
  switch index {
  case 0: if value.IntVal!=41 {t.Fatal(value)}
  case 1: if value.IntVal!=99 {t.Fatal(value)}
  case 2: if value.StrVal()!="retained" {t.Fatal(value)}
  case 3: if !math.Signbit(value.FloatVal()) || value.FloatVal()!=0 {t.Fatal("negative zero",value)}
  case 4: if !math.IsNaN(value.FloatVal()) {t.Fatal("NaN",value)}
  }
  for replay:=0;replay<2;replay++ {
   ${owner}_Trace=nil;call()
   if ${owner}_Last!=initial[index] {t.Fatal("the converted literal was not reused",index)}
   if !reflect.DeepEqual(${owner}_Trace,[]string{"factory","effect"}) {t.Fatal("effects were cached",${owner}_Trace)}
  }
 }
 for _,input:=range []int64{13,27,-9} {
  ${owner}_Trace=nil;dynamic(input)
  value,ok:=${owner}_Last.(gopurs_runtime.Value)
  if !ok || value.IntVal!=input {t.Fatal("dynamic argument was cached or changed ABI",${owner}_Last)}
  if !reflect.DeepEqual(${owner}_Trace,[]string{"factory","effect"}) {t.Fatal(${owner}_Trace)}
 }
}
`);
  const result = spawnSync('go', ['test', '.'], { cwd: directory, encoding: 'utf8', env: { ...process.env, GOWORK: 'off' } });
  assert.equal(result.status, 0, result.stderr || result.stdout);
});
