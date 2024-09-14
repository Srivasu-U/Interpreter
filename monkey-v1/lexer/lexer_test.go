package lexer

import (
	"Learning-Go/monkey-v1/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	var input string = `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLit  string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	stringUnderTest := New(input)

	for i, testToken := range tests {
		token := stringUnderTest.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("testing value [%d] - tokentype wrong. expected=%q, got=%q", i+1, testToken.expectedType, token.Type)
		}

		if token.Literal != testToken.expectedLit {
			t.Fatalf("testing value [%d] - literal wrong. expected=%q, got=%q", i+1, testToken.expectedLit, token.Literal)
		}
	}
}
