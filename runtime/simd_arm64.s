//go:build arm64 && go1.27

#include "textflag.h"

// func sumInt32Asm(ptr *int32, n int) int32
// Wrapping int32 sum using 128-bit NEON vector additions (32 int32s / 128 B per iteration).
TEXT ·sumInt32Asm(SB), NOSPLIT, $0-20
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1

	VEOR	V0.B16, V0.B16, V0.B16
	SUBS	$32, R1, R1
	BLT	sum_restore32

	VEOR	V1.B16, V1.B16, V1.B16
	VEOR	V2.B16, V2.B16, V2.B16
	VEOR	V3.B16, V3.B16, V3.B16
	VEOR	V16.B16, V16.B16, V16.B16
	VEOR	V17.B16, V17.B16, V17.B16
	VEOR	V18.B16, V18.B16, V18.B16
	VEOR	V19.B16, V19.B16, V19.B16

sum_loop32:
	VLD1.P	64(R0), [V4.S4, V5.S4, V6.S4, V7.S4]
	VLD1.P	64(R0), [V20.S4, V21.S4, V22.S4, V23.S4]
	VADD	V4.S4, V0.S4, V0.S4
	VADD	V5.S4, V1.S4, V1.S4
	VADD	V6.S4, V2.S4, V2.S4
	VADD	V7.S4, V3.S4, V3.S4
	VADD	V20.S4, V16.S4, V16.S4
	VADD	V21.S4, V17.S4, V17.S4
	VADD	V22.S4, V18.S4, V18.S4
	VADD	V23.S4, V19.S4, V19.S4
	SUBS	$32, R1, R1
	BGE	sum_loop32

	VADD	V16.S4, V0.S4, V0.S4
	VADD	V17.S4, V1.S4, V1.S4
	VADD	V18.S4, V2.S4, V2.S4
	VADD	V19.S4, V3.S4, V3.S4
	VADD	V1.S4, V0.S4, V0.S4
	VADD	V3.S4, V2.S4, V2.S4
	VADD	V2.S4, V0.S4, V0.S4

sum_restore32:
	ADD	$32, R1
	CMP	$16, R1
	BLT	sum_tail4_check

	VLD1.P	64(R0), [V4.S4, V5.S4, V6.S4, V7.S4]
	VADD	V4.S4, V0.S4, V0.S4
	VADD	V5.S4, V6.S4, V5.S4
	VADD	V7.S4, V0.S4, V0.S4
	VADD	V5.S4, V0.S4, V0.S4
	SUB	$16, R1

sum_tail4_check:
	CMP	$4, R1
	BLT	sum_reduce

sum_loop4:
	VLD1.P	16(R0), [V4.S4]
	VADD	V4.S4, V0.S4, V0.S4
	SUB	$4, R1
	CMP	$4, R1
	BGE	sum_loop4

sum_reduce:
	VMOV	V0.S[0], R2
	VMOV	V0.S[1], R3
	VMOV	V0.S[2], R4
	VMOV	V0.S[3], R5
	ADDW	R3, R2, R2
	ADDW	R5, R4, R4
	ADDW	R4, R2, R2

sum_tail1:
	CBZ	R1, sum_done
	MOVW.P	4(R0), R4
	ADDW	R4, R2, R2
	SUB	$1, R1
	B	sum_tail1

sum_done:
	MOVW	R2, ret+16(FP)
	RET

// func isNonIncreasingInt32Asm(ptr *int32, n int) bool
// Returns true iff xs[i] >= xs[i+1] for all 0 <= i < n-1.
TEXT ·isNonIncreasingInt32Asm(SB), NOSPLIT, $0-17
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1

	CMP	$2, R1
	BLT	ni_true

	MOVW	(R0), R2
	MOVW	4(R0), R3
	CMPW	R2, R3
	BGT	ni_false

	CMP	$20, R1
	BLT	ni_scalar_from_1

	MOVW	8(R0), R4
	CMPW	R3, R4
	BGT	ni_false
	MOVW	12(R0), R5
	CMPW	R4, R5
	BGT	ni_false

	VLD1.P	16(R0), [V0.S4]
	SUB	$4, R1
	VEOR	V17.B16, V17.B16, V17.B16

	CMP	$64, R1
	BLT	ni_check16

ni_loop64:
	MOVD	$4, R6
