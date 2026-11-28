package gopurs_runtime

import "strings"

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

// IndexByteASCII locates the first occurrence of c in s and reports whether
// every byte before that match is ASCII (< 0x80).
func IndexByteASCII(s string, c byte) (int, bool) {
	idx := strings.IndexByte(s, c)
	if idx < 0 {
		return -1, false
	}
	prefix := s[:idx]
	i := 0
	for ; i <= len(prefix)-8; i += 8 {
		if (prefix[i] | prefix[i+1] | prefix[i+2] | prefix[i+3] |
			prefix[i+4] | prefix[i+5] | prefix[i+6] | prefix[i+7]) >= 0x80 {
			return idx, false
		}
	}
	for ; i < len(prefix); i++ {
		if prefix[i] >= 0x80 {
			return idx, false
		}
	}
	return idx, true
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
	i := 0
	for ; i <= len(xs)-4; i += 4 {
		c := xs[i : i+4]
		if c[0] == target || c[1] == target || c[2] == target || c[3] == target {
			return true
		}
	}
	for ; i < len(xs); i++ {
		if xs[i] == target {
			return true
		}
	}
	return false
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
