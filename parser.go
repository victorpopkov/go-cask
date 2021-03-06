package cask

import (
	"fmt"

	"github.com/pkg/errors"
)

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

	// insideIfElse specifies if the parser is currently inside of the if
	// statement.
	insideIfElse bool
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

// ParseCask parses the input from Parse.lexer into the Parse.cask.
func (p *Parser) ParseCask(cask *Cask) error {
	p.cask = cask

	for !p.currentTokenIs(EOF) {
		p.parseStatement()

		err := p.nextToken()
		if err != nil {
			break
		}
	}

	if p.currentCaskVariant != nil {
		p.cask.AddVariant(p.currentCaskVariant)
		p.currentCaskVariant = nil
	}

	first := p.cask.Variants[0]
	last := p.cask.Variants[len(p.cask.Variants)-1]
	for _, v := range p.cask.Variants {
		// version
		if v.Version == nil && last.Version != nil && last.Version.IsGlobal {
			v.Version = last.Version
		} else if v.Version == nil && first.Version != nil && first.Version.IsGlobal {
			v.Version = first.Version
		}

		// sha256
		if v.SHA256 == nil && last.SHA256 != nil && last.SHA256.IsGlobal {
			v.SHA256 = last.SHA256
		}

		// url
		if v.URL == nil && last.URL != nil && last.URL.IsGlobal {
			v.URL = last.URL
		}

		// appcast
		if v.Appcast == nil && last.Appcast != nil && last.Appcast.IsGlobal {
			v.Appcast = last.Appcast
		}

		// name
		if len(v.Names) == 0 && len(last.Names) != 0 {
			for _, n := range last.Names {
				if n.IsGlobal {
					v.Names = last.Names
				}
			}
		}

		// homepage
		if v.Homepage == nil && last.Homepage != nil && last.Homepage.IsGlobal {
			v.Homepage = last.Homepage
		}

		// artifact
		if len(v.Artifacts) == 0 {
			for _, a := range last.Artifacts {
				v.AddArtifact(a)
			}
		}
	}

	if len(p.errors) != 0 {
		return NewErrors("Parsing errors", p.errors...)
	}

	return nil
}

// parseStatement parses a single statement.
func (p *Parser) parseStatement() {
	switch p.currentToken.Type {
	case ILLEGAL:
		p.errors = append(p.errors, fmt.Errorf("%s", p.currentToken.Literal))
	case EOF:
		p.errors = append(p.errors, &unexpectedTokenError{
			expectedTokens: []TokenType{NEWLINE},
			actualToken:    EOF,
		})
	default:
		p.parseExpressionStatement()
	}
}