ni_loop16_in64:
	VLD1.P	64(R0), [V1.S4, V2.S4, V3.S4, V4.S4]
	VEXT	$4, V1.B16, V0.B16, V5.B16
	VEXT	$4, V2.B16, V1.B16, V6.B16
	VEXT	$4, V3.B16, V2.B16, V7.B16
	VEXT	$4, V4.B16, V3.B16, V16.B16
	VCMGT	V0.S4, V5.S4, V5.S4
	VCMGT	V1.S4, V6.S4, V6.S4
	VCMGT	V2.S4, V7.S4, V7.S4
	VCMGT	V3.S4, V16.S4, V16.S4
	VORR	V5.B16, V6.B16, V5.B16
	VORR	V7.B16, V16.B16, V7.B16
	VORR	V5.B16, V7.B16, V5.B16
	VORR	V5.B16, V17.B16, V17.B16
	VORR	V4.B16, V4.B16, V0.B16
	SUBS	$1, R6, R6
	BNE	ni_loop16_in64
	VMOV	V17.D[0], R2
	VMOV	V17.D[1], R3
	ORR	R2, R3, R2
	CBNZ	R2, ni_false
	SUB	$64, R1
	CMP	$64, R1
	BGE	ni_loop64

ni_check16:
	CMP	$16, R1
	BLT	ni_tail_setup

ni_loop16:
	VLD1.P	64(R0), [V1.S4, V2.S4, V3.S4, V4.S4]
	VEXT	$4, V1.B16, V0.B16, V5.B16
	VEXT	$4, V2.B16, V1.B16, V6.B16
	VEXT	$4, V3.B16, V2.B16, V7.B16
	VEXT	$4, V4.B16, V3.B16, V16.B16
	VCMGT	V0.S4, V5.S4, V5.S4
	VCMGT	V1.S4, V6.S4, V6.S4
	VCMGT	V2.S4, V7.S4, V7.S4
	VCMGT	V3.S4, V16.S4, V16.S4
	VORR	V5.B16, V6.B16, V5.B16
	VORR	V7.B16, V16.B16, V7.B16
	VORR	V5.B16, V7.B16, V5.B16
	VORR	V5.B16, V17.B16, V17.B16
	VORR	V4.B16, V4.B16, V0.B16
	SUB	$16, R1
	CMP	$16, R1
	BGE	ni_loop16

	VMOV	V17.D[0], R2
	VMOV	V17.D[1], R3
	ORR	R2, R3, R2
	CBNZ	R2, ni_false

ni_tail_setup:
	SUB	$16, R0
	ADD	$4, R1
	MOVW.P	4(R0), R2
	SUB	$1, R1
	B	ni_scalar_loop

ni_scalar_from_1:
	MOVW	R3, R2
	ADD	$8, R0
	SUB	$2, R1

ni_scalar_loop:
	CBZ	R1, ni_true
	MOVW.P	4(R0), R3
	CMPW	R2, R3
	BGT	ni_false
	MOVW	R3, R2
	SUB	$1, R1
	B	ni_scalar_loop

ni_true:
	MOVD	$1, R2
	MOVB	R2, ret+16(FP)
	RET

ni_false:
	MOVB	ZR, ret+16(FP)
	RET

// func isNonDecreasingInt32Asm(ptr *int32, n int) bool
// Returns true iff xs[i] <= xs[i+1] for all 0 <= i < n-1.
TEXT ·isNonDecreasingInt32Asm(SB), NOSPLIT, $0-17
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1

	CMP	$2, R1
	BLT	nd_true

	MOVW	(R0), R2
	MOVW	4(R0), R3
	CMPW	R2, R3
	BLT	nd_false

	CMP	$20, R1
	BLT	nd_scalar_from_1

	MOVW	8(R0), R4
	CMPW	R3, R4
	BLT	nd_false
	MOVW	12(R0), R5
	CMPW	R4, R5
	BLT	nd_false

	VLD1.P	16(R0), [V0.S4]
	SUB	$4, R1
	VEOR	V17.B16, V17.B16, V17.B16

	CMP	$64, R1
	BLT	nd_check16

nd_loop64:
	MOVD	$4, R6
