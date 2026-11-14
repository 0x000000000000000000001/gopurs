import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { generateFfiBridge } from '../output/Gopurs.FfiBridge/index.js';
import { prepareFfi } from '../output/Gopurs.FfiSupport/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
const foreign = new C.ADT('Foreign.Foreign', ['Foreign', 'Foreign'], []);
const unit = C.Unit.value;
const str = C.String.value;
const effect = new C.ADT('Effect.Effect', ['Effect', 'Effect'], [unit]);
const eff1 = new C.ADT('Effect.Uncurried.EffectFn1', ['Effect', 'Uncurried', 'EffectFn1'], [str, unit]);
const fn = (args, ret) => new C.Func(args, ret);
const countRecord = new C.Record(new C.Row([new Tuple('count', C.Int.value)], Nothing.value));
const handle = new C.ADT('Fixture.Handle', ['Fixture', 'Handle'], []);
const source = `package fixture
func Dynamic(v any) any { return v }
func Dynamics(v []any) []any { return v }
func EchoForeignCallback(cb func(any) any, value any) any { return cb(value) }
func Callback(cb func(string)) { cb("typed") }
func Tick(cb func()) { cb() }
func LegacyBind(a func(any) any, f func(any) func(any) any, unit any) any { return f(a(nil))(unit) }
func LegacyEffect(unit any) int64 { return 42 }
func GenericFn[R any](callback func(string, string) R) R { return callback("left", "right") }
type handle struct { closed bool }
func Handle(v *handle) *handle { return v }
func ForeignCallback(cb func(int64) any) bool { _, ok := cb(1).(string); return ok }
func ForeignEffectCallback(cb func() any) bool { _, ok := cb().(string); return ok }
func Matrix(values [][]int64) [][]int64 { return values }
type Payload struct { Count int64 }
func StructCallback(cb func(int64) Payload) int64 { return cb(1).Count }
func StructArrayCallback(cb func() []Payload) int64 { return cb()[0].Count }
func ForeignArrayCallback(cb func() []any) bool { _, ok := cb()[0].(string); return ok }
func LegacyArrayCallback(cb func() []any) []any { return cb() }
`;
const foreigns = [
  ['dynamic', fn([foreign], foreign)], ['dynamics', fn([new C.Array(foreign)], new C.Array(foreign))],
  ['echoForeignCallback', fn([fn([foreign], foreign), foreign], foreign)],
  ['callback', fn([fn([str], effect)], effect)], ['uncurried', fn([eff1], effect)],
  ['tick', fn([effect], effect)], ['handle', fn([handle], handle)],
  ['legacyBind', fn([effect, fn([unit], effect)], effect)],
  ['foreignCallback', fn([fn([C.Int.value], foreign)], C.Boolean.value)],
  ['foreignEffectCallback', fn([new C.ADT('Effect.Effect', ['Effect', 'Effect'], [foreign])], new C.ADT('Effect.Effect', ['Effect', 'Effect'], [C.Boolean.value]))],
  ['matrix', fn([new C.Array(new C.Array(C.Int.value))], new C.Array(new C.Array(C.Int.value)))],
  ['structCallback', fn([fn([C.Int.value], countRecord)], C.Int.value)],
  ['structArrayCallback', fn([new C.ADT('Effect.Effect', ['Effect', 'Effect'], [new C.Array(countRecord)])], new C.ADT('Effect.Effect', ['Effect', 'Effect'], [C.Int.value]))],
  ['foreignArrayCallback', fn([new C.ADT('Effect.Effect', ['Effect', 'Effect'], [new C.Array(foreign)])], new C.ADT('Effect.Effect', ['Effect', 'Effect'], [C.Boolean.value]))],
  ['legacyArrayCallback', fn([new C.ADT('Effect.Effect', ['Effect', 'Effect'], [new C.Array(new C.TypeVar('a'))])], new C.ADT('Effect.Effect', ['Effect', 'Effect'], [new C.Array(new C.TypeVar('a'))]))],
  ['legacyEffect', new C.ADT('Effect.Effect', ['Effect', 'Effect'], [C.Int.value])],
  ['genericFn', new C.ForAll(['a'], fn([new C.ADT('Data.Function.Uncurried.Fn2', ['Data', 'Function', 'Uncurried', 'Fn2'], [str, str, new C.TypeVar('a')])], new C.TypeVar('a')))],
].map(([name, type]) => new Tuple(name, new Just(type)));

