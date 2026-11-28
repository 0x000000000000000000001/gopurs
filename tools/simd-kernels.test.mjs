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

func refNonIncreasing(xs []int32) bool {
	for i := 1; i < len(xs); i++ {
		if xs[i] > xs[i-1] {
			return false
		}
	}
	return true
}

func refNonDecreasing(xs []int32) bool {
	for i := 1; i < len(xs); i++ {
		if xs[i] < xs[i-1] {
			return false
		}
	}
	return true
}

func refContains(xs []int32, target int32) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
	}
	return false
}

func TestMonotoneAndContainsInt32(t *testing.T) {
	lengths := make([]int, 0, 270)
	for n := 0; n <= 257; n++ {
		lengths = append(lengths, n)
	}
	lengths = append(lengths, 511, 512, 1000, 1001, 4096, 10000)

	for _, n := range lengths {
		desc := make([]int32, n)
		asc := make([]int32, n)
		eq := make([]int32, n)
		for i := 0; i < n; i++ {
			desc[i] = int32(n - i)
			asc[i] = int32(i + 1)
			eq[i] = 7
		}
		if !IsNonIncreasingInt32(desc) || !IsNonIncreasingInt32(eq) {
			t.Fatalf("IsNonIncreasingInt32 false negative n=%d", n)
		}
		if !IsNonDecreasingInt32(asc) || !IsNonDecreasingInt32(eq) {
			t.Fatalf("IsNonDecreasingInt32 false negative n=%d", n)
		}
		if n >= 2 {
			if IsNonIncreasingInt32(asc) {
				t.Fatalf("IsNonIncreasingInt32 false positive on asc n=%d", n)
			}
			if IsNonDecreasingInt32(desc) {
				t.Fatalf("IsNonDecreasingInt32 false positive on desc n=%d", n)
			}
		}
		// Single violation at every boundary position (start, near 4, 16, 20, 64, end)
		for _, pos := range []int{0, 1, 2, 3, 4, 15, 16, 17, 19, 20, 63, 64, 65, n / 2, n - 2} {
			if pos >= 0 && pos+1 < n {
				// Non-increasing with a single bump
				buf := append([]int32(nil), desc...)
				buf[pos+1] = buf[pos] + 1
				if got, want := IsNonIncreasingInt32(buf), refNonIncreasing(buf); got != want {
					t.Fatalf("IsNonIncreasingInt32 bump n=%d pos=%d: got %v want %v", n, pos, got, want)
				}
				// Duplicate allowed (non-increasing)
				buf[pos+1] = buf[pos]
				if got, want := IsNonIncreasingInt32(buf), refNonIncreasing(buf); got != want {
					t.Fatalf("IsNonIncreasingInt32 dup n=%d pos=%d: got %v want %v", n, pos, got, want)
				}
				// Non-decreasing with a single dip
				buf2 := append([]int32(nil), asc...)
				buf2[pos+1] = buf2[pos] - 1
				if got, want := IsNonDecreasingInt32(buf2), refNonDecreasing(buf2); got != want {
					t.Fatalf("IsNonDecreasingInt32 dip n=%d pos=%d: got %v want %v", n, pos, got, want)
				}
				buf2[pos+1] = buf2[pos]
				if got, want := IsNonDecreasingInt32(buf2), refNonDecreasing(buf2); got != want {
					t.Fatalf("IsNonDecreasingInt32 dup n=%d pos=%d: got %v want %v", n, pos, got, want)
				}
			}
		}
		// ContainsInt32: absent, first, middle, last, boundaries
		for _, target := range []int32{0, -1, int32(n + 1), math.MinInt32, math.MaxInt32, 1, 2, 4, 16, 17, int32(n / 2), int32(n)} {
			if got, want := ContainsInt32(asc, target), refContains(asc, target); got != want {
				t.Fatalf("ContainsInt32 n=%d target=%d: got %v want %v", n, target, got, want)
			}
		}
		// Unaligned sub-slices
		if n >= 5 {
			for off := 1; off <= 3; off++ {
				if !IsNonIncreasingInt32(desc[off:]) || !IsNonDecreasingInt32(asc[off:]) {
					t.Fatalf("unaligned monotone failed n=%d off=%d", n, off)
				}
				sub := asc[off:]
				if !ContainsInt32(sub, sub[len(sub)-1]) || ContainsInt32(sub, 0) {
					t.Fatalf("unaligned ContainsInt32 failed n=%d off=%d", n, off)
				}
			}
		}
	}

	// Signed int32 extreme comparisons (MinInt32 vs MaxInt32)
	extDesc := make([]int32, 32)
	extAsc := make([]int32, 32)
	for i := 0; i < 16; i++ {
		extDesc[i] = math.MaxInt32
		extAsc[i] = math.MinInt32
	}
	for i := 16; i < 32; i++ {
		extDesc[i] = math.MinInt32
		extAsc[i] = math.MaxInt32
	}
	if !IsNonIncreasingInt32(extDesc) || IsNonDecreasingInt32(extDesc) {
		t.Fatalf("signed extreme desc failed")
	}
	if !IsNonDecreasingInt32(extAsc) || IsNonIncreasingInt32(extAsc) {
		t.Fatalf("signed extreme asc failed")
	}
	if !ContainsInt32(extDesc, math.MinInt32) || !ContainsInt32(extDesc, math.MaxInt32) || ContainsInt32(extDesc, 0) {
		t.Fatalf("signed extreme contains failed")
	}
}

