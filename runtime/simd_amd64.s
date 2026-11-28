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

// func isNonIncreasingInt32Asm(ptr *int32, n int) bool
// Returns true iff xs[i] >= xs[i+1] for all 0 <= i < n-1.
TEXT ·isNonIncreasingInt32Asm(SB), NOSPLIT, $0-17
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX

	CMPQ	CX, $2
	JL	ni_true

	MOVL	0(SI), AX
	MOVL	4(SI), DX
	CMPL	DX, AX
	JG	ni_false

	CMPQ	CX, $17
	JL	ni_scalar_from_1

ni_loop16:
	MOVOU	0(SI), X0
	MOVOU	4(SI), X1
	MOVOU	16(SI), X2
	MOVOU	20(SI), X3
	MOVOU	32(SI), X4
	MOVOU	36(SI), X5
	MOVOU	48(SI), X6
	MOVOU	52(SI), X7
	PCMPGTL	X0, X1
	PCMPGTL	X2, X3
	PCMPGTL	X4, X5
	PCMPGTL	X6, X7
	POR	X3, X1
	POR	X7, X5
	POR	X5, X1
	PMOVMSKB X1, AX
	TESTL	AX, AX
	JNZ	ni_false
	ADDQ	$64, SI
	SUBQ	$16, CX
	CMPQ	CX, $17
	JGE	ni_loop16

	MOVL	0(SI), AX
	ADDQ	$4, SI
	DECQ	CX
	JMP	ni_scalar_loop

ni_scalar_from_1:
	MOVL	DX, AX
	ADDQ	$8, SI
	SUBQ	$2, CX

ni_scalar_loop:
	TESTQ	CX, CX
	JLE	ni_true
	MOVL	0(SI), DX
	CMPL	DX, AX
	JG	ni_false
	MOVL	DX, AX
	ADDQ	$4, SI
	DECQ	CX
	JMP	ni_scalar_loop

ni_true:
	MOVB	$1, ret+16(FP)
	RET

ni_false:
	MOVB	$0, ret+16(FP)
	RET

// func isNonDecreasingInt32Asm(ptr *int32, n int) bool
// Returns true iff xs[i] <= xs[i+1] for all 0 <= i < n-1.
TEXT ·isNonDecreasingInt32Asm(SB), NOSPLIT, $0-17
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX

	CMPQ	CX, $2
	JL	nd_true

	MOVL	0(SI), AX
	MOVL	4(SI), DX
	CMPL	DX, AX
	JL	nd_false

	CMPQ	CX, $17
	JL	nd_scalar_from_1

nd_loop16:
	MOVOU	0(SI), X0
	MOVOU	4(SI), X1
	MOVOU	16(SI), X2
	MOVOU	20(SI), X3
	MOVOU	32(SI), X4
	MOVOU	36(SI), X5
	MOVOU	48(SI), X6
	MOVOU	52(SI), X7
	PCMPGTL	X1, X0
	PCMPGTL	X3, X2
	PCMPGTL	X5, X4
	PCMPGTL	X7, X6
	POR	X2, X0
	POR	X6, X4
	POR	X4, X0
	PMOVMSKB X0, AX
	TESTL	AX, AX
	JNZ	nd_false
	ADDQ	$64, SI
	SUBQ	$16, CX
	CMPQ	CX, $17
	JGE	nd_loop16

	MOVL	0(SI), AX
	ADDQ	$4, SI
	DECQ	CX
	JMP	nd_scalar_loop

nd_scalar_from_1:
	MOVL	DX, AX
	ADDQ	$8, SI
	SUBQ	$2, CX

nd_scalar_loop:
	TESTQ	CX, CX
	JLE	nd_true
	MOVL	0(SI), DX
	CMPL	DX, AX
	JL	nd_false
	MOVL	DX, AX
	ADDQ	$4, SI
	DECQ	CX
	JMP	nd_scalar_loop

nd_true:
	MOVB	$1, ret+16(FP)
	RET

nd_false:
	MOVB	$0, ret+16(FP)
	RET

