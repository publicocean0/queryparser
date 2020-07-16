// Code generated by goyacc -o queryparser.go -p Query queryparser.y. DO NOT EDIT.

// queryparser.y:5

package parser

import __yyfmt__ "fmt"

// queryparser.y:6

import (
	"github.com/publicocean0/queryparser/common"
	"regexp"
	"strconv"
	"time"
)

// queryparser.y:23
type QuerySymType struct {
	yys   int
	expr  common.Expression
	cond  common.Condition
	token common.Token
}

const IDENT = 57346
const STRING = 57347
const INT = 57348
const FLOAT = 57349
const BOOL = 57350
const EQ = 57351
const NEQ = 57352
const LT = 57353
const GT = 57354
const LTE = 57355
const GTE = 57356
const LPAREN = 57357
const RPAREN = 57358
const AND = 57359
const NOT = 57360
const OR = 57361
const LIKE = 57362
const NLIKE = 57363
const DATE = 57364
const DATETIME = 57365
const TIME = 57366

var QueryToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"STRING",
	"INT",
	"FLOAT",
	"BOOL",
	"EQ",
	"NEQ",
	"LT",
	"GT",
	"LTE",
	"GTE",
	"LPAREN",
	"RPAREN",
	"AND",
	"NOT",
	"OR",
	"LIKE",
	"NLIKE",
	"DATE",
	"DATETIME",
	"TIME",
}
var QueryStatenames = [...]string{}

const QueryEofCode = 1
const QueryErrCode = 2
const QueryInitialStackSize = 16

// queryparser.y:299
/*  start  of  programs  */

type QueryLexerImpl struct {
	src       string
	pos       int
	re        *regexp.Regexp
	Exception *common.Exception
	AST       common.Expression
}

func (l *QueryLexerImpl) Init(src string) {
	l.src = src
	l.pos = 0
	l.re = regexp.MustCompile(`^((?P<op>((OR)|(NOT)|(AND)))|(?P<comp>((\!\=)|(\!\~)|(\<\=)|(\>\=)|(\=)|(\~)|(\>)|(\<)))|(?P<bool>((true)|(false)))|(?P<string>\"(?:[^"\\]|\\.)*\")|(?P<ident>[A-Za-z]\w*)|(?P<time>((\d{4}\-\d{2}\-\d{2})(T(\d{1,2}\:\d{2}((\:\d{2}(\:\d{2})?)?)(([-+]\d{2}\:\d{2})|Z)?))?)|(\d{1,2}\:\d{2}((\:\d{2}(\:\d{2})?)?)(([-+]\d{2}\:\d{2})|Z)?))|(?P<number>([-+]?\d+)(\.\d+)?))`)
}