nd_loop16_in64:
	VLD1.P	64(R0), [V1.S4, V2.S4, V3.S4, V4.S4]
	VEXT	$4, V1.B16, V0.B16, V5.B16
	VEXT	$4, V2.B16, V1.B16, V6.B16
	VEXT	$4, V3.B16, V2.B16, V7.B16
	VEXT	$4, V4.B16, V3.B16, V16.B16
	VCMGT	V5.S4, V0.S4, V5.S4
	VCMGT	V6.S4, V1.S4, V6.S4
	VCMGT	V7.S4, V2.S4, V7.S4
	VCMGT	V16.S4, V3.S4, V16.S4
	VORR	V5.B16, V6.B16, V5.B16
	VORR	V7.B16, V16.B16, V7.B16
	VORR	V5.B16, V7.B16, V5.B16
	VORR	V5.B16, V17.B16, V17.B16
	VORR	V4.B16, V4.B16, V0.B16
	SUBS	$1, R6, R6
	BNE	nd_loop16_in64
	VMOV	V17.D[0], R2
	VMOV	V17.D[1], R3
	ORR	R2, R3, R2
	CBNZ	R2, nd_false
	SUB	$64, R1
	CMP	$64, R1
	BGE	nd_loop64

nd_check16:
	CMP	$16, R1
	BLT	nd_tail_setup

nd_loop16:
	VLD1.P	64(R0), [V1.S4, V2.S4, V3.S4, V4.S4]
	VEXT	$4, V1.B16, V0.B16, V5.B16
	VEXT	$4, V2.B16, V1.B16, V6.B16
	VEXT	$4, V3.B16, V2.B16, V7.B16
	VEXT	$4, V4.B16, V3.B16, V16.B16
	VCMGT	V5.S4, V0.S4, V5.S4
	VCMGT	V6.S4, V1.S4, V6.S4
	VCMGT	V7.S4, V2.S4, V7.S4
	VCMGT	V16.S4, V3.S4, V16.S4
	VORR	V5.B16, V6.B16, V5.B16
	VORR	V7.B16, V16.B16, V7.B16
	VORR	V5.B16, V7.B16, V5.B16
	VORR	V5.B16, V17.B16, V17.B16
	VORR	V4.B16, V4.B16, V0.B16
	SUB	$16, R1
	CMP	$16, R1
	BGE	nd_loop16

	VMOV	V17.D[0], R2
	VMOV	V17.D[1], R3
	ORR	R2, R3, R2
	CBNZ	R2, nd_false

nd_tail_setup:
	SUB	$16, R0
	ADD	$4, R1
	MOVW.P	4(R0), R2
	SUB	$1, R1
	B	nd_scalar_loop

nd_scalar_from_1:
	MOVW	R3, R2
	ADD	$8, R0
	SUB	$2, R1

nd_scalar_loop:
	CBZ	R1, nd_true
	MOVW.P	4(R0), R3
	CMPW	R2, R3
	BLT	nd_false
	MOVW	R3, R2
	SUB	$1, R1
	B	nd_scalar_loop

nd_true:
	MOVD	$1, R2
	MOVB	R2, ret+16(FP)
	RET

nd_false:
	MOVB	ZR, ret+16(FP)
	RET

// func indexOfInt32Asm(ptr *int32, n int, target int32) int
// Returns the lowest index i in [0, n) with ptr[i] == target, or -1.
TEXT ·indexOfInt32Asm(SB), NOSPLIT, $0-32
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1
	MOVW	target+16(FP), R2

	MOVD	R0, R10
	VDUP	R2, V0.S4

	CMP	$16, R1
	BLT	io_tail4_check

io_loop16:
	VLD1.P	64(R0), [V1.S4, V2.S4, V3.S4, V4.S4]
	VCMEQ	V0.S4, V1.S4, V1.S4
	VCMEQ	V0.S4, V2.S4, V2.S4
	VCMEQ	V0.S4, V3.S4, V3.S4
	VCMEQ	V0.S4, V4.S4, V4.S4
	VORR	V2.B16, V1.B16, V1.B16
	VORR	V4.B16, V3.B16, V3.B16
	VORR	V3.B16, V1.B16, V1.B16
	VMOV	V1.D[0], R3
	VMOV	V1.D[1], R4
	ORR	R3, R4, R3
	CBNZ	R3, io_hit16
	SUB	$16, R1
	CMP	$16, R1
	BGE	io_loop16

io_tail4_check:
	CMP	$4, R1
	BLT	io_tail1