// func containsInt32Asm(ptr *int32, n int, target int32) bool
// Returns true iff target occurs in ptr[0:n].
TEXT ·containsInt32Asm(SB), NOSPLIT, $0-25
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX
	MOVL	target+16(FP), AX

	MOVD	AX, X0
	PSHUFD	$0x00, X0, X0

	CMPQ	CX, $16
	JL	c_tail4_check

c_loop16:
	MOVOU	0(SI), X1
	MOVOU	16(SI), X2
	MOVOU	32(SI), X3
	MOVOU	48(SI), X4
	PCMPEQL	X0, X1
	PCMPEQL	X0, X2
	PCMPEQL	X0, X3
	PCMPEQL	X0, X4
	POR	X2, X1
	POR	X4, X3
	POR	X3, X1
	PMOVMSKB X1, DX
	TESTL	DX, DX
	JNZ	c_true
	ADDQ	$64, SI
	SUBQ	$16, CX
	CMPQ	CX, $16
	JGE	c_loop16

c_tail4_check:
	CMPQ	CX, $4
	JL	c_tail1

c_loop4:
	MOVOU	0(SI), X1
	PCMPEQL	X0, X1
	PMOVMSKB X1, DX
	TESTL	DX, DX
	JNZ	c_true
	ADDQ	$16, SI
	SUBQ	$4, CX
	CMPQ	CX, $4
	JGE	c_loop4

c_tail1:
	TESTQ	CX, CX
	JLE	c_false
	CMPL	0(SI), AX
	JE	c_true
	ADDQ	$4, SI
	DECQ	CX
	JMP	c_tail1

c_true:
	MOVB	$1, ret+24(FP)
	RET

c_false:
	MOVB	$0, ret+24(FP)
	RET

// Baseline SSE2 constants for 16-byte UTF-8 validation & code-point summation.
DATA u8_80<>+0(SB)/8, $0x8080808080808080
DATA u8_80<>+8(SB)/8, $0x8080808080808080
GLOBL u8_80<>(SB), RODATA, $16

DATA u8_c0<>+0(SB)/8, $0xc0c0c0c0c0c0c0c0
DATA u8_c0<>+8(SB)/8, $0xc0c0c0c0c0c0c0c0
GLOBL u8_c0<>(SB), RODATA, $16

DATA u8_3f<>+0(SB)/8, $0x3f3f3f3f3f3f3f3f
DATA u8_3f<>+8(SB)/8, $0x3f3f3f3f3f3f3f3f
GLOBL u8_3f<>(SB), RODATA, $16

DATA u8_5f<>+0(SB)/8, $0x5f5f5f5f5f5f5f5f
DATA u8_5f<>+8(SB)/8, $0x5f5f5f5f5f5f5f5f
GLOBL u8_5f<>(SB), RODATA, $16

DATA u8_6f<>+0(SB)/8, $0x6f6f6f6f6f6f6f6f
DATA u8_6f<>+8(SB)/8, $0x6f6f6f6f6f6f6f6f
GLOBL u8_6f<>(SB), RODATA, $16

DATA u8_74<>+0(SB)/8, $0x7474747474747474
DATA u8_74<>+8(SB)/8, $0x7474747474747474
GLOBL u8_74<>(SB), RODATA, $16

DATA u8_fe<>+0(SB)/8, $0xfefefefefefefefe
DATA u8_fe<>+8(SB)/8, $0xfefefefefefefefe
GLOBL u8_fe<>(SB), RODATA, $16

DATA u8_e0<>+0(SB)/8, $0xe0e0e0e0e0e0e0e0
DATA u8_e0<>+8(SB)/8, $0xe0e0e0e0e0e0e0e0
GLOBL u8_e0<>(SB), RODATA, $16

DATA u8_ed<>+0(SB)/8, $0xedededededededed
DATA u8_ed<>+8(SB)/8, $0xedededededededed
GLOBL u8_ed<>(SB), RODATA, $16

DATA u8_f0<>+0(SB)/8, $0xf0f0f0f0f0f0f0f0
DATA u8_f0<>+8(SB)/8, $0xf0f0f0f0f0f0f0f0
GLOBL u8_f0<>(SB), RODATA, $16