func (l *QueryLexerImpl) Lex(lval *QuerySymType) int {
	var t int = -1
	len := len(l.src)
	// remove all spaces on the left
	for l.pos < len && (l.src[l.pos] == ' ' || l.src[l.pos] == '\t') {
		l.pos++

	}
	if l.pos == len {
		return -1
	}

	if l.src[l.pos] == '(' {
		t = LPAREN
		lval.token = common.Token{Token: t, Literal: "("}
		l.pos++
		return t
	} else if l.src[l.pos] == ')' {
		t = RPAREN
		lval.token = common.Token{Token: t, Literal: ")"}
		l.pos++
		return t
	}
	//find the leftmost token (and its subtokens)
	result := l.re.FindSubmatchIndex([]byte(l.src[l.pos:]))
	if result == nil {
		l.Exception = &common.Exception{}
		l.Exception.Init(l.pos, "invalid syntax at "+l.src[l.pos:])
		return -1
	}
	for pairIndex := 2; t == -1 && pairIndex < 88; pairIndex += 2 {

		rstart := result[pairIndex]
		if rstart != -1 {
			start := l.pos + result[pairIndex]
			switch pairIndex {
			// comparator
			case 18:
				t = NEQ
				lval.token = common.Token{Token: t, Literal: "!="}
				break
			case 20:
				t = NLIKE
				lval.token = common.Token{Token: t, Literal: "!~"}
				break
			case 22:
				t = LTE
				lval.token = common.Token{Token: t, Literal: "<="}
				break
			case 24:
				t = GTE
				lval.token = common.Token{Token: t, Literal: ">="}
				break
			case 26:
				t = EQ
				lval.token = common.Token{Token: t, Literal: "="}
				break
			case 28:
				t = LIKE
				lval.token = common.Token{Token: t, Literal: "~"}
				break
			case 30:
				t = GT
				lval.token = common.Token{Token: t, Literal: ">"}
				break
			case 32:
				t = LT
				lval.token = common.Token{Token: t, Literal: "<"}
				break
			case 62:
				if result[66] != -1 {
					t = FLOAT
				} else {
					t = INT
				}
				lval.token = common.Token{Token: t, Literal: l.src[start : l.pos+result[pairIndex+1]]}
				break
				// string
			case 42: //SHIFT
				t = STRING
				lval.token = common.Token{Token: t, Literal: l.src[start+1 : l.pos+result[pairIndex+1]-1]}
				break
			case 8:
				t = OR
				lval.token = common.Token{Token: t, Literal: "OR"}
				break
			case 10:
				t = NOT
				lval.token = common.Token{Token: t, Literal: "NOT"}
				break
			case 12:
				t = AND
				lval.token = common.Token{Token: t, Literal: "AND"}
				break
			case 36:
				t = BOOL
				lval.token = common.Token{Token: t, Literal: l.src[start : l.pos+result[pairIndex+1]]}
				break
			case 44: // gia' sfhittato
				t = IDENT
				lval.token = common.Token{Token: t, Literal: l.src[start : l.pos+result[pairIndex+1]]}
				break
			case 66:
				t = TIME
				s := ""
				if result[70] == -1 {
					s += ":00"
					if result[74] != -1 {
						s += l.src[l.pos+result[74] : l.pos+result[75]]
						lval.token = common.Token{Token: t, Literal: l.src[start:l.pos+result[74]] + s}
					}
				}
				if result[74] == -1 {
					s += "+00:00"
				}
				lval.token = common.Token{Token: t, Literal: l.src[start:l.pos+result[pairIndex+1]] + s}
				break
			case 48:
				s := ""
				if result[52] != -1 {
					t = DATETIME

					if result[58] == -1 {
						s = ":00"
						if result[62] != -1 {
							s += l.src[l.pos+result[62] : l.pos+result[63]]
							lval.token = common.Token{Token: t, Literal: l.src[start:l.pos+result[62]] + s}
						}
					}
					if result[62] == -1 {
						s += "+00:00"
						lval.token = common.Token{Token: t, Literal: l.src[start:l.pos+result[pairIndex+1]] + s}
					}

				} else {
					t = DATE
					s += "T00:00:00+00:00"
					lval.token = common.Token{Token: t, Literal: l.src[start:l.pos+result[pairIndex+1]] + s}
				}

				break

			}
			if t != -1 {
				l.pos += result[pairIndex+1]
				if l.pos != len && l.src[l.pos] != ' ' {
					t = -1
					l.pos -= result[pairIndex+1]
					//l.Error("invalid token at "+l.src[l.pos:])
					return t
				}
			}

		}

	}

	return t
}

func (l *QueryLexerImpl) Error(e string) {
	l.Exception = &common.Exception{}
	l.Exception.Init(l.pos, e+" at "+l.src[l.pos:])
}

// yacctab:1
var QueryExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const QueryPrivate = 57344

const QueryLast = 65

var QueryAct = [...]int{

	28, 30, 31, 29, 21, 23, 24, 22, 52, 53,
	7, 47, 48, 2, 42, 43, 36, 32, 33, 34,
	35, 25, 26, 27, 54, 55, 56, 49, 50, 51,
	44, 45, 46, 37, 38, 10, 11, 15, 14, 17,
	16, 20, 6, 5, 7, 0, 12, 13, 0, 39,
	40, 41, 0, 6, 3, 7, 1, 4, 0, 0,
	8, 9, 0, 18, 19,
}
var QueryPact = [...]int{

	39, 36, -1000, 39, 39, 26, 39, 39, 25, -1000,
	-1, -5, 15, 11, 27, 8, 5, 2, -9, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var QueryPgo = [...]int{

	0, 56, 13,
}
var QueryR1 = [...]int{

	0, 1, 1, 1, 1, 1, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2,
}
var QueryR2 = [...]int{

	0, 1, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3,
}
var QueryChk = [...]int{

	-1000, -1, -2, 15, 18, 4, 17, 19, -1, -1,
	9, 10, 20, 21, 12, 11, 14, 13, -1, -1,
	16, 5, 8, 6, 7, 22, 23, 24, 5, 8,
	6, 7, 22, 23, 24, 5, 5, 6, 7, 22,
	23, 24, 6, 7, 22, 23, 24, 6, 7, 22,
	23, 24, 6, 7, 22, 23, 24,
}
var QueryDef = [...]int{

	0, -2, 1, 0, 0, 0, 0, 0, 0, 3,
	0, 0, 0, 0, 0, 0, 0, 0, 4, 5,
	2, 6, 10, 12, 14, 24, 25, 26, 7, 11,
	13, 15, 27, 28, 29, 8, 9, 16, 18, 33,
	34, 35, 17, 19, 30, 31, 32, 20, 22, 39,
	40, 41, 21, 23, 36, 37, 38,
}
var QueryTok1 = [...]int{

	1,
}
var QueryTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24,
}
var QueryTok3 = [...]int{
	0,
}

var QueryErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

// yaccpar:1

/*	parser for yacc output	*/

var (
	QueryDebug        = 0
	QueryErrorVerbose = false
)

type QueryLexer interface {
	Lex(lval *QuerySymType) int
	Error(s string)
}

type QueryParser interface {
	Parse(QueryLexer) int
	Lookahead() int
}

type QueryParserImpl struct {
	lval  QuerySymType
	stack [QueryInitialStackSize]QuerySymType
	char  int
}

func (p *QueryParserImpl) Lookahead() int {
	return p.char
}

func QueryNewParser() QueryParser {
	return &QueryParserImpl{}
}

const QueryFlag = -1000

func QueryTokname(c int) string {
	if c >= 1 && c-1 < len(QueryToknames) {
		if QueryToknames[c-1] != "" {
			return QueryToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func QueryStatname(s int) string {
	if s >= 0 && s < len(QueryStatenames) {
		if QueryStatenames[s] != "" {
			return QueryStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func QueryErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !QueryErrorVerbose {
		return "syntax error"
	}

	for _, e := range QueryErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + QueryTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := QueryPact[state]
	for tok := TOKSTART; tok-1 < len(QueryToknames); tok++ {
		if n := base + tok; n >= 0 && n < QueryLast && QueryChk[QueryAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if QueryDef[state] == -2 {
		i := 0
		for QueryExca[i] != -1 || QueryExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; QueryExca[i] >= 0; i += 2 {
			tok := QueryExca[i]
			if tok < TOKSTART || QueryExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if QueryExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += QueryTokname(tok)
	}
	return res
}

func Querylex1(lex QueryLexer, lval *QuerySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = QueryTok1[0]
		goto out
	}
	if char < len(QueryTok1) {
		token = QueryTok1[char]
		goto out
	}
	if char >= QueryPrivate {
		if char < QueryPrivate+len(QueryTok2) {
			token = QueryTok2[char-QueryPrivate]
			goto out
		}
	}
	for i := 0; i < len(QueryTok3); i += 2 {
		token = QueryTok3[i+0]
		if token == char {
			token = QueryTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = QueryTok2[1] /* unknown char */
	}
	if QueryDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", QueryTokname(token), uint(char))
	}
	return char, token
}

func QueryParse(Querylex QueryLexer) int {
	return QueryNewParser().Parse(Querylex)
}

func (Queryrcvr *QueryParserImpl) Parse(Querylex QueryLexer) int {
	var Queryn int
	var QueryVAL QuerySymType
	var QueryDollar []QuerySymType
	_ = QueryDollar // silence set and not used
	QueryS := Queryrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Querystate := 0
	Queryrcvr.char = -1
	Querytoken := -1 // Queryrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Querystate = -1
		Queryrcvr.char = -1
		Querytoken = -1
	}()
	Queryp := -1
	goto Querystack

ret0:
	return 0

ret1:
	return 1

Querystack:
	/* put a state and value onto the stack */
	if QueryDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", QueryTokname(Querytoken), QueryStatname(Querystate))
	}

	Queryp++
	if Queryp >= len(QueryS) {
		nyys := make([]QuerySymType, len(QueryS)*2)
		copy(nyys, QueryS)
		QueryS = nyys
	}
	QueryS[Queryp] = QueryVAL
	QueryS[Queryp].yys = Querystate

Querynewstate:
	Queryn = QueryPact[Querystate]
	if Queryn <= QueryFlag {
		goto Querydefault /* simple state */
	}
	if Queryrcvr.char < 0 {
		Queryrcvr.char, Querytoken = Querylex1(Querylex, &Queryrcvr.lval)
	}
	Queryn += Querytoken
	if Queryn < 0 || Queryn >= QueryLast {
		goto Querydefault
	}
	Queryn = QueryAct[Queryn]
	if QueryChk[Queryn] == Querytoken { /* valid shift */
		Queryrcvr.char = -1
		Querytoken = -1
		QueryVAL = Queryrcvr.lval
		Querystate = Queryn
		if Errflag > 0 {
			Errflag--
		}
		goto Querystack
	}

Querydefault:
	/* default state action */
	Queryn = QueryDef[Querystate]
	if Queryn == -2 {
		if Queryrcvr.char < 0 {
			Queryrcvr.char, Querytoken = Querylex1(Querylex, &Queryrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if QueryExca[xi+0] == -1 && QueryExca[xi+1] == Querystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Queryn = QueryExca[xi+0]
			if Queryn < 0 || Queryn == Querytoken {
				break
			}
		}
		Queryn = QueryExca[xi+1]
		if Queryn < 0 {
			goto ret0
		}
	}
	if Queryn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Querylex.Error(QueryErrorMessage(Querystate, Querytoken))
			Nerrs++
			if QueryDebug >= 1 {
				__yyfmt__.Printf("%s", QueryStatname(Querystate))
				__yyfmt__.Printf(" saw %s\n", QueryTokname(Querytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Queryp >= 0 {
				Queryn = QueryPact[QueryS[Queryp].yys] + QueryErrCode
				if Queryn >= 0 && Queryn < QueryLast {
					Querystate = QueryAct[Queryn] /* simulate a shift of "error" */
					if QueryChk[Querystate] == QueryErrCode {
						goto Querystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if QueryDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", QueryS[Queryp].yys)
				}
				Queryp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if QueryDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", QueryTokname(Querytoken))
			}
			if Querytoken == QueryEofCode {
				goto ret1
			}
			Queryrcvr.char = -1
			Querytoken = -1
			goto Querynewstate /* try again in the same state */
		}
	}

	/* reduction by production Queryn */
	if QueryDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Queryn, QueryStatname(Querystate))
	}

	Querynt := Queryn
	Querypt := Queryp
	_ = Querypt // guard against "declared and not used"

	Queryp -= QueryR2[Queryn]
	// Queryp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Queryp+1 >= len(QueryS) {
		nyys := make([]QuerySymType, len(QueryS)*2)
		copy(nyys, QueryS)
		QueryS = nyys
	}
	QueryVAL = QueryS[Queryp+1]

	/* consult goto table to find next state */
	Queryn = QueryR1[Queryn]
	Queryg := QueryPgo[Queryn]
	Queryj := Queryg + QueryS[Queryp].yys + 1

	if Queryj >= QueryLast {
		Querystate = QueryAct[Queryg]
	} else {
		Querystate = QueryAct[Queryj]
		if QueryChk[Querystate] != -Queryn {
			Querystate = QueryAct[Queryg]
		}
	}
	// dummy call; replaced with literal code
	switch Querynt {

	case 1:
		QueryDollar = QueryS[Querypt-1 : Querypt+1]
		// queryparser.y:62
		{
			QueryVAL.expr = QueryDollar[1].cond
			Querylex.(*QueryLexerImpl).AST = QueryVAL.expr

		}
	case 2:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:68
		{
			QueryVAL.expr = common.SubExpression{Expr: QueryDollar[2].expr}
			Querylex.(*QueryLexerImpl).AST = QueryVAL.expr

		}
	case 3:
		QueryDollar = QueryS[Querypt-2 : Querypt+1]
		// queryparser.y:73
		{
			QueryVAL.expr = common.NotExpression{Expr: QueryDollar[2].expr}
			Querylex.(*QueryLexerImpl).AST = QueryVAL.expr
		}
	case 4:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:77
		{
			QueryVAL.expr = common.BiExpression{BooleanOperator: QueryDollar[2].token, Left: QueryDollar[1].expr, Right: QueryDollar[3].expr}
			Querylex.(*QueryLexerImpl).AST = QueryVAL.expr

		}
	case 5:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:82
		{
			QueryVAL.expr = common.BiExpression{BooleanOperator: QueryDollar[2].token, Left: QueryDollar[1].expr, Right: QueryDollar[3].expr}
			Querylex.(*QueryLexerImpl).AST = QueryVAL.expr
		}
	case 6:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:90
		{
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: QueryDollar[3].token.Literal}}
		}
	case 7:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:94
		{
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: QueryDollar[3].token.Literal}}
		}
	case 8:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:98
		{
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: QueryDollar[3].token.Literal}}
		}
	case 9:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:102
		{
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: QueryDollar[3].token.Literal}}
		}
	case 10:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:106
		{
			b, _ := strconv.ParseBool(QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: b}}
		}
	case 11:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:111
		{
			b, _ := strconv.ParseBool(QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: b}}
		}
	case 12:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:116
		{
			i, _ := strconv.ParseInt(QueryDollar[3].token.Literal, 10, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: i}}

		}
	case 13:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:122
		{
			i, _ := strconv.ParseInt(QueryDollar[3].token.Literal, 10, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: i}}

		}
	case 14:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:128
		{
			f, _ := strconv.ParseFloat(QueryDollar[3].token.Literal, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: f}}

		}
	case 15:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:134
		{
			f, _ := strconv.ParseFloat(QueryDollar[3].token.Literal, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: f}}

		}
	case 16:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:140
		{
			i, _ := strconv.ParseInt(QueryDollar[3].token.Literal, 10, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: i}}

		}
	case 17:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:146
		{
			i, _ := strconv.ParseInt(QueryDollar[3].token.Literal, 10, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: i}}

		}
	case 18:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:152
		{
			f, _ := strconv.ParseFloat(QueryDollar[3].token.Literal, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: f}}

		}
	case 19:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:158
		{
			f, _ := strconv.ParseFloat(QueryDollar[3].token.Literal, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: f}}

		}
	case 20:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:164
		{
			i, _ := strconv.ParseInt(QueryDollar[3].token.Literal, 10, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: i}}

		}
	case 21:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:170
		{
			i, _ := strconv.ParseInt(QueryDollar[3].token.Literal, 10, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: i}}

		}
	case 22:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:176
		{
			f, _ := strconv.ParseFloat(QueryDollar[3].token.Literal, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: f}}

		}
	case 23:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:182
		{
			f, _ := strconv.ParseFloat(QueryDollar[3].token.Literal, 64)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: f}}

		}
	case 24:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:188
		{
			t, _ := time.Parse(time.RFC3339, QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 25:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:194
		{
			t, _ := time.Parse(time.RFC3339, QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 26:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:200
		{
			t, _ := time.Parse(time.RFC3339, "1970-01-01T"+QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 27:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:206
		{
			t, _ := time.Parse("1970-01-01", QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 28:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:212
		{
			t, _ := time.Parse(time.RFC3339, QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 29:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:218
		{
			t, _ := time.Parse(time.RFC3339, "1970-01-01T"+QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 30:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:224
		{
			t, _ := time.Parse("1970-01-01", QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 31:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:230
		{
			t, _ := time.Parse(time.RFC3339, QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 32:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:236
		{
			t, _ := time.Parse(time.RFC3339, "1970-01-01T"+QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 33:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:242
		{
			t, _ := time.Parse("1970-01-01", QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 34:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:248
		{
			t, _ := time.Parse(time.RFC3339, QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 35:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:254
		{
			t, _ := time.Parse(time.RFC3339, "1970-01-01T"+QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 36:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:260
		{
			t, _ := time.Parse("1970-01-01", QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 37:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:266
		{
			t, _ := time.Parse(time.RFC3339, QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 38:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:272
		{
			t, _ := time.Parse(time.RFC3339, "1970-01-01T"+QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 39:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:278
		{
			t, _ := time.Parse("1970-01-01", QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 40:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:284
		{
			t, _ := time.Parse(time.RFC3339, QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	case 41:
		QueryDollar = QueryS[Querypt-3 : Querypt+1]
		// queryparser.y:290
		{
			t, _ := time.Parse(time.RFC3339, "1970-01-01T"+QueryDollar[3].token.Literal)
			QueryVAL.cond = common.Condition{Variable: QueryDollar[1].token, Comparator: QueryDollar[2].token, Value: common.TokenValue{Token: QueryDollar[3].token.Token, Content: t}}

		}
	}
	goto Querystack /* stack new state and value */
}