io_loop4:
	VLD1.P	16(R0), [V1.S4]
	VCMEQ	V0.S4, V1.S4, V1.S4
	VMOV	V1.D[0], R3
	VMOV	V1.D[1], R4
	ORR	R3, R4, R3
	CBNZ	R3, io_hit4
	SUB	$4, R1
	CMP	$4, R1
	BGE	io_loop4

io_tail1:
	CBZ	R1, io_miss
	MOVW.P	4(R0), R3
	CMPW	R2, R3
	BEQ	io_found
	SUB	$1, R1
	B	io_tail1

io_hit16:
	SUB	$64, R0
io_scan:
	MOVW.P	4(R0), R3
	CMPW	R2, R3
	BNE	io_scan
	B	io_found

io_hit4:
	SUB	$16, R0
	B	io_scan

io_found:
	SUB	$4, R0
	SUB	R10, R0, R0
	LSR	$2, R0, R0
	MOVD	R0, ret+24(FP)
	RET

io_miss:
	MOVD	$-1, R0
	MOVD	R0, ret+24(FP)
	RET

// Contiguous 128-byte constant block for 16-byte NEON RFC 3629 UTF-8 validation:
//   +0:   0xBF   +16:  0xDF   +32:  0xEF   +48:  0xC0   +64:  0x0F
//   +80:  p1_hi  +96:  p1_lo  +112: cur_hi
DATA utf8_consts<>+0(SB)/8, $0xbfbfbfbfbfbfbfbf
DATA utf8_consts<>+8(SB)/8, $0xbfbfbfbfbfbfbfbf
DATA utf8_consts<>+16(SB)/8, $0xdfdfdfdfdfdfdfdf
DATA utf8_consts<>+24(SB)/8, $0xdfdfdfdfdfdfdfdf
DATA utf8_consts<>+32(SB)/8, $0xefefefefefefefef
DATA utf8_consts<>+40(SB)/8, $0xefefefefefefefef
DATA utf8_consts<>+48(SB)/8, $0xc0c0c0c0c0c0c0c0
DATA utf8_consts<>+56(SB)/8, $0xc0c0c0c0c0c0c0c0
DATA utf8_consts<>+64(SB)/8, $0x0f0f0f0f0f0f0f0f
DATA utf8_consts<>+72(SB)/8, $0x0f0f0f0f0f0f0f0f
DATA utf8_consts<>+80(SB)/8, $0x0000000000000000
DATA utf8_consts<>+88(SB)/8, $0x5806002000000000
DATA utf8_consts<>+96(SB)/8, $0x404040100000202a
DATA utf8_consts<>+104(SB)/8, $0x4040444040404040
DATA utf8_consts<>+112(SB)/8, $0x6060606060606060
DATA utf8_consts<>+120(SB)/8, $0x606060607474726a
GLOBL utf8_consts<>(SB), RODATA, $128

// func countUTF8Asm(ptr *byte, n int) (count int64, consumed int)
// Validates RFC 3629 UTF-8 (bailing on surrogates 0xED 0xA0..0xBF or invalid
// bytes) in 16-byte NEON chunks (with a 64-byte pure-ASCII fast path) and
// returns (non-continuation count, bytes consumed).
TEXT ·countUTF8Asm(SB), NOSPLIT, $0-32
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1

	MOVD	$0, R2		// total count
	MOVD	$0, R3		// total consumed

	CMP	$64, R1
	BLT	u8c_utf8_setup

	MOVD	$0x8080808080808080, R9
	VLD1	(R0), [V1.B16, V2.B16, V3.B16, V4.B16]
	VORR	V2.B16, V1.B16, V5.B16
	VORR	V4.B16, V3.B16, V6.B16
	VORR	V6.B16, V5.B16, V5.B16
	VMOV	V5.D[0], R4
	VMOV	V5.D[1], R7
	ORR	R4, R7, R4
	TST	R9, R4
	BNE	u8c_utf8_setup

u8c_ascii_block:
	LSR	$6, R1, R5
	CBZ	R5, u8c_utf8_setup
	CMP	$256, R5
	BLE	u8c_ascii_start
	MOVD	$256, R5

u8c_ascii_start:
	MOVD	R0, R10
	MOVD	R5, R6
	VEOR	V28.B16, V28.B16, V28.B16