DATA u8_f4<>+0(SB)/8, $0xf4f4f4f4f4f4f4f4
DATA u8_f4<>+8(SB)/8, $0xf4f4f4f4f4f4f4f4
GLOBL u8_f4<>(SB), RODATA, $16

DATA u8_20<>+0(SB)/8, $0x2020202020202020
DATA u8_20<>+8(SB)/8, $0x2020202020202020
GLOBL u8_20<>(SB), RODATA, $16

DATA u8_1f<>+0(SB)/8, $0x1f1f1f1f1f1f1f1f
DATA u8_1f<>+8(SB)/8, $0x1f1f1f1f1f1f1f1f
GLOBL u8_1f<>(SB), RODATA, $16

DATA u8_10<>+0(SB)/8, $0x1010101010101010
DATA u8_10<>+8(SB)/8, $0x1010101010101010
GLOBL u8_10<>(SB), RODATA, $16

DATA u8_0f<>+0(SB)/8, $0x0f0f0f0f0f0f0f0f
DATA u8_0f<>+8(SB)/8, $0x0f0f0f0f0f0f0f0f
GLOBL u8_0f<>(SB), RODATA, $16

DATA u8_7f<>+0(SB)/8, $0x7f7f7f7f7f7f7f7f
DATA u8_7f<>+8(SB)/8, $0x7f7f7f7f7f7f7f7f
GLOBL u8_7f<>(SB), RODATA, $16

DATA u8_40<>+0(SB)/8, $0x4040404040404040
DATA u8_40<>+8(SB)/8, $0x4040404040404040
GLOBL u8_40<>(SB), RODATA, $16

DATA u8_08<>+0(SB)/8, $0x0808080808080808
DATA u8_08<>+8(SB)/8, $0x0808080808080808
GLOBL u8_08<>(SB), RODATA, $16

// func countUTF8Asm(ptr *byte, n int) (count int64, consumed int)
TEXT ·countUTF8Asm(SB), NOSPLIT, $0-32
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX

	XORQ	R8, R8		// total count
	XORQ	R9, R9		// total consumed

	CMPQ	CX, $16
	JL	u8c_done

u8c_ascii_loop:
	CMPQ	CX, $16
	JL	u8c_done
	MOVOU	0(SI), X1
	PMOVMSKB X1, AX
	TESTL	AX, AX
	JNZ	u8c_utf8_setup
	ADDQ	$16, SI
	ADDQ	$16, R8
	ADDQ	$16, R9
	SUBQ	$16, CX
	JMP	u8c_ascii_loop

u8c_utf8_setup:
	PXOR	X0, X0		// prev = 0
	MOVOU	u8_c0<>(SB), X12
	MOVOU	u8_80<>(SB), X13

u8c_block:
	MOVQ	CX, R10
	SHRQ	$4, R10
	TESTQ	R10, R10
	JZ	u8c_done
	CMPQ	R10, $240
	JLE	u8c_block_start
	MOVQ	$240, R10

u8c_block_start:
	MOVQ	R10, R11
	PXOR	X14, X14	// errAcc
	PXOR	X15, X15	// contAcc

