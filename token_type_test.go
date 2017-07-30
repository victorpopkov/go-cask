package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupIdent(t *testing.T) {
	testCases := map[string]TokenType{
		// Keywords
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

		// CONST
		"Test":    CONST,
		"Example": CONST,

		// IDENT
		"test":    IDENT,
		"example": IDENT,
	}

	for testCase, expected := range testCases {
		assert.Equal(t, expected, LookupIdent(testCase))
	}
}
