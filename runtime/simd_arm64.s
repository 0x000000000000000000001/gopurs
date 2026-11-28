//go:build arm64

#include "textflag.h"

// func sumEvenOddInt32Asm(ptr *int32, n int) (even int32, odd int32)
// Accumulates wrapping int32 sums of even-indexed and odd-indexed elements
// using 128-bit NEON vector additions (32 int32s / 128 B per loop iteration).
TEXT ·sumEvenOddInt32Asm(SB), NOSPLIT, $0-24
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1

	VEOR	V0.B16, V0.B16, V0.B16
	SUBS	$32, R1, R1
	BLT	restore32

	VEOR	V1.B16, V1.B16, V1.B16
	VEOR	V2.B16, V2.B16, V2.B16
	VEOR	V3.B16, V3.B16, V3.B16
	VEOR	V16.B16, V16.B16, V16.B16
	VEOR	V17.B16, V17.B16, V17.B16
	VEOR	V18.B16, V18.B16, V18.B16
	VEOR	V19.B16, V19.B16, V19.B16

loop32:
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
	BGE	loop32

	VADD	V16.S4, V0.S4, V0.S4
	VADD	V17.S4, V1.S4, V1.S4
	VADD	V18.S4, V2.S4, V2.S4
	VADD	V19.S4, V3.S4, V3.S4
	VADD	V1.S4, V0.S4, V0.S4
	VADD	V3.S4, V2.S4, V2.S4
	VADD	V2.S4, V0.S4, V0.S4

restore32:
	ADD	$32, R1
	CMP	$16, R1
	BLT	tail4_check

	VLD1.P	64(R0), [V4.S4, V5.S4, V6.S4, V7.S4]
	VADD	V4.S4, V0.S4, V0.S4
	VADD	V5.S4, V6.S4, V5.S4
	VADD	V7.S4, V0.S4, V0.S4
	VADD	V5.S4, V0.S4, V0.S4
	SUB	$16, R1

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

// func containsInt32Asm(ptr *int32, n int, target int32) bool
// Returns true iff target occurs in ptr[0:n].
TEXT ·containsInt32Asm(SB), NOSPLIT, $0-25
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1
	MOVW	target+16(FP), R2

	VDUP	R2, V0.S4

	CMP	$16, R1
	BLT	c_tail4_check

c_loop16:
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
	CBNZ	R3, c_true
	SUB	$16, R1
	CMP	$16, R1
	BGE	c_loop16

c_tail4_check:
	CMP	$4, R1
	BLT	c_tail1

c_loop4:
	VLD1.P	16(R0), [V1.S4]
	VCMEQ	V0.S4, V1.S4, V1.S4
	VMOV	V1.D[0], R3
	VMOV	V1.D[1], R4
	ORR	R3, R4, R3
	CBNZ	R3, c_true
	SUB	$4, R1
	CMP	$4, R1
	BGE	c_loop4

c_tail1:
	CBZ	R1, c_false
	MOVW.P	4(R0), R3
	CMPW	R2, R3
	BEQ	c_true
	SUB	$1, R1
	B	c_tail1

c_true:
	MOVD	$1, R3
	MOVB	R3, ret+24(FP)
	RET

c_false:
	MOVB	ZR, ret+24(FP)
	RET

// Contiguous 144-byte constant block for 16-byte NEON RFC 3629 UTF-8 validation
// & SIMD code-point summation, loaded with 3 multi-register VLD1 instructions:
//   +0:   0xBF   +16:  0xDF   +32:  0xEF   +48:  0xC0   +64:  0x0F
//   +80:  p1_hi  +96:  p1_lo  +112: cur_hi +128: pmask
// Combined 3-VTBL RFC 3629 + surrogate rejection tables on (prev1, cur):
// Bit 0x02: E0 + 80..9F   Bit 0x04: ED + A0..BF (WTF-8 surrogates)
// Bit 0x08: F0 + 80..8F   Bit 0x10: F4 + 90..BF
// Bit 0x20: C0..C1 + any  Bit 0x40: F5..FF + any
// Payload bitmasks indexed by (cur >> 4):
// 0..7: 0x7F, 8..B: 0x3F, C..D: 0x1F, E: 0x0F, F: 0x07
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
DATA utf8_consts<>+128(SB)/8, $0x7f7f7f7f7f7f7f7f
DATA utf8_consts<>+136(SB)/8, $0x070f1f1f3f3f3f3f
GLOBL utf8_consts<>(SB), RODATA, $144

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

	// Chunk B: cur = V0 (becomes next prev), cur_hi = V26 (becomes next prev_hi), prev = V1, prev_hi = V10
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