u8c_loop:
	MOVOU	0(SI), X1	// cur
	ADDQ	$16, SI

	// prev1 (X2), prev2 (X3), prev3 (X4)
	MOVO	X0, X2
	PSRLDQ	$15, X2
	MOVO	X1, X5
	PSLLDQ	$1, X5
	POR	X5, X2

	MOVO	X0, X3
	PSRLDQ	$14, X3
	MOVO	X1, X5
	PSLLDQ	$2, X5
	POR	X5, X3

	MOVO	X0, X4
	PSRLDQ	$13, X4
	MOVO	X1, X5
	PSLLDQ	$3, X5
	POR	X5, X4

	// is_cont (X5): 0xC0 >s cur
	MOVO	X12, X5
	PCMPGTB	X1, X5
	PSUBB	X5, X15		// contAcc += is_cont

	// expect (X6) = (prev1_s > 0x3F) | (prev2_s > 0x5F) | (prev3_s > 0x6F)
	MOVO	X2, X6
	PXOR	X13, X6
	MOVOU	u8_3f<>(SB), X7
	PCMPGTB	X7, X6

	MOVO	X3, X7
	PXOR	X13, X7
	MOVOU	u8_5f<>(SB), X8
	PCMPGTB	X8, X7
	POR	X7, X6

	MOVO	X4, X7
	PXOR	X13, X7
	MOVOU	u8_6f<>(SB), X8
	PCMPGTB	X8, X7
	POR	X7, X6

	PXOR	X5, X6		// structural err

	// cur_s = cur ^ 0x80
	MOVO	X1, X8
	PXOR	X13, X8

	// cur >u 0xF4 (cur_s >s 0x74)
	MOVO	X8, X7
	MOVOU	u8_74<>(SB), X9
	PCMPGTB	X9, X7
	POR	X7, X6

	// (cur & 0xFE) == 0xC0
	MOVO	X1, X7
	MOVOU	u8_fe<>(SB), X9
	PAND	X9, X7
	PCMPEQB	X12, X7
	POR	X7, X6

	// prev1 == 0xE0 && cur_s < 0x20
	MOVOU	u8_20<>(SB), X7
	PCMPGTB	X8, X7
	MOVO	X2, X9
	MOVOU	u8_e0<>(SB), X10
	PCMPEQB	X10, X9
	PAND	X9, X7
	POR	X7, X6

	// prev1 == 0xED && cur_s > 0x1F (WTF-8 surrogates)
	MOVO	X8, X7
	MOVOU	u8_1f<>(SB), X9
	PCMPGTB	X9, X7
	MOVO	X2, X9
	MOVOU	u8_ed<>(SB), X10
	PCMPEQB	X10, X9
	PAND	X9, X7
	POR	X7, X6

	// prev1 == 0xF0 && cur_s < 0x10
	MOVOU	u8_10<>(SB), X7
	PCMPGTB	X8, X7
	MOVO	X2, X9
	MOVOU	u8_f0<>(SB), X10
	PCMPEQB	X10, X9
	PAND	X9, X7
	POR	X7, X6

	// prev1 == 0xF4 && cur_s > 0x0F
	MOVO	X8, X7
	MOVOU	u8_0f<>(SB), X9
	PCMPGTB	X9, X7
	MOVO	X2, X9
	MOVOU	u8_f4<>(SB), X10
	PCMPEQB	X10, X9
	PAND	X9, X7
	POR	X7, X6

	POR	X6, X14
	MOVO	X1, X0		// prev = cur

	DECQ	R10
	JNZ	u8c_loop

	PMOVMSKB X14, AX
	TESTL	AX, AX
	JNZ	u8c_done

	PXOR	X7, X7
	PSADBW	X7, X15
	MOVQ	X15, AX
	PSRLDQ	$8, X15
	MOVQ	X15, DX
	ADDQ	DX, AX		// total continuations

	MOVQ	R11, DX
	SHLQ	$4, DX		// bytes = chunks * 16
	SUBQ	AX, DX		// cps = bytes - cont
	ADDQ	DX, R8
	SHLQ	$4, R11
	ADDQ	R11, R9
	SUBQ	R11, CX
	JMP	u8c_block

u8c_done:
	MOVQ	R8, count+16(FP)
	MOVQ	R9, consumed+24(FP)
	RET

// func sumCodePointsLimitAsm(ptr *byte, n int, limit int64) (sum int64, consumed int, count int64)
// Baseline SSE2 16-byte/iter UTF-8 validator + parallel linear code-point summation via PSADBW.
TEXT ·sumCodePointsLimitAsm(SB), NOSPLIT, $0-48
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX
	MOVQ	limit+16(FP), BX

	XORQ	R8, R8		// total sum
	XORQ	R9, R9		// total consumed
	XORQ	R10, R10	// total count

	PXOR	X11, X11	// zero for PSADBW

