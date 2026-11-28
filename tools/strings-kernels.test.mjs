import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';

const runtimeDir = new URL('../runtime/', import.meta.url);
const stringsDir = new URL('../../gopurs-strings/src/Data/String/', import.meta.url);

test('SWAR CodePoints.Length and adaptive CodeUnits.indexOf match exact scalar references across UTF-8, WTF-8, surrogates and malformed bytes', t => {
  const dir = mkdtempSync(join(tmpdir(), 'gopurs-strings-oracle-'));
  t.after(() => rmSync(dir, { recursive: true, force: true }));
  mkdirSync(join(dir, 'gopurs_runtime'));
  mkdirSync(join(dir, 'codepoints'));
  mkdirSync(join(dir, 'codeunits'));
  writeFileSync(join(dir, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  for (const name of ['runtime.go', 'simd_generic.go']) {
    copyFileSync(new URL(name, runtimeDir), join(dir, 'gopurs_runtime', name));
  }
  copyFileSync(new URL('CodePoints.go', stringsDir), join(dir, 'codepoints/CodePoints.go'));
  writeFileSync(
    join(dir, 'codeunits/CodeUnits.go'),
    'package codeunits\n' + readFileSync(new URL('CodeUnits.go', stringsDir), 'utf8'),
  );
  writeFileSync(
    join(dir, 'codepoints/oracle_test.go'),
    `package CodePoints

import (
	"math/rand"
	"strings"
	"testing"
)

func refLength(s string) int64 {
	var n int64
	for i := 0; i < len(s); {
		_, w := decodeCodePoint(s[i:])
		n++
		i += w
	}
	return n
}

func TestSWARLengthOracle(t *testing.T) {
	corpus := []string{
		"", "a", "ab", "abcdefgh", "abcdefghi",
		"é", "a😀é", strings.Repeat("a😀é", 1000),
		strings.Repeat("日本語テスト", 200),
		strings.Repeat("😀", 100),
		"\\xed\\xa0\\x80", "\\xed\\xb0\\x80", "\\xed\\xa0\\xbd\\xed\\xb8\\x80",
		"\\xed\\x9f\\xbf", // U+D7FF (valid 0xED < 0xA0)
		"\\xe0\\x9f\\x80", // overlong 3-byte
		"\\xe0\\xa0\\x80", // valid U+0800
		"\\xf0\\x8f\\x80\\x80", // overlong 4-byte
		"\\xf0\\x90\\x80\\x80", // valid U+10000
		"\\xf4\\x8f\\xbf\\xbf", // valid U+10FFFF
		"\\xf4\\x90\\x80\\x80", // > U+10FFFF
		"\\xf5\\x80\\x80\\x80", "\\xf8\\x80\\x80\\x80", "\\xff",
		"\\xc0\\x80", "\\xc1\\xbf", "\\xc2\\x80",
		"\\x80", "\\xbf", "aaaaaaa\\x80", "aaaaaa\\xc2", "aaaaa\\xe2\\x82", "aaaa\\xf0\\x9f\\x98",
	}
	for _, s := range corpus {
		for pad := 0; pad <= 9; pad++ {
			p := strings.Repeat("x", pad) + s + strings.Repeat("y", pad)
			if got, want := Length(p), refLength(p); got != want {
				t.Fatalf("Length(%q): got %d, want %d", p, got, want)
			}
		}
	}
	// Test every single split boundary for 2/3/4-byte sequences crossing an 8-byte boundary
	tokens := []string{"a", "é", "€", "😀", "\\xed\\xa0\\x80", "\\xed\\xb0\\x80", "\\xf0\\x8f\\x80\\x80", "\\xf0\\x90\\x80\\x80", "\\xe0\\xa0\\x80", "\\xe0\\x9f\\x80", "\\xf4\\x8f\\xbf\\xbf", "\\xf4\\x90\\x80\\x80", "\\xc0\\x80", "\\xc1\\x80", "\\x80", "\\xff"}
	rng := rand.New(rand.NewSource(999))
	for iter := 0; iter < 5000; iter++ {
		var b strings.Builder
		k := rng.Intn(25)
		for j := 0; j < k; j++ {
			b.WriteString(tokens[rng.Intn(len(tokens))])
		}
		s := b.String()
		if got, want := Length(s), refLength(s); got != want {
			t.Fatalf("token fuzz %d (%q): got %d, want %d", iter, s, got, want)
		}
	}
	for iter := 0; iter < 5000; iter++ {
		buf := make([]byte, rng.Intn(40))
		rng.Read(buf)
		s := string(buf)
		if got, want := Length(s), refLength(s); got != want {
			t.Fatalf("random bytes %d (%x): got %d, want %d", iter, buf, got, want)
		}
	}
	// Valid random UTF-8 including all boundary code points
	validRunes := []rune{'a', 'z', 0x7F, 0x80, 0x7FF, 0x800, 0xD7FF, 0xE000, 0xFFFF, 0x10000, 0x1F600, 0x10FFFF}
	for iter := 0; iter < 2000; iter++ {
		var b strings.Builder
		k := rng.Intn(40)
		for j := 0; j < k; j++ {
			b.WriteRune(validRunes[rng.Intn(len(validRunes))])
		}
		s := b.String()
		if got, want := Length(s), refLength(s); got != want {
			t.Fatalf("valid utf8 %d: got %d, want %d", iter, got, want)
		}
	}
}
`,
  );
  writeFileSync(
    join(dir, 'codeunits/oracle_test.go'),
    `package codeunits

import (
	"math/rand"
	"strings"
	"testing"
)

func refIndexOf(x, s string, start int) int {
	input, pattern := stringUnits(s), stringUnits(x)
	if start < 0 || start > len(input) || start+len(pattern) > len(input) {
		return -1
	}
	return findUnits(pattern, input, start, false)
}

func TestAdaptiveIndexOfOracle(t *testing.T) {
	just := func(i int) interface{} { return i }
	nothing := -1
	cases := [][2]string{
		{"absent", strings.Repeat("a", 10000) + "needle"},
		{"needle", strings.Repeat("a", 10000) + "needle"},
		{"aaaaaa", strings.Repeat("a", 10000) + "needle"},
		{"aaaaab", strings.Repeat("a", 200) + "b"},
		{"ababac", strings.Repeat("ab", 500) + "ac"},
		{"ababaz", strings.Repeat("ab", 500) + "ac"},
		{"", "hello"},
		{"x", ""},
		{"é", strings.Repeat("a😀é", 100)},
		{"😀", strings.Repeat("a😀é", 100)},
		{"\\xed\\xa0\\xbd", strings.Repeat("a😀é", 10)},
		{"\\xed\\xb8\\x80", strings.Repeat("a😀é", 10)},
		{"\\xed\\x9f\\xbf", strings.Repeat("a\\xed\\x9f\\xbf", 50)},
		{"\\xff", strings.Repeat("a", 100) + "\\xff"},
	}
	for _, tc := range cases {
		for _, start := range []int{-1, 0, 1, 2, 5, 100, 100000} {
			got := _IndexOfStartingAt(just, nothing, tc[0], start, tc[1]).(int)
			want := refIndexOf(tc[0], tc[1], start)
			if got != want {
				t.Fatalf("indexOf(%q, len=%d, start=%d): got %d, want %d", tc[0], len(tc[1]), start, got, want)
			}
		}
	}
	rng := rand.New(rand.NewSource(2026))
	alphabet := []string{"a", "b", "c", "é", "😀"}
	for iter := 0; iter < 500; iter++ {
		var sb, pb strings.Builder
		slen := rng.Intn(200)
		plen := rng.Intn(12)
		for i := 0; i < slen; i++ {
			sb.WriteString(alphabet[rng.Intn(len(alphabet))])
		}
		for i := 0; i < plen; i++ {
			pb.WriteString(alphabet[rng.Intn(len(alphabet))])
		}
		s, p := sb.String(), pb.String()
		for _, start := range []int{0, 1, 7} {
			got := _IndexOfStartingAt(just, nothing, p, start, s).(int)
			want := refIndexOf(p, s, start)
			if got != want {
				t.Fatalf("fuzz %d: p=%q start=%d got %d want %d", iter, p, start, got, want)
			}
		}
	}
}
`,
  );

  const res = spawnSync('go', ['test', '-race', '-count=1', './codepoints', './codeunits'], {
    cwd: dir,
    encoding: 'utf8',
    timeout: 60_000,
    env: { ...process.env, GOWORK: 'off' },
  });
  assert.ifError(res.error);
  assert.equal(res.status, 0, res.stdout + res.stderr);
});
