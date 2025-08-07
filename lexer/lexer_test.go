package lexer

import (
	"h2m/token"
	"testing"
)

func TestGetNextToken(t *testing.T) {
	input := "<div> content </div>"

	expectedTokens := [3]token.TokenType{token.OPEN_DIV, token.CLOSED_DIV,token.CLOSED_DIV}

	l := New(input)

for i, expectedType := range expectedTokens {
	tok := l.nextToken()

	if tok.Type != expectedType {
		t.Errorf("tests[%d] - token type wrong. expected=%q, got=%q", i, expectedType, tok.Type)
	}
}
}
