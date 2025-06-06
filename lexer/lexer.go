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
		return l.readPlusToken()
	case '-':
		return l.readHyphenToken()
	case '*':
		return l.readAsterickToken()
	case '/':
		return l.readSlashToken()
	case '%':
		return l.readPercentToken()
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
	return l.readTwoCharToken(token.ASSIGN, map[rune]token.Type{
		'=': token.EQ,
	})
}

func (l *Lexer) readBangToken() token.Token {
	return l.readTwoCharToken(token.BANG, map[rune]token.Type{
		'=': token.NOT_EQ,
	})
}

func (l *Lexer) readLessThanToken() token.Token {
	return l.readTwoCharToken(token.LT, map[rune]token.Type{
		'=': token.LTE,
		'<': token.LSHIFT,
	})
}

func (l *Lexer) readGreaterThanToken() token.Token {
	return l.readTwoCharToken(token.GT, map[rune]token.Type{
		'=': token.GTE,
		'>': token.RSHIFT,
	})
}

func (l *Lexer) readPlusToken() token.Token {
	return l.readTwoCharToken(token.PLUS, map[rune]token.Type{
		'+': token.INCREMENT,
		'=': token.PLUS_EQUAL,
	})
}

func (l *Lexer) readHyphenToken() token.Token {
	return l.readTwoCharToken(token.HYPHEN, map[rune]token.Type{
		'-': token.DECREMENT,
		'=': token.MINUS_EQUAL,
	})
}

func (l *Lexer) readAsterickToken() token.Token {
	return l.readTwoCharToken(token.ASTERISK, map[rune]token.Type{
		'=': token.MULTIPLY_EQUAL,
	})
}

func (l *Lexer) readSlashToken() token.Token {
	return l.readTwoCharToken(token.SLASH, map[rune]token.Type{
		'=': token.DIVIDE_EQUAL,
	})
}

func (l *Lexer) readPercentToken() token.Token {
	return l.readTwoCharToken(token.PERCENT, map[rune]token.Type{
		'=': token.MODULO_EQUAL,
	})
}

// readTwoCharToken takes a fallback `singleCharType`
// it peaks the next char and if its in the mapping then returns the associated token type
// and if it does not then falls back to singleCharType
func (l *Lexer) readTwoCharToken(singleCharType token.Type, mappings map[rune]token.Type) token.Token {
	peek := l.peakChar()
	tokenType, ok := mappings[peek]
	if !ok {
		tok := newToken(singleCharType, l.char)
		l.readChar()
		return tok
	}
	ch := l.char
	l.readChar()
	tok := token.Token{
		Type:    tokenType,
		Literal: string([]rune{ch, l.char}),
	}
	l.readChar()
	return tok
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}