u8s_ascii_loop:
	MOVQ	BX, R11
	SUBQ	R10, R11
	CMPQ	R11, CX
	JLE	u8s_ascii_chk
	MOVQ	CX, R11
u8s_ascii_chk:
	CMPQ	R11, $16
	JL	u8s_done
	MOVOU	0(SI), X1
	PMOVMSKB X1, AX
	TESTL	AX, AX
	JNZ	u8s_utf8_setup
	PSADBW	X11, X1
	MOVQ	X1, AX
	PSRLDQ	$8, X1
	MOVQ	X1, DX
	ADDQ	AX, R8
	ADDQ	DX, R8
	ADDQ	$16, SI
	ADDQ	$16, R9
	ADDQ	$16, R10
	SUBQ	$16, CX
	JMP	u8s_ascii_loop

u8s_utf8_setup:
	PXOR	X0, X0		// prev = 0
	MOVOU	u8_c0<>(SB), X12
	MOVOU	u8_80<>(SB), X13

u8s_block:
	MOVQ	BX, R11
	SUBQ	R10, R11	// remLimit = limit - count
	CMPQ	R11, CX
	JLE	u8s_use_lim
	MOVQ	CX, R11
u8s_use_lim:
	SHRQ	$4, R11
	TESTQ	R11, R11
	JZ	u8s_done
	CMPQ	R11, $240
	JLE	u8s_block_start
	MOVQ	$240, R11

u8s_block_start:
	MOVQ	R11, R12	// savedChunks
	PXOR	X14, X14	// errAcc
	PXOR	X15, X15	// contAcc
	XORQ	R13, R13	// S_payload
	XORQ	R14, R14	// S1
	XORQ	R15, R15	// S2
	XORQ	DI, DI		// S3

