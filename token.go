package cask

// A Token represents a known token used in Lexer with its literal
// representation.
type Token struct {
	// Type specifies the recognized token type.
	Type TokenType

	// Literal specifies the literal representation of the token which is the
	// token value itself.
	Literal string

	// Position specifies the position where token is found.
	Position int
}

// NewToken returns a new Token.
func NewToken(t TokenType, literal string, position int) *Token {
	return &Token{t, literal, position}
}
