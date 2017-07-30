package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	// preparations
	token := NewToken(EQ, "==", 0)

	// test
	assert.IsType(t, Token{}, *token)
	assert.Equal(t, EQ, token.Type)
	assert.Equal(t, "==", token.Literal)
	assert.Equal(t, 0, token.Position)
}
