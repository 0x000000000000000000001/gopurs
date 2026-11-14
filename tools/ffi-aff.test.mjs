import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { Just } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { ffiProblems, generateFfiBridge } from '../output/Gopurs.FfiBridge/index.js';
import { prepareFfi } from '../output/Gopurs.FfiSupport/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';

const aff = result => new C.ADT('Effect.Aff.Compat.EffectFnAff', ['Effect', 'Aff', 'Compat', 'EffectFnAff'], [result]);
const fn = (args, result) => new C.Func(args, result);
const str = C.String.value;
const unit = C.Unit.value;
const source = `package fixture
import (
  "context"
  "errors"
  "sync/atomic"
)
var Calls atomic.Int64
func Hello(s string) (string, error) { Calls.Add(1); if s == "error" { return "", errors.New("expected") }; return "hello " + s, nil }
func Notify() error { Calls.Add(1); return nil }
func Plain() string { Calls.Add(1); return "plain" }
func Void() { Calls.Add(1) }
func Context(ctx context.Context) (string, error) { if ctx == nil { panic("nil context") }; return "context", nil }
func Count() int64 { return Calls.Load() }
func Acquire() (string, func(), error) { return "resource", func() { Calls.Add(100) }, nil }
func Subscribe() (func(), error) { return func() { Calls.Add(100) }, nil }
func SubscribeFail() (func(), error) { return func() { Calls.Add(100) }, errors.New("subscribe failed") }
`;
const foreigns = [
  ['hello', fn([str], aff(str))], ['notify', aff(unit)], ['plain', aff(str)],
  ['void', aff(unit)], ['context', aff(str)], ['acquire', aff(str)], ['subscribe', aff(unit)], ['subscribeFail', aff(unit)],
  ['count', new C.ADT('Effect.Effect', ['Effect', 'Effect'], [C.Int.value])],
].map(([name, type]) => new Tuple(name, new Just(type)));

test('ordinary Go results produce lazy, repeatable EffectFnAff actions', t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-idiomatic-ffi-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\ngo 1.27.0\n');
  mkdirSync(join(directory, 'ordinary'));
  writeFileSync(join(directory, 'ordinary/fixture.go'), source);
  writeFileSync(join(directory, 'ordinary/fixture_test.go'), `package fixture
import "testing"
func TestOrdinary(t *testing.T) { if value, err := Hello("Go"); value != "hello Go" || err != nil { t.Fatal(value, err) }; if _, err := Hello("error"); err == nil { t.Fatal("missing error") } }
`);
  const parsed = prepareFfi({ moduleName: 'Fixture', path: 'fixture.go' })('Fixture_')(source)();
  assert.deepEqual(ffiProblems('Fixture')(parsed.decls)(foreigns), []);
  const bridge = generateFfiBridge('Fixture')([])(parsed.decls)(foreigns);
  assert.match(bridge, /EffectFnAff\(func/);
  mkdirSync(join(directory, 'gopurs_runtime'));
  copyFileSync(new URL('../runtime/runtime.go', import.meta.url), join(directory, 'gopurs_runtime/runtime.go'));
  writeFileSync(join(directory, 'bridge.go'), 'package fixture\nimport "gopurs/output/gopurs_runtime"\n' + parsed.content + '\n' + bridge);
  writeFileSync(join(directory, 'bridge_test.go'), `package fixture
import (
  "testing"
  "time"
  rt "gopurs/output/gopurs_runtime"
)
func execute(t *testing.T, action rt.Value) (rt.Value, error) {
  t.Helper()
  type result struct { value rt.Value; err error }
  done := make(chan result, 2)
  rt.Apply2(action, rt.Func(func(err rt.Value) rt.Value { done <- result{err: rt.Unbox[error](err)}; return rt.Value{} }), rt.Func(func(value rt.Value) rt.Value { done <- result{value: value}; return rt.Value{} }))
  select { case result := <-done: return result.value, result.err; case <-time.After(time.Second): t.Fatal("timeout"); return rt.Value{}, nil }
}
func TestActions(t *testing.T) {
  action := rt.Apply(_Gopurs_Fixture_Hello, rt.Str("Go"))
  if Fixture_Calls.Load() != 0 { t.Fatal("eager execution") }
  for i := 0; i < 2; i++ { value, err := execute(t, action); if err != nil || rt.Unbox[string](value) != "hello Go" { t.Fatal(value, err) } }
  if Fixture_Calls.Load() != 2 { t.Fatal("action memoized") }
  if _, err := execute(t, rt.Apply(_Gopurs_Fixture_Hello, rt.Str("error"))); err == nil || err.Error() != "expected" { t.Fatal(err) }
  for _, action := range []rt.Value{_Gopurs_Fixture_Notify, _Gopurs_Fixture_Plain, _Gopurs_Fixture_Void, _Gopurs_Fixture_Context} { if _, err := execute(t, action); err != nil { t.Fatal(err) } }
  if got := rt.Apply(_Gopurs_Fixture_Count, rt.Value{}); rt.Unbox[int64](got) != 6 { t.Fatal("Effect count", got) }
  if value, err := execute(t, _Gopurs_Fixture_Acquire); err != nil || rt.Unbox[string](value)!="resource" { t.Fatal(value,err) }
  if _, err := execute(t, _Gopurs_Fixture_Subscribe); err != nil { t.Fatal(err) }
  if Fixture_Calls.Load()!=6 { t.Fatal("transferred acquisition released") }
  if _, err := execute(t, _Gopurs_Fixture_SubscribeFail); err == nil || err.Error()!="subscribe failed" { t.Fatal(err) }
  if Fixture_Calls.Load()!=106 { t.Fatal("failed acquisition not released once") }
}
`);
  const result = spawnSync('go', ['test', '-race', '-count=1', './...'], { cwd: directory, encoding: 'utf8', timeout: 60000, env: { ...process.env, GOWORK: 'off' } });
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stdout + result.stderr);
});

