//go:build arm64

#include "textflag.h"

// func sumEvenOddInt32Asm(ptr *int32, n int) (even int32, odd int32)
// Accumulates wrapping int32 sums of even-indexed and odd-indexed elements
// using 128-bit NEON vector additions (16 int32s per loop iteration).
TEXT ·sumEvenOddInt32Asm(SB), NOSPLIT, $0-24
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1

	VEOR	V0.B16, V0.B16, V0.B16
	VEOR	V1.B16, V1.B16, V1.B16
	VEOR	V2.B16, V2.B16, V2.B16
	VEOR	V3.B16, V3.B16, V3.B16

	CMP	$16, R1
	BLT	tail4_check

loop16:
	VLD1.P	64(R0), [V4.S4, V5.S4, V6.S4, V7.S4]
	VADD	V4.S4, V0.S4, V0.S4
	VADD	V5.S4, V1.S4, V1.S4
	VADD	V6.S4, V2.S4, V2.S4
	VADD	V7.S4, V3.S4, V3.S4
	SUB	$16, R1
	CMP	$16, R1
	BGE	loop16

	VADD	V1.S4, V0.S4, V0.S4
	VADD	V3.S4, V2.S4, V2.S4
	VADD	V2.S4, V0.S4, V0.S4

tail4_check:
	CMP	$4, R1
	BLT	reduce

loop4:
	VLD1.P	16(R0), [V4.S4]
	VADD	V4.S4, V0.S4, V0.S4
	SUB	$4, R1
	CMP	$4, R1
	BGE	loop4

reduce:
	VMOV	V0.S[0], R2
	VMOV	V0.S[1], R3
	VMOV	V0.S[2], R4
	VMOV	V0.S[3], R5
	ADDW	R4, R2, R2
	ADDW	R5, R3, R3

	CBZ	R1, done
	MOVW	(R0), R4
	ADDW	R4, R2, R2
	CMP	$1, R1
	BEQ	done
	MOVW	4(R0), R4
	ADDW	R4, R3, R3
	CMP	$2, R1
	BEQ	done
	MOVW	8(R0), R4
	ADDW	R4, R2, R2

done:
	MOVW	R2, even+16(FP)
	MOVW	R3, odd+20(FP)
	RET
