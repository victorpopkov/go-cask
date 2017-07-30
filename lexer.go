package cask

// The Lexer has been stolen from:
// https://github.com/goruby/goruby/blob/master/lexer

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

const eof = -1

// Lexer is the engine to process input and emit Tokens.
type Lexer struct {
	// input specifies the string being scanned.
	input string

	// state specifies the next lexing function to enter.
	state StateFn

	// position specifies current position in the input.
	position int

	// lines specifies the number of lines that have been lexed in input.
	lines int

	// start specifies the position of this item.
	start int

	// width specifies the width of the last rune read from input.
	width int

	// tokens specifies the channel of scanned tokens.
	tokens chan Token
}

// LexStartFn represents the entrypoint the Lexer uses to start processing the
// input.
var LexStartFn = startLexer

// StateFn represents a function which is capable of lexing parts of the
// input. It returns another StateFn to proceed with.
//
// Typically a state function would get called from LexStartFn and should
// return LexStartFn to go back to the decision loop. It also could return
// another non start state function if the partial input to parse is abiguous.
type StateFn func(*Lexer) StateFn

// NewLexer creates a new Lexer instance and returns its pointer. Requires an
// input to be passed as an argument that is ready to be processed.
func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		state:  startLexer,
		tokens: make(chan Token, 2), // two tokens is sufficient.
	}
}

// NextToken will return the next token processed from the lexer.
func (l *Lexer) NextToken() Token {
	for {
		select {
		case item, ok := <-l.tokens:
			if ok {
				return item
			}
			// panic(fmt.Errorf("No items left"))
		default:
			l.state = l.state(l)
			// if l.state == nil {
			// 	close(l.tokens)
			// }
		}
	}
}

// HasNext returns true if there are tokens left, false if EOF has reached.
func (l *Lexer) HasNext() bool {
	return l.state != nil
}

// emit passes a token back to the client.
func (l *Lexer) emit(t TokenType) {
	l.tokens <- *NewToken(t, l.input[l.start:l.position], l.start)
	l.start = l.position
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	if l.position >= len(l.input) {
		l.width = 0
		return eof
	}
	var r rune
	r, l.width = utf8.DecodeRuneInString(l.input[l.position:])
	l.position += l.width
	return r
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	l.start = l.position
}

// backup steps back one rune.
func (l *Lexer) backup() {
	l.position -= l.width
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.run.
func (l *Lexer) errorf(format string, args ...interface{}) StateFn {
	l.tokens <- *NewToken(ILLEGAL, fmt.Sprintf(format, args...), l.start)
	return nil
}

// startLexer starts the Lexer processing.
func startLexer(l *Lexer) StateFn {
	r := l.next()
	if isWhitespace(r) {
		l.ignore()
		return startLexer
	}
	switch r {
	case '$':
		return lexGlobal
	case '\n':
		l.lines++
		l.emit(NEWLINE)
		return startLexer
	case '\'':
		return lexSingleQuoteString
	case '"':
		return lexString
	case ':':
		if p := l.peek(); p == ':' {
			l.next()
			l.emit(SCOPE)
			return startLexer
		}
		return lexSymbol
	case '.':
		l.emit(DOT)
		return startLexer
	case '=':
		if l.peek() == '=' {
			l.next()
			l.emit(EQ)
		} else {
			l.emit(ASSIGN)
		}
		return startLexer
	case '+':
		l.emit(PLUS)
		return startLexer
	case '-':
		l.emit(MINUS)
		return startLexer
	case '!':
		if l.peek() == '=' {
			l.next()
			l.emit(NOTEQ)
		} else {
			l.emit(BANG)
		}
		return startLexer
	case '/':
		l.emit(SLASH)
		return startLexer
	case '*':
		l.emit(ASTERISK)
		return startLexer
	case '<':
		l.emit(LT)
		return startLexer
	case '>':
		l.emit(GT)
		return startLexer
	case '(':
		l.emit(LPAREN)
		return startLexer
	case ')':
		l.emit(RPAREN)
		return startLexer
	case '{':
		l.emit(LBRACE)
		return startLexer
	case '}':
		l.emit(RBRACE)
		return startLexer
	case '[':
		l.emit(LBRACKET)
		return startLexer
	case ']':
		l.emit(RBRACKET)
		return startLexer
	case ',':
		l.emit(COMMA)
		return startLexer
	case ';':
		l.emit(SEMICOLON)
		return startLexer
	case eof:
		l.emit(EOF)
		return startLexer
	case '#':
		return commentLexer
	case '|':
		l.emit(PIPE)
		return startLexer
	default:
		if isLetter(r) {
			return lexIdentifier
		}

		if isDigit(r) {
			return lexDigit
		}

		return l.errorf("Illegal character at %d: '%c'", l.start, r)
	}
}

// lexIdentifier lexes the identifier.
func lexIdentifier(l *Lexer) StateFn {
	legalIdentifierCharacters := []byte{'?', '!'}

	r := l.next()
	for isLetter(r) || isDigit(r) || bytes.ContainsRune(legalIdentifierCharacters, r) {
		r = l.next()
	}

	l.backup()
	literal := l.input[l.start:l.position]
	l.emit(LookupIdent(literal))

	return startLexer
}

// lexIdentifier lexes the digit.
func lexDigit(l *Lexer) StateFn {
	r := l.next()
	for isDigitOrUnderscore(r) {
		r = l.next()
	}

	l.backup()
	l.emit(INT)

	return startLexer
}

// lexSingleQuoteString lexes the single quote string.
func lexSingleQuoteString(l *Lexer) StateFn {
	l.ignore()
	r := l.next()

	for r != '\'' {
		r = l.next()
	}
	l.backup()
	l.emit(STRING)
	l.next()
	l.ignore()

	return startLexer
}

// lexString lexes the string.
func lexString(l *Lexer) StateFn {
	l.ignore()

	r := l.next()
	for r != '"' {
		r = l.next()
	}

	l.backup()
	l.emit(STRING)
	l.next()
	l.ignore()

	return startLexer
}

// lexSymbol lexes the symbol.
func lexSymbol(l *Lexer) StateFn {
	l.ignore()

	r := l.next()
	for isLetter(r) || isDigit(r) {
		r = l.next()
	}

	l.backup()
	l.emit(SYMBOL)

	return startLexer
}

// lexGlobal lexes the global.
func lexGlobal(l *Lexer) StateFn {
	r := l.next()

	if isExpressionDelimiter(r) {
		return l.errorf("Illegal character at %d: '%c'", l.position, r)
	}

	if isWhitespace(r) {
		return l.errorf("Illegal character at %d: '%c'", l.position, r)
	}

	for !isWhitespace(r) && !isExpressionDelimiter(r) {
		r = l.next()
	}

	l.backup()
	l.emit(GLOBAL)

	return startLexer
}

// commentLexer lexes the comment.
func commentLexer(l *Lexer) StateFn {
	r := l.next()
	for r != '\n' {
		r = l.next()
	}

	l.ignore()

	return startLexer
}

// isDigit checks whether the specified rune is a whitespace character.
func isWhitespace(r rune) bool {
	return unicode.IsSpace(r) && r != '\n'
}

// isDigit checks whether the specified rune is a letter.
func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '_'
}

// isDigit checks whether the specified rune is a digit.
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// isDigitOrUnderscore checks whether the specified rune is a digit or an
// underscore.
func isDigitOrUnderscore(r rune) bool {
	return isDigit(r) || r == '_'
}

// isExpressionDelimiter checks whether the specified rune is an expression
// delimiter.
func isExpressionDelimiter(r rune) bool {
	return r == '\n' || r == ';' || r == eof
}
