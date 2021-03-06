// Code generated by goyacc - DO NOT EDIT.

package eval

import __yyfmt__ "fmt"

type yySymType struct {
	yys   int
	value uint64
	ast   *ast
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57347
	yyEofCode = 57344
	NUMBER    = 57346
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -10
)

var (
	yyPrec = map[int]int{
		'+': 0,
		'-': 0,
		'*': 1,
		'/': 1,
		'%': 1,
		'^': 1,
	}

	yyXLAT = map[int]int{
		57344: 0,  // $end (10x)
		37:    1,  // '%' (10x)
		42:    2,  // '*' (10x)
		43:    3,  // '+' (10x)
		45:    4,  // '-' (10x)
		47:    5,  // '/' (10x)
		94:    6,  // '^' (10x)
		41:    7,  // ')' (9x)
		40:    8,  // '(' (8x)
		57348: 9,  // expression (8x)
		57346: 10, // NUMBER (8x)
		57349: 11, // top (1x)
		57347: 12, // $default (0x)
		57345: 13, // error (0x)
	}

	yySymNames = []string{
		"$end",
		"'%'",
		"'*'",
		"'+'",
		"'-'",
		"'/'",
		"'^'",
		"')'",
		"'('",
		"expression",
		"NUMBER",
		"top",
		"$default",
		"error",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0: {0, 1},
		1: {11, 1},
		2: {9, 3},
		3: {9, 3},
		4: {9, 3},
		5: {9, 3},
		6: {9, 3},
		7: {9, 3},
		8: {9, 3},
		9: {9, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [19][]uint8{
		// 0
		{8: 13, 12, 14, 11},
		{10},
		{9, 21, 19, 17, 18, 20, 22},
		{8: 13, 15, 14},
		{1, 1, 1, 1, 1, 1, 1, 1},
		// 5
		{1: 21, 19, 17, 18, 20, 22, 16},
		{8, 8, 8, 8, 8, 8, 8, 8},
		{8: 13, 28, 14},
		{8: 13, 27, 14},
		{8: 13, 26, 14},
		// 10
		{8: 13, 25, 14},
		{8: 13, 24, 14},
		{8: 13, 23, 14},
		{2, 2, 2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3, 3, 3},
		// 15
		{4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5},
		{6, 21, 19, 6, 6, 20, 22, 6},
		{7, 21, 19, 7, 7, 20, 22, 7},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 13

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			if l, ok := yylex.(*simpleLex); ok {
				l.ast = yyS[yypt-0].ast
			}
		}
	case 2:
		{
			yyVAL.ast = yyS[yypt-1].ast
		}
	case 3:
		{
			yyVAL.ast = genOpNode('+', yyS[yypt-2].ast, yyS[yypt-0].ast)
		}
	case 4:
		{
			yyVAL.ast = genOpNode('-', yyS[yypt-2].ast, yyS[yypt-0].ast)
		}
	case 5:
		{
			yyVAL.ast = genOpNode('*', yyS[yypt-2].ast, yyS[yypt-0].ast)
		}
	case 6:
		{
			yyVAL.ast = genOpNode('/', yyS[yypt-2].ast, yyS[yypt-0].ast)
		}
	case 7:
		{
			yyVAL.ast = genOpNode('%', yyS[yypt-2].ast, yyS[yypt-0].ast)
		}
	case 8:
		{
			yyVAL.ast = genOpNode('^', yyS[yypt-2].ast, yyS[yypt-0].ast)
		}
	case 9:
		{
			yyVAL.ast = genValueNode(yyS[yypt-0].value)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
