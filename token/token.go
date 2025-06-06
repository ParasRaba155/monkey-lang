package token

//go:generate stringer -type=Type
type Type int

const (
	ILLEGAL Type = iota
	EOF
	// Operators
	ASSIGN
	PLUS
	HYPHEN
	ASTERISK
	SLASH
	BANG
	LT
	GT
	PERCENT
	// two char operators
	EQ
	NOT_EQ
	LTE
	GTE
	LSHIFT
	RSHIFT
	// two char assignment operator
	PLUS_EQUAL
	MINUS_EQUAL
	MULTIPLY_EQUAL
	DIVIDE_EQUAL
	MODULO_EQUAL
	// incr & decr
	INCREMENT
	DECREMENT
	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	// Language specific keywords
	LET
	IDENT
	INT
	FUNCTION
	IF
	ELSE
	RETURN
)

type Token struct {
	Type    Type
	Literal string
}

func LookupIdent(ident string) Type {
	switch ident {
	case "fn":
		return FUNCTION
	case "let":
		return LET
	case "if":
		return IF
	case "else":
		return ELSE
	case "return":
		return RETURN
	default:
		return IDENT
	}
}
