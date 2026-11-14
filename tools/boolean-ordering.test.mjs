import assert from "node:assert/strict";
import test from "node:test";
import { Just } from "../output/Data.Maybe/index.js";
import { empty as emptySet } from "../output/Data.Set/index.js";
import * as Ref from "../output/Effect.Ref/index.js";
import * as Go from "../output/Gopurs.GoAst/index.js";
import { binary } from "../output/Gopurs.PrimitiveExprs/index.js";
import { printGoExpr } from "../output/Gopurs.Printer/index.js";
import { runtimeGoCode } from "../output/Gopurs.Runtime.Source/index.js";
import * as S from "../output/PureScript.Backend.Optimizer.Syntax/index.js";
import { goModule, runGo, writeGoFile } from "./lib/go.mjs";

test("boolean ordering handles native and boxed operands with eager single evaluation", t => {
  const state = Ref.new({ reboxPairs: emptySet })();
  const operators = [
    [S.OpEq.value, "=="], [S.OpNotEq.value, "!="],
    [S.OpLt.value, "<"], [S.OpLte.value, "<="],
    [S.OpGt.value, ">"], [S.OpGte.value, ">="],
  ];
  const operand = (name, value, boxed) => {
    const call = new Go.GoCall(new Go.GoVar(name), [new Go.GoVar(value)]);
    return boxed
      ? { expr: new Go.GoCall(new Go.GoSelector(new Go.GoVar("gopurs_runtime"), "Bool"), [call]), exprType: Go.TypeValue.value }
      : { expr: call, exprType: Go.TypeBool.value };
  };
  const declarations = [];
  const checks = [];
  for (const [operator, symbol] of operators) {
    const emitter = binary(state)("Test")(new S.OpBooleanOrd(operator));
    assert.ok(emitter instanceof Just);
    for (const leftBoxed of [false, true]) {
      for (const rightBoxed of [false, true]) {
        const name = `compare${declarations.length}`;
        const result = emitter.value0(operand("left", "a", leftBoxed))(operand("right", "b", rightBoxed));
        assert.deepEqual(result.exprType, Go.TypeBool.value);
        declarations.push(`func ${name}(a, b bool) bool { return ${printGoExpr(result.expr)} }`);
        checks.push(`trace = ""
        if got := ${name}(a, b); got != (i ${symbol} j) || trace != "LR" {
          panic(fmt.Sprintf("${name}(%t,%t): got %t, trace %q", a, b, got, trace))
        }`);
      }
    }
  }
  const directory = goModule(t, "gopurs-boolean-ordering-", { runtime: runtimeGoCode });
  writeGoFile(directory, "main.go", `package main
import (
  "fmt"
  "gopurs/output/gopurs_runtime"
)
var trace string
func left(value bool) bool { trace += "L"; return value }
func right(value bool) bool { trace += "R"; return value }
${declarations.join("\n")}
func main() {
  for i, a := range []bool{false, true} {
    for j, b := range []bool{false, true} {
      ${checks.join("\n")}
    }
  }
  fmt.Print("96 cases passed")
}
`);
  assert.equal(runGo(directory, ["run", "."]).stdout, "96 cases passed");
});
