package cask

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrors(t *testing.T) {
	// preparations
	e := NewErrors("Test", errors.New("Error"))

	// test
	assert.IsType(t, Errors{}, *e)
	assert.Equal(t, "Test", e.context)
	assert.Len(t, e.errors, 1)
	assert.Equal(t, "Error", e.errors[0].Error())
}

func TestErrorsError(t *testing.T) {
	// preparations
	e := NewErrors("Test", errors.New("Error"))

	// test
	assert.Equal(t, "Test:\nError\n", e.Error())
}

func TestUnexpectedTokenErrorError(t *testing.T) {
	// preparations
	e := &unexpectedTokenError{
		expectedTokens: []TokenType{CONST},
		actualToken:    GLOBAL,
	}

	// test
	assert.Equal(t, "expected next token to be of type [CONST], got GLOBAL instead", e.Error())
}
