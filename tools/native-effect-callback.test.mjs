import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import * as Maps from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Effects from '../output/Gopurs.EffectExprs/index.js';
import * as Statements from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { SymbolReference, foreignSymbol } from '../output/Gopurs.GoSymbol/index.js';
import { translate as translateModule } from '../output/Gopurs.CodeGen/index.js';
import { boxedArgument } from '../output/Gopurs.FfiBridge.Values/index.js';
import * as F from '../output/Gopurs.FfiTypes/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const effect = type => new C.ADT('Effect.Effect', ['Effect', 'Effect'], [type]);
const ir = syntax => new TcoExpr(null, syntax);
const typed = (type, value) => ir(new S.Typed(type, value));
const integer = value => typed(C.Int.value, ir(new S.Lit(new C.LitInt(value))));
const action = type => typed(effect(type), ir(new S.EffectPure(integer(1))));
const metadata = withReboxFields({
  pointerAdtPaths: Maps.empty, pointerAdtNodes: emptySet, pointerAdtLeaves: Maps.empty,
  enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet, ctorTypes: Maps.empty,
  classDeclsFields: Maps.empty, globalTypes: Maps.empty, globalFunctions: Maps.empty, ffiFunctions: Maps.empty,
});
const context = () => ({
  metadata, codegenStateRef: { value: { declarations: [], globalId: 0, reboxPairs: emptySet } },
  depth: 0, modNameStr: 'Effects', recVars: [], moduleFunctions: Maps.empty, bound: Maps.empty,
  tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
  mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
});

function generatedAction() {
  let translations = 0;
  const translate = ctx => nextId => _expression => {
    translations++;
    assert.deepEqual(ctx.options, { isTail: false, inEffectBlock: true });
    return { stmts: new Statements.StmtLeaf(Go.rawGo('mark("effect")')),
      expr: Go.rawGo('captured + readCounter()'), exprType: Go.TypeInt64.value, nextId: nextId + 1 };
  };
  const result = Effects.wrapCachedInt(translate)(context())(0)(action(C.Int.value));
  assert.ok(result instanceof Just);
  assert.equal(translations, 1, 'the effect body is translated once');
  assert.equal(result.value0.nextId, 1);
  return printGoExpr(result.value0.expr);
}

const adapter = (native, result = C.Int.value) => boxedArgument([])(new Just(new C.Func([effect(result)], C.Int.value)))(0)(
  new F.TFunc([], new Just(new F.TNamed(native)))).join('\n');

