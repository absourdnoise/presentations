// https://golang.org/src/sync/atomic/asm_amd64.s 

// func StoreUint32(addr *uint32, val uint32)
TEXT ·StoreUint32(SB),NOSPLIT,$0-12
	MOVQ	addr+0(FP), BP
	MOVL	val+8(FP), AX
	XCHGL	AX, 0(BP)
	RET

// func LoadUint32(addr *uint32) (val uint32)
TEXT ·LoadUint32(SB),NOSPLIT,$0-12
	MOVQ	addr+0(FP), AX
	MOVL	0(AX), AX
	MOVL	AX, val+8(FP)
	RET