func refCountUTF8(s string) int64 {
	var c int64
	for i := 0; i < len(s); {
		_, w := DecodeStringCodePoint(s[i:])
		c++
		i += w
	}
	return c
}

func refSumCodePointsLimit(s string, limit int64) (int64, bool) {
	var sum, cnt int64
	for i := 0; cnt < limit && i < len(s); cnt++ {
		r, w := DecodeStringCodePoint(s[i:])
		sum += int64(r)
		i += w
	}
	return sum, cnt > 0
}

func TestUTF8CountAndSumSIMD(t *testing.T) {
	// 1. Exhaustive test of ALL 65,536 2-byte combinations embedded across a 16-byte SIMD chunk!
	var buf [32]byte
	for i := range buf {
		buf[i] = 'a'
	}
	for b0 := 0; b0 < 256; b0++ {
		for b1 := 0; b1 < 256; b1++ {
			buf[14] = byte(b0)
			buf[15] = byte(b1)
			s := string(buf[:])
			if got, want := CountUTF8CodePoints(s), refCountUTF8(s); got != want {
				t.Fatalf("CountUTF8 pair (%02x,%02x) at 14: got %d want %d", b0, b1, got, want)
			}
			gs, gt := SumCodePoints(s)
			ws, wt := refSumCodePointsLimit(s, int64(len(s)))
			if gs != ws || gt != wt {
				t.Fatalf("SumCodePoints pair (%02x,%02x) at 14: got (%d,%v) want (%d,%v)", b0, b1, gs, gt, ws, wt)
			}
			// Straddling chunk boundary (15, 16)
			buf[14] = 'a'
			buf[15] = byte(b0)
			buf[16] = byte(b1)
			s2 := string(buf[:])
			if got, want := CountUTF8CodePoints(s2), refCountUTF8(s2); got != want {
				t.Fatalf("CountUTF8 pair (%02x,%02x) at 15: got %d want %d", b0, b1, got, want)
			}
			gs2, gt2 := SumCodePoints(s2)
			ws2, wt2 := refSumCodePointsLimit(s2, int64(len(s2)))
			if gs2 != ws2 || gt2 != wt2 {
				t.Fatalf("SumCodePoints pair (%02x,%02x) at 15: got (%d,%v) want (%d,%v)", b0, b1, gs2, gt2, ws2, wt2)
			}
			buf[16] = 'a'
		}
	}

	// 2. Corpus of valid Unicode, Hangul (0xED < 0xA0), WTF-8 surrogates, overlongs, and malformed bytes
	tokens := []string{
		"a", "é", "€", "한", "😀", "𝄞",
		"\\xed\\x9f\\xbf", "\\xed\\xa0\\x80", "\\xed\\xb0\\x80", "\\xed\\xa0\\xbd\\xed\\xb8\\x80",
		"\\xe0\\x9f\\x80", "\\xe0\\xa0\\x80", "\\xf0\\x8f\\x80\\x80", "\\xf0\\x90\\x80\\x80",
		"\\xf4\\x8f\\xbf\\xbf", "\\xf4\\x90\\x80\\x80", "\\xf5\\x80\\x80\\x80", "\\xc0\\x80", "\\xc1\\xbf", "\\x80", "\\xff",
	}
	rng := rand.New(rand.NewSource(2026))
	for iter := 0; iter < 2000; iter++ {
		var b []byte
		k := rng.Intn(40)
		for j := 0; j < k; j++ {
			b = append(b, tokens[rng.Intn(len(tokens))]...)
		}
		s := string(b)
		if got, want := CountUTF8CodePoints(s), refCountUTF8(s); got != want {
			t.Fatalf("CountUTF8 token fuzz %d: got %d want %d", iter, got, want)
		}
		for _, lim := range []int64{-1, 0, 1, 15, 16, 17, 32, int64(len(s))} {
			gs, gt := SumCodePointsLimit(s, lim)
			ws, wt := refSumCodePointsLimit(s, lim)
			if gs != ws || gt != wt {
				t.Fatalf("SumCodePointsLimit fuzz %d lim=%d: got (%d,%v) want (%d,%v)", iter, lim, gs, gt, ws, wt)
			}
		}
	}
	// Long multi-block valid strings (> 4096 bytes)
	var longBuf []byte
	for i := 0; i < 1000; i++ {
		longBuf = append(longBuf, "a😀é한€"...)
	}
	longStr := string(longBuf)
	if got, want := CountUTF8CodePoints(longStr), refCountUTF8(longStr); got != want {
		t.Fatalf("CountUTF8 long: got %d want %d", got, want)
	}
	for _, lim := range []int64{0, 1, 16, 100, 1500, 4999, 5000, 5001} {
		gs, gt := SumCodePointsLimit(longStr, lim)
		ws, wt := refSumCodePointsLimit(longStr, lim)
		if gs != ws || gt != wt {
			t.Fatalf("SumCodePointsLimit long lim=%d: got (%d,%v) want (%d,%v)", lim, gs, gt, ws, wt)
		}
	}
}

