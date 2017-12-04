// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"SELECT",
		"*",
		"var",
		"FROM",
		"string",
		"WHERE",
		"{",
		"}",
		".",
		"uri",
		"url",
		"|",
		"/",
		"a",
		"(",
		")",
		"?",
		"+",
		"empty",
		"UNION",
	},

	idMap: map[string]Type{
		"INVALID": 0,
		"$":       1,
		"SELECT":  2,
		"*":       3,
		"var":     4,
		"FROM":    5,
		"string":  6,
		"WHERE":   7,
		"{":       8,
		"}":       9,
		".":       10,
		"uri":     11,
		"url":     12,
		"|":       13,
		"/":       14,
		"a":       15,
		"(":       16,
		")":       17,
		"?":       18,
		"+":       19,
		"empty":   20,
		"UNION":   21,
	},
}
