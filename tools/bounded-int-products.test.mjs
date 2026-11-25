import assert from 'node:assert/strict';
import test from 'node:test';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { StmtEmpty, StmtLeaf } from '../output/Gopurs.ExprContext/index.js';
import { boundedIntProduct, binary } from '../output/Gopurs.PrimitiveExprs/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

const expr = syntax => new TcoExpr(null, syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const integer = n => typed(C.Int.value, expr(new S.Lit(new C.LitInt(n))));
const local = level => typed(C.Int.value, expr(new S.Local(Nothing.value, level)));
const op = (operator, a, b) => typed(C.Int.value, expr(new S.PrimOp(new S.Op2(operator, a, b))));
const normalized = a => op(S.OpIntBitOr.value, a, integer(0));
const mask = (a, n) => op(S.OpIntBitAnd.value, a, integer(n));
const shift = n => op(S.OpIntBitShiftRight.value, local(0), integer(n));
const multiply = new S.OpIntNum(S.OpMultiply.value);
const context = { codegenStateRef: Ref.new({ reboxPairs: emptySet })(), modNameStr: 'RenamedBounds', recVars: [], depth: 0 };
const dummy = () => nextId => () => ({ stmts: StmtEmpty.value, expr: new Go.GoVar('value'), exprType: Go.TypeInt64.value, nextId });
const emit = (left, right, translate = dummy) => boundedIntProduct(translate)(context)(0)(multiply)(left)(right);

for (const [label, left, right] of [
    ['low 16 bits', mask(local(0), 65535), normalized(local(1))],
    ['zero mask', mask(local(0), 0), normalized(local(1))],
    ['high signed 16 bits', shift(16), normalized(local(1))],
    ['sign bit shift', shift(31), normalized(local(1))],
    ['swapped mask operands', op(S.OpIntBitAnd.value, integer(65535), local(0)), normalized(local(1))],
    ['exact positive 2^53', integer(-2147483648), integer(-4194304)],
    ['exact negative 2^53', integer(-2147483648), integer(4194304)],
    ['negative literals', integer(-65535), normalized(local(1))],
    ['normalizing addition', mask(local(0), 65535), op(new S.OpIntNum(S.OpAdd.value), local(1), integer(1))],
]) test(`bounded product admits ${label}`, () => assert.ok(emit(left, right) instanceof Just));

for (const [label, left, right] of [
    ['unproven typed locals', local(0), local(1)],
    ['two full Int32 values', normalized(local(0)), normalized(local(1))],
    ['magnitude beyond 2^53', integer(2147483647), integer(4194305)],
    ['wide literal', integer(4294967296), integer(1)],
    ['wide division result', op(new S.OpIntNum(S.OpDivide.value), local(0), integer(-1)), integer(65535)],
    ['unsigned shift result', op(S.OpIntBitZeroFillShiftRight.value, local(0), integer(0)), integer(65535)],
    ['shift by 32', shift(32), normalized(local(1))],
    ['negative shift', shift(-16), normalized(local(1))],
    ['unknown shift count', op(S.OpIntBitShiftRight.value, local(0), local(1)), normalized(local(1))],
    ['a Number annotation', typed(C.Number.value, mask(local(0), 65535)), normalized(local(1))],
    ['an unknown call', expr(new S.App(local(0), [integer(1)])), integer(65535)],
]) test(`bounded product retains ${label}`, () => {
    const fail = () => { throw new Error('rejected proof must not translate operands'); };
    assert.ok(emit(left, right, fail) instanceof Nothing);
});

test('native products match JS rounding, constants, signed boundaries and strict operand order', t => {
    const values = [-2147483648, -2147483647, -65536, -1, 0, 1, 65535, 65536, 2147483647, 4294967296];
    const shapes = [
        [mask(local(0), 65535), normalized(local(1)), 'gopurs_runtime.IntBitAnd(left(a),65535)', 'gopurs_runtime.IntBitOr(right(b),0)', (a,b) => Math.trunc(((a & 65535) * (b | 0))) | 0],
        [shift(16), normalized(local(1)), 'gopurs_runtime.IntShr(left(a),16)', 'gopurs_runtime.IntBitOr(right(b),0)', (a,b) => ((a >> 16) * (b | 0)) | 0],
        [mask(local(0), 0), normalized(local(1)), 'gopurs_runtime.IntBitAnd(left(a),0)', 'gopurs_runtime.IntBitOr(right(b),0)', () => 0],
        [shift(31), normalized(local(1)), 'gopurs_runtime.IntShr(left(a),31)', 'gopurs_runtime.IntBitOr(right(b),0)', (a,b) => ((a >> 31) * (b | 0)) | 0],
    ];
    const declarations = [], checks = [];
    for (const [left, right, leftGo, rightGo, oracle] of shapes) {
        const translate = () => nextId => expression => ({
            stmts: new StmtLeaf(Go.rawGo(`trace += "${expression === left ? 'l' : 'r'}"`)),
            expr: Go.rawGo(expression === left ? leftGo : rightGo), exprType: Go.TypeInt64.value, nextId: nextId + 1,
        });
        const result = emit(left, right, translate).value0;
        assert.equal(result.nextId, 2);
        assert.equal(result.stmts, StmtEmpty.value);
        const name = `bounded${declarations.length}`;
        declarations.push(`func ${name}(a,b int64) int64 { return ${printGoExpr(result.expr)} }`);
        for (const a of values) for (const b of values) checks.push(`check(${name},${a},${b},${oracle(a,b)})`);
    }
    // Constant-only lowering must compile even when the exact product cannot
    // be represented by a Go int32 constant. Ordinary large products retain
    // IntMul's JS rounding rather than silently adopting modular arithmetic.
    for (const [a,b,admitted] of [[-2147483648,-4194304,true],[-2147483648,4194304,true],[2147483647,4194305,false],[2147483647,2147483647,false]]) {
        const left = integer(a), right = integer(b);
        const translate = () => nextId => expression => ({ stmts: StmtEmpty.value,
            expr: new Go.GoInt(expression === left ? a : b), exprType: Go.TypeInt64.value, nextId });
        const result = emit(left,right,translate);
        assert.equal(result instanceof Just, admitted);
        const product = admitted ? result.value0 : binary(context.codegenStateRef)('RenamedBounds')(multiply).value0(
            {expr:new Go.GoInt(a),exprType:Go.TypeInt64.value})({expr:new Go.GoInt(b),exprType:Go.TypeInt64.value});
        checks.push(`if got:= ${printGoExpr(product.expr)}; got != ${a*b|0} {panic("constant product")}`);
    }
    const directory = goModule(t,'gopurs-bounded-products-', {runtime:runtimeGoCode});
    writeGoFile(directory,'main.go',`package main
import("fmt";"gopurs/output/gopurs_runtime")
var trace string
var failLeft bool
func left(a int64)int64{trace+="L";if failLeft {panic("left failed")};return a}
func right(b int64)int64{trace+="R";return b}
${declarations.join('\n')}
func check(operation func(int64,int64)int64,a,b,want int64){trace="";got:=operation(a,b);if got!=want||trace!="lLrR"{panic(fmt.Sprintf("%d,%d: got %d want %d trace %s",a,b,got,want,trace))}}
func checkFailure(){trace="";failLeft=true;defer func(){if recover()==nil||trace!="lL"{panic("right operand evaluated after left failure")}}();bounded0(1,2)}
func main(){${checks.join('\n')};checkFailure();fmt.Print("404 products and strict failure verified")}
`);
    assert.equal(runGo(directory,['run','.']).stdout,'404 products and strict failure verified');
});
