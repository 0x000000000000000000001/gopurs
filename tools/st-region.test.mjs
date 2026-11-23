import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import test from "node:test";
import * as Map from "../output/Data.Map/index.js";
import { Just, Nothing } from "../output/Data.Maybe/index.js";
import { ordString } from "../output/Data.Ord/index.js";
import { empty as emptySet } from "../output/Data.Set/index.js";
import { Tuple } from "../output/Data.Tuple/index.js";
import * as Ref from "../output/Effect.Ref/index.js";
import { StmtEmpty, flattenStmts } from "../output/Gopurs.ExprContext/index.js";
import * as Go from "../output/Gopurs.GoAst/index.js";
import { admit, emit } from "../output/Gopurs.PrivateSTRegions/index.js";
import { printGoExpr } from "../output/Gopurs.Printer/index.js";
import { runtimeGoCode } from "../output/Gopurs.Runtime.Source/index.js";
import { TcoExpr } from "../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js";
import * as C from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";
import { localId } from "../output/PureScript.Backend.Optimizer.FreeVars/index.js";
import * as S from "../output/PureScript.Backend.Optimizer.Syntax/index.js";
import { withReboxFields } from "./codegen-metadata.mjs";

const expr = (syntax) => new TcoExpr(null, syntax);
const global = (owner, name) => expr(new S.Var(new C.Qualified(owner === null ? Nothing.value : new Just(owner), name)));
const local = (level, name = "local") => expr(new S.Local(new Just(name), level));
const app = (fn, ...args) => expr(new S.App(fn, args));
const effectApp = (owner, name, ...args) => expr(new S.UncurriedEffectApp(global(owner, name), args));
const bind = (level, value, body, name = "local") => expr(new S.EffectBind(new Just(name), level, value, body));
const letIn = (level, value, body) => expr(new S.Let(new Just("alias"), level, value, body));
const pure = (value) => expr(new S.EffectPure(value));
const refNew = (value) => expr(new S.PrimEffect(new S.EffectRefNew(value)));
const refRead = (value) => expr(new S.PrimEffect(new S.EffectRefRead(value)));
const refWrite = (ref, value) => expr(new S.PrimEffect(new S.EffectRefWrite(ref, value)));
const int = (value) => expr(new S.Lit(new C.LitInt(value)));
const bool = (value) => expr(new S.Lit(new C.LitBoolean(value)));
const array = (...values) => expr(new S.Lit(new C.LitArray(values)));
const lambda = (body) => expr(new S.Abs([new Tuple(new Just("argument"), 30)], body));
const run = (body, owner = "Control.Monad.ST.Internal") => app(global(owner, "run"), body);
const while_ = (condition, body) => app(global("Control.Monad.ST.Internal", "while"), condition, body);
const push = (value, target = local(0)) => effectApp("Data.Array.ST", "pushImpl", value, target);
const freeze = (target = local(0)) => effectApp("Data.Array.ST", "unsafeFreezeImpl", target);
const scoped = (body, initializer = int(7), newCell = refNew) => run(bind(0, global("Data.Array.ST", "new"),
  bind(1, newCell(initializer), bind(2, newCell(bool(true)), body))));

function canonical(newCell = refNew) {
  const stop = bind(5, freeze(), bind(6, refWrite(local(3), local(5)), refWrite(local(2), bool(false))));
  const step = bind(7, push(local(4)), refWrite(local(1), app(global("Probe", "next"), local(4))));
  const body = bind(4, refRead(local(1)), expr(new S.Branch([
    new S.Pair(app(global("Probe", "finished"), local(4)), stop),
  ], step)));
  return scoped(bind(3, newCell(array()), bind(8, while_(refRead(local(2)), body), refRead(local(3)))), int(7), newCell);
}