func TestTypeStringArrayRoundTrip(t *testing.T) {
	sa := StringArray([]string{"ab", "cd", "é"})
	if sa.Type != TypeStringArray || ArrayLength(sa) != 3 || ArrayAccess(sa, 2).StrVal() != "é" {
		t.Fatalf("TypeStringArray basic ops failed")
	}
	vals := ArrayValues(sa)
	if len(vals) != 3 || vals[0].StrVal() != "ab" || vals[2].StrVal() != "é" {
		t.Fatalf("ArrayValues(TypeStringArray) failed")
	}
	unboxed := Unbox[[]string](sa)
	if len(unboxed) != 3 || unboxed[1] != "cd" {
		t.Fatalf("Unbox[[]string] failed")
	}
	boxed := Box([]string{"x", "y"})
	if boxed.Type != TypeStringArray || ArrayLength(boxed) != 2 {
		t.Fatalf("Box([]string) did not produce TypeStringArray")
	}
	sum, anyB, idx, off := SumASCIIStringArray("hi", []string{"ab", "cd"})
	if sum != int64('h'+'i'+'a'+'b'+'c'+'d') || !anyB || idx != -2 || off != 0 {
		t.Fatalf("SumASCIIStringArray ASCII failed: %d %v %d %d", sum, anyB, idx, off)
	}
	sum2, _, idx2, off2 := SumASCIIStringArray("", []string{"ab", "xé"})
	if sum2 != int64('a'+'b'+'x') || idx2 != 1 || off2 != 1 {
		t.Fatalf("SumASCIIStringArray non-ASCII failed: %d %d %d", sum2, idx2, off2)
	}
	rep := "hello"
	parts := make([]string, 50)
	for i := range parts {
		parts[i] = rep
	}
	sum3, any3, idx3, off3 := SumASCIIStringArray("!", parts)
	want3 := int64('!') + 50*int64('h'+'e'+'l'+'l'+'o')
	if sum3 != want3 || !any3 || idx3 != -2 || off3 != 0 {
		t.Fatalf("SumASCIIStringArray repeated failed: %d want %d (%v %d %d)", sum3, want3, any3, idx3, off3)
	}
	parts2 := []string{rep, rep, "", rep[:3], rep, "z\x80"}
	sum4, any4, idx4, off4 := SumASCIIStringArray("", parts2)
	want4 := 3*int64('h'+'e'+'l'+'l'+'o') + int64('h'+'e'+'l') + int64('z')
	if sum4 != want4 || !any4 || idx4 != 5 || off4 != 1 {
		t.Fatalf("SumASCIIStringArray mixed slices failed: %d want %d (%v %d %d)", sum4, want4, any4, idx4, off4)
	}
	if got := SumOddProgression(500, 0); got != 21536 {
		t.Fatalf("SumOddProgression(500, 0) = %d want 21536", got)
	}
	if got := SumOddProgression(2, 17); got != 19 {
		t.Fatalf("SumOddProgression(2, 17) = %d want 19", got)
	}
}

