package cask

// The Lexer has been stolen from:
// https://github.com/goruby/goruby/blob/master/lexer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNextToken(t *testing.T, lexer *Lexer, expectedType TokenType, expectedLiteral string, position int) {
	token := lexer.NextToken()

	// type
	assert.Equal(t, expectedType, token.Type, fmt.Sprintf(
		"Expected token with type %q at position %d, got type %q\n",
		expectedType,
		position,
		token.Type,
	))

	// literal
	assert.Equal(t, expectedLiteral, token.Literal, fmt.Sprintf(
		"Expected token with literal %q at position %d, got literal %q\n",
		expectedLiteral,
		position,
		token.Literal,
	))
}

func assertSingleNextToken(t *testing.T, input string, expectedType TokenType, expectedLiteral string) {
	lexer := NewLexer(input)

	if lexer.HasNext() {
		token := lexer.NextToken()

		// type
		assert.Equal(t, expectedType, token.Type, fmt.Sprintf(
			"Expected token with type %q, got type %q\n",
			expectedType,
			token.Type,
		))

		// literal
		assert.Equal(t, expectedLiteral, token.Literal, fmt.Sprintf(
			"Expected token with literal %q, got literal %q\n",
			expectedLiteral,
			token.Literal,
		))
	}
}

func TestLexerNextToken(t *testing.T) {
	// test (successful)
	input := `
# just comment

five = 5
fifty = 50
result = add(five, fifty)

!-/*5;
10 == 10
10 != 9
4 % 2
5 < 10 > 5

[1, 2]
A::B
:symbol

$foo;
$Foo
$@
$a

def add(x, y)
	x + y
end

def nil?
end

def run!
end

module Abc
end

class Abc
end

if 5 < 10 then
	true
else
	false
end

add do |x|
end

add { |x| x }
`

	testCases := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{IDENT, "five"},
		{ASSIGN, "="},
		{INT, "5"},
		{NEWLINE, "\n"},

		{IDENT, "fifty"},
		{ASSIGN, "="},
		{INT, "50"},
		{NEWLINE, "\n"},

		{IDENT, "result"},
		{ASSIGN, "="},
		{IDENT, "add"},
		{LPAREN, "("},
		{IDENT, "five"},
		{COMMA, ","},
		{IDENT, "fifty"},
		{RPAREN, ")"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{BANG, "!"},
		{MINUS, "-"},
		{SLASH, "/"},
		{ASTERISK, "*"},
		{INT, "5"},
		{SEMICOLON, ";"},
		{NEWLINE, "\n"},

		{INT, "10"},
		{EQ, "=="},
		{INT, "10"},
		{NEWLINE, "\n"},

		{INT, "10"},
		{NOTEQ, "!="},
		{INT, "9"},
		{NEWLINE, "\n"},

		{INT, "4"},
		{MODULUS, "%"},
		{INT, "2"},
		{NEWLINE, "\n"},

		{INT, "5"},
		{LT, "<"},
		{INT, "10"},
		{GT, ">"},
		{INT, "5"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{LBRACKET, "["},
		{INT, "1"},
		{COMMA, ","},
		{INT, "2"},
		{RBRACKET, "]"},
		{NEWLINE, "\n"},

		{CONST, "A"},
		{SCOPE, "::"},
		{CONST, "B"},
		{NEWLINE, "\n"},

		{SYMBOL, "symbol"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		// Globals

		{GLOBAL, "$foo"},
		{SEMICOLON, ";"},
		{NEWLINE, "\n"},

		{GLOBAL, "$Foo"},
		{NEWLINE, "\n"},

		{GLOBAL, "$@"},
		{NEWLINE, "\n"},

		{GLOBAL, "$a"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		// Blocks

		{DEF, "def"},
		{IDENT, "add"},
		{LPAREN, "("},
		{IDENT, "x"},
		{COMMA, ","},
		{IDENT, "y"},
		{RPAREN, ")"},
		{NEWLINE, "\n"},
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "y"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{DEF, "def"},
		{IDENT, "nil?"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{DEF, "def"},
		{IDENT, "run!"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{MODULE, "module"},
		{CONST, "Abc"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{CLASS, "class"},
		{CONST, "Abc"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{IF, "if"},
		{INT, "5"},
		{LT, "<"},
		{INT, "10"},
		{THEN, "then"},
		{NEWLINE, "\n"},
		{TRUE, "true"},
		{NEWLINE, "\n"},
		{ELSE, "else"},
		{NEWLINE, "\n"},
		{FALSE, "false"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{IDENT, "add"},
		{DO, "do"},
		{PIPE, "|"},
		{IDENT, "x"},
		{PIPE, "|"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},

		{IDENT, "add"},
		{LBRACE, "{"},
		{PIPE, "|"},
		{IDENT, "x"},
		{PIPE, "|"},
		{IDENT, "x"},
		{RBRACE, "}"},
		{NEWLINE, "\n"},

		{EOF, ""},
	}

	lexer := NewLexer(input)

	// test
	for position, testCase := range testCases {
		assertNextToken(t, lexer, testCase.expectedType, testCase.expectedLiteral, position)
	}

	// test (error)
	assertSingleNextToken(t, "$ ", ILLEGAL, "Illegal character at 2: ' '")
	assertSingleNextToken(t, "$;", ILLEGAL, "Illegal character at 2: ';'")
	assertSingleNextToken(t, "\\", ILLEGAL, "Illegal character at 0: '\\'")
}

func TestLexerPercentNotationRegexp(t *testing.T) {
	// test (successful)
	input := `
%r(regex)
%r[regex]
%r{regex}
%r<regex>
`

	testCases := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{NEWLINE, "\n"},

		{PNREGEXP, "%r"},
		{PNSTART, "("},
		{REGEXP, "regex"},
		{PNEND, ")"},
		{NEWLINE, "\n"},

		{PNREGEXP, "%r"},
		{PNSTART, "["},
		{REGEXP, "regex"},
		{PNEND, "]"},
		{NEWLINE, "\n"},

		{PNREGEXP, "%r"},
		{PNSTART, "{"},
		{REGEXP, "regex"},
		{PNEND, "}"},
		{NEWLINE, "\n"},

		{PNREGEXP, "%r"},
		{PNSTART, "<"},
		{REGEXP, "regex"},
		{PNEND, ">"},
		{NEWLINE, "\n"},

		{EOF, ""},
	}

	lexer := NewLexer(input)

	// test
	for position, testCase := range testCases {
		assertNextToken(t, lexer, testCase.expectedType, testCase.expectedLiteral, position)
	}
}

func TestLexerStrings(t *testing.T) {
	// test (successful)
	input := `
""
''
"double quoted string"
'single quoted string'
"double quoted string (\") with escaped double quote"
'single quoted string (\') with escaped single quote'
`

	testCases := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{NEWLINE, "\n"},

		{STRING, ""},
		{NEWLINE, "\n"},

		{STRING, ""},
		{NEWLINE, "\n"},

		{STRING, "double quoted string"},
		{NEWLINE, "\n"},

		{STRING, "single quoted string"},
		{NEWLINE, "\n"},

		{STRING, "double quoted string (\\\") with escaped double quote"},
		{NEWLINE, "\n"},

		{STRING, "single quoted string (\\') with escaped single quote"},
		{NEWLINE, "\n"},

		{EOF, ""},
	}

	lexer := NewLexer(input)

	// test
	for position, testCase := range testCases {
		assertNextToken(t, lexer, testCase.expectedType, testCase.expectedLiteral, position)
	}
}

func TestLexerDelimiters(t *testing.T) {
	assertSingleNextToken(t, ",", COMMA, ",")
	assertSingleNextToken(t, "\n", NEWLINE, "\n")
	assertSingleNextToken(t, ";", SEMICOLON, ";")
	assertSingleNextToken(t, ".", DOT, ".")
	assertSingleNextToken(t, "{", LBRACE, "{")
	assertSingleNextToken(t, "[", LBRACKET, "[")
	assertSingleNextToken(t, "(", LPAREN, "(")
	assertSingleNextToken(t, "|", PIPE, "|")
	assertSingleNextToken(t, "}", RBRACE, "}")
	assertSingleNextToken(t, "]", RBRACKET, "]")
	assertSingleNextToken(t, ")", RPAREN, ")")
	assertSingleNextToken(t, "::", SCOPE, "::")
}

func TestLexerKeywords(t *testing.T) {
	assertSingleNextToken(t, "class", CLASS, "class")
	assertSingleNextToken(t, "def", DEF, "def")
	assertSingleNextToken(t, "do", DO, "do")
	assertSingleNextToken(t, "else", ELSE, "else")
	assertSingleNextToken(t, "end", END, "end")
	assertSingleNextToken(t, "false", FALSE, "false")
	assertSingleNextToken(t, "if", IF, "if")
	assertSingleNextToken(t, "elsif", ELSEIF, "elsif")
	assertSingleNextToken(t, "module", MODULE, "module")
	assertSingleNextToken(t, "nil", NIL, "nil")
	assertSingleNextToken(t, "return", RETURN, "return")
	assertSingleNextToken(t, "self", SELF, "self")
	assertSingleNextToken(t, "then", THEN, "then")
	assertSingleNextToken(t, "true", TRUE, "true")
	assertSingleNextToken(t, "yield", YIELD, "yield")
}