u8s_loop:
	MOVOU	0(SI), X1	// cur
	ADDQ	$16, SI

	// prev1 (X2), prev2 (X3), prev3 (X4)
	MOVO	X0, X2
	PSRLDQ	$15, X2
	MOVO	X1, X5
	PSLLDQ	$1, X5
	POR	X5, X2

	MOVO	X0, X3
	PSRLDQ	$14, X3
	MOVO	X1, X5
	PSLLDQ	$2, X5
	POR	X5, X3

	MOVO	X0, X4
	PSRLDQ	$13, X4
	MOVO	X1, X5
	PSLLDQ	$3, X5
	POR	X5, X4

	// is_cont (X5)
	MOVO	X12, X5
	PCMPGTB	X1, X5
	PSUBB	X5, X15

	// expect (X6)
	MOVO	X2, X6
	PXOR	X13, X6
	MOVOU	u8_3f<>(SB), X7
	PCMPGTB	X7, X6

	MOVO	X3, X7
	PXOR	X13, X7
	MOVOU	u8_5f<>(SB), X8
	PCMPGTB	X8, X7
	POR	X7, X6

	MOVO	X4, X7
	PXOR	X13, X7
	MOVOU	u8_6f<>(SB), X8
	PCMPGTB	X8, X7
	POR	X7, X6

	PXOR	X5, X6		// structural err

	// cur_s = cur ^ 0x80 (X8)
	MOVO	X1, X8
	PXOR	X13, X8

	// cur >u 0xF4
	MOVO	X8, X7
	MOVOU	u8_74<>(SB), X9
	PCMPGTB	X9, X7
	POR	X7, X6

	// (cur & 0xFE) == 0xC0
	MOVO	X1, X7
	MOVOU	u8_fe<>(SB), X9
	PAND	X9, X7
	PCMPEQB	X12, X7
	POR	X7, X6

	// prev1 == 0xE0 && cur_s < 0x20
	MOVOU	u8_20<>(SB), X7
	PCMPGTB	X8, X7
	MOVO	X2, X9
	MOVOU	u8_e0<>(SB), X10
	PCMPEQB	X10, X9
	PAND	X9, X7
	POR	X7, X6

	// prev1 == 0xED && cur_s > 0x1F
	MOVO	X8, X7
	MOVOU	u8_1f<>(SB), X9
	PCMPGTB	X9, X7
	MOVO	X2, X9
	MOVOU	u8_ed<>(SB), X10
	PCMPEQB	X10, X9
	PAND	X9, X7
	POR	X7, X6

	// prev1 == 0xF0 && cur_s < 0x10
	MOVOU	u8_10<>(SB), X7
	PCMPGTB	X8, X7
	MOVO	X2, X9
	MOVOU	u8_f0<>(SB), X10
	PCMPEQB	X10, X9
	PAND	X9, X7
	POR	X7, X6

	// prev1 == 0xF4 && cur_s > 0x0F
	MOVO	X8, X7
	MOVOU	u8_0f<>(SB), X9
	PCMPGTB	X9, X7
	MOVO	X2, X9
	MOVOU	u8_f4<>(SB), X10
	PCMPEQB	X10, X9
	PAND	X9, X7
	POR	X7, X6

	POR	X6, X14		// errAcc |= err

	// --- Parallel linear payload extraction (baseline SSE2) ---
	// cur_234 (X4), cur_34 (X5), cur_4 (X6)
	MOVO	X8, X4
	MOVOU	u8_3f<>(SB), X7
	PCMPGTB	X7, X4		// cur >= 0xC0

	MOVO	X8, X5
	MOVOU	u8_5f<>(SB), X7
	PCMPGTB	X7, X5		// cur >= 0xE0

	MOVO	X8, X6
	MOVOU	u8_6f<>(SB), X7
	PCMPGTB	X7, X6		// cur >= 0xF0 (= m3!)

	// payload (X7) = cur with UTF-8 prefix bits cleared:
	// clear bit7 always; bit6 if cur>=0x80; bit5 if cur_234; bit4 if cur_34; bit3 if cur_4
	MOVO	X11, X9
	PCMPGTB	X1, X9		// 0xFF iff 0 >s cur (i.e. cur >= 0x80)
	MOVOU	u8_40<>(SB), X10
	PAND	X10, X9		// bit 6 mask
	MOVO	X4, X10
	MOVOU	u8_20<>(SB), X7
	PAND	X7, X10
	POR	X10, X9		// | bit 5
	MOVO	X5, X10
	MOVOU	u8_10<>(SB), X7
	PAND	X7, X10
	POR	X10, X9		// | bit 4
	MOVO	X6, X10
	MOVOU	u8_08<>(SB), X7
	PAND	X7, X10
	POR	X10, X9		// | bit 3
	MOVO	X1, X7
	MOVOU	u8_7f<>(SB), X10
	PAND	X10, X7
	PANDN	X7, X9		// X9 = (~X9) & X7 = payload!

	// S_payload += sum(X9)
	MOVO	X9, X7
	PSADBW	X11, X7
	MOVQ	X7, AX
	PSRLDQ	$8, X7
	MOVQ	X7, DX
	ADDQ	AX, R13
	ADDQ	DX, R13

	// S3 += sum(payload & cur_4)
	MOVO	X9, X7
	PAND	X6, X7
	PSADBW	X11, X7
	MOVQ	X7, AX
	PSRLDQ	$8, X7
	MOVQ	X7, DX
	ADDQ	AX, DI
	ADDQ	DX, DI

	// prev1_s = prev1 ^ 0x80 (X2), prev2_s = prev2 ^ 0x80 (X3)
	PXOR	X13, X2
	PXOR	X13, X3

	// p1_4 (X8), p1_34 (X10), p2_4 (X3)
	MOVO	X2, X8
	MOVOU	u8_6f<>(SB), X7
	PCMPGTB	X7, X8		// p1_4
	MOVO	X2, X10
	MOVOU	u8_5f<>(SB), X7
	PCMPGTB	X7, X10		// p1_34
	MOVOU	u8_6f<>(SB), X7
	PCMPGTB	X7, X3		// p2_4

	// m2 = (cur_34 &~ cur_4) | p1_4
	// Plan 9 PANDN src, dst => dst = (~dst) & src
	MOVO	X6, X7		// cur_4
	PANDN	X5, X7		// (~cur_4) & cur_34 = cur_3
	POR	X8, X7		// m2
	PAND	X9, X7		// p2 = payload & m2
	PSADBW	X11, X7
	MOVQ	X7, AX
	PSRLDQ	$8, X7
	MOVQ	X7, DX
	ADDQ	AX, R15
	ADDQ	DX, R15

	// m1 = (cur_234 &~ cur_34) | (p1_34 &~ p1_4) | p2_4
	MOVO	X5, X7		// cur_34
	PANDN	X4, X7		// (~cur_34) & cur_234 = cur_2
	MOVO	X8, X2		// p1_4
	PANDN	X10, X2		// (~p1_4) & p1_34 = p1_3
	POR	X2, X7
	POR	X3, X7		// m1
	PAND	X9, X7		// p1 = payload & m1
	PSADBW	X11, X7
	MOVQ	X7, AX
	PSRLDQ	$8, X7
	MOVQ	X7, DX
	ADDQ	AX, R14
	ADDQ	DX, R14

	MOVO	X1, X0		// prev = cur

	DECQ	R11
	JNZ	u8s_loop

	PMOVMSKB X14, AX
	TESTL	AX, AX
	JNZ	u8s_done

	// totalSum += S_payload + 63*S1 + 4095*S2 + 262143*S3
	ADDQ	R13, R8
	SUBQ	R14, R8
	SHLQ	$6, R14
	ADDQ	R14, R8
	SUBQ	R15, R8
	SHLQ	$12, R15
	ADDQ	R15, R8
	SUBQ	DI, R8
	SHLQ	$18, DI
	ADDQ	DI, R8

	// continuations
	PSADBW	X11, X15
	MOVQ	X15, AX
	PSRLDQ	$8, X15
	MOVQ	X15, DX
	ADDQ	DX, AX

	MOVQ	R12, DX
	SHLQ	$4, DX		// bytes = chunks * 16
	SUBQ	AX, DX		// cps
	ADDQ	DX, R10
	SHLQ	$4, R12
	ADDQ	R12, R9
	SUBQ	R12, CX
	JMP	u8s_block

