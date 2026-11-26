import assert from 'node:assert/strict';
import test from 'node:test';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { StmtEmpty, StmtLeaf } from '../output/Gopurs.ExprContext/index.js';
import { boundedIntProduct, guardedIntProduct, powerOfTwoModulo } from '../output/Gopurs.PrimitiveExprs/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

const e = syntax => new TcoExpr(null, syntax);
const typed = (type, value) => e(new S.Typed(type, value));
const int = value => typed(C.Int.value, e(new S.Lit(new C.LitInt(value))));
const local = level => typed(C.Int.value, e(new S.Local(Nothing.value, level)));
const operation = (operator, a, b) => typed(C.Int.value, e(new S.PrimOp(new S.Op2(operator, a, b))));
const num = (operator, a, b) => operation(new S.OpIntNum(operator), a, b);
const mod = (a, b) => num(S.OpMod.value, a, b);
const add = (a, b) => num(S.OpAdd.value, a, b);
const weight = add(int(1), mod(local(0), int(2)));
const multiply = new S.OpIntNum(S.OpMultiply.value), modulo = new S.OpIntNum(S.OpMod.value);
const context = { codegenStateRef: Ref.new({ reboxPairs: emptySet })(), modNameStr: 'Arithmetic', recVars: [], depth: 0 };
const dummy = () => nextId => () => ({ stmts: StmtEmpty.value, expr: new Go.GoVar('value'), exprType: Go.TypeInt64.value, nextId });
const guarded = (a, b, translate = dummy) => guardedIntProduct(translate)(context)(0)(multiply)(a)(b);
const bounded = (a, b) => boundedIntProduct(dummy)(context)(0)(multiply)(a)(b);
const masked = (a, b, translate = dummy) => powerOfTwoModulo(translate)(context)(0)(modulo)(a)(b);

test('guarded multiplication uses proved small factors without assuming opaque Int values are Int32', () => {
  for (const factor of [int(0), int(1), int(-2), int(2097152), int(-2097152), weight,
    mod(local(0), int(3)), num(S.OpSubtract.value, mod(local(0), int(8)), int(3))]) {
    for (const [left, right] of [[factor, local(1)], [local(1), factor]]) assert.ok(guarded(left, right) instanceof Just);
  }
  const expression = printGoExpr(guarded(weight, local(1)).value0.expr);
  assert.match(expression, /guardedRight/); assert.match(expression, /<< \(52\)/);
  assert.match(expression, /int64\(int32\(\(guardedLeft\) \* \(guardedRight\)\)\)/);
  assert.match(expression, /gopurs_runtime.IntMul\(guardedLeft, guardedRight\)/);
  const normalized = operation(S.OpIntBitOr.value, local(1), int(0));
  assert.ok(bounded(weight, normalized) instanceof Just, 'a fully proved product needs no dynamic guard');
  assert.equal(bounded(weight, local(1)), Nothing.value);
});

test('unknown factors, incompatible annotations, large factors and noncanonical modulus literals remain fallback', () => {
  const never = () => { throw Error('rejected arithmetic must not translate operands'); };
  for (const [a, b] of [[local(0), local(1)], [int(2097153), local(1)], [local(0), int(-2097153)],
    [typed(C.Number.value, local(0)), int(1)], [weight, typed(C.Boolean.value, local(1))]]) {
    assert.equal(guarded(a, b, never), Nothing.value);
  }
  for (const divisor of [0, -1, -2, 3, 2147483647, 4294967296, 4294967298, 9007199254740992]) {
    assert.equal(masked(local(0), int(divisor), never), Nothing.value);
  }
  for (const divisor of [4294967296, 4294967298, 9007199254740992]) {
    assert.equal(guarded(mod(local(0), int(divisor)), local(1), never), Nothing.value);
    assert.equal(bounded(mod(local(0), int(divisor)), int(1)), Nothing.value);
  }
  assert.equal(masked(local(0), local(1), never), Nothing.value);
  assert.equal(masked(typed(C.Number.value, local(0)), int(2), never), Nothing.value);
  assert.equal(masked(local(0), typed(C.Number.value, int(2)), never), Nothing.value);
  for (const divisor of [1, 2, 4, 8, 1024, 1073741824]) assert.ok(masked(local(0), int(divisor)) instanceof Just);
});