// func sumCodePointsLimitAsm(ptr *byte, n int, limit int64) (sum int64, consumed int, count int64)
// Validates RFC 3629 UTF-8 and sums code points in parallel across 16-byte NEON
// chunks, with a 64-byte pure-ASCII fast path.
TEXT ·sumCodePointsLimitAsm(SB), NOSPLIT, $0-48
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1
	MOVD	limit+16(FP), R2

	MOVD	$0, R3		// total sum
	MOVD	$0, R4		// total consumed
	MOVD	$0, R9		// total count

	MOVD	R2, R5
	CMP	R1, R5
	BLE	u8s_chk_ascii
	MOVD	R1, R5
u8s_chk_ascii:
	CMP	$64, R5
	BLT	u8s_utf8_setup

	MOVD	$0x8080808080808080, R12
	VLD1	(R0), [V1.B16, V2.B16, V3.B16, V4.B16]
	VORR	V2.B16, V1.B16, V5.B16
	VORR	V4.B16, V3.B16, V6.B16
	VORR	V6.B16, V5.B16, V5.B16
	VMOV	V5.D[0], R7
	VMOV	V5.D[1], R8
	ORR	R7, R8, R7
	TST	R12, R7
	BNE	u8s_utf8_setup

u8s_ascii_block:
	SUB	R9, R2, R5
	CMP	R1, R5
	BLE	u8s_ascii_lim
	MOVD	R1, R5
u8s_ascii_lim:
	LSR	$6, R5, R5
	CBZ	R5, u8s_utf8_setup
	CMP	$64, R5
	BLE	u8s_ascii_start
	MOVD	$64, R5

u8s_ascii_start:
	MOVD	R0, R10
	MOVD	R5, R6
	VEOR	V29.B16, V29.B16, V29.B16
	VEOR	V28.B16, V28.B16, V28.B16

u8s_ascii_loop:
	VLD1.P	64(R0), [V1.B16, V2.B16, V3.B16, V4.B16]
	VORR	V2.B16, V1.B16, V5.B16
	VORR	V4.B16, V3.B16, V6.B16
	VORR	V6.B16, V5.B16, V5.B16
	VORR	V5.B16, V28.B16, V28.B16
	VADD	V2.B16, V1.B16, V1.B16
	VADD	V4.B16, V3.B16, V3.B16
	VUADDW	V1.B8, V29.H8, V29.H8
	VUADDW2	V1.B16, V29.H8, V29.H8
	VUADDW	V3.B8, V29.H8, V29.H8
	VUADDW2	V3.B16, V29.H8, V29.H8
	SUBS	$1, R5, R5
	BNE	u8s_ascii_loop

	VMOV	V28.D[0], R7
	VMOV	V28.D[1], R8
	ORR	R7, R8, R7
	TST	R12, R7
	BNE	u8s_ascii_bail

	VUADDLV	V29.H8, V5
	VMOV	V5.S[0], R7
	ADD	R7, R3, R3
	LSL	$6, R6, R7
	ADD	R7, R4, R4
	ADD	R7, R9, R9
	SUB	R7, R1, R1
	B	u8s_ascii_block

u8s_ascii_bail:
	MOVD	R10, R0

