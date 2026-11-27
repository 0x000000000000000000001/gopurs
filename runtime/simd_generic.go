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