u8s_done:
	MOVQ	R8, sum+24(FP)
	MOVQ	R9, consumed+32(FP)
	MOVQ	R10, count+40(FP)
	RET

// func filterMaskEvenOddInt32Asm(ptr *int32, n int, mask int32, target int32) (evenSum, oddSum, evenCnt, oddCnt int32)
// Pure-SSE2 branchless filtered even/odd parity sum using 4-lane prefix-XOR.
TEXT ·filterMaskEvenOddInt32Asm(SB), NOSPLIT, $0-40
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX
	MOVL	mask+16(FP), R8
	MOVL	target+20(FP), R9

	MOVD	R8, X0
	PSHUFD	$0x00, X0, X0		// mask
	MOVD	R9, X1
	PSHUFD	$0x00, X1, X1		// target
	PXOR	X11, X11		// running parity (0 or 0xFFFFFFFF)

	PXOR	X12, X12		// evenSum
	PXOR	X13, X13		// totalSum
	PXOR	X14, X14		// totalCnt

fm_loop4:
	CMPQ	CX, $4
	JL	fm_reduce
	MOVOU	0(SI), X2
	MOVO	X2, X3
	PAND	X0, X3
	PCMPEQL	X1, X3			// match mask
	PSUBL	X3, X14			// totalCnt += popcount
	MOVO	X3, X4
	PSLLDQ	$4, X4
	PXOR	X3, X4
	MOVO	X4, X5
	PSLLDQ	$8, X5
	PXOR	X5, X4			// P = inclusive 4-lane prefix XOR
	MOVO	X4, X5
	PXOR	X11, X5			// P ^ curPar
	PSHUFD	$0xFF, X4, X6		// broadcast total parity of this 4-elem vector
	PXOR	X6, X11			// curPar ^= P[3]
	PAND	X2, X3			// matched values
	PADDD	X3, X13			// totalSum
	PAND	X3, X5			// even-indexed matched values
	PADDD	X5, X12			// evenSum
	ADDQ	$16, SI
	SUBQ	$4, CX
	JMP	fm_loop4