u8c_ascii_loop:
	VLD1.P	64(R0), [V1.B16, V2.B16, V3.B16, V4.B16]
	VORR	V2.B16, V1.B16, V5.B16
	VORR	V4.B16, V3.B16, V6.B16
	VORR	V6.B16, V5.B16, V5.B16
	VORR	V5.B16, V28.B16, V28.B16
	SUBS	$1, R5, R5
	BNE	u8c_ascii_loop

	VMOV	V28.D[0], R4
	VMOV	V28.D[1], R7
	ORR	R4, R7, R4
	TST	R9, R4
	BNE	u8c_ascii_bail

	LSL	$6, R6, R7
	ADD	R7, R2, R2
	ADD	R7, R3, R3
	SUB	R7, R1, R1
	B	u8c_ascii_block

u8c_ascii_bail:
	MOVD	R10, R0

u8c_utf8_setup:
	CMP	$16, R1
	BLT	u8c_done

	MOVD	$utf8_consts<>(SB), R4
	VLD1.P	48(R4), [V16.B16, V17.B16, V18.B16]
	VLD1.P	32(R4), [V20.B16, V21.B16]
	VLD1	(R4), [V22.B16, V23.B16, V24.B16]

	VEOR	V0.B16, V0.B16, V0.B16		// prev = 0
	VEOR	V26.B16, V26.B16, V26.B16	// prev_hi = 0

u8c_block:
	MOVD	R1, R5
	LSR	$5, R5			// pairs of 16 bytes (32 bytes/iter)
	CBZ	R5, u8c_check16
	CMP	$120, R5
	BLE	u8c_block_start
	MOVD	$120, R5

u8c_block_start:
	MOVD	R5, R6
	VEOR	V29.B16, V29.B16, V29.B16	// errAcc
	VEOR	V30.B16, V30.B16, V30.B16	// contAcc

u8c_loop32:
	// Chunk A: cur = V1, prev = V0, prev_hi = V26
	VLD1.P	16(R0), [V1.B16]
	VUSHR	$4, V1.B16, V10.B16

	VEXT	$15, V1.B16, V0.B16, V2.B16
	VEXT	$14, V1.B16, V0.B16, V3.B16
	VEXT	$13, V1.B16, V0.B16, V4.B16
	VEXT	$15, V10.B16, V26.B16, V11.B16

	VCMGT	V1.B16, V20.B16, V5.B16
	VSUB	V5.B16, V30.B16, V30.B16

	VCMHI	V16.B16, V2.B16, V6.B16
	VCMHI	V17.B16, V3.B16, V7.B16
	VCMHI	V18.B16, V4.B16, V8.B16
	VORR	V7.B16, V6.B16, V6.B16
	VORR	V8.B16, V6.B16, V6.B16
	VEOR	V5.B16, V6.B16, V6.B16

	VAND	V21.B16, V2.B16, V12.B16
	VTBL	V11.B16, [V22.B16], V7.B16
	VTBL	V12.B16, [V23.B16], V8.B16
	VTBL	V10.B16, [V24.B16], V9.B16
	VAND	V8.B16, V7.B16, V7.B16
	VAND	V9.B16, V7.B16, V7.B16
	VORR	V7.B16, V6.B16, V6.B16
	VORR	V6.B16, V29.B16, V29.B16

	// Chunk B: cur = V0 (becomes next prev), cur_hi = V26 (becomes next prev_hi)
	VLD1.P	16(R0), [V0.B16]
	VUSHR	$4, V0.B16, V26.B16

	VEXT	$15, V0.B16, V1.B16, V2.B16
	VEXT	$14, V0.B16, V1.B16, V3.B16
	VEXT	$13, V0.B16, V1.B16, V4.B16
	VEXT	$15, V26.B16, V10.B16, V11.B16

	VCMGT	V0.B16, V20.B16, V5.B16
	VSUB	V5.B16, V30.B16, V30.B16

	VCMHI	V16.B16, V2.B16, V6.B16
	VCMHI	V17.B16, V3.B16, V7.B16
	VCMHI	V18.B16, V4.B16, V8.B16
	VORR	V7.B16, V6.B16, V6.B16
	VORR	V8.B16, V6.B16, V6.B16
	VEOR	V5.B16, V6.B16, V6.B16

	VAND	V21.B16, V2.B16, V12.B16
	VTBL	V11.B16, [V22.B16], V7.B16
	VTBL	V12.B16, [V23.B16], V8.B16
	VTBL	V26.B16, [V24.B16], V9.B16
	VAND	V8.B16, V7.B16, V7.B16
	VAND	V9.B16, V7.B16, V7.B16
	VORR	V7.B16, V6.B16, V6.B16
	VORR	V6.B16, V29.B16, V29.B16

	SUBS	$1, R5, R5
	BNE	u8c_loop32

	VMOV	V29.D[0], R4
	VMOV	V29.D[1], R7
	ORR	R4, R7, R4
	CBNZ	R4, u8c_done

	VUADDLV	V30.B16, V7
	VMOV	V7.H[0], R4

	LSL	$5, R6, R7
	SUB	R4, R7, R4
	ADD	R4, R2, R2
	ADD	R7, R3, R3
	SUB	R7, R1, R1
	B	u8c_block

