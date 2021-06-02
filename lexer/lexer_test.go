package lexer_test

import (
	"testing"

	"github.com/vegarsti/calc/lexer"
	"github.com/vegarsti/calc/token"
)

func TestExpressionValue(t *testing.T) {
	input := "1 + 2 * (30 / 5) - 1"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "1"},
		{token.PLUS, "+"},
		{token.INT, "2"},
		{token.MULTIPLY, "*"},
		{token.LPAREN, "("},
		{token.INT, "30"},
		{token.DIVIDE, "/"},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.MINUS, "-"},
		{token.INT, "1"},
	}
	l := lexer.New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
	}
}
