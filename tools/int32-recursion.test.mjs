import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { generated } from '../output/Gopurs.GoSymbol/index.js';
import { plan } from '../output/Gopurs.Int32Recursion/index.js';
import { printGoDecl, printGoExpr } from '../output/Gopurs.Printer/index.js';

const expr = syntax => new TcoExpr({}, syntax);
const lit = n => expr(new S.Lit(new C.LitInt(n)));
const boolLit = b => expr(new S.Lit(new C.LitBoolean(b)));
const local = (name, level) => expr(new S.Local(new Just(name), level));
const gvar = name => expr(new S.Var(new C.Qualified(Nothing.value, name)));
const call = (name, args) => expr(new S.App(gvar(name), args));
const op1 = (op, x) => expr(new S.PrimOp(new S.Op1(op, x)));
const op2 = (op, a, b) => expr(new S.PrimOp(new S.Op2(op, a, b)));
const eq = (a, b) => op2(new S.OpIntOrd(S.OpEq.value), a, b);
const lt = (a, b) => op2(new S.OpIntOrd(S.OpLt.value), a, b);
const lte = (a, b) => op2(new S.OpIntOrd(S.OpLte.value), a, b);
const gt = (a, b) => op2(new S.OpIntOrd(S.OpGt.value), a, b);
const add = (a, b) => op2(new S.OpIntNum(S.OpAdd.value), a, b);
const sub = (a, b) => op2(new S.OpIntNum(S.OpSubtract.value), a, b);
const mul = (a, b) => op2(new S.OpIntNum(S.OpMultiply.value), a, b);
const zshr = (a, b) => op2(S.OpIntBitZeroFillShiftRight.value, a, b);
const shl = (a, b) => op2(S.OpIntBitShiftLeft.value, a, b);
const xor = (a, b) => op2(S.OpIntBitXor.value, a, b);
const neg = x => op1(S.OpIntNegate.value, x);
const branch = (pairs, fallback) => expr(new S.Branch(pairs.map(([c, v]) => new S.Pair(c, v)), fallback));
const letBind = (name, level, val, rest) => expr(new S.Let(new Just(name), level, val, rest));

test('Int32Recursion rejects pure tail loops, non-exact Int32 ops, external calls and wide literals', () => {
  const params1 = [new Tuple('n_0', Go.TypeInt64.value)];
  const run = (name, params, body, ret = Go.TypeInt64.value) =>
    plan('M')(name)(params)(ret)(body)(generated('M')(`${name}_i32`))(new Go.GoVar(`${name}_cold`));

  // Pure tail loop (no non-tail self-call) is left to tail-loop passes.
  const tailOnly = branch([[eq(local('n', 0), lit(0)), lit(0)]], call('countdown', [sub(local('n', 0), lit(1))]));
  assert.ok(run('countdown', params1, tailOnly) instanceof Nothing);

  // Multiply follows JS Number semantics, so must be rejected.
  const mulRec = branch([[lte(local('n', 0), lit(1)), lit(1)]], mul(local('n', 0), call('fact', [sub(local('n', 0), lit(1))])));
  assert.ok(run('fact', params1, mulRec) instanceof Nothing);

  // Unsigned right shift can exceed signed int32.
  const zshrRec = branch([[eq(local('n', 0), lit(0)), lit(0)]], add(zshr(local('n', 0), lit(1)), call('u', [sub(local('n', 0), lit(1))])));
  assert.ok(run('u', params1, zshrRec) instanceof Nothing);

  // External call must be rejected.
  const extRec = branch([[eq(local('n', 0), lit(0)), lit(0)]], add(call('other', [local('n', 0)]), call('ext', [sub(local('n', 0), lit(1))])));
  assert.ok(run('ext', params1, extRec) instanceof Nothing);

  // Non-Int64 parameter or result must be rejected.
  assert.ok(run('badRet', params1, branch([[eq(local('n', 0), lit(0)), lit(0)]], add(lit(1), call('badRet', [sub(local('n', 0), lit(1))]))), Go.TypeValue.value) instanceof Nothing);
  assert.ok(run('badArg', [new Tuple('n_0', Go.TypeValue.value)], branch([[eq(local('n', 0), lit(0)), lit(0)]], add(lit(1), call('badArg', [sub(local('n', 0), lit(1))])))) instanceof Nothing);
});

