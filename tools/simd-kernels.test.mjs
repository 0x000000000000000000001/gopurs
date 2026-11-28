import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';

const runtimeDir = new URL('../runtime/', import.meta.url);
const runtimeFiles = ['runtime.go', 'simd_generic.go'];

test('Portable pure-Go runtime kernels match scalar oracles, pass vet/race and cross-compile everywhere', t => {
  const dir = mkdtempSync(join(tmpdir(), 'gopurs-runtime-kernels-'));
  t.after(() => rmSync(dir, { recursive: true, force: true }));
  mkdirSync(join(dir, 'gopurs_runtime'));
  writeFileSync(join(dir, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  for (const name of runtimeFiles) {
    copyFileSync(new URL(name, runtimeDir), join(dir, 'gopurs_runtime', name));
  }
  writeFileSync(
    join(dir, 'gopurs_runtime/kernels_test.go'),
    `package gopurs_runtime

import (
	"math/rand"
	"strings"
	"testing"
)

func refSum(xs []int32) int32 {
	var s int32
	for _, x := range xs {
		s += x
	}
	return s
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
	for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 15, 16, 17, 31, 32, 33, 127, 128, 1000} {
		xs := make([]int32, n)
		for i := range xs {
			xs[i] = int32(rng.Uint32())
		}
		if got, want := SumInt32(xs), refSum(xs); got != want {
			t.Fatalf("SumInt32 n=%d: got %d, want %d", n, got, want)
		}
	}
}

func TestMonotoneAndContainsInt32(t *testing.T) {
	for _, n := range []int{0, 1, 2, 3, 4, 5, 19, 20, 21, 64, 100, 1000} {
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
		if n > 0 {
			if !ContainsInt32(asc, asc[0]) || !ContainsInt32(asc, asc[n-1]) {
				t.Fatalf("ContainsInt32 missed element n=%d", n)
			}
			if ContainsInt32(asc, int32(n+100)) {
				t.Fatalf("ContainsInt32 false positive n=%d", n)
			}
		}
	}
}

func TestUTF8AndIndexByteASCII(t *testing.T) {
	corpus := []string{
		"", "a", "abcdefgh", "abcdefghi", "é", "a😀é",
		strings.Repeat("a😀é", 200),
		strings.Repeat("hello world! ", 200) + "🎯",
		"\\xed\\xa0\\x80", "\\xed\\xb0\\x80", "\\xe0\\x9f\\x80", "\\xf0\\x8f\\x80\\x80",
		"\\xf4\\x90\\x80\\x80", "\\xc0\\x80", "\\x80", "\\xff",
	}
	for _, s := range corpus {
		if got, want := CountUTF8CodePoints(s), refCountUTF8(s); got != want {
			t.Fatalf("CountUTF8CodePoints(%q): got %d, want %d", s, got, want)
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
	for iter := 0; iter < 2000; iter++ {
		buf := make([]byte, rng.Intn(80))
		rng.Read(buf)
		s := string(buf)
		if got, want := CountUTF8CodePoints(s), refCountUTF8(s); got != want {
			t.Fatalf("fuzz CountUTF8CodePoints: got %d, want %d", got, want)
		}
		c := byte(rng.Intn(256))
		gi, ga := IndexByteASCII(s, c)
		wi, wa := refIndexByteASCII(s, c)
		if gi != wi || ga != wa {
			t.Fatalf("fuzz IndexByteASCII: got (%d,%v), want (%d,%v)", gi, ga, wi, wa)
		}
	}
}
`,
  );

  for (const [label, args, envExtra] of [
    ['go vet', ['vet', './gopurs_runtime'], {}],
    ['go test -race', ['test', '-race', '-count=1', './gopurs_runtime'], {}],
    ['cross linux/arm64', ['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'arm64' }],
    ['cross linux/amd64', ['build', './gopurs_runtime'], { GOOS: 'linux', GOARCH: 'amd64' }],
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
});