// parseExpressionStatement parses a single expression statement.
func (p *Parser) parseExpressionStatement() {
	if p.currentCaskVariant == nil {
		p.currentCaskVariant = NewVariant()
	}

	switch p.currentToken.Type {
	case IDENT:
		if p.currentToken.Literal == "cask" {
			if p.peekTokenIs(STRING) {
				p.accept(STRING)
				p.cask.Token = p.currentToken.Literal
			}
		}

		if p.peekTokenIs(STRING) {
			switch p.currentToken.Literal {
			case "sha256":
				if p.currentCaskVariant.SHA256 != nil {
					p.mergeCurrentCaskVariantIfNotEmpty(p.currentCaskVariant.SHA256.Value)
				}

				s := NewSHA256(p.peekToken.Literal)
				if !p.insideIfElse {
					s.IsGlobal = true
				}
				p.currentCaskVariant.SHA256 = s
			case "url":
				if p.currentCaskVariant.URL != nil {
					p.mergeCurrentCaskVariantIfNotEmpty(p.currentCaskVariant.URL.Value)
				}

				u := NewURL(p.peekToken.Literal)
				if !p.insideIfElse {
					u.IsGlobal = true
				}
				p.currentCaskVariant.URL = u
			case "appcast":
				if p.currentCaskVariant.Appcast != nil {
					p.mergeCurrentCaskVariantIfNotEmpty(p.currentCaskVariant.Appcast.URL)
				}

				a, err := p.parseAppcast()
				if err == nil {
					if !p.insideIfElse {
						a.IsGlobal = true
					}
					p.currentCaskVariant.Appcast = a
				}
			case "name":
				p.mergeCurrentCaskVariantIfNotEmpty(p.currentCaskVariant.Names)

				n := NewName(p.peekToken.Literal)
				if !p.insideIfElse {
					n.IsGlobal = true
				}
				p.currentCaskVariant.AddName(n)
			case "homepage":
				if p.currentCaskVariant.Homepage != nil {
					p.mergeCurrentCaskVariantIfNotEmpty(p.currentCaskVariant.Homepage.Value)
				}

				h := NewHomepage(p.peekToken.Literal)
				if !p.insideIfElse {
					h.IsGlobal = true
				}
				p.currentCaskVariant.Homepage = h
			case "version":
				if p.currentCaskVariant.Version != nil {
					p.mergeCurrentCaskVariantIfNotEmpty(p.currentCaskVariant.Version.Value)
				}

				v, err := p.parseVersion()
				if err == nil {
					if !p.insideIfElse {
						v.IsGlobal = true
					}
					p.currentCaskVariant.Version = v
				}
			}

			// artifacts
			if p.currentTokenLiteralOneOf("app", "pkg", "binary") {
				if p.currentIfVariant != nil {
					p.mergeCurrentCaskVariantIfNotEmpty(p.currentCaskVariant.Artifacts)
				}

				a, err := p.ParseArtifact()
				if err == nil && a != nil {
					p.currentCaskVariant.AddArtifact(a)
				}
			}
		}

		if p.peekTokenIs(SYMBOL) {
			switch p.currentToken.Literal {
			case "version":
				if p.currentCaskVariant.Version != nil {
					p.mergeCurrentCaskVariantIfNotEmpty(p.currentCaskVariant.Version.Value)
				}

				v, err := p.parseVersion()
				if err == nil {
					if !p.insideIfElse {
						v.IsGlobal = true
					}
					p.currentCaskVariant.Version = v
				}
			}
		}
	case IF:
		p.parseIfExpression()
	case ELSEIF:
		p.parseIfExpression()
	case ELSE:
		p.insideIfElse = true
		p.parseBlockStatement()
	}

	if p.peekTokenOneOf(SEMICOLON, NEWLINE, COMMA) {
		p.nextToken()
	}
}

func (p *Parser) parseIfExpression() {
	p.nextToken()

	p.parseIfCondition()

	if p.peekTokenIs(THEN) {
		p.accept(THEN)
	}

	if !p.peekTokenOneOf(NEWLINE, SEMICOLON) {
		err := errors.Wrap(
			&unexpectedTokenError{
				expectedTokens: []TokenType{NEWLINE, SEMICOLON},
				actualToken:    p.peekToken.Type,
			},
			fmt.Sprintf(
				"could not parse if expression: unexpected token %s: '%s'",
				p.peekToken.Type.String(),
				p.peekToken.Literal,
			),
		)

		p.errors = append(p.errors, err)

		return
	}

	p.parseBlockStatement(ELSE, ELSEIF)

	if p.currentIfVariant != nil {
		p.currentCaskVariant.MinimumSupportedMacOS = p.currentIfVariant.MinimumSupportedMacOS
		p.currentCaskVariant.MaximumSupportedMacOS = p.currentIfVariant.MaximumSupportedMacOS
		p.currentIfVariant = nil
	}

	p.insideIfElse = false

	return
}

// parseIfCondition parses the if condition if it's supported.
func (p *Parser) parseIfCondition() {
	p.currentIfVariant = NewVariant()
	p.insideIfElse = true

	min, max, err := p.ParseConditionMacOS()
	if err == nil {
		p.currentIfVariant.MinimumSupportedMacOS = min
		p.currentIfVariant.MaximumSupportedMacOS = max
		return
	}
}

