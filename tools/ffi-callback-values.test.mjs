import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import test from "node:test";
import { Just } from "../output/Data.Maybe/index.js";
import { toValue } from "../output/Gopurs.FfiBridge.Values/index.js";
import { TNamed } from "../output/Gopurs.FfiTypes/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import { goModule, runGo, writeGoFile } from "./lib/go.mjs";

test("polymorphic callbacks preserve runtime values and native curried functions", (t) => {
  // A callback's polymorphic result may itself be a function, as in
  // NonEmptyArray.traverse1 passing consList through the supplied map.
  const callback = new C.Func([new C.TypeVar("a")], new C.TypeVar("b"));
  const bridges = ["any", "interface{}"]
    .map(
      (goType, index) =>
        `func Bridge${index}(callback ${goType}) gopurs_runtime.Value { return ${toValue([])(new TNamed(goType))(
          new Just(callback),
        )("callback")} }`,
    )
    .join("\n");
  const directory = goModule(t, "gopurs-ffi-callback-values-", {
    runtime: readFileSync(new URL("../runtime/runtime.go", import.meta.url), "utf8"),
  });
  writeGoFile(directory, "bridge.go", 'package fixture\nimport "gopurs/output/gopurs_runtime"\n' + bridges);
  writeGoFile(
    directory,
    "bridge_test.go",
    `package fixture
import (
    "testing"
    rt "gopurs/output/gopurs_runtime"
)
func TestPolymorphicCallbacks(t *testing.T) {
    cases := []struct {
        name string
        arity int
        make func(*int) any
    }{
        {"runtime-value", 1, func(calls *int) any {
            return rt.Func(func(value rt.Value) rt.Value { *calls++; return value })
        }},
        {"native-unary", 1, func(calls *int) any {
            return func(value any) any { *calls++; return value }
        }},
        {"native-curried-two", 2, func(calls *int) any {
            return func(value any) func(any) any {
                *calls++
                return func(_ any) any { *calls++; return value }
            }
        }},
        {"native-curried-three", 3, func(calls *int) any {
            return func(value any) func(any) func(any) any {
                *calls++
                return func(_ any) func(any) any {
                    *calls++
                    return func(_ any) any { *calls++; return value }
                }
            }
        }},
    }
    for name, bridge := range map[string]func(any) rt.Value{"any": Bridge0, "interface": Bridge1} {
        for _, row := range cases {
            t.Run(name+"/"+row.name, func(t *testing.T) {
                calls := 0
                wrapped := bridge(row.make(&calls))
                if calls != 0 { t.Fatal("wrapping invoked the callback") }
                result := rt.Apply(wrapped, rt.Int(42))
                for depth := 1; depth < row.arity; depth++ { result = rt.Apply(result, rt.Int(7)) }
                if got := rt.Unbox[int32](result); got != 42 { t.Fatalf("result: %d", got) }
                if calls != row.arity { t.Fatalf("callback calls: %d, want %d", calls, row.arity) }
            })
        }
    }
}
`,
  );
  const result = runGo(directory, ["test", "-count=1", "-v", "."]);
  assert.match(result.stdout, /--- PASS: TestPolymorphicCallbacks/);
  for (const form of ["runtime-value", "native-unary", "native-curried-two", "native-curried-three"]) {
    assert.match(result.stdout, new RegExp(`--- PASS: TestPolymorphicCallbacks/any/${form}`));
    assert.match(result.stdout, new RegExp(`--- PASS: TestPolymorphicCallbacks/interface/${form}`));
  }
});
