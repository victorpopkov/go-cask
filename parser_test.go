package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		lexer:  NewLexer("five = 5"),
		errors: []error{},
	}

	// test
	assert.Equal(t, EOF, p.currentToken.Type)
	assert.Equal(t, EOF, p.peekToken.Type)
	p.nextToken()
	p.nextToken()
	assert.Equal(t, IDENT, p.currentToken.Type)
	assert.Equal(t, ASSIGN, p.peekToken.Type)
}
