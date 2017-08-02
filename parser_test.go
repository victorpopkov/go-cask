package cask

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTokenTestParser() *Parser {
	p := &Parser{
		lexer: NewLexer("five = 5"),
	}

	p.nextToken()
	p.nextToken()

	return p
}

func createCaskTestParser() *Parser {
	c := NewCask("cask 'example' do\nend\n")

	return c.parser
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

func TestParseStatement(t *testing.T) {
	testCases := map[string]interface{}{
		// token
		"cask 'example-one' do\nend": nil,

		// cask stanzas
		"version 'test'":  nil,
		"version :latest": nil,
		"sha256 'test'":   nil,
		"url 'test'":      nil,
		"appcast 'test'":  nil,
		"name 'test'":     nil,
		"homepage 'test'": nil,
		"app 'test'":      nil,

		// if/elsif
		"if MacOS.release == :tiger\nfive = 5\nend":    nil,
		"elsif MacOS.release == :tiger\nfive = 5\nend": nil,

		// errors
		"":   "expected next token to be of type [NEWLINE], got EOF instead",
		"\\": "Illegal character at 0: '\\'",
	}

	// preparations
	for input, err := range testCases {
		// preparations
		c := NewCask(input)
		p := c.parser

		// test
		if err == nil {
			assert.Len(t, p.errors, 0)
			p.parseStatement()
			assert.Len(t, p.errors, 0)
		} else {
			assert.Len(t, p.errors, 0)
			p.parseStatement()
			assert.Len(t, p.errors, 1)
			assert.Error(t, p.errors[0])
			assert.Equal(t, err, p.errors[0].Error())
		}
	}
}

func TestParseIfExpression(t *testing.T) {
	testCases := map[string]*Token{
		// successful
		"if MacOS.release == :tiger\nfive = 5\nend":      nil,
		"if MacOS.release == :tiger then\nfive = 5; end": nil,
		"if MacOS.release == :tiger then; five = 5; end": nil,

		// errors
		"if": {EOF, "", 0},
		"if MacOS.release == :tiger": {EOF, "", 0},

		// unknown conditions
		"if five == 5\nfive = 5\nend":      {EQ, "==", 0},
		"if five == 5; then five = 5; end": {EQ, "==", 0},
	}

	// preparations
	for input, err := range testCases {
		// preparations
		l := NewLexer(input)
		p := NewParser(l)

		// test
		if err == nil {
			assert.Len(t, p.errors, 0)
			p.parseIfExpression()
			assert.Len(t, p.errors, 0)
		} else {
			assert.Len(t, p.errors, 0)
			p.parseIfExpression()
			assert.Len(t, p.errors, 1)
			assert.Error(t, p.errors[0])
			assert.Equal(
				t,
				fmt.Sprintf(
					"could not parse if expression: unexpected token %s: '%s': expected next token to be of type [NEWLINE SEMICOLON], got %s instead",
					err.Type.String(),
					err.Literal,
					err.Type.String(),
				),
				p.errors[0].Error(),
			)
		}
	}
}

func TestParseVersion(t *testing.T) {
	// test (successful)
	testCases := map[string]string{
		`version "1.0.0"`: "1.0.0",
		"version '1.0.0'": "1.0.0",
		"version :latest": "latest",
	}

	for testCase, expected := range testCases {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.parseVersion()
		assert.Nil(t, err)
		assert.IsType(t, &Version{}, actual)
		assert.Equal(t, expected, actual.Value)
	}

	// test (error)
	testCasesErrors := map[string]string{
		"invalid":       "version not found",
		"version 1.0.0": "version not found",
	}

	for testCase, expected := range testCasesErrors {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.parseVersion()
		assert.Empty(t, actual)
		assert.Error(t, err)
		assert.Equal(t, expected, err.Error())
	}
}

func TestParseAppcast(t *testing.T) {
	// test (successful)
	testCases := map[string]Appcast{
		`appcast "https://example.com/appcast.xml"`: {
			Stanza: Stanza{
				Value:    "https://example.com/appcast.xml",
				IsGlobal: true,
			},
			Checkpoint: "",
		},
		"appcast 'https://example.com/appcast.xml'": {
			Stanza: Stanza{
				Value:    "https://example.com/appcast.xml",
				IsGlobal: true,
			},
			Checkpoint: "",
		},
		"appcast 'https://example.com/appcast.xml', checkpoint: '2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1'": {
			Stanza: Stanza{
				Value:    "https://example.com/appcast.xml",
				IsGlobal: true,
			},
			Checkpoint: "2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1",
		},
		"appcast 'https://example.com/appcast.xml',\ncheckpoint: '2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1'": {
			Stanza: Stanza{
				Value:    "https://example.com/appcast.xml",
				IsGlobal: true,
			},
			Checkpoint: "2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1",
		},
	}

	for testCase, expected := range testCases {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.parseAppcast()
		assert.Nil(t, err)
		assert.IsType(t, Appcast{}, *actual)
		assert.Equal(t, expected.Value, actual.Value)
		assert.Equal(t, expected.Checkpoint, actual.Checkpoint)
	}

	// test (error)
	testCasesErrors := map[string]string{
		"invalid": "appcast not found",
		"appcast https://example.com/appcast.xml": "appcast not found",
	}

	for testCase, expected := range testCasesErrors {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.parseAppcast()
		assert.Nil(t, actual)
		assert.Error(t, err)
		assert.Equal(t, expected, err.Error())
	}
}

func TestParseArtifactApp(t *testing.T) {
	// test (successful)
	testCases := map[string]Artifact{
		// app
		`app "Test.app"`: {
			Type:   ArtifactApp,
			Value:  "Test.app",
			Target: "",
		},
		"app 'Test.app'": {
			Type:   ArtifactApp,
			Value:  "Test.app",
			Target: "",
		},
		"app 'Test.app', target: 'Test Target.app'": {
			Type:   ArtifactApp,
			Value:  "Test.app",
			Target: "Test Target.app",
		},
		"app 'Test.app',\ntarget: 'Test Target.app'": {
			Type:   ArtifactApp,
			Value:  "Test.app",
			Target: "Test Target.app",
		},
	}

	for testCase, expected := range testCases {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.ParseArtifact()
		assert.Nil(t, err)
		assert.IsType(t, Artifact{}, *actual)
		assert.Equal(t, expected.Type, actual.Type)
		assert.Equal(t, expected.Value, actual.Value)
		assert.Equal(t, expected.Target, actual.Target)
		assert.False(t, actual.AllowUntrusted)
	}

	// test (error)
	testCasesErrors := map[string]string{
		"invalid":     "artifact not found",
		"app invalid": `error parsing "app" artifact`,
	}

	for testCase, expected := range testCasesErrors {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.ParseArtifact()
		assert.Nil(t, actual)
		assert.Error(t, err)
		assert.Equal(t, expected, err.Error())
	}
}

func TestParseArtifactPkg(t *testing.T) {
	// test (successful)
	testCases := map[string]Artifact{
		// app
		`pkg "test.pkg"`: {
			Type:           ArtifactPkg,
			Value:          "test.pkg",
			AllowUntrusted: false,
		},
		"pkg 'test.pkg'": {
			Type:           ArtifactPkg,
			Value:          "test.pkg",
			AllowUntrusted: false,
		},
		"pkg 'test.pkg', allow_untrusted: true": {
			Type:           ArtifactPkg,
			Value:          "test.pkg",
			AllowUntrusted: true,
		},
		"pkg 'test.pkg',\nallow_untrusted: true": {
			Type:           ArtifactPkg,
			Value:          "test.pkg",
			AllowUntrusted: true,
		},
	}

	for testCase, expected := range testCases {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.ParseArtifact()
		assert.Nil(t, err)
		assert.IsType(t, Artifact{}, *actual)
		assert.Equal(t, expected.Type, actual.Type)
		assert.Equal(t, expected.Value, actual.Value)
		assert.Equal(t, expected.AllowUntrusted, actual.AllowUntrusted)
	}

	// test (error)
	testCasesErrors := map[string]string{
		"invalid":     "artifact not found",
		"pkg invalid": `error parsing "pkg" artifact`,
	}

	for testCase, expected := range testCasesErrors {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.ParseArtifact()
		assert.Nil(t, actual)
		assert.Error(t, err)
		assert.Equal(t, expected, err.Error())
	}
}

func TestParseArtifactBinary(t *testing.T) {
	// test (successful)
	testCases := map[string]Artifact{
		// app
		`binary "test"`: {
			Type:   ArtifactBinary,
			Value:  "test",
			Target: "",
		},
		"binary 'test'": {
			Type:   ArtifactBinary,
			Value:  "test",
			Target: "",
		},
		"binary 'test', target: 'test-target'": {
			Type:   ArtifactBinary,
			Value:  "test",
			Target: "test-target",
		},
		"binary 'test',\ntarget: 'test-target'": {
			Type:   ArtifactBinary,
			Value:  "test",
			Target: "test-target",
		},
	}

	for testCase, expected := range testCases {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.ParseArtifact()
		assert.Nil(t, err)
		assert.IsType(t, Artifact{}, *actual)
		assert.Equal(t, expected.Type, actual.Type)
		assert.Equal(t, expected.Value, actual.Value)
		assert.Equal(t, expected.Target, actual.Target)
	}

	// test (error)
	testCasesErrors := map[string]string{
		"invalid":        "artifact not found",
		"binary invalid": `error parsing "binary" artifact`,
	}

	for testCase, expected := range testCasesErrors {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.ParseArtifact()
		assert.Nil(t, actual)
		assert.Error(t, err)
		assert.Equal(t, expected, err.Error())
	}
}

func TestParseConditionMacOS(t *testing.T) {
	// test (successful)
	testCases := map[string][]MacOS{
		// EQ (==)
		"MacOS.release == :high_sierra":   {MacOSHighSierra, MacOSHighSierra},
		"MacOS.release == :sierra":        {MacOSSierra, MacOSSierra},
		"MacOS.release == :el_capitan":    {MacOSElCapitan, MacOSElCapitan},
		"MacOS.release == :yosemite":      {MacOSYosemite, MacOSYosemite},
		"MacOS.release == :mavericks":     {MacOSMavericks, MacOSMavericks},
		"MacOS.release == :mountain_lion": {MacOSMountainLion, MacOSMountainLion},
		"MacOS.release == :lion":          {MacOSLion, MacOSLion},
		"MacOS.release == :snow_leopard":  {MacOSSnowLeopard, MacOSSnowLeopard},
		"MacOS.release == :leopard":       {MacOSLeopard, MacOSLeopard},
		"MacOS.release == :tiger":         {MacOSTiger, MacOSTiger},

		// GT (>)
		"MacOS.release > :el_capitan":  {MacOSSierra, MacOSHighSierra},
		"MacOS.release > :high_sierra": {MacOSHighSierra, MacOSHighSierra},

		// LT (<)
		"MacOS.release < :el_capitan": {MacOSTiger, MacOSYosemite},
		"MacOS.release < :tiger":      {MacOSTiger, MacOSTiger},

		// GT and EQ (>=)
		"MacOS.release >= :el_capitan":  {MacOSElCapitan, MacOSHighSierra},
		"MacOS.release >= :high_sierra": {MacOSHighSierra, MacOSHighSierra},

		// LT and EQ (<=)
		"MacOS.release <= :el_capitan": {MacOSTiger, MacOSElCapitan},
		"MacOS.release <= :tiger":      {MacOSTiger, MacOSTiger},
	}

	for testCase, expected := range testCases {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		min, max, err := p.ParseConditionMacOS()
		assert.Nil(t, err)
		assert.Equal(t, expected[0], min)
		assert.Equal(t, expected[1], max)
	}

	// test (error)
	testCasesErrors := map[string]string{
		"MacOS.release == :invalid": "MacOS condition is unknown",
		"invalid":                   "MacOS condition not found",
	}

	for testCase, expected := range testCasesErrors {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		min, max, err := p.ParseConditionMacOS()
		assert.Equal(t, MacOSHighSierra, min)
		assert.Equal(t, MacOSHighSierra, max)
		assert.Error(t, err)
		assert.Equal(t, expected, err.Error())
	}
}

func TestMergeCurrentCaskVariantIfNotEmpty(t *testing.T) {
	// preparations
	p := createCaskTestParser()
	p.currentCaskVariant = NewVariant()

	// test (string)
	assert.Len(t, p.cask.Variants, 0)
	p.mergeCurrentCaskVariantIfNotEmpty("test")
	assert.Len(t, p.cask.Variants, 1)

	// test (strings array)
	p.mergeCurrentCaskVariantIfNotEmpty([]string{"test", "test"})
	assert.Len(t, p.cask.Variants, 2)

	// test (strings array)
	p.mergeCurrentCaskVariantIfNotEmpty([]Artifact{
		*NewArtifact(ArtifactApp, "test"),
		*NewArtifact(ArtifactApp, "test"),
	})
	assert.Len(t, p.cask.Variants, 3)
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

func TestCurrentTokenLiteralIs(t *testing.T) {
	// preparations
	p := createTokenTestParser()

	// test
	assert.True(t, p.currentTokenLiteralIs("five"))
}

func TestCurrentTokenLiteralOneOf(t *testing.T) {
	// preparations
	p := createTokenTestParser()

	// test
	assert.False(t, p.currentTokenLiteralOneOf("="))
	assert.True(t, p.currentTokenLiteralOneOf("five"))
	assert.True(t, p.currentTokenLiteralOneOf("five", "="))
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
