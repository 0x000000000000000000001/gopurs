import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';

const runtimeDir = new URL('../runtime/', import.meta.url);
const runtimeFiles = ['runtime.go', 'simd_asm.go', 'simd_generic.go', 'simd_arm64.s', 'simd_amd64.s'];

const kernelTestSource = `package gopurs_runtime

import (
	"math/rand"
	"strings"
	"testing"
	"unicode/utf8"
)

func refSum(xs []int32) int32 {
	var s int32
	for _, x := range xs {
		s += x
	}
	return s
}

func refIndexOf(xs []int32, target int32) int {
	for i, x := range xs {
		if x == target {
			return i
		}
	}
	return -1
}

func refFirstNonASCII(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] >= 0x80 {
			return i
		}
	}
	return len(s)
}

func refCountUTF8(s string) int64 {
	var n int64
	for i := 0; i < len(s); {
		_, w := DecodeStringCodePoint(s[i:])
		n++
		i += w
	}
	return n
}

func refAdvanceUTF8(s string, n int64) (int, bool) {
	if n <= 0 || len(s) == 0 {
		return 0, false
	}
	off := 0
	malformed := false
	for count := int64(0); count < n && off < len(s); count++ {
		r, w := DecodeStringCodePoint(s[off:])
		if r == utf8.RuneError && w == 1 {
			malformed = true
		}
		off += w
	}
	return off, malformed
}

func refIndexByteASCII(s string, c byte) (int, bool) {
	allASCII := true
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i, allASCII
		}
		if s[i] >= 0x80 {
			allASCII = false
		}
	}
	return -1, false
}

func TestSumInt32(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 15, 16, 17, 31, 32, 33, 63, 64, 65, 127, 128, 1000} {
		xs := make([]int32, n)
		for i := range xs {
			xs[i] = int32(rng.Uint32())
		}
		if got, want := SumInt32(xs), refSum(xs); got != want {
			t.Fatalf("SumInt32 n=%d: got %d, want %d", n, got, want)
		}
	}
}

func TestMonotoneAndIndexOfInt32(t *testing.T) {
	for _, n := range []int{0, 1, 2, 3, 4, 5, 15, 16, 17, 19, 20, 21, 63, 64, 65, 84, 100, 1000} {
		desc := make([]int32, n)
		asc := make([]int32, n)
		for i := 0; i < n; i++ {
			desc[i] = int32(n - i)
			asc[i] = int32(i)
		}
		if !IsNonIncreasingInt32(desc) {
			t.Fatalf("IsNonIncreasingInt32 false on desc n=%d", n)
		}
		if !IsNonDecreasingInt32(asc) {
			t.Fatalf("IsNonDecreasingInt32 false on asc n=%d", n)
		}
		if n >= 2 {
			desc[n/2] = desc[n/2] + 5
			if IsNonIncreasingInt32(desc) {
				t.Fatalf("IsNonIncreasingInt32 missed break n=%d", n)
			}
			asc[n/2] = asc[n/2] - 5
			if IsNonDecreasingInt32(asc) {
				t.Fatalf("IsNonDecreasingInt32 missed break n=%d", n)
			}
		}
		for _, target := range []int32{0, 1, int32(n / 2), int32(n - 1), int32(n + 100), -5} {
			got := IndexOfInt32(asc, target)
			want := refIndexOf(asc, target)
			if got != want {
				t.Fatalf("IndexOfInt32 n=%d target=%d: got %d, want %d", n, target, got, want)
			}
			if ContainsInt32(asc, target) != (want >= 0) {
				t.Fatalf("ContainsInt32 mismatch n=%d target=%d", n, target)
			}
		}
		// Duplicate elements must return the lowest index.
		if n >= 4 {
			dups := make([]int32, n)
			for i := range dups {
				dups[i] = int32(i % 7)
			}
			for _, target := range []int32{0, 3, 6, 9} {
				if got, want := IndexOfInt32(dups, target), refIndexOf(dups, target); got != want {
					t.Fatalf("IndexOfInt32 dups n=%d target=%d: got %d, want %d", n, target, got, want)
				}
			}
		}
	}
}

func TestUTF8ASCIIAndAdvance(t *testing.T) {
	corpus := []string{
		"", "a", "abcdefgh", "abcdefghi",
		strings.Repeat("a", 63), strings.Repeat("a", 64), strings.Repeat("a", 65),
		strings.Repeat("a", 128) + "é",
		"é", "a😀é",
		strings.Repeat("a😀é", 200),
		strings.Repeat("日本語テスト", 300),
		strings.Repeat("hello world! ", 200) + "🎯",
		"\\xed\\xa0\\x80", "\\xed\\xb0\\x80", "\\xed\\xa0\\xbd\\xed\\xb8\\x80",
		"\\xed\\x9f\\xbf",
		"\\xe0\\x9f\\x80", "\\xf0\\x8f\\x80\\x80", "\\xf4\\x90\\x80\\x80",
		"\\xc0\\x80", "\\x80", "\\xff",
		strings.Repeat("a😀é", 30) + "\\xed\\xa0\\xbd\\xed\\xb8\\x80" + strings.Repeat("é", 40),
		strings.Repeat("a😀é", 30) + "\\xff" + strings.Repeat("é", 40),
	}
	for _, s := range corpus {
		if got, want := FirstNonASCII(s), refFirstNonASCII(s); got != want {
			t.Fatalf("FirstNonASCII(%q): got %d, want %d", s, got, want)
		}
		if IsASCII(s) != (refFirstNonASCII(s) == len(s)) {
			t.Fatalf("IsASCII mismatch")
		}
		if got, want := CountUTF8CodePoints(s), refCountUTF8(s); got != want {
			t.Fatalf("CountUTF8CodePoints(%q): got %d, want %d", s, got, want)
		}
		for _, limit := range []int64{-1, 0, 1, 2, 15, 16, 17, 31, 32, 33, 64, 100, 500, 1500, 100000} {
			goff, gmal := AdvanceUTF8CodePoints(s, limit)
			woff, wmal := refAdvanceUTF8(s, limit)
			if goff != woff || gmal != wmal {
				t.Fatalf("AdvanceUTF8CodePoints(len=%d, n=%d): got (%d,%v), want (%d,%v)", len(s), limit, goff, gmal, woff, wmal)
			}
		}
		for _, c := range []byte{'a', '!', 'z', 0, 0x80, 0xff} {
			gi, ga := IndexByteASCII(s, c)
			wi, wa := refIndexByteASCII(s, c)
			if gi != wi || ga != wa {
				t.Fatalf("IndexByteASCII(%q, %d): got (%d,%v), want (%d,%v)", s, c, gi, ga, wi, wa)
			}
		}
	}
	rng := rand.New(rand.NewSource(999))
	validRunes := []rune{'a', 'z', 0x7F, 0x80, 0x7FF, 0x800, 0xD7FF, 0xE000, 0xFFFF, 0x10000, 0x1F600, 0x10FFFF}
	for iter := 0; iter < 1500; iter++ {
		var b strings.Builder
		k := rng.Intn(120)
		for j := 0; j < k; j++ {
			b.WriteRune(validRunes[rng.Intn(len(validRunes))])
		}
		s := b.String()
		if got, want := FirstNonASCII(s), refFirstNonASCII(s); got != want {
			t.Fatalf("valid FirstNonASCII: got %d, want %d", got, want)
		}
		if got, want := CountUTF8CodePoints(s), refCountUTF8(s); got != want {
			t.Fatalf("valid CountUTF8CodePoints: got %d, want %d", got, want)
		}
		n := int64(rng.Intn(k + 5))
		goff, gmal := AdvanceUTF8CodePoints(s, n)
		woff, wmal := refAdvanceUTF8(s, n)
		if goff != woff || gmal != wmal {
			t.Fatalf("valid AdvanceUTF8CodePoints(n=%d): got (%d,%v), want (%d,%v)", n, goff, gmal, woff, wmal)
		}
	}
	for iter := 0; iter < 1500; iter++ {
		buf := make([]byte, rng.Intn(120))
		rng.Read(buf)
		s := string(buf)
		if got, want := FirstNonASCII(s), refFirstNonASCII(s); got != want {
			t.Fatalf("fuzz FirstNonASCII: got %d, want %d", got, want)
		}
		if got, want := CountUTF8CodePoints(s), refCountUTF8(s); got != want {
			t.Fatalf("fuzz CountUTF8CodePoints: got %d, want %d", got, want)
		}
		n := int64(rng.Intn(len(buf) + 5))
		goff, gmal := AdvanceUTF8CodePoints(s, n)
		woff, wmal := refAdvanceUTF8(s, n)
		if goff != woff || gmal != wmal {
			t.Fatalf("fuzz AdvanceUTF8CodePoints: got (%d,%v), want (%d,%v)", goff, gmal, woff, wmal)
		}
		c := byte(rng.Intn(256))
		gi, ga := IndexByteASCII(s, c)
		wi, wa := refIndexByteASCII(s, c)
		if gi != wi || ga != wa {
			t.Fatalf("fuzz IndexByteASCII: got (%d,%v), want (%d,%v)", gi, ga, wi, wa)
		}
	}
}
`;