test("private ST loops admit synchronous branches and callbacks receiving only values", () => {
  assert.equal(admit(canonical()), true);
  assert.equal(admit(canonical((value) => app(global("Control.Monad.ST.Internal", "newImpl"), value))), true,
    "foreign newImpl calls in the observed lexer IR");
  assert.equal(admit(scoped(bind(3, refRead(local(1)), refWrite(local(1), local(3))))), true);
  assert.equal(admit(scoped(bind(3, push(int(9)), freeze()))), true);
  assert.equal(admit(scoped(bind(3, while_(pure(bool(false)), push(int(9))), freeze()))), true);
  assert.equal(admit(run(bind(1, app(global("Control.Monad.ST.Internal", "newImpl"), int(1)),
    refRead(local(1))))), true);
  // Copying the frozen slice also makes a later push safe for the published value.
  assert.equal(admit(scoped(bind(3, freeze(), bind(4, push(int(9)), pure(local(3)))))), true);
});

test("the proof follows lexical levels and preserves annotation wrappers", () => {
  assert.equal(admit(expr(new S.Typed(C.Any.value, canonical()))), true);
  assert.equal(admit(expr(new S.TypeApp(canonical(), C.Any.value))), true);
  // All binders intentionally share a spelling; their lexical levels differ.
  assert.equal(admit(scoped(bind(3, refRead(local(1)), pure(local(3))))), true);
  assert.equal(admit(scoped(pure(local(99)))), true, "unrelated local with the same spelling");
});

test("while never moves strict argument preparation into the repeated computation", () => {
  const deferred = (body) => expr(new S.EffectDefer(body));
  const prepare = (body) => letIn(10, app(global("Probe", "strictPrepare"), int(1)), body);
  const preparedBody = while_(pure(bool(false)), prepare(deferred(push(local(10)))));
  const preparedCondition = while_(prepare(deferred(pure(bool(false)))), deferred(push(int(1))));
  const repeatedPreparation = while_(pure(bool(false)), deferred(prepare(push(local(10)))));
  assert.equal(admit(scoped(preparedBody)), false,
    "body preparation executes once even when the condition starts false");
  assert.equal(admit(scoped(preparedCondition)), false,
    "condition preparation executes once, outside the repeated condition");
  assert.equal(admit(scoped(repeatedPreparation)), true,
    "preparation inside EffectDefer belongs to each executed iteration");
  const opaqueAction = app(global("Probe", "makeAction"), int(1));
  assert.equal(admit(scoped(while_(pure(bool(false)), opaqueAction))), false,
    "an opaque body action is constructed once, before the loop");
  assert.equal(admit(scoped(while_(opaqueAction, deferred(push(int(1)))))), false,
    "an opaque condition action is constructed once, before the loop");
  assert.equal(admit(scoped(deferred(opaqueAction))), false,
    "EffectDefer must not turn a returned opaque action into an executed action");
  assert.equal(admit(scoped(deferred(local(99)))), false,
    "EffectDefer must not execute a returned independent local action");
});

test("references and arrays cannot escape through values, aliases, callbacks or containers", () => {
  const uses = [
    ["returned reference", pure(local(1))],
    ["returned mutable array", pure(local(0))],
    ["reference alias", letIn(10, local(1), refRead(local(10)))],
    ["array alias", letIn(10, local(0), freeze(local(10)))],
    ["returned closure", pure(lambda(refRead(local(1))))],
    ["returned uncurried effect closure", pure(expr(new S.UncurriedEffectAbs([new Tuple(new Just("argument"), 30)], refRead(local(1)))))],
    ["callback capture", pure(app(global("Probe", "callback"), lambda(refRead(local(1)))))],
    ["unknown reference consumer", app(global("Probe", "read"), local(1))],
    ["unknown array consumer", effectApp("Probe", "consume", local(0))],
    ["array container", pure(array(local(1)))],
    ["reference stored in another cell", refWrite(local(2), local(1))],
    ["array stored in a cell", refWrite(local(1), local(0))],
    ["new cell containing old reference", bind(10, refNew(local(1)), refRead(local(10)))],
    ["reference stored in mutable array", push(local(1))],
    ["self-containing mutable array", push(local(0))],
    ["effect closure", pure(expr(new S.EffectDefer(refRead(local(1)))))],
    ["recursive capture", expr(new S.LetRec(20, [new Tuple("loop", lambda(refRead(local(1))))], pure(int(0))))],
    ["rebound private level", bind(1, refNew(int(9)), refRead(local(1)))],
  ];
  for (const [name, body] of uses) assert.equal(admit(scoped(body)), false, name);
});

