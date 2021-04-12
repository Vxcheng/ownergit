"".main STEXT size=53 args=0x0 locals=0x8 funcid=0x0
	0x0000 00000 (debug.go:5)	TEXT	"".main(SB), ABIInternal, $8-0
	0x0000 00000 (debug.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (debug.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (debug.go:5)	PCDATA	$0, $-2
	0x000d 00013 (debug.go:5)	JLS	46
	0x000f 00015 (debug.go:5)	PCDATA	$0, $-1
	0x000f 00015 (debug.go:5)	SUBQ	$8, SP
	0x0013 00019 (debug.go:5)	MOVQ	BP, (SP)
	0x0017 00023 (debug.go:5)	LEAQ	(SP), BP
	0x001b 00027 (debug.go:5)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (debug.go:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (debug.go:6)	PCDATA	$1, $0
	0x001b 00027 (debug.go:6)	NOP
	0x0020 00032 (debug.go:6)	CALL	"".debugG(SB)
	0x0025 00037 (debug.go:7)	MOVQ	(SP), BP
	0x0029 00041 (debug.go:7)	ADDQ	$8, SP
	0x002d 00045 (debug.go:7)	RET
	0x002e 00046 (debug.go:7)	NOP
	0x002e 00046 (debug.go:5)	PCDATA	$1, $-1
	0x002e 00046 (debug.go:5)	PCDATA	$0, $-2
	0x002e 00046 (debug.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x0033 00051 (debug.go:5)	PCDATA	$0, $-1
	0x0033 00051 (debug.go:5)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 1f 48  dH..%....H;a.v.H
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f 1f 44 00 00  ...H.,$H.,$..D..
	0x0020 e8 00 00 00 00 48 8b 2c 24 48 83 c4 08 c3 e8 00  .....H.,$H......
	0x0030 00 00 00 eb cb                                   .....
	rel 5+4 t=17 TLS+0
	rel 33+4 t=8 "".debugG+0
	rel 47+4 t=8 runtime.morestack_noctxt+0
"".debugG STEXT size=207 args=0x0 locals=0x70 funcid=0x0
	0x0000 00000 (debug.go:9)	TEXT	"".debugG(SB), ABIInternal, $112-0
	0x0000 00000 (debug.go:9)	MOVQ	(TLS), CX
	0x0009 00009 (debug.go:9)	CMPQ	SP, 16(CX)
	0x000d 00013 (debug.go:9)	PCDATA	$0, $-2
	0x000d 00013 (debug.go:9)	JLS	197
	0x0013 00019 (debug.go:9)	PCDATA	$0, $-1
	0x0013 00019 (debug.go:9)	SUBQ	$112, SP
	0x0017 00023 (debug.go:9)	MOVQ	BP, 104(SP)
	0x001c 00028 (debug.go:9)	LEAQ	104(SP), BP
	0x0021 00033 (debug.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0021 00033 (debug.go:9)	FUNCDATA	$1, gclocals·ff19ed39bdde8a01a800918ac3ef0ec7(SB)
	0x0021 00033 (debug.go:9)	FUNCDATA	$2, "".debugG.stkobj(SB)
	0x0021 00033 (debug.go:9)	XORL	AX, AX
	0x0023 00035 (debug.go:10)	JMP	177
	0x0028 00040 (debug.go:10)	MOVQ	AX, "".i+64(SP)
	0x002d 00045 (debug.go:11)	MOVQ	AX, (SP)
	0x0031 00049 (debug.go:11)	PCDATA	$1, $0
	0x0031 00049 (debug.go:11)	CALL	runtime.convT64(SB)
	0x0036 00054 (debug.go:11)	MOVQ	8(SP), AX
	0x003b 00059 (debug.go:11)	XORPS	X0, X0
	0x003e 00062 (debug.go:11)	MOVUPS	X0, ""..autotmp_15+72(SP)
	0x0043 00067 (debug.go:11)	MOVUPS	X0, ""..autotmp_15+88(SP)
	0x0048 00072 (debug.go:11)	LEAQ	type.string(SB), CX
	0x004f 00079 (debug.go:11)	MOVQ	CX, ""..autotmp_15+72(SP)
	0x0054 00084 (debug.go:11)	LEAQ	""..stmp_0(SB), DX
	0x005b 00091 (debug.go:11)	MOVQ	DX, ""..autotmp_15+80(SP)
	0x0060 00096 (debug.go:11)	LEAQ	type.int(SB), BX
	0x0067 00103 (debug.go:11)	MOVQ	BX, ""..autotmp_15+88(SP)
	0x006c 00108 (debug.go:11)	MOVQ	AX, ""..autotmp_15+96(SP)
	0x0071 00113 (<unknown line number>)	NOP
	0x0071 00113 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), AX
	0x0078 00120 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), SI
	0x007f 00127 ($GOROOT/src/fmt/print.go:274)	MOVQ	SI, (SP)
	0x0083 00131 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, 8(SP)
	0x0088 00136 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_15+72(SP), AX
	0x008d 00141 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, 16(SP)
	0x0092 00146 ($GOROOT/src/fmt/print.go:274)	MOVQ	$2, 24(SP)
	0x009b 00155 ($GOROOT/src/fmt/print.go:274)	MOVQ	$2, 32(SP)
	0x00a4 00164 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x00a9 00169 (debug.go:10)	MOVQ	"".i+64(SP), AX
	0x00ae 00174 (debug.go:10)	INCQ	AX
	0x00b1 00177 (debug.go:10)	CMPQ	AX, $4
	0x00b5 00181 (debug.go:10)	JLT	40
	0x00bb 00187 (debug.go:10)	PCDATA	$1, $-1
	0x00bb 00187 (debug.go:10)	MOVQ	104(SP), BP
	0x00c0 00192 (debug.go:10)	ADDQ	$112, SP
	0x00c4 00196 (debug.go:10)	RET
	0x00c5 00197 (debug.go:10)	NOP
	0x00c5 00197 (debug.go:9)	PCDATA	$1, $-1
	0x00c5 00197 (debug.go:9)	PCDATA	$0, $-2
	0x00c5 00197 (debug.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x00ca 00202 (debug.go:9)	PCDATA	$0, $-1
	0x00ca 00202 (debug.go:9)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 b2  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 70 48 89 6c 24 68 48 8d 6c 24  ...H..pH.l$hH.l$
	0x0020 68 31 c0 e9 89 00 00 00 48 89 44 24 40 48 89 04  h1......H.D$@H..
	0x0030 24 e8 00 00 00 00 48 8b 44 24 08 0f 57 c0 0f 11  $.....H.D$..W...
	0x0040 44 24 48 0f 11 44 24 58 48 8d 0d 00 00 00 00 48  D$H..D$XH......H
	0x0050 89 4c 24 48 48 8d 15 00 00 00 00 48 89 54 24 50  .L$HH......H.T$P
	0x0060 48 8d 1d 00 00 00 00 48 89 5c 24 58 48 89 44 24  H......H.\$XH.D$
	0x0070 60 48 8b 05 00 00 00 00 48 8d 35 00 00 00 00 48  `H......H.5....H
	0x0080 89 34 24 48 89 44 24 08 48 8d 44 24 48 48 89 44  .4$H.D$.H.D$HH.D
	0x0090 24 10 48 c7 44 24 18 02 00 00 00 48 c7 44 24 20  $.H.D$.....H.D$ 
	0x00a0 02 00 00 00 e8 00 00 00 00 48 8b 44 24 40 48 ff  .........H.D$@H.
	0x00b0 c0 48 83 f8 04 0f 8c 6d ff ff ff 48 8b 6c 24 68  .H.....m...H.l$h
	0x00c0 48 83 c4 70 c3 e8 00 00 00 00 e9 31 ff ff ff     H..p.......1...
	rel 3+0 t=25 type.string+0
	rel 3+0 t=25 type.int+0
	rel 3+0 t=25 type.*os.File+0
	rel 5+4 t=17 TLS+0
	rel 50+4 t=8 runtime.convT64+0
	rel 75+4 t=16 type.string+0
	rel 87+4 t=16 ""..stmp_0+0
	rel 99+4 t=16 type.int+0
	rel 116+4 t=16 os.Stdout+0
	rel 123+4 t=16 go.itab.*os.File,io.Writer+0
	rel 165+4 t=8 fmt.Fprintln+0
	rel 198+4 t=8 runtime.morestack_noctxt+0
os.(*File).close STEXT dupok nosplit size=26 args=0x18 locals=0x0 funcid=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	os.(*File).close(SB), DUPOK|NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (<autogenerated>:1)	FUNCDATA	$0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
	0x0000 00000 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0000 00000 (<autogenerated>:1)	MOVQ	""..this+8(SP), AX
	0x0005 00005 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0008 00008 (<autogenerated>:1)	MOVQ	AX, ""..this+8(SP)
	0x000d 00013 (<autogenerated>:1)	XORPS	X0, X0
	0x0010 00016 (<autogenerated>:1)	MOVUPS	X0, "".~r0+16(SP)
	0x0015 00021 (<autogenerated>:1)	JMP	os.(*file).close(SB)
	0x0000 48 8b 44 24 08 48 8b 00 48 89 44 24 08 0f 57 c0  H.D$.H..H.D$..W.
	0x0010 0f 11 44 24 10 e9 00 00 00 00                    ..D$......
	rel 22+4 t=8 os.(*file).close+0
type..eq.[2]interface {} STEXT dupok size=170 args=0x18 locals=0x30 funcid=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq.[2]interface {}(SB), DUPOK|ABIInternal, $48-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	PCDATA	$0, $-2
	0x000d 00013 (<autogenerated>:1)	JLS	160
	0x0013 00019 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0013 00019 (<autogenerated>:1)	SUBQ	$48, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 40(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	40(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (<autogenerated>:1)	MOVQ	"".q+64(SP), AX
	0x0026 00038 (<autogenerated>:1)	MOVQ	"".p+56(SP), CX
	0x002b 00043 (<autogenerated>:1)	XORL	DX, DX
	0x002d 00045 (<autogenerated>:1)	JMP	66
	0x002f 00047 (<autogenerated>:1)	MOVQ	""..autotmp_6+32(SP), BX
	0x0034 00052 (<autogenerated>:1)	LEAQ	1(BX), DX
	0x0038 00056 (<autogenerated>:1)	MOVQ	"".q+64(SP), AX
	0x003d 00061 (<autogenerated>:1)	MOVQ	"".p+56(SP), CX
	0x0042 00066 (<autogenerated>:1)	CMPQ	DX, $2
	0x0046 00070 (<autogenerated>:1)	JGE	149
	0x0048 00072 (<autogenerated>:1)	MOVQ	DX, BX
	0x004b 00075 (<autogenerated>:1)	SHLQ	$4, DX
	0x004f 00079 (<autogenerated>:1)	MOVQ	(CX)(DX*1), SI
	0x0053 00083 (<autogenerated>:1)	MOVQ	(AX)(DX*1), DI
	0x0057 00087 (<autogenerated>:1)	MOVQ	8(DX)(CX*1), R8
	0x005c 00092 (<autogenerated>:1)	MOVQ	8(DX)(AX*1), DX
	0x0061 00097 (<autogenerated>:1)	CMPQ	DI, SI
	0x0064 00100 (<autogenerated>:1)	JNE	133
	0x0066 00102 (<autogenerated>:1)	MOVQ	BX, ""..autotmp_6+32(SP)
	0x006b 00107 (<autogenerated>:1)	MOVQ	SI, (SP)
	0x006f 00111 (<autogenerated>:1)	MOVQ	R8, 8(SP)
	0x0074 00116 (<autogenerated>:1)	MOVQ	DX, 16(SP)
	0x0079 00121 (<autogenerated>:1)	PCDATA	$1, $0
	0x0079 00121 (<autogenerated>:1)	CALL	runtime.efaceeq(SB)
	0x007e 00126 (<autogenerated>:1)	CMPB	24(SP), $0
	0x0083 00131 (<autogenerated>:1)	JNE	47
	0x0085 00133 (<autogenerated>:1)	XORL	AX, AX
	0x0087 00135 (<autogenerated>:1)	MOVB	AL, "".r+72(SP)
	0x008b 00139 (<autogenerated>:1)	MOVQ	40(SP), BP
	0x0090 00144 (<autogenerated>:1)	ADDQ	$48, SP
	0x0094 00148 (<autogenerated>:1)	RET
	0x0095 00149 (<autogenerated>:1)	MOVL	$1, AX
	0x009a 00154 (<autogenerated>:1)	JMP	135
	0x009c 00156 (<autogenerated>:1)	NOP
	0x009c 00156 (<autogenerated>:1)	PCDATA	$1, $-1
	0x009c 00156 (<autogenerated>:1)	PCDATA	$0, $-2
	0x009c 00156 (<autogenerated>:1)	NOP
	0x00a0 00160 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x00a5 00165 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00a5 00165 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 8d  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
	0x0020 28 48 8b 44 24 40 48 8b 4c 24 38 31 d2 eb 13 48  (H.D$@H.L$81...H
	0x0030 8b 5c 24 20 48 8d 53 01 48 8b 44 24 40 48 8b 4c  .\$ H.S.H.D$@H.L
	0x0040 24 38 48 83 fa 02 7d 4d 48 89 d3 48 c1 e2 04 48  $8H...}MH..H...H
	0x0050 8b 34 11 48 8b 3c 10 4c 8b 44 0a 08 48 8b 54 02  .4.H.<.L.D..H.T.
	0x0060 08 48 39 f7 75 1f 48 89 5c 24 20 48 89 34 24 4c  .H9.u.H.\$ H.4$L
	0x0070 89 44 24 08 48 89 54 24 10 e8 00 00 00 00 80 7c  .D$.H.T$.......|
	0x0080 24 18 00 75 aa 31 c0 88 44 24 48 48 8b 6c 24 28  $..u.1..D$HH.l$(
	0x0090 48 83 c4 30 c3 b8 01 00 00 00 eb eb 0f 1f 40 00  H..0..........@.
	0x00a0 e8 00 00 00 00 e9 56 ff ff ff                    ......V...
	rel 5+4 t=17 TLS+0
	rel 122+4 t=8 runtime.efaceeq+0
	rel 161+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 04 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 11  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 11 6e 00 01 00 00 00 00 11  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=24 type.[]interface {}+0
	rel 0+0 t=24 type.error+0
	rel 0+0 t=24 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
go.string."hello: " SRODATA dupok size=7
	0x0000 68 65 6c 6c 6f 3a 20                             hello: 
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..eqfunc.[2]interface {} SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq.[2]interface {}+0
type..namedata.*[2]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 32 5d 69 6e 74 65 72 66 61 63 65  ...*[2]interface
	0x0010 20 7b 7d                                          {}
type.*[2]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 be 73 2d 71 08 08 08 36 00 00 00 00 00 00 00 00  .s-q...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[2]interface {}-+0
	rel 48+8 t=1 type.[2]interface {}+0
runtime.gcbits.0a SRODATA dupok size=1
	0x0000 0a                                               .
type.[2]interface {} SRODATA dupok size=72
	0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
	0x0010 2c 59 a4 f1 02 08 08 11 00 00 00 00 00 00 00 00  ,Y..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..eqfunc.[2]interface {}+0
	rel 32+8 t=1 runtime.gcbits.0a+0
	rel 40+4 t=5 type..namedata.*[2]interface {}-+0
	rel 44+4 t=6 type.*[2]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
""..stmp_0 SRODATA static size=16
	0x0000 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."hello: "+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=1 os.(*File).Write+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·ff19ed39bdde8a01a800918ac3ef0ec7 SRODATA dupok size=9
	0x0000 01 00 00 00 04 00 00 00 00                       .........
"".debugG.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 e0 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[2]interface {}+0
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
