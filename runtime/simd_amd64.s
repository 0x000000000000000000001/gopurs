//go:build amd64

#include "textflag.h"

// func sumEvenOddInt32Asm(ptr *int32, n int) (even int32, odd int32)
// Accumulates wrapping int32 sums of even-indexed and odd-indexed elements
// using baseline SSE2 vector additions (16 int32s per loop iteration).
TEXT ·sumEvenOddInt32Asm(SB), NOSPLIT, $0-24
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX

	PXOR	X0, X0
	PXOR	X1, X1
	PXOR	X2, X2
	PXOR	X3, X3

	CMPQ	CX, $16
	JL	tail4_check

loop16:
	MOVOU	0(SI), X4
	MOVOU	16(SI), X5
	MOVOU	32(SI), X6
	MOVOU	48(SI), X7
	PADDD	X4, X0
	PADDD	X5, X1
	PADDD	X6, X2
	PADDD	X7, X3
	ADDQ	$64, SI
	SUBQ	$16, CX
	CMPQ	CX, $16
	JGE	loop16

	PADDD	X1, X0
	PADDD	X3, X2
	PADDD	X2, X0

tail4_check:
	CMPQ	CX, $4
	JL	reduce

loop4:
	MOVOU	0(SI), X4
	PADDD	X4, X0
	ADDQ	$16, SI
	SUBQ	$4, CX
	CMPQ	CX, $4
	JGE	loop4

reduce:
	PSHUFD	$0xEE, X0, X1
	PADDD	X1, X0
	MOVQ	X0, AX
	MOVL	AX, DX
	SHRQ	$32, AX

	TESTQ	CX, CX
	JZ	done
	ADDL	0(SI), DX
	CMPQ	CX, $1
	JE	done
	ADDL	4(SI), AX
	CMPQ	CX, $2
	JE	done
	ADDL	8(SI), DX

done:
	MOVL	DX, even+16(FP)
	MOVL	AX, odd+20(FP)
	RET
