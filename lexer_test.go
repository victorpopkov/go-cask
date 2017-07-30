package cask

// The Lexer has been stolen from:
// https://github.com/goruby/goruby/blob/master/lexer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testSingleNextToken(t *testing.T, input string, expectedType TokenType, expectedLiteral string) {
	lexer := NewLexer(input)
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

func TestLexerNextToken(t *testing.T) {
	// test (successful)
	input := `five = 5
# just comment
fifty = 5_0
ten = 10

def add(x, y)
	x + y
end

result = add(five, ten)
!-/*5;
5 < 10 > 5
return

if 5 < 10 then
	true
else
	false
end

10 == 10
10 != 9
""
"foobar"
'foobar'
"foo bar"
'foo bar'
:sym
.

def nil?
end

def run!
end

[1, 2]
nil
self
Ten = 10
module Abc
end
class Abc
end
add { |x| x }
add do |x|
end
yield
A::B
$foo;
$Foo
$@
$a`

	testCases := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{IDENT, "five"},
		{ASSIGN, "="},
		{INT, "5"},
		{NEWLINE, "\n"},
		{IDENT, "fifty"},
		{ASSIGN, "="},
		{INT, "5_0"},
		{NEWLINE, "\n"},
		{IDENT, "ten"},
		{ASSIGN, "="},
		{INT, "10"},
		{NEWLINE, "\n"},
		{NEWLINE, "\n"},
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
		{IDENT, "result"},
		{ASSIGN, "="},
		{IDENT, "add"},
		{LPAREN, "("},
		{IDENT, "five"},
		{COMMA, ","},
		{IDENT, "ten"},
		{RPAREN, ")"},
		{NEWLINE, "\n"},
		{BANG, "!"},
		{MINUS, "-"},
		{SLASH, "/"},
		{ASTERISK, "*"},
		{INT, "5"},
		{SEMICOLON, ";"},
		{NEWLINE, "\n"},
		{INT, "5"},
		{LT, "<"},
		{INT, "10"},
		{GT, ">"},
		{INT, "5"},
		{NEWLINE, "\n"},
		{RETURN, "return"},
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
		{INT, "10"},
		{EQ, "=="},
		{INT, "10"},
		{NEWLINE, "\n"},
		{INT, "10"},
		{NOTEQ, "!="},
		{INT, "9"},
		{NEWLINE, "\n"},
		{STRING, ""},
		{NEWLINE, "\n"},
		{STRING, "foobar"},
		{NEWLINE, "\n"},
		{STRING, "foobar"},
		{NEWLINE, "\n"},
		{STRING, "foo bar"},
		{NEWLINE, "\n"},
		{STRING, "foo bar"},
		{NEWLINE, "\n"},
		{SYMBOL, "sym"},
		{NEWLINE, "\n"},
		{DOT, "."},
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
		{LBRACKET, "["},
		{INT, "1"},
		{COMMA, ","},
		{INT, "2"},
		{RBRACKET, "]"},
		{NEWLINE, "\n"},
		{NIL, "nil"},
		{NEWLINE, "\n"},
		{SELF, "self"},
		{NEWLINE, "\n"},
		{CONST, "Ten"},
		{ASSIGN, "="},
		{INT, "10"},
		{NEWLINE, "\n"},
		{MODULE, "module"},
		{CONST, "Abc"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{CLASS, "class"},
		{CONST, "Abc"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{IDENT, "add"},
		{LBRACE, "{"},
		{PIPE, "|"},
		{IDENT, "x"},
		{PIPE, "|"},
		{IDENT, "x"},
		{RBRACE, "}"},
		{NEWLINE, "\n"},
		{IDENT, "add"},
		{DO, "do"},
		{PIPE, "|"},
		{IDENT, "x"},
		{PIPE, "|"},
		{NEWLINE, "\n"},
		{END, "end"},
		{NEWLINE, "\n"},
		{YIELD, "yield"},
		{NEWLINE, "\n"},
		{CONST, "A"},
		{SCOPE, "::"},
		{CONST, "B"},
		{NEWLINE, "\n"},
		{GLOBAL, "$foo"},
		{SEMICOLON, ";"},
		{NEWLINE, "\n"},
		{GLOBAL, "$Foo"},
		{NEWLINE, "\n"},
		{GLOBAL, "$@"},
		{NEWLINE, "\n"},
		{GLOBAL, "$a"},
		{EOF, ""},
	}

	lexer := NewLexer(input)

	// test
	for pos, testCase := range testCases {
		token := lexer.NextToken()

		// type
		assert.Equal(t, testCase.expectedType, token.Type, fmt.Sprintf(
			"Expected token with type %q at position %d, got type %q\n",
			testCase.expectedType,
			pos,
			token.Type,
		))

		// literal
		assert.Equal(t, testCase.expectedLiteral, token.Literal, fmt.Sprintf(
			"Expected token with literal %q at position %d, got literal %q\n",
			testCase.expectedLiteral,
			pos,
			token.Literal,
		))
	}

	// test (error)
	testSingleNextToken(t, "$ ", ILLEGAL, "Illegal character at 2: ' '")
	testSingleNextToken(t, "$;", ILLEGAL, "Illegal character at 2: ';'")
	testSingleNextToken(t, "\\", ILLEGAL, "Illegal character at 0: '\\'")
}
