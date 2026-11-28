//go:build !(arm64 || amd64) || !go1.27

package gopurs_runtime

import (
	"strings"
	"unicode/utf8"
)

func SumInt32(xs []int32) int32 {
	var s0, s1, s2, s3 int32
	i := 0
	for ; i <= len(xs)-4; i += 4 {
		c := xs[i : i+4]
		s0 += c[0]
		s1 += c[1]
		s2 += c[2]
		s3 += c[3]
	}
	sum := (s0 + s1) + (s2 + s3)
	for ; i < len(xs); i++ {
		sum += xs[i]
	}
	return sum
}

func IsNonIncreasingInt32(xs []int32) bool {
	for i := 1; i < len(xs); i++ {
		if xs[i] > xs[i-1] {
			return false
		}
	}
	return true
}

func IsNonDecreasingInt32(xs []int32) bool {
	for i := 1; i < len(xs); i++ {
		if xs[i] < xs[i-1] {
			return false
		}
	}
	return true
}

func IndexOfInt32(xs []int32, target int32) int {
	for i, x := range xs {
		if x == target {
			return i
		}
	}
	return -1
}

func ContainsInt32(xs []int32, target int32) bool {
	return IndexOfInt32(xs, target) >= 0
}

// FirstNonASCII returns the byte offset of the first byte >= 0x80 in s,
// or len(s) if every byte is ASCII.
func FirstNonASCII(s string) int {
	i := 0
	for ; i <= len(s)-8; i += 8 {
		if (s[i] | s[i+1] | s[i+2] | s[i+3] |
			s[i+4] | s[i+5] | s[i+6] | s[i+7]) >= 0x80 {
			for j := i; j < i+8; j++ {
				if s[j] >= 0x80 {
					return j
				}
			}
		}
	}
	for ; i < len(s); i++ {
		if s[i] >= 0x80 {
			return i
		}
	}
	return len(s)
}

func IsASCII(s string) bool {
	return FirstNonASCII(s) == len(s)
}

// IndexByteASCII locates the first occurrence of c in s and reports whether
// every byte strictly before that match is ASCII (< 0x80).
func IndexByteASCII(s string, c byte) (int, bool) {
	idx := strings.IndexByte(s, c)
	if idx < 0 {
		return -1, false
	}
	return idx, FirstNonASCII(s[:idx]) == idx
}

func CountUTF8CodePoints(s string) int64 {
	var count int64
	n := len(s)
	i := 0
	for i <= n-8 {
		if (s[i] | s[i+1] | s[i+2] | s[i+3] |
			s[i+4] | s[i+5] | s[i+6] | s[i+7]) < 0x80 {
			count += 8
			i += 8
			continue
		}
		b := s[i]
		count++
		if b < 0x80 {
			i++
		} else if b < 0xE0 {
			if b >= 0xC2 && s[i+1]&0xC0 == 0x80 {
				i += 2
			} else {
				i++
			}
		} else if b < 0xF0 {
			if b != 0xED && ((s[i+1]^0x80)|(s[i+2]^0x80)) < 0x40 && (b != 0xE0 || s[i+1] >= 0xA0) {
				i += 3
			} else {
				_, w := DecodeStringCodePoint(s[i:])
				i += w
			}
		} else if b <= 0xF4 && ((s[i+1]^0x80)|(s[i+2]^0x80)|(s[i+3]^0x80)) < 0x40 && (b != 0xF0 || s[i+1] >= 0x90) && (b != 0xF4 || s[i+1] <= 0x8F) {
			i += 4
		} else {
			i++
		}
	}
	for i < n {
		_, w := DecodeStringCodePoint(s[i:])
		count++
		i += w
	}
	return count
}

// AdvanceUTF8CodePoints returns the byte offset after n code points (clamped to
// [0, len(s)]) using exact UTF-8/WTF-8 semantics, and reports whether any
// malformed byte sequence was encountered inside that prefix.
func AdvanceUTF8CodePoints(s string, n int64) (int, bool) {
	if n <= 0 || len(s) == 0 {
		return 0, false
	}
	off := FirstNonASCII(s)
	if int64(off) >= n {
		return int(n), false
	}
	count := int64(off)
	malformed := false
	for count < n && off < len(s) {
		r, w := DecodeStringCodePoint(s[off:])
		if r == utf8.RuneError && w == 1 {
			malformed = true
		}
		off += w
		count++
	}
	return off, malformed
}
