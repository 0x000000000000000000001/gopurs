//go:build arm64 || amd64

package gopurs_runtime

import "unsafe"

//go:noescape
func sumEvenOddInt32Asm(ptr *int32, n int) (even int32, odd int32)

//go:noescape
func isNonIncreasingInt32Asm(ptr *int32, n int) bool

//go:noescape
func isNonDecreasingInt32Asm(ptr *int32, n int) bool

//go:noescape
func containsInt32Asm(ptr *int32, n int, target int32) bool

//go:noescape
func countUTF8Asm(ptr *byte, n int) (count int64, consumed int)

//go:noescape
func sumCodePointsLimitAsm(ptr *byte, n int, limit int64) (sum int64, consumed int, count int64)

//go:noescape
func filterMaskEvenOddInt32Asm(ptr *int32, n int, mask, target int32) (evenSum, oddSum, evenCnt, oddCnt int32)

//go:noescape
func indexByteASCIIAsm(ptr *byte, n int, c byte) (idx int, prefixASCII bool)

func IndexByteASCII(s string, c byte) (int, bool) {
	if len(s) == 0 {
		return -1, false
	}
	return indexByteASCIIAsm(unsafe.StringData(s), len(s), c)
}

func SumInt32(xs []int32) int32 {
	even, odd := sumEvenOddInt32Asm(unsafe.SliceData(xs), len(xs))
	return even + odd
}

func AffineParitySumInt32(xs []int32, aEven, bEven, aOdd, bOdd int32) int32 {
	n := len(xs)
	if n == 0 {
		return 0
	}
	even, odd := sumEvenOddInt32Asm(unsafe.SliceData(xs), n)
	nEven := int32((uint(n) + 1) >> 1)
	nOdd := int32(uint(n) >> 1)
	return aEven*even + aOdd*odd + bEven*nEven + bOdd*nOdd
}

func FilterMaskAffineParitySumInt32(xs []int32, mask, target, aEven, bEven, aOdd, bOdd int32) (int32, int64) {
	n := len(xs)
	if n == 0 {
		return 0, 0
	}
	evenSum, oddSum, evenCnt, oddCnt := filterMaskEvenOddInt32Asm(unsafe.SliceData(xs), n, mask, target)
	totalCnt := int64(uint32(evenCnt)) + int64(uint32(oddCnt))
	return aEven*evenSum + aOdd*oddSum + bEven*evenCnt + bOdd*oddCnt, totalCnt
}

func IsNonIncreasingInt32(xs []int32) bool {
	return isNonIncreasingInt32Asm(unsafe.SliceData(xs), len(xs))
}

func IsNonDecreasingInt32(xs []int32) bool {
	return isNonDecreasingInt32Asm(unsafe.SliceData(xs), len(xs))
}

func ContainsInt32(xs []int32, target int32) bool {
	if len(xs) == 0 {
		return false
	}
	if xs[0] == target {
		return true
	}
	return containsInt32Asm(unsafe.SliceData(xs), len(xs), target)
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

func SumCodePoints(s string) (int64, bool) {
	return SumCodePointsLimit(s, int64(len(s)))
}

func SumCodePointsLimit(s string, limit int64) (int64, bool) {
	n := len(s)
	if limit <= 0 || n == 0 {
		return 0, false
	}
	sum, i, count := sumCodePointsLimitAsm(unsafe.StringData(s), n, limit)
	if i > 0 {
		if b := s[i-1]; b >= 0xC0 {
			if b < 0xE0 {
				sum -= int64(b&0x1F) << 6
			} else if b < 0xF0 {
				sum -= int64(b&0x0F) << 12
			} else {
				sum -= int64(b&0x07) << 18
			}
			count--
			i -= 1
		} else if i >= 2 && s[i-2] >= 0xE0 {
			b0, b1 := s[i-2], s[i-1]
			if b0 < 0xF0 {
				sum -= (int64(b0&0x0F) << 12) + (int64(b1&0x3F) << 6)
			} else {
				sum -= (int64(b0&0x07) << 18) + (int64(b1&0x3F) << 12)
			}
			count--
			i -= 2
		} else if i >= 3 && s[i-3] >= 0xF0 {
			b0, b1, b2 := s[i-3], s[i-2], s[i-1]
			sum -= (int64(b0&0x07) << 18) + (int64(b1&0x3F) << 12) + (int64(b2&0x3F) << 6)
			count--
			i -= 3
		}
	}
	for count < limit && i < n {
		r, w := DecodeStringCodePoint(s[i:])
		sum += int64(r)
		i += w
		count++
	}
	return sum, count > 0
}