test('SIMD (arm64/amd64 Go 1.27+) and generic fallback kernels match scalar oracles, pass vet/race and cross-compile everywhere', t => {
  const dir = mkdtempSync(join(tmpdir(), 'gopurs-simd-kernels-'));
  t.after(() => rmSync(dir, { recursive: true, force: true }));
  mkdirSync(join(dir, 'gopurs_runtime'));
  writeFileSync(join(dir, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  for (const name of runtimeFiles) {
    copyFileSync(new URL(name, runtimeDir), join(dir, 'gopurs_runtime', name));
  }
  writeFileSync(join(dir, 'gopurs_runtime/kernels_test.go'), kernelTestSource);

  for (const [label, args, envExtra] of [
    ['go vet', ['vet', './gopurs_runtime'], {}],
    ['go test -race (native)', ['test', '-race', '-count=1', './gopurs_runtime'], {}],
    ['go test (amd64 via Rosetta)', ['test', '-count=1', './gopurs_runtime'], { GOARCH: 'amd64' }],
    ['cross linux/arm64', ['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'arm64' }],
    ['cross linux/amd64', ['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'amd64' }],
    ['cross linux/riscv64', ['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'riscv64' }],
    ['cross js/wasm', ['build', './gopurs_runtime'], { GOOS: 'js', GOARCH: 'wasm' }],
  ]) {
    const res = spawnSync('go', args, {
      cwd: dir,
      encoding: 'utf8',
      timeout: 60000,
      env: { ...process.env, GOWORK: 'off', ...envExtra },
    });
    assert.equal(res.status, 0, `${label}:\n${res.stdout}\n${res.stderr}`);
  }

  // Also test the pure-Go fallback (simd_generic.go) directly as Go <= 1.26 / non-SIMD targets do.
  const genericDir = mkdtempSync(join(tmpdir(), 'gopurs-simd-generic-'));
  t.after(() => rmSync(genericDir, { recursive: true, force: true }));
  mkdirSync(join(genericDir, 'gopurs_runtime'));
  writeFileSync(join(genericDir, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  copyFileSync(new URL('runtime.go', runtimeDir), join(genericDir, 'gopurs_runtime/runtime.go'));
  const genericSrc = readFileSync(new URL('simd_generic.go', runtimeDir), 'utf8')
    .replace(/^\/\/go:build .*$/m, '');
  writeFileSync(join(genericDir, 'gopurs_runtime/simd_generic.go'), genericSrc);
  writeFileSync(join(genericDir, 'gopurs_runtime/kernels_test.go'), kernelTestSource);

  const genericRes = spawnSync('go', ['test', '-race', '-count=1', './gopurs_runtime'], {
    cwd: genericDir,
    encoding: 'utf8',
    timeout: 60000,
    env: { ...process.env, GOWORK: 'off' },
  });
  assert.equal(genericRes.status, 0, `pure-Go fallback:\n${genericRes.stdout}\n${genericRes.stderr}`);
});
