package dft

func DFT12(ri, ii, ro, io []float64) {
	T1 := ri[0]
	T2 := ri[4]
	T3 := ri[8]
	T4 := T2 + T3
	T5 := T1 + T4
	T53 := T1 - (KP500000000 * T4)
	T36 := KP866025403 * (T3 - T2)
	T24 := ii[0]
	T25 := ii[4]
	T26 := ii[8]
	T27 := T25 + T26
	T28 := T24 + T27
	T54 := KP866025403 * (T25 - T26)
	T35 := T24 - (KP500000000 * T27)
	T6 := ri[6]
	T7 := ri[10]
	T8 := ri[2]
	T9 := T7 + T8
	T10 := T6 + T9
	T56 := T6 - (KP500000000 * T9)
	T39 := KP866025403 * (T8 - T7)
	T29 := ii[6]
	T30 := ii[10]
	T31 := ii[2]
	T32 := T30 + T31
	T33 := T29 + T32
	T57 := KP866025403 * (T30 - T31)
	T38 := T29 - (KP500000000 * T32)
	T12 := ri[3]
	T13 := ri[7]
	T14 := ri[11]
	T15 := T13 + T14
	T16 := T12 + T15
	T72 := KP866025403 * (T14 - T13)
	T42 := T12 - (KP500000000 * T15)
	T73 := ii[3]
	T43 := ii[7]
	T44 := ii[11]
	T74 := T43 + T44
	T45 := KP866025403 * (T43 - T44)
	T92 := T73 + T74
	T75 := T73 - (KP500000000 * T74)
	T17 := ri[9]
	T18 := ri[1]
	T19 := ri[5]
	T20 := T18 + T19
	T21 := T17 + T20
	T77 := KP866025403 * (T19 - T18)
	T47 := T17 - (KP500000000 * T20)
	T78 := ii[9]
	T48 := ii[1]
	T49 := ii[5]
	T79 := T48 + T49
	T50 := KP866025403 * (T48 - T49)
	T93 := T78 + T79
	T80 := T78 - (KP500000000 * T79)
	T11 := T5 + T10
	T22 := T16 + T21
	ro[6] = T11 - T22
	ro[0] = T11 + T22
	T95 := T28 + T33
	T96 := T92 + T93
	io[6] = T95 - T96
	io[0] = T95 + T96
	T23 := T16 - T21
	T34 := T28 - T33
	io[3] = T23 + T34
	io[9] = T34 - T23
	T91 := T5 - T10
	T94 := T92 - T93
	ro[3] = T91 - T94
	ro[9] = T91 + T94
	T61 := T36 + T35
	T62 := T39 + T38
	T63 := T61 - T62
	T83 := T61 + T62
	T76 := T72 + T75
	T81 := T77 + T80
	T82 := T76 - T81
	T84 := T76 + T81
	T64 := T42 + T45
	T65 := T47 + T50
	T66 := T64 - T65
	T70 := T64 + T65
	T67 := T53 + T54
	T68 := T56 + T57
	T69 := T67 + T68
	T71 := T67 - T68
	io[1] = T63 - T66
	ro[1] = T71 + T82
	io[7] = T63 + T66
	ro[7] = T71 - T82
	ro[10] = T69 - T70
	io[10] = T83 - T84
	ro[4] = T69 + T70
	io[4] = T83 + T84
	T37 := T35 - T36
	T40 := T38 - T39
	T41 := T37 - T40
	T89 := T37 + T40
	T86 := T75 - T72
	T87 := T80 - T77
	T88 := T86 - T87
	T90 := T86 + T87
	T46 := T42 - T45
	T51 := T47 - T50
	T52 := T46 - T51
	T60 := T46 + T51
	T55 := T53 - T54
	T58 := T56 - T57
	T59 := T55 + T58
	T85 := T55 - T58
	io[5] = T41 - T52
	ro[5] = T85 + T88
	io[11] = T41 + T52
	ro[11] = T85 - T88
	ro[2] = T59 - T60
	io[2] = T89 - T90
	ro[8] = T59 + T60
	io[8] = T89 + T90
}
