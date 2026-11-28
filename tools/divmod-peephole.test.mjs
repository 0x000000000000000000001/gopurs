import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { loop } from '../output/Gopurs.Int32Loops/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';

const i64 = Go.TypeInt64.value;
const variable = name => new Go.GoVar(name);
const cast = (name, value) => new Go.GoCall(variable(name), [value]);
const integer = n => cast('int64', new Go.GoInt(n));
const runtime = (name, ...args) => new Go.GoCall(new Go.GoSelector(variable('gopurs_runtime'), name), args);
const binary = (op, left, right) => new Go.GoBinOp(op, left, right);

test('Int32Loops narrows constant positive IntModPos64, IntDivPos64 and power-of-two mask on native slots', () => {
  const n = variable('n'), acc = variable('acc');
  const params = [new Tuple('n', i64), new Tuple('acc', i64)];
  const stmts = [
    new Go.GoIfElse(binary('==', n, integer(0)), [new Go.GoReturn(acc)], []),
    new Go.GoMutate('n_loop', runtime('IntSub', n, integer(1))),
    new Go.GoMutate('acc_loop', runtime('IntAdd', acc, runtime('IntModPos64', n, integer(3)))),
    new Go.GoContinue('again'),
  ];
  const code = printGoExpr(loop('again')(params)(i64)(stmts));
  assert.match(code, /IntModPos32\(n, int32\(3\)\)/);
  assert.doesNotMatch(code, /IntModPos64\(int64\(n\)/);

  const divStmts = [
    new Go.GoIfElse(binary('==', n, integer(0)), [new Go.GoReturn(acc)], []),
    new Go.GoMutate('n_loop', runtime('IntSub', n, integer(1))),
    new Go.GoMutate('acc_loop', runtime('IntAdd', acc, runtime('IntDivPos64', n, integer(5)))),
    new Go.GoContinue('again'),
  ];
  const divCode = printGoExpr(loop('again')(params)(i64)(divStmts));
  assert.match(divCode, /IntDivPos32\(n, int32\(5\)\)/);

  const maskStmts = [
    new Go.GoIfElse(binary('==', n, integer(0)), [new Go.GoReturn(acc)], []),
    new Go.GoMutate('n_loop', runtime('IntSub', n, integer(1))),
    new Go.GoMutate('acc_loop', runtime('IntAdd', acc, binary('&', n, new Go.GoInt(7)))),
    new Go.GoContinue('again'),
  ];
  const maskCode = printGoExpr(loop('again')(params)(i64)(maskStmts));
  assert.match(maskCode, /\(n\) & \(int32\(7\)\)/);
});

test('Constant positive div/mod helpers, power-of-two shifts/masks, and IntModZero match IntDiv/IntMod across all edge values', t => {
  const dir = mkdtempSync(join(tmpdir(), 'gopurs-divmod-oracle-'));
  t.after(() => rmSync(dir, { recursive: true, force: true }));
  mkdirSync(join(dir, 'gopurs_runtime'));
  writeFileSync(join(dir, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(dir, 'gopurs_runtime/runtime.go'), runtimeGoCode);
  writeFileSync(
    join(dir, 'gopurs_runtime/divmod_test.go'),
    `package gopurs_runtime

import (
	"math"
	"testing"
)

func TestDivModHelpersExact(t *testing.T) {
	i32Values := []int32{
		math.MinInt32, math.MinInt32 + 1, -100000, -17, -10, -9, -8, -7, -5, -4, -3, -2, -1,
		0, 1, 2, 3, 4, 5, 7, 8, 9, 10, 17, 100000, math.MaxInt32 - 1, math.MaxInt32,
	}
	i64Values := []int64{
		math.MinInt64, math.MinInt64 + 1, -4294967296, -2147483649,
		math.MinInt32, math.MinInt32 + 1, -100000, -17, -10, -9, -8, -7, -5, -4, -3, -2, -1,
		0, 1, 2, 3, 4, 5, 7, 8, 9, 10, 17, 100000, math.MaxInt32 - 1, math.MaxInt32,
		2147483648, 4294967295, 4294967296, math.MaxInt64 - 1, math.MaxInt64,
	}
	posDivisors := []int32{1, 2, 3, 4, 5, 7, 8, 10, 16, 17, 64, 100, 1024, 1000003, 1 << 30, math.MaxInt32}

	for _, a := range i32Values {
		for _, d := range posDivisors {
			if got, want := int64(IntModPos32(a, d)), IntMod(int64(a), int64(d)); got != want {
				t.Fatalf("IntModPos32(%d, %d): got %d want %d", a, d, got, want)
			}
			if got, want := int64(IntDivPos32(a, d)), IntDiv(int64(a), int64(d)); got != want {
				t.Fatalf("IntDivPos32(%d, %d): got %d want %d", a, d, got, want)
			}
		}
	}

	for _, a := range i64Values {
		for _, d := range posDivisors {
			d64 := int64(d)
			if got, want := IntModPos64(a, d64), IntMod(a, d64); got != want {
				t.Fatalf("IntModPos64(%d, %d): got %d want %d", a, d64, got, want)
			}
			if got, want := IntDivPos64(a, d64), IntDiv(a, d64); got != want {
				t.Fatalf("IntDivPos64(%d, %d): got %d want %d", a, d64, got, want)
			}
			if d&(d-1) == 0 {
				if got, want := a&(d64-1), IntMod(a, d64); got != want {
					t.Fatalf("pow2 mask(%d, %d): got %d want %d", a, d64, got, want)
				}
				shift := uint(0)
				for x := d; x > 1; x >>= 1 {
					shift++
				}
				if got, want := a>>shift, IntDiv(a, d64); got != want {
					t.Fatalf("pow2 shift(%d, %d): got %d want %d", a, d64, got, want)
				}
			}
			if got, want := (a%d64 == 0), (IntMod(a, d64) == 0); got != want {
				t.Fatalf("const mod zero(%d, %d): got %v want %v", a, d64, got, want)
			}
		}
		for _, b := range []int64{math.MinInt64, -100, -7, -3, -2, -1, 0, 1, 2, 3, 7, 100, math.MaxInt64} {
			if got, want := IntModZero(a, b), (IntMod(a, b) == 0); got != want {
				t.Fatalf("IntModZero(%d, %d): got %v want %v", a, b, got, want)
			}
		}
	}

	// Exhaustive boundary + 200,000 deterministic pseudo-random int32 pairs for Lemire Divisor32.
	div32Edge := []int32{
		math.MinInt32, math.MinInt32 + 1, -1000003, -65537, -100, -17, -8, -7, -4, -3, -2, -1,
		0, 1, 2, 3, 4, 7, 8, 17, 100, 65537, 1000003, math.MaxInt32 - 1, math.MaxInt32,
	}
	for _, d := range div32Edge {
		div := NewDivisor32(d)
		for _, a := range div32Edge {
			want := IntMod(int64(a), int64(d)) == 0
			if got := div.Divides(a); got != want {
				t.Fatalf("Divisor32(%d).Divides(%d): got %v want %v", d, a, got, want)
			}
		}
	}
	state := uint32(0x12345678)
	nextRand := func() int32 {
		state ^= state << 13
		state ^= state >> 17
		state ^= state << 5
		return int32(state)
	}
	for i := 0; i < 200000; i++ {
		d := nextRand()
		if i%4 == 0 {
			// Exercise small divisors heavily as well as full-range random divisors.
			d = d % 1024
		}
		a := nextRand()
		want := IntMod(int64(a), int64(d)) == 0
		if got := NewDivisor32(d).Divides(a); got != want {
			t.Fatalf("random Divisor32(%d).Divides(%d): got %v want %v", d, a, got, want)
		}
	}
}
`,
  );

  for (const [label, args, envExtra] of [
    ['arm64 -race', ['test', '-race', '-count=1', './gopurs_runtime'], {}],
    ['amd64 Rosetta', ['test', '-count=1', './gopurs_runtime'], { GOARCH: 'amd64' }],
  ]) {
    const res = spawnSync('go', args, {
      cwd: dir,
      encoding: 'utf8',
      timeout: 60_000,
      env: { ...process.env, GOWORK: 'off', ...envExtra },
    });
    assert.equal(res.status, 0, `${label}:\n${res.stdout}\n${res.stderr}`);
  }
});