test("only exact primitive owners, arities and local operand kinds are admitted", () => {
  for (const [name, body] of [
    ["foreign read", app(global("Other", "read"), local(1))],
    ["unqualified read", app(global(null, "read"), local(1))],
    ["foreign push", effectApp("Other", "pushImpl", int(1), local(0))],
    ["unqualified freeze", effectApp(null, "unsafeFreezeImpl", local(0))],
    ["partial push", effectApp("Data.Array.ST", "pushImpl", local(0))],
    ["extra push argument", effectApp("Data.Array.ST", "pushImpl", int(1), local(0), int(2))],
    ["partial freeze", effectApp("Data.Array.ST", "unsafeFreezeImpl")],
    ["extra freeze argument", effectApp("Data.Array.ST", "unsafeFreezeImpl", local(0), int(2))],
    ["read array as a reference", refRead(local(0))],
    ["write array as a reference", refWrite(local(0), int(1))],
    ["push into a reference", push(int(1), local(1))],
    ["freeze a reference", freeze(local(1))],
    ["unknown while implementation", app(global("Other", "while"), refRead(local(2)), push(int(1)))],
  ]) assert.equal(admit(scoped(body)), false, name);
  for (const owner of ["Other", null]) {
    assert.equal(admit(run(bind(1, refNew(int(0)), refRead(local(1))), owner)), false);
  }
  assert.equal(admit(app(global("Control.Monad.ST.Internal", "run"))), false);
  assert.equal(admit(app(global("Control.Monad.ST.Internal", "run"), canonical(), int(1))), false);
});