u8s_utf8_setup:
	CMP	$16, R1
	BLT	u8s_done
	SUB	R9, R2, R5
	CMP	$16, R5
	BLT	u8s_done

	MOVD	$utf8_consts<>(SB), R5
	VLD1.P	64(R5), [V16.B16, V17.B16, V18.B16, V19.B16]
	VLD1.P	16(R5), [V20.B16]
	VLD1	(R5), [V22.B16, V23.B16, V24.B16, V25.B16]

	VEOR	V0.B16, V0.B16, V0.B16		// prev = 0
	VEOR	V26.B16, V26.B16, V26.B16	// prev_hi = 0
	VEOR	V15.B16, V15.B16, V15.B16	// prev_c4 = 0
	VEOR	V14.B16, V14.B16, V14.B16	// prev_m4096 = 0
	VEOR	V13.B16, V13.B16, V13.B16	// prev_m64 = 0

u8s_block:
	SUB	R9, R2, R5			// remLimit = limit - count
	CMP	R1, R5
	BLE	u8s_use_lim
	MOVD	R1, R5
u8s_use_lim:
	LSR	$5, R5, R6			// pairs of 16 bytes (32 bytes/iter)
	CBZ	R6, u8s_check16

	// 120-pair accumulation window: keep acc0..acc3 & errAcc live across
	// geometric remLimit sub-steps, flushing .H8 sums once per 120 pairs.
	MOVD	$120, R13			// remaining pair budget in this window
	MOVD	R1, R14				// window remaining bytes
	MOVD	$0, R15				// window processed bytes
	MOVD	$0, R12				// window processed code points
	VEOR	V27.B16, V27.B16, V27.B16	// errAcc
	VEOR	V29.B16, V29.B16, V29.B16	// acc0 (.H8, *1)
	VEOR	V30.B16, V30.B16, V30.B16	// acc1 (.H8, *63)
	VEOR	V31.B16, V31.B16, V31.B16	// acc2 (.H8, *4032)
	VEOR	V21.B16, V21.B16, V21.B16	// acc3 (.H8, *258048)

u8s_sub_start:
	CMP	R13, R6
	BLE	u8s_sub_go
	MOVD	R13, R6
u8s_sub_go:
	MOVD	R6, R5
	VEOR	V28.B16, V28.B16, V28.B16	// contAcc (.B16, <= 240)