fm_reduce:
	MOVD	X11, R11		// 0 (even) or 0xFFFFFFFF (odd)

	MOVO	X12, X2
	PSRLDQ	$8, X2
	PADDD	X2, X12
	MOVO	X12, X2
	PSRLDQ	$4, X2
	PADDD	X2, X12
	MOVD	X12, AX			// evenSum

	MOVO	X13, X2
	PSRLDQ	$8, X2
	PADDD	X2, X13
	MOVO	X13, X2
	PSRLDQ	$4, X2
	PADDD	X2, X13
	MOVD	X13, DX			// totalSum
	SUBL	AX, DX			// oddSum = totalSum - evenSum

	MOVO	X14, X2
	PSRLDQ	$8, X2
	PADDD	X2, X14
	MOVO	X14, X2
	PSRLDQ	$4, X2
	PADDD	X2, X14
	MOVD	X14, BX			// totalCnt
	MOVL	BX, R12
	SHRL	$1, R12			// oddCnt
	INCL	BX
	SHRL	$1, BX			// evenCnt

fm_tail1:
	TESTQ	CX, CX
	JLE	fm_done
	MOVL	0(SI), R13
	MOVL	R13, R14
	ANDL	R8, R14
	CMPL	R14, R9
	JNE	fm_next
	TESTQ	R11, R11
	JNZ	fm_odd
	ADDL	R13, AX
	INCL	BX
	MOVQ	$1, R11
	JMP	fm_next
fm_odd:
	ADDL	R13, DX
	INCL	R12
	XORQ	R11, R11
fm_next:
	ADDQ	$4, SI
	DECQ	CX
	JMP	fm_tail1

fm_done:
	MOVL	AX, evenSum+24(FP)
	MOVL	DX, oddSum+28(FP)
	MOVL	BX, evenCnt+32(FP)
	MOVL	R12, oddCnt+36(FP)
	RET

// func indexByteASCIIAsm(ptr *byte, n int, c byte) (idx int, prefixASCII bool)
TEXT ·indexByteASCIIAsm(SB), NOSPLIT, $0-33
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX
	MOVBLZX	c+16(FP), R8

	MOVQ	SI, R10
	MOVQ	R8, AX
	IMULQ	$0x01010101, AX
	MOVD	AX, X0
	PSHUFD	$0x00, X0, X0		// broadcast c
	PXOR	X15, X15		// prefix high-bit OR accumulator

iba_loop16:
	CMPQ	CX, $16
	JL	iba_tail1
	MOVOU	0(SI), X1
	MOVO	X1, X2
	PCMPEQB	X0, X2
	PMOVMSKB X2, DX
	TESTL	DX, DX
	JNZ	iba_hit16
	POR	X1, X15
	ADDQ	$16, SI
	SUBQ	$16, CX
	JMP	iba_loop16

iba_tail1:
	XORL	R9, R9
iba_tail1_loop:
	TESTQ	CX, CX
	JLE	iba_miss
	MOVBLZX	0(SI), AX
	INCQ	SI
	CMPL	AX, R8
	JE	iba_hit_done
	ORL	AX, R9
	DECQ	CX
	JMP	iba_tail1_loop

iba_hit16:
	XORL	R9, R9
iba_hit16_scan:
	MOVBLZX	0(SI), AX
	INCQ	SI
	CMPL	AX, R8
	JE	iba_hit_done
	ORL	AX, R9
	JMP	iba_hit16_scan

iba_hit_done:
	DECQ	SI
	SUBQ	R10, SI
	MOVQ	SI, idx+24(FP)
	PMOVMSKB X15, AX
	TESTL	AX, AX
	JNZ	iba_non_ascii
	TESTL	$0x80, R9
	JNZ	iba_non_ascii
	MOVB	$1, prefixASCII+32(FP)
	RET

iba_non_ascii:
	MOVB	$0, prefixASCII+32(FP)
	RET

iba_miss:
	MOVQ	$-1, AX
	MOVQ	AX, idx+24(FP)
	MOVB	$0, prefixASCII+32(FP)
	RET
