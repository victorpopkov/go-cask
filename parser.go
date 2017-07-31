package cask

// A Parser represents the parser that uses the emitted token provided by Lexer.
type Parser struct {
	// cask specifies the parsed Cask.
	cask *Cask

	// lexer specifies Lexer pointer.
	lexer *Lexer

	// currentToken specifies the current emitted Lexer token.
	currentToken Token

	// peekToken specifies the next Lexer token.
	peekToken Token

	// errors specify an array of errors.
	errors []error

	// currentCaskVariant specifies the temporary cask Variant that currently
	// being parsed.
	currentCaskVariant *Variant

	// currentIfVariant specifies the temporary cask Variant that holds data
	// extracted in if condition.
	currentIfVariant *Variant
}

// NewParser creates a new Parser instance and returns its pointer. Requires a
// Lexer and a Cask to be specified as arguments.
func NewParser(lexer *Lexer) *Parser {
	p := &Parser{
		lexer:  lexer,
		errors: []error{},
	}

	// read two tokens, so both currentToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken updates the Parser.currentToken and Parser.peekToken values to
// match the next Lexer token.
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	if p.lexer.HasNext() {
		p.peekToken = p.lexer.NextToken()
	}
}

// currentTokenIs checks whether the current Token matches the specified
// TokenType.
func (p *Parser) currentTokenIs(t TokenType) bool {
	return p.currentToken.Type == t
}

// currentTokenOneOf checks whether the current Token is from valid TokenType
// set.
func (p *Parser) currentTokenOneOf(types ...TokenType) bool {
	for _, t := range types {
		if p.currentToken.Type == t {
			return true
		}
	}
	return false
}

// peekTokenIs checks whether the next Token matches the specified TokenType.
func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

// peekTokenOneOf checks whether the next Token is from valid TokenType set.
func (p *Parser) peekTokenOneOf(types ...TokenType) bool {
	for _, t := range types {
		if p.peekToken.Type == t {
			return true
		}
	}
	return false
}