u8c_check16:
	CMP	$16, R1
	BLT	u8c_done

	VLD1.P	16(R0), [V1.B16]
	VUSHR	$4, V1.B16, V10.B16

	VEXT	$15, V1.B16, V0.B16, V2.B16
	VEXT	$14, V1.B16, V0.B16, V3.B16
	VEXT	$13, V1.B16, V0.B16, V4.B16
	VEXT	$15, V10.B16, V26.B16, V11.B16

	VCMGT	V1.B16, V20.B16, V5.B16
	VEOR	V30.B16, V30.B16, V30.B16
	VSUB	V5.B16, V30.B16, V30.B16

	VCMHI	V16.B16, V2.B16, V6.B16
	VCMHI	V17.B16, V3.B16, V7.B16
	VCMHI	V18.B16, V4.B16, V8.B16
	VORR	V7.B16, V6.B16, V6.B16
	VORR	V8.B16, V6.B16, V6.B16
	VEOR	V5.B16, V6.B16, V6.B16

	VAND	V21.B16, V2.B16, V12.B16
	VTBL	V11.B16, [V22.B16], V7.B16
	VTBL	V12.B16, [V23.B16], V8.B16
	VTBL	V10.B16, [V24.B16], V9.B16
	VAND	V8.B16, V7.B16, V7.B16
	VAND	V9.B16, V7.B16, V7.B16
	VORR	V7.B16, V6.B16, V6.B16

	VMOV	V6.D[0], R4
	VMOV	V6.D[1], R7
	ORR	R4, R7, R4
	CBNZ	R4, u8c_done

	VUADDLV	V30.B16, V7
	VMOV	V7.H[0], R4
	MOVD	$16, R7
	SUB	R4, R7, R4
	ADD	R4, R2, R2
	ADD	$16, R3, R3

u8c_done:
	MOVD	R2, count+16(FP)
	MOVD	R3, consumed+24(FP)
	RET

// func indexByteASCIIAsm(ptr *byte, n int, c byte) (idx int, prefixASCII bool)
// Single-pass 64-byte/iter NEON byte search that simultaneously checks whether
// all bytes strictly before the match index are pure ASCII (< 0x80).
TEXT ·indexByteASCIIAsm(SB), NOSPLIT, $0-33
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1
	MOVBU	c+16(FP), R2

	MOVD	R0, R10
	VDUP	R2, V0.B16
	VEOR	V16.B16, V16.B16, V16.B16

	CMP	$64, R1
	BLT	iba_check16

iba_loop64:
	VLD1.P	64(R0), [V1.B16, V2.B16, V3.B16, V4.B16]
	VCMEQ	V0.B16, V1.B16, V5.B16
	VCMEQ	V0.B16, V2.B16, V6.B16
	VCMEQ	V0.B16, V3.B16, V7.B16
	VCMEQ	V0.B16, V4.B16, V8.B16
	VORR	V6.B16, V5.B16, V5.B16
	VORR	V8.B16, V7.B16, V7.B16
	VORR	V7.B16, V5.B16, V5.B16
	VMOV	V5.D[0], R3
	VMOV	V5.D[1], R4
	ORR	R3, R4, R3
	CBNZ	R3, iba_hit64
	VORR	V2.B16, V1.B16, V1.B16
	VORR	V4.B16, V3.B16, V3.B16
	VORR	V3.B16, V1.B16, V1.B16
	VORR	V1.B16, V16.B16, V16.B16
	SUB	$64, R1
	CMP	$64, R1
	BGE	iba_loop64

