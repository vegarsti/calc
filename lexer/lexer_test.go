package lexer_test

import (
	"testing"

	"github.com/vegarsti/calc/lexer"
	"github.com/vegarsti/calc/token"
)

func TestExpressionValue(t *testing.T) {
	input := "+-*/"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.MULTIPLY, "*"},
		{token.DIVIDE, "/"},
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
