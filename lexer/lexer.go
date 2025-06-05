package lexer

import (
	"unicode"

	"github.com/ParasRaba155/monkey-lang/token"
)

type Lexer struct {
	input   []rune
	currPos int
	readPos int
	char    rune
}

func New(input []rune) *Lexer {
	l := Lexer{
		input:   input,
		currPos: 0,
		readPos: 0,
		char:    0,
	}
	l.readChar()
	return &l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}
	l.currPos = l.readPos
	l.readPos++
}

func (l *Lexer) peakChar() rune {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

func (l *Lexer) atEOF() bool {
	return l.currPos == len(l.input)
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhiteSpace()
	var tok token.Token

	switch l.char {
	// All the single chars will need to read the next char to advance the `l.char`
	// however the complex tokens will have always advance the `l.char`

	// all the delimiters
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
		l.readChar()
	case ',':
		tok = newToken(token.COMMA, l.char)
		l.readChar()
	case '(':
		tok = newToken(token.LPAREN, l.char)
		l.readChar()
	case ')':
		tok = newToken(token.RPAREN, l.char)
		l.readChar()
	case '{':
		tok = newToken(token.LBRACE, l.char)
		l.readChar()
	case '}':
		tok = newToken(token.RBRACE, l.char)
		l.readChar()
	// all the operators
	case '=':
		return l.readEqualToken()
	case '+':
		tok = newToken(token.PLUS, l.char)
		l.readChar()
	case '-':
		tok = newToken(token.HYPHEN, l.char)
		l.readChar()
	case '*':
		tok = newToken(token.ASTERISK, l.char)
		l.readChar()
	case '/':
		tok = newToken(token.SLASH, l.char)
		l.readChar()
	case '!':
		return l.readBangToken()
	case '<':
		return l.readLessThanToken()
	case '>':
		return l.readGreaterThanToken()
	default:
		return l.readComplexToken()
	}
	return tok
}

func newToken(tokenType token.Type, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	pos := l.currPos
	for isLetter(l.char) {
		l.readChar()
	}
	return string(l.input[pos:l.currPos])
}

func (l *Lexer) readNumber() string {
	pos := l.currPos
	for unicode.IsDigit(l.char) {
		l.readChar()
	}
	return string(l.input[pos:l.currPos])
}

func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(l.char) {
		l.readChar()
	}
}

func (l *Lexer) readComplexToken() token.Token {
	if isLetter(l.char) {
		ident := l.readIdentifier()
		return token.Token{
			Type:    token.LookupIdent(ident),
			Literal: ident,
		}
	}

	if unicode.IsDigit(l.char) {
		num := l.readNumber()
		return token.Token{
			Type:    token.INT,
			Literal: num,
		}
	}

	if l.atEOF() {
		return newToken(token.EOF, 0)
	}

	tok := newToken(token.ILLEGAL, l.char)
	return tok
}

func (l *Lexer) readEqualToken() token.Token {
	if l.peakChar() != '=' {
		tok := newToken(token.ASSIGN, l.char)
		l.readChar()
		return tok
	}
	ch := l.char
	l.readChar()
	tok := token.Token{
		Type:    token.EQ,
		Literal: string([]rune{ch, l.char}),
	}
	l.readChar()
	return tok
}

func (l *Lexer) readBangToken() token.Token {
	if l.peakChar() != '=' {
		tok := newToken(token.BANG, l.char)
		l.readChar()
		return tok
	}
	ch := l.char
	l.readChar()
	tok := token.Token{
		Type:    token.NOT_EQ,
		Literal: string([]rune{ch, l.char}),
	}
	l.readChar()
	return tok
}

func (l *Lexer) readLessThanToken() token.Token {
	nextChar := l.peakChar()
	if nextChar != '=' && nextChar != '<' {
		tok := newToken(token.LT, l.char)
		l.readChar()
		return tok
	}

	if nextChar == '=' {
		ch := l.char
		l.readChar()
		tok := token.Token{
			Type:    token.LTE,
			Literal: string([]rune{ch, l.char}),
		}
		l.readChar()
		return tok
	}

	ch := l.char
	l.readChar()
	tok := token.Token{
		Type:    token.LSHIFT,
		Literal: string([]rune{ch, l.char}),
	}
	l.readChar()
	return tok
}

func (l *Lexer) readGreaterThanToken() token.Token {
	nextChar := l.peakChar()
	if nextChar != '=' && nextChar != '>' {
		tok := newToken(token.GT, l.char)
		l.readChar()
		return tok
	}
	if nextChar == '=' {
		ch := l.char
		l.readChar()
		tok := token.Token{
			Type:    token.GTE,
			Literal: string([]rune{ch, l.char}),
		}
		l.readChar()
		return tok
	}

	ch := l.char
	l.readChar()
	tok := token.Token{
		Type:    token.RSHIFT,
		Literal: string([]rune{ch, l.char}),
	}
	l.readChar()
	return tok
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}
