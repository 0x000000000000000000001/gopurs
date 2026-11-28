//go:build (arm64 || amd64) && go1.27

package gopurs_runtime

import (
	"unicode/utf8"
	"unsafe"
)

//go:noescape
func sumInt32Asm(ptr *int32, n int) int32

//go:noescape
func isNonIncreasingInt32Asm(ptr *int32, n int) bool

//go:noescape
func isNonDecreasingInt32Asm(ptr *int32, n int) bool

//go:noescape
func indexOfInt32Asm(ptr *int32, n int, target int32) int

//go:noescape
func countUTF8Asm(ptr *byte, n int) (count int64, consumed int)

//go:noescape
func indexByteASCIIAsm(ptr *byte, n int, c byte) (idx int, prefixASCII bool)

//go:noescape
func firstNonASCIIAsm(ptr *byte, n int) int

func SumInt32(xs []int32) int32 {
	if len(xs) == 0 {
		return 0
	}
	return sumInt32Asm(unsafe.SliceData(xs), len(xs))
}

func IsNonIncreasingInt32(xs []int32) bool {
	return isNonIncreasingInt32Asm(unsafe.SliceData(xs), len(xs))
}

func IsNonDecreasingInt32(xs []int32) bool {
	return isNonDecreasingInt32Asm(unsafe.SliceData(xs), len(xs))
}

func IndexOfInt32(xs []int32, target int32) int {
	if len(xs) == 0 {
		return -1
	}
	if xs[0] == target {
		return 0
	}
	return indexOfInt32Asm(unsafe.SliceData(xs), len(xs), target)
}

func ContainsInt32(xs []int32, target int32) bool {
	return IndexOfInt32(xs, target) >= 0
}

// FirstNonASCII returns the byte offset of the first byte >= 0x80 in s,
// or len(s) if every byte is ASCII.
func FirstNonASCII(s string) int {
	if len(s) == 0 {
		return 0
	}
	return firstNonASCIIAsm(unsafe.StringData(s), len(s))
}

func IsASCII(s string) bool {
	return FirstNonASCII(s) == len(s)
}

// IndexByteASCII locates the first occurrence of c in s and reports whether
// every byte strictly before that match is ASCII (< 0x80).
func IndexByteASCII(s string, c byte) (int, bool) {
	if len(s) == 0 {
		return -1, false
	}
	return indexByteASCIIAsm(unsafe.StringData(s), len(s), c)
}

func CountUTF8CodePoints(s string) int64 {
	n := len(s)
	if n == 0 {
		return 0
	}
	count, i := countUTF8Asm(unsafe.StringData(s), n)
	if i > 0 {
		if s[i-1] >= 0xC0 {
			count--
			i -= 1
		} else if i >= 2 && s[i-2] >= 0xE0 {
			count--
			i -= 2
		} else if i >= 3 && s[i-3] >= 0xF0 {
			count--
			i -= 3
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
	total := len(s)
	if n <= 0 || total == 0 {
		return 0, false
	}
	off := firstNonASCIIAsm(unsafe.StringData(s), total)
	if int64(off) >= n {
		return int(n), false
	}
	count := int64(off)
	base := unsafe.Pointer(unsafe.StringData(s))
	// In any strict RFC 3629 UTF-8 span, k bytes contain at most k code points.
	// Passing min(remainingBytes, remainingCodePoints) to countUTF8Asm therefore
	// never overshoots n, scans each byte once at SIMD speed, and guarantees
	// malformed == false for everything it consumes.
	for {
		remCP := n - count
		remBytes := total - off
		if remCP < 16 || remBytes < 16 {
			break
		}
		budget := remBytes
		if int64(budget) > remCP {
			budget = int(remCP)
		}
		c, consumed := countUTF8Asm((*byte)(unsafe.Add(base, off)), budget)
		if consumed == 0 {
			break
		}
		end := off + consumed
		if s[end-1] >= 0xC0 {
			c--
			consumed -= 1
		} else if consumed >= 2 && s[end-2] >= 0xE0 {
			c--
			consumed -= 2
		} else if consumed >= 3 && s[end-3] >= 0xF0 {
			c--
			consumed -= 3
		}
		if consumed <= 0 {
			break
		}
		off += consumed
		count += c
	}
	malformed := false
	for count < n && off < total {
		r, w := DecodeStringCodePoint(s[off:])
		if r == utf8.RuneError && w == 1 {
			malformed = true
		}
		off += w
		count++
	}
	return off, malformed
}
