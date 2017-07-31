package cask

import "github.com/pkg/errors"

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

// parseVersion parses the version if the Parser.peekToken matches the cask
// requirements. If the ":latest" symbol is found, the version will become the
// "latest" string.
func (p *Parser) parseVersion() (string, error) {
	if p.peekTokenIs(STRING) {
		p.accept(STRING)
		return p.currentToken.Literal, nil
	}

	if p.peekTokenIs(SYMBOL) && p.peekToken.Literal == "latest" {
		p.accept(SYMBOL)
		return "latest", nil
	}

	return "", errors.New("version not found")
}

// parseAppcast parses the appcast if the Parser.peekToken matches the cask
// requirements. Supports both with and without checkpoint.
func (p *Parser) parseAppcast() (*Appcast, error) {
	if p.peekTokenIs(STRING) {
		p.accept(STRING)

		url := p.currentToken.Literal
		checkpoint := ""

		if p.peekTokenIs(COMMA) {
			p.accept(COMMA)
		}

		if p.peekTokenIs(NEWLINE) {
			p.accept(NEWLINE)
		}

		if p.peekTokenIs(IDENT) && p.peekToken.Literal == "checkpoint" {
			p.accept(IDENT)
			p.accept(SYMBOL)

			if p.peekTokenIs(STRING) {
				p.accept(STRING)
				checkpoint = p.currentToken.Literal
			}
		}

		return NewAppcast(url, checkpoint), nil
	}

	return nil, errors.New("appcast not found")
}

// ParseArtifact parses the supported artifact if the Parser.currentToken
// literal value matches the supported one. It runs the corresponding artifact
// specific parsing function. Returns an "artifact not found" error if the
// Parser.currentToken literal value doesn't match any supported one.
func (p *Parser) ParseArtifact() (*Artifact, error) {
	switch p.currentToken.Literal {
	case "app":
		return p.ParseArtifactApp()
	case "pkg":
		return p.ParseArtifactPkg()
	case "binary":
		return p.ParseArtifactBinary()
	default:
		return nil, errors.New("artifact not found")
	}
}

// ParseArtifactApp parses the "app" artifact if the Parser.currentToken matches
// the requirements.
func (p *Parser) ParseArtifactApp() (*Artifact, error) {
	if p.currentTokenIs(IDENT) && p.currentToken.Literal == "app" {
		if p.peekTokenIs(STRING) {
			p.accept(STRING)

			a := NewArtifact(ArtifactApp, p.currentToken.Literal)

			if p.peekTokenIs(COMMA) {
				p.accept(COMMA)
			}

			if p.peekTokenIs(NEWLINE) {
				p.accept(NEWLINE)
			}

			if p.peekTokenIs(IDENT) && p.peekToken.Literal == "target" {
				p.accept(IDENT)
				p.accept(SYMBOL)

				if p.peekTokenIs(STRING) {
					p.accept(STRING)
					a.Target = p.currentToken.Literal
				}
			}

			return a, nil
		}
	}

	return nil, errors.New(`error parsing "app" artifact`)
}

// ParseArtifactPkg parses the "pkg" artifact if the Parser.currentToken matches
// the requirements.
func (p *Parser) ParseArtifactPkg() (*Artifact, error) {
	if p.currentTokenIs(IDENT) && p.currentToken.Literal == "pkg" {
		if p.peekTokenIs(STRING) {
			p.accept(STRING)

			a := NewArtifact(ArtifactApp, p.currentToken.Literal)

			if p.peekTokenIs(COMMA) {
				p.accept(COMMA)
			}

			if p.peekTokenIs(NEWLINE) {
				p.accept(NEWLINE)
			}

			if p.peekTokenIs(IDENT) && p.peekToken.Literal == "allow_untrusted" {
				p.accept(IDENT)
				p.accept(SYMBOL)

				if p.peekTokenIs(TRUE) {
					p.accept(TRUE)
					a.AllowUntrusted = true
				}
			}

			return a, nil
		}
	}

	return nil, errors.New(`error parsing "pkg" artifact`)
}

// ParseArtifactBinary parses the "binary" artifact if the Parser.currentToken
// matches the requirements.
func (p *Parser) ParseArtifactBinary() (*Artifact, error) {
	if p.currentTokenIs(IDENT) && p.currentToken.Literal == "binary" {
		if p.peekTokenIs(STRING) {
			p.accept(STRING)

			a := NewArtifact(ArtifactBinary, p.currentToken.Literal)

			if p.peekTokenIs(COMMA) {
				p.accept(COMMA)
			}

			if p.peekTokenIs(NEWLINE) {
				p.accept(NEWLINE)
			}

			if p.peekTokenIs(IDENT) && p.peekToken.Literal == "target" {
				p.accept(IDENT)
				p.accept(SYMBOL)

				if p.peekTokenIs(STRING) {
					p.accept(STRING)
					a.Target = p.currentToken.Literal
				}
			}

			return a, nil
		}
	}

	return nil, errors.New(`error parsing "binary" artifact`)
}