test("emitted regions execute native loops and preserve operation results and published snapshots", (t) => {
  const metadata = withReboxFields({
    pointerAdtPaths: Map.empty, pointerAdtNodes: emptySet, pointerAdtLeaves: Map.empty,
    enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet, ctorTypes: Map.empty,
    classDeclsFields: Map.empty, globalTypes: Map.empty, globalFunctions: Map.empty,
  });
  const context = () => ({
    metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 0, modNameStr: "Probe", recVars: [], moduleFunctions: Map.empty, bound: Map.empty,
    tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
  });
  // Deliberately supports only ordinary scalar values and bound snapshots.
  // Every ST operation must therefore be translated by the native region emitter.
  const translate = (ctx) => (nextId) => (value) => {
    const node = value.value1;
    let expression, type;
    if (node instanceof S.Lit && node.value0 instanceof C.LitInt) {
      expression = Go.rawGo(String(node.value0.value0));
      type = Go.TypeInt64.value;
    } else if (node instanceof S.Lit && node.value0 instanceof C.LitBoolean) {
      expression = Go.rawGo(String(node.value0.value0));
      type = Go.TypeBool.value;
    } else if (node instanceof S.Local) {
      const binding = Map.lookup(ordString)(localId(node.value0)(node.value1))(ctx.bound);
      assert.ok(binding instanceof Just, "missing local binding");
      expression = new Go.GoVar(binding.value0.name);
      type = binding.value0.goType;
    } else if (node instanceof S.Var && node.value0.value0 instanceof Just
      && node.value0.value0.value0 === "Probe" && ["condition", "body"].includes(node.value0.value1)) {
      expression = Go.rawGo(node.value0.value1 === "condition" ? "prepareCondition()" : "prepareBody()");
      type = Go.TypeValue.value;
    } else {
      assert.fail(`unexpected fallback translation: ${JSON.stringify(node)}`);
    }
    return { stmts: StmtEmpty.value, expr: expression, exprType: type, nextId };
  };
  const withArray = (body) => bind(0, global("Data.Array.ST", "new"), body);
  const programs = [
    ["results", withArray(bind(1, refNew(int(7)),
      bind(2, refWrite(local(1), int(12)), bind(3, push(local(2)),
        bind(4, push(local(3)), freeze())))))],
    ["snapshot", withArray(bind(1, push(int(9)), bind(2, freeze(),
      bind(3, push(int(17)), pure(local(2))))))],
    ["readSnapshot", withArray(bind(1, refNew(int(7)), bind(2, refRead(local(1)),
      bind(3, refWrite(local(1), int(99)), bind(4, push(local(2)), freeze())))))],
    ["hygiene", withArray(bind(1, refNew(local(99)), bind(2, refRead(local(1)),
      bind(3, push(local(2)), freeze())))), true],
    ["loop", withArray(bind(1, refNew(bool(true)),
      bind(3, while_(refRead(local(1)), bind(2, push(int(21)), refWrite(local(1), bool(false)))), freeze())))],
    ["skipped", withArray(bind(1, while_(pure(bool(false)), push(int(99))), freeze()))],
    ["externalActions", withArray(bind(1, while_(global("Probe", "condition"), global("Probe", "body")), freeze()))],
  ];
  const declarations = programs.map(([name, body, hygiene]) => {
    const ctx = context();
    if (hygiene) ctx.bound = Map.insert(ordString)(localId(new Just("local"))(99))({
      name: "__st_ref_1", goType: Go.TypeInt64.value,
    })(ctx.bound);
    const result = emit(translate)(ctx)(0)(global("Control.Monad.ST.Internal", "run"))([body]);
    assert.ok(result instanceof Just, name);
    const generated = [...flattenStmts(result.value0.stmts), new Go.GoReturn(result.value0.expr)]
      .map(printGoExpr).join("\n");
    assert.doesNotMatch(generated, /NewSTRef|ReadSTRef|WriteSTRef|\.Func\(/);
    if (name !== "externalActions") assert.doesNotMatch(generated, /\.Apply\(/);
    if (name === "loop" || name === "skipped") assert.match(generated, /for\s*\{/);
    return `func ${name}(${hygiene ? "__st_ref_1 int64" : ""}) []gopurs_runtime.Value {\n${generated}\n}`;
  });
  const directory = mkdtempSync(join(tmpdir(), "gopurs-private-st-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, "gopurs_runtime"));
  writeFileSync(join(directory, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  writeFileSync(join(directory, "gopurs_runtime/runtime.go"), runtimeGoCode);
  writeFileSync(join(directory, "main.go"), `package main
import ("fmt"; "runtime"; "strings"; "gopurs/output/gopurs_runtime")
${declarations.join("\n")}
var trace []string
var limit, totalTests int
func prepareCondition() gopurs_runtime.Value {
  trace = append(trace, "initCondition")
  remaining := limit
  return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
    trace = append(trace, "runCondition")
    totalTests++
    if totalTests > 20 { panic("condition action was reconstructed") }
    result := remaining > 0
    remaining--
    return gopurs_runtime.Bool(result)
  })
}
func prepareBody() gopurs_runtime.Value {
  trace = append(trace, "initBody")
  return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
    trace = append(trace, "runBody")
    return gopurs_runtime.Value{}
  })
}
func check(actual []gopurs_runtime.Value, expected ...int64) {
  if len(actual) != len(expected) { panic("length") }
  for i, value := range actual { if value.IntVal != expected[i] { panic("value/order") } }
}
func main() {
  check(results(), 12, 1)
  check(snapshot(), 9)
  check(readSnapshot(), 7)
  check(hygiene(19), 19)
  check(loop(), 21)
  check(skipped())
  check(externalActions())
  if strings.Join(trace, ",") != "initCondition,initBody,runCondition" { panic(fmt.Sprint(trace)) }
  trace, limit = nil, 2
  check(externalActions())
  if strings.Join(trace, ",") != "initCondition,initBody,runCondition,runBody,runCondition,runBody,runCondition" { panic(fmt.Sprint(trace)) }
  kept := snapshot()
  for i := 0; i < 100; i++ { check(results(), 12, 1); check(snapshot(), 9); check(loop(), 21) }
  runtime.GC()
  check(kept, 9)
  fmt.Println("private ST semantics passed")
}
`);
  const result = spawnSync("go", ["run", "."], {
    cwd: directory, encoding: "utf8", timeout: 60_000,
    env: { ...process.env, GOWORK: "off", GOTOOLCHAIN: "local" },
  });
  assert.ifError(result.error);
  assert.equal(result.status, 0, result.stdout + result.stderr);
  assert.equal(result.stdout.trim(), "private ST semantics passed");
});
