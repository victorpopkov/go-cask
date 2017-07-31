package cask

import (
	"bytes"
	"fmt"
)

// Errors represents a group of errors and its context.
type Errors struct {
	// context specifies the errors context.
	context string

	// errors specify a group of errors represented as an array.
	errors []error
}

// An unexpectedTokenError represents the error that occurs when an the expected
// tokens doesn't match the actual ones.
type unexpectedTokenError struct {
	// expectedTokens specify a group of expected token types.
	expectedTokens []TokenType

	// actualToken specifies the actual token type.
	actualToken TokenType
}

// NewErrors creates a new Errors instance and returns its pointer. Requires
// both Errors.context and Errors.errors to be passed as arguments.
func NewErrors(context string, errors ...error) *Errors {
	return &Errors{context, errors}
}

// Error returns all error messages represented as a single string. All errors
// include trailing newlines and prepended context.
func (e *Errors) Error() string {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, "%s:\n", e.context)
	for _, err := range e.errors {
		fmt.Fprintf(&buffer, "%s\n", err.Error())
	}

	return buffer.String()
}

// Error returns a string representation of the unexpectedTokenError.
func (u *unexpectedTokenError) Error() string {
	return fmt.Sprintf(
		"expected next token to be of type %v, got %s instead",
		u.expectedTokens,
		u.actualToken,
	)
}