test('dynamic values, native handles and typed callbacks cross without handwritten runtime access', t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-values-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  const parsed = prepareFfi({ moduleName: 'Fixture', path: 'fixture.go' })('Fixture_')(source + '\nfunc Uncurried(cb func(string)) { cb("uncurried") }')();
  mkdirSync(join(directory, 'gopurs_runtime'));
  copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\ngo 1.27.0\n');
  writeFileSync(join(directory, 'bridge.go'), 'package fixture\nimport "gopurs/output/gopurs_runtime"\n' + parsed.content + '\n' + generateFfiBridge('Fixture')([])(parsed.decls)(foreigns));
  writeFileSync(join(directory, 'bridge_test.go'), `package fixture
import (
 "testing"
 "reflect"
 rt "gopurs/output/gopurs_runtime"
)
func TestDynamic(t *testing.T) {
 values := []rt.Value{rt.Any(nil), rt.Bool(false), rt.Int(0), rt.Int(9007199254740993), rt.Box(1.25), rt.Str(""), rt.Array([]rt.Value{}), rt.Record(map[string]rt.Value{"nil":rt.Any(nil), "false":rt.Bool(false)})}
 for _, value := range values {
   want := rt.ForeignValue(value)
   got := rt.ForeignValue(rt.Apply(_Gopurs_Fixture_Dynamic, value))
   if !reflect.DeepEqual(got, want) { t.Fatalf("%T %#v != %T %#v", got, got, want, want) }
 }
 result := rt.Apply(_Gopurs_Fixture_Dynamics, rt.Array(values))
 if !reflect.DeepEqual(rt.ForeignValue(result), rt.ForeignValue(rt.Array(values))) { t.Fatal("array lost values") }
 var nilSlice []any
 if got := rt.ForeignValue(rt.Apply(_Gopurs_Fixture_Dynamic, rt.Any(nilSlice))); !reflect.ValueOf(got).IsNil() { t.Fatal("nil slice became empty") }
 raw := map[string]any{"huge": uint64(18446744073709551615)}
 got := rt.ForeignValue(rt.Apply(_Gopurs_Fixture_Dynamic, rt.Any(raw))).(map[string]any)
 if got["huge"] != raw["huge"] { t.Fatal("integer precision") }
 if _, exists := got["absent"]; exists { t.Fatal("missing key materialized") }
}
func TestNativeCallbackResults(t *testing.T) {
 callback := rt.Func(func(rt.Value) rt.Value { return rt.Str("native") })
 if !rt.Unbox[bool](rt.Apply(_Gopurs_Fixture_ForeignCallback, callback)) { t.Fatal("Foreign callback leaked Value") }
 action := rt.Apply(_Gopurs_Fixture_ForeignEffectCallback, callback)
 if !rt.Unbox[bool](rt.Apply(action, rt.Value{})) { t.Fatal("Effect Foreign callback leaked Value") }
 record := rt.Func(func(rt.Value) rt.Value { return rt.RecordDict1("count", rt.Int(42)) })
 if value := rt.Apply(_Gopurs_Fixture_StructCallback, record); value.IntVal != 42 { t.Fatal("record callback was reinterpreted as a Go struct", value) }
 array := rt.Func(func(rt.Value) rt.Value { return rt.Array([]rt.Value{rt.RecordDict1("count", rt.Int(42))}) })
 if value := rt.Apply(rt.Apply(_Gopurs_Fixture_StructArrayCallback, array), rt.Value{}); value.IntVal != 42 { t.Fatal("array of struct callback", value) }
 dynamicArray := rt.Func(func(rt.Value) rt.Value { return rt.Array([]rt.Value{rt.Str("native")}) })
 if !rt.Unbox[bool](rt.Apply(rt.Apply(_Gopurs_Fixture_ForeignArrayCallback, dynamicArray), rt.Value{})) { t.Fatal("Foreign array callback leaked Value") }
 legacy := rt.Apply(rt.Apply(_Gopurs_Fixture_LegacyArrayCallback, dynamicArray), rt.Value{})
 if rt.ArrayAccess(legacy, 0).StrVal() != "native" { t.Fatal("legacy array callback lost boxed values") }
}
func TestNestedNativeArrays(t *testing.T) {
 input := rt.Array([]rt.Value{rt.Array([]rt.Value{rt.Int(42), rt.Int(-7)}), rt.Array([]rt.Value{})})
 output := rt.Apply(_Gopurs_Fixture_Matrix, input)
 if !reflect.DeepEqual(rt.ForeignValue(input), rt.ForeignValue(output)) { t.Fatal("nested array reinterpreted Value memory", rt.ForeignValue(output)) }
}
func TestUndefinedRoundTrip(t *testing.T) {
 if got := rt.Apply(_Gopurs_Fixture_Dynamic, rt.Value{}); got.Type != 0 { t.Fatal("undefined changed", got.Type) }
 if got := rt.Apply(_Gopurs_Fixture_Dynamic, rt.Any(nil)); got.Type == 0 { t.Fatal("null became undefined") }
 input := rt.RecordDict2("undefined", "null", rt.Value{}, rt.Any(nil))
 result := rt.ForeignValue(rt.Apply(_Gopurs_Fixture_Dynamic, input)).(map[string]any)
 if value, present := result["undefined"]; !present || !rt.IsForeignUndefined(value) { t.Fatal("undefined property lost") }
 if value, present := result["null"]; !present || value != nil { t.Fatal("null property changed") }
 if _, present := result["absent"]; present { t.Fatal("absent property created") }
 callback := rt.Func(func(value rt.Value) rt.Value { if value.Type != 0 { t.Fatal("callback did not receive undefined") }; return value })
 if got := rt.Apply2(_Gopurs_Fixture_EchoForeignCallback, callback, rt.Value{}); got.Type != 0 { t.Fatal("callback result changed") }
}
func TestHandles(t *testing.T) {
 for _, value := range []*handle{nil, {}} { if got := rt.Unbox[*handle](rt.Apply(_Gopurs_Fixture_Handle, rt.Any(value))); got != value { t.Fatal("handle identity") } }
 value := &handle{closed:true}
 if got := rt.Unbox[*handle](rt.Apply(_Gopurs_Fixture_Handle, rt.Any(value))); !got.closed { t.Fatal("handle state") }
}
func TestCallbacks(t *testing.T) {
 calls := 0
 cb := rt.Func(func(value rt.Value) rt.Value { if rt.Unbox[string](value)!="typed" { t.Fatal(value) }; return rt.Func(func(_ rt.Value) rt.Value { calls++; return rt.Value{} }) })
 action := rt.Apply(_Gopurs_Fixture_Callback, cb)
 if calls != 0 { t.Fatal("callback eager") }; rt.Apply(action,rt.Value{}); if calls!=1 { t.Fatal("callback effect not run") }
 uncurried := rt.Func(func(value rt.Value) rt.Value { if rt.Unbox[string](value)!="uncurried" { t.Fatal(value) }; calls++; return rt.Value{} })
 rt.Apply(rt.Apply(_Gopurs_Fixture_Uncurried, uncurried), rt.Value{})
 rt.Apply(rt.Apply(_Gopurs_Fixture_Tick, rt.Func(func(_ rt.Value) rt.Value { calls++; return rt.Value{} })), rt.Value{})
 if calls!=3 { t.Fatal("callback count",calls) }
}
func TestLegacyBindPreservesExplicitEffectThunk(t *testing.T) {
 calls := 0
 first := rt.Func(func(rt.Value) rt.Value { calls++; return rt.Value{} })
 continuation := rt.Func(func(rt.Value) rt.Value { return rt.Func(func(rt.Value) rt.Value { calls++; return rt.Value{} }) })
 bound := rt.Apply2(_Gopurs_Fixture_LegacyBind, first, continuation)
 if calls != 0 { t.Fatal("legacy bind evaluated before the effect ran") }
 rt.Apply(bound, rt.Value{})
 if calls != 2 { t.Fatal("legacy callback effect thunk was forced twice", calls) }
}
func TestLegacyEffectKeepsExplicitUnitParameter(t *testing.T) {
 if result := rt.Apply(_Gopurs_Fixture_LegacyEffect, rt.Value{}); rt.Unbox[int64](result) != 42 { t.Fatal("legacy Effect returned an extra thunk", result) }
}
func TestGenericUncurriedCallbackIsNotAnOpaqueHandle(t *testing.T) {
 callback := rt.Func2(func(left, right rt.Value) rt.Value { return rt.Str(rt.Unbox[string](left) + rt.Unbox[string](right)) })
 if result := rt.Apply(_Gopurs_Fixture_GenericFn, callback); rt.Unbox[string](result) != "leftright" { t.Fatal("generic Fn2 callback", result) }
}
`);
  const result = spawnSync('go', ['test', '-race', './...'], { cwd: directory, encoding: 'utf8', timeout: 60000, env: { ...process.env, GOWORK: 'off' } });
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stdout + result.stderr);
});