// parseBlockStatement parses the block statement if the Parser.peekToken
// matches the requirements.
func (p *Parser) parseBlockStatement(t ...TokenType) {
	terminatorTokens := append(
		[]TokenType{
			END,
			EOF,
		},
		t...,
	)

	for !p.peekTokenOneOf(terminatorTokens...) {
		p.nextToken()
		p.parseExpressionStatement()
	}

	p.insideIfElse = false
}

// parseVersion parses the version if the Parser.peekToken matches the cask
// requirements. If the ":latest" symbol is found, the Version will have the
// "latest" string value.
func (p *Parser) parseVersion() (*Version, error) {
	if p.peekTokenIs(STRING) {
		p.accept(STRING)
		return NewVersion(p.currentToken.Literal), nil
	}

	if p.peekTokenIs(SYMBOL) && p.peekToken.Literal == "latest" {
		p.accept(SYMBOL)
		return NewVersion("latest"), nil
	}

	return nil, errors.New("version not found")
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
		return p.parseArtifactApp()
	case "pkg":
		return p.parseArtifactPkg()
	case "binary":
		return p.parseArtifactBinary()
	default:
		return nil, errors.New("artifact not found")
	}
}

// parseArtifactApp parses the "app" artifact if the Parser.currentToken matches
// the requirements.
func (p *Parser) parseArtifactApp() (*Artifact, error) {
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

// parseArtifactPkg parses the "pkg" artifact if the Parser.currentToken matches
// the requirements.
func (p *Parser) parseArtifactPkg() (*Artifact, error) {
	if p.currentTokenIs(IDENT) && p.currentToken.Literal == "pkg" {
		if p.peekTokenIs(STRING) {
			p.accept(STRING)

			a := NewArtifact(ArtifactPkg, p.currentToken.Literal)

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

// parseArtifactBinary parses the "binary" artifact if the Parser.currentToken
// matches the requirements.
func (p *Parser) parseArtifactBinary() (*Artifact, error) {
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

// ParseConditionMacOS parses the "MacOS.version" condition statement. Returns
// both the minimum and maximum macOS releases. By default, the minimum is
// MacOSTiger and the maximum matches the latest macOS release which is
// MacOSHighSierra.
func (p *Parser) ParseConditionMacOS() (min MacOS, max MacOS, err error) {
	var comparison TokenType
	var hasEqual bool
	var mac MacOS

	if p.currentTokenIs(CONST) && p.currentToken.Literal == "MacOS" {
		p.accept(DOT)

		// version
		if p.peekTokenIs(IDENT) && p.peekToken.Literal == "version" {
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

// mergeCurrentCaskVariantIfNotEmpty is a Parser.mergeCurrentCaskVariant
// convenience function for checking if the provided interface is not
// empty. Supported types: string, []string, []Artifact.
func (p *Parser) mergeCurrentCaskVariantIfNotEmpty(i interface{}) {
	switch i.(type) {
	case string:
		p.mergeCurrentCaskVariant(i.(string) != "")
	case []string:
		p.mergeCurrentCaskVariant(len(i.([]string)) > 0)
	case []Artifact:
		p.mergeCurrentCaskVariant(len(i.([]Artifact)) > 0)
	}
}

// mergeCurrentCaskVariant adds Parser.currentCaskVariant to the Cask.Variants
// array and create a new empty Parser.currentCaskVariant if the provided
// condition is true.
func (p *Parser) mergeCurrentCaskVariant(condition bool) {
	if condition {
		p.cask.AddVariant(p.currentCaskVariant)
		p.currentCaskVariant = NewVariant()
	}
}

// nextToken updates the Parser.currentToken and Parser.peekToken values to
// match the next Lexer token. If Lexer doesn't have any token left, returns the
// "No tokens left" error.
func (p *Parser) nextToken() error {
	p.currentToken = p.peekToken
	if p.lexer.HasNext() {
		p.peekToken = p.lexer.NextToken()
	} else {
		return errors.New("No tokens left")
	}

	return nil
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
