package cask

import (
	"bytes"
	"unicode"
)

// TokenType represents a known token type.
type TokenType int

// Different token types that can be recognized.
const (
	EOF     TokenType = iota // end of input
	ILLEGAL                  // an illegal/unknown character

	// Identifier + literals

	CONST
	GLOBAL
	IDENT
	INT
	STRING
	SYMBOL // :symbol

	// Operators

	ASSIGN   // =
	ASTERISK // *
	BANG     // !
	MINUS    // -
	PLUS     // +
	SLASH    // /

	EQ    // ==
	GT    // >
	LT    // <
	NOTEQ // !=

	// Delimiters

	COMMA
	NEWLINE // \n
	SEMICOLON

	COLON    // :
	DOT      // .
	LBRACE   // {
	LBRACKET // [
	LPAREN   // (
	PIPE     // |
	RBRACE   // }
	RBRACKET // ]
	RPAREN   // )

	SCOPE // ::

	// Keywords

	CLASS
	DEF
	DO
	ELSE
	ELSEIF
	END
	FALSE
	IF
	MODULE
	NIL
	RETURN
	SELF
	THEN
	TRUE
	YIELD
)

var keywords = map[string]TokenType{
	"class":  CLASS,
	"def":    DEF,
	"do":     DO,
	"else":   ELSE,
	"end":    END,
	"false":  FALSE,
	"if":     IF,
	"elsif":  ELSEIF,
	"module": MODULE,
	"nil":    NIL,
	"return": RETURN,
	"self":   SELF,
	"then":   THEN,
	"true":   TRUE,
	"yield":  YIELD,
}

// LookupIdent returns a TokenType keyword if ident is in the keywords map. If
// specified ident starts with an upper character it will return a CONST
// TokenType. Otherwise, it returns IDENT.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	if unicode.IsUpper(bytes.Runes([]byte(ident))[0]) {
		return CONST
	}

	return IDENT
}