// ParseConditionMacOS parses the "MacOS.release" condition statement. Returns
// both the minimum and maximum macOS release. By default, the minimum is
// MacOSTiger and the maximum matches the latest macOS release which is
// MacOSHighSierra.
func (p *Parser) ParseConditionMacOS() (min MacOS, max MacOS, err error) {
	var comparison TokenType
	var hasEqual bool
	var mac MacOS

	if p.currentTokenIs(CONST) && p.currentToken.Literal == "MacOS" {
		p.accept(DOT)

		// release
		if p.peekTokenIs(IDENT) && p.peekToken.Literal == "release" {
			p.accept(IDENT)

			// comparison
			if p.peekTokenOneOf(EQ, GT, LT) {
				p.acceptOneOf(EQ, GT, LT)
				comparison = p.currentToken.Type

				if p.peekTokenIs(ASSIGN) {
					p.accept(ASSIGN)
					hasEqual = true
				}
			}

			// macOS
			if p.peekTokenIs(SYMBOL) {
				p.accept(SYMBOL)
				switch p.currentToken.Literal {
				case "high_sierra":
					mac = MacOSHighSierra
				case "sierra":
					mac = MacOSSierra
				case "el_capitan":
					mac = MacOSElCapitan
				case "yosemite":
					mac = MacOSYosemite
				case "mavericks":
					mac = MacOSMavericks
				case "mountain_lion":
					mac = MacOSMountainLion
				case "lion":
					mac = MacOSLion
				case "snow_leopard":
					mac = MacOSSnowLeopard
				case "leopard":
					mac = MacOSLeopard
				case "tiger":
					mac = MacOSTiger
				default:
					return MacOSHighSierra, MacOSHighSierra, errors.New("MacOS condition is unknown")
				}
			}

			// comparison with macOS
			switch comparison {
			case EQ:
				return mac, mac, nil
			case GT:
				min = mac - 1
				max = MacOSHighSierra
				if hasEqual || min < 0 {
					min = mac
				}
				return min, max, nil
			case LT:
				min = MacOSTiger
				max = mac + 1
				if hasEqual || max > MacOSTiger {
					max = mac
				}
				return min, max, nil
			}
		}
	}

	// by default should return the latest
	return MacOSHighSierra, MacOSHighSierra, errors.New("MacOS condition not found")
}

// nextToken updates the Parser.currentToken and Parser.peekToken values to
// match the next Lexer token.
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	if p.lexer.HasNext() {
		p.peekToken = p.lexer.NextToken()
	}
}

// currentTokenIs checks whether the current Token.Type matches the specified
// TokenType.
func (p *Parser) currentTokenIs(t TokenType) bool {
	return p.currentToken.Type == t
}

// currentTokenOneOf checks whether the current Token.Type is from valid
// TokenType set.
func (p *Parser) currentTokenOneOf(types ...TokenType) bool {
	for _, t := range types {
		if p.currentToken.Type == t {
			return true
		}
	}
	return false
}

// currentTokenLiteralIs checks whether the current Token.Literal matches the
// specified value.
func (p *Parser) currentTokenLiteralIs(l string) bool {
	return p.currentToken.Literal == l
}

// currentTokenLiteralOneOf checks whether the current Token.Literal is from
// valid values set.
func (p *Parser) currentTokenLiteralOneOf(literals ...string) bool {
	for _, l := range literals {
		if p.currentToken.Literal == l {
			return true
		}
	}
	return false
}

// peekTokenIs checks whether the next Token.Type matches the specified
// TokenType.
func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

// peekTokenOneOf checks whether the next Token.Type is from valid TokenType
// set.
func (p *Parser) peekTokenOneOf(types ...TokenType) bool {
	for _, t := range types {
		if p.peekToken.Type == t {
			return true
		}
	}
	return false
}

// accept moves to the next Token if it's from the valid TokenType set.
func (p *Parser) accept(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

// acceptOneOf moves to the next Token if it's from the valid TokenType set.
func (p *Parser) acceptOneOf(t ...TokenType) bool {
	if p.peekTokenOneOf(t...) {
		p.nextToken()
		return true
	}

	p.peekError(t...)
	return false
}

// peekError adds a new unexpectedTokenError to Parser.errors.
func (p *Parser) peekError(t ...TokenType) {
	p.errors = append(p.errors, &unexpectedTokenError{
		expectedTokens: t,
		actualToken:    p.peekToken.Type,
	})
}

// Errors returns all errors which happened during the Parser.input parsing.
func (p *Parser) Errors() []error {
	return p.errors
}
