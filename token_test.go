package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	token := NewToken(EQ, "==", 0)
	assert.IsType(t, Token{}, *token)
	assert.Equal(t, EQ, token.Type)
	assert.Equal(t, "==", token.Literal)
	assert.Equal(t, 0, token.Position)
}
