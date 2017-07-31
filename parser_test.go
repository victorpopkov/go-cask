package cask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTokenTestParser() *Parser {
	// preparations
	p := &Parser{
		lexer: NewLexer("five = 5"),
	}

	p.nextToken()
	p.nextToken()

	return p
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
		assert.Equal(t, expected, actual)
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
			URL:        "https://example.com/appcast.xml",
			Checkpoint: "",
		},
		"appcast 'https://example.com/appcast.xml'": {
			URL:        "https://example.com/appcast.xml",
			Checkpoint: "",
		},
		"appcast 'https://example.com/appcast.xml', checkpoint: '2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1'": {
			URL:        "https://example.com/appcast.xml",
			Checkpoint: "2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1",
		},
		`appcast 'https://example.com/appcast.xml',
		         checkpoint: '2ffedc4898df88e05a6e8f5519e11159d967153f75f8d4e8c9a0286d347ea1e1'
    `: {
			URL:        "https://example.com/appcast.xml",
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
		assert.Equal(t, expected.URL, actual.URL)
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
		actual, err := p.parseArtifact()
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
		actual, err := p.parseArtifact()
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
			Type:           ArtifactApp,
			Value:          "test.pkg",
			AllowUntrusted: false,
		},
		"pkg 'test.pkg'": {
			Type:           ArtifactApp,
			Value:          "test.pkg",
			AllowUntrusted: false,
		},
		"pkg 'test.pkg', allow_untrusted: true": {
			Type:           ArtifactApp,
			Value:          "test.pkg",
			AllowUntrusted: true,
		},
		"pkg 'test.pkg',\nallow_untrusted: true": {
			Type:           ArtifactApp,
			Value:          "test.pkg",
			AllowUntrusted: true,
		},
	}

	for testCase, expected := range testCases {
		// preparations
		l := NewLexer(testCase)
		p := NewParser(l)

		// test
		actual, err := p.parseArtifact()
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
		actual, err := p.parseArtifact()
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
		actual, err := p.parseArtifact()
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
		actual, err := p.parseArtifact()
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
		min, max, err := p.parseConditionMacOS()
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
		min, max, err := p.parseConditionMacOS()
		assert.Equal(t, MacOSHighSierra, min)
		assert.Equal(t, MacOSHighSierra, max)
		assert.Error(t, err)
		assert.Equal(t, expected, err.Error())
	}
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
