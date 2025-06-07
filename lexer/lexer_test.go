package lexer_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/ParasRaba155/monkey-lang/lexer"
	"github.com/ParasRaba155/monkey-lang/token"
)

func readfileInRunes(path string) ([]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("readfileInRunes: open file: %w", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	result := []rune{}

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 1 {
			return nil, fmt.Errorf("got len(text) != 0: %d with text %q", len(text), text)
		}
		r := []rune(text)
		result = append(result, r...)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading runes failed: %w", err)
	}
	return result, nil
}

func TestNextToken(t *testing.T) {
	input, err := readfileInRunes("./../testdata/lexer/01_simple.txt")
	if err != nil {
		t.Errorf("did not expect error in reading file: %v", err)
	}
	lexer := lexer.New(input)

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, "\x00"},
	}
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - wrong token type. Expected %q got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - wrong literal. Expected %q got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestInvalidUTF8File(t *testing.T) {
	input, err := readfileInRunes("./../testdata/lexer/02_invalid_utf8.txt")
	if err == nil {
		t.Errorf("expected error in reading file: %v", err)
	}
	if len(input) != 0 {
		t.Errorf("expected 0 input rune slice got %d", len(input))
	}
}

func TestLangaugeFile(t *testing.T) {
	input, err := readfileInRunes("./../testdata/lexer/03_lang.txt")
	if err != nil {
		t.Errorf("did not expect error in reading file: %v", err)
	}
	lexer := lexer.New(input)
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, "\x00"},
	}
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - wrong token type. Expected %q got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - wrong literal. Expected %q got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestMoreOperators(t *testing.T) {
	input, err := readfileInRunes("./../testdata/lexer/04_additional_operators.txt")
	if err != nil {
		t.Errorf("did not expect error in reading file: %v", err)
	}
	lexer := lexer.New(input)
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.HYPHEN, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, "\x00"},
	}
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - wrong token type. Expected %q got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - wrong literal. Expected %q got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestIfElseReturn(t *testing.T) {
	input, err := readfileInRunes("./../testdata/lexer/05_if_else_return.txt")
	if err != nil {
		t.Errorf("did not expect error in reading file: %v", err)
	}
	lexer := lexer.New(input)

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "some_var"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "some_var"},
		{token.LT, "<"},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "some_var"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - wrong token type. Expected %q got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - wrong literal. Expected %q got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestTwoCharOperators(t *testing.T) {
	input, err := readfileInRunes("./../testdata/lexer/06_two_char_operators.txt")
	if err != nil {
		t.Errorf("did not expect error in reading file: %v", err)
	}
	lexer := lexer.New(input)

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.GTE, ">="},
		{token.INT, "120"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.LTE, "<="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.RSHIFT, ">>"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.LSHIFT, "<<"},
		{token.INT, "102"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.INCREMENT, "++"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.PLUS_EQUAL, "+="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.MINUS_EQUAL, "-="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.MULTIPLY_EQUAL, "*="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.DIVIDE_EQUAL, "/="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.MODULO_EQUAL, "%="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.DECREMENT, "--"},
		{token.SEMICOLON, ";"},
		{token.INT, "2"},
		{token.AMPERSAND, "&"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.INT, "2"},
		{token.PIPE, "|"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.INT, "2"},
		{token.OR, "||"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.INT, "2"},
		{token.AND, "&&"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.INT, "3"},
		{token.BINARY_AND_EQUAL, "&="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.INT, "3"},
		{token.BINARY_OR_EQUAL, "|="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.EOF, "\x00"},
	}
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - wrong token type. Expected %q got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - wrong literal. Expected %q got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
