package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTokenTestParser() *Parser {
	// preparations
	p := &Parser{
		lexer: NewLexer("five = 5"),
	}

	p.nextToken()
	p.nextToken()

	return p
}

func TestNewParser(t *testing.T) {
	// preparations
	l := NewLexer("cask 'example' do\nend\n")
	p := NewParser(l)

	// test
	assert.IsType(t, Parser{}, *p)
	assert.Equal(t, l, p.lexer)
	assert.Len(t, p.errors, 0)
}

func TestNextToken(t *testing.T) {
	// preparations
	p := &Parser{
		lexer: NewLexer("five = 5"),
	}

	// test
	assert.Equal(t, EOF, p.currentToken.Type)
	assert.Equal(t, EOF, p.peekToken.Type)
	p.nextToken()
	p.nextToken()
	assert.Equal(t, IDENT, p.currentToken.Type)
	assert.Equal(t, ASSIGN, p.peekToken.Type)
}

func TestCurrentTokenIs(t *testing.T) {
	// preparations
	p := createTokenTestParser()

	// test
	assert.False(t, p.currentTokenIs(EOF))
	assert.True(t, p.currentTokenIs(IDENT))
}

func TestCurrentTokenOneOf(t *testing.T) {
	// preparations
	p := createTokenTestParser()

	// test
	assert.False(t, p.currentTokenOneOf(ASSIGN))
	assert.True(t, p.currentTokenOneOf(IDENT))
	assert.True(t, p.currentTokenOneOf(IDENT, ASSIGN))
}

func TestPeekTokenIs(t *testing.T) {
	// preparations
	p := createTokenTestParser()

	// test
	assert.False(t, p.peekTokenIs(IDENT))
	assert.True(t, p.peekTokenIs(ASSIGN))
}

func TestPeekTokenOneOf(t *testing.T) {
	// preparations
	p := createTokenTestParser()

	// test
	assert.False(t, p.peekTokenOneOf(IDENT))
	assert.True(t, p.peekTokenOneOf(ASSIGN))
	assert.True(t, p.peekTokenOneOf(ASSIGN, IDENT))
}

func TestAccept(t *testing.T) {
	// preparations
	p := createTokenTestParser()

	// test (successful)
	assert.True(t, p.peekTokenIs(ASSIGN))
	p.accept(ASSIGN)
	assert.True(t, p.peekTokenIs(INT))

	// test (error)
	assert.Len(t, p.errors, 0)
	p.accept(GLOBAL)
	assert.Len(t, p.errors, 1)
	assert.Equal(t, "expected next token to be of type [GLOBAL], got INT instead", p.errors[0].Error())
}

func TestAcceptOneOf(t *testing.T) {
	// preparations
	p := createTokenTestParser()

	// test
	assert.True(t, p.peekTokenIs(ASSIGN))
	p.acceptOneOf(ASSIGN, IDENT)
	assert.True(t, p.peekTokenIs(INT))

	// test (error)
	assert.Len(t, p.errors, 0)
	p.acceptOneOf(GLOBAL, IDENT)
	assert.Len(t, p.errors, 1)
	assert.Equal(t, "expected next token to be of type [GLOBAL IDENT], got INT instead", p.errors[0].Error())
}

func TestParserErrors(t *testing.T) {
	// preparations
	p := createTokenTestParser()

	// test
	assert.IsType(t, []error{}, p.Errors())
}