func TestFilterMaskAffineParityInt32AndASCIIUTF8(t *testing.T) {
	rng := rand.New(rand.NewSource(99))
	coeffs := [][4]int32{
		{1, 0, 2, 0},
		{1, 0, 1, 0},
		{3, -5, -7, 11},
		{math.MaxInt32, 1, math.MinInt32, -1},
	}
	masks := [][2]int32{{1, 0}, {1, 1}, {3, 0}, {3, 2}, {7, 5}, {0, 0}, {0, 1}, {15, 13}}
	for _, n := range []int{0, 1, 2, 3, 4, 5, 15, 16, 17, 31, 32, 33, 63, 64, 65, 257, 1000} {
		xs := make([]int32, n)
		for i := range xs {
			xs[i] = int32(rng.Uint32())
		}
		if n >= 4 {
			xs[0], xs[1], xs[2], xs[3] = math.MinInt32, math.MaxInt32, -1, 0
		}
		for _, m := range masks {
			for _, c := range coeffs {
				gotSum, gotCnt := FilterMaskAffineParitySumInt32(xs, m[0], m[1], c[0], c[1], c[2], c[3])
				var wantSum int32
				var wantCnt int64
				for _, v := range xs {
					if (v & m[0]) == m[1] {
						if wantCnt&1 == 0 {
							wantSum += c[0]*v + c[1]
						} else {
							wantSum += c[2]*v + c[3]
						}
						wantCnt++
					}
				}
				if gotSum != wantSum || gotCnt != wantCnt {
					t.Fatalf("FilterMask n=%d mask=%v c=%v: got (%d,%d) want (%d,%d)", n, m, c, gotSum, gotCnt, wantSum, wantCnt)
				}
			}
		}
	}
	// Pure-ASCII (including max-byte 0x7F torture for uint16 lane headroom) and ASCII-prefix+Unicode/malformed tests
	maxAscii := make([]byte, 10000)
	for i := range maxAscii {
		maxAscii[i] = 0x7F
	}
	if gs, gt := SumCodePointsLimit(string(maxAscii), 10000); gs != 10000*0x7F || !gt {
		t.Fatalf("max-ASCII 0x7F SumCodePointsLimit: got (%d,%v) want (%d,true)", gs, gt, 10000*0x7F)
	}
	asciiBuf := make([]byte, 10000)
	for i := range asciiBuf {
		asciiBuf[i] = byte(i % 128)
	}
	for _, n := range []int{0, 15, 16, 63, 64, 65, 127, 128, 129, 3999, 4000, 8192, 8193, 9999} {
		s := string(asciiBuf[:n])
		if got := CountUTF8CodePoints(s); got != int64(n) {
			t.Fatalf("ASCII CountUTF8 n=%d: got %d", n, got)
		}
		for _, lim := range []int64{0, 1, 15, 16, 63, 64, 65, int64(n) / 2, int64(n), int64(n) + 10} {
			gs, gt := SumCodePointsLimit(s, lim)
			ws, wt := refSumCodePointsLimit(s, lim)
			if gs != ws || gt != wt {
				t.Fatalf("ASCII SumCodePointsLimit n=%d lim=%d: got (%d,%v) want (%d,%v)", n, lim, gs, gt, ws, wt)
			}
		}
	}
	// ASCII prefix followed by Unicode or malformed/surrogate bytes at various offsets (triggers bailout + rewind)
	for _, prefixLen := range []int{0, 1, 15, 16, 63, 64, 65, 128, 250, 9000} {
		for _, tail := range []string{"é", "😀", "\xED\xA0\x80", "\xFFabc", "é" + string(asciiBuf[:200])} {
			s := string(asciiBuf[:prefixLen]) + tail
			if got, want := CountUTF8CodePoints(s), refCountUTF8(s); got != want {
				t.Fatalf("mixed CountUTF8 pref=%d: got %d want %d", prefixLen, got, want)
			}
			for _, lim := range []int64{int64(prefixLen), int64(prefixLen) + 1, int64(len(s))} {
				gs, gt := SumCodePointsLimit(s, lim)
				ws, wt := refSumCodePointsLimit(s, lim)
				if gs != ws || gt != wt {
					t.Fatalf("mixed SumCodePointsLimit pref=%d lim=%d: got (%d,%v) want (%d,%v)", prefixLen, lim, gs, gt, ws, wt)
				}
			}
		}
	}
	// Exhaustive IndexByteASCII tests (ASCII hit/miss, non-ASCII before/at/after hit, boundary lengths)
	for _, n := range []int{0, 1, 15, 16, 17, 63, 64, 65, 127, 128, 129, 10000} {
		buf := make([]byte, n)
		for i := range buf {
			buf[i] = 'a'
		}
		if idx, ascii := IndexByteASCII(string(buf), 'z'); idx != -1 || ascii {
			t.Fatalf("IndexByteASCII miss n=%d: got (%d,%v)", n, idx, ascii)
		}
		if n > 0 {
			for _, pos := range []int{0, n / 2, n - 1} {
				buf[pos] = 'z'
				if idx, ascii := IndexByteASCII(string(buf), 'z'); idx != pos || !ascii {
					t.Fatalf("IndexByteASCII pure ASCII n=%d pos=%d: got (%d,%v)", n, pos, idx, ascii)
				}
				// Non-ASCII strictly AFTER hit must still report prefixASCII=true
				if pos+1 < n {
					buf[pos+1] = 0xE2
					if idx, ascii := IndexByteASCII(string(buf), 'z'); idx != pos || !ascii {
						t.Fatalf("IndexByteASCII non-ASCII after hit n=%d pos=%d: got (%d,%v)", n, pos, idx, ascii)
					}
					buf[pos+1] = 'a'
				}
				// Non-ASCII strictly BEFORE hit must report prefixASCII=false
				if pos > 0 {
					for _, bad := range []int{0, pos - 1} {
						buf[bad] = 0xC3
						if idx, ascii := IndexByteASCII(string(buf), 'z'); idx != pos || ascii {
							t.Fatalf("IndexByteASCII non-ASCII before hit n=%d pos=%d bad=%d: got (%d,%v)", n, pos, bad, idx, ascii)
						}
						buf[bad] = 'a'
					}
				}
				buf[pos] = 'a'
			}
		}
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
  // Verify amd64 SSE2 kernels execute cleanly under Rosetta 2 on darwin/arm64
  if (process.platform === 'darwin' && process.arch === 'arm64') {
    run(['test', '-count=1', './gopurs_runtime'], { GOOS: 'darwin', GOARCH: 'amd64' });
  }
  // Verify pure-Go fallback compiles cleanly (riscv64 and wasm) AND amd64 SSE2 assembles cleanly
  run(['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'amd64' });
  run(['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'arm64' });
  run(['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'riscv64' });
  run(['build', './gopurs_runtime'], { GOOS: 'js', GOARCH: 'wasm' });
});