test('cached Effect Int emits independent boxed/native entries from one translation', () => {
  const source = generatedAction();
  assert.match(source, /WithFunctionData/);
  assert.match(source, /func\(\) int64/);
  assert.equal((source.match(/mark\("effect"\)/g) ?? []).length, 2);
  assert.doesNotMatch(source, /FunctionData\[/, 'the ordinary body never dispatches through metadata');
  for (const expr of [action(C.String.value), action(new C.TypeVar('a')),
    ir(new S.EffectPure(integer(1))),
    typed(effect(C.Int.value), ir(new S.Var(new C.Qualified(Nothing.value, 'alias')))),
    typed(effect(C.Int.value), ir(new S.App(ir(new S.Var(new C.Qualified(Nothing.value, 'factory'))), [integer(1)]))),
    typed(effect(C.Int.value), ir(new S.EffectDefer(action(C.Int.value))))]) {
    const forbidden = () => { throw new Error('refused action was translated'); };
    assert.equal(Effects.wrapCachedInt(forbidden)(context())(0)(expr), Nothing.value);
  }
});

test('boxed dynamic action results keep complete Unbox semantics without retranslating', () => {
  const runtimeCall = (name, args = [new Go.GoVar('input')]) => new Go.GoCall(new Go.GoSelector(new Go.GoVar('gopurs_runtime'), name), args);
  for (const [expr, eligible] of [
    [runtimeCall('Int'), true], [runtimeCall('Float'), false], [runtimeCall('Any'), false],
    [Go.rawGo('dynamicResult'), false],
    [new Go.GoCall(new Go.GoSelector(new Go.GoVar('other'), 'Int'), [new Go.GoVar('input')]), false],
    [runtimeCall('Int', []), false],
  ]) {
    let translations = 0;
    const translate = () => nextId => () => {
      translations++;
      return { stmts: Statements.StmtEmpty.value, expr, exprType: Go.TypeValue.value, nextId: nextId + 1 };
    };
    const result = Effects.wrapCachedInt(translate)(context())(0)(action(C.Int.value));
    assert.ok(result instanceof Just);
    assert.equal(translations, 1);
    assert.equal(/WithFunctionData/.test(printGoExpr(result.value0.expr)), eligible);
    assert.doesNotMatch(printGoExpr(result.value0.expr), /IntVal/);
  }
});

function moduleSource() {
  const e = syntax => new NeutralExpr(syntax);
  const ann = (type, value) => e(new S.Typed(type, value));
  const literal = ann(C.Int.value, e(new S.Lit(new C.LitInt(7))));
  const pending = e(new S.App(e(new S.Var(new C.Qualified(new Just('Effects'), 'boundary'))), [literal]));
  const cached = ann(effect(C.Int.value), e(new S.Let(new Just('pending'), 1, pending,
    e(new S.EffectBind(new Just('value'), 2, e(new S.Local(new Just('pending'), 1)), e(new S.EffectPure(literal)))))));
  const alias = ann(effect(C.Int.value), e(new S.Var(new C.Qualified(new Just('Effects'), 'cached'))));
  const local = ann(C.Int.value, e(new S.Local(new Just('capture'), 0)));
  const factory = ann(new C.Func([C.Int.value], effect(C.Int.value)),
    e(new S.Abs([new Tuple(new Just('capture'), 0)], ann(effect(C.Int.value), e(new S.EffectPure(local))))));
  const applied = ann(effect(C.Int.value), e(new S.App(e(new S.Var(new C.Qualified(new Just('Effects'), 'foreignFactory'))), [literal])));
  const letFactory = ann(effect(C.Int.value), e(new S.Let(new Just('saved'), 1, applied,
    e(new S.Local(new Just('saved'), 1)))));
  const ffiFunctions = Maps.insert(ordString)('foreignFactory')({
    reference: new SymbolReference(foreignSymbol('Effects')('Effects_ForeignFactory')),
    fArgs: [Go.TypeInt64.value], fRet: Go.TypeValue.value, arity: 1,
  })(Maps.singleton('$gopursRunEffect$boundary')({
    reference: new SymbolReference(foreignSymbol('Effects')('Effects_Boundary_runEffectWorker')),
    fArgs: [Go.TypeValue.value], fRet: Go.TypeValue.value, arity: 1,
  }));
  return translateModule({ ...metadata, ffiFunctions })({
    name: 'Effects', bindings: [{ recursive: false, bindings: [
      new Tuple('cached', cached), new Tuple('alias', alias), new Tuple('create', factory),
      new Tuple('fromFactory', applied), new Tuple('fromLetFactory', letFactory),
      new Tuple('factoryAlias', ann(effect(C.Int.value), e(new S.Var(new C.Qualified(new Just('Effects'), 'fromFactory'))))),
      new Tuple('pureAction', ann(effect(C.Int.value), e(new S.EffectPure(literal)))),
    ] }], comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
    dataTypes: Maps.empty, dataDecls: [], classDecls: [], foreign: Maps.singleton('foreignFactory')(new Just(new C.Func([C.Int.value], effect(C.Int.value)))),
    implementations: Maps.empty, directives: Maps.empty,
  });
}

test('only module cached values receive metadata, preserving local factories and aliases', () => {
  const source = moduleSource();
  assert.equal((source.match(/WithFunctionData/g) ?? []).length, 2);
  assert.match(source, /func Call_Effects_create\(/);
  assert.match(source, /func Get_Effects_alias\(/);
});

test('cached getters retain factory timing and local actions remain ordinary during repeated construction', t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-cached-effect-factory-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, 'gopurs_runtime'));
  mkdirSync(join(directory, 'purescript'));
  copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(directory, 'purescript/Effects.go'), moduleSource());
  writeFileSync(join(directory, 'purescript/effects_test.go'), `package purescript
import("testing";"gopurs/output/gopurs_runtime")
var factories,effects int
var _Gopurs_Effects_ForeignFactory gopurs_runtime.Value
func Effects_Boundary_runEffectWorker(value gopurs_runtime.Value)gopurs_runtime.Value{effects++;return value}
func Effects_ForeignFactory(value int64)gopurs_runtime.Value{factories++;return gopurs_runtime.Func(func(gopurs_runtime.Value)gopurs_runtime.Value{effects++;return gopurs_runtime.Int(value)})}
func TestGetterTiming(t *testing.T){
 if factories!=0||effects!=0{t.Fatal("eager module initialization")}
 cached:=Get_Effects_cached();alias:=Get_Effects_alias();pure:=Get_Effects_pureAction()
 if factories!=0||effects!=0{t.Fatal("cached construction executed effects")}
 if _,ok:=gopurs_runtime.FunctionData[func()int64](cached);!ok{t.Fatal("cached metadata")}
 if _,ok:=gopurs_runtime.FunctionData[func()int64](alias);!ok{t.Fatal("alias lost metadata")}
 if gopurs_runtime.Apply(pure,gopurs_runtime.Value{}).IntVal!=7||effects!=0{t.Fatal("pure action")}
 fromFactory:=Get_Effects_fromFactory()
 if factories!=1||effects!=0{t.Fatal("factory moved past getter",factories,effects)}
 Get_Effects_fromFactory();factoryAlias:=Get_Effects_factoryAlias()
 if factories!=1{t.Fatal("factory replayed at getter")}
 for _,action:=range []gopurs_runtime.Value{fromFactory,factoryAlias}{
  if _,ok:=gopurs_runtime.FunctionData[func()int64](action);ok{t.Fatal("factory result gained metadata")}
  if gopurs_runtime.Apply(action,gopurs_runtime.Value{}).IntVal!=7{t.Fatal("saved factory result")}
 }
 if factories!=1||effects!=2{t.Fatal("saved replay",factories,effects)}
 fromLet:=Get_Effects_fromLetFactory()
 if factories!=2||effects!=2{t.Fatal("let factory moved",factories,effects)}
 gopurs_runtime.Apply(fromLet,gopurs_runtime.Value{})
 for round:=0;round<10;round++{
  actions:=make([]gopurs_runtime.Value,0,4)
  for capture:=int64(0);capture<4;capture++{actions=append(actions,Call_Effects_create(capture))}
  for capture,action:=range actions{
   if _,ok:=gopurs_runtime.FunctionData[func()int64](action);ok{t.Fatal("local action allocation changed")}
   for replay:=0;replay<2;replay++{if got:=gopurs_runtime.Apply(action,gopurs_runtime.Value{}).IntVal;got!=int64(capture){t.Fatal("capture",got,capture)}}
  }
 }
}
`);
  const result = spawnSync('go', ['test', '-count=1', './purescript'], { cwd: directory, encoding: 'utf8', timeout: 60000,
    env: { ...process.env, GOWORK: 'off' } });
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stdout + result.stderr);
});

