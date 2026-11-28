//go:build !arm64 && !amd64

package gopurs_runtime

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

func AffineParitySumInt32(xs []int32, aEven, bEven, aOdd, bOdd int32) int32 {
	n := len(xs)
	if n == 0 {
		return 0
	}
	var e0, o0, e1, o1 int32
	i := 0
	for ; i <= n-4; i += 4 {
		c := xs[i : i+4]
		e0 += c[0]
		o0 += c[1]
		e1 += c[2]
		o1 += c[3]
	}
	even := e0 + e1
	odd := o0 + o1
	if i < n {
		even += xs[i]
		i++
	}
	if i < n {
		odd += xs[i]
		i++
	}
	if i < n {
		even += xs[i]
	}
	nEven := int32((uint(n) + 1) >> 1)
	nOdd := int32(uint(n) >> 1)
	return aEven*even + aOdd*odd + bEven*nEven + bOdd*nOdd
}

func FilterMaskAffineParitySumInt32(xs []int32, mask, target, aEven, bEven, aOdd, bOdd int32) (int32, int64) {
	var evenSum, oddSum, evenCnt, oddCnt int32
	parity := 0
	for _, v := range xs {
		if (v & mask) == target {
			if parity == 0 {
				evenSum += v
				evenCnt++
				parity = 1
			} else {
				oddSum += v
				oddCnt++
				parity = 0
			}
		}
	}
	totalCnt := int64(uint32(evenCnt)) + int64(uint32(oddCnt))
	return aEven*evenSum + aOdd*oddSum + bEven*evenCnt + bOdd*oddCnt, totalCnt
}

func IndexByteASCII(s string, c byte) (int, bool) {
	allASCII := true
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == c {
			return i, allASCII
		}
		if b >= 0x80 {
			allASCII = false
		}
	}
	return -1, false
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

func ContainsInt32(xs []int32, target int32) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
	}
	return false
}

func CountUTF8CodePoints(s string) int64 {
	var count int64
	n := len(s)
	i := 0
	for i <= n-4 {
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

func SumCodePoints(s string) (int64, bool) {
	return SumCodePointsLimit(s, int64(len(s)))
}

func SumCodePointsLimit(s string, limit int64) (int64, bool) {
	n := len(s)
	if limit <= 0 || n == 0 {
		return 0, false
	}
	var sum, count int64
	i := 0
	for count < limit && i <= n-4 {
		b0 := s[i]
		if b0 < 0x80 {
			sum += int64(b0)
			i++
		} else if b0 < 0xE0 {
			b1 := s[i+1]
			if b0 >= 0xC2 && b1&0xC0 == 0x80 {
				sum += int64(b0&0x1F)<<6 | int64(b1&0x3F)
				i += 2
			} else {
				sum += 0xFFFD
				i++
			}
		} else if b0 < 0xF0 {
			b1, b2 := s[i+1], s[i+2]
			if b0 != 0xED && ((b1^0x80)|(b2^0x80)) < 0x40 && (b0 != 0xE0 || b1 >= 0xA0) {
				sum += int64(b0&0x0F)<<12 | int64(b1&0x3F)<<6 | int64(b2&0x3F)
				i += 3
			} else {
				r, w := DecodeStringCodePoint(s[i:])
				sum += int64(r)
				i += w
			}
		} else {
			b1, b2, b3 := s[i+1], s[i+2], s[i+3]
			if b0 <= 0xF4 && ((b1^0x80)|(b2^0x80)|(b3^0x80)) < 0x40 && (b0 != 0xF0 || b1 >= 0x90) && (b0 != 0xF4 || b1 <= 0x8F) {
				sum += int64(b0&0x07)<<18 | int64(b1&0x3F)<<12 | int64(b2&0x3F)<<6 | int64(b3&0x3F)
				i += 4
			} else {
				sum += 0xFFFD
				i++
			}
		}
		count++
	}
	for count < limit && i < n {
		r, w := DecodeStringCodePoint(s[i:])
		sum += int64(r)
		i += w
		count++
	}
	return sum, count > 0
}
