#include "textflag.h"

// func Add(a, b int) int
TEXT ·Add(SB),NOSPLIT,$0-24
	MOVQ a+0(FP), AX
	MOVQ b+8(FP), CX
	ADDQ CX, AX
	MOVQ AX, ret+16(FP)
	RET