iba_check16:
	CMP	$16, R1
	BLT	iba_tail1

iba_loop16:
	VLD1.P	16(R0), [V1.B16]
	VCMEQ	V0.B16, V1.B16, V5.B16
	VMOV	V5.D[0], R3
	VMOV	V5.D[1], R4
	ORR	R3, R4, R3
	CBNZ	R3, iba_hit16
	VORR	V1.B16, V16.B16, V16.B16
	SUB	$16, R1
	CMP	$16, R1
	BGE	iba_loop16

iba_tail1:
	MOVD	$0, R5
iba_tail1_loop:
	CBZ	R1, iba_miss
	MOVBU.P	1(R0), R3
	CMPW	R2, R3
	BEQ	iba_hit_done
	ORR	R3, R5, R5
	SUB	$1, R1
	B	iba_tail1_loop

iba_hit64:
	SUB	$64, R0
	MOVD	$0, R5
iba_hit64_scan:
	MOVBU.P	1(R0), R3
	CMPW	R2, R3
	BEQ	iba_hit_done
	ORR	R3, R5, R5
	B	iba_hit64_scan

iba_hit16:
	SUB	$16, R0
	MOVD	$0, R5
iba_hit16_scan:
	MOVBU.P	1(R0), R3
	CMPW	R2, R3
	BEQ	iba_hit_done
	ORR	R3, R5, R5
	B	iba_hit16_scan

iba_hit_done:
	SUB	$1, R0
	SUB	R10, R0, R0
	MOVD	R0, idx+24(FP)
	VMOV	V16.D[0], R3
	VMOV	V16.D[1], R4
	ORR	R3, R4, R3
	MOVD	$0x8080808080808080, R9
	TST	R9, R3
	BNE	iba_non_ascii
	TST	$0x80, R5
	BNE	iba_non_ascii
	MOVD	$1, R3
	MOVB	R3, prefixASCII+32(FP)
	RET

iba_non_ascii:
	MOVB	ZR, prefixASCII+32(FP)
	RET

iba_miss:
	MOVD	$-1, R0
	MOVD	R0, idx+24(FP)
	MOVB	ZR, prefixASCII+32(FP)
	RET

// func firstNonASCIIAsm(ptr *byte, n int) int
// Returns the byte offset of the first byte >= 0x80, or n if all bytes are ASCII.
TEXT ·firstNonASCIIAsm(SB), NOSPLIT, $0-24
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1

	MOVD	R0, R10
	MOVD	$0x8080808080808080, R9

	CMP	$64, R1
	BLT	fna_check16

fna_loop64:
	VLD1.P	64(R0), [V1.B16, V2.B16, V3.B16, V4.B16]
	VORR	V2.B16, V1.B16, V1.B16
	VORR	V4.B16, V3.B16, V3.B16
	VORR	V3.B16, V1.B16, V1.B16
	VMOV	V1.D[0], R3
	VMOV	V1.D[1], R4
	ORR	R3, R4, R3
	TST	R9, R3
	BNE	fna_hit64
	SUB	$64, R1
	CMP	$64, R1
	BGE	fna_loop64

fna_check16:
	CMP	$16, R1
	BLT	fna_tail1

fna_loop16:
	VLD1.P	16(R0), [V1.B16]
	VMOV	V1.D[0], R3
	VMOV	V1.D[1], R4
	ORR	R3, R4, R3
	TST	R9, R3
	BNE	fna_hit16
	SUB	$16, R1
	CMP	$16, R1
	BGE	fna_loop16

fna_tail1:
	CBZ	R1, fna_all_ascii
	MOVBU.P	1(R0), R3
	TST	$0x80, R3
	BNE	fna_found
	SUB	$1, R1
	B	fna_tail1

fna_hit64:
	SUB	$64, R0
fna_scan:
	MOVBU.P	1(R0), R3
	TST	$0x80, R3
	BEQ	fna_scan
	B	fna_found

fna_hit16:
	SUB	$16, R0
	B	fna_scan

fna_found:
	SUB	$1, R0
	SUB	R10, R0, R0
	MOVD	R0, ret+16(FP)
	RET

fna_all_ascii:
	SUB	R10, R0, R0
	MOVD	R0, ret+16(FP)
	RET