test('generated guards and masks preserve JS rounding, int64 fallback, operand order and failure', t => {
  const declarations = [], checks = [];
  const compileGuard = (left, right, leftGo, rightGo) => {
    const translate = () => nextId => expression => ({ stmts: new StmtLeaf(Go.rawGo(`trace += "${expression === left ? 'l' : 'r'}"`)),
      expr: Go.rawGo(expression === left ? leftGo : rightGo), exprType: Go.TypeInt64.value, nextId: nextId + 1 });
    const result = guarded(left, right, translate).value0;
    assert.equal(result.stmts, StmtEmpty.value); assert.equal(result.nextId, 2);
    const name = `guard${declarations.length}`;
    declarations.push(`func ${name}(a,b int64)int64{return ${printGoExpr(result.expr)}}`);
    return name;
  };
  const left = compileGuard(weight, local(1), 'gopurs_runtime.IntAdd(1,gopurs_runtime.IntMod(left(a),2))', 'right(b)');
  const right = compileGuard(local(1), weight, 'left(a)', 'gopurs_runtime.IntAdd(1,gopurs_runtime.IntMod(right(b),2))');
  const signed = compileGuard(num(S.OpSubtract.value, int(0), weight), local(1), '-gopurs_runtime.IntAdd(1,gopurs_runtime.IntMod(left(a),2))', 'right(b)');
  const boundary = compileGuard(int(2097152), local(1), 'left(2097152)', 'right(b)');
  const values = [-9007199254740992, -9007199254740991, -4503599627370497, -4503599627370496, -4503599627370495,
    -4294967296, -2147483648, -1, 0, 1, 2147483647, 2147483648, 4294967295, 4294967296,
    4503599627370495, 4503599627370496, 4503599627370497, 9007199254740991, 9007199254740992];
  const remainder = (a,b) => ((a % b) + b) % b;
  for (const index of [-9007199254740992, -3, -1, 0, 1, 2, 2147483648, 4294967295]) for (const v of values) {
    const weight = 1 + remainder(index,2);
    checks.push(`check(${left},${index},${v},${weight*v|0})`, `check(${right},${v},${index},${v*weight|0})`,
      `check(${signed},${index},${v},${-weight*v|0})`);
  }
  for (const v of [-4294967297,-4294967296,-4294967295,4294967295,4294967296,4294967297,9007199254740991]) {
    checks.push(`check(${boundary},0,${v},${2097152*v|0})`);
  }
  for (const divisor of [1, 2, 4, 8, 1024, 1073741824]) {
    const translated = masked(local(0), int(divisor), () => nextId => () => ({
      stmts: new StmtLeaf(Go.rawGo('trace += "l"')), expr: Go.rawGo('left(a)'), exprType: Go.TypeInt64.value, nextId: nextId+1 })).value0;
    const name = `mask${divisor}`;
    declarations.push(`func ${name}(a int64)int64{return ${printGoExpr(translated.expr)}}`);
    for (const v of values) checks.push(`trace="";if got:=${name}(${v});got!=${remainder(v,divisor)}||trace!="lL"{panic("modulo/trace")}`);
    checks.push(`for _,v:=range []int64{-9223372036854775808,9223372036854775807,-9007199254740993,9007199254740993}{trace="";if ${name}(v)!=gopurs_runtime.IntMod(v,${divisor}){panic("int64 mask fallback oracle")}}`);
  }
  const dir = goModule(t, 'gopurs-guarded-arithmetic-', { runtime: runtimeGoCode });
  writeGoFile(dir, 'main.go', `package main
import("fmt";"gopurs/output/gopurs_runtime")
var trace string
var failLeft,failRight bool
func left(v int64)int64{trace+="L";if failLeft{panic("left")};return v}
func right(v int64)int64{trace+="R";if failRight{panic("right")};return v}
${declarations.join('\n')}
func check(f func(int64,int64)int64,a,b,want int64){trace="";if got:=f(a,b);got!=want||trace!="lLrR"{panic(fmt.Sprintf("%d %d: %d != %d trace %s",a,b,got,want,trace))}}
func failure(rightSide bool){trace="";failLeft=!rightSide;failRight=rightSide;defer func(){message:=recover();want:="lL";if rightSide{want="lLrR"};if message==nil||trace!=want{panic("strict failure order")};failLeft=false;failRight=false}();${left}(0,7)}
func main(){${checks.join('\n')}
 for _,v:=range []int64{-9223372036854775808,9223372036854775807,-9007199254740993,9007199254740993}{for _,i:=range []int64{-3,0,1,2}{w:=gopurs_runtime.IntAdd(1,gopurs_runtime.IntMod(i,2));check(${left},i,v,gopurs_runtime.IntMul(w,v));check(${right},v,i,gopurs_runtime.IntMul(v,w));check(${signed},i,v,gopurs_runtime.IntMul(-w,v))}}
 failure(false);failure(true);fmt.Println("guarded products, modulo, wide fallback and strict order: ok")}
`);
  assert.match(runGo(dir, ['run', '.']).stdout, /wide fallback and strict order: ok/);
});