u8s_loop32:
	// Chunk A: cur = V1, state from (V0, V26, V15, V14, V13) -> (V1, V10, V8, V7, V6)
	VLD1.P	16(R0), [V1.B16]
	VUSHR	$4, V1.B16, V10.B16

	VEXT	$15, V1.B16, V0.B16, V2.B16
	VEXT	$15, V10.B16, V26.B16, V11.B16

	VCMGT	V1.B16, V19.B16, V5.B16
	VSUB	V5.B16, V28.B16, V28.B16

	// Shared-shift telescoping masks:
	//   c4       = (cur > 0xEF)           = m_262144
	//   m4096    = (cur > 0xDF) | p1(c4)
	//   m64      = (cur > 0xBF) | p1(m4096)
	//   expected = p1(m64)
	VCMHI	V16.B16, V1.B16, V6.B16
	VCMHI	V17.B16, V1.B16, V7.B16
	VCMHI	V18.B16, V1.B16, V8.B16
	VEXT	$15, V8.B16, V15.B16, V9.B16
	VORR	V9.B16, V7.B16, V7.B16
	VEXT	$15, V7.B16, V14.B16, V9.B16
	VORR	V9.B16, V6.B16, V6.B16
	VEXT	$15, V6.B16, V13.B16, V9.B16
	VEOR	V5.B16, V9.B16, V9.B16

	VAND	V20.B16, V2.B16, V12.B16
	VTBL	V11.B16, [V22.B16], V3.B16
	VTBL	V12.B16, [V23.B16], V4.B16
	VTBL	V10.B16, [V24.B16], V5.B16
	VAND	V4.B16, V3.B16, V3.B16
	VAND	V5.B16, V3.B16, V3.B16
	VORR	V3.B16, V9.B16, V9.B16
	VORR	V9.B16, V27.B16, V27.B16

	VTBL	V10.B16, [V25.B16], V2.B16
	VAND	V2.B16, V1.B16, V2.B16		// A_pay
	VAND	V6.B16, V2.B16, V3.B16		// A_p1
	VAND	V7.B16, V2.B16, V4.B16		// A_p2
	VAND	V8.B16, V2.B16, V5.B16		// A_p3

	// Chunk B: cur = V0, state from (V1, V10, V8, V7, V6) -> (V0, V26, V15, V14, V13)
	VLD1.P	16(R0), [V0.B16]
	VUSHR	$4, V0.B16, V26.B16

	VEXT	$15, V0.B16, V1.B16, V12.B16
	VEXT	$15, V26.B16, V10.B16, V11.B16

	VCMGT	V0.B16, V19.B16, V1.B16
	VSUB	V1.B16, V28.B16, V28.B16

	VCMHI	V16.B16, V0.B16, V13.B16
	VCMHI	V17.B16, V0.B16, V14.B16
	VCMHI	V18.B16, V0.B16, V15.B16
	VEXT	$15, V15.B16, V8.B16, V9.B16
	VORR	V9.B16, V14.B16, V14.B16
	VEXT	$15, V14.B16, V7.B16, V9.B16
	VORR	V9.B16, V13.B16, V13.B16
	VEXT	$15, V13.B16, V6.B16, V9.B16
	VEOR	V1.B16, V9.B16, V9.B16

	VAND	V20.B16, V12.B16, V12.B16
	VTBL	V11.B16, [V22.B16], V6.B16
	VTBL	V12.B16, [V23.B16], V7.B16
	VTBL	V26.B16, [V24.B16], V8.B16
	VAND	V7.B16, V6.B16, V6.B16
	VAND	V8.B16, V6.B16, V6.B16
	VORR	V6.B16, V9.B16, V9.B16
	VORR	V9.B16, V27.B16, V27.B16

	VTBL	V26.B16, [V25.B16], V1.B16
	VAND	V1.B16, V0.B16, V1.B16		// B_pay
	VAND	V13.B16, V1.B16, V6.B16		// B_p1
	VAND	V14.B16, V1.B16, V7.B16		// B_p2
	VAND	V15.B16, V1.B16, V8.B16		// B_p3

	// Combine A + B in uint8 (<= 254, 126, 126, 8) then widen once per 32 bytes
	VADD	V1.B16, V2.B16, V2.B16
	VADD	V6.B16, V3.B16, V3.B16
	VADD	V7.B16, V4.B16, V4.B16
	VADD	V8.B16, V5.B16, V5.B16

	VUADDW	V2.B8, V29.H8, V29.H8
	VUADDW2	V2.B16, V29.H8, V29.H8
	VUADDW	V3.B8, V30.H8, V30.H8
	VUADDW2	V3.B16, V30.H8, V30.H8
	VUADDW	V4.B8, V31.H8, V31.H8
	VUADDW2	V4.B16, V31.H8, V31.H8
	VUADDW	V5.B8, V21.H8, V21.H8
	VUADDW2	V5.B16, V21.H8, V21.H8

	SUBS	$1, R5, R5
	BNE	u8s_loop32

	VUADDLV	V28.B16, V5
	VMOV	V5.H[0], R7
	LSL	$5, R6, R8
	SUB	R7, R8, R7
	ADD	R7, R12, R12
	ADD	R8, R15, R15
	SUB	R8, R14, R14
	SUBS	R6, R13, R13
	BEQ	u8s_flush

	ADD	R9, R12, R7
	SUB	R7, R2, R5
	CMP	R14, R5
	BLE	u8s_next_lim
	MOVD	R14, R5
u8s_next_lim:
	LSR	$5, R5, R6
	CBNZ	R6, u8s_sub_start

