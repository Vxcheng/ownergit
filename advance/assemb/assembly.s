// assembly.s
TEXT ·assemblyFunction(SB), NOSPLIT, $0
    MOVQ $1, AX
    MOVQ $2, BX
    ADDQ AX, BX
    RET
