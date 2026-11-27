import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';

const runtimeDir = new URL('../runtime/', import.meta.url);
const runtimeFiles = ['runtime.go', 'simd_asm.go', 'simd_generic.go', 'simd_arm64.s', 'simd_amd64.s'];

test('SIMD int32 kernels match scalar oracle, pass vet/race and cross-compile across arm64, amd64 and generic fallback', t => {
  const dir = mkdtempSync(join(tmpdir(), 'gopurs-simd-kernels-'));
  t.after(() => rmSync(dir, { recursive: true, force: true }));
  mkdirSync(join(dir, 'gopurs_runtime'));
  writeFileSync(join(dir, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  for (const name of runtimeFiles) {
    copyFileSync(new URL(name, runtimeDir), join(dir, 'gopurs_runtime', name));
  }
  writeFileSync(
    join(dir, 'gopurs_runtime/simd_test.go'),
    `package gopurs_runtime

import (
	"math"
	"math/rand"
	"testing"
)

func refSum(xs []int32) int32 {
	var acc int32
	for _, x := range xs {
		acc += x
	}
	return acc
}

func refAffine(xs []int32, aEven, bEven, aOdd, bOdd int32) int32 {
	var acc int32
	for i, x := range xs {
		if i&1 == 0 {
			acc += aEven*x + bEven
		} else {
			acc += aOdd*x + bOdd
		}
	}
	return acc
}

func TestSumAndAffineParityInt32(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	coeffs := [][4]int32{
		{1, 0, 1, 0},
		{1, 0, 2, 0},
		{1, 1, 2, 2},
		{-1, 7, 3, -11},
		{0, 5, 0, -9},
		{2097152, math.MaxInt32, -2097152, math.MinInt32},
		{math.MaxInt32, math.MinInt32, math.MinInt32, math.MaxInt32},
	}
	lengths := make([]int, 0, 265)
	for n := 0; n <= 257; n++ {
		lengths = append(lengths, n)
	}
	lengths = append(lengths, 1000, 1001, 4095, 4096, 10000)

	for _, n := range lengths {
		for mode := 0; mode < 4; mode++ {
			xs := make([]int32, n)
			for i := range xs {
				switch mode {
				case 0:
					xs[i] = int32(i + 1)
				case 1:
					xs[i] = math.MaxInt32
				case 2:
					xs[i] = math.MinInt32
				default:
					xs[i] = int32(rng.Uint32())
				}
			}
			if got, want := SumInt32(xs), refSum(xs); got != want {
				t.Fatalf("SumInt32 n=%d mode=%d: got %d, want %d", n, mode, got, want)
			}
			for _, c := range coeffs {
				got := AffineParitySumInt32(xs, c[0], c[1], c[2], c[3])
				want := refAffine(xs, c[0], c[1], c[2], c[3])
				if got != want {
					t.Fatalf("AffineParitySumInt32 n=%d mode=%d c=%v: got %d, want %d", n, mode, c, got, want)
				}
			}
			// Test unaligned sub-slices (offset 1, 2, 3) when n >= 4
			if n >= 4 && mode == 3 {
				for off := 1; off <= 3; off++ {
					sub := xs[off:]
					if got, want := SumInt32(sub), refSum(sub); got != want {
						t.Fatalf("unaligned SumInt32 n=%d off=%d: got %d, want %d", n, off, got, want)
					}
					got := AffineParitySumInt32(sub, 1, 1, 2, 2)
					want := refAffine(sub, 1, 1, 2, 2)
					if got != want {
						t.Fatalf("unaligned AffineParitySumInt32 n=%d off=%d: got %d, want %d", n, off, got, want)
					}
				}
			}
		}
	}
}

func TestTypeInt32ArrayRoundTrip(t *testing.T) {
	compact := IntArray([]int64{math.MinInt32, -1, 0, 1, math.MaxInt32})
	if compact.Type != TypeInt32Array {
		t.Fatalf("expected TypeInt32Array, got %d", compact.Type)
	}
	if ArrayLength(compact) != 5 || ArrayAccess(compact, 0).IntVal != math.MinInt32 || ArrayAccess(compact, 4).IntVal != math.MaxInt32 {
		t.Fatalf("ArrayLength/ArrayAccess mismatch")
	}
	wide := IntArray([]int64{1, math.MaxInt32 + 1})
	if wide.Type != TypeIntArray {
		t.Fatalf("expected TypeIntArray for wide int64, got %d", wide.Type)
	}
	boxed := Array([]Value{Int(math.MinInt32), Int(-1), Int(0), Int(1), Int(math.MaxInt32)})
	if !EqualIntArrayValues(compact, boxed) || !EqualIntArrayValues(boxed, compact) {
		t.Fatalf("EqualIntArrayValues mismatch between compact and boxed")
	}
	if !EqualIntArrayWithValue(UnboxIntArray(compact), compact) {
		t.Fatalf("EqualIntArrayWithValue mismatch")
	}
	borrowed := BorrowIntArray(wide)
	if len(borrowed) != 2 || &borrowed[0] != &(*(*[]int64)(wide.UnsafePtr))[0] {
		t.Fatalf("BorrowIntArray must zero-copy TypeIntArray")
	}
}
`,
  );

  const run = (args, extraEnv = {}) => {
    const res = spawnSync('go', args, {
      cwd: dir,
      encoding: 'utf8',
      timeout: 60_000,
      env: { ...process.env, GOWORK: 'off', ...extraEnv },
    });
    assert.ifError(res.error);
    assert.equal(res.status, 0, `go ${args.join(' ')} (${JSON.stringify(extraEnv)}):\n${res.stdout}\n${res.stderr}`);
    return res;
  };

  run(['vet', './gopurs_runtime']);
  run(['test', '-race', '-count=1', './gopurs_runtime']);
  // Verify pure-Go fallback compiles cleanly (riscv64 and wasm) AND amd64 SSE2 assembles cleanly
  run(['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'amd64' });
  run(['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'arm64' });
  run(['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'riscv64' });
  run(['build', './gopurs_runtime'], { GOOS: 'js', GOARCH: 'wasm' });
});