u8s_flush:
	VMOV	V27.D[0], R7
	VMOV	V27.D[1], R8
	ORR	R7, R8, R7
	CBNZ	R7, u8s_done

	VUADDLV	V29.H8, V5
	VMOV	V5.S[0], R7
	VUADDLV	V30.H8, V6
	VMOV	V6.S[0], R8
	VUADDLV	V31.H8, V7
	VMOV	V7.S[0], R10
	VUADDLV	V21.H8, V8
	VMOV	V8.S[0], R11

	// Telescoping sum: S0 + 63*S1 + 4032*S2 + 258048*S3
	ADD	R7, R3, R3
	SUB	R8, R3, R3
	ADD	R8<<6, R3, R3
	SUB	R10<<6, R3, R3
	ADD	R10<<12, R3, R3
	SUB	R11<<12, R3, R3
	ADD	R11<<18, R3, R3

	ADD	R12, R9, R9
	ADD	R15, R4, R4
	MOVD	R14, R1
	B	u8s_block

u8s_check16:
	CMP	$16, R5
	BLT	u8s_done

	VLD1.P	16(R0), [V1.B16]
	VUSHR	$4, V1.B16, V10.B16

	VEXT	$15, V1.B16, V0.B16, V2.B16
	VEXT	$15, V10.B16, V26.B16, V11.B16

	VCMGT	V1.B16, V19.B16, V5.B16
	VEOR	V28.B16, V28.B16, V28.B16
	VSUB	V5.B16, V28.B16, V28.B16

	VCMHI	V16.B16, V1.B16, V6.B16
	VCMHI	V17.B16, V1.B16, V7.B16
	VCMHI	V18.B16, V1.B16, V8.B16
	VEXT	$15, V8.B16, V15.B16, V9.B16
	VORR	V9.B16, V7.B16, V7.B16
	VEXT	$15, V7.B16, V14.B16, V9.B16
	VORR	V9.B16, V6.B16, V6.B16
	VEXT	$15, V6.B16, V13.B16, V9.B16
	VEOR	V5.B16, V9.B16, V9.B16

	VAND	V20.B16, V2.B16, V12.B16
	VTBL	V11.B16, [V22.B16], V3.B16
	VTBL	V12.B16, [V23.B16], V4.B16
	VTBL	V10.B16, [V24.B16], V5.B16
	VAND	V4.B16, V3.B16, V3.B16
	VAND	V5.B16, V3.B16, V3.B16
	VORR	V3.B16, V9.B16, V9.B16

	VMOV	V9.D[0], R7
	VMOV	V9.D[1], R8
	ORR	R7, R8, R7
	CBNZ	R7, u8s_done

	VTBL	V10.B16, [V25.B16], V2.B16
	VAND	V2.B16, V1.B16, V2.B16
	VAND	V6.B16, V2.B16, V3.B16
	VAND	V7.B16, V2.B16, V4.B16
	VAND	V8.B16, V2.B16, V5.B16

	VUADDLV	V2.B16, V9
	VMOV	V9.H[0], R7
	VUADDLV	V3.B16, V9
	VMOV	V9.H[0], R8
	VUADDLV	V4.B16, V9
	VMOV	V9.H[0], R10
	VUADDLV	V5.B16, V9
	VMOV	V9.H[0], R11

	ADD	R7, R3, R3
	SUB	R8, R3, R3
	ADD	R8<<6, R3, R3
	SUB	R10<<6, R3, R3
	ADD	R10<<12, R3, R3
	SUB	R11<<12, R3, R3
	ADD	R11<<18, R3, R3

	VUADDLV	V28.B16, V5
	VMOV	V5.H[0], R7
	MOVD	$16, R8
	SUB	R7, R8, R7
	ADD	R7, R9, R9
	ADD	$16, R4, R4
	SUB	$16, R1, R1

	VORR	V1.B16, V1.B16, V0.B16
	VORR	V10.B16, V10.B16, V26.B16
	VORR	V8.B16, V8.B16, V15.B16
	VORR	V7.B16, V7.B16, V14.B16
	VORR	V6.B16, V6.B16, V13.B16
	B	u8s_block

u8s_done:
	MOVD	R3, sum+24(FP)
	MOVD	R4, consumed+32(FP)
	MOVD	R9, count+40(FP)
	RET

