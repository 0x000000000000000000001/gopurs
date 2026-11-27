//go:build arm64 || amd64

package gopurs_runtime

import "unsafe"

//go:noescape
func sumEvenOddInt32Asm(ptr *int32, n int) (even int32, odd int32)

func SumInt32(xs []int32) int32 {
	if len(xs) == 0 {
		return 0
	}
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