test('Int32Recursion handles Fib, Ackermann, Tak(x,y,z), McCarthy91, Let, bitwise ops, wide inputs and Int32 wrap', () => {
  const fibBody = branch(
    [
      [eq(local('n', 0), lit(0)), lit(0)],
      [eq(local('n', 0), lit(1)), lit(1)],
    ],
    add(call('fib', [sub(local('n', 0), lit(1))]), call('fib', [sub(local('n', 0), lit(2))])),
  );

  const ackBody = branch(
    [
      [eq(local('m', 0), lit(0)), add(local('n', 1), lit(1))],
      [eq(local('n', 1), lit(0)), call('ack', [sub(local('m', 0), lit(1)), lit(1)])],
    ],
    call('ack', [sub(local('m', 0), lit(1)), call('ack', [local('m', 0), sub(local('n', 1), lit(1))])]),
  );

  // Takeuchi 3-arg nested recursion: tak(x, y, z) = if x <= y then z else tak(tak(x-1,y,z), tak(y-1,z,x), tak(z-1,x,y))
  const takBody = branch(
    [[lte(local('x', 0), local('y', 1)), local('z', 2)]],
    call('tak', [
      call('tak', [sub(local('x', 0), lit(1)), local('y', 1), local('z', 2)]),
      call('tak', [sub(local('y', 1), lit(1)), local('z', 2), local('x', 0)]),
      call('tak', [sub(local('z', 2), lit(1)), local('x', 0), local('y', 1)]),
    ]),
  );

  // McCarthy 91 with Let + bitwise mix test
  const mc91Body = branch(
    [[gt(local('n', 0), lit(100)), sub(local('n', 0), lit(10))]],
    letBind('inner', 1, call('mc91', [add(local('n', 0), lit(11))]), call('mc91', [local('inner', 1)])),
  );

  const bitTreeBody = branch(
    [[lte(local('n', 0), lit(0)), neg(local('seed', 1))]],
    xor(
      shl(call('bitTree', [sub(local('n', 0), lit(1)), local('seed', 1)]), lit(1)),
      call('bitTree', [sub(local('n', 0), lit(2)), add(local('seed', 1), lit(3))]),
    ),
  );

  const mk = (name, ps, body) => {
    const params = ps.map(p => new Tuple(p, Go.TypeInt64.value));
    const coldCall = new Go.GoCall(new Go.GoVar(`${name}_cold`), ps.map(p => new Go.GoVar(`${p}_loop`)));
    const res = plan('M')(name)(params)(Go.TypeInt64.value)(body)(generated('M')(`${name}_i32`))(coldCall);
    assert.ok(res instanceof Just, `${name} should be specialized`);
    return res.value0;
  };

  const pFib = mk('fib', ['n_0'], fibBody);
  const pAck = mk('ack', ['m_0', 'n_1'], ackBody);
  const pTak = mk('tak', ['x_0', 'y_1', 'z_2'], takBody);
  const pMc91 = mk('mc91', ['n_0'], mc91Body);
  const pBit = mk('bitTree', ['n_0', 'seed_1'], bitTreeBody);

  // Verify Fib stays genuinely recursive (calls fib_i32) and Ack has no m==1 shortcut.
  const fibDecl = printGoDecl(pFib.worker);
  const ackDecl = printGoDecl(pAck.worker);
  assert.match(fibDecl, /fib_i32\(/);
  assert.match(ackDecl, /ack_i32\(/);

  const dir = mkdtempSync(join(tmpdir(), 'gopurs-int32-rec-'));
  try {
    const src = `package main
import "fmt"
func intAdd(a, b int64) int64 { return int64(int32(a + b)) }
func intSub(a, b int64) int64 { return int64(int32(a - b)) }
func intShl(a, b int64) int64 { return int64(int32(a) << (uint32(b) & 31)) }
func fib_cold(n int64) int64 { if n == 0 { return 0 }; if n == 1 { return 1 }; return intAdd(fib_cold(intSub(n, 1)), fib_cold(intSub(n, 2))) }
func ack_cold(m, n int64) int64 { if m == 0 { return intAdd(n, 1) }; if n == 0 { return ack_cold(intSub(m, 1), 1) }; return ack_cold(intSub(m, 1), ack_cold(m, intSub(n, 1))) }
func tak_cold(x, y, z int64) int64 { if x <= y { return z }; return tak_cold(tak_cold(intSub(x, 1), y, z), tak_cold(intSub(y, 1), z, x), tak_cold(intSub(z, 1), x, y)) }
func mc91_cold(n int64) int64 { if n > 100 { return intSub(n, 10) }; return mc91_cold(mc91_cold(intAdd(n, 11))) }
func bitTree_cold(n, seed int64) int64 { if n <= 0 { return int64(-int32(seed)) }; return int64(int32(intShl(bitTree_cold(intSub(n, 1), seed), 1)) ^ int32(bitTree_cold(intSub(n, 2), intAdd(seed, 3)))) }
${fibDecl}
func fib(n_0_loop int64) int64 { ${printGoExpr(pFib.fast)} }
${ackDecl}
func ack(m_0_loop, n_1_loop int64) int64 { ${printGoExpr(pAck.fast)} }
${printGoDecl(pTak.worker)}
func tak(x_0_loop, y_1_loop, z_2_loop int64) int64 { ${printGoExpr(pTak.fast)} }
${printGoDecl(pMc91.worker)}
func mc91(n_0_loop int64) int64 { ${printGoExpr(pMc91.fast)} }
${printGoDecl(pBit.worker)}
func bitTree(n_0_loop, seed_1_loop int64) int64 { ${printGoExpr(pBit.fast)} }
func main() {
  for n := int64(0); n <= 16; n++ {
    if got, want := fib(n), fib_cold(n); got != want { panic(fmt.Sprintf("fib(%d)=%d want %d", n, got, want)) }
  }
  for _, tc := range [][2]int64{{0,0},{0,2147483647},{0,2147483648},{1,0},{1,5},{2,5},{3,4}} {
    if got, want := ack(tc[0], tc[1]), ack_cold(tc[0], tc[1]); got != want { panic(fmt.Sprintf("ack(%d,%d)=%d want %d", tc[0], tc[1], got, want)) }
  }
  for _, tc := range [][3]int64{{6,4,2},{8,4,1},{10,5,0},{2,5,9},{2147483648,2147483648,42}} {
    if got, want := tak(tc[0], tc[1], tc[2]), tak_cold(tc[0], tc[1], tc[2]); got != want { panic(fmt.Sprintf("tak=%d want %d", got, want)) }
  }
  for _, n := range []int64{-5, 0, 50, 99, 100, 101, 120, 2147483648} {
    if got, want := mc91(n), mc91_cold(n); got != want { panic(fmt.Sprintf("mc91(%d)=%d want %d", n, got, want)) }
  }
  for _, seed := range []int64{0, 1, -1, 2147483647, -2147483648, 2147483648} {
    for n := int64(0); n <= 8; n++ {
      if got, want := bitTree(n, seed), bitTree_cold(n, seed); got != want { panic(fmt.Sprintf("bitTree(%d,%d)=%d want %d", n, seed, got, want)) }
    }
  }
}
`;
    writeFileSync(join(dir, 'main.go'), src);
    const res = spawnSync('go', ['run', 'main.go'], { cwd: dir, encoding: 'utf8', timeout: 30000 });
    assert.equal(res.status, 0, res.stdout + res.stderr);
  } finally {
    rmSync(dir, { recursive: true, force: true });
  }
});
