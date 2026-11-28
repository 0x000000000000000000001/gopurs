//go:build amd64 && go1.27

#include "textflag.h"

// func sumInt32Asm(ptr *int32, n int) int32
// Wrapping int32 sum using baseline SSE2 vector additions (16 int32s per iteration).
TEXT ·sumInt32Asm(SB), NOSPLIT, $0-20
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX

	PXOR	X0, X0
	PXOR	X1, X1
	PXOR	X2, X2
	PXOR	X3, X3

	CMPQ	CX, $16
	JL	sum_tail4_check

sum_loop16:
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
	JGE	sum_loop16

	PADDD	X1, X0
	PADDD	X3, X2
	PADDD	X2, X0

sum_tail4_check:
	CMPQ	CX, $4
	JL	sum_reduce

sum_loop4:
	MOVOU	0(SI), X4
	PADDD	X4, X0
	ADDQ	$16, SI
	SUBQ	$4, CX
	CMPQ	CX, $4
	JGE	sum_loop4

sum_reduce:
	PSHUFD	$0xEE, X0, X1
	PADDD	X1, X0
	MOVQ	X0, AX
	MOVL	AX, DX
	SHRQ	$32, AX
	ADDL	AX, DX

sum_tail1:
	TESTQ	CX, CX
	JLE	sum_done
	ADDL	0(SI), DX
	ADDQ	$4, SI
	DECQ	CX
	JMP	sum_tail1

sum_done:
	MOVL	DX, ret+16(FP)
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

// func indexOfInt32Asm(ptr *int32, n int, target int32) int
// Returns the lowest index i in [0, n) with ptr[i] == target, or -1.
TEXT ·indexOfInt32Asm(SB), NOSPLIT, $0-32
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX
	MOVL	target+16(FP), AX

	MOVQ	SI, R10
	MOVD	AX, X0
	PSHUFD	$0x00, X0, X0

	CMPQ	CX, $16
	JL	io_tail4_check

io_loop16:
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
	JNZ	io_scan
	ADDQ	$64, SI
	SUBQ	$16, CX
	CMPQ	CX, $16
	JGE	io_loop16

io_tail4_check:
	CMPQ	CX, $4
	JL	io_tail1

io_loop4:
	MOVOU	0(SI), X1
	PCMPEQL	X0, X1
	PMOVMSKB X1, DX
	TESTL	DX, DX
	JNZ	io_scan
	ADDQ	$16, SI
	SUBQ	$4, CX
	CMPQ	CX, $4
	JGE	io_loop4

io_tail1:
	TESTQ	CX, CX
	JLE	io_miss
	CMPL	0(SI), AX
	JE	io_found
	ADDQ	$4, SI
	DECQ	CX
	JMP	io_tail1

io_scan:
	CMPL	0(SI), AX
	JE	io_found
	ADDQ	$4, SI
	JMP	io_scan

io_found:
	SUBQ	R10, SI
	SHRQ	$2, SI
	MOVQ	SI, ret+24(FP)
	RET

io_miss:
	MOVQ	$-1, AX
	MOVQ	AX, ret+24(FP)
	RET

// Baseline SSE2 constants for 16-byte RFC 3629 UTF-8 validation.
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

// func firstNonASCIIAsm(ptr *byte, n int) int
// Returns the byte offset of the first byte >= 0x80, or n if all bytes are ASCII.
TEXT ·firstNonASCIIAsm(SB), NOSPLIT, $0-24
	MOVQ	ptr+0(FP), SI
	MOVQ	n+8(FP), CX
	MOVQ	SI, R10

fna_loop16:
	CMPQ	CX, $16
	JL	fna_tail1
	MOVOU	0(SI), X1
	PMOVMSKB X1, AX
	TESTL	AX, AX
	JNZ	fna_hit16
	ADDQ	$16, SI
	SUBQ	$16, CX
	JMP	fna_loop16

fna_tail1:
	TESTQ	CX, CX
	JLE	fna_all_ascii
	MOVBLZX	0(SI), AX
	TESTL	$0x80, AX
	JNZ	fna_found
	INCQ	SI
	DECQ	CX
	JMP	fna_tail1

fna_hit16:
	MOVBLZX	0(SI), AX
	TESTL	$0x80, AX
	JNZ	fna_found
	INCQ	SI
	JMP	fna_hit16

fna_found:
	SUBQ	R10, SI
	MOVQ	SI, ret+16(FP)
	RET

fna_all_ascii:
	SUBQ	R10, SI
	MOVQ	SI, ret+16(FP)
	RET