test('multiple results cannot silently enter a pure or incompatible effect binding', () => {
  const parsed = prepareFfi({ moduleName: 'Invalid', path: '/fixture/invalid.go' })('Invalid_')('func Read() (string, error) { return "", nil }')();
  for (const type of [str, new C.ADT('Effect.Effect', ['Effect', 'Effect'], [str])]) {
    const errors = ffiProblems('Invalid')(parsed.decls)([new Tuple('read', new Just(type))]);
    assert.equal(errors.length, 1);
    assert.match(errors[0], /read at 1:6.*Go.*PureScript/);
  }
});

test('incompatible native scalar and ambiguous dynamic Aff results fail explicitly', () => {
  for (const source of ['func Read() (int64, error) { return 1, nil }', 'func Read() any { return "value" }']) {
    const parsed = prepareFfi({ moduleName: 'Invalid', path: '/fixture/invalid.go' })('Invalid_')(source)();
    const errors = ffiProblems('Invalid')(parsed.decls)([new Tuple('read', new Just(aff(str)))]);
    assert.equal(errors.length, 1);
    assert.match(errors[0], /scalar|ambiguous/);
  }
});

test('unqualified ordinary Aff and constrained imports are rejected before emission', () => {
  const parsed = prepareFfi({ moduleName: 'Invalid', path: '/fixture/invalid.go' })('Invalid_')('func Read() string { return "value" }')();
  const direct = new C.ADT('Effect.Aff.Aff', ['Effect', 'Aff', 'Aff'], [str]);
  const constrained = new C.ConstrainedType([new Tuple(['Data', 'Show', 'Show'], [str])], new C.ADT('Effect.Effect', ['Effect', 'Effect'], [str]));
  for (const [type, message] of [[direct, /use EffectFnAff/], [constrained, /type-class dictionaries/]]) {
    const errors = ffiProblems('Invalid')(parsed.decls)([new Tuple('read', new Just(type))]);
    assert.equal(errors.length, 1);
    assert.match(errors[0], message);
  }
});