// 32-entry x 16-byte LUT for branchless SIMD filtered even/odd parity splitting.
// index = (startParity << 4) | 4-bit match mask.
DATA filter_even_lut<>+0(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+8(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+16(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+24(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+32(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+40(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+48(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+56(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+64(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+72(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+80(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+88(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+96(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+104(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+112(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+120(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+128(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+136(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+144(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+152(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+160(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+168(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+176(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+184(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+192(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+200(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+208(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+216(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+224(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+232(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+240(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+248(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+256(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+264(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+272(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+280(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+288(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+296(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+304(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+312(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+320(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+328(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+336(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+344(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+352(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+360(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+368(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+376(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+384(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+392(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+400(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+408(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+416(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+424(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+432(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+440(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+448(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+456(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+464(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+472(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+480(SB)/8, $0x0000000000000000
DATA filter_even_lut<>+488(SB)/8, $0x00000000ffffffff
DATA filter_even_lut<>+496(SB)/8, $0xffffffff00000000
DATA filter_even_lut<>+504(SB)/8, $0xffffffff00000000
GLOBL filter_even_lut<>(SB), RODATA, $512

DATA filter_bitw<>+0(SB)/8, $0x8040201008040201
DATA filter_bitw<>+8(SB)/8, $0x8040201008040201
GLOBL filter_bitw<>(SB), RODATA, $16

// func filterMaskEvenOddInt32Asm(ptr *int32, n int, mask int32, target int32) (evenSum, oddSum, evenCnt, oddCnt int32)
// Hybrid NEON+LUT 16-element/iter branchless filtered even/odd parity sum:
// uses VUZP1+VADDP single-VMOV 16-bit mask packing, .B16 match counting, and
// a single-instruction VCMTST+VBIC fast path whenever target == 0.
TEXT ·filterMaskEvenOddInt32Asm(SB), NOSPLIT, $0-40
	MOVD	ptr+0(FP), R0
	MOVD	n+8(FP), R1
	MOVW	mask+16(FP), R2
	MOVW	target+20(FP), R3

	VDUP	R2, V16.S4
	VDUP	R3, V17.S4
	MOVD	$filter_bitw<>(SB), R4
	VLD1	(R4), [V18.B16]
	MOVD	$filter_even_lut<>(SB), R5
	MOVD	$0x6996, R6
	MOVD	$0, R7				// running parity (0 or 16)

	VEOR	V24.B16, V24.B16, V24.B16	// evenSum (.S4)
	VEOR	V25.B16, V25.B16, V25.B16	// totalSum (.S4)
	MOVD	$0, R19				// totalCnt

	CBNZW	R3, fm_g_block

fm_z_block:
	LSR	$4, R1, R20
	CBZ	R20, fm_tail_reduce
	CMP	$240, R20
	BLE	fm_z_start
	MOVD	$240, R20

fm_z_start:
	MOVD	R20, R21
	LSL	$4, R20, R22
	SUB	R22, R1, R1
	VEOR	V26.B16, V26.B16, V26.B16	// block non-match count (.B16, <= 240)

fm_z_loop16:
	VLD1.P	64(R0), [V0.S4, V1.S4, V2.S4, V3.S4]

	// VCMTST: 0xFFFFFFFF iff (x & mask) != 0 (inverted match mask in 1 op!)
	VCMTST	V16.S4, V0.S4, V4.S4
	VCMTST	V16.S4, V1.S4, V5.S4
	VCMTST	V16.S4, V2.S4, V6.S4
	VCMTST	V16.S4, V3.S4, V7.S4

	// Pack 16 inverted match masks -> V8.B16 and extract 16-bit mask via 1 VMOV
	VUZP1	V5.H8, V4.H8, V8.H8
	VUZP1	V7.H8, V6.H8, V9.H8
	VUZP1	V9.B16, V8.B16, V8.B16
	VSUB	V8.B16, V26.B16, V26.B16
	VAND	V18.B16, V8.B16, V9.B16
	VADDP	V9.B16, V9.B16, V9.B16
	VADDP	V9.B16, V9.B16, V9.B16
	VADDP	V9.B16, V9.B16, V9.B16
	VMOV	V9.H[0], R8

	// Matched values = V_k & ~V_{4+k} via VBIC (zero VMVN!)
	VBIC	V4.B16, V0.B16, V4.B16
	VBIC	V5.B16, V1.B16, V5.B16
	VBIC	V6.B16, V2.B16, V6.B16
	VBIC	V7.B16, V3.B16, V7.B16
	VADD	V4.S4, V5.S4, V10.S4
	VADD	V6.S4, V7.S4, V11.S4
	VADD	V10.S4, V11.S4, V10.S4
	VADD	V10.S4, V25.S4, V25.S4

	// Invert 16-bit mask on GP side and compute 4 LUT addresses + running parity
	EOR	$0xFFFF, R8, R8
	LSR	$4, R8, R9
	LSR	$8, R8, R10
	LSR	$12, R8, R11
	AND	$15, R8, R8
	AND	$15, R9, R9
	AND	$15, R10, R10

	LSR	R8, R6, R12
	LSR	R9, R6, R13
	LSR	R10, R6, R14
	LSR	R11, R6, R15
	AND	$1, R12
	AND	$1, R13
	AND	$1, R14
	AND	$1, R15
	LSL	$4, R12
	LSL	$4, R13
	LSL	$4, R14
	LSL	$4, R15

	EOR	R13, R12, R16
	EOR	R15, R14, R17
	EOR	R12, R7, R12
	EOR	R16, R7, R13
	EOR	R14, R13, R14
	ADD	R7, R8, R8
	EOR	R17, R13, R7
	ADD	R12, R9, R9
	ADD	R13, R10, R10
	ADD	R14, R11, R11

	ADD	R8<<4, R5, R8
	ADD	R9<<4, R5, R9
	ADD	R10<<4, R5, R10
	ADD	R11<<4, R5, R11
	VLD1	(R8), [V12.S4]
	VLD1	(R9), [V13.S4]
	VLD1	(R10), [V14.S4]
	VLD1	(R11), [V15.S4]

	VAND	V12.B16, V4.B16, V12.B16
	VAND	V13.B16, V5.B16, V13.B16
	VAND	V14.B16, V6.B16, V14.B16
	VAND	V15.B16, V7.B16, V15.B16
	VADD	V12.S4, V13.S4, V12.S4
	VADD	V14.S4, V15.S4, V14.S4
	VADD	V12.S4, V14.S4, V12.S4
	VADD	V12.S4, V24.S4, V24.S4

	SUBS	$1, R21, R21
	BNE	fm_z_loop16

	VUADDLV	V26.B16, V0
	VMOV	V0.H[0], R8
	SUB	R8, R22, R8
	ADD	R8, R19, R19
	B	fm_z_block

fm_g_block:
	LSR	$4, R1, R20
	CBZ	R20, fm_tail_reduce
	CMP	$240, R20
	BLE	fm_g_start
	MOVD	$240, R20

fm_g_start:
	MOVD	R20, R21
	LSL	$4, R20, R22
	SUB	R22, R1, R1
	VEOR	V26.B16, V26.B16, V26.B16	// block match count (.B16, <= 240)

fm_g_loop16:
	VLD1.P	64(R0), [V0.S4, V1.S4, V2.S4, V3.S4]

	VAND	V16.B16, V0.B16, V4.B16
	VAND	V16.B16, V1.B16, V5.B16
	VAND	V16.B16, V2.B16, V6.B16
	VAND	V16.B16, V3.B16, V7.B16
	VCMEQ	V17.S4, V4.S4, V4.S4
	VCMEQ	V17.S4, V5.S4, V5.S4
	VCMEQ	V17.S4, V6.S4, V6.S4
	VCMEQ	V17.S4, V7.S4, V7.S4

	VUZP1	V5.H8, V4.H8, V8.H8
	VUZP1	V7.H8, V6.H8, V9.H8
	VUZP1	V9.B16, V8.B16, V8.B16
	VSUB	V8.B16, V26.B16, V26.B16
	VAND	V18.B16, V8.B16, V9.B16
	VADDP	V9.B16, V9.B16, V9.B16
	VADDP	V9.B16, V9.B16, V9.B16
	VADDP	V9.B16, V9.B16, V9.B16
	VMOV	V9.H[0], R8

	VAND	V4.B16, V0.B16, V4.B16
	VAND	V5.B16, V1.B16, V5.B16
	VAND	V6.B16, V2.B16, V6.B16
	VAND	V7.B16, V3.B16, V7.B16
	VADD	V4.S4, V5.S4, V10.S4
	VADD	V6.S4, V7.S4, V11.S4
	VADD	V10.S4, V11.S4, V10.S4
	VADD	V10.S4, V25.S4, V25.S4

	LSR	$4, R8, R9
	LSR	$8, R8, R10
	LSR	$12, R8, R11
	AND	$15, R8, R8
	AND	$15, R9, R9
	AND	$15, R10, R10

	LSR	R8, R6, R12
	LSR	R9, R6, R13
	LSR	R10, R6, R14
	LSR	R11, R6, R15
	AND	$1, R12
	AND	$1, R13
	AND	$1, R14
	AND	$1, R15
	LSL	$4, R12
	LSL	$4, R13
	LSL	$4, R14
	LSL	$4, R15

	EOR	R13, R12, R16
	EOR	R15, R14, R17
	EOR	R12, R7, R12
	EOR	R16, R7, R13
	EOR	R14, R13, R14
	ADD	R7, R8, R8
	EOR	R17, R13, R7
	ADD	R12, R9, R9
	ADD	R13, R10, R10
	ADD	R14, R11, R11

	ADD	R8<<4, R5, R8
	ADD	R9<<4, R5, R9
	ADD	R10<<4, R5, R10
	ADD	R11<<4, R5, R11
	VLD1	(R8), [V12.S4]
	VLD1	(R9), [V13.S4]
	VLD1	(R10), [V14.S4]
	VLD1	(R11), [V15.S4]

	VAND	V12.B16, V4.B16, V12.B16
	VAND	V13.B16, V5.B16, V13.B16
	VAND	V14.B16, V6.B16, V14.B16
	VAND	V15.B16, V7.B16, V15.B16
	VADD	V12.S4, V13.S4, V12.S4
	VADD	V14.S4, V15.S4, V14.S4
	VADD	V12.S4, V14.S4, V12.S4
	VADD	V12.S4, V24.S4, V24.S4

	SUBS	$1, R21, R21
	BNE	fm_g_loop16

	VUADDLV	V26.B16, V0
	VMOV	V0.H[0], R8
	ADD	R8, R19, R19
	B	fm_g_block

fm_tail_reduce:
	VADDV	V24.S4, V24
	VADDV	V25.S4, V25
	VMOV	V24.S[0], R8
	VMOV	V25.S[0], R9
	SUBW	R8, R9, R9
	MOVD	R19, R10
	LSRW	$1, R10, R11
	ADDW	$1, R10, R10
	LSRW	$1, R10, R10

fm_tail1:
	CBZ	R1, fm_done
	MOVW.P	4(R0), R12
	ANDW	R2, R12, R13
	CMPW	R3, R13
	BNE	fm_next
	CBNZ	R7, fm_odd
	ADDW	R12, R8, R8
	ADDW	$1, R10, R10
	MOVD	$16, R7
	B	fm_next
fm_odd:
	ADDW	R12, R9, R9
	ADDW	$1, R11, R11
	MOVD	$0, R7
fm_next:
	SUB	$1, R1
	B	fm_tail1

fm_done:
	MOVW	R8, evenSum+24(FP)
	MOVW	R9, oddSum+28(FP)
	MOVW	R10, evenCnt+32(FP)
	MOVW	R11, oddCnt+36(FP)
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
