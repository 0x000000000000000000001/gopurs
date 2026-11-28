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
import { plan } from '../output/Gopurs.NestedRecursion/index.js';
import { printGoDecl, printGoExpr } from '../output/Gopurs.Printer/index.js';

const expr = syntax => new TcoExpr({}, syntax);
const lit = n => expr(new S.Lit(new C.LitInt(n)));
const local = (name, level) => expr(new S.Local(new Just(name), level));
const gvar = name => expr(new S.Var(new C.Qualified(Nothing.value, name)));
const call = (name, args) => expr(new S.App(gvar(name), args));
const op2 = (op, a, b) => expr(new S.PrimOp(new S.Op2(op, a, b)));
const eq = (a, b) => op2(new S.OpIntOrd(S.OpEq.value), a, b);
const add = (a, b) => op2(new S.OpIntNum(S.OpAdd.value), a, b);
const sub = (a, b) => op2(new S.OpIntNum(S.OpSubtract.value), a, b);
const branch = (pairs, fallback) => expr(new S.Branch(pairs.map(([c, v]) => new S.Pair(c, v)), fallback));

test('NestedRecursion handles standard, custom-step, swapped, wide and overflow cases identically to JS', () => {
  const ackBody = branch(
    [
      [eq(local('m', 0), lit(0)), add(local('n', 1), lit(1))],
      [eq(local('n', 1), lit(0)), call('ack', [sub(local('m', 0), lit(1)), lit(1)])],
    ],
    call('ack', [sub(local('m', 0), lit(1)), call('ack', [local('m', 0), sub(local('n', 1), lit(1))])]),
  );
  // Custom affine base + swapped params + incrementing b toward b0=5 + negative k=-3 + maxInt32 reset
  const customBody = branch(
    [
      [eq(local('a', 1), lit(2)), sub(local('b', 0), lit(3))],
      [eq(local('b', 0), lit(5)), call('custom', [lit(2), sub(local('a', 1), lit(2))])],
    ],
    call('custom', [call('custom', [add(local('b', 0), lit(1)), local('a', 1)]), sub(local('a', 1), lit(2))]),
  );

  const pAck = plan('M')('ack')([new Tuple('m_0', Go.TypeInt64.value), new Tuple('n_1', Go.TypeInt64.value)])(Go.TypeInt64.value)(ackBody)(generated('M')('ack_i32'))(new Go.GoCall(new Go.GoVar('ack_cold'), [new Go.GoVar('m_0_loop'), new Go.GoVar('n_1_loop')]));
  const pCustom = plan('M')('custom')([new Tuple('b_0', Go.TypeInt64.value), new Tuple('a_1', Go.TypeInt64.value)])(Go.TypeInt64.value)(customBody)(generated('M')('custom_i32'))(new Go.GoCall(new Go.GoVar('custom_cold'), [new Go.GoVar('b_0_loop'), new Go.GoVar('a_1_loop')]));
  assert.ok(pAck instanceof Just, 'ack recognized');
  assert.ok(pCustom instanceof Just, 'custom swapped recognized');

  const dir = mkdtempSync(join(tmpdir(), 'gopurs-nested-rec-'));
  try {
    const src = `package main
import "fmt"
func intAdd(a, b int64) int64 { return int64(int32(a + b)) }
func intSub(a, b int64) int64 { return int64(int32(a - b)) }
func ack_cold(m, n int64) int64 {
  if m == 0 { return intAdd(n, 1) }
  if n == 0 { return ack_cold(intSub(m, 1), 1) }
  return ack_cold(intSub(m, 1), ack_cold(m, intSub(n, 1)))
}
func custom_cold(b, a int64) int64 {
  if a == 2 { return intSub(b, 3) }
  if b == 5 { return custom_cold(2, intSub(a, 2)) }
  return custom_cold(custom_cold(intAdd(b, 1), a), intSub(a, 2))
}
${printGoDecl(pAck.value0.worker)}
func ack(m_0_loop, n_1_loop int64) int64 { ${printGoExpr(pAck.value0.fast)} }
${printGoDecl(pCustom.value0.worker)}
func custom(b_0_loop, a_1_loop int64) int64 { ${printGoExpr(pCustom.value0.fast)} }
func main() {
  for _, tc := range [][2]int64{{0,0},{0,5},{0,2147483647},{0,2147483648},{1,0},{1,1},{1,20},{2,0},{2,5},{3,0},{3,4},{3,5}} {
    if got, want := ack(tc[0], tc[1]), ack_cold(tc[0], tc[1]); got != want {
      panic(fmt.Sprintf("ack(%d,%d)=%d want %d", tc[0], tc[1], got, want))
    }
  }
  // Overflow on large n at m=1
  if got, want := ack(1, 2147483647), intAdd(2147483647, 2); got != want {
    panic(fmt.Sprintf("ack(1,MaxInt32)=%d want %d", got, want))
  }
  for _, tc := range [][2]int64{{0,2},{5,2},{2147483648,2},{5,4},{4,4},{0,4},{-10,4},{5,6},{4,6},{2,6}} {
    if got, want := custom(tc[0], tc[1]), custom_cold(tc[0], tc[1]); got != want {
      panic(fmt.Sprintf("custom(%d,%d)=%d want %d", tc[0], tc[1], got, want))
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
