╭─yuchanns@yuchannsdeMacBook-Pro.local ~
╰─➤  go build -gcflags="-N -l -S" range_clause.go
# command-line-arguments
"".RangeClause STEXT size=842 args=0x0 locals=0x158
        0x0000 00000 (range_clause.go:5)     TEXT    "".RangeClause(SB), ABIInternal, $344-0
        0x0000 00000 (range_clause.go:5)     MOVQ    (TLS), CX
        0x0009 00009 (range_clause.go:5)     LEAQ    -216(SP), AX
        0x0011 00017 (range_clause.go:5)     CMPQ    AX, 16(CX)
        0x0015 00021 (range_clause.go:5)     JLS     832
        0x001b 00027 (range_clause.go:5)     SUBQ    $344, SP
        0x0022 00034 (range_clause.go:5)     MOVQ    BP, 336(SP)
        0x002a 00042 (range_clause.go:5)     LEAQ    336(SP), BP
        0x0032 00050 (range_clause.go:5)     FUNCDATA        $0, gclocals·f0a67958015464e4cc8847ce0df60843(SB)
        0x0032 00050 (range_clause.go:5)     FUNCDATA        $1, gclocals·1be50b3ff1c6bce621b19ced5cafc212(SB)
        0x0032 00050 (range_clause.go:5)     FUNCDATA        $2, gclocals·160a1dd0c9595e8bcf8efc4c6b948d91(SB)
        0x0032 00050 (range_clause.go:5)     FUNCDATA        $3, "".RangeClause.stkobj(SB)
        0x0032 00050 (range_clause.go:6)     PCDATA  $0, $1
        0x0032 00050 (range_clause.go:6)     PCDATA  $1, $0
        0x0032 00050 (range_clause.go:6)     LEAQ    ""..autotmp_9+120(SP), AX
        0x0037 00055 (range_clause.go:6)     PCDATA  $1, $1
        0x0037 00055 (range_clause.go:6)     MOVQ    AX, ""..autotmp_8+152(SP)
        0x003f 00063 (range_clause.go:6)     PCDATA  $0, $0
        0x003f 00063 (range_clause.go:6)     TESTB   AL, (AX)
        0x0041 00065 (range_clause.go:6)     MOVQ    ""..stmp_0(SB), AX
        0x0048 00072 (range_clause.go:6)     MOVQ    AX, ""..autotmp_9+120(SP)
        0x004d 00077 (range_clause.go:6)     MOVUPS  ""..stmp_0+8(SB), X0
        0x0054 00084 (range_clause.go:6)     MOVUPS  X0, ""..autotmp_9+128(SP)
        0x005c 00092 (range_clause.go:6)     PCDATA  $0, $1
        0x005c 00092 (range_clause.go:6)     PCDATA  $1, $0
        0x005c 00092 (range_clause.go:6)     MOVQ    ""..autotmp_8+152(SP), AX
        0x0064 00100 (range_clause.go:6)     TESTB   AL, (AX)
        0x0066 00102 (range_clause.go:6)     JMP     104
        0x0068 00104 (range_clause.go:6)     PCDATA  $0, $0
        0x0068 00104 (range_clause.go:6)     PCDATA  $1, $2
        0x0068 00104 (range_clause.go:6)     MOVQ    AX, "".arr+240(SP)
        0x0070 00112 (range_clause.go:6)     MOVQ    $3, "".arr+248(SP)
        0x007c 00124 (range_clause.go:6)     MOVQ    $3, "".arr+256(SP)
        0x0088 00136 (range_clause.go:7)     PCDATA  $1, $3
        0x0088 00136 (range_clause.go:7)     MOVQ    $0, "".newArr+216(SP)
        0x0094 00148 (range_clause.go:7)     XORPS   X0, X0
        0x0097 00151 (range_clause.go:7)     MOVUPS  X0, "".newArr+224(SP)
        0x009f 00159 (range_clause.go:8)     PCDATA  $0, $1
        0x009f 00159 (range_clause.go:8)     LEAQ    type.int(SB), AX
        0x00a6 00166 (range_clause.go:8)     PCDATA  $0, $0
        0x00a6 00166 (range_clause.go:8)     MOVQ    AX, (SP)
        0x00aa 00170 (range_clause.go:8)     CALL    runtime.newobject(SB)
        0x00af 00175 (range_clause.go:8)     PCDATA  $0, $1
        0x00af 00175 (range_clause.go:8)     MOVQ    8(SP), AX
        0x00b4 00180 (range_clause.go:8)     PCDATA  $0, $0
        0x00b4 00180 (range_clause.go:8)     PCDATA  $1, $4
        0x00b4 00180 (range_clause.go:8)     MOVQ    AX, "".&v+192(SP)
        0x00bc 00188 (range_clause.go:8)     MOVQ    "".arr+256(SP), AX
        0x00c4 00196 (range_clause.go:8)     MOVQ    "".arr+248(SP), CX
        0x00cc 00204 (range_clause.go:8)     PCDATA  $0, $2
        0x00cc 00204 (range_clause.go:8)     PCDATA  $1, $5
        0x00cc 00204 (range_clause.go:8)     MOVQ    "".arr+240(SP), DX
        0x00d4 00212 (range_clause.go:8)     PCDATA  $0, $0
        0x00d4 00212 (range_clause.go:8)     PCDATA  $1, $6
        0x00d4 00212 (range_clause.go:8)     MOVQ    DX, ""..autotmp_5+288(SP)
        0x00dc 00220 (range_clause.go:8)     MOVQ    CX, ""..autotmp_5+296(SP)
        0x00e4 00228 (range_clause.go:8)     MOVQ    AX, ""..autotmp_5+304(SP)
        0x00ec 00236 (range_clause.go:8)     MOVQ    $0, ""..autotmp_10+112(SP)
        0x00f5 00245 (range_clause.go:8)     MOVQ    ""..autotmp_5+296(SP), AX
        0x00fd 00253 (range_clause.go:8)     MOVQ    AX, ""..autotmp_11+104(SP)
        0x0102 00258 (range_clause.go:8)     JMP     260
        0x0104 00260 (range_clause.go:8)     MOVQ    ""..autotmp_11+104(SP), CX
        0x0109 00265 (range_clause.go:8)     CMPQ    ""..autotmp_10+112(SP), CX
        0x010e 00270 (range_clause.go:8)     JLT     277
        0x0110 00272 (range_clause.go:8)     JMP     516
        0x0115 00277 (range_clause.go:8)     MOVQ    ""..autotmp_10+112(SP), CX
        0x011a 00282 (range_clause.go:8)     SHLQ    $3, CX
        0x011e 00286 (range_clause.go:8)     PCDATA  $0, $3
        0x011e 00286 (range_clause.go:8)     ADDQ    ""..autotmp_5+288(SP), CX
        0x0126 00294 (range_clause.go:8)     PCDATA  $0, $0
        0x0126 00294 (range_clause.go:8)     MOVQ    (CX), CX
        0x0129 00297 (range_clause.go:8)     MOVQ    CX, ""..autotmp_12+96(SP)
        0x012e 00302 (range_clause.go:8)     PCDATA  $0, $2
        0x012e 00302 (range_clause.go:8)     MOVQ    "".&v+192(SP), DX
        0x0136 00310 (range_clause.go:8)     PCDATA  $0, $0
        0x0136 00310 (range_clause.go:8)     MOVQ    CX, (DX)
        0x0139 00313 (range_clause.go:9)     PCDATA  $0, $3
        0x0139 00313 (range_clause.go:9)     MOVQ    "".&v+192(SP), CX
        0x0141 00321 (range_clause.go:9)     PCDATA  $0, $0
        0x0141 00321 (range_clause.go:9)     PCDATA  $1, $7
        0x0141 00321 (range_clause.go:9)     MOVQ    CX, ""..autotmp_13+184(SP)
        0x0149 00329 (range_clause.go:9)     MOVQ    "".newArr+232(SP), CX
        0x0151 00337 (range_clause.go:9)     MOVQ    "".newArr+224(SP), DX
        0x0159 00345 (range_clause.go:9)     PCDATA  $0, $4
        0x0159 00345 (range_clause.go:9)     PCDATA  $1, $8
        0x0159 00345 (range_clause.go:9)     MOVQ    "".newArr+216(SP), BX
        0x0161 00353 (range_clause.go:9)     LEAQ    1(DX), SI
        0x0165 00357 (range_clause.go:9)     CMPQ    SI, CX
        0x0168 00360 (range_clause.go:9)     JLS     364
        0x016a 00362 (range_clause.go:9)     JMP     446
        0x016c 00364 (range_clause.go:9)     PCDATA  $0, $-2
        0x016c 00364 (range_clause.go:9)     PCDATA  $1, $-2
        0x016c 00364 (range_clause.go:9)     JMP     366
        0x016e 00366 (range_clause.go:9)     PCDATA  $0, $5
        0x016e 00366 (range_clause.go:9)     PCDATA  $1, $9
        0x016e 00366 (range_clause.go:9)     MOVQ    ""..autotmp_13+184(SP), AX
        0x0176 00374 (range_clause.go:9)     PCDATA  $0, $6
        0x0176 00374 (range_clause.go:9)     LEAQ    (BX)(DX*8), DI
        0x017a 00378 (range_clause.go:9)     PCDATA  $0, $-2
        0x017a 00378 (range_clause.go:9)     PCDATA  $1, $-2
        0x017a 00378 (range_clause.go:9)     CMPL    runtime.writeBarrier(SB), $0
        0x0181 00385 (range_clause.go:9)     JEQ     389
        0x0183 00387 (range_clause.go:9)     JMP     439
        0x0185 00389 (range_clause.go:9)     MOVQ    AX, (BX)(DX*8)
        0x0189 00393 (range_clause.go:9)     JMP     395
        0x018b 00395 (range_clause.go:9)     PCDATA  $0, $0
        0x018b 00395 (range_clause.go:9)     PCDATA  $1, $6
        0x018b 00395 (range_clause.go:9)     MOVQ    BX, "".newArr+216(SP)
        0x0193 00403 (range_clause.go:9)     MOVQ    SI, "".newArr+224(SP)
        0x019b 00411 (range_clause.go:9)     MOVQ    CX, "".newArr+232(SP)
        0x01a3 00419 (range_clause.go:9)     JMP     421
        0x01a5 00421 (range_clause.go:8)     MOVQ    ""..autotmp_10+112(SP), CX
        0x01aa 00426 (range_clause.go:8)     INCQ    CX
        0x01ad 00429 (range_clause.go:8)     MOVQ    CX, ""..autotmp_10+112(SP)
        0x01b2 00434 (range_clause.go:8)     JMP     260
        0x01b7 00439 (range_clause.go:9)     PCDATA  $0, $-2
        0x01b7 00439 (range_clause.go:9)     PCDATA  $1, $-2
        0x01b7 00439 (range_clause.go:9)     CALL    runtime.gcWriteBarrier(SB)
        0x01bc 00444 (range_clause.go:9)     JMP     395
        0x01be 00446 (range_clause.go:9)     PCDATA  $0, $4
        0x01be 00446 (range_clause.go:9)     PCDATA  $1, $8
        0x01be 00446 (range_clause.go:9)     MOVQ    DX, ""..autotmp_21+64(SP)
        0x01c3 00451 (range_clause.go:9)     PCDATA  $0, $5
        0x01c3 00451 (range_clause.go:9)     LEAQ    type.*int(SB), AX
        0x01ca 00458 (range_clause.go:9)     PCDATA  $0, $4
        0x01ca 00458 (range_clause.go:9)     MOVQ    AX, (SP)
        0x01ce 00462 (range_clause.go:9)     PCDATA  $0, $0
        0x01ce 00462 (range_clause.go:9)     MOVQ    BX, 8(SP)
        0x01d3 00467 (range_clause.go:9)     MOVQ    DX, 16(SP)
        0x01d8 00472 (range_clause.go:9)     MOVQ    CX, 24(SP)
        0x01dd 00477 (range_clause.go:9)     MOVQ    SI, 32(SP)
        0x01e2 00482 (range_clause.go:9)     CALL    runtime.growslice(SB)
        0x01e7 00487 (range_clause.go:9)     PCDATA  $0, $4
        0x01e7 00487 (range_clause.go:9)     MOVQ    40(SP), BX
        0x01ec 00492 (range_clause.go:9)     MOVQ    48(SP), AX
        0x01f1 00497 (range_clause.go:9)     MOVQ    56(SP), CX
        0x01f6 00502 (range_clause.go:9)     LEAQ    1(AX), SI
        0x01fa 00506 (range_clause.go:9)     MOVQ    ""..autotmp_21+64(SP), DX
        0x01ff 00511 (range_clause.go:9)     JMP     366
        0x0204 00516 (range_clause.go:11)    PCDATA  $0, $0
        0x0204 00516 (range_clause.go:11)    PCDATA  $1, $10
        0x0204 00516 (range_clause.go:11)    MOVQ    "".newArr+232(SP), AX
        0x020c 00524 (range_clause.go:11)    MOVQ    "".newArr+224(SP), CX
        0x0214 00532 (range_clause.go:11)    PCDATA  $0, $2
        0x0214 00532 (range_clause.go:11)    PCDATA  $1, $0
        0x0214 00532 (range_clause.go:11)    MOVQ    "".newArr+216(SP), DX
        0x021c 00540 (range_clause.go:11)    PCDATA  $0, $0
        0x021c 00540 (range_clause.go:11)    PCDATA  $1, $11
        0x021c 00540 (range_clause.go:11)    MOVQ    DX, ""..autotmp_6+264(SP)
        0x0224 00548 (range_clause.go:11)    MOVQ    CX, ""..autotmp_6+272(SP)
        0x022c 00556 (range_clause.go:11)    MOVQ    AX, ""..autotmp_6+280(SP)
        0x0234 00564 (range_clause.go:11)    MOVQ    $0, ""..autotmp_14+88(SP)
        0x023d 00573 (range_clause.go:11)    MOVQ    ""..autotmp_6+272(SP), AX
        0x0245 00581 (range_clause.go:11)    MOVQ    AX, ""..autotmp_15+80(SP)
        0x024a 00586 (range_clause.go:11)    JMP     588
        0x024c 00588 (range_clause.go:11)    MOVQ    ""..autotmp_15+80(SP), AX
        0x0251 00593 (range_clause.go:11)    CMPQ    ""..autotmp_14+88(SP), AX
        0x0256 00598 (range_clause.go:11)    JLT     605
        0x0258 00600 (range_clause.go:11)    JMP     816
        0x025d 00605 (range_clause.go:11)    MOVQ    ""..autotmp_14+88(SP), AX
        0x0262 00610 (range_clause.go:11)    SHLQ    $3, AX
        0x0266 00614 (range_clause.go:11)    PCDATA  $0, $1
        0x0266 00614 (range_clause.go:11)    ADDQ    ""..autotmp_6+264(SP), AX
        0x026e 00622 (range_clause.go:11)    MOVQ    (AX), AX
        0x0271 00625 (range_clause.go:11)    MOVQ    AX, ""..autotmp_16+176(SP)
        0x0279 00633 (range_clause.go:11)    MOVQ    AX, "".v+144(SP)
        0x0281 00641 (range_clause.go:12)    TESTB   AL, (AX)
        0x0283 00643 (range_clause.go:12)    PCDATA  $0, $0
        0x0283 00643 (range_clause.go:12)    MOVQ    (AX), AX
        0x0286 00646 (range_clause.go:12)    MOVQ    AX, ""..autotmp_17+72(SP)
        0x028b 00651 (range_clause.go:12)    MOVQ    AX, (SP)
        0x028f 00655 (range_clause.go:12)    CALL    runtime.convT64(SB)
        0x0294 00660 (range_clause.go:12)    PCDATA  $0, $1
        0x0294 00660 (range_clause.go:12)    MOVQ    8(SP), AX
        0x0299 00665 (range_clause.go:12)    PCDATA  $0, $0
        0x0299 00665 (range_clause.go:12)    PCDATA  $1, $12
        0x0299 00665 (range_clause.go:12)    MOVQ    AX, ""..autotmp_18+168(SP)
        0x02a1 00673 (range_clause.go:12)    PCDATA  $1, $13
        0x02a1 00673 (range_clause.go:12)    XORPS   X0, X0
        0x02a4 00676 (range_clause.go:12)    MOVUPS  X0, ""..autotmp_7+200(SP)
        0x02ac 00684 (range_clause.go:12)    PCDATA  $0, $1
        0x02ac 00684 (range_clause.go:12)    PCDATA  $1, $12
        0x02ac 00684 (range_clause.go:12)    LEAQ    ""..autotmp_7+200(SP), AX
        0x02b4 00692 (range_clause.go:12)    MOVQ    AX, ""..autotmp_20+160(SP)
        0x02bc 00700 (range_clause.go:12)    TESTB   AL, (AX)
        0x02be 00702 (range_clause.go:12)    PCDATA  $0, $7
        0x02be 00702 (range_clause.go:12)    PCDATA  $1, $11
        0x02be 00702 (range_clause.go:12)    MOVQ    ""..autotmp_18+168(SP), CX
        0x02c6 00710 (range_clause.go:12)    PCDATA  $0, $8
        0x02c6 00710 (range_clause.go:12)    LEAQ    type.int(SB), DX
        0x02cd 00717 (range_clause.go:12)    PCDATA  $0, $7
        0x02cd 00717 (range_clause.go:12)    MOVQ    DX, ""..autotmp_7+200(SP)
        0x02d5 00725 (range_clause.go:12)    PCDATA  $0, $1
        0x02d5 00725 (range_clause.go:12)    MOVQ    CX, ""..autotmp_7+208(SP)
        0x02dd 00733 (range_clause.go:12)    TESTB   AL, (AX)
        0x02df 00735 (range_clause.go:12)    JMP     737
        0x02e1 00737 (range_clause.go:12)    MOVQ    AX, ""..autotmp_19+312(SP)
        0x02e9 00745 (range_clause.go:12)    MOVQ    $1, ""..autotmp_19+320(SP)
        0x02f5 00757 (range_clause.go:12)    MOVQ    $1, ""..autotmp_19+328(SP)
        0x0301 00769 (range_clause.go:12)    PCDATA  $0, $0
        0x0301 00769 (range_clause.go:12)    MOVQ    AX, (SP)
        0x0305 00773 (range_clause.go:12)    MOVQ    $1, 8(SP)
        0x030e 00782 (range_clause.go:12)    MOVQ    $1, 16(SP)
        0x0317 00791 (range_clause.go:12)    CALL    fmt.Println(SB)
        0x031c 00796 (range_clause.go:12)    JMP     798
        0x031e 00798 (range_clause.go:11)    MOVQ    ""..autotmp_14+88(SP), AX
        0x0323 00803 (range_clause.go:11)    INCQ    AX
        0x0326 00806 (range_clause.go:11)    MOVQ    AX, ""..autotmp_14+88(SP)
        0x032b 00811 (range_clause.go:11)    JMP     588
        0x0330 00816 (<unknown line number>)    PCDATA  $1, $0
        0x0330 00816 (<unknown line number>)    MOVQ    336(SP), BP
        0x0338 00824 (<unknown line number>)    ADDQ    $344, SP
        0x033f 00831 (<unknown line number>)    RET
        0x0340 00832 (<unknown line number>)    NOP
        0x0340 00832 (range_clause.go:5)     PCDATA  $1, $-1
        0x0340 00832 (range_clause.go:5)     PCDATA  $0, $-1
        0x0340 00832 (range_clause.go:5)     CALL    runtime.morestack_noctxt(SB)
        0x0345 00837 (range_clause.go:5)     JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 84 24 28 ff ff  eH..%....H..$(..
        0x0010 ff 48 3b 41 10 0f 86 25 03 00 00 48 81 ec 58 01  .H;A...%...H..X.
        0x0020 00 00 48 89 ac 24 50 01 00 00 48 8d ac 24 50 01  ..H..$P...H..$P.
        0x0030 00 00 48 8d 44 24 78 48 89 84 24 98 00 00 00 84  ..H.D$xH..$.....
        0x0040 00 48 8b 05 00 00 00 00 48 89 44 24 78 0f 10 05  .H......H.D$x...
        0x0050 00 00 00 00 0f 11 84 24 80 00 00 00 48 8b 84 24  .......$....H..$
        0x0060 98 00 00 00 84 00 eb 00 48 89 84 24 f0 00 00 00  ........H..$....
        0x0070 48 c7 84 24 f8 00 00 00 03 00 00 00 48 c7 84 24  H..$........H..$
        0x0080 00 01 00 00 03 00 00 00 48 c7 84 24 d8 00 00 00  ........H..$....
        0x0090 00 00 00 00 0f 57 c0 0f 11 84 24 e0 00 00 00 48  .....W....$....H
        0x00a0 8d 05 00 00 00 00 48 89 04 24 e8 00 00 00 00 48  ......H..$.....H
        0x00b0 8b 44 24 08 48 89 84 24 c0 00 00 00 48 8b 84 24  .D$.H..$....H..$
        0x00c0 00 01 00 00 48 8b 8c 24 f8 00 00 00 48 8b 94 24  ....H..$....H..$
        0x00d0 f0 00 00 00 48 89 94 24 20 01 00 00 48 89 8c 24  ....H..$ ...H..$
        0x00e0 28 01 00 00 48 89 84 24 30 01 00 00 48 c7 44 24  (...H..$0...H.D$
        0x00f0 70 00 00 00 00 48 8b 84 24 28 01 00 00 48 89 44  p....H..$(...H.D
        0x0100 24 68 eb 00 48 8b 4c 24 68 48 39 4c 24 70 7c 05  $h..H.L$hH9L$p|.
        0x0110 e9 ef 00 00 00 48 8b 4c 24 70 48 c1 e1 03 48 03  .....H.L$pH...H.
        0x0120 8c 24 20 01 00 00 48 8b 09 48 89 4c 24 60 48 8b  .$ ...H..H.L$`H.
        0x0130 94 24 c0 00 00 00 48 89 0a 48 8b 8c 24 c0 00 00  .$....H..H..$...
        0x0140 00 48 89 8c 24 b8 00 00 00 48 8b 8c 24 e8 00 00  .H..$....H..$...
        0x0150 00 48 8b 94 24 e0 00 00 00 48 8b 9c 24 d8 00 00  .H..$....H..$...
        0x0160 00 48 8d 72 01 48 39 ce 76 02 eb 52 eb 00 48 8b  .H.r.H9.v..R..H.
        0x0170 84 24 b8 00 00 00 48 8d 3c d3 83 3d 00 00 00 00  .$....H.<..=....
        0x0180 00 74 02 eb 32 48 89 04 d3 eb 00 48 89 9c 24 d8  .t..2H.....H..$.
        0x0190 00 00 00 48 89 b4 24 e0 00 00 00 48 89 8c 24 e8  ...H..$....H..$.
        0x01a0 00 00 00 eb 00 48 8b 4c 24 70 48 ff c1 48 89 4c  .....H.L$pH..H.L
        0x01b0 24 70 e9 4d ff ff ff e8 00 00 00 00 eb cd 48 89  $p.M..........H.
        0x01c0 54 24 40 48 8d 05 00 00 00 00 48 89 04 24 48 89  T$@H......H..$H.
        0x01d0 5c 24 08 48 89 54 24 10 48 89 4c 24 18 48 89 74  \$.H.T$.H.L$.H.t
        0x01e0 24 20 e8 00 00 00 00 48 8b 5c 24 28 48 8b 44 24  $ .....H.\$(H.D$
        0x01f0 30 48 8b 4c 24 38 48 8d 70 01 48 8b 54 24 40 e9  0H.L$8H.p.H.T$@.
        0x0200 6a ff ff ff 48 8b 84 24 e8 00 00 00 48 8b 8c 24  j...H..$....H..$
        0x0210 e0 00 00 00 48 8b 94 24 d8 00 00 00 48 89 94 24  ....H..$....H..$
        0x0220 08 01 00 00 48 89 8c 24 10 01 00 00 48 89 84 24  ....H..$....H..$
        0x0230 18 01 00 00 48 c7 44 24 58 00 00 00 00 48 8b 84  ....H.D$X....H..
        0x0240 24 10 01 00 00 48 89 44 24 50 eb 00 48 8b 44 24  $....H.D$P..H.D$
        0x0250 50 48 39 44 24 58 7c 05 e9 d3 00 00 00 48 8b 44  PH9D$X|......H.D
        0x0260 24 58 48 c1 e0 03 48 03 84 24 08 01 00 00 48 8b  $XH...H..$....H.
        0x0270 00 48 89 84 24 b0 00 00 00 48 89 84 24 90 00 00  .H..$....H..$...
        0x0280 00 84 00 48 8b 00 48 89 44 24 48 48 89 04 24 e8  ...H..H.D$HH..$.
        0x0290 00 00 00 00 48 8b 44 24 08 48 89 84 24 a8 00 00  ....H.D$.H..$...
        0x02a0 00 0f 57 c0 0f 11 84 24 c8 00 00 00 48 8d 84 24  ..W....$....H..$
        0x02b0 c8 00 00 00 48 89 84 24 a0 00 00 00 84 00 48 8b  ....H..$......H.
        0x02c0 8c 24 a8 00 00 00 48 8d 15 00 00 00 00 48 89 94  .$....H......H..
        0x02d0 24 c8 00 00 00 48 89 8c 24 d0 00 00 00 84 00 eb  $....H..$.......
        0x02e0 00 48 89 84 24 38 01 00 00 48 c7 84 24 40 01 00  .H..$8...H..$@..
        0x02f0 00 01 00 00 00 48 c7 84 24 48 01 00 00 01 00 00  .....H..$H......
        0x0300 00 48 89 04 24 48 c7 44 24 08 01 00 00 00 48 c7  .H..$H.D$.....H.
        0x0310 44 24 10 01 00 00 00 e8 00 00 00 00 eb 00 48 8b  D$............H.
        0x0320 44 24 58 48 ff c0 48 89 44 24 58 e9 1c ff ff ff  D$XH..H.D$X.....
        0x0330 48 8b ac 24 50 01 00 00 48 81 c4 58 01 00 00 c3  H..$P...H..X....
        0x0340 e8 00 00 00 00 e9 b6 fc ff ff                    ..........
        rel 5+4 t=16 TLS+0
        rel 68+4 t=15 ""..stmp_0+0
        rel 80+4 t=15 ""..stmp_0+8
        rel 162+4 t=15 type.int+0
        rel 171+4 t=8 runtime.newobject+0
        rel 380+4 t=15 runtime.writeBarrier+-1
        rel 440+4 t=8 runtime.gcWriteBarrier+0
        rel 454+4 t=15 type.*int+0
        rel 483+4 t=8 runtime.growslice+0
        rel 656+4 t=8 runtime.convT64+0
        rel 713+4 t=15 type.int+0
        rel 792+4 t=8 fmt.Println+0
        rel 833+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.producer.command-line-arguments SDWARFINFO dupok size=0
        0x0000 2d 4e 20 2d 6c                                   -N -l
go.cuinfo.packagename.command-line-arguments SDWARFINFO dupok size=0
        0x0000 61 73 73 65 6d 62 6c 79                          assembly
go.loc."".RangeClause SDWARFLOC size=0
go.info."".RangeClause SDWARFINFO size=132
        0x0000 03 22 22 2e 52 61 6e 67 65 43 6c 61 75 73 65 00  ."".RangeClause.
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0020 01 9c 00 00 00 00 01 0a 6e 65 77 41 72 72 00 07  ........newArr..
        0x0030 00 00 00 00 03 91 f8 7e 0a 61 72 72 00 06 00 00  .......~.arr....
        0x0040 00 00 03 91 90 7f 15 00 00 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 0a 26 76 00 08 00 00 00 00  ........&v......
        0x0060 03 91 e0 7e 00 15 00 00 00 00 00 00 00 00 00 00  ...~............
        0x0070 00 00 00 00 00 00 0a 76 00 0b 00 00 00 00 03 91  .......v........
        0x0080 b0 7e 00 00                                      .~..
        rel 16+8 t=1 "".RangeClause+0
        rel 24+8 t=1 "".RangeClause+842
        rel 34+4 t=29 gofile..range_clause.go+0
        rel 48+4 t=28 go.info.[]*int+0
        rel 62+4 t=28 go.info.[]int+0
        rel 71+8 t=1 "".RangeClause+159
        rel 79+8 t=1 "".RangeClause+516
        rel 92+4 t=28 go.info.*int+0
        rel 102+8 t=1 "".RangeClause+516
        rel 110+8 t=1 "".RangeClause+816
        rel 122+4 t=28 go.info.*int+0
go.range."".RangeClause SDWARFRANGE size=0
go.isstmt."".RangeClause SDWARFMISC size=0
        0x0000 04 1b 04 17 03 05 01 51 02 0c 01 0b 02 07 01 04  .......Q........
        0x0010 02 05 01 55 02 05 01 30 02 08 01 64 02 05 01 38  ...U...0...d...8
        0x0020 02 05 01 1d 02 08 01 40 02 05 01 30 02 02 01 0c  .......@...0....
        0x0030 02 05 01 83 01 02 05 01 02 02 05 01 1d 02 0a 00  ................
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
        0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+80
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
        rel 24+8 t=1 runtime.algarray+144
        rel 32+8 t=1 runtime.gcbits.02+0
        rel 40+4 t=5 type..namedata.*interface {}-+0
        rel 44+4 t=6 type.*interface {}+0
        rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
        0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
        0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+80
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]interface {}-+0
        rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]interface {}-+0
        rel 44+4 t=6 type.*[]interface {}+0
        rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
        0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
        0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+80
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
        rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 01 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+144
        rel 32+8 t=1 runtime.gcbits.02+0
        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
        rel 44+4 t=6 type.*[1]interface {}+0
        rel 48+8 t=1 type.interface {}+0
        rel 56+8 t=1 type.[]interface {}+0
""..inittask SNOPTRDATA size=32
        0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 fmt..inittask+0
""..stmp_0 SRODATA size=24
        0x0000 01 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
        0x0010 03 00 00 00 00 00 00 00                          ........
type..namedata.*[]int- SRODATA dupok size=9
        0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 1b 31 52 88 00 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+80
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]int-+0
        rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]int-+0
        rel 44+4 t=6 type.*[]int+0
        rel 48+8 t=1 type.int+0
type..hashfunc24 SRODATA dupok size=16
        0x0000 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
        rel 0+8 t=1 runtime.memhash_varlen+0
type..eqfunc24 SRODATA dupok size=16
        0x0000 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
        rel 0+8 t=1 runtime.memequal_varlen+0
type..alg24 SRODATA dupok size=16
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 0+8 t=1 type..hashfunc24+0
        rel 8+8 t=1 type..eqfunc24+0
runtime.gcbits. SRODATA dupok size=0
type..namedata.*[3]int- SRODATA dupok size=10
        0x0000 00 00 07 2a 5b 33 5d 69 6e 74                    ...*[3]int
type.[3]int SRODATA dupok size=72
        0x0000 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 63 c4 1a 46 02 08 08 11 00 00 00 00 00 00 00 00  c..F............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 03 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 type..alg24+0
        rel 32+8 t=1 runtime.gcbits.+0
        rel 40+4 t=5 type..namedata.*[3]int-+0
        rel 44+4 t=6 type.*[3]int+0
        rel 48+8 t=1 type.int+0
        rel 56+8 t=1 type.[]int+0
type.*[3]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 25 30 9a e3 00 08 08 36 00 00 00 00 00 00 00 00  %0.....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+80
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[3]int-+0
        rel 48+8 t=1 type.[3]int+0
type..namedata.*[]*int- SRODATA dupok size=10
        0x0000 00 00 07 2a 5b 5d 2a 69 6e 74                    ...*[]*int
type.*[]*int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 18 3d d8 bc 00 08 08 36 00 00 00 00 00 00 00 00  .=.....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+80
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]*int-+0
        rel 48+8 t=1 type.[]*int+0
type.[]*int SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 41 59 d9 aa 02 08 08 17 00 00 00 00 00 00 00 00  AY..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]*int-+0
        rel 44+4 t=6 type.*[]*int+0
        rel 48+8 t=1 type.*int+0
type..importpath.fmt. SRODATA dupok size=6
        0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·f0a67958015464e4cc8847ce0df60843 SRODATA dupok size=8
        0x0000 0e 00 00 00 00 00 00 00                          ........
gclocals·1be50b3ff1c6bce621b19ced5cafc212 SRODATA dupok size=50
        0x0000 0e 00 00 00 18 00 00 00 00 00 00 02 00 00 00 10  ................
        0x0010 00 00 12 00 40 12 00 40 02 00 40 02 04 60 02 04  ....@..@..@..`..
        0x0020 60 00 04 40 00 04 00 02 00 00 80 00 08 80 00 08  `..@............
        0x0030 81 00                                            ..
gclocals·160a1dd0c9595e8bcf8efc4c6b948d91 SRODATA dupok size=17
        0x0000 09 00 00 00 07 00 00 00 00 01 04 02 08 09 49 03  ..............I.
        0x0010 07                                               .
"".RangeClause.stkobj SRODATA dupok size=24
        0x0000 01 00 00 00 00 00 00 00 78 ff ff ff ff ff ff ff  ........x.......
        0x0010 00 00 00 00 00 00 00 00                          ........
        rel 16+8 t=1 type.[1]interface {}+0