test('FFI lookup selects only zero-argument native int/int64 callbacks with nominal Effect Int', () => {
  for (const native of ['int', 'int64']) assert.match(adapter(native), /FunctionData\[func\(\) int64\]/);
  for (const [native, result] of [['any', C.Int.value], ['interface{}', C.Int.value], ['gopurs_runtime.Value', C.Int.value],
    ['int', new C.TypeVar('a')], ['bool', C.Boolean.value], ['int32', C.Int.value]]) {
    assert.doesNotMatch(adapter(native, result), /FunctionData/);
  }
});

test('native callbacks preserve deferred execution, replay, captures, fallback and polymorphic results', t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-native-effect-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, 'gopurs_runtime'));
  copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(directory, 'main_test.go'), `package main
import("fmt";"reflect";"testing";"gopurs/output/gopurs_runtime")
var trace []string
var counter int64
func mark(label string){trace=append(trace,label)}
func readCounter()int64{counter++;return counter}
func create(captured int64)gopurs_runtime.Value{return ${generatedAction()}}
func adaptInt(arg0 gopurs_runtime.Value)func()int{${adapter('int')};return go_arg0}
func adaptInt64(arg0 gopurs_runtime.Value)func()int64{${adapter('int64')};return go_arg0}
func adaptAny(arg0 gopurs_runtime.Value)func()any{${adapter('any', new C.TypeVar('a'))};return go_arg0}
func TestEntries(t *testing.T){
 first,second:=create(10),create(100)
 directInt64,ok:=gopurs_runtime.FunctionData[func()int64](first)
 if !ok{t.Fatal("missing int64 metadata")}
 a,b:=adaptInt(first),adaptInt64(second)
 if ^uint(0)>>63==1 && reflect.ValueOf(a).Pointer()!=reflect.ValueOf(directInt64).Pointer(){t.Fatal("adaptInt wrapped direct func() int64 in a trampoline")}
 if len(trace)!=0||counter!=0{t.Fatal("eager action",trace,counter)}
 for _,step:=range []struct{call func()int64;want int64}{
  {func()int64{return int64(a())},11},{b,102},
  {func()int64{return gopurs_runtime.Apply(first,gopurs_runtime.Value{}).IntVal},13},
  {func()int64{return int64(a())},14},{b,105},
 }{if got:=step.call();got!=step.want{t.Fatal(got,step.want)}}
 if !reflect.DeepEqual(trace,[]string{"effect","effect","effect","effect","effect"}){t.Fatal(trace)}
 nativeOnly:=gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(gopurs_runtime.Value)gopurs_runtime.Value{panic("fallback executed")}),func()int64{return 42})
 if adaptInt(nativeOnly)()!=42||adaptInt64(nativeOnly)()!=42{t.Fatal("native entry")}
 ordinary:=gopurs_runtime.Func(func(gopurs_runtime.Value)gopurs_runtime.Value{return gopurs_runtime.Int(17)})
 wrong:=gopurs_runtime.WithFunctionData(ordinary,func()string{panic("wrong metadata")})
 if adaptInt(ordinary)()!=17||adaptInt64(wrong)()!=17{t.Fatal("fallback")}
 returned:=adaptAny(ordinary)()
 if value,ok:=returned.(gopurs_runtime.Value);!ok||value.IntVal!=17{t.Fatalf("polymorphic ABI: %T %v",returned,returned)}
 marker:=&struct{}{}
 failing:=gopurs_runtime.WithFunctionData(gopurs_runtime.Func(func(gopurs_runtime.Value)gopurs_runtime.Value{panic(marker)}),func()int64{panic(marker)})
 for _,call:=range []func(){func(){adaptInt(failing)()},func(){gopurs_runtime.Apply(failing,gopurs_runtime.Value{})}}{
  func(){defer func(){if recover()!=marker{t.Fatal("panic identity")}}();call()}()
 }
 invalid:=gopurs_runtime.Value{Type:gopurs_runtime.TypeFunc}
 for _,value:=range []gopurs_runtime.Value{invalid,gopurs_runtime.WithFunctionData(invalid,func()int64{return 0})}{
  func(){defer func(){want:=fmt.Sprintf("PANIC in Apply: f.Type = %v, f.UnsafePtr is nil",gopurs_runtime.TypeFunc);if recover()!=want{t.Fatal("nil panic changed")}}();gopurs_runtime.Apply(value,gopurs_runtime.Value{})}()
 }
}
`);
  const result = spawnSync('go', ['test', '-count=1', '.'], { cwd: directory, encoding: 'utf8', timeout: 60000,
    env: { ...process.env, GOWORK: 'off' } });
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stdout + result.stderr);
});
