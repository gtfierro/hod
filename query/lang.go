//line lang.y:2
package lang

import __yyfmt__ "fmt"

//line lang.y:2
import (
	"fmt"
	turtle "github.com/gtfierro/hod/goraptor"
	"io"
)

//line lang.y:13
type yySymType struct {
	yys      int
	str      string
	val      turtle.URI
	pred     []PathPattern
	triple   Filter
	triples  []Filter
	varlist  []turtle.URI
	distinct bool
	count    bool
}

const SELECT = 57346
const COUNT = 57347
const DISTINCT = 57348
const WHERE = 57349
const COMMA = 57350
const LBRACE = 57351
const RBRACE = 57352
const DOT = 57353
const SEMICOLON = 57354
const SLASH = 57355
const PLUS = 57356
const QUESTION = 57357
const ASTERISK = 57358
const VAR = 57359
const URI = 57360

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SELECT",
	"COUNT",
	"DISTINCT",
	"WHERE",
	"COMMA",
	"LBRACE",
	"RBRACE",
	"DOT",
	"SEMICOLON",
	"SLASH",
	"PLUS",
	"QUESTION",
	"ASTERISK",
	"VAR",
	"URI",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line lang.y:131

const eof = 0

type lexer struct {
	scanner  *Scanner
	error    error
	varlist  []turtle.URI
	triples  []Filter
	distinct bool
	count    bool
	pos      int
}

func newlexer(r io.Reader) *lexer {
	scanner := NewScanner(
		[]Definition{
			{Token: LBRACE, Pattern: "\\{"},
			{Token: RBRACE, Pattern: "\\}"},
			{Token: COMMA, Pattern: "\\,"},
			{Token: SEMICOLON, Pattern: ";"},
			{Token: DOT, Pattern: "\\."},
			{Token: SELECT, Pattern: "SELECT"},
			{Token: COUNT, Pattern: "COUNT"},
			{Token: DISTINCT, Pattern: "DISTINCT"},
			{Token: WHERE, Pattern: "WHERE"},
			{Token: URI, Pattern: "[a-zA-Z]+:[a-zA-Z0-9_\\-#%$@]+"},
			{Token: VAR, Pattern: "\\?[a-zA-Z0-9_]+"},
			{Token: QUESTION, Pattern: "\\?"},
			{Token: SLASH, Pattern: "/"},
			{Token: PLUS, Pattern: "\\+"},
			{Token: ASTERISK, Pattern: "\\*"},
		})
	scanner.SetInput(r)
	return &lexer{
		scanner: scanner,
		pos:     0,
	}
}

func (l *lexer) Lex(lval *yySymType) int {
	r := l.scanner.Next()
	if r.Token == Error {
		if len(r.Value) > 0 {
			fmt.Println("ERROR", string(r.Value))
		}
		return eof
	}
	lval.str = string(r.Value)
	l.pos += len(r.Value)
	return int(r.Token)
}

func (l *lexer) Error(s string) {
	l.error = fmt.Errorf("Error parsing: %s. Current line %d:%d. Recent token '%s'", s, l.scanner.lineNumber, l.pos, l.scanner.tokenizer.Text())
}

func TokenName(t Token) string {
	return yyTokname(int(t) - 57342)
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 20
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 40

var yyAct = [...]int{

	15, 21, 16, 17, 18, 24, 23, 7, 8, 13,
	17, 18, 29, 30, 31, 28, 26, 25, 8, 37,
	33, 36, 27, 6, 20, 19, 10, 32, 9, 5,
	34, 11, 12, 35, 3, 4, 22, 14, 2, 1,
}
var yyPact = [...]int{

	30, -1000, 22, 1, -9, 17, -1000, -9, -9, -1000,
	-7, -1000, -1000, 15, -7, -12, -14, -1000, -1000, 4,
	-1000, -14, 2, -2, -1000, -12, -1000, 9, -12, -1000,
	-1000, -1000, -14, -1000, -1000, 11, 8, -1000,
}
var yyPgo = [...]int{

	0, 39, 38, 9, 23, 37, 0, 1, 36,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 4, 4, 3, 3, 5,
	5, 7, 7, 8, 8, 8, 8, 8, 6, 6,
}
var yyR2 = [...]int{

	0, 6, 2, 3, 2, 1, 2, 1, 2, 4,
	6, 1, 3, 1, 1, 2, 2, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, 4, 5, 7, -4, 6, 17, -4,
	9, -4, -4, -3, -5, -6, 9, 17, 18, 10,
	-3, -7, -8, 18, 17, -6, 12, -6, 13, 14,
	15, 16, -7, 11, -7, -6, 10, 11,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 0, 2, 0, 5, 4,
	0, 3, 6, 0, 7, 0, 0, 18, 19, 0,
	8, 0, 11, 13, 14, 0, 1, 0, 0, 15,
	16, 17, 0, 9, 12, 0, 0, 10,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line lang.y:31
		{
			yylex.(*lexer).varlist = yyDollar[1].varlist
			yylex.(*lexer).distinct = yyDollar[1].distinct
			yylex.(*lexer).triples = yyDollar[4].triples
			yylex.(*lexer).distinct = yyDollar[1].distinct
			yylex.(*lexer).count = yyDollar[1].count
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:41
		{
			yyVAL.varlist = yyDollar[2].varlist
			yyVAL.distinct = false
			yyVAL.count = false
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:47
		{
			yyVAL.varlist = yyDollar[3].varlist
			yyVAL.distinct = true
			yyVAL.count = false
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:53
		{
			yyVAL.varlist = yyDollar[2].varlist
			yyVAL.distinct = false
			yyVAL.count = true
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:61
		{
			yyVAL.varlist = []turtle.URI{turtle.ParseURI(yyDollar[1].str)}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:65
		{
			yyVAL.varlist = append([]turtle.URI{turtle.ParseURI(yyDollar[1].str)}, yyDollar[2].varlist...)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:71
		{
			yyVAL.triples = []Filter{yyDollar[1].triple}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:75
		{
			yyVAL.triples = append(yyDollar[2].triples, yyDollar[1].triple)
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:81
		{
			yyVAL.triple = Filter{yyDollar[1].val, yyDollar[2].pred, yyDollar[3].val}
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line lang.y:85
		{
			yyVAL.triple = Filter{yyDollar[2].val, yyDollar[3].pred, yyDollar[4].val}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:91
		{
			yyVAL.pred = yyDollar[1].pred
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:95
		{
			yyVAL.pred = append(yyDollar[1].pred, yyDollar[3].pred...)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:101
		{
			yyVAL.pred = []PathPattern{{Predicate: turtle.ParseURI(yyDollar[1].str), Pattern: PATTERN_SINGLE}}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:105
		{
			yyVAL.pred = []PathPattern{{Predicate: turtle.ParseURI(yyDollar[1].str), Pattern: PATTERN_SINGLE}}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:109
		{
			yyVAL.pred = []PathPattern{{Predicate: turtle.ParseURI(yyDollar[1].str), Pattern: PATTERN_ONE_PLUS}}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:113
		{
			yyVAL.pred = []PathPattern{{Predicate: turtle.ParseURI(yyDollar[1].str), Pattern: PATTERN_ZERO_ONE}}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:117
		{
			yyVAL.pred = []PathPattern{{Predicate: turtle.ParseURI(yyDollar[1].str), Pattern: PATTERN_ZERO_PLUS}}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:123
		{
			yyVAL.val = turtle.ParseURI(yyDollar[1].str)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:127
		{
			yyVAL.val = turtle.ParseURI(yyDollar[1].str)
		}
	}
	goto yystack /* stack new state and value */
}
