import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import test from "node:test";
import { Just } from "../output/Data.Maybe/index.js";
import { Tuple } from "../output/Data.Tuple/index.js";
import { generateFfiBridge, ffiProblems, ffiValueWorkers } from "../output/Gopurs.FfiBridge/index.js";
import { prepareFfi } from "../output/Gopurs.FfiSupport/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";

test("FFI adapts arguments and residual results at the actual native Go arity", (t) => {
  const int = C.Int.value;
  const fn = (args, result) => new C.Func(args, result);
  const uncurried = (args, result) =>
    new C.ADT(
      `Data.Function.Uncurried.Fn${args.length}`,
      ["Data", "Function", "Uncurried", `Fn${args.length}`],
      [...args, result],
    );
  const fourth = fn([int], int);
  const fn3ReturningFn = uncurried([int, int, int], fourth);
  const foreign = new C.ADT("Foreign.Foreign", ["Foreign", "Foreign"], []);
  const a = new C.TypeVar("a");
  const unitToInt = fn([C.Unit.value], int);
  const zeroForeigns = [
    ["zeroAny", unitToInt],
    ["zeroInt", unitToInt],
    ["oneIgnore", unitToInt],
    ["fnZero", uncurried([], int)],
    ["fnZeroReturning", uncurried([], fourth)],
  ].map(([name, type]) => new Tuple(name, new Just(type)));
  const foreigns = [
    ["flat", fn3ReturningFn],
    ["curried", fn3ReturningFn],
    ["dynamicCurried", fn3ReturningFn],
    ["quantified", new C.ForAll(["a"], uncurried([int, int, int], fn([a], a)))],
    ["nested", fn([int], fn([int], fn([int, int], int)))],
    ["nestedFn", uncurried([int, int], uncurried([int, int], int))],
    ["partial", fn([int, int, int, int], int)],
    ["foreignTail", uncurried([int, int, int], fn([foreign], foreign))],
    ["legacy", fn3ReturningFn],
    ["worker", fn([fn([int], int), int], fn([int], fourth))],
  ]
    .map(([name, type]) => new Tuple(name, new Just(type)))
    .concat(zeroForeigns);
  const parsed = prepareFfi({ moduleName: "ReturnArity", path: "ReturnArity.go" })("ReturnArity_")(`
func Flat(a, b, c, d int64) any { return a + b + c + d }
func Curried(a, b, c int64) func(int64) int64 {
    return func(d int64) int64 { return a + b + c + d }
}
func DynamicCurried(a, b, c int64) any {
    return func(d any) any { return a + b + c + gopurs_runtime.Unbox[int64](d) }
}
func Quantified(a, b, c int64, value any) any { return value }
func Nested(a, b, c, d int64) any { return a + b + c + d }
func NestedFn(a, b, c, d int64) any { return a + b + c + d }
func Partial(a, b int64) func(int64) func(int64) int64 {
    return func(c int64) func(int64) int64 {
        return func(d int64) int64 { return a + b + c + d }
    }
}
func ForeignTail(a, b, c int64, value any) any {
    return map[string]any{"answer": value.(map[string]any)["answer"]}
}
func Legacy(a, b, c, d gopurs_runtime.Value) gopurs_runtime.Value {
    return gopurs_runtime.Int(a.IntVal + b.IntVal + c.IntVal + d.IntVal)
}
func Worker(callback func(int64) int64, a, b int64) func(int64) any {
    return func(c int64) any { return callback(a + b + c) }
}
func ZeroAny() any { return int64(42) }
func ZeroInt() int64 { return 42 }
func OneIgnore(_ any) any { return int64(42) }
func FnZero() any { return int64(42) }
func FnZeroReturning() func(int64) int64 {
    return func(value int64) int64 { return value + 1 }
}
`)();
  assert.deepEqual(ffiProblems("ReturnArity")(parsed.decls)(zeroForeigns), []);
  const directory = mkdtempSync(join(tmpdir(), "gopurs-ffi-return-arity-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, "gopurs_runtime"));
  copyFileSync(new URL("../runtime/runtime.go", import.meta.url), join(directory, "gopurs_runtime/runtime.go"));
  writeFileSync(join(directory, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  writeFileSync(
    join(directory, "bridge.go"),
    'package main\nimport "gopurs/output/gopurs_runtime"\n' +
      parsed.content +
      "\n" +
      generateFfiBridge("ReturnArity")([])(parsed.decls)(foreigns) +
      "\n" +
      ffiValueWorkers("ReturnArity")([])(foreigns)(parsed.decls),
  );
  writeFileSync(
    join(directory, "bridge_test.go"),
    `package main
import (
    "testing"
    rt "gopurs/output/gopurs_runtime"
)
func TestNativeArity(t *testing.T) {
    for name, bridge := range map[string]rt.Value{
        "flattened nominal Fn3": _Gopurs_ReturnArity_Flat,
        "native returned function": _Gopurs_ReturnArity_Curried,
        "dynamic returned function": _Gopurs_ReturnArity_DynamicCurried,
        "nested functions": _Gopurs_ReturnArity_Nested,
        "nested nominal functions": _Gopurs_ReturnArity_NestedFn,
        "partial first argument group": _Gopurs_ReturnArity_Partial,
        "legacy runtime values": _Gopurs_ReturnArity_Legacy,
    } {
        t.Run(name, func(t *testing.T) {
            got := rt.Apply4(bridge, rt.Int(10), rt.Int(10), rt.Int(10), rt.Int(12))
            if got.Type != rt.TypeInt || got.IntVal != 42 { t.Fatalf("got %+v; expected Int(42)", got) }
        })
    }
    quantified := rt.Apply4(_Gopurs_ReturnArity_Quantified, rt.Int(0), rt.Int(0), rt.Int(0), rt.Str("preserved"))
    if got := rt.Unbox[string](quantified); got != "preserved" { t.Fatalf("polymorphic result: %q", got) }
    native := rt.Any(map[string]any{"answer": int64(42)})
    returned := rt.Apply4(_Gopurs_ReturnArity_ForeignTail, rt.Int(0), rt.Int(0), rt.Int(0), native)
    if got := rt.ForeignValue(returned).(map[string]any)["answer"]; got != int64(42) {
        t.Fatalf("Foreign argument/result after nominal Fn3: %v", got)
    }
    callback := rt.Func(func(value rt.Value) rt.Value { return rt.Int(value.IntVal + 10) })
    for name, result := range map[string]rt.Value{
        "wrapper with a returned function": rt.Apply4(_Gopurs_ReturnArity_Worker, callback, rt.Int(10), rt.Int(10), rt.Int(12)),
        "direct native worker with a returned function": rt.Apply(ReturnArity_Worker_nativeWorker(callback, 10, 10), rt.Int(12)),
    } {
        if result.Type != rt.TypeInt || result.IntVal != 42 { t.Fatalf("%s: got %+v; expected Int(42)", name, result) }
    }
}
func TestZeroArity(t *testing.T) {
    for _, item := range []struct { name string; function rt.Value }{
        {"UnitToInt_Go0Any", _Gopurs_ReturnArity_ZeroAny},
        {"UnitToInt_Go0Int64", _Gopurs_ReturnArity_ZeroInt},
        {"UnitToInt_Go1Ignore", _Gopurs_ReturnArity_OneIgnore},
        {"Fn0Int_Go0Any", _Gopurs_ReturnArity_FnZero},
    } {
        t.Run(item.name, func(t *testing.T) {
            got := rt.Apply(item.function, rt.Value{})
            if got.Type != rt.TypeInt || got.IntVal != 42 { t.Fatalf("got %+v; expected Int(42)", got) }
        })
    }
    t.Run("Fn0ReturningFunction", func(t *testing.T) {
        returned := rt.Apply(_Gopurs_ReturnArity_FnZeroReturning, rt.Value{})
        got := rt.Apply(returned, rt.Int(41))
        if got.Type != rt.TypeInt || got.IntVal != 42 { t.Fatalf("got %+v; expected Int(42)", got) }
    })
}
`,
  );
  const result = spawnSync("go", ["test", "-count=1", "-v", "./..."], {
    cwd: directory,
    encoding: "utf8",
    timeout: 120_000,
    env: { ...process.env, GOMAXPROCS: "2", GOFLAGS: "-p=2" },
  });
  assert.equal(result.status, 0, result.stdout + result.stderr + (result.error ?? ""));
});